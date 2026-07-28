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

// Package controller has handlers and their request/response bodies for validation APIs
package controller

import (
	"fmt"
	"net/http"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/validation"
	"github.com/labstack/echo/v4"

	"github.com/rs/zerolog/log"
)

type ValidateInfraRequest struct {
	cloudmodel.RecommendedInfra
}

// ValidateInfra godoc
// @ID ValidateInfra
// @Summary (Preview) Validate a target infrastructure model before migration
// @Description Runs, without creating or modifying any resource, the same checks
// @Description migration execution performs immediately before provisioning:
// @Description
// @Description - **Naming & referential integrity**: names must be 3-63 alphanumeric/hyphen
// @Description   characters, and internal references must resolve within the submitted model
// @Description   (e.g. a NodeGroup's `securityGroupIds` must match a `name` in `targetSecurityGroupList`).
// @Description - **Required fields**, which differ by `useExisting`:
// @Description   - `true`: each NodeGroup's `vNetId`, `sshKeyId`, `securityGroupIds` must be set.
// @Description   - `false`: `targetVNet.name`, `targetSshKey.name`, and each security group's `name` must be set.
// @Description - **Resource name collision / availability** against Tumblebug, which also differs by `useExisting`:
// @Description   - `false`: the VNet/SSH key/security groups to be created must NOT already exist in the
// @Description     namespace (e.g. `targetVNet.name: "vnet-01"` fails if a VNet named `vnet-01` already exists).
// @Description   - `true`: an existing resource must be under the same CSP/region connection the NodeGroup
// @Description     requests (e.g. reusing a VNet provisioned under connection `aws-ap-northeast-2` while the
// @Description     NodeGroup's `connectionName` is `gcp-asia-northeast3` fails); a resource that doesn't exist
// @Description     yet must have enough data alongside it to create it instead (e.g. `targetVNet.cidrBlock`).
// @Description - **VM spec/image compatibility** per NodeGroup: `specId`, `imageId`, and `connectionName` must
// @Description   be set, `connectionName` must be in `csp-region` format (e.g. `aws-ap-northeast-2`), and the
// @Description   resolved spec/image pair must be compatible for that CSP (e.g. an `x86_64` spec paired with
// @Description   an `arm64` image fails).
// @Description - **Infra (MCI) name collision**: `targetInfra.name` must not already exist in the namespace.
// @Description
// @Description Always returns HTTP 200: the response body's `valid` field and
// @Description `issues` list carry the outcome, since a failed check is a normal,
// @Description successfully-answered result rather than a malformed request.
// @Description 400 is reserved for request body/parameter errors.
// @Description
// @Description Because Tumblebug/CSP state can change afterward, a `valid: true`
// @Description result is a best-effort snapshot, not a guarantee - the migration
// @Description API re-runs this same validation immediately before provisioning.
// @Tags [Validation] Target Cloud Configuration (Preview), [Migration] Infrastructure
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param useExisting query bool false "Validate as if reusing existing resources (VNet, SSH Key, Security Group) instead of creating new ones (default: true); should match the `useExisting` value intended for the actual migration call"
// @Param infraInfo body ValidateInfraRequest true "The target infrastructure model to validate"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[validation.ValidationResult] "Validation outcome: valid flag plus zero or more issues"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request body or parameters"
// @Router /validation/ns/{nsId}/infra [post]
func ValidateInfra(c echo.Context) error {

	// [Input]
	nsId := c.Param("nsId")
	if nsId == "" {
		err := fmt.Errorf("invalid request, namespace ID (nsId: %s) is required", nsId)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	req := new(ValidateInfraRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	useExistingStr := c.QueryParam("useExisting")
	useExisting := true
	if useExistingStr == "false" {
		useExisting = false
	}

	// [Process]
	result := validation.ValidateTargetInfra(nsId, &req.RecommendedInfra, useExisting)

	// [Output]
	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}
