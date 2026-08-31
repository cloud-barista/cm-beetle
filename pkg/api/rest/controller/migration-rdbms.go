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

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/cloud-barista/cm-beetle/pkg/ratelimit"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// Request Models
// ============================================================================

// MigrateRDBMSRequest represents a request for managed RDBMS migration
type MigrateRDBMSRequest struct {
	rdbmsmodel.RecommendedRDBMS
}

// ============================================================================
// RDBMS Migration API
// ============================================================================

// MigrateRDBMS godoc
// @ID MigrateRDBMS
// @Summary Migrate Managed RDBMS (RDS) instances to cloud
// @Description Provision and migrate managed RDBMS / RDS instances in target cloud based on recommendation results (supports AWS RDS, GCP Cloud SQL, Azure Database, NCP Cloud DB, NHN RDS, Alibaba ApsaraDB, TencentDB, IBM Databases)
// @Description
// @Description [Note]
// @Description - This API provisions managed RDBMS instances in the target cloud within the specified namespace.
// @Description - Input should be the output from the RecommendRDBMS API.
// @Description - Connection name is automatically resolved from CSP and region in the request body.
// @Description
// @Description [Note] `nameSeed` enables dynamic naming via **Late Binding**.
// @Description - If `nameSeed` query param is set (e.g., `?nameSeed=my`), instance names are prefixed: `my-rdbms-01`.
// @Description
// @Description By default this API runs synchronously. Send header `Prefer: respond-async` to run it
// @Description asynchronously instead (recommended due to CSP RDS provisioning time of 5-10 minutes): receive 202 Accepted with a reqId.
// @Tags [Migration] Managed RDBMS
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param nameSeed query string false "Optional prefix for instance names (e.g., 'my' → 'my-rdbms-01')"
// @Param request body MigrateRDBMSRequest true "RDBMS migration request (use RecommendRDBMS response)"
// @Param X-Request-Id header string false "Unique request ID"
// @Param Prefer header string false "Set to 'respond-async' to run this migration asynchronously" Enums(respond-async)
// @Success 201 "Created - Managed RDBMS instances created successfully"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Migration started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during RDBMS creation"
// @Failure 503 {object} model.ApiResponse[any] "Too many requests"
// @Router /migration/middleware/ns/{nsId}/rdbms [post]
func MigrateRDBMS(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	var req MigrateRDBMSRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if req.TargetCloud.Csp == "" || req.TargetCloud.Region == "" {
		log.Warn().Msg("CSP and region are required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP and region required"))
	}

	if len(req.TargetRDBMSInstances) == 0 {
		log.Warn().Msg("At least one target RDBMS instance must be provided")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one target RDBMS instance required"))
	}

	nameSeed := c.QueryParam("nameSeed")
	if ok, detail := common.IsValidNameSeed(nameSeed); !ok {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid nameSeed: "+detail))
	}

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (map[string]any, error) {
			if err := migration.CreateRDBMS(nsId, req.RecommendedRDBMS, nameSeed); err != nil {
				return nil, err
			}
			return map[string]any{"message": "Managed RDBMS instances created successfully"}, nil
		})
		if !started {
			c.Response().Header().Set("Retry-After", "5")
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many async jobs in progress; retry shortly, or retry without Prefer: respond-async"))
		}
		c.Response().Header().Set("Preference-Applied", "respond-async")
		return c.JSON(http.StatusAccepted, model.SuccessResponseWithMessage(
			model.AsyncJobResponse{
				ReqID:     reqID,
				Status:    common.RequestStatusHandling,
				StatusURL: fmt.Sprintf("/beetle/request/%s", reqID),
			},
			"Managed RDBMS migration started. Use GET /request/{reqId} to check status."))
	}

	if err := migration.CreateRDBMS(nsId, req.RecommendedRDBMS, nameSeed); err != nil {
		if retryAfter, ok := ratelimit.RetryAfter(err); ok {
			c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many requests to the underlying infrastructure provider; retry after the given time"))
		}
		log.Error().Err(err).Msg("Managed RDBMS migration failed")
		if strings.Contains(err.Error(), "invalid cloud configuration") {
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SimpleSuccessResponse("Managed RDBMS instances created successfully"))
}

// ValidateMigrateRDBMS godoc
// @ID ValidateMigrateRDBMS
// @Summary Validate Managed RDBMS (RDS) creation request against target cloud
// @Description Validate managed RDBMS configuration against target cloud constraints before actual migration
// @Tags [Migration] Managed RDBMS
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param request body rdbmsmodel.RDBMSCreateRequest true "RDBMS creation request to validate"
// @Success 200 {object} model.ApiResponse[rdbmsmodel.RDBMSCreateRequest] "Successfully validated RDBMS creation configuration"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/rdbms/validate [post]
func ValidateMigrateRDBMS(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
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

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(validated, "Successfully validated RDBMS configuration"))
}

// ListRDBMS godoc
// @ID ListRDBMS
// @Summary List migrated Managed RDBMS (RDS) instances
// @Description Retrieve the list of all migrated managed RDBMS / RDS instances in the namespace
// @Tags [Migration] Managed RDBMS
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Success 200 {object} model.ApiResponse[rdbmsmodel.RDBMSListResponse] "Successfully retrieved RDBMS list"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/rdbms [get]
func ListRDBMS(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	res, err := migration.ListRDBMS(nsId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Msg("Failed to list managed RDBMS instances")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(res, "Successfully retrieved managed RDBMS list"))
}

