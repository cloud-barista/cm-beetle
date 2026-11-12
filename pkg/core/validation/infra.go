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
	"context"
	"fmt"
	"slices"
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
//   - useExisting=false: MigrateInfra creates fresh VNet/SshKey/SecurityGroups, so
//     none of them may already exist.
//   - useExisting=true: MigrateInfraWithExisting reuses a resource by ID if found,
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
func ValidateTargetInfra(ctx context.Context, nsId string, targetInfraModel *cloudmodel.RecommendedInfra, useExisting bool) ValidationResult {
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
		issues = append(issues, checkAvailabilityForExisting(ctx, nsId, targetInfraModel)...)
	} else {
		issues = append(issues, validateFreshCreationNamesPresent(*targetInfraModel)...)
		issues = append(issues, checkNotAlreadyExistsForFreshCreation(ctx, nsId, *targetInfraModel)...)
	}

	// 3. Spec/Image compatibility and subnet-zone alignment - required in both modes.
	resolvedSpecs := make([]*cloudmodel.SpecInfo, len(targetInfraModel.TargetInfra.NodeGroups))
	for i, ng := range targetInfraModel.TargetInfra.NodeGroups {
		path := fmt.Sprintf("targetInfra.nodeGroups[%d]", i)
		spec, issue := checkSpecImageCompatibility(ctx, ng, path)
		if issue != nil {
			issues = append(issues, *issue)
		}
		resolvedSpecs[i] = enrichSpecAvailabilityFromModel(spec, ng.SpecId, targetInfraModel.TargetSpecList)
	}
	issues = append(issues, validateSubnetZoneAlignment(ctx, nsId, targetInfraModel, resolvedSpecs, useExisting)...)

	// 4. Infra (MCI) name collision - must not already exist, regardless of mode.
	if targetInfraModel.TargetInfra.Name != "" {
		infraInfo, err := tbclient.NewSession().SetContext(ctx).ReadInfra(nsId, targetInfraModel.TargetInfra.Name)
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

// validateFreshCreationNamesPresent checks the name fields MigrateInfra needs to create fresh resources.
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

// checkNotAlreadyExistsForFreshCreation ensures the resources MigrateInfra is about
// to create do not already exist in the namespace.
func checkNotAlreadyExistsForFreshCreation(ctx context.Context, nsId string, target cloudmodel.RecommendedInfra) []ValidationIssue {
	var issues []ValidationIssue

	if target.TargetVNet.Name != "" {
		if vNetInfo, err := tbclient.NewSession().SetContext(ctx).ReadVNet(nsId, target.TargetVNet.Name); err == nil && vNetInfo.Id != "" {
			issues = append(issues, ValidationIssue{
				Code: CodeResourceAlreadyExists, Severity: SeverityError, Path: "targetVNet.name",
				Message: fmt.Sprintf("vNet '%s' already exists in namespace '%s'", target.TargetVNet.Name, nsId),
			})
		}
	}

	if target.TargetSshKey.Name != "" {
		if sshKeyInfo, err := tbclient.NewSession().SetContext(ctx).ReadSshKey(nsId, target.TargetSshKey.Name); err == nil && sshKeyInfo.Id != "" {
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
		if sgInfo, err := tbclient.NewSession().SetContext(ctx).ReadSecurityGroup(nsId, sg.Name); err == nil && sgInfo.Id != "" {
			issues = append(issues, ValidationIssue{
				Code: CodeResourceAlreadyExists, Severity: SeverityError, Path: fmt.Sprintf("targetSecurityGroupList[%d].name", i),
				Message: fmt.Sprintf("security group '%s' already exists in namespace '%s'", sg.Name, nsId),
			})
		}
	}

	return issues
}

// validateNodeGroupIdsPresent checks the ID fields MigrateInfraWithExisting needs per NodeGroup.
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
// MigrateInfraWithExisting makes for each VNet/SshKey/SecurityGroup, without
// creating anything: it is an error only when a required resource is missing
// AND the accompanying creation data needed to fall back to creating it is
// also missing or invalid.
func checkAvailabilityForExisting(ctx context.Context, nsId string, target *cloudmodel.RecommendedInfra) []ValidationIssue {
	var issues []ValidationIssue

	for _, netRequirement := range DeriveNetworkRequirements(target.TargetInfra.NodeGroups) {
		if _, issue := CheckNetworkAvailability(ctx, nsId, netRequirement, target.TargetVNet); issue != nil {
			issues = append(issues, *issue)
		}
	}
	for _, sshKeyRequirement := range DeriveSshKeyRequirements(target.TargetInfra.NodeGroups) {
		if _, issue := CheckSshKeyAvailability(ctx, nsId, sshKeyRequirement, target.TargetSshKey); issue != nil {
			issues = append(issues, *issue)
		}
	}
	for _, sgRequirement := range DeriveSecurityGroupRequirements(target.TargetInfra.NodeGroups) {
		if _, issue := CheckSecurityGroupAvailability(ctx, nsId, sgRequirement, target.TargetSecurityGroupList); issue != nil {
			issues = append(issues, *issue)
		}
	}

	return issues
}

// checkSpecImageCompatibility validates that a NodeGroup's spec and image are
// well-formed and compatible with each other on their CSP.
func checkSpecImageCompatibility(ctx context.Context, ng cloudmodel.CreateNodeGroupReq, path string) (*cloudmodel.SpecInfo, *ValidationIssue) {
	specId := strings.TrimSpace(ng.SpecId)
	imageId := strings.TrimSpace(ng.ImageId)
	connectionName := strings.TrimSpace(ng.ConnectionName)

	if specId == "" || specId == "empty" {
		return nil, ptr(requiredFieldIssue(path+".specId", fmt.Sprintf("invalid specId '%s' in nodegroup '%s'", specId, ng.Name)))
	}
	if imageId == "" || imageId == "empty" {
		return nil, ptr(requiredFieldIssue(path+".imageId", fmt.Sprintf("invalid imageId '%s' in nodegroup '%s'", imageId, ng.Name)))
	}
	if connectionName == "" {
		return nil, ptr(requiredFieldIssue(path+".connectionName", fmt.Sprintf("invalid connectionName '%s' in nodegroup '%s'", connectionName, ng.Name)))
	}

	connectionParts := strings.Split(connectionName, "-")
	if len(connectionParts) < 2 {
		return nil, &ValidationIssue{
			Code: CodeInvalidConnectionName, Severity: SeverityError, Path: path + ".connectionName",
			Message: fmt.Sprintf("invalid connection name format '%s' in nodegroup '%s', expected format: 'csp-region'", connectionName, ng.Name),
		}
	}
	csp := connectionParts[0]

	specInfo, err := tbclient.NewSession().SetContext(ctx).ReadVmSpec("system", specId)
	if err != nil {
		return nil, &ValidationIssue{
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
	imageInfo, err := tbclient.NewSession().SetContext(ctx).ReadVmOsImage("system", imageKey)
	if err != nil {
		return nil, &ValidationIssue{
			Code: CodeSpecOrImageLookupFailed, Severity: SeverityError, Path: path + ".imageId",
			Message: fmt.Sprintf("failed to read VM OS image '%s': %s", imageKey, err.Error()),
		}
	}

	specInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.SpecInfo, cloudmodel.SpecInfo](specInfo)
	if err != nil {
		return nil, &ValidationIssue{
			Code: CodeSpecOrImageLookupFailed, Severity: SeverityError, Path: path + ".specId",
			Message: fmt.Sprintf("failed to convert spec info for compatibility check (specId: %s): %s", specId, err.Error()),
		}
	}
	imageInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.ImageInfo, cloudmodel.ImageInfo](imageInfo)
	if err != nil {
		return nil, &ValidationIssue{
			Code: CodeSpecOrImageLookupFailed, Severity: SeverityError, Path: path + ".imageId",
			Message: fmt.Sprintf("failed to convert image info for compatibility check (imageId: %s): %s", imageId, err.Error()),
		}
	}

	if !recommendation.CheckSpecImageCompatibility(csp, specInfoConverted, imageInfoConverted) {
		return &specInfoConverted, &ValidationIssue{
			Code: CodeSpecImageIncompatible, Severity: SeverityError, Path: path,
			Message: fmt.Sprintf("VM spec '%s' and image '%s' are incompatible for CSP '%s' in nodegroup '%s'", specId, imageId, csp, ng.Name),
		}
	}
	return &specInfoConverted, nil
}

// enrichSpecAvailabilityFromModel fills runtime availability fields from TargetSpecList when ReadVmSpec omitted them.
func enrichSpecAvailabilityFromModel(spec *cloudmodel.SpecInfo, specId string, modelSpecs []cloudmodel.SpecInfo) *cloudmodel.SpecInfo {
	if spec == nil {
		return nil
	}
	if spec.AvailabilityStatus != "" || len(spec.AvailableZones) > 0 || len(spec.UnavailableZones) > 0 {
		return spec
	}
	trimmedId := strings.TrimSpace(specId)
	for _, ms := range modelSpecs {
		if ms.Id == trimmedId || (ms.CspSpecName != "" && (ms.CspSpecName == spec.CspSpecName || ms.CspSpecName == trimmedId)) {
			merged := *spec
			merged.AvailabilityStatus = ms.AvailabilityStatus
			merged.AvailabilityReason = ms.AvailabilityReason
			merged.AvailableZones = slices.Clone(ms.AvailableZones)
			merged.UnavailableZones = slices.Clone(ms.UnavailableZones)
			return &merged
		}
	}
	return spec
}

// subnetZoneRef records the resolved zone and JSON path of a target subnet.
type subnetZoneRef struct {
	name string
	zone string
	path string
}

// validateSubnetZoneAlignment verifies zonal subnet availability and regional subnet zone pinning rules.
func validateSubnetZoneAlignment(
	ctx context.Context,
	nsId string,
	target *cloudmodel.RecommendedInfra,
	resolvedSpecs []*cloudmodel.SpecInfo,
	useExisting bool,
) []ValidationIssue {
	var issues []ValidationIssue

	modelCSP := extractProviderName(target.TargetCloud.Csp, target.TargetVNet.ConnectionName)
	if modelCSP == "" && len(target.TargetInfra.NodeGroups) > 0 {
		modelCSP = extractProviderName("", target.TargetInfra.NodeGroups[0].ConnectionName)
	}

	if recommendation.IsRegionalSubnetWarningCSP(modelCSP) {
		for idx, sub := range target.TargetVNet.SubnetInfoList {
			if zone := strings.TrimSpace(sub.Zone); zone != "" {
				issues = append(issues, ValidationIssue{
					Code:     CodeRegionalSubnetZonePinned,
					Severity: SeverityWarning,
					Path:     fmt.Sprintf("targetVNet.subnetInfoList[%d].zone", idx),
					Message: fmt.Sprintf("subnet '%s' pins zone '%s' on regional-subnet CSP '%s', which may cause over-constrained allocation failures",
						sub.Name, zone, modelCSP),
				})
			}
		}
	}

	reqSubnets := make(map[string]subnetZoneRef, len(target.TargetVNet.SubnetInfoList))
	for idx, sub := range target.TargetVNet.SubnetInfoList {
		if sub.Name != "" {
			reqSubnets[sub.Name] = subnetZoneRef{
				name: sub.Name,
				zone: strings.TrimSpace(sub.Zone),
				path: fmt.Sprintf("targetVNet.subnetInfoList[%d].zone", idx),
			}
		}
	}

	existingVNets := make(map[string]*tbmodel.VNetInfo)
	resolveSubnet := func(vNetId, subnetId, ngPath string) (subnetZoneRef, bool) {
		if useExisting && vNetId != "" {
			vNetInfo, cached := existingVNets[vNetId]
			if !cached {
				if fetched, err := tbclient.NewSession().SetContext(ctx).ReadVNet(nsId, vNetId); err == nil && fetched.Id != "" {
					vNetInfo = &fetched
				}
				existingVNets[vNetId] = vNetInfo
			}
			if vNetInfo != nil {
				for _, sub := range vNetInfo.SubnetInfoList {
					if sub.Name == subnetId {
						return subnetZoneRef{
							name: sub.Name,
							zone: strings.TrimSpace(sub.Zone),
							path: ngPath + ".subnetId",
						}, true
					}
				}
			}
		}
		if ref, ok := reqSubnets[subnetId]; ok {
			return ref, true
		}
		return subnetZoneRef{}, false
	}

	type sharedUnpinnedGroup struct {
		ref   subnetZoneRef
		specs []*cloudmodel.SpecInfo
	}
	unpinnedGroups := make(map[string]*sharedUnpinnedGroup)

	for i, ng := range target.TargetInfra.NodeGroups {
		ngPath := fmt.Sprintf("targetInfra.nodeGroups[%d]", i)
		ngCSP := extractProviderName("", ng.ConnectionName)
		if ngCSP == "" {
			ngCSP = modelCSP
		}
		if !recommendation.IsZonalSubnetCSP(ngCSP) {
			continue
		}

		spec := resolvedSpecs[i]
		if spec == nil {
			continue
		}

		switch spec.AvailabilityStatus {
		case cloudmodel.AvailabilityUnsupported, cloudmodel.AvailabilityQuotaExceeded, cloudmodel.AvailabilitySoldOut:
			issues = append(issues, ValidationIssue{
				Code:     CodeSubnetZoneUnavailable,
				Severity: SeverityError,
				Path:     ngPath + ".specId",
				Message:  fmt.Sprintf("VM spec '%s' in nodegroup '%s' is not available (%s)", ng.SpecId, ng.Name, spec.AvailabilityStatus),
			})
			continue
		}

		subRef, found := resolveSubnet(ng.VNetId, ng.SubnetId, ngPath)
		if !found {
			continue
		}

		availZones := normalizeZones(spec.AvailableZones)
		unavailZones := normalizeZones(spec.UnavailableZones)

		if subRef.zone == "" {
			if len(unavailZones) > 0 {
				issues = append(issues, ValidationIssue{
					Code:     CodeSubnetZoneUnavailable,
					Severity: SeverityError,
					Path:     subRef.path,
					Message: fmt.Sprintf("subnet '%s' must specify an available zone for spec '%s' in nodegroup '%s' on zonal CSP '%s' (unavailable zones: %s)",
						subRef.name, ng.SpecId, ng.Name, ngCSP, strings.Join(unavailZones, ", ")),
				})
				continue
			}
			key := ng.VNetId + "/" + subRef.name
			grp, exists := unpinnedGroups[key]
			if !exists {
				grp = &sharedUnpinnedGroup{ref: subRef}
				unpinnedGroups[key] = grp
			}
			grp.specs = append(grp.specs, spec)
			continue
		}

		if slices.Contains(unavailZones, subRef.zone) || (len(availZones) > 0 && !slices.Contains(availZones, subRef.zone)) {
			issues = append(issues, ValidationIssue{
				Code:     CodeSubnetZoneUnavailable,
				Severity: SeverityError,
				Path:     subRef.path,
				Message: fmt.Sprintf("subnet '%s' zone '%s' is not available for spec '%s' in nodegroup '%s' (available zones: %v)",
					subRef.name, subRef.zone, ng.SpecId, ng.Name, availZones),
			})
		}
	}

	for _, grp := range unpinnedGroups {
		if len(grp.specs) < 2 {
			continue
		}
		if hasDisjointAvailableZones(grp.specs) {
			issues = append(issues, ValidationIssue{
				Code:     CodeSubnetZoneUnavailable,
				Severity: SeverityError,
				Path:     grp.ref.path,
				Message: fmt.Sprintf("nodegroups sharing subnet '%s' have disjoint available zones and cannot be placed in a single zonal subnet",
					grp.ref.name),
			})
		}
	}

	return issues
}

// extractProviderName resolves a lowercase CSP name from an explicit csp field or connectionName prefix.
func extractProviderName(cspField, connectionName string) string {
	if trimmed := strings.TrimSpace(cspField); trimmed != "" {
		return strings.ToLower(trimmed)
	}
	parts := strings.SplitN(strings.TrimSpace(connectionName), "-", 2)
	if len(parts) == 2 && parts[0] != "" {
		return strings.ToLower(parts[0])
	}
	return ""
}

// normalizeZones returns deduplicated trimmed non-empty zone names.
func normalizeZones(zones []string) []string {
	if len(zones) == 0 {
		return nil
	}
	seen := make(map[string]bool, len(zones))
	out := make([]string, 0, len(zones))
	for _, raw := range zones {
		z := strings.TrimSpace(raw)
		if z == "" || seen[z] {
			continue
		}
		seen[z] = true
		out = append(out, z)
	}
	return out
}

// hasDisjointAvailableZones reports whether two or more specs have non-empty AvailableZones with an empty intersection.
func hasDisjointAvailableZones(specs []*cloudmodel.SpecInfo) bool {
	var common []string
	constrainedCount := 0
	for _, spec := range specs {
		avail := normalizeZones(spec.AvailableZones)
		if len(avail) == 0 {
			continue
		}
		constrainedCount++
		if constrainedCount == 1 {
			common = avail
			continue
		}
		next := make([]string, 0, len(common))
		for _, z := range common {
			if slices.Contains(avail, z) {
				next = append(next, z)
			}
		}
		common = next
	}
	return constrainedCount >= 2 && len(common) == 0
}

func ptr(issue ValidationIssue) *ValidationIssue {
	return &issue
}
