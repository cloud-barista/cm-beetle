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

package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog/log"
)

/*
 * Multi-target Infrastructure Recommendation (cross-CSP comparison)
 */

// RecommendMultiInfraRequest is the request body for POST /recommendation/multiInfra.
type RecommendMultiInfraRequest struct {
	DesiredCspAndRegionPairs []cloudmodel.CloudProperty `json:"desiredCspAndRegionPairs" validate:"required,min=2,max=10,dive"`
	SourceInfra              onpremmodel.OnpremInfra    `json:"sourceInfra" validate:"required"`
}

// validateMultiInfraTargets checks the target count bounds and rejects duplicate CSP/region
// pairs, so each response item maps back to exactly one requested target unambiguously.
func validateMultiInfraTargets(pairs []cloudmodel.CloudProperty) (string, bool) {
	if len(pairs) < recommendation.MinMultiInfraTargets {
		return fmt.Sprintf("At least %d target CSP/region pairs required", recommendation.MinMultiInfraTargets), false
	}
	if len(pairs) > recommendation.MaxMultiInfraTargets {
		return fmt.Sprintf("At most %d target CSP/region pairs allowed", recommendation.MaxMultiInfraTargets), false
	}

	seen := make(map[string]bool, len(pairs))
	for _, pair := range pairs {
		if pair.Csp == "" || pair.Region == "" {
			return "Each target requires both csp and region", false
		}
		key := strings.ToLower(pair.Csp) + "/" + strings.ToLower(pair.Region)
		if seen[key] {
			return fmt.Sprintf("Duplicate target: %s", key), false
		}
		seen[key] = true
	}

	return "", true
}

// parseMinMatchRate parses the minMatchRate query param, defaulting to 90.0.
func parseMinMatchRate(c echo.Context) float64 {
	minMatchRate := 90.0
	if qp := c.QueryParam("minMatchRate"); qp != "" {
		if parsed, err := strconv.ParseFloat(qp, 64); err != nil {
			log.Warn().Err(err).Msgf("invalid minMatchRate value: %s, using default 90.0", qp)
		} else if parsed < 0 || parsed > 100 {
			log.Warn().Msgf("minMatchRate out of range [0-100]: %.1f, using default 90.0", parsed)
		} else {
			minMatchRate = parsed
		}
	}
	return minMatchRate
}

// RecommendMultiInfraCandidates godoc
// @ID RecommendMultiInfraCandidates
// @Summary Recommend the best-match infrastructure per target cloud, for cross-CSP comparison
// @Description Recommend a single best-effort infrastructure candidate for each of several target CSP/region pairs.
// @Description
// @Description Use this API to compare candidate clouds before committing to one; once a target is chosen,
// @Description use `POST /recommendation/infra` against that single CSP/region to explore multiple candidates.
// @Description
// @Description **[Required Parameter: `desiredCspAndRegionPairs`]** 2 to 10 target CSP/region pairs (project scope: 10 supported CSPs).
// @Description Duplicate pairs are rejected.
// @Description
// @Description **[Response]** Always returns exactly one item per requested target, in request order
// @Description (`len(data) == len(desiredCspAndRegionPairs)`), so items map back to targets via `targetCloud`
// @Description without needing to be re-sorted or grouped. A target that fails validation or yields no
// @Description compatible infrastructure still produces one item, with `status` set to `failed` or
// @Description `nothing-to-recommend` and `targetInfra` left empty.
// @Description
// @Description **[Optional Parameter: `minMatchRate`]** Minimum match rate threshold for highly-matched classification (default: 90.0, range: 0-100)
// @Description
// @Description [Note] Each target costs roughly as much as one `/recommendation/infra` call; requests with
// @Description several targets can take a while. `Prefer: respond-async` is strongly recommended.
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param request body RecommendMultiInfraRequest true "Target CSP/region pairs and the source infrastructure to be migrated"
// @Param minMatchRate query number false "Minimum match rate for highly-matched classification (default: 90.0, range: 0-100)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this recommendation asynchronously (RFC 7240)" Enums(respond-async)
// @Success 200 {object} model.ApiResponse[[]cloudmodel.RecommendedInfra] "One recommended (or failed/nothing-to-recommend) candidate per target, in request order"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Recommendation started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /recommendation/multiInfra [post]
func RecommendMultiInfraCandidates(c echo.Context) error {

	// [Input]
	reqt := &RecommendMultiInfraRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Warn().Err(err).Msg("failed to bind a request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if msg, ok := validateMultiInfraTargets(reqt.DesiredCspAndRegionPairs); !ok {
		log.Warn().Msg(msg)
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(msg))
	}

	minMatchRate := parseMinMatchRate(c)
	pairs := reqt.DesiredCspAndRegionPairs
	sourceInfra := reqt.SourceInfra

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() ([]cloudmodel.RecommendedInfra, error) {
			return recommendation.RecommendMultiInfraCandidates(pairs, sourceInfra, minMatchRate)
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

	// [Process]
	results, err := recommendation.RecommendMultiInfraCandidates(pairs, sourceInfra, minMatchRate)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend multi-target infrastructure candidates")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Recommendation failed"))
	}

	// [Output]
	return c.JSON(http.StatusOK, model.SuccessListResponseWithMessage(results,
		fmt.Sprintf("Recommended infrastructure for %d target(s)", len(results))))
}

