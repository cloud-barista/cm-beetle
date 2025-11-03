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

// GenerateMarkdownSummary converts TargetInfraSummary to markdown format
func GenerateMarkdownSummary(summary *TargetInfraSummary) string {
	var md strings.Builder

	// Title and Metadata
	md.WriteString("# Infrastructure Summary\n\n")
	md.WriteString(fmt.Sprintf("**Generated At:** %s\n\n", summary.SummaryMetadata.GeneratedAt.Format("2006-01-02 15:04:05")))
	md.WriteString(fmt.Sprintf("**Namespace:** %s\n\n", summary.SummaryMetadata.Namespace))
	md.WriteString(fmt.Sprintf("**MCI Name:** %s\n\n", summary.SummaryMetadata.MciName))
	md.WriteString(fmt.Sprintf("**Summary Version:** %s\n\n", summary.SummaryMetadata.SummaryVersion))
	md.WriteString("---\n\n")

	// Overview Section
	md.WriteString("## Overview\n\n")
	md.WriteString(generateOverviewMarkdown(&summary.Overview))
	md.WriteString("\n")

	// Compute Resources Section
	md.WriteString("## Compute Resources\n\n")
	md.WriteString(generateComputeResourcesMarkdown(&summary.ComputeResources))
	md.WriteString("\n")

	// Network Resources Section
	md.WriteString("## Network Resources\n\n")
	md.WriteString(generateNetworkResourcesMarkdown(&summary.NetworkResources))
	md.WriteString("\n")

	// Security Resources Section
	md.WriteString("## Security Resources\n\n")
	md.WriteString(generateSecurityResourcesMarkdown(&summary.SecurityResources))
	md.WriteString("\n")

	// Cost Estimation Section
	md.WriteString("## Cost Estimation\n\n")
	md.WriteString(generateCostEstimationMarkdown(&summary.CostEstimation))
	md.WriteString("\n")

	return md.String()
}

// generateOverviewMarkdown generates markdown for infrastructure overview
func generateOverviewMarkdown(overview *TargetInfraOverview) string {
	var md strings.Builder

	md.WriteString("| Property | Value |\n")
	md.WriteString("|----------|-------|\n")
	md.WriteString(fmt.Sprintf("| **MCI Name** | %s |\n", overview.MciName))
	md.WriteString(fmt.Sprintf("| **Description** | %s |\n", overview.MciDescription))
	md.WriteString(fmt.Sprintf("| **Status** | %s |\n", overview.Status))
	md.WriteString(fmt.Sprintf("| **Target Cloud** | %s |\n", overview.TargetCloud))
	md.WriteString(fmt.Sprintf("| **Target Region** | %s |\n", overview.TargetRegion))
	md.WriteString(fmt.Sprintf("| **Total VMs** | %d |\n", overview.TotalVmCount))
	md.WriteString(fmt.Sprintf("| **Running VMs** | %d |\n", overview.RunningVmCount))
	md.WriteString(fmt.Sprintf("| **Stopped VMs** | %d |\n", overview.StoppedVmCount))
	if len(overview.Label) > 0 {
		labelStr := ""
		for k, v := range overview.Label {
			if labelStr != "" {
				labelStr += ", "
			}
			labelStr += fmt.Sprintf("%s=%s", k, v)
		}
		md.WriteString(fmt.Sprintf("| **Label** | %s |\n", labelStr))
	}
	md.WriteString(fmt.Sprintf("| **Monitoring Agent** | %s |\n", overview.InstallMonAgent))

	return md.String()
}

