package recommendation

import (
	"fmt"
	"sort"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/compat"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
	"github.com/rs/zerolog/log"
)

// Recommendation limits constants for VM specs
const (
	defaultSpecsLimit = 30
)

// GetDefaultSpecsLimit returns the default VM specs recommendation limit
func GetDefaultSpecsLimit() int {
	return defaultSpecsLimit
}

// RecommendVmSpecsForImage recommends appropriate VM specs for the server and image
func RecommendVmSpecsForImage(csp string, region string, server onpremmodel.ServerProperty, limit int, image cloudmodel.ImageInfo) (vmSpecList []cloudmodel.SpecInfo, length int, err error) {

	if limit <= 0 {
		err := fmt.Errorf("invalid 'limit' value: %d, set default: %d", limit, defaultSpecsLimit)
		log.Warn().Msgf("%s", err.Error())
		limit = defaultSpecsLimit
	}

	vmSpecList, length, err = RecommendVmSpecs(csp, region, server, limit)
	if err != nil {
		log.Warn().Err(err).Msg("failed to recommend VM specs")
		return nil, 0, err
	}

	// Use unified compatibility filtering instead of CSP-specific switches
	compatibleSpecs := make([]cloudmodel.SpecInfo, 0, len(vmSpecList))

	for _, spec := range vmSpecList {
		if isCompatible := compat.CheckCompatibility(strings.ToLower(csp), spec, image); isCompatible {
			compatibleSpecs = append(compatibleSpecs, spec)
		} else {
			log.Debug().Msgf("Filtered incompatible spec: %s for image: %s on CSP: %s",
				spec.CspSpecName, image.CspImageName, csp)
		}
	}

	if len(compatibleSpecs) == 0 {
		log.Warn().Msgf("No compatible specs found for image %s on CSP %s, returning original list",
			image.CspImageName, csp)
		return vmSpecList, length, nil
	}

	log.Info().Msgf("Filtered %d specs to %d compatible specs for image %s on CSP %s",
		len(vmSpecList), len(compatibleSpecs), image.CspImageName, csp)

	return compatibleSpecs, len(compatibleSpecs), nil
}

