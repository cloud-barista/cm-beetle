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
	"strings"

	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog/log"
)

type MigrateInfraWithDefaultsRequest struct {
	// [NOTE] Failed to embed the struct in CB-Tumblebug as follows:
	// mci.MciDynamicReq

	cloudmodel.MciDynamicReq
}

type MigrateInfraWithDefaultsResponse struct {
	cloudmodel.VmInfraInfo
}

// MigrateInfraWithDefaults godoc
// @ID MigrateInfraWithDefaults
// @Summary Migrate an infrastructure to the multi-cloud infrastructure (MCI) with defaults
// @Description Migrate an infrastructure to the multi-cloud infrastructure (MCI) with defaults.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciInfo body MigrateInfraWithDefaultsRequest true "Specify the information for the targeted mulci-cloud infrastructure (MCI)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 201 {object} model.ApiResponse[MigrateInfraWithDefaultsResponse] "Successfully migrated to the multi-cloud infrastructure"
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/mciWithDefaults [post]
func MigrateInfraWithDefaults(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	req := new(MigrateInfraWithDefaultsRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	log.Debug().Msgf("req: %v", req)
	log.Debug().Msgf("req.MciDynamicReq: %v", req.MciDynamicReq)

	// [Process]
	// Create the VM infrastructure for migration
	mciInfo, err := migration.CreateVMInfraWithDefaults(nsId, &req.MciDynamicReq)

	log.Debug().Msgf("mciInfo: %v", mciInfo)

	// [Output]
	if err != nil {
		log.Error().Err(err).Msg("failed to create VM infrastructure")

		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SuccessResponse(mciInfo))
}

// TODO: Check and dev the request and response bodies for the following API

type MigrateInfraRequest struct {
	cloudmodel.RecommendedVmInfra
}

type MigrateInfraResponse struct {
	cloudmodel.VmInfraInfo
}

// MigrateInfra godoc
// @ID MigrateInfra
// @Summary Migrate an infrastructure to the multi-cloud infrastructure (MCI) with defaults
// @Description Migrate an infrastructure to the multi-cloud infrastructure (MCI) with defaults.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciInfo body MigrateInfraRequest true "Specify the information for the targeted mulci-cloud infrastructure (MCI)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 201 {object} model.ApiResponse[MigrateInfraResponse] "Successfully migrated to the multi-cloud infrastructure"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/mci [post]
func MigrateInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	req := new(MigrateInfraRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	// log.Debug().Msgf("req: %+v", req)
	log.Debug().Msgf("req.RecommendedVmInfra: %+v", req.RecommendedVmInfra)

	// [Process]
	// Apply NameSeed (Late Binding) before migration
	infraToMigrate := common.ApplyNameSeed(req.RecommendedVmInfra)

	// Validate names and referential integrity
	if ok, detail := common.ValidateComposedNames(infraToMigrate); !ok {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Naming/Reference validation failed: "+detail))
	}

	// Create the VM infrastructure for migration
	mciInfo, err := migration.CreateVMInfra(nsId, &infraToMigrate)

	log.Debug().Msgf("mciInfo: %+v", mciInfo)

	// [Output]
	if err != nil {
		log.Error().Err(err).Msg("failed to create VM infrastructure")

		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SuccessResponse(mciInfo))
}

// ListInfra godoc
// @ID ListInfra
// @Summary Get the migrated multi-cloud infrastructure (MCI)
// @Description Get the migrated multi-cloud infrastructure (MCI)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param option query string false "Option for getting the migrated multi-cloud infrastructure" Enums(id)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[cloudmodel.MciInfoList] "The info list of the migrated multi-cloud infrastructure (MCI)"
// @Success 200 {object} model.ApiResponse[cloudmodel.IdList] "The ID list of The migrated multi-cloud infrastructure (MCI)"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/mci [get]
func ListInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the nanespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	option := c.QueryParam("option")
	if option != "" && option != "id" {
		err := fmt.Errorf("invalid request, the option (option: %s) is invalid", option)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// [Process] List the migrated multi-cloud infrastructures as the option
	switch option {
	case "id":
		idList, err := migration.ListVMInfraIDs(nsId, option)
		if err != nil {
			log.Error().Err(err).Msg("failed to get the migrated multi-cloud infrastructure IDs")
			return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, model.SuccessResponse(idList))
	default:
		infraInfoList, err := migration.ListAllVMInfraInfo(nsId)
		if err != nil {
			log.Error().Err(err).Msg("failed to get the migrated multi-cloud infrastructures")
			return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, model.SuccessResponse(infraInfoList))
	}
	// return c.JSON(http.StatusInternalServerError, nil)
}

// GetInfra godoc
// @ID GetInfra
// @Summary Get the migrated multi-cloud infrastructure (MCI)
// @Description Get the migrated multi-cloud infrastructure (MCI)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciId path string true "Migrated Multi-Cloud Infrastructure (MCI) ID" default(mci101)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[MigrateInfraResponse] "The migrated multi-cloud infrastructure (MCI) information"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/mci/{mciId} [get]
func GetInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the nanespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	mciId := c.Param("mciId")
	if mciId == "" {
		err := fmt.Errorf("invalid request, the multi-cloud infrastructure ID (mciId: %s) is required", mciId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// [Process]
	vmInfraInfo, err := migration.GetVMInfra(nsId, mciId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get the migrated multi-cloud infrastructure")

		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "does not exist") {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse("Infrastructure not found"))
		}

		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	// [Ouput]
	return c.JSON(http.StatusOK, model.SuccessResponse(vmInfraInfo))
}

// DeleteInfra godoc
// @ID DeleteInfra
// @Summary Delete the migrated mult-cloud infrastructure (MCI)
// @Description Delete the migrated mult-cloud infrastructure (MCI)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciId path string true "Migrated Multi-Cloud Infrastructure (MCI) ID" default(mci101)
// @Param option query string false "Option for deletion" Enums(terminate,force) default(terminate)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[any] "The result of deleting the migrated multi-cloud infrastructure (MCI)"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/mci/{mciId} [delete]
func DeleteInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	mciId := c.Param("mciId")
	if mciId == "" {
		err := fmt.Errorf("invalid request, the multi-cloud infrastructure ID (mciId: %s) is required", mciId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	option := c.QueryParam("option")
	if option != "" && option != "terminate" && option != "force" {
		err := fmt.Errorf("invalid request, the option (option: %s) is invalid", option)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// [Process]
	retMsg, err := migration.DeleteVMInfra(nsId, mciId, option)

	if err != nil {
		log.Error().Err(err).Msg("failed to delete the migrated multi-cloud infrastructure")

		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "does not exist") {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse("Infrastructure not found"))
		}

		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	// [Ouput]
	return c.JSON(http.StatusOK, model.SimpleSuccessResponse(retMsg.Message))
}
