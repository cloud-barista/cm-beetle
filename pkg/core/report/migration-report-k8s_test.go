/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package report

import (
	"strings"
	"testing"
	"time"

	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
)

func TestPartitionSourceNodes(t *testing.T) {
	nodes := []onpremmodel.NodeProperty{
		{
			Hostname:  "k8s-master-01",
			MachineId: "m1",
			Role:      "control-plane",
		},
		{
			Hostname:  "k8s-worker-01",
			MachineId: "w1",
			Role:      "worker",
		},
		{
			Hostname:  "k8s-worker-02",
			MachineId: "w2",
			Role:      "Worker",
		},
		{
			Hostname:  "master-node-02",
			MachineId: "m2",
			Role:      "", // empty role, but hostname indicates master
		},
	}

	workers, excluded := partitionSourceNodes(nodes)

	if len(workers) != 2 {
		t.Fatalf("expected 2 workers, got %d", len(workers))
	}
	if len(excluded) != 2 {
		t.Fatalf("expected 2 excluded control-plane nodes, got %d", len(excluded))
	}

	if excluded[0].Hostname != "k8s-master-01" || excluded[0].Reason == "" {
		t.Errorf("unexpected excluded item: %+v", excluded[0])
	}
}

func TestCorrelateWorkersToNodeGroups(t *testing.T) {
	workers := []onpremmodel.NodeProperty{
		{
			Hostname:  "worker-small",
			MachineId: "w-small",
			Role:      "worker",
			CPU:       onpremmodel.CpuProperty{Cpus: 1, Cores: 1, Threads: 1},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 2},
			RootDisk:  onpremmodel.DiskProperty{TotalSize: 50, Type: "SSD"},
		},
		{
			Hostname:  "worker-standard",
			MachineId: "w-std",
			Role:      "worker",
			CPU:       onpremmodel.CpuProperty{Cpus: 1, Cores: 4, Threads: 4},
			Memory:    onpremmodel.MemoryProperty{TotalSize: 16},
			RootDisk:  onpremmodel.DiskProperty{TotalSize: 100, Type: "SSD"},
		},
	}

	nodeGroups := []summary.SummaryNodeGroupInfo{
		{
			Name:            "workers-tier1",
			SpecId:          "c5a.xlarge",
			Vcpu:            4,
			MemoryGib:       8.0,
			DesiredNodeSize: 1,
			MonthlyCostUSD:  110.88,
		},
		{
			Name:            "workers-tier2",
			SpecId:          "c5a.2xlarge",
			Vcpu:            8,
			MemoryGib:       16.0,
			DesiredNodeSize: 1,
			MonthlyCostUSD:  221.76,
		},
	}

	mappings, sizingItems := correlateWorkersToNodeGroups(workers, nodeGroups)

	if len(mappings) != 2 {
		t.Fatalf("expected 2 node group mappings, got %d", len(mappings))
	}
	if len(sizingItems) != 2 {
		t.Fatalf("expected 2 sizing items, got %d", len(sizingItems))
	}

	// worker-small (1 vCPU, 2 GiB) should have floor upscaled rationale
	if !strings.Contains(sizingItems[0].SizingReason, "floor") {
		t.Errorf("expected floor upscaled sizing reason, got: %s", sizingItems[0].SizingReason)
	}

	// worker-standard (4 vCPU, 16 GiB) should be assigned to workers-tier2
	if sizingItems[1].TargetSpecId != "c5a.2xlarge" {
		t.Errorf("expected worker-standard to map to c5a.2xlarge, got: %s", sizingItems[1].TargetSpecId)
	}
}

func TestBuildK8sResourceComparison(t *testing.T) {
	workers := []onpremmodel.NodeProperty{
		{
			Hostname: "w1",
			Role:     "worker",
			CPU:      onpremmodel.CpuProperty{Cpus: 1, Cores: 2, Threads: 2},
			Memory:   onpremmodel.MemoryProperty{TotalSize: 4},
			RootDisk: onpremmodel.DiskProperty{TotalSize: 50},
		},
	}

	targetSummary := &summary.TargetK8sInfraSummary{
		Overview: summary.TargetK8sInfraOverview{
			TotalWorkerNodeCount: 2,
		},
		ComputeOverview: summary.K8sComputeOverview{
			TotalVcpus:     8,
			TotalMemoryGib: 16.0,
			TotalStorageGB: 100,
		},
	}

	comp := buildK8sResourceComparison(workers, targetSummary)

	if comp.WorkerCountDelta.ChangeType != "Upgrade" {
		t.Errorf("expected Upgrade for worker count, got %s", comp.WorkerCountDelta.ChangeType)
	}
	if comp.VcpuDelta.ChangeRatio != 4.0 {
		t.Errorf("expected 4.0x vcpu change ratio, got %f", comp.VcpuDelta.ChangeRatio)
	}
	if comp.MemoryGibDelta.ChangeRatio != 4.0 {
		t.Errorf("expected 4.0x memory change ratio, got %f", comp.MemoryGibDelta.ChangeRatio)
	}
}

