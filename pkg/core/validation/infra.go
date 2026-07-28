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

package validation

import (
	"fmt"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	"github.com/rs/zerolog/log"
)

// ValidateTargetInfra checks whether targetInfraModel is internally consistent
// and can be migrated into namespace nsId, given how resources are provisioned
// under useExisting:
//   - useExisting=false: CreateInfra creates fresh VNet/SshKey/SecurityGroups, so
//     none of them may already exist.
//   - useExisting=true: CreateInfraWithExisting reuses a resource by ID if found,
//     otherwise falls back to creating it from the accompanying Target*Req data,
//     so a missing resource is only an error when that fallback data is absent.
//
// In both modes the check performs Tumblebug reads only - no resource is created,
// modified, or deleted. Because state can change between this call and an actual
// migration, a "valid" result is a best-effort snapshot, not a guarantee; the
// migration path re-runs this same validation immediately before provisioning.
//
// All applicable checks run to completion and their issues are accumulated,
// rather than stopping at the first failure, so a caller (e.g. a Portal UI) can
// surface every problem found in a single call.
func ValidateTargetInfra(nsId string, targetInfraModel *cloudmodel.RecommendedInfra, useExisting bool) ValidationResult {
	if targetInfraModel == nil {
		return newResult([]ValidationIssue{{
			Code:     CodeRequiredFieldMissing,
			Severity: SeverityError,
			Path:     "target",
			Message:  "target infrastructure model is nil",
		}})
	}

	var issues []ValidationIssue

	// 1. Name format and referential integrity (shared with /naming/validation).
	if ok, detail := common.ValidateComposedNames(*targetInfraModel); !ok {
		issues = append(issues, ValidationIssue{
			Code:     CodeReferentialIntegrity,
			Severity: SeverityError,
			Path:     "target",
			Message:  detail,
		})
	}

	if targetInfraModel.TargetInfra.Name == "" {
		issues = append(issues, requiredFieldIssue("targetInfra.name", "target Infra (MCI) name is required"))
	}

	// 2. Mode-specific required fields + resource existence/availability.
	if useExisting {
		issues = append(issues, validateNodeGroupIdsPresent(targetInfraModel.TargetInfra.NodeGroups)...)
		issues = append(issues, checkAvailabilityForExisting(nsId, targetInfraModel)...)
	} else {
		issues = append(issues, validateFreshCreationNamesPresent(*targetInfraModel)...)
		issues = append(issues, checkNotAlreadyExistsForFreshCreation(nsId, *targetInfraModel)...)
	}

	// 3. Spec/Image compatibility - required in both modes.
	for i, ng := range targetInfraModel.TargetInfra.NodeGroups {
		path := fmt.Sprintf("targetInfra.nodeGroups[%d]", i)
		if issue := checkSpecImageCompatibility(ng, path); issue != nil {
			issues = append(issues, *issue)
		}
	}

	// 4. Infra (MCI) name collision - must not already exist, regardless of mode.
	if targetInfraModel.TargetInfra.Name != "" {
		infraInfo, err := tbclient.NewSession().ReadInfra(nsId, targetInfraModel.TargetInfra.Name)
		if err == nil && infraInfo.Id != "" {
			issues = append(issues, ValidationIssue{
				Code:     CodeResourceAlreadyExists,
				Severity: SeverityError,
				Path:     "targetInfra.name",
				Message:  fmt.Sprintf("Infra '%s' already exists in namespace '%s'", targetInfraModel.TargetInfra.Name, nsId),
			})
		}
	}

	result := newResult(issues)
	log.Debug().Msgf("validated target infra model (nsId: %s, useExisting: %t, valid: %t, issues: %d)",
		nsId, useExisting, result.Valid, len(result.Issues))
	return result
}

func requiredFieldIssue(path, message string) ValidationIssue {
	return ValidationIssue{Code: CodeRequiredFieldMissing, Severity: SeverityError, Path: path, Message: message}
}

// validateFreshCreationNamesPresent checks the name fields CreateInfra needs to create fresh resources.
func validateFreshCreationNamesPresent(target cloudmodel.RecommendedInfra) []ValidationIssue {
	var issues []ValidationIssue
	if target.TargetVNet.Name == "" {
		issues = append(issues, requiredFieldIssue("targetVNet.name", "target VNet name is required"))
	}
	if target.TargetSshKey.Name == "" {
		issues = append(issues, requiredFieldIssue("targetSshKey.name", "target SSH key name is required"))
	}
	for i, sg := range target.TargetSecurityGroupList {
		if sg.Name == "" {
			issues = append(issues, requiredFieldIssue(fmt.Sprintf("targetSecurityGroupList[%d].name", i), "target security group name is required"))
		}
	}
	return issues
}