// generateNetworkResourcesMarkdown generates markdown for network resources
func generateNetworkResourcesMarkdown(resources *SummaryNetworkResources) string {
	var md strings.Builder

	md.WriteString("### Virtual Networks\n\n")

	if len(resources.VNets) == 0 {
		md.WriteString("*No virtual networks found.*\n\n")
		return md.String()
	}

	for _, vnet := range resources.VNets {
		md.WriteString(fmt.Sprintf("#### VNet: %s\n\n", vnet.Name))
		md.WriteString("| Property | Value |\n")
		md.WriteString("|----------|-------|\n")
		md.WriteString(fmt.Sprintf("| **Name** | %s |\n", vnet.Name))
		md.WriteString(fmt.Sprintf("| **CSP VNet ID** | %s |\n", vnet.CspVNetId))
		md.WriteString(fmt.Sprintf("| **CIDR Block** | %s |\n", vnet.CidrBlock))
		md.WriteString(fmt.Sprintf("| **Connection** | %s |\n", vnet.ConnectionName))
		md.WriteString(fmt.Sprintf("| **Subnet Count** | %d |\n", vnet.SubnetCount))
		md.WriteString("\n")

		if len(vnet.Subnets) > 0 {
			md.WriteString("**Subnets:**\n\n")
			md.WriteString("| Name | CSP Subnet ID | CIDR Block | Zone |\n")
			md.WriteString("|------|---------------|------------|------|\n")
			for _, subnet := range vnet.Subnets {
				md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
					subnet.Name, subnet.CspSubnetId, subnet.CidrBlock, subnet.Zone))
			}
			md.WriteString("\n")
		}
	}

	return md.String()
}

// generateSecurityResourcesMarkdown generates markdown for security resources
func generateSecurityResourcesMarkdown(resources *SummarySecurityResources) string {
	var md strings.Builder

	// SSH Keys Section
	md.WriteString("### SSH Keys\n\n")
	if len(resources.SshKeys) == 0 {
		md.WriteString("*No SSH keys found.*\n\n")
	} else {
		md.WriteString("| Name | CSP SSH Key ID | Username | Fingerprint |\n")
		md.WriteString("|------|----------------|----------|-------------|\n")
		for _, sshKey := range resources.SshKeys {
			md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
				sshKey.Name, sshKey.CspSshKeyId, sshKey.Username, sshKey.Fingerprint))
		}
		md.WriteString("\n")
	}

	// Security Groups Section
	md.WriteString("### Security Groups\n\n")
	if len(resources.SecurityGroups) == 0 {
		md.WriteString("*No security groups found.*\n\n")
		return md.String()
	}

	for _, sg := range resources.SecurityGroups {
		md.WriteString(fmt.Sprintf("#### Security Group: %s\n\n", sg.Name))
		md.WriteString("| Property | Value |\n")
		md.WriteString("|----------|-------|\n")
		md.WriteString(fmt.Sprintf("| **Name** | %s |\n", sg.Name))
		md.WriteString(fmt.Sprintf("| **CSP Security Group ID** | %s |\n", sg.CspSecurityGroupId))
		md.WriteString(fmt.Sprintf("| **VNet** | %s |\n", sg.VNetName))
		md.WriteString(fmt.Sprintf("| **Rule Count** | %d |\n", sg.RuleCount))
		md.WriteString("\n")

		if len(sg.Rules) > 0 {
			md.WriteString("**Security Group Rules:**\n\n")
			md.WriteString("| Direction | Protocol | Port | CIDR |\n")
			md.WriteString("|-----------|----------|------|------|\n")
			for _, rule := range sg.Rules {
				port := rule.FromPort
				if rule.FromPort != rule.ToPort {
					port = fmt.Sprintf("%s-%s", rule.FromPort, rule.ToPort)
				}
				md.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
					rule.Direction, rule.Protocol, port, rule.Cidr))
			}
			md.WriteString("\n")
		}
	}

	return md.String()
}

