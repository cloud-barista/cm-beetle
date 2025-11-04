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

	model "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/beetle"
	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// GenerateTargetInfraSummary godoc
// @ID GenerateTargetInfraSummary
// @Summary Generate target infrastructure summary
// @Description Generate a comprehensive target infrastructure summary in JSON or Markdown format
// @Tags [Summary/Report] Infrastructure Analysis for Migration
// @Accept  json
// @Produce  json
// @Produce  text/markdown
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciId path string true "Multi-Cloud Infrastructure (MCI) ID" default(mmci01)
// @Param format query string false "Summary format: json or md" Enums(json,md) default(md)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} JSONResult{[MARKDOWN]=string,[JSON]=summary.TargetInfraSummary} "Different return types: json format returns TargetInfraSummary object, md format returns markdown string"
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /summary/target/ns/{nsId}/mci/{mciId} [get]
func GenerateTargetInfraSummary(c echo.Context) error {

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
	log.Info().Msgf("Generating infrastructure summary (nsId: %s, mciId: %s, format: %s)", nsId, mciId, format)

	// Generate the infrastructure summary
	infraSummary, err := summary.GenerateInfraSummary(nsId, mciId)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate infrastructure summary")
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// [Output]
	if format == "md" {
		// Return markdown format
		markdownSummary := summary.GenerateMarkdownSummary(infraSummary)
		return c.String(http.StatusOK, markdownSummary)
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
// @Description Generate a comprehensive source infrastructure summary from on-premise data in JSON or Markdown format
// @Tags [Summary/Report] Infrastructure Analysis for Migration
// @Accept  json
// @Produce  json
// @Produce  text/markdown
// @Param format query string false "Summary format: json or md" Enums(json,md) default(md)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Param Request body controller.GenerateSourceInfraSummaryRequest true "Source infrastructure data"
// @Success 200 {object} JSONResult{[MARKDOWN]=string,[JSON]=summary.SourceInfraSummary} "Different return types: json format returns SourceInfraSummary object, md format returns markdown string"
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /summary/source [post]
func GenerateSourceInfraSummary(c echo.Context) error {

	// [Input]
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

	// Bind request body
	var req GenerateSourceInfraSummaryRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		res := model.Response{
			Success: false,
			Text:    fmt.Sprintf("invalid request body: %v", err),
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Process]
	// Use a default infrastructure name if not provided
	infraName := "on-premise-infra"
	if len(req.OnpremiseInfraModel.Servers) > 0 {
		infraName = fmt.Sprintf("infra-%d-servers", len(req.OnpremiseInfraModel.Servers))
	}

	log.Info().Msgf("Generating source infrastructure summary (infraName: %s, format: %s)", infraName, format)

	// Generate the source infrastructure summary
	sourceSummary, err := summary.GenerateSourceInfraSummary(infraName, req.OnpremiseInfraModel)
	if err != nil {
		log.Error().Err(err).Msg("failed to generate source infrastructure summary")
		res := model.Response{
			Success: false,
			Text:    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// [Output]
	if format == "md" {
		// Return markdown format
		markdownSummary := summary.GenerateSourceMarkdownSummary(sourceSummary)
		return c.String(http.StatusOK, markdownSummary)
	}

	// Return JSON format (default)
	return c.JSON(http.StatusOK, sourceSummary)
}