// GetRDBMS godoc
// @ID GetRDBMS
// @Summary Get details of a migrated Managed RDBMS (RDS) instance
// @Description Retrieve details of a specific migrated managed RDBMS / RDS instance in the namespace
// @Tags [Migration] Managed RDBMS
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param rdbmsId path string true "RDBMS Instance ID"
// @Success 200 {object} model.ApiResponse[rdbmsmodel.RDBMSInfo] "Successfully retrieved RDBMS details"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/rdbms/{rdbmsId} [get]
func GetRDBMS(c echo.Context) error {
	nsId := c.Param("nsId")
	rdbmsId := c.Param("rdbmsId")
	if nsId == "" || rdbmsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId and rdbmsId required"))
	}

	res, err := migration.GetRDBMS(nsId, rdbmsId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Failed to get managed RDBMS details")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(res, "Successfully retrieved managed RDBMS details"))
}

// DeleteRDBMS godoc
// @ID DeleteRDBMS
// @Summary Delete a migrated Managed RDBMS (RDS) instance
// @Description Delete a specific migrated managed RDBMS / RDS instance in the namespace
// @Tags [Migration] Managed RDBMS
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param rdbmsId path string true "RDBMS Instance ID"
// @Param option query string false "Deletion option (e.g., 'force')"
// @Success 200 {object} model.ApiResponse[any] "Successfully deleted managed RDBMS instance"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/rdbms/{rdbmsId} [delete]
func DeleteRDBMS(c echo.Context) error {
	nsId := c.Param("nsId")
	rdbmsId := c.Param("rdbmsId")
	option := c.QueryParam("option")

	if nsId == "" || rdbmsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId and rdbmsId required"))
	}

	if err := migration.DeleteRDBMS(nsId, rdbmsId, option); err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Failed to delete managed RDBMS instance")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.SimpleSuccessResponse(fmt.Sprintf("Managed RDBMS '%s' deleted successfully", rdbmsId)))
}

// ============================================================================
// Inner Database Management Handlers
// ============================================================================

// CreateRDBMSDatabase godoc
// @ID CreateRDBMSDatabase
// @Summary Create a logical database inside a Managed RDBMS (RDS) instance
// @Description Create a new logical database inside an existing managed RDBMS instance
// @Tags [Migration] Managed RDBMS
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param rdbmsId path string true "RDBMS Instance ID"
// @Param request body rdbmsmodel.RDBMSDatabaseCreateReq true "Database creation request"
// @Success 201 {object} model.ApiResponse[any] "Successfully created logical database"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database [post]
func CreateRDBMSDatabase(c echo.Context) error {
	nsId := c.Param("nsId")
	rdbmsId := c.Param("rdbmsId")
	if nsId == "" || rdbmsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId and rdbmsId required"))
	}

	var req rdbmsmodel.RDBMSDatabaseCreateReq
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if err := migration.CreateRDBMSDatabase(nsId, rdbmsId, req); err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Failed to create logical database")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SimpleSuccessResponse(fmt.Sprintf("Logical database '%s' created successfully", req.DatabaseName)))
}

// ListRDBMSDatabases godoc
// @ID ListRDBMSDatabases
// @Summary List logical databases inside a Managed RDBMS (RDS) instance
// @Description Retrieve the list of logical databases inside an existing managed RDBMS instance
// @Tags [Migration] Managed RDBMS
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param rdbmsId path string true "RDBMS Instance ID"
// @Param X-Admin-User-Password header string false "Admin User Password"
// @Success 200 {object} model.ApiResponse[rdbmsmodel.RDBMSDatabaseListResponse] "Successfully retrieved logical databases"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database [get]
func ListRDBMSDatabases(c echo.Context) error {
	nsId := c.Param("nsId")
	rdbmsId := c.Param("rdbmsId")
	adminPass := c.Request().Header.Get("X-Admin-User-Password")

	if nsId == "" || rdbmsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId and rdbmsId required"))
	}

	res, err := migration.ListRDBMSDatabases(nsId, rdbmsId, adminPass)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Failed to list logical databases")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(res, "Successfully retrieved logical databases"))
}

// DeleteRDBMSDatabase godoc
// @ID DeleteRDBMSDatabase
// @Summary Delete a logical database inside a Managed RDBMS (RDS) instance
// @Description Delete a logical database inside an existing managed RDBMS instance
// @Tags [Migration] Managed RDBMS
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param rdbmsId path string true "RDBMS Instance ID"
// @Param dbName path string true "Database Name"
// @Param X-Admin-User-Password header string false "Admin User Password"
// @Success 200 {object} model.ApiResponse[any] "Successfully deleted logical database"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/rdbms/{rdbmsId}/database/{dbName} [delete]
func DeleteRDBMSDatabase(c echo.Context) error {
	nsId := c.Param("nsId")
	rdbmsId := c.Param("rdbmsId")
	dbName := c.Param("dbName")
	adminPass := c.Request().Header.Get("X-Admin-User-Password")

	if nsId == "" || rdbmsId == "" || dbName == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId, rdbmsId, and dbName required"))
	}

	if err := migration.DeleteRDBMSDatabase(nsId, rdbmsId, dbName, adminPass); err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Str("dbName", dbName).Msg("Failed to delete logical database")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.SimpleSuccessResponse(fmt.Sprintf("Logical database '%s' deleted successfully", dbName)))
}
