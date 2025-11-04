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

// Package report provides migration report generation logic
package report

import (
	"fmt"
	"strings"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
	"github.com/rs/zerolog/log"
)

// GenerateMigrationReport generates a comprehensive migration report
func GenerateMigrationReport(nsId, mciId string, sourceInfra onpremmodel.OnpremInfra) (*MigrationReport, error) {
	log.Info().Msgf("Generating migration report (nsId: %s, mciId: %s)", nsId, mciId)

	// Step 1: Generate source infrastructure summary
	sourceInfraName := fmt.Sprintf("infra-%d-servers", len(sourceInfra.Servers))
	sourceSummary, err := summary.GenerateSourceInfraSummary(sourceInfraName, sourceInfra)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate source infrastructure summary")
		return nil, fmt.Errorf("failed to generate source infrastructure summary: %w", err)
	}

	// Step 2: Generate target infrastructure summary
	targetSummary, err := summary.GenerateInfraSummary(nsId, mciId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate target infrastructure summary")
		return nil, fmt.Errorf("failed to generate target infrastructure summary: %w", err)
	}

	// Step 3: Build report metadata
	metadata := ReportMetadata{
		GeneratedAt:   time.Now(),
		MigrationID:   fmt.Sprintf("%s/%s", nsId, mciId),
		Namespace:     nsId,
		MciID:         mciId,
		ReportVersion: "1.0",
	}

	// Step 4: Build executive summary
	executiveSummary := buildExecutiveSummary(sourceSummary, targetSummary)

	// Step 5: Build migration mappings (match source servers to target VMs)
	migrationMappings := buildMigrationMappings(sourceSummary, targetSummary)

	// Step 6: Build network analysis
	networkAnalysis := buildNetworkAnalysis(sourceSummary, targetSummary)

	// Step 7: Build security analysis
	securityAnalysis := buildSecurityAnalysis(sourceSummary, targetSummary, migrationMappings)

	// Step 8: Build cost summary
	costSummary := buildCostSummary(targetSummary, migrationMappings)

	// Step 9: Generate recommendations
	recommendations := generateRecommendations(sourceSummary, targetSummary, migrationMappings)

	// Step 10: Assemble final report
	report := &MigrationReport{
		Metadata:          metadata,
		ExecutiveSummary:  executiveSummary,
		MigrationMappings: migrationMappings,
		NetworkAnalysis:   networkAnalysis,
		SecurityAnalysis:  securityAnalysis,
		CostSummary:       costSummary,
		Recommendations:   recommendations,
		SourceDetails:     sourceSummary,
		TargetDetails:     targetSummary,
	}

	log.Info().Msgf("Successfully generated migration report (nsId: %s, mciId: %s)", nsId, mciId)
	return report, nil
}

// buildExecutiveSummary builds the executive summary section
func buildExecutiveSummary(sourceSummary *summary.SourceInfraSummary, targetSummary *summary.TargetInfraSummary) ExecutiveSummary {
	totalServers := sourceSummary.Overview.TotalServerCount
	migratedServers := targetSummary.Overview.TotalVmCount
	failedServers := totalServers - migratedServers
	if failedServers < 0 {
		failedServers = 0
	}

	migrationStatus := "Completed"
	if failedServers > 0 {
		migrationStatus = "Partially Completed"
	}
	if migratedServers == 0 {
		migrationStatus = "Failed"
	}

	return ExecutiveSummary{
		MigrationStatus: migrationStatus,
		TotalServers:    totalServers,
		MigratedServers: migratedServers,
		FailedServers:   failedServers,
		TargetCloud:     targetSummary.Overview.TargetCloud,
		TargetRegion:    targetSummary.Overview.TargetRegion,
		MonthlyCostUSD:  float64(targetSummary.CostEstimation.TotalCostPerMonth),
	}
}