func TestGenerateK8sMigrationReportMarkdown(t *testing.T) {
	report := &K8sMigrationReport{
		Metadata: ReportMetadata{
			GeneratedAt:   time.Now(),
			MigrationID:   "mig01/k8s-cluster",
			Namespace:     "mig01",
			InfraID:       "k8s-cluster",
			ReportVersion: "1.0",
		},
		ExecutiveSummary: K8sExecutiveSummary{
			MigrationStatus:      "Completed",
			SourceClusterName:    "onprem-cluster",
			SourceTotalNodes:     3,
			SourceWorkerNodes:    2,
			TargetCloud:          "AWS",
			TargetRegion:         "ap-northeast-2",
			TargetClusterName:    "target-k8s",
			TargetNodeGroupCount: 1,
			TargetTotalNodes:     2,
			MonthlyCostUSD:       221.76,
		},
		VersionAnalysis: K8sVersionAnalysis{
			SourceVersion: "1.28.0",
			TargetVersion: "1.34",
			UpgradeStatus: "Upgraded (Compatible)",
			Notes:         "Upgraded to target CSP supported minimum stable version",
		},
		ResourceComparison: K8sResourceComparison{
			WorkerCountDelta: ResourceChange{
				ResourceType: "Worker Count", SourceValue: "2.0 nodes", TargetValue: "2.0 nodes", ChangeType: "Same", ChangeRatio: 1.0, Description: "+0 nodes (1.0x same)",
			},
			VcpuDelta: ResourceChange{
				ResourceType: "Total vCPU", SourceValue: "4.0 vCPU", TargetValue: "8.0 vCPU", ChangeType: "Upgrade", ChangeRatio: 2.0, Description: "+4 vCPU (2.0x upgrade)",
			},
			MemoryGibDelta: ResourceChange{
				ResourceType: "Total Memory", SourceValue: "8.0 GiB", TargetValue: "16.0 GiB", ChangeType: "Upgrade", ChangeRatio: 2.0, Description: "+8.0 GiB (2.0x upgrade)",
			},
			StorageGbDelta: ResourceChange{
				ResourceType: "Total Storage", SourceValue: "100.0 GB", TargetValue: "100.0 GB", ChangeType: "Same", ChangeRatio: 1.0, Description: "+0 GB (1.0x same)",
			},
		},
		NodeGroupMappings: []WorkerNodeGroupMapping{
			{
				NodeGroupName:   "workers1",
				TargetSpecId:    "c5a.xlarge",
				TargetNodeCount: 2,
				AssignedWorkers: []SourceServerBrief{
					{
						Hostname: "w1", MachineID: "mid-1", CPUModel: "Intel", CPUs: 1, CPUThreads: 2, MemoryGB: 4, DiskGB: 50, DiskType: "SSD", OSName: "Ubuntu 22.04", PrimaryIP: "10.0.0.1",
					},
				},
			},
		},
		SizingAnalysis: []WorkerSizingItem{
			{
				WorkerHostname: "w1", SourceVcpu: 2, SourceMemoryGb: 4.0, TargetSpecId: "c5a.xlarge", TargetVcpu: 4, TargetMemoryGb: 8.0, SizingReason: "Rounded to closest catalog spec with headroom (c5a.xlarge)",
			},
		},
		ExcludedWorkers: []ExcludedWorkerItem{
			{
				Hostname: "m1", Role: "control-plane", Reason: "Control-plane abstracted into CSP managed control plane",
			},
		},
		CostSummary: CostSummary{
			TotalHourlyCost:  0.308,
			TotalDailyCost:   7.392,
			TotalMonthlyCost: 221.76,
			TotalYearlyCost:  2661.12,
			CostByComponent: []ComponentCost{
				{ComponentName: "NodeGroup: workers1 (2 nodes)", SpecName: "c5a.xlarge", MonthlyCost: 221.76, CostPercentage: 100.0},
			},
		},
		Recommendations: []Recommendation{
			{
				Category: "Scalability", Priority: "Medium", Title: "Enable Auto-Scaling", Description: "Configure Cluster Autoscaler", ActionItems: []string{"Enable on workers1"},
			},
		},
	}

	md := GenerateK8sMigrationReportMarkdown(report)

	expectedSections := []string{
		"# 🚀 Kubernetes Infrastructure Migration Report",
		"## 📊 Migration Summary",
		"## ☸️ Kubernetes Version Analysis",
		"## 🔄 Resource Delta & Sizing Comparison",
		"## 📦 Worker Node Group Consolidation Mappings",
		"## 🎯 Worker Sizing Rationale (4-Category Analysis)",
		"## 🛡️ Abstracted / Excluded Control-Plane Nodes",
		"## 💰 Cost Summary",
		"## 💡 Architectural Recommendations & Next Steps",
	}

	for _, section := range expectedSections {
		if !strings.Contains(md, section) {
			t.Errorf("markdown report missing expected section: %s", section)
		}
	}
}

