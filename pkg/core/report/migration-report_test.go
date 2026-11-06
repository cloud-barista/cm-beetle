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
	"encoding/json"
	"strings"
	"testing"

	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
)

// TestBuildMigrationMappingsNlbGroup reproduces the 3-server sample used across
// the test-cli reports: one solo node whose NodeGroup has its own machine ID
// label, plus two NLB backend nodes sharing a single NodeGroup (and therefore
// the same comma-separated "cm-source-machine-ids" label). All three must resolve
// to a mapping — previously the two NLB-group members were silently dropped,
// which is what caused the Cost Breakdown table to only show 1 of 3 VMs.
func TestBuildMigrationMappingsNlbGroup(t *testing.T) {
	sourceSummary := &summary.SourceInfraSummary{
		ComputeResources: summary.SourceSummaryComputeResources{
			Servers: []summary.SourceServerInfo{
				{Hostname: "ip-10-0-1-30", MachineId: "ec268ed7-821e-9d73-e79f-961262161624", CPU: summary.SourceCPUInfo{Cores: 1}, Memory: summary.SourceMemoryInfo{TotalGB: 2}, Disk: summary.SourceDiskInfo{TotalGB: 30}},
				{Hostname: "ip-10-0-1-221", MachineId: "ec2d32b5-98fb-5a96-7913-d3db1ec18932", CPU: summary.SourceCPUInfo{Cores: 1}, Memory: summary.SourceMemoryInfo{TotalGB: 8}, Disk: summary.SourceDiskInfo{TotalGB: 30}},
				{Hostname: "ip-10-0-1-138", MachineId: "ec288dd0-c6fa-8a49-2f60-bc898311febf", CPU: summary.SourceCPUInfo{Cores: 2}, Memory: summary.SourceMemoryInfo{TotalGB: 8}, Disk: summary.SourceDiskInfo{TotalGB: 30}},
			},
		},
	}

	nlbLabel := map[string]string{
		"cm-source-machine-ids": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
	}
	targetSummary := &summary.TargetInfraSummary{
		ComputeResources: summary.SummaryComputeResources{
			Vms: []summary.SummaryVmInfo{
				{
					Name:   "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
					Status: "Running",
					Spec:   summary.SummaryVmSpecInfo{Name: "t3a.small", VCpus: 2, MemoryGiB: 2},
					Label:  map[string]string{"cm-source-machine-ids": "ec268ed7-821e-9d73-e79f-961262161624"},
				},
				{
					Name:   "my-ng-influxdb-back-1",
					Status: "Running",
					Spec:   summary.SummaryVmSpecInfo{Name: "t3a.xlarge", VCpus: 4, MemoryGiB: 16},
					Label:  nlbLabel,
				},
				{
					Name:   "my-ng-influxdb-back-2",
					Status: "Running",
					Spec:   summary.SummaryVmSpecInfo{Name: "t3a.xlarge", VCpus: 4, MemoryGiB: 16},
					Label:  nlbLabel,
				},
			},
		},
		CostEstimation: summary.SummaryCostEstimation{
			ByVm: []summary.SummaryCostByVm{
				{VmName: "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1", CostPerMonth: 16.85},
				{VmName: "my-ng-influxdb-back-1", CostPerMonth: 134.78},
				{VmName: "my-ng-influxdb-back-2", CostPerMonth: 134.78},
			},
		},
	}

	mappings := buildMigrationMappings(sourceSummary, targetSummary)

	if len(mappings) != 3 {
		t.Fatalf("buildMigrationMappings() returned %d mappings, want 3 (got: %+v)", len(mappings), mappings)
	}

	byHostname := make(map[string]SourceTargetMapping)
	for _, m := range mappings {
		byHostname[m.SourceServer.Hostname] = m
	}

	cases := []struct {
		hostname   string
		wantVmName string
		wantCost   float64
	}{
		{"ip-10-0-1-30", "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1", 16.85},
		{"ip-10-0-1-221", "my-ng-influxdb-back-1", 134.78},
		{"ip-10-0-1-138", "my-ng-influxdb-back-2", 134.78},
	}

	for _, c := range cases {
		m, ok := byHostname[c.hostname]
		if !ok {
			t.Errorf("no mapping found for source hostname %q", c.hostname)
			continue
		}
		if m.TargetVM.InstanceName != c.wantVmName {
			t.Errorf("hostname %q mapped to VM %q, want %q", c.hostname, m.TargetVM.InstanceName, c.wantVmName)
		}
		// CostPerMonth passes through a float32 (SummaryCostByVm) on its way to
		// this float64 field, so compare at float32 precision.
		if float32(m.CostPerMonth) != float32(c.wantCost) {
			t.Errorf("hostname %q CostPerMonth = %v, want %v", c.hostname, m.CostPerMonth, c.wantCost)
		}
	}
}

