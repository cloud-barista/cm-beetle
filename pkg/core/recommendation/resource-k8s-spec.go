/*
Copyright 2024 The Cloud-Barista Authors.
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
	"fmt"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	"github.com/rs/zerolog/log"
)

// K8s minimum node requirements, matching CB-Tumblebug's validateK8sMinimumRequirements.
// Shared across RecommendK8sSpecs and applyK8sMinimums.
const (
	minK8sVCPU   = uint32(2)
	minK8sMemGiB = uint32(4)
)

// RecommendK8sSpecs returns a sorted list of K8s-compatible node specs for the given node.
// K8s minimum requirements (vCPU >= 2, memoryGiB >= 4) are applied before querying.
// Uses the same POST /tumblebug/recommendSpec endpoint as VM spec recommendation.
// NCP KVM hypervisor filter from RecommendVmSpecs is intentionally excluded for Managed K8s.
func RecommendK8sSpecs(csp, region string, node onpremmodel.NodeProperty, limit int) ([]cloudmodel.SpecInfo, int, error) {
	const defaultArchitecture = "x86_64"

	emptyResp := []cloudmodel.SpecInfo{}

	if limit <= 0 {
		limit = defaultSpecsLimit
	}

	// Apply K8s minimum requirements before building FilterPolicy
	node = applyK8sMinimums(node)

	threads := node.CPU.Threads
	if threads == 0 {
		threads = 1
	}
	vcpusCalculated := node.CPU.Cpus * threads
	memory := uint32(node.Memory.TotalSize)

	providerName := strings.ToLower(csp)
	regionName := strings.ToLower(region)

	architecture := node.CPU.Architecture
	if architecture == "" || architecture == "amd64" {
		architecture = defaultArchitecture
	}

	// Same FilterPolicy template as VM spec recommendation
	const planTemplate = `{
		"filter": {
			"policy": [
				{
					"condition": [
						{"operand": "%d", "operator": ">="},
						{"operand": "%d", "operator": "<="}
					],
					"metric": "vCPU"
				},
				{
					"condition": [
						{"operand": "%d", "operator": ">="},
						{"operand": "%d", "operator": "<="}
					],
					"metric": "memoryGiB"
				},
				{
					"condition": [{"operand": "%s"}],
					"metric": "providerName"
				},
				{
					"condition": [{"operand": "%s"}],
					"metric": "regionName"
				},
				{
					"condition": [{"operand": "%s"}],
					"metric": "architecture"
				}
			]
		},
		"limit": %d,
		"priority": {
			"policy": [{"metric": "cost"}]
		}
	}`

	const (
		initialRangeWeight = 1
		maxRangeWeight     = 5
		// K8s minimums are enforced in FilterPolicy below; see package-level minK8sVCPU / minK8sMemGiB.
	)

	var specInfoList []tbmodel.SpecInfo

	for rangeWeight := initialRangeWeight; rangeWeight <= maxRangeWeight; rangeWeight++ {
		vcpusMin, vcpusMax, memoryMin, memoryMax := calculateOptimalRange(vcpusCalculated, memory, rangeWeight)

		// Enforce K8s minimums in the FilterPolicy before sending to TB.
		if vcpusMin < minK8sVCPU {
			vcpusMin = minK8sVCPU
		}
		if memoryMin < minK8sMemGiB {
			memoryMin = minK8sMemGiB
		}

		plan := fmt.Sprintf(planTemplate,
			vcpusMin, vcpusMax,
			memoryMin, memoryMax,
			providerName, regionName, architecture,
			limit,
		)

		log.Debug().
			Str("hostname", node.Hostname).
			Int("rangeWeight", rangeWeight).
			Uint32("vcpu", vcpusCalculated).
			Uint32("memoryGiB", memory).
			Uint32("filterVcpuMin", vcpusMin).
			Uint32("filterMemoryMin", memoryMin).
			Str("provider", providerName).
			Str("region", regionName).
			Msgf("Querying K8s specs (attempt %d/%d)", rangeWeight, maxRangeWeight)

		var err error
		specInfoList, err = tbclient.NewSession().InfraRecommendSpec(plan)
		if err != nil {
			return emptyResp, -1, fmt.Errorf("failed to get K8s spec recommendations: %w", err)
		}

		// Filter specs with valid (non-negative) cost only.
		// K8s minimums are already enforced in the FilterPolicy above.
		valid := make([]tbmodel.SpecInfo, 0, len(specInfoList))
		for _, spec := range specInfoList {
			if spec.CostPerHour >= 0 {
				valid = append(valid, spec)
			}
		}
		specInfoList = valid

		if len(specInfoList) > 0 {
			break
		}

		log.Warn().
			Str("hostname", node.Hostname).
			Int("rangeWeight", rangeWeight).
			Msg("No K8s specs found; retrying with wider range")
	}

	if len(specInfoList) == 0 {
		return emptyResp, -1, fmt.Errorf("no K8s specs found for node (hostname: %s, vcpu: %d, memory: %d GiB, csp: %s, region: %s)",
			node.Hostname, vcpusCalculated, memory, csp, region)
	}

	if limit < len(specInfoList) {
		specInfoList = specInfoList[:limit]
	}

	converted, err := modelconv.ConvertWithValidation[[]tbmodel.SpecInfo, []cloudmodel.SpecInfo](specInfoList)
	if err != nil {
		return emptyResp, -1, fmt.Errorf("failed to convert K8s spec list: %w", err)
	}

	sortByProximityWithCost(vcpusCalculated, memory, converted)

	log.Info().
		Str("hostname", node.Hostname).
		Int("count", len(converted)).
		Msg("K8s spec recommendation completed")

	return converted, len(converted), nil
}
