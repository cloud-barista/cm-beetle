package recommendation

import (
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMatchRate_TeslaT4_SupportedCSPs validates 6-dimension evaluation for an NVIDIA Tesla T4 node.
func TestMatchRate_TeslaT4_SupportedCSPs(t *testing.T) {
	nodeT4 := onpremmodel.NodeProperty{
		MachineId: "t4-srv-01",
		CPU: onpremmodel.CpuProperty{
			Cpus:         4,
			Threads:      1,
			Architecture: "amd64",
		},
		Memory: onpremmodel.MemoryProperty{
			TotalSize: 16,
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
				Model:         "Tesla T4",
				Type:          "GPU",
				MemoryTotalGB: 16.0,
			},
		},
	}

	standardGpuImage := cloudmodel.ImageInfo{
		Id:             "img-gpu-ubuntu-2204",
		CspImageName:   "ubuntu-22.04-cuda12-gpu",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "ubuntu jammy 22.04 nvidia cuda",
		IsGPUImage:     true,
	}

	azureGpuImage := cloudmodel.ImageInfo{
		Id:             "img-azure-ubuntu-2204",
		CspImageName:   "Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:latest",
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "ubuntu jammy 22.04 nvidia cuda",
		IsGPUImage:     true,
	}

	cases := []struct {
		csp             string
		caseName        string
		spec            cloudmodel.SpecInfo
		customImage     *cloudmodel.ImageInfo
		expectedCPU     float64
		expectedMemory  float64
		expectedImage   float64
		expectedGPUCnt  float64
		expectedGPUVram float64
		expectedGPUMod  float64
	}{
		{
			csp:      "aws",
			caseName: "AWS g4dn.xlarge (Tesla T4 16GB - Tier 1 exact match)",
			spec: cloudmodel.SpecInfo{
				Id:                  "aws+ap-northeast-2+g4dn.xlarge",
				CspSpecName:         "g4dn.xlarge",
				VCPU:                4,
				MemoryGiB:           16,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 16,
				AcceleratorModel:    "NVIDIA Tesla T4",
			},
			expectedCPU:     100.0,
			expectedMemory:  100.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 100.0,
			expectedGPUMod:  100.0,
		},
		{
			csp:      "aws",
			caseName: "AWS g5.2xlarge (A10G 24GB - Tier 2 same vendor, VRAM upsizing)",
			spec: cloudmodel.SpecInfo{
				Id:                  "aws+ap-northeast-2+g5.2xlarge",
				CspSpecName:         "g5.2xlarge",
				VCPU:                8,
				MemoryGiB:           32,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 24,
				AcceleratorModel:    "NVIDIA A10G",
			},
			expectedCPU:     50.0,
			expectedMemory:  50.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 100.0,
			expectedGPUMod:  75.0,
		},
		{
			csp:         "azure",
			caseName:    "Azure Standard_NC4as_T4_v3 (Tesla T4 16GB - Tier 1 exact match)",
			customImage: &azureGpuImage,
			spec: cloudmodel.SpecInfo{
				Id:                  "azure+koreasouth+Standard_NC4as_T4_v3",
				CspSpecName:         "Standard_NC4as_T4_v3",
				VCPU:                4,
				MemoryGiB:           28,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 16,
				AcceleratorModel:    "NVIDIA Tesla T4",
			},
			expectedCPU:     100.0,
			expectedMemory:  (16.0 / 28.0) * 100.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 100.0,
			expectedGPUMod:  100.0,
		},
		{
			csp:      "gcp",
			caseName: "GCP g2-standard-8 (L4 24GB - Tier 2 same vendor, VRAM upsizing)",
			spec: cloudmodel.SpecInfo{
				Id:                  "gcp+asia-northeast3+g2-standard-8",
				CspSpecName:         "g2-standard-8",
				VCPU:                8,
				MemoryGiB:           32,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 24,
				AcceleratorModel:    "NVIDIA L4",
			},
			expectedCPU:     50.0,
			expectedMemory:  50.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 100.0,
			expectedGPUMod:  75.0,
		},
		{
			csp:      "alibaba",
			caseName: "Alibaba ecs.gn6i-c4g1.xlarge (Tesla T4 16GB - Tier 1 exact match)",
			spec: cloudmodel.SpecInfo{
				Id:                  "alibaba+ap-northeast-2+ecs.gn6i-c4g1.xlarge",
				CspSpecName:         "ecs.gn6i-c4g1.xlarge",
				VCPU:                4,
				MemoryGiB:           15,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 16,
				AcceleratorModel:    "NVIDIA Tesla T4",
			},
			expectedCPU:     100.0,
			expectedMemory:  (15.0 / 16.0) * 100.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 100.0,
			expectedGPUMod:  100.0,
		},
		{
			csp:      "tencent",
			caseName: "Tencent GN7.2XLARGE32 (Tesla T4 16GB - Tier 1 exact match)",
			spec: cloudmodel.SpecInfo{
				Id:                  "tencent+ap-seoul+GN7.2XLARGE32",
				CspSpecName:         "GN7.2XLARGE32",
				VCPU:                8,
				MemoryGiB:           32,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 16,
				AcceleratorModel:    "NVIDIA Tesla T4",
			},
			expectedCPU:     50.0,
			expectedMemory:  50.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 100.0,
			expectedGPUMod:  100.0,
		},
		{
			csp:      "ibm",
			caseName: "IBM gx2-8x64x1v100 (V100 16GB - Tier 2 same vendor, identical VRAM)",
			spec: cloudmodel.SpecInfo{
				Id:                  "ibm+au-syd+gx2-8x64x1v100",
				CspSpecName:         "gx2-8x64x1v100",
				VCPU:                8,
				MemoryGiB:           64,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 16,
				AcceleratorModel:    "NVIDIA Tesla V100",
			},
			expectedCPU:     50.0,
			expectedMemory:  25.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 100.0,
			expectedGPUMod:  75.0,
		},
		{
			csp:      "ncp",
			caseName: "NCP g2 (V100 32GB - Tier 2 same vendor, VRAM upsizing)",
			spec: cloudmodel.SpecInfo{
				Id:                  "ncp+kr+g2",
				CspSpecName:         "g2",
				VCPU:                8,
				MemoryGiB:           32,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 32,
				AcceleratorModel:    "NVIDIA Tesla V100",
			},
			expectedCPU:     50.0,
			expectedMemory:  50.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 100.0,
			expectedGPUMod:  75.0,
		},
		{
			csp:      "nhn",
			caseName: "NHN m2.c8m32.gpu (V100 32GB - Tier 2 same vendor, VRAM upsizing)",
			spec: cloudmodel.SpecInfo{
				Id:                  "nhn+kr1+m2.c8m32.gpu",
				CspSpecName:         "m2.c8m32.gpu",
				VCPU:                8,
				MemoryGiB:           32,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 32,
				AcceleratorModel:    "NVIDIA Tesla V100",
			},
			expectedCPU:     50.0,
			expectedMemory:  50.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 100.0,
			expectedGPUMod:  75.0,
		},
		{
			csp:      "aws",
			caseName: "AWS g4ad.2xlarge (AMD Radeon Pro V520 - Tier 3 cross-vendor, VRAM downsizing)",
			spec: cloudmodel.SpecInfo{
				Id:                  "aws+ap-northeast-2+g4ad.2xlarge",
				CspSpecName:         "g4ad.2xlarge",
				VCPU:                8,
				MemoryGiB:           32,
				AcceleratorCount:    1,
				AcceleratorMemoryGB: 8,
				AcceleratorModel:    "AMD Radeon Pro V520",
			},
			expectedCPU:     50.0,
			expectedMemory:  50.0,
			expectedImage:   100.0,
			expectedGPUCnt:  100.0,
			expectedGPUVram: 50.0, // 8 / 16 * 100
			expectedGPUMod:  50.0, // Tier 3
		},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			img := standardGpuImage
			if tc.customImage != nil {
				img = *tc.customImage
			}
			vec := calculateMatchRateVector(tc.csp, nodeT4, tc.spec, img)

			assert.True(t, vec.HasGPU)
			require.Len(t, vec.ActiveDimensions(), 6)
			assert.InDelta(t, tc.expectedCPU, vec.CPU, 0.01)
			assert.InDelta(t, tc.expectedMemory, vec.Memory, 0.01)
			assert.InDelta(t, tc.expectedImage, vec.Image, 0.01)
			assert.InDelta(t, tc.expectedGPUCnt, vec.GPUCount, 0.01)
			assert.InDelta(t, tc.expectedGPUVram, vec.GPUVRAM, 0.01)
			assert.InDelta(t, tc.expectedGPUMod, vec.GPUModel, 0.01)
		})
	}
}

