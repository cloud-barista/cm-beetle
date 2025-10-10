package recommendation

import (
	"fmt"
	"strings"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"

	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/rs/zerolog/log"
)

func isSupportedCSP(csp string) bool {
	supportedCSPs := map[string]bool{
		"aws":     true,
		"azure":   true,
		"gcp":     true,
		"alibaba": true,
		// "tencent": true,
		// "ibm":   true,
		"ncp": true,
		// "nhn": true,
		// "kt": true,
		// "openstack": true,
	}

	return supportedCSPs[csp]
}

func IsValidCspAndRegion(csp string, region string) (bool, error) {

	isValid := false
	cspName := strings.ToLower(csp)
	regionName := strings.ToLower(region)
	supportedCsp := isSupportedCSP(cspName)

	if !supportedCsp {
		err := fmt.Errorf("not supported yet (provider: %s)", cspName)
		log.Warn().Msgf("%s", err.Error())
		return isValid, err
	}

	apiConfig := tbclient.ApiConfig{
		RestUrl:  config.Tumblebug.RestUrl,
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
	}

	tbCli := tbclient.NewClient(apiConfig)

	// Check if the region is valid for the specified CSP
	_, err := tbCli.ReadRegionInfo(cspName, regionName)
	if err != nil {
		log.Warn().Msgf("failed to read region info for CSP %s and region %s: %v", cspName, regionName, err)
		return isValid, err
	}

	isValid = true

	return isValid, nil
}

