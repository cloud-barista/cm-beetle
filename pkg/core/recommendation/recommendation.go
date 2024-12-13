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

func IsValidProviderAndRegion(provider string, region string) (bool, error) {

	isValid := false

	providerName := strings.ToLower(provider)

	regionName := strings.ToLower(region)

	supportedCsp := isSupportedCSP(providerName)

	if !supportedCsp {
		err := fmt.Errorf("not supported yet (provider: %s)", providerName)
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
	url := fmt.Sprintf("%s/provider/%s/region/%s", epTumblebug, providerName, regionName)

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

// Recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration
func Recommend(desiredProvider string, desiredRegion string, srcInfra inframodel.OnpremInfra) (RecommendedInfraInfo, error) {

	var emptyResp RecommendedInfraInfo
	var recommendedInfraInfo RecommendedInfraInfo

	// Set target infra
	recommendedInfraInfo.TargetInfra = tbmodel.TbMciDynamicReq{
		Description:     "A cloud infra recommended by CM-Beetle",
		InstallMonAgent: "no",
		Name:            "",
		SystemLabel:     "",
		Vm:              []tbmodel.TbVmDynamicReq{},
	}

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	// Set a deployment plan to recommand virtual machines
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
				"metric": "performance"
			}
		]
	}
}`

	// Recommand VMs
	for _, server := range srcInfra.Servers {

		// Set VM info
		recommendedVm := tbmodel.TbVmDynamicReq{
			ConnectionName: "",
			CommonImage:    "", // Search and set an appropriate VM OS image
			CommonSpec:     "", // Search and set an appropriate VM spec
			Description:    "a recommended virtual machine",
			Name:           fmt.Sprintf("rehosted-%s", server.Hostname),
			RootDiskSize:   "", // TBD
			RootDiskType:   "", // TBD
			SubGroupSize:   "",
			VmUserPassword: "",
		}

		/*
			Search an appropriate VM spec for the server by /mciRecommendVm API
		*/

		// Extract server info from source computing infra info
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

		providerName := desiredProvider
		regionName := desiredRegion

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
		log.Debug().Msgf("deployment plan to search proper VMs: %s", planToSearchProperVm)

		// Search and set a target VM spec
		method := "POST"
		url := fmt.Sprintf("%s/mciRecommendVm", epTumblebug)

		// Request body
		reqRecommVm := new(tbmodel.DeploymentPlan)
		err := json.Unmarshal([]byte(planToSearchProperVm), reqRecommVm)
		if err != nil {
			log.Error().Err(err).Msg("")
			return emptyResp, err
		}
		log.Trace().Msgf("deployment plan for the VM recommendation: %+v", reqRecommVm)

		// Response body
		resRecommVmList := []tbmodel.TbSpecInfo{}

		err = common.ExecuteHttpRequest(
			client,
			method,
			url,
			nil,
			common.SetUseBody(*reqRecommVm),
			reqRecommVm,
			&resRecommVmList,
			common.VeryShortDuration,
		)

		if err != nil {
			log.Error().Err(err).Msg("")
			return emptyResp, err
		}

		numRecommenedVm := len(resRecommVmList)

		log.Debug().Msgf("the number of recommended VM specs: %d (for the inserted PM/VM with spec (cores: %d, memory (GiB): %d))", numRecommenedVm, cores, memory)
		log.Trace().Msgf("recommendedVmList for the inserted PM/VM with spec (cores: %d, memory (GiB): %d): %+v", cores, memory, resRecommVmList)

		if numRecommenedVm == 0 {
			log.Warn().Msgf("no VM spec recommended for the inserted PM/VM with spec (cores: %d, memory (GiB): %d)", cores, memory)
			recommendedInfraInfo.TargetInfra.Vm = append(recommendedInfraInfo.TargetInfra.Vm, recommendedVm)
			continue
		}
		log.Debug().Msgf("select the 1st recommended virtual machine: %+v", resRecommVmList[0])
		recommendedSpec := resRecommVmList[0].Id

		// Assign the recommended spec
		recommendedVm.CommonSpec = recommendedSpec

		/*
			Search an appropriate VM OS image for the server by /mciDynamicCheckRequest API
		*/

		// Search and set target VM image (e.g. ubuntu22.04)
		method = "POST"
		url = fmt.Sprintf("%s/mciDynamicCheckRequest", epTumblebug)

		// Request body
		reqMciDynamicCheck := new(tbmodel.MciConnectionConfigCandidatesReq)
		reqMciDynamicCheck.CommonSpecs = []string{recommendedSpec}

		// Response body
		resMciDynamicCheck := new(tbmodel.CheckMciDynamicReqInfo)

		err = common.ExecuteHttpRequest(
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
			return emptyResp, err
		}

		// for pretty logging
		prettyImages, err := json.MarshalIndent(resMciDynamicCheck.ReqCheck[0].Image, "", "  ")
		if err != nil {
			log.Error().Err(err).Msg("failed to marshal response")
			return emptyResp, err
		}
		log.Debug().Msgf("resMciDynamicCheck.ReqCheck[0].Image: %s", prettyImages)

		if len(resMciDynamicCheck.ReqCheck) == 0 {
			log.Warn().Msg("no VM OS image recommended for the inserted PM/VM")
			recommendedInfraInfo.TargetInfra.Vm = append(recommendedInfraInfo.TargetInfra.Vm, recommendedVm)
			continue
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
		vmOsImageId := FindBestVmOsImage(keywords, delimiters1, resMciDynamicCheck.ReqCheck[0].Image, delimiters2)

		// Assign the recommended VM OS image
		recommendedVm.CommonImage = vmOsImageId

		recommendedInfraInfo.TargetInfra.Vm = append(recommendedInfraInfo.TargetInfra.Vm, recommendedVm)
	}
	log.Debug().Msgf("the recommended infra info: %+v", recommendedInfraInfo)

	status := checkOverallVmStatus(recommendedInfraInfo.TargetInfra.Vm)
	recommendedInfraInfo.Status = status
	if status == string(NothingRecommended) {
		recommendedInfraInfo.Description = "Could not find approprate VMs."
	} else if status == string(FullyRecommended) {
		recommendedInfraInfo.Description = "Target infra is recommended."
	} else {
		recommendedInfraInfo.Description = "Some VMs are recommended. Please check and fill the required information."
	}

	return recommendedInfraInfo, nil
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
		log.Debug().Msgf("VmImageName: %s, score: %f, description: %s", image.GuestOS, score, image.Description)

	}
	log.Debug().Msgf("bestVmOsImageID: %s, highestScore: %f", bestVmOsImageID, highestScore)

	return bestVmOsImageID
}
