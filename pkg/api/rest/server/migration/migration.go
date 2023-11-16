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

////////////////////////

type Network struct {
	Name          string `json:"name"`
	Id            string `json:"id"`
	IPv4CIDRBlock string `json:"ipv4CidrBlock"`
	IPv6CIDRBlock string `json:"ipv6CidrBlock"`
}

type Subnet struct {
	Network
	ParentNetworkId string `json:"parentNetworkId"`
}

type DummyNetwork struct {
	Network
	Subnets []Subnet `json:"subnets"`
}

type MigrateNetworkRequest struct {
	DummyNetwork
}

type MigrateNetworkResponse struct {
	DummyNetwork
}

// MigrateNetwork godoc
// @Summary (Skeleton) Migrate network on a cloud platform
// @Description It migrates network on a cloud platform. Network includes name, ID, IPv4 CIDR block, IPv6 CIDR block, and so on.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param NetworkInfo body MigrateNetworkRequest true "Specify name, IPv4 CIDR block, etc."
// @Success 200 {object} MigrateNetworkResponse "Successfully migrated network on a cloud platform"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /migration/infra/network [post]
func (rh *Handlers) MigrateNetwork(c echo.Context) error {

	// [Note] Input section
	req := &MigrateNetworkRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	fmt.Printf("RequestBody: %v\n", req)
	fmt.Print(req)
	fmt.Print(req.DummyNetwork)

	// [Note] Process section
	// Something to process here like,
	// Perform some functions,
	// Calls external APIs and so on

	res := &MigrateNetworkResponse{}
	fmt.Print(res)
	fmt.Print(res.DummyNetwork)

	// This is an intentionally created variable.
	// You will have to delete this later.
	var err error = nil

	// [Note] Ouput section
	if err != nil {
		common.CBLog.Error(err)
		mapA := map[string]string{"message": err.Error()}
		return c.JSON(http.StatusInternalServerError, &mapA)
	}

	return c.JSON(http.StatusOK, res)

}

////////////////////////

////////////////////////

type Storage struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Type string `json:"type"`
	Size string `json:"size"`
}

type DummyStorage struct {
	Storage
	NetworkID string `json:"NetworkId"`
}

type MigrateStorageRequest struct {
	DummyStorage
}

type MigrateStorageResponse struct {
	DummyStorage
}

// MigrateStorage godoc
// @Summary (Skeleton) Migrate storage on a cloud platform
// @Description It migrates storage on a cloud platform. Storage includes name, ID, type, size, and so on.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param StorageInfo body MigrateStorageRequest true "Specify name, type, size, affiliated Network ID, and so on."
// @Success 200 {object} MigrateStorageResponse "Successfully migrated storage on a cloud platform"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /migration/infra/storage [post]
func (rh *Handlers) MigrateStorage(c echo.Context) error {

	// [Note] Input section
	req := &MigrateStorageRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	fmt.Printf("RequestBody: %v\n", req)
	fmt.Print(req)
	fmt.Print(req.DummyStorage)

	// [Note] Process section
	// Something to process here like,
	// Perform some functions,
	// Calls external APIs and so on

	res := &MigrateStorageResponse{}
	fmt.Print(res)
	fmt.Print(res.DummyStorage)

	// This is an intentionally created variable.
	// You will have to delete this later.
	var err error = nil

	// [Note] Ouput section
	if err != nil {
		common.CBLog.Error(err)
		mapA := map[string]string{"message": err.Error()}
		return c.JSON(http.StatusInternalServerError, &mapA)
	}

	return c.JSON(http.StatusOK, res)

}

////////////////////////

////////////////////////

type Instance struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Spec string `json:"type"`
	OS   string `json:"os"`
}

type DummyInstance struct {
	Instance
	NetworkID string `json:"NetworkId"`
}

type MigrateInstanceRequest struct {
	DummyInstance
}

type MigrateInstanceResponse struct {
	DummyInstance
}

// MigrateInstance godoc
// @Summary (Skeleton) Migrate instance on a cloud platform
// @Description It migrates instance on a cloud platform. Storage includes name, spec, OS, and so on.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param InstanceInfo body MigrateInstanceRequest true "Specify name, spec, OS, and so on."
// @Success 200 {object} MigrateInstanceResponse "Successfully migrated storage on a cloud platform"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /migration/infra/instance [post]
func (rh *Handlers) MigrateInstance(c echo.Context) error {

	// [Note] Input section
	req := &MigrateInstanceRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	fmt.Printf("RequestBody: %v\n", req)
	fmt.Print(req)
	fmt.Print(req.DummyInstance)

	// [Note] Process section
	// Something to process here like,
	// Perform some functions,
	// Calls external APIs and so on

	res := &MigrateInstanceResponse{}
	fmt.Print(res)
	fmt.Print(res.DummyInstance)

	// This is an intentionally created variable.
	// You will have to delete this later.
	var err error = nil

	// [Note] Ouput section
	if err != nil {
		common.CBLog.Error(err)
		mapA := map[string]string{"message": err.Error()}
		return c.JSON(http.StatusInternalServerError, &mapA)
	}

	return c.JSON(http.StatusOK, res)

}

////////////////////////
