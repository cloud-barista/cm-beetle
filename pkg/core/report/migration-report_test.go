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
	"testing"

	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
)

// TestBuildMigrationMappingsNlbGroup reproduces the 3-server sample used across
// the test-cli reports: one solo node whose NodeGroup has its own machine ID
// label, plus two NLB backend nodes sharing a single NodeGroup (and therefore
// the same comma-separated "sourceMachineIds" label). All three must resolve
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
		"sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
	}
	targetSummary := &summary.TargetInfraSummary{
		ComputeResources: summary.SummaryComputeResources{
			Vms: []summary.SummaryVmInfo{
				{
					Name:   "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
					Status: "Running",
					Spec:   summary.SummaryVmSpecInfo{Name: "t3a.small", VCpus: 2, MemoryGiB: 2},
					Label:  map[string]string{"sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"},
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
