package recommendation

import (
	"encoding/json"
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasGpu(t *testing.T) {
	tests := []struct {
		name     string
		node     onpremmodel.NodeProperty
		expected bool
	}{
		{
			name:     "nil or empty GPUCards",
			node:     onpremmodel.NodeProperty{GPUCards: nil},
			expected: false,
		},
		{
			name:     "empty GPUCards slice",
			node:     onpremmodel.NodeProperty{GPUCards: []onpremmodel.GpuCardProperty{}},
			expected: false,
		},
		{
			name: "GPU cards present with 1 card",
			node: onpremmodel.NodeProperty{
				GPUCards: []onpremmodel.GpuCardProperty{
					{
						DriverIndex:   "0",
						Vendor:        "NVIDIA",
						Model:         "Tesla T4",
						MemoryTotalGB: 16,
					},
				},
			},
			expected: true,
		},
		{
			name: "GPU cards present with 4 cards",
			node: onpremmodel.NodeProperty{
				GPUCards: []onpremmodel.GpuCardProperty{
					{DriverIndex: "0", Vendor: "NVIDIA", Model: "A100-SXM4-80GB", MemoryTotalGB: 80},
					{DriverIndex: "1", Vendor: "NVIDIA", Model: "A100-SXM4-80GB", MemoryTotalGB: 80},
					{DriverIndex: "2", Vendor: "NVIDIA", Model: "A100-SXM4-80GB", MemoryTotalGB: 80},
					{DriverIndex: "3", Vendor: "NVIDIA", Model: "A100-SXM4-80GB", MemoryTotalGB: 80},
				},
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, hasGpu(tc.node))
		})
	}
}

func TestDetectGpuVendor(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"NVIDIA", GpuVendorNVIDIA},
		{"nvidia", GpuVendorNVIDIA},
		{"Tesla T4", GpuVendorNVIDIA},
		{"A100-SXM4-80GB", GpuVendorNVIDIA},
		{"NVIDIA A10G", GpuVendorNVIDIA},
		{"NVIDIA L4", GpuVendorNVIDIA},
		{"GeForce RTX 4090", GpuVendorNVIDIA},
		{"Quadro RTX 6000", GpuVendorNVIDIA},
		{"Tesla V100-PCIE-32GB", GpuVendorNVIDIA},
		{"AMD", GpuVendorAMD},
		{"Radeon Pro V520", GpuVendorAMD},
		{"AMD INSTINCT MI300X", GpuVendorAMD},
		{"Instinct MI350", GpuVendorAMD},
		{"AMD Radeon Pro V620", GpuVendorAMD},
		{"Intel", GpuVendorIntel},
		{"Intel Gaudi3", GpuVendorIntel},
		{"Intel Arc Pro A60", GpuVendorIntel},
		{"Google TPU v4", GpuVendorGoogle},
		{"TPU7X", GpuVendorGoogle},
		{"UnknownVendor", GpuVendorOther},
		{"", GpuVendorOther},
		{"   ", GpuVendorOther},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.expected, detectGpuVendor(tc.input))
		})
	}
}

func TestRightSizeUpVramTier(t *testing.T) {
	tests := []struct {
		input    float32
		expected float32
	}{
		{4, 8},
		{8, 8},
		{12, 16},
		{16, 16},
		{20, 24},
		{24, 24},
		{32, 32},
		{40, 40},
		{45, 48},
		{70, 80},
		{80, 80},
		{100, 144},
		{200, 288},
		{300, 300}, // Beyond max tier: returns requested
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, rightSizeUpVramTier(tc.input))
	}
}

func TestRightSizeUpGpuCount(t *testing.T) {
	assert.Equal(t, uint32(1), rightSizeUpGpuCount(0))
	assert.Equal(t, uint32(1), rightSizeUpGpuCount(1))
	assert.Equal(t, uint32(2), rightSizeUpGpuCount(2))
	assert.Equal(t, uint32(4), rightSizeUpGpuCount(3)) // 3 -> 4
	assert.Equal(t, uint32(4), rightSizeUpGpuCount(4))
	assert.Equal(t, uint32(8), rightSizeUpGpuCount(5)) // 5 -> 8
	assert.Equal(t, uint32(8), rightSizeUpGpuCount(6)) // 6 -> 8
	assert.Equal(t, uint32(8), rightSizeUpGpuCount(7)) // 7 -> 8
	assert.Equal(t, uint32(8), rightSizeUpGpuCount(8))
}

