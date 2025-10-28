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

	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	// "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"

	// "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/infra"
	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog/log"
)

/*
 * VM Infrastructure Recommendation
 */

type RecommendVmInfraWithDefaultsRequest struct {
	DesiredCspAndRegionPair cloudmodel.CloudProperty `json:"desiredCspAndRegionPair"`
	OnpremiseInfraModel     onpremmodel.OnpremInfra
}

type RecommendVmInfraWithDefaultsResponse struct {
	cloudmodel.RecommendedVmInfraDynamicList
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
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendVmInfraWithDefaultsRequest true "Specify the your infrastructure to be migrated"
// @Param desiredCsp query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendVmInfraWithDefaultsResponse "The result of recommended infrastructure"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/mciWithDefaults [post]
func RecommendVMInfraWithDefaults(c echo.Context) error {

	// [Input]
	desiredCsp := c.QueryParam("desiredCsp")
	desiredRegion := c.QueryParam("desiredRegion")

	reqt := &RecommendVmInfraWithDefaultsRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Warn().Err(err).Msg("failed to bind a request body")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}
	log.Trace().Msgf("reqt: %v\n", reqt)

	if reqt.DesiredCspAndRegionPair.Csp == "" && desiredCsp == "" {
		err := fmt.Errorf("invalid request: 'desiredCsp' is required")
		log.Warn().Err(err).Msg("invalid request: 'desiredCsp' is required")
		resp := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}
	if reqt.DesiredCspAndRegionPair.Region == "" && desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
		log.Warn().Err(err).Msg("invalid request: 'desiredRegion' is required")
		resp := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
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
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Process]
	recommendedInfraInfoList, err := recommendation.RecommendVmInfraWithDefaults(csp, region, sourceInfra)
	// recommendedInfraInfoList.TargetInfra.Name = "mmci01"

	// [Ouput]
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	return c.JSON(http.StatusOK, recommendedInfraInfoList)
}

type RecommendVmInfraRequest struct {
	DesiredCspAndRegionPair cloudmodel.CloudProperty `json:"desiredCspAndRegionPair"`
	OnpremiseInfraModel     onpremmodel.OnpremInfra
}

type RecommendVmInfraResponse struct {
	cloudmodel.RecommendedVmInfra
}

// RecommendVMInfra godoc
// @ID RecommendVMInfra
// @Summary Recommend an appropriate VM infrastructure (i.e., MCI, multi-cloud infrastructure) for cloud migration
// @Description Recommend an appropriate VM infrastructure (i.e., MCI, multi-cloud infrastructure) for cloud migration
// @Description
// @Description [Note] `desiredCsp` and `desiredRegion` are required.
// @Description - `desiredCsp` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredCsp and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendVmInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredCsp query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendVmInfraResponse "The result of recommended infrastructure"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/mci [post]
func RecommendVMInfra(c echo.Context) error {

	// [Input]
	desiredCsp := c.QueryParam("desiredCsp")
	desiredRegion := c.QueryParam("desiredRegion")

	reqt := &RecommendVmInfraRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Warn().Err(err).Msg("failed to bind a request body")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}
	log.Trace().Msgf("reqt: %v\n", reqt)

	if reqt.DesiredCspAndRegionPair.Csp == "" && desiredCsp == "" {
		err := fmt.Errorf("invalid request: 'desiredCsp' is required")
		log.Warn().Err(err).Msg("invalid request: 'desiredCsp' is required")
		resp := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}
	if reqt.DesiredCspAndRegionPair.Region == "" && desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
		log.Warn().Err(err).Msg("invalid request: 'desiredRegion' is required")
		resp := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
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
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Process]
	recommendedInfra, err := recommendation.RecommendVmInfra(csp, region, sourceInfra)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	// [Ouput]
	//

	return c.JSON(http.StatusOK, recommendedInfra)
}

