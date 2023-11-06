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
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
}

type MigrateInfraRequest struct {
	// [NOTE] Failed to embed the struct in CB-Tumblebug as follows:
	// mcis.TbMcisDynamicReq

	TbMcisDynamicReq
}

type MigrateInfraResponse struct {
	TbMcisInfo
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

	// [Note] Input section
	req := &MigrateInfraRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	fmt.Printf("RequestBody: %v\n", req)
	fmt.Print(req)
	fmt.Print(req.TbMcisDynamicReq)

	// [Note] Process section
	// Call CB-Tumblebug API, which can be "/mcisDynamic"
	// Default nsId is "ns01"
	nsId := "ns01"
	result, err := createVMInfra(nsId, &req.TbMcisDynamicReq)

	fmt.Print(result)

	// [Note] Ouput section
	if err != nil {
		common.CBLog.Error(err)
		mapA := map[string]string{"message": err.Error()}
		return c.JSON(http.StatusInternalServerError, &mapA)
	}

	res := result

	return c.JSON(http.StatusOK, res)

}

func createVMInfra(nsId string, infraModel *TbMcisDynamicReq) (TbMcisInfo, error) {

	client := resty.New()
	client.SetBasicAuth("default", "default")
	method := "POST"

	// CB-Tumblebug API endpoint
	cbTumblebugApiEndpoint := "http://localhost:1323/tumblebug"
	url := cbTumblebugApiEndpoint + fmt.Sprintf("/ns/%s/mcisDynamic", nsId)
	// url := fmt.Sprintf("%s/ns/{nsId}/mcisDynamic%s", cbTumblebugApiEndpoint, idDetails.IdInSp)

	// Set request body
	requestBody := *infraModel

	// Set response body
	responseBody := TbMcisInfo{}

	client.SetTimeout(5 * time.Minute)

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		nil,
		common.SetUseBody(requestBody),
		&requestBody,
		&responseBody,
		common.MediumDuration,
	)

	if err != nil {
		// common.CBLog.Error(err)
		return TbMcisInfo{}, err
	}

	return responseBody, nil
}
