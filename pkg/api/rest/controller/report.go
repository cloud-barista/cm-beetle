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

// Package controller has handlers for report APIs
package controller

import (
	"fmt"
	"net/http"

	model "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/beetle"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// GetInfraReport godoc
// @ID GetInfraReport
// @Summary Get migration infrastructure report
// @Description Get a comprehensive migration infrastructure report in JSON or Markdown format
// @Tags [Report] Source/Target Infrastructure
// @Accept  json
// @Produce  json
// @Produce  text/markdown
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciId path string true "Multi-Cloud Infrastructure (MCI) ID" default(mmci01)
// @Param format query string false "Report format: json or md" Enums(json,md) default(json)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} migration.MigrationReport "Successfully generated migration report (JSON format)"
// @Success 200 {string} string "Successfully generated migration report (Markdown format)"
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /report/ns/{nsId}/mci/{mciId} [get]
func GetInfraReport(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	mciId := c.Param("mciId")
	if mciId == "" {
		err := fmt.Errorf("invalid request, the multi-cloud infrastructure ID (mciId: %s) is required", mciId)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	format := c.QueryParam("format")
	if format == "" {
		format = "json" // default format
	}
	if format != "json" && format != "md" {
		err := fmt.Errorf("invalid request, the format (format: %s) must be 'json' or 'md'", format)
		log.Warn().Msg(err.Error())
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Process]
	log.Info().Msgf("Generating migration report (nsId: %s, mciId: %s, format: %s)", nsId, mciId, format)

	// Generate the migration report
	report, err := migration.GenerateMigrationReport(nsId, mciId)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate migration report")
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// [Output]
	if format == "md" {
		// Return markdown format
		markdownReport := migration.GenerateMarkdownReport(report)
		return c.String(http.StatusOK, markdownReport)
	}

	// Return JSON format (default)
	return c.JSON(http.StatusOK, report)
}
