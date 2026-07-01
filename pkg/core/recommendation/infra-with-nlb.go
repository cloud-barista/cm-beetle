package recommendation

import (
	"fmt"
	"net"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// Constants
// ============================================================================

const (
	defaultHealthCheckInterval  = 10
	defaultHealthCheckTimeout   = 10
	defaultHealthCheckThreshold = 3
	defaultNlbScope             = "REGION"

	// sizingPolicy controls how target cloud specs are selected relative to source node requirements.
	// "upsizing" picks the smallest spec that meets or exceeds the source node's CPU/memory.
	// TODO: expose as a user-configurable parameter in a future version.
	sizingPolicy = "upsizing"
)

// ============================================================================
// Package-level types (used across functions)
// ============================================================================

// resolvedNlb holds a source NLB after its backend server IPs have been correlated
// with source Nodes and the representative backend port has been determined.
type resolvedNlb struct {
	sourceNlb        onpremmodel.NlbProperty
	backendPort      int      // majority-voted representative port from backend servers
	memberMachineIds []string // machineIds of Nodes that are backend servers of this NLB
}

// nodeGroupBlueprint bundles a skeleton NodeGroup (no spec/image yet) with the
// representative source node used for spec/image recommendation, and a flag
// indicating whether this NodeGroup was formed from NLB backend nodes.
type nodeGroupBlueprint struct {
	skeleton           cloudmodel.CreateNodeGroupReq
	representativeNode onpremmodel.NodeProperty
	isNlbRelated       bool
}

// ============================================================================
// RecommendInfraWithNlbCandidates — entry point
// ============================================================================

// RecommendInfraWithNlbCandidates recommends multiple NLB-aware infrastructure
// candidates following the same Pareto-frontier pattern as RecommendVmInfraCandidates.
//
// Processing phases:
//  1. Build lookup indexes (IP → Node, MachineId → Node)
//  2. Correlate NLB backend server IPs with source Nodes; normalize backend ports
//  3. Build shared skeleton (vNet, SSH key)
//  4. Group source nodes into NodeGroups: NLB-related (N:1) and unrelated (1:1)
//  5. Find ranked spec-image pairs per NodeGroup (sizingPolicy = upsizing)
//  6. Build target NLB list — identical for all candidates
//  7. Assemble candidates: candidate i assigns the i-th ranked pair to each NodeGroup
func RecommendInfraWithNlbCandidates(desiredCsp, desiredRegion string, srcInfra onpremmodel.OnpremInfra, limit int, minMatchRate float64) ([]cloudmodel.RecommendedInfra, error) {
	if len(srcInfra.NLBs) == 0 {
		return nil, fmt.Errorf("sourceInfra.nlbs is empty")
	}

	csp := strings.ToLower(desiredCsp)
	region := strings.ToLower(desiredRegion)

	limitSpecs := GetDefaultSpecsLimit()
	limitImages := GetDefaultImagesLimit()

	// ── Phase 1: build lookup indexes ──────────────────────────────────────────
	nodeByIP := buildNodeByIPIndex(srcInfra.Nodes)
	nodeByMachineId := buildNodeByMachineIdIndex(srcInfra.Nodes)
	log.Debug().Int("nodeCount", len(srcInfra.Nodes)).Int("ipEntries", len(nodeByIP)).Msg("IP index built")

	// ── Phase 2: correlate NLB backend IPs with source Nodes; normalize backend ports ──
	var resolvedNlbs []resolvedNlb
	var warnings []string
	// Used in Phase 4 to skip NLB-member nodes when building unrelated (1:1) NodeGroups.
	// A node CAN appear in multiple NLB backends and will be in each of their NodeGroups.
	nlbMemberMachineIds := map[string]string{} // machineId → first backendName that claimed it

	for _, nlb := range srcInfra.NLBs {
		backendName := nlb.Backend.Name
		var matchedMachineIds []string

		for _, srv := range nlb.Backend.Servers {
			node, ok := nodeByIP[srv.IP]
			if !ok {
				warnings = append(warnings, fmt.Sprintf(
					"NLB backend '%s': server IP %s not found in Nodes[] — skipping this server",
					backendName, srv.IP))
				continue
			}
			matchedMachineIds = append(matchedMachineIds, node.MachineId)
		}

		if len(matchedMachineIds) == 0 {
			warnings = append(warnings, fmt.Sprintf(
				"NLB backend '%s': no backend IPs matched any Node — skipping this NLB", backendName))
			continue
		}

		backendPort, portWarning := resolvePort(nlb.Backend.Servers)
		if portWarning != "" {
			warnings = append(warnings, fmt.Sprintf("NLB backend '%s': %s", backendName, portWarning))
		}

		for _, machineId := range matchedMachineIds {
			if _, alreadyMember := nlbMemberMachineIds[machineId]; !alreadyMember {
				nlbMemberMachineIds[machineId] = backendName
			} else {
				// Node serves multiple NLB backends → it appears in each backend's NodeGroup.
				// In the cloud, a separate VM is provisioned per NodeGroup (role separation).
				warnings = append(warnings, fmt.Sprintf(
					"Node %s belongs to multiple NLB backends ('%s' and '%s'). "+
						"A separate VM will be created for each backend in the target cloud.",
					machineId, nlbMemberMachineIds[machineId], backendName))
			}
		}

		if nlb.Backend.Balance != "" {
			warnings = append(warnings, fmt.Sprintf(
				"NLB backend '%s': load-balancing algorithm '%s' cannot be directly mapped to cloud NLB. CSP default algorithm will be used.",
				backendName, nlb.Backend.Balance))
		}

		resolvedNlbs = append(resolvedNlbs, resolvedNlb{
			sourceNlb:        nlb,
			backendPort:      backendPort,
			memberMachineIds: matchedMachineIds,
		})
	}

	if len(resolvedNlbs) == 0 {
		return nil, fmt.Errorf("no NLBs could be resolved from source configuration: %s", strings.Join(warnings, "; "))
	}

	// ── Phase 3: build shared skeleton (vNet, SSH key) ─────────────────────────
	connectionName := fmt.Sprintf("%s-%s", csp, region)

	skeleton := cloudmodel.RecommendedInfra{
		TargetCloud: cloudmodel.CloudProperty{Csp: csp, Region: region},
		TargetInfra: cloudmodel.InfraReq{
			Name:        "infra101",
			Description: "NLB-aware recommended infrastructure for cloud migration",
			NodeGroups:  []cloudmodel.CreateNodeGroupReq{},
		},
	}

	recommendedVNetList, err := RecommendVNet(csp, region, srcInfra)
	if err != nil || len(recommendedVNetList) == 0 {
		log.Warn().Err(err).Msg("failed to recommend vNet for NLB-aware infra")
	}
	if len(recommendedVNetList) > 0 {
		skeleton.TargetVNet = recommendedVNetList[0]
	}
	skeleton.TargetVNet.Name = "mig-vnet-01"
	skeleton.TargetVNet.Description = "a recommended vNet for migration"
	for i := range skeleton.TargetVNet.SubnetInfoList {
		skeleton.TargetVNet.SubnetInfoList[i].Name = fmt.Sprintf("mig-subnet-%02d", i+1)
		skeleton.TargetVNet.SubnetInfoList[i].Description = "a recommended subnet for migration"
	}
	firstSubnetId := ""
	if len(skeleton.TargetVNet.SubnetInfoList) > 0 {
		firstSubnetId = skeleton.TargetVNet.SubnetInfoList[0].Name
	}

	skeleton.TargetSshKey = cloudmodel.SshKeyReq{
		Name:           "mig-sshkey-01",
		ConnectionName: connectionName,
		Description:    "SSH key pair for migration (Note: provided ONLY once, MUST be downloaded)",
	}

	// ── Phase 4: group source nodes into NodeGroups; build SecurityGroups ──────
	// NLB-related (N:1): one NodeGroup per NLB backend; NodeGroupSize = backend member count.
	// Unrelated   (1:1): one NodeGroup per node not in any NLB backend; NodeGroupSize = 1.

	var ngBlueprints []nodeGroupBlueprint
	var deduplicatedSgList []cloudmodel.SecurityGroupReq

	// NLB-related NodeGroups — added first so indexes remain stable when NLB target list references them.
	// NodeGroup name "ng-<backendName>" must match the nodeGroupId used in buildTargetNlbList.
	for _, rnlb := range resolvedNlbs {
		backendName := rnlb.sourceNlb.Backend.Name
		nodeGroupSize := len(rnlb.memberMachineIds)
		ngId := "ng-" + sanitizeName(backendName)
		// Synthetic node: CPU from max-vCPU member; memory and root disk are per-dimension maxima.
		syntheticNode := synthesizeGroupRepresentativeNode(rnlb.memberMachineIds, nodeByMachineId)
		nodeWithMergedFirewallRules := mergeNodesFirewallRules(rnlb.memberMachineIds, nodeByMachineId, syntheticNode)

		// Cloud NLBs (AWS, GCP, Azure) operate in pass-through mode: they preserve the original
		// client source IP instead of SNAT-ing to their own IP. In the source HAProxy setup the
		// backend port (8086) was restricted to the internal VPC CIDR because haproxy was a proxy
		// that always forwarded from its own internal IP. In the cloud NLB equivalent the backend
		// VMs receive packets whose source IP is the external client IP, so the SG must allow
		// the target port from 0.0.0.0/0. (Health checks still originate from VPC-internal NLB
		// node IPs, so the existing VPC-CIDR rule continues to satisfy them.)
		nodeWithMergedFirewallRules = ensurePortOpenToPublic(nodeWithMergedFirewallRules, rnlb.backendPort)

		sg, sgErr := RecommendSecurityGroup(csp, region, nodeWithMergedFirewallRules)
		if sgErr != nil {
			log.Warn().Err(sgErr).Str("backend", backendName).Msg("failed to recommend SG for NLB NodeGroup")
		}
		exists, _, existingSg := containSg(deduplicatedSgList, sg)
		if !exists {
			sg.Name = fmt.Sprintf("mig-sg-%02d", len(deduplicatedSgList)+1)
			sg.ConnectionName = connectionName
			sg.Description = fmt.Sprintf("Recommended security group for NLB backend %s", backendName)
			sg.VNetId = skeleton.TargetVNet.Name
			deduplicatedSgList = append(deduplicatedSgList, sg)
		} else {
			sg = existingSg
		}

		rootDiskSize := max(int(syntheticNode.RootDisk.TotalSize), getCspMinRootDiskSizeGB(csp))

		ngBlueprints = append(ngBlueprints, nodeGroupBlueprint{
			representativeNode: syntheticNode,
			isNlbRelated:       true,
			skeleton: cloudmodel.CreateNodeGroupReq{
				ConnectionName:   connectionName,
				Name:             ngId,
				VNetId:           skeleton.TargetVNet.Name,
				SubnetId:         firstSubnetId,
				SecurityGroupIds: []string{sg.Name},
				SshKeyId:         skeleton.TargetSshKey.Name,
				RootDiskType:     "",
				RootDiskSize:     rootDiskSize,
				NodeGroupSize:    nodeGroupSize,
				Description:      fmt.Sprintf("Recommended VM for NLB backend %s (%d nodes)", backendName, nodeGroupSize),
				Label: map[string]string{
					"sourceMachineIds": strings.Join(rnlb.memberMachineIds, ","),
					"nlbBackend":       backendName,
				},
			},
		})
	}

	// NLB-unrelated NodeGroups — one per source node; NLB-member nodes are skipped here.
	for i, node := range srcInfra.Nodes {
		if _, isMember := nlbMemberMachineIds[node.MachineId]; isMember {
			continue
		}

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

		rootDiskSize := max(int(node.RootDisk.TotalSize), getCspMinRootDiskSizeGB(csp))

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
				Description:      fmt.Sprintf("Recommended VM %02d for %s", i+1, node.MachineId),
				Label:            map[string]string{"sourceMachineIds": node.MachineId},
			},
		})
	}

	skeleton.TargetSecurityGroupList = deduplicatedSgList

	// ── Phase 5: find compatible spec-image pairs per NodeGroup ─────────────
	// pairsByGroup[ngIdx] holds the ranked (spec, image) pairs for ngBlueprints[ngIdx].
	// Specs are selected using sizingPolicy (currently hardcoded to "upsizing").
	// TODO: pass sizingPolicy as a user-configurable parameter when RecommendVmSpecs supports it.
	pairsByGroup := make([][]CompatibleSpecImagePair, len(ngBlueprints))
	for ngIdx, bp := range ngBlueprints {
		specList, _, specErr := RecommendVmSpecs(csp, region, bp.representativeNode, limitSpecs)
		if specErr != nil {
			log.Warn().Err(specErr).Str("machineId", bp.representativeNode.MachineId).Str("sizingPolicy", sizingPolicy).Msg("failed to recommend VM specs")
		}
		imageList, imageErr := RecommendVmOsImages(csp, region, bp.representativeNode, limitImages)
		if imageErr != nil {
			log.Warn().Err(imageErr).Str("machineId", bp.representativeNode.MachineId).Msg("failed to recommend OS images")
		}

		if len(specList) == 0 || len(imageList) == 0 {
			log.Warn().Str("machineId", bp.representativeNode.MachineId).Msg("no compatible spec or image found for NodeGroup")
			continue
		}

		pairs, pairErr := FindCompatibleVmSpecAndImagePairs(specList, imageList, csp)
		if pairErr != nil {
			log.Warn().Err(pairErr).Str("machineId", bp.representativeNode.MachineId).Msg("failed to find compatible pairs; will use first available")
			if len(specList) > 0 && len(imageList) > 0 {
				pairsByGroup[ngIdx] = []CompatibleSpecImagePair{{Spec: specList[0], Image: imageList[0]}}
			}
		} else {
			pairsByGroup[ngIdx] = pairs
		}

		log.Debug().
			Str("machineId", bp.representativeNode.MachineId).
			Bool("isNlbRelated", bp.isNlbRelated).
			Int("pairCount", len(pairsByGroup[ngIdx])).
			Msg("Compatible pairs found for NodeGroup")
	}

	// ── Phase 6: build target NLB list — identical for all candidates ─────────
	targetNlbList := buildTargetNlbList(resolvedNlbs)
	targetNlbList = sanitizeNlbListByCsp(targetNlbList, csp)

	// ── Phase 7: assemble candidates — candidate i uses the i-th ranked pair per NodeGroup ──
	// maxCandidates = max pairs available across all NodeGroups, capped at limit.
	// NodeGroups with fewer than i+1 pairs are skipped for candidate i.
	maxCandidates := 0
	for _, pairs := range pairsByGroup {
		if len(pairs) > maxCandidates {
			maxCandidates = len(pairs)
		}
	}
	if maxCandidates > limit {
		maxCandidates = limit
	}

	log.Debug().Int("requestedLimit", limit).Int("maxCandidates", maxCandidates).Msg("Candidate limit determined")

	var candidates []cloudmodel.RecommendedInfra

	for candidateIdx := 0; candidateIdx < maxCandidates; candidateIdx++ {
		// Start from the skeleton NodeGroups; fill in spec/image for this candidate.
		candidateNodeGroups := make([]cloudmodel.CreateNodeGroupReq, len(ngBlueprints))
		for ngIdx := range ngBlueprints {
			candidateNodeGroups[ngIdx] = ngBlueprints[ngIdx].skeleton
		}

		var candidateSpecList []cloudmodel.SpecInfo
		var candidateImageList []cloudmodel.ImageInfo

		for ngIdx, bp := range ngBlueprints {
			pairs := pairsByGroup[ngIdx]
			if len(pairs) == 0 {
				log.Warn().Str("machineId", bp.representativeNode.MachineId).
					Msgf("candidate %d: no pairs available for this NodeGroup, skipping", candidateIdx+1)
				continue
			}
			if candidateIdx >= len(pairs) {
				log.Warn().Str("machineId", bp.representativeNode.MachineId).
					Msgf("candidate %d: only %d pairs available, skipping", candidateIdx+1, len(pairs))
				continue
			}

			selectedPair := pairs[candidateIdx]
			selectedSpec := selectedPair.Spec
			selectedImage := selectedPair.Image

			matchRateVec := calculateMatchRateVector(csp, bp.representativeNode, selectedSpec, selectedImage)

			log.Debug().
				Str("machineId", bp.representativeNode.MachineId).
				Int("candidateIdx", candidateIdx).
				Str("specId", selectedSpec.Id).
				Str("imageId", selectedImage.Id).
				Float64("cpuMatchRate", matchRateVec.CPU).
				Float64("memoryMatchRate", matchRateVec.Memory).
				Float64("imageMatchRate", matchRateVec.Image).
				Msg("Selected spec-image pair for NodeGroup candidate")

			// Preflight check: resolve latest CSP image name and suggest disk type
			precheck, prefErr := PreflightCheckCspProvisioning(
				selectedSpec.Id, selectedImage.Id, selectedImage.CspImageName, "",
			)
			if prefErr != nil {
				log.Warn().Err(prefErr).Str("machineId", bp.representativeNode.MachineId).Msg("preflight check failed; using cached image")
			} else {
				if precheck.SuggestedSystemDisk != "" {
					candidateNodeGroups[ngIdx].RootDiskType = precheck.SuggestedSystemDisk
				}
				if precheck.ResolvedCspImageName != "" && precheck.ResolvedCspImageName != selectedImage.CspImageName {
					log.Info().Str("machineId", bp.representativeNode.MachineId).
						Str("from", selectedImage.CspImageName).
						Str("to", precheck.ResolvedCspImageName).
						Msg("image resolved to latest")
					candidateNodeGroups[ngIdx].CspImageName = precheck.ResolvedCspImageName
				}
			}

			candidateNodeGroups[ngIdx].SpecId = selectedSpec.Id
			candidateNodeGroups[ngIdx].ImageId = selectedImage.Id
			candidateNodeGroups[ngIdx].Description = fmt.Sprintf(
				"%s | Match Rate: CPU=%.1f%% Memory=%.1f%% Image=%.1f%%",
				ngBlueprints[ngIdx].skeleton.Description,
				matchRateVec.CPU, matchRateVec.Memory, matchRateVec.Image,
			)

			// Deduplicate spec and image across NodeGroups in this candidate
			specAlreadyAdded := false
			for _, s := range candidateSpecList {
				if s.CspSpecName == selectedSpec.CspSpecName {
					specAlreadyAdded = true
					selectedSpec = s
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
					selectedImage = img
					break
				}
			}
			if !imageAlreadyAdded {
				candidateImageList = append(candidateImageList, selectedImage)
			}
		}

		// Build a synthetic source infra that has one representative node per NodeGroup,
		// matching the ordering of candidateNodeGroups, for match-rate calculation.
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

		nlbWarningNote := ""
		if len(warnings) > 0 {
			nlbWarningNote = fmt.Sprintf(" | %d NLB warning(s): %s", len(warnings), strings.Join(warnings, "; "))
		}
		candidate.Description = fmt.Sprintf(
			"Candidate #%d | %s | %d NLB(s) | Overall Match Rate: Min=%.1f%% Max=%.1f%% Avg=%.1f%% | %s%s",
			candidateIdx+1,
			overallStatus,
			len(targetNlbList),
			summary.MinMatchRate, summary.MaxMatchRate, summary.AvgMatchRate,
			overallStatusDesc,
			nlbWarningNote,
		)

		candidates = append(candidates, candidate)

		log.Debug().
			Int("candidateIdx", candidateIdx).
			Str("status", overallStatus).
			Float64("avgMatchRate", summary.AvgMatchRate).
			Int("nlbs", len(targetNlbList)).
			Msg("Candidate assembled")
	}

	log.Info().
		Int("candidates", len(candidates)).
		Int("resolvedNlbs", len(resolvedNlbs)).
		Int("nodeGroups", len(ngBlueprints)).
		Int("warnings", len(warnings)).
		Msg("infraWithNlb candidates recommendation completed")

	return candidates, nil
}

