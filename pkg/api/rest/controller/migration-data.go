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

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
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
// @Success 200 {object} common.SimpleMsg "OK"
// @Failure 400 {object} common.SimpleMsg "Bad Request"
// @Failure 404 {object} common.SimpleMsg "Not Found"
// @Failure 500 {object} common.SimpleMsg "Internal Server Error"
// @Router /migration/data [post]
func MigrateData(c echo.Context) error {

	req := new(transx.DataMigrationModel)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("failed to bind the request")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	err := transx.Validate(*req)
	if err != nil {
		log.Error().Err(err).Msg("invalid request")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	if !(req.Source.IsRemote() && req.Destination.IsRemote()) {
		err := fmt.Errorf("%s", "both source and destination should be remote endpoints")
		log.Error().Err(err).Msg("invalid request")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	log.Info().Msgf("Migrate data from %s (%s) to %s (%s)", req.Source.GetEndpoint(), req.Source.DataPath, req.Destination.GetEndpoint(), req.Destination.DataPath)

	err = transx.Transfer(*req)
	if err != nil {
		log.Error().Err(err).Msg("failed to migrate data")
		return c.JSON(http.StatusInternalServerError, common.SimpleMsg{Message: err.Error()})
	}

	log.Info().Msg("Data migration completed successfully")
	return c.JSON(http.StatusOK, common.SimpleMsg{Message: "Data migration completed successfully"})

}
