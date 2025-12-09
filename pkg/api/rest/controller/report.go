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
	"fmt"
	"net/http"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/report"
	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
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
// @Description Generate a comprehensive migration report comparing source infrastructure with target cloud VMs, including resource mappings, network/security analysis, cost summary, and recommendations in Markdown or HTML format
// @Tags [Summary/Report] Infrastructure Analysis for Migration
// @Accept json
// @Produce text/markdown
// @Produce text/html
// @Param nsId path string true "Namespace ID" example("mig01") default(mig01)
// @Param mciId path string true "MCI ID" example("mmci01") default(mmci01)
// @Param format query string false "Report format: md or html" Enums(md,html) default(md)
// @Param download query string false "Download as file: true for file download, false for inline display (only affects browsers/Swagger UI, not curl)" Enums(true,false) default(false)
// @Param onpremiseInfraModel body controller.GenerateMigrationReportRequest true "Source infrastructure data from on-premise"
// @Success 200 {string} string "Migration report in markdown or HTML format"
// @Header 200 {string} Content-Disposition "inline; filename=\"migration-report.md\" or \"migration-report.html\" (or attachment when download=true)"
// @Header 200 {string} Content-Type "text/markdown; charset=utf-8 or text/html; charset=utf-8"
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

	// Extract query parameters
	format := c.QueryParam("format")
	if format == "" {
		format = "md" // default format
	}
	if format != "md" && format != "html" {
		err := fmt.Errorf("invalid request, the format (format: %s) must be 'md' or 'html'", format)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Text:    err.Error(),
		})
	}

	download := c.QueryParam("download")
	if download == "" {
		download = "false" // default: inline display
	}
	if download != "true" && download != "false" {
		err := fmt.Errorf("invalid request, the download (download: %s) must be 'true' or 'false'", download)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			Success: false,
			Text:    err.Error(),
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
		Str("format", format).
		Str("download", download).
		Msg("Generating migration report")

	migrationReport, err := report.GenerateMigrationReport(nsId, mciId, req.OnpremiseInfraModel)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate migration report")
		return c.JSON(http.StatusInternalServerError, model.Response{
			Success: false,
			Text:    "Failed to generate migration report: " + err.Error(),
		})
	}

	// [Output]
	// Generate markdown report
	markdownReport := report.GenerateMigrationReportMarkdown(migrationReport)

	var content []byte
	var contentType string
	var fileExtension string

	if format == "html" {
		// Convert markdown to HTML
		content = summary.ConvertMarkdownToHTML([]byte(markdownReport))
		contentType = "text/html; charset=utf-8"
		fileExtension = "html"
	} else {
		// Return as markdown (default)
		content = []byte(markdownReport)
		contentType = "text/markdown; charset=utf-8"
		fileExtension = "md"
	}

	// Set Content-Disposition header based on download parameter
	// - "inline": displays content in browser/Swagger UI (allows both viewing and downloading)
	// - "attachment": forces file download in browser (content not displayed in Swagger UI response body)
	filename := fmt.Sprintf("migration-report-%s-%s.%s", nsId, mciId, fileExtension)
	dispositionType := "inline"
	if download == "true" {
		dispositionType = "attachment"
	}
	disposition := dispositionType + "; filename=\"" + filename + "\""
	c.Response().Header().Set(echo.HeaderContentDisposition, disposition)

	// Return with proper Content-Type
	return c.Blob(http.StatusOK, contentType, content)
}