func TestClusterGpuCards(t *testing.T) {
	// Mixed cards: 2x NVIDIA A100 40GB, 1x NVIDIA T4 16GB, 1x AMD MI350 288GB
	cards := []onpremmodel.GpuCardProperty{
		{DriverIndex: "0", Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
		{DriverIndex: "1", Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
		{DriverIndex: "2", Vendor: "NVIDIA", Model: "T4", MemoryTotalGB: 16},
		{DriverIndex: "card0", Vendor: "AMD", Model: "MI350", MemoryTotalGB: 288},
	}

	clusters := clusterGpuCards(cards)
	require.Len(t, clusters, 3)

	// Primary cluster must have highest Count (2x A100)
	assert.Equal(t, "NVIDIA", clusters[0].Vendor)
	assert.Equal(t, "A100", clusters[0].Model)
	assert.Equal(t, uint32(2), clusters[0].Count)
	assert.Equal(t, float32(40), clusters[0].MemoryTotalGB)

	// Second cluster tie-broken by MemoryTotalGB desc (AMD MI350 288GB vs T4 16GB)
	assert.Equal(t, "AMD", clusters[1].Vendor)
	assert.Equal(t, "MI350", clusters[1].Model)
	assert.Equal(t, uint32(1), clusters[1].Count)
	assert.Equal(t, float32(288), clusters[1].MemoryTotalGB)

	// Third cluster (T4 16GB)
	assert.Equal(t, "NVIDIA", clusters[2].Vendor)
	assert.Equal(t, "T4", clusters[2].Model)
	assert.Equal(t, uint32(1), clusters[2].Count)
	assert.Equal(t, float32(16), clusters[2].MemoryTotalGB)
}

func TestBuildGpuDeploymentPlan(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "gpu-node-01",
		CPU: onpremmodel.CpuProperty{
			Cpus:         8,
			Threads:      2,
			Architecture: "x86_64",
		},
		Memory: onpremmodel.MemoryProperty{
			TotalSize: 64,
		},
		GPUCards: []onpremmodel.GpuCardProperty{
			{DriverIndex: "0", Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
			{DriverIndex: "1", Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
		},
	}

	plan, vcpusMin, vcpusMax, memMin, memMax := buildGpuDeploymentPlan(
		node, "aws", "ap-northeast-2", "x86_64", 16, 64, 1, 30,
	)

	assert.NotEmpty(t, plan)
	assert.Greater(t, vcpusMax, vcpusMin)
	assert.Greater(t, memMax, memMin)

	// Validate that the output is valid JSON
	var parsed map[string]interface{}
	err := json.Unmarshal([]byte(plan), &parsed)
	require.NoError(t, err, "GPU deployment plan must be valid JSON")

	filter, ok := parsed["filter"].(map[string]interface{})
	require.True(t, ok)
	policy, ok := filter["policy"].([]interface{})
	require.True(t, ok)

	// Verify required metrics are present
	metricsFound := make(map[string]bool)
	for _, p := range policy {
		m, ok := p.(map[string]interface{})
		if ok {
			metricName, ok := m["metric"].(string)
			if ok {
				metricsFound[metricName] = true
			}
		}
	}

	assert.True(t, metricsFound["acceleratorCount"], "must have acceleratorCount filter")
	assert.True(t, metricsFound["acceleratorMemoryGB"], "must have acceleratorMemoryGB filter")
	assert.True(t, metricsFound["vCPU"], "must have vCPU filter")
	assert.True(t, metricsFound["memoryGiB"], "must have memoryGiB filter")
	assert.True(t, metricsFound["providerName"], "must have providerName filter")
	assert.True(t, metricsFound["regionName"], "must have regionName filter")
	assert.True(t, metricsFound["architecture"], "must have architecture filter")
}

func TestSortGpuByProximityWithCost_CountProximity(t *testing.T) {
	// Source node wants 2 GPUs
	node := onpremmodel.NodeProperty{
		MachineId: "node-2gpu",
		CPU:       onpremmodel.CpuProperty{Cpus: 4, Threads: 2},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 32},
		GPUCards: []onpremmodel.GpuCardProperty{
			{DriverIndex: "0", Vendor: "NVIDIA", MemoryTotalGB: 24},
			{DriverIndex: "1", Vendor: "NVIDIA", MemoryTotalGB: 24},
		},
	}

	specs := []cloudmodel.SpecInfo{
		{
			CspSpecName:         "spec-4gpu",
			VCPU:                16,
			MemoryGiB:           64,
			AcceleratorCount:    4,
			AcceleratorMemoryGB: 24,
			AcceleratorModel:    "NVIDIA A10G",
			CostPerHour:         3.0, // Cheaper than 2-GPU spec to test that count proximity dominates
		},
		{
			CspSpecName:         "spec-2gpu-exact",
			VCPU:                12,
			MemoryGiB:           48,
			AcceleratorCount:    2,
			AcceleratorMemoryGB: 24,
			AcceleratorModel:    "NVIDIA A10G",
			CostPerHour:         5.0,
		},
		{
			CspSpecName:         "spec-1gpu",
			VCPU:                8,
			MemoryGiB:           32,
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 24,
			AcceleratorModel:    "NVIDIA A10G",
			CostPerHour:         1.5,
		},
	}

	sortGpuByProximityWithCost(specs, node, "aws")

	// Exact match (2 GPUs) must rank first
	assert.Equal(t, "spec-2gpu-exact", specs[0].CspSpecName, "Exact GPU count match must rank #1")
}

func TestSortGpuByProximityWithCost_VramProximity(t *testing.T) {
	// Source node wants 1 GPU with 80GB VRAM
	node := onpremmodel.NodeProperty{
		MachineId: "node-a100-80gb",
		CPU:       onpremmodel.CpuProperty{Cpus: 8, Threads: 2},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 64},
		GPUCards: []onpremmodel.GpuCardProperty{
			{DriverIndex: "0", Vendor: "NVIDIA", MemoryTotalGB: 80},
		},
	}

	specs := []cloudmodel.SpecInfo{
		{
			CspSpecName:         "spec-40gb",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 40,
			AcceleratorModel:    "NVIDIA A100-40GB",
			CostPerHour:         4.0, // Cheaper
		},
		{
			CspSpecName:         "spec-80gb-exact",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 80,
			AcceleratorModel:    "NVIDIA A100-80GB",
			CostPerHour:         6.0,
		},
		{
			CspSpecName:         "spec-96gb",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 96,
			AcceleratorModel:    "NVIDIA H100",
			CostPerHour:         9.0,
		},
	}

	sortGpuByProximityWithCost(specs, node, "gcp")

	// Exact 80GB must win over 40GB and 96GB
	assert.Equal(t, "spec-80gb-exact", specs[0].CspSpecName, "Exact VRAM match must rank #1")
}

func TestSortGpuByProximityWithCost_VendorMatch(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "node-nvidia",
		CPU:       onpremmodel.CpuProperty{Cpus: 8, Threads: 2},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 64},
		GPUCards: []onpremmodel.GpuCardProperty{
			{DriverIndex: "0", Vendor: "NVIDIA", MemoryTotalGB: 32},
		},
	}

	specs := []cloudmodel.SpecInfo{
		{
			CspSpecName:         "spec-amd-cheaper",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 32,
			AcceleratorModel:    "AMD Instinct MI210",
			CostPerHour:         3.0,
		},
		{
			CspSpecName:         "spec-nvidia-match",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 32,
			AcceleratorModel:    "NVIDIA V100 32GB",
			CostPerHour:         4.0,
		},
	}

	sortGpuByProximityWithCost(specs, node, "azure")

	assert.Equal(t, "spec-nvidia-match", specs[0].CspSpecName, "Vendor match (NVIDIA) must rank #1 over mismatched vendor")
}

