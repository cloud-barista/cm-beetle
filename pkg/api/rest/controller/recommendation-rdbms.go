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

// Package controller has handlers and their request/response bodies for recommendation APIs
package controller

import (
	"fmt"
	"net/http"
	"strings"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// Dummy blank identifier to retain rdbmsmodel import for Swagger annotations
var _ = rdbmsmodel.RDBMSSupportResponse{}
var _ = rdbmsmodel.RDBMSCapabilityResponse{}

// GetRDBMSSupport godoc
// @ID GetRDBMSSupport
// @Summary Get CSP support map for Managed RDBMS (RDS)
// @Description Retrieve CSP feature support matrix for managed RDBMS / RDS (Amazon RDS, Google Cloud SQL, Azure Database, NCP Cloud DB, etc.) via CB-Tumblebug
// @Tags [Recommendation] Managed RDBMS
// @Accept json
// @Produce json
// @Param providerName query string false "CSP Type filter (e.g., aws, azure, gcp)"
// @Success 200 {object} rdbmsmodel.RDBMSSupportResponse
// @Router /recommendation/middleware/rdbms/support [get]
func GetRDBMSSupport(c echo.Context) error {
	providerName := c.QueryParam("providerName")

	result, err := recommendation.GetRDBMSSupport(providerName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get RDBMS support matrix")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to get RDBMS support matrix: %v", err)))
	}

	return c.JSON(http.StatusOK, result)
}

// GetRDBMSCapability godoc
// @ID GetRDBMSCapability
// @Summary Get real-time capability and spec options for Managed RDBMS (RDS)
// @Description Retrieve real-time engine versions, DB instance specs, and storage options for Managed RDBMS / RDS (Amazon RDS, Cloud SQL, Azure Database, etc.) from CB-Tumblebug
// @Tags [Recommendation] Managed RDBMS
// @Accept json
// @Produce json
// @Param connectionName query string true "Connection Name (e.g., aws-ap-northeast-2)"
// @Success 200 {object} rdbmsmodel.RDBMSCapabilityResponse
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /recommendation/middleware/rdbms/capability [get]
func GetRDBMSCapability(c echo.Context) error {
	connectionName := c.QueryParam("connectionName")
	if connectionName == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("connectionName is required"))
	}

	result, err := recommendation.GetRDBMSCapability(connectionName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get RDBMS capability")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to get RDBMS capability: %v", err)))
	}

	return c.JSON(http.StatusOK, result)
}

// ValidateRDBMS godoc
// @ID ValidateRDBMSRecommendation
// @Summary Validate Managed RDBMS (RDS) configuration against target cloud
// @Description Perform strict dry-run validation for a Managed RDBMS / RDS configuration against target cloud constraints
// @Tags [Recommendation] Managed RDBMS
// @Accept json
// @Produce json
// @Param nsId query string false "Namespace ID (defaults to 'mig01')" default(mig01)
// @Param request body rdbmsmodel.RDBMSCreateRequest true "RDBMS creation request to validate"
// @Success 200 {object} model.ApiResponse[rdbmsmodel.RDBMSCreateRequest] "Successfully validated RDBMS configuration"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Validation error"
// @Router /recommendation/middleware/rdbms/validate [post]
func ValidateRDBMS(c echo.Context) error {
	nsId := c.QueryParam("nsId")
	if strings.TrimSpace(nsId) == "" {
		nsId = common.DefaultNamespaceId
	}

	var req rdbmsmodel.RDBMSCreateRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	req.AutoFillDefaults = false
	validated, err := recommendation.ValidateRDBMS(nsId, req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to validate RDBMS configuration")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(fmt.Sprintf("Validation failed: %v", err)))
	}

	res := model.SuccessResponseWithMessage(validated, "RDBMS configuration is valid")
	return c.JSON(http.StatusOK, res)
}

// ============================================================================
// Request Models
// ============================================================================

// RecommendRDBMSRequest represents a request for managed RDBMS migration recommendations
type RecommendRDBMSRequest struct {
	DesiredCloud         rdbmsmodel.CloudProperty         `json:"desiredCloud" validate:"required"`
	SourceRDBMSInstances []rdbmsmodel.SourceRDBMSProperty `json:"sourceRDBMSInstances" validate:"required,min=1"`
}

// RecommendRDBMS godoc
// @ID RecommendRDBMS
// @Summary Recommend Managed RDBMS (RDS) for cloud migration
// @Description Recommend optimal Managed RDBMS / RDS (MySQL, MariaDB) instance specs, engine versions, and storage configurations for target cloud migration (supports AWS RDS, GCP Cloud SQL, Azure Database, Alibaba ApsaraDB, TencentDB, IBM Databases, NCP Cloud DB, NHN RDS)
// @Description
// @Description [Note] `desiredCsp` and `desiredRegion` are required.
// @Description - `desiredCsp` and `desiredRegion` can be set in the query parameter or the request body.
// @Description - If set in the request body, the query parameter values will be overridden.
// @Description
// @Description [Note] The recommended instance names use default patterns (`mig-rdbms-01`, `mig-rdbms-02`, ...).
// @Description - To apply a naming prefix at migration time, use the `nameSeed` query parameter on the migration API.
// @Tags [Recommendation] Managed RDBMS
// @Accept json
// @Produce json
// @Param request body RecommendRDBMSRequest true "Specify source RDBMS instances to be migrated"
// @Param desiredCsp query string false "CSP (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,tencent,ibm,openstack,ncp,nhn) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[rdbmsmodel.RecommendedRDBMS] "Successfully recommended managed RDBMS"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Router /recommendation/middleware/rdbms [post]
func RecommendRDBMS(c echo.Context) error {
	// [Input]
	var req RecommendRDBMSRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	desiredCsp := c.QueryParam("desiredCsp")
	desiredRegion := c.QueryParam("desiredRegion")

	if desiredCsp == "" {
		desiredCsp = req.DesiredCloud.Csp
	}
	if desiredRegion == "" {
		desiredRegion = req.DesiredCloud.Region
	}

	if desiredCsp == "" || desiredRegion == "" {
		log.Warn().Msg("desiredCsp and desiredRegion are required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP and region required"))
	}

	if len(req.SourceRDBMSInstances) == 0 {
		log.Warn().Msg("At least one source RDBMS instance must be provided")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one source RDBMS instance required"))
	}

	log.Info().
		Str("desiredCsp", desiredCsp).
		Str("region", desiredRegion).
		Int("sourceInstances", len(req.SourceRDBMSInstances)).
		Msg("Processing managed RDBMS recommendation request")

	// [Process]
	recommended, err := recommendation.RecommendRDBMS(desiredCsp, desiredRegion, req.SourceRDBMSInstances)
	if err != nil {
		log.Error().Err(err).Msg("Failed to recommend managed RDBMS")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	successMsg := fmt.Sprintf("Successfully recommended %d managed RDBMS configuration(s) for %s (%s)",
		len(recommended.TargetRDBMSInstances), desiredCsp, desiredRegion)
	res := model.SuccessResponseWithMessage(recommended, successMsg)

	return c.JSON(http.StatusOK, res)
}
