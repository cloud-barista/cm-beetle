package recommendation

import (
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMatchRateVector_CpuOnlyNode_3Dimensions(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "cpu-srv-01",
		CPU: onpremmodel.CpuProperty{
			Cpus:         8,
			Threads:      1,
			Architecture: "amd64",
		},
		Memory: onpremmodel.MemoryProperty{
			TotalSize: 32,
		},
		OS: onpremmodel.OsProperty{
			ID:              "ubuntu",
			VersionID:       "22.04",
			VersionCodename: "jammy",
		},
	}

	spec := cloudmodel.SpecInfo{
		Id:          "aws+ap-northeast-2+m6i.2xlarge",
		CspSpecName: "m6i.2xlarge",
		VCPU:        8,
		MemoryGiB:   32,
	}

	image := cloudmodel.ImageInfo{
		Id:             "ami-ubuntu-2204",
		CspImageName:   "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "ubuntu jammy 22.04",
		IsGPUImage:     false,
	}

	vec := calculateMatchRateVector("aws", node, spec, image)

	assert.False(t, vec.HasGPU)
	require.Len(t, vec.ActiveDimensions(), 3)
	assert.InDelta(t, 100.0, vec.CPU, 0.01)
	assert.InDelta(t, 100.0, vec.Memory, 0.01)
	assert.InDelta(t, 100.0, vec.Image, 0.01)
	assert.InDelta(t, 100.0, vec.MinMatchRate(), 0.01)
	assert.InDelta(t, 100.0, vec.MaxMatchRate(), 0.01)
	assert.InDelta(t, 100.0, vec.AverageMatchRate(), 0.01)
	assert.True(t, vec.IsMatched(90.0))
	assert.Equal(t, "Match Rate: CPU=100.0% Memory=100.0% Image=100.0%", vec.FormatMatchRateSummary())
}

func TestMatchRateVector_GpuNode_6Dimensions_HighlyMatched(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "gpu-srv-01",
		CPU: onpremmodel.CpuProperty{
			Cpus:         16,
			Threads:      2,
			Architecture: "amd64",
		},
		Memory: onpremmodel.MemoryProperty{
			TotalSize: 128,
		},
		OS: onpremmodel.OsProperty{
			ID:              "ubuntu",
			VersionID:       "22.04",
			VersionCodename: "jammy",
		},
		GPUCards: []onpremmodel.GpuCardProperty{
			{
				DriverIndex:   "0",
				Vendor:        "NVIDIA",
				Model:         "NVIDIA A100-PCIE-40GB",
				MemoryTotalGB: 40.2,
			},
		},
	}

	spec := cloudmodel.SpecInfo{
		Id:                  "gcp+asia-northeast3+a2-highgpu-1g",
		CspSpecName:         "a2-highgpu-1g",
		VCPU:                32,
		MemoryGiB:           128,
		AcceleratorCount:    1,
		AcceleratorMemoryGB: 40,
		AcceleratorModel:    "NVIDIA A100 40GB",
	}

	// Note: vmImage does not contain "A100-PCIE-40GB", only IsGPUImage=true and Ubuntu 22.04 metadata.
	image := cloudmodel.ImageInfo{
		Id:             "gcp-ubuntu-2204-gpu",
		CspImageName:   "ubuntu-2204-jammy-v20260731",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "Ubuntu 22.04 jammy",
		IsGPUImage:     true,
	}

	vec := calculateMatchRateVector("gcp", node, spec, image)

	assert.True(t, vec.HasGPU)
	require.Len(t, vec.ActiveDimensions(), 6)
	assert.InDelta(t, 100.0, vec.CPU, 0.01)
	assert.InDelta(t, 100.0, vec.Memory, 0.01)
	assert.InDelta(t, 100.0, vec.Image, 0.01)
	assert.InDelta(t, 100.0, vec.GPUCount, 0.01)
	assert.InDelta(t, 100.0, vec.GPUVRAM, 0.01)
	assert.InDelta(t, 100.0, vec.GPUModel, 0.01)
	assert.True(t, vec.IsMatched(90.0))
	assert.Equal(
		t,
		"Match Rate: CPU=100.0% Memory=100.0% Image=100.0% GPUCount=100.0% GPUVRAM=100.0% GPUModel=100.0%",
		vec.FormatMatchRateSummary(),
	)
}

