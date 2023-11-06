/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package migration is to handle REST API for migration
package migration

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handlers struct {
}

// type Infrastructure struct {
// 	Network        string
// 	Disk           string
// 	Compute        string
// 	SecurityGroup  string
// 	VirtualMachine string
// }

// TbMcisDynamicReq is sturct for requirements to create MCIS dynamically (with default resource option)
type TbMcisDynamicReq struct {
	Name string `json:"name" validate:"required" example:"mcis01"`

	// InstallMonAgent Option for CB-Dragonfly agent installation ([yes/no] default:yes)
	InstallMonAgent string `json:"installMonAgent" example:"no" default:"yes" enums:"yes,no"` // yes or no

	// Label is for describing the mcis in a keyword (any string can be used)
	Label string `json:"label" example:"DynamicVM" default:""`

	// SystemLabel is for describing the mcis in a keyword (any string can be used) for special System purpose
	SystemLabel string `json:"systemLabel" example:"" default:""`

	Description string `json:"description" example:"Made in CB-TB"`

	Vm []TbVmDynamicReq `json:"vm" validate:"required"`
}

// TbVmDynamicReq is struct to get requirements to create a new server instance dynamically (with default resource option)
type TbVmDynamicReq struct {
	// VM name or subGroup name if is (not empty) && (> 0). If it is a group, actual VM name will be generated with -N postfix.
	Name string `json:"name" example:"g1-1"`

	// if subGroupSize is (not empty) && (> 0), subGroup will be gernetad. VMs will be created accordingly.
	SubGroupSize string `json:"subGroupSize" example:"3" default:""`

	Label string `json:"label" example:"DynamicVM"`

	Description string `json:"description" example:"Description"`

	// CommonSpec is field for id of a spec in common namespace
	CommonSpec string `json:"commonSpec" validate:"required" example:"aws-ap-northeast-2-t2-small"`
	// CommonImage is field for id of a image in common namespace
	CommonImage string `json:"commonImage" validate:"required" example:"ubuntu18.04"`

	RootDiskType string `json:"rootDiskType,omitempty" example:"default, TYPE1, ..."`  // "", "default", "TYPE1", AWS: ["standard", "gp2", "gp3"], Azure: ["PremiumSSD", "StandardSSD", "StandardHDD"], GCP: ["pd-standard", "pd-balanced", "pd-ssd", "pd-extreme"], ALIBABA: ["cloud_efficiency", "cloud", "cloud_essd"], TENCENT: ["CLOUD_PREMIUM", "CLOUD_SSD"]
	RootDiskSize string `json:"rootDiskSize,omitempty" example:"default, 30, 42, ..."` // "default", Integer (GB): ["50", ..., "1000"]

	VmUserPassword string `json:"vmUserPassword default:""`
	// if ConnectionName is given, the VM tries to use associtated credential.
	// if not, it will use predefined ConnectionName in Spec objects
	ConnectionName string `json:"connectionName,omitempty" default:""`
}

type MigrateInfraRequest struct {
	// [NOTE] Failed to embed the struct in CB-Tumblebug as follows:
	// mcis.TbMcisDynamicReq

	TbMcisDynamicReq
}

type MigrateInfraResponse struct {
	ResponseText string
}

// MigrateInfra godoc
// @Summary Migrate an infrastructure on a cloud platform
// @Description It migrates an infrastructure on a cloud platform. Infrastructure includes network, storage, compute, and so on.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param InfrastructureInfo body MigrateInfraRequest true "Specify network, disk, compute, security group, virtual machine, etc."
// @Success 200 {object} MigrateInfraResponse "Successfully migrated infrastructure on a cloud platform"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /migration/infra [post]
func (rh *Handlers) MigrateInfra(c echo.Context) error {

	// Input
	req := &MigrateInfraRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	fmt.Print(req)

	res := &MigrateInfraResponse{}
	// Process

	// Call CB-Tumblebug API, which can be "/mcisDynamic"

	// Ouput

	// if err != nil {
	// 	common.CBLog.Error(err)
	// 	mapA := map[string]string{"message": err.Error()}
	// 	return c.JSON(http.StatusInternalServerError, &mapA)
	// }

	return c.JSON(http.StatusOK, res)

}