// ============================================================================
// NLB target mapping
// ============================================================================

// buildTargetNlbList converts resolved NLB entries to cloudmodel.NlbReq list.
// The list is identical for all candidates since NLB configuration does not vary by spec/image choice.
func buildTargetNlbList(resolvedNlbs []resolvedNlb) []cloudmodel.NlbReq {
	var targetNlbList []cloudmodel.NlbReq

	for _, rnlb := range resolvedNlbs {
		nlbType := "PUBLIC"
		if rnlb.sourceNlb.Listener.BindAddress != "" && rnlb.sourceNlb.Listener.BindAddress != "*" {
			nlbType = "INTERNAL"
		}

		listenerProtocol := normalizeProtocol(rnlb.sourceNlb.Listener.Protocol)
		targetGroupProtocol := normalizeProtocol(rnlb.sourceNlb.Backend.Protocol)
		nodeGroupId := "ng-" + sanitizeName(rnlb.sourceNlb.Backend.Name)

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

		targetNlbList = append(targetNlbList, cloudmodel.NlbReq{
			Description: fmt.Sprintf("Migrated from HAProxy backend: %s", rnlb.sourceNlb.Backend.Name),
			Scope:       defaultNlbScope,
			Type:        nlbType,
			Listener: cloudmodel.NlbListenerReq{
				Protocol: listenerProtocol,
				Port:     fmt.Sprintf("%d", rnlb.sourceNlb.Listener.Port),
			},
			TargetGroup: cloudmodel.NlbTargetGroupReq{
				Protocol:    targetGroupProtocol,
				Port:        fmt.Sprintf("%d", rnlb.backendPort),
				NodeGroupId: nodeGroupId,
			},
			HealthChecker: hc,
		})
	}

	return targetNlbList
}

