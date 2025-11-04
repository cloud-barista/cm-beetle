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

// Package controller provides REST API controllers for migration reports
package controller

import (
	"net/http"

	model "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/beetle"
	"github.com/cloud-barista/cm-beetle/pkg/core/report"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// GenerateMigrationReportRequest represents the request body for generating a migration report
type GenerateMigrationReportRequest struct {
	OnpremiseInfraModel onpremmodel.OnpremInfra `json:"onpremiseInfraModel" validate:"required"`
}

// GenerateMigrationReport godoc
// @ID GenerateMigrationReport
// @Summary Generate migration report (with source-target correlation analysis)
// @Description Generate a comprehensive migration report comparing source infrastructure with target cloud VMs, including resource mappings, network/security analysis, cost summary, and recommendations
// @Tags [Summary/Report] Infrastructure Analysis for Migration
// @Accept json
// @Produce plain
// @Param nsId path string true "Namespace ID" example("mig01") default(mig01)
// @Param mciId path string true "MCI ID" example("mmci01") default(mmci01)
// @Param onpremiseInfraModel body controller.GenerateMigrationReportRequest true "Source infrastructure data from on-premise"
// @Success 200 {string} string "Markdown formatted migration report"
// @Failure 400 {object} model.Response "Invalid request"
// @Failure 500 {object} model.Response "Internal server error"
// @Router /report/migration/ns/{nsId}/mci/{mciId} [post]
func GenerateMigrationReport(c echo.Context) error {
	// Extract path parameters
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Text:    "nsId is required",
		})
	}

	mciId := c.Param("mciId")
	if mciId == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Text:    "mciId is required",
		})
	}

	// Parse request body
	var req GenerateMigrationReportRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Text:    "Invalid request body: " + err.Error(),
		})
	}

	// Validate source infrastructure
	if len(req.OnpremiseInfraModel.Servers) == 0 {
		return c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Text:    "Source infrastructure must contain at least one server",
		})
	}

	// Generate migration report
	log.Info().
		Str("nsId", nsId).
		Str("mciId", mciId).
		Int("sourceServers", len(req.OnpremiseInfraModel.Servers)).
		Msg("Generating migration report")

	migrationReport, err := report.GenerateMigrationReport(nsId, mciId, req.OnpremiseInfraModel)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate migration report")
		return c.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Text:    "Failed to generate migration report: " + err.Error(),
		})
	}

	// Generate markdown
	markdown := report.GenerateMigrationReportMarkdown(migrationReport)

	// Return markdown with proper content type
	return c.String(http.StatusOK, markdown)
}
