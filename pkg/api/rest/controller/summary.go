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

// Package controller has handlers for infrastructure summary APIs
package controller

import (
	"fmt"
	"net/http"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// GenerateTargetInfraSummary godoc
// @ID GenerateTargetInfraSummary
// @Summary Generate target infrastructure summary
// @Description Generate a comprehensive target infrastructure summary in multiple formats based on 'format' query parameter:
// @Description
// @Description **Response Format by 'format' Parameter:**
// @Description - `format=md` (default): Returns markdown string with Content-Type: text/markdown; charset=utf-8
// @Description - `format=html`: Returns HTML string with Content-Type: text/html; charset=utf-8
// @Description - `format=json`: Returns ApiResponse[TargetInfraSummary] with Content-Type: application/json
// @Description
// @Description **Note:** API documentation shows JSON schema for reference, but actual default response is markdown format.
// @Description
// @Description **Markdown example**: https://github.com/cloud-barista/cm-beetle/blob/main/cmd/test-cli/testresult/beetle-summary-target-aws.md
// @Description
// @Description **Download Behavior:**
// @Description - `download=false` (default): Content displayed inline (viewable in browser/Swagger UI)
// @Description - `download=true`: Content downloaded as file (Content-Disposition: attachment)
// @Tags [Summary/Report] Infrastructure Analysis for Migration
// @Accept  json
// @Produce  json
// @Produce  text/markdown
// @Produce  text/html
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciId path string true "Multi-Cloud Infrastructure (MCI) ID" default(mmci01)
// @Param format query string false "Summary format: md, html, or json" Enums(md,html,json) default(md)
// @Param download query string false "Download as file: true for file download, false for inline display (only affects browsers/Swagger UI, not curl)" Enums(true,false) default(false)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[summary.TargetInfraSummary] "Successfully generated target infrastructure summary (format varies by 'format' parameter)"
// @Header 200 {string} Content-Disposition "inline; filename=\"target-summary.md\" or \"target-summary.html\" (or attachment when download=true)"
// @Header 200 {string} Content-Type "text/markdown; charset=utf-8 or text/html; charset=utf-8"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during summary generation"
// @Router /summary/target/ns/{nsId}/mci/{mciId} [get]
func GenerateTargetInfraSummary(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("Namespace ID is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Namespace ID required"))
	}

	mciId := c.Param("mciId")
	if mciId == "" {
		log.Warn().Msg("MCI ID is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("MCI ID required"))
	}

	format := c.QueryParam("format")
	if format == "" {
		format = "md" // default format
	}
	if format != "json" && format != "md" && format != "html" {
		log.Warn().Msgf("Invalid format: %s", format)
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Format must be 'json', 'md', or 'html'"))
	}

	download := c.QueryParam("download")
	if download == "" {
		download = "false" // default: inline display
	}
	if download != "true" && download != "false" {
		log.Warn().Msgf("Invalid download parameter: %s", download)
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Download parameter must be 'true' or 'false'"))
	}

	// [Process]
	log.Info().Msgf("Generating infrastructure summary (nsId: %s, mciId: %s, format: %s, download: %s)", nsId, mciId, format, download)

	// Generate the infrastructure summary
	infraSummary, err := summary.GenerateInfraSummary(nsId, mciId)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate infrastructure summary")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Summary generation failed"))
	}

	// [Output]
	if format == "md" || format == "html" {
		// Generate markdown summary
		markdownSummary := summary.GenerateMarkdownSummary(infraSummary)

		var content []byte
		var contentType string
		var fileExtension string

		if format == "html" {
			// Convert markdown to HTML
			content = summary.ConvertMarkdownToHTML([]byte(markdownSummary))
			contentType = "text/html; charset=utf-8"
			fileExtension = "html"
		} else {
			// Return as markdown
			content = []byte(markdownSummary)
			contentType = "text/markdown; charset=utf-8"
			fileExtension = "md"
		}

		// Set Content-Disposition header based on download parameter
		// - "inline": displays content in browser/Swagger UI (allows both viewing and downloading)
		// - "attachment": forces file download in browser (content not displayed in Swagger UI response body)
		filename := fmt.Sprintf("target-summary-%s-%s.%s", nsId, mciId, fileExtension)
		dispositionType := "inline"
		if download == "true" {
			dispositionType = "attachment"
		}
		disposition := dispositionType + "; filename=\"" + filename + "\""
		c.Response().Header().Set(echo.HeaderContentDisposition, disposition)

		// Return with proper Content-Type
		return c.Blob(http.StatusOK, contentType, content)
	}

	// Return JSON format (default)
	return c.JSON(http.StatusOK, infraSummary)
}

// GenerateSourceInfraSummaryRequest represents the request body for source infrastructure summary
type GenerateSourceInfraSummaryRequest struct {
	OnpremiseInfraModel onpremmodel.OnpremInfra `json:"onpremiseInfraModel" validate:"required"`
}

