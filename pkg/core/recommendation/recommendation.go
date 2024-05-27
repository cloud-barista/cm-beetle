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
	planDocstring := `{
	"filter": {
		"policy": [
			{
				"condition": [
					{
						"operand": "%d",
						"operator": "<="
					},
					{
						"operand": "%d",
						"operator": ">="
					}
				],
				"metric": "vCPU"
			},
			{
				"condition": [
					{
						"operand": "%d",
						"operator": "<="
					},
					{
						"operand": "%d",
						"operator": ">="
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
	"limit": "5"
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
		method := "POST"
		url := fmt.Sprintf("%s/mcisRecommendVm", epTumblebug)

		// Extract server info from source computing infra info
		cores := server.Compute.ComputeResource.CPU.Cores
		memory := MBtoGiB(float64(server.Compute.ComputeResource.Memory.Size))

		coreUpperLimit := cores + 1
		var coreLowerLimit uint
		if cores > 1 {
			coreLowerLimit = cores - 1
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
			Uint32("memoryUpperLimit", memoryUpperLimit).
			Uint32("memoryLowerLimit", memoryLowerLimit).
			Str("providerName", providerName).
			Str("regionName", regionName).
			Str("osNameWithVersion", osNameWithVersion).
			Msg("Source computing infrastructure info")

		// Set a deployment plan with the server info
		plan := fmt.Sprintf(planDocstring,
			coreUpperLimit,
			coreLowerLimit,
			memoryUpperLimit,
			memoryLowerLimit,
			providerName,
			regionName,
		)

		// Request body
		reqRecommVm := new(mcis.DeploymentPlan)
		err := json.Unmarshal([]byte(plan), reqRecommVm)
		if err != nil {
			log.Err(err).Msg("")
			return cloudmodel.InfraMigrationReq{}, err
		}
		log.Debug().Msgf("deployment plan for the VM recommendation: %+v", reqRecommVm)

		// Response body
		resRecommVm := new(mcir.TbSpecInfo)

		err = common.ExecuteHttpRequest(
			client,
			method,
			url,
			nil,
			common.SetUseBody(*reqRecommVm),
			reqRecommVm,
			resRecommVm,
			common.VeryShortDuration,
		)

		if err != nil {
			log.Err(err).Msg("")
			return cloudmodel.InfraMigrationReq{}, err
		}
		log.Debug().Msgf("resRecommVm: %+v", resRecommVm)

		// Set target VM
		image := fmt.Sprintf("%s+%s+%s", providerName, regionName, osNameWithVersion)
		spec := resRecommVm.CspSpecName

		vm := cloudmodel.HostMigrationReq{
			ConnectionName: "",
			CommonImage:    image,
			CommonSpec:     spec,
			Description:    "a recommended virtual machine",
			Label:          "rehosted virtual machine",
			Name:           "recomm-vm",
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