// RecommendVmInfraWithDefaults an appropriate multi-cloud infrastructure (MCI) for cloud migration
func RecommendVmInfraWithDefaults(desiredCsp string, desiredRegion string, srcInfra onpremmodel.OnpremInfra) (cloudmodel.RecommendedVmInfraDynamicList, error) {

	// var emptyResp RecommendedVmInfraInfoList
	var recommendedVmInfraInfoList cloudmodel.RecommendedVmInfraDynamicList

	// TODO: To be updated, a user will input the desired number of recommended VMs
	var defaultSpecsLimit int = GetDefaultSpecsLimit()

	// Initialize the response body
	recommendedVmInfraInfoList = cloudmodel.RecommendedVmInfraDynamicList{
		Description:       "This is a list of recommended target infrastructures. Please review and use them.",
		Count:             0,
		TargetVmInfraList: []cloudmodel.RecommendedVmInfraDynamic{},
	}

	// // Set VM info
	// recommendedVm := tbmodel.TbVmDynamicReq{
	// 	ConnectionName: "",
	// 	CommonImage:    "", // Lookup and set an appropriate VM OS image
	// 	CommonSpec:     "", // Lookup and set an appropriate VM spec
	// 	Description:    "a recommended virtual machine",
	// 	Name:           fmt.Sprintf("migrated-%s", server.MachineId),
	// 	RootDiskSize:   "", // TBD
	// 	RootDiskType:   "", // TBD
	// 	SubGroupSize:   "",
	// 	VmUserPassword: "",
	// }

	/*
	 * [Process]
	 */
	type RecommendedVmInfo struct {
		vmSpecId    string
		vmOsImageId string
	}

	recommendedVmInfoList := [][]RecommendedVmInfo{}

	// Recommand VM specs and OS images for servers in the source computing infrastructure
	for _, server := range srcInfra.Servers {

		// Lookup the appropriate VM specs for the server
		vmSpecList, _, err := RecommendVmSpecs(desiredCsp, desiredRegion, server, defaultSpecsLimit)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM specs for server %s: %v", server.MachineId, err)
			continue
		}

		// Lookup the appropriate VM OS images for the server
		vmOsImageIdList := []string{}
		for range vmSpecList {
			osImgId, err := RecommendVmOsImageId(desiredCsp, desiredRegion, server)
			if err != nil {
				log.Warn().Msgf("failed to recommend VM OS image for server %s: %v", server.MachineId, err)
				vmOsImageIdList = append(vmOsImageIdList, "")
			} else {
				vmOsImageIdList = append(vmOsImageIdList, osImgId)
			}
		}

		// Set the recommended VM specs and OS images to the response body
		recommendedVmInfo := []RecommendedVmInfo{}
		for i, vmSpec := range vmSpecList {
			recommendedVmInfo = append(recommendedVmInfo, RecommendedVmInfo{
				vmSpecId:    vmSpec.CspSpecName,
				vmOsImageId: vmOsImageIdList[i],
			})
		}
		recommendedVmInfoList = append(recommendedVmInfoList, recommendedVmInfo)
	}

	// Debug log
	// log.Debug().Msgf("the number of recommended VM specs and OS images: %d", len(recommendedVmInfoList))
	// for i, vmInfoList := range recommendedVmInfoList {
	// 	log.Debug().Msgf("the number of recommended VM specs and OS images for server %d: %d", i, len(vmInfoList))
	// }
	// log.Debug().Msgf("recommended VM specs and OS images: %+v", recommendedVmInfoList)

	/*
	 * [Output]
	 */

	// Transpose the matrix to change from "VM recommendations per server" to "servers per VM recommendation".
	// Before: [Server1's VM recommendations, Server2's VM recommendations, ...]
	// After: [Recommendation1 for all servers, Recommendation2 for all servers, ...]
	transposed := transposeMatrix(recommendedVmInfoList)
	// log.Debug().Msgf("transposed recommended VM specs and OS images: %+v", transposed)

	// Build response body which includes multiple recommended infrastructures
	recommenedVmInfraInfoList := []cloudmodel.RecommendedVmInfraDynamic{}

	for i, vmInfoList := range transposed {

		tempVmInfraInfo := cloudmodel.RecommendedVmInfraDynamic{
			Status:      string(NothingRecommended),
			Description: "This is a recommended target infrastructure.",
			TargetVmInfra: cloudmodel.MciDynamicReq{
				Name:        fmt.Sprintf("migrated-%02d", i),
				Description: "a recommended multi-cloud infrastructure",
				SubGroups:   []cloudmodel.CreateSubGroupDynamicReq{},
			},
		}

		for j, subgroupInfo := range vmInfoList {
			tempCreateSubgroupReq := cloudmodel.CreateSubGroupDynamicReq{
				ConnectionName: fmt.Sprintf("%s-%s", desiredCsp, desiredRegion),
				ImageId:        subgroupInfo.vmOsImageId,
				SpecId:         subgroupInfo.vmSpecId,
				Description:    "a recommended virtual machine",
				Name:           fmt.Sprintf("migrated-%s", srcInfra.Servers[j].MachineId), // Set MachineId to identify the source server
				RootDiskSize:   "",                                                        // TBD
				RootDiskType:   "",                                                        // TBD
				SubGroupSize:   "",
				VmUserPassword: "",
			}
			tempVmInfraInfo.TargetVmInfra.SubGroups = append(tempVmInfraInfo.TargetVmInfra.SubGroups, tempCreateSubgroupReq)
		}

		status := checkOverallSubGroupStatus(tempVmInfraInfo.TargetVmInfra.SubGroups)
		tempVmInfraInfo.Status = status
		if status == string(NothingRecommended) {
			tempVmInfraInfo.Description = "Could not find approprate VMs."
		} else if status == string(FullyRecommended) {
			tempVmInfraInfo.Description = "Target infra is recommended."
		} else {
			tempVmInfraInfo.Description = "Some VMs are recommended. Please check and fill the required information."
		}

		recommenedVmInfraInfoList = append(recommenedVmInfraInfoList, tempVmInfraInfo)
	}

	// Assign the target infrastructure list to the response
	recommendedVmInfraInfoList.TargetVmInfraList = recommenedVmInfraInfoList
	recommendedVmInfraInfoList.Count = len(recommenedVmInfraInfoList)

	log.Trace().Msgf("the recommended infra info: %+v", recommendedVmInfraInfoList)

	return recommendedVmInfraInfoList, nil
}