// sanitizeNlbListByCsp applies CSP-specific platform constraints and adjustments to the target NLB list.
func sanitizeNlbListByCsp(nlbList []cloudmodel.NlbReq, csp string) []cloudmodel.NlbReq {
	for i := range nlbList {
		// Azure Load Balancers do not support custom health check timeouts.
		// Setting Timeout to -1 tells Tumblebug/Spider to omit it.
		if csp == "azure" {
			nlbList[i].HealthChecker.Timeout = -1
		}

		// IBM Cloud Load Balancer requires the health check timeout to be strictly
		// less than the health check interval (delay).
		if csp == "ibm" {
			if nlbList[i].HealthChecker.Timeout >= nlbList[i].HealthChecker.Interval {
				// Adjust timeout to be half of the interval (capped at 1 or greater)
				newTimeout := nlbList[i].HealthChecker.Interval / 2
				if newTimeout < 1 {
					newTimeout = 1
				}
				nlbList[i].HealthChecker.Timeout = newTimeout
			}
		}

		// GCP External Passthrough NLB uses target pools and does NOT perform port translation.
		// Traffic arrives at backend VMs on the same destination port as the forwarding rule
		// (the listener port), not the application's backend port.
		// To ensure traffic reaches the application, the listener port must equal the backend port.
		// We use the backend (application) port as the authoritative value since the application
		// cannot be reconfigured as part of the migration.
		if csp == "gcp" {
			nlbList[i].Listener.Port = nlbList[i].TargetGroup.Port
		}
	}
	return nlbList
}

