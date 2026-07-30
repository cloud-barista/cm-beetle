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
	"time"

	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog/log"
)

type MigrateInfraWithDefaultsRequest struct {
	// [NOTE] Failed to embed the struct in CB-Tumblebug as follows:
	// infra.InfraDynamicReq

	cloudmodel.InfraDynamicReq
}

type MigrateInfraWithDefaultsResponse struct {
	cloudmodel.VmInfraInfo
}

// MigrateInfraWithDefaults godoc
// @ID MigrateInfraWithDefaults
// @Summary Migrate an infrastructure to the multi-cloud infrastructure (MCI) with defaults (sync by default; async via Prefer: respond-async)
// @Description Migrate an infrastructure to the multi-cloud infrastructure (MCI) with defaults.
// @Description
// @Description By default this API runs synchronously. Send header `Prefer: respond-async` to run it
// @Description asynchronously instead: receive 202 Accepted with a reqId, then poll GET /request/{reqId}
// @Description (status flow: Handling → Success / Error). Only the "respond-async" token is recognized.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param mciInfo body MigrateInfraWithDefaultsRequest true "Specify the information for the targeted mulci-cloud infrastructure (MCI)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this migration asynchronously (RFC 7240)" Enums(respond-async)
// @Success 201 {object} model.ApiResponse[MigrateInfraWithDefaultsResponse] "Successfully migrated to the multi-cloud infrastructure"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Migration started asynchronously - use GET /request/{reqId} to check status"
// @Failure 500 {object} model.ApiResponse[any]
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /migration/ns/{nsId}/infraWithDefaults [post]
func MigrateInfraWithDefaults(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	req := new(MigrateInfraWithDefaultsRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	log.Debug().Msgf("req: %v", req)
	log.Debug().Msgf("req.InfraDynamicReq: %v", req.InfraDynamicReq)

	// [Process]
	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (cloudmodel.VmInfraInfo, error) {
			return migration.CreateInfraWithDefaults(nsId, &req.InfraDynamicReq)
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
			"Migration started. Use GET /request/{reqId} to check status."))
	}

	// Create the infrastructure for migration
	mciInfo, err := migration.CreateInfraWithDefaults(nsId, &req.InfraDynamicReq)

	log.Debug().Msgf("mciInfo: %v", mciInfo)

	// [Output]
	if err != nil {
		log.Error().Err(err).Msg("failed to create the infrastructure")

		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SuccessResponse(mciInfo))
}

// TODO: Check and dev the request and response bodies for the following API

type MigrateInfraRequest struct {
	cloudmodel.RecommendedInfra
}

type MigrateInfraResponse struct {
	cloudmodel.VmInfraInfo
}

// MigrateInfra godoc
// @ID MigrateInfra
// @Summary Migrate an infrastructure to the multi-cloud infrastructure (MCI) with defaults (sync by default; async via Prefer: respond-async)
// @Description Migrate an infrastructure to the multi-cloud infrastructure (MCI) with defaults.
// @Description
// @Description **[Request Field: `nodeGroups[].cspImageName`]** Optional CSP-native image identifier.
// @Description - **Non-empty**: TumbleBug sends this to Spider directly, bypassing the per-node image DB lookup (prevents stale image failures, e.g., Alibaba alibase images).
// @Description - **Empty**: TumbleBug uses `imageId` for the standard DB lookup (may encounter stale images for some CSPs).
// @Description - Recommended: pass the recommendation API response as-is to use the latest resolved image.
// @Description
// @Description By default this API runs synchronously. Send header `Prefer: respond-async` to run it
// @Description asynchronously instead: receive 202 Accepted with a reqId, then poll GET /request/{reqId}
// @Description (status flow: Handling → Success / Error). Only the "respond-async" token is recognized.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param nameSeed query string false "Optional prefix for all resource names (e.g., 'blue' → 'blue-infra101', 'blue-vnet-01'). Applied at migration time."
// @Param useExisting query bool false "Reuse existing resources (VNet, SSH Key, Security Group) if they already exist, instead of creating new ones (default: true)"
// @Param infraInfo body MigrateInfraRequest true "Specify the information for the targeted multi-cloud infrastructure"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this migration asynchronously (RFC 7240)" Enums(respond-async)
// @Success 201 {object} model.ApiResponse[MigrateInfraResponse] "Successfully migrated to the multi-cloud infrastructure"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Migration started asynchronously - use GET /request/{reqId} to check status"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /migration/ns/{nsId}/infra [post]
func MigrateInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	req := new(MigrateInfraRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	// log.Debug().Msgf("req: %+v", req)
	log.Debug().Msgf("req.RecommendedInfra: %+v", req.RecommendedInfra)

	// [Process]
	// Apply NameSeed (Late Binding) from query param before migration.
	// Query param takes precedence; if empty, no prefix is applied.
	nameSeed := c.QueryParam("nameSeed")
	if ok, detail := common.IsValidNameSeed(nameSeed); !ok {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid nameSeed: "+detail))
	}
	infraToMigrate := common.ApplyNameSeed(req.RecommendedInfra, nameSeed)

	// Parse useExisting parameter (default: true)
	useExistingStr := c.QueryParam("useExisting")
	useExisting := true
	if useExistingStr == "false" {
		useExisting = false
	}

	// Validate names and referential integrity
	if ok, detail := common.ValidateComposedNames(infraToMigrate); !ok {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Naming/Reference validation failed: "+detail))
	}

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (cloudmodel.VmInfraInfo, error) {
			if useExisting {
				return migration.CreateInfraWithExisting(nsId, &infraToMigrate)
			}
			return migration.CreateInfra(nsId, &infraToMigrate)
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
			"Migration started. Use GET /request/{reqId} to check status."))
	}

	// Create the infrastructure for migration
	var mciInfo cloudmodel.VmInfraInfo
	var err error
	if useExisting {
		mciInfo, err = migration.CreateInfraWithExisting(nsId, &infraToMigrate)
	} else {
		mciInfo, err = migration.CreateInfra(nsId, &infraToMigrate)
	}

	log.Debug().Msgf("mciInfo: %+v", mciInfo)

	// [Output]
	if err != nil {
		log.Error().Err(err).Msg("failed to create the infrastructure")

		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SuccessResponse(mciInfo))
}

