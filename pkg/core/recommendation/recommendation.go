package recommendation

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cloud-barista/cb-tumblebug/src/core/mcir"
	"github.com/cloud-barista/cb-tumblebug/src/core/mcis"
	cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

func Recommend(srcInfra []infra.Infra) (cloudmodel.InfraMigrationReq, error) {

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := os.Getenv("API_USERNAME")
	apiPass := os.Getenv("API_PASSWORD")
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
		return cloudmodel.InfraMigrationReq{}, err
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
	targetInfra := cloudmodel.InfraMigrationReq{
		Description:     "A cloud infra recommended by CM-Beetle",
		InstallMonAgent: "no",
		Label:           "DynamicVM",
		Name:            "",
		SystemLabel:     "",
		Vm:              []cloudmodel.HostMigrationReq{},
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
			return cloudmodel.InfraMigrationReq{}, err
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
			return cloudmodel.InfraMigrationReq{}, err
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

		name := fmt.Sprintf("rehosted-%s-%s", server.Compute.OS.Node.Hostname, server.Compute.OS.Node.Machineid)

		////////////////////////////////////////
		// Search and set target OS image (e.g. ubuntu22.04)
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
			return cloudmodel.InfraMigrationReq{}, err
		}

		log.Trace().Msgf("resMcisDynamicCheck: %+v", resMcisDynamicCheck)

		// candidateImages := resMcisDynamicCheck.ReqCheck[0].Image

		// [TBD] Select VM image by TextSimilarity

		image := fmt.Sprintf("%s+%s+%s", providerName, regionName, osNameWithVersion)

		vm := cloudmodel.HostMigrationReq{
			ConnectionName: "",
			CommonImage:    image,
			CommonSpec:     recommendedSpec,
			Description:    "a recommended virtual machine",
			Label:          "rehosted",
			Name:           name,
			RootDiskSize:   "default",
			RootDiskType:   "default",
			SubGroupSize:   "1",
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
