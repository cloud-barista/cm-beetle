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

// Package summary provides target K8s infrastructure summary generation logic
package summary

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// GenerateTargetK8sInfraSummary generates a comprehensive K8s infrastructure summary.
func GenerateTargetK8sInfraSummary(nsId, clusterId string) (*TargetK8sInfraSummary, error) {
	log.Info().Msgf("Generating K8s infrastructure summary (nsId: %s, clusterId: %s)", nsId, clusterId)

	// Step 1: Collect K8s cluster information
	clusterInfo, err := tbclient.NewSession().ReadK8sCluster(nsId, clusterId)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to retrieve K8s cluster: %s", clusterId)
		return nil, fmt.Errorf("failed to retrieve K8s cluster: %w", err)
	}

	// Step 2: Collect network resources
	var vnetIds []string
	if clusterInfo.Network.VNetId != "" {
		vnetIds = append(vnetIds, clusterInfo.Network.VNetId)
	}
	networkResources, err := collectNetworkResources(nsId, vnetIds)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to collect network resources for K8s cluster")
	}

	// Step 3: Collect security resources
	uniqueSshKeyIds := extractUniqueK8sSshKeyIds(clusterInfo.K8sNodeGroupList)
	securityResources, err := collectSecurityResources(nsId, uniqueSshKeyIds, clusterInfo.Network.SecurityGroupIds)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to collect security resources for K8s cluster")
	}

	// Step 4: Collect node groups and compute overview
	nodeGroupList, computeOverview, totalHourlyCost := collectK8sNodeGroupResources(&clusterInfo)

	// Step 5: Build cost estimation
	costEstimation := buildK8sCostEstimation(&clusterInfo, totalHourlyCost, nodeGroupList)

	// Step 6: Build summary metadata
	metadata := TargetK8sSummaryMetadata{
		GeneratedAt:    time.Now(),
		Namespace:      nsId,
		ClusterId:      clusterId,
		ClusterName:    clusterInfo.Name,
		SummaryVersion: "1.0",
	}

	// Step 7: Build infrastructure overview
	overview := buildTargetK8sInfraOverview(&clusterInfo, nodeGroupList)

	// Step 8: Assemble final summary
	k8sSummary := &TargetK8sInfraSummary{
		SummaryMetadata:   metadata,
		Overview:          overview,
		NetworkResources:  networkResources,
		SecurityResources: securityResources,
		NodeGroupList:     nodeGroupList,
		ComputeOverview:   computeOverview,
		CostEstimation:    costEstimation,
	}

	log.Info().Msgf("Successfully generated K8s infrastructure summary for cluster: %s", clusterId)
	return k8sSummary, nil
}

// collectK8sNodeGroupResources gathers node group specs, nodes, compute metrics, and costs.
func collectK8sNodeGroupResources(clusterInfo *tbmodel.K8sClusterInfo) ([]SummaryNodeGroupInfo, K8sComputeOverview, float64) {
	var ngList []SummaryNodeGroupInfo
	var compute K8sComputeOverview
	var totalHourlyCost float64

	for _, ng := range clusterInfo.K8sNodeGroupList {
		var vcpu int
		var memGib float64
		var hourlyCostPerNode float64

		if ng.SpecId != "" {
			specInfo, err := tbclient.NewSession().ReadVmSpec("system", ng.SpecId)
			if err == nil {
				vcpu = int(specInfo.VCPU)
				memGib = float64(specInfo.MemoryGiB)
				hourlyCostPerNode = float64(specInfo.CostPerHour)
			} else {
				log.Warn().Err(err).Str("specId", ng.SpecId).Msg("Failed to read spec info for node group")
			}
		}

		nodeCount := ng.DesiredNodeSize
		if nodeCount <= 0 {
			nodeCount = len(ng.K8sNodes)
		}

		ngHourlyCost := hourlyCostPerNode * float64(nodeCount)
		totalHourlyCost += ngHourlyCost

		compute.TotalVcpus += vcpu * nodeCount
		compute.TotalMemoryGib += memGib * float64(nodeCount)
		compute.TotalStorageGB += ng.RootDiskSize * nodeCount

		var nodes []SummaryK8sNodeInfo
		for _, node := range ng.K8sNodes {
			nodes = append(nodes, SummaryK8sNodeInfo{
				CspResourceName: node.CspResourceName,
				CspResourceId:   node.CspResourceId,
			})
		}

		ngList = append(ngList, SummaryNodeGroupInfo{
			Name:            ng.Name,
			Status:          string(ng.Status),
			SpecId:          ng.SpecId,
			Vcpu:            vcpu,
			MemoryGib:       memGib,
			RootDiskType:    ng.RootDiskType,
			RootDiskSizeGB:  ng.RootDiskSize,
			OnAutoScaling:   strconv.FormatBool(ng.OnAutoScaling),
			DesiredNodeSize: ng.DesiredNodeSize,
			MinNodeSize:     ng.MinNodeSize,
			MaxNodeSize:     ng.MaxNodeSize,
			Nodes:           nodes,
			HourlyCostUSD:   ngHourlyCost,
			MonthlyCostUSD:  ngHourlyCost * 24 * 30,
		})
	}

	return ngList, compute, totalHourlyCost
}