// buildMigrationMappings builds source-to-target mappings
func buildMigrationMappings(sourceSummary *summary.SourceInfraSummary, targetSummary *summary.TargetInfraSummary) []SourceTargetMapping {
	var mappings []SourceTargetMapping

	// Create a map of machine IDs from target VMs
	vmByMachineID := make(map[string]summary.SummaryVmInfo)
	for _, vm := range targetSummary.ComputeResources.Vms {
		// Extract sourceMachineId from VM name or labels
		machineID := extractSourceMachineID(vm.Name)
		if machineID != "" {
			vmByMachineID[machineID] = vm
		}
	}

	// Match source servers to target VMs
	mappingID := 1
	for _, sourceServer := range sourceSummary.ComputeResources.Servers {
		// Try to find matching VM by machine ID (now directly from source server data)
		machineID := sourceServer.MachineId
		targetVM, found := vmByMachineID[machineID]

		if !found {
			// Try alternative matching strategies
			log.Warn().Msgf("No matching target VM found for source server: %s (MachineId: %s)", sourceServer.Hostname, machineID)
			continue
		}

		// Build source server brief
		sourceFirewallRules := 0
		for _, serverFirewall := range sourceSummary.SecurityResources.ServerFirewalls {
			if serverFirewall.Hostname == sourceServer.Hostname {
				sourceFirewallRules = len(serverFirewall.FirewallRules)
				break
			}
		}

		sourceBrief := SourceServerBrief{
			Hostname:      sourceServer.Hostname,
			MachineID:     machineID,
			CPUModel:      sourceServer.CPU.Model,
			CPUs:          sourceServer.CPU.CPUs,
			CPUThreads:    sourceServer.CPU.Threads,
			MemoryGB:      sourceServer.Memory.TotalGB,
			DiskGB:        sourceServer.Disk.TotalGB,
			DiskType:      sourceServer.Disk.Type,
			OSName:        fmt.Sprintf("%s %s", sourceServer.OS.Name, sourceServer.OS.Version),
			PrimaryIP:     sourceServer.Network.IPAddress,
			FirewallRules: sourceFirewallRules,
		}

		// Build target VM brief
		targetBrief := TargetVMBrief{
			InstanceName:    targetVM.Name,
			InstanceID:      targetVM.CspVmId,
			SourceMachineID: machineID,
			Status:          targetVM.Status,
			SpecName:        targetVM.Spec.Name,
			VCPUs:           targetVM.Spec.VCpus,
			MemoryGB:        targetVM.Spec.MemoryGiB,
			RootDiskGB:      50,    // Default, would need to get from VM details
			RootDiskType:    "gp2", // Default, would need to get from VM details
			PublicIP:        targetVM.Misc.PublicIp,
			PrivateIP:       targetVM.Misc.PrivateIp,
			SecurityGroups:  targetVM.Misc.SecurityGroups,
		}

		// Analyze resource changes
		resourceChanges := analyzeResourceChanges(sourceServer, targetVM)

		// Calculate cost for this VM
		costPerMonth := 0.0
		for _, vmCost := range targetSummary.CostEstimation.ByVm {
			if vmCost.VmName == targetVM.Name {
				costPerMonth = float64(vmCost.CostPerMonth)
				break
			}
		}

		mapping := SourceTargetMapping{
			MappingID:       mappingID,
			SourceServer:    sourceBrief,
			TargetVM:        targetBrief,
			ResourceChanges: resourceChanges,
			MigrationStatus: "Success",
			CostPerMonth:    costPerMonth,
		}

		mappings = append(mappings, mapping)
		mappingID++
	}

	return mappings
}

