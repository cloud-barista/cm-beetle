package recommendation

import (
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
	targetInfra := model.TbMcisDynamicReq{
		Description:     "A cloud infra recommended by CM-Beetle",
		InstallMonAgent: "no",
		Label:           "DynamicVM",
		Name:            "",
		SystemLabel:     "",
		Vm:              []model.TbVmDynamicReq{},
	}

	targetInfra.Vm = append(targetInfra.Vm, targetVM)

	return targetInfra, nil
}
