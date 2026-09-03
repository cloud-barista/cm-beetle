package recommendation

import (
	"encoding/json"
	"testing"

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
			name:     "nil GPU",
			node:     onpremmodel.NodeProperty{GPU: nil},
			expected: false,
		},
		{
			name: "GPU present but count 0",
			node: onpremmodel.NodeProperty{
				GPU: &onpremmodel.GpuProperty{Count: 0},
			},
			expected: false,
		},
		{
			name: "GPU present with count 1",
			node: onpremmodel.NodeProperty{
				GPU: &onpremmodel.GpuProperty{
					Count:         1,
					Vendor:        "NVIDIA",
					Model:         "Tesla T4",
					TotalMemoryGB: 16,
				},
			},
			expected: true,
		},
		{
			name: "GPU present with count 4",
			node: onpremmodel.NodeProperty{
				GPU: &onpremmodel.GpuProperty{
					Count:         4,
					Vendor:        "NVIDIA",
					Model:         "A100-SXM4-80GB",
					TotalMemoryGB: 320,
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
		GPU: &onpremmodel.GpuProperty{
			Count:         2,
			Vendor:        "NVIDIA",
			Model:         "A100",
			TotalMemoryGB: 80, // 40GB per GPU
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
		GPU: &onpremmodel.GpuProperty{
			Count:         2,
			Vendor:        "NVIDIA",
			TotalMemoryGB: 48, // 24GB per GPU
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
		GPU: &onpremmodel.GpuProperty{
			Count:         1,
			Vendor:        "NVIDIA",
			TotalMemoryGB: 80,
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
		GPU: &onpremmodel.GpuProperty{
			Count:         1,
			Vendor:        "NVIDIA",
			TotalMemoryGB: 32,
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
		GPU: &onpremmodel.GpuProperty{
			Count:         2,
			Vendor:        "NVIDIA",
			TotalMemoryGB: 48,
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
		GPU: &onpremmodel.GpuProperty{
			Count:         1,
			Vendor:        "NVIDIA",
			TotalMemoryGB: 16,
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
