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

	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// Request Models
// ============================================================================

// MigrateObjectStorageRequest represents a request for object storage migration
type MigrateObjectStorageRequest struct {
	storagemodel.RecommendedObjectStorage
}

// ============================================================================
// Object Storage Migration API
// ============================================================================

// MigrateObjectStorage godoc
// @ID MigrateObjectStorage
// @Summary Migrate object storages to cloud
// @Description Migrate object storages to cloud based on recommendation results
// @Description
// @Description [Note]
// @Description - This API creates object storages (buckets) in the target cloud within the specified namespace
// @Description - Input should be the output from RecommendObjectStorage API
// @Description - Connection name is automatically generated from CSP and region in the request body
// @Description
// @Description [Note] `nameSeed` enables dynamic naming via **Late Binding**.
// @Description - If `nameSeed` is set (e.g., `my`), bucket names are prefixed at migration time: `my-os-01`.
// @Description - If `nameSeed` is empty, bucket names are used as-is from the recommendation result.
// @Description
// @Description [Examples]
// @Description * Test results: https://github.com/cloud-barista/cm-beetle/blob/main/docs/test-results-data-migration.md
// @Description
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce	json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param request body MigrateObjectStorageRequest true "Object storage migration request (use RecommendObjectStorage response)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 201 "Created - Object storages created successfully"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during object storage creation"
// @Router /migration/middleware/ns/{nsId}/objectStorage [post]
func MigrateObjectStorage(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	var req MigrateObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if req.TargetCloud.Csp == "" || req.TargetCloud.Region == "" {
		log.Warn().Msg("CSP and region are required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP and region required"))
	}

	if len(req.TargetObjectStorages) == 0 {
		log.Warn().Msg("At least one target object storage must be provided")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one target bucket required"))
	}

	if err := migration.CreateObjectStorage(nsId, req.RecommendedObjectStorage); err != nil {
		log.Error().Err(err).Msg("Object storage migration failed")
		if strings.Contains(err.Error(), "invalid cloud configuration") {
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusCreated)
}

// ============================================================================
// Object Storage Management APIs
// ============================================================================

// ListObjectStorages godoc
// @ID ListObjectStorages
// @Summary List object storages (buckets)
// @Description Get the list of all object storages (buckets) in the namespace
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[migration.MigratedObjectStorageListResponse] "Successfully retrieved object storage list"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during list operation"
// @Router /migration/middleware/ns/{nsId}/objectStorage [get]
func ListObjectStorages(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	result, err := migration.ListObjectStorages(nsId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Msg("Failed to list object storages")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to list object storages: %v", err)))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}

// GetObjectStorage godoc
// @ID GetObjectStorage
// @Summary Get object storage (bucket) details
// @Description Get details of a specific object storage (bucket)
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[migration.MigratedObjectStorageInfo] "Successfully retrieved object storage details"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during get operation"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId} [get]
func GetObjectStorage(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	osId := c.Param("osId")
	if osId == "" {
		log.Warn().Msg("osId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("osId required"))
	}

	result, err := migration.GetObjectStorage(nsId, osId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to get object storage")
		if strings.Contains(strings.ToLower(err.Error()), "not found") || strings.Contains(strings.ToLower(err.Error()), "does not exist") {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object storage '%s' not found", osId)))
		}
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to get object storage '%s': %v", osId, err)))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}

// ExistObjectStorage godoc
// @ID ExistObjectStorage
// @Summary Check object storage (bucket) existence
// @Description Check if a specific object storage (bucket) exists
// @Description
// @Description [Note]
// @Description - Returns 200 OK if the bucket exists, 404 Not Found if it doesn't exist
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 "OK - Object storage exists"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during existence check"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId} [head]
func ExistObjectStorage(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	osId := c.Param("osId")
	if osId == "" {
		log.Warn().Msg("osId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("osId required"))
	}

	exists, err := migration.ExistObjectStorage(nsId, osId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to check object storage existence")
		if strings.Contains(strings.ToLower(err.Error()), "not found") || strings.Contains(strings.ToLower(err.Error()), "does not exist") {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object storage '%s' not found", osId)))
		}
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to check object storage '%s' existence: %v", osId, err)))
	}

	if !exists {
		return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object storage '%s' not found", osId)))
	}

	return c.NoContent(http.StatusOK)
}

// DeleteObjectStorage godoc
// @ID DeleteObjectStorage
// @Summary Delete object storage (bucket)
// @Description Delete a specific object storage (bucket)
// @Description
// @Description [Note]
// @Description - The bucket must be empty before deletion
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 204 "Object storage deleted successfully"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during deletion"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId} [delete]
func DeleteObjectStorage(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	osId := c.Param("osId")
	if osId == "" {
		log.Warn().Msg("osId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("osId required"))
	}

	if err := migration.DeleteObjectStorage(nsId, osId); err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to delete object storage")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}
