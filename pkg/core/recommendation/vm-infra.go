package recommendation

import (
	"encoding/json"
	"fmt"
	"strings"

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

// RecommendVmInfra an appropriate multi-cloud infrastructure (MCI) for cloud migration
func RecommendVmInfra(desiredCsp string, desiredRegion string, srcInfra inframodel.OnpremInfra) (RecommendedVmInfraInfoList, error) {

	// var emptyResp RecommendedVmInfraInfoList
	var recommendedVmInfraInfoList RecommendedVmInfraInfoList

	// ! To be updated, a user will input the desired number of recommended VMs
	var max int = 5
	// Initialize the response body
	recommendedVmInfraInfoList = RecommendedVmInfraInfoList{
		Description:       "This is a list of recommended target infrastructures. Please review and use them.",
		Count:             0,
		TargetVmInfraList: []RecommendedVmInfraInfo{},
	}

	// // Set VM info
	// recommendedVm := tbmodel.TbVmDynamicReq{
	// 	ConnectionName: "",
	// 	CommonImage:    "", // Lookup and set an appropriate VM OS image
	// 	CommonSpec:     "", // Lookup and set an appropriate VM spec
	// 	Description:    "a recommended virtual machine",
	// 	Name:           fmt.Sprintf("rehosted-%s", server.Hostname),
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
			osImgId, err := RecommendVmOsImage(desiredCsp, desiredRegion, server, vmSpec.Id)
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
	recommenedVmInfraInfoList := []RecommendedVmInfraInfo{}

	for i, vmInfoList := range transposed {

		tempVmInfraInfo := RecommendedVmInfraInfo{
			Status:      string(NothingRecommended),
			Description: "This is a recommended target infrastructure.",
			TargetVmInfra: tbmodel.TbMciDynamicReq{
				Name:        fmt.Sprintf("rehosted-%02d", i),
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
				Name:           fmt.Sprintf("rehosted-%s", srcInfra.Servers[j].Hostname),
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

// RecommendVmOsImage recommends an appropriate VM OS image (e.g., Ubuntu 22.04) for the given VM spec
func RecommendVmOsImage(csp string, region string, server inframodel.ServerProperty, vmSpecId string) (string, error) {

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
	vmOsImageId = FindBestVmOsImage(keywords, delimiters1, resMciDynamicCheck.ReqCheck[0].Image, delimiters2)

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
func FindBestVmOsImage(keywords string, kwDelimiters []string, vmImages []tbmodel.TbImageInfo, imgDelimiters []string) string {

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
