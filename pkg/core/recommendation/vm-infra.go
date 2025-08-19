package recommendation

import (
	"fmt"
	"math/big"
	"net"
	"sort"
	"strings"

	"github.com/cloud-barista/cb-tumblebug/src/core/common/netutil"
	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/compat"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"

	// "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/infra"
	// "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"

	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/similarity"
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
		"ncpvpc": true,
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
		log.Warn().Msgf(err.Error())
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
	var max int = 5
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
		vmSpecList, _, err := RecommendVmSpecs(desiredCsp, desiredRegion, server, max)
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
			TargetVmInfra: cloudmodel.TbMciDynamicReq{
				Name:        fmt.Sprintf("migrated-%02d", i),
				Description: "a recommended multi-cloud infrastructure",
				Vm:          []cloudmodel.TbVmDynamicReq{},
			},
		}

		for j, vmInfo := range vmInfoList {
			tempVmReq := cloudmodel.TbVmDynamicReq{
				ConnectionName: fmt.Sprintf("%s-%s", desiredCsp, desiredRegion),
				CommonImage:    vmInfo.vmOsImageId,
				CommonSpec:     vmInfo.vmSpecId,
				Description:    "a recommended virtual machine",
				Name:           fmt.Sprintf("migrated-%s", srcInfra.Servers[j].MachineId), // Set MachineId to identify the source server
				RootDiskSize:   "",                                                        // TBD
				RootDiskType:   "",                                                        // TBD
				SubGroupSize:   "",
				VmUserPassword: "",
			}
			tempVmInfraInfo.TargetVmInfra.Vm = append(tempVmInfraInfo.TargetVmInfra.Vm, tempVmReq)
		}

		status := checkOverallVmStatus(tempVmInfraInfo.TargetVmInfra.Vm)
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
	var limitSpecs int = 10
	var limitImages int = 20

	// Initialize the response body
	recommendedVmInfra = cloudmodel.RecommendedVmInfra{
		Description: "This is a list of recommended target infrastructures. Please review and use them.",
		Status:      "",
		TargetVmInfra: cloudmodel.TbMciReq{
			Name:        "mmci01",
			Description: "a recommended multi-cloud infrastructure",
			Vm:          []cloudmodel.TbVmReq{},
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
	// var recommendedSshKey = tbmodel.TbSshKeyReq{}
	// * Set a name to indicate a dependency between resources.
	recommendedVmInfra.TargetSshKey.Name = "mig-sshkey-01"
	recommendedVmInfra.TargetSshKey.ConnectionName = fmt.Sprintf("%s-%s", csp, region)
	recommendedVmInfra.TargetSshKey.Description = "a SSH Key pair for migration (Note - provided ONLY once, MUST be downloaded"

	// 3. Recommend VM specs, OS images, and security groups, and
	// recommend VMs by removing duplicates of VM specs, OS images, and security groups and specifying them.
	// Note: Don't need to register specs and OS images.
	var recommendedVmList = []cloudmodel.TbVmReq{}
	var recommendedVmSpecList = []cloudmodel.TbSpecInfo{}
	var recommendedVmOsImageList = []cloudmodel.TbImageInfo{}
	var recommendedSecurityGroupList = []cloudmodel.TbSecurityGroupReq{}

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

		log.Debug().Msgf("recommendedVmSpecInfoList: %+v", recommendedVmSpecInfoList)
		log.Debug().Msgf("recommendedVmOsImageInfoList: %+v", recommendedVmOsImageInfoList)
		var selectedVmSpec cloudmodel.TbSpecInfo
		var selectedVmOsImage cloudmodel.TbImageInfo
		if len(recommendedVmSpecInfoList) == 0 || len(recommendedVmOsImageInfoList) == 0 {
			log.Warn().Msgf("no recommended VM specs or OS images found for server %s", server.MachineId)
		} else {
			// Find compatible spec and image pair
			tempSelectedVmSpec, tempSelectedVmOsImage, err := FindCompatibleSpecAndImage(recommendedVmSpecInfoList, recommendedVmOsImageInfoList, csp)
			if err != nil {
				log.Warn().Msgf("failed to find compatible spec-image pair for server %s: %v", server.MachineId, err)
				// Use fallback selection (first spec, first image)
			} else {
				selectedVmSpec = tempSelectedVmSpec
				selectedVmOsImage = tempSelectedVmOsImage
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
		exists, _, _ = containSg(recommendedSecurityGroupList, recommendedSg)
		if !exists {
			// If the security group does not exist, set a name to indicate a dependency between resources.
			recommendedSg.Name = fmt.Sprintf("mig-sg-%02d", len(recommendedSecurityGroupList)+1)
			recommendedSg.ConnectionName = fmt.Sprintf("%s-%s", csp, region)
			recommendedSg.Description = fmt.Sprintf("Recommended security group for %s", server.MachineId) // Set MachineId to identify the source server

			// * Set name to indicate a dependency between resources.
			recommendedSg.VNetId = recommendedVmInfra.TargetVNet.Name // Set the vNet ID to the security group

			// Set the security group to the response body
			recommendedSecurityGroupList = append(recommendedSecurityGroupList, recommendedSg)
		}

		/*
		 * Recommend VM by specifying the recommended VM specs, OS images, and security groups
		 */
		// TODO: Select a subnet by the server's network information
		// xxx

		// * Set names to indicate a dependency between resources.
		tempVmReq := cloudmodel.TbVmReq{
			ConnectionName:   fmt.Sprintf("%s-%s", csp, region),
			Description:      fmt.Sprintf("a recommended virtual machine %02d for %s", i+1, server.MachineId), // Set MachineId to identify the source server
			SpecId:           selectedVmSpec.CspSpecName,
			ImageId:          selectedVmOsImage.CspImageName,
			VNetId:           recommendedVmInfra.TargetVNet.Name,
			SubnetId:         recommendedVmInfra.TargetVNet.SubnetInfoList[0].Name, // Set the first subnet for simplicity
			SecurityGroupIds: []string{recommendedSg.Name},                         // Set the security group ID
			Name:             fmt.Sprintf("migrated-%s", server.MachineId),         // Set MachineId to identify the source server
			RootDiskSize:     "",                                                   // TBD
			RootDiskType:     "",                                                   // TBD
			SshKeyId:         recommendedVmInfra.TargetSshKey.Name,                 // Set the SSH key ID
			VmUserName:       "",                                                   // TBD: Set the VM user name if needed
			VmUserPassword:   "",                                                   // TBD
			SubGroupSize:     "",                                                   // TBD
		}

		// Append the VM request to the list
		recommendedVmList = append(recommendedVmList, tempVmReq)
	}

	/*
	 * [Output]
	 */
	recommendedVmInfra.TargetVmInfra.Vm = recommendedVmList
	recommendedVmInfra.TargetVmSpecList = recommendedVmSpecList
	recommendedVmInfra.TargetVmOsImageList = recommendedVmOsImageList
	recommendedVmInfra.TargetSecurityGroupList = recommendedSecurityGroupList

	log.Trace().Msgf("the recommended infra info: %+v", recommendedVmInfra)

	return recommendedVmInfra, nil
}

func RecommendVNet(csp string, region string, srcInfra onpremmodel.OnpremInfra) ([]cloudmodel.TbVNetReq, error) {

	var emptyRes []cloudmodel.TbVNetReq
	var recommendedVNets []cloudmodel.TbVNetReq

	// [Input]
	ok, err := IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid csp (%s) or region (%s)", csp, region)
		return emptyRes, err
	}

	if len(srcInfra.Network.IPv4Networks.CidrBlocks) == 0 && len(srcInfra.Network.IPv4Networks.DefaultGateways) == 0 {
		err := fmt.Errorf("no network information found in the source computing infrastructure")
		log.Error().Err(err).Msg("failed to recommend a virtual network")
		return emptyRes, err
	}

	var srcNetworks []string
	// * Note: srcInfra.Network.IPv4Networks.CidrBlocks is the input from the user (e.g., network admin)
	if len(srcInfra.Network.IPv4Networks.CidrBlocks) != 0 {
		srcNetworks = srcInfra.Network.IPv4Networks.CidrBlocks
	} else if len(srcInfra.Network.IPv4Networks.DefaultGateways) != 0 {
		// * Note: To estimate the network address space of the source computing infrastructure,
		// * Source networks are derived by combining the default gateway and network interface information of each server.
		srcNetworks, err = deriveSourceNetworksFromDefaultGateways(srcInfra)
		if err != nil {
			log.Error().Err(err).Msg("failed to derive CIDR blocks from default gateways")
			return emptyRes, err
		}
	} else {
		log.Warn().Msg("no network information found in the source computing infrastructure")
		return emptyRes, fmt.Errorf("no network information found in the source computing infrastructure")
	}
	log.Debug().Msgf("Source networks (CIDR blocks): %v", srcNetworks)

	// [Process] Recommend the vNet and subnets
	// * Note:
	// * At least 1 subnet is required.
	// * Derive a super network that includes user's all networks and set it as a vNet
	// * Set user's networks as subnets

	// ? Assumption: a network in on-premise infrastructure is designed and configured with various network segments or types.
	// * Thus, it must be selected which of these network segments will be the vNet.
	// ? If so, is grouping the network segments required?

	// Categorizes the entered CIDR blocks by private network (i.e., 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16)
	var cidrs10 []string
	var cidrs172 []string
	var cidrs192 []string

	for _, srcNetwork := range srcNetworks {
		identifiedNet, err := netutil.WhichPrivateNetworkByCidr(srcNetwork)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to identify the network %s", srcNetwork)
			continue
		}
		log.Debug().Msgf("identified network: %s", identifiedNet)

		switch identifiedNet {
		case netutil.PrivateNetwork10Dot:
			cidrs10 = append(cidrs10, srcNetwork)
		case netutil.PrivateNetwork172Dot:
			cidrs172 = append(cidrs172, srcNetwork)
		case netutil.PrivateNetwork192Dot:
			cidrs192 = append(cidrs192, srcNetwork)
		default:
			log.Warn().Msgf("skipped because CIDR block (%s) is not a private network", srcNetwork)
			continue
		}
	}
	log.Debug().Msgf("CIDR blocks for %s: %v", netutil.PrivateNetwork10Dot, cidrs10)
	log.Debug().Msgf("CIDR blocks for %s: %v", netutil.PrivateNetwork172Dot, cidrs172)
	log.Debug().Msgf("CIDR blocks for %s: %v", netutil.PrivateNetwork192Dot, cidrs192)

	// Calculate the super network of each group
	var supernet10, supernet172, supernet192 string = "", "", ""

	if len(cidrs10) > 0 {
		supernet10, err = netutil.CalculateSupernet(cidrs10)
		if err != nil {
			log.Warn().Err(err).Msg("failed to calculate supernet")
		}
		log.Debug().Msgf("supernet10: %s\n", supernet10)
	}

	if len(cidrs172) > 0 {
		supernet172, err = netutil.CalculateSupernet(cidrs172)
		if err != nil {
			log.Warn().Err(err).Msg("failed to calculate supernet")
		}
		log.Debug().Msgf("supernet172: %s\n", supernet172)
	}

	if len(cidrs192) > 0 {
		supernet192, err = netutil.CalculateSupernet(cidrs192)
		if err != nil {
			log.Warn().Err(err).Msg("failed to calculate supernet")
		}
		log.Debug().Msgf("supernet192: %s\n", supernet192)
	}

	// Estimate the more :D super network for each private network
	// TODO: Set the number of networks to be included in the super network
	estimateNumNetworks := 4
	if len(supernet10) > 0 {
		supernet10, err = estimateSupernet(supernet10, estimateNumNetworks)
		if err != nil {
			log.Warn().Err(err).Msg("failed to estimate supernet for 10.x.x.x")
		}
	}
	if len(supernet172) > 0 {
		supernet172, err = estimateSupernet(supernet172, estimateNumNetworks)
		if err != nil {
			log.Warn().Err(err).Msg("failed to estimate supernet for 172.x.x.x")
		}
	}
	if len(supernet192) > 0 {
		supernet192, err = estimateSupernet(supernet192, estimateNumNetworks)
		if err != nil {
			log.Warn().Err(err).Msg("failed to estimate supernet for 192.x.x.x")
		}
	}

	// Select a super network for the vNet
	// ? But how to select the super network?
	// * Currrently, a list of recommended networks is returned.

	if supernet10 != "" {
		// Set tempSubnets by the CIDR blocks from the source computing infra
		tempSubnets := []cloudmodel.TbSubnetReq{}
		for _, cidr := range cidrs10 {
			networkAddr, err := toNetworkAddress(cidr)
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse CIDR block %s", cidr)
				continue
			}

			tempSubnets = append(tempSubnets, cloudmodel.TbSubnetReq{
				Name:        "INSERT_YOUR_SUBNET_NAME", // TODO: Set a name for the subnet
				Description: "subnet from source computing infra",
				IPv4_CIDR:   networkAddr,
			})
		}

		// Set the calculated supernet as the tempVNet
		tempVNet := cloudmodel.TbVNetReq{
			Name:           "INSERT_YOUR_VNET_NAME", // TODO: Set a name for the vNet
			ConnectionName: fmt.Sprintf("%s-%s", csp, region),
			Description:    "Recommended vNet for " + netutil.PrivateNetwork10Dot,
			CidrBlock:      supernet10,
			SubnetInfoList: tempSubnets,
		}

		// Append recommended virtual network info to the list
		recommendedVNets = append(recommendedVNets, tempVNet)
	}

	if supernet172 != "" {

		// Set tempSubnets by the CIDR blocks from the source computing infra
		tempSubnets := []cloudmodel.TbSubnetReq{}
		for _, cidr := range cidrs172 {
			networkAddr, err := toNetworkAddress(cidr)
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse CIDR block %s", cidr)
				continue
			}

			tempSubnets = append(tempSubnets, cloudmodel.TbSubnetReq{
				Name:        "INSERT_YOUR_SUBNET_NAME", // TODO: Set a name for the subnet
				Description: "subnet from source computing infra",
				IPv4_CIDR:   networkAddr,
			})
		}

		tempVNet := cloudmodel.TbVNetReq{
			Name:           "INSERT_YOUR_VNET_NAME", // TODO: Set a name for the vNet
			ConnectionName: fmt.Sprintf("%s-%s", csp, region),
			Description:    "Recommended vNet for " + netutil.PrivateNetwork172Dot,
			CidrBlock:      supernet172,
			SubnetInfoList: tempSubnets,
		}

		// Append recommended virtual network info to the list
		recommendedVNets = append(recommendedVNets, tempVNet)
	}

	if supernet192 != "" {

		// Set tempSubnets by the CIDR blocks from the source computing infra
		tempSubnets := []cloudmodel.TbSubnetReq{}
		for _, cidr := range cidrs192 {

			networkAddr, err := toNetworkAddress(cidr)
			if err != nil {
				log.Warn().Err(err).Msgf("failed to parse CIDR block %s", cidr)
				continue
			}

			tempSubnets = append(tempSubnets, cloudmodel.TbSubnetReq{
				Name:        "INSERT_YOUR_SUBNET_NAME", // TODO: Set a name for the subnet
				Description: "subnet from source computing infra",
				IPv4_CIDR:   networkAddr,
			})
		}

		// Set the calculated supernet as the vNet
		tempVNet := cloudmodel.TbVNetReq{
			Name:           "INSERT_YOUR_VNET_NAME", // TODO: Set a name for the vNet
			ConnectionName: fmt.Sprintf("%s-%s", csp, region),
			Description:    "Recommended vNet for " + netutil.PrivateNetwork192Dot,
			CidrBlock:      supernet192,
			SubnetInfoList: tempSubnets,
		}

		// Append recommended virtual network info to the list
		recommendedVNets = append(recommendedVNets, tempVNet)
	}

	// [Output]
	if len(recommendedVNets) == 0 {
		return emptyRes, fmt.Errorf("no recommended virtual network found for the source computing infra")
	}

	return recommendedVNets, nil
}

