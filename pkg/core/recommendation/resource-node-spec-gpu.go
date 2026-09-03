package recommendation

import (
	"fmt"
	"math"
	"sort"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/rs/zerolog/log"
)

// hasGpu returns true if the node property contains at least one physical GPU accelerator.
func hasGpu(node onpremmodel.NodeProperty) bool {
	return node.GPU != nil && node.GPU.Count > 0
}

// buildGpuDeploymentPlan constructs the deployment plan JSON for GPU accelerator node spec recommendation.
// It searches for specs with at least the requested GPU count and per-device VRAM, within host vCPU/RAM ranges.
func buildGpuDeploymentPlan(
	node onpremmodel.NodeProperty,
	csp string,
	region string,
	architecture string,
	vcpusCalculated uint32,
	memory uint32,
	rangeWeight int,
	limit int,
) (string, uint32, uint32, uint32, uint32) {
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
					"condition": [{"operand": "%d", "operator": ">="}],
					"metric": "acceleratorCount"
				},
				{
					"condition": [{"operand": "%.1f", "operator": ">="}],
					"metric": "acceleratorMemoryGB"
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

	vcpusMin, vcpusMax, memoryMin, memoryMax := calculateOptimalRange(vcpusCalculated, memory, rangeWeight)

	targetCount := uint32(1)
	if node.GPU != nil && node.GPU.Count > 0 {
		targetCount = node.GPU.Count
	}

	targetVramPerGpu := float32(0)
	if node.GPU != nil {
		if node.GPU.TotalMemoryGB > 0 && node.GPU.Count > 0 {
			targetVramPerGpu = node.GPU.TotalMemoryGB / float32(node.GPU.Count)
		} else if len(node.GPU.Details) > 0 && node.GPU.Details[0].MemoryTotal > 0 {
			targetVramPerGpu = node.GPU.Details[0].MemoryTotal
		}
	}

	providerName := strings.ToLower(csp)
	regionName := strings.ToLower(region)

	plan := fmt.Sprintf(planTemplate,
		vcpusMin, vcpusMax,
		memoryMin, memoryMax,
		targetCount,
		targetVramPerGpu,
		providerName, regionName, architecture,
		limit,
	)

	return plan, vcpusMin, vcpusMax, memoryMin, memoryMax
}

// gpuRankingContext carries the evaluation metrics for multi-dimensional GPU ranking.
type gpuRankingContext struct {
	targetCount  uint8
	targetVram   float32
	targetVendor string
	vcpus        uint32
	memory       uint32
	csp          string
}

// gpuCriterion defines comparison between two GPU specs on a specific ranking dimension.
type gpuCriterion func(ctx gpuRankingContext, a, b cloudmodel.SpecInfo) int

// gpuCountProximity gives highest priority to specs with the exact GPU count (e.g., 2 GPUs -> 2 GPUs).
// If neither matches exactly, the spec with fewer surplus GPUs ranks first to avoid wasteful overprovisioning.
func gpuCountProximity(ctx gpuRankingContext, a, b cloudmodel.SpecInfo) int {
	exactA := a.AcceleratorCount == ctx.targetCount
	exactB := b.AcceleratorCount == ctx.targetCount

	if exactA && !exactB {
		return -1
	}
	if !exactA && exactB {
		return 1
	}

	// If both match or both don't match, rank closer count first
	diffA := abs(int32(a.AcceleratorCount) - int32(ctx.targetCount))
	diffB := abs(int32(b.AcceleratorCount) - int32(ctx.targetCount))
	return int(diffA - diffB)
}

// gpuVramProximity prioritizes the spec with VRAM closest to the requested per-GPU capacity (minimum surplus first).
func gpuVramProximity(ctx gpuRankingContext, a, b cloudmodel.SpecInfo) int {
	if ctx.targetVram <= 0 {
		return 0
	}

	diffA := math.Abs(float64(a.AcceleratorMemoryGB - ctx.targetVram))
	diffB := math.Abs(float64(b.AcceleratorMemoryGB - ctx.targetVram))

	const epsilon = 0.01
	if math.Abs(diffA-diffB) < epsilon {
		return 0
	}
	if diffA < diffB {
		return -1
	}
	return 1
}

