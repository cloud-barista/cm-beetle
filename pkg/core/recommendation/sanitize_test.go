package recommendation

import (
	"testing"

	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/stretchr/testify/assert"
)

func TestSanitizeSourceNode_MissingOrZeroValues(t *testing.T) {
	rawNode := onpremmodel.NodeProperty{
		Hostname:  "test-zero-node",
		MachineId: "machine-zero-01",
		CPU: onpremmodel.CpuProperty{
			Architecture: "x86_64",
			Cpus:         0,
			Cores:        0,
			Threads:      0,
		},
		Memory: onpremmodel.MemoryProperty{
			Type:      "DDR4",
			TotalSize: 0,
		},
		RootDisk: onpremmodel.DiskProperty{
			Label:     "/",
			Type:      "SSD",
			TotalSize: 0,
		},
	}

	sanitized := SanitizeSourceNode(rawNode)

	// Verify that zero values are safely mitigated to minimum defaults.
	assert.Equal(t, uint32(1), sanitized.CPU.Cpus)
	assert.Equal(t, uint32(1), sanitized.CPU.Cores)
	assert.Equal(t, uint32(1), sanitized.CPU.Threads)
	assert.Equal(t, uint64(1), sanitized.Memory.TotalSize)
	assert.Equal(t, uint64(1), sanitized.RootDisk.TotalSize)
	assert.Equal(t, "test-zero-node", sanitized.Hostname)
	assert.Equal(t, "machine-zero-01", sanitized.MachineId)
}

func TestSanitizeSourceNode_ValidValuesPreserved(t *testing.T) {
	validNode := onpremmodel.NodeProperty{
		Hostname:  "valid-worker-01",
		MachineId: "machine-valid-01",
		CPU: onpremmodel.CpuProperty{
			Architecture: "x86_64",
			Cpus:         2,
			Cores:        8,
			Threads:      16,
		},
		Memory: onpremmodel.MemoryProperty{
			Type:      "DDR4",
			TotalSize: 64,
		},
		RootDisk: onpremmodel.DiskProperty{
			Label:     "/",
			Type:      "SSD",
			TotalSize: 500,
		},
	}

	sanitized := SanitizeSourceNode(validNode)

	// Verify that valid node specifications are preserved without change.
	assert.Equal(t, uint32(2), sanitized.CPU.Cpus)
	assert.Equal(t, uint32(8), sanitized.CPU.Cores)
	assert.Equal(t, uint32(16), sanitized.CPU.Threads)
	assert.Equal(t, uint64(64), sanitized.Memory.TotalSize)
	assert.Equal(t, uint64(500), sanitized.RootDisk.TotalSize)
}

func TestSanitizeSourceInfra(t *testing.T) {
	srcInfra := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{
			{
				MachineId: "node-1",
				Memory:    onpremmodel.MemoryProperty{TotalSize: 0},
				CPU:       onpremmodel.CpuProperty{Cpus: 0},
			},
			{
				MachineId: "node-2",
				Memory:    onpremmodel.MemoryProperty{TotalSize: 32},
				CPU:       onpremmodel.CpuProperty{Cpus: 4, Cores: 4, Threads: 8},
				RootDisk:  onpremmodel.DiskProperty{TotalSize: 100},
			},
		},
	}

	sanitizedInfra := SanitizeSourceInfra(srcInfra)

	// Verify all nodes are properly processed in the infra.
	assert.Len(t, sanitizedInfra.Nodes, 2)
	assert.Equal(t, uint64(1), sanitizedInfra.Nodes[0].Memory.TotalSize)
	assert.Equal(t, uint32(1), sanitizedInfra.Nodes[0].CPU.Cpus)
	assert.Equal(t, uint64(32), sanitizedInfra.Nodes[1].Memory.TotalSize)
	assert.Equal(t, uint32(4), sanitizedInfra.Nodes[1].CPU.Cpus)
}

func TestSanitizeSourceInfra_EmptyNodes(t *testing.T) {
	emptyInfra := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{},
	}

	sanitized := SanitizeSourceInfra(emptyInfra)

	// Verify that empty nodes slice is handled gracefully.
	assert.Empty(t, sanitized.Nodes)
}

