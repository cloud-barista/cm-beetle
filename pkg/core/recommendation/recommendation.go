package recommendation

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloud-barista/cb-tumblebug/src/core/mcir"
	"github.com/cloud-barista/cb-tumblebug/src/core/mcis"

	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"

	"github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/infra"
	// "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/strcomp"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// func Recommend(srcInfra []infra.Infra) (cloudmodel.InfraMigrationReq, error) {
func Recommend(srcInfra []infra.Infra) (mcis.TbMcisDynamicReq, error) {

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := viper.GetString("api.username")
	apiPass := viper.GetString("api.password")
	client.SetBasicAuth(apiUser, apiPass)

	// set endpoint
	epTumblebug := common.TumblebugRestUrl

	// check readyz
	method := "GET"
	url := fmt.Sprintf("%s/readyz", epTumblebug)
	reqReadyz := common.NoBody
	resReadyz := new(common.SimpleMsg)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(reqReadyz),
		&reqReadyz,
		resReadyz,
		common.VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return mcis.TbMcisDynamicReq{}, err
	}
	log.Debug().Msgf("resReadyz: %+v", resReadyz.Message)

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

	// A target infrastructure by recommendation
	targetInfra := mcis.TbMcisDynamicReq{
		Description:     "A cloud infra recommended by CM-Beetle",
		InstallMonAgent: "no",
		Label:           "rehosted-infra",
		Name:            "",
		SystemLabel:     "",
		Vm:              []mcis.TbVmDynamicReq{},
	}

	// Recommand VMs
	for _, server := range srcInfra {

		// Extract server info from source computing infra info
		cores := server.Compute.ComputeResource.CPU.Cores
		memory := MBtoGiB(float64(server.Compute.ComputeResource.Memory.Size))

		coreUpperLimit := cores << 1
		var coreLowerLimit uint
		if cores > 1 {
			coreLowerLimit = cores >> 1
		} else {
			coreLowerLimit = 1
		}

		memoryUpperLimit := memory << 1
		var memoryLowerLimit uint32
		if memory > 1 {
			memoryLowerLimit = memory >> 1
		} else {
			memoryLowerLimit = 1
		}

		providerName := "aws"
		regionName := "ap-northeast-2"

		osVendor := server.Compute.OS.OS.Vendor
		osVersion := server.Compute.OS.OS.Release
		osNameWithVersion := strings.ToLower(osVendor + osVersion)

		log.Debug().
			Uint("coreUpperLimit", coreUpperLimit).
			Uint("coreLowerLimit", coreLowerLimit).
			Uint32("memoryUpperLimit (GiB)", memoryUpperLimit).
			Uint32("memoryLowerLimit (GiB)", memoryLowerLimit).
			Str("providerName", providerName).
			Str("regionName", regionName).
			Str("osNameWithVersion", osNameWithVersion).
			Msg("Source computing infrastructure info")

		// To search proper VMs with the server info, set a deployment plan
		planToSearchProperVm := fmt.Sprintf(planDocstring,
			coreLowerLimit,
			coreUpperLimit,
			memoryLowerLimit,
			memoryUpperLimit,
			providerName,
			regionName,
		)

		////////////////////////////////////////
		// Search and set a target VM spec
		method := "POST"
		url := fmt.Sprintf("%s/mcisRecommendVm", epTumblebug)

		// Request body
		reqRecommVm := new(mcis.DeploymentPlan)
		err := json.Unmarshal([]byte(planToSearchProperVm), reqRecommVm)
		if err != nil {
			log.Err(err).Msg("")
			return mcis.TbMcisDynamicReq{}, err
		}
		log.Trace().Msgf("deployment plan for the VM recommendation: %+v", reqRecommVm)

		// Response body
		resRecommVmList := []mcir.TbSpecInfo{}

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
			log.Err(err).Msg("")
			return mcis.TbMcisDynamicReq{}, err
		}

		numRecommenedVm := len(resRecommVmList)

		log.Debug().Msgf("the number of recommended VM specs: %d (for the inserted PM/VM with spec (cores: %d, memory (GiB): %d))", numRecommenedVm, cores, memory)
		log.Trace().Msgf("recommendedVmList for the inserted PM/VM with spec (cores: %d, memory (GiB): %d): %+v", cores, memory, resRecommVmList)

		if numRecommenedVm == 0 {
			log.Warn().Msgf("no VM spec recommended for the inserted PM/VM with spec (cores: %d, memory (GiB): %d)", cores, memory)
			continue
		}
		log.Debug().Msgf("select the 1st recommended virtual machine: %+v", resRecommVmList[0])
		recommendedSpec := resRecommVmList[0].Id

		// name := fmt.Sprintf("rehosted-%s-%s", server.Compute.OS.Node.Hostname, server.Compute.OS.Node.Machineid)
		name := fmt.Sprintf("rehosted-%s", server.Compute.OS.Node.Hostname)

		////////////////////////////////////////
		// Search and set target VM image (e.g. ubuntu22.04)
		method = "POST"
		url = fmt.Sprintf("%s/mcisDynamicCheckRequest", epTumblebug)

		// Request body
		reqMcisDynamicCheck := new(mcis.McisConnectionConfigCandidatesReq)
		reqMcisDynamicCheck.CommonSpecs = []string{recommendedSpec}

		// Response body
		resMcisDynamicCheck := new(mcis.CheckMcisDynamicReqInfo)

		err = common.ExecuteHttpRequest(
			client,
			method,
			url,
			nil,
			common.SetUseBody(*reqMcisDynamicCheck),
			reqMcisDynamicCheck,
			resMcisDynamicCheck,
			common.VeryShortDuration,
		)

		if err != nil {
			log.Err(err).Msg("")
			return mcis.TbMcisDynamicReq{}, err
		}

		log.Trace().Msgf("resMcisDynamicCheck: %+v", resMcisDynamicCheck)

		if len(resMcisDynamicCheck.ReqCheck) == 0 {
			log.Warn().Msg("no VM OS image recommended for the inserted PM/VM")
			continue
		}

		keywords := fmt.Sprintf("%s %s %s %s",
			server.Compute.OS.OS.Vendor,
			server.Compute.OS.OS.Version,
			server.Compute.OS.OS.Architecture,
			server.Compute.ComputeResource.RootDisk.Type)
		log.Debug().Msg("keywords for the VM OS image recommendation: " + keywords)

		// Select VM OS image via LevenshteinDistance-based text similarity
		delimiters1 := []string{" ", "-", "_", ",", "(", ")", "[", "]", "/"}
		delimiters2 := delimiters1
		vmOsImageId := FindBestVmOsImage(keywords, delimiters1, resMcisDynamicCheck.ReqCheck[0].Image, delimiters2)

		// vmOsImage := fmt.Sprintf("%s+%s+%s", providerName, regionName, osNameWithVersion)

		vm := mcis.TbVmDynamicReq{
			ConnectionName: "",
			CommonImage:    vmOsImageId,
			CommonSpec:     recommendedSpec,
			Description:    "a recommended virtual machine",
			Label:          "rehosted-vm",
			Name:           name,
			RootDiskSize:   "",
			RootDiskType:   "",
			SubGroupSize:   "",
			VmUserPassword: "",
		}

		targetInfra.Vm = append(targetInfra.Vm, vm)
	}

	log.Trace().Msgf("targetInfra: %+v", targetInfra)

	return targetInfra, nil
}

func MBtoGiB(mb float64) uint32 {
	const bytesInMB = 1000000.0
	const bytesInGiB = 1073741824.0
	gib := (mb * bytesInMB) / bytesInGiB
	return uint32(gib)
}

// FindBestVmOsImage finds the best matching image based on the similarity scores
func FindBestVmOsImage(keywords string, kwDelimiters []string, vmImages []mcir.TbImageInfo, imgDelimiters []string) string {

	var bestVmOsImageID string
	var highestScore float64

	for _, image := range vmImages {
		score := strcomp.CalculateSimilarity(keywords, kwDelimiters, image.CspImageName, imgDelimiters)
		if score > highestScore {
			highestScore = score
			bestVmOsImageID = image.Id
		}
		log.Trace().Msgf("VmImageName: %s, score: %f", image.CspImageName, score)

	}
	log.Debug().Msgf("bestVmOsImageID: %s, highestScore: %f", bestVmOsImageID, highestScore)

	return bestVmOsImageID
}