// checkNotAlreadyExistsForFreshCreation ensures the resources CreateInfra is about
// to create do not already exist in the namespace.
func checkNotAlreadyExistsForFreshCreation(nsId string, target cloudmodel.RecommendedInfra) []ValidationIssue {
	var issues []ValidationIssue

	if target.TargetVNet.Name != "" {
		if vNetInfo, err := tbclient.NewSession().ReadVNet(nsId, target.TargetVNet.Name); err == nil && vNetInfo.Id != "" {
			issues = append(issues, ValidationIssue{
				Code: CodeResourceAlreadyExists, Severity: SeverityError, Path: "targetVNet.name",
				Message: fmt.Sprintf("vNet '%s' already exists in namespace '%s'", target.TargetVNet.Name, nsId),
			})
		}
	}

	if target.TargetSshKey.Name != "" {
		if sshKeyInfo, err := tbclient.NewSession().ReadSshKey(nsId, target.TargetSshKey.Name); err == nil && sshKeyInfo.Id != "" {
			issues = append(issues, ValidationIssue{
				Code: CodeResourceAlreadyExists, Severity: SeverityError, Path: "targetSshKey.name",
				Message: fmt.Sprintf("SSH key '%s' already exists in namespace '%s'", target.TargetSshKey.Name, nsId),
			})
		}
	}

	for i, sg := range target.TargetSecurityGroupList {
		if sg.Name == "" {
			continue
		}
		if sgInfo, err := tbclient.NewSession().ReadSecurityGroup(nsId, sg.Name); err == nil && sgInfo.Id != "" {
			issues = append(issues, ValidationIssue{
				Code: CodeResourceAlreadyExists, Severity: SeverityError, Path: fmt.Sprintf("targetSecurityGroupList[%d].name", i),
				Message: fmt.Sprintf("security group '%s' already exists in namespace '%s'", sg.Name, nsId),
			})
		}
	}

	return issues
}

// validateNodeGroupIdsPresent checks the ID fields CreateInfraWithExisting needs per NodeGroup.
func validateNodeGroupIdsPresent(nodeGroups []cloudmodel.CreateNodeGroupReq) []ValidationIssue {
	var issues []ValidationIssue
	for i, ng := range nodeGroups {
		path := fmt.Sprintf("targetInfra.nodeGroups[%d]", i)
		if ng.VNetId == "" {
			issues = append(issues, requiredFieldIssue(path+".vNetId", fmt.Sprintf("VNet ID is required for nodegroup '%s' in useExisting mode", ng.Name)))
		}
		if ng.SshKeyId == "" {
			issues = append(issues, requiredFieldIssue(path+".sshKeyId", fmt.Sprintf("SSH key ID is required for nodegroup '%s' in useExisting mode", ng.Name)))
		}
		if len(ng.SecurityGroupIds) == 0 {
			issues = append(issues, requiredFieldIssue(path+".securityGroupIds", fmt.Sprintf("security group IDs are required for nodegroup '%s' in useExisting mode", ng.Name)))
		}
	}
	return issues
}

// checkAvailabilityForExisting mirrors the use-or-create decision that
// CreateInfraWithExisting makes for each VNet/SshKey/SecurityGroup, without
// creating anything: it is an error only when a required resource is missing
// AND the accompanying creation data needed to fall back to creating it is
// also missing or invalid.
func checkAvailabilityForExisting(nsId string, target *cloudmodel.RecommendedInfra) []ValidationIssue {
	var issues []ValidationIssue

	for _, netRequirement := range DeriveNetworkRequirements(target.TargetInfra.NodeGroups) {
		if _, issue := CheckNetworkAvailability(nsId, netRequirement, target.TargetVNet); issue != nil {
			issues = append(issues, *issue)
		}
	}
	for _, sshKeyRequirement := range DeriveSshKeyRequirements(target.TargetInfra.NodeGroups) {
		if _, issue := CheckSshKeyAvailability(nsId, sshKeyRequirement, target.TargetSshKey); issue != nil {
			issues = append(issues, *issue)
		}
	}
	for _, sgRequirement := range DeriveSecurityGroupRequirements(target.TargetInfra.NodeGroups) {
		if _, issue := CheckSecurityGroupAvailability(nsId, sgRequirement, target.TargetSecurityGroupList); issue != nil {
			issues = append(issues, *issue)
		}
	}

	return issues
}

