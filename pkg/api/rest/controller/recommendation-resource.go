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

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// type RecommendVNetRequest struct {
// 	DesiredProvider string   `json:"desiredProvider" example:"aws"`
// 	DesiredRegion   string   `json:"desiredRegion" example:"ap-northeast-2"`
// 	CidrBlocks      []string `json:"cidrBlocks" example:""`
// }

type RecommendVNetResponse struct {
	recommendation.RecommendedVNetList
}

// RecommendVNet godoc
// @ID RecommendVNet
// @Summary Recommend an appropriate virtual network for cloud migration
// @Description Recommend an appropriate virtual network for cloud migration
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Description - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredProvider and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Tags [Recommendation] Resources for VM infrastructure
// @Accept json
// @Produce	json
// @Param UserInfra body RecommendVmInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendVNetResponse "The result of recommended vNet"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/resources/vNet [post]
func RecommendVNet(c echo.Context) error {

	// [Input]
	var req RecommendVmInfraRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}
	log.Trace().Msgf("req: %v\n", req)

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	// Validate the input
	if desiredProvider == "" {
		err := fmt.Errorf("invalid request: 'desiredProvider' is required")
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}
	if desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	// [Process]
	ret, err := recommendation.RecommendVNet(desiredProvider, desiredRegion, req.OnpremiseInfraModel)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend vNet")
		res := common.SimpleMsg{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// [Output]
	res := RecommendVNetResponse{}
	res.Description = "Recommended vNet information list"
	res.Count = len(ret)

	tempList := []recommendation.RecommendedVNet{}
	for _, vNet := range ret {
		tempList = append(tempList, recommendation.RecommendedVNet{
			Status:      string(recommendation.FullyRecommended),
			Description: vNet.Description,
			TargetVNet:  vNet,
		})
	}
	res.TargetVNetList = tempList

	return c.JSON(http.StatusOK, res)
}

type RecommendSecurityGroupRequest struct {
	// ! To be replaced with the actual model
	// FirewallRules []inframodel.FirewallRuleProperty `json:"firewallRules" example:""`
	FirewallRules []FirewallRuleProperty `json:"firewallRules" example:""`
}

// To be replaced with the actual model
type FirewallRuleProperty struct { // note: reference command `sudo ufw status verbose`
	SrcCIDR   string `json:"srcCIDR,omitempty"`
	DstCIDR   string `json:"dstCIDR,omitempty"`
	SrcPorts  string `json:"srcPorts,omitempty"`
	DstPorts  string `json:"dstPorts,omitempty"`
	Protocol  string `json:"protocol,omitempty"`  // TCP, UDP, ICMP
	Direction string `json:"direction,omitempty"` // inbound, outbound
	Action    string `json:"action,omitempty"`    // allow, deny
}

type RecommendSecurityGroupResponse struct {
	recommendation.RecommendedSecurityGroupList
}

// RecommendSecurityGroups godoc
// @ID RecommendSecurityGroups
// @Summary Recommend an appropriate security group for cloud migration
// @Description Recommend an appropriate security group for cloud migration
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Description - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredProvider and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Tags [Recommendation] Resources for VM infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendVmInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendSecurityGroupResponse "The result of recommended security groups"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/resources/securityGroups [post]
func RecommendSecurityGroups(c echo.Context) error {

	// [Input]
	var req RecommendVmInfraRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}
	log.Trace().Msgf("req: %v\n", req)

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	// Validate the input
	if desiredProvider == "" {
		err := fmt.Errorf("invalid request: 'desiredProvider' is required")
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}
	if desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	// [Process]
	ret, err := recommendation.RecommendSecurityGroups(desiredProvider, desiredRegion, req.OnpremiseInfraModel.Servers)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend security groups")
		res := common.SimpleMsg{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, res)
	}

	// [Output]
	log.Debug().Msgf("recommendedSecurityGroupsList: %v", ret)

	return c.JSON(http.StatusOK, ret)
}

type RecommendVmSpecResponse struct {
	recommendation.RecommendedVmSpecList
}

