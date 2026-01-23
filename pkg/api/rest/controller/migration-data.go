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
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/transx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// MigrateData godoc
// @ID MigrateData
// @Summary Migrate data from source to target
// @Description Migrate data from source to target
// @Description
// @Description [Note]
// @Description * Both source and destination must be remote endpoints (SSH or object storage).
// @Description * Local filesystem access is not allowed for security reasons.
// @Description * Strategy options: auto (default), direct, relay.
// @Description * For SSH endpoints, supports PrivateKey content or PrivateKeyPath.
// @Description
// @Description [Note]
// @Description * Examples(test result): https://github.com/cloud-barista/cm-beetle/blob/main/docs/test-results-data-migration.md
// @Description
// @Tags [Migration] Data (incubating)
// @Accept  json
// @Produce  json
// @Param reqBody body transx.DataMigrationModel true "Request Body"
// @Success 200 {object} model.ApiResponse[string] "Data migrated successfully"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during data migration"
// @Router /migration/data [post]
func MigrateData(c echo.Context) error {

	req := new(transx.DataMigrationModel)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("failed to bind the request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	err := transx.Validate(*req)
	if err != nil {
		log.Error().Err(err).Msg("invalid request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// Security check: Prevent access to local filesystem
	// API users must not access the server's local filesystem
	if req.Source.IsLocal() {
		log.Warn().Msg("rejected: source uses local filesystem")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Local filesystem access not allowed for source; use SSH or object storage"))
	}
	if req.Destination.IsLocal() {
		log.Warn().Msg("rejected: destination uses local filesystem")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Local filesystem access not allowed for destination; use SSH or object storage"))
	}

	log.Info().
		Str("sourceType", req.Source.StorageType).
		Str("sourcePath", req.Source.Path).
		Str("destType", req.Destination.StorageType).
		Str("destPath", req.Destination.Path).
		Str("strategy", req.Strategy).
		Msg("Starting data migration")

	// Start time measurement
	startTime := time.Now()

	// Execute migration
	err = transx.Transfer(*req)

	// Calculate elapsed time
	elapsedTime := time.Since(startTime)

	if err != nil {
		log.Error().Err(err).Dur("elapsedTime", elapsedTime).Msg("failed to migrate data")
		errorMsg := fmt.Sprintf("Data migration failed: %v (%s)", err, elapsedTime.Round(time.Millisecond))
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(errorMsg))
	}

	log.Info().Dur("elapsedTime", elapsedTime).Msg("Data migration completed successfully")
	successMsg := fmt.Sprintf("Data migrated successfully (%s)", elapsedTime.Round(time.Millisecond))
	return c.JSON(http.StatusOK, model.SimpleSuccessResponse(successMsg))
}
