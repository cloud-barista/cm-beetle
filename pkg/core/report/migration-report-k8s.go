/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package report provides migration report generation logic for K8s infrastructure
package report

import (
	"fmt"
	"math"
	"strings"
	"time"

	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
	"github.com/rs/zerolog/log"
)

// GenerateK8sMigrationReport generates a comprehensive migration report for K8s infrastructure.
func GenerateK8sMigrationReport(nsId, clusterId string, sourceInfra onpremmodel.OnpremInfra) (*K8sMigrationReport, error) {
	log.Info().Msgf("Generating K8s migration report (nsId: %s, clusterId: %s)", nsId, clusterId)

	// Step 1: Generate source infrastructure summary
	sourceClusterName := "onprem-k8s"
	if sourceInfra.K8sCluster != nil && sourceInfra.K8sCluster.Name != "" {
		sourceClusterName = sourceInfra.K8sCluster.Name
	}
	sourceSummary, err := summary.GenerateSourceInfraSummary(sourceClusterName, sourceInfra)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate source infrastructure summary for K8s report")
		return nil, fmt.Errorf("failed to generate source infrastructure summary: %w", err)
	}

	// Step 2: Generate target K8s infrastructure summary
	targetSummary, err := summary.GenerateTargetK8sInfraSummary(nsId, clusterId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate target K8s infrastructure summary for report")
		return nil, fmt.Errorf("failed to generate target K8s infrastructure summary: %w", err)
	}

	// Step 3: Build report metadata
	metadata := ReportMetadata{
		GeneratedAt:   time.Now(),
		MigrationID:   fmt.Sprintf("%s/%s", nsId, clusterId),
		Namespace:     nsId,
		InfraID:       clusterId,
		ReportVersion: "1.0",
	}

	// Step 4: Partition worker nodes and excluded control-plane nodes
	workerNodes, excludedWorkers := partitionSourceNodes(sourceInfra.Nodes)

	// Step 5: Build executive summary
	executiveSummary := buildK8sExecutiveSummary(sourceClusterName, len(sourceInfra.Nodes), workerNodes, targetSummary)

	// Step 6: Build version analysis
	versionAnalysis := buildK8sVersionAnalysis(sourceInfra.K8sCluster, targetSummary.Overview.K8sVersion)

	// Step 7: Correlate worker nodes to node groups and analyze sizing
	nodeGroupMappings, sizingAnalysis := correlateWorkersToNodeGroups(workerNodes, targetSummary.NodeGroupList)

	// Step 8: Calculate resource comparisons and deltas
	resourceComparison := buildK8sResourceComparison(workerNodes, targetSummary)

	// Step 9: Build cost summary
	costSummary := buildK8sCostSummary(targetSummary)

	// Step 10: Generate architectural recommendations
	recommendations := generateK8sRecommendations(targetSummary, resourceComparison)

	report := &K8sMigrationReport{
		Metadata:           metadata,
		ExecutiveSummary:   executiveSummary,
		VersionAnalysis:    versionAnalysis,
		NodeGroupMappings:  nodeGroupMappings,
		SizingAnalysis:     sizingAnalysis,
		ExcludedWorkers:    excludedWorkers,
		ResourceComparison: resourceComparison,
		CostSummary:        costSummary,
		Recommendations:    recommendations,
		SourceDetails:      sourceSummary,
		TargetDetails:      targetSummary,
	}

	log.Info().Msgf("Successfully generated K8s migration report (nsId: %s, clusterId: %s)", nsId, clusterId)
	return report, nil
}

// partitionSourceNodes separates worker nodes from control-plane or master nodes.
func partitionSourceNodes(nodes []onpremmodel.NodeProperty) ([]onpremmodel.NodeProperty, []ExcludedWorkerItem) {
	var workers []onpremmodel.NodeProperty
	var excluded []ExcludedWorkerItem

	for _, n := range nodes {
		role := n.Role
		if role == "" {
			lowerHost := strings.ToLower(n.Hostname)
			if strings.Contains(lowerHost, "master") || strings.Contains(lowerHost, "control") {
				role = "control-plane"
			} else {
				role = "worker"
			}
		}

		if strings.EqualFold(role, "worker") {
			workers = append(workers, n)
		} else {
			excluded = append(excluded, ExcludedWorkerItem{
				Hostname: n.Hostname,
				Role:     role,
				Reason:   "Control-plane abstracted into CSP managed control plane",
			})
		}
	}

	return workers, excluded
}

