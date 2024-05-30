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

// Package controller has handlers and their request/response bodies for migration APIs
package controller

import (
	"fmt"
	"net/http"

	model "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/beetle"
	cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/labstack/echo/v4"

	// Black import (_) is for running a package's init() function without using its other contents.
	_ "github.com/cloud-barista/cm-beetle/pkg/logger"
	"github.com/rs/zerolog/log"
)

type MigrateInfraRequest struct {
	// [NOTE] Failed to embed the struct in CB-Tumblebug as follows:
	// mcis.TbMcisDynamicReq

	cloudmodel.InfraMigrationReq
}

type MigrateInfraResponse struct {
	cloudmodel.TbMcisInfo
}

// MigrateInfra godoc
// @Summary Migrate an infrastructure on a cloud platform
// @Description It migrates an infrastructure on a cloud platform.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param InfrastructureInfo body MigrateInfraRequest true "Specify network, disk, compute, security group, virtual machine, etc."
// @Success 200 {object} MigrateInfraResponse "Successfully migrated infrastructure on a cloud platform"
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /migration/infra [post]
func MigrateInfra(c echo.Context) error {

	// [Note] Input section
	// nsId := c.Param("nsId")
	// if nsId == "" {
	// 	err := fmt.Errorf("invalid request, namespace ID (nsId: %s) is required", nsId)
	// 	log.Warn().Msg(err.Error())
	// 	res := model.Response{
	// 		Success: false,
	// 		Text:    err.Error(),
	// 	}
	// 	return c.JSON(http.StatusBadRequest, res)
	// }
	nsId := common.DefaulNamespaceId

	req := &MigrateInfraRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}

	log.Trace().Msgf("req: %v\n", req)
	log.Trace().Msgf("req.TbMcisDynamicReq: %v\n", req.InfraMigrationReq)

	nsInfo, err := common.GetNamespace(nsId)
	if err != nil {
		// [temporary code block] Create a namespace as a default
		log.Warn().Msgf("failed to get the namespace (nsId: %s)", nsId)
		log.Info().Msgf("create a namespace as a default (nsId: %s)", common.DefaulNamespaceId)
		nsReq := common.NsReq{
			Name: common.DefaulNamespaceId,
		}
		nsInfo, err = common.CreateNamespace(nsReq)

	}

	if nsInfo.Id == "" {
		err := fmt.Errorf("not found the namespace (nsId: %s), create a namespace first", nsId)
		log.Error().Err(err).Msg("")
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	// Create the VM infrastructure for migration
	infraInfo, err := migration.CreateVMInfra(nsId, &req.InfraMigrationReq)

	log.Trace().Msgf("infraInfo: %v\n", infraInfo)

	// [Note] Ouput section
	if err != nil {
		log.Error().Err(err).Msg("failed to create VM infrastructure")

		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := infraInfo
	return c.JSON(http.StatusOK, res)

}

// GetInfra godoc
// @Summary Get the migrated infrastructure on a cloud platform
// @Description It gets the migrated infrastructure on a cloud platform.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param infraId path string true "a infrastructure ID created for migration"
// @Success 200 {object} MigrateInfraResponse "Successfully got the migrated infrastructure on a cloud platform"
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /migration/infra/{infraId} [get]
func GetInfra(c echo.Context) error {

	// [Note] Input section
	// nsId := c.Param("nsId")
	// if nsId == "" {
	// 	err := fmt.Errorf("invalid request, the nanespace ID (nsId: %s) is required", nsId)
	// 	log.Warn().Msg(err.Error())
	// 	res := model.Response{
	// 		Success: false,
	// 		Text:    err.Error(),
	// 	}
	// 	return c.JSON(http.StatusBadRequest, res)
	// }
	nsId := common.DefaulNamespaceId

	infraId := c.Param("infraId")
	if infraId == "" {
		err := fmt.Errorf("invalid request, the infrastructure ID (infraId: %s) is required", infraId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Note] Process section
	vmInfraInfo, err := migration.GetVMInfra(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get the migrated infrastructure")
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// [Note] Ouput section
	return c.JSON(http.StatusOK, vmInfraInfo)
}

// DeleteInfra godoc
// @Summary Delete the migrated infrastructure on a cloud platform
// @Description It deletes the migrated infrastructure on a cloud platform.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param infraId path string true "a infrastructure ID created for migration"
// @Success 200 {object} model.Response "Successfully deleted the migrated infrastructure on a cloud platform"
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /migration/infra/{infraId} [delete]
func DeleteInfra(c echo.Context) error {

	// [Note] Input section
	// nsId := c.Param("nsId")
	// if nsId == "" {
	// 	err := fmt.Errorf("invalid request, the nanespace ID (nsId: %s) is required", nsId)
	// 	log.Warn().Msg(err.Error())
	// 	res := model.Response{
	// 		Success: false,
	// 		Text:    err.Error(),
	// 	}
	// 	return c.JSON(http.StatusBadRequest, res)
	// }
	nsId := common.DefaulNamespaceId

	infraId := c.Param("infraId")
	if infraId == "" {
		err := fmt.Errorf("invalid request, the infrastructure ID (infraId: %s) is required", infraId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Note] Process section
	retMsg, err := migration.DeleteVMInfra(nsId, infraId)

	if err != nil {
		log.Error().Err(err).Msg("failed to delete the migrated infrastructure")
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// [Note] Ouput section
	res := model.Response{
		Success: true,
		Text:    retMsg.Message,
	}

	return c.JSON(http.StatusOK, res)
}

////////////////////////

// type MigrateNetworkRequest struct {
// 	cloudmodel.DummyNetwork
// }

// type MigrateNetworkResponse struct {
// 	cloudmodel.DummyNetwork
// }

// // MigrateNetwork godoc
// // @Summary (Skeleton) Migrate network on a cloud platform
// // @Description It migrates network on a cloud platform. Network includes name, ID, IPv4 CIDR block, IPv6 CIDR block, and so on.
// // @Tags [Migration] Infrastructure
// // @Accept  json
// // @Produce  json
// // @Param NetworkInfo body MigrateNetworkRequest true "Specify name, IPv4 CIDR block, etc."
// // @Success 200 {object} MigrateNetworkResponse "Successfully migrated network on a cloud platform"
// // @Failure 404 {object} common.SimpleMsg
// // @Failure 500 {object} common.SimpleMsg
// // @Router /migration/infra/network [post]
// func MigrateNetwork(c echo.Context) error {

// 	// [Note] Input section
// 	req := &MigrateNetworkRequest{}
// 	if err := c.Bind(req); err != nil {
// 		return err
// 	}

// 	log.Trace().Msgf("req: %v\n", req)
// 	log.Trace().Msgf("req.DummyNetwork: %v\n", req.DummyNetwork)

// 	// [Note] Process section
// 	// Something to process here like,
// 	// Perform some functions,
// 	// Calls external APIs and so on

// 	res := &MigrateNetworkResponse{}
// 	log.Trace().Msgf("res: %v\n", res)
// 	log.Trace().Msgf("res.DummyNetwork: %v\n", res.DummyNetwork)

// 	// This is an intentionally created variable.
// 	// You will have to delete this later.
// 	var err error = nil

// 	// [Note] Ouput section
// 	if err != nil {
// 		log.Error().Err(err).Msg("Failed to migrate network on a cloud platform")
// 		mapA := map[string]string{"message": err.Error()}
// 		return c.JSON(http.StatusInternalServerError, &mapA)
// 	}

// 	return c.JSON(http.StatusOK, res)

// }

// ////////////////////////

// ////////////////////////

// type MigrateStorageRequest struct {
// 	cloudmodel.DummyStorage
// }

// type MigrateStorageResponse struct {
// 	cloudmodel.DummyStorage
// }

// // MigrateStorage godoc
// // @Summary (Skeleton) Migrate storage on a cloud platform
// // @Description It migrates storage on a cloud platform. Storage includes name, ID, type, size, and so on.
// // @Tags [Migration] Infrastructure
// // @Accept  json
// // @Produce  json
// // @Param StorageInfo body MigrateStorageRequest true "Specify name, type, size, affiliated Network ID, and so on."
// // @Success 200 {object} MigrateStorageResponse "Successfully migrated storage on a cloud platform"
// // @Failure 404 {object} common.SimpleMsg
// // @Failure 500 {object} common.SimpleMsg
// // @Router /migration/infra/storage [post]
// func MigrateStorage(c echo.Context) error {

// 	// [Note] Input section
// 	req := &MigrateStorageRequest{}
// 	if err := c.Bind(req); err != nil {
// 		return err
// 	}

// 	log.Trace().Msgf("req: %v\n", req)
// 	log.Trace().Msgf("req.DummyStorage: %v\n", req.DummyStorage)

// 	// [Note] Process section
// 	// Something to process here like,
// 	// Perform some functions,
// 	// Calls external APIs and so on

// 	res := &MigrateStorageResponse{}
// 	log.Trace().Msgf("res: %v\n", res)
// 	log.Trace().Msgf("res.DummyStorage: %v\n", res.DummyStorage)

// 	// This is an intentionally created variable.
// 	// You will have to delete this later.
// 	var err error = nil

// 	// [Note] Ouput section
// 	if err != nil {
// 		log.Error().Err(err).Msg("Failed to migrate storage on a cloud platform")
// 		mapA := map[string]string{"message": err.Error()}
// 		return c.JSON(http.StatusInternalServerError, &mapA)
// 	}

// 	return c.JSON(http.StatusOK, res)

// }

// ////////////////////////

// ////////////////////////

// type MigrateInstanceRequest struct {
// 	cloudmodel.DummyInstance
// }

// type MigrateInstanceResponse struct {
// 	cloudmodel.DummyInstance
// }

// // MigrateInstance godoc
// // @Summary (Skeleton) Migrate instance on a cloud platform
// // @Description It migrates instance on a cloud platform. Storage includes name, spec, OS, and so on.
// // @Tags [Migration] Infrastructure
// // @Accept  json
// // @Produce  json
// // @Param InstanceInfo body MigrateInstanceRequest true "Specify name, spec, OS, and so on."
// // @Success 200 {object} MigrateInstanceResponse "Successfully migrated storage on a cloud platform"
// // @Failure 404 {object} common.SimpleMsg
// // @Failure 500 {object} common.SimpleMsg
// // @Router /migration/infra/instance [post]
// func MigrateInstance(c echo.Context) error {

// 	// [Note] Input section
// 	req := &MigrateInstanceRequest{}
// 	if err := c.Bind(req); err != nil {
// 		return err
// 	}

// 	log.Trace().Msgf("req: %v\n", req)
// 	log.Trace().Msgf("req.DummyInstance: %v\n", req.DummyInstance)

// 	// [Note] Process section
// 	// Something to process here like,
// 	// Perform some functions,
// 	// Calls external APIs and so on

// 	res := &MigrateInstanceResponse{}
// 	log.Trace().Msgf("res: %v\n", res)
// 	log.Trace().Msgf("res.DummyInstance: %v\n", res.DummyInstance)

// 	// This is an intentionally created variable.
// 	// You will have to delete this later.
// 	var err error = nil

// 	// [Note] Ouput section
// 	if err != nil {
// 		log.Error().Err(err).Msg("Failed to migrate instance on a cloud platform")
// 		mapA := map[string]string{"message": err.Error()}
// 		return c.JSON(http.StatusInternalServerError, &mapA)
// 	}

// 	return c.JSON(http.StatusOK, res)

// }

// ////////////////////////