// buildSyntheticSrcInfra builds an OnpremInfra containing one representative node per
// NodeGroup (in the same order as ngBlueprints). This is passed to
// calculateCandidateMatchRateWithDetails, which expects one source node per NodeGroup.
func buildSyntheticSrcInfra(ngBlueprints []nodeGroupBlueprint) onpremmodel.OnpremInfra {
	var representativeNodes []onpremmodel.NodeProperty
	for _, bp := range ngBlueprints {
		representativeNodes = append(representativeNodes, bp.representativeNode)
	}
	return onpremmodel.OnpremInfra{Nodes: representativeNodes}
}

// ============================================================================
// Helpers — node selection and merging
// ============================================================================

// buildNodeByMachineIdIndex builds a machineId → NodeProperty lookup index.
func buildNodeByMachineIdIndex(nodes []onpremmodel.NodeProperty) map[string]onpremmodel.NodeProperty {
	index := make(map[string]onpremmodel.NodeProperty, len(nodes))
	for _, node := range nodes {
		index[node.MachineId] = node
	}
	return index
}

// selectRepresentativeNode picks the node with the highest vCPU count (CPUs × threads).
// On tie, higher memory wins.
func selectRepresentativeNode(memberMachineIds []string, nodeByMachineId map[string]onpremmodel.NodeProperty) onpremmodel.NodeProperty {
	var bestNode onpremmodel.NodeProperty
	bestVCPU := uint32(0)
	for _, machineId := range memberMachineIds {
		node, ok := nodeByMachineId[machineId]
		if !ok {
			continue
		}
		threads := node.CPU.Threads
		if threads == 0 {
			threads = 1
		}
		vCPU := node.CPU.Cpus * threads
		if vCPU > bestVCPU || (vCPU == bestVCPU && node.Memory.TotalSize > bestNode.Memory.TotalSize) {
			bestNode = node
			bestVCPU = vCPU
		}
	}
	return bestNode
}

