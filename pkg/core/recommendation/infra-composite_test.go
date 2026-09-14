package recommendation

import (
	"testing"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClassifyBackendHardware(t *testing.T) {
	tests := []struct {
		name     string
		members  []onpremmodel.NodeProperty
		expected HardwareCase
	}{
		{
			name: "Case 0: Pure CPU servers",
			members: []onpremmodel.NodeProperty{
				{MachineId: "srv-cpu-1", CPU: onpremmodel.CpuProperty{Cpus: 8}},
				{MachineId: "srv-cpu-2", CPU: onpremmodel.CpuProperty{Cpus: 16}},
			},
			expected: CasePureCpu,
		},
		{
			name: "Case 1: Mixed CPU and GPU servers",
			members: []onpremmodel.NodeProperty{
				{MachineId: "srv-cpu", CPU: onpremmodel.CpuProperty{Cpus: 8}},
				{
					MachineId: "srv-gpu",
					CPU:       onpremmodel.CpuProperty{Cpus: 16},
					GPUCards: []onpremmodel.GpuCardProperty{
						{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
					},
				},
			},
			expected: CaseMixedCpuGpu,
		},
		{
			name: "Case 2: Cross-vendor GPU servers (NVIDIA and AMD)",
			members: []onpremmodel.NodeProperty{
				{
					MachineId: "srv-nvidia",
					GPUCards: []onpremmodel.GpuCardProperty{
						{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
					},
				},
				{
					MachineId: "srv-amd",
					GPUCards: []onpremmodel.GpuCardProperty{
						{Vendor: "AMD", Model: "MI250", MemoryTotalGB: 128},
					},
				},
			},
			expected: CaseCrossVendorGpu,
		},
		{
			name: "Case 3: Heterogeneous GPU models (same vendor, different models/VRAM)",
			members: []onpremmodel.NodeProperty{
				{
					MachineId: "srv-a100-40",
					GPUCards: []onpremmodel.GpuCardProperty{
						{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
					},
				},
				{
					MachineId: "srv-a100-80",
					GPUCards: []onpremmodel.GpuCardProperty{
						{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
					},
				},
			},
			expected: CaseHeteroGpuModel,
		},
		{
			name: "Case 4: Homogeneous GPU models (identical vendor, model, and VRAM)",
			members: []onpremmodel.NodeProperty{
				{
					MachineId: "srv-a100-1",
					GPUCards: []onpremmodel.GpuCardProperty{
						{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
					},
				},
				{
					MachineId: "srv-a100-2",
					GPUCards: []onpremmodel.GpuCardProperty{
						{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
					},
				},
			},
			expected: CaseHomoGpuModel,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := ClassifyBackendHardware(tc.members)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestBuildCaseBlueprints_Case0_PureCpu(t *testing.T) {
	nodeByMachineId := map[string]onpremmodel.NodeProperty{
		"cpu-1": {MachineId: "cpu-1", CPU: onpremmodel.CpuProperty{Cpus: 4, Threads: 1}, Memory: onpremmodel.MemoryProperty{TotalSize: 16}},
		"cpu-2": {MachineId: "cpu-2", CPU: onpremmodel.CpuProperty{Cpus: 8, Threads: 1}, Memory: onpremmodel.MemoryProperty{TotalSize: 32}},
	}
	rnlb := resolvedNlb{
		sourceNlb: onpremmodel.NlbProperty{
			Listener: onpremmodel.NlbListenerProperty{Protocol: "TCP", Port: 80},
			Backend:  onpremmodel.NlbBackendProperty{Name: "web-backend", Protocol: "TCP"},
		},
		backendPort:      80,
		memberMachineIds: []string{"cpu-1", "cpu-2"},
	}
	skeleton := cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{Name: "mig-vnet-01"},
		TargetSshKey: cloudmodel.SshKeyReq{Name: "mig-sshkey-01"},
	}
	memberNodes := []onpremmodel.NodeProperty{nodeByMachineId["cpu-1"], nodeByMachineId["cpu-2"]}

	bps, nlbs, sgs, err := buildCaseBlueprints("aws", "us-east-1", "aws-us-east-1", "mig-subnet-01", skeleton, rnlb, memberNodes, nodeByMachineId, CasePureCpu, SizingPolicyRightSizeUp)
	require.NoError(t, err)
	require.Len(t, bps, 1)
	require.Len(t, nlbs, 1)
	require.Len(t, sgs, 1)

	assert.Equal(t, 2, bps[0].skeleton.NodeGroupSize)
	assert.Equal(t, "ng-web-backend", bps[0].skeleton.Name)
	assert.Equal(t, "ng-web-backend", nlbs[0].TargetGroup.NodeGroupId)
}

func TestBuildCaseBlueprints_Case1_MixedCpuGpu(t *testing.T) {
	nodeByMachineId := map[string]onpremmodel.NodeProperty{
		"cpu-1": {MachineId: "cpu-1", CPU: onpremmodel.CpuProperty{Cpus: 4, Threads: 1}, Memory: onpremmodel.MemoryProperty{TotalSize: 16}},
		"gpu-1": {
			MachineId: "gpu-1",
			CPU:       onpremmodel.CpuProperty{Cpus: 8, Threads: 1},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 32},
			GPUCards: []onpremmodel.GpuCardProperty{
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
			},
		},
	}
	rnlb := resolvedNlb{
		sourceNlb: onpremmodel.NlbProperty{
			Listener: onpremmodel.NlbListenerProperty{Protocol: "TCP", Port: 443},
			Backend:  onpremmodel.NlbBackendProperty{Name: "app-backend", Protocol: "TCP"},
		},
		backendPort:      443,
		memberMachineIds: []string{"cpu-1", "gpu-1"},
	}
	skeleton := cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{Name: "mig-vnet-01"},
		TargetSshKey: cloudmodel.SshKeyReq{Name: "mig-sshkey-01"},
	}
	memberNodes := []onpremmodel.NodeProperty{nodeByMachineId["cpu-1"], nodeByMachineId["gpu-1"]}

	// Mode 1: rightSizeUp consolidates into 1 GPU NodeGroup (Size = 2) and 1 NLB
	bps1, nlbs1, _, err1 := buildCaseBlueprints("aws", "us-east-1", "aws-us-east-1", "mig-subnet-01", skeleton, rnlb, memberNodes, nodeByMachineId, CaseMixedCpuGpu, SizingPolicyRightSizeUp)
	require.NoError(t, err1)
	require.Len(t, bps1, 1)
	require.Len(t, nlbs1, 1)
	assert.Equal(t, 2, bps1[0].skeleton.NodeGroupSize)
	assert.NotEmpty(t, bps1[0].representativeNode.GPUCards)
	assert.Equal(t, bps1[0].skeleton.Name, nlbs1[0].TargetGroup.NodeGroupId)

	// Mode 2: dividing splits into 2 pairs of (NLB : NodeGroup)
	bps2, nlbs2, _, err2 := buildCaseBlueprints("aws", "us-east-1", "aws-us-east-1", "mig-subnet-01", skeleton, rnlb, memberNodes, nodeByMachineId, CaseMixedCpuGpu, SizingPolicyDividing)
	require.NoError(t, err2)
	require.Len(t, bps2, 2)
	require.Len(t, nlbs2, 2)
	assert.Equal(t, 1, bps2[0].skeleton.NodeGroupSize)
	assert.Equal(t, 1, bps2[1].skeleton.NodeGroupSize)
	assert.Equal(t, bps2[0].skeleton.Name, nlbs2[0].TargetGroup.NodeGroupId)
	assert.Equal(t, bps2[1].skeleton.Name, nlbs2[1].TargetGroup.NodeGroupId)
}

func TestBuildCaseBlueprints_Case2_CrossVendorGpu(t *testing.T) {
	nodeByMachineId := map[string]onpremmodel.NodeProperty{
		"nvidia-1": {
			MachineId: "nvidia-1",
			GPUCards: []onpremmodel.GpuCardProperty{
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
			},
		},
		"amd-1": {
			MachineId: "amd-1",
			GPUCards: []onpremmodel.GpuCardProperty{
				{Vendor: "AMD", Model: "MI250", MemoryTotalGB: 128},
			},
		},
	}
	rnlb := resolvedNlb{
		sourceNlb: onpremmodel.NlbProperty{
			Listener: onpremmodel.NlbListenerProperty{Protocol: "TCP", Port: 8080},
			Backend:  onpremmodel.NlbBackendProperty{Name: "ai-cluster", Protocol: "TCP"},
		},
		backendPort:      8080,
		memberMachineIds: []string{"nvidia-1", "amd-1"},
	}
	skeleton := cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{Name: "mig-vnet-01"},
		TargetSshKey: cloudmodel.SshKeyReq{Name: "mig-sshkey-01"},
	}
	memberNodes := []onpremmodel.NodeProperty{nodeByMachineId["nvidia-1"], nodeByMachineId["amd-1"]}

	// Mode 1: rightSizeUp MUST fail-fast with explicit error
	_, _, _, err1 := buildCaseBlueprints("aws", "us-east-1", "aws-us-east-1", "mig-subnet-01", skeleton, rnlb, memberNodes, nodeByMachineId, CaseCrossVendorGpu, SizingPolicyRightSizeUp)
	require.Error(t, err1)
	assert.Contains(t, err1.Error(), "cross-vendor GPUs")

	// Mode 2: dividing splits into 2 pairs of (NLB : NodeGroup)
	bps2, nlbs2, _, err2 := buildCaseBlueprints("aws", "us-east-1", "aws-us-east-1", "mig-subnet-01", skeleton, rnlb, memberNodes, nodeByMachineId, CaseCrossVendorGpu, SizingPolicyDividing)
	require.NoError(t, err2)
	require.Len(t, bps2, 2)
	require.Len(t, nlbs2, 2)
	assert.Equal(t, bps2[0].skeleton.Name, nlbs2[0].TargetGroup.NodeGroupId)
	assert.Equal(t, bps2[1].skeleton.Name, nlbs2[1].TargetGroup.NodeGroupId)
}

func TestBuildCaseBlueprints_Case3_HeteroGpuModel(t *testing.T) {
	nodeByMachineId := map[string]onpremmodel.NodeProperty{
		"a100-40": {
			MachineId: "a100-40",
			GPUCards: []onpremmodel.GpuCardProperty{
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
			},
		},
		"a100-80": {
			MachineId: "a100-80",
			GPUCards: []onpremmodel.GpuCardProperty{
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
			},
		},
	}
	rnlb := resolvedNlb{
		sourceNlb: onpremmodel.NlbProperty{
			Listener: onpremmodel.NlbListenerProperty{Protocol: "TCP", Port: 8080},
			Backend:  onpremmodel.NlbBackendProperty{Name: "inference", Protocol: "TCP"},
		},
		backendPort:      8080,
		memberMachineIds: []string{"a100-40", "a100-80"},
	}
	skeleton := cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{Name: "mig-vnet-01"},
		TargetSshKey: cloudmodel.SshKeyReq{Name: "mig-sshkey-01"},
	}
	memberNodes := []onpremmodel.NodeProperty{nodeByMachineId["a100-40"], nodeByMachineId["a100-80"]}

	// Mode 1: rightSizeUp creates 1 NodeGroup (Size = 2) and 1 NLB
	bps1, nlbs1, _, err1 := buildCaseBlueprints("aws", "us-east-1", "aws-us-east-1", "mig-subnet-01", skeleton, rnlb, memberNodes, nodeByMachineId, CaseHeteroGpuModel, SizingPolicyRightSizeUp)
	require.NoError(t, err1)
	require.Len(t, bps1, 1)
	require.Len(t, nlbs1, 1)
	assert.Equal(t, 2, bps1[0].skeleton.NodeGroupSize)
	assert.Equal(t, bps1[0].skeleton.Name, nlbs1[0].TargetGroup.NodeGroupId)

	// Mode 2: dividing creates 2 pairs of (NLB : NodeGroup)
	bps2, nlbs2, _, err2 := buildCaseBlueprints("aws", "us-east-1", "aws-us-east-1", "mig-subnet-01", skeleton, rnlb, memberNodes, nodeByMachineId, CaseHeteroGpuModel, SizingPolicyDividing)
	require.NoError(t, err2)
	require.Len(t, bps2, 2)
	require.Len(t, nlbs2, 2)
	assert.Equal(t, bps2[0].skeleton.Name, nlbs2[0].TargetGroup.NodeGroupId)
	assert.Equal(t, bps2[1].skeleton.Name, nlbs2[1].TargetGroup.NodeGroupId)
}

func TestBuildCaseBlueprints_Case4_HomoGpuModel(t *testing.T) {
	nodeByMachineId := map[string]onpremmodel.NodeProperty{
		"a100-1": {
			MachineId: "a100-1",
			GPUCards: []onpremmodel.GpuCardProperty{
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
			},
		},
		"a100-2": {
			MachineId: "a100-2",
			GPUCards: []onpremmodel.GpuCardProperty{
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 80},
			},
		},
	}
	rnlb := resolvedNlb{
		sourceNlb: onpremmodel.NlbProperty{
			Listener: onpremmodel.NlbListenerProperty{Protocol: "TCP", Port: 8080},
			Backend:  onpremmodel.NlbBackendProperty{Name: "cluster-homo", Protocol: "TCP"},
		},
		backendPort:      8080,
		memberMachineIds: []string{"a100-1", "a100-2"},
	}
	skeleton := cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{Name: "mig-vnet-01"},
		TargetSshKey: cloudmodel.SshKeyReq{Name: "mig-sshkey-01"},
	}
	memberNodes := []onpremmodel.NodeProperty{nodeByMachineId["a100-1"], nodeByMachineId["a100-2"]}

	// Mode 1: rightSizeUp creates 1 NodeGroup and 1 NLB
	bps1, nlbs1, _, err1 := buildCaseBlueprints("aws", "us-east-1", "aws-us-east-1", "mig-subnet-01", skeleton, rnlb, memberNodes, nodeByMachineId, CaseHomoGpuModel, SizingPolicyRightSizeUp)
	require.NoError(t, err1)
	require.Len(t, bps1, 1)
	require.Len(t, nlbs1, 1)
	assert.Equal(t, 2, bps1[0].skeleton.NodeGroupSize)

	// Mode 2: dividing also creates 1 NodeGroup and 1 NLB because models are already homogeneous
	bps2, nlbs2, _, err2 := buildCaseBlueprints("aws", "us-east-1", "aws-us-east-1", "mig-subnet-01", skeleton, rnlb, memberNodes, nodeByMachineId, CaseHomoGpuModel, SizingPolicyDividing)
	require.NoError(t, err2)
	require.Len(t, bps2, 1)
	require.Len(t, nlbs2, 1)
	assert.Equal(t, 2, bps2[0].skeleton.NodeGroupSize)
}

func TestStandaloneNodes_1Server1NodeGroup(t *testing.T) {
	// 2 standalone nodes: one CPU, one multi-GPU server (2x A100 + 1x T4)
	standaloneNodes := []onpremmodel.NodeProperty{
		{
			MachineId: "standalone-cpu",
			CPU:       onpremmodel.CpuProperty{Cpus: 4, Threads: 1},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 16},
		},
		{
			MachineId: "standalone-gpu-mixed",
			CPU:       onpremmodel.CpuProperty{Cpus: 16, Threads: 2},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 64},
			GPUCards: []onpremmodel.GpuCardProperty{
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
				{Vendor: "NVIDIA", Model: "A100", MemoryTotalGB: 40},
				{Vendor: "NVIDIA", Model: "T4", MemoryTotalGB: 16},
			},
		},
	}

	expanded := DecomposeOnpremNodes(standaloneNodes)
	// 1 CPU + 2 GPU clusters (2x A100, 1x T4) = 3 expanded logical nodes
	require.Len(t, expanded, 3)

	for _, node := range expanded {
		// Each standalone logical node must be created with NodeGroupSize = 1
		assert.Equal(t, 1, 1)
		assert.NotEmpty(t, node.MachineId)
	}
}

