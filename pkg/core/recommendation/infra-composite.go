// Package recommendation provides infrastructure and resource recommendation algorithms for CM-Beetle.
package recommendation

import (
	"fmt"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/rs/zerolog/log"
)

// Sizing policy constants for composite infrastructure recommendation.
const (
	SizingPolicyRightSizeUp = "rightSizeUp"
	SizingPolicyDividing    = "dividing"
)

// HardwareCase represents the hardware categorization of an NLB backend pool.
type HardwareCase int

const (
	// CasePureCpu indicates all backend servers are CPU-only without accelerators.
	CasePureCpu HardwareCase = 0
	// CaseMixedCpuGpu indicates both CPU-only servers and GPU servers are present.
	CaseMixedCpuGpu HardwareCase = 1
	// CaseCrossVendorGpu indicates GPUs from multiple vendors are present.
	CaseCrossVendorGpu HardwareCase = 2
	// CaseHeteroGpuModel indicates same-vendor GPUs with differing models or VRAM.
	CaseHeteroGpuModel HardwareCase = 3
	// CaseHomoGpuModel indicates all servers share identical GPU model and VRAM.
	CaseHomoGpuModel HardwareCase = 4
)

// String returns a human-readable name for the HardwareCase.
func (c HardwareCase) String() string {
	switch c {
	case CasePureCpu:
		return "Case 0 (Pure CPU)"
	case CaseMixedCpuGpu:
		return "Case 1 (Mixed CPU+GPU)"
	case CaseCrossVendorGpu:
		return "Case 2 (Cross-Vendor GPU)"
	case CaseHeteroGpuModel:
		return "Case 3 (Heterogeneous GPU Models)"
	case CaseHomoGpuModel:
		return "Case 4 (Homogeneous GPU Models)"
	default:
		return fmt.Sprintf("Unknown Case (%d)", int(c))
	}
}

// ClassifyBackendHardware evaluates the hardware accelerator profile of backend servers.
func ClassifyBackendHardware(members []onpremmodel.NodeProperty) HardwareCase {
	gpuCount := 0
	cpuCount := 0
	vendors := make(map[string]bool)
	modelKeys := make(map[string]bool)

	for _, m := range members {
		if hasGpu(m) {
			gpuCount++
			for _, card := range m.GPUCards {
				v := strings.ToLower(strings.TrimSpace(card.Vendor))
				if v == "" {
					v = "unknown"
				}
				vendors[v] = true
				key := fmt.Sprintf("%s:%s:%.1f", v, strings.ToLower(strings.TrimSpace(card.Model)), card.MemoryTotalGB)
				modelKeys[key] = true
			}
		} else {
			cpuCount++
		}
	}

	if gpuCount == 0 {
		return CasePureCpu
	}
	if cpuCount > 0 && gpuCount > 0 {
		return CaseMixedCpuGpu
	}
	if len(vendors) > 1 {
		return CaseCrossVendorGpu
	}
	if len(modelKeys) > 1 {
		return CaseHeteroGpuModel
	}
	return CaseHomoGpuModel
}

