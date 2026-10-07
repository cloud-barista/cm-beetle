package recommendation

import (
	"fmt"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/rs/zerolog/log"
)

// HasAnyGpu returns true if at least one node in the infrastructure contains physical GPU accelerator cards.
func HasAnyGpu(srcInfra onpremmodel.OnpremInfra) bool {
	for _, node := range srcInfra.Nodes {
		if hasGpu(node) {
			return true
		}
	}
	return false
}

// DecomposeOnpremNodes decomposes heterogeneous on-premise physical servers into homogeneous logical nodes.
// For nodes without GPUs or with homogeneous cards (only 1 cluster), the node is preserved 1:1.
// For nodes with mixed/heterogeneous GPU cards (e.g. 2x A100 + 1x T4 + 1x MI350), the node is decomposed
// into K distinct logical nodes (1:K), each containing one homogeneous GPU cluster.
func DecomposeOnpremNodes(nodes []onpremmodel.NodeProperty) []onpremmodel.NodeProperty {
	decomposed := make([]onpremmodel.NodeProperty, 0, len(nodes))

	for _, node := range nodes {
		if !hasGpu(node) {
			decomposed = append(decomposed, node)
			continue
		}

		clusters := clusterGpuCards(node.GPUCards)
		if len(clusters) <= 1 {
			// Already homogeneous (0 or 1 GPU cluster)
			decomposed = append(decomposed, node)
			continue
		}

		// Heterogeneous: decompose into K logical nodes
		for clusterIdx, cluster := range clusters {
			// Find all physical cards matching this cluster
			matchedCards := make([]onpremmodel.GpuCardProperty, 0, cluster.Count)
			for _, card := range node.GPUCards {
				if strings.EqualFold(card.Vendor, cluster.Vendor) &&
					strings.EqualFold(card.Model, cluster.Model) &&
					card.MemoryTotalGB == cluster.MemoryTotalGB {
					matchedCards = append(matchedCards, card)
				}
			}

			// Format unique machine ID while preserving traceability
			sanitizedModel := sanitizeName(cluster.Model)
			if sanitizedModel == "" {
				sanitizedModel = fmt.Sprintf("cluster-%d", clusterIdx+1)
			}
			logicalMachineId := fmt.Sprintf("%s-gpu-%s", node.MachineId, sanitizedModel)
			if clusterIdx == 0 {
				logicalMachineId = fmt.Sprintf("%s-primary-%s", node.MachineId, sanitizedModel)
			}

			logicalNode := node
			logicalNode.MachineId = logicalMachineId
			logicalNode.GPUCards = matchedCards

			decomposed = append(decomposed, logicalNode)
		}
	}

	return decomposed
}