// analyzeResourceChanges analyzes changes between source and target resources
func analyzeResourceChanges(sourceServer summary.SourceServerInfo, targetVM summary.SummaryVmInfo) ResourceChangeAnalysis {
	// CPU Change
	cpuChange := ResourceChange{
		ResourceType: "CPU",
		SourceValue:  fmt.Sprintf("%d cores", sourceServer.CPU.Cores),
		TargetValue:  fmt.Sprintf("%d vCPU", targetVM.Spec.VCpus),
		ChangeType:   determineChangeType(float64(sourceServer.CPU.Cores), float64(targetVM.Spec.VCpus)),
		ChangeRatio:  float64(targetVM.Spec.VCpus) / float64(sourceServer.CPU.Cores),
		Description:  fmt.Sprintf("%+d cores (%.1fx)", targetVM.Spec.VCpus-sourceServer.CPU.Cores, float64(targetVM.Spec.VCpus)/float64(sourceServer.CPU.Cores)),
	}

	// Memory Change
	memoryChange := ResourceChange{
		ResourceType: "Memory",
		SourceValue:  fmt.Sprintf("%d GB", sourceServer.Memory.TotalGB),
		TargetValue:  fmt.Sprintf("%.1f GB", targetVM.Spec.MemoryGiB),
		ChangeType:   determineChangeType(float64(sourceServer.Memory.TotalGB), float64(targetVM.Spec.MemoryGiB)),
		ChangeRatio:  float64(targetVM.Spec.MemoryGiB) / float64(sourceServer.Memory.TotalGB),
		Description:  "Same",
	}
	if sourceServer.Memory.TotalGB != int(targetVM.Spec.MemoryGiB) {
		memoryChange.Description = fmt.Sprintf("%+.1f GB", float64(targetVM.Spec.MemoryGiB)-float64(sourceServer.Memory.TotalGB))
	}

	// Storage Change
	rootDiskGB := 50 // Default assumption
	storageChange := ResourceChange{
		ResourceType: "Storage",
		SourceValue:  fmt.Sprintf("%d GB %s", sourceServer.Disk.TotalGB, sourceServer.Disk.Type),
		TargetValue:  fmt.Sprintf("%d GB (root disk)", rootDiskGB),
		ChangeType:   determineChangeType(float64(sourceServer.Disk.TotalGB), float64(rootDiskGB)),
		ChangeRatio:  float64(rootDiskGB) / float64(sourceServer.Disk.TotalGB),
		Description:  fmt.Sprintf("%+d GB (root disk only)", rootDiskGB-sourceServer.Disk.TotalGB),
	}

	// Network Change
	networkChange := NetworkChange{
		SourceIP:        sourceServer.Network.IPAddress,
		TargetPrivateIP: targetVM.Misc.PrivateIp,
		TargetPublicIP:  targetVM.Misc.PublicIp,
		ChangeType:      "Public IP Added",
		Description:     "Added public IP for internet access",
	}
	if targetVM.Misc.PublicIp == "" {
		networkChange.ChangeType = "Private Only"
		networkChange.Description = "Private IP only, no public IP assigned"
	}

	// Security Change (placeholder - will be populated from security analysis)
	securityChange := SecurityChange{
		SourceRules:      0, // Will be updated from source firewall rules
		TargetRules:      0, // Will be updated from target security groups
		ConversionStatus: "Converted",
		Description:      "Security rules converted to cloud-native format",
	}

	return ResourceChangeAnalysis{
		CPUChange:      cpuChange,
		MemoryChange:   memoryChange,
		StorageChange:  storageChange,
		NetworkChange:  networkChange,
		SecurityChange: securityChange,
	}
}

// determineChangeType determines if a change is an upgrade, downgrade, or same
func determineChangeType(sourceValue, targetValue float64) string {
	if targetValue > sourceValue*1.1 { // More than 10% increase
		return "Upgrade"
	} else if targetValue < sourceValue*0.9 { // More than 10% decrease
		return "Downgrade"
	}
	return "Same"
}