// RecommendVmInfra an appropriate multi-cloud infrastructure (MCI) for cloud migration
func RecommendVmInfra(desiredCsp string, desiredRegion string, srcInfra onpremmodel.OnpremInfra) (cloudmodel.RecommendedVmInfra, error) {

	// var emptyResp RecommendedVmInfra
	var recommendedVmInfra cloudmodel.RecommendedVmInfra

	// TODO: To be updated, a user will input the desired number of recommended VMs
	var limitSpecs int = GetDefaultSpecsLimit()
	var limitImages int = GetDefaultImagesLimit()

	// Initialize the response body
	recommendedVmInfra = cloudmodel.RecommendedVmInfra{
		Description: "This is a list of recommended target infrastructures. Please review and use them.",
		Status:      "",
		TargetCloud: cloudmodel.CloudProperty{
			Csp:    desiredCsp,
			Region: desiredRegion,
		},
		TargetVmInfra: cloudmodel.MciReq{
			Name:        "mmci01",
			Description: "a recommended multi-cloud infrastructure",
			SubGroups:   []cloudmodel.CreateSubGroupReq{},
		},
	}

	csp := strings.ToLower(desiredCsp)
	region := strings.ToLower(desiredRegion)

	/*
	 * [Process]
	 */

	// 1. Recommend vNet and subnets (Note: vNet can be a VPC or a VNet depending on the CSP)
	recommendedVNetInfoList, err := RecommendVNet(csp, region, srcInfra)
	if err != nil {
		log.Warn().Err(err).Msg("failed to recommend a virtual network for the source computing infrastructure")
	}

	if len(recommendedVNetInfoList) == 0 {
		log.Warn().Msg("no recommended virtual network found for the source computing infrastructure")
	}

	// Assign the recommended virtural network to the response body
	// TODO: Consider the other index in the recommended virtual network
	recommendedVmInfra.TargetVNet = recommendedVNetInfoList[0]

	// * Set a name to indicate a dependency between resources.
	recommendedVmInfra.TargetVNet.Name = "mig-vnet-01"
	recommendedVmInfra.TargetVNet.Description = "a recommended vNet for migration"
	for i := range recommendedVmInfra.TargetVNet.SubnetInfoList {
		recommendedVmInfra.TargetVNet.SubnetInfoList[i].Name = fmt.Sprintf("mig-subnet-%02d", i+1)
		recommendedVmInfra.TargetVNet.SubnetInfoList[i].Description = "a recommended subnet for migration"
	}

	// 2. Recommend(?) SSH key pair
	// var recommendedSshKey = tbmodel.SshKeyReq{}
	// * Set a name to indicate a dependency between resources.
	recommendedVmInfra.TargetSshKey.Name = "mig-sshkey-01"
	recommendedVmInfra.TargetSshKey.ConnectionName = fmt.Sprintf("%s-%s", csp, region)
	recommendedVmInfra.TargetSshKey.Description = "a SSH Key pair for migration (Note - provided ONLY once, MUST be downloaded"

	// 3. Recommend VM specs, OS images, and security groups, and
	// recommend VMs by removing duplicates of VM specs, OS images, and security groups and specifying them.
	// Note: Don't need to register specs and OS images.
	var recommendedSubgroupList = []cloudmodel.CreateSubGroupReq{}
	var recommendedVmSpecList = []cloudmodel.SpecInfo{}
	var recommendedVmOsImageList = []cloudmodel.ImageInfo{}
	var recommendedSecurityGroupList = []cloudmodel.SecurityGroupReq{}

	for i, server := range srcInfra.Servers {

		/*
		 * Recommend VM specs, OS images, and security groups
		 */

		// Lookup the appropriate VM specs for the server
		recommendedVmSpecInfoList, _, err := RecommendVmSpecs(csp, region, server, limitSpecs)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM specs for server %s: %v", server.MachineId, err)
		}

		// Lookup the appropriate VM OS images for the server
		// recommendedVmOsImageInfo, err := RecommendVmOsImage(csp, region, server)
		recommendedVmOsImageInfoList, err := RecommendVmOsImages(csp, region, server, limitImages)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM OS images for server %s: %v", server.MachineId, err)
		}

		// Generete security group from the server's firewall rules (or firewall table)
		recommendedSg, err := RecommendSecurityGroup(csp, region, server)
		if err != nil {
			log.Warn().Msgf("failed to recommend security group for server %s: %v", server.MachineId, err)
		}

		lenSpecList := len(recommendedVmSpecInfoList)
		lenImageList := len(recommendedVmOsImageInfoList)
		log.Debug().Msgf("length of recommendedVmSpecInfoList: %d", lenSpecList)
		log.Debug().Msgf("length of recommendedVmOsImageInfoList: %d", lenImageList)

		// Logging the first 3 items to avoid excessive output
		loggingLimit := 3
		for i := 0; i < lenSpecList && i < loggingLimit; i++ {
			log.Debug().Msgf("(logging up to 3 specs) recommendedVmSpecInfoList[%d]: %+v", i, recommendedVmSpecInfoList[i])
		}
		for i := 0; i < lenImageList && i < loggingLimit; i++ {
			log.Debug().Msgf("(logging up to 3 images) recommendedVmOsImageInfoList[%d]: %+v", i, recommendedVmOsImageInfoList[i])
		}

		var selectedVmSpec cloudmodel.SpecInfo
		var selectedVmOsImage cloudmodel.ImageInfo
		if len(recommendedVmSpecInfoList) == 0 || len(recommendedVmOsImageInfoList) == 0 {
			log.Warn().Msgf("no recommended VM specs or OS images found for server %s", server.MachineId)
		} else {

			// * Note: (opinion) Find multiple compatible pairs and use them as needed in the later process
			// Find compatible spec and image pair
			tempSelectedVmSpec, tempSelectedVmOsImage, err := FindCompatibleSpecAndImage(recommendedVmSpecInfoList, recommendedVmOsImageInfoList, csp)
			if err != nil {
				log.Warn().Msgf("failed to find compatible spec-image pair for server %s: %v", server.MachineId, err)
				// Use fallback selection (first spec, first image)
			} else {
				selectedVmSpec = tempSelectedVmSpec
				selectedVmOsImage = tempSelectedVmOsImage

				// Log CPU comparison
				log.Debug().
					Str("machineId", server.MachineId).
					Str("specCspName", selectedVmSpec.CspSpecName).
					Str("specId", selectedVmSpec.Id).
					Uint32("originalCPUs", server.CPU.Cpus).
					Uint32("recommendedVCPU", uint32(selectedVmSpec.VCPU)).
					Msg("CPU comparison")

				// Log Memory comparison
				log.Debug().
					Str("machineId", server.MachineId).
					Str("specCspName", selectedVmSpec.CspSpecName).
					Str("specId", selectedVmSpec.Id).
					Uint32("originalMemoryGB", uint32(server.Memory.TotalSize)).
					Float32("recommendedMemoryGiB", selectedVmSpec.MemoryGiB).
					Msg("Memory comparison")

				// Log OS comparison
				log.Debug().
					Str("machineId", server.MachineId).
					Str("imageCspName", selectedVmOsImage.CspImageName).
					Str("imageId", selectedVmOsImage.Id).
					Str("originalOS", server.OS.Name+" "+server.OS.Version).
					Str("recommendedOSImage", selectedVmOsImage.CspImageName).
					Msg("OS comparison")
			}
		}

		/*
		 * Check duplicate and append the recommended VM specs, OS images, and security groups
		 */
		// Check duplicates and append the recommended VM specs
		// * Note: Use the name of the VM spec managed by Tumblebug
		exists := false
		// If the recommended VM spec already exists in the list, select the existing spec
		for _, vmSpec := range recommendedVmSpecList {
			if vmSpec.CspSpecName == selectedVmSpec.CspSpecName {
				exists = true
				selectedVmSpec = vmSpec
				break
			}
		}
		if !exists {
			recommendedVmSpecList = append(recommendedVmSpecList, selectedVmSpec)
		}

		// Check duplicates and append the recommended VM OS images
		// * Note: Use the name of the VM OS image managed by Tumblebug
		log.Debug().Msgf("selectedVmOsImage: %+v", selectedVmOsImage)
		exists = false
		// If the recommended VM OS image already exists in the list, select the existing OS image
		for _, vmOsImage := range recommendedVmOsImageList {
			if vmOsImage.CspImageName == selectedVmOsImage.CspImageName {
				exists = true
				selectedVmOsImage = vmOsImage
				break
			}
		}
		if !exists {
			recommendedVmOsImageList = append(recommendedVmOsImageList, selectedVmOsImage)
		}

		// Check duplicates and append the recommended security groups
		exists, _, existingSg := containSg(recommendedSecurityGroupList, recommendedSg)
		if !exists {
			// If the security group does not exist, set a name to indicate a dependency between resources.
			recommendedSg.Name = fmt.Sprintf("mig-sg-%02d", len(recommendedSecurityGroupList)+1)
			recommendedSg.ConnectionName = fmt.Sprintf("%s-%s", csp, region)
			recommendedSg.Description = fmt.Sprintf("Recommended security group for %s", server.MachineId) // Set MachineId to identify the source server

			// * Set name to indicate a dependency between resources.
			recommendedSg.VNetId = recommendedVmInfra.TargetVNet.Name // Set the vNet ID to the security group

			// Set the security group to the response body
			recommendedSecurityGroupList = append(recommendedSecurityGroupList, recommendedSg)
		} else {
			recommendedSg = existingSg
		}

		/*
		 * Recommend VM by specifying the recommended VM specs, OS images, and security groups
		 */
		// TODO: Select a subnet by the server's network information (for now, select the first one)

		// Ref: https://github.com/cloud-barista/cb-spider/blob/master/cloud-driver-libs/cloudos_meta.yaml
		// Note: "TYPE1" for RootDiskType is the first in the list
		// - AWS: ["standard", "gp2", "gp3"],
		// - Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"],
		// - GCP: [ "pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"],
		// - ALIBABA: ["cloud_essd", "cloud_efficiency", "cloud", "cloud_ssd"],
		// - TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
		// - NCP: ["HDD"]
		// - NHN: ["General_HDD", "General_SSD"]
		// - KT: ["HDD", "SSD"]

		// * Set names to indicate a dependency between resources.
		tempCreateSubGroupReq := cloudmodel.CreateSubGroupReq{
			ConnectionName:   fmt.Sprintf("%s-%s", csp, region),
			Description:      fmt.Sprintf("a recommended virtual machine %02d for %s", i+1, server.MachineId), // Set MachineId to identify the source server
			SpecId:           selectedVmSpec.Id,
			ImageId:          selectedVmOsImage.Id,
			VNetId:           recommendedVmInfra.TargetVNet.Name,
			SubnetId:         recommendedVmInfra.TargetVNet.SubnetInfoList[0].Name, // Set the first subnet for simplicity (TBD, select the appropriate subnet)
			SecurityGroupIds: []string{recommendedSg.Name},                         // Set the security group ID
			Name:             fmt.Sprintf("migrated-%s", server.MachineId),         // Set MachineId to identify the source server
			RootDiskType:     "",                                                   // Set "" or default to use CSP's default
			RootDiskSize:     "30",                                                 // Set 30 GiB as a default value
			SshKeyId:         recommendedVmInfra.TargetSshKey.Name,                 // Set the SSH key ID
			VmUserName:       "",                                                   // TBD: Set the VM user name if needed
			VmUserPassword:   "",                                                   // TBD
			SubGroupSize:     "",                                                   // TBD
			Label: map[string]string{
				"sourceMachineId": server.MachineId,
			},
		}

		// ! Set the root disk type for Alibaba Cloud
		if csp == "alibaba" && tempCreateSubGroupReq.RootDiskType == "" {
			log.Warn().Msg("set the root disk type to 'cloud_essd' for Alibaba Cloud")
			tempCreateSubGroupReq.RootDiskType = "TYPE1" // "cloud_essd"
		}

		// Append the VM request to the list
		recommendedSubgroupList = append(recommendedSubgroupList, tempCreateSubGroupReq)
	}

	/*
	 * [Output]
	 */
	recommendedVmInfra.TargetVmInfra.SubGroups = recommendedSubgroupList
	recommendedVmInfra.TargetVmSpecList = recommendedVmSpecList
	recommendedVmInfra.TargetVmOsImageList = recommendedVmOsImageList
	recommendedVmInfra.TargetSecurityGroupList = recommendedSecurityGroupList

	log.Trace().Msgf("the recommended infra info: %+v", recommendedVmInfra)

	return recommendedVmInfra, nil
}