func deriveSourceNetworksFromDefaultGateways(srcInfra onpremmodel.OnpremInfra) ([]string, error) {
	if len(srcInfra.Network.IPv4Networks.DefaultGateways) == 0 {
		return nil, fmt.Errorf("no network information found in the source computing infrastructure")
	}

	var sourceNetworks []string
	// 1. Find the server that has the same "machine ID" as the gateway
	for _, gateway := range srcInfra.Network.IPv4Networks.DefaultGateways {
		for _, server := range srcInfra.Servers {
			if server.MachineId == gateway.MachineId {

				// 2. Find the network interface that has the same network "name" as the gateway
				for _, nic := range server.Interfaces {
					if nic.Name == gateway.InterfaceName {

						// 3. Get "prefix length" from the network interface
						if nic.IPv4CidrBlocks == nil && len(nic.IPv4CidrBlocks) == 0 {
							log.Warn().Msgf("no IPv4 CIDR blocks found in the network interface %s of the server %s", nic.Name, server.MachineId)
							continue
						}

						cidrBlock := nic.IPv4CidrBlocks[0]
						_, ipNet, err := net.ParseCIDR(cidrBlock)
						if err != nil {
							log.Warn().Err(err).Msgf("failed to parse CIDR block %s", cidrBlock)
							continue
						}

						prefixLen, _ := ipNet.Mask.Size()

						// 4. Derive the CIDR block from the gateway IP and prefix length
						gatewayCIDR := fmt.Sprintf("%s/%d", gateway.IP, prefixLen)

						// 5. Append the derived CIDR block to the list
						sourceNetworks = append(sourceNetworks, gatewayCIDR)
					}
				}
			}
		}
	}

	// Deduplicate the source networks
	sourceNetworks = deduplicateSlice(sourceNetworks)

	return sourceNetworks, nil
}