// buildNetworkAnalysis builds network migration analysis
func buildNetworkAnalysis(sourceSummary *summary.SourceInfraSummary, targetSummary *summary.TargetInfraSummary) NetworkMigrationAnalysis {
	// Extract source network info
	sourceNetwork := SourceNetworkInfo{
		CIDR:             "",
		Gateway:          "",
		ConnectedServers: sourceSummary.Overview.TotalServerCount,
	}

	if len(sourceSummary.NetworkResources.Networks) > 0 {
		sourceNetwork.CIDR = sourceSummary.NetworkResources.Networks[0].NetworkCIDR
		sourceNetwork.Gateway = sourceSummary.NetworkResources.Networks[0].Gateway
	}

	// Extract target network info
	targetNetwork := TargetNetworkInfo{}
	if len(targetSummary.NetworkResources.VNets) > 0 {
		vnet := targetSummary.NetworkResources.VNets[0]
		targetNetwork.VNetName = vnet.Name
		targetNetwork.VNetCIDR = vnet.CidrBlock
		targetNetwork.CSPVNetID = vnet.CspVNetId
		if len(vnet.Subnets) > 0 {
			targetNetwork.SubnetCIDR = vnet.Subnets[0].CidrBlock
		}
	}

	// Build IP mappings
	var ipMappings []IPMapping
	for _, server := range sourceSummary.ComputeResources.Servers {
		machineID := server.MachineId

		// Find matching target VM
		for _, vm := range targetSummary.ComputeResources.Vms {
			if strings.Contains(vm.Name, machineID) {
				mapping := IPMapping{
					SourceIP:        server.Network.IPAddress,
					SourceHostname:  server.Hostname,
					TargetPrivateIP: vm.Misc.PrivateIp,
					TargetPublicIP:  vm.Misc.PublicIp,
				}
				ipMappings = append(ipMappings, mapping)
				break
			}
		}
	}

	// Check if CIDR is preserved
	cidrPreserved := strings.Contains(targetNetwork.SubnetCIDR, sourceNetwork.CIDR) ||
		strings.Contains(sourceNetwork.CIDR, targetNetwork.SubnetCIDR)

	description := "Network configuration successfully migrated to cloud VNet/Subnet"
	if cidrPreserved {
		description = "Network CIDR preserved in target subnet"
	}

	return NetworkMigrationAnalysis{
		SourceNetwork: sourceNetwork,
		TargetNetwork: targetNetwork,
		IPMappings:    ipMappings,
		CIDRPreserved: cidrPreserved,
		Description:   description,
	}
}

// buildSecurityAnalysis builds security migration analysis
func buildSecurityAnalysis(sourceSummary *summary.SourceInfraSummary, targetSummary *summary.TargetInfraSummary, mappings []SourceTargetMapping) SecurityMigrationAnalysis {
	var conversions []SecurityConversion

	for _, mapping := range mappings {
		sourceHostname := mapping.SourceServer.Hostname
		sourceRules := mapping.SourceServer.FirewallRules

		// Find target security groups for this VM
		targetSGs := mapping.TargetVM.SecurityGroups
		totalTargetRules := 0

		for _, sgName := range targetSGs {
			for _, sg := range targetSummary.SecurityResources.SecurityGroups {
				if sg.Name == sgName {
					totalTargetRules += len(sg.Rules)
				}
			}
		}

		if sourceRules > 0 || len(targetSGs) > 0 {
			conversion := SecurityConversion{
				SourceHostname: sourceHostname,
				SourceRules:    sourceRules,
				TargetSGName:   strings.Join(targetSGs, ", "),
				TargetRules:    totalTargetRules,
				ConversionType: "iptables to Cloud Security Group",
				Status:         "Converted",
			}

			if sourceRules == 0 {
				conversion.Status = "Added"
				conversion.ConversionType = "New Security Group"
			}

			conversions = append(conversions, conversion)
		}
	}

	summary := fmt.Sprintf("Converted %d server firewall configurations to cloud-native security groups", len(conversions))

	return SecurityMigrationAnalysis{
		Conversions: conversions,
		Summary:     summary,
	}
}

// buildCostSummary builds cost summary
func buildCostSummary(targetSummary *summary.TargetInfraSummary, mappings []SourceTargetMapping) CostSummary {
	totalMonthlyCost := float64(targetSummary.CostEstimation.TotalCostPerMonth)

	var costByComponent []ComponentCost
	for _, mapping := range mappings {
		percentage := 0.0
		if totalMonthlyCost > 0 {
			percentage = (mapping.CostPerMonth / totalMonthlyCost) * 100
		}

		component := ComponentCost{
			ComponentName:  fmt.Sprintf("%s (migrated)", mapping.SourceServer.Hostname),
			SpecName:       mapping.TargetVM.SpecName,
			MonthlyCost:    mapping.CostPerMonth,
			CostPercentage: percentage,
		}
		costByComponent = append(costByComponent, component)
	}

	return CostSummary{
		TotalHourlyCost:  float64(targetSummary.CostEstimation.TotalCostPerHour),
		TotalDailyCost:   float64(targetSummary.CostEstimation.TotalCostPerDay),
		TotalMonthlyCost: totalMonthlyCost,
		TotalYearlyCost:  totalMonthlyCost * 12,
		CostByComponent:  costByComponent,
	}
}