// buildK8sExecutiveSummary builds the executive summary for K8s migration.
func buildK8sExecutiveSummary(clusterName string, totalNodes int, workers []onpremmodel.NodeProperty, targetSummary *summary.TargetK8sInfraSummary) K8sExecutiveSummary {
	migrationStatus := "Completed"
	if len(workers) == 0 && totalNodes > 0 {
		migrationStatus = "No Workers Migrated"
	}

	return K8sExecutiveSummary{
		MigrationStatus:      migrationStatus,
		SourceClusterName:    clusterName,
		SourceTotalNodes:     totalNodes,
		SourceWorkerNodes:    len(workers),
		TargetCloud:          targetSummary.Overview.TargetCloud,
		TargetRegion:         targetSummary.Overview.TargetRegion,
		TargetClusterName:    targetSummary.Overview.ClusterName,
		TargetNodeGroupCount: targetSummary.Overview.TotalNodeGroupCount,
		TargetTotalNodes:     targetSummary.Overview.TotalWorkerNodeCount,
		MonthlyCostUSD:       float64(targetSummary.CostEstimation.TotalCostPerMonth),
	}
}

// buildK8sVersionAnalysis builds K8s version delta and upgrade status.
func buildK8sVersionAnalysis(clusterProp *onpremmodel.K8sClusterProperty, targetVersion string) K8sVersionAnalysis {
	srcVer := "Unknown"
	if clusterProp != nil && clusterProp.Version != "" {
		srcVer = clusterProp.Version
	}

	upgradeStatus := "Aligned"
	notes := "Kubernetes versions aligned with target CSP compatibility"
	if srcVer != "Unknown" && targetVersion != "" {
		if srcVer == targetVersion {
			upgradeStatus = "Identical"
			notes = "Source and target Kubernetes versions are identical"
		} else {
			upgradeStatus = "Upgraded (Compatible)"
			notes = "Upgraded to target CSP supported minimum stable version"
		}
	}

	return K8sVersionAnalysis{
		SourceVersion: srcVer,
		TargetVersion: targetVersion,
		UpgradeStatus: upgradeStatus,
		Notes:         notes,
	}
}

// correlateWorkersToNodeGroups correlates source worker nodes to target cloud node groups.
func correlateWorkersToNodeGroups(workers []onpremmodel.NodeProperty, nodeGroups []summary.SummaryNodeGroupInfo) ([]WorkerNodeGroupMapping, []WorkerSizingItem) {
	mappings := make([]WorkerNodeGroupMapping, len(nodeGroups))
	for i, ng := range nodeGroups {
		mappings[i] = WorkerNodeGroupMapping{
			NodeGroupName:   ng.Name,
			TargetSpecId:    ng.SpecId,
			TargetNodeCount: ng.DesiredNodeSize,
			AssignedWorkers: make([]SourceServerBrief, 0),
		}
	}

	sizingItems := make([]WorkerSizingItem, 0, len(workers))
	if len(nodeGroups) == 0 {
		return mappings, sizingItems
	}

	// Map each worker to the closest matching node group
	for _, w := range workers {
		srcVcpu := getSourceNodeVcpu(w)
		srcMem := float64(w.Memory.TotalSize)

		bestNgIdx := selectBestNodeGroupIndex(srcVcpu, srcMem, nodeGroups)
		matchedNg := nodeGroups[bestNgIdx]

		brief := toSourceServerBrief(w)
		mappings[bestNgIdx].AssignedWorkers = append(mappings[bestNgIdx].AssignedWorkers, brief)

		sizingReason := determineSizingReason(srcVcpu, srcMem, matchedNg)
		sizingItems = append(sizingItems, WorkerSizingItem{
			WorkerHostname: w.Hostname,
			SourceVcpu:     srcVcpu,
			SourceMemoryGb: srcMem,
			TargetSpecId:   matchedNg.SpecId,
			TargetVcpu:     matchedNg.Vcpu,
			TargetMemoryGb: matchedNg.MemoryGib,
			SizingReason:   sizingReason,
		})
	}

	return mappings, sizingItems
}

// selectBestNodeGroupIndex finds the best fitting node group index for given worker specs.
func selectBestNodeGroupIndex(srcVcpu int, srcMem float64, nodeGroups []summary.SummaryNodeGroupInfo) int {
	bestIdx := 0
	minDistance := math.MaxFloat64

	for i, ng := range nodeGroups {
		// Distance metric favoring node groups satisfying capacity with minimal surplus
		vcpuDiff := float64(ng.Vcpu - srcVcpu)
		memDiff := ng.MemoryGib - srcMem

		if vcpuDiff < 0 {
			vcpuDiff = math.Abs(vcpuDiff) * 3.0
		}
		if memDiff < 0 {
			memDiff = math.Abs(memDiff) * 3.0
		}

		distance := vcpuDiff*2.0 + memDiff
		if distance < minDistance {
			minDistance = distance
			bestIdx = i
		}
	}

	return bestIdx
}

