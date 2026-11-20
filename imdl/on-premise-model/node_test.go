package onpremisemodel

import (
	"encoding/json"
	"testing"
)

func TestNodePropertyWithGPU(t *testing.T) {
	nodeJSON := `{
		"hostname": "gpu-node-01",
		"machineId": "m-12345678",
		"cpu": {
			"architecture": "x86_64",
			"cpus": 2,
			"cores": 16,
			"threads": 32,
			"maxSpeed": 3.2,
			"vendor": "GenuineIntel",
			"model": "Intel(R) Xeon(R) Gold 6226R"
		},
		"memory": {
			"type": "DDR4",
			"totalSize": 128,
			"available": 110,
			"used": 18
		},
		"rootDisk": {
			"label": "/",
			"type": "SSD",
			"totalSize": 1024,
			"available": 800,
			"used": 224
		},
		"dataDisks": [
			{
				"label": "/data",
				"type": "SSD",
				"totalSize": 2048,
				"available": 1800,
				"used": 248
			}
		],
		"interfaces": [],
		"routingTable": [],
		"os": {
			"prettyName": "Ubuntu 22.04.4 LTS"
		},
		"gpuCards": [
			{
				"driverIndex": "0",
				"uuid": "GPU-11111111-2222-3333-4444-555555555555",
				"vendor": "NVIDIA",
				"model": "NVIDIA A100-PCIE-40GB",
				"type": "GPU",
				"architecture": "Ampere",
				"driverVersion": "535.129.03",
				"cudaVersion": "12.2",
				"slot": "PCIe Slot 1",
				"pciBusId": "0000:01:00.0",
				"memoryTotalGB": 40,
				"memoryFreeGB": 38,
				"memoryUsedGB": 2
			},
			{
				"driverIndex": "1",
				"uuid": "GPU-66666666-7777-8888-9999-000000000000",
				"vendor": "NVIDIA",
				"model": "NVIDIA A100-PCIE-40GB",
				"type": "GPU",
				"architecture": "Ampere",
				"driverVersion": "535.129.03",
				"cudaVersion": "12.2",
				"slot": "PCIe Slot 2",
				"pciBusId": "0000:02:00.0",
				"memoryTotalGB": 40,
				"memoryFreeGB": 39,
				"memoryUsedGB": 1
			}
		]
	}`

	var node NodeProperty
	err := json.Unmarshal([]byte(nodeJSON), &node)
	if err != nil {
		t.Fatalf("failed to unmarshal NodeProperty with GPUCards: %v", err)
	}

	if len(node.GPUCards) != 2 {
		t.Fatalf("expected 2 GPU cards, got %d", len(node.GPUCards))
	}

	card0 := node.GPUCards[0]
	if card0.Vendor != "NVIDIA" {
		t.Errorf("expected GPU vendor 'NVIDIA', got '%s'", card0.Vendor)
	}

	if card0.Model != "NVIDIA A100-PCIE-40GB" {
		t.Errorf("expected Model 'NVIDIA A100-PCIE-40GB', got '%s'", card0.Model)
	}

	if card0.CudaVersion != "12.2" {
		t.Errorf("expected CUDA version '12.2', got '%s'", card0.CudaVersion)
	}

	if card0.MemoryTotalGB != 40 {
		t.Errorf("expected MemoryTotalGB 40, got %f", card0.MemoryTotalGB)
	}

	if card0.Slot != "PCIe Slot 1" {
		t.Errorf("expected Slot 'PCIe Slot 1', got '%s'", card0.Slot)
	}

	if card0.DriverIndex != "0" {
		t.Errorf("expected DriverIndex '0', got '%s'", card0.DriverIndex)
	}
}

func TestNodePropertyWithoutGPUBackwardCompatibility(t *testing.T) {
	nodeJSON := `{
		"hostname": "cpu-node-01",
		"machineId": "m-87654321",
		"cpu": {
			"architecture": "x86_64",
			"cpus": 1,
			"cores": 8,
			"threads": 16
		},
		"memory": {
			"type": "DDR4",
			"totalSize": 64
		},
		"rootDisk": {
			"label": "/",
			"type": "SSD",
			"totalSize": 512
		},
		"interfaces": [],
		"routingTable": [],
		"os": {
			"prettyName": "Ubuntu 22.04 LTS"
		}
	}`

	var node NodeProperty
	err := json.Unmarshal([]byte(nodeJSON), &node)
	if err != nil {
		t.Fatalf("failed to unmarshal NodeProperty without GPU: %v", err)
	}

	if len(node.GPUCards) != 0 {
		t.Errorf("expected node.GPUCards to be empty/nil for CPU node, got %+v", node.GPUCards)
	}

	// Marshalling should omit "gpuCards" field
	data, err := json.Marshal(node)
	if err != nil {
		t.Fatalf("failed to marshal NodeProperty: %v", err)
	}

	var rawMap map[string]interface{}
	if err := json.Unmarshal(data, &rawMap); err != nil {
		t.Fatalf("failed to unmarshal into map: %v", err)
	}

	if _, exists := rawMap["gpuCards"]; exists {
		t.Errorf("expected 'gpuCards' key to be omitted when nil/empty")
	}
}

