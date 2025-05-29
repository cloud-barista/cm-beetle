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
	inframodel "github.com/cloud-barista/cm-model/infra/onprem"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog/log"
)

// type Infrastructure struct {
// 	Network        string
// 	Disk           string
// 	Compute        string
// 	SecurityGroup  string
// 	VirtualMachine string
// }

type RecommendVmInfraRequest struct {
	DesiredCspAndRegionPair recommendation.CspRegionPair `json:"desiredCspAndRegionPair"`
	inframodel.OnpremiseInfraModel
}

type RecommendVmInfraResponse struct {
	recommendation.RecommendedInfraInfo
}

// RecommendVMInfra godoc
// @ID RecommendVMInfra
// @Summary Recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration
// @Description Recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration
// @Description
// @Description [Note] `desiredCsp` and `desiredRegion` are required.
// @Description - `desiredCsp` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredCsp and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendVmInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredCsp query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendVmInfraResponse "The result of recommended infrastructure"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/mci [post]
func RecommendVMInfra(c echo.Context) error {

	desiredCsp := c.QueryParam("desiredCsp")
	desiredRegion := c.QueryParam("desiredRegion")

	// [Input]
	reqt := &RecommendVmInfraRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Error().Err(err).Msg("failed to bind a request body")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}
	log.Trace().Msgf("reqt: %v\n", reqt)

	if reqt.DesiredCspAndRegionPair.Csp == "" && desiredCsp == "" {
		err := fmt.Errorf("invalid request: 'desiredCsp' is required")
		resp := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}
	if reqt.DesiredCspAndRegionPair.Region == "" && desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
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
	sourceInfra := reqt.OnpremiseInfraModel.OnpremiseInfraModel

	// Replace "ncp" with "ncpvpc"
	// TODO: improve it when "ncp" and "ncpvpc" are updated.
	if csp == "ncp" {
		csp = csp + "vpc"
	}

	ok, err := recommendation.IsValidCspAndRegion(csp, region)
	if !ok {
		log.Error().Err(err).Msg("failed to validate CSP and region")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Process]
	recommendedInfraInfoList, err := recommendation.RecommendVmInfraDynamic(csp, region, sourceInfra)
	// recommendedInfraInfoList.TargetInfra.Name = "mmci01"

	// [Ouput]
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	return c.JSON(http.StatusOK, recommendedInfraInfoList)
}

type RecommendInfraRequest struct {
	DesiredProvider string `json:"desiredProvider" example:"aws"`
	DesiredRegion   string `json:"desiredRegion" example:"ap-northeast-2"`
	inframodel.OnpremiseInfraModel
}

type RecommendInfraResponse struct {
	recommendation.RecommendedInfraInfo
}

// RecommendContainerInfra godoc
// @ID RecommendContainerInfra
// @Summary Recommend an appropriate container infrastructure for cloud migration
// @Description Recommend an appropriate container infrastructure for container-based workloads
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Description - `desiredProvider` and `desiredRegion` can be set in the query parameter or the request body.
// @Description - If both are set, the values in the request body take precedence.
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendInfraRequest true "Specify the source container infrastructure"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendInfraResponse "The result of recommended container infrastructure"
// @Failure 400 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/containerInfra [post]
func RecommendContainerInfra(c echo.Context) error {

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	reqt := &RecommendInfraRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Error().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	if reqt.DesiredProvider == "" && desiredProvider == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: "'desiredProvider' is required"})
	}
	if reqt.DesiredRegion == "" && desiredRegion == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: "'desiredRegion' is required"})
	}

	provider := reqt.DesiredProvider
	if provider == "" {
		provider = desiredProvider
	}
	region := reqt.DesiredRegion
	if region == "" {
		region = desiredRegion
	}
	sourceInfra := reqt.OnpremiseInfraModel.OnpremiseInfraModel

	if provider == "ncp" {
		provider = provider + "vpc"
	}

	ok, err := recommendation.IsValidCspAndRegion(provider, region)
	if !ok {
		log.Error().Err(err).Msg("invalid provider or region")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	result, err := recommendation.RecommendContainer(provider, region, sourceInfra)
	if err != nil {
		log.Error().Err(err).Msg("failed to call RecommendContainer")
		return c.JSON(http.StatusInternalServerError, common.SimpleMsg{Message: "container recommendation failed"})
	}

	return c.JSON(http.StatusOK, result)
}