// determineSizingReason classifies the sizing rationale into one of four standard categories.
func determineSizingReason(srcVcpu int, srcMem float64, ng summary.SummaryNodeGroupInfo) string {
	if srcVcpu < 2 || srcMem < 4.0 {
		return "Upscaled to minimum viable K8s worker floor (2vCPU / 4GiB)"
	}
	if ng.Vcpu > srcVcpu && ng.MemoryGib > srcMem {
		return fmt.Sprintf("Consolidated into uniform node group tier (%s)", ng.SpecId)
	}
	if ng.Vcpu > srcVcpu || ng.MemoryGib > srcMem {
		return fmt.Sprintf("Rounded to closest catalog spec with headroom (%s)", ng.SpecId)
	}
	return fmt.Sprintf("Direct capacity match (%s)", ng.SpecId)
}

// buildK8sResourceComparison computes source vs target resource deltas.
func buildK8sResourceComparison(workers []onpremmodel.NodeProperty, targetSummary *summary.TargetK8sInfraSummary) K8sResourceComparison {
	srcWorkerCount := len(workers)
	srcTotalVcpu := 0
	srcTotalMem := 0.0
	srcTotalDisk := 0

	for _, w := range workers {
		srcTotalVcpu += getSourceNodeVcpu(w)
		srcTotalMem += float64(w.Memory.TotalSize)
		srcTotalDisk += int(w.RootDisk.TotalSize)
		for _, d := range w.DataDisks {
			srcTotalDisk += int(d.TotalSize)
		}
	}

	tgtWorkerCount := targetSummary.Overview.TotalWorkerNodeCount
	tgtTotalVcpu := targetSummary.ComputeOverview.TotalVcpus
	tgtTotalMem := targetSummary.ComputeOverview.TotalMemoryGib
	tgtTotalDisk := targetSummary.ComputeOverview.TotalStorageGB

	return K8sResourceComparison{
		WorkerCountDelta: computeResourceDelta("Worker Count", float64(srcWorkerCount), float64(tgtWorkerCount), "nodes"),
		VcpuDelta:        computeResourceDelta("Total vCPU", float64(srcTotalVcpu), float64(tgtTotalVcpu), "vCPU"),
		MemoryGibDelta:   computeResourceDelta("Total Memory", srcTotalMem, tgtTotalMem, "GiB"),
		StorageGbDelta:   computeResourceDelta("Total Storage", float64(srcTotalDisk), float64(tgtTotalDisk), "GB"),
	}
}

// computeResourceDelta creates a standardized ResourceChange representation.
func computeResourceDelta(resourceType string, srcVal, tgtVal float64, unit string) ResourceChange {
	changeType := "Same"
	if tgtVal > srcVal {
		changeType = "Upgrade"
	} else if tgtVal < srcVal {
		changeType = "Downgrade"
	}

	ratio := 1.0
	if srcVal > 0 {
		ratio = tgtVal / srcVal
	}

	diff := tgtVal - srcVal
	desc := fmt.Sprintf("%+.1f %s (%.1fx %s)", diff, unit, ratio, strings.ToLower(changeType))
	if unit == "nodes" || unit == "vCPU" || unit == "GB" {
		desc = fmt.Sprintf("%+d %s (%.1fx %s)", int(diff), unit, ratio, strings.ToLower(changeType))
	}

	return ResourceChange{
		ResourceType: resourceType,
		SourceValue:  fmt.Sprintf("%.1f %s", srcVal, unit),
		TargetValue:  fmt.Sprintf("%.1f %s", tgtVal, unit),
		ChangeType:   changeType,
		ChangeRatio:  ratio,
		Description:  desc,
	}
}

// buildK8sCostSummary constructs cost breakdown by node group components.
func buildK8sCostSummary(targetSummary *summary.TargetK8sInfraSummary) CostSummary {
	totalMonthly := float64(targetSummary.CostEstimation.TotalCostPerMonth)
	components := make([]ComponentCost, 0, len(targetSummary.NodeGroupList))

	for _, ng := range targetSummary.NodeGroupList {
		pct := 0.0
		if totalMonthly > 0 {
			pct = (ng.MonthlyCostUSD / totalMonthly) * 100.0
		}
		components = append(components, ComponentCost{
			ComponentName:  fmt.Sprintf("NodeGroup: %s (%d nodes)", ng.Name, ng.DesiredNodeSize),
			SpecName:       ng.SpecId,
			MonthlyCost:    ng.MonthlyCostUSD,
			CostPercentage: pct,
		})
	}

	return CostSummary{
		TotalHourlyCost:  float64(targetSummary.CostEstimation.TotalCostPerHour),
		TotalDailyCost:   float64(targetSummary.CostEstimation.TotalCostPerDay),
		TotalMonthlyCost: totalMonthly,
		TotalYearlyCost:  totalMonthly * 12.0,
		CostByComponent:  components,
	}
}