// ListInfra godoc
// @ID ListInfra
// @Summary Get the migrated multi-cloud infrastructure (MCI)
// @Description Get the migrated multi-cloud infrastructure (MCI)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param option query string false "Option for getting the migrated multi-cloud infrastructure" Enums(id)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[cloudmodel.InfraInfoList] "The info list of the migrated multi-cloud infrastructure (Infra)"
// @Success 200 {object} model.ApiResponse[cloudmodel.IdList] "The ID list of The migrated multi-cloud infrastructure (Infra)"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/infra [get]
func ListInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the nanespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	option := c.QueryParam("option")
	if option != "" && option != "id" {
		err := fmt.Errorf("invalid request, the option (option: %s) is invalid", option)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// [Process] List the migrated multi-cloud infrastructures as the option
	switch option {
	case "id":
		idList, err := migration.ListInfraIDs(nsId, option)
		if err != nil {
			log.Error().Err(err).Msg("failed to get the migrated multi-cloud infrastructure IDs")
			return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, model.SuccessResponse(idList))
	default:
		infraInfoList, err := migration.ListAllInfraInfo(nsId)
		if err != nil {
			log.Error().Err(err).Msg("failed to get the migrated multi-cloud infrastructures")
			return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, model.SuccessResponse(infraInfoList))
	}
	// return c.JSON(http.StatusInternalServerError, nil)
}

// GetInfra godoc
// @ID GetInfra
// @Summary Get the migrated multi-cloud infrastructure (MCI)
// @Description Get the migrated multi-cloud infrastructure (MCI)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param infraId path string true "Migrated Cloud Infrastructure ID (the actual ID returned by the migration API; includes NameSeed prefix if used, e.g., 'test-infra101')"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[MigrateInfraResponse] "The migrated multi-cloud infrastructure information"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/infra/{infraId} [get]
func GetInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the nanespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	infraId := c.Param("infraId")
	if infraId == "" {
		err := fmt.Errorf("invalid request, the cloud infrastructure ID (infraId: %s) is required", infraId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// [Process]
	infraInfo, err := migration.GetInfra(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get the migrated multi-cloud infrastructure")

		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "does not exist") {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse("Infrastructure not found"))
		}

		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	// [Ouput]
	return c.JSON(http.StatusOK, model.SuccessResponse(infraInfo))
}