func TestBuildMigrationMappingsDecomposedGpu(t *testing.T) {
	// 1:N decomposed GPU server: one physical server decomposed into two target VMs
	// sharing the same physical MachineId in cm-source-machine-ids.
	sourceSummary := &summary.SourceInfraSummary{
		ComputeResources: summary.SourceSummaryComputeResources{
			Servers: []summary.SourceServerInfo{
				{
					Hostname:  "ai-gpu-01",
					MachineId: "c1fb45d7-f1a0-465c-96f9-269c172d84df",
					CPU:       summary.SourceCPUInfo{Cores: 32},
					Memory:    summary.SourceMemoryInfo{TotalGB: 128},
					Disk:      summary.SourceDiskInfo{TotalGB: 500},
				},
			},
		},
	}

	targetSummary := &summary.TargetInfraSummary{
		ComputeResources: summary.SummaryComputeResources{
			Vms: []summary.SummaryVmInfo{
				{
					Name:   "ng-ai-gpu-01-a100-1",
					Status: "Running",
					Spec:   summary.SummaryVmSpecInfo{Name: "p4d.24xlarge", VCpus: 96, MemoryGiB: 1152},
					Label:  map[string]string{"cm-source-machine-ids": "c1fb45d7-f1a0-465c-96f9-269c172d84df"},
				},
				{
					Name:   "ng-ai-gpu-01-t4-1",
					Status: "Running",
					Spec:   summary.SummaryVmSpecInfo{Name: "g4dn.4xlarge", VCpus: 16, MemoryGiB: 64},
					Label:  map[string]string{"cm-source-machine-ids": "c1fb45d7-f1a0-465c-96f9-269c172d84df"},
				},
			},
		},
	}

	mappings := buildMigrationMappings(sourceSummary, targetSummary)
	if len(mappings) != 2 {
		t.Fatalf("buildMigrationMappings() returned %d mappings for 1:N GPU server, want 2", len(mappings))
	}
	if mappings[0].TargetVM.InstanceName != "ng-ai-gpu-01-a100-1" || mappings[1].TargetVM.InstanceName != "ng-ai-gpu-01-t4-1" {
		t.Errorf("unexpected target VM names: %q, %q", mappings[0].TargetVM.InstanceName, mappings[1].TargetVM.InstanceName)
	}
}

