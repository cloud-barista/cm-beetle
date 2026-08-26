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
	"strconv"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	"github.com/rs/zerolog/log"
)

// maxK8sSpecRangeWeight bounds the widening retry loop, mirroring RecommendVmSpecs.
const maxK8sSpecRangeWeight = 5

// RecommendK8sNodeGroupSpecs recommends a worker spec for every node group derived from the
// source workers, and assembles the L0 response.
//
// Workers are recommended individually and then merged by the target spec they resolve to, because
// managed K8s node groups are homogeneous; one recommendation set is produced per merged group,
// with the group's members reported in SourceServers.
//
// Unlike RecommendK8sInfra, this does NOT require srcInfra.K8sCluster — spec recommendation only
// needs worker nodes. Cluster info is a version-selection input, not a spec input.
func RecommendK8sNodeGroupSpecs(provider, region string, srcInfra onpremmodel.OnpremInfra, limit int) (cloudmodel.RecommendedSpecList, error) {

	ret := cloudmodel.RecommendedSpecList{}

	workers := collectWorkerNodes(srcInfra.Nodes)
	if len(workers) == 0 {
		return ret, fmt.Errorf("no worker nodes found in source infra (nodes with role=worker are required)")
	}

	specsLimit := limit
	if specsLimit <= 0 {
		specsLimit = GetDefaultSpecsLimit()
	}

	// Share the resolve-then-merge pipeline with RecommendK8sInfra so both APIs report the same
	// node group breakdown for the same input. They used to derive groups independently, which let
	// them disagree.
	targets := resolveWorkerTargets(provider, region, workers, specsLimit)
	groups, failed := mergeIntoNodeGroups(targets)
	logExcludedWorkers(failed)

	for i, g := range groups {
		members := machineIdsOf(g.nodes)

		desc := fmt.Sprintf("Recommended worker spec for node group %d (%d node(s))", i+1, len(g.nodes))
		if len(g.notes) > 0 {
			desc += " " + strings.Join(g.notes, " ")
		}

		// A node group's candidate list can overlap another's; merge their source servers rather
		// than emitting duplicate entries. Same dedup rule as the VM spec controller.
		for _, spec := range g.specs {
			if idx := indexOfRecommendedSpec(ret.RecommendedSpecList, spec.Id); idx >= 0 {
				ret.RecommendedSpecList[idx].SourceServers = append(ret.RecommendedSpecList[idx].SourceServers, members...)
				continue
			}
			ret.RecommendedSpecList = append(ret.RecommendedSpecList, cloudmodel.RecommendedSpec{
				Status:        string(FullyRecommended),
				SourceServers: members,
				Description:   desc,
				TargetSpec:    spec,
			})
		}
	}

	// Report excluded workers grouped by cause, so one shared reason is stated once.
	for _, f := range groupFailuresByCause(failed) {
		ret.RecommendedSpecList = append(ret.RecommendedSpecList, cloudmodel.RecommendedSpec{
			Status:        string(NothingRecommended),
			SourceServers: f.machineIds,
			Description:   fmt.Sprintf("failed to recommend worker specs: %s", f.cause),
			TargetSpec:    cloudmodel.SpecInfo{},
		})
	}

	ret.Count = len(ret.RecommendedSpecList)
	ret.Status, ret.Description = summarizeRecommendedSpecList(ret.RecommendedSpecList)

	log.Info().
		Str("provider", provider).Str("region", region).
		Int("workerCount", len(workers)).Int("specCount", ret.Count).
		Msg("K8s node group spec recommendation completed")

	return ret, nil
}

