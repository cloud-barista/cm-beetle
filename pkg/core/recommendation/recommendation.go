package recommendation

import (
	"strings"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model/source/infra"
)

func Recommend(source infra.Infra) (model.TbMcisDynamicReq, error) {

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

	// Instance with deafult values
	targetVM := model.TbVmDynamicReq{
		ConnectionName: "string",
		Description:    "Description",
		Label:          "DynamicVM",
		Name:           "RecommVM01",
		RootDiskSize:   "default",
		RootDiskType:   "default",
		SubGroupSize:   "3",
		VmUserPassword: "string",
	}

	// Match source and target

	// Do something for matching VM image
	// (example) lower case and remove space
	// (example) select common image from a list of images with the lower cased OS name
	// Sample code
	lowerCaseOSName := strings.ToLower(source.Compute.OS.OS.Name)
	targetVM.CommonImage = lowerCaseOSName

	// Do something
	// Do something for matching instance spec
	// (example) get the number of cores from the source instance spec
	// (example) select the instance spec with the with similar or appropriate number of cores from a list of instance specs

	// Sample code
	// if source.Compute.ComputeResource.CPU.Cores == 1 && source.Compute.ComputeResource.Memory.Size == 2 {
	targetVM.CommonSpec = "aws-ap-northeast-2-t3-small"
	// }

	// Instance with deafult values
	targetInfra := model.TbMcisDynamicReq{
		Description:     "Made in CB-TB",
		InstallMonAgent: "no",
		Label:           "DynamicVM",
		Name:            "",
		SystemLabel:     "",
		Vm:              []model.TbVmDynamicReq{},
	}

	targetInfra.Name = "RecommInfra01"
	targetInfra.Vm = append(targetInfra.Vm, targetVM)

	return targetInfra, nil
}