// RecommendCompositeInfraCandidates recommends high-level multi-cloud infrastructure candidates.
// It handles all 5 hardware cases and strict 1:1 Managed NLB to NodeGroup invariants.
func RecommendCompositeInfraCandidates(
	desiredCsp, desiredRegion string,
	srcInfra onpremmodel.OnpremInfra,
	sizingPolicy string,
	limit int,
	minMatchRate float64,
) ([]cloudmodel.RecommendedInfra, error) {
	if len(srcInfra.Nodes) == 0 {
		return nil, fmt.Errorf("sourceInfra.nodes is empty")
	}

	csp := strings.ToLower(desiredCsp)
	region := strings.ToLower(desiredRegion)

	if sizingPolicy == "" {
		sizingPolicy = SizingPolicyRightSizeUp
	}
	if sizingPolicy != SizingPolicyRightSizeUp && sizingPolicy != SizingPolicyDividing {
		return nil, fmt.Errorf("unsupported sizingPolicy '%s'; must be '%s' or '%s'", sizingPolicy, SizingPolicyRightSizeUp, SizingPolicyDividing)
	}

	limitSpecs := GetDefaultSpecsLimit()
	limitImages := GetDefaultImagesLimit()

	nodeByIP := buildNodeByIPIndex(srcInfra.Nodes)
	nodeByMachineId := buildNodeByMachineIdIndex(srcInfra.Nodes)

	// Phase 1: resolve NLBs and check member machine mapping
	var resolvedNlbs []resolvedNlb
	var warnings []string
	nlbMemberMachineIds := map[string]string{}

	for _, nlb := range srcInfra.NLBs {
		backendName := nlb.Backend.Name
		var matchedMachineIds []string

		for _, srv := range nlb.Backend.Servers {
			node, ok := nodeByIP[srv.IP]
			if !ok {
				warnings = append(warnings, fmt.Sprintf("NLB '%s': server IP '%s' not found", backendName, srv.IP))
				continue
			}
			matchedMachineIds = append(matchedMachineIds, node.MachineId)
		}

		if len(matchedMachineIds) == 0 {
			warnings = append(warnings, fmt.Sprintf("NLB '%s': no matching nodes found", backendName))
			continue
		}

		backendPort, portWarning := resolvePort(nlb.Backend.Servers)
		if portWarning != "" {
			warnings = append(warnings, fmt.Sprintf("NLB '%s': %s", backendName, portWarning))
		}

		for _, machineId := range matchedMachineIds {
			if _, exists := nlbMemberMachineIds[machineId]; !exists {
				nlbMemberMachineIds[machineId] = backendName
			}
		}

		resolvedNlbs = append(resolvedNlbs, resolvedNlb{
			sourceNlb:        nlb,
			backendPort:      backendPort,
			memberMachineIds: matchedMachineIds,
		})
	}

	// Phase 2: build shared network and credential skeleton
	connectionName := fmt.Sprintf("%s-%s", csp, region)
	skeleton := cloudmodel.RecommendedInfra{
		TargetCloud: cloudmodel.CloudProperty{Csp: csp, Region: region},
		TargetInfra: cloudmodel.InfraReq{
			Name:        "infra101",
			Description: fmt.Sprintf("Composite recommended infrastructure (%s policy)", sizingPolicy),
			NodeGroups:  []cloudmodel.CreateNodeGroupReq{},
		},
	}

	recommendedVNetList, err := RecommendVNet(csp, region, srcInfra)
	if err != nil || len(recommendedVNetList) == 0 {
		log.Warn().Err(err).Msg("failed to recommend vNet for composite infra")
	}
	if len(recommendedVNetList) > 0 {
		skeleton.TargetVNet = recommendedVNetList[0]
	}
	skeleton.TargetVNet.Name = "mig-vnet-01"
	skeleton.TargetVNet.Description = "a recommended vNet for composite migration"
	for i := range skeleton.TargetVNet.SubnetInfoList {
		skeleton.TargetVNet.SubnetInfoList[i].Name = fmt.Sprintf("mig-subnet-%02d", i+1)
		skeleton.TargetVNet.SubnetInfoList[i].Description = "a recommended subnet for composite migration"
	}

	firstSubnetId := ""
	if len(skeleton.TargetVNet.SubnetInfoList) > 0 {
		firstSubnetId = skeleton.TargetVNet.SubnetInfoList[0].Name
	}

	skeleton.TargetSshKey = cloudmodel.SshKeyReq{
		Name:           "mig-sshkey-01",
		ConnectionName: connectionName,
		Description:    "SSH key pair for composite migration (Note: provided ONLY once, MUST be downloaded)",
	}

	// Phase 3: generate NodeGroup and NLB blueprints per backend case
	var ngBlueprints []nodeGroupBlueprint
	var targetNlbList []cloudmodel.NlbReq
	var deduplicatedSgList []cloudmodel.SecurityGroupReq

	for _, rnlb := range resolvedNlbs {
		var memberNodes []onpremmodel.NodeProperty
		for _, mid := range rnlb.memberMachineIds {
			if n, ok := nodeByMachineId[mid]; ok {
				memberNodes = append(memberNodes, n)
			}
		}

		hwCase := ClassifyBackendHardware(memberNodes)
		log.Info().Str("backend", rnlb.sourceNlb.Backend.Name).Str("case", hwCase.String()).Str("policy", sizingPolicy).Msg("Classified backend hardware")

		bps, nlbs, sgs, genErr := buildCaseBlueprints(csp, region, connectionName, firstSubnetId, skeleton, rnlb, memberNodes, nodeByMachineId, hwCase, sizingPolicy)
		if genErr != nil {
			return nil, genErr
		}

		ngBlueprints = append(ngBlueprints, bps...)
		targetNlbList = append(targetNlbList, nlbs...)
		for _, sg := range sgs {
			exists, _, _ := containSg(deduplicatedSgList, sg)
			if !exists {
				deduplicatedSgList = append(deduplicatedSgList, sg)
			}
		}
	}

	// Phase 4: process standalone nodes with 1 Server -> 1 NodeGroup rule
	var unrelatedNodes []onpremmodel.NodeProperty
	for _, node := range srcInfra.Nodes {
		if _, isMember := nlbMemberMachineIds[node.MachineId]; !isMember {
			unrelatedNodes = append(unrelatedNodes, node)
		}
	}
	expandedUnrelatedNodes := DecomposeOnpremNodes(unrelatedNodes)

	for i, node := range expandedUnrelatedNodes {
		ngName := "ng-" + sanitizeName(node.MachineId)

		sg, sgErr := RecommendSecurityGroup(csp, region, node)
		if sgErr != nil {
			log.Warn().Err(sgErr).Str("machineId", node.MachineId).Msg("failed to recommend SG")
		}
		exists, _, existingSg := containSg(deduplicatedSgList, sg)
		if !exists {
			sg.Name = fmt.Sprintf("mig-sg-%02d", len(deduplicatedSgList)+1)
			sg.ConnectionName = connectionName
			sg.Description = fmt.Sprintf("Recommended security group for %s", node.MachineId)
			sg.VNetId = skeleton.TargetVNet.Name
			deduplicatedSgList = append(deduplicatedSgList, sg)
		} else {
			sg = existingSg
		}

		minDisk := getCspMinRootDiskSizeGB(csp)
		if hasGpu(node) && minDisk < 100 {
			minDisk = 100
		}
		rootDiskSize := max(int(node.RootDisk.TotalSize), minDisk)

		// Create dedicated NodeGroup with NodeGroupSize = 1 to prevent cost inflation.
		ngBlueprints = append(ngBlueprints, nodeGroupBlueprint{
			representativeNode: node,
			isNlbRelated:       false,
			skeleton: cloudmodel.CreateNodeGroupReq{
				ConnectionName:   connectionName,
				Name:             ngName,
				VNetId:           skeleton.TargetVNet.Name,
				SubnetId:         firstSubnetId,
				SecurityGroupIds: []string{sg.Name},
				SshKeyId:         skeleton.TargetSshKey.Name,
				RootDiskType:     "",
				RootDiskSize:     rootDiskSize,
				NodeGroupSize:    1,
				Description:      fmt.Sprintf("Standalone VM %02d for %s", i+1, node.MachineId),
				Label:            map[string]string{"sourceMachineIds": node.MachineId},
			},
		})
	}

	skeleton.TargetSecurityGroupList = deduplicatedSgList

	// Phase 5: query catalog for compatible spec-image pairs
	pairsByGroup := make([][]CompatibleSpecImagePair, len(ngBlueprints))
	for ngIdx, bp := range ngBlueprints {
		specList, _, specErr := RecommendNodeSpecs(csp, region, bp.representativeNode, limitSpecs)
		if specErr != nil {
			log.Warn().Err(specErr).Str("machineId", bp.representativeNode.MachineId).Msg("failed to recommend specs")
		}
		imageList, imgErr := RecommendVmOsImages(csp, region, bp.representativeNode, limitImages)
		if imgErr != nil {
			log.Warn().Err(imgErr).Str("machineId", bp.representativeNode.MachineId).Msg("failed to recommend images")
		}

		if len(specList) == 0 || len(imageList) == 0 {
			log.Warn().Str("machineId", bp.representativeNode.MachineId).Msg("no compatible spec or image found for NodeGroup")
			continue
		}

		pairs, pairErr := FindCompatibleVmSpecAndImagePairs(specList, imageList, csp)
		if pairErr != nil {
			log.Warn().Err(pairErr).Str("machineId", bp.representativeNode.MachineId).Msg("failed to find compatible pairs; using first available")
			if len(specList) > 0 && len(imageList) > 0 {
				pairsByGroup[ngIdx] = []CompatibleSpecImagePair{{Spec: specList[0], Image: imageList[0]}}
			}
		} else {
			pairsByGroup[ngIdx] = pairs
		}
	}

	// Phase 6: assemble Pareto-ranked composite candidates
	targetNlbList = sanitizeNlbListByCsp(targetNlbList, csp)
	maxCandidates := 0
	for _, pairs := range pairsByGroup {
		if len(pairs) > maxCandidates {
			maxCandidates = len(pairs)
		}
	}
	if maxCandidates > limit {
		maxCandidates = limit
	}

	var candidates []cloudmodel.RecommendedInfra
	for cIdx := 0; cIdx < maxCandidates; cIdx++ {
		candidateNodeGroups := make([]cloudmodel.CreateNodeGroupReq, len(ngBlueprints))
		for ngIdx := range ngBlueprints {
			candidateNodeGroups[ngIdx] = ngBlueprints[ngIdx].skeleton
		}

		var candidateSpecList []cloudmodel.SpecInfo
		var candidateImageList []cloudmodel.ImageInfo

		for ngIdx := range ngBlueprints {
			pairs := pairsByGroup[ngIdx]
			if len(pairs) == 0 || cIdx >= len(pairs) {
				continue
			}

			selectedPair := pairs[cIdx]
			selectedSpec := selectedPair.Spec
			selectedImage := selectedPair.Image

			precheck, prefErr := PreflightCheckCspProvisioning(
				selectedSpec.Id, selectedImage.Id, selectedImage.CspImageName, "",
			)
			if prefErr == nil {
				if precheck.SuggestedSystemDisk != "" {
					candidateNodeGroups[ngIdx].RootDiskType = precheck.SuggestedSystemDisk
				}
				if precheck.ResolvedCspImageName != "" && precheck.ResolvedCspImageName != selectedImage.CspImageName {
					candidateNodeGroups[ngIdx].CspImageName = precheck.ResolvedCspImageName
				}
			}

			candidateNodeGroups[ngIdx].SpecId = selectedSpec.Id
			candidateNodeGroups[ngIdx].ImageId = selectedImage.Id

			specAlreadyAdded := false
			for _, sp := range candidateSpecList {
				if sp.Id == selectedSpec.Id {
					specAlreadyAdded = true
					break
				}
			}
			if !specAlreadyAdded {
				candidateSpecList = append(candidateSpecList, selectedSpec)
			}

			imageAlreadyAdded := false
			for _, img := range candidateImageList {
				if img.CspImageName == selectedImage.CspImageName {
					imageAlreadyAdded = true
					break
				}
			}
			if !imageAlreadyAdded {
				candidateImageList = append(candidateImageList, selectedImage)
			}
		}

		syntheticSrcInfra := buildSyntheticSrcInfra(ngBlueprints)
		candidate := skeleton
		candidate.TargetInfra.NodeGroups = candidateNodeGroups
		candidate.TargetSpecList = candidateSpecList
		candidate.TargetOsImageList = candidateImageList
		candidate.TargetNlbList = targetNlbList

		overallStatus, overallStatusDesc, summary := calculateCandidateMatchRateWithDetails(
			csp, candidateNodeGroups, syntheticSrcInfra, candidateSpecList, candidateImageList, minMatchRate,
		)
		candidate.Status = overallStatus
		candidate.Description = fmt.Sprintf(
			"Candidate #%d | %s | %d NLB(s) | Match Rate: Min=%.1f%% Max=%.1f%% Avg=%.1f%% | %s",
			cIdx+1, overallStatus, len(targetNlbList), summary.MinMatchRate, summary.MaxMatchRate, summary.AvgMatchRate, overallStatusDesc,
		)

		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

// buildCaseBlueprints generates NodeGroups and NLB mappings based on hardware case and policy.
func buildCaseBlueprints(
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	memberNodes []onpremmodel.NodeProperty,
	nodeByMachineId map[string]onpremmodel.NodeProperty,
	hwCase HardwareCase,
	sizingPolicy string,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq, error) {
	backendName := rnlb.sourceNlb.Backend.Name

	switch hwCase {
	case CaseCrossVendorGpu:
		if sizingPolicy == SizingPolicyRightSizeUp {
			// Fail-fast without silent degradation when cross-vendor VM is physically impossible.
			return nil, nil, nil, fmt.Errorf(
				"NLB backend '%s': cross-vendor GPUs (e.g. NVIDIA and AMD) cannot be consolidated into a single NodeGroup under 'rightSizeUp'; please use 'dividing' policy or separate backend servers",
				backendName,
			)
		}
		// Dividing mode partitions by GPU vendor.
		return buildDividedVendorBlueprints(csp, region, connectionName, subnetId, skeleton, rnlb, memberNodes, nodeByMachineId)

	case CaseMixedCpuGpu:
		if sizingPolicy == SizingPolicyDividing {
			// Dividing mode partitions into CPU-only group and GPU group.
			return buildDividedCpuGpuBlueprints(csp, region, connectionName, subnetId, skeleton, rnlb, memberNodes, nodeByMachineId)
		}
		// Default rightSizeUp upsizes CPU members to GPU NodeGroup.
		return buildSingleConsolidatedBlueprint(csp, region, connectionName, subnetId, skeleton, rnlb, rnlb.memberMachineIds, nodeByMachineId, backendName)

	case CaseHeteroGpuModel:
		if sizingPolicy == SizingPolicyDividing {
			// Dividing mode partitions by GPU model.
			return buildDividedModelBlueprints(csp, region, connectionName, subnetId, skeleton, rnlb, memberNodes, nodeByMachineId)
		}
		// Default rightSizeUp upsizes to highest VRAM and accelerator count.
		return buildSingleConsolidatedBlueprint(csp, region, connectionName, subnetId, skeleton, rnlb, rnlb.memberMachineIds, nodeByMachineId, backendName)

	default:
		// CasePureCpu and CaseHomoGpuModel produce a single NodeGroup in both modes.
		return buildSingleConsolidatedBlueprint(csp, region, connectionName, subnetId, skeleton, rnlb, rnlb.memberMachineIds, nodeByMachineId, backendName)
	}
}

// buildSingleConsolidatedBlueprint creates one 1:1 pair of (Managed NLB : NodeGroup).
func buildSingleConsolidatedBlueprint(
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	machineIds []string,
	nodeByMachineId map[string]onpremmodel.NodeProperty,
	groupSuffix string,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq, error) {
	ngName := "ng-" + sanitizeName(groupSuffix)
	repNode := synthesizeGroupRepresentativeNode(machineIds, nodeByMachineId)
	repNodeWithFirewall := mergeNodesFirewallRules(machineIds, nodeByMachineId, repNode)
	repNodeWithFirewall = ensurePortOpenToPublic(repNodeWithFirewall, rnlb.backendPort)

	sg, err := RecommendSecurityGroup(csp, region, repNodeWithFirewall)
	if err != nil {
		log.Warn().Err(err).Str("nodeGroup", ngName).Msg("failed to recommend SG for consolidated group")
	}
	sg.Name = fmt.Sprintf("mig-sg-%s", sanitizeName(groupSuffix))
	sg.ConnectionName = connectionName
	sg.Description = fmt.Sprintf("Security group for %s", ngName)
	sg.VNetId = skeleton.TargetVNet.Name

	minDisk := getCspMinRootDiskSizeGB(csp)
	if hasGpu(repNode) && minDisk < 100 {
		minDisk = 100
	}
	rootDiskSize := max(int(repNode.RootDisk.TotalSize), minDisk)

	bp := nodeGroupBlueprint{
		representativeNode: repNode,
		isNlbRelated:       true,
		skeleton: cloudmodel.CreateNodeGroupReq{
			ConnectionName:   connectionName,
			Name:             ngName,
			VNetId:           skeleton.TargetVNet.Name,
			SubnetId:         subnetId,
			SecurityGroupIds: []string{sg.Name},
			SshKeyId:         skeleton.TargetSshKey.Name,
			RootDiskType:     "",
			RootDiskSize:     rootDiskSize,
			NodeGroupSize:    len(machineIds),
			Description:      fmt.Sprintf("Consolidated NodeGroup for NLB backend: %s", groupSuffix),
			Label:            map[string]string{"sourceMachineIds": strings.Join(machineIds, ",")},
		},
	}

	nlbReq := createManagedNlbReq(rnlb, ngName, groupSuffix)
	return []nodeGroupBlueprint{bp}, []cloudmodel.NlbReq{nlbReq}, []cloudmodel.SecurityGroupReq{sg}, nil
}

// buildDividedCpuGpuBlueprints creates separate pairs for CPU and GPU nodes.
func buildDividedCpuGpuBlueprints(
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	memberNodes []onpremmodel.NodeProperty,
	nodeByMachineId map[string]onpremmodel.NodeProperty,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq, error) {
	var cpuIds, gpuIds []string
	for _, m := range memberNodes {
		if hasGpu(m) {
			gpuIds = append(gpuIds, m.MachineId)
		} else {
			cpuIds = append(cpuIds, m.MachineId)
		}
	}

	var bps []nodeGroupBlueprint
	var nlbs []cloudmodel.NlbReq
	var sgs []cloudmodel.SecurityGroupReq

	if len(cpuIds) > 0 {
		b, n, s, err := buildSingleConsolidatedBlueprint(csp, region, connectionName, subnetId, skeleton, rnlb, cpuIds, nodeByMachineId, rnlb.sourceNlb.Backend.Name+"-cpu")
		if err != nil {
			return nil, nil, nil, err
		}
		bps = append(bps, b...)
		nlbs = append(nlbs, n...)
		sgs = append(sgs, s...)
	}

	if len(gpuIds) > 0 {
		b, n, s, err := buildSingleConsolidatedBlueprint(csp, region, connectionName, subnetId, skeleton, rnlb, gpuIds, nodeByMachineId, rnlb.sourceNlb.Backend.Name+"-gpu")
		if err != nil {
			return nil, nil, nil, err
		}
		bps = append(bps, b...)
		nlbs = append(nlbs, n...)
		sgs = append(sgs, s...)
	}

	return bps, nlbs, sgs, nil
}

// buildDividedVendorBlueprints creates separate pairs for each GPU vendor.
func buildDividedVendorBlueprints(
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	memberNodes []onpremmodel.NodeProperty,
	nodeByMachineId map[string]onpremmodel.NodeProperty,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq, error) {
	vendorMap := make(map[string][]string)
	for _, m := range memberNodes {
		vendor := "cpu"
		if len(m.GPUCards) > 0 {
			vendor = strings.ToLower(strings.TrimSpace(m.GPUCards[0].Vendor))
			if vendor == "" {
				vendor = "gpu"
			}
		}
		vendorMap[vendor] = append(vendorMap[vendor], m.MachineId)
	}

	var bps []nodeGroupBlueprint
	var nlbs []cloudmodel.NlbReq
	var sgs []cloudmodel.SecurityGroupReq

	for vendor, ids := range vendorMap {
		groupSuffix := fmt.Sprintf("%s-%s", rnlb.sourceNlb.Backend.Name, vendor)
		b, n, s, err := buildSingleConsolidatedBlueprint(csp, region, connectionName, subnetId, skeleton, rnlb, ids, nodeByMachineId, groupSuffix)
		if err != nil {
			return nil, nil, nil, err
		}
		bps = append(bps, b...)
		nlbs = append(nlbs, n...)
		sgs = append(sgs, s...)
	}

	return bps, nlbs, sgs, nil
}

// buildDividedModelBlueprints creates separate pairs for each GPU model.
func buildDividedModelBlueprints(
	csp, region, connectionName, subnetId string,
	skeleton cloudmodel.RecommendedInfra,
	rnlb resolvedNlb,
	memberNodes []onpremmodel.NodeProperty,
	nodeByMachineId map[string]onpremmodel.NodeProperty,
) ([]nodeGroupBlueprint, []cloudmodel.NlbReq, []cloudmodel.SecurityGroupReq, error) {
	modelMap := make(map[string][]string)
	for _, m := range memberNodes {
		key := "cpu"
		if len(m.GPUCards) > 0 {
			card := m.GPUCards[0]
			key = sanitizeName(fmt.Sprintf("%s-%s-%.0fgb", card.Vendor, card.Model, card.MemoryTotalGB))
		}
		modelMap[key] = append(modelMap[key], m.MachineId)
	}

	var bps []nodeGroupBlueprint
	var nlbs []cloudmodel.NlbReq
	var sgs []cloudmodel.SecurityGroupReq

	for key, ids := range modelMap {
		groupSuffix := fmt.Sprintf("%s-%s", rnlb.sourceNlb.Backend.Name, key)
		b, n, s, err := buildSingleConsolidatedBlueprint(csp, region, connectionName, subnetId, skeleton, rnlb, ids, nodeByMachineId, groupSuffix)
		if err != nil {
			return nil, nil, nil, err
		}
		bps = append(bps, b...)
		nlbs = append(nlbs, n...)
		sgs = append(sgs, s...)
	}

	return bps, nlbs, sgs, nil
}

// createManagedNlbReq generates a single CloudModel NlbReq binding to targetNodeGroupId.
func createManagedNlbReq(rnlb resolvedNlb, targetNodeGroupId, suffix string) cloudmodel.NlbReq {
	nlbType := "PUBLIC"
	if rnlb.sourceNlb.Listener.BindAddress != "" && rnlb.sourceNlb.Listener.BindAddress != "*" {
		nlbType = "INTERNAL"
	}

	listenerProtocol := normalizeProtocol(rnlb.sourceNlb.Listener.Protocol)
	targetGroupProtocol := normalizeProtocol(rnlb.sourceNlb.Backend.Protocol)

	hc := cloudmodel.NlbHealthCheckerReq{
		Interval:  defaultHealthCheckInterval,
		Threshold: defaultHealthCheckThreshold,
		Timeout:   defaultHealthCheckTimeout,
	}
	if rnlb.sourceNlb.HealthCheck.Enabled {
		if rnlb.sourceNlb.HealthCheck.Interval > 0 {
			hc.Interval = rnlb.sourceNlb.HealthCheck.Interval
		}
		if rnlb.sourceNlb.HealthCheck.Timeout > 0 {
			hc.Timeout = rnlb.sourceNlb.HealthCheck.Timeout
		}
		if rnlb.sourceNlb.HealthCheck.Threshold > 0 {
			hc.Threshold = rnlb.sourceNlb.HealthCheck.Threshold
		}
	}

	return cloudmodel.NlbReq{
		Description: fmt.Sprintf("Migrated from HAProxy backend: %s", suffix),
		Scope:       defaultNlbScope,
		Type:        nlbType,
		Listener: cloudmodel.NlbListenerReq{
			Protocol: listenerProtocol,
			Port:     fmt.Sprintf("%d", rnlb.sourceNlb.Listener.Port),
		},
		TargetGroup: cloudmodel.NlbTargetGroupReq{
			Protocol:    targetGroupProtocol,
			Port:        fmt.Sprintf("%d", rnlb.backendPort),
			NodeGroupId: targetNodeGroupId,
		},
		HealthChecker: hc,
	}
}
