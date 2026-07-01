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

// Package common is to handle REST API for common funcitonalities
package controller

import (
	"fmt"
	"net/http"
	"strconv"

	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	// "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"

	// "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/infra"
	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog/log"
)

/*
 * VM Infrastructure Recommendation
 */

type RecommendInfraWithDefaultsRequest struct {
	DesiredCspAndRegionPair cloudmodel.CloudProperty `json:"desiredCspAndRegionPair"`
	OnpremiseInfraModel     onpremmodel.OnpremInfra
}

type RecommendInfraWithDefaultsResponse struct {
	cloudmodel.RecommendedInfraDynamicList
}

// RecommendVMInfraWithDefaults godoc
// @ID RecommendVMInfraWithDefaults
// @Summary (To be updated) Recommend an appropriate VM infrastructure (i.e., MCI, multi-cloud infrastructure) with defaults for cloud migration
// @Description Recommend an appropriate VM infrastructure (i.e., MCI, multi-cloud infrastructure) with defaults for cloud migration
// @Description
// @Description [Note] `desiredCsp` and `desiredRegion` are required.
// @Description - `desiredCsp` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredCsp and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Description
// @Description **[Response Field: `nodeGroups[].cspImageName`]** Set only when the spec-image review resolved a newer image than the DB cache.
// @Description - **Non-empty**: TumbleBug sends this to Spider directly, bypassing the per-VM image DB lookup (prevents stale image failures, e.g., Alibaba alibase images).
// @Description - **Empty**: TumbleBug uses `imageId` for the standard DB lookup path.
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendInfraWithDefaultsRequest true "Specify the source infrastructure to be migrated"
// @Param desiredCsp query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[RecommendInfraWithDefaultsResponse] "The result of recommended infrastructure"
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /recommendation/infraWithDefaults [post]
func RecommendVMInfraWithDefaults(c echo.Context) error {

	// [Input]
	desiredCsp := c.QueryParam("desiredCsp")
	desiredRegion := c.QueryParam("desiredRegion")

	reqt := &RecommendInfraWithDefaultsRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Warn().Err(err).Msg("failed to bind a request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}
	log.Trace().Msgf("reqt: %v\n", reqt)

	if reqt.DesiredCspAndRegionPair.Csp == "" && desiredCsp == "" {
		log.Warn().Msg("desiredCsp is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Provider required"))
	}
	if reqt.DesiredCspAndRegionPair.Region == "" && desiredRegion == "" {
		log.Warn().Msg("desiredRegion is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Region required"))
	}

	csp := reqt.DesiredCspAndRegionPair.Csp
	if csp == "" {
		csp = desiredCsp
	}
	region := reqt.DesiredCspAndRegionPair.Region
	if region == "" {
		region = desiredRegion
	}
	sourceInfra := reqt.OnpremiseInfraModel

	ok, err := recommendation.IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msg("failed to validate CSP and region")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid provider or region"))
	}

	// [Process]
	recommendedInfraInfoList, err := recommendation.RecommendVmInfraWithDefaults(csp, region, sourceInfra)
	// recommendedInfraInfoList.TargetInfra.Name = "mci101"

	// [Ouput]
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration")
		return c.JSON(http.StatusNotFound, model.SimpleErrorResponse("Recommendation failed"))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(recommendedInfraInfoList))
}

type RecommendInfraRequest struct {
	DesiredCspAndRegionPair cloudmodel.CloudProperty `json:"desiredCspAndRegionPair"`
	OnpremiseInfraModel     onpremmodel.OnpremInfra
}

type RecommendInfraResponse struct {
	cloudmodel.RecommendedInfra
}