// RecommendK8sNodeSpecs recommends cost-ranked specs for a single K8s worker node.
//
// It differs from RecommendVmSpecs in three ways, all driven by the same constraint — a worker
// smaller than its source leaves pods unschedulable:
//   - the search range is clamped so it never dips below the source (nor below the minimum viable
//     worker spec), whereas the VM path deliberately allows downsizing,
//   - ranking ignores CPU vendor (see sortK8sSpecsByProximity): with every candidate already at or
//     above the source, a vendor match can only buy a larger, costlier node,
//   - NCP hypervisor filtering is skipped (NKS is pinned to XEN, so the VM path's KVM filter would
//     exclude the very specs NKS can use),
//   - an upscale note is returned when the floor raised the request, so callers can surface it.
//
// The request goes to Tumblebug's K8s endpoint, whose validateK8sMinimumRequirements enforces the
// K8s node minimums; the clamp below keeps the lower bounds at or above those minimums so the
// request passes validation instead of being rejected.
func RecommendK8sNodeSpecs(provider, region string, worker onpremmodel.NodeProperty, limit int) ([]cloudmodel.SpecInfo, string, error) {

	if limit <= 0 {
		limit = GetDefaultSpecsLimit()
	}

	// Source sizing uses the shared sourceVcpuOf rule, so the value driving this search is the
	// same one that keys the recommendation cache. Deriving it twice, differently, is how a node
	// could previously be keyed by one vCPU count and sized by another.
	srcVcpu := sourceVcpuOf(worker)
	srcMem := uint32(worker.Memory.TotalSize)
	arch := normalizeArch(worker.CPU.Architecture)

	// Floor at the minimum viable worker spec, and report the change so it is not silent.
	minVcpu, minMem := srcVcpu, srcMem
	if minVcpu < minViableWorkerVcpu {
		minVcpu = minViableWorkerVcpu
	}
	if uint64(minMem) < minViableWorkerMemGiB {
		minMem = uint32(minViableWorkerMemGiB)
	}

	upscaleNote := ""
	if minVcpu != srcVcpu || minMem != srcMem {
		upscaleNote = fmt.Sprintf(
			"(worker spec upscaled from source %dvCPU/%dGiB to %dvCPU/%dGiB — the minimum node size "+
				"accepted by the target K8s node recommendation. A node this small leaves little allocatable "+
				"capacity after kubelet, kube-proxy, CNI, and system pods take their reserved share, so pods "+
				"may fail to schedule at the source size.)",
			srcVcpu, srcMem, minVcpu, minMem)
		log.Warn().
			Uint32("sourceVcpu", srcVcpu).Uint32("sourceMemGiB", srcMem).
			Uint32("minVcpu", minVcpu).Uint32("minMemGiB", minMem).
			Msg("Source worker below minimum viable K8s worker spec; upscaling recommendation")
	}

	for rangeWeight := 1; rangeWeight <= maxK8sSpecRangeWeight; rangeWeight++ {

		// Reuse the VM range model (workload-type aware), then clamp the lower bound. The clamp is
		// the crux: calculateComputeIntensiveRange returns memoryRangeMin = 0, which for a K8s
		// worker would admit specs with less memory than the source.
		vcpuMin, vcpuMax, memMin, memMax := calculateOptimalRange(minVcpu, minMem, rangeWeight)
		if vcpuMin < minVcpu {
			vcpuMin = minVcpu
		}
		if memMin < minMem {
			memMin = minMem
		}
		if vcpuMax < vcpuMin {
			vcpuMax = vcpuMin
		}
		if memMax < memMin {
			memMax = memMin
		}

		req := buildK8sSpecSearchReq(k8sSpecSearchBounds{
			vcpuMin: vcpuMin, vcpuMax: vcpuMax,
			memMin: memMin, memMax: memMax,
			provider: provider, region: region, arch: arch, limit: limit,
		})

		log.Debug().
			Uint32("vcpuMin", vcpuMin).Uint32("vcpuMax", vcpuMax).
			Uint32("memMin", memMin).Uint32("memMax", memMax).
			Str("arch", arch).Int("rangeWeight", rangeWeight).
			Msg("Searching K8s worker specs")

		found, err := tbclient.NewSession().K8sClusterRecommendNode(req)
		if err != nil {
			return nil, upscaleNote, fmt.Errorf("K8s worker spec search failed: %w", err)
		}

		// Drop unpriced specs only when pricing data exists at all, so regions without loaded
		// cost data still yield recommendations.
		found = filterPricedSpecs(found)

		// Note: the VM path's NCP KVM filter is intentionally NOT applied here.

		if len(found) > 0 {
			converted, err := modelconv.ConvertWithValidation[[]tbmodel.SpecInfo, []cloudmodel.SpecInfo](found)
			if err != nil {
				return nil, upscaleNote, fmt.Errorf("failed to convert K8s worker spec list: %w", err)
			}

			sortK8sSpecsByProximity(converted, minVcpu, minMem)
			if len(converted) > limit {
				converted = converted[:limit]
			}

			log.Info().
				Int("specsFound", len(converted)).Int("rangeWeight", rangeWeight).
				Str("topSpecId", converted[0].Id).
				Msg("K8s worker specs selected")
			return converted, upscaleNote, nil
		}

		log.Warn().
			Int("rangeWeight", rangeWeight).Int("maxRangeWeight", maxK8sSpecRangeWeight).
			Msg("No K8s worker spec found in range, widening and retrying")
	}

	return nil, upscaleNote, fmt.Errorf(
		"no spec found for K8s worker node (vCPU>=%d, memory>=%dGiB) in %s %s after %d attempts",
		minVcpu, minMem, provider, region, maxK8sSpecRangeWeight)
}

