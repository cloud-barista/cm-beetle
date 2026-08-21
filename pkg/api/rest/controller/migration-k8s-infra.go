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

// Package controller has handlers for K8s infra migration APIs.
package controller

import (
	"fmt"
	"net/http"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// MigrateK8sInfraRequest is the request body for POST /migration/ns/{nsId}/k8sCluster.
// Use the RecommendedInfra returned by POST /recommendation/k8sCluster as-is.
type MigrateK8sInfraRequest struct {
	cloudmodel.RecommendedInfra
}

// MigrateK8sInfraResponse wraps the created K8s cluster info.
type MigrateK8sInfraResponse struct {
	tbmodel.K8sClusterInfo
}

// MigrateK8sInfra godoc
// @ID MigrateK8sInfra
// @Summary Migrate an on-premise K8s cluster to a K8s infra on the target CSP (sync by default; async via Prefer: respond-async)
// @Description Migrate an on-premise Kubernetes cluster to a managed K8s service (e.g., AWS EKS).
// @Description
// @Description **Workflow**:
// @Description 1. Call POST /recommendation/k8sCluster to get a RecommendK8sInfraResponse.
// @Description 2. Pass the recommendation result as-is to this endpoint.
// @Description
// @Description **Internal steps**: VNet → SSH Key → Security Group → K8s Cluster → Node Groups
// @Description
// @Description **Long-running** — This API waits for the cluster to become Active, up to **40 min**
// @Description (typically **~10-20 min** on EKS), then adds the node groups, so it can take longer still.
// @Description
// @Description By default it runs synchronously, so configure your own HTTP client to wait that long —
// @Description most client defaults cut the connection first. Send header `Prefer: respond-async` to run it
// @Description asynchronously instead: receive 202 Accepted with a reqId, then poll GET /request/{reqId}
// @Description (status flow: Handling → Success / Error). Only the "respond-async" token is recognized.
// @Description
// @Description If a synchronous connection drops the migration keeps running on the server; poll
// @Description GET /migration/ns/{nsId}/k8sCluster/{k8sClusterId} with the cluster name you supplied.
// @Tags [Migration] K8s Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param nameSeed query string false "Optional prefix for all resource names (e.g., 'blue' -> 'blue-k8s-vpc', 'blue-on-prem-k8s-cluster'). Node group names are left unprefixed because they are scoped inside the cluster and bound by per-CSP length limits. Applied at migration time."
// @Param k8sInfraInfo body MigrateK8sInfraRequest true "K8s infra recommendation (from POST /recommendation/k8sCluster)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this migration asynchronously (RFC 7240)" Enums(respond-async)
// @Success 201 {object} model.ApiResponse[MigrateK8sInfraResponse] "K8s infra migrated successfully"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Migration started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /migration/ns/{nsId}/k8sCluster [post]
func MigrateK8sInfra(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Namespace ID required"))
	}

	req := new(MigrateK8sInfraRequest)
	if err := c.Bind(req); err != nil {
		log.Warn().Err(err).Msg("failed to bind K8s migration request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if req.TargetK8sCluster.Name == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Cluster name required"))
	}

	// Late-binding names let several migrations share one namespace: without a seed every run
	// would reuse the same fixed resource names and collide with the previous one.
	nameSeed := c.QueryParam("nameSeed")
	if ok, detail := common.IsValidNameSeed(nameSeed); !ok {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid nameSeed: "+detail))
	}
	infraToMigrate := common.ApplyNameSeed(req.RecommendedInfra, nameSeed)

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (tbmodel.K8sClusterInfo, error) {
			return migration.CreateK8sInfra(nsId, &infraToMigrate)
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
			"K8s infra migration started. Use GET /request/{reqId} to check status."))
	}

	clusterInfo, err := migration.CreateK8sInfra(nsId, &infraToMigrate)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Msg("failed to migrate K8s infra")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SuccessResponse(clusterInfo))
}

// ListK8sClusters godoc
// @ID ListK8sClusters
// @Summary List all migrated K8s clusters
// @Description List all K8s clusters created in the namespace via cm-beetle migration.
// @Description
// @Description Use `option=id` to get only the cluster IDs. The full listing refreshes every
// @Description cluster's state through the CSP, so its cost grows with the cluster count; the
// @Description ID listing reads stored metadata only and stays cheap.
// @Tags [Migration] K8s Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param option query string false "Option for listing the migrated K8s clusters" Enums(id)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[[]tbmodel.K8sClusterInfo] "List of migrated K8s clusters (or IDs when option=id)"
// @Failure 400 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/k8sCluster [get]
func ListK8sClusters(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Namespace ID required"))
	}

	option := c.QueryParam("option")
	if option != "" && option != "id" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(
			fmt.Sprintf("Invalid option: %s (only 'id' is supported)", option)))
	}

	if option == "id" {
		idList, err := migration.ListK8sClusterIds(nsId)
		if err != nil {
			log.Error().Err(err).Str("nsId", nsId).Msg("failed to list K8s cluster IDs")
			return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, model.SuccessResponse(idList))
	}

	clusters, err := migration.ListK8sClusters(nsId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Msg("failed to list K8s clusters")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(clusters))
}

// GetK8sCluster godoc
// @ID GetK8sCluster
// @Summary Get a migrated K8s cluster
// @Description Retrieve information about a specific K8s cluster created via cm-beetle migration.
// @Tags [Migration] K8s Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param k8sClusterId path string true "K8s Cluster ID"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[MigrateK8sInfraResponse] "K8s cluster information"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/k8sCluster/{k8sClusterId} [get]
func GetK8sCluster(c echo.Context) error {
	nsId := c.Param("nsId")
	k8sClusterId := c.Param("k8sClusterId")
	if nsId == "" || k8sClusterId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(
			fmt.Sprintf("Namespace ID and K8s cluster ID required (nsId=%s, k8sClusterId=%s)", nsId, k8sClusterId)))
	}

	clusterInfo, err := migration.GetK8sCluster(nsId, k8sClusterId)
	if err != nil {
		log.Error().Err(err).Str("k8sClusterId", k8sClusterId).Msg("failed to get K8s cluster")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(clusterInfo))
}

// DeleteK8sCluster godoc
// @ID DeleteK8sCluster
// @Summary Delete a migrated K8s cluster
// @Description Delete a K8s cluster created via cm-beetle migration. This removes the cluster and all its node groups.
// @Tags [Migration] K8s Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param k8sClusterId path string true "K8s Cluster ID"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[any] "K8s cluster deleted"
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/k8sCluster/{k8sClusterId} [delete]
func DeleteK8sCluster(c echo.Context) error {
	nsId := c.Param("nsId")
	k8sClusterId := c.Param("k8sClusterId")
	if nsId == "" || k8sClusterId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(
			fmt.Sprintf("Namespace ID and K8s cluster ID required (nsId=%s, k8sClusterId=%s)", nsId, k8sClusterId)))
	}

	if err := migration.DeleteK8sCluster(nsId, k8sClusterId); err != nil {
		log.Error().Err(err).Str("k8sClusterId", k8sClusterId).Msg("failed to delete K8s cluster")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(map[string]string{"message": fmt.Sprintf("K8s cluster %s deleted", k8sClusterId)}, fmt.Sprintf("K8s cluster %s deleted", k8sClusterId)))
}
