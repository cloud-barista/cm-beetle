package recommendation

import (
	"fmt"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/cloud-barista/cm-beetle/pkg/compat"
	"github.com/cloud-barista/cm-beetle/pkg/similarity"
	"github.com/rs/zerolog/log"
)

// GPU model match rate score tiers (0-100%) for GPU-accelerated nodes.
const (
	gpuModelScoreExactMatch      = 100.0
	gpuModelScoreSameVendor      = 75.0
	gpuModelScoreDifferentVendor = 50.0
	gpuModelScoreNoGpu           = 0.0
)

// MatchRateVector represents multi-dimensional match rate scores (0-100%) for a node.
// CPU-only nodes evaluate 3 dimensions; GPU nodes evaluate 6 dimensions.
type MatchRateVector struct {
	CPU      float64
	Memory   float64
	Image    float64
	HasGPU   bool
	GPUCount float64
	GPUVRAM  float64
	GPUModel float64
}

// ActiveDimensions returns the 3 active dimensions for CPU nodes or 6 for GPU nodes.
func (q MatchRateVector) ActiveDimensions() []float64 {
	if q.HasGPU {
		return []float64{q.CPU, q.Memory, q.Image, q.GPUCount, q.GPUVRAM, q.GPUModel}
	}
	return []float64{q.CPU, q.Memory, q.Image}
}

