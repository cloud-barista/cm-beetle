package recommendation

import (
	"fmt"
	"sort"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/compat"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
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

// RecommendVmSpecsForImage recommends appropriate VM specs for the node and image
func RecommendVmSpecsForImage(csp string, region string, node onpremmodel.NodeProperty, limit int, image cloudmodel.ImageInfo) (vmSpecList []cloudmodel.SpecInfo, length int, err error) {

	if limit <= 0 {
		err := fmt.Errorf("invalid 'limit' value: %d, set default: %d", limit, defaultSpecsLimit)
		log.Warn().Msgf("%s", err.Error())
		limit = defaultSpecsLimit
	}

	vmSpecList, length, err = RecommendNodeSpecs(csp, region, node, limit)
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

// RecommendNodeSpecs recommends appropriate node specs (VM specs) for the given node
func RecommendNodeSpecs(csp string, region string, node onpremmodel.NodeProperty, limit int) (vmSpecList []cloudmodel.SpecInfo, length int, err error) {

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

	// Deployment plan template for node spec recommendation
	// * Note:
	// * ">=": greater than or equal to
	// * "<=": less than or equal to
	// * The plan is designed to recommend node specs based on vCPU and memory ranges.
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
		"limit": %d,
		"priority": {
			"policy": [{"metric": "cost"}]
		}
	}`

	// Extract node specifications from source computing envrionment
	// * Note: vcpus = cpus * cpuThreads
	cpus := node.CPU.Cpus
	threads := node.CPU.Threads
	if threads == 0 {
		threads = 1 // Default to 1 thread if not specified
	}

	vcpusCalculated := uint32(cpus * threads)
	memory := uint32(node.Memory.TotalSize)

	// Set provider and region names
	providerName := strings.ToLower(csp)
	regionName := strings.ToLower(region)

	// Set architecture (default: "x86_64")
	architecture := node.CPU.Architecture
	if architecture == "" || architecture == "amd64" {
		architecture = defaultArchitecture
	}

	// Iterative search with increasing rangeWeight to find suitable node specs
	const (
		initialRangeWeight = 1
		maxRangeWeight     = 5
	)

	var (
		nodeSpecInfoList     []tbmodel.SpecInfo
		vcpusMin, vcpusMax   uint32
		memoryMin, memoryMax uint32
	)

	// Retry loop: increase rangeWeight if no specs are found
	for rangeWeight := initialRangeWeight; rangeWeight <= maxRangeWeight; rangeWeight++ {
		// Calculate optimal vCPU and memory ranges based on AWS, GCP, and NCP instance patterns
		vcpusMin, vcpusMax, memoryMin, memoryMax = calculateOptimalRange(vcpusCalculated, memory, rangeWeight)

		log.Debug().
			Str("machineId", node.MachineId).
			Int("rangeWeight", rangeWeight).
			Uint32("originalCpu*Threads", vcpusCalculated).
			Uint32("originalMemory", memory).
			Float64("memoryCpuThreadsRatio", float64(memory)/float64(vcpusCalculated)).
			Uint32("vcpuRange", vcpusMax-vcpusMin).
			Uint32("memoryRange", memoryMax-memoryMin).
			Str("provider", providerName).
			Str("region", regionName).
			Str("architecture", architecture).
			Msgf("Calculating node spec recommendations for machine: %s (attempt %d/%d)", node.MachineId, rangeWeight, maxRangeWeight)

		// Create a plan to search proper node specs with calculated parameters
		planToSearchProperNode := fmt.Sprintf(planTemplate,
			vcpusMin, vcpusMax,
			memoryMin, memoryMax,
			providerName, regionName, architecture,
			limit,
		)
		log.Debug().Msgf("Deployment plan for machine %s: %s", node.MachineId, planToSearchProperNode)

		// Call Tumblebug API to get recommended node specs
		var err error
		nodeSpecInfoList, err = tbclient.NewSession().InfraRecommendSpec(planToSearchProperNode)
		if err != nil {
			log.Error().Err(err).
				Str("machineId", node.MachineId).
				Str("provider", providerName).
				Str("region", regionName).
				Int("rangeWeight", rangeWeight).
				Msg("Failed to get node spec recommendations from Tumblebug")
			return emptyResp, -1, fmt.Errorf("failed to get node spec recommendations for machine %s: %w", node.MachineId, err)
		}

		// Filter specs with valid cost
		validNodeSpecs := make([]tbmodel.SpecInfo, 0, len(nodeSpecInfoList))
		for _, spec := range nodeSpecInfoList {
			if spec.CostPerHour >= 0 {
				validNodeSpecs = append(validNodeSpecs, spec)
			} else {
				log.Debug().Msgf("Filtered node spec with negative cost: %s (CostPerHour: %.4f)",
					spec.CspSpecName, spec.CostPerHour)
			}
		}
		nodeSpecInfoList = validNodeSpecs

		// NCP-specific filtering for KVM hypervisor
		if strings.Contains(strings.ToLower(csp), "ncp") {
			log.Debug().
				Str("machineId", node.MachineId).
				Int("rangeWeight", rangeWeight).
				Msg("Filtering node specs for KVM hypervisor (NCP)")

			kvmNodeSpecs := make([]tbmodel.SpecInfo, 0, len(nodeSpecInfoList))
			for _, nodeSpec := range nodeSpecInfoList {
				for _, detail := range nodeSpec.Details {
					if detail.Key == "HypervisorType" && strings.Contains(strings.ToLower(detail.Value), "kvm") {
						kvmNodeSpecs = append(kvmNodeSpecs, nodeSpec)
						break
					}
				}
			}

			if len(kvmNodeSpecs) > 0 {
				nodeSpecInfoList = kvmNodeSpecs
				log.Debug().
					Str("machineId", node.MachineId).
					Int("kvmSpecs", len(kvmNodeSpecs)).
					Int("rangeWeight", rangeWeight).
					Msg("Filtered to KVM-compatible specs for NCP")
			} else {
				log.Debug().
					Str("machineId", node.MachineId).
					Int("rangeWeight", rangeWeight).
					Msg("No KVM-compatible specs found for NCP at this rangeWeight, will retry with increased range")
				// Continue to retry with increased rangeWeight
				continue
			}
		}

		// Check if any node specs were found
		if len(nodeSpecInfoList) > 0 {
			log.Info().
				Str("machineId", node.MachineId).
				Int("specsFound", len(nodeSpecInfoList)).
				Int("rangeWeight", rangeWeight).
				Uint32("vcpusCalculated", vcpusCalculated).
				Uint32("memory", memory).
				Msgf("Found %d node spec recommendations for machine: %s with rangeWeight: %d", len(nodeSpecInfoList), node.MachineId, rangeWeight)
			break // Exit loop if specs found
		}

		log.Warn().
			Str("machineId", node.MachineId).
			Int("rangeWeight", rangeWeight).
			Int("maxRangeWeight", maxRangeWeight).
			Msgf("No node specs found for machine %s with rangeWeight %d, retrying with increased range...", node.MachineId, rangeWeight)
	}

	// Final check after all retry attempts
	numOfNodeSpecs := len(nodeSpecInfoList)
	if numOfNodeSpecs == 0 {
		err := fmt.Errorf("no node specs recommended for machine %s after %d attempts (vcpusCalculated: %d, memory: %d GiB)",
			node.MachineId, maxRangeWeight, vcpusCalculated, memory)
		log.Warn().Err(err).
			Str("machineId", node.MachineId).
			Uint32("vcpusCalculated", vcpusCalculated).
			Uint32("memory", memory).
			Msg("No node specifications found")
		return emptyResp, -1, err
	}

	// [Output]
	// Apply limit to results
	finalSpecCount := len(nodeSpecInfoList)
	if limit < finalSpecCount {
		nodeSpecInfoList = nodeSpecInfoList[:limit]
		finalSpecCount = limit
	}

	log.Debug().
		Str("machineId", node.MachineId).
		Int("finalSpecCount", finalSpecCount).
		Msg("Finalized node spec recommendations")

	// Convert model types with validation
	convertedNodeSpecList, err := modelconv.ConvertWithValidation[[]tbmodel.SpecInfo, []cloudmodel.SpecInfo](nodeSpecInfoList)
	if err != nil {
		log.Error().Err(err).
			Str("machineId", node.MachineId).
			Msg("Failed to convert node spec list model")
		return emptyResp, -1, fmt.Errorf("failed to convert node spec list model for machine %s: %w", node.MachineId, err)
	}

	// Sort specs by proximity with cost consideration
	sortByProximityWithCost(convertedNodeSpecList, vcpusCalculated, memory, providerName, extractCpuVendor(node.CPU.Vendor))

	// // ! Logging section for research purpose
	// log.Info().Msgf("No.,Provider,Region,node Spec ID,vCPU,MemoryGiB,CostPerHour")
	// for i, nodeSpec := range convertedNodeSpecList {
	// 	log.Info().Msgf("%d,%s,%s,%s,%d,%.2f,%.4f",
	// 		i+1, nodeSpec.ProviderName, nodeSpec.RegionName, nodeSpec.CspSpecName, nodeSpec.VCPU, nodeSpec.MemoryGiB, nodeSpec.CostPerHour)
	// }

	log.Info().
		Str("machineId", node.MachineId).
		Int("recommendedSpecs", len(convertedNodeSpecList)).
		Msgf("Successfully recommended %d node specifications for machine: %s", len(convertedNodeSpecList), node.MachineId)

	return convertedNodeSpecList, numOfNodeSpecs, nil
}

// specRankingContext carries whatever inputs the ranking criteria below need. Add a field here -
// not a new sortByProximityWithCost parameter - when a future criterion needs new input.
type specRankingContext struct {
	vcpus        uint32
	memory       uint32
	csp          string
	cpuVendor    string
	vendorByName map[string]string // precomputed once per call, keyed by CspSpecName
}

// specCriterion compares two specs on one ranking dimension: negative if a ranks first, positive
// if b does, zero if tied (falls through to the next criterion in the chain).
type specCriterion func(ctx specRankingContext, a, b cloudmodel.SpecInfo) int

// vcpuProximity ranks the spec whose vCPU count is closer to the target first.
func vcpuProximity(ctx specRankingContext, a, b cloudmodel.SpecInfo) int {
	return int(abs(int32(a.VCPU)-int32(ctx.vcpus))) - int(abs(int32(b.VCPU)-int32(ctx.vcpus)))
}

// memoryProximity ranks the spec whose memory size is closer to the target first.
func memoryProximity(ctx specRankingContext, a, b cloudmodel.SpecInfo) int {
	return int(abs(int32(a.MemoryGiB)-int32(ctx.memory))) - int(abs(int32(b.MemoryGiB)-int32(ctx.memory)))
}

// manhattanProximity ranks by combined vCPU+memory distance (L1 norm), preferred over Euclidean
// distance here because vCPU and memory are independent resources on different scales.
func manhattanProximity(ctx specRankingContext, a, b cloudmodel.SpecInfo) int {
	da := abs(int32(a.VCPU)-int32(ctx.vcpus)) + abs(int32(a.MemoryGiB)-int32(ctx.memory))
	db := abs(int32(b.VCPU)-int32(ctx.vcpus)) + abs(int32(b.MemoryGiB)-int32(ctx.memory))
	return int(da - db)
}

// vendorMatch ranks the spec matching the source server's CPU vendor first; a no-op when the
// vendor is unknown on either side (see vendorRank).
func vendorMatch(ctx specRankingContext, a, b cloudmodel.SpecInfo) int {
	return vendorRank(ctx.cpuVendor, ctx.vendorByName, a) - vendorRank(ctx.cpuVendor, ctx.vendorByName, b)
}

// vendorRank returns 0 if spec's vendor matches sourceVendor, 1 otherwise (including unknown
// vendors) - unmatched and unknown are treated the same since we only have positive match evidence.
func vendorRank(sourceVendor string, vendorByName map[string]string, spec cloudmodel.SpecInfo) int {
	if sourceVendor == "" || vendorByName == nil {
		return 0
	}
	if vendorByName[spec.CspSpecName] == sourceVendor {
		return 0
	}
	return 1
}

// cpuVendorAliases maps a lowercase substring to its canonical vendor name; add an entry here to
// recognize a new vendor. "ampere" and "arm" both canonicalize to "arm" since Ampere chips are ARM-based.
var cpuVendorAliases = []struct {
	substring string
	vendor    string
}{
	{"intel", "intel"},
	{"amd", "amd"},
	{"ampere", "arm"},
	{"arm", "arm"},
}

// extractCpuVendor extracts a canonical vendor ("intel"/"amd"/"arm"/"") from a raw on-prem CPU
// vendor/model string (e.g. "GenuineIntel"), matching compat.GetCpuVendor's vocabulary.
func extractCpuVendor(vendor string) string {
	v := strings.ToLower(strings.TrimSpace(vendor))
	for _, alias := range cpuVendorAliases {
		if strings.Contains(v, alias.substring) {
			return alias.vendor
		}
	}
	return ""
}

// byCost ranks the cheaper spec first.
func byCost(ctx specRankingContext, a, b cloudmodel.SpecInfo) int {
	switch {
	case a.CostPerHour < b.CostPerHour:
		return -1
	case a.CostPerHour > b.CostPerHour:
		return 1
	default:
		return 0
	}
}

// rankSpecs sorts vmSpecs in place: the first criterion that isn't tied decides the order; ties
// cascade to the next criterion in the list.
func rankSpecs(ctx specRankingContext, vmSpecs []cloudmodel.SpecInfo, criteria ...specCriterion) {
	sort.Slice(vmSpecs, func(i, j int) bool {
		for _, c := range criteria {
			if d := c(ctx, vmSpecs[i], vmSpecs[j]); d != 0 {
				return d < 0
			}
		}
		return false
	})
}

// sortByProximityWithCost sorts VM specs by CPU vendor match with the source server (when known),
// then by proximity to the desired resource allocation, then by cost. When the source vendor is
// unknown, vendorMatch always ties (see vendorRank) and this reduces to proximity-then-cost.
func sortByProximityWithCost(vmSpecs []cloudmodel.SpecInfo, vcpus, memory uint32, csp string, cpuVendor string) {

	// Derive node's spec type (i.e. compute intensive type, memory intensive type, general purpose type)
	machineType := deriveMachineType(vcpus, memory)

	log.Debug().Msgf("Sorting VM specs for machine type: %s (vcpus: %d, memory: %d GiB)", machineType, vcpus, memory)

	ctx := specRankingContext{vcpus: vcpus, memory: memory, csp: csp, cpuVendor: cpuVendor}
	if cpuVendor != "" {
		ctx.vendorByName = make(map[string]string, len(vmSpecs))
		for _, s := range vmSpecs {
			ctx.vendorByName[s.CspSpecName] = compat.GetCpuVendor(csp, s)
		}
	}

	// Sort based on the machine type; within each type, vendor match dominates when the source
	// vendor is known, then proximity, then cost as the final tie-break.
	switch machineType {
	case "compute-intensive":
		rankSpecs(ctx, vmSpecs, vendorMatch, vcpuProximity, memoryProximity, byCost)
	case "memory-intensive":
		rankSpecs(ctx, vmSpecs, vendorMatch, memoryProximity, vcpuProximity, byCost)
	default: // "general-purpose"
		rankSpecs(ctx, vmSpecs, vendorMatch, manhattanProximity, byCost)
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
func calculateOptimalRange(vcpus uint32, memory uint32, rangeWeight int) (vcpusMin, vcpusMax, memoryMin, memoryMax uint32) {
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
		return calculateComputeIntensiveRange(vcpus, memory, rangeWeight)
	case memoryToVcpuRatio >= memoryIntensiveRatioThreshold: // Memory Intensive (1:8)
		return calculateMemoryIntensiveRange(vcpus, memory, rangeWeight)
	default: // General Purpose (1:4)
		return calculateGeneralPurposeRange(vcpus, memory, rangeWeight)
	}
}

func calculateComputeIntensiveRange(vcpus, memory uint32, rangeWeight int) (vcpusRangeMin, vcpusRangeMax, memoryRangeMin, memoryRangeMax uint32) {

	log.Debug().Msgf("Classified as Compute Intensive workload (vcpus: %d, memory: %d GiB)", vcpus, memory)

	const (
		memoryMultiplier             = 4 // Memory multiplier for max calculation
		primeNumSearchIterationCount = 2 // Number of prime number search iterations per direction (previously: prev-prev for min, next-next for max)
	)

	// Note: The value of 2 for primeNumSearchIterationCount was determined heuristically
	// to provide an optimal balance between range flexibility and recommendation precision.
	// This allows searching 2 prime numbers backward (for min) and forward (for max),
	// which empirically yields appropriate VM spec recommendations across various workloads.
	vcpusRangeMin = vcpus
	vcpusRangeMax = vcpus
	for i := 0; i < primeNumSearchIterationCount*rangeWeight; i++ {
		vcpusRangeMin = findPreviousPrimeOrDecrementOne(vcpusRangeMin)
		vcpusRangeMax = findNextPrimeNumber(vcpusRangeMax)
	}

	// Expand the range if it's too narrow
	if vcpusRangeMax-vcpus < 4 {
		vcpusRangeMax = findNextPrimeNumber(vcpusRangeMax)
	}

	// Set a wide search range for memory for compute-intensive workloads
	memoryRangeMin = 0
	memoryRangeMax = vcpusRangeMax * memoryMultiplier

	return vcpusRangeMin, vcpusRangeMax, memoryRangeMin, memoryRangeMax
}

func calculateMemoryIntensiveRange(vcpus, memory uint32, rangeWeight int) (vcpusMin, vcpusMax, memoryRangeMin, memoryRangeMax uint32) {

	log.Debug().Msgf("Classified as Memory Intensive workload (vcpus: %d, memory: %d GiB)", vcpus, memory)

	const (
		memoryToCpuRatio             = 7 // memory to CPU ratio for calculation (Standard: 8)
		primeNumSearchIterationCount = 2 // Number of prime number search iterations per direction (previously: prev-prev for min, next-next for max)
	)

	// Note: The value of 2 for primeNumSearchIterationCount was determined heuristically
	// to provide an optimal balance between range flexibility and recommendation precision.
	// This allows searching 2 prime numbers backward (for min) and forward (for max),
	// which empirically yields appropriate VM spec recommendations across various workloads.
	memoryRangeMin = memory
	memoryRangeMax = memory
	for i := 0; i < primeNumSearchIterationCount*rangeWeight; i++ {
		memoryRangeMin = findPreviousPrimeOrDecrementOne(memoryRangeMin)
		memoryRangeMax = findNextPrimeNumber(memoryRangeMax)
	}

	// Expand the range if it's too narrow
	if memoryRangeMax-memory < 4 {
		memoryRangeMax = findNextPrimeNumber(memoryRangeMax)
	}

	// Set a wide search range for vCPU for memory-intensive workloads
	vcpusMin = 0
	vcpusMax = memoryRangeMax / memoryToCpuRatio

	return vcpusMin, vcpusMax, memoryRangeMin, memoryRangeMax
}

func calculateGeneralPurposeRange(vcpus, memory uint32, rangeWeight int) (vcpusMin, vcpusMax, memoryMin, memoryMax uint32) {

	log.Debug().Msgf("Classified as General Purpose workload (vcpus: %d, memory: %d GiB)", vcpus, memory)
	// For General Purpose workloads, provide balanced flexibility for both vCPU and memory
	// The input has already been classified as General Purpose in calculateOptimalRange

	const primeNumSearchIterationCount = 2 // Number of prime number search iterations per direction (previously: prev-prev for min, next-next for max)

	// Note: The value of 2 for primeNumSearchIterationCount was determined heuristically
	// to provide an optimal balance between range flexibility and recommendation precision.
	// This allows searching 2 prime numbers backward (for min) and forward (for max),
	// which empirically yields appropriate VM spec recommendations across various workloads.
	vcpusMin = vcpus
	vcpusMax = vcpus
	for i := 0; i < primeNumSearchIterationCount*rangeWeight; i++ {
		vcpusMin = findPreviousPrimeOrDecrementOne(vcpusMin)
		vcpusMax = findNextPrimeNumber(vcpusMax)
	}

	// Expand the range if it's too narrow
	if vcpusMax-vcpus < 4 {
		vcpusMax = findNextPrimeNumber(vcpusMax)
	}

	memoryMin = memory
	memoryMax = memory
	for i := 0; i < primeNumSearchIterationCount*rangeWeight; i++ {
		memoryMin = findPreviousPrimeOrDecrementOne(memoryMin)
		memoryMax = findNextPrimeNumber(memoryMax)
	}
	// Expand the range if it's too narrow
	if memoryMax-memory < 4 {
		memoryMax = findNextPrimeNumber(memoryMax)
	}

	return vcpusMin, vcpusMax, memoryMin, memoryMax
}

// // calculateRangeMin calculates the minimum value for a range based on a given number
// func calculateRangeMin(n uint32) uint32 {

// 	// Calculate previous previous prime number
// 	min := findPreviousPrimeOrDecrementOne(n)
// 	min = findPreviousPrimeOrDecrementOne(min)

// 	return min
// }

// // calculateRangeMax calculates the maximum value for a range based on a given number
// func calculateRangeMax(n uint32) uint32 {

// 	// Calculate next next prime number
// 	max := findNextPrimeNumber(n)
// 	max = findNextPrimeNumber(max)

// 	// Expand the range if it's too narrow
// 	if max-n < 4 {
// 		max = findNextPrimeNumber(max)
// 	}

// 	return max
// }

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
