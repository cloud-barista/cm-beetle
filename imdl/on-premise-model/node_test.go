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
			"model": "Intel(R) Xeon(R) Gold 6248R CPU @ 3.00GHz"
		},
		"memory": {
			"type": "DDR4",
			"totalSize": 256
		},
		"rootDisk": {
			"label": "/",
			"type": "SSD",
			"totalSize": 1024
		},
		"interfaces": [
			{"name": "eth0", "ipv4CidrBlocks": ["192.168.1.10/24"]}
		],
		"routingTable": [],
		"os": {
			"prettyName": "Ubuntu 22.04.4 LTS"
		},
		"gpu": {
			"count": 2,
			"vendor": "NVIDIA",
			"model": "NVIDIA A100-PCIE-40GB",
			"type": "GPU",
			"totalMemoryGB": 80,
			"driverVersion": "535.129.03",
			"cudaVersion": "12.2",
			"architecture": "Ampere",
			"details": [
				{
					"index": 0,
					"uuid": "GPU-11111111-2222-3333-4444-555555555555",
					"model": "NVIDIA A100-PCIE-40GB",
					"pciBusId": "0000:01:00.0",
					"memoryTotal": 40,
					"memoryFree": 38,
					"memoryUsed": 2
				},
				{
					"index": 1,
					"uuid": "GPU-66666666-7777-8888-9999-000000000000",
					"model": "NVIDIA A100-PCIE-40GB",
					"pciBusId": "0000:02:00.0",
					"memoryTotal": 40,
					"memoryFree": 39,
					"memoryUsed": 1
				}
			]
		}
	}`

	var node NodeProperty
	err := json.Unmarshal([]byte(nodeJSON), &node)
	if err != nil {
		t.Fatalf("failed to unmarshal NodeProperty with GPU: %v", err)
	}

	if node.GPU == nil {
		t.Fatalf("expected node.GPU to be non-nil")
	}

	if node.GPU.Count != 2 {
		t.Errorf("expected GPU count 2, got %d", node.GPU.Count)
	}

	if node.GPU.Vendor != "NVIDIA" {
		t.Errorf("expected GPU vendor 'NVIDIA', got '%s'", node.GPU.Vendor)
	}

	if node.GPU.CudaVersion != "12.2" {
		t.Errorf("expected CUDA version '12.2', got '%s'", node.GPU.CudaVersion)
	}

	if len(node.GPU.Details) != 2 {
		t.Errorf("expected 2 GPU details, got %d", len(node.GPU.Details))
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

	if node.GPU != nil {
		t.Errorf("expected node.GPU to be nil for CPU node, got %+v", node.GPU)
	}

	// Marshalling should omit "gpu" field
	data, err := json.Marshal(node)
	if err != nil {
		t.Fatalf("failed to marshal NodeProperty: %v", err)
	}

	var rawMap map[string]interface{}
	if err := json.Unmarshal(data, &rawMap); err != nil {
		t.Fatalf("failed to unmarshal into map: %v", err)
	}

	if _, exists := rawMap["gpu"]; exists {
		t.Errorf("expected 'gpu' key to be omitted when nil")
	}
}
