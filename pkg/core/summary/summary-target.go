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
func GenerateInfraSummary(nsId, mciId string) (*TargetInfraSummary, error) {
	log.Info().Msgf("Generating infrastructure summary for MCI (nsId: %s, mciId: %s)", nsId, mciId)

	// Step 1: Collect MCI information
	mciInfo, err := tbclient.NewSession().ReadMci(nsId, mciId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve MCI information")
		return nil, fmt.Errorf("failed to retrieve MCI information: %w", err)
	}

	// Step 2: Extract unique resource IDs
	uniqueVNetIds := extractUniqueVNetIds(mciInfo.Vm)
	uniqueSshKeyIds := extractUniqueSshKeyIds(mciInfo.Vm)
	uniqueSecurityGroupIds := extractUniqueSecurityGroupIds(mciInfo.Vm)
	uniqueSpecIds := extractUniqueSpecIds(mciInfo.Vm)
	uniqueImageIds := extractUniqueImageIds(mciInfo.Vm)

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
	computeResources, err := collectComputeResources(nsId, &mciInfo, uniqueSpecIds, uniqueImageIds)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to collect some compute resources")
	}

	// Step 6: Calculate cost estimation
	costEstimation := calculateCostEstimation(computeResources)

	// Step 7: Build summary metadata
	metadata := TargetSummaryMetadata{
		GeneratedAt:    time.Now(),
		Namespace:      nsId,
		MciId:          mciId,
		MciName:        mciInfo.Name,
		SummaryVersion: "1.0",
	}

	// Step 8: Build infrastructure overview
	overview := buildInfraOverview(&mciInfo)

	// Step 9: Assemble final summary
	infraSummary := &TargetInfraSummary{
		SummaryMetadata:   metadata,
		Overview:          overview,
		NetworkResources:  networkResources,
		SecurityResources: securityResources,
		ComputeResources:  computeResources,
		CostEstimation:    costEstimation,
	}

	log.Info().Msgf("Successfully generated infrastructure summary for MCI: %s", mciId)
	return infraSummary, nil
}

// extractUniqueVNetIds extracts unique VNet IDs from VMs
func extractUniqueVNetIds(vms []tbmodel.VmInfo) []string {
	idMap := make(map[string]struct{})
	for _, vm := range vms {
		if vm.VNetId != "" {
			idMap[vm.VNetId] = struct{}{}
		}
	}

	var ids []string
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// extractUniqueSshKeyIds extracts unique SSH Key IDs from VMs
func extractUniqueSshKeyIds(vms []tbmodel.VmInfo) []string {
	idMap := make(map[string]struct{})
	for _, vm := range vms {
		if vm.SshKeyId != "" {
			idMap[vm.SshKeyId] = struct{}{}
		}
	}

	var ids []string
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// extractUniqueSecurityGroupIds extracts unique Security Group IDs from VMs
func extractUniqueSecurityGroupIds(vms []tbmodel.VmInfo) []string {
	idMap := make(map[string]struct{})
	for _, vm := range vms {
		for _, sgId := range vm.SecurityGroupIds {
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

// extractUniqueSpecIds extracts unique Spec IDs from VMs
func extractUniqueSpecIds(vms []tbmodel.VmInfo) []string {
	idMap := make(map[string]struct{})
	for _, vm := range vms {
		if vm.SpecId != "" {
			idMap[vm.SpecId] = struct{}{}
		}
	}

	var ids []string
	for id := range idMap {
		ids = append(ids, id)
	}
	return ids
}

// extractUniqueImageIds extracts unique Image IDs from VMs
func extractUniqueImageIds(vms []tbmodel.VmInfo) []string {
	idMap := make(map[string]struct{})
	for _, vm := range vms {
		if vm.ImageId != "" {
			idMap[vm.ImageId] = struct{}{}
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

// collectComputeResources collects Spec, Image, and VM information
func collectComputeResources(nsId string, mciInfo *tbmodel.MciInfo, specIds, imageIds []string) (SummaryComputeResources, error) {
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

	// Count usage and build VM list
	for _, vm := range mciInfo.Vm {
		specUsage[vm.SpecId]++
		imageUsage[vm.ImageId]++
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

	// Build VM list for report
	for _, vm := range mciInfo.Vm {
		spec := specMap[vm.SpecId]
		image := imageMap[vm.ImageId]

		reportVm := SummaryVmInfo{
			Name:    vm.Name,
			CspVmId: vm.CspResourceId,
			Status:  vm.Status,
			Spec: SummaryVmSpecInfo{
				Name:         extractShortSpecName(vm.CspSpecName),
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
				VNet:           vm.VNetId,
				Subnet:         vm.SubnetId,
				PublicIp:       vm.PublicIP,
				PrivateIp:      vm.PrivateIP,
				SecurityGroups: vm.SecurityGroupIds,
				SshKey:         vm.SshKeyId,
				ConnectionName: vm.ConnectionName,
			},
			Region: vm.Region.Region,
			Zone:   vm.Region.Zone,
		}

		resources.Vms = append(resources.Vms, reportVm)
	}

	return resources, nil
}

// buildInfraOverview builds the migration summary from MCI info
func buildInfraOverview(mciInfo *tbmodel.MciInfo) TargetInfraOverview {
	runningCount := 0
	stoppedCount := 0

	for _, vm := range mciInfo.Vm {
		if strings.EqualFold(vm.Status, "running") {
			runningCount++
		} else if strings.EqualFold(vm.Status, "stopped") || strings.EqualFold(vm.Status, "terminated") {
			stoppedCount++
		}
	}

	targetCloud := "Unknown"
	targetRegion := "Unknown"
	if len(mciInfo.Vm) > 0 {
		targetCloud = strings.ToUpper(mciInfo.Vm[0].ConnectionConfig.ProviderName)
		targetRegion = mciInfo.Vm[0].Region.Region
	}

	return TargetInfraOverview{
		MciName:         mciInfo.Name,
		MciDescription:  mciInfo.Description,
		Status:          mciInfo.Status,
		TargetCloud:     targetCloud,
		TargetRegion:    targetRegion,
		TotalVmCount:    len(mciInfo.Vm),
		RunningVmCount:  runningCount,
		StoppedVmCount:  stoppedCount,
		Label:           mciInfo.Label,
		InstallMonAgent: mciInfo.InstallMonAgent,
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