// gpuVendorMatch ranks matching GPU vendor/family (e.g. NVIDIA) first.
func gpuVendorMatch(ctx gpuRankingContext, a, b cloudmodel.SpecInfo) int {
	if ctx.targetVendor == "" {
		return 0
	}

	modelA := strings.ToLower(a.AcceleratorModel)
	modelB := strings.ToLower(b.AcceleratorModel)
	vendor := strings.ToLower(ctx.targetVendor)

	matchA := strings.Contains(modelA, vendor)
	matchB := strings.Contains(modelB, vendor)

	if matchA && !matchB {
		return -1
	}
	if !matchA && matchB {
		return 1
	}
	return 0
}

// gpuHostResourceProximity ranks by combined vCPU and host memory distance (L1 Manhattan norm).
func gpuHostResourceProximity(ctx gpuRankingContext, a, b cloudmodel.SpecInfo) int {
	da := abs(int32(a.VCPU)-int32(ctx.vcpus)) + abs(int32(a.MemoryGiB)-int32(ctx.memory))
	db := abs(int32(b.VCPU)-int32(ctx.vcpus)) + abs(int32(b.MemoryGiB)-int32(ctx.memory))
	return int(da - db)
}

// sortGpuByProximityWithCost sorts GPU VM specs using a multi-dimensional ranking hierarchy:
// 1. GPU Vendor Match (e.g., NVIDIA -> NVIDIA: absolute prerequisite for driver/CUDA binary compatibility)
// 2. GPU Count Proximity (Exact count match first, preventing architecture & parallelism mismatch)
// 3. GPU VRAM Proximity (Closest per-device VRAM to avoid OOM or excessive overprovisioning)
// 4. Host Resource Distance (vCPU + Memory Manhattan distance)
// 5. CostPerHour (Lowest cost as final tie-breaker)
func sortGpuByProximityWithCost(vmSpecs []cloudmodel.SpecInfo, node onpremmodel.NodeProperty, csp string) {
	if len(vmSpecs) == 0 {
		return
	}

	targetCount := uint8(1)
	if node.GPU != nil && node.GPU.Count > 0 {
		targetCount = uint8(node.GPU.Count)
	}

	targetVramPerGpu := float32(0)
	targetVendor := ""
	if node.GPU != nil {
		targetVendor = node.GPU.Vendor
		if node.GPU.TotalMemoryGB > 0 && node.GPU.Count > 0 {
			targetVramPerGpu = node.GPU.TotalMemoryGB / float32(node.GPU.Count)
		} else if len(node.GPU.Details) > 0 && node.GPU.Details[0].MemoryTotal > 0 {
			targetVramPerGpu = node.GPU.Details[0].MemoryTotal
		}
	}

	// Calculate host vCPUs
	cpus := node.CPU.Cpus
	threads := node.CPU.Threads
	if threads == 0 {
		threads = 1
	}
	vcpusCalculated := uint32(cpus * threads)
	memory := uint32(node.Memory.TotalSize)

	ctx := gpuRankingContext{
		targetCount:  targetCount,
		targetVram:   targetVramPerGpu,
		targetVendor: targetVendor,
		vcpus:        vcpusCalculated,
		memory:       memory,
		csp:          csp,
	}

	log.Debug().
		Uint8("targetGpuCount", targetCount).
		Float32("targetVramPerGpu", targetVramPerGpu).
		Str("targetVendor", targetVendor).
		Int("specsToSort", len(vmSpecs)).
		Msg("Sorting GPU VM specs with multi-dimensional proximity ranking (Vendor > Count > VRAM > Host > Cost)")

	sort.Slice(vmSpecs, func(i, j int) bool {
		criteria := []gpuCriterion{
			gpuVendorMatch,
			gpuCountProximity,
			gpuVramProximity,
			gpuHostResourceProximity,
			func(ctx gpuRankingContext, a, b cloudmodel.SpecInfo) int {
				switch {
				case a.CostPerHour < b.CostPerHour:
					return -1
				case a.CostPerHour > b.CostPerHour:
					return 1
				default:
					return 0
				}
			},
		}

		for _, criterion := range criteria {
			if d := criterion(ctx, vmSpecs[i], vmSpecs[j]); d != 0 {
				return d < 0
			}
		}
		return false
	})
}
