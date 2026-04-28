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

// Package controller is to handle REST API for beetle
package controller

import (
	"net/http"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/labstack/echo/v4"
)

// AlignNames godoc
// @ID AlignNames
// @Summary Propagate a resource name change to all dependent child resources
// @Description When a parent/primary resource is renamed (e.g., VNet), this API updates all
// @Description child/dependent references in the model (e.g., SecurityGroup.VNetId, SubGroup.VNetId).
// @Description
// @Description **Supported resourceType values** (cb-tumblebug convention):
// @Description - `vNet` : Rename VNet → propagates to SecurityGroup.VNetId, SubGroup.VNetId
// @Description - `subnet` : Rename Subnet → propagates to SubGroup.SubnetId
// @Description - `sshKey` : Rename SSH Key → propagates to SubGroup.SshKeyId
// @Description - `securityGroup` : Rename SecurityGroup → propagates to SubGroup.SecurityGroupIds
// @Description - `mci` : Rename MCI (no child propagation)
// @Description
// @Description After propagation, names are validated with NameSeed applied (pre-flight check).
// @Description The returned model uses **base names only** (NameSeed is applied at migration time).
// @Description
// @Description See also: [API Guide: Align Names](https://github.com/cloud-barista/cm-beetle/blob/main/docs/api-guide-align-names.md)
// @Description
// @Tags [Infrastructure] Resource Naming
// @Accept  json
// @Produce  json
// @Param resourceType query string true "Resource type to rename" Enums(vNet,subnet,securityGroup,sshKey,mci)
// @Param oldName query string true "Current name of the resource (before change)"
// @Param newName query string true "New name of the resource (after change)"
// @Param UserInfra body cloudmodel.RecommendedVmInfra true "The recommendation model to update"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[cloudmodel.RecommendedVmInfra] "Updated and validated model (base names)"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request or referential integrity failure"
// @Router /naming/alignment [post]
func AlignNames(c echo.Context) error {
	resourceType := c.QueryParam("resourceType")
	oldName := c.QueryParam("oldName")
	newName := c.QueryParam("newName")

	if resourceType == "" || oldName == "" || newName == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("query params 'resourceType', 'oldName', and 'newName' are all required"))
	}

	req := &cloudmodel.RecommendedVmInfra{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request body format"))
	}

	// 1. Propagate the name change from parent to child resources
	propagated := common.PropagateNameChange(*req, resourceType, oldName, newName)

	// 2. Pre-flight validation with NameSeed
	// Apply NameSeed temporarily to validate that names + seed will be valid at migration time.
	// The unseeded base-name model is returned so users can still inspect/modify names before migration.
	seeded := common.ApplyNameSeed(propagated)
	if ok, detail := common.ValidateComposedNames(seeded); !ok {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Referential integrity or name validation failed (with NameSeed applied): "+detail))
	}

	// 3. Return the unseeded model (NameSeed is applied at migration time - Late Binding)
	return c.JSON(http.StatusOK, model.SuccessResponse(propagated))
}

// ValidateNames godoc
// @ID ValidateNames
// @Summary Validate resource names and referential integrity
// @Description Validates that all internal references within a RecommendedVmInfra model
// @Description are consistent and point to existing resources.
// @Description NameSeed is NOT applied here; this validates the base names only.
// @Description
// @Tags [Infrastructure] Resource Naming
// @Accept  json
// @Produce  json
// @Param UserInfra body cloudmodel.RecommendedVmInfra true "The recommendation model to validate"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[any] "Naming and referential integrity are valid"
// @Failure 400 {object} model.ApiResponse[any] "Referential integrity validation failure"
// @Router /naming/validation [post]
func ValidateNames(c echo.Context) error {
	req := &cloudmodel.RecommendedVmInfra{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// Validate referential integrity (base names, no NameSeed applied)
	if ok, detail := common.ValidateReferentialIntegrity(*req); !ok {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Referential integrity validation failed: "+detail))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse[any](nil))
}