// RecommendVmInfraCandidates godoc
// @ID RecommendVmInfraCandidates
// @Summary Recommend multiple VM infrastructure candidates for cloud migration
// @Description Recommend best-effort VM infrastructure (MCI) candidates for migrating on-premise workloads to cloud environments.
// @Description
// @Description - See overview and examples on https://github.com/cloud-barista/cm-beetle/discussions/256
// @Description
// @Description **[Required Parameters: `desiredCsp`, `desiredRegion`]** The desired cloud service provider and region for the recommended infrastructure.
// @Description - if **desiredCsp** and **desiredRegion** are set on request body, the values in the query parameter will be ignored.
// @Description
// @Description **[Optional Parameters: `limit`]** Maximum number of recommended infrastructures to return (default: 3)
// @Description
// @Description **[Optional Parameters: `minMatchRate`]** Minimum match rate threshold for highly-matched classification (default: 90.0, range: 0-100)
// @Description
// @Description **[Response Field: `status`]** Candidate status based on the match rate threshold
// @Description - **highly-matched**: Candidates meet or exceed the match rate threshold
// @Description - **partially-matched**: Valid candidates below the match rate threshold
// @Description
// @Description **[Response Field: `description`]** Summary containing Candidate ID, status, match rate statistics (Min/Max/Avg), and VM counts
// @Description - Example: "Candidate #1 | partially-matched | Overall Match Rate: Min=88.9% Max=100.0% Avg=98.7% | VMs: 3 total, 2 matched, 1 acceptable"
// @Description
// @Description **[Response Field: `nodeGroups[].cspImageName`]** Set only when the spec-image review resolved a newer image than the DB cache.
// @Description - **Non-empty**: TumbleBug sends this to Spider directly, bypassing the per-VM image DB lookup (prevents stale image failures, e.g., Alibaba alibase images).
// @Description - **Empty**: TumbleBug uses `imageId` for the standard DB lookup path.
// @Description - Pass the recommendation response as-is to the migration API to ensure the resolved image is used.
// @Description
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendInfraRequest true "Specify the source infrastructure to be migrated"
// @Param desiredCsp query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param limit query int false "Limit (default: 3) the number of recommended infrastructures"
// @Param minMatchRate query number false "Minimum match rate for highly-matched classification (default: 90.0, range: 0-100)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[[]cloudmodel.RecommendedInfra] "Successfully recommended infrastructure candidates"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Router /recommendation/infra [post]
func RecommendVmInfraCandidates(c echo.Context) error {

	// [Input]
	desiredCsp := c.QueryParam("desiredCsp")
	desiredRegion := c.QueryParam("desiredRegion")
	limitStr := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 3 // default value
	}

	// Parse minMatchRate parameter (default: 90.0)
	minMatchRateStr := c.QueryParam("minMatchRate")
	minMatchRate := 90.0 // default value
	if minMatchRateStr != "" {
		parsedRate, err := strconv.ParseFloat(minMatchRateStr, 64)
		if err != nil {
			log.Warn().Err(err).Msgf("invalid minMatchRate value: %s, using default 90.0", minMatchRateStr)
		} else if parsedRate < 0 || parsedRate > 100 {
			log.Warn().Msgf("minMatchRate out of range [0-100]: %.1f, using default 90.0", parsedRate)
		} else {
			minMatchRate = parsedRate
		}
	}

	reqt := &RecommendInfraRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Warn().Err(err).Msg("failed to bind a request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}
	log.Trace().Msgf("reqt: %v\n", reqt)

	if reqt.DesiredCspAndRegionPair.Csp == "" && desiredCsp == "" {
		log.Warn().Msg("desiredCsp is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Provider required"))
	}
	if reqt.DesiredCspAndRegionPair.Region == "" && desiredRegion == "" {
		log.Warn().Msg("desiredRegion is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Region required"))
	}

	csp := reqt.DesiredCspAndRegionPair.Csp
	if csp == "" {
		csp = desiredCsp
	}
	region := reqt.DesiredCspAndRegionPair.Region
	if region == "" {
		region = desiredRegion
	}
	sourceInfra := reqt.OnpremiseInfraModel

	ok, err := recommendation.IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msg("failed to validate CSP and region")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid provider or region"))
	}

	// [Process]
	recommendedInfraCandidates, err := recommendation.RecommendVmInfraCandidates(csp, region, sourceInfra, limit, minMatchRate)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend multiple candidates of appropriate multi-cloud infrastructure (MCI) for cloud migration")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Recommendation failed"))
	}

	// [Output]
	// Returns base names only. NameSeed is applied at migration time via query param on the migration API.
	return c.JSON(http.StatusOK, model.SuccessListResponse(recommendedInfraCandidates))
}

/*
 * NLB-aware Infrastructure Recommendation
 */

// RecommendInfraWithNlbRequest is the request body for POST /recommendation/infraWithNlb.
type RecommendInfraWithNlbRequest struct {
	DesiredCsp    string                  `json:"desiredCsp"`    // Target CSP (e.g., "aws")
	DesiredRegion string                  `json:"desiredRegion"` // Target region (e.g., "ap-northeast-2")
	SourceInfra   onpremmodel.OnpremInfra `json:"sourceInfra"   validate:"required"`
}

