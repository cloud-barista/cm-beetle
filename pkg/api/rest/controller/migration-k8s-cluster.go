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

// Package controller has handlers for K8s cluster migration APIs.
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

// MigrateK8sClusterRequest is the request body for POST /migration/ns/{nsId}/k8sCluster.
// Use the RecommendedInfra returned by POST /recommendation/k8sCluster as-is.
type MigrateK8sClusterRequest struct {
	cloudmodel.RecommendedInfra
}

// MigrateK8sClusterResponse wraps the created K8s cluster info.
type MigrateK8sClusterResponse struct {
	tbmodel.K8sClusterInfo
}

// MigrateK8sCluster godoc
// @ID MigrateK8sCluster
// @Summary Migrate an on-premise K8s cluster to the target CSP
// @Description Migrate an on-premise Kubernetes cluster to a managed K8s service (e.g., AWS EKS).
// @Description
// @Description **Workflow**:
// @Description 1. Call POST /recommendation/k8sCluster to get a RecommendedK8sInfra.
// @Description 2. Pass the recommendation result as-is to this endpoint.
// @Description
// @Description **Internal steps**: VNet → SSH Key → Security Group → K8s Cluster → Node Groups
// @Description
// @Description By default this API runs synchronously (EKS takes ~10-20 min). Send header
// @Description `Prefer: respond-async` to run it asynchronously: receive 202 Accepted with a
// @Description reqId, then poll GET /request/{reqId} to check status.
// @Tags [Migration] K8s Cluster
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param k8sClusterInfo body MigrateK8sClusterRequest true "K8s cluster recommendation (from POST /recommendation/k8sCluster)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this migration asynchronously (RFC 7240)" Enums(respond-async)
// @Success 201 {object} model.ApiResponse[MigrateK8sClusterResponse] "K8s cluster migrated successfully"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Migration started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /migration/ns/{nsId}/k8sCluster [post]
func MigrateK8sCluster(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Namespace ID required"))
	}

	req := new(MigrateK8sClusterRequest)
	if err := c.Bind(req); err != nil {
		log.Warn().Err(err).Msg("failed to bind K8s migration request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if req.TargetK8sCluster.Name == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Cluster name required"))
	}

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (tbmodel.K8sClusterInfo, error) {
			return migration.CreateK8sCluster(nsId, &req.RecommendedInfra)
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
			"K8s cluster migration started. Use GET /request/{reqId} to check status."))
	}

	clusterInfo, err := migration.CreateK8sCluster(nsId, &req.RecommendedInfra)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Msg("failed to migrate K8s cluster")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SuccessResponse(clusterInfo))
}

// ListK8sClusters godoc
// @ID ListK8sClusters
// @Summary List all migrated K8s clusters
// @Description List all K8s clusters created in the namespace via cm-beetle migration.
// @Tags [Migration] K8s Cluster
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[[]tbmodel.K8sClusterInfo] "List of migrated K8s clusters"
// @Failure 500 {object} model.ApiResponse[any]
// @Router /migration/ns/{nsId}/k8sCluster [get]
func ListK8sClusters(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Namespace ID required"))
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
// @Tags [Migration] K8s Cluster
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param k8sClusterId path string true "K8s Cluster ID"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[MigrateK8sClusterResponse] "K8s cluster information"
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
// @Tags [Migration] K8s Cluster
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