func TestSortGpuByProximityWithCost_VendorDominatesCountAndCost(t *testing.T) {
	// Source node wants 2 NVIDIA GPUs
	node := onpremmodel.NodeProperty{
		MachineId: "node-nvidia-2gpu",
		CPU:       onpremmodel.CpuProperty{Cpus: 8, Threads: 2},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 64},
		GPUCards: []onpremmodel.GpuCardProperty{
			{DriverIndex: "0", Vendor: "NVIDIA", MemoryTotalGB: 24},
			{DriverIndex: "1", Vendor: "NVIDIA", MemoryTotalGB: 24},
		},
	}

	specs := []cloudmodel.SpecInfo{
		{
			CspSpecName:         "spec-amd-exact-count-cheaper",
			AcceleratorCount:    2,  // Exact count match!
			AcceleratorMemoryGB: 24, // Exact VRAM!
			AcceleratorModel:    "AMD Instinct MI210",
			CostPerHour:         2.0, // Cheaper!
		},
		{
			CspSpecName:         "spec-nvidia-4gpu-expensive",
			AcceleratorCount:    4,  // Inexact count
			AcceleratorMemoryGB: 24,
			AcceleratorModel:    "NVIDIA A10G",
			CostPerHour:         6.0, // More expensive
		},
	}

	sortGpuByProximityWithCost(specs, node, "aws")

	// NVIDIA must rank #1 even though AMD had exact count and was cheaper,
	// because driver/CUDA binary compatibility is an absolute prerequisite.
	assert.Equal(t, "spec-nvidia-4gpu-expensive", specs[0].CspSpecName, "Vendor match (NVIDIA) must dominate count proximity and cost")
}

