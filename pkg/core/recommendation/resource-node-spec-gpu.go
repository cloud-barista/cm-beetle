package recommendation

import (
	"fmt"
	"math"
	"sort"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	"github.com/rs/zerolog/log"
)

// Canonical GPU Vendor identities
const (
	GpuVendorNVIDIA = "NVIDIA"
	GpuVendorAMD    = "AMD"
	GpuVendorIntel  = "Intel"
	GpuVendorGoogle = "Google"
	GpuVendorOther  = "Other"
)

// Metric names for Tumblebug deployment plans
const (
	MetricAcceleratorCount    = "acceleratorCount"
	MetricAcceleratorMemoryGB = "acceleratorMemoryGB"
	MetricAcceleratorType     = "acceleratorType"
	MetricAcceleratorModel    = "acceleratorModel"
	AcceleratorTypeGPU        = "gpu"
)

// gpuVendorAliases maps model tokens to canonical GPU vendor names.
// Kept in code matching the established cpuVendorAliases pattern in resource-node-spec.go.
var gpuVendorAliases = []struct {
	token  string
	vendor string
}{
	{"nvidia", GpuVendorNVIDIA},
	{"tesla", GpuVendorNVIDIA},
	{"geforce", GpuVendorNVIDIA},
	{"quadro", GpuVendorNVIDIA},
	{"a100", GpuVendorNVIDIA},
	{"h100", GpuVendorNVIDIA},
	{"b200", GpuVendorNVIDIA},
	{"l40", GpuVendorNVIDIA},
	{"l4", GpuVendorNVIDIA},
	{"a10", GpuVendorNVIDIA},
	{"t4", GpuVendorNVIDIA},
	{"v100", GpuVendorNVIDIA},
	{"k80", GpuVendorNVIDIA},
	{"m60", GpuVendorNVIDIA},
	{"amd", GpuVendorAMD},
	{"radeon", GpuVendorAMD},
	{"instinct", GpuVendorAMD},
	{"mi300", GpuVendorAMD},
	{"mi350", GpuVendorAMD},
	{"mi250", GpuVendorAMD},
	{"mi210", GpuVendorAMD},
	{"mi100", GpuVendorAMD},
	{"v520", GpuVendorAMD},
	{"v620", GpuVendorAMD},
	{"v710", GpuVendorAMD},
	{"intel", GpuVendorIntel},
	{"gaudi", GpuVendorIntel},
	{"xeon phi", GpuVendorIntel},
	{"arc", GpuVendorIntel},
	{"google", GpuVendorGoogle},
	{"tpu", GpuVendorGoogle},
}

// detectGpuVendor canonicalizes arbitrary GPU vendor or model names into a standard vendor identifier.
func detectGpuVendor(raw string) string {
	lower := strings.ToLower(strings.TrimSpace(raw))
	for _, entry := range gpuVendorAliases {
		if strings.Contains(lower, entry.token) {
			return entry.vendor
		}
	}
	return GpuVendorOther
}

// hasGpu returns true if the node property contains at least one physical GPU accelerator card.
func hasGpu(node onpremmodel.NodeProperty) bool {
	return len(node.GPUCards) > 0
}

// GpuCardCluster represents a homogeneous cluster of identical GPU cards installed on the node.
type GpuCardCluster struct {
	Vendor        string
	Model         string
	MemoryTotalGB float32
	Count         uint32
}