// DeleteInfra godoc
// @ID DeleteInfra
// @Summary Delete the migrated mult-cloud infrastructure (MCI) (sync by default; async via Prefer: respond-async)
// @Description Delete the migrated mult-cloud infrastructure (MCI).
// @Description
// @Description This operation can take a long time (multiple settle-time waits and vNet-deletion
// @Description retries). By default it runs synchronously. Send header `Prefer: respond-async` to run
// @Description it asynchronously instead: receive 202 Accepted with a reqId, then poll GET /request/{reqId}
// @Description (status flow: Handling → Success / Error). Only the "respond-async" token is recognized.
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param infraId path string true "Migrated Cloud Infrastructure ID (the actual ID returned by the migration API; includes NameSeed prefix if used, e.g., 'test-infra101')"
// @Param option query string false "Option for deletion" Enums(terminate,force) default(terminate)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this deletion asynchronously (RFC 7240)" Enums(respond-async)
// @Success 200 {object} model.ApiResponse[any] "The result of deleting the migrated multi-cloud infrastructure"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Deletion started asynchronously - use GET /request/{reqId} to check status"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /migration/ns/{nsId}/infra/{infraId} [delete]
func DeleteInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	// nsId := common.DefaultNamespaceId

	infraId := c.Param("infraId")
	if infraId == "" {
		err := fmt.Errorf("invalid request, the cloud infrastructure ID (infraId: %s) is required", infraId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	option := c.QueryParam("option")
	if option != "" && option != "terminate" && option != "force" {
		err := fmt.Errorf("invalid request, the option (option: %s) is invalid", option)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (common.SimpleMsg, error) {
			return migration.DeleteInfra(nsId, infraId, option)
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
			"Deletion started. Use GET /request/{reqId} to check status."))
	}

	// [Process]
	retMsg, err := migration.DeleteInfra(nsId, infraId, option)

	if err != nil {
		log.Error().Err(err).Msg("failed to delete the migrated multi-cloud infrastructure")

		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "does not exist") {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse("Infrastructure not found"))
		}

		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	// [Ouput]
	return c.JSON(http.StatusOK, model.SimpleSuccessResponse(retMsg.Message))
}

// NodeSSHStatus represents the SSH readiness status of a single node
type NodeSSHStatus struct {
	ID        string `json:"id" example:"node-01"`          // Node ID
	Name      string `json:"name" example:"node-01"`        // Node Name
	PublicIP  string `json:"publicIP" example:"1.2.3.4"`    // Public IP address
	PrivateIP string `json:"privateIP" example:"10.0.1.10"` // Private IP address
	Username  string `json:"username" example:"cb-user"`    // SSH username
	Status    string `json:"status" example:"Running"`      // Node status (Running, Creating, etc.)
	SSHReady  bool   `json:"sshReady" example:"true"`       // Whether SSH port is accessible
	SSHPort   int    `json:"sshPort" example:"22"`          // SSH port number
	Error     string `json:"error,omitempty" example:""`    // Error message if SSH check failed
}

// CheckSSHReadyResponse represents the response for SSH readiness check
type CheckSSHReadyResponse struct {
	Ready           bool            `json:"ready"`                     // Overall readiness (true if all nodes are ready)
	TotalNodes      int             `json:"totalNodes"`                // Total number of nodes in the infrastructure
	ReadyNodes      int             `json:"readyNodes"`                // Number of nodes that are SSH-ready
	NodeStatus      []NodeSSHStatus `json:"nodeStatus,omitempty"`      // Detailed status for each node (only included when option=detail)
	Message         string          `json:"message"`                   // Summary message
	CheckedAt       string          `json:"checkedAt"`                 // Timestamp when check was performed
	NextAllowedTime string          `json:"nextAllowedTime,omitempty"` // Next allowed check time (for rate limiting)
}