// RecommendInfraWithNlbCandidates godoc
// @ID RecommendInfraWithNlbCandidates
// @Summary (Preview) Recommend infrastructure candidates with NLB for cloud migration
// @Description Perform NLB-aware infrastructure recommendation and return multiple Pareto-optimal candidates.
// @Description
// @Description The recommendation engine:
// @Description 1. Correlates NLB backend server IPs with source Node IPs
// @Description 2. Normalizes backend ports via majority vote when ports differ
// @Description 3. Assigns NLB-related nodes to shared NodeGroups (N:1), unrelated nodes to individual NodeGroups (1:1)
// @Description 4. Finds ranked compatible spec-image pairs per NodeGroup (representative node for NLB groups)
// @Description 5. Generates up to `limit` candidates — candidate i uses the i-th ranked pair per NodeGroup
// @Description 6. Maps source NLB configuration to target cloud NLB model (same for all candidates)
// @Description
// @Description [Note] `sourceInfra.nlbs` must be populated (HAProxy frontend-backend pairs from cm-honeybee).
// @Description
// @Description [Note] The returned `targetInfra.nodeGroups[].name` values are referenced by `targetNlbList[].targetGroup.nodeGroupId`.
// @Description Use the same NodeGroup IDs when calling POST /migration/infra so that the NLB migration can reference them immediately.
// @Description
// @Description ---
// @Description ## CSP-Specific NLB Notes
// @Description
// @Description AWS:
// @Description - Port translation supported (e.g., listener 9999 → backend 8086).
// @Description - DNS endpoint; allow ~5 min for propagation after creation.
// @Description - [Auto] SG rule for backend port opened from 0.0.0.0/0.
// @Description
// @Description Azure:
// @Description - Port translation supported. DNS + static IP endpoint.
// @Description - [Auto] Health check timeout omitted (not supported by Azure).
// @Description
// @Description GCP:
// @Description - Port translation NOT supported; traffic arrives at backend VMs on the listener port.
// @Description - [Auto] Listener port is forced equal to the backend port — clients must connect on the application port (e.g., 8086, not 9999).
// @Description - IP-only endpoint (no DNS name).
// @Description
// @Description IBM:
// @Description - Port translation supported.
// @Description - Listener address is assigned asynchronously; re-query if the address is empty after migration.
// @Description - [Auto] Health check timeout forced strictly less than the interval.
// @Description ---
// @Tags [Recommendation] Infrastructure
// @Accept json
// @Produce json
// @Param desiredCsp query string false "Target CSP (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Target region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param limit query int false "Maximum number of candidates to return" default(5)
// @Param minMatchRate query number false "Minimum match rate (0-100) for highly-matched classification" default(90.0)
// @Param request body RecommendInfraWithNlbRequest true "Source infra including NLBs (from cm-honeybee)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided)"
// @Success 200 {object} model.ApiResponse[[]cloudmodel.RecommendedInfra] "NLB-aware recommendation candidates"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /recommendation/infraWithNlb [post]
func RecommendInfraWithNlbCandidates(c echo.Context) error {

	// [Input]
	var req RecommendInfraWithNlbRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind RecommendInfraWithNlbCandidates request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// Query params take priority over request body values.
	if qp := c.QueryParam("desiredCsp"); qp != "" {
		req.DesiredCsp = qp
	}
	if qp := c.QueryParam("desiredRegion"); qp != "" {
		req.DesiredRegion = qp
	}

	limit := 5
	if qp := c.QueryParam("limit"); qp != "" {
		n, err := strconv.Atoi(qp)
		if err != nil || n < 1 {
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("limit must be a positive integer"))
		}
		limit = n
	}

	minMatchRate := 90.0
	if qp := c.QueryParam("minMatchRate"); qp != "" {
		r, err := strconv.ParseFloat(qp, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("minMatchRate must be a number (0-100)"))
		}
		minMatchRate = r
	}

	if len(req.SourceInfra.Nodes) == 0 {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("sourceInfra.nodes is required"))
	}
	if len(req.SourceInfra.NLBs) == 0 {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(
			"sourceInfra.nlbs is required for infraWithNlb; use /recommendation/infra for NLB-free recommendation"))
	}

	log.Info().
		Str("desiredCsp", req.DesiredCsp).
		Str("desiredRegion", req.DesiredRegion).
		Int("nodes", len(req.SourceInfra.Nodes)).
		Int("nlbs", len(req.SourceInfra.NLBs)).
		Int("limit", limit).
		Float64("minMatchRate", minMatchRate).
		Msg("Processing infraWithNlb recommendation request")

	// [Process]
	candidates, err := recommendation.RecommendInfraWithNlbCandidates(
		req.DesiredCsp, req.DesiredRegion, req.SourceInfra, limit, minMatchRate,
	)
	if err != nil {
		log.Error().Err(err).Msg("infraWithNlb recommendation failed")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(err.Error()))
	}

	// [Output]
	nlbCount := 0
	ngCount := 0
	if len(candidates) > 0 {
		nlbCount = len(candidates[0].TargetNlbList)
		ngCount = len(candidates[0].TargetInfra.NodeGroups)
	}
	successMsg := fmt.Sprintf(
		"%d candidate(s) recommended — each with %d NLB(s) and %d NodeGroup(s)",
		len(candidates), nlbCount, ngCount)

	log.Info().
		Int("candidates", len(candidates)).
		Int("nodeGroups", ngCount).
		Int("nlbs", nlbCount).
		Msg("infraWithNlb recommendation completed")

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(candidates, successMsg))
}