func TestCalculateGpuModelMatchRate_4Tiers(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "gpu-node",
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24},
		},
	}

	tests := []struct {
		name     string
		spec     cloudmodel.SpecInfo
		expected float64
	}{
		{
			name: "Tier 1: Same vendor and same canonical chip model (100%)",
			spec: cloudmodel.SpecInfo{
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 24,
				AcceleratorModel:    "NVIDIA L4",
			},
			expected: 100.0,
		},
		{
			name: "Tier 2: Same vendor but different chip model (75%)",
			spec: cloudmodel.SpecInfo{
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 24,
				AcceleratorModel:    "NVIDIA A10G",
			},
			expected: 75.0,
		},
		{
			name: "Tier 3: Cross-vendor GPU allocated (50%, non-zero)",
			spec: cloudmodel.SpecInfo{
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 192,
				AcceleratorModel:    "AMD Instinct MI300X",
			},
			expected: 50.0,
		},
		{
			name: "Tier 4: No GPU allocated on target spec (0%)",
			spec: cloudmodel.SpecInfo{
				AcceleratorCount:    0,
				AcceleratorMemoryGB: 0,
				AcceleratorModel:    "",
			},
			expected: 0.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			score := calculateGpuModelMatchRate(node, tc.spec)
			assert.InDelta(t, tc.expected, score, 0.01)
		})
	}
}

func TestCalculateGpuVramMatchRate_ToleranceAndRelative(t *testing.T) {
	node := onpremmodel.NodeProperty{
		MachineId: "gpu-vram-node",
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24.2},
		},
	}

	// Within 0.5 GB rounding tolerance (24.2 GB vs 24.0 GB) -> 100%
	specExactTier := cloudmodel.SpecInfo{AcceleratorCount: 1, AcceleratorMemoryGB: 24.0}
	assert.InDelta(t, 100.0, calculateGpuVramMatchRate(node, specExactTier), 0.01)

	// Target upsizing (24.0 GB source vs 48.0 GB target) -> 100.0% (satisfies VRAM requirement)
	nodeExact24 := onpremmodel.NodeProperty{
		GPUCards: []onpremmodel.GpuCardProperty{{Vendor: "NVIDIA", Model: "L4", MemoryTotalGB: 24.0}},
	}
	spec48 := cloudmodel.SpecInfo{AcceleratorCount: 1, AcceleratorMemoryGB: 48.0}
	assert.InDelta(t, 100.0, calculateGpuVramMatchRate(nodeExact24, spec48), 0.01)

	// Target downsizing (24.0 GB source vs 12.0 GB target) -> 50.0% relative match
	spec12 := cloudmodel.SpecInfo{AcceleratorCount: 1, AcceleratorMemoryGB: 12.0}
	assert.InDelta(t, 50.0, calculateGpuVramMatchRate(nodeExact24, spec12), 0.01)
}

