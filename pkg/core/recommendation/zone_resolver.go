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

package recommendation

import (
	"slices"
	"sort"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/csp"
	"github.com/rs/zerolog/log"
)

// IsZonalSubnetCSP reports whether the provider binds each subnet to a single availability zone.
func IsZonalSubnetCSP(cspName string) bool {
	switch strings.ToLower(strings.TrimSpace(cspName)) {
	case csp.AWS, csp.Alibaba, csp.Tencent, csp.IBM, csp.NCP, "ncpvpc", csp.KT, "ktvpc":
		return true
	default:
		return false
	}
}

// IsRegionalSubnetWarningCSP reports whether pinning Subnet.Zone on the provider risks allocation failure.
func IsRegionalSubnetWarningCSP(cspName string) bool {
	switch strings.ToLower(strings.TrimSpace(cspName)) {
	case csp.Azure, csp.GCP:
		return true
	default:
		return false
	}
}

// isSpecProvisionable checks whether runtime availability metadata allows recommending a spec.
func isSpecProvisionable(isZonal bool, status string, availableZones, unavailableZones []string) bool {
	switch strings.TrimSpace(status) {
	case string(cloudmodel.AvailabilityUnsupported),
		string(cloudmodel.AvailabilityQuotaExceeded),
		string(cloudmodel.AvailabilitySoldOut):
		return false
	}
	if isZonal && len(unavailableZones) > 0 && len(availableZones) == 0 {
		return false
	}
	return true
}

// filterAvailableTbSpecs excludes Tumblebug specs that are unsupported, quota-exceeded, or sold out.
func filterAvailableTbSpecs(cspName string, specs []tbmodel.SpecInfo) []tbmodel.SpecInfo {
	if len(specs) == 0 {
		return specs
	}
	isZonal := IsZonalSubnetCSP(cspName)
	filtered := make([]tbmodel.SpecInfo, 0, len(specs))
	for _, spec := range specs {
		if !isSpecProvisionable(isZonal, string(spec.AvailabilityStatus), spec.AvailableZones, spec.UnavailableZones) {
			log.Debug().
				Str("csp", cspName).
				Str("specId", spec.Id).
				Str("availabilityStatus", string(spec.AvailabilityStatus)).
				Msg("Excluded unprovisionable spec from recommendation candidates")
			continue
		}
		filtered = append(filtered, spec)
	}
	return filtered
}

// filterAvailableSpecs excludes converted cloudmodel specs that are unsupported, quota-exceeded, or sold out.
func filterAvailableSpecs(cspName string, specs []cloudmodel.SpecInfo) []cloudmodel.SpecInfo {
	if len(specs) == 0 {
		return specs
	}
	isZonal := IsZonalSubnetCSP(cspName)
	filtered := make([]cloudmodel.SpecInfo, 0, len(specs))
	for _, spec := range specs {
		if !isSpecProvisionable(isZonal, string(spec.AvailabilityStatus), spec.AvailableZones, spec.UnavailableZones) {
			continue
		}
		filtered = append(filtered, spec)
	}
	return filtered
}

// availabilityPriority ranks verified Available specs ahead of Unverified specs.
func availabilityPriority(spec cloudmodel.SpecInfo) int {
	switch spec.AvailabilityStatus {
	case cloudmodel.AvailabilityAvailable:
		return 0
	case "", cloudmodel.AvailabilityUnverified:
		return 1
	default:
		return 2
	}
}

// alignCandidateSubnetsWithSpecZones assigns a valid common AvailableZone to the default subnet on zonal CSPs.
func alignCandidateSubnetsWithSpecZones(candidate *cloudmodel.RecommendedInfra) {
	if candidate == nil || len(candidate.TargetVNet.SubnetInfoList) == 0 {
		return
	}

	// Clone SubnetInfoList so candidate-specific zone updates never mutate shared skeleton slices.
	clonedSubnets := make([]cloudmodel.SubnetReq, len(candidate.TargetVNet.SubnetInfoList))
	copy(clonedSubnets, candidate.TargetVNet.SubnetInfoList)
	candidate.TargetVNet.SubnetInfoList = clonedSubnets

	cspName := resolveCandidateCSP(*candidate)
	if !IsZonalSubnetCSP(cspName) {
		for i := range candidate.TargetVNet.SubnetInfoList {
			candidate.TargetVNet.SubnetInfoList[i].Zone = ""
		}
		return
	}

	commonZones, hasZoneConstraint := intersectNodeGroupAvailableZones(
		candidate.TargetInfra.NodeGroups,
		candidate.TargetSpecList,
	)
	if !hasZoneConstraint {
		return
	}
	if len(commonZones) == 0 {
		candidate.TargetVNet.SubnetInfoList[0].Zone = ""
		log.Warn().
			Str("csp", cspName).
			Str("infraName", candidate.TargetInfra.Name).
			Msg("NodeGroups have disjoint AvailableZones; leaving single subnet zone unset for validation")
		return
	}

	candidate.TargetVNet.SubnetInfoList[0].Zone = commonZones[0]
}

// resolveCandidateCSP extracts the target CSP identifier from CloudProperty or ConnectionName.
func resolveCandidateCSP(candidate cloudmodel.RecommendedInfra) string {
	if trimmed := strings.TrimSpace(candidate.TargetCloud.Csp); trimmed != "" {
		return strings.ToLower(trimmed)
	}
	if parts := strings.SplitN(strings.TrimSpace(candidate.TargetVNet.ConnectionName), "-", 2); len(parts) == 2 && parts[0] != "" {
		return strings.ToLower(parts[0])
	}
	for _, ng := range candidate.TargetInfra.NodeGroups {
		if parts := strings.SplitN(strings.TrimSpace(ng.ConnectionName), "-", 2); len(parts) == 2 && parts[0] != "" {
			return strings.ToLower(parts[0])
		}
	}
	return ""
}

// intersectNodeGroupAvailableZones computes the sorted common AvailableZones across all NodeGroups.
func intersectNodeGroupAvailableZones(nodeGroups []cloudmodel.CreateNodeGroupReq, specs []cloudmodel.SpecInfo) ([]string, bool) {
	var common []string
	initialized := false
	unavailableSet := make(map[string]bool)

	for _, ng := range nodeGroups {
		spec, ok := findSpecForNodeGroup(ng.SpecId, specs)
		if !ok {
			continue
		}
		for _, uz := range spec.UnavailableZones {
			if z := strings.TrimSpace(uz); z != "" {
				unavailableSet[z] = true
			}
		}
		avail := normalizeZoneList(spec.AvailableZones)
		if len(avail) == 0 {
			continue
		}
		if !initialized {
			common = avail
			initialized = true
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

	if !initialized {
		return nil, false
	}

	filtered := make([]string, 0, len(common))
	for _, z := range common {
		if !unavailableSet[z] {
			filtered = append(filtered, z)
		}
	}
	sort.Strings(filtered)
	return filtered, true
}

// findSpecForNodeGroup matches a NodeGroup SpecId against SpecInfo.Id or SpecInfo.CspSpecName.
func findSpecForNodeGroup(specId string, specs []cloudmodel.SpecInfo) (cloudmodel.SpecInfo, bool) {
	trimmed := strings.TrimSpace(specId)
	if trimmed == "" {
		return cloudmodel.SpecInfo{}, false
	}
	for _, s := range specs {
		if s.Id == trimmed || (s.CspSpecName != "" && s.CspSpecName == trimmed) {
			return s, true
		}
	}
	return cloudmodel.SpecInfo{}, false
}

// normalizeZoneList returns deduplicated non-empty trimmed zone names.
func normalizeZoneList(zones []string) []string {
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