// k8sSpecSearchBounds carries the resolved search window for one spec-search request.
type k8sSpecSearchBounds struct {
	vcpuMin, vcpuMax uint32
	memMin, memMax   uint32
	provider, region string
	arch             string
	limit            int
}

// buildK8sSpecSearchReq assembles the Tumblebug spec-search request.
//
// Built as a struct rather than a JSON string template: Operand is a string on the wire even for
// numeric metrics, and exact-match metrics carry no operator, both of which are easy to get wrong
// in a format string and impossible for the compiler to check.
func buildK8sSpecSearchReq(b k8sSpecSearchBounds) tbmodel.RecommendSpecReq {

	rangeCond := func(metric string, min, max uint32) tbmodel.FilterCondition {
		return tbmodel.FilterCondition{
			Metric: metric,
			Condition: []tbmodel.Operation{
				{Operator: ">=", Operand: strconv.FormatUint(uint64(min), 10)},
				{Operator: "<=", Operand: strconv.FormatUint(uint64(max), 10)},
			},
		}
	}
	exactCond := func(metric, value string) tbmodel.FilterCondition {
		return tbmodel.FilterCondition{
			Metric:    metric,
			Condition: []tbmodel.Operation{{Operand: value}},
		}
	}

	return tbmodel.RecommendSpecReq{
		Filter: tbmodel.FilterInfo{Policy: []tbmodel.FilterCondition{
			rangeCond("vCPU", b.vcpuMin, b.vcpuMax),
			rangeCond("memoryGiB", b.memMin, b.memMax),
			exactCond("providerName", strings.ToLower(b.provider)),
			exactCond("regionName", strings.ToLower(b.region)),
			exactCond("architecture", b.arch),
		}},
		Priority: tbmodel.PriorityInfo{Policy: []tbmodel.PriorityCondition{{Metric: "cost"}}},
		Limit:    b.limit,
	}
}

// filterPricedSpecs drops specs with unknown pricing (-1), but only when at least one spec carries
// real pricing. If none do, cost data is simply not loaded for this region and filtering would
// discard every candidate.
func filterPricedSpecs(specs []tbmodel.SpecInfo) []tbmodel.SpecInfo {
	hasPricing := false
	for _, s := range specs {
		if s.CostPerHour >= 0 {
			hasPricing = true
			break
		}
	}
	if !hasPricing {
		return specs
	}

	priced := make([]tbmodel.SpecInfo, 0, len(specs))
	for _, s := range specs {
		if s.CostPerHour >= 0 {
			priced = append(priced, s)
		}
	}
	return priced
}

// machineIdsOf collects the source machine IDs backing a worker group.
func machineIdsOf(nodes []onpremmodel.NodeProperty) []string {
	ids := make([]string, 0, len(nodes))
	for _, n := range nodes {
		ids = append(ids, n.MachineId)
	}
	return ids
}

// indexOfRecommendedSpec returns the index of an entry already recommending specId, or -1.
func indexOfRecommendedSpec(list []cloudmodel.RecommendedSpec, specId string) int {
	for i, entry := range list {
		if entry.TargetSpec.Id == specId {
			return i
		}
	}
	return -1
}

// summarizeRecommendedSpecList derives the overall status and description, using the same rule as
// the VM spec controller: no failures is ok, all failures is none, anything else is partial.
func summarizeRecommendedSpecList(list []cloudmodel.RecommendedSpec) (status, description string) {
	countFailed := 0
	for _, entry := range list {
		if entry.Status == string(NothingRecommended) {
			countFailed++
		}
	}

	switch {
	case len(list) == 0 || countFailed == len(list):
		return string(NothingRecommended), "No K8s worker node specs available"
	case countFailed == 0:
		return string(FullyRecommended), fmt.Sprintf("Recommended %d K8s worker node spec(s)", len(list))
	default:
		return string(PartiallyRecommended), fmt.Sprintf(
			"Recommended %d of %d K8s worker node spec(s)", len(list)-countFailed, len(list))
	}
}