func TestCalculateImageMatchRateScore_CpuVsGpuNodes(t *testing.T) {
	gpuNode := onpremmodel.NodeProperty{
		MachineId: "gpu-srv",
		CPU:       onpremmodel.CpuProperty{Architecture: "amd64"},
		OS:        onpremmodel.OsProperty{ID: "ubuntu", VersionID: "22.04", VersionCodename: "jammy"},
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "NVIDIA H100-80GB-HBM3", MemoryTotalGB: 80},
		},
	}

	nonGpuBasicImage := cloudmodel.ImageInfo{
		CspImageName:   "ubuntu-22.04-server-standard",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "Ubuntu 22.04 jammy",
		IsGPUImage:     false,
	}

	gpuEnabledImage := cloudmodel.ImageInfo{
		CspImageName:   "ubuntu-22.04-server-gpu",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "Ubuntu 22.04 jammy",
		IsGPUImage:     true,
	}

	// GPU node with non-GPU basic image is penalized by 0.5 (50%)
	scoreNonGpuImg := calculateImageMatchRateScore("aws", gpuNode, nonGpuBasicImage) * 100.0
	assert.InDelta(t, 50.0, scoreNonGpuImg, 0.01)

	// GPU node with GPU-enabled image scores 100% without penalty from raw card model string
	scoreGpuImg := calculateImageMatchRateScore("aws", gpuNode, gpuEnabledImage) * 100.0
	assert.InDelta(t, 100.0, scoreGpuImg, 0.01)
}

func TestCalculateCandidateMatchRateWithDetails_MixedCpuAndGpuNodes(t *testing.T) {
	// Composite infrastructure with 1 CPU-only node (3 dims) and 1 GPU node (6 dims) = 9 total dimensions
	cpuNode := onpremmodel.NodeProperty{
		MachineId: "web-cpu-01",
		CPU:       onpremmodel.CpuProperty{Cpus: 4, Threads: 1, Architecture: "amd64"},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 16},
		OS:        onpremmodel.OsProperty{ID: "ubuntu", VersionID: "22.04", VersionCodename: "jammy"},
	}
	gpuNode := onpremmodel.NodeProperty{
		MachineId: "ai-gpu-01",
		CPU:       onpremmodel.CpuProperty{Cpus: 8, Threads: 1, Architecture: "amd64"},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 32},
		OS:        onpremmodel.OsProperty{ID: "ubuntu", VersionID: "22.04", VersionCodename: "jammy"},
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24},
		},
	}

	srcInfra := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{cpuNode, gpuNode},
	}

	specs := []cloudmodel.SpecInfo{
		{
			Id:          "spec-cpu-exact",
			CspSpecName: "m6i.xlarge",
			VCPU:        4,
			MemoryGiB:   16,
		},
		{
			Id:                  "spec-gpu-a10g",
			CspSpecName:         "g5.2xlarge",
			VCPU:                8,
			MemoryGiB:           32,
			AcceleratorCount:    1,
			AcceleratorMemoryGB: 24,
			AcceleratorModel:    "NVIDIA A10G", // Same vendor, different chip model -> GPUModel = 75%
		},
	}

	images := []cloudmodel.ImageInfo{
		{
			Id:             "img-cpu",
			CspImageName:   "ubuntu-22.04-amd64",
			OSType:         "ubuntu 22.04",
			OSArchitecture: "amd64",
			OSDistribution: "ubuntu jammy 22.04",
			IsGPUImage:     false,
		},
		{
			Id:             "img-gpu",
			CspImageName:   "ubuntu-22.04-gpu-amd64",
			OSType:         "ubuntu 22.04",
			OSArchitecture: "amd64",
			OSDistribution: "ubuntu jammy 22.04",
			IsGPUImage:     true,
		},
	}

	nodeGroups := []cloudmodel.CreateNodeGroupReq{
		{Name: "ng-web", SpecId: "spec-cpu-exact", ImageId: "img-cpu"},
		{Name: "ng-ai", SpecId: "spec-gpu-a10g", ImageId: "img-gpu"},
	}

	status, desc, summary := calculateCandidateMatchRateWithDetails(
		"aws", nodeGroups, srcInfra, specs, images, 90.0,
	)

	// CPU node is 100% across 3 dims (sum = 300).
	// GPU node has 100% on 5 dims and 75% on GPUModel (sum = 575).
	// Total across 9 active dimensions = 875 / 9 = 97.222%.
	// Because GPUModel (75%) < 90%, GPU node is not matched -> overall status must be "partially-matched".
	assert.Equal(t, "partially-matched", status)
	assert.Equal(t, "VMs: 2 total, 1 matched, 1 acceptable", desc)
	assert.Equal(t, 2, summary.TotalVMs)
	assert.InDelta(t, 75.0, summary.MinMatchRate, 0.01)
	assert.InDelta(t, 100.0, summary.MaxMatchRate, 0.01)
	assert.InDelta(t, 875.0/9.0, summary.AvgMatchRate, 0.01)
	assert.InDelta(t, 50.0, summary.BestEffortRate, 0.01)
}