// RecommendMultiInfraWithNlbRequest is the request body for POST /recommendation/multiInfraWithNlb.
type RecommendMultiInfraWithNlbRequest struct {
	DesiredCspAndRegionPairs []cloudmodel.CloudProperty `json:"desiredCspAndRegionPairs" validate:"required,min=2,max=10,dive"`
	SourceInfra              onpremmodel.OnpremInfra    `json:"sourceInfra" validate:"required"`
}

// RecommendMultiInfraWithNlbCandidates godoc
// @ID RecommendMultiInfraWithNlbCandidates
// @Summary (Preview) Recommend the best-match NLB-aware infrastructure per target cloud, for cross-CSP comparison
// @Description NLB-aware counterpart of `POST /recommendation/multiInfra`. Recommends a single best-effort
// @Description infrastructure candidate (including NLB mapping) for each of several target CSP/region pairs.
// @Description
// @Description Use this API to compare candidate clouds before committing to one; once a target is chosen,
// @Description use `POST /recommendation/infraWithNlb` against that single CSP/region to explore multiple candidates.
// @Description
// @Description **[Required Parameter: `desiredCspAndRegionPairs`]** 2 to 10 target CSP/region pairs (project scope: 10 supported CSPs).
// @Description Duplicate pairs are rejected.
// @Description
// @Description [Note] `sourceInfra.nlbs` must be populated (HAProxy frontend-backend pairs from cm-honeybee).
// @Description
// @Description **[Response]** Always returns exactly one item per requested target, in request order
// @Description (`len(data) == len(desiredCspAndRegionPairs)`), so items map back to targets via `targetCloud`
// @Description without needing to be re-sorted or grouped. A target that fails validation or yields no
// @Description compatible infrastructure still produces one item, with `status` set to `failed` or
// @Description `nothing-to-recommend` and `targetInfra` left empty.
// @Description
// @Description **[Optional Parameter: `minMatchRate`]** Minimum match rate threshold for highly-matched classification (default: 90.0, range: 0-100)
// @Description
// @Description [Note] Each target costs roughly as much as one `/recommendation/infraWithNlb` call; requests with
// @Description several targets can take a while. `Prefer: respond-async` is strongly recommended.
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param request body RecommendMultiInfraWithNlbRequest true "Target CSP/region pairs and the source infra including NLBs"
// @Param minMatchRate query number false "Minimum match rate for highly-matched classification (default: 90.0, range: 0-100)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this recommendation asynchronously (RFC 7240)" Enums(respond-async)
// @Success 200 {object} model.ApiResponse[[]cloudmodel.RecommendedInfra] "One recommended (or failed/nothing-to-recommend) candidate per target, in request order"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Recommendation started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /recommendation/multiInfraWithNlb [post]
func RecommendMultiInfraWithNlbCandidates(c echo.Context) error {

	// [Input]
	reqt := &RecommendMultiInfraWithNlbRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Warn().Err(err).Msg("failed to bind a request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if msg, ok := validateMultiInfraTargets(reqt.DesiredCspAndRegionPairs); !ok {
		log.Warn().Msg(msg)
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(msg))
	}

	if len(reqt.SourceInfra.Nodes) == 0 {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("sourceInfra.nodes is required"))
	}
	if len(reqt.SourceInfra.NLBs) == 0 {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(
			"sourceInfra.nlbs is required for multiInfraWithNlb; use /recommendation/multiInfra for NLB-free recommendation"))
	}

	minMatchRate := parseMinMatchRate(c)
	pairs := reqt.DesiredCspAndRegionPairs
	sourceInfra := reqt.SourceInfra

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() ([]cloudmodel.RecommendedInfra, error) {
			return recommendation.RecommendMultiInfraWithNlbCandidates(pairs, sourceInfra, minMatchRate)
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

	// [Process]
	results, err := recommendation.RecommendMultiInfraWithNlbCandidates(pairs, sourceInfra, minMatchRate)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend multi-target NLB-aware infrastructure candidates")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Recommendation failed"))
	}

	// [Output]
	return c.JSON(http.StatusOK, model.SuccessListResponseWithMessage(results,
		fmt.Sprintf("Recommended infrastructure for %d target(s)", len(results))))
}