// generateRecommendations generates recommendations based on the migration
func generateRecommendations(sourceSummary *summary.SourceInfraSummary, targetSummary *summary.TargetInfraSummary, mappings []SourceTargetMapping) []Recommendation {
	var recommendations []Recommendation

	// Storage recommendation
	totalSourceDisk := sourceSummary.Overview.TotalDiskGB
	totalTargetDisk := 0
	for range targetSummary.ComputeResources.Vms {
		totalTargetDisk += 50 // Assuming 50GB root disks
	}

	if totalSourceDisk > totalTargetDisk*2 {
		recommendations = append(recommendations, Recommendation{
			Category:    "Storage Optimization",
			Priority:    "High",
			Title:       "Consider adding additional EBS volumes",
			Description: fmt.Sprintf("Source infrastructure has %d GB total disk, but target only has %d GB. Consider attaching additional volumes for data storage.", totalSourceDisk, totalTargetDisk),
			ActionItems: []string{
				"Review data storage requirements",
				"Attach additional EBS volumes as needed",
				"Set up automated backups",
			},
		})
	}

	// Cost optimization recommendation
	if targetSummary.CostEstimation.TotalCostPerMonth > 300 {
		savingsEstimate := float64(targetSummary.CostEstimation.TotalCostPerMonth) * 0.35
		recommendations = append(recommendations, Recommendation{
			Category:    "Cost Optimization",
			Priority:    "Medium",
			Title:       "Consider Reserved Instances for cost savings",
			Description: fmt.Sprintf("Reserved Instances can save up to 35%% (~$%.2f/month)", savingsEstimate),
			ActionItems: []string{
				"Analyze usage patterns",
				"Compare Reserved Instance pricing",
				"Plan for 1 or 3 year commitment",
			},
		})
	}

	// Performance monitoring recommendation
	recommendations = append(recommendations, Recommendation{
		Category:    "Monitoring",
		Priority:    "High",
		Title:       "Set up comprehensive monitoring",
		Description: "Configure CloudWatch monitoring and alerts for the migrated infrastructure",
		ActionItems: []string{
			"Enable detailed CloudWatch monitoring",
			"Set up CPU and memory alerts",
			"Configure log aggregation",
			"Create operational dashboards",
		},
	})

	// Security recommendation
	recommendations = append(recommendations, Recommendation{
		Category:    "Security",
		Priority:    "High",
		Title:       "Review and optimize security group rules",
		Description: "Ensure security groups follow the principle of least privilege",
		ActionItems: []string{
			"Review all security group rules",
			"Remove overly permissive rules (0.0.0.0/0)",
			"Implement network segmentation",
			"Enable VPC flow logs",
		},
	})

	return recommendations
}

// Helper functions

// extractSourceMachineID extracts machine ID from VM name
// Example: "migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1" -> "0036e4b9-c8b4-e811-906e-000ffee02d5c"
func extractSourceMachineID(vmName string) string {
	// Remove "migrated-" prefix and "-1" suffix
	name := strings.TrimPrefix(vmName, "migrated-")
	parts := strings.Split(name, "-")
	if len(parts) >= 5 {
		// Reconstruct the GUID format
		return strings.Join(parts[:5], "-")
	}
	return ""
}

// extractMachineIDFromHostname attempts to extract machine ID from source server
// This is a placeholder - actual implementation depends on source data structure
func extractMachineIDFromHostname(hostname string) string {
	// In real implementation, this would look up machine ID from source data
	// For now, we'll use a placeholder that matches the VM naming pattern

	// This is a simplified approach - in production, you'd need actual machine IDs
	// from the source infrastructure data
	return "" // Will be populated from actual source data
}