// synthesizeGroupRepresentativeNode creates a synthetic NodeProperty for spec/image recommendation.
// CPU is taken from the node with the highest vCPU (which also serves as the OS reference).
// Memory and root disk are per-dimension maxima across all member nodes, so the recommended
// spec satisfies every member's resource requirements under the upsizing policy.
func synthesizeGroupRepresentativeNode(memberMachineIds []string, nodeByMachineId map[string]onpremmodel.NodeProperty) onpremmodel.NodeProperty {
	// Base: node with highest vCPU — provides CPU and OS reference.
	synthetic := selectRepresentativeNode(memberMachineIds, nodeByMachineId)

	// Override memory and root disk with per-dimension maxima.
	for _, machineId := range memberMachineIds {
		node, ok := nodeByMachineId[machineId]
		if !ok {
			continue
		}
		if node.Memory.TotalSize > synthetic.Memory.TotalSize {
			synthetic.Memory.TotalSize = node.Memory.TotalSize
		}
		if node.RootDisk.TotalSize > synthetic.RootDisk.TotalSize {
			synthetic.RootDisk.TotalSize = node.RootDisk.TotalSize
		}
	}

	return synthetic
}

// mergeNodesFirewallRules returns a copy of representativeNode whose FirewallTable is the union
// of all firewall rules from every member node. Used so the Security Group covers all members.
func mergeNodesFirewallRules(memberMachineIds []string, nodeByMachineId map[string]onpremmodel.NodeProperty, representativeNode onpremmodel.NodeProperty) onpremmodel.NodeProperty {
	merged := representativeNode
	ruleSet := make(map[string]onpremmodel.FirewallRuleProperty)
	for _, rule := range representativeNode.FirewallTable {
		ruleSet[firewallRuleKey(rule)] = rule
	}
	for _, machineId := range memberMachineIds {
		node, ok := nodeByMachineId[machineId]
		if !ok || node.MachineId == representativeNode.MachineId {
			continue
		}
		for _, rule := range node.FirewallTable {
			ruleSet[firewallRuleKey(rule)] = rule
		}
	}
	var mergedRules []onpremmodel.FirewallRuleProperty
	for _, rule := range ruleSet {
		mergedRules = append(mergedRules, rule)
	}
	merged.FirewallTable = mergedRules
	return merged
}

