package recommendation

import (
	"fmt"
	"strings"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/csp"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"

	"github.com/rs/zerolog/log"
)

// supportedInfraCSPs defines the CSPs supported for computing infrastructure recommendation.
// KT and OpenStack are currently not supported for computing infrastructure.
var supportedInfraCSPs = map[string]bool{
	csp.AWS:       true,
	csp.Azure:     true,
	csp.GCP:       true,
	csp.Alibaba:   true,
	csp.Tencent:   true,
	csp.IBM:       true,
	csp.NCP:       true,
	csp.NHN:       true,
	csp.OpenStack: false,
	csp.KT:        false,
}

func isSupportedCSP(cspName string) bool {
	return supportedInfraCSPs[strings.ToLower(strings.TrimSpace(cspName))]
}

// getCspMinRootDiskSizeGB returns the CSP-specific safe minimum root disk size in GB.
// Values are conservative to prevent VM creation failures near the lower bound.
//
// Notes: IBM/NCP use fixed sizes (API ignores the field); Azure derives size from the image.
// Ref: https://github.com/cloud-barista/cb-spider/wiki/API-and-features-of-VM-rootfs-volume-configuration
//
// TODO: Replace hardcoded values with dynamic lookup via CB-Spider's GET /cloudos/metainfo/{CloudOSName}
// (returns RootDiskSize as a string enum list). Consider caching at startup to avoid per-request overhead.
// Ref: https://github.com/cloud-barista/cb-spider/blob/master/cloud-driver-libs/cloudos_meta.yaml
func getCspMinRootDiskSizeGB(cspName string) int {
	switch strings.ToLower(cspName) {
	case csp.IBM:
		// Fixed at 100 GB by default; root disk size is not configurable via API.
		return 100
	case csp.NCP, "ncpvpc":
		// Fixed at 50 GB for Linux; root disk size cannot be changed via API or console.
		return 50
	case csp.Tencent:
		// Minimum per cloudos_meta.yaml: CLOUD_PREMIUM|50 GB, CLOUD_SSD|50 GB.
		return 50
	case csp.KT, "ktvpc":
		// Minimum per cloudos_meta.yaml: HDD|50 GB, SSD|50 GB.
		return 50
	case csp.Alibaba:
		// Minimum: cloud_essd|20 GB, cloud_efficiency|20 GB; cloud_auto requires 40 GB.
		return 40
	case csp.Azure:
		// Root disk size depends on the selected image and is not configurable via API.
		return 30
	case csp.NHN:
		// Minimum per cloudos_meta.yaml: General_HDD|20 GB, General_SSD|20 GB.
		return 20
	case "ktclassic":
		// Minimum per cloudos_meta.yaml: HDD|20 GB, SSD|20 GB.
		return 20
	case csp.AWS, csp.GCP:
		// AWS: gp2/gp3 practical OS minimum ~8 GB. GCP: pd-standard minimum 10 GB.
		return 10
	default:
		return 50 // Conservative safe default for unknown or future CSPs
	}
}

func IsValidCspAndRegion(cspName string, region string) (bool, error) {

	isValid := false
	cspLower := strings.ToLower(cspName)
	regionName := strings.ToLower(region)
	supportedCsp := isSupportedCSP(cspLower)

	if !supportedCsp {
		err := fmt.Errorf("not supported yet (provider: %s)", cspLower)
		log.Warn().Msgf("%s", err.Error())
		return isValid, err
	}

	// Check if the region is valid for the specified CSP
	_, err := tbclient.NewSession().ReadRegionInfo(cspLower, regionName)
	if err != nil {
		log.Warn().Msgf("failed to read region info for CSP %s and region %s: %v", cspLower, regionName, err)
		return isValid, err
	}

	isValid = true

	return isValid, nil
}