func deduplicateSlice[T comparable](slice []T) []T {
	// Create a map to track unique elements
	uniqueMap := make(map[T]struct{})
	for _, item := range slice {
		uniqueMap[item] = struct{}{}
	}

	// Convert the map keys back to a slice
	result := make([]T, 0, len(uniqueMap))
	for item := range uniqueMap {
		result = append(result, item)
	}
	return result
}

// estimateSupernet finds the smallest supernet that contains a given number
// of consecutive networks, starting from a given CIDR.
func estimateSupernet(startCIDR string, numNetworks int) (string, error) {
	// 1. Parse the starting CIDR.
	ip, ipNet, err := net.ParseCIDR(startCIDR)
	if err != nil {
		return "", fmt.Errorf("invalid CIDR: %v", err)
	}

	// Ensure it's an IPv4 address.
	ipv4 := ip.To4()
	if ipv4 == nil {
		return "", fmt.Errorf("only IPv4 addresses are supported")
	}

	// 2. Calculate the total IP range.
	// Number of addresses in the start network (e.g., /24 -> 256).
	prefixLen, bits := ipNet.Mask.Size()
	numAddressesPerNet := 1 << (bits - prefixLen)

	// Total number of addresses to cover.
	totalAddresses := numAddressesPerNet * numNetworks

	// Convert the starting IP to an integer for calculation.
	startIPint := big.NewInt(0)
	startIPint.SetBytes(ipv4)

	// Calculate the last IP address in the entire range.
	// Last IP = Start IP + Total Addresses - 1.
	offset := big.NewInt(int64(totalAddresses - 1))
	endIPint := big.NewInt(0)
	endIPint.Add(startIPint, offset)

	// Convert the integer back to a net.IP.
	firstIP := ipv4
	lastIP := net.IP(endIPint.Bytes())

	// 3. Find the common supernet.
	// Iterate from the initial prefix down to 0, finding the first prefix
	// length where both the first and last IPs belong to the same network.
	for newPrefixLen := prefixLen; newPrefixLen >= 0; newPrefixLen-- {
		mask := net.CIDRMask(newPrefixLen, bits)
		network1 := firstIP.Mask(mask)
		network2 := lastIP.Mask(mask)

		// If both IPs belong to the same network, we've found our supernet.
		if network1.Equal(network2) {
			return (&net.IPNet{IP: network1, Mask: mask}).String(), nil
		}
	}

	return "", fmt.Errorf("could not find a common supernet")
}

func toNetworkAddress(cidr string) (string, error) {
	_, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", fmt.Errorf("failed to parse CIDR block %s: %v", cidr, err)
	}
	return subnet.String(), nil
}

// RecommendVmSpecsForImage recommends appropriate VM specs for the server and image
func RecommendVmSpecsForImage(csp string, region string, server onpremmodel.ServerProperty, limit int, image cloudmodel.TbImageInfo) (vmSpecList []cloudmodel.TbSpecInfo, length int, err error) {

	vmSpecList, length, err = RecommendVmSpecs(csp, region, server, limit)
	if err != nil {
		log.Warn().Err(err).Msg("failed to recommend VM specs")
		return nil, 0, err
	}

	// Use unified compatibility filtering instead of CSP-specific switches
	compatibleSpecs := make([]cloudmodel.TbSpecInfo, 0, len(vmSpecList))

	for _, spec := range vmSpecList {
		if isCompatible := compat.CheckCompatibility(strings.ToLower(csp), spec, image); isCompatible {
			compatibleSpecs = append(compatibleSpecs, spec)
		} else {
			log.Debug().Msgf("Filtered incompatible spec: %s for image: %s on CSP: %s",
				spec.CspSpecName, image.CspImageName, csp)
		}
	}

	if len(compatibleSpecs) == 0 {
		log.Warn().Msgf("No compatible specs found for image %s on CSP %s, returning original list",
			image.CspImageName, csp)
		return vmSpecList, length, nil
	}

	log.Info().Msgf("Filtered %d specs to %d compatible specs for image %s on CSP %s",
		len(vmSpecList), len(compatibleSpecs), image.CspImageName, csp)

	return compatibleSpecs, len(compatibleSpecs), nil
}

