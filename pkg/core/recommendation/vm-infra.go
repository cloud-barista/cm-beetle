package recommendation

import (
	"fmt"
	"strings"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/similarity"

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
			RootDiskSize:     "50",                                                 // Set 50 GB as a default value
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

// RecommendVmInfraCandidates an appropriate multi-cloud infrastructure (MCI) for cloud migration
func RecommendVmInfraCandidates(desiredCsp string, desiredRegion string, srcInfra onpremmodel.OnpremInfra, limit int) ([]cloudmodel.RecommendedVmInfra, error) {

	// * To recommend multiple infra candidates (i.e., multiple VM spec and OS image combinations),
	// * this function estimates, recommends or just generates vNets, subnets, SSH key pair, and security groups
	// * and then, recommends compatible pairs of VM specs and OS images from the best option to the alternative ones.
	// * All those are corresponding to the source servers.

	// var emptyResp RecommendedVmInfra
	var recommendedVmInfraCandidates []cloudmodel.RecommendedVmInfra

	// TODO: To be updated, a user will input the desired number of recommended VMs
	var limitSpecs int = GetDefaultSpecsLimit()
	var limitImages int = GetDefaultImagesLimit()

	// Initialize the response body
	skeletonVmInfra := cloudmodel.RecommendedVmInfra{
		Description: "This is a recommended target infrastructures and resources. Please review and use them.",
		Status:      "",
		TargetCloud: cloudmodel.CloudProperty{
			Csp:    desiredCsp,
			Region: desiredRegion,
		},
		TargetVmInfra: cloudmodel.MciReq{
			Name:      "mmci01",
			SubGroups: []cloudmodel.CreateSubGroupReq{},
			// Description: "Recommended VMs comprising the multi-cloud infrastructure",
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
	skeletonVmInfra.TargetVNet = recommendedVNetInfoList[0]

	// * Set a name to indicate a dependency between resources.
	skeletonVmInfra.TargetVNet.Name = "mig-vnet-01"
	skeletonVmInfra.TargetVNet.Description = "a recommended vNet for migration"
	for i := range skeletonVmInfra.TargetVNet.SubnetInfoList {
		skeletonVmInfra.TargetVNet.SubnetInfoList[i].Name = fmt.Sprintf("mig-subnet-%02d", i+1)
		skeletonVmInfra.TargetVNet.SubnetInfoList[i].Description = "a recommended subnet for migration"
	}

	// 2. Recommend(?) SSH key pair
	// var recommendedSshKey = tbmodel.SshKeyReq{}
	// * Set a name to indicate a dependency between resources.
	skeletonVmInfra.TargetSshKey.Name = "mig-sshkey-01"
	skeletonVmInfra.TargetSshKey.ConnectionName = fmt.Sprintf("%s-%s", csp, region)
	skeletonVmInfra.TargetSshKey.Description = "a SSH Key pair for migration (Note - provided ONLY once, MUST be downloaded"

	// 3. Generate a skeleton of SubGroup List for VMs
	var skeletonSubgroupList = []cloudmodel.CreateSubGroupReq{}

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
	for i, server := range srcInfra.Servers {
		// * Set names to indicate a dependency between resources.
		tempCreateSubGroupReq := cloudmodel.CreateSubGroupReq{
			ConnectionName: fmt.Sprintf("%s-%s", csp, region),
			Description:    fmt.Sprintf("a recommended virtual machine %02d for %s", i+1, server.MachineId), // Set MachineId to identify the source server
			VNetId:         skeletonVmInfra.TargetVNet.Name,
			SubnetId:       skeletonVmInfra.TargetVNet.SubnetInfoList[0].Name, // Set the first subnet for simplicity (TBD, select the appropriate subnet)
			Name:           fmt.Sprintf("migrated-%s", server.MachineId),      // Set MachineId to identify the source server
			RootDiskType:   "",                                                // Set "" or default to use CSP's default
			RootDiskSize:   "50",                                              // Set 50 GB as a default value
			SshKeyId:       skeletonVmInfra.TargetSshKey.Name,                 // Set the SSH key ID
			VmUserName:     "",                                                // TBD: Set the VM user name if needed
			VmUserPassword: "",                                                // TBD
			SubGroupSize:   "",                                                // TBD
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
		skeletonSubgroupList = append(skeletonSubgroupList, tempCreateSubGroupReq)
	}

	// 4. Recommend security groups with removing duplicates,
	// and set the recommended security groups to the skeleton SubGroup List
	var deduplicatedSecurityGroupList = []cloudmodel.SecurityGroupReq{}
	for i, server := range srcInfra.Servers {

		// Generete security group from the server's firewall rules (or firewall table)
		recommendedSg, err := RecommendSecurityGroup(csp, region, server)
		if err != nil {
			log.Warn().Msgf("failed to recommend security group for server %s: %v", server.MachineId, err)
		}

		// Check duplicates and append the recommended security groups
		exists, _, existingSg := containSg(deduplicatedSecurityGroupList, recommendedSg)
		if !exists {
			// If the security group does not exist, set a name to indicate a dependency between resources.
			recommendedSg.Name = fmt.Sprintf("mig-sg-%02d", len(deduplicatedSecurityGroupList)+1)
			recommendedSg.ConnectionName = fmt.Sprintf("%s-%s", csp, region)
			recommendedSg.Description = fmt.Sprintf("Recommended security group for %s", server.MachineId) // Set MachineId to identify the source server

			// * Set name to indicate a dependency between resources.
			recommendedSg.VNetId = skeletonVmInfra.TargetVNet.Name // Set the vNet ID to the security group

			// Set the security group to the response body
			deduplicatedSecurityGroupList = append(deduplicatedSecurityGroupList, recommendedSg)
		} else {
			recommendedSg = existingSg
		}

		// * Set the security group ID to the skeleton SubGroup List
		skeletonSubgroupList[i].SecurityGroupIds = []string{recommendedSg.Name}
	}

	/*
	 *
	 */

	// 5. Recommend the compatible pairs of VM specs and OS images with removing duplicates,
	// Note: Don't need to register specs and OS images.
	var compatiblePairsForEachServer = make([][]CompatibleSpecImagePair, len(srcInfra.Servers))

	// Find compatible pairs of VM specs and OS images for servers
	for i, server := range srcInfra.Servers {

		// Lookup the appropriate VM specs for the server
		recommendedVmSpecInfoList, _, err := RecommendVmSpecs(csp, region, server, limitSpecs)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM specs for server %s: %v", server.MachineId, err)
		}

		// Lookup the appropriate VM OS images for the server
		recommendedVmOsImageInfoList, err := RecommendVmOsImages(csp, region, server, limitImages)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM OS images for server %s: %v", server.MachineId, err)
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

		if len(recommendedVmSpecInfoList) == 0 || len(recommendedVmOsImageInfoList) == 0 {
			log.Warn().Msgf("no recommended VM specs or OS images found for server %s", server.MachineId)
		} else {
			// Find compatible VM spec and image pairs
			compatiblePairsForEachServer[i], err = FindCompatibleVmSpecAndImagePairs(recommendedVmSpecInfoList, recommendedVmOsImageInfoList, csp)
			if err != nil {
				log.Warn().Msgf("failed to find compatible spec-image pair for server %s: %v", server.MachineId, err)
				// Use fallback selection (first spec, first image)
			} else {
				// Log details about found compatible pairs for this server
				log.Debug().
					Str("machineId", server.MachineId).
					Int("serverIndex", i).
					Int("compatiblePairsCount", len(compatiblePairsForEachServer[i])).
					Msg("Found compatible pairs for server")

				// Log first few pairs for debugging
				loggingLimit := 3
				for pairIdx := 0; pairIdx < len(compatiblePairsForEachServer[i]) && pairIdx < loggingLimit; pairIdx++ {
					pair := compatiblePairsForEachServer[i][pairIdx]
					log.Debug().
						Str("machineId", server.MachineId).
						Int("pairIndex", pairIdx).
						Str("specId", pair.Spec.Id).
						Str("specName", pair.Spec.CspSpecName).
						Str("imageId", pair.Image.Id).
						Str("imageName", pair.Image.CspImageName).
						Msg("Compatible pair details")
				}
			}
		}
	}

	var deduplicatedVmSpecList = []cloudmodel.SpecInfo{}
	var deduplicatedVmOsImageList = []cloudmodel.ImageInfo{}

	// Calculate the actual limit based on the maximum number of compatible pairs across all servers
	actualLimit := 0
	for _, compatiblePairs := range compatiblePairsForEachServer {
		if len(compatiblePairs) > actualLimit {
			actualLimit = len(compatiblePairs)
		}
	}

	// Use the smaller value between requested limit and available pairs
	if actualLimit > limit {
		actualLimit = limit
	}

	log.Debug().
		Int("requestedLimit", limit).
		Int("maxAvailablePairs", actualLimit).
		Msg("Determined actual candidate limit")

	// For each candidate up to the actual limit
	for i := 0; i < actualLimit; i++ {

		// Create a copy of the skeleton SubGroup List
		tempSubGroupList := make([]cloudmodel.CreateSubGroupReq, len(skeletonSubgroupList))
		copy(tempSubGroupList, skeletonSubgroupList)

		var selectedVmSpec cloudmodel.SpecInfo
		var selectedVmOsImage cloudmodel.ImageInfo

		// For each server, select the i-th compatible pair of VM spec and OS image
		for j, server := range srcInfra.Servers {

			// Select compatible pairs for the j-th server
			compatiblePairs := compatiblePairsForEachServer[j]
			if len(compatiblePairs) == 0 {
				log.Warn().Msgf("no compatible VM spec and OS image pairs found for server %s", server.MachineId)
				continue
			}

			// If the i-th pair exists, select it; otherwise skip this server for this candidate
			var pair CompatibleSpecImagePair
			if i < len(compatiblePairs) {
				pair = compatiblePairs[i]
			} else {
				log.Warn().Msgf("candidate %d: server %s has only %d pairs available, skipping this server for this candidate", i+1, server.MachineId, len(compatiblePairs))
				continue
			}

			selectedVmSpec = pair.Spec
			selectedVmOsImage = pair.Image

			// Log candidate and spec selection details
			log.Debug().
				Str("machineId", server.MachineId).
				Int("candidateIndex", i).
				Int("serverIndex", j).
				Int("pairIndex", i).
				Int("totalPairsForServer", len(compatiblePairs)).
				Str("selectedSpecId", selectedVmSpec.Id).
				Str("selectedSpecName", selectedVmSpec.CspSpecName).
				Str("selectedImageId", selectedVmOsImage.Id).
				Str("selectedImageName", selectedVmOsImage.CspImageName).
				Msg("Selected spec-image pair for candidate")

			// Log CPU comparison
			log.Debug().
				Str("machineId", server.MachineId).
				Str("specCspName", selectedVmSpec.CspSpecName).
				Str("specId", selectedVmSpec.Id).
				Uint32("originalCPUs", server.CPU.Cpus).
				Uint32("recommendedVCPU", uint32(selectedVmSpec.VCPU)).
				Msg("CPU comparison")

			// Log Memory comparison
			log.Trace().
				Str("machineId", server.MachineId).
				Str("specCspName", selectedVmSpec.CspSpecName).
				Str("specId", selectedVmSpec.Id).
				Uint32("originalMemoryGB", uint32(server.Memory.TotalSize)).
				Float32("recommendedMemoryGiB", selectedVmSpec.MemoryGiB).
				Msg("Memory comparison")

			// Log OS comparison
			log.Trace().
				Str("machineId", server.MachineId).
				Str("imageCspName", selectedVmOsImage.CspImageName).
				Str("imageId", selectedVmOsImage.Id).
				Str("originalOS", server.OS.Name+" "+server.OS.Version).
				Str("recommendedOSImage", selectedVmOsImage.CspImageName).
				Msg("OS comparison")

			// * Set the selected spec and image IDs to the corresponding SubGroup
			tempSubGroupList[j].SpecId = selectedVmSpec.Id
			tempSubGroupList[j].ImageId = selectedVmOsImage.Id

			// Check duplicates and append the recommended VM specs
			// * Note: Use the name of the VM spec managed by Tumblebug
			exists := false
			// If the recommended VM spec already exists in the list, select the existing spec
			for _, vmSpec := range deduplicatedVmSpecList {
				if vmSpec.CspSpecName == selectedVmSpec.CspSpecName {
					exists = true
					selectedVmSpec = vmSpec
					break
				}
			}
			if !exists {
				deduplicatedVmSpecList = append(deduplicatedVmSpecList, selectedVmSpec)
			}

			// Check duplicates and append the recommended VM OS images
			// * Note: Use the name of the VM OS image managed by Tumblebug
			log.Debug().Msgf("selectedVmOsImage: %+v", selectedVmOsImage)
			exists = false
			// If the recommended VM OS image already exists in the list, select the existing OS image
			for _, vmOsImage := range deduplicatedVmOsImageList {
				if vmOsImage.CspImageName == selectedVmOsImage.CspImageName {
					exists = true
					selectedVmOsImage = vmOsImage
					break
				}
			}
			if !exists {
				deduplicatedVmOsImageList = append(deduplicatedVmOsImageList, selectedVmOsImage)
			}
		}

		// Create a candidate infrastructure based on skeleton and current tempSubGroupList
		candidateInfra := skeletonVmInfra
		candidateInfra.TargetVmInfra.SubGroups = tempSubGroupList
		candidateInfra.TargetVmSpecList = deduplicatedVmSpecList
		candidateInfra.TargetVmOsImageList = deduplicatedVmOsImageList
		candidateInfra.TargetSecurityGroupList = deduplicatedSecurityGroupList
		candidateInfra.TargetVmInfra.Description = fmt.Sprintf("candidate #%d: recommended VMs comprising multi-cloud infrastructure", i+1)

		// Calculate overall quality for the candidate infrastructure
		overallQuality, overallQualityDesc, _, _ := calculateCandidateQuality(tempSubGroupList, srcInfra, deduplicatedVmSpecList, deduplicatedVmOsImageList)

		// Set the status and description based on the overall quality
		candidateInfra.Status = overallQuality
		candidateInfra.Description = fmt.Sprintf("Infrastructure recommendation for candidate %d - %s", i+1, overallQualityDesc)

		recommendedVmInfraCandidates = append(recommendedVmInfraCandidates, candidateInfra)
	}

	/*
	 * [Output]
	 */
	log.Trace().Msgf("recommended infrastructure candidates: %+v", recommendedVmInfraCandidates)

	return recommendedVmInfraCandidates, nil
}

// calculateComprehensiveQualityScore calculates comprehensive quality score using Manhattan distance
// Returns: (weightedScore, specScore, imageScore)
func calculateComprehensiveQualityScore(server onpremmodel.ServerProperty, vmSpec cloudmodel.SpecInfo, vmImage cloudmodel.ImageInfo) (float64, float64, float64) {
	// Log server and VM specifications for comparison
	log.Debug().
		Str("machineId", server.MachineId).
		Uint32("serverCPUs", server.CPU.Cpus).
		Uint32("serverThreads", server.CPU.Threads).
		Uint32("serverMemoryGB", uint32(server.Memory.TotalSize)).
		Str("serverArchitecture", server.CPU.Architecture).
		Str("serverOS", fmt.Sprintf("%s %s %s", server.OS.Name, server.OS.Version, server.OS.VersionCodename)).
		Uint32("vmSpecVCPU", uint32(vmSpec.VCPU)).
		Float32("vmSpecMemoryGiB", vmSpec.MemoryGiB).
		Str("vmSpecName", vmSpec.CspSpecName).
		Str("vmImageName", vmImage.CspImageName).
		Str("vmImageOS", fmt.Sprintf("%s %s %s", vmImage.OSType, vmImage.OSDistribution, vmImage.OSArchitecture)).
		Msg("Server and VM specification comparison")

	// 1. Calculate Manhattan distance between server and VM spec resources
	// Server: vCPUs = CPUs * Threads (fallback to CPUs if Threads is 0), Memory = TotalSize
	serverThreads := server.CPU.Threads
	if serverThreads == 0 {
		serverThreads = 1 // Default to 1 thread per CPU if not specified
	}
	serverVCPUs := float64(server.CPU.Cpus * serverThreads)
	serverMemoryGB := float64(server.Memory.TotalSize)

	// VM Spec: vCPUs, Memory in GiB
	vmSpecVCPUs := float64(vmSpec.VCPU)
	vmSpecMemoryGiB := float64(vmSpec.MemoryGiB)

	// Calculate Manhattan distance: |server_vcpus - vmspec_vcpus| + |server_memory - vmspec_memory|
	cpuDistance := absFloat64(serverVCPUs - vmSpecVCPUs)
	memoryDistance := absFloat64(serverMemoryGB - vmSpecMemoryGiB)

	// Manhattan distance (absolute differences)
	manhattanDistance := cpuDistance + memoryDistance

	// Normalize to percentage: convert Manhattan distance to similarity score (0-100%)
	// Perfect match (distance=0) = 100% similarity, larger distances = lower similarity
	// Use the maximum of server and VM spec magnitude for better scaling
	serverMagnitude := serverVCPUs + serverMemoryGB
	vmSpecMagnitude := vmSpecVCPUs + vmSpecMemoryGiB
	maxMagnitude := serverMagnitude
	if vmSpecMagnitude > maxMagnitude {
		maxMagnitude = vmSpecMagnitude
	}
	if maxMagnitude == 0 {
		maxMagnitude = 1.0 // Avoid division by zero
	}

	// Calculate similarity: 100% - (distance/maxMagnitude * 100%)
	specSimilarityScore := 100.0 - ((manhattanDistance / maxMagnitude) * 100.0)

	// Ensure similarity score is within 0-100% range
	if specSimilarityScore < 0 {
		specSimilarityScore = 0
	} else if specSimilarityScore > 100 {
		specSimilarityScore = 100
	}

	// 2. Calculate image similarity quality (lower similarity = higher difference score)
	imageQualityScore := calculateImageQualityScore(server, vmImage)
	imageQualityScore = imageQualityScore * 100.0 // Convert to percentage

	// 3. Combine spec and image quality with weights
	// Weight: 70% for spec quality, 30% for image quality
	comprehensiveScore := (specSimilarityScore * 0.7) + (imageQualityScore * 0.3)
	log.Debug().
		Str("machineId", server.MachineId).
		Msgf("Quality Score - total weighted score: %.1f%%, spec score: %.1f%%, image score: %.1f%%",
			comprehensiveScore, specSimilarityScore, imageQualityScore)
	return comprehensiveScore, specSimilarityScore, imageQualityScore
}

// calculateImageQualityScore calculates image quality score based on OS similarity
func calculateImageQualityScore(server onpremmodel.ServerProperty, vmImage cloudmodel.ImageInfo) float64 {
	// Set keywords and delimiters similar to existing image recommendation logic
	keywords, kwDelimiters, imgDelimiters := SetKeywordsAndDelimeters(server)

	// Create image keywords for similarity calculation
	vmImgKeywords := fmt.Sprintf("%s %s %s %s",
		vmImage.OSType,
		vmImage.OSArchitecture,
		vmImage.OSDiskType,
		vmImage.OSDistribution,
	)

	log.Debug().
		Str("machineId", server.MachineId).
		Str("osId", server.OS.ID).
		Str("osVersionId", server.OS.VersionID).
		Str("osVersionCodename", server.OS.VersionCodename).
		Str("cpuArchitecture", server.CPU.Architecture).
		Str("rootDiskType", server.RootDisk.Type).
		Str("vmImageOSType", vmImage.OSType).
		Str("vmImageOSArchitecture", string(vmImage.OSArchitecture)).
		Str("vmImageOSDiskType", vmImage.OSDiskType).
		Str("vmImageOSDistribution", vmImage.OSDistribution).
		Msg("Image similarity input details")

	// Calculate similarity score (0.0 to 1.0, where 1.0 is perfect match)
	similarityScore := similarity.CalcResourceSimilarity(keywords, kwDelimiters, vmImgKeywords, imgDelimiters)

	log.Debug().
		Str("machineId", server.MachineId).
		Str("serverKeywords", keywords).
		Str("imageKeywords", vmImgKeywords).
		Float64("similarity", similarityScore).
		Msg("Image quality calculation")

	return similarityScore
}

// absFloat64 returns the absolute value of a float64
func absFloat64(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// determineRecommendationQuality determines if the recommendation is "best effort" or "alternative" based on similarity score
func determineRecommendationQuality(similarityScore, threshold float64) string {
	if similarityScore >= threshold {
		return "best-effort"
	}
	return "alternative"
}

// calculateCandidateQuality calculates overall quality for a candidate infrastructure
func calculateCandidateQuality(tempSubGroupList []cloudmodel.CreateSubGroupReq, srcInfra onpremmodel.OnpremInfra, deduplicatedVmSpecList []cloudmodel.SpecInfo, deduplicatedVmOsImageList []cloudmodel.ImageInfo) (string, string, int, int) {
	// Calculate overall quality based on all servers - only "best-effort" if ALL servers are best-effort
	var overallQuality string
	var overallQualityDesc string
	var bestOptionCount int
	var alternativeCount int

	if len(tempSubGroupList) > 0 {
		// Calculate individual quality for each server
		var totalQualityScore float64
		var totalSpecScore float64
		var totalImageScore float64
		var validServerCount int

		for j, subGroup := range tempSubGroupList {
			if subGroup.SpecId != "" && subGroup.ImageId != "" {
				// Find the corresponding spec and image for quality calculation
				server := srcInfra.Servers[j]

				// Find the spec and image from deduplicated lists
				var selectedSpec cloudmodel.SpecInfo
				var selectedImage cloudmodel.ImageInfo

				for _, spec := range deduplicatedVmSpecList {
					if spec.Id == subGroup.SpecId {
						selectedSpec = spec
						break
					}
				}

				for _, image := range deduplicatedVmOsImageList {
					if image.Id == subGroup.ImageId {
						selectedImage = image
						break
					}
				}

				if selectedSpec.Id != "" && selectedImage.Id != "" {
					qualityScore, specScore, imageScore := calculateComprehensiveQualityScore(server, selectedSpec, selectedImage)
					totalQualityScore += qualityScore
					totalSpecScore += specScore
					totalImageScore += imageScore
					validServerCount++

					// Check if this individual server is best-effort
					// For similarity score: higher score = better match, use >= threshold
					var individualQuality string
					if qualityScore == 100.0 {
						// Perfect match - definitely best-effort
						individualQuality = "best-effort"
					} else {
						// Apply threshold for non-perfect matches (70% similarity threshold)
						individualQuality = determineRecommendationQuality(qualityScore, 70.0) // 70% similarity threshold
					}
					if individualQuality == "best-effort" {
						bestOptionCount++
					} else {
						alternativeCount++
					}

					log.Trace().
						Str("machineId", server.MachineId).
						Float64("qualityScore", qualityScore).
						Str("individualQuality", individualQuality).
						Str("specId", selectedSpec.Id).
						Str("imageId", selectedImage.Id).
						Msg("Individual server quality assessment")
				}
			}
		}

		if validServerCount > 0 {
			averageQualityScore := totalQualityScore / float64(validServerCount)

			// Overall quality is "best-effort" only if ALL valid servers are best-effort
			if bestOptionCount == validServerCount {
				overallQuality = "best-effort"
			} else {
				overallQuality = "alternative"
			}

			overallQualityDesc = fmt.Sprintf("VM total: %d, Best effort option: %d, Alternatives: %d (avg score: %.1f/10.0)",
				validServerCount, bestOptionCount, alternativeCount, averageQualityScore/10.0)

			log.Info().
				Float64("averageQualityScore", averageQualityScore).
				Str("overallQuality", overallQuality).
				Int("bestOptionCount", bestOptionCount).
				Int("alternativeCount", alternativeCount).
				Int("validServerCount", validServerCount).
				Msg("Overall candidate quality assessment")
		} else {
			overallQuality = "nothing-to-recommend"
			overallQualityDesc = "VM total: 0, Best effort option: 0, Alternatives: 0 (Unable to assess quality - no valid spec-image pairs)"
			bestOptionCount = 0
			alternativeCount = 0
		}
	} else {
		// No SubGroups available
		overallQuality = "nothing-to-recommend"
		overallQualityDesc = "VM total: 0, Best effort option: 0, Alternatives: 0 (no recommendations available)"
		bestOptionCount = 0
		alternativeCount = 0
	}

	return overallQuality, overallQualityDesc, bestOptionCount, alternativeCount
}