// RecommendGpuInfraCandidates recommends multiple target infrastructure candidates
// for source infrastructure containing GPU accelerators.
// It adheres to the Go Single Responsibility Principle by encapsulating heterogeneous GPU decomposition,
// shared intra-node networking, and GPU-aware infrastructure assembly.
func RecommendGpuInfraCandidates(
	desiredCsp string,
	desiredRegion string,
	srcInfra onpremmodel.OnpremInfra,
	limit int,
	minMatchRate float64,
) ([]cloudmodel.RecommendedInfra, error) {
	srcInfra = SanitizeSourceInfra(srcInfra)

	var recommendedVmInfraCandidates []cloudmodel.RecommendedInfra

	var limitSpecs int = GetDefaultSpecsLimit()
	var limitImages int = GetDefaultImagesLimit()
	if limit <= 0 {
		limit = 3
	}

	csp := strings.ToLower(desiredCsp)
	region := strings.ToLower(desiredRegion)
	connectionName := fmt.Sprintf("%s-%s", csp, region)

	if !isSupportedCSP(csp) {
		log.Warn().Msgf("unsupported CSP for GPU infrastructure recommendation: %s", csp)
	}

	// 1. Recommend vNet and subnets
	recommendedVNetInfoList, err := RecommendVNet(csp, region, srcInfra)
	if err != nil {
		log.Warn().Err(err).Msg("failed to recommend a virtual network for GPU infrastructure")
	}

	skeletonVmInfra := cloudmodel.RecommendedInfra{
		Description: "Recommended target GPU-accelerated multi-cloud infrastructure.",
		Status:      "",
		TargetCloud: cloudmodel.CloudProperty{
			Csp:    csp,
			Region: region,
		},
		TargetInfra: cloudmodel.InfraReq{
			Name:       "gpu-infra-01",
			NodeGroups: []cloudmodel.CreateNodeGroupReq{},
		},
	}

	if len(recommendedVNetInfoList) > 0 {
		skeletonVmInfra.TargetVNet = recommendedVNetInfoList[0]
		skeletonVmInfra.TargetVNet.Name = "vnet-01"
		skeletonVmInfra.TargetVNet.Description = "a recommended vNet for GPU workload migration"
		for i := range skeletonVmInfra.TargetVNet.SubnetInfoList {
			skeletonVmInfra.TargetVNet.SubnetInfoList[i].Name = fmt.Sprintf("subnet-%02d", i+1)
			skeletonVmInfra.TargetVNet.SubnetInfoList[i].Description = "a recommended subnet for GPU migration"
		}
	}

	// 2. Recommend SSH key pair
	skeletonVmInfra.TargetSshKey.Name = "sshkey-01"
	skeletonVmInfra.TargetSshKey.ConnectionName = connectionName
	skeletonVmInfra.TargetSshKey.Description = "SSH Key pair for GPU migration"

	// 3. Decompose heterogeneous GPU servers into homogeneous logical nodes
	expandedNodes := DecomposeOnpremNodes(srcInfra.Nodes)
	log.Info().
		Int("sourceNodes", len(srcInfra.Nodes)).
		Int("expandedNodes", len(expandedNodes)).
		Msg("Decomposed on-premise nodes for GPU infrastructure recommendation")

	// 4. Generate NodeGroups and Security Groups for each expanded node
	var skeletonNodegroupList = make([]cloudmodel.CreateNodeGroupReq, 0, len(expandedNodes))
	var deduplicatedSecurityGroupList = []cloudmodel.SecurityGroupReq{}

	// Enforce safe minimum root disk for GPU nodes to hold model weights/checkpoints
	const gpuMinRootDiskSizeGB = 100

	firstSubnetId := ""
	if len(skeletonVmInfra.TargetVNet.SubnetInfoList) > 0 {
		firstSubnetId = skeletonVmInfra.TargetVNet.SubnetInfoList[0].Name
	}

	for i, node := range expandedNodes {
		minDisk := getCspMinRootDiskSizeGB(csp)
		if hasGpu(node) && minDisk < gpuMinRootDiskSizeGB {
			minDisk = gpuMinRootDiskSizeGB
		}
		rootDiskSize := max(int(node.RootDisk.TotalSize), minDisk)

		// Recommend security group for the node
		recommendedSg, sgErr := RecommendSecurityGroup(csp, region, node)
		if sgErr != nil {
			log.Warn().Err(sgErr).Msgf("failed to recommend security group for GPU node %s", node.MachineId)
		}

		exists, _, existingSg := containSg(deduplicatedSecurityGroupList, recommendedSg)
		if !exists {
			recommendedSg.Name = fmt.Sprintf("sg-%02d", len(deduplicatedSecurityGroupList)+1)
			recommendedSg.ConnectionName = connectionName
			recommendedSg.Description = fmt.Sprintf("Recommended security group for %s", node.MachineId)
			recommendedSg.VNetId = skeletonVmInfra.TargetVNet.Name
			deduplicatedSecurityGroupList = append(deduplicatedSecurityGroupList, recommendedSg)
		} else {
			recommendedSg = existingSg
		}

		tempCreateNodeGroupReq := cloudmodel.CreateNodeGroupReq{
			ConnectionName:   connectionName,
			Description:      fmt.Sprintf("Recommended VM %02d for %s", i+1, node.MachineId),
			VNetId:           skeletonVmInfra.TargetVNet.Name,
			SubnetId:         firstSubnetId,
			Name:             fmt.Sprintf("vm-%s", sanitizeName(node.MachineId)),
			RootDiskType:     "",
			RootDiskSize:     rootDiskSize,
			SshKeyId:         skeletonVmInfra.TargetSshKey.Name,
			NodeGroupSize:    1,
			SecurityGroupIds: []string{recommendedSg.Name},
			Label: map[string]string{
				lineage.LabelSourceMachineIds: node.MachineId,
			},
		}

		skeletonNodegroupList = append(skeletonNodegroupList, tempCreateNodeGroupReq)
	}

	skeletonVmInfra.TargetSecurityGroupList = deduplicatedSecurityGroupList

	// 5. Recommend compatible VM spec and OS image pairs per expanded node
	compatiblePairsForEachNode := make([][]CompatibleSpecImagePair, len(expandedNodes))

	for i, node := range expandedNodes {
		specList, _, specErr := RecommendNodeSpecs(csp, region, node, limitSpecs)
		if specErr != nil {
			log.Warn().Err(specErr).Msgf("failed to recommend specs for GPU node %s", node.MachineId)
		}

		imageList, imgErr := RecommendVmOsImages(csp, region, node, limitImages)
		if imgErr != nil {
			log.Warn().Err(imgErr).Msgf("failed to recommend images for GPU node %s", node.MachineId)
		}

		if len(specList) == 0 || len(imageList) == 0 {
			log.Warn().Msgf("no recommended VM specs or images found for GPU node %s", node.MachineId)
			continue
		}

		pairs, pairErr := FindCompatibleVmSpecAndImagePairs(specList, imageList, csp)
		if pairErr != nil {
			log.Warn().Err(pairErr).Msgf("failed to find compatible pairs for node %s; using fallback", node.MachineId)
			if len(specList) > 0 && len(imageList) > 0 {
				pairs = []CompatibleSpecImagePair{{Spec: specList[0], Image: imageList[0]}}
			}
		}
		compatiblePairsForEachNode[i] = pairs
	}

	// 6. Determine candidate count
	maxCandidates := 0
	for _, pairs := range compatiblePairsForEachNode {
		if len(pairs) > maxCandidates {
			maxCandidates = len(pairs)
		}
	}
	if maxCandidates > limit {
		maxCandidates = limit
	}

	if maxCandidates == 0 {
		log.Warn().Msg("no valid compatible pairs across any nodes; returning skeleton GPU infrastructure")
		skeletonVmInfra.Status = "nothing-to-recommend"
		skeletonVmInfra.Description = "No compatible VM spec and OS image pairs found for GPU infrastructure."
		return []cloudmodel.RecommendedInfra{skeletonVmInfra}, nil
	}

	// 7. Assemble candidates
	for candidateIdx := 0; candidateIdx < maxCandidates; candidateIdx++ {
		candidateNodeGroups := make([]cloudmodel.CreateNodeGroupReq, len(skeletonNodegroupList))
		copy(candidateNodeGroups, skeletonNodegroupList)

		var candidateSpecList []cloudmodel.SpecInfo
		var candidateImageList []cloudmodel.ImageInfo

		for nodeIdx := range expandedNodes {
			pairs := compatiblePairsForEachNode[nodeIdx]
			if len(pairs) == 0 {
				continue
			}

			pairIdx := candidateIdx
			if pairIdx >= len(pairs) {
				pairIdx = 0 // Wrap to best available pair
			}

			selectedPair := pairs[pairIdx]
			candidateNodeGroups[nodeIdx].SpecId = selectedPair.Spec.Id
			candidateNodeGroups[nodeIdx].ImageId = selectedPair.Image.Id
			candidateNodeGroups[nodeIdx].CspImageName = selectedPair.Image.CspImageName

			// Deduplicate spec list in candidate
			specFound := false
			for _, s := range candidateSpecList {
				if s.CspSpecName == selectedPair.Spec.CspSpecName {
					specFound = true
					break
				}
			}
			if !specFound {
				candidateSpecList = append(candidateSpecList, selectedPair.Spec)
			}

			// Deduplicate image list in candidate
			imageFound := false
			for _, img := range candidateImageList {
				if img.CspImageName == selectedPair.Image.CspImageName {
					imageFound = true
					break
				}
			}
			if !imageFound {
				candidateImageList = append(candidateImageList, selectedPair.Image)
			}
		}

		syntheticInfra := onpremmodel.OnpremInfra{Nodes: expandedNodes}
		overallStatus, overallStatusDesc, summary := calculateCandidateMatchRateWithDetails(
			csp, candidateNodeGroups, syntheticInfra, candidateSpecList, candidateImageList, minMatchRate,
		)

		candidate := skeletonVmInfra
		candidate.TargetInfra.NodeGroups = candidateNodeGroups
		candidate.TargetSpecList = candidateSpecList
		candidate.TargetOsImageList = candidateImageList
		candidate.Status = overallStatus
		// Skip duplicate candidate if all node groups match an existing candidate
		if isDuplicateInfraCandidate(recommendedVmInfraCandidates, candidate) {
			log.Debug().Int("candidateIndex", candidateIdx).Msg("skipping duplicate candidate infrastructure")
			continue
		}

		candidate.Description = fmt.Sprintf(
			"GPU Candidate #%d | %s | Overall Match Rate: Min=%.1f%% Max=%.1f%% Avg=%.1f%% | %s",
			len(recommendedVmInfraCandidates)+1,
			overallStatus,
			summary.MinMatchRate,
			summary.MaxMatchRate,
			summary.AvgMatchRate,
			overallStatusDesc,
		)

		recommendedVmInfraCandidates = append(recommendedVmInfraCandidates, candidate)
	}

	log.Info().
		Int("candidateCount", len(recommendedVmInfraCandidates)).
		Msg("Successfully recommended GPU infrastructure candidates")

	return recommendedVmInfraCandidates, nil
}
