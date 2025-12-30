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
// @Description * Only relay mode is supported for now (both source and destination should be remote endpoints).
// @Description * Supported methods: rsync, object storage
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
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid data migration configuration"))
	}

	if !(req.Source.IsRemote() && req.Destination.IsRemote()) {
		err := fmt.Errorf("both source and destination must be remote endpoints")
		log.Error().Err(err).Msg("invalid request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	log.Info().Msgf("Migrate data from %s (%s) to %s (%s)", req.Source.GetEndpoint(), req.Source.DataPath, req.Destination.GetEndpoint(), req.Destination.DataPath)

	// Start time measurement
	startTime := time.Now()

	err = transx.Transfer(*req)

	// Calculate elapsed time
	elapsedTime := time.Since(startTime)

	if err != nil {
		log.Error().Err(err).Dur("elapsedTime", elapsedTime).Msg("failed to migrate data")
		errorMsg := fmt.Sprintf("Data migration failed (%s)", elapsedTime.Round(time.Millisecond))
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(errorMsg))
	}

	log.Info().Dur("elapsedTime", elapsedTime).Msg("Data migration completed successfully")
	successMsg := fmt.Sprintf("Data migrated successfully (%s)", elapsedTime.Round(time.Millisecond))
	return c.JSON(http.StatusOK, model.SimpleSuccessResponse(successMsg))
}
