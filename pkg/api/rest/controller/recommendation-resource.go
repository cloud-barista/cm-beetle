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

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

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
// @Param UserInfra body RecommendInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[cloudmodel.RecommendedVNetList] "Successfully recommended vNet(s)"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Router /recommendation/resources/vNet [post]
func RecommendVNet(c echo.Context) error {

	// [Input]
	var req RecommendInfraRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}
	log.Trace().Msgf("req: %v\n", req)

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	// Validate the input
	if desiredProvider == "" {
		log.Warn().Msg("desiredProvider is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Provider required"))
	}
	if desiredRegion == "" {
		log.Warn().Msg("desiredRegion is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Region required"))
	}

	// [Process]
	ret, err := recommendation.RecommendVNet(desiredProvider, desiredRegion, req.OnpremiseInfraModel)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend vNet")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("VNet recommendation failed"))
	}

	// [Output]
	RecommendVNetList := cloudmodel.RecommendedVNetList{}
	RecommendVNetList.Description = "Recommended vNet information list"
	RecommendVNetList.Count = len(ret)

	tempList := []cloudmodel.RecommendedVNet{}
	for _, vNet := range ret {
		tempList = append(tempList, cloudmodel.RecommendedVNet{
			Status:      string(recommendation.FullyRecommended),
			Description: vNet.Description,
			TargetVNet:  vNet,
		})
	}
	RecommendVNetList.TargetVNetList = tempList

	successMsg := fmt.Sprintf("Recommended %d vNet(s) for %s %s", len(ret), desiredProvider, desiredRegion)
	res := model.SuccessResponseWithMessage(RecommendVNetList, successMsg)

	return c.JSON(http.StatusOK, res)
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
// @Param UserInfra body RecommendInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[cloudmodel.RecommendedSecurityGroupList] "Successfully recommended security group(s)"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Router /recommendation/resources/securityGroups [post]
func RecommendSecurityGroups(c echo.Context) error {

	// [Input]
	var req RecommendInfraRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}
	log.Trace().Msgf("req: %v\n", req)

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	// Validate the input
	if desiredProvider == "" {
		log.Warn().Msg("desiredProvider is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Provider required"))
	}
	if desiredRegion == "" {
		log.Warn().Msg("desiredRegion is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Region required"))
	}

	// [Process]
	ret, err := recommendation.RecommendSecurityGroups(desiredProvider, desiredRegion, req.OnpremiseInfraModel.Servers)
	if err != nil {
		log.Error().Err(err).Msg("failed to recommend security groups")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Security group recommendation failed"))
	}

	// [Output]
	log.Debug().Msgf("recommendedSecurityGroupsList: %v", ret)
	successMsg := fmt.Sprintf("Recommended security group(s) for %s %s", desiredProvider, desiredRegion)
	res := model.SuccessResponseWithMessage(ret, successMsg)

	return c.JSON(http.StatusOK, res)
}

