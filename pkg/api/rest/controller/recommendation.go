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

/*
 * K8s Infrastructure Recommendation With Defaults (K8sClusterDynamic path)
 */

// RecommendK8sInfraWithDefaultsRequest is the request body for RecommendK8sInfraWithDefaults.
type RecommendK8sInfraWithDefaultsRequest struct {
	DesiredCspAndRegionPair cloudmodel.CloudProperty `json:"desiredCspAndRegionPair"`
	onpremmodel.OnpremInfra
}

// RecommendK8sInfraWithDefaults godoc
// @ID RecommendK8sInfraWithDefaults
// @Summary Recommend K8s infrastructure with defaults for cloud migration
// @Description Recommend a target managed K8s cluster configuration based on source on-premise K8s infra.
// @Description Returns K8sClusterDynamicReq (+ additional node groups) ready for cb-tumblebug k8sClusterDynamic API.
// @Description TB auto-manages vNet/SG/SSHKey — no separate resource recommendation needed.
// @Tags [Recommendation] K8s Infrastructure
// @Accept  json
// @Produce  json
// @Param UserK8sInfra body RecommendK8sInfraWithDefaultsRequest true "Source on-premise K8s infrastructure"
// @Param desiredCsp query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[cloudmodel.RecommendedK8sCluster] "The recommended K8s cluster configuration"
// @Failure 400 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /recommendation/k8sInfraWithDefaults [post]
func RecommendK8sInfraWithDefaults(c echo.Context) error {

	desiredCsp := c.QueryParam("desiredCsp")
	desiredRegion := c.QueryParam("desiredRegion")

	reqt := &RecommendK8sInfraWithDefaultsRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	csp := reqt.DesiredCspAndRegionPair.Csp
	if csp == "" {
		csp = desiredCsp
	}
	region := reqt.DesiredCspAndRegionPair.Region
	if region == "" {
		region = desiredRegion
	}

	if csp == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Provider required"))
	}
	if region == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Region required"))
	}
	if len(reqt.Nodes) == 0 {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one source node required"))
	}
	if reqt.K8sCluster == nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("K8s cluster info required (k8sCluster field)"))
	}

	ok, err := recommendation.IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msg("invalid provider or region")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid provider or region"))
	}

	result, err := recommendation.RecommendK8sInfraWithDefaults(csp, region, reqt.OnpremInfra)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend K8s infrastructure")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("K8s infrastructure recommendation failed"))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}
