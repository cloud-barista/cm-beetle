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

// Package summary provides infrastructure summary markdown generation
package summary

import (
	"fmt"
	"strings"
)

// GenerateMarkdownK8sInfraSummary converts TargetK8sInfraSummary to markdown format.
func GenerateMarkdownK8sInfraSummary(summary *TargetK8sInfraSummary) string {
	var md strings.Builder

	// Title and Metadata
	md.WriteString("# Target Cloud Kubernetes Infrastructure Summary\n\n")
	md.WriteString(fmt.Sprintf("**Generated At:** %s\n\n", summary.SummaryMetadata.GeneratedAt.Format("2006-01-02 15:04:05")))
	md.WriteString(fmt.Sprintf("**Namespace:** %s\n\n", summary.SummaryMetadata.Namespace))
	md.WriteString(fmt.Sprintf("**Cluster Name:** %s\n\n", summary.SummaryMetadata.ClusterName))
	md.WriteString("---\n\n")

	// Overview Section
	md.WriteString("## Overview\n\n")
	md.WriteString(generateK8sOverviewMarkdown(&summary.Overview))
	md.WriteString("\n")

	// Node Groups & Compute Resources Section
	md.WriteString("## Kubernetes Node Groups & Compute Resources\n\n")
	md.WriteString(generateK8sNodeGroupsMarkdown(summary.NodeGroupList, &summary.ComputeOverview))
	md.WriteString("\n")

	// Network Resources Section (reuses shared formatter)
	md.WriteString("## Network Resources\n\n")
	md.WriteString(generateNetworkResourcesMarkdown(&summary.NetworkResources))
	md.WriteString("\n")

	// Security Resources Section (reuses shared formatter)
	md.WriteString("## Security Resources\n\n")
	md.WriteString(generateSecurityResourcesMarkdown(&summary.SecurityResources))
	md.WriteString("\n")

	// Cost Estimation Section (reuses shared formatter)
	md.WriteString("## Cost Estimation\n\n")
	md.WriteString(generateCostEstimationMarkdown(&summary.CostEstimation))
	md.WriteString("\n")

	return md.String()
}

// generateK8sOverviewMarkdown generates markdown table for K8s cluster overview.
func generateK8sOverviewMarkdown(overview *TargetK8sInfraOverview) string {
	var md strings.Builder

	md.WriteString("| Property | Value |\n")
	md.WriteString("|----------|-------|\n")
	md.WriteString(fmt.Sprintf("| **Cluster Name** | %s |\n", overview.ClusterName))
	md.WriteString(fmt.Sprintf("| **Status** | %s |\n", overview.Status))
	md.WriteString(fmt.Sprintf("| **Target Cloud** | %s |\n", overview.TargetCloud))
	md.WriteString(fmt.Sprintf("| **Target Region** | %s |\n", overview.TargetRegion))
	md.WriteString(fmt.Sprintf("| **Connection Name** | %s |\n", overview.ConnectionName))
	md.WriteString(fmt.Sprintf("| **Kubernetes Version** | %s |\n", overview.K8sVersion))
	if overview.CspClusterId != "" {
		md.WriteString(fmt.Sprintf("| **CSP Cluster ID** | %s |\n", overview.CspClusterId))
	}
	md.WriteString(fmt.Sprintf("| **Node Groups Count** | %d |\n", overview.TotalNodeGroupCount))
	md.WriteString(fmt.Sprintf("| **Total Desired Workers** | %d |\n", overview.TotalWorkerNodeCount))
	md.WriteString(fmt.Sprintf("| **Active Nodes** | %d |\n", overview.ActiveNodeCount))

	return md.String()
}

// generateK8sNodeGroupsMarkdown generates markdown for compute overview and node groups.
func generateK8sNodeGroupsMarkdown(nodeGroups []SummaryNodeGroupInfo, compute *K8sComputeOverview) string {
	var md strings.Builder

	// Total Compute Aggregates
	md.WriteString("### Compute Totals\n\n")
	md.WriteString("| Metric | Value |\n")
	md.WriteString("|--------|-------|\n")
	md.WriteString(fmt.Sprintf("| **Total vCPUs** | %d cores |\n", compute.TotalVcpus))
	md.WriteString(fmt.Sprintf("| **Total Memory** | %.1f GiB |\n", compute.TotalMemoryGib))
	md.WriteString(fmt.Sprintf("| **Total Storage** | %d GB |\n", compute.TotalStorageGB))
	md.WriteString("\n")

	// Node Groups Table
	md.WriteString("### Node Groups\n\n")
	if len(nodeGroups) == 0 {
		md.WriteString("*No node groups found.*\n\n")
		return md.String()
	}

	md.WriteString("| Node Group | Spec ID | Desired (Min/Max) | vCPU / RAM | Root Disk | Autoscaling | Status | Cost/Month (USD) |\n")
	md.WriteString("|------------|---------|-------------------|------------|-----------|-------------|--------|------------------|\n")
	for _, ng := range nodeGroups {
		scalingRange := fmt.Sprintf("%d (%d/%d)", ng.DesiredNodeSize, ng.MinNodeSize, ng.MaxNodeSize)
		specDetail := fmt.Sprintf("%d vCPU, %.1f GiB", ng.Vcpu, ng.MemoryGib)
		diskDetail := fmt.Sprintf("%d GB (%s)", ng.RootDiskSizeGB, ng.RootDiskType)
		md.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | $%.2f |\n",
			ng.Name, ng.SpecId, scalingRange, specDetail, diskDetail, ng.OnAutoScaling, ng.Status, ng.MonthlyCostUSD))
	}
	md.WriteString("\n")

	// Worker Nodes Breakdown
	md.WriteString("### Worker Node Instances\n\n")
	for _, ng := range nodeGroups {
		md.WriteString(fmt.Sprintf("#### Node Group: `%s`\n\n", ng.Name))
		if len(ng.Nodes) == 0 {
			md.WriteString("*No provisioned nodes reported for this node group.*\n\n")
			continue
		}

		md.WriteString("| CSP Resource Name | CSP Resource ID |\n")
		md.WriteString("|-------------------|-----------------|\n")
		for _, n := range ng.Nodes {
			md.WriteString(fmt.Sprintf("| %s | %s |\n", n.CspResourceName, n.CspResourceId))
		}
		md.WriteString("\n")
	}

	return md.String()
}
