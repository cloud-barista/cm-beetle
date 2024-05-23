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

	cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model/onprem/infra"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"

	// Black import (_) is for running a package's init() function without using its other contents.
	_ "github.com/cloud-barista/cm-beetle/pkg/logger"
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
	infra.Infra
}

type RecommendInfraResponse struct {
	cloudmodel.TbMcisDynamicReq
}

// RecommendInfra godoc
// @Summary Recommend an appropriate infrastructure for cloud migration
// @Description It recommends a cloud infrastructure most similar to the input. Infrastructure includes network, storage, compute, and so on.
// @Tags [Recommendation] Infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfrastructure body RecommendInfraRequest true "Specify network, disk, compute, security group, virtual machine, etc."
// @Success 200 {object} RecommendInfraResponse "Successfully recommended an appropriate infrastructure for cloud migration"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/infra [post]
func RecommendInfra(c echo.Context) error {

	// Input
	req := &RecommendInfraRequest{}
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("Failed to bind a request body")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusBadRequest, res)
	}

	log.Trace().Msgf("req: %v\n", req)
	log.Trace().Msgf("req.Infra.Compute: %v\n", req.Infra.Compute)
	// log.Trace().Msgf("req.Infra.Network: %v\n", req.Infra.Network)
	// log.Trace().Msgf("req.Infra.GPU: %v\n", req.Infra.GPU)

	// Process
	recommendedInfra, err := recommendation.Recommend(req.Infra)
	recommendedInfra.Name = "recomm-infra01"

	// Ouput
	if err != nil {
		log.Error().Err(err).Msg("Failed to recommend an appropriate infrastructure for cloud migration")
		res := common.SimpleMsg{Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	return c.JSON(http.StatusOK, recommendedInfra)
}
