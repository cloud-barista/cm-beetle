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
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// NLB Migration API
// ============================================================================

// MigrateNlbs godoc
// @ID MigrateNlbs
// @Summary (Preview) Migrate NLBs to a cloud infra
// @Description Migrate NLBs to the target cloud infra based on recommendation results.
// @Description
// @Description [Prerequisites]
// @Description - The target Namespace (nsId) must exist.
// @Description - The target Infra (infraId) must exist and have at least one NodeGroup in Running state.
// @Description - Each `targetNlbList[].targetGroup.nodeGroupId` must reference an existing NodeGroup in the Infra.
// @Description
// @Description [Note] Input should be the `targetNlbList` field from the POST /recommendation/infraWithNlb response.
// @Description Ensure `targetGroup.nodeGroupId` matches the NodeGroup IDs created during infra migration.
// @Description
// @Description [Note] All NLBs are attempted independently. Partial success is possible.
// @Description
// @Description By default this API runs synchronously. Send header `Prefer: respond-async` to run it
// @Description asynchronously instead: receive 202 Accepted with a reqId, then poll GET /request/{reqId}
// @Description (status flow: Handling → Success / Error). Only the "respond-async" token is recognized.
// @Tags [Migration] Managed Network Load Balancer (NLB) - preview
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param infraId path string true "Infra ID (target infra with NodeGroups already created)"
// @Param useExisting query bool false "Reuse existing NLB if one targeting the same nodeGroupId already exists, instead of creating a new one (default: true)"
// @Param request body cloudmodel.RecommendedNlb true "NLB migration request (use targetNlbList[] from /recommendation/infraWithNlb)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided)"
// @Param Prefer header string false "Set to 'respond-async' to run this migration asynchronously (RFC 7240)" Enums(respond-async)
// @Success 201 {object} model.ApiResponse[cloudmodel.MigratedNlbResult] "NLBs created successfully"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Migration started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during NLB creation"
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /migration/middleware/ns/{nsId}/infra/{infraId}/nlb [post]
func MigrateNlbs(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	infraId := c.Param("infraId")
	if infraId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("infraId required"))
	}

	var req cloudmodel.RecommendedNlb
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind MigrateNlbs request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if len(req.TargetNlbList) == 0 {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("targetNlbList is required and must not be empty"))
	}

	log.Info().
		Str("nsId", nsId).
		Str("infraId", infraId).
		Int("count", len(req.TargetNlbList)).
		Msg("Starting NLB migration")

	// Parse useExisting parameter (default: true) — same pattern as infra migration
	useExisting := true
	if c.QueryParam("useExisting") == "false" {
		useExisting = false
	}

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (cloudmodel.MigratedNlbResult, error) {
			return migration.CreateNlbs(nsId, infraId, req, useExisting)
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
			"NLB migration started. Use GET /request/{reqId} to check status."))
	}

	result, err := migration.CreateNlbs(nsId, infraId, req, useExisting)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("infraId", infraId).Msg("NLB migration failed")
		if strings.Contains(err.Error(), "all NLB migrations failed") {
			return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SuccessResponseWithMessage(result, result.Description))
}

// ============================================================================
// NLB Management APIs
// ============================================================================

// ListNlbs godoc
// @ID ListNlbs
// @Summary List NLBs in a cloud infra
// @Description Get the list of all NLBs in the specified namespace and infra
// @Tags [Migration] Managed Network Load Balancer (NLB) - preview
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param infraId path string true "Infra ID"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[[]cloudmodel.NLBInfo] "NLB list"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/infra/{infraId}/nlb [get]
func ListNlbs(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	infraId := c.Param("infraId")
	if infraId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("infraId required"))
	}

	infos, err := migration.ListNlbs(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("infraId", infraId).Msg("Failed to list NLBs")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to list NLBs: %v", err)))
	}

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(infos,
		fmt.Sprintf("Listed %d NLB(s)", len(infos))))
}