// MinMatchRate returns the minimum match rate score across all active dimensions.
func (q MatchRateVector) MinMatchRate() float64 {
	dims := q.ActiveDimensions()
	minVal := dims[0]
	for _, val := range dims[1:] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

// MaxMatchRate returns the maximum match rate score across all active dimensions.
func (q MatchRateVector) MaxMatchRate() float64 {
	dims := q.ActiveDimensions()
	maxVal := dims[0]
	for _, val := range dims[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// IsMatched checks if all active dimensions meet or exceed the match rate threshold.
func (q MatchRateVector) IsMatched(threshold float64) bool {
	for _, val := range q.ActiveDimensions() {
		if val < threshold {
			return false
		}
	}
	return true
}

// AverageMatchRate returns the arithmetic mean across all active dimensions.
func (q MatchRateVector) AverageMatchRate() float64 {
	dims := q.ActiveDimensions()
	total := 0.0
	for _, val := range dims {
		total += val
	}
	return total / float64(len(dims))
}

// FormatMatchRateSummary formats the active dimensions into a human-readable summary string.
func (q MatchRateVector) FormatMatchRateSummary() string {
	if q.HasGPU {
		return fmt.Sprintf(
			"Match Rate: CPU=%.1f%% Memory=%.1f%% Image=%.1f%% GPUCount=%.1f%% GPUVRAM=%.1f%% GPUModel=%.1f%%",
			q.CPU, q.Memory, q.Image, q.GPUCount, q.GPUVRAM, q.GPUModel,
		)
	}
	return fmt.Sprintf(
		"Match Rate: CPU=%.1f%% Memory=%.1f%% Image=%.1f%%",
		q.CPU, q.Memory, q.Image,
	)
}

// InfraMatchRateSummary represents overall candidate infrastructure match rate metrics.
type InfraMatchRateSummary struct {
	MinMatchRate   float64
	MaxMatchRate   float64
	AvgMatchRate   float64
	BestEffortRate float64
	TotalVMs       int
}

// calculateCandidateMatchRateWithDetails evaluates all node groups in a candidate against minMatchRate.
func calculateCandidateMatchRateWithDetails(
	csp string,
	tempNodeGroupList []cloudmodel.CreateNodeGroupReq,
	srcInfra onpremmodel.OnpremInfra,
	deduplicatedVmSpecList []cloudmodel.SpecInfo,
	deduplicatedVmOsImageList []cloudmodel.ImageInfo,
	minMatchRate float64,
) (string, string, InfraMatchRateSummary) {
	var overallStatus string
	var overallStatusDesc string
	var summary InfraMatchRateSummary

	if len(tempNodeGroupList) == 0 {
		overallStatus = "nothing-to-recommend"
		overallStatusDesc = "No VMs available for recommendation"
		summary = InfraMatchRateSummary{MinMatchRate: 0, MaxMatchRate: 0, AvgMatchRate: 0, BestEffortRate: 0, TotalVMs: 0}
		return overallStatus, overallStatusDesc, summary
	}

	lowestMatchRate := 100.0
	highestMatchRate := 0.0
	totalMatchRate := 0.0
	totalDimensionCount := 0
	bestEffortCount := 0
	validServerCount := 0

	for j, nodeGroup := range tempNodeGroupList {
		if j >= len(srcInfra.Nodes) {
			break
		}
		if nodeGroup.SpecId == "" || nodeGroup.ImageId == "" {
			continue
		}

		node := srcInfra.Nodes[j]

		var selectedSpec cloudmodel.SpecInfo
		for _, spec := range deduplicatedVmSpecList {
			if spec.Id == nodeGroup.SpecId {
				selectedSpec = spec
				break
			}
		}

		var selectedImage cloudmodel.ImageInfo
		for _, image := range deduplicatedVmOsImageList {
			if image.Id == nodeGroup.ImageId {
				selectedImage = image
				break
			}
		}

		if selectedSpec.Id == "" || selectedImage.Id == "" {
			continue
		}

		matchRateVec := calculateMatchRateVector(csp, node, selectedSpec, selectedImage)
		validServerCount++

		vmMinMatchRate := matchRateVec.MinMatchRate()
		vmMaxMatchRate := matchRateVec.MaxMatchRate()
		vmAvgMatchRate := matchRateVec.AverageMatchRate()

		if vmMinMatchRate < lowestMatchRate {
			lowestMatchRate = vmMinMatchRate
		}
		if vmMaxMatchRate > highestMatchRate {
			highestMatchRate = vmMaxMatchRate
		}

		activeDims := matchRateVec.ActiveDimensions()
		for _, dimScore := range activeDims {
			totalMatchRate += dimScore
		}
		totalDimensionCount += len(activeDims)

		if matchRateVec.IsMatched(minMatchRate) {
			bestEffortCount++
		}

		log.Trace().
			Str("machineId", node.MachineId).
			Bool("hasGpu", matchRateVec.HasGPU).
			Float64("cpuMatchRate", matchRateVec.CPU).
			Float64("memoryMatchRate", matchRateVec.Memory).
			Float64("imageMatchRate", matchRateVec.Image).
			Float64("gpuCountMatchRate", matchRateVec.GPUCount).
			Float64("gpuVramMatchRate", matchRateVec.GPUVRAM).
			Float64("gpuModelMatchRate", matchRateVec.GPUModel).
			Float64("vmMinMatchRate", vmMinMatchRate).
			Float64("vmMaxMatchRate", vmMaxMatchRate).
			Float64("vmAvgMatchRate", vmAvgMatchRate).
			Str("specId", selectedSpec.Id).
			Str("imageId", selectedImage.Id).
			Msg("Individual node match rate assessment")
	}

	if validServerCount == 0 || totalDimensionCount == 0 {
		overallStatus = "nothing-to-recommend"
		overallStatusDesc = "No valid VM spec-image pairs found"
		summary = InfraMatchRateSummary{MinMatchRate: 0, MaxMatchRate: 0, AvgMatchRate: 0, BestEffortRate: 0, TotalVMs: 0}
		return overallStatus, overallStatusDesc, summary
	}

	avgMatchRate := totalMatchRate / float64(totalDimensionCount)
	bestEffortRate := float64(bestEffortCount) / float64(validServerCount) * 100.0

	if bestEffortCount == validServerCount {
		overallStatus = "highly-matched"
	} else {
		overallStatus = "partially-matched"
	}

	overallStatusDesc = fmt.Sprintf(
		"VMs: %d total, %d matched, %d acceptable",
		validServerCount,
		bestEffortCount,
		validServerCount-bestEffortCount,
	)

	summary = InfraMatchRateSummary{
		MinMatchRate:   lowestMatchRate,
		MaxMatchRate:   highestMatchRate,
		AvgMatchRate:   avgMatchRate,
		BestEffortRate: bestEffortRate,
		TotalVMs:       validServerCount,
	}

	log.Info().
		Float64("lowestMatchRate", lowestMatchRate).
		Float64("highestMatchRate", highestMatchRate).
		Float64("avgMatchRate", avgMatchRate).
		Float64("bestEffortRate", bestEffortRate).
		Str("overallStatus", overallStatus).
		Int("bestEffortCount", bestEffortCount).
		Int("validServerCount", validServerCount).
		Int("totalDimensionCount", totalDimensionCount).
		Msg("Overall candidate match rate assessment")

	return overallStatus, overallStatusDesc, summary
}

// calculateMatchRateVector evaluates 3 dimensions for CPU-only nodes and 6 dimensions for GPU nodes.
func calculateMatchRateVector(
	csp string,
	node onpremmodel.NodeProperty,
	vmSpec cloudmodel.SpecInfo,
	vmImage cloudmodel.ImageInfo,
) MatchRateVector {
	serverThreads := node.CPU.Threads
	if serverThreads == 0 {
		serverThreads = 1
	}
	serverVCPUs := float64(node.CPU.Cpus * serverThreads)
	vmSpecVCPUs := float64(vmSpec.VCPU)
	cpuMatchRate := calculateRelativeMatch(serverVCPUs, vmSpecVCPUs)

	serverMemoryGB := float64(node.Memory.TotalSize)
	vmSpecMemoryGiB := float64(vmSpec.MemoryGiB)
	memoryMatchRate := calculateRelativeMatch(serverMemoryGB, vmSpecMemoryGiB)

	imageMatchRate := calculateImageMatchRateScore(csp, node, vmImage) * 100.0

	matchRateVec := MatchRateVector{
		CPU:    cpuMatchRate,
		Memory: memoryMatchRate,
		Image:  imageMatchRate,
		HasGPU: hasGpu(node),
	}

	if matchRateVec.HasGPU {
		matchRateVec.GPUCount = calculateGpuCountMatchRate(node, vmSpec)
		matchRateVec.GPUVRAM = calculateGpuVramMatchRate(node, vmSpec)
		matchRateVec.GPUModel = calculateGpuModelMatchRate(node, vmSpec)
	}

	log.Debug().
		Str("machineId", node.MachineId).
		Bool("hasGpu", matchRateVec.HasGPU).
		Float64("cpuMatchRate", matchRateVec.CPU).
		Float64("memoryMatchRate", matchRateVec.Memory).
		Float64("imageMatchRate", matchRateVec.Image).
		Float64("gpuCountMatchRate", matchRateVec.GPUCount).
		Float64("gpuVramMatchRate", matchRateVec.GPUVRAM).
		Float64("gpuModelMatchRate", matchRateVec.GPUModel).
		Msg("Match rate vector calculated")

	return matchRateVec
}

// calculateRelativeMatch computes scale-independent relative match as (min(a,b) / max(a,b)) * 100.
// Example: 8 CPUs vs 16 CPUs = 50.0%, 16 GB vs 16 GB = 100.0%.
func calculateRelativeMatch(serverValue, vmValue float64) float64 {
	if serverValue == 0 && vmValue == 0 {
		return 100.0
	}
	if serverValue == 0 || vmValue == 0 {
		return 0.0
	}

	minVal := serverValue
	maxVal := vmValue
	if vmValue < serverValue {
		minVal = vmValue
		maxVal = serverValue
	}

	return (minVal / maxVal) * 100.0
}

// calculateGpuCountMatchRate compares the primary source GPU cluster count against target AcceleratorCount.
func calculateGpuCountMatchRate(node onpremmodel.NodeProperty, vmSpec cloudmodel.SpecInfo) float64 {
	clusters := clusterGpuCards(node.GPUCards)
	if len(clusters) == 0 {
		if vmSpec.AcceleratorCount == 0 {
			return 100.0
		}
		return 0.0
	}

	primary := clusters[0]
	return calculateRelativeMatch(float64(primary.Count), float64(vmSpec.AcceleratorCount))
}

// calculateGpuVramMatchRate compares per-device GPU VRAM with vramTierToleranceGB (0.5 GB) tolerance.
// Upsizing (target >= source - 0.5 GB) awards 100.0%, while downsizing evaluates relative match (target / source * 100).
func calculateGpuVramMatchRate(node onpremmodel.NodeProperty, vmSpec cloudmodel.SpecInfo) float64 {
	clusters := clusterGpuCards(node.GPUCards)
	if len(clusters) == 0 {
		if vmSpec.AcceleratorMemoryGB == 0 {
			return 100.0
		}
		return 0.0
	}

	sourceVram := float64(clusters[0].MemoryTotalGB)
	targetVram := float64(vmSpec.AcceleratorMemoryGB)

	if sourceVram <= 0 {
		if targetVram == 0 {
			return 100.0
		}
		return 0.0
	}

	// Upsizing: target VRAM satisfies source VRAM requirement within tolerance.
	if targetVram+float64(vramTierToleranceGB) >= sourceVram {
		return 100.0
	}

	// Downsizing: compute relative proportion to reflect VRAM deficiency.
	return (targetVram / sourceVram) * 100.0
}

// calculateGpuModelMatchRate evaluates GPU vendor and chip model compatibility using a 4-tier score.
// Tiers: 100% (same vendor + model), 75% (same vendor + different model), 50% (cross-vendor GPU), 0% (no GPU).
func calculateGpuModelMatchRate(node onpremmodel.NodeProperty, vmSpec cloudmodel.SpecInfo) float64 {
	clusters := clusterGpuCards(node.GPUCards)
	if len(clusters) == 0 {
		return 100.0
	}

	if vmSpec.AcceleratorCount == 0 && strings.TrimSpace(vmSpec.AcceleratorModel) == "" {
		return gpuModelScoreNoGpu
	}

	primary := clusters[0]
	sourceVendor := detectGpuVendor(primary.Vendor)
	if sourceVendor == GpuVendorOther {
		sourceVendor = detectGpuVendor(primary.Model)
	}
	targetVendor := detectGpuVendor(vmSpec.AcceleratorModel)

	if sourceVendor != GpuVendorOther && sourceVendor == targetVendor {
		sourceModel := detectGpuModel(primary.Model)
		targetModel := detectGpuModel(vmSpec.AcceleratorModel)

		if sourceModel != "" && sourceModel == targetModel {
			return gpuModelScoreExactMatch
		}
		if sourceModel == "" && targetModel == "" &&
			strings.EqualFold(strings.TrimSpace(primary.Model), strings.TrimSpace(vmSpec.AcceleratorModel)) {
			return gpuModelScoreExactMatch
		}
		return gpuModelScoreSameVendor
	}

	return gpuModelScoreDifferentVendor
}

// calculateImageMatchRateScore computes OS image similarity (0.0 to 1.0) for CPU-only and CPU+GPU nodes.
func calculateImageMatchRateScore(csp string, node onpremmodel.NodeProperty, vmImage cloudmodel.ImageInfo) float64 {
	keywords, kwDelimiters, imgDelimiters := SetKeywordsAndDelimeters(node)

	vmImgKeywords := fmt.Sprintf("%s %s %s %s",
		vmImage.OSType,
		vmImage.OSArchitecture,
		vmImage.OSDiskType,
		vmImage.OSDistribution,
	)

	similarityScore := similarity.CalcResourceSimilarity(keywords, kwDelimiters, vmImgKeywords, imgDelimiters)

	priority := compat.GetImagePriority(csp, vmImage)
	if priority > 5 {
		similarityScore *= 0.5
	}

	if hasGpu(node) && !isGpuSuitableImage(node, vmImage) {
		similarityScore *= 0.5
	}

	return similarityScore
}

// isGpuSuitableImage checks whether the target OS image is GPU-enabled or carries GPU driver support.
func isGpuSuitableImage(node onpremmodel.NodeProperty, vmImage cloudmodel.ImageInfo) bool {
	if vmImage.IsGPUImage || vmImage.IsBasicGpuImage {
		return true
	}

	combined := strings.ToLower(fmt.Sprintf(
		"%s %s %s",
		vmImage.OSDistribution,
		vmImage.CspImageName,
		vmImage.Description,
	))

	gpuIndicators := []string{"gpu", "cuda", "nvidia", "rocm", "deep learning", "deeplearning", "pytorch", "tensorflow"}
	for _, token := range gpuIndicators {
		if strings.Contains(combined, token) {
			return true
		}
	}

	if len(node.GPUCards) > 0 {
		vendor := strings.ToLower(strings.TrimSpace(node.GPUCards[0].Vendor))
		if vendor != "" && strings.Contains(combined, vendor) {
			return true
		}
	}

	return false
}