func TestMatchRateVector_GpuNode_StrictAllDimensionParity(t *testing.T) {
	gpuNode := onpremmodel.NodeProperty{
		MachineId: "gpu-strict-node",
		CPU:       onpremmodel.CpuProperty{Cpus: 8, Threads: 1, Architecture: "amd64"},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 32},
		OS:        onpremmodel.OsProperty{ID: "ubuntu", VersionID: "22.04", VersionCodename: "jammy"},
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24},
		},
	}

	gpuImage := cloudmodel.ImageInfo{
		Id:             "img-gpu",
		CspImageName:   "ubuntu-22.04-gpu",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "ubuntu jammy 22.04",
		IsGPUImage:     true,
	}

	// Host Memory mismatch (32 GiB vs 64 GiB = 50%) despite 100% GPU match must be partially-matched
	specHostMemMismatch := cloudmodel.SpecInfo{
		Id:                  "spec-mem-64",
		CspSpecName:         "g6.4xlarge",
		VCPU:                8,
		MemoryGiB:           64,
		AcceleratorCount:    1,
		AcceleratorMemoryGB: 24,
		AcceleratorModel:    "NVIDIA L4",
	}
	vecMem := calculateMatchRateVector("aws", gpuNode, specHostMemMismatch, gpuImage)
	assert.InDelta(t, 50.0, vecMem.Memory, 0.01)
	assert.InDelta(t, 100.0, vecMem.GPUCount, 0.01)
	assert.InDelta(t, 100.0, vecMem.GPUVRAM, 0.01)
	assert.InDelta(t, 100.0, vecMem.GPUModel, 0.01)
	assert.False(t, vecMem.IsMatched(90.0))

	// Cross-vendor GPU (AMD instead of NVIDIA = 50% GPUModel) keeps MinMatchRate >= 50% (never 0%)
	specCrossVendor := cloudmodel.SpecInfo{
		Id:                  "spec-amd",
		CspSpecName:         "g4ad.2xlarge",
		VCPU:                8,
		MemoryGiB:           32,
		AcceleratorCount:    1,
		AcceleratorMemoryGB: 24,
		AcceleratorModel:    "AMD Radeon Pro V520",
	}
	vecCross := calculateMatchRateVector("aws", gpuNode, specCrossVendor, gpuImage)
	assert.InDelta(t, 50.0, vecCross.GPUModel, 0.01)
	assert.InDelta(t, 50.0, vecCross.MinMatchRate(), 0.01)
	assert.False(t, vecCross.IsMatched(90.0))
}

func TestSynthesizeGroupRepresentativeNode_GpuUpsizing(t *testing.T) {
	// cpu-heavy has 16 vCPU (no GPU), gpu-member has 8 vCPU and 2x A100 80GB
	nodeByMachineId := map[string]onpremmodel.NodeProperty{
		"cpu-heavy": {
			MachineId: "cpu-heavy",
			CPU:       onpremmodel.CpuProperty{Cpus: 16, Threads: 1, Architecture: "amd64"},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 64},
		},
		"gpu-member": {
			MachineId: "gpu-member",
			CPU:       onpremmodel.CpuProperty{Cpus: 8, Threads: 1, Architecture: "amd64"},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 128},
			GPUCards: []onpremmodel.GpuCardProperty{
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
			},
		},
	}

	synth := synthesizeGroupRepresentativeNode([]string{"cpu-heavy", "gpu-member"}, nodeByMachineId)
	assert.Equal(t, uint32(16), synth.CPU.Cpus)
	assert.Equal(t, uint64(128), synth.Memory.TotalSize)
	require.Len(t, synth.GPUCards, 2)
	assert.Equal(t, "A100", synth.GPUCards[0].Model)
	assert.Equal(t, float32(80), synth.GPUCards[0].MemoryTotalGB)
}