// TestMatchRate_TeslaT4_SummaryFormatting verifies 6-dimension format string.
func TestMatchRate_TeslaT4_SummaryFormatting(t *testing.T) {
	nodeT4 := onpremmodel.NodeProperty{
		MachineId: "t4-node",
		CPU:       onpremmodel.CpuProperty{Cpus: 4, Threads: 1, Architecture: "amd64"},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 16},
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "Tesla T4", MemoryTotalGB: 16},
		},
	}

	specT4 := cloudmodel.SpecInfo{
		VCPU:                4,
		MemoryGiB:           16,
		AcceleratorCount:    1,
		AcceleratorMemoryGB: 16,
		AcceleratorModel:    "NVIDIA Tesla T4",
	}

	imageGPU := cloudmodel.ImageInfo{
		OSType:         "ubuntu 22.04",
		OSArchitecture: "amd64",
		OSDistribution: "ubuntu jammy 22.04",
		IsGPUImage:     true,
	}

	vec := calculateMatchRateVector("aws", nodeT4, specT4, imageGPU)
	expectedStr := "Match Rate: CPU=100.0% Memory=100.0% Image=100.0% GPUCount=100.0% GPUVRAM=100.0% GPUModel=100.0%"
	assert.Equal(t, expectedStr, vec.FormatMatchRateSummary())
}