// GenerateSourceInfraSummary godoc
// @ID GenerateSourceInfraSummary
// @Summary Generate source infrastructure summary
// @Description Generate a comprehensive source infrastructure summary from on-premise data in multiple formats based on 'format' query parameter:
// @Description
// @Description **Response Format by 'format' Parameter:**
// @Description - `format=json`: Returns ApiResponse[SourceInfraSummary] with Content-Type: application/json
// @Description - `format=md` (default): Returns markdown string with Content-Type: text/markdown; charset=utf-8
// @Description - `format=html`: Returns HTML string with Content-Type: text/html; charset=utf-8
// @Description
// @Description **Note:** API documentation shows JSON schema for reference, but actual default response is markdown format.
// @Description
// @Description **Markdown example**: https://github.com/cloud-barista/cm-beetle/blob/main/cmd/test-cli/testresult/beetle-summary-source.md
// @Description
// @Description **Download Behavior:**
// @Description - `download=false` (default): Content displayed inline (viewable in browser/Swagger UI)
// @Description - `download=true`: Content downloaded as file (Content-Disposition: attachment)
// @Tags [Summary/Report] Infrastructure Analysis for Migration
// @Accept  json
// @Produce  json
// @Produce  text/markdown
// @Produce  text/html
// @Param format query string false "Summary format: md, html, or json" Enums(md,html,json) default(md)
// @Param download query string false "Download as file: true for file download, false for inline display (only affects browsers/Swagger UI, not curl)" Enums(true,false) default(false)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Request body controller.GenerateSourceInfraSummaryRequest true "Source infrastructure data"
// @Success 200 {object} model.ApiResponse[summary.SourceInfraSummary] "Successfully generated source infrastructure summary (format varies by 'format' parameter)"
// @Header 200 {string} Content-Disposition "inline; filename=\"source-summary.md\" or \"source-summary.html\" (or attachment when download=true)"
// @Header 200 {string} Content-Type "text/markdown; charset=utf-8 or text/html; charset=utf-8"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during summary generation"
// @Router /summary/source [post]
func GenerateSourceInfraSummary(c echo.Context) error {

	// [Input]
	format := c.QueryParam("format")
	if format == "" {
		format = "md" // default format
	}
	if format != "json" && format != "md" && format != "html" {
		log.Warn().Msgf("Invalid format: %s", format)
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Format must be 'json', 'md', or 'html'"))
	}

	download := c.QueryParam("download")
	if download == "" {
		download = "false" // default: inline display
	}
	if download != "true" && download != "false" {
		log.Warn().Msgf("Invalid download parameter: %s", download)
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Download parameter must be 'true' or 'false'"))
	}

	// Bind request body
	var req GenerateSourceInfraSummaryRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// [Process]
	// Use a default infrastructure name if not provided
	infraName := "on-premise-infra"
	if len(req.OnpremiseInfraModel.Servers) > 0 {
		infraName = fmt.Sprintf("infra-%d-servers", len(req.OnpremiseInfraModel.Servers))
	}

	log.Info().Msgf("Generating source infrastructure summary (infraName: %s, format: %s, download: %s)", infraName, format, download)

	// Generate the source infrastructure summary
	sourceSummary, err := summary.GenerateSourceInfraSummary(infraName, req.OnpremiseInfraModel)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate source infrastructure summary")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Summary generation failed"))
	}

	// [Output]
	if format == "md" || format == "html" {
		// Generate markdown summary
		markdownSummary := summary.GenerateSourceMarkdownSummary(sourceSummary)

		var content []byte
		var contentType string
		var fileExtension string

		if format == "html" {
			// Convert markdown to HTML
			content = summary.ConvertMarkdownToHTML([]byte(markdownSummary))
			contentType = "text/html; charset=utf-8"
			fileExtension = "html"
		} else {
			// Return as markdown
			content = []byte(markdownSummary)
			contentType = "text/markdown; charset=utf-8"
			fileExtension = "md"
		}

		// Set Content-Disposition header based on download parameter
		// - "inline": displays content in browser/Swagger UI (allows both viewing and downloading)
		// - "attachment": forces file download in browser (content not displayed in Swagger UI response body)
		filename := fmt.Sprintf("source-summary-%s.%s", infraName, fileExtension)
		dispositionType := "inline"
		if download == "true" {
			dispositionType = "attachment"
		}
		disposition := dispositionType + "; filename=\"" + filename + "\""
		c.Response().Header().Set(echo.HeaderContentDisposition, disposition)

		// Return with proper Content-Type
		return c.Blob(http.StatusOK, contentType, content)
	}

	// Return JSON format (default)
	return c.JSON(http.StatusOK, sourceSummary)
}
