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

// Package summary provides infrastructure summary generation logic
package summary

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// GenerateInfraSummary generates a comprehensive infrastructure summary
// Note: This function does NOT modify existing tbmodel structs, only reads from them
func GenerateInfraSummary(nsId, infraId string) (*TargetInfraSummary, error) {
	log.Info().Msgf("Generating infrastructure summary for Infra (nsId: %s, infraId: %s)", nsId, infraId)

	// Step 1: Collect Infra information
	infraInfo, err := tbclient.NewSession().ReadInfra(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve Infra information")
		return nil, fmt.Errorf("failed to retrieve Infra information: %w", err)
	}

	// Step 2: Extract unique resource IDs
	uniqueVNetIds := extractUniqueVNetIds(infraInfo.Node)
	uniqueSshKeyIds := extractUniqueSshKeyIds(infraInfo.Node)
	uniqueSecurityGroupIds := extractUniqueSecurityGroupIds(infraInfo.Node)
	uniqueSpecIds := extractUniqueSpecIds(infraInfo.Node)
	uniqueImageIds := extractUniqueImageIds(infraInfo.Node)

	log.Debug().Msgf("Unique resource counts - VNets: %d, SSHKeys: %d, SecurityGroups: %d, Specs: %d, Images: %d",
		len(uniqueVNetIds), len(uniqueSshKeyIds), len(uniqueSecurityGroupIds), len(uniqueSpecIds), len(uniqueImageIds))

	// Step 3: Collect network resources
	networkResources, err := collectNetworkResources(nsId, uniqueVNetIds)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to collect some network resources")
	}

	// Step 4: Collect security resources
	securityResources, err := collectSecurityResources(nsId, uniqueSshKeyIds, uniqueSecurityGroupIds)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to collect some security resources")
	}

	// Step 5: Collect compute resources
	computeResources, err := collectComputeResources(nsId, &infraInfo, uniqueSpecIds, uniqueImageIds)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to collect some compute resources")
	}

	// Step 6: Calculate cost estimation
	costEstimation := calculateCostEstimation(computeResources)

	// Step 7: Build summary metadata
	metadata := TargetSummaryMetadata{
		GeneratedAt:    time.Now(),
		Namespace:      nsId,
		InfraId:        infraId,
		InfraName:      infraInfo.Name,
		SummaryVersion: "1.0",
	}

	// Step 8: Build infrastructure overview
	overview := buildInfraOverview(&infraInfo)

	// Step 9: Assemble final summary
	infraSummary := &TargetInfraSummary{
		SummaryMetadata:   metadata,
		Overview:          overview,
		NetworkResources:  networkResources,
		SecurityResources: securityResources,
		ComputeResources:  computeResources,
		CostEstimation:    costEstimation,
	}

	log.Info().Msgf("Successfully generated infrastructure summary for Infra: %s", infraId)
	return infraSummary, nil
}