func TestSortGpuByProximityWithCost_CostTieBreak(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "node-tie",
		CPU:       onpremmodel.CpuProperty{Cpus: 4, Threads: 1},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 16},
		GPUCards: []onpremmodel.GpuCardProperty{
			{DriverIndex: "0", Vendor: "NVIDIA", MemoryTotalGB: 16},
		},
	}

	specs := []cloudmodel.SpecInfo{
		{
			CspSpecName:         "spec-expensive",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 16,
			AcceleratorModel:    "Tesla T4",
			VCPU:                4,
			MemoryGiB:           16,
			CostPerHour:         1.2,
		},
		{
			CspSpecName:         "spec-cheaper",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 16,
			AcceleratorModel:    "Tesla T4",
			VCPU:                4,
			MemoryGiB:           16,
			CostPerHour:         0.8,
		},
	}

	sortGpuByProximityWithCost(specs, node, "aws")

	assert.Equal(t, "spec-cheaper", specs[0].CspSpecName, "Cheaper spec must win tie-break")
	assert.Equal(t, "spec-expensive", specs[1].CspSpecName)
}

func TestDetectGpuModel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"NVIDIA L4", "l4"},
		{"NVIDIA L40S", "l40s"},
		{"NVIDIA A10G", "a10g"},
		{"A100-SXM4-80GB", "a100"},
		{"Tesla T4", "t4"},
		{"NVIDIA H100 80GB", "h100"},
		{"AMD Instinct MI300X", "mi300"},
		{"Unknown GPU", ""},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, detectGpuModel(tc.input))
	}
}

func TestSortGpuByProximityWithCost_ModelMatch(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "node-l4-32vcpu",
		CPU:       onpremmodel.CpuProperty{Cpus: 1, Threads: 32},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 128},
		GPUCards: []onpremmodel.GpuCardProperty{
			{DriverIndex: "0", Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24},
		},
	}

	specs := []cloudmodel.SpecInfo{
		{
			CspSpecName:         "g5.8xlarge",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 24,
			AcceleratorModel:    "NVIDIA A10G",
			VCPU:                32,
			MemoryGiB:           128,
			CostPerHour:         1.8,
		},
		{
			CspSpecName:         "g6.8xlarge",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 24,
			AcceleratorModel:    "NVIDIA L4",
			VCPU:                32,
			MemoryGiB:           128,
			CostPerHour:         2.0144,
		},
	}

	sortGpuByProximityWithCost(specs, node, "aws")

	// Exact GPU chip model (NVIDIA L4 -> g6.8xlarge) must rank ahead of A10G (g5.8xlarge).
	assert.Equal(t, "g6.8xlarge", specs[0].CspSpecName)
}

func TestSortGpuByProximityWithCost_VramSufficiency(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "node-24gb",
		CPU:       onpremmodel.CpuProperty{Cpus: 1, Threads: 8},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 32},
		GPUCards: []onpremmodel.GpuCardProperty{
			{DriverIndex: "0", Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24},
		},
	}

	specs := []cloudmodel.SpecInfo{
		{
			CspSpecName:         "spec-16gb-insufficient",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 16,
			AcceleratorModel:    "Tesla T4",
			VCPU:                8,
			MemoryGiB:           32,
			CostPerHour:         0.7,
		},
		{
			CspSpecName:         "spec-32gb-sufficient",
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 32,
			AcceleratorModel:    "Tesla V100",
			VCPU:                8,
			MemoryGiB:           32,
			CostPerHour:         3.0,
		},
	}

	sortGpuByProximityWithCost(specs, node, "aws")

	// 32GB (sufficient VRAM) must beat 16GB (insufficient VRAM) despite equal 8GB absolute difference.
	assert.Equal(t, "spec-32gb-sufficient", specs[0].CspSpecName)
}