func TestNodePropertyWithMixedGPU(t *testing.T) {
	// Node with 4 physical GPU cards: 2x NVIDIA A100-PCIE-40GB, 1x NVIDIA Tesla T4 16GB, and 1x AMD Instinct MI350 288GB
	mixedGpuJSON := `{
		"hostname": "mixed-gpu-server",
		"machineId": "m-mixed-001",
		"cpu": {
			"architecture": "x86_64",
			"cpus": 2,
			"cores": 32,
			"threads": 64
		},
		"memory": {
			"type": "DDR5",
			"totalSize": 512
		},
		"rootDisk": {
			"label": "/",
			"type": "NVMe",
			"totalSize": 2048
		},
		"interfaces": [],
		"routingTable": [],
		"os": {
			"prettyName": "Ubuntu 24.04 LTS"
		},
		"gpuCards": [
			{
				"driverIndex": "0",
				"uuid": "GPU-a100-01",
				"vendor": "NVIDIA",
				"model": "NVIDIA A100-PCIE-40GB",
				"type": "GPU",
				"architecture": "Ampere",
				"driverVersion": "550.54.14",
				"cudaVersion": "12.4",
				"slot": "PCIe Slot 1",
				"pciBusId": "0000:01:00.0",
				"memoryTotalGB": 40,
				"memoryFreeGB": 39,
				"memoryUsedGB": 1
			},
			{
				"driverIndex": "1",
				"uuid": "GPU-a100-02",
				"vendor": "NVIDIA",
				"model": "NVIDIA A100-PCIE-40GB",
				"type": "GPU",
				"architecture": "Ampere",
				"driverVersion": "550.54.14",
				"cudaVersion": "12.4",
				"slot": "PCIe Slot 2",
				"pciBusId": "0000:02:00.0",
				"memoryTotalGB": 40,
				"memoryFreeGB": 39,
				"memoryUsedGB": 1
			},
			{
				"driverIndex": "2",
				"uuid": "GPU-t4-01",
				"vendor": "NVIDIA",
				"model": "Tesla T4",
				"type": "GPU",
				"architecture": "Turing",
				"driverVersion": "550.54.14",
				"cudaVersion": "12.4",
				"slot": "PCIe Slot 3",
				"pciBusId": "0000:03:00.0",
				"memoryTotalGB": 16,
				"memoryFreeGB": 15,
				"memoryUsedGB": 1
			},
			{
				"driverIndex": "card0",
				"uuid": "GPU-mi350-01",
				"vendor": "AMD",
				"model": "AMD Instinct MI350",
				"type": "GPU",
				"architecture": "CDNA 4",
				"driverVersion": "ROCm 6.2",
				"slot": "PCIe Slot 4",
				"pciBusId": "0000:04:00.0",
				"memoryTotalGB": 288,
				"memoryFreeGB": 280,
				"memoryUsedGB": 8
			}
		]
	}`

	var node NodeProperty
	err := json.Unmarshal([]byte(mixedGpuJSON), &node)
	if err != nil {
		t.Fatalf("failed to unmarshal mixed GPU JSON: %v", err)
	}

	if len(node.GPUCards) != 4 {
		t.Fatalf("expected 4 GPU cards, got %d", len(node.GPUCards))
	}

	// Verify card 0 (A100)
	c0 := node.GPUCards[0]
	if c0.Vendor != "NVIDIA" || c0.Model != "NVIDIA A100-PCIE-40GB" || c0.MemoryTotalGB != 40 || c0.Architecture != "Ampere" {
		t.Errorf("unexpected card 0: %+v", c0)
	}

	// Verify card 2 (T4)
	c2 := node.GPUCards[2]
	if c2.Vendor != "NVIDIA" || c2.Model != "Tesla T4" || c2.MemoryTotalGB != 16 || c2.Architecture != "Turing" {
		t.Errorf("unexpected card 2: %+v", c2)
	}

	// Verify card 3 (MI350)
	c3 := node.GPUCards[3]
	if c3.Vendor != "AMD" || c3.Model != "AMD Instinct MI350" || c3.MemoryTotalGB != 288 || c3.Architecture != "CDNA 4" {
		t.Errorf("unexpected card 3: %+v", c3)
	}
	if c3.DriverIndex != "card0" {
		t.Errorf("expected card 3 DriverIndex 'card0', got '%s'", c3.DriverIndex)
	}
}