// extractUniqueVNetIds extracts unique VNet IDs from Nodes
func extractUniqueVNetIds(nodes []tbmodel.NodeInfo) []string {
	idMap := make(map[string]struct{})
	for _, node := range nodes {
		if node.VNetId != "" {
			idMap[node.VNetId] = struct{}{}
		}
	}

	var ids []string
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// extractUniqueSshKeyIds extracts unique SSH Key IDs from Nodes
func extractUniqueSshKeyIds(nodes []tbmodel.NodeInfo) []string {
	idMap := make(map[string]struct{})
	for _, node := range nodes {
		if node.SshKeyId != "" {
			idMap[node.SshKeyId] = struct{}{}
		}
	}

	var ids []string
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// extractUniqueSecurityGroupIds extracts unique Security Group IDs from Nodes
func extractUniqueSecurityGroupIds(nodes []tbmodel.NodeInfo) []string {
	idMap := make(map[string]struct{})
	for _, node := range nodes {
		for _, sgId := range node.SecurityGroupIds {
			if sgId != "" {
				idMap[sgId] = struct{}{}
			}
		}
	}

	var ids []string
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// extractUniqueSpecIds extracts unique Spec IDs from Nodes
func extractUniqueSpecIds(nodes []tbmodel.NodeInfo) []string {
	idMap := make(map[string]struct{})
	for _, node := range nodes {
		if node.SpecId != "" {
			idMap[node.SpecId] = struct{}{}
		}
	}

	var ids []string
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// extractUniqueImageIds extracts unique Image IDs from Nodes
func extractUniqueImageIds(nodes []tbmodel.NodeInfo) []string {
	idMap := make(map[string]struct{})
	for _, node := range nodes {
		if node.ImageId != "" {
			idMap[node.ImageId] = struct{}{}
		}
	}

	var ids []string
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// collectNetworkResources collects VNet and Subnet information
func collectNetworkResources(nsId string, vnetIds []string) (SummaryNetworkResources, error) {
	var resources SummaryNetworkResources
	resources.VNets = []SummaryVNetInfo{}

	for _, vnetId := range vnetIds {
		vnetInfo, err := tbclient.NewSession().ReadVNet(nsId, vnetId)
		if err != nil {
			log.Warn().Err(err).Msgf("Failed to retrieve VNet: %s", vnetId)
			continue
		}

		// Convert subnets
		var subnets []SummarySubnetInfo
		for _, subnet := range vnetInfo.SubnetInfoList {
			subnets = append(subnets, SummarySubnetInfo{
				Name:        subnet.Name,
				CspSubnetId: subnet.CspResourceId,
				CidrBlock:   subnet.IPv4_CIDR,
				Zone:        subnet.Zone,
			})
		}

		reportVNet := SummaryVNetInfo{
			Name:           vnetInfo.Name,
			CspVNetId:      vnetInfo.CspResourceId,
			CidrBlock:      vnetInfo.CidrBlock,
			Region:         "", // Region info not available in VNetInfo
			Subnets:        subnets,
			SubnetCount:    len(subnets),
			ConnectionName: vnetInfo.ConnectionName,
		}

		resources.VNets = append(resources.VNets, reportVNet)
	}

	return resources, nil
}

// collectSecurityResources collects SSH Key and Security Group information
func collectSecurityResources(nsId string, sshKeyIds, securityGroupIds []string) (SummarySecurityResources, error) {
	var resources SummarySecurityResources
	resources.SshKeys = []SummarySshKeyInfo{}
	resources.SecurityGroups = []SummarySecurityGroupInfo{}

	// Collect SSH Keys
	for _, sshKeyId := range sshKeyIds {
		sshKeyInfo, err := tbclient.NewSession().ReadSshKey(nsId, sshKeyId)
		if err != nil {
			log.Warn().Err(err).Msgf("Failed to retrieve SSH Key: %s", sshKeyId)
			continue
		}

		// Truncate public key for security
		publicKey := sshKeyInfo.PublicKey
		if len(publicKey) > 50 {
			publicKey = publicKey[:50] + "..."
		}

		reportSshKey := SummarySshKeyInfo{
			Name:        sshKeyInfo.Name,
			CspSshKeyId: sshKeyInfo.CspResourceId,
			Username:    sshKeyInfo.Username,
			PublicKey:   publicKey,
			Fingerprint: sshKeyInfo.Fingerprint,
		}

		resources.SshKeys = append(resources.SshKeys, reportSshKey)
	}

	// Collect Security Groups
	for _, sgId := range securityGroupIds {
		sgInfo, err := tbclient.NewSession().ReadSecurityGroup(nsId, sgId)
		if err != nil {
			log.Warn().Err(err).Msgf("Failed to retrieve Security Group: %s", sgId)
			continue
		}

		// Convert firewall rules
		var rules []SummaryFirewallRule
		for _, rule := range sgInfo.FirewallRules {
			rules = append(rules, SummaryFirewallRule{
				Direction: rule.Direction,
				Protocol:  rule.Protocol,
				FromPort:  rule.Port,
				ToPort:    rule.Port,
				Cidr:      rule.CIDR,
			})
		}

		reportSg := SummarySecurityGroupInfo{
			Name:               sgInfo.Name,
			CspSecurityGroupId: sgInfo.CspResourceId,
			VNetName:           sgInfo.VNetId,
			Rules:              rules,
			RuleCount:          len(rules),
		}

		resources.SecurityGroups = append(resources.SecurityGroups, reportSg)
	}

	// Sort security groups by name for consistent ordering
	sort.Slice(resources.SecurityGroups, func(i, j int) bool {
		return resources.SecurityGroups[i].Name < resources.SecurityGroups[j].Name
	})

	return resources, nil
}

// collectComputeResources collects Spec, Image, and Node information
func collectComputeResources(nsId string, infraInfo *tbmodel.InfraInfo, specIds, imageIds []string) (SummaryComputeResources, error) {
	var resources SummaryComputeResources

	// Collect specs with usage count
	specMap := make(map[string]*tbmodel.SpecInfo)
	specUsage := make(map[string]int)

	for _, specId := range specIds {
		// Specs are stored in system namespace
		specInfo, err := tbclient.NewSession().ReadVmSpec("system", specId)
		if err != nil {
			log.Warn().Err(err).Msgf("Failed to retrieve Spec: %s", specId)
			continue
		}
		specMap[specId] = &specInfo
	}

	// Collect images with usage count
	imageMap := make(map[string]*tbmodel.ImageInfo)
	imageUsage := make(map[string]int)

	for _, imageId := range imageIds {
		// Images are stored in system namespace
		imageInfo, err := tbclient.NewSession().ReadVmOsImage("system", imageId)
		if err != nil {
			log.Warn().Err(err).Msgf("Failed to retrieve Image: %s", imageId)
			continue
		}
		imageMap[imageId] = &imageInfo
	}

	// Count usage and build Node list
	for _, node := range infraInfo.Node {
		specUsage[node.SpecId]++
		imageUsage[node.ImageId]++
	}

	// Build spec list for report - using tbmodel.SpecInfo directly
	for specId, spec := range specMap {
		specWithUsage := SummarySpecInfoWithUsage{
			SpecInfo:   *spec, // Use full SpecInfo from CB-Tumblebug
			UsageCount: specUsage[specId],
		}
		resources.Specs = append(resources.Specs, specWithUsage)
	}

	// Build image list for report - using tbmodel.ImageInfo directly
	for imageId, image := range imageMap {
		imageWithUsage := SummaryImageInfoWithUsage{
			ImageInfo:  *image, // Use full ImageInfo from CB-Tumblebug
			UsageCount: imageUsage[imageId],
		}
		resources.Images = append(resources.Images, imageWithUsage)
	}

	// Build Node list for report
	for _, node := range infraInfo.Node {
		spec := specMap[node.SpecId]
		image := imageMap[node.ImageId]

		reportVm := SummaryVmInfo{
			Name:    node.Name,
			CspVmId: node.CspResourceId,
			Status:  node.Status,
			Spec: SummaryVmSpecInfo{
				Name:         extractShortSpecName(node.CspSpecName),
				VCpus:        getSpecVCpus(spec),
				MemoryGiB:    getSpecMemory(spec),
				Architecture: getSpecArchitecture(spec),
			},
			Image: SummaryVmImageInfo{
				Name:         extractShortImageName(image),
				Id:           getImageId(image),
				OsType:       getImageOsType(image),
				Distribution: getImageDistribution(image),
				OsVersion:    getImageOsVersion(image),
			},
			Misc: SummaryVmMiscInfo{
				VNet:           node.VNetId,
				Subnet:         node.SubnetId,
				PublicIp:       node.PublicIP,
				PrivateIp:      node.PrivateIP,
				SecurityGroups: node.SecurityGroupIds,
				SshKey:         node.SshKeyId,
				ConnectionName: node.ConnectionName,
			},
			Region: node.Region.Region,
			Zone:   node.Region.Zone,
		}

		resources.Vms = append(resources.Vms, reportVm)
	}

	return resources, nil
}

// buildInfraOverview builds the migration summary from Infra info
func buildInfraOverview(infraInfo *tbmodel.InfraInfo) TargetInfraOverview {
	runningCount := 0
	stoppedCount := 0

	for _, node := range infraInfo.Node {
		if strings.EqualFold(node.Status, "running") {
			runningCount++
		} else if strings.EqualFold(node.Status, "stopped") || strings.EqualFold(node.Status, "terminated") {
			stoppedCount++
		}
	}

	targetCloud := "Unknown"
	targetRegion := "Unknown"
	if len(infraInfo.Node) > 0 {
		targetCloud = strings.ToUpper(infraInfo.Node[0].ConnectionConfig.ProviderName)
		targetRegion = infraInfo.Node[0].Region.Region
	}

	return TargetInfraOverview{
		InfraName:       infraInfo.Name,
		InfraDescription: infraInfo.Description,
		Status:          infraInfo.Status,
		TargetCloud:     targetCloud,
		TargetRegion:    targetRegion,
		TotalVmCount:    len(infraInfo.Node),
		RunningVmCount:  runningCount,
		StoppedVmCount:  stoppedCount,
		Label:           infraInfo.Label,
		InstallMonAgent: infraInfo.InstallMonAgent,
	}
}

// calculateCostEstimation calculates cost estimation
func calculateCostEstimation(resources SummaryComputeResources) SummaryCostEstimation {
	var totalCostPerHour float32 = 0
	var byRegionMap = make(map[string]*SummaryCostByRegion)
	var byVmList []SummaryCostByVm

	// Calculate cost per VM
	for _, vm := range resources.Vms {
		// Find spec cost
		var specCost float32 = 0
		for _, spec := range resources.Specs {
			if spec.CspSpecName == vm.Spec.Name {
				specCost = spec.CostPerHour
				break
			}
		}

		totalCostPerHour += specCost

		// Group by region
		regionKey := vm.Region
		if _, exists := byRegionMap[regionKey]; !exists {
			byRegionMap[regionKey] = &SummaryCostByRegion{
				Csp:    strings.ToUpper(strings.Split(vm.Misc.ConnectionName, "-")[0]),
				Region: vm.Region,
			}
		}
		byRegionMap[regionKey].VmCount++
		byRegionMap[regionKey].CostPerHour += specCost

		// Add to by-VM list
		byVmList = append(byVmList, SummaryCostByVm{
			VmName:       vm.Name,
			SpecName:     vm.Spec.Name,
			CostPerHour:  specCost,
			CostPerMonth: specCost * 24 * 30,
		})
	}

	// Convert region map to slice
	var byRegionList []SummaryCostByRegion
	for _, region := range byRegionMap {
		region.CostPerMonth = region.CostPerHour * 24 * 30
		byRegionList = append(byRegionList, *region)
	}

	return SummaryCostEstimation{
		Currency:          "USD",
		TotalCostPerHour:  totalCostPerHour,
		TotalCostPerDay:   totalCostPerHour * 24,
		TotalCostPerMonth: totalCostPerHour * 24 * 30,
		ByRegion:          byRegionList,
		ByVm:              byVmList,
	}
}

// Helper functions to extract information from tbmodel structs

// Helper function to get spec architecture safely
func getSpecArchitecture(spec *tbmodel.SpecInfo) string {
	if spec == nil {
		return ""
	}
	return spec.Architecture
}

// Helper function to get image distribution safely
func getImageDistribution(image *tbmodel.ImageInfo) string {
	if image == nil {
		return ""
	}

	// OSDistribution is the primary field (e.g., "Ubuntu 22.04~")
	if image.OSDistribution != "" {
		return image.OSDistribution
	}

	// Use OSType as fallback (e.g., "ubuntu 22.04")
	if image.OSType != "" {
		return image.OSType
	}

	// Last resort: use OSPlatform (e.g., "Linux/UNIX")
	if image.OSPlatform != "" {
		return string(image.OSPlatform)
	}

	return ""
}

// Helper function to get image OS version safely
func getImageOsVersion(image *tbmodel.ImageInfo) string {
	if image == nil {
		return ""
	}

	// OSDistribution contains distribution with version (e.g., "Ubuntu 22.04~")
	if image.OSDistribution != "" {
		return image.OSDistribution
	}

	// Fallback to OSType if OSDistribution is not available
	if image.OSType != "" {
		return image.OSType
	}

	return ""
}

// Helper function to get image ID safely
func getImageId(image *tbmodel.ImageInfo) string {
	if image == nil {
		return ""
	}
	// Use CspImageId (e.g., "ami-010be25c3775061c9" for AWS)
	if image.CspImageId != "" {
		return image.CspImageId
	}
	// Fallback to Id field
	if image.Id != "" {
		return image.Id
	}
	return ""
}

// Helper function to get image OS type safely
func getImageOsType(image *tbmodel.ImageInfo) string {
	if image == nil {
		return ""
	}
	// OSType contains OS type with version (e.g., "Ubuntu 22.04")
	if image.OSType != "" {
		return image.OSType
	}
	return ""
}

func extractShortSpecName(fullSpecName string) string {
	// Extract short name from full name (e.g., "aws+ap-northeast-2+t3a.xlarge" -> "t3a.xlarge")
	parts := strings.Split(fullSpecName, "+")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return fullSpecName
}

func extractShortImageName(image *tbmodel.ImageInfo) string {
	if image == nil {
		return ""
	}
	// Extract short name from full name
	return image.Name
}

func getSpecVCpus(spec *tbmodel.SpecInfo) int {
	if spec == nil {
		return 0
	}
	return int(spec.VCPU)
}

func getSpecMemory(spec *tbmodel.SpecInfo) float32 {
	if spec == nil {
		return 0
	}
	return spec.MemoryGiB
}