// firewallRuleKey produces a deduplication key for a firewall rule.
func firewallRuleKey(rule onpremmodel.FirewallRuleProperty) string {
	return fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s",
		rule.Action, rule.Direction, rule.Protocol, rule.SrcCIDR, rule.SrcPorts, rule.DstCIDR, rule.DstPorts)
}

// ============================================================================
// Helpers — IP index and port resolution
// ============================================================================

// buildNodeByIPIndex builds a map from each node's IP (CIDR-stripped) to its NodeProperty.
func buildNodeByIPIndex(nodes []onpremmodel.NodeProperty) map[string]onpremmodel.NodeProperty {
	index := make(map[string]onpremmodel.NodeProperty, len(nodes)*2)
	for _, node := range nodes {
		for _, iface := range node.Interfaces {
			for _, cidr := range iface.IPv4CidrBlocks {
				ip, _, err := net.ParseCIDR(cidr)
				if err != nil {
					ip = net.ParseIP(cidr)
				}
				if ip != nil {
					index[ip.String()] = node
				}
			}
		}
	}
	return index
}

// resolvePort picks the representative backend port using a three-tier priority:
//
//  1. Majority vote   — the port used by the most servers wins outright.
//  2. Highest weight  — on a frequency tie, the port belonging to the highest-weight server wins.
//  3. First server    — if weight is also tied, fall back to the first server's port.
//
// A non-empty description is returned whenever the result required tie-breaking.
func resolvePort(servers []onpremmodel.NlbServerProperty) (int, string) {
	if len(servers) == 0 {
		return 0, "no servers"
	}

	// Collect frequency and maximum weight per port.
	portFrequency := make(map[int]int)
	portMaxWeight := make(map[int]int)
	for _, srv := range servers {
		portFrequency[srv.Port]++
		if srv.Weight > portMaxWeight[srv.Port] {
			portMaxWeight[srv.Port] = srv.Weight
		}
	}

	// All servers agree on one port — no tie-breaking needed.
	if len(portFrequency) == 1 {
		for port := range portFrequency {
			return port, ""
		}
	}

	// ── Priority 1: majority vote ──────────────────────────────────────────
	maxFreq := 0
	for _, count := range portFrequency {
		if count > maxFreq {
			maxFreq = count
		}
	}

	var majorityCandidates []int
	for port, count := range portFrequency {
		if count == maxFreq {
			majorityCandidates = append(majorityCandidates, port)
		}
	}

	if len(majorityCandidates) == 1 {
		return majorityCandidates[0], fmt.Sprintf(
			"backend server ports differ; resolved by majority vote — using port %d", majorityCandidates[0])
	}

	// ── Priority 2: highest weight ─────────────────────────────────────────
	maxWeight := -1
	for _, port := range majorityCandidates {
		if portMaxWeight[port] > maxWeight {
			maxWeight = portMaxWeight[port]
		}
	}

	var weightCandidates []int
	for _, port := range majorityCandidates {
		if portMaxWeight[port] == maxWeight {
			weightCandidates = append(weightCandidates, port)
		}
	}

	if len(weightCandidates) == 1 {
		return weightCandidates[0], fmt.Sprintf(
			"backend server ports differ with no majority; resolved by highest-weight server (weight=%d) — using port %d",
			maxWeight, weightCandidates[0])
	}

	// ── Priority 3: first server's port ───────────────────────────────────
	firstPort := servers[0].Port
	return firstPort, fmt.Sprintf(
		"backend server ports differ with no majority or weight advantage; using first server's port %d — manual review recommended",
		firstPort)
}