// clusterGpuCards groups identical physical GPU cards by (Vendor, Model, MemoryTotalGB).
func clusterGpuCards(cards []onpremmodel.GpuCardProperty) []GpuCardCluster {
	if len(cards) == 0 {
		return nil
	}

	clusters := make([]GpuCardCluster, 0)
	indexMap := make(map[string]int)

	for _, card := range cards {
		key := fmt.Sprintf("%s|%s|%.1f", strings.ToLower(card.Vendor), strings.ToLower(card.Model), card.MemoryTotalGB)
		if idx, found := indexMap[key]; found {
			clusters[idx].Count++
		} else {
			indexMap[key] = len(clusters)
			clusters = append(clusters, GpuCardCluster{
				Vendor:        card.Vendor,
				Model:         card.Model,
				MemoryTotalGB: card.MemoryTotalGB,
				Count:         1,
			})
		}
	}

	// Sort clusters: highest card count first, then highest VRAM
	sort.Slice(clusters, func(i, j int) bool {
		if clusters[i].Count != clusters[j].Count {
			return clusters[i].Count > clusters[j].Count
		}
		return clusters[i].MemoryTotalGB > clusters[j].MemoryTotalGB
	})

	return clusters
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
	targetVramPerGpu := float32(0)

	clusters := clusterGpuCards(node.GPUCards)
	if len(clusters) > 0 {
		primary := clusters[0]
		targetCount = primary.Count
		targetVramPerGpu = primary.MemoryTotalGB
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

// gpuModelTokens lists canonical GPU chip identifiers ordered from specific to general.
var gpuModelTokens = []string{
	"b200", "h200", "h100", "a100", "l40s", "l40", "l4", "a10g", "a10", "t4", "v100", "p100", "k80", "m60",
	"mi350", "mi300", "mi250", "mi210", "mi100", "v710", "v620", "v520", "gaudi3", "gaudi2", "gaudi",
}

// detectGpuModel extracts a canonical GPU chip token from a raw model string.
func detectGpuModel(raw string) string {
	lower := strings.ToLower(strings.TrimSpace(raw))
	for _, token := range gpuModelTokens {
		if strings.Contains(lower, token) {
			return token
		}
	}
	return ""
}

// gpuRankingContext carries the evaluation metrics for multi-dimensional GPU ranking.
type gpuRankingContext struct {
	targetCount  uint8
	targetVram   float32
	targetVendor string
	targetModel  string
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

// gpuVramProximity prioritizes specs meeting target VRAM first, then closest capacity.
func gpuVramProximity(ctx gpuRankingContext, a, b cloudmodel.SpecInfo) int {
	if ctx.targetVram <= 0 {
		return 0
	}

	// Prefer specs that meet or exceed target VRAM (with 0.5 GB tolerance) to prevent LLM OOM.
	suffA := a.AcceleratorMemoryGB+vramTierToleranceGB >= ctx.targetVram
	suffB := b.AcceleratorMemoryGB+vramTierToleranceGB >= ctx.targetVram
	if suffA && !suffB {
		return -1
	}
	if !suffA && suffB {
		return 1
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

	targetCanonical := detectGpuVendor(ctx.targetVendor)
	vendorA := detectGpuVendor(a.AcceleratorModel)
	vendorB := detectGpuVendor(b.AcceleratorModel)

	matchA := (vendorA == targetCanonical && vendorA != GpuVendorOther)
	matchB := (vendorB == targetCanonical && vendorB != GpuVendorOther)

	if matchA && !matchB {
		return -1
	}
	if !matchA && matchB {
		return 1
	}
	return 0
}

// gpuModelMatch ranks specs with the same canonical GPU chip model (e.g. L4 -> L4) first.
func gpuModelMatch(ctx gpuRankingContext, a, b cloudmodel.SpecInfo) int {
	targetCanonical := detectGpuModel(ctx.targetModel)
	if targetCanonical == "" {
		return 0
	}

	modelA := detectGpuModel(a.AcceleratorModel)
	modelB := detectGpuModel(b.AcceleratorModel)

	matchA := (modelA == targetCanonical)
	matchB := (modelB == targetCanonical)

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
// 4. GPU Model Match (Same canonical GPU chip family, e.g., L4 -> L4 over A10G)
// 5. Host Resource Distance (vCPU + Memory Manhattan distance)
// 6. CostPerHour (Lowest cost as final tie-breaker)
func sortGpuByProximityWithCost(vmSpecs []cloudmodel.SpecInfo, node onpremmodel.NodeProperty, csp string) {
	if len(vmSpecs) == 0 {
		return
	}

	targetCount := uint8(1)
	targetVramPerGpu := float32(0)
	targetVendor := ""
	targetModel := ""

	clusters := clusterGpuCards(node.GPUCards)
	if len(clusters) > 0 {
		primary := clusters[0]
		targetCount = uint8(primary.Count)
		targetVramPerGpu = primary.MemoryTotalGB
		targetVendor = primary.Vendor
		targetModel = primary.Model
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
		targetModel:  targetModel,
		vcpus:        vcpusCalculated,
		memory:       memory,
		csp:          csp,
	}

	log.Debug().
		Uint8("targetGpuCount", targetCount).
		Float32("targetVramPerGpu", targetVramPerGpu).
		Str("targetVendor", targetVendor).
		Str("targetModel", targetModel).
		Int("specsToSort", len(vmSpecs)).
		Msg("Sorting GPU VM specs with multi-dimensional proximity ranking (Vendor > Count > VRAM > Model > Host > Cost)")

	sort.Slice(vmSpecs, func(i, j int) bool {
		criteria := []gpuCriterion{
			gpuVendorMatch,
			gpuCountProximity,
			gpuVramProximity,
			gpuModelMatch,
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

// standardVramTiers defines common discrete VRAM capacities in public cloud GPU instances (in GB).
var standardVramTiers = []float32{8, 16, 24, 32, 40, 48, 80, 96, 144, 192, 288}

// vramTierToleranceGB absorbs minor MB-to-GiB conversion rounding (e.g., 24.2 GB -> 24 GB tier).
const vramTierToleranceGB float32 = 0.5

// minGpuFetchLimit ensures large-host GPU specs (e.g., g6.8xlarge) are not truncated by cost-sorted limits.
const minGpuFetchLimit = 100

// rightSizeUpVramTier sizes a requested VRAM up to the nearest standard cloud GPU tier.
func rightSizeUpVramTier(vram float32) float32 {
	for _, tier := range standardVramTiers {
		if tier+vramTierToleranceGB >= vram {
			return tier
		}
	}
	return vram
}

// rightSizeUpGpuCount sizes non-standard physical GPU card counts up to standard power-of-two CSP topologies.
func rightSizeUpGpuCount(count uint32) uint32 {
	switch {
	case count <= 1:
		return 1
	case count == 2:
		return 2
	case count <= 4:
		return 4
	default:
		return 8
	}
}

// recommendGpuNodeSpec recommends appropriate GPU-accelerated node specs for the given node.
// It executes a targeted, single-shot query with discrete VRAM and topology right-sizing up,
// followed by multi-dimensional proximity ranking (Vendor > Count > VRAM > Model > Host L1 > Cost).
func recommendGpuNodeSpec(
	csp string,
	region string,
	node onpremmodel.NodeProperty,
	limit int,
) ([]cloudmodel.SpecInfo, int, error) {
	emptyResp := []cloudmodel.SpecInfo{}

	clusters := clusterGpuCards(node.GPUCards)
	if len(clusters) == 0 {
		return emptyResp, 0, fmt.Errorf("no GPU cards found for machine %s", node.MachineId)
	}

	primary := clusters[0]
	targetCount := rightSizeUpGpuCount(primary.Count)
	targetVram := rightSizeUpVramTier(primary.MemoryTotalGB)

	providerName := strings.ToLower(csp)
	regionName := strings.ToLower(region)

	architecture := node.CPU.Architecture
	if architecture == "" || architecture == "amd64" {
		architecture = defaultArchitecture
	}

	// Single-shot targeted query plan (simple, deterministic, no convoluted fallback loop)
	const planTemplate = `{
		"filter": {
			"policy": [
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

	fetchLimit := limit * 5
	if fetchLimit < minGpuFetchLimit {
		fetchLimit = minGpuFetchLimit
	}

	plan := fmt.Sprintf(planTemplate,
		targetCount,
		targetVram,
		providerName,
		regionName,
		architecture,
		fetchLimit,
	)

	log.Debug().
		Str("machineId", node.MachineId).
		Uint32("targetGpuCount", targetCount).
		Float32("targetVramGB", targetVram).
		Str("provider", providerName).
		Str("region", regionName).
		Str("architecture", architecture).
		Msg("Querying Tumblebug for GPU node spec recommendations")

	rawSpecs, err := tbclient.NewSession().InfraRecommendSpec(plan)
	if err != nil {
		log.Error().Err(err).
			Str("machineId", node.MachineId).
			Str("provider", providerName).
			Str("region", regionName).
			Msg("Failed to get GPU node spec recommendations from Tumblebug")
		return emptyResp, -1, fmt.Errorf("failed to get GPU node spec recommendations for machine %s: %w", node.MachineId, err)
	}

	// Filter specs with valid cost
	validSpecs := make([]tbmodel.SpecInfo, 0, len(rawSpecs))
	for _, spec := range rawSpecs {
		if spec.CostPerHour >= 0 {
			validSpecs = append(validSpecs, spec)
		}
	}

	if len(validSpecs) == 0 {
		log.Warn().
			Str("machineId", node.MachineId).
			Uint32("targetGpuCount", targetCount).
			Float32("targetVramGB", targetVram).
			Msg("No GPU node specs found matching requirements")
		return emptyResp, 0, nil
	}

	// Convert model types with validation
	convertedSpecs, err := modelconv.ConvertWithValidation[[]tbmodel.SpecInfo, []cloudmodel.SpecInfo](validSpecs)
	if err != nil {
		log.Error().Err(err).
			Str("machineId", node.MachineId).
			Msg("Failed to convert GPU node spec list model")
		return emptyResp, -1, fmt.Errorf("failed to convert GPU node spec list model for machine %s: %w", node.MachineId, err)
	}

	// Rank GPU specs (Vendor Match > Count > VRAM > Host L1 > Cost)
	sortGpuByProximityWithCost(convertedSpecs, node, csp)

	// Apply requested limit
	if limit > 0 && len(convertedSpecs) > limit {
		convertedSpecs = convertedSpecs[:limit]
	}

	if len(clusters) > 1 {
		log.Info().
			Str("machineId", node.MachineId).
			Int("heterogeneousClusters", len(clusters)).
			Msg("Heterogeneous GPU node detected: primary cluster recommended; secondary cluster(s) require multi-node decomposition")
	}

	return convertedSpecs, len(convertedSpecs), nil
}
