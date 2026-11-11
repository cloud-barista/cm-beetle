/*
Copyright 2024 The Cloud-Barista Authors.
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

// Package controller has handlers for K8s infra recommendation APIs.
package controller

import (
	"context"
	"fmt"
	"net/http"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog/log"
)

/*
 * K8s Infra Recommendation
 */

type RecommendK8sInfraRequest struct {
	cloudmodel.CloudProperty
	onpremmodel.OnpremiseInfraModel
}

type RecommendK8sInfraResponse struct {
	cloudmodel.RecommendedInfra
}

// RecommendK8sInfra godoc
// @ID RecommendK8sInfra
// @Summary Recommend a K8s infra configuration for cloud migration
// @Description Recommend a complete K8s infra configuration based on on-premise infra data from cm-honeybee.
// @Description Returns a RecommendK8sInfraResponse that can be passed directly to the K8s migration API.
// @Description
// @Description **Required Parameters**: `desiredProvider`, `desiredRegion`
// @Description **Input**: On-premise infra model (from honeybee /infra/refined) with K8s cluster info and node roles
// @Tags [Recommendation] K8s Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendK8sInfraRequest true "Source on-premise infra (must include k8sCluster and nodes with role=worker)"
// @Param desiredProvider query string true "Provider (e.g., aws)" Enums(aws,azure,gcp,alibaba,tencent,ibm,ncp,nhn)
// @Param desiredRegion query string true "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this recommendation asynchronously (RFC 7240)" Enums(respond-async)
// @Success 200 {object} model.ApiResponse[RecommendK8sInfraResponse] "K8s infra recommendation (pass directly to POST /migration/ns/{nsId}/k8sCluster)"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Recommendation started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /recommendation/k8sCluster [post]
func RecommendK8sInfra(c echo.Context) error {
	ctx := c.Request().Context()
	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	reqt := &RecommendK8sInfraRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// Query params override request body if body fields are empty
	if reqt.Csp == "" {
		reqt.Csp = desiredProvider
	}
	if reqt.Region == "" {
		reqt.Region = desiredRegion
	}

	if reqt.Csp == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Provider required"))
	}
	if reqt.Region == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Region required"))
	}

	ok, err := recommendation.IsValidCspAndRegion(ctx, reqt.Csp, reqt.Region)
	if !ok {
		log.Error().Err(err).Msg("invalid provider or region")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid provider or region"))
	}

	csp := reqt.Csp
	region := reqt.Region
	sourceInfra := reqt.OnpremiseInfraModel.OnpremiseInfraModel

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(ctx, reqID, func(bgCtx context.Context) (cloudmodel.RecommendedInfra, error) {
			return recommendation.RecommendK8sInfra(bgCtx, csp, region, sourceInfra)
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
			"Recommendation started. Use GET /request/{reqId} to check status."))
	}

	result, err := recommendation.RecommendK8sInfra(ctx, csp, region, sourceInfra)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend K8s infra")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}
