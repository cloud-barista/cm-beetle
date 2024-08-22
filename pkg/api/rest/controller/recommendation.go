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

	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	// "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"
	"github.com/cloud-barista/cb-tumblebug/src/core/mci"
	"github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/infra"

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
	Servers []infra.Infra `json:"servers" validate:"required"`
}

type RecommendInfraResponse struct {
	mci.TbMciDynamicReq
}

// RecommendInfra godoc
// @Summary Recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration
// @Description Recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendInfraRequest true "Specify the your infrastructure to be migrated"
// @Success 200 {object} RecommendInfraResponse "The result of recommended infrastructure"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/mci [post]
func RecommendInfra(c echo.Context) error {

	// Input
	req := &RecommendInfraRequest{}
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("failed to bind a request body")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}

	log.Trace().Msgf("req: %v\n", req)

	// Process
	recommendedInfra, err := recommendation.Recommend(req.Servers)
	recommendedInfra.Name = "mmci01"

	// Ouput
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend an appropriate multi-cloud infrastructure (MCI) for cloud migration")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	return c.JSON(http.StatusOK, recommendedInfra)
}