/*
 * K8s Cluster Control Plane and Node Group Recommendation
 */

type RecommendK8sClusterResponse struct {
	tbmodel.K8sClusterDynamicReq
}

// RecommendK8sControlPlane godoc
// @ID RecommendK8sControlPlane
// @Summary Recommend K8s control plane configuration
// @Description Get recommendation for K8s control plane based on honeybee source cluster data
// @Description Returns configuration that can be directly used with cb-tumblebug k8sClusterDynamic API
// @Tags [Recommendation] K8s Cluster (prototype)
// @Accept  json
// @Produce  json
// @Param UserK8sInfra body recommendation.KubernetesInfoList true "Source cluster information from honeybee"
// @Param desiredProvider query string true "Provider (e.g., aws)" Enums(aws)
// @Param desiredRegion query string true "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[tbmodel.K8sClusterDynamicReq] "K8s control plane recommendation (ready for cb-tumblebug API)"
// @Failure 400 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /recommendation/k8sControlPlane [post]
func RecommendK8sControlPlane(c echo.Context) error {
	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	if desiredProvider == "" || desiredRegion == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("'desiredProvider' and 'desiredRegion' query parameters are required"))
	}

	reqt := &recommendation.KubernetesInfoList{}
	if err := c.Bind(reqt); err != nil {
		log.Error().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if len(reqt.Servers) == 0 {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one cluster required"))
	}

	ok, err := recommendation.IsValidCspAndRegion(desiredProvider, desiredRegion)
	if !ok {
		log.Error().Err(err).Msg("invalid provider or region")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid provider or region"))
	}

	k8sInfoList := recommendation.KubernetesInfoList{
		Servers: reqt.Servers,
	}

	result, err := recommendation.RecommendK8sControlPlane(desiredProvider, desiredRegion, k8sInfoList)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend K8s control plane")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("K8s control plane recommendation failed"))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}

// RecommendK8sNodeGroup godoc
// @ID RecommendK8sNodeGroup
// @Summary Recommend K8s worker node group configuration
// @Description Get recommendation for K8s worker node group based on honeybee source cluster data
// @Description Returns configuration that can be directly used with cb-tumblebug k8sNodeGroupDynamic API
// @Tags [Recommendation] K8s Cluster (prototype)
// @Accept  json
// @Produce  json
// @Param UserK8sInfra body recommendation.KubernetesInfoList true "Source cluster information from honeybee"
// @Param desiredProvider query string true "Provider (e.g., aws)" Enums(aws)
// @Param desiredRegion query string true "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[tbmodel.K8sNodeGroupReq] "K8s worker node group recommendation (ready for cb-tumblebug API)"
// @Failure 400 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /recommendation/k8sNodeGroup [post]
func RecommendK8sNodeGroup(c echo.Context) error {
	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	if desiredProvider == "" || desiredRegion == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("'desiredProvider' and 'desiredRegion' query parameters are required"))
	}

	reqt := &recommendation.KubernetesInfoList{}
	if err := c.Bind(reqt); err != nil {
		log.Error().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if len(reqt.Servers) == 0 {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one cluster required"))
	}

	ok, err := recommendation.IsValidCspAndRegion(desiredProvider, desiredRegion)
	if !ok {
		log.Error().Err(err).Msg("invalid provider or region")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid provider or region"))
	}

	k8sInfoList := recommendation.KubernetesInfoList{
		Servers: reqt.Servers,
	}

	result, err := recommendation.RecommendK8sNodeGroup(desiredProvider, desiredRegion, k8sInfoList)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend K8s node group")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("K8s node group recommendation failed"))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}
