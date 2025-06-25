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
// @Param machineID query string false "Machine ID (e.g., TBD, xxxxxx)"
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
	// machineID := c.QueryParam("machineID")

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
	// if machineID == "" {
	// 	err := fmt.Errorf("invalid request: 'machineID' is required")
	// 	log.Warn().Msg(err.Error())
	// 	return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	// }

	// // Validate by finding the machine ID in the request body
	// found := false
	// machine := inframodel.ServerProperty{}
	// for _, server := range req.OnpremiseInfraModel.Servers {
	// 	if server.MachineID == machineID {
	// 		found = true
	// 		machine = server
	// 		break
	// 	}
	// }
	// if !found {
	// 	err := fmt.Errorf("invalid request: 'machineID' not found in the request body")
	// 	log.Warn().Msg(err.Error())
	// 	return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	// }

	// [Process]
	recommendedVmSpecList := recommendation.RecommendedVmSpecList{}
	for i, server := range req.OnpremiseInfraModel.Servers {

		// Recommend VM specs for the server
		specList, count, err := recommendation.RecommendVmSpecs(desiredProvider, desiredRegion, server, 5)

		// Handle errors and empty recommendations
		if err != nil {
			log.Error().Err(err).Msg("failed to recommend VM specs")

			temp := recommendation.RecommendedVmSpec{
				SourceServers: []string{server.Hostname}, //TODO replace Hostname with MachineID
				Description:   fmt.Sprintf("failed to recommend VM specs for server %d: %s", i+1, server.Hostname),
				Status:        string(recommendation.NothingRecommended),
				TargetVmSpec:  tbmodel.TbSpecInfo{},
			}
			recommendedVmSpecList.RecommendedVmSpecList = append(recommendedVmSpecList.RecommendedVmSpecList, temp)
			continue
		}
		log.Trace().Msgf("specList: %v, count: %d", specList, count)
		if count == 0 {
			log.Warn().Msgf("no VM specs recommended for server: %s", server.Hostname)

			temp := recommendation.RecommendedVmSpec{
				SourceServers: []string{server.Hostname}, //TODO replace Hostname with MachineID
				Description:   fmt.Sprintf("no VM specs recommended for server %d: %s", i+1, server.Hostname),
				Status:        string(recommendation.NothingRecommended),
				TargetVmSpec:  tbmodel.TbSpecInfo{},
			}
			recommendedVmSpecList.RecommendedVmSpecList = append(recommendedVmSpecList.RecommendedVmSpecList, temp)
			continue
		}

		// Recursively check duplicates and append the recommended specs
		for _, spec := range specList {
			// Check if the spec already exists in the list
			exists := false
			idx := -1
			for i, existingSpec := range recommendedVmSpecList.RecommendedVmSpecList {
				if existingSpec.TargetVmSpec.Id == spec.Id {
					exists = true
					idx = i
					break
				}
			}

			// If the spec already exists, append the server to the existing list
			// Otherwise, create a new entry
			if exists {
				recommendedVmSpecList.RecommendedVmSpecList[idx].SourceServers = append(
					recommendedVmSpecList.RecommendedVmSpecList[idx].SourceServers,
					server.Hostname, //TODO replace Hostname with MachineID
				)
			} else {
				temp := recommendation.RecommendedVmSpec{
					Status:        string(recommendation.FullyRecommended),
					SourceServers: []string{server.Hostname}, //TODO replace Hostname with MachineID
					Description:   fmt.Sprintf("Recommended VM spec for server %d: %s", i+1, server.Hostname),
					TargetVmSpec:  spec,
				}
				recommendedVmSpecList.RecommendedVmSpecList = append(recommendedVmSpecList.RecommendedVmSpecList, temp)
			}
		}
	}

	// [Output]
	countFailed := 0
	for _, spec := range recommendedVmSpecList.RecommendedVmSpecList {
		if spec.Status == string(recommendation.NothingRecommended) {
			countFailed++
		}
	}

	recommendedVmSpecList.Count = len(recommendedVmSpecList.RecommendedVmSpecList)
	switch countFailed {
	case 0:
		recommendedVmSpecList.Status = string(recommendation.FullyRecommended)
		recommendedVmSpecList.Description = "Successfully recommended VM specs for all servers in the source infrastructure"
	case recommendedVmSpecList.Count:
		recommendedVmSpecList.Status = string(recommendation.NothingRecommended)
		recommendedVmSpecList.Description = "Unable to recommend any VM specs for the servers in the source infrastructure"
	default:
		recommendedVmSpecList.Status = string(recommendation.PartiallyRecommended)
		recommendedVmSpecList.Description = fmt.Sprintf(
			"Partially recommended VM specs: successful for %d servers, failed for %d servers",
			recommendedVmSpecList.Count-countFailed, countFailed,
		)
	}

	log.Debug().Msgf("recommendedVmSpecList: %+v", recommendedVmSpecList)

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

		vmOsImageList, err := recommendation.RecommendVmOsImages(desiredProvider, desiredRegion, server, 3)

		// Handle errors and empty recommendations
		if err != nil {
			log.Error().Err(err).Msg("failed to recommend VM OS images")

			temp := recommendation.RecommendedVmOsImage{
				Status:          string(recommendation.NothingRecommended),
				SourceServers:   []string{server.Hostname}, //TODO replace Hostname with MachineID
				Description:     fmt.Sprintf("Recommended VM OS images for server %d: %s", i+1, server.Hostname),
				TargetVmOsImage: tbmodel.TbImageInfo{},
			}
			recommendedOsImageList.RecommendedVmOsImageList = append(recommendedOsImageList.RecommendedVmOsImageList, temp)
			continue
		}

		if len(vmOsImageList) == 0 {
			log.Warn().Msgf("no VM OS images recommended for server: %s", server.Hostname)

			temp := recommendation.RecommendedVmOsImage{
				Status:          string(recommendation.NothingRecommended),
				SourceServers:   []string{server.Hostname}, //TODO replace Hostname with MachineID
				Description:     fmt.Sprintf("No VM OS images recommended for server %d: %s", i+1, server.Hostname),
				TargetVmOsImage: tbmodel.TbImageInfo{},
			}
			recommendedOsImageList.RecommendedVmOsImageList = append(recommendedOsImageList.RecommendedVmOsImageList, temp)
			continue
		}

		// Recursively check duplicates and append the recommended OS images
		for _, vmOsImageAndScore := range vmOsImageList {
			vmOsImage := vmOsImageAndScore.VmOsImageInfo
			// Check if the OS image already exists in the list
			exists := false
			idx := -1
			for i, existingOsImage := range recommendedOsImageList.RecommendedVmOsImageList {
				if existingOsImage.TargetVmOsImage.Id == vmOsImage.Id {
					exists = true
					idx = i
					break
				}
			}
			// If the OS image already exists, append the server to the existing list
			// Otherwise, create a new entry
			if exists {
				recommendedOsImageList.RecommendedVmOsImageList[idx].SourceServers = append(
					recommendedOsImageList.RecommendedVmOsImageList[idx].SourceServers,
					server.Hostname, //TODO replace Hostname with MachineID
				)
			} else {
				temp := recommendation.RecommendedVmOsImage{
					Status:          string(recommendation.FullyRecommended),
					SourceServers:   []string{server.Hostname}, //TODO replace Hostname with	 MachineID
					Description:     fmt.Sprintf("Recommended VM OS image for server %d: %s", i+1, server.Hostname),
					TargetVmOsImage: vmOsImage,
				}
				recommendedOsImageList.RecommendedVmOsImageList = append(recommendedOsImageList.RecommendedVmOsImageList, temp)
			}
		}
	}

	// [Output]
	countFailed := 0
	for _, osImage := range recommendedOsImageList.RecommendedVmOsImageList {
		if osImage.Status == string(recommendation.NothingRecommended) {
			countFailed++
		}
	}
	recommendedOsImageList.Count = len(recommendedOsImageList.RecommendedVmOsImageList)
	switch countFailed {
	case 0:
		recommendedOsImageList.Status = string(recommendation.FullyRecommended)
		recommendedOsImageList.Description = "Successfully recommended VM OS images for the servers in the source computing infra"
	case recommendedOsImageList.Count:
		recommendedOsImageList.Status = string(recommendation.NothingRecommended)
		recommendedOsImageList.Description = "Unable to recommend any VM OS images for the servers in the source computing infra"
	default:
		recommendedOsImageList.Status = string(recommendation.PartiallyRecommended)
		recommendedOsImageList.Description = fmt.Sprintf(
			"Partially recommended VM OS images: successful for %d servers, failed for %d servers",
			recommendedOsImageList.Count-countFailed, countFailed,
		)
	}

	return c.JSON(http.StatusOK, recommendedOsImageList)
}