// normalizeProtocol converts to Tumblebug-expected uppercase. Defaults to "TCP".
func normalizeProtocol(proto string) string {
	switch strings.ToUpper(proto) {
	case "UDP":
		return "UDP"
	case "HTTP":
		return "HTTP"
	default:
		return "TCP"
	}
}

// ensurePortOpenToPublic returns a copy of node whose FirewallTable is guaranteed to contain
// an inbound TCP rule that allows the given port from 0.0.0.0/0.
// If such a rule already exists (exact port or wildcard "*") it is a no-op.
func ensurePortOpenToPublic(node onpremmodel.NodeProperty, port int) onpremmodel.NodeProperty {
	portStr := fmt.Sprintf("%d", port)
	for _, rule := range node.FirewallTable {
		if rule.Direction != "inbound" || rule.Action != "allow" {
			continue
		}
		proto := strings.ToLower(rule.Protocol)
		if proto != "tcp" && proto != "*" {
			continue
		}
		if rule.DstPorts != portStr && rule.DstPorts != "*" {
			continue
		}
		if rule.SrcCIDR == "0.0.0.0/0" || rule.SrcCIDR == "*" || rule.SrcCIDR == "" {
			return node // already open to public
		}
	}
	node.FirewallTable = append(node.FirewallTable, onpremmodel.FirewallRuleProperty{
		Action:    "allow",
		Direction: "inbound",
		Protocol:  "tcp",
		SrcCIDR:   "0.0.0.0/0",
		SrcPorts:  "*",
		DstCIDR:   "0.0.0.0/0",
		DstPorts:  portStr,
	})
	return node
}

// sanitizeName converts a name to a safe NodeGroup ID component (lowercase, hyphens only).
func sanitizeName(name string) string {
	s := strings.ToLower(name)
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			b.WriteRune(r)
		} else {
			b.WriteRune('-')
		}
	}
	result := b.String()
	for strings.Contains(result, "--") {
		result = strings.ReplaceAll(result, "--", "-")
	}
	return strings.Trim(result, "-")
}