// generateK8sRecommendations creates strategic recommendations for K8s infrastructure.
func generateK8sRecommendations(targetSummary *summary.TargetK8sInfraSummary, comparison K8sResourceComparison) []Recommendation {
	var recs []Recommendation

	// Auto-scaling recommendation
	hasAutoScaling := false
	for _, ng := range targetSummary.NodeGroupList {
		if strings.EqualFold(ng.OnAutoScaling, "true") {
			hasAutoScaling = true
			break
		}
	}
	if !hasAutoScaling {
		recs = append(recs, Recommendation{
			Category:    "Scalability",
			Priority:    "Medium",
			Title:       "Enable Node Group Auto-Scaling",
			Description: "Configure Cluster Autoscaler on worker node groups to dynamically adapt to varying container workloads.",
			ActionItems: []string{
				"Enable AutoScaling on target node groups",
				"Configure min/max bounds based on peak traffic estimates",
				"Set up Horizontal Pod Autoscaler (HPA) for application deployments",
			},
		})
	}

	// Cost optimization recommendation
	totalCostPerMonth := float64(targetSummary.CostEstimation.TotalCostPerMonth)
	if totalCostPerMonth > 200.0 {
		savings := totalCostPerMonth * 0.35
		recs = append(recs, Recommendation{
			Category:    "Cost Optimization",
			Priority:    "High",
			Title:       "Evaluate Compute Savings Plans or Reserved Instances",
			Description: fmt.Sprintf("Worker nodes run continuous container workloads; 1-year commitments can save up to 35%% (~$%.2f/month).", savings),
			ActionItems: []string{
				"Assess baseline worker node capacity requirements",
				"Apply 1-year or 3-year Compute Savings Plans or Reserved Instances",
				"Consider Spot / Preemptible node groups for non-critical batch jobs",
			},
		})
	}

	// Persistent storage and CSI recommendation
	recs = append(recs, Recommendation{
		Category:    "Storage",
		Priority:    "Medium",
		Title:       "Verify CSP Container Storage Interface (CSI) Storage Classes",
		Description: "Ensure appropriate default StorageClass and volume snapshot policies are configured for stateful workloads.",
		ActionItems: []string{
			"Verify CSI driver installation and default StorageClass",
			"Configure volume expansion policies in PersistentVolumeClaims",
			"Establish volume snapshot schedules for backup and disaster recovery",
		},
	})

	// Addon and version lifecycle recommendation
	recs = append(recs, Recommendation{
		Category:    "Operations",
		Priority:    "High",
		Title:       "Plan Managed K8s Version Upgrade Lifecycle",
		Description: fmt.Sprintf("Cluster is running Kubernetes %s; monitor CSP end-of-life dates and plan regular minor version updates.", targetSummary.Overview.K8sVersion),
		ActionItems: []string{
			"Track CSP managed Kubernetes support calendar",
			"Validate Helm charts and API deprecations before upgrading",
			"Test minor version upgrades in staging node groups first",
		},
	})

	return recs
}

// getSourceNodeVcpu calculates total vCPUs for an onpremise node.
func getSourceNodeVcpu(n onpremmodel.NodeProperty) int {
	cpus := int(n.CPU.Cpus)
	if cpus == 0 {
		cpus = 1
	}
	threads := int(n.CPU.Threads)
	if threads == 0 {
		threads = int(n.CPU.Cores)
	}
	if threads == 0 {
		threads = 1
	}
	return cpus * threads
}

// toSourceServerBrief converts an onpremise NodeProperty to SourceServerBrief.
func toSourceServerBrief(n onpremmodel.NodeProperty) SourceServerBrief {
	primaryIp := ""
	for _, iface := range n.Interfaces {
		if len(iface.IPv4CidrBlocks) > 0 {
			primaryIp = strings.Split(iface.IPv4CidrBlocks[0], "/")[0]
			break
		}
	}

	return SourceServerBrief{
		Hostname:      n.Hostname,
		MachineID:     n.MachineId,
		CPUModel:      n.CPU.Model,
		CPUs:          int(n.CPU.Cpus),
		CPUThreads:    int(n.CPU.Threads),
		MemoryGB:      int(n.Memory.TotalSize),
		DiskGB:        int(n.RootDisk.TotalSize),
		DiskType:      n.RootDisk.Type,
		OSName:        fmt.Sprintf("%s %s", n.OS.Name, n.OS.Version),
		PrimaryIP:     primaryIp,
		FirewallRules: len(n.FirewallTable),
	}
}
