package recommendation

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloud-barista/cb-tumblebug/src/core/common/netutil"
	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"

	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"

	// "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/infra"
	// "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"

	inframodel "github.com/cloud-barista/cm-model/infra/onprem"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/similarity"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

func isSupportedCSP(csp string) bool {
	supportedCSPs := map[string]bool{
		"aws":   true,
		"azure": true,
		"gcp":   true,
		// "alibaba": true,
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

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	// [via Tumblebug] Check if the provider and region are valid
	method := "GET"
	url := fmt.Sprintf("%s/provider/%s/region/%s", epTumblebug, cspName, regionName)

	// Request body
	tbReqt := common.NoBody
	tbResp := tbmodel.RegionDetail{}

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(tbReqt),
		&tbReqt,
		&tbResp,
		common.VeryShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("")
		return isValid, err
	}

	isValid = true

	return isValid, nil
}

// RecommendVmInfraDynamic an appropriate multi-cloud infrastructure (MCI) for cloud migration
func RecommendVmInfraDynamic(desiredCsp string, desiredRegion string, srcInfra inframodel.OnpremInfra) (RecommendedVmInfraDynamicList, error) {

	// var emptyResp RecommendedVmInfraInfoList
	var recommendedVmInfraInfoList RecommendedVmInfraDynamicList

	// ! To be updated, a user will input the desired number of recommended VMs
	var max int = 5
	// Initialize the response body
	recommendedVmInfraInfoList = RecommendedVmInfraDynamicList{
		Description:       "This is a list of recommended target infrastructures. Please review and use them.",
		Count:             0,
		TargetVmInfraList: []RecommendedVmInfraDynamic{},
	}

	// // Set VM info
	// recommendedVm := tbmodel.TbVmDynamicReq{
	// 	ConnectionName: "",
	// 	CommonImage:    "", // Lookup and set an appropriate VM OS image
	// 	CommonSpec:     "", // Lookup and set an appropriate VM spec
	// 	Description:    "a recommended virtual machine",
	// 	Name:           fmt.Sprintf("migrated-%s", server.Hostname),
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
			log.Warn().Msgf("failed to recommend VM specs for server %s: %v", server.Hostname, err)
			continue
		}

		// Lookup the appropriate VM OS images for the server
		vmOsImageIdList := []string{}
		for _, vmSpec := range vmSpecList {
			osImgId, err := RecommendVmOsImageId(desiredCsp, desiredRegion, server, vmSpec.Id)
			if err != nil {
				log.Warn().Msgf("failed to recommend VM OS image for server %s: %v", server.Hostname, err)
				vmOsImageIdList = append(vmOsImageIdList, "")
			}
			vmOsImageIdList = append(vmOsImageIdList, osImgId)
		}

		// Set the recommended VM specs and OS images to the response body
		recommendedVmInfo := []RecommendedVmInfo{}
		for i, vmSpec := range vmSpecList {
			recommendedVmInfo = append(recommendedVmInfo, RecommendedVmInfo{
				vmSpecId:    vmSpec.Id,
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
	recommenedVmInfraInfoList := []RecommendedVmInfraDynamic{}

	for i, vmInfoList := range transposed {

		tempVmInfraInfo := RecommendedVmInfraDynamic{
			Status:      string(NothingRecommended),
			Description: "This is a recommended target infrastructure.",
			TargetVmInfra: tbmodel.TbMciDynamicReq{
				Name:        fmt.Sprintf("migrated-%02d", i),
				Description: "a recommended multi-cloud infrastructure",
				Vm:          []tbmodel.TbVmDynamicReq{},
			},
		}

		for j, vmInfo := range vmInfoList {
			tempVmReq := tbmodel.TbVmDynamicReq{
				ConnectionName: "",
				CommonImage:    vmInfo.vmOsImageId,
				CommonSpec:     vmInfo.vmSpecId,
				Description:    "a recommended virtual machine",
				Name:           fmt.Sprintf("migrated-%s", srcInfra.Servers[j].Hostname),
				RootDiskSize:   "", // TBD
				RootDiskType:   "", // TBD
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
func RecommendVmInfra(desiredCsp string, desiredRegion string, srcInfra inframodel.OnpremInfra) (RecommendedVmInfra, error) {

	var emptyResp RecommendedVmInfra
	var recommendedVmInfra RecommendedVmInfra

	// ! To be updated, a user will input the desired number of recommended VMs
	var max int = 5

	// Initialize the response body
	recommendedVmInfra = RecommendedVmInfra{
		Description: "This is a list of recommended target infrastructures. Please review and use them.",
		Status:      "",
		TargetVmInfra: tbmodel.TbMciReq{
			Name:        "mmci-01",
			Description: "a recommended multi-cloud infrastructure",
			Vm:          []tbmodel.TbVmReq{},
		},
	}

	csp := strings.ToLower(desiredCsp)
	region := strings.ToLower(desiredRegion)

	/*
	 * [Process]
	 */

	// 1. Recommend vNet and subnets (Note: vNet can be a VPC or a VNet depending on the CSP)
	recommendedNetwork, err := RecommendNetwork(csp, region, srcInfra)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend a network for the source computing infrastructure")
		return emptyResp, err
	}

	// * Set a name to indicate a dependency between resources.
	recommendedNetwork.Name = "mig-vnet-01"
	recommendedNetwork.Description = "a recommended vNet for migration"
	for i, _ := range recommendedNetwork.SubnetInfoList {
		recommendedNetwork.SubnetInfoList[i].Name = fmt.Sprintf("mig-subnet-%02d", i+1)
		recommendedNetwork.SubnetInfoList[i].Description = "a recommended subnet for migration"
	}

	// Assign the recommended network to the response body
	recommendedVmInfra.TargetNetwork = recommendedNetwork

	// ? When I register my spec and OS images, can I use custom names for them?

	// 2. Recommend VM specs, OS images, and security groups, and
	// recommend VMs by removing duplicates of VM specs, OS images, and security groups and specifying them.
	var recommendedVmList = []tbmodel.TbVmReq{}
	var recommendedVmSpecList = []tbmodel.TbSpecReq{}
	var recommendedVmOsImageList = []tbmodel.TbImageReq{}
	var recommendedSecurityGroupList = []tbmodel.TbSecurityGroupReq{}

	for i, server := range srcInfra.Servers {

		/*
		 * Recommend VM specs, OS images, and security groups
		 */
		// Lookup the appropriate VM specs for the server
		vmSpecList, _, err := RecommendVmSpecs(csp, region, server, max)
		if err != nil {
			log.Warn().Msgf("failed to recommend VM specs for server %s: %v", server.Hostname, err)
			continue
		}

		// Lookup the appropriate VM OS images for the server
		vmOsImageInfoList := []tbmodel.TbImageInfo{}
		for _, vmSpec := range vmSpecList {
			osImgInfo, err := RecommendVmOsImageInfo(csp, region, server, vmSpec.Id)
			if err != nil {
				log.Warn().Msgf("failed to recommend VM OS image for server %s: %v", server.Hostname, err)
				vmOsImageInfoList = append(vmOsImageInfoList, tbmodel.TbImageInfo{})
			}
			vmOsImageInfoList = append(vmOsImageInfoList, osImgInfo)
		}

		// Generete security group from the server's firewall rules (or firewall table)
		sg, err := RecommendSecurityGroup(csp, region, server)
		if err != nil {
			log.Warn().Msgf("failed to recommend security group for server %s: %v", server.Hostname, err)
			continue
		}

		/*
		 * Check and deduplicate the recommended VM specs, OS images, and security groups
		 */

		// Deduplicate the recommended VM specs
		var tempVmSpec = tbmodel.TbSpecReq{}
		exists := false
		if len(vmSpecList) > 0 {
			for _, vmSpec := range recommendedVmSpecList {
				if vmSpec.CspSpecName == vmSpecList[0].CspSpecName {
					exists = true
					tempVmSpec = vmSpec
					break
				}
			}
		}
		if !exists && len(vmSpecList) > 0 {
			tempVmSpec = tbmodel.TbSpecReq{
				Name:           fmt.Sprintf("mig-spec-%02d", len(recommendedVmSpecList)+1),
				ConnectionName: fmt.Sprintf("%s-%s", csp, region),
				CspSpecName:    vmSpecList[0].CspSpecName,
				Description:    fmt.Sprintf("Recommended VM spec for %s", server.Hostname),
			}
			recommendedVmSpecList = append(recommendedVmSpecList, tempVmSpec)
		}

		// Deduplicate the recommended VM OS images
		var tempVmOsImage = tbmodel.TbImageReq{}
		exists = false
		if len(vmOsImageInfoList) > 0 && vmOsImageInfoList[0].CspImageName != "" {
			for _, vmOsImage := range recommendedVmOsImageList {
				if vmOsImage.CspImageName == vmOsImageInfoList[0].CspImageName {
					exists = true
					tempVmOsImage = vmOsImage
					break
				}
			}
		}
		if !exists && len(vmOsImageInfoList) > 0 && vmOsImageInfoList[0].CspImageName != "" {
			tempVmOsImage = tbmodel.TbImageReq{
				Name:           fmt.Sprintf("mig-image-%02d", len(recommendedVmOsImageList)+1),
				ConnectionName: fmt.Sprintf("%s-%s", csp, region),
				CspImageName:   vmOsImageInfoList[0].CspImageName,
				Description:    fmt.Sprintf("Recommended OS image for %s", server.Hostname),
			}
			recommendedVmOsImageList = append(recommendedVmOsImageList, tempVmOsImage)
		}

		// Deduplicate the recommended security groups
		exists = false
		for _, recommendedSg := range recommendedSecurityGroupList {
			// Both SGs have rules defined
			if recommendedSg.FirewallRules != nil && sg.FirewallRules != nil {
				// Quick check if they have the same number of rules
				if len(*recommendedSg.FirewallRules) == len(*sg.FirewallRules) {
					areAllRulesSame := true

					// Create maps for each rule in both security groups for comparison
					sgRulesMap := make(map[string]bool)
					for _, rule := range *sg.FirewallRules {
						// Create a unique key for each rule
						key := fmt.Sprintf("%s-%s-%s-%s-%s",
							rule.Direction, rule.IPProtocol, rule.CIDR, rule.FromPort, rule.ToPort)
						sgRulesMap[key] = true
					}

					// Check if all rules in the recommended SG exist in the current SG
					for _, rule := range *recommendedSg.FirewallRules {
						key := fmt.Sprintf("%s-%s-%s-%s-%s",
							rule.Direction, rule.IPProtocol, rule.CIDR, rule.FromPort, rule.ToPort)
						if !sgRulesMap[key] {
							areAllRulesSame = false
							break
						}
					}

					if areAllRulesSame {
						exists = true
						sg = recommendedSg
						break
					}
				}
			}
		}

		if !exists {
			// If the security group does not exist, set a name to indicate a dependency between resources.
			sg.Name = fmt.Sprintf("mig-sg-%02d", len(recommendedSecurityGroupList)+1)
			sg.ConnectionName = fmt.Sprintf("%s-%s", csp, region)
			sg.Description = fmt.Sprintf("Recommended security group for %s", server.Hostname)

			// * Set a name to indicate a dependency between resources.
			sg.VNetId = recommendedVmInfra.TargetNetwork.Name // Set the vNet ID to the security group

			// Set the security group to the response body
			recommendedSecurityGroupList = append(recommendedSecurityGroupList, sg)
		}

		/*
		 * Recommend VM by specifying the recommended VM specs, OS images, and security groups
		 */

		// * Set names to indicate a dependency between resources.
		tempVmReq := tbmodel.TbVmReq{
			ConnectionName:   "",
			Description:      fmt.Sprintf("a recommended virtual machine %02d for %s", i+1, server.Hostname),
			SpecId:           tempVmSpec.Name,
			ImageId:          tempVmOsImage.Name,
			VNetId:           recommendedVmInfra.TargetNetwork.Name,
			SubnetId:         recommendedVmInfra.TargetNetwork.SubnetInfoList[0].Name, // Use the first subnet for simplicity
			SecurityGroupIds: []string{sg.Name},                                       // Set the security group ID
			Name:             fmt.Sprintf("migrated-%s", server.Hostname),
			RootDiskSize:     "", // TBD
			RootDiskType:     "", // TBD
			SshKeyId:         "", // TBD: Set the SSH key ID if needed
			VmUserName:       "", // TBD: Set the VM user name if needed
			VmUserPassword:   "", // TBD
			SubGroupSize:     "", // TBD
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

func RecommendNetwork(csp string, region string, srcInfra inframodel.OnpremInfra) (tbmodel.TbVNetReq, error) {

	var emptyRes tbmodel.TbVNetReq
	var recommandedNetwork tbmodel.TbVNetReq

	// [Input]

	ok, err := IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid csp (%s) or region (%s)", csp, region)
		return emptyRes, err
	}

	// ! TBD: Validate req if needed
	//

	// * To be updated, the network in onpremise model.
	// srcInfra.Network.IPv4Networks

	// ! It's a dummy data. It should be replaced with the actual model.

	cidrBlocks := []string{
		"192.168.0.0/24",
		"192.168.1.0/24",
	}
	srcNetworks := cidrBlocks

	// [Process] Recommend the network vNet and subnets
	// * Note:
	// * At least 1 subnet is required.
	// * Derive a super network that includes user's all networks and set it as a vNet
	// * Set user's networks as subnets

	// ? Assumption: a network in on-premise infrastructure is designed and configured with various network segments or types.
	// * Thus, it must be selected which of these network segments will be the vNet.
	// ? If so, is grouping the network segments required?

	// Categorizes the entered CIDR blocks by private network (i.e., 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16)
	cidrs10 := []string{}
	cidrs172 := []string{}
	cidrs192 := []string{}

	for _, srcNetwork := range srcNetworks {
		identifiedNet, err := netutil.WhichPrivateNetworkByCidr(srcNetwork)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to identify the network %s", srcNetwork)
			continue
		}

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

	// Select a super network for the vNet
	// ? But how to select the super network?
	// * Currrently, a list of recommended networks is returned.

	recommendedNetworks := []RecommendedNetwork{}
	if supernet10 != "" {

		// Set subnets by the CIDR blocks from the source computing infra
		subnets := []tbmodel.TbSubnetReq{}
		for i, cidr := range cidrs10 {
			subnets = append(subnets, tbmodel.TbSubnetReq{
				Name:        fmt.Sprintf("subnet-%02d", i+1),
				Description: "subnet from source computing infra",
				IPv4_CIDR:   cidr,
			})
		}

		// Set the calculated supernet as the vNet
		tempNetworkInfo := RecommendedNetwork{
			Status:      "",
			Description: "Recommended network information",
			TargetNetwork: tbmodel.TbVNetReq{
				Name:           "vnet-01",
				Description:    "Recommended vNet for " + netutil.PrivateNetwork10Dot,
				CidrBlock:      supernet10,
				SubnetInfoList: subnets,
			},
		}

		// Append recommended network info to the list
		recommendedNetworks = append(recommendedNetworks, tempNetworkInfo)
	}

	if supernet172 != "" {

		// Set subnets by the CIDR blocks from the source computing infra
		subnets := []tbmodel.TbSubnetReq{}
		for i, cidr := range cidrs172 {
			subnets = append(subnets, tbmodel.TbSubnetReq{
				Name:        fmt.Sprintf("subnet-%02d", i+1),
				Description: "subnet from source computing infra",
				IPv4_CIDR:   cidr,
			})
		}

		tempNetworkInfo := RecommendedNetwork{
			Status:      "",
			Description: "Recommended network information",
			TargetNetwork: tbmodel.TbVNetReq{
				Name:           "vnet-01",
				Description:    "Recommended vNet for " + netutil.PrivateNetwork172Dot,
				CidrBlock:      supernet172,
				SubnetInfoList: subnets,
			},
		}
		// Append recommended network info to the list
		recommendedNetworks = append(recommendedNetworks, tempNetworkInfo)
	}

	if supernet192 != "" {

		// Set subnets by the CIDR blocks from the source computing infra
		subnets := []tbmodel.TbSubnetReq{}
		for i, cidr := range cidrs192 {
			subnets = append(subnets, tbmodel.TbSubnetReq{
				Name:        fmt.Sprintf("subnet-%02d", i+1),
				Description: "subnet from source computing infra",
				IPv4_CIDR:   cidr,
			})
		}

		// Set the calculated supernet as the vNet
		tempNetworkInfo := RecommendedNetwork{
			Status:      "",
			Description: "Recommended network information",
			TargetNetwork: tbmodel.TbVNetReq{
				Name:           "vnet-01",
				Description:    "Recommended vNet for " + netutil.PrivateNetwork192Dot,
				CidrBlock:      supernet192,
				SubnetInfoList: subnets,
			},
		}

		// Append recommended network info to the list
		recommendedNetworks = append(recommendedNetworks, tempNetworkInfo)
	}

	// [Output]
	if len(recommendedNetworks) == 0 {
		return emptyRes, fmt.Errorf("no recommended network found for the source computing infra")
	}

	// Select the first recommended network as the target network
	recommandedNetwork = recommendedNetworks[0].TargetNetwork
	log.Debug().Msgf("recommended network: %+v", recommandedNetwork)

	return recommandedNetwork, nil
}

// RecommendVmSpecs recommends appropriate VM specs for the given server
func RecommendVmSpecs(csp string, region string, server inframodel.ServerProperty, max int) (specList []tbmodel.TbSpecInfo, length int, err error) {

	var emptyResp []tbmodel.TbSpecInfo
	vmSpecInfoList := []tbmodel.TbSpecInfo{}

	if max <= 0 {
		err := fmt.Errorf("invalid max value: %d, set default: 5", max)
		log.Warn().Msgf(err.Error())
		max = 5
	}

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	// Set a deployment plan to recommand virtual machines
	// [Note]
	// * ">=" means greater than or equal to the operand
	// * "<=" means less than or equal to the operand
	// Ref: https://github.com/cloud-barista/cb-tumblebug/discussions/1234
	planDocstring := `{
	"filter": {
		"policy": [
			{
				"condition": [
					{
						"operand": "%d",
						"operator": ">="
					},
					{
						"operand": "%d",
						"operator": "<="
					}
				],
				"metric": "vCPU"
			},
			{
				"condition": [
					{
						"operand": "%d",
						"operator": ">="
					},
					{
						"operand": "%d",
						"operator": "<="
					}
				],
				"metric": "memoryGiB"
			},
			{
				"condition": [
					{
						"operand": "%s"
					}
				],
				"metric": "providerName"
			},
			{
				"condition": [
					{
						"operand": "%s"
					}
				],
				"metric": "regionName"
			}
		]
	},
	"limit": "5",
	"priority": {
		"policy": [
			{
				"metric": "performance",
				"weight": "0.5"
			}
		]
	}
}`

	// Get server info from source computing infra info
	cores := server.CPU.Cores
	// memory := MBtoGiB(float64(server.Memory.TotalSize))
	memory := uint32(server.Memory.TotalSize)

	coresMax := cores << 1
	var coresMin uint32
	if cores > 1 {
		coresMin = cores >> 1
	} else {
		coresMin = 1
	}

	memoryMax := memory << 1
	var memoryMin uint32
	if memory > 1 {
		memoryMin = memory >> 1
	} else {
		memoryMin = 1
	}

	providerName := csp
	regionName := region

	osNameAndVersion := server.OS.Name + " " + server.OS.Version
	osNameWithVersion := strings.ToLower(osNameAndVersion)

	log.Debug().
		Uint32("coreLowerLimit", coresMin).
		Uint32("coreUpperLimit", coresMax).
		Uint32("memoryLowerLimit (GiB)", memoryMin).
		Uint32("memoryUpperLimit (GiB)", memoryMax).
		Str("providerName", providerName).
		Str("regionName", regionName).
		Str("osNameWithVersion", osNameWithVersion).
		Msg("Source computing infrastructure info")

	// Set a deployment plan to search VMs having appropriate specs
	planToSearchProperVm := fmt.Sprintf(planDocstring,
		coresMin,
		coresMax,
		memoryMin,
		memoryMax,
		providerName,
		regionName,
	)
	log.Trace().Msgf("deployment plan to search proper VMs: %s", planToSearchProperVm)

	// Lookup VM specs
	method := "POST"
	url := fmt.Sprintf("%s/mciRecommendVm", epTumblebug)

	// Request body
	reqRecommVm := new(tbmodel.DeploymentPlan)
	err = json.Unmarshal([]byte(planToSearchProperVm), reqRecommVm)
	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyResp, -1, err
	}
	log.Trace().Msgf("deployment plan for the VM recommendation: %+v", reqRecommVm)

	// Response body

	err = common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(*reqRecommVm),
		reqRecommVm,
		&vmSpecInfoList,
		common.VeryShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyResp, -1, err
	}

	numOfVmSpecs := len(vmSpecInfoList)
	log.Debug().Msgf("the number of recommended VM specs: %d (for the inserted PM/VM with spec (cores: %d, memory (GiB): %d))", numOfVmSpecs, cores, memory)
	log.Trace().Msgf("recommendedVmList for the inserted PM/VM with spec (cores: %d, memory (GiB): %d): %+v", cores, memory, vmSpecInfoList)

	if numOfVmSpecs == 0 {
		err := fmt.Errorf("no VM spec recommended for the inserted PM/VM with spec (cores: %d, memory (GiB): %d)", cores, memory)
		log.Warn().Msgf(err.Error())
		return emptyResp, -1, err
	}

	vmSpecInfoList = vmSpecInfoList[:max]

	return vmSpecInfoList, numOfVmSpecs, nil
}

// RecommendVmOsImageInfo recommends an appropriate VM OS image (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImageInfo(csp string, region string, server inframodel.ServerProperty, vmSpecId string) (tbmodel.TbImageInfo, error) {

	var emptyRes = tbmodel.TbImageInfo{}
	var vmOsImage = tbmodel.TbImageInfo{}

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl
	method := "POST"
	url := fmt.Sprintf("%s/mciDynamicCheckRequest", epTumblebug)

	// Request body
	reqMciDynamicCheck := new(tbmodel.MciConnectionConfigCandidatesReq)
	reqMciDynamicCheck.CommonSpecs = []string{vmSpecId}

	// Response body
	resMciDynamicCheck := new(tbmodel.CheckMciDynamicReqInfo)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(*reqMciDynamicCheck),
		reqMciDynamicCheck,
		resMciDynamicCheck,
		common.VeryShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyRes, err
	}

	// for pretty logging
	prettyImages, err := json.MarshalIndent(resMciDynamicCheck.ReqCheck[0].Image, "", "  ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
		return emptyRes, err
	}
	log.Trace().Msgf("resMciDynamicCheck.ReqCheck[0].Image: %s", prettyImages)

	if len(resMciDynamicCheck.ReqCheck) == 0 || len(resMciDynamicCheck.ReqCheck[0].Image) == 0 {
		err := fmt.Errorf("no VM OS image recommended for the inserted PM/VM")
		log.Warn().Msgf(err.Error())
		return emptyRes, err
	}

	keywords := fmt.Sprintf("%s %s %s %s",
		server.OS.Name,
		server.OS.Version,
		server.CPU.Architecture,
		server.RootDisk.Type)
	log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

	// Select VM OS image via LevenshteinDistance-based text similarity
	delimiters1 := []string{" ", "-", "_", ",", "(", ")", "[", "]", "/"}
	delimiters2 := delimiters1
	vmOsImage = FindBestVmOsImage(keywords, delimiters1, resMciDynamicCheck.ReqCheck[0].Image, delimiters2)

	return vmOsImage, nil
}

// RecommendVmOsImageId recommends an appropriate VM OS image (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImageId(csp string, region string, server inframodel.ServerProperty, vmSpecId string) (string, error) {

	var emptyRes string = ""
	var vmOsImageId string = ""

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl
	method := "POST"
	url := fmt.Sprintf("%s/mciDynamicCheckRequest", epTumblebug)

	// Request body
	reqMciDynamicCheck := new(tbmodel.MciConnectionConfigCandidatesReq)
	reqMciDynamicCheck.CommonSpecs = []string{vmSpecId}

	// Response body
	resMciDynamicCheck := new(tbmodel.CheckMciDynamicReqInfo)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(*reqMciDynamicCheck),
		reqMciDynamicCheck,
		resMciDynamicCheck,
		common.VeryShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyRes, err
	}

	// for pretty logging
	prettyImages, err := json.MarshalIndent(resMciDynamicCheck.ReqCheck[0].Image, "", "  ")
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
		return emptyRes, err
	}
	log.Trace().Msgf("resMciDynamicCheck.ReqCheck[0].Image: %s", prettyImages)

	if len(resMciDynamicCheck.ReqCheck) == 0 || len(resMciDynamicCheck.ReqCheck[0].Image) == 0 {
		err := fmt.Errorf("no VM OS image recommended for the inserted PM/VM")
		log.Warn().Msgf(err.Error())
		return emptyRes, err
	}

	keywords := fmt.Sprintf("%s %s %s %s",
		server.OS.Name,
		server.OS.Version,
		server.CPU.Architecture,
		server.RootDisk.Type)
	log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

	// Select VM OS image via LevenshteinDistance-based text similarity
	delimiters1 := []string{" ", "-", "_", ",", "(", ")", "[", "]", "/"}
	delimiters2 := delimiters1
	vmOsImageId = FindBestVmOsImageId(keywords, delimiters1, resMciDynamicCheck.ReqCheck[0].Image, delimiters2)

	return vmOsImageId, nil
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
func checkOverallVmStatus(vms []tbmodel.TbVmDynamicReq) string {
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

// FindBestVmOsImage finds the best matching image based on the similarity scores
func FindBestVmOsImage(keywords string, kwDelimiters []string, vmImages []tbmodel.TbImageInfo, imgDelimiters []string) tbmodel.TbImageInfo {

	var bestVmOsImageID tbmodel.TbImageInfo
	var highestScore float64 = 0.0

	for _, image := range vmImages {
		score := similarity.CalcResourceSimilarity(keywords, kwDelimiters, image.GuestOS, imgDelimiters)
		if score > highestScore {
			highestScore = score
			bestVmOsImageID = image
		}
		// log.Debug().Msgf("VmImageName: %s, score: %f, description: %s", image.GuestOS, score, image.Description)

	}
	log.Debug().Msgf("bestVmOsImageID: %s, highestScore: %f", bestVmOsImageID, highestScore)

	return bestVmOsImageID
}

// FindBestVmOsImageId finds the best matching image based on the similarity scores
func FindBestVmOsImageId(keywords string, kwDelimiters []string, vmImages []tbmodel.TbImageInfo, imgDelimiters []string) string {

	var bestVmOsImageID string
	var highestScore float64 = 0.0

	for _, image := range vmImages {
		score := similarity.CalcResourceSimilarity(keywords, kwDelimiters, image.GuestOS, imgDelimiters)
		if score > highestScore {
			highestScore = score
			bestVmOsImageID = image.Id
		}
		// log.Debug().Msgf("VmImageName: %s, score: %f, description: %s", image.GuestOS, score, image.Description)

	}
	log.Debug().Msgf("bestVmOsImageID: %s, highestScore: %f", bestVmOsImageID, highestScore)

	return bestVmOsImageID
}

func RecommendSecurityGroup(csp string, region string, server inframodel.ServerProperty) (tbmodel.TbSecurityGroupReq, error) {

	var emptyRes = tbmodel.TbSecurityGroupReq{}
	var recommendedSecurityGroup = tbmodel.TbSecurityGroupReq{}

	// [Input]
	ok, err := IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msgf("invalid provider (%s) or region (%s)", csp, region)
		return emptyRes, err
	}

	// ! To be updated, the security group in onpremise model.
	// server.FirewallRules = dummyFirewallRules

	firewallRules := dummyFirewallRules

	// Use the provided firewall rules or fall back to dummy data if empty
	if len(firewallRules) == 0 {
		log.Warn().Msg("no firewall rules provided, using sample data")
		firewallRules = dummyFirewallRules
	}

	// [Process] Recommend the security group

	// Create security group recommendations
	// ! To be updated with the actual model and real data
	// ! A list of firewall rules(i.e., firewall table) will be entered (currently, it's a dummy single firewall table)

	sgRules := generateSecurityGroupRules(firewallRules)

	// [Output]
	// Create a security group for all rules
	recommendedSecurityGroup = tbmodel.TbSecurityGroupReq{
		Name:          "mig-sg-01",
		Description:   "Recommended security group for migration",
		FirewallRules: sgRules,
	}

	log.Debug().Msgf("recommendedSecurityGroup: %+v\n")

	return recommendedSecurityGroup, nil
}

// ! To be replaced with the actual model
type FirewallRuleProperty struct { // note: reference command `sudo ufw status verbose`
	SrcCIDR   string `json:"srcCIDR,omitempty"`
	DstCIDR   string `json:"dstCIDR,omitempty"`
	SrcPorts  string `json:"srcPorts,omitempty"`
	DstPorts  string `json:"dstPorts,omitempty"`
	Protocol  string `json:"protocol,omitempty"`  // TCP, UDP, ICMP
	Direction string `json:"direction,omitempty"` // inbound, outbound
	Action    string `json:"action,omitempty"`    // allow, deny
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
func generateSecurityGroupRules(rules []FirewallRuleProperty) *[]tbmodel.TbFirewallRuleInfo {
	var tbRules []tbmodel.TbFirewallRuleInfo

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

			// * NOTE: 3 cases for destination ports
			// 1. Comma-separated ports (e.g., 22,23,24)
			// 2. Port range with colon notation (e.g., 30000:40000)
			// 3. Single port (e.g., 22)

			// Handle destination ports based on format
			if rule.DstPorts == "" {
				// Skip rules without port information
				log.Debug().Msgf("Skipping inbound rule without port information: %+v", rule)
				continue

			} else if strings.Contains(rule.DstPorts, ",") {
				// Handle comma-separated ports (e.g., 22,23,24) - create multiple rules
				ports := strings.Split(rule.DstPorts, ",")

				for _, port := range ports {
					portTrimmed := strings.TrimSpace(port)
					tbRule := tbmodel.TbFirewallRuleInfo{
						Direction:  rule.Direction,
						IPProtocol: rule.Protocol,
						CIDR:       srcCIDR,
						FromPort:   portTrimmed,
						ToPort:     portTrimmed,
					}
					tbRules = append(tbRules, tbRule)
					log.Debug().Msgf("Created inbound rule for comma-separated port %s: %+v", portTrimmed, tbRule)
				}

			} else if strings.Contains(rule.DstPorts, ":") {
				// Handle port range with colon notation (e.g., 30000:40000)
				portRange := strings.Split(rule.DstPorts, ":")
				if len(portRange) == 2 {
					tbRule := tbmodel.TbFirewallRuleInfo{
						Direction:  rule.Direction,
						IPProtocol: rule.Protocol,
						CIDR:       srcCIDR,
						FromPort:   strings.TrimSpace(portRange[0]),
						ToPort:     strings.TrimSpace(portRange[1]),
					}
					tbRules = append(tbRules, tbRule)
					log.Debug().Msgf("Created inbound rule for port range %s: %+v", rule.DstPorts, tbRule)
				} else {
					log.Warn().Msgf("Invalid port range format in rule.DstPorts: %s - skipping rule", rule.DstPorts)
					continue
				}
			} else {
				// Handle single port
				tbRule := tbmodel.TbFirewallRuleInfo{
					Direction:  rule.Direction,
					IPProtocol: rule.Protocol,
					CIDR:       srcCIDR,
					FromPort:   rule.DstPorts,
					ToPort:     rule.DstPorts,
				}
				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created inbound rule for single port %s: %+v", rule.DstPorts, tbRule)
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

			// Now handle the ports similar to inbound case
			if rule.DstPorts == "" {
				// Skip rules without port information
				log.Debug().Msgf("Skipping outbound rule without port information: %+v", rule)
				continue
			} else if strings.Contains(rule.DstPorts, ",") {
				// Handle comma-separated ports - create multiple rules
				ports := strings.Split(rule.DstPorts, ",")

				for _, port := range ports {
					portTrimmed := strings.TrimSpace(port)
					tbRule := tbmodel.TbFirewallRuleInfo{
						Direction:  rule.Direction,
						IPProtocol: rule.Protocol,
						CIDR:       dstCIDR,
						FromPort:   portTrimmed,
						ToPort:     portTrimmed,
					}
					tbRules = append(tbRules, tbRule)
					log.Debug().Msgf("Created outbound rule for comma-separated port %s: %+v", portTrimmed, tbRule)
				}
			} else if strings.Contains(rule.DstPorts, ":") {
				// Handle port range with colon notation
				portRange := strings.Split(rule.DstPorts, ":")
				if len(portRange) == 2 {
					tbRule := tbmodel.TbFirewallRuleInfo{
						Direction:  rule.Direction,
						IPProtocol: rule.Protocol,
						CIDR:       dstCIDR,
						FromPort:   strings.TrimSpace(portRange[0]),
						ToPort:     strings.TrimSpace(portRange[1]),
					}
					tbRules = append(tbRules, tbRule)
					log.Debug().Msgf("Created outbound rule for port range %s: %+v", rule.DstPorts, tbRule)
				} else {
					log.Warn().Msgf("Invalid port range format: %s - skipping rule", rule.DstPorts)
					continue
				}
			} else {
				// Handle single port
				tbRule := tbmodel.TbFirewallRuleInfo{
					Direction:  rule.Direction,
					IPProtocol: rule.Protocol,
					CIDR:       dstCIDR,
					FromPort:   rule.DstPorts,
					ToPort:     rule.DstPorts,
				}
				tbRules = append(tbRules, tbRule)
				log.Debug().Msgf("Created outbound rule for single port %s: %+v", rule.DstPorts, tbRule)
			}

		default:
			log.Warn().Msgf("Unknown direction '%s' in rule: %+v", rule.Direction, rule)
		}

		log.Debug().Msgf("Original FirewallRule: %+v", rule)
	}

	return &tbRules
}

var dummyFirewallRules = []FirewallRuleProperty{
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "22",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "80,443",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "10.0.0.0/16",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "3306",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "172.16.0.0/12",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "5432",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "25",
		Protocol:  "TCP",
		Direction: "inbound",
		Action:    "deny",
	},
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "53",
		Protocol:  "UDP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "0.0.0.0/0",
		DstCIDR:   "192.168.1.10/32",
		DstPorts:  "1194",
		Protocol:  "UDP",
		Direction: "inbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "192.168.1.10/32",
		DstCIDR:   "0.0.0.0/0",
		SrcPorts:  "32768-60999",
		Protocol:  "TCP",
		Direction: "outbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "192.168.1.10/32",
		DstCIDR:   "8.8.8.8/32",
		DstPorts:  "53",
		Protocol:  "UDP",
		Direction: "outbound",
		Action:    "allow",
	},
	{
		SrcCIDR:   "192.168.1.10/32",
		DstCIDR:   "0.0.0.0/0",
		Protocol:  "ICMP",
		Direction: "outbound",
		Action:    "allow",
	},
}