// CheckSSHReady godoc
// @ID CheckSSHReady
// @Summary Check SSH readiness for migrated infrastructure nodes
// @Description Check if all nodes in the migrated infrastructure are SSH-accessible.
// @Description
// @Description "Running" status doesn't mean cloud-init (SSH user setup) is done; timing varies by
// @Description CSP (IBM Cloud VPC: up to ~3 min in testing). Works for any CSP.
// @Description
// @Description **Rate Limiting**: To prevent SSH server abuse, this API can only be called
// @Description once every 3 minutes for the same infrastructure. If called too soon, it returns
// @Description HTTP 429 with the time to wait before the next check is allowed.
// @Description
// @Description **Check Method**: This API runs a lightweight command on each node via Tumblebug's
// @Description remote command API (not a direct connection from CM-Beetle), so it reuses Tumblebug's
// @Description existing SSH/bastion setup for the node's CSP.
// @Description
// @Description **Response Options**:
// @Description - Default (no option): Returns summary information (ready, totalNodes, readyNodes, message)
// @Description - option=detail: Returns summary + detailed node status array (nodeStatus) for troubleshooting (per-node SSH readiness)
// @Tags [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param infraId path string true "Migrated Cloud Infrastructure ID"
// @Param option query string false "Response format (detail: include per-node SSH readiness information)" Enums(detail) default()
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[CheckSSHReadyResponse] "SSH readiness check completed"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Infrastructure not found"
// @Failure 429 {object} model.ApiResponse[CheckSSHReadyResponse] "Rate limited - check too soon after previous check"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /migration/ns/{nsId}/infra/{infraId}/ssh-ready [get]
func CheckSSHReady(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, the namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	infraId := c.Param("infraId")
	if infraId == "" {
		err := fmt.Errorf("invalid request, the cloud infrastructure ID (infraId: %s) is required", infraId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// Parse option parameter
	option := c.QueryParam("option")
	if option != "" && option != "detail" {
		err := fmt.Errorf("invalid option value: %s (supported: detail)", option)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// Check rate limiting
	limiter := migration.GetSSHCheckRateLimiter()
	allowed, retryAfter, err := limiter.CheckAllowed(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Msg("Rate limiter error")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	if !allowed {
		// Rate limited
		retryAfterSeconds := int(retryAfter.Seconds())
		c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", retryAfterSeconds))

		response := CheckSSHReadyResponse{
			Ready:           false,
			Message:         fmt.Sprintf("Rate limited - please wait %v before checking again", retryAfter.Round(time.Second)),
			NextAllowedTime: time.Now().Add(retryAfter).Format(time.RFC3339),
		}

		return c.JSON(http.StatusTooManyRequests, model.SuccessResponseWithMessage(
			response,
			fmt.Sprintf("Rate limited - SSH readiness check for this infrastructure can only be performed once every 3 minutes. Please wait %v.", retryAfter.Round(time.Second)),
		))
	}

	// [Process] Perform SSH readiness check
	log.Info().Msgf("Performing SSH readiness check (nsId: %s, infraId: %s)", nsId, infraId)

	// Check with 30 second timeout (single attempt)
	const checkTimeout = 30 * time.Second
	const checkInterval = 10 * time.Second

	nodeStatusList, err := migration.CheckSSHReadinessWithDetails(nsId, infraId, checkTimeout, checkInterval)

	// Convert internal status to API response format
	apiNodeStatus := make([]NodeSSHStatus, len(nodeStatusList))
	for i, nodeStatus := range nodeStatusList {
		apiNodeStatus[i] = NodeSSHStatus{
			ID:        nodeStatus.ID,
			Name:      nodeStatus.Name,
			PublicIP:  nodeStatus.PublicIP,
			PrivateIP: nodeStatus.PrivateIP,
			Username:  nodeStatus.Username,
			Status:    nodeStatus.Status,
			SSHReady:  nodeStatus.SSHReady,
			SSHPort:   nodeStatus.SSHPort,
			Error:     nodeStatus.Error,
		}
	}

	// Count ready nodes from status list
	totalNodes := len(apiNodeStatus)
	readyNodes := 0
	for _, nodeStatus := range apiNodeStatus {
		if nodeStatus.SSHReady {
			readyNodes++
		}
	}

	if err != nil {
		// SSH check failed or timed out
		log.Warn().Err(err).Msgf("SSH readiness check incomplete (nsId: %s, infraId: %s)", nsId, infraId)

		response := CheckSSHReadyResponse{
			Ready:      false,
			TotalNodes: totalNodes,
			ReadyNodes: readyNodes,
			Message:    fmt.Sprintf("Nodes not yet ready: %v", err),
			CheckedAt:  time.Now().Format(time.RFC3339),
		}

		// Include node status only if option=detail
		if option == "detail" {
			response.NodeStatus = apiNodeStatus
		}

		return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(
			response,
			"SSH check incomplete - Some nodes are not yet SSH-accessible. Cloud-init setup timing varies by CSP; retry later.",
		))
	}

	// [Output] SSH readiness confirmed
	response := CheckSSHReadyResponse{
		Ready:      true,
		TotalNodes: totalNodes,
		ReadyNodes: readyNodes,
		Message:    "All nodes are SSH-accessible",
		CheckedAt:  time.Now().Format(time.RFC3339),
	}

	// Include node status only if option=detail
	if option == "detail" {
		response.NodeStatus = apiNodeStatus
	}

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(
		response,
		"SSH ready - All nodes in the infrastructure are SSH-accessible.",
	))
}