// GetNlb godoc
// @ID GetNlb
// @Summary Get NLB details
// @Description Get details of a specific NLB
// @Tags [Migration] Managed Network Load Balancer (NLB) - preview
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param infraId path string true "Infra ID"
// @Param nlbId path string true "NLB ID"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[cloudmodel.NLBInfo] "NLB details"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "NLB not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId} [get]
func GetNlb(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	infraId := c.Param("infraId")
	if infraId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("infraId required"))
	}

	nlbId := c.Param("nlbId")
	if nlbId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nlbId required"))
	}

	info, err := migration.GetNlb(nsId, infraId, nlbId)
	if err != nil {
		log.Error().Err(err).Str("nlbId", nlbId).Msg("Failed to get NLB")
		errLower := strings.ToLower(err.Error())
		if strings.Contains(errLower, "not found") || strings.Contains(errLower, "does not exist") {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("NLB '%s' not found", nlbId)))
		}
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to get NLB '%s': %v", nlbId, err)))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(info))
}

// GetNlbHealth godoc
// @ID GetNlbHealth
// @Summary Get NLB health status (live CSP check)
// @Description Perform a live health check on NLB backend targets via the CSP.
// @Description Unlike GET /nlb/{nlbId} (which returns cached state), this endpoint queries the CSP directly
// @Description to retrieve the current health status of VM targets in the NLB target group.
// @Tags [Migration] Managed Network Load Balancer (NLB) - preview
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param infraId path string true "Infra ID"
// @Param nlbId path string true "NLB ID"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[cloudmodel.NLBInfo] "NLB health info (live from CSP)"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "NLB not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}/healthz [get]
func GetNlbHealth(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	infraId := c.Param("infraId")
	if infraId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("infraId required"))
	}

	nlbId := c.Param("nlbId")
	if nlbId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nlbId required"))
	}

	info, err := migration.GetNlbHealth(nsId, infraId, nlbId)
	if err != nil {
		log.Error().Err(err).Str("nlbId", nlbId).Msg("Failed to get NLB health")
		errLower := strings.ToLower(err.Error())
		if strings.Contains(errLower, "not found") || strings.Contains(errLower, "does not exist") {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("NLB '%s' not found", nlbId)))
		}
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to get NLB health '%s': %v", nlbId, err)))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(info))
}

// DeleteNlb godoc
// @ID DeleteNlb
// @Summary Delete an NLB
// @Description Delete a specific NLB from the target infra.
// @Description
// @Description [Note] Some CSPs delete NLBs asynchronously — the API returns success before ENIs are fully released.
// @Description Deleting VNet/subnets immediately after NLB deletion may cause dependency errors (e.g., DependencyViolation on AWS).
// @Description CM-Beetle waits a short period (e.g., 15s) after a successful deletion response to allow CSP-side cleanup to complete.
// @Description
// @Description By default this API runs synchronously (always includes the 15s settle wait). Send header
// @Description `Prefer: respond-async` to run it asynchronously instead: receive 202 Accepted with a reqId,
// @Description then poll GET /request/{reqId} (status flow: Handling → Success / Error).
// @Tags [Migration] Managed Network Load Balancer (NLB) - preview
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param infraId path string true "Infra ID"
// @Param nlbId path string true "NLB ID"
// @Param X-Request-Id header string false "Unique request ID"
// @Param Prefer header string false "Set to 'respond-async' to run this deletion asynchronously (RFC 7240)" Enums(respond-async)
// @Success 204 "NLB deleted (includes 15s settle wait for CSP async cleanup)"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Deletion started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId} [delete]
func DeleteNlb(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	infraId := c.Param("infraId")
	if infraId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("infraId required"))
	}

	nlbId := c.Param("nlbId")
	if nlbId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nlbId required"))
	}

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (map[string]any, error) {
			if err := migration.DeleteNlb(nsId, infraId, nlbId); err != nil {
				return nil, err
			}
			return map[string]any{"message": fmt.Sprintf("NLB '%s' deleted successfully", nlbId)}, nil
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
			"NLB deletion started. Use GET /request/{reqId} to check status."))
	}

	if err := migration.DeleteNlb(nsId, infraId, nlbId); err != nil {
		log.Error().Err(err).Str("nlbId", nlbId).Msg("Failed to delete NLB")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}
