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

	"github.com/labstack/echo/v4"
)

type Infrastructure struct {
	Network        string
	Disk           string
	Compute        string
	SecurityGroup  string
	VirtualMachine string
}

type RecommendInfraRequest struct {
	Infrastructure
}

type RecommendInfraResponse struct {
	Infrastructure
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
		return err
	}

	fmt.Print(req.Network)
	fmt.Print(req.Disk)
	fmt.Print(req.Compute)
	fmt.Print(req.SecurityGroup)
	fmt.Print(req.VirtualMachine)

	res := &RecommendInfraResponse{}
	// Process

	// Ouput

	// if err != nil {
	// 	common.CBLog.Error(err)
	// 	mapA := map[string]string{"message": err.Error()}
	// 	return c.JSON(http.StatusInternalServerError, &mapA)
	// }

	return c.JSON(http.StatusOK, res)

}