// RecommendVmSpecs recommends appropriate VM specs for the given server
func RecommendVmSpecs(csp string, region string, server onpremmodel.ServerProperty, limit int) (vmSpecList []cloudmodel.SpecInfo, length int, err error) {

	// Constants
	const (
		defaultArchitecture = "x86_64"
	)

	var emptyResp = []cloudmodel.SpecInfo{}

	// Validate and set default limit
	if limit <= 0 {
		log.Warn().Msgf("Invalid limit value: %d, setting to default: %d", limit, defaultSpecsLimit)
		limit = defaultSpecsLimit
	}

	// Deployment plan template for VM spec recommendation
	// * Note:
	// * ">=": greater than or equal to
	// * "<=": less than or equal to
	// * The plan is designed to recommend VM specs based on vCPU and memory ranges.
	// Reference: https://github.com/cloud-barista/cb-tumblebug/discussions/1234
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
		"limit": "%d",
		"priority": {
			"policy": [{"metric": "cost"}]
		}
	}`

	// Extract server specifications from source computing envrionment
	// * Note: vcpus = cpus * cpuThreads
	vcpusCalculated := uint32(server.CPU.Cpus * server.CPU.Threads)
	memory := uint32(server.Memory.TotalSize)

	// Calculate optimal vCPU and memory ranges based on AWS, GCP, and NCP instance patterns
	vcpusMin, vcpusMax, memoryMin, memoryMax := calculateOptimalRange(vcpusCalculated, memory)

	// Set provider and region names
	providerName := strings.ToLower(csp)
	regionName := strings.ToLower(region)

	// Set architecture (default: "x86_64")
	architecture := server.CPU.Architecture
	if architecture == "" || architecture == "amd64" {
		architecture = defaultArchitecture
	}

	// Set OS name and version
	osNameAndVersion := server.OS.Name + " " + server.OS.Version
	osNameWithVersion := strings.ToLower(osNameAndVersion)

	log.Debug().
		Str("machineId", server.MachineId).
		Uint32("originalCpu*Threads", vcpusCalculated).
		Uint32("originalMemory", memory).
		Float64("memoryCpuThreadsRatio", float64(memory)/float64(vcpusCalculated)).
		Uint32("vcpuRange", vcpusMax-vcpusMin).
		Uint32("memoryRange", memoryMax-memoryMin).
		Str("provider", providerName).
		Str("region", regionName).
		Str("architecture", architecture).
		Str("osNameWithVersion", osNameWithVersion).
		Msgf("Calculating VM spec recommendations for machine: %s", server.MachineId)

	// Create deployment plan with calculated parameters
	deploymentPlan := fmt.Sprintf(planTemplate,
		vcpusMin, vcpusMax,
		memoryMin, memoryMax,
		providerName, regionName, architecture,
		limit,
	)
	log.Debug().Msgf("Deployment plan for machine %s: %s", server.MachineId, deploymentPlan)

	// Initialize Tumblebug API client
	tbCli := tbclient.NewClient(tbclient.ApiConfig{
		RestUrl:  config.Tumblebug.RestUrl,
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
	})

	// Call Tumblebug API to get recommended VM specs
	vmSpecInfoList, err := tbCli.MciRecommendSpec(deploymentPlan)
	if err != nil {
		log.Error().Err(err).
			Str("machineId", server.MachineId).
			Str("provider", providerName).
			Str("region", region).
			Msg("Failed to get VM spec recommendations from Tumblebug")
		return emptyResp, -1, fmt.Errorf("failed to get VM spec recommendations for machine %s: %w", server.MachineId, err)
	}

	// Filter if CostPerHour is less then 0
	validVmSpecs := make([]tbmodel.SpecInfo, 0, len(vmSpecInfoList))
	for _, spec := range vmSpecInfoList {
		if spec.CostPerHour >= 0 {
			validVmSpecs = append(validVmSpecs, spec)
		} else {
			log.Debug().Msgf("Filtered spec with negative cost: %s (CostPerHour: %.4f)",
				spec.CspSpecName, spec.CostPerHour)
		}
	}
	vmSpecInfoList = validVmSpecs

	// Check if any VM specs were found
	numOfVmSpecs := len(vmSpecInfoList)
	if numOfVmSpecs == 0 {
		err := fmt.Errorf("no VM specs recommended for machine %s (vcpusCalculated: %d, memory: %d GiB)",
			server.MachineId, vcpusCalculated, memory)
		log.Warn().Err(err).
			Str("machineId", server.MachineId).
			Uint32("vcpusCalculated", vcpusCalculated).
			Uint32("memory", memory).
			Msg("No VM specifications found")
		return emptyResp, -1, err
	}

	log.Info().
		Str("machineId", server.MachineId).
		Int("specsFound", numOfVmSpecs).
		Uint32("vcpusCalculated", vcpusCalculated).
		Uint32("memory", memory).
		Msgf("Found %d VM spec recommendations for machine: %s", numOfVmSpecs, server.MachineId)

	// NCP-specific filtering for KVM hypervisor
	if strings.Contains(strings.ToLower(csp), "ncp") {
		log.Debug().
			Str("machineId", server.MachineId).
			Msg("Filtering VM specs for KVM hypervisor (NCP)")

		kvmVmSpecs := make([]tbmodel.SpecInfo, 0, len(vmSpecInfoList))
		for _, vmSpec := range vmSpecInfoList {
			for _, detail := range vmSpec.Details {
				if detail.Key == "HypervisorType" && strings.Contains(strings.ToLower(detail.Value), "kvm") {
					kvmVmSpecs = append(kvmVmSpecs, vmSpec)
					break
				}
			}
		}

		if len(kvmVmSpecs) > 0 {
			vmSpecInfoList = kvmVmSpecs
			log.Debug().
				Str("machineId", server.MachineId).
				Int("kvmSpecs", len(kvmVmSpecs)).
				Msg("Filtered to KVM-compatible specs for NCP")
		} else {
			log.Warn().
				Str("machineId", server.MachineId).
				Msg("No KVM-compatible specs found for NCP, using all available specs")
		}
	}

	// [Output]
	// Apply limit to results
	finalSpecCount := len(vmSpecInfoList)
	if limit < finalSpecCount {
		vmSpecInfoList = vmSpecInfoList[:limit]
		finalSpecCount = limit
	}

	log.Debug().
		Str("machineId", server.MachineId).
		Int("finalSpecCount", finalSpecCount).
		Msg("Finalized VM spec recommendations")

	// Convert model types with validation
	convertedVmSpecList, err := modelconv.ConvertWithValidation[[]tbmodel.SpecInfo, []cloudmodel.SpecInfo](vmSpecInfoList)
	if err != nil {
		log.Error().Err(err).
			Str("machineId", server.MachineId).
			Msg("Failed to convert VM spec list model")
		return emptyResp, -1, fmt.Errorf("failed to convert VM spec list model for machine %s: %w", server.MachineId, err)
	}

	// Sort specs by proximity with cost consideration
	sortByProximityWithCost(vcpusCalculated, memory, convertedVmSpecList)

	// // ! Logging section for research purpose
	// log.Info().Msgf("No.,Provider,Region,VM Spec ID,vCPU,MemoryGiB,CostPerHour")
	// for i, vmSpec := range convertedVmSpecList {
	// 	log.Info().Msgf("%d,%s,%s,%s,%d,%.2f,%.4f",
	// 		i+1, vmSpec.ProviderName, vmSpec.RegionName, vmSpec.CspSpecName, vmSpec.VCPU, vmSpec.MemoryGiB, vmSpec.CostPerHour)
	// }

	log.Info().
		Str("machineId", server.MachineId).
		Int("recommendedSpecs", len(convertedVmSpecList)).
		Msgf("Successfully recommended %d VM specifications for machine: %s", len(convertedVmSpecList), server.MachineId)

	return convertedVmSpecList, numOfVmSpecs, nil
}

// Sort VM specs by proximity to the desired resource allocation with cost consideration
func sortByProximityWithCost(vcpus, memory uint32, vmSpecs []cloudmodel.SpecInfo) {

	// Derive server's spec type (i.e. compute intensive type, memory intensive type, general purpose type)
	machineType := deriveMachineType(vcpus, memory)

	log.Debug().Msgf("Sorting VM specs for machine type: %s (vcpus: %d, memory: %d GiB)", machineType, vcpus, memory)

	// Sort based on the machine type
	switch machineType {
	case "compute-intensive":
		// Sort by proximity to desired values with cost as secondary criterion
		// 1. First sort by vCPU proximity (closest to target first)
		// 2. Within same vCPU values, sort by memory proximity (closest to target first)
		// 3. If both are same, sort by cost per hour (lowest cost first)
		sort.Slice(vmSpecs, func(i, j int) bool {
			vcpuDiffI := abs(int32(vmSpecs[i].VCPU) - int32(vcpus))
			vcpuDiffJ := abs(int32(vmSpecs[j].VCPU) - int32(vcpus))

			// If vCPU differences are different, sort by vCPU proximity
			if vcpuDiffI != vcpuDiffJ {
				return vcpuDiffI < vcpuDiffJ
			}

			// If vCPU differences are same, sort by memory proximity
			memDiffI := abs(int32(vmSpecs[i].MemoryGiB) - int32(memory))
			memDiffJ := abs(int32(vmSpecs[j].MemoryGiB) - int32(memory))
			if memDiffI != memDiffJ {
				return memDiffI < memDiffJ
			}

			// If both vCPU and memory differences are same, sort by cost per hour
			return vmSpecs[i].CostPerHour < vmSpecs[j].CostPerHour
		})
	case "memory-intensive":
		// Sort by proximity to desired values with cost as secondary criterion
		// 1. First sort by memory proximity (closest to target first)
		// 2. Within same memory values, sort by vCPU proximity (closest to target first)
		// 3. If both are same, sort by cost per hour (lowest cost first)
		sort.Slice(vmSpecs, func(i, j int) bool {
			memDiffI := abs(int32(vmSpecs[i].MemoryGiB) - int32(memory))
			memDiffJ := abs(int32(vmSpecs[j].MemoryGiB) - int32(memory))

			// If memory differences are different, sort by memory proximity
			if memDiffI != memDiffJ {
				return memDiffI < memDiffJ
			}

			// If memory differences are same, sort by vCPU proximity
			vcpuDiffI := abs(int32(vmSpecs[i].VCPU) - int32(vcpus))
			vcpuDiffJ := abs(int32(vmSpecs[j].VCPU) - int32(vcpus))
			if vcpuDiffI != vcpuDiffJ {
				return vcpuDiffI < vcpuDiffJ
			}

			// If both memory and vCPU differences are same, sort by cost per hour
			return vmSpecs[i].CostPerHour < vmSpecs[j].CostPerHour
		})
	default: // "general-purpose"
		// * Note: Manhattan Distance is preferred over Euclidean distance for VM specs
		// because CPU and memory are independent resources with different scales

		// Sort by Manhattan distance (L1 norm) for balanced workloads with cost as secondary criterion
		sort.Slice(vmSpecs, func(i, j int) bool {
			vcpuDiffI := abs(int32(vmSpecs[i].VCPU) - int32(vcpus))
			memDiffI := abs(int32(vmSpecs[i].MemoryGiB) - int32(memory))
			totalDiffI := vcpuDiffI + memDiffI

			vcpuDiffJ := abs(int32(vmSpecs[j].VCPU) - int32(vcpus))
			memDiffJ := abs(int32(vmSpecs[j].MemoryGiB) - int32(memory))
			totalDiffJ := vcpuDiffJ + memDiffJ

			// If total differences are different, sort by total difference
			if totalDiffI != totalDiffJ {
				return totalDiffI < totalDiffJ
			}

			// If total differences are same, sort by cost per hour (lowest cost first)
			return vmSpecs[i].CostPerHour < vmSpecs[j].CostPerHour
		})
	}
}

// abs returns the absolute value of x
func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// MBtoGiB converts megabytes to gibibytes
func MBtoGiB(mb float64) uint32 {
	const bytesInMB = 1000000.0
	const bytesInGiB = 1073741824.0
	gib := (mb * bytesInMB) / bytesInGiB
	return uint32(gib)
}

// deriveMachineType derives the machine type based on vCPU and memory
func deriveMachineType(vcpus uint32, memory uint32) (machineType string) {
	const (
		computeIntensiveRatioThreshold = 3.0 // 1:2 ratio instances
		memoryIntensiveRatioThreshold  = 7.0 // 1:8 ratio instances
	)

	memoryToVcpuRatio := float64(memory) / float64(vcpus)

	switch {
	case memoryToVcpuRatio <= computeIntensiveRatioThreshold: // Compute Intensive (1:2)
		return "compute-intensive"
	case memoryToVcpuRatio >= memoryIntensiveRatioThreshold: // Memory Intensive (1:8)
		return "memory-intensive"
	default: // General Purpose (1:4)
		return "general-purpose"
	}
}

// calculateOptimalRange calculates optimal vCPU and memory ranges based on AWS instance patterns
func calculateOptimalRange(vcpus uint32, memory uint32) (vcpusMin, vcpusMax, memoryMin, memoryMax uint32) {
	// Constants for instance type thresholds and ratios
	const (
		computeIntensiveRatioThreshold = 3.0 // 1:2 ratio instances
		memoryIntensiveRatioThreshold  = 7.0 // 1:8 ratio instances
		// minMemoryBound                 = 2   // Minimum memory requirement
		// minVcpuBound                   = 1   // Minimum vCPU requirement
		// maxVcpuForMemoryIntensive      = 10  // Maximum vCPU for memory intensive
	)

	memoryToVcpuRatio := float64(memory) / float64(vcpus)

	switch {
	case memoryToVcpuRatio <= computeIntensiveRatioThreshold: // Compute Intensive (1:2)
		return calculateComputeIntensiveRange(vcpus, memory)
	case memoryToVcpuRatio >= memoryIntensiveRatioThreshold: // Memory Intensive (1:8)
		return calculateMemoryIntensiveRange(vcpus, memory)
	default: // General Purpose (1:4)
		return calculateGeneralPurposeRange(vcpus, memory)
	}
}

func calculateComputeIntensiveRange(vcpus, memory uint32) (vcpusRangeMin, vcpusRangeMax, memoryRangeMin, memoryRangeMax uint32) {
	const (
		memoryMultiplier = 4 // Memory multiplier for max calculation
	)

	vcpusRangeMin = findPreviousPrimeOrDecrementOne(vcpus)
	vcpusRangeMax = calculateRangeMax(vcpus) // find the next next prime number

	// Set a wide search range for memory for compute-intensive workloads
	memoryRangeMin = 0
	memoryRangeMax = vcpusRangeMax * memoryMultiplier

	return vcpusRangeMin, vcpusRangeMax, memoryRangeMin, memoryRangeMax
}

func calculateMemoryIntensiveRange(vcpus, memory uint32) (vcpusMin, vcpusMax, memoryRangeMin, memoryRangeMax uint32) {
	const (
		memoryToCpuRatio = 7 // memory to CPU ratio for calculation (Standard: 8)
	)

	memoryRangeMin = calculateRangeMin(memory)
	memoryRangeMax = calculateRangeMax(memory)

	// Set a wide search range for vCPU for memory-intensive workloads
	vcpusMin = 0
	vcpusMax = memoryRangeMax / memoryToCpuRatio

	return vcpusMin, vcpusMax, memoryRangeMin, memoryRangeMax
}

func calculateGeneralPurposeRange(vcpus, memory uint32) (vcpusMin, vcpusMax, memoryMin, memoryMax uint32) {
	// For General Purpose workloads, provide balanced flexibility for both vCPU and memory
	// The input has already been classified as General Purpose in calculateOptimalRange

	vcpusMin = findPreviousPrimeOrDecrementOne(vcpus)
	vcpusMax = calculateRangeMax(vcpus) // find the next next prime number

	memoryMin = calculateRangeMin(memory)
	memoryMax = calculateRangeMax(memory)

	return vcpusMin, vcpusMax, memoryMin, memoryMax
}

// calculateRangeMin calculates the minimum value for a range based on a given number
func calculateRangeMin(n uint32) uint32 {

	// Calculate previous previous prime number
	min := findPreviousPrimeOrDecrementOne(n)
	min = findPreviousPrimeOrDecrementOne(min)

	return min
}

// calculateRangeMax calculates the maximum value for a range based on a given number
func calculateRangeMax(n uint32) uint32 {

	// Calculate next next prime number
	max := findNextPrimeNumber(n)
	max = findNextPrimeNumber(max)

	// Expand the range if it's too narrow
	if max-n < 4 {
		max = findNextPrimeNumber(max)
	}

	return max
}

// isPrimeNumber checks if a number is prime
func isPrimeNumber(n uint32) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := uint32(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// findPreviousPrimeOrDecrementOne finds the largest prime number smaller than n,
// returns 1 if n=2, returns 0 if n=1
func findPreviousPrimeOrDecrementOne(n uint32) uint32 {

	// Return 1 when n is 2
	if n == 2 {
		return 1
	}

	// Return 0 when n is 1 or less
	if n <= 1 {
		return 0
	}

	// Find the prime number smaller than n
	for i := n - 1; i >= 2; i-- {
		if isPrimeNumber(i) {
			return i
		}
	}
	return 0 // Return 0 as fallback minimum value
}

// findNextPrimeNumber finds the smallest prime number larger than n
func findNextPrimeNumber(n uint32) uint32 {
	candidate := n + 1
	for {
		if isPrimeNumber(candidate) {
			return candidate
		}
		candidate++
	}
}