// RecommendVmSpecs godoc
// @ID RecommendVmSpecs
// @Summary Recommend an appropriate VM specification for cloud migration
// @Description Recommend an appropriate VM specification for cloud migration
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Description - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredProvider and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Tags [Recommendation] Resources for VM infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendVmInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendVmSpecResponse "The result of recommended VM specifications"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/resources/vmSpecs [post]
func RecommendVmSpecs(c echo.Context) error {

	// [Input]
	var req RecommendVmInfraRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}
	log.Trace().Msgf("req: %v\n", req)

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	// Validate the input
	if desiredProvider == "" {
		err := fmt.Errorf("invalid request: 'desiredProvider' is required")
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}
	if desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	// [Process]
	recommendedVmSpecList := recommendation.RecommendedVmSpecList{}
	for i, server := range req.OnpremiseInfraModel.Servers {

		// Initialize a temporary RecommendedVmSpec object
		temp := recommendation.RecommendedVmSpec{
			SourceServer:     server.Hostname,
			Description:      fmt.Sprintf("Recommended VM specs for server %d: %s", i+1, server.Hostname),
			Status:           string(recommendation.NothingRecommended),
			Count:            0,
			TargetVmSpecList: []tbmodel.TbSpecInfo{},
		}

		// Recommend VM specs for the server
		specList, count, err := recommendation.RecommendVmSpecs(desiredProvider, desiredRegion, server, 5)
		// Handle errors and empty recommendations
		if err != nil {
			log.Error().Err(err).Msg("failed to recommend VM specs")
			recommendedVmSpecList.RecommendedVmSpecList = append(recommendedVmSpecList.RecommendedVmSpecList, temp)
			continue
		}
		log.Trace().Msgf("specList: %v, count: %d", specList, count)
		if count == 0 {
			log.Warn().Msgf("no VM specs recommended for server: %s", server.Hostname)
			recommendedVmSpecList.RecommendedVmSpecList = append(recommendedVmSpecList.RecommendedVmSpecList, temp)
			continue
		}
		// Update the temporary object with the recommended specs
		temp.Status = string(recommendation.FullyRecommended)
		temp.Count = count
		temp.TargetVmSpecList = specList
		recommendedVmSpecList.RecommendedVmSpecList = append(recommendedVmSpecList.RecommendedVmSpecList, temp)
	}

	// [Output]
	recommendedVmSpecList.Description = "A collection of recommended VM specs across multiple source servers"
	recommendedVmSpecList.Count = len(recommendedVmSpecList.RecommendedVmSpecList)

	return c.JSON(http.StatusOK, recommendedVmSpecList)
}

type RecommendVmOsImageResponse struct {
	recommendation.RecommendedVmOsImageList
}

// RecommendVmOsImages godoc
// @ID RecommendVmOsImages
// @Summary Recommend an appropriate OS image for cloud migration
// @Description Recommend an appropriate OS image for cloud migration
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Description - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredProvider and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Tags [Recommendation] Resources for VM infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendVmInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendVmOsImageResponse "The result of recommended VM OS images"
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /recommendation/resources/vmOsImages [post]
func RecommendVmOsImages(c echo.Context) error {
	// [Input]
	var req RecommendVmInfraRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}
	log.Trace().Msgf("req: %v\n", req)

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	// Validate the input
	if desiredProvider == "" {
		err := fmt.Errorf("invalid request: 'desiredProvider' is required")
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}
	if desiredRegion == "" {
		err := fmt.Errorf("invalid request: 'desiredRegion' is required")
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	}

	// [Process]
	recommendedOsImageList := recommendation.RecommendedVmOsImageList{}
	for i, server := range req.OnpremiseInfraModel.Servers {

		temp := recommendation.RecommendedVmOsImage{
			Status:          string(recommendation.NothingRecommended),
			SourceServer:    server.Hostname,
			Description:     fmt.Sprintf("Recommended VM OS images for server %d: %s", i+1, server.Hostname),
			TargetVmOsImage: tbmodel.TbImageInfo{},
		}

		vmOsImage, err := recommendation.RecommendVmOsImage(desiredProvider, desiredRegion, server)
		if err != nil {
			log.Error().Err(err).Msg("failed to recommend VM OS images")
			recommendedOsImageList.RecommendedVmOsImageList = append(recommendedOsImageList.RecommendedVmOsImageList, temp)
			continue
		}
		log.Trace().Msgf("vmOsImage: %v", vmOsImage)

		temp.Status = string(recommendation.FullyRecommended)
		temp.TargetVmOsImage = vmOsImage
		recommendedOsImageList.RecommendedVmOsImageList = append(recommendedOsImageList.RecommendedVmOsImageList, temp)
	}
	// [Output]
	recommendedOsImageList.Description = "A collection of recommended VM OS images across multiple source servers"
	recommendedOsImageList.Count = len(recommendedOsImageList.RecommendedVmOsImageList)
	return c.JSON(http.StatusOK, recommendedOsImageList)
}