// RecommendInfraWithDefaults recommends an appropriate target infrastructure for cloud migration using defaults
func RecommendInfraWithDefaults(desiredCsp string, desiredRegion string, srcInfra onpremmodel.OnpremInfra) (cloudmodel.RecommendedInfraDynamicList, error) {
	srcInfra = SanitizeSourceInfra(srcInfra)

	// var emptyResp RecommendedVmInfraInfoList
	var recommendedVmInfraInfoList cloudmodel.RecommendedInfraDynamicList

	// TODO: To be updated, a user will input the desired number of recommended VMs
	var defaultSpecsLimit int = GetDefaultSpecsLimit()

	// Initialize the response body
	recommendedVmInfraInfoList = cloudmodel.RecommendedInfraDynamicList{
		Description:     "This is a list of recommended target infrastructures. Please review and use them.",
		Count:           0,
		TargetInfraList: []cloudmodel.RecommendedInfraDynamic{},
	}

	// // Set VM info
	// recommendedVm := tbmodel.TbVmDynamicReq{
	// 	ConnectionName: "",
	// 	CommonImage:    "", // Lookup and set an appropriate VM OS image
	// 	CommonSpec:     "", // Lookup and set an appropriate VM spec
	// 	Description:    "a recommended virtual machine",
	// 	Name:           fmt.Sprintf("migrated-%s", node.MachineId),
	// 	RootDiskSize:   "", // TBD
	// 	RootDiskType:   "", // TBD
	// 	NodeGroupSize:  "",
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
	for _, node := range srcInfra.Nodes {

		// Lookup the appropriate VM specs for the node
		vmSpecList, _, err := RecommendNodeSpecs(desiredCsp, desiredRegion, node, defaultSpecsLimit)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM specs for node %s: %v", node.MachineId, err)
			continue
		}

		// Lookup the appropriate VM OS images for the node
		vmOsImageIdList := []string{}
		for range vmSpecList {
			osImgId, err := RecommendVmOsImageId(desiredCsp, desiredRegion, node)
			if err != nil {
				log.Warn().Msgf("failed to recommend VM OS image for node %s: %v", node.MachineId, err)
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
	// 	log.Debug().Msgf("the number of recommended VM specs and OS images for node %d: %d", i, len(vmInfoList))
	// }
	// log.Debug().Msgf("recommended VM specs and OS images: %+v", recommendedVmInfoList)

	/*
	 * [Output]
	 */

	// Transpose the matrix to change from "VM recommendations per node" to "servers per VM recommendation".
	// Before: [Server1's VM recommendations, Server2's VM recommendations, ...]
	// After: [Recommendation1 for all servers, Recommendation2 for all servers, ...]
	transposed := transposeMatrix(recommendedVmInfoList)
	// log.Debug().Msgf("transposed recommended VM specs and OS images: %+v", transposed)

	// Build response body which includes multiple recommended infrastructures
	recommenedVmInfraInfoList := []cloudmodel.RecommendedInfraDynamic{}

	for i, vmInfoList := range transposed {

		tempVmInfraInfo := cloudmodel.RecommendedInfraDynamic{
			Status:      string(NothingRecommended),
			Description: "This is a recommended target infrastructure.",
			TargetInfra: cloudmodel.InfraDynamicReq{
				Name:        fmt.Sprintf("migrated-%02d", i),
				Description: "a recommended multi-cloud infrastructure",
				NodeGroups:  []cloudmodel.CreateNodeGroupDynamicReq{},
			},
		}

		for j, nodegroupInfo := range vmInfoList {
			tempCreateNodegroupReq := cloudmodel.CreateNodeGroupDynamicReq{
				ConnectionName: fmt.Sprintf("%s-%s", desiredCsp, desiredRegion),
				ImageId:        nodegroupInfo.vmOsImageId,
				SpecId:         nodegroupInfo.vmSpecId,
				Description:    "a recommended virtual machine",
				Name:           fmt.Sprintf("migrated-%s", srcInfra.Nodes[j].MachineId), // Set MachineId to identify the source node
				RootDiskSize:   0,                                                       // TBD
				RootDiskType:   "",                                                      // TBD
				NodeGroupSize:  1,                                                       // TBD
			}
			tempVmInfraInfo.TargetInfra.NodeGroups = append(tempVmInfraInfo.TargetInfra.NodeGroups, tempCreateNodegroupReq)
		}

		status := checkOverallNodeGroupStatus(tempVmInfraInfo.TargetInfra.NodeGroups)
		tempVmInfraInfo.Status = status
		if status == string(NothingRecommended) {
			tempVmInfraInfo.Description = "Could not find appropriate Node (VM)s."
		} else if status == string(FullyRecommended) {
			tempVmInfraInfo.Description = "Target infra is recommended."
		} else {
			tempVmInfraInfo.Description = "Some Node (VM)s are recommended. Please check and fill the required information."
		}

		recommenedVmInfraInfoList = append(recommenedVmInfraInfoList, tempVmInfraInfo)
	}

	// Assign the target infrastructure list to the response
	recommendedVmInfraInfoList.TargetInfraList = recommenedVmInfraInfoList
	recommendedVmInfraInfoList.Count = len(recommenedVmInfraInfoList)

	log.Trace().Msgf("the recommended infra info: %+v", recommendedVmInfraInfoList)

	return recommendedVmInfraInfoList, nil
}

// RecommendInfra recommends an appropriate target infrastructure for cloud migration
func RecommendInfra(desiredCsp string, desiredRegion string, srcInfra onpremmodel.OnpremInfra) (cloudmodel.RecommendedInfra, error) {
	srcInfra = SanitizeSourceInfra(srcInfra)

	// var emptyResp RecommendedVmInfra
	var recommendedVmInfra cloudmodel.RecommendedInfra

	// TODO: To be updated, a user will input the desired number of recommended VMs
	var limitSpecs int = GetDefaultSpecsLimit()
	var limitImages int = GetDefaultImagesLimit()

	// Initialize the response body
	recommendedVmInfra = cloudmodel.RecommendedInfra{
		Description: "This is a list of recommended target infrastructures. Please review and use them.",
		Status:      "",
		TargetCloud: cloudmodel.CloudProperty{
			Csp:    desiredCsp,
			Region: desiredRegion,
		},
		TargetInfra: cloudmodel.InfraReq{
			Name:        "infra101",
			Description: "a recommended multi-cloud infrastructure",
			NodeGroups:  []cloudmodel.CreateNodeGroupReq{},
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

	if len(recommendedVNetInfoList) > 0 {
		recommendedVmInfra.TargetVNet = recommendedVNetInfoList[0]
		recommendedVmInfra.TargetVNet.Name = "mig-vnet-01"
		recommendedVmInfra.TargetVNet.Description = "a recommended vNet for migration"
		for i := range recommendedVmInfra.TargetVNet.SubnetInfoList {
			recommendedVmInfra.TargetVNet.SubnetInfoList[i].Name = fmt.Sprintf("mig-subnet-%02d", i+1)
			recommendedVmInfra.TargetVNet.SubnetInfoList[i].Description = "a recommended subnet for migration"
		}
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
	var recommendedNodegroupList = []cloudmodel.CreateNodeGroupReq{}
	var recommendedVmSpecList = []cloudmodel.SpecInfo{}
	var recommendedVmOsImageList = []cloudmodel.ImageInfo{}
	var recommendedSecurityGroupList = []cloudmodel.SecurityGroupReq{}

	firstSubnetId := ""
	if len(recommendedVmInfra.TargetVNet.SubnetInfoList) > 0 {
		firstSubnetId = recommendedVmInfra.TargetVNet.SubnetInfoList[0].Name
	}

	for i, node := range srcInfra.Nodes {

		/*
		 * Recommend VM specs, OS images, and security groups
		 */

		// Lookup the appropriate VM specs for the node
		recommendedVmSpecInfoList, _, err := RecommendNodeSpecs(csp, region, node, limitSpecs)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM specs for node %s: %v", node.MachineId, err)
		}

		// Lookup the appropriate VM OS images for the node
		// recommendedVmOsImageInfo, err := RecommendVmOsImage(csp, region, node)
		recommendedVmOsImageInfoList, err := RecommendVmOsImages(csp, region, node, limitImages)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM OS images for node %s: %v", node.MachineId, err)
		}

		// Generete security group from the node's firewall rules (or firewall table)
		recommendedSg, err := RecommendSecurityGroup(csp, region, node)
		if err != nil {
			log.Warn().Msgf("failed to recommend security group for node %s: %v", node.MachineId, err)
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
		resolvedCspImageName := "" // set only when review resolves a newer image than the DB value
		suggestedSystemDisk := ""  // set when specImagePairReview confirms an available disk type
		if len(recommendedVmSpecInfoList) == 0 || len(recommendedVmOsImageInfoList) == 0 {
			log.Warn().Msgf("no recommended VM specs or OS images found for node %s", node.MachineId)
		} else {

			// * Note: (opinion) Find multiple compatible pairs and use them as needed in the later process
			// Find compatible spec and image pair
			tempSelectedVmSpec, tempSelectedVmOsImage, err := FindCompatibleSpecAndImage(recommendedVmSpecInfoList, recommendedVmOsImageInfoList, csp)
			if err != nil {
				log.Warn().Msgf("failed to find compatible spec-image pair for node %s: %v", node.MachineId, err)
				// Use fallback selection (first spec, first image)
			} else {
				selectedVmSpec = tempSelectedVmSpec
				selectedVmOsImage = tempSelectedVmOsImage

				// Resolve the latest CSP image name and confirm available system disk.
				precheck, reviewErr := PreflightCheckCspProvisioning(
					selectedVmSpec.Id, selectedVmOsImage.Id, selectedVmOsImage.CspImageName, "",
				)
				if reviewErr != nil {
					log.Warn().Err(reviewErr).Msgf("preflight check failed for node %s; using cached image", node.MachineId)
				} else if !precheck.IsAvailable {
					log.Warn().Msgf("image %s may be unavailable for node %s; using cached image", selectedVmOsImage.Id, node.MachineId)
				} else {
					if precheck.SuggestedSystemDisk != "" {
						log.Debug().Msgf("using suggested system disk for node %s: %s", node.MachineId, precheck.SuggestedSystemDisk)
						suggestedSystemDisk = precheck.SuggestedSystemDisk
					}
					if precheck.ResolvedCspImageName != selectedVmOsImage.CspImageName {
						log.Info().Msgf("image resolved to latest for node %s: %s -> %s", node.MachineId, selectedVmOsImage.CspImageName, precheck.ResolvedCspImageName)
						selectedVmOsImage.CspImageName = precheck.ResolvedCspImageName
						resolvedCspImageName = precheck.ResolvedCspImageName
					}
				}

				// Log CPU comparison
				log.Debug().
					Str("machineId", node.MachineId).
					Str("specCspName", selectedVmSpec.CspSpecName).
					Str("specId", selectedVmSpec.Id).
					Uint32("originalCPUs", node.CPU.Cpus).
					Uint32("recommendedVCPU", uint32(selectedVmSpec.VCPU)).
					Msg("CPU comparison")

				// Log Memory comparison
				log.Debug().
					Str("machineId", node.MachineId).
					Str("specCspName", selectedVmSpec.CspSpecName).
					Str("specId", selectedVmSpec.Id).
					Uint32("originalMemoryGB", uint32(node.Memory.TotalSize)).
					Float32("recommendedMemoryGiB", selectedVmSpec.MemoryGiB).
					Msg("Memory comparison")

				// Log OS comparison
				log.Debug().
					Str("machineId", node.MachineId).
					Str("imageCspName", selectedVmOsImage.CspImageName).
					Str("imageId", selectedVmOsImage.Id).
					Str("originalOS", node.OS.Name+" "+node.OS.Version).
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
			recommendedSg.Description = fmt.Sprintf("Recommended security group for %s", node.MachineId) // Set MachineId to identify the source node

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
		// TODO: Select a subnet by the node's network information (for now, select the first one)

		// Ref: https://github.com/cloud-barista/cb-spider/blob/master/cloud-driver-libs/cloudos_meta.yaml
		// Note: "TYPE1" for RootDiskType is the first in the list
		// - AWS: ["standard", "gp2", "gp3"],
		// - Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"],
		// - GCP: [ "pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"],
		// - ALIBABA: ["cloud_auto", "cloud_essd", "cloud_efficiency", "cloud", "cloud_ssd"],
		// - TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
		// - NCP: ["HDD"]
		// - NHN: ["General_HDD", "General_SSD"]
		// - KT: ["HDD", "SSD"]

		// * Set names to indicate a dependency between resources.
		// Use the source disk size as the base, with a CSP-specific minimum floor to avoid
		// over-provisioning and keep costs low (e.g., AWS/GCP support as low as 10 GB).
		rootDiskSize := max(int(node.RootDisk.TotalSize), getCspMinRootDiskSizeGB(csp))
		tempCreateNodeGroupReq := cloudmodel.CreateNodeGroupReq{
			ConnectionName:   fmt.Sprintf("%s-%s", csp, region),
			Description:      fmt.Sprintf("a recommended virtual machine %02d for %s", i+1, node.MachineId), // Set MachineId to identify the source node
			SpecId:           selectedVmSpec.Id,
			ImageId:          selectedVmOsImage.Id,
			CspImageName:     resolvedCspImageName, // Set only when review resolved a newer image; TumbleBug sends this to Spider instead of looking up via ImageId.
			VNetId:           recommendedVmInfra.TargetVNet.Name,
			SubnetId:         firstSubnetId,
			SecurityGroupIds: []string{recommendedSg.Name},                         // Set the security group ID
			Name:             fmt.Sprintf("migrated-%s", node.MachineId),           // Set MachineId to identify the source node
			RootDiskType:     suggestedSystemDisk,                                  // Confirmed-available disk type from specImagePairReview; empty → CSP default
			RootDiskSize:     rootDiskSize,                                         // max(source disk size, CSP minimum) to keep costs low
			SshKeyId:         recommendedVmInfra.TargetSshKey.Name,                 // Set the SSH key ID
			NodeUserName:     "",                                                   // TBD: Set the VM user name if needed
			NodeUserPassword: "",                                                   // TBD
			NodeGroupSize:    1,                                                    // TBD
			Label: map[string]string{
				lineage.LabelSourceMachineIds: node.MachineId,
			},
		}

		// Append the VM request to the list
		recommendedNodegroupList = append(recommendedNodegroupList, tempCreateNodeGroupReq)
	}

	/*
	 * [Output]
	 */
	recommendedVmInfra.TargetInfra.NodeGroups = recommendedNodegroupList
	recommendedVmInfra.TargetSpecList = recommendedVmSpecList
	recommendedVmInfra.TargetOsImageList = recommendedVmOsImageList
	recommendedVmInfra.TargetSecurityGroupList = recommendedSecurityGroupList

	log.Trace().Msgf("the recommended infra info: %+v", recommendedVmInfra)

	return recommendedVmInfra, nil
}

// RecommendInfraCandidates recommends appropriate target infrastructure candidates for cloud migration
func RecommendInfraCandidates(desiredCsp string, desiredRegion string, srcInfra onpremmodel.OnpremInfra, limit int, minMatchRate float64) ([]cloudmodel.RecommendedInfra, error) {
	srcInfra = SanitizeSourceInfra(srcInfra)

	// * To recommend multiple infra candidates (i.e., multiple VM spec and OS image combinations),
	// * this function estimates, recommends or just generates vNets, subnets, SSH key pair, and security groups
	// * and then, recommends compatible pairs of VM specs and OS images from the best option to the alternative ones.
	// * All those are corresponding to the source servers.

	// var emptyResp RecommendedVmInfra
	var recommendedVmInfraCandidates []cloudmodel.RecommendedInfra

	// TODO: To be updated, a user will input the desired number of recommended VMs
	var limitSpecs int = GetDefaultSpecsLimit()
	var limitImages int = GetDefaultImagesLimit()

	csp := strings.ToLower(desiredCsp)
	region := strings.ToLower(desiredRegion)

	// Initialize the response body
	skeletonVmInfra := cloudmodel.RecommendedInfra{
		Description: "This is a recommended target infrastructures and resources. Please review and use them.",
		Status:      "",
		TargetCloud: cloudmodel.CloudProperty{
			Csp:    csp,
			Region: region,
		},
		TargetInfra: cloudmodel.InfraReq{
			Name:       "infra101",
			NodeGroups: []cloudmodel.CreateNodeGroupReq{},
			// Description: "Recommended VMs comprising the multi-cloud infrastructure",
		},
	}

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

	if len(recommendedVNetInfoList) > 0 {
		skeletonVmInfra.TargetVNet = recommendedVNetInfoList[0]
		skeletonVmInfra.TargetVNet.Name = "vnet-01"
		skeletonVmInfra.TargetVNet.Description = "a recommended vNet for migration"
		for i := range skeletonVmInfra.TargetVNet.SubnetInfoList {
			skeletonVmInfra.TargetVNet.SubnetInfoList[i].Name = fmt.Sprintf("subnet-%02d", i+1)
			skeletonVmInfra.TargetVNet.SubnetInfoList[i].Description = "a recommended subnet for migration"
		}
	}

	// 2. Recommend(?) SSH key pair
	// var recommendedSshKey = tbmodel.SshKeyReq{}
	// * Set a name to indicate a dependency between resources.
	skeletonVmInfra.TargetSshKey.Name = "sshkey-01"
	skeletonVmInfra.TargetSshKey.ConnectionName = fmt.Sprintf("%s-%s", csp, region)
	skeletonVmInfra.TargetSshKey.Description = "a SSH Key pair for migration (Note - provided ONLY once, MUST be downloaded"

	// 3. Generate a skeleton of NodeGroup List for VMs
	var skeletonNodegroupList = []cloudmodel.CreateNodeGroupReq{}

	firstSubnetId := ""
	if len(skeletonVmInfra.TargetVNet.SubnetInfoList) > 0 {
		firstSubnetId = skeletonVmInfra.TargetVNet.SubnetInfoList[0].Name
	}

	// TODO: Select a subnet by the node's network information (for now, select the first one)
	// Ref: https://github.com/cloud-barista/cb-spider/blob/master/cloud-driver-libs/cloudos_meta.yaml
	// Note: "TYPE1" for RootDiskType is the first in the list
	// - AWS: ["standard", "gp2", "gp3"],
	// - Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"],
	// - GCP: [ "pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"],
	// - ALIBABA: ["cloud_auto", "cloud_essd", "cloud_efficiency", "cloud", "cloud_ssd"],
	// - TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	// - NCP: ["HDD"]
	// - NHN: ["General_HDD", "General_SSD"]
	// - KT: ["HDD", "SSD"]

	// * Set names to indicate a dependency between resources.
	for i, node := range srcInfra.Nodes {
		// * Set names to indicate a dependency between resources.
		// Use the source disk size as the base, with a CSP-specific minimum floor to avoid
		// over-provisioning and keep costs low (e.g., AWS/GCP support as low as 10 GB).
		rootDiskSize := max(int(node.RootDisk.TotalSize), getCspMinRootDiskSizeGB(csp))
		tempCreateNodeGroupReq := cloudmodel.CreateNodeGroupReq{
			ConnectionName:   fmt.Sprintf("%s-%s", csp, region),
			Description:      fmt.Sprintf("Recommended VM %02d for %s", i+1, node.MachineId), // Set MachineId to identify the source node
			VNetId:           skeletonVmInfra.TargetVNet.Name,
			SubnetId:         firstSubnetId,
			Name:             fmt.Sprintf("vm-%s", node.MachineId),              // Set MachineId to identify the source node
			RootDiskType:     "",                                                // Set "" or default to use CSP's default
			RootDiskSize:     rootDiskSize,                                      // max(source disk size, CSP minimum) to keep costs low
			SshKeyId:         skeletonVmInfra.TargetSshKey.Name,                 // Set the SSH key ID
			NodeUserName:     "",                                                // TBD: Set the VM user name if needed
			NodeUserPassword: "",                                                // TBD
			NodeGroupSize:    1,                                                 // Default: 1
			Label: map[string]string{
				lineage.LabelSourceMachineIds: node.MachineId,
			},
		}

		// Append the VM request to the list
		skeletonNodegroupList = append(skeletonNodegroupList, tempCreateNodeGroupReq)
	}

	// 4. Recommend security groups with removing duplicates,
	// and set the recommended security groups to the skeleton NodeGroup List
	var deduplicatedSecurityGroupList = []cloudmodel.SecurityGroupReq{}
	for i, node := range srcInfra.Nodes {

		// Generete security group from the node's firewall rules (or firewall table)
		recommendedSg, err := RecommendSecurityGroup(csp, region, node)
		if err != nil {
			log.Warn().Msgf("failed to recommend security group for node %s: %v", node.MachineId, err)
		}

		// Check duplicates and append the recommended security groups
		exists, _, existingSg := containSg(deduplicatedSecurityGroupList, recommendedSg)
		if !exists {
			// If the security group does not exist, set a name to indicate a dependency between resources.
			recommendedSg.Name = fmt.Sprintf("sg-%02d", len(deduplicatedSecurityGroupList)+1)
			recommendedSg.ConnectionName = fmt.Sprintf("%s-%s", csp, region)
			recommendedSg.Description = fmt.Sprintf("Recommended security group for %s", node.MachineId) // Set MachineId to identify the source node

			// * Set name to indicate a dependency between resources.
			recommendedSg.VNetId = skeletonVmInfra.TargetVNet.Name // Set the vNet ID to the security group

			// Set the security group to the response body
			deduplicatedSecurityGroupList = append(deduplicatedSecurityGroupList, recommendedSg)
		} else {
			recommendedSg = existingSg
		}

		// * Set the security group ID to the skeleton NodeGroup List
		skeletonNodegroupList[i].SecurityGroupIds = []string{recommendedSg.Name}
	}

	/*
	 *
	 */

	// 5. Recommend the compatible pairs of VM specs and OS images with removing duplicates,
	// Note: Don't need to register specs and OS images.
	var compatiblePairsForEachServer = make([][]CompatibleSpecImagePair, len(srcInfra.Nodes))

	// Find compatible pairs of VM specs and OS images for servers
	for i, node := range srcInfra.Nodes {

		// Lookup the appropriate VM specs for the node
		recommendedVmSpecInfoList, _, err := RecommendNodeSpecs(csp, region, node, limitSpecs)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM specs for node %s: %v", node.MachineId, err)
		}

		// Lookup the appropriate VM OS images for the node
		recommendedVmOsImageInfoList, err := RecommendVmOsImages(csp, region, node, limitImages)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM OS images for node %s: %v", node.MachineId, err)
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
			log.Warn().Msgf("no recommended VM specs or OS images found for node %s", node.MachineId)
		} else {
			// Find compatible VM spec and image pairs
			compatiblePairsForEachServer[i], err = FindCompatibleVmSpecAndImagePairs(recommendedVmSpecInfoList, recommendedVmOsImageInfoList, csp)

			// * Uncomment to check specs of compatible pairs
			// tempLoggingLimit := 5
			// for idx := 0; idx < len(compatiblePairsForEachServer[i]) && idx < tempLoggingLimit; idx++ {
			// 	spec := compatiblePairsForEachServer[i][idx].Spec
			// 	log.Debug().Msgf("(temp logging) compatiblePairsForEachServer[%d][%d]: %s (%d vCPU, %.1f GiB Memory)", i, idx, spec.Id, spec.VCPU, spec.MemoryGiB)
			// }

			if err != nil {
				log.Warn().Msgf("failed to find compatible spec-image pair for node %s: %v", node.MachineId, err)
				// Use fallback selection (first spec, first image)
			} else {
				// Log details about found compatible pairs for this node
				log.Debug().
					Str("machineId", node.MachineId).
					Int("serverIndex", i).
					Int("compatiblePairsCount", len(compatiblePairsForEachServer[i])).
					Msg("Found compatible pairs for node")

				// Log first few pairs for debugging
				loggingLimit := 3
				for pairIdx := 0; pairIdx < len(compatiblePairsForEachServer[i]) && pairIdx < loggingLimit; pairIdx++ {
					pair := compatiblePairsForEachServer[i][pairIdx]
					log.Debug().
						Str("machineId", node.MachineId).
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

	// Generate multiple infrastructure candidates based on Pareto efficiency principles.
	// For each source node, compatible spec-image pairs are ranked by match rate
	// (CPU, Memory, Image similarity to the original node).
	// - Candidate 0: Best match pairs (highest similarity to source servers)
	// - Candidate i: i-th ranked pairs (alternative solutions with different characteristics)
	// This provides a Pareto frontier of solutions exploring performance-availability trade-offs.
	// Note: Cost optimization is planned for future implementation.

	// For each candidate up to the actual limit
	for i := 0; i < actualLimit; i++ {

		// Create a copy of the skeleton NodeGroup List
		tempNodeGroupList := make([]cloudmodel.CreateNodeGroupReq, len(skeletonNodegroupList))
		copy(tempNodeGroupList, skeletonNodegroupList)

		var selectedVmSpec cloudmodel.SpecInfo
		var selectedVmOsImage cloudmodel.ImageInfo

		// For each node, select the i-th compatible pair of VM spec and OS image
		for j, node := range srcInfra.Nodes {

			// Select compatible pairs for the j-th node
			compatiblePairs := compatiblePairsForEachServer[j]
			if len(compatiblePairs) == 0 {
				log.Warn().Msgf("no compatible VM spec and OS image pairs found for node %s", node.MachineId)
				continue
			}

			// Pareto optimal selection: Select the i-th ranked pair for this candidate.
			// Pairs are ranked by match rate to the source node (CPU, Memory, Image similarity).
			// This explores alternative solutions in the multi-dimensional space.

			// Select the i-th pair, or reuse the optimal pair (Rank 1) when alternatives are exhausted
			pairIdx := i
			var pair CompatibleSpecImagePair
			if i < len(compatiblePairs) {
				pair = compatiblePairs[i]
			} else {
				pairIdx = 0
				pair = compatiblePairs[0]
				log.Debug().Msgf("candidate %d: node %s has only %d pairs available, reusing optimal pair", i+1, node.MachineId, len(compatiblePairs))
			}

			selectedVmSpec = pair.Spec
			selectedVmOsImage = pair.Image

			// Calculate match rate vector for this node-VM pair
			matchRateVec := calculateMatchRateVector(csp, node, selectedVmSpec, selectedVmOsImage)

			// Log candidate and spec selection details with match rate
			log.Debug().
				Str("machineId", node.MachineId).
				Int("candidateIndex", i).
				Int("serverIndex", j).
				Int("pairIndex", pairIdx).
				Int("totalPairsForServer", len(compatiblePairs)).
				Str("selectedSpecId", selectedVmSpec.Id).
				Str("selectedSpecName", selectedVmSpec.CspSpecName).
				Str("selectedImageId", selectedVmOsImage.Id).
				Str("selectedImageName", selectedVmOsImage.CspImageName).
				Float64("cpuMatchRate", matchRateVec.CPU).
				Float64("memoryMatchRate", matchRateVec.Memory).
				Float64("imageMatchRate", matchRateVec.Image).
				Msg("Selected spec-image pair for candidate")

			// Log CPU comparison
			log.Debug().
				Str("machineId", node.MachineId).
				Str("specCspName", selectedVmSpec.CspSpecName).
				Str("specId", selectedVmSpec.Id).
				Uint32("originalCPUs", node.CPU.Cpus).
				Uint32("recommendedVCPU", uint32(selectedVmSpec.VCPU)).
				Msg("CPU comparison")

			// Log Memory comparison
			log.Trace().
				Str("machineId", node.MachineId).
				Str("specCspName", selectedVmSpec.CspSpecName).
				Str("specId", selectedVmSpec.Id).
				Uint32("originalMemoryGB", uint32(node.Memory.TotalSize)).
				Float32("recommendedMemoryGiB", selectedVmSpec.MemoryGiB).
				Msg("Memory comparison")

			// Log OS comparison
			log.Trace().
				Str("machineId", node.MachineId).
				Str("imageCspName", selectedVmOsImage.CspImageName).
				Str("imageId", selectedVmOsImage.Id).
				Str("originalOS", node.OS.Name+" "+node.OS.Version).
				Str("recommendedOSImage", selectedVmOsImage.CspImageName).
				Msg("OS comparison")

			// * Set the selected spec and image IDs to the corresponding NodeGroup
			tempNodeGroupList[j].SpecId = selectedVmSpec.Id
			tempNodeGroupList[j].ImageId = selectedVmOsImage.Id

			// Resolve the latest CSP image name and confirm available system disk.
			precheck, reviewErr := PreflightCheckCspProvisioning(
				selectedVmSpec.Id, selectedVmOsImage.Id, selectedVmOsImage.CspImageName, "",
			)
			if reviewErr != nil {
				log.Warn().Err(reviewErr).Msgf("preflight check failed for node %s; using cached image", node.MachineId)
			} else {
				if precheck.SuggestedSystemDisk != "" {
					log.Debug().Msgf("using suggested system disk for node %s: %s", node.MachineId, precheck.SuggestedSystemDisk)
					tempNodeGroupList[j].RootDiskType = precheck.SuggestedSystemDisk
				}
				if precheck.ResolvedCspImageName != selectedVmOsImage.CspImageName {
					log.Info().Msgf("image resolved to latest for node %s: %s -> %s", node.MachineId, selectedVmOsImage.CspImageName, precheck.ResolvedCspImageName)
					tempNodeGroupList[j].CspImageName = precheck.ResolvedCspImageName
				}
			}

			// Include match rate vector in NodeGroup description for transparency
			tempNodeGroupList[j].Description = fmt.Sprintf(
				"Recommended VM for %s | %s",
				node.MachineId,
				matchRateVec.FormatMatchRateSummary(),
			)

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

		// Create a candidate infrastructure based on skeleton and current tempNodeGroupList
		candidateInfra := skeletonVmInfra
		candidateInfra.TargetInfra.NodeGroups = tempNodeGroupList
		candidateInfra.TargetSpecList = deduplicatedVmSpecList
		candidateInfra.TargetOsImageList = deduplicatedVmOsImageList
		candidateInfra.TargetSecurityGroupList = deduplicatedSecurityGroupList
		candidateInfra.TargetInfra.Description = "Recommended VMs comprising multi-cloud infrastructure"

		// Calculate overall match rate with detailed information
		overallStatus, overallStatusDesc, infraMatchRateSummary := calculateCandidateMatchRateWithDetails(csp, tempNodeGroupList, srcInfra, deduplicatedVmSpecList, deduplicatedVmOsImageList, minMatchRate)

		// Skip duplicate candidate if all node groups match an existing candidate
		if isDuplicateInfraCandidate(recommendedVmInfraCandidates, candidateInfra) {
			log.Debug().Int("candidateIndex", i).Msg("skipping duplicate candidate infrastructure")
			continue
		}

		// Set the status and enhanced description with match rate summary
		candidateInfra.Status = overallStatus
		candidateInfra.Description = fmt.Sprintf(
			"Candidate #%d | %s | Overall Match Rate: Min=%.1f%% Max=%.1f%% Avg=%.1f%% | %s",
			len(recommendedVmInfraCandidates)+1,
			overallStatus,
			infraMatchRateSummary.MinMatchRate,
			infraMatchRateSummary.MaxMatchRate,
			infraMatchRateSummary.AvgMatchRate,
			overallStatusDesc,
		)

		recommendedVmInfraCandidates = append(recommendedVmInfraCandidates, candidateInfra)
	}

	/*
	 * [Output]
	 */
	log.Trace().Msgf("recommended infrastructure candidates: %+v", recommendedVmInfraCandidates)

	return recommendedVmInfraCandidates, nil
}