func TestResolveGpuAcceleratorModel(t *testing.T) {
	ncpL4Spec := cloudmodel.SpecInfo{
		CspSpecName:      "gp1l4-g3",
		AcceleratorModel: "NVIDIA NA",
		Details: []cloudmodel.KeyValue{
			{Key: "ServerSpecDescription", Value: "NVIDIA L4 GPU 1EA, GPUMemory 24GB, vCPU 4EA, Memory 48GB"},
		},
	}
	ncpV100Spec := cloudmodel.SpecInfo{
		CspSpecName:      "gp1v8-g1-h50",
		AcceleratorModel: "NVIDIA NA",
		Details: []cloudmodel.KeyValue{
			{Key: "ServerSpecDescription", Value: "Tesla V100 GPU 1EA, GPUMemory 32GB, vCPU 8EA, Memory 90GB, Disk 50GB"},
		},
	}
	normalSpec := cloudmodel.SpecInfo{
		CspSpecName:      "g4dn.xlarge",
		AcceleratorModel: "NVIDIA T4",
	}

	specs := []cloudmodel.SpecInfo{ncpL4Spec, ncpV100Spec, normalSpec}
	sanitizeGpuSpecsMetadata(specs)

	assert.Equal(t, "NVIDIA L4", specs[0].AcceleratorModel)
	assert.Equal(t, "Tesla V100", specs[1].AcceleratorModel)
	assert.Equal(t, "NVIDIA T4", specs[2].AcceleratorModel)
}

func TestNcpKvmAndPricedSpecsRelaxation(t *testing.T) {
	rawSpecs := []tbmodel.SpecInfo{
		{
			CspSpecName: "gp1t4-g1-h50",
			CostPerHour: 1.1567,
			Details:     []tbmodel.KeyValue{{Key: "HypervisorType", Value: "XEN"}},
		},
		{
			CspSpecName: "gp1l4-g3",
			CostPerHour: -1,
			Details:     []tbmodel.KeyValue{{Key: "HypervisorType", Value: "KVM"}},
		},
		{
			CspSpecName: "gp1l8-g3",
			CostPerHour: -1,
			Details:     []tbmodel.KeyValue{{Key: "HypervisorType", Value: "KVM"}},
		},
	}

	// Step 1: NCP KVM filter drops XEN T4 specs, leaving KVM L4 specs.
	kvmSpecs := filterNcpKvmTbSpecs("kr", rawSpecs)
	require.Len(t, kvmSpecs, 2)
	assert.Equal(t, "gp1l4-g3", kvmSpecs[0].CspSpecName)

	// Step 2: Because all remaining KVM specs have CostPerHour == -1, filterPricedSpecs retains them.
	validSpecs := filterPricedSpecs(kvmSpecs)
	require.Len(t, validSpecs, 2)
	assert.Equal(t, "gp1l4-g3", validSpecs[0].CspSpecName)

	// Once Tumblebug populates pricing for gp1l4-g3, strict cost filtering applies automatically.
	kvmSpecs[0].CostPerHour = 1.85
	strictSpecs := filterPricedSpecs(kvmSpecs)
	require.Len(t, strictSpecs, 1)
	assert.Equal(t, "gp1l4-g3", strictSpecs[0].CspSpecName)
}

func TestRecommendGpuNodeSpec_LiveTumblebug(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "gpu-node-live-test",
		CPU: onpremmodel.CpuProperty{
			Cpus:         8,
			Threads:      1,
			Architecture: "x86_64",
		},
		Memory: onpremmodel.MemoryProperty{
			TotalSize: 32,
		},
		GPUCards: []onpremmodel.GpuCardProperty{
			{
				DriverIndex:   "0",
				Vendor:        "NVIDIA",
				Model:         "Tesla T4",
				MemoryTotalGB: 16,
			},
		},
	}

	specs, count, err := RecommendNodeSpecs("aws", "ap-northeast-2", node, 3)
	if err != nil {
		t.Skipf("Skipping live Tumblebug test (daemon not reachable or error): %v", err)
		return
	}

	if count == 0 {
		t.Skip("Skipping assertion: no specs returned from live Tumblebug in region")
		return
	}

	assert.NotEmpty(t, specs)
	assert.GreaterOrEqual(t, count, 1)
	assert.Equal(t, GpuVendorNVIDIA, detectGpuVendor(specs[0].AcceleratorModel))
	assert.GreaterOrEqual(t, specs[0].AcceleratorCount, uint8(1))
}