// generateComputeResourcesMarkdown generates markdown for compute resources
func generateComputeResourcesMarkdown(resources *SummaryComputeResources) string {
	var md strings.Builder

	// VM Specifications Section
	md.WriteString("### VM Specifications\n\n")
	if len(resources.Specs) == 0 {
		md.WriteString("*No VM specifications found.*\n\n")
	} else {
		md.WriteString("| Name | vCPUs | Memory (GiB) | GPU | Architecture | Disk Type | Cost/Hour (USD) | Usage Count |\n")
		md.WriteString("|------|-------|--------------|-----|--------------|-----------|-----------------|-------------|\n")
		for _, spec := range resources.Specs {
			gpuInfo := "-"
			if spec.AcceleratorCount > 0 {
				gpuInfo = fmt.Sprintf("%d x %s (%.2f GB)", spec.AcceleratorCount, spec.AcceleratorModel, spec.AcceleratorMemoryGB)
			}
			md.WriteString(fmt.Sprintf("| %s | %d | %.1f | %s | %s | %s | $%.4f | %d |\n",
				spec.CspSpecName, int(spec.VCPU), spec.MemoryGiB, gpuInfo, spec.Architecture, spec.RootDiskType, spec.CostPerHour, spec.UsageCount))
		}
		md.WriteString("\n")
	}

	// VM Images Section
	md.WriteString("### VM Images\n\n")
	if len(resources.Images) == 0 {
		md.WriteString("*No VM images found.*\n\n")
	} else {
		md.WriteString("| Name | Distribution | OS Type | OS Platform | Architecture | Disk Type | Disk Size | Usage Count |\n")
		md.WriteString("|------|--------------|---------|-------------|--------------|-----------|-----------|-------------|\n")
		for _, image := range resources.Images {
			diskSize := fmt.Sprintf("%.0f GB", image.OSDiskSizeGB)
			md.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %d |\n",
				image.CspImageName, image.OSDistribution, image.OSType, string(image.OSPlatform),
				string(image.OSArchitecture), image.OSDiskType, diskSize, image.UsageCount))
		}
		md.WriteString("\n")
	}

	// Virtual Machines Section
	md.WriteString("### Virtual Machines\n\n")
	if len(resources.Vms) == 0 {
		md.WriteString("*No virtual machines found.*\n\n")
		return md.String()
	}

	md.WriteString("| Instance Name | CSP Instance ID | Status | Spec (vCPU, Memory GiB) | Image | Misc |\n")
	md.WriteString("|---------------|-----------------|--------|-------------------------|-------|------|\n")
	for _, vm := range resources.Vms {
		specInfo := fmt.Sprintf("%d vCPU, %.1f GiB", vm.Spec.VCpus, vm.Spec.MemoryGiB)
		imageInfo := fmt.Sprintf("%s (%s)", vm.Image.Distribution, vm.Image.OsVersion)
		miscInfo := fmt.Sprintf("VNet: %s, Subnet: %s, Public IP: %s, Private IP: %s, SGs: %s, SSH: %s",
			vm.Misc.VNet, vm.Misc.Subnet, vm.Misc.PublicIp, vm.Misc.PrivateIp,
			strings.Join(vm.Misc.SecurityGroups, ", "), vm.Misc.SshKey)
		md.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
			vm.Name, vm.CspVmId, vm.Status, specInfo, imageInfo, miscInfo))
	}
	md.WriteString("\n")

	return md.String()
}

// generateCostEstimationMarkdown generates markdown for cost estimation
func generateCostEstimationMarkdown(cost *SummaryCostEstimation) string {
	var md strings.Builder

	md.WriteString("### Total Cost Summary\n\n")
	md.WriteString("| Period | Cost (USD) |\n")
	md.WriteString("|--------|------------|\n")
	md.WriteString(fmt.Sprintf("| **Per Hour** | $%.4f |\n", cost.TotalCostPerHour))
	md.WriteString(fmt.Sprintf("| **Per Day** | $%.2f |\n", cost.TotalCostPerDay))
	md.WriteString(fmt.Sprintf("| **Per Month (30 days)** | $%.2f |\n", cost.TotalCostPerMonth))
	md.WriteString("\n")

	// Cost by Region
	if len(cost.ByRegion) > 0 {
		md.WriteString("### Cost by Region\n\n")
		md.WriteString("| CSP | Region | VM Count | Cost/Hour (USD) | Cost/Month (USD) |\n")
		md.WriteString("|-----|--------|----------|-----------------|------------------|\n")
		for _, region := range cost.ByRegion {
			md.WriteString(fmt.Sprintf("| %s | %s | %d | $%.4f | $%.2f |\n",
				region.Csp, region.Region, region.VmCount, region.CostPerHour, region.CostPerMonth))
		}
		md.WriteString("\n")
	}

	// Cost by VM
	if len(cost.ByVm) > 0 {
		md.WriteString("### Cost by Virtual Machine\n\n")
		md.WriteString("| VM Name | Spec | Cost/Hour (USD) | Cost/Month (USD) |\n")
		md.WriteString("|---------|------|-----------------|------------------|\n")
		for _, vm := range cost.ByVm {
			md.WriteString(fmt.Sprintf("| %s | %s | $%.4f | $%.2f |\n",
				vm.VmName, vm.SpecName, vm.CostPerHour, vm.CostPerMonth))
		}
		md.WriteString("\n")
	}

	return md.String()
}