type RecommendVmSpecResponse struct {
	cloudmodel.RecommendedSpecList
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
// @Description - If `targetMachineId` is provided, only that specific machine will be processed.
// @Tags [Recommendation] Resources for VM infrastructure
// @Accept  json
// @Produce  json
// @Param UserInfra body RecommendInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param targetMachineId query string false "Target Machine ID to focus recommendation on (optional)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[cloudmodel.RecommendedSpecList] "Successfully recommended VM spec(s)"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Router /recommendation/resources/specs [post]
func RecommendVmSpecs(c echo.Context) error {

	// [Input]
	var req RecommendInfraRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}
	log.Trace().Msgf("req: %v\n", req)

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")
	targetMachineId := c.QueryParam("targetMachineId") // Add targetMachineId parameter

	// Validate the input
	if desiredProvider == "" {
		log.Warn().Msg("desiredProvider is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Provider required"))
	}
	if desiredRegion == "" {
		log.Warn().Msg("desiredRegion is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Region required"))
	}
	// if machineID == "" {
	// 	err := fmt.Errorf("invalid request: 'machineID' is required")
	// 	log.Warn().Msg(err.Error())
	// 	return c.JSON(http.StatusBadRequest, common.SimpleMsg{Message: err.Error()})
	// }

	// Handle targetMachineId filtering if provided
	var serversToProcess []onpremmodel.ServerProperty
	if targetMachineId != "" {
		// Validate by finding the machine ID in the request body
		targetMachine := onpremmodel.ServerProperty{}
		found := false
		for _, server := range req.OnpremiseInfraModel.Servers {
			if server.MachineId == targetMachineId {
				found = true
				targetMachine = server
				break
			}
		}
		if !found {
			log.Warn().Msgf("targetMachineId '%s' not found in request body", targetMachineId)
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(fmt.Sprintf("Machine ID '%s' not found", targetMachineId)))
		}
		// Process only the target machine
		serversToProcess = []onpremmodel.ServerProperty{targetMachine}
		log.Info().Msgf("Processing VM specs for target machine: %s", targetMachineId)
	} else {
		// Process all servers in the infrastructure
		serversToProcess = req.OnpremiseInfraModel.Servers
		log.Info().Msgf("Processing VM specs for all servers (%d total)", len(serversToProcess))
	}

	// [Process]
	recommendedVmSpecList := cloudmodel.RecommendedSpecList{}
	for i, server := range serversToProcess {

		specsLimit := recommendation.GetDefaultSpecsLimit()
		// Recommend VM specs for the server
		specList, count, err := recommendation.RecommendVmSpecs(desiredProvider, desiredRegion, server, specsLimit)

		// Handle errors and empty recommendations
		if err != nil {
			log.Error().Err(err).Msg("failed to recommend VM specs")

			temp := cloudmodel.RecommendedSpec{
				SourceServers: []string{server.MachineId}, // Set MachineId to identify the source server
				Description:   fmt.Sprintf("failed to recommend VM specs for server %d: %s", i+1, server.MachineId),
				Status:        string(recommendation.NothingRecommended),
				TargetSpec:    cloudmodel.SpecInfo{},
			}
			recommendedVmSpecList.RecommendedSpecList = append(recommendedVmSpecList.RecommendedSpecList, temp)
			continue
		}
		log.Trace().Msgf("specList: %v, count: %d", specList, count)
		if count == 0 {
			log.Warn().Msgf("no VM specs recommended for server: %s", server.MachineId)

			temp := cloudmodel.RecommendedSpec{
				SourceServers: []string{server.MachineId}, // Set MachineId to identify the source server
				Description:   fmt.Sprintf("no VM specs recommended for server %d: %s", i+1, server.MachineId),
				Status:        string(recommendation.NothingRecommended),
				TargetSpec:    cloudmodel.SpecInfo{},
			}
			recommendedVmSpecList.RecommendedSpecList = append(recommendedVmSpecList.RecommendedSpecList, temp)
			continue
		}

		// Recursively check duplicates and append the recommended specs
		for _, spec := range specList {
			// Check if the spec already exists in the list
			exists := false
			idx := -1
			for i, existingSpec := range recommendedVmSpecList.RecommendedSpecList {
				if existingSpec.TargetSpec.Id == spec.Id {
					exists = true
					idx = i
					break
				}
			}

			// If the spec already exists, append the server to the existing list
			// Otherwise, create a new entry
			if exists {
				recommendedVmSpecList.RecommendedSpecList[idx].SourceServers = append(
					recommendedVmSpecList.RecommendedSpecList[idx].SourceServers,
					server.MachineId, // Set MachineId to identify the source server
				)
			} else {
				temp := cloudmodel.RecommendedSpec{
					Status:        string(recommendation.FullyRecommended),
					SourceServers: []string{server.MachineId}, // Set MachineId to identify the source server
					Description:   fmt.Sprintf("Recommended VM spec for server %d: %s", i+1, server.MachineId),
					TargetSpec:    spec,
				}
				recommendedVmSpecList.RecommendedSpecList = append(recommendedVmSpecList.RecommendedSpecList, temp)
			}
		}
	}

	// [Output]
	countFailed := 0
	for _, spec := range recommendedVmSpecList.RecommendedSpecList {
		if spec.Status == string(recommendation.NothingRecommended) {
			countFailed++
		}
	}

	recommendedVmSpecList.Count = len(recommendedVmSpecList.RecommendedSpecList)
	switch countFailed {
	case 0:
		recommendedVmSpecList.Status = string(recommendation.FullyRecommended)
		if targetMachineId != "" {
			recommendedVmSpecList.Description = fmt.Sprintf("Recommended VM spec(s) for machine '%s'", targetMachineId)
		} else {
			recommendedVmSpecList.Description = fmt.Sprintf("Recommended VM spec(s) for %d server(s)", len(serversToProcess))
		}
	case recommendedVmSpecList.Count:
		recommendedVmSpecList.Status = string(recommendation.NothingRecommended)
		if targetMachineId != "" {
			recommendedVmSpecList.Description = fmt.Sprintf("No VM specs available for machine '%s'", targetMachineId)
		} else {
			recommendedVmSpecList.Description = "No VM specs available for any server"
		}
	default:
		recommendedVmSpecList.Status = string(recommendation.PartiallyRecommended)
		successCount := recommendedVmSpecList.Count - countFailed
		if targetMachineId != "" {
			recommendedVmSpecList.Description = fmt.Sprintf(
				"Recommended %d of %d VM spec(s) for machine '%s'",
				successCount, recommendedVmSpecList.Count, targetMachineId,
			)
		} else {
			recommendedVmSpecList.Description = fmt.Sprintf(
				"Recommended %d of %d VM spec(s)",
				successCount, recommendedVmSpecList.Count,
			)
		}
	}

	log.Debug().Msgf("recommendedVmSpecList: %+v", recommendedVmSpecList)

	successMsg := fmt.Sprintf("Recommended VM spec(s) for %s %s", desiredProvider, desiredRegion)
	res := model.SuccessResponseWithMessage(recommendedVmSpecList, successMsg)

	return c.JSON(http.StatusOK, res)
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
// @Param UserInfra body RecommendInfraRequest true "Specify the your infrastructure to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[cloudmodel.RecommendedOsImageList] "Successfully recommended VM OS image(s)"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Router /recommendation/resources/osImages [post]
func RecommendVmOsImages(c echo.Context) error {
	// [Input]
	var req RecommendInfraRequest
	if err := c.Bind(&req); err != nil {
		log.Warn().Err(err).Msg("failed to bind request body")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}
	log.Trace().Msgf("req: %v\n", req)

	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	// Validate the input
	if desiredProvider == "" {
		log.Warn().Msg("desiredProvider is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Provider required"))
	}
	if desiredRegion == "" {
		log.Warn().Msg("desiredRegion is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Region required"))
	}

	// [Process]
	recommendedOsImageList := cloudmodel.RecommendedOsImageList{}
	for i, server := range req.OnpremiseInfraModel.Servers {

		imagesLimit := recommendation.GetDefaultImagesLimit()
		vmOsImageList, err := recommendation.RecommendVmOsImages(desiredProvider, desiredRegion, server, imagesLimit)

		// Handle errors and empty recommendations
		if err != nil {
			log.Error().Err(err).Msg("failed to recommend VM OS images")

			temp := cloudmodel.RecommendedOsImage{
				Status:        string(recommendation.NothingRecommended),
				SourceServers: []string{server.MachineId}, // Set MachineId to identify the source server
				Description:   fmt.Sprintf("Failed to recommend VM OS images for server %d: %s", i+1, server.MachineId),
				TargetOsImage: cloudmodel.ImageInfo{},
			}
			recommendedOsImageList.RecommendedOsImageList = append(recommendedOsImageList.RecommendedOsImageList, temp)
			continue
		}

		if len(vmOsImageList) == 0 {
			log.Warn().Msgf("no VM OS images recommended for server: %s", server.MachineId)

			temp := cloudmodel.RecommendedOsImage{
				Status:        string(recommendation.NothingRecommended),
				SourceServers: []string{server.MachineId}, // Set MachineId to identify the source server
				Description:   fmt.Sprintf("No VM OS images recommended for server %d: %s", i+1, server.MachineId),
				TargetOsImage: cloudmodel.ImageInfo{},
			}
			recommendedOsImageList.RecommendedOsImageList = append(recommendedOsImageList.RecommendedOsImageList, temp)
			continue
		}

		// Recursively check duplicates and append the recommended OS images
		for _, vmOsImage := range vmOsImageList {
			// Check if the OS image already exists in the list
			exists := false
			idx := -1
			for i, existingOsImage := range recommendedOsImageList.RecommendedOsImageList {
				if existingOsImage.TargetOsImage.Id == vmOsImage.Id {
					exists = true
					idx = i
					break
				}
			}
			// If the OS image already exists, append the server to the existing list
			// Otherwise, create a new entry
			if exists {
				recommendedOsImageList.RecommendedOsImageList[idx].SourceServers = append(
					recommendedOsImageList.RecommendedOsImageList[idx].SourceServers,
					server.MachineId, // Set MachineId to identify the source server
				)
			} else {
				temp := cloudmodel.RecommendedOsImage{
					Status:        string(recommendation.FullyRecommended),
					SourceServers: []string{server.MachineId}, // Set MachineId to identify the source server
					Description:   fmt.Sprintf("Recommended VM OS image for server %d: %s", i+1, server.MachineId),
					TargetOsImage: vmOsImage,
				}
				recommendedOsImageList.RecommendedOsImageList = append(recommendedOsImageList.RecommendedOsImageList, temp)
			}
		}
	}

	// [Output]
	countFailed := 0
	for _, osImage := range recommendedOsImageList.RecommendedOsImageList {
		if osImage.Status == string(recommendation.NothingRecommended) {
			countFailed++
		}
	}
	recommendedOsImageList.Count = len(recommendedOsImageList.RecommendedOsImageList)
	successCount := recommendedOsImageList.Count - countFailed
	switch countFailed {
	case 0:
		recommendedOsImageList.Status = string(recommendation.FullyRecommended)
		recommendedOsImageList.Description = fmt.Sprintf("Recommended OS image(s) for %d server(s)", len(req.OnpremiseInfraModel.Servers))
	case recommendedOsImageList.Count:
		recommendedOsImageList.Status = string(recommendation.NothingRecommended)
		recommendedOsImageList.Description = "No OS images available for any server"
	default:
		recommendedOsImageList.Status = string(recommendation.PartiallyRecommended)
		recommendedOsImageList.Description = fmt.Sprintf(
			"Recommended %d of %d OS image(s)",
			successCount, recommendedOsImageList.Count,
		)
	}

	log.Debug().Msgf("recommendedVmOsImageList: %+v", recommendedOsImageList)

	successMsg := fmt.Sprintf("Recommended VM OS image(s) for %s %s", desiredProvider, desiredRegion)
	res := model.SuccessResponseWithMessage(recommendedOsImageList, successMsg)

	return c.JSON(http.StatusOK, res)
}