func TestMigrationReportGPUOmitemptyAndMarkdown(t *testing.T) {
	// 1. CPU-only mapping: GPU fields must be omitted in JSON
	cpuSource := &summary.SourceInfraSummary{
		Overview: summary.SourceSummaryInfraOverview{TotalServerCount: 1},
		ComputeResources: summary.SourceSummaryComputeResources{
			Servers: []summary.SourceServerInfo{
				{Hostname: "web-01", MachineId: "m-web-01", CPU: summary.SourceCPUInfo{Cores: 4}, Memory: summary.SourceMemoryInfo{TotalGB: 16}, Disk: summary.SourceDiskInfo{TotalGB: 100}},
			},
		},
	}
	cpuTarget := &summary.TargetInfraSummary{
		Overview: summary.TargetInfraOverview{TotalVmCount: 1, TargetCloud: "AWS", TargetRegion: "ap-northeast-2"},
		ComputeResources: summary.SummaryComputeResources{
			Vms: []summary.SummaryVmInfo{
				{
					Name:   "ng-web-01-1",
					Status: "Running",
					Spec:   summary.SummaryVmSpecInfo{Name: "t3.xlarge", VCpus: 4, MemoryGiB: 16},
					Label:  map[string]string{"cm-source-machine-ids": "m-web-01"},
				},
			},
		},
	}
	cpuMappings := buildMigrationMappings(cpuSource, cpuTarget)
	cpuExec := buildExecutiveSummary(cpuSource, cpuTarget)
	cpuReport := &MigrationReport{
		ExecutiveSummary:  cpuExec,
		MigrationMappings: cpuMappings,
		SourceDetails:     cpuSource,
		TargetDetails:     cpuTarget,
	}
	cpuBytes, err := json.Marshal(cpuReport)
	if err != nil {
		t.Fatalf("json.Marshal(cpuReport) failed: %v", err)
	}
	cpuJson := string(cpuBytes)
	for _, forbidden := range []string{`"totalSourceGpus"`, `"totalTargetGpus"`, `"gpuCount"`, `"gpuSummary"`, `"gpuChange"`, `"nlbMappings"`} {
		if strings.Contains(cpuJson, forbidden) {
			t.Errorf("CPU-only report JSON should omit %s, got: %s", forbidden, cpuJson)
		}
	}

	// 2. GPU mapping: GPU fields must be included in JSON and Markdown
	gpuSource := &summary.SourceInfraSummary{
		Overview: summary.SourceSummaryInfraOverview{TotalServerCount: 1, TotalGPUs: 1},
		ComputeResources: summary.SourceSummaryComputeResources{
			Servers: []summary.SourceServerInfo{
				{
					Hostname:  "gpu-01",
					MachineId: "m-gpu-01",
					CPU:       summary.SourceCPUInfo{CPUs: 1, Cores: 8, Threads: 16},
					Memory:    summary.SourceMemoryInfo{TotalGB: 64},
					Disk:      summary.SourceDiskInfo{TotalGB: 200},
					GPUCards: []summary.SourceGPUInfo{
						{Vendor: "NVIDIA", Model: "NVIDIA Tesla T4", MemoryTotalGB: 16},
					},
				},
			},
		},
	}
	gpuTarget := &summary.TargetInfraSummary{
		Overview: summary.TargetInfraOverview{TotalVmCount: 1, TotalGpuCount: 1, TargetCloud: "AWS", TargetRegion: "ap-northeast-2"},
		ComputeResources: summary.SummaryComputeResources{
			Vms: []summary.SummaryVmInfo{
				{
					Name:   "ng-gpu-01-1",
					Status: "Running",
					Spec: summary.SummaryVmSpecInfo{
						Name:                "g4dn.2xlarge",
						VCpus:               8,
						MemoryGiB:           32,
						AcceleratorType:     "gpu",
						AcceleratorModel:    "Tesla T4",
						AcceleratorCount:    1,
						AcceleratorMemoryGB: 16,
					},
					Label: map[string]string{"cm-source-machine-ids": "m-gpu-01"},
				},
			},
		},
	}
	gpuMappings := buildMigrationMappings(gpuSource, gpuTarget)
	gpuExec := buildExecutiveSummary(gpuSource, gpuTarget)
	gpuReport := &MigrationReport{
		ExecutiveSummary:  gpuExec,
		MigrationMappings: gpuMappings,
		SourceDetails:     gpuSource,
		TargetDetails:     gpuTarget,
	}
	gpuBytes, err := json.Marshal(gpuReport)
	if err != nil {
		t.Fatalf("json.Marshal(gpuReport) failed: %v", err)
	}
	gpuJson := string(gpuBytes)
	for _, want := range []string{`"totalSourceGpus":1`, `"totalTargetGpus":1`, `"gpuCount":1`, `"gpuSummary"`, `"gpuChange"`} {
		if !strings.Contains(gpuJson, want) {
			t.Errorf("GPU report JSON missing %s, got: %s", want, gpuJson)
		}
	}
	gpuMD := GenerateMigrationReportMarkdown(gpuReport)
	for _, wantMD := range []string{"**Total Source GPUs:** 1", "**Total Target GPUs:** 1", "**GPU:** 1 x Tesla T4 (16 GB)", "**GPU:** 1 card(s): NVIDIA Tesla T4 (16 GB)"} {
		if !strings.Contains(gpuMD, wantMD) {
			t.Errorf("GPU report markdown missing %q\n--- got ---\n%s", wantMD, gpuMD)
		}
	}
}