// checkSpecImageCompatibility validates that a NodeGroup's spec and image are
// well-formed and compatible with each other on their CSP.
func checkSpecImageCompatibility(ng cloudmodel.CreateNodeGroupReq, path string) *ValidationIssue {
	specId := strings.TrimSpace(ng.SpecId)
	imageId := strings.TrimSpace(ng.ImageId)
	connectionName := strings.TrimSpace(ng.ConnectionName)

	if specId == "" || specId == "empty" {
		return ptr(requiredFieldIssue(path+".specId", fmt.Sprintf("invalid specId '%s' in nodegroup '%s'", specId, ng.Name)))
	}
	if imageId == "" || imageId == "empty" {
		return ptr(requiredFieldIssue(path+".imageId", fmt.Sprintf("invalid imageId '%s' in nodegroup '%s'", imageId, ng.Name)))
	}
	if connectionName == "" {
		return ptr(requiredFieldIssue(path+".connectionName", fmt.Sprintf("invalid connectionName '%s' in nodegroup '%s'", connectionName, ng.Name)))
	}

	connectionParts := strings.Split(connectionName, "-")
	if len(connectionParts) < 2 {
		return &ValidationIssue{
			Code: CodeInvalidConnectionName, Severity: SeverityError, Path: path + ".connectionName",
			Message: fmt.Sprintf("invalid connection name format '%s' in nodegroup '%s', expected format: 'csp-region'", connectionName, ng.Name),
		}
	}
	csp := connectionParts[0]

	specInfo, err := tbclient.NewSession().ReadVmSpec("system", specId)
	if err != nil {
		return &ValidationIssue{
			Code: CodeSpecOrImageLookupFailed, Severity: SeverityError, Path: path + ".specId",
			Message: fmt.Sprintf("failed to read VM spec '%s': %s", specId, err.Error()),
		}
	}

	// Note - current imageId format: csp+cspImageName (e.g., alibaba+ubuntu_22_04_x64_20G_alibase_20250722.vhd)
	// ref: https://github.com/cloud-barista/cb-tumblebug/pull/2130#issuecomment-3243624048
	imageKey := imageId
	if !strings.Contains(imageKey, "+") {
		imageKey = fmt.Sprintf("%s+%s", csp, imageId)
	}
	imageInfo, err := tbclient.NewSession().ReadVmOsImage("system", imageKey)
	if err != nil {
		return &ValidationIssue{
			Code: CodeSpecOrImageLookupFailed, Severity: SeverityError, Path: path + ".imageId",
			Message: fmt.Sprintf("failed to read VM OS image '%s': %s", imageKey, err.Error()),
		}
	}

	specInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.SpecInfo, cloudmodel.SpecInfo](specInfo)
	if err != nil {
		return &ValidationIssue{
			Code: CodeSpecOrImageLookupFailed, Severity: SeverityError, Path: path + ".specId",
			Message: fmt.Sprintf("failed to convert spec info for compatibility check (specId: %s): %s", specId, err.Error()),
		}
	}
	imageInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.ImageInfo, cloudmodel.ImageInfo](imageInfo)
	if err != nil {
		return &ValidationIssue{
			Code: CodeSpecOrImageLookupFailed, Severity: SeverityError, Path: path + ".imageId",
			Message: fmt.Sprintf("failed to convert image info for compatibility check (imageId: %s): %s", imageId, err.Error()),
		}
	}

	if !recommendation.CheckSpecImageCompatibility(csp, specInfoConverted, imageInfoConverted) {
		return &ValidationIssue{
			Code: CodeSpecImageIncompatible, Severity: SeverityError, Path: path,
			Message: fmt.Sprintf("VM spec '%s' and image '%s' are incompatible for CSP '%s' in nodegroup '%s'", specId, imageId, csp, ng.Name),
		}
	}
	return nil
}

func ptr(issue ValidationIssue) *ValidationIssue {
	return &issue
}