// RecommendVmSpecs recommends appropriate VM specs for the given server
func RecommendVmSpecs(csp string, region string, server onpremmodel.ServerProperty, limit int) (vmSpecList []cloudmodel.TbSpecInfo, length int, err error) {

	// Constants
	const (
		defaultLimit        = 5
		defaultArchitecture = "x86_64"
	)

	var emptyResp = []cloudmodel.TbSpecInfo{}

	// Validate and set default limit
	if limit <= 0 {
		log.Warn().Msgf("Invalid limit value: %d, setting to default: %d", limit, defaultLimit)
		limit = defaultLimit
	}

	// Deployment plan template for VM spec recommendation
	// * Note:
	// * ">=": greater than or equal to
	// * "<=": less than or equal to
	// * The plan is designed to recommend VM specs based on vCPU and memory ranges.
	// Reference: https://github.com/cloud-barista/cb-tumblebug/discussions/1234
	const planTemplate = `{
		"filter": {
			"policy": [
				{
					"condition": [
						{"operand": "%d", "operator": ">="},
						{"operand": "%d", "operator": "<="}
					],
					"metric": "vCPU"
				},
				{
					"condition": [
						{"operand": "%d", "operator": ">="},
						{"operand": "%d", "operator": "<="}
					],
					"metric": "memoryGiB"
				},
				{
					"condition": [{"operand": "%s"}],
					"metric": "providerName"
				},
				{
					"condition": [{"operand": "%s"}],
					"metric": "regionName"
				},
				{
					"condition": [{"operand": "%s"}],
					"metric": "architecture"
				}
			]
		},
		"limit": "%d",
		"priority": {
			"policy": [{"metric": "cost"}]
		}
	}`

	// Extract server specifications from source computing envrionment
	vcpus := server.CPU.Cpus
	memory := uint32(server.Memory.TotalSize)

	// Calculate optimal vCPU and memory ranges based on AWS, GCP, and NCP instance patterns
	vcpusMin, vcpusMax, memoryMin, memoryMax := calculateOptimalRange(vcpus, memory)

	// Set provider and region names
	providerName := strings.ToLower(csp)
	regionName := strings.ToLower(region)

	// Set architecture (default: "x86_64")
	architecture := server.CPU.Architecture
	if architecture == "" || architecture == "amd64" {
		architecture = defaultArchitecture
	}

	// Set OS name and version
	osNameAndVersion := server.OS.Name + " " + server.OS.Version
	osNameWithVersion := strings.ToLower(osNameAndVersion)

	log.Debug().
		Str("machineId", server.MachineId).
		Uint32("originalVcpus", vcpus).
		Uint32("originalMemory", memory).
		Float64("memoryVcpuRatio", float64(memory)/float64(vcpus)).
		Uint32("vcpuRange", vcpusMax-vcpusMin).
		Uint32("memoryRange", memoryMax-memoryMin).
		Str("provider", providerName).
		Str("region", regionName).
		Str("architecture", architecture).
		Str("osNameWithVersion", osNameWithVersion).
		Msgf("Calculating VM spec recommendations for machine: %s", server.MachineId)

	// Create deployment plan with calculated parameters
	deploymentPlan := fmt.Sprintf(planTemplate,
		vcpusMin, vcpusMax,
		memoryMin, memoryMax,
		providerName, regionName, architecture,
		limit,
	)
	log.Debug().Msgf("Deployment plan for machine %s: %s", server.MachineId, deploymentPlan)

	// Initialize Tumblebug API client
	tbCli := tbclient.NewClient(tbclient.ApiConfig{
		RestUrl:  config.Tumblebug.RestUrl,
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
	})

	// Call Tumblebug API to get recommended VM specs
	vmSpecInfoList, err := tbCli.MciRecommendVm(deploymentPlan)
	if err != nil {
		log.Error().Err(err).
			Str("machineId", server.MachineId).
			Str("provider", providerName).
			Str("region", region).
			Msg("Failed to get VM spec recommendations from Tumblebug")
		return emptyResp, -1, fmt.Errorf("failed to get VM spec recommendations for machine %s: %w", server.MachineId, err)
	}

	numOfVmSpecs := len(vmSpecInfoList)
	if numOfVmSpecs == 0 {
		err := fmt.Errorf("no VM specs recommended for machine %s (vcpus: %d, memory: %d GiB)",
			server.MachineId, vcpus, memory)
		log.Warn().Err(err).
			Str("machineId", server.MachineId).
			Uint32("vcpus", vcpus).
			Uint32("memory", memory).
			Msg("No VM specifications found")
		return emptyResp, -1, err
	}

	log.Info().
		Str("machineId", server.MachineId).
		Int("specsFound", numOfVmSpecs).
		Uint32("vcpus", vcpus).
		Uint32("memory", memory).
		Msgf("Found %d VM spec recommendations for machine: %s", numOfVmSpecs, server.MachineId)

	// NCP-specific filtering for KVM hypervisor
	if strings.Contains(strings.ToLower(csp), "ncp") {
		log.Debug().
			Str("machineId", server.MachineId).
			Msg("Filtering VM specs for KVM hypervisor (NCP)")

		kvmVmSpecs := make([]tbmodel.TbSpecInfo, 0, len(vmSpecInfoList))
		for _, vmSpec := range vmSpecInfoList {
			for _, detail := range vmSpec.Details {
				if detail.Key == "HypervisorType" && strings.Contains(strings.ToLower(detail.Value), "kvm") {
					kvmVmSpecs = append(kvmVmSpecs, vmSpec)
					break
				}
			}
		}

		if len(kvmVmSpecs) > 0 {
			vmSpecInfoList = kvmVmSpecs
			log.Debug().
				Str("machineId", server.MachineId).
				Int("kvmSpecs", len(kvmVmSpecs)).
				Msg("Filtered to KVM-compatible specs for NCP")
		} else {
			log.Warn().
				Str("machineId", server.MachineId).
				Msg("No KVM-compatible specs found for NCP, using all available specs")
		}
	}

	// [Output]
	// Apply limit to results
	finalSpecCount := len(vmSpecInfoList)
	if limit < finalSpecCount {
		vmSpecInfoList = vmSpecInfoList[:limit]
		finalSpecCount = limit
	}

	log.Debug().
		Str("machineId", server.MachineId).
		Int("finalSpecCount", finalSpecCount).
		Msg("Finalized VM spec recommendations")

	// Convert model types with validation
	convertedVmSpecList, err := modelconv.ConvertWithValidation[[]tbmodel.TbSpecInfo, []cloudmodel.TbSpecInfo](vmSpecInfoList)
	if err != nil {
		log.Error().Err(err).
			Str("machineId", server.MachineId).
			Msg("Failed to convert VM spec list model")
		return emptyResp, -1, fmt.Errorf("failed to convert VM spec list for machine %s: %w", server.MachineId, err)
	}

	for i, vmSpec := range convertedVmSpecList {
		log.Debug().Msgf("Recommended VM specification %d: %+v", i+1, vmSpec)
	}

	log.Info().
		Str("machineId", server.MachineId).
		Int("recommendedSpecs", len(convertedVmSpecList)).
		Msgf("Successfully recommended %d VM specifications for machine: %s", len(convertedVmSpecList), server.MachineId)

	return convertedVmSpecList, numOfVmSpecs, nil
}

// RecommendVmOsImage recommends an appropriate VM OS image (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImage(csp string, region string, server onpremmodel.ServerProperty) (cloudmodel.TbImageInfo, error) {

	var emptyRes cloudmodel.TbImageInfo

	imageList, err := RecommendVmOsImages(csp, region, server, 20)
	if err != nil {
		log.Error().Err(err).Msg("Failed to recommend VM OS images")
		return emptyRes, err
	}

	// Set keywords and delimiters to calculate text similarity
	keywords, kwDelimiters, imgDelimiters := setKeywordsAndDelimeters(server)
	log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

	// Find the best VM OS image
	bestVmOsImage := FindBestVmOsImage(keywords, kwDelimiters, imageList, imgDelimiters)

	log.Debug().Msgf("Best VM OS image found: %+v", bestVmOsImage)

	return bestVmOsImage, nil
}

// RecommendVmOsImageId recommends an appropriate VM OS image (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImageId(csp string, region string, server onpremmodel.ServerProperty) (string, error) {

	imageList, err := RecommendVmOsImages(csp, region, server, 20)
	if err != nil {
		log.Error().Err(err).Msg("Failed to recommend VM OS images")
		return "", err
	}

	// Set keywords and delimiters to calculate text similarity
	keywords, kwDelimiters, imgDelimiters := setKeywordsAndDelimeters(server)
	log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

	vmOsImageId := FindBestVmOsImageNameUsedInCsp(keywords, kwDelimiters, imageList, imgDelimiters)

	log.Debug().Msgf("Best VM OS image ID found: %s", vmOsImageId)

	return vmOsImageId, nil
}

// RecommendVmOsImages recommends an appropriate VM OS image (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImages(csp string, region string, server onpremmodel.ServerProperty, limit int) ([]cloudmodel.TbImageInfo, error) {

	var emptyRes = []cloudmodel.TbImageInfo{}
	var vmOsImageInfoList = []cloudmodel.TbImageInfo{}

	if limit <= 0 {
		err := fmt.Errorf("invalid 'limit' value: %d, set default: 5", limit)
		log.Warn().Msgf(err.Error())
		limit = 5
	}

	// Request body
	falseValue := false
	trueValue := true
	searchImageReq := tbmodel.SearchImageRequest{
		// DetailSearchKeys:       []string{},
		// IncludeDeprecatedImage: &falseValue,
		// IsKubernetesImage:      &falseValue, // The only image in the Azure (ubuntu 22.04) is both for K8s nodes and gerneral VMs.
		// IsRegisteredByAsset:    &falseValue,
		IncludeBasicImageOnly: &trueValue,
		OSArchitecture:        tbmodel.OSArchitecture(server.CPU.Architecture),
		OSType:                server.OS.Name + " " + server.OS.VersionID,
		ProviderName:          csp,
		RegionName:            region,
	}

	// TODO: Add condition to check if searchImageReq.IsGPUImage is set, when GPU information is confirmed in the source model
	searchImageReq.IsGPUImage = &falseValue

	log.Debug().Msgf("searchImageReq: %+v", searchImageReq)

	// Call Tumblebug API to search VM OS images
	apiConfig := tbclient.ApiConfig{
		RestUrl:  config.Tumblebug.RestUrl,
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
	}
	tbCli := tbclient.NewClient(apiConfig)
	nsId := "system" // default

	resSearchImage, err := tbCli.SearchVmOsImage(nsId, searchImageReq)
	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyRes, err
	}

	// Debug logging up to 3 images to avoid excessive output
	if len(resSearchImage.ImageList) > 3 {
		for i := range 3 {
			log.Debug().Msgf("[Response from Tumblebug] resSearchImage.ImageList[%d]: %+v", i, resSearchImage.ImageList[i])
		}
	} else {
		for i := range resSearchImage.ImageList {
			log.Debug().Msgf("[Response from Tumblebug] resSearchImage.ImageList[%d]: %+v", i, resSearchImage.ImageList[i])
		}
	}

	// Filter VM OS images to support stability
	var filteredImages []tbmodel.TbImageInfo
	for _, img := range resSearchImage.ImageList {
		if strings.Contains(strings.ToLower(img.CspImageName), "uefi") {
			continue
		}
		// Add more filters as needed

		filteredImages = append(filteredImages, img)
	}

	// Convert model from '[]tbmodel.TbImageInfo' to '[]cloudmodel.TbImageInfo'
	imageList, err := modelconv.ConvertWithValidation[[]tbmodel.TbImageInfo, []cloudmodel.TbImageInfo](filteredImages)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert VM OS image list")
		return emptyRes, err
	}

	// Set keywords and delimiters to calculate text similarity
	keywords, kwDelimiters, imgDelimiters := setKeywordsAndDelimeters(server)
	log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

	// Select VM OS image via LevenshteinDistance-based text similarity
	vmOsImageInfoList = FindAndSortVmOsImageInfoListBySimilarity(keywords, kwDelimiters, imageList, imgDelimiters)

	count := len(vmOsImageInfoList)
	if count == 0 {
		err := fmt.Errorf("no VM OS image recommended for the inserted PM/VM")
		log.Warn().Msgf(err.Error())
		return emptyRes, err
	}

	// [Output]
	// Limit the number of VM specs
	if limit < count {
		log.Debug().Msgf("Limiting the number of recommended VM OS images to %d", limit)
		// * Note: If the number of recommended VM OS images is less than the limit, it will not be truncated.
		// * This is to ensure that the user can see all available images.
		vmOsImageInfoList = vmOsImageInfoList[:limit]
	}

	log.Debug().Msgf("Found %d VM OS images for the given server: %s", len(vmOsImageInfoList), server.MachineId)

	return vmOsImageInfoList, nil
}