// buildTargetK8sInfraOverview constructs overview information from K8s cluster details.
func buildTargetK8sInfraOverview(clusterInfo *tbmodel.K8sClusterInfo, nodeGroups []SummaryNodeGroupInfo) TargetK8sInfraOverview {
	targetCloud := "Unknown"
	targetRegion := "Unknown"

	if clusterInfo.ConnectionConfig.ProviderName != "" {
		targetCloud = strings.ToUpper(clusterInfo.ConnectionConfig.ProviderName)
	}
	if clusterInfo.ConnectionConfig.RegionDetail.RegionName != "" {
		targetRegion = clusterInfo.ConnectionConfig.RegionDetail.RegionName
	}
	if targetCloud == "Unknown" && clusterInfo.ConnectionName != "" {
		parts := strings.Split(clusterInfo.ConnectionName, "-")
		if len(parts) >= 2 {
			targetCloud = strings.ToUpper(parts[0])
			targetRegion = strings.Join(parts[1:], "-")
		}
	}

	totalWorkers := 0
	activeNodes := 0
	for _, ng := range nodeGroups {
		totalWorkers += ng.DesiredNodeSize
		if strings.EqualFold(ng.Status, "active") || strings.EqualFold(ng.Status, "ready") {
			if len(ng.Nodes) > 0 {
				activeNodes += len(ng.Nodes)
			} else {
				activeNodes += ng.DesiredNodeSize
			}
		}
	}

	return TargetK8sInfraOverview{
		ClusterName:          clusterInfo.Name,
		Status:               string(clusterInfo.Status),
		TargetCloud:          targetCloud,
		TargetRegion:         targetRegion,
		ConnectionName:       clusterInfo.ConnectionName,
		K8sVersion:           clusterInfo.Version,
		CspClusterId:         clusterInfo.CspResourceId,
		TotalNodeGroupCount:  len(clusterInfo.K8sNodeGroupList),
		TotalWorkerNodeCount: totalWorkers,
		ActiveNodeCount:      activeNodes,
	}
}

// buildK8sCostEstimation calculates aggregated cost estimation for the K8s cluster.
func buildK8sCostEstimation(clusterInfo *tbmodel.K8sClusterInfo, totalHourlyCost float64, nodeGroups []SummaryNodeGroupInfo) SummaryCostEstimation {
	csp := "Unknown"
	region := "Unknown"
	if clusterInfo.ConnectionName != "" {
		parts := strings.Split(clusterInfo.ConnectionName, "-")
		if len(parts) >= 2 {
			csp = strings.ToUpper(parts[0])
			region = strings.Join(parts[1:], "-")
		}
	}

	var byVmList []SummaryCostByVm
	for _, ng := range nodeGroups {
		byVmList = append(byVmList, SummaryCostByVm{
			VmName:       fmt.Sprintf("nodegroup:%s (%d nodes)", ng.Name, ng.DesiredNodeSize),
			SpecName:     ng.SpecId,
			CostPerHour:  float32(ng.HourlyCostUSD),
			CostPerMonth: float32(ng.MonthlyCostUSD),
		})
	}

	byRegion := []SummaryCostByRegion{
		{
			Csp:          csp,
			Region:       region,
			VmCount:      len(nodeGroups),
			CostPerHour:  float32(totalHourlyCost),
			CostPerMonth: float32(totalHourlyCost * 24 * 30),
		},
	}

	return SummaryCostEstimation{
		Currency:          "USD",
		TotalCostPerHour:  float32(totalHourlyCost),
		TotalCostPerDay:   float32(totalHourlyCost * 24),
		TotalCostPerMonth: float32(totalHourlyCost * 24 * 30),
		ByRegion:          byRegion,
		ByVm:              byVmList,
	}
}

// extractUniqueK8sSshKeyIds extracts unique SSH Key IDs from node groups.
func extractUniqueK8sSshKeyIds(nodeGroups []tbmodel.K8sNodeGroupInfo) []string {
	idMap := make(map[string]struct{})
	for _, ng := range nodeGroups {
		if ng.SshKeyId != "" {
			idMap[ng.SshKeyId] = struct{}{}
		}
	}

	ids := make([]string, 0, len(idMap))
	for id := range idMap {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