func TestNodePropertyWithListeningPorts(t *testing.T) {
	nodeJSON := `{
		"hostname": "web-node-01",
		"machineId": "m-web-001",
		"cpu": {
			"architecture": "x86_64",
			"cpus": 1,
			"cores": 4,
			"threads": 8
		},
		"memory": {
			"type": "DDR4",
			"totalSize": 16
		},
		"rootDisk": {
			"label": "/",
			"type": "SSD",
			"totalSize": 256
		},
		"interfaces": [],
		"routingTable": [],
		"listeningPorts": [
			{
				"protocol": "tcp",
				"bindAddress": "0.0.0.0",
				"port": 8080,
				"processName": "nginx"
			},
			{
				"protocol": "udp",
				"bindAddress": "0.0.0.0",
				"port": 53,
				"processName": "named"
			},
			{
				"protocol": "tcp",
				"port": 22
			}
		],
		"os": {
			"prettyName": "Ubuntu 22.04 LTS"
		}
	}`

	var node NodeProperty
	if err := json.Unmarshal([]byte(nodeJSON), &node); err != nil {
		t.Fatalf("failed to unmarshal NodeProperty with ListeningPorts: %v", err)
	}

	if len(node.ListeningPorts) != 3 {
		t.Fatalf("expected 3 listening ports, got %d", len(node.ListeningPorts))
	}

	p0 := node.ListeningPorts[0]
	if p0.Protocol != "tcp" || p0.BindAddress != "0.0.0.0" || p0.Port != 8080 || p0.ProcessName != "nginx" {
		t.Errorf("unexpected listening port 0: %+v", p0)
	}

	p1 := node.ListeningPorts[1]
	if p1.Protocol != "udp" || p1.BindAddress != "0.0.0.0" || p1.Port != 53 || p1.ProcessName != "named" {
		t.Errorf("unexpected listening port 1: %+v", p1)
	}

	p2 := node.ListeningPorts[2]
	if p2.Protocol != "tcp" || p2.BindAddress != "" || p2.Port != 22 || p2.ProcessName != "" {
		t.Errorf("unexpected listening port 2: %+v", p2)
	}

	data, err := json.Marshal(node)
	if err != nil {
		t.Fatalf("failed to marshal NodeProperty: %v", err)
	}

	var roundTrip NodeProperty
	if err := json.Unmarshal(data, &roundTrip); err != nil {
		t.Fatalf("failed to round-trip NodeProperty: %v", err)
	}

	if len(roundTrip.ListeningPorts) != 3 || roundTrip.ListeningPorts[0] != p0 || roundTrip.ListeningPorts[2] != p2 {
		t.Errorf("round-trip mismatch: got %+v", roundTrip.ListeningPorts)
	}
}

func TestNodePropertyWithoutListeningPortsBackwardCompatibility(t *testing.T) {
	nodeJSON := `{
		"hostname": "legacy-node-01",
		"machineId": "m-legacy-001",
		"cpu": {
			"architecture": "x86_64",
			"cpus": 1,
			"cores": 4,
			"threads": 8
		},
		"memory": {
			"type": "DDR4",
			"totalSize": 16
		},
		"rootDisk": {
			"label": "/",
			"type": "SSD",
			"totalSize": 256
		},
		"interfaces": [],
		"routingTable": [],
		"os": {
			"prettyName": "Ubuntu 22.04 LTS"
		}
	}`

	var node NodeProperty
	if err := json.Unmarshal([]byte(nodeJSON), &node); err != nil {
		t.Fatalf("failed to unmarshal NodeProperty without ListeningPorts: %v", err)
	}

	if len(node.ListeningPorts) != 0 {
		t.Errorf("expected empty ListeningPorts, got %+v", node.ListeningPorts)
	}

	data, err := json.Marshal(node)
	if err != nil {
		t.Fatalf("failed to marshal NodeProperty: %v", err)
	}

	var rawMap map[string]interface{}
	if err := json.Unmarshal(data, &rawMap); err != nil {
		t.Fatalf("failed to unmarshal into map: %v", err)
	}

	if _, exists := rawMap["listeningPorts"]; exists {
		t.Errorf("expected 'listeningPorts' key to be omitted when empty")
	}
}
