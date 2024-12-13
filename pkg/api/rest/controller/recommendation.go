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

type RecommendInfraRequest struct {
	DesiredProvider string `json:"desiredProvider" example:"aws"`
	DesiredRegion   string `json:"desiredRegion" example:"ap-northeast-2"`
	inframodel.OnpremiseInfraModel
}

type RecommendInfraResponse struct {
	recommendation.RecommendedInfraInfo
}

// RecommendInfra godoc
// @ID RecommendInfra
// @Summary Recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration
// @Description Recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Description - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredProvider and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendInfraResponse "The result of recommended infrastructure"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/mci [post]
func RecommendInfra(c echo.Context) error {

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	// [Input]
	reqt := &RecommendInfraRequest{}
	if err := c.Bind(reqt); err != nil {
		log.Error().Err(err).Msg("failed to bind a request body")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}

	log.Trace().Msgf("reqt: %v\n", reqt)

	if reqt.DesiredProvider == "" && desiredProvider == "" {
		err := fmt.Errorf("invalid request: 'desiredProvider' is required")
		resp := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
	}
	if reqt.DesiredRegion == "" && desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
		resp := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, resp)
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

	// Replace "ncp" with "ncpvpc"
	// TODO: improve it when "ncp" and "ncpvpc" are updated.
	if provider == "ncp" {
		provider = provider + "vpc"
	}

	ok, err := recommendation.IsValidProviderAndRegion(provider, region)
	if !ok {
		log.Error().Err(err).Msg("failed to validate provider and region")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}

	// [Process]
	recommendedInfraInfo, err := recommendation.Recommend(provider, region, sourceInfra)
	recommendedInfraInfo.TargetInfra.Name = "mmci01"

	// [Ouput]
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	return c.JSON(http.StatusOK, recommendedInfraInfo)
}
