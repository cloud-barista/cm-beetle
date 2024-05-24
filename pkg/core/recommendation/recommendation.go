package recommendation

import (
	"fmt"
	"os"

	"github.com/cloud-barista/cb-tumblebug/src/core/mcir"
	"github.com/cloud-barista/cb-tumblebug/src/core/mcis"
	cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

func Recommend(source infra.Infra) (cloudmodel.TbMcisDynamicReq, error) {

	// {
	// 	"description": "Made in CB-TB",
	// 	"installMonAgent": "no",
	// 	"label": "DynamicVM",
	// 	"name": "mcis01",
	// 	"systemLabel": "",
	// 	"vm": [
	// 		{
	// 		"commonImage": "ubuntu18.04",
	// 		"commonSpec": "aws-ap-northeast-2-t2-small",
	// 		"connectionName": "string",
	// 		"description": "Description",
	// 		"label": "DynamicVM",
	// 		"name": "g1-1",
	// 		"rootDiskSize": "default, 30, 42, ...",
	// 		"rootDiskType": "default, TYPE1, ...",
	// 		"subGroupSize": "3",
	// 		"vmUserPassword default:": "string"
	// 		}
	// 	]
	// }

	// Extract server info from source computing infra info

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
		return cloudmodel.TbMcisDynamicReq{}, err
	}
	log.Debug().Msgf("resReadyz: %+v", resReadyz.Message)

	// Recommand VMs

	method = "POST"
	url = fmt.Sprintf("%s/mcisRecommendVm", epTumblebug)

	// Set the deployment plan
	reqRecommVm := new(mcis.DeploymentPlan)
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
		return cloudmodel.TbMcisDynamicReq{}, err
	}
	log.Debug().Msgf("resRearesRecommVmdyz: %+v", resRecommVm)

	// Instance with deafult values
	targetVM := cloudmodel.TbVmDynamicReq{
		ConnectionName: "",
		Description:    "Description",
		Label:          "DynamicVM",
		Name:           "recomm-vm",
		RootDiskSize:   "default",
		RootDiskType:   "default",
		SubGroupSize:   "1",
		VmUserPassword: "",
	}

	// Match source and target

	// Do something for matching VM image
	// (example) lower case and remove space
	// (example) select common image from a list of images with the lower cased OS name
	// Sample code
	// lowerCaseOSName := strings.ToLower(source.Compute.OS.OS.Name)
	targetVM.CommonImage = "ubuntu22.04"

	// Do something
	// Do something for matching instance spec
	// (example) get the number of cores from the source instance spec
	// (example) select the instance spec with the with similar or appropriate number of cores from a list of instance specs

	// Sample code
	// if source.Compute.ComputeResource.CPU.Cores == 1 && source.Compute.ComputeResource.Memory.Size == 2 {
	targetVM.CommonSpec = "aws-ap-northeast-2-t3-small"
	// }

	// Instance with deafult values
	targetInfra := cloudmodel.TbMcisDynamicReq{
		Description:     "A cloud infra recommended by CM-Beetle",
		InstallMonAgent: "no",
		Label:           "DynamicVM",
		Name:            "",
		SystemLabel:     "",
		Vm:              []cloudmodel.TbVmDynamicReq{},
	}

	targetInfra.Vm = append(targetInfra.Vm, targetVM)

	return targetInfra, nil
}