func setKeywordsAndDelimeters(server onpremmodel.ServerProperty) (string, []string, []string) {
	keywords := fmt.Sprintf("%s %s %s %s %s",
		server.OS.Name,
		server.OS.VersionID,
		server.OS.VersionCodename,
		server.CPU.Architecture,
		server.RootDisk.Type)

	kwDelimiters := []string{" ", "-", ",", "(", ")", "[", "]", "/"}
	imgDelimiters := []string{" ", "-", ",", "(", ")", "[", "]", "/"}

	return keywords, kwDelimiters, imgDelimiters
}

func transposeMatrix[T any](matrix [][]T) [][]T {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]T{}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]T, cols)
	for i := range result {
		result[i] = make([]T, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

// Function to check overall status for the entire list of VMs
func checkOverallVmStatus(vms []cloudmodel.TbVmDynamicReq) string {
	allOk := true
	allNone := true

	for _, vm := range vms {
		if vm.CommonImage == "" || vm.CommonSpec == "" {
			allOk = false // At least one VM is not fully populated
		}
		if vm.CommonImage != "" || vm.CommonSpec != "" {
			allNone = false // At least one VM has a value
		}
	}

	// Determine overall status
	if allNone {
		return string(NothingRecommended)
	} else if allOk {
		return string(FullyRecommended)
	} else {
		return string(PartiallyRecommended)
	}
}

func MBtoGiB(mb float64) uint32 {
	const bytesInMB = 1000000.0
	const bytesInGiB = 1073741824.0
	gib := (mb * bytesInMB) / bytesInGiB
	return uint32(gib)
}

// isPrimeNumber checks if a number is prime
func isPrimeNumber(n uint32) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := uint32(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// findPreviousPrimeNumberOrOne finds the largest prime number smaller than n
func findPreviousPrimeNumberOrOne(n uint32) uint32 {
	if n <= 2 {
		return 1 // Return 1 as minimum vCPU count if no smaller prime exists
	}

	for i := n - 1; i >= 2; i-- {
		if isPrimeNumber(i) {
			return i
		}
	}
	return 1 // Return 1 as fallback minimum vCPU count
}

// findNextPrimeNumber finds the smallest prime number larger than n
func findNextPrimeNumber(n uint32) uint32 {
	candidate := n + 1
	for {
		if isPrimeNumber(candidate) {
			return candidate
		}
		candidate++
	}
}

// calculateOptimalRange calculates optimal vCPU and memory ranges based on AWS instance patterns
func calculateOptimalRange(vcpus uint32, memory uint32) (vcpusMin, vcpusMax, memoryMin, memoryMax uint32) {
	// Constants for instance type thresholds and ratios
	const (
		computeIntensiveRatioThreshold = 3.0 // 1:2 ratio instances
		memoryIntensiveRatioThreshold  = 7.0 // 1:8 ratio instances
		minMemoryBound                 = 2   // Minimum memory requirement
		minVcpuBound                   = 1   // Minimum vCPU requirement
		maxVcpuForMemoryIntensive      = 10  // Maximum vCPU for memory intensive
	)

	memoryToVcpuRatio := float64(memory) / float64(vcpus)

	switch {
	case memoryToVcpuRatio <= computeIntensiveRatioThreshold: // Compute Intensive (1:2)
		return calculateComputeIntensiveRange(vcpus, memory, minMemoryBound)
	case memoryToVcpuRatio >= memoryIntensiveRatioThreshold: // Memory Intensive (1:8)
		return calculateMemoryIntensiveRange(vcpus, memory, minVcpuBound, maxVcpuForMemoryIntensive)
	default: // General Purpose (1:4)
		return calculateGeneralPurposeRange(vcpus, memory, minVcpuBound, minMemoryBound)
	}
}

func calculateComputeIntensiveRange(vcpus, memory, minMemoryBound uint32) (vcpusMin, vcpusMax, memoryMin, memoryMax uint32) {
	const (
		vcpuRangeBuffer  = 2 // Buffer for vCPU range expansion
		memoryMultiplier = 4 // Memory multiplier for max calculation
	)

	vcpusMin = findPreviousPrimeNumberOrOne(vcpus)
	vcpusMax = findNextPrimeNumber(vcpus + vcpuRangeBuffer)

	// Set a wide search range for memory for compute-intensive workloads
	memoryMin = minMemoryBound
	memoryMax = vcpusMax * memoryMultiplier

	return vcpusMin, vcpusMax, memoryMin, memoryMax
}

func calculateMemoryIntensiveRange(vcpus, memory, minVcpuBound, maxVcpuForMemoryIntensive uint32) (vcpusMin, vcpusMax, memoryMin, memoryMax uint32) {
	const (
		memoryRangeBuffer = 4 // Buffer for memory range expansion
		memoryToCpuRatio  = 8 // Standard memory to CPU ratio for calculation
	)

	memoryMin = findPreviousPrimeNumberOrOne(memory)
	memoryMax = findNextPrimeNumber(memory + memoryRangeBuffer)

	// Set a wide search range for vCPU for memory-intensive workloads
	vcpusMin = minVcpuBound
	vcpusMax = memoryMax / memoryToCpuRatio
	if vcpusMax < maxVcpuForMemoryIntensive {
		vcpusMax = maxVcpuForMemoryIntensive
	}

	return vcpusMin, vcpusMax, memoryMin, memoryMax
}

func calculateGeneralPurposeRange(vcpus, memory, minVcpuBound, minMemoryBound uint32) (vcpusMin, vcpusMax, memoryMin, memoryMax uint32) {
	const (
		vcpuRangeBuffer     = 2 // Buffer for vCPU range expansion
		memoryRangeBuffer   = 4 // Buffer for memory range expansion
		minMemoryToCpuRatio = 2 // Minimum memory to CPU ratio (1:2)
		maxCpuToMemoryRatio = 2 // Maximum CPU to memory ratio (2:1)
	)

	// For General Purpose, provide balanced flexibility for both vCPU and memory
	// Allow moderate range expansion while maintaining 1:4 ratio as the center point
	vcpusMin = findPreviousPrimeNumberOrOne(vcpus)
	vcpusMax = findNextPrimeNumber(vcpus + vcpuRangeBuffer) // Slightly wider vCPU range

	memoryMin = findPreviousPrimeNumberOrOne(memory)
	memoryMax = findNextPrimeNumber(memory + memoryRangeBuffer) // Moderate memory range

	// Apply minimum bounds
	if vcpusMin < minVcpuBound {
		vcpusMin = minVcpuBound
	}
	if memoryMin < minMemoryBound {
		memoryMin = minMemoryBound
	}

	// Ensure reasonable relationship between vCPU and memory
	// Allow 1:2 to 1:8 ratio range for General Purpose workloads
	if memoryMax < vcpusMin*minMemoryToCpuRatio {
		memoryMax = vcpusMin * minMemoryToCpuRatio
	}
	if vcpusMax > memoryMax/maxCpuToMemoryRatio {
		vcpusMax = memoryMax / maxCpuToMemoryRatio
		if vcpusMax < vcpusMin {
			vcpusMax = vcpusMin
		}
	}

	return vcpusMin, vcpusMax, memoryMin, memoryMax
}

// FindBestVmOsImage finds the best matching image based on the similarity scores
func FindBestVmOsImage(keywords string, kwDelimiters []string, vmImages []cloudmodel.TbImageInfo, imgDelimiters []string) cloudmodel.TbImageInfo {

	var bestVmOsImage cloudmodel.TbImageInfo
	var highestScore float64 = 0.0

	for _, image := range vmImages {

		vmImgKeywords := fmt.Sprintf("%s %s %s %s",
			image.OSType,
			image.OSArchitecture,
			image.OSDiskType,
			image.OSDistribution,
		)

		score := similarity.CalcResourceSimilarity(keywords, kwDelimiters, vmImgKeywords, imgDelimiters)
		if score > highestScore {
			highestScore = score
			bestVmOsImage = image
		}
		// log.Debug().Msgf("VmImageName: %s, score: %f, description: %s", image.OSDistribution, score, image.Description)

	}
	log.Debug().Msgf("highestScore: %f, bestVmOsImage: %v", highestScore, bestVmOsImage)

	return bestVmOsImage
}

type VmOsImageInfoWithScore struct {
	VmOsImageInfo   cloudmodel.TbImageInfo
	SimilarityScore float64
}

// FindAndSortVmOsImageInfoListBySimilarity finds VM OS images that match the keywords and sorts them by similarity score
func FindAndSortVmOsImageInfoListBySimilarity(keywords string, kwDelimiters []string, vmImages []cloudmodel.TbImageInfo, imgDelimiters []string) []cloudmodel.TbImageInfo {

	var imageInfoListForSorting []VmOsImageInfoWithScore
	var imageInfoList []cloudmodel.TbImageInfo

	for _, image := range vmImages {

		vmImgKeywords := fmt.Sprintf("%s %s %s %s",
			image.OSType,
			image.OSArchitecture,
			image.OSDiskType,
			image.OSDistribution,
		)

		score := similarity.CalcResourceSimilarity(keywords, kwDelimiters, vmImgKeywords, imgDelimiters)
		imageInfo := VmOsImageInfoWithScore{
			VmOsImageInfo:   image,
			SimilarityScore: score,
		}
		imageInfoListForSorting = append(imageInfoListForSorting, imageInfo)

	}

	// Sort the imageInfoList by highestScore in descending order
	sort.Slice(imageInfoListForSorting, func(i, j int) bool {
		return imageInfoListForSorting[i].SimilarityScore > imageInfoListForSorting[j].SimilarityScore
	})

	// List the sorted images
	for _, imageInfo := range imageInfoListForSorting {
		log.Debug().Msgf("VmImageName: %s, score: %f, description: %s", imageInfo.VmOsImageInfo.Name, imageInfo.SimilarityScore, imageInfo.VmOsImageInfo.Description)
		imageInfoList = append(imageInfoList, imageInfo.VmOsImageInfo)
	}

	return imageInfoList
}

// FindBestVmOsImageNameUsedInCsp finds the best matching image based on the similarity scores
func FindBestVmOsImageNameUsedInCsp(keywords string, kwDelimiters []string, vmImages []cloudmodel.TbImageInfo, imgDelimiters []string) string {

	var bestVmOsImageNameUsedInCsp string
	var highestScore float64 = 0.0

	for _, image := range vmImages {
		vmImgKeywords := fmt.Sprintf("%s %s %s %s",
			image.OSType,
			image.OSArchitecture,
			image.OSDiskType,
			image.OSDistribution,
		)

		score := similarity.CalcResourceSimilarity(keywords, kwDelimiters, vmImgKeywords, imgDelimiters)
		if score > highestScore {
			highestScore = score
			bestVmOsImageNameUsedInCsp = image.CspImageName
		}
		// log.Debug().Msgf("VmImageName: %s, score: %f, description: %s", image.OSDistribution, score, image.Description)

	}
	log.Debug().Msgf("bestVmOsImageID: %s, highestScore: %f", bestVmOsImageNameUsedInCsp, highestScore)

	return bestVmOsImageNameUsedInCsp
}

func RecommendSecurityGroup(csp string, region string, server onpremmodel.ServerProperty) (cloudmodel.TbSecurityGroupReq, error) {

	var emptyRes = cloudmodel.TbSecurityGroupReq{}
	var recommendedSecurityGroup = cloudmodel.TbSecurityGroupReq{}

	// [Input]
	ok, err := IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid provider (%s) or region (%s)", csp, region)
		return emptyRes, err
	}

	firewallRules := server.FirewallTable
	log.Debug().Msgf("firewallRules: %+v", firewallRules)

	// Default rules
	// * Note: Spider supports this rule. Do not set this rule to avoid duplication error.
	// ruleToAllowAllOutboundTraffic := cloudmodel.TbFirewallRuleInfo{
	// 	Direction:  "outbound",
	// 	IPProtocol: "all",
	// 	CIDR:       "0.0.0.0/0",
	// 	FromPort:   "0",
	// 	ToPort:     "0",
	// }
	ruleToAllowSSHInboundTraffic := cloudmodel.TbFirewallRuleInfo{
		Direction: "inbound",
		Protocol:  "tcp",
		CIDR:      "0.0.0.0/0",
		Ports:     "22",
	}

	// [Process] Recommend the security group
	// Create security group recommendations
	var sgRules []cloudmodel.TbFirewallRuleInfo
	// 1. Set default security group rules if no firewall rules are provided
	if len(firewallRules) == 0 {
		log.Warn().Msg("no firewall rules provided, using default rules")
		// Allow all outbound traffic and deny all inbound traffic
		// TODO: Check if the default rules are OK on testing.
		// sgRules = append(sgRules, ruleToAllowAllOutboundTraffic)
		sgRules = append(sgRules, ruleToAllowSSHInboundTraffic)
	} else {
		sgRules = generateSecurityGroupRules(firewallRules)
	}

	log.Debug().Msgf("sgRules: %+v", sgRules)

	// [Output]
	// Create a security group for all rules
	recommendedSecurityGroup = cloudmodel.TbSecurityGroupReq{
		Name:           "INSERT_YOUR_SECURITY_GROUP_NAME",
		VNetId:         "INSERT_YOUR_VNET_ID",
		ConnectionName: fmt.Sprintf("%s-%s", csp, region),
		Description:    fmt.Sprintf("Recommended security group for %s", server.MachineId), // Set MachineId to identify the source server
		FirewallRules:  &sgRules,
	}

	log.Debug().Msgf("recommendedSecurityGroup: %+v", recommendedSecurityGroup)

	return recommendedSecurityGroup, nil
}

func RecommendSecurityGroups(csp string, region string, servers []onpremmodel.ServerProperty) (cloudmodel.RecommendedSecurityGroupList, error) {

	var emptyRet = cloudmodel.RecommendedSecurityGroupList{}
	var recommendedSecurityGroupList = cloudmodel.RecommendedSecurityGroupList{}

	// [Input]
	ok, err := IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid provider (%s) or region (%s)", csp, region)
		return emptyRet, err
	}

	// [Process] Recommend the security group for each server
	var tempRecSgList []cloudmodel.TbSecurityGroupReq
	var targetSecurityGroupList []cloudmodel.RecommendedSecurityGroup

	for _, server := range servers {
		// Recommend a security group for the server
		recommendedTargetSg, err := RecommendSecurityGroup(csp, region, server)
		if err != nil {
			log.Error().Err(err).Msgf("failed to recommend security group for server: %+v", server)
			recommendedTargetSg.Description = fmt.Sprintf("Failed to recommend security group for %s", server.MachineId) // Set MachineId to identify the source server
			recommendedTargetSg.FirewallRules = nil                                                                      // No rules if recommendation fails
		}

		// Check duplicates and append the recommended security group
		exists, idx, _ := containSg(tempRecSgList, recommendedTargetSg)

		// If not exists, append the recommended security group
		// If exists, just append the MachineId to the existing security group
		if !exists {
			// Note: This is a temporary list for checking duplicates
			tempRecSgList = append(tempRecSgList, recommendedTargetSg)

			// Create a temporary recommended security group
			tempRecommendedSecurityGroup := cloudmodel.RecommendedSecurityGroup{
				SourceServers:       []string{server.MachineId}, // Set MachineId to identify the source server
				Description:         "Recommended security group",
				TargetSecurityGroup: recommendedTargetSg,
			}

			// Set status
			if recommendedTargetSg.FirewallRules != nil {
				tempRecommendedSecurityGroup.Status = string(FullyRecommended)
			} else {
				tempRecommendedSecurityGroup.Status = string(NothingRecommended)
			}

			// Append the recommended security group to the list
			targetSecurityGroupList = append(targetSecurityGroupList, tempRecommendedSecurityGroup)
		} else {
			// Just append the MachineId to the existing security group
			targetSecurityGroupList[idx].SourceServers = append(targetSecurityGroupList[idx].SourceServers, server.MachineId)
		}
	}

	// [Output]
	countFailed := 0
	for _, recSg := range targetSecurityGroupList {
		if recSg.Status == string(NothingRecommended) {
			countFailed++
		}
	}

	recommendedSecurityGroupList.Count = len(targetSecurityGroupList)
	switch countFailed {
	case 0:
		recommendedSecurityGroupList.Status = string(FullyRecommended)
		recommendedSecurityGroupList.Description = "Successfully recommended and deduplicated security groups for all servers"
	case recommendedSecurityGroupList.Count:
		recommendedSecurityGroupList.Status = string(NothingRecommended)
		recommendedSecurityGroupList.Description = "Unable to recommend any security groups for the servers in the source infrastructure"
	default:
		recommendedSecurityGroupList.Status = string(PartiallyRecommended)
		recommendedSecurityGroupList.Description = fmt.Sprintf("Partially recommended security groups: %d of %d server groups have recommendations",
			recommendedSecurityGroupList.Count-countFailed, recommendedSecurityGroupList.Count)
	}

	recommendedSecurityGroupList.TargetSecurityGroupList = targetSecurityGroupList

	log.Debug().Msgf("recommendedSecurityGroupList: %+v", recommendedSecurityGroupList)

	return recommendedSecurityGroupList, nil
}

func containSg(sgList []cloudmodel.TbSecurityGroupReq, sg cloudmodel.TbSecurityGroupReq) (bool, int, cloudmodel.TbSecurityGroupReq) {

	// Check duplicates and append the recommended security group
	temp := cloudmodel.TbSecurityGroupReq{}
	exists := false
	idx := -1
	for i, sgItem := range sgList {
		// Both SGs have rules defined
		if sgItem.FirewallRules != nil && sg.FirewallRules != nil {
			// Quick check if they have the same number of rules
			if len(*sgItem.FirewallRules) == len(*sg.FirewallRules) {
				areAllRulesSame := true

				// Create maps for each rule in both security groups for comparison
				sgRulesMap := make(map[string]bool)
				for _, rule := range *sg.FirewallRules {
					// Create a unique key for each rule
					key := fmt.Sprintf("%s-%s-%s-%s",
						rule.Direction, rule.Protocol, rule.CIDR, rule.Ports)
					sgRulesMap[key] = true
				}

				// Check if all rules in the recommended SG exist in the current SG
				for _, rule := range *sgItem.FirewallRules {
					key := fmt.Sprintf("%s-%s-%s-%s",
						rule.Direction, rule.Protocol, rule.CIDR, rule.Ports)
					if !sgRulesMap[key] {
						areAllRulesSame = false
						break
					}
				}

				if areAllRulesSame {
					exists = true
					temp = sgItem
					idx = i
					break
				}
			}
		}
	}

	return exists, idx, temp
}

// formatCIDR formats the CIDR string:
// - If it's "anywhere", return "0.0.0.0/0"
// - If it doesn't have a prefix (like "/24"), add "/32"
// - Otherwise return as is
func formatCIDR(cidr string) string {
	if cidr == "anywhere" {
		return "0.0.0.0/0"
	}

	// Check if the CIDR has a prefix
	if !strings.Contains(cidr, "/") {
		// If it's a valid IP without prefix, add "/32"
		return cidr + "/32"
	}

	return cidr
}

// generateSecurityGroupRules converts FirewallRuleProperty to tbmodel.TbFirewallRuleInfo
func generateSecurityGroupRules(rules []onpremmodel.FirewallRuleProperty) []cloudmodel.TbFirewallRuleInfo {
	var tbRules []cloudmodel.TbFirewallRuleInfo

	for _, rule := range rules {
		// Skip 'deny' rules (note: SecurityGroup does not support adding 'deny' rules)
		if rule.Action == "deny" {
			continue
		}

		// Skip rules with no protocol specified
		if rule.Protocol == "" {
			log.Warn().Msgf("Protocol is not specified in rule: %+v - skipping rule", rule)
			continue
		}

		// Skip IPv6 rules (currently not supported)
		if isIPv6Rule(rule) {
			log.Warn().Msgf("IPv6 rule detected but not currently supported: %+v - skipping rule", rule)
			continue
		}

		// Handle protocol wildcard
		protocol := rule.Protocol
		if protocol == "*" {
			protocol = "ALL"
		}

		switch rule.Direction {
		case "inbound":
			// Set CIDR block for source - For inbound, use source CIDR (where traffic comes from)
			if rule.SrcCIDR == "" {
				log.Warn().Msgf("SrcCIDR is not specified in rule: %+v - skipping rule", rule)
				continue
			}

			// Format the CIDR correctly
			srcCIDR := formatCIDR(rule.SrcCIDR)
			log.Debug().Msgf("Formatted SrcCIDR from '%s' to '%s'", rule.SrcCIDR, srcCIDR)

			// ! Skip default outbound rule that allows all traffic because it is automatically created by cloud providers, CB-Spider, or CB-Tumblebug.
			// TODO: To be updated if the default rule is needed in the future.
			if strings.ToLower(protocol) == "all" && srcCIDR == "0.0.0.0/0" {
				log.Debug().Msgf("Skipping default inbound ALL traffic rule (may conflict with existing rules): %+v", rule)
				continue
			}

			// Handle destination ports based on format
			if rule.DstPorts == "" {
				// Skip rules without port information for non-ICMP/ALL protocols
				log.Debug().Msgf("Skipping inbound rule without port information: %+v", rule)
				continue
			}

			// * NOTE: Handle destination ports (where traffic is going to)
			// Special cases based on CB-Spider specification:
			// - (protocol: TCP/UDP/ALL) "*" port from the source is converted to "1-65535" for the target ports.
			// - (protocol: ICMP) "*" ports from the source is omitted in the target ports.
			// - Comma-separated ports (e.g., 22,23,24)
			// - Port range with colon notation (e.g., 30000:40000)
			// - Single port (e.g., 22)

			protocolLower := strings.ToLower(protocol)
			switch protocolLower {
			case "icmp":
				tbRule := cloudmodel.TbFirewallRuleInfo{
					Direction: rule.Direction,
					Protocol:  protocol,
					CIDR:      srcCIDR,
				}
				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created inbound rule for 'ICMP' protocol: %+v", tbRule)

			case "tcp", "udp", "all":
				var dstPorts string
				// Handle wildcard ports based on protocol
				if rule.DstPorts == "*" {
					dstPorts = "1:65535" // TCP/UDP use 1-65535 range
				} else {
					dstPorts = rule.DstPorts
				}

				// * Note: CB-Tumblebug will handle comma-separated ports and port ranges.
				// In here, just convert : to - for port ranges.
				if strings.Contains(dstPorts, ":") {
					// Handle port range with colon notation (e.g., 30000:40000)
					portRange := strings.Split(dstPorts, ":")
					if len(portRange) == 2 {
						dstPorts = strings.TrimSpace(portRange[0]) + "-" + strings.TrimSpace(portRange[1])
					} else {
						log.Warn().Msgf("Invalid port range format in rule.DstPorts: %s - skipping rule", dstPorts)
						continue
					}
				}

				tbRule := cloudmodel.TbFirewallRuleInfo{
					Direction: rule.Direction,
					Protocol:  protocol,
					CIDR:      srcCIDR,
					Ports:     dstPorts,
				}

				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created inbound rule for '%s' protocol: %+v", protocol, tbRule)
			default:
				log.Warn().Msgf("Unsupported protocol '%s' in inbound rule: %+v - skipping rule", protocol, rule)
				continue
			}

		case "outbound":
			// Set CIDR block for destination
			if rule.DstCIDR == "" {
				// Skip rule if no CIDR is specified
				log.Warn().Msgf("No CIDR specified for outbound rule: %+v - skipping rule", rule)
				continue
			}

			// Format the CIDR correctly
			dstCIDR := formatCIDR(rule.DstCIDR)
			log.Debug().Msgf("Formatted outbound CIDR from '%s' to '%s'", rule.DstCIDR, dstCIDR)

			// ! Skip default outbound rule that allows all traffic because it is automatically created by cloud providers, CB-Spider, or CB-Tumblebug.
			// TODO: To be updated if the default rule is needed in the future.
			if strings.ToLower(protocol) == "all" && dstCIDR == "0.0.0.0/0" {
				log.Debug().Msgf("Skipping default outbound ALL traffic rule (may conflict with existing rules): %+v", rule)
				continue
			}

			// Handle destination ports based on format
			if rule.DstPorts == "" {
				// Skip rules without port information for non-ICMP/ALL protocols
				log.Debug().Msgf("Skipping inbound rule without port information: %+v", rule)
				continue
			}

			// * NOTE: Handle destination ports (where traffic is going to)
			// Special cases based on CB-Spider specification:
			// - (protocol: TCP/UDP/ALL) "*" port from the source is converted to "1-65535" for the target ports.
			// - (protocol: ICMP) "*" ports from the source is omitted in the target ports.
			// - Comma-separated ports (e.g., 22,23,24)
			// - Port range with colon notation (e.g., 30000:40000)
			// - Single port (e.g., 22)

			protocolLower := strings.ToLower(protocol)
			switch protocolLower {
			case "icmp":
				// Special case for ICMP protocol - no ports needed, just CIDR
				tbRule := cloudmodel.TbFirewallRuleInfo{
					Direction: rule.Direction,
					Protocol:  protocol,
					CIDR:      dstCIDR,
				}
				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created outbound rule for 'ICMP' protocol: %+v", tbRule)

			case "tcp", "udp", "all": // Handle destination ports with wildcard support based on CB-Spider specification

				var dstPorts string
				// Handle wildcard ports based on protocol
				if rule.DstPorts == "*" {
					dstPorts = "1:65535" // TCP/UDP use 1-65535 range
				} else {
					dstPorts = rule.DstPorts
				}

				// * Note: CB-Tumblebug will handle comma-separated ports and port ranges.
				// In here, just convert : to - for port ranges.
				if strings.Contains(dstPorts, ":") {
					// Handle port range with colon notation
					portRange := strings.Split(dstPorts, ":")
					if len(portRange) == 2 {
						dstPorts = strings.TrimSpace(portRange[0]) + "-" + strings.TrimSpace(portRange[1])
					} else {
						log.Warn().Msgf("Invalid port range format: %s - skipping rule", dstPorts)
						continue
					}
				}

				tbRule := cloudmodel.TbFirewallRuleInfo{
					Direction: rule.Direction,
					Protocol:  protocol,
					CIDR:      dstCIDR,
					Ports:     dstPorts,
				}
				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created outbound rule for '%s' protocol: %+v", protocol, tbRule)
			default:
				log.Warn().Msgf("Unsupported protocol '%s' in outbound rule: %+v - skipping rule", protocol, rule)
				continue
			}

		default:
			log.Warn().Msgf("Unknown direction '%s' in rule: %+v", rule.Direction, rule)
		}

		log.Debug().Msgf("Original FirewallRule: %+v", rule)
	}

	return tbRules
}

// isIPv6Rule checks if the firewall rule contains IPv6 elements
func isIPv6Rule(rule onpremmodel.FirewallRuleProperty) bool {
	// Check for IPv6 CIDR blocks (contains ":")
	if strings.Contains(rule.SrcCIDR, ":") || strings.Contains(rule.DstCIDR, ":") {
		return true
	}

	// Check for IPv6-specific protocols
	protocol := strings.ToLower(rule.Protocol)
	if protocol == "icmpv6" || protocol == "ipv6" {
		return true
	}

	return false
}

// FindCompatibleSpecAndImage finds a compatible VM spec and image pair by performing CSP-specific compatibility checks
func FindCompatibleSpecAndImage(specs []cloudmodel.TbSpecInfo, images []cloudmodel.TbImageInfo, csp string) (cloudmodel.TbSpecInfo, cloudmodel.TbImageInfo, error) {
	var emptySpec cloudmodel.TbSpecInfo
	var emptyImage cloudmodel.TbImageInfo

	if len(specs) == 0 {
		return emptySpec, emptyImage, fmt.Errorf("no VM specs provided")
	}
	if len(images) == 0 {
		return emptySpec, emptyImage, fmt.Errorf("no VM images provided")
	}

	log.Debug().Msgf("Finding compatible spec and image for CSP: %s, specs: %d, images: %d", csp, len(specs), len(images))

	// Pre-filter specs and images based on CSP-specific rules
	filteredSpecs, filteredImages := preFilterByCsp(csp, specs, images)

	if len(filteredSpecs) == 0 {
		return emptySpec, emptyImage, fmt.Errorf("no compatible VM specs found after CSP-specific filtering")
	}
	if len(filteredImages) == 0 {
		return emptySpec, emptyImage, fmt.Errorf("no compatible VM images found after CSP-specific filtering")
	}

	log.Debug().Msgf("After pre-filtering - specs: %d, images: %d", len(filteredSpecs), len(filteredImages))

	// Find best compatible pair without scoring
	bestSpec, bestImage, err := findCompatiblePair(csp, filteredSpecs, filteredImages)
	if err != nil {
		return emptySpec, emptyImage, fmt.Errorf("failed to find compatible spec-image pair: %w", err)
	}

	log.Info().Msgf("Found compatible pair - Spec: %s, Image: %s", bestSpec.CspSpecName, bestImage.CspImageName)
	return bestSpec, bestImage, nil
}

// preFilterByCsp performs CSP-specific pre-filtering with integrated logic
func preFilterByCsp(csp string, specs []cloudmodel.TbSpecInfo, images []cloudmodel.TbImageInfo) ([]cloudmodel.TbSpecInfo, []cloudmodel.TbImageInfo) {
	cspLower := strings.ToLower(csp)

	switch cspLower {
	case "aws":
		// Filter out UEFI images for AWS
		filteredImages := make([]cloudmodel.TbImageInfo, 0, len(images))
		for _, img := range images {
			if !strings.Contains(strings.ToLower(img.CspImageName), "uefi") {
				filteredImages = append(filteredImages, img)
			}
		}
		log.Debug().Msgf("AWS pre-filtering: %d images filtered to %d", len(images), len(filteredImages))
		return specs, filteredImages

	case "ncp":
		// Filter specs for KVM hypervisor only
		filteredSpecs := filterNcpVmSpecsByHypervisor(specs)
		log.Debug().Msgf("NCP pre-filtering: %d specs filtered to %d KVM specs", len(specs), len(filteredSpecs))
		return filteredSpecs, images

	default:
		// No specific filtering for GCP, Azure, and others
		log.Debug().Msgf("No specific pre-filtering rules for CSP: %s", csp)
		return specs, images
	}
}

// findCompatiblePair finds the first compatible spec-image pair using comprehensive compatibility checks
func findCompatiblePair(csp string, specs []cloudmodel.TbSpecInfo, images []cloudmodel.TbImageInfo) (cloudmodel.TbSpecInfo, cloudmodel.TbImageInfo, error) {
	var emptySpec cloudmodel.TbSpecInfo
	var emptyImage cloudmodel.TbImageInfo

	cspLower := strings.ToLower(csp)

	// Use standard compatibility check for all CSPs
	for _, spec := range specs {
		for _, image := range images {
			if isCompatible := compat.CheckCompatibility(cspLower, spec, image); isCompatible {
				log.Info().Msgf("Found compatible pair - Spec: %s, Image: %s",
					spec.CspSpecName, image.CspImageName)
				log.Debug().Msgf("Compatible pair (spec): %v", spec)
				log.Debug().Msgf("Compatible pair (image): %v", image)
				return spec, image, nil
			}
		}
	}

	return emptySpec, emptyImage, fmt.Errorf("no compatible spec-image pairs found")
}

// filterNcpVmSpecsByHypervisor filters NCP VM specs to include only KVM hypervisor specs
func filterNcpVmSpecsByHypervisor(vmSpecs []cloudmodel.TbSpecInfo) []cloudmodel.TbSpecInfo {
	if len(vmSpecs) == 0 {
		return vmSpecs
	}

	log.Debug().Msgf("NCP filtering: checking %d VM specs for KVM hypervisor", len(vmSpecs))

	var filteredSpecs []cloudmodel.TbSpecInfo

	for _, spec := range vmSpecs {
		hasKvmHypervisor := false

		// Check if this spec has KVM hypervisor
		for _, detail := range spec.Details {
			if strings.EqualFold(detail.Key, "hypervisortype") &&
				strings.Contains(strings.ToUpper(detail.Value), "KVM") {
				hasKvmHypervisor = true
				break
			}
		}

		if hasKvmHypervisor {
			filteredSpecs = append(filteredSpecs, spec)
			log.Debug().Msgf("NCP: included VM spec %s (KVM hypervisor found)", spec.CspSpecName)
		} else {
			log.Debug().Msgf("NCP: filtered out VM spec %s (no KVM hypervisor)", spec.CspSpecName)
		}
	}

	log.Debug().Msgf("NCP filtering: %d VM specs filtered to %d KVM specs", len(vmSpecs), len(filteredSpecs))

	// If no KVM specs found, return original list with warning
	if len(filteredSpecs) == 0 {
		log.Warn().Msg("No KVM hypervisor specs found for NCP, returning all specs")
		return vmSpecs
	}

	return filteredSpecs
}