/*
 * K8s Cluster and Node Group Recommendation
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
// @Param K8sRequest body recommendation.KubernetesInfoList true "Source cluster information from honeybee"
// @Param desiredProvider query string true "Provider (e.g., aws)" Enums(aws)
// @Param desiredRegion query string true "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID"
// @Success 200 {object} tbmodel.K8sClusterDynamicReq "K8s control plane recommendation (ready for cb-tumblebug API)"
// @Failure 400 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/k8sControlPlane [post]
func RecommendK8sControlPlane(c echo.Context) error {
	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	if desiredProvider == "" || desiredRegion == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: "'desiredProvider' and 'desiredRegion' query parameters are required"})
	}

	reqt := &recommendation.KubernetesInfoList{}
	if err := c.Bind(reqt); err != nil {
		log.Error().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	if len(reqt.Servers) == 0 {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: "at least one server information is required"})
	}

	ok, err := recommendation.IsValidCspAndRegion(desiredProvider, desiredRegion)
	if !ok {
		log.Error().Err(err).Msg("invalid provider or region")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	k8sInfoList := recommendation.KubernetesInfoList{
		Servers: reqt.Servers,
	}

	result, err := recommendation.RecommendK8sControlPlane(desiredProvider, desiredRegion, k8sInfoList)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend K8s control plane")
		return c.JSON(http.StatusInternalServerError, common.SimpleMsg{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// RecommendK8sNodeGroup godoc
// @ID RecommendK8sNodeGroup
// @Summary Recommend K8s worker node group configuration
// @Description Get recommendation for K8s worker node group based on honeybee source cluster data
// @Description Returns configuration that can be directly used with cb-tumblebug k8sNodeGroupDynamic API
// @Tags [Recommendation] K8s Cluster (prototype)
// @Accept  json
// @Produce  json
// @Param K8sRequest body recommendation.KubernetesInfoList true "Source cluster information from honeybee"
// @Param desiredProvider query string true "Provider (e.g., aws)" Enums(aws)
// @Param desiredRegion query string true "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID"
// @Success 200 {object} tbmodel.K8sNodeGroupReq "K8s worker node group recommendation (ready for cb-tumblebug API)"
// @Failure 400 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/k8sNodeGroup [post]
func RecommendK8sNodeGroup(c echo.Context) error {
	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	if desiredProvider == "" || desiredRegion == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: "'desiredProvider' and 'desiredRegion' query parameters are required"})
	}

	reqt := &recommendation.KubernetesInfoList{}
	if err := c.Bind(reqt); err != nil {
		log.Error().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	if len(reqt.Servers) == 0 {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: "at least one server information is required"})
	}

	ok, err := recommendation.IsValidCspAndRegion(desiredProvider, desiredRegion)
	if !ok {
		log.Error().Err(err).Msg("invalid provider or region")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	k8sInfoList := recommendation.KubernetesInfoList{
		Servers: reqt.Servers,
	}

	result, err := recommendation.RecommendK8sNodeGroup(desiredProvider, desiredRegion, k8sInfoList)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend K8s node group")
		return c.JSON(http.StatusInternalServerError, common.SimpleMsg{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

/*
 * Container Infrastructure Recommendation (Legacy - will be deprecated)
 */
type RecommendInfraRequest struct {
	DesiredProvider string                      `json:"desiredProvider" example:"aws"`
	DesiredRegion   string                      `json:"desiredRegion" example:"ap-northeast-2"`
	Servers         []recommendation.Kubernetes `json:"servers"`
}

// recommendation.KubernetesInfoList is defined in pkg/core/recommendation/container-infra.go

type RecommendInfraResponse struct {
	recommendation.RecommendedInfraInfo
}

// RecommendContainerInfra godoc
// @ID RecommendContainerInfra
// @Summary (Deprecated) Recommend an appropriate container infrastructure for cloud migration
// @Description [DEPRECATED] This endpoint is deprecated. Use /recommendation/k8sCluster and /recommendation/k8sNodeGroup instead.
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendInfraRequest true "Specify the source container infrastructure"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID"
// @Success 200 {object} common.SimpleMsg "Deprecated endpoint notice"
// @Failure 400 {object} common.SimpleMsg
// @Router /recommendation/containerInfra [post]
// @Deprecated
func RecommendContainerInfra(c echo.Context) error {
	return c.JSON(http.StatusGone, common.SimpleMsg{
		Message: "This endpoint is deprecated. Please use /recommendation/k8sCluster for control plane and /recommendation/k8sNodeGroup for worker nodes.",
	})
}