func TestCalculateGpuVramMatchRate_UpsizingAndDownsizing(t *testing.T) {
	node16GB := onpremmodel.NodeProperty{
		MachineId: "gpu-node-16gb",
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "T4", MemoryTotalGB: 16},
		},
	}

	// Case 1: Exact match -> 100%
	spec16GB := cloudmodel.SpecInfo{AcceleratorCount: 1, AcceleratorMemoryGB: 16}
	assert.InDelta(t, 100.0, calculateGpuVramMatchRate(node16GB, spec16GB), 0.01)

	// Case 2: Upsizing (16GB -> 24GB) -> 100% (sufficient VRAM to prevent OOM)
	spec24GB := cloudmodel.SpecInfo{AcceleratorCount: 1, AcceleratorMemoryGB: 24}
	assert.InDelta(t, 100.0, calculateGpuVramMatchRate(node16GB, spec24GB), 0.01)

	// Case 3: Downsizing (24GB -> 16GB) -> 66.67%
	node24GB := onpremmodel.NodeProperty{
		MachineId: "gpu-node-24gb",
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "L4", MemoryTotalGB: 24},
		},
	}
	assert.InDelta(t, 66.67, calculateGpuVramMatchRate(node24GB, spec16GB), 0.01)

	// Case 4: No GPU in target spec -> 0%
	spec0GB := cloudmodel.SpecInfo{AcceleratorCount: 0, AcceleratorMemoryGB: 0}
	assert.InDelta(t, 0.0, calculateGpuVramMatchRate(node24GB, spec0GB), 0.01)
}

func TestFindAndSortVmOsImageInfoListBySimilarity_GpuSuitabilityPriority(t *testing.T) {
	gpuNode := onpremmodel.NodeProperty{
		MachineId: "gpu-worker",
		OS: onpremmodel.OsProperty{
			ID:              "ubuntu",
			VersionID:       "22.04",
			VersionCodename: "jammy",
		},
		CPU: onpremmodel.CpuProperty{Architecture: "amd64"},
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "L4", MemoryTotalGB: 24},
		},
	}

	standardImg := cloudmodel.ImageInfo{
		Id:             "img-standard",
		CspImageName:   "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "ubuntu jammy 22.04",
		IsGPUImage:     false,
	}

	gpuImg := cloudmodel.ImageInfo{
		Id:             "img-gpu-driver",
		CspImageName:   "Deep Learning OSS Nvidia Driver AMI GPU PyTorch (Ubuntu 22.04)",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "ubuntu jammy 22.04 nvidia cuda",
		IsGPUImage:     true,
	}

	keywords, kwDelimiters, imgDelimiters := SetKeywordsAndDelimeters(gpuNode)

	// When evaluating for GPU node, GPU-suitable image must sort before standard image.
	sortedForGpu := FindAndSortVmOsImageInfoListBySimilarity("aws", keywords, kwDelimiters, []cloudmodel.ImageInfo{standardImg, gpuImg}, imgDelimiters, gpuNode)
	require.Len(t, sortedForGpu, 2)
	assert.Equal(t, "img-gpu-driver", sortedForGpu[0].Id)

	// For CPU-only node, canonical shorter standard base image is preferred.
	cpuNode := gpuNode
	cpuNode.GPUCards = nil
	sortedForCpu := FindAndSortVmOsImageInfoListBySimilarity("aws", keywords, kwDelimiters, []cloudmodel.ImageInfo{standardImg, gpuImg}, imgDelimiters, cpuNode)
	require.Len(t, sortedForCpu, 2)
	assert.Equal(t, "img-standard", sortedForCpu[0].Id)
}