func TestSanitizeSourceNode_CrossRestore(t *testing.T) {
	// Threads are 0, Cores are 4: Threads should be restored to 4.
	nodeWithCoresOnly := onpremmodel.NodeProperty{
		CPU: onpremmodel.CpuProperty{Cores: 4, Threads: 0},
	}
	sanitized1 := SanitizeSourceNode(nodeWithCoresOnly)
	assert.Equal(t, uint32(4), sanitized1.CPU.Threads)
	assert.Equal(t, uint32(4), sanitized1.CPU.Cores)
	assert.Equal(t, uint32(1), sanitized1.CPU.Cpus)

	// Cores are 0, Threads are 8: Cores should be restored to 8.
	nodeWithThreadsOnly := onpremmodel.NodeProperty{
		CPU: onpremmodel.CpuProperty{Cores: 0, Threads: 8},
	}
	sanitized2 := SanitizeSourceNode(nodeWithThreadsOnly)
	assert.Equal(t, uint32(8), sanitized2.CPU.Cores)
	assert.Equal(t, uint32(8), sanitized2.CPU.Threads)
	assert.Equal(t, uint32(1), sanitized2.CPU.Cpus)
}

func TestBuildNodeGroupNameAndLineage(t *testing.T) {
	nodes := []onpremmodel.NodeProperty{
		{
			Hostname:  "ai-gpu-01",
			MachineId: "c1fb45d7-f1a0-465c-96f9-269c172d84df",
		},
		{
			Hostname:  "ubuntu",
			MachineId: "11111111-2222-3333-4444-aaaa55556666",
		},
		{
			Hostname:  "ubuntu",
			MachineId: "99999999-8888-7777-6666-bbbb55556666",
		},
		{
			Hostname:  "",
			MachineId: "00000000-0000-0000-0000-1234abcd5678",
		},
		{
			Hostname:  "web-01.corp.internal",
			MachineId: "abcdef12-3456-7890-abcd-ef1234567890",
		},
	}

	shortIds := buildShortMachineIdMap(nodes)
	usedNames := make(map[string]bool)

	// 1:1 standalone node with hostname and UUID
	name1, orig1 := buildNodeGroupNameAndLineage(nodes[0], shortIds, usedNames)
	assert.Equal(t, "ng-ai-gpu-01-172d84df", name1)
	assert.Equal(t, "c1fb45d7-f1a0-465c-96f9-269c172d84df", orig1)

	// Duplicate hostname "ubuntu" disambiguated by 8-char short ID
	name2, orig2 := buildNodeGroupNameAndLineage(nodes[1], shortIds, usedNames)
	assert.Equal(t, "ng-ubuntu-55556666", name2)
	assert.Equal(t, "11111111-2222-3333-4444-aaaa55556666", orig2)

	// Second node shares the last 8 chars ("55556666") so shortIds falls back to first 8 chars ("99999999")
	name3, orig3 := buildNodeGroupNameAndLineage(nodes[2], shortIds, usedNames)
	assert.Equal(t, "ng-ubuntu-99999999", name3)
	assert.Equal(t, "99999999-8888-7777-6666-bbbb55556666", orig3)

	// Empty hostname falls back to ng-{shortId8}
	name4, orig4 := buildNodeGroupNameAndLineage(nodes[3], shortIds, usedNames)
	assert.Equal(t, "ng-abcd5678", name4)
	assert.Equal(t, "00000000-0000-0000-0000-1234abcd5678", orig4)

	// FQDN strips domain suffix
	name5, orig5 := buildNodeGroupNameAndLineage(nodes[4], shortIds, usedNames)
	assert.Equal(t, "ng-web-01-34567890", name5)
	assert.Equal(t, "abcdef12-3456-7890-abcd-ef1234567890", orig5)

	// 1:N decomposed GPU nodes use ng-{hostname}-{gpuModel} and restore physical MachineId
	decomposedNode1 := onpremmodel.NodeProperty{
		Hostname:  "ai-gpu-01",
		MachineId: "c1fb45d7-f1a0-465c-96f9-269c172d84df-primary-a100",
	}
	decomposedNode2 := onpremmodel.NodeProperty{
		Hostname:  "ai-gpu-01",
		MachineId: "c1fb45d7-f1a0-465c-96f9-269c172d84df-gpu-t4",
	}

	gpuName1, gpuOrig1 := buildNodeGroupNameAndLineage(decomposedNode1, shortIds, usedNames)
	assert.Equal(t, "ng-ai-gpu-01-a100", gpuName1)
	assert.Equal(t, "c1fb45d7-f1a0-465c-96f9-269c172d84df", gpuOrig1)

	gpuName2, gpuOrig2 := buildNodeGroupNameAndLineage(decomposedNode2, shortIds, usedNames)
	assert.Equal(t, "ng-ai-gpu-01-t4", gpuName2)
	assert.Equal(t, "c1fb45d7-f1a0-465c-96f9-269c172d84df", gpuOrig2)
}


