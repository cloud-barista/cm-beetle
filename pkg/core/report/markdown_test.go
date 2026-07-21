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

func testSourceDetails() *summary.SourceInfraSummary {
	return &summary.SourceInfraSummary{
		ComputeResources: summary.SourceSummaryComputeResources{
			Servers: []summary.SourceServerInfo{
				{Hostname: "ip-10-0-1-30", MachineId: "ec268ed7-821e-9d73-e79f-961262161624"},
				{Hostname: "ip-10-0-1-221", MachineId: "ec2d32b5-98fb-5a96-7913-d3db1ec18932"},
				{Hostname: "ip-10-0-1-138", MachineId: "ec288dd0-c6fa-8a49-2f60-bc898311febf"},
			},
		},
	}
}

func TestVmGroupIndex(t *testing.T) {
	cases := []struct {
		vmName string
		want   int
	}{
		{"my-ng-influxdb-back-1", 1},
		{"my-ng-influxdb-back-2", 2},
		{"my-ng-influxdb-back-12", 12},
		{"my-ng-ec268ed7-821e-9d73-e79f-961262161624-1", 1},
		{"no-trailing-number", 0},
		{"", 0},
		{"trailing-dash-", 0},
	}

	for _, c := range cases {
		if got := vmGroupIndex(c.vmName); got != c.want {
			t.Errorf("vmGroupIndex(%q) = %d, want %d", c.vmName, got, c.want)
		}
	}
}

func TestResolveSourceMachineID(t *testing.T) {
	cases := []struct {
		name string
		vm   summary.SummaryVmInfo
		want string
	}{
		{
			name: "1:1 NodeGroup - single machine ID in label",
			vm: summary.SummaryVmInfo{
				Name:  "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
				Label: map[string]string{"sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"},
			},
			want: "ec268ed7-821e-9d73-e79f-961262161624",
		},
		{
			name: "NLB backend NodeGroup - 1st member picks 1st machine ID by position",
			vm: summary.SummaryVmInfo{
				Name:  "my-ng-influxdb-back-1",
				Label: map[string]string{"sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"},
			},
			want: "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
		},
		{
			name: "NLB backend NodeGroup - 2nd member picks 2nd machine ID by position",
			vm: summary.SummaryVmInfo{
				Name:  "my-ng-influxdb-back-2",
				Label: map[string]string{"sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"},
			},
			want: "ec288dd0-c6fa-8a49-2f60-bc898311febf",
		},
		{
			name: "no label",
			vm:   summary.SummaryVmInfo{Name: "my-ng-influxdb-back-1"},
			want: "",
		},
		{
			name: "multi-value label but VM name has no group index",
			vm: summary.SummaryVmInfo{
				Name:  "my-ng-influxdb-back",
				Label: map[string]string{"sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"},
			},
			want: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := resolveSourceMachineID(c.vm); got != c.want {
				t.Errorf("resolveSourceMachineID(%+v) = %q, want %q", c.vm, got, c.want)
			}
		})
	}
}

func TestFindSourceServer(t *testing.T) {
	sourceDetails := testSourceDetails()

	cases := []struct {
		name         string
		vm           summary.SummaryVmInfo
		wantHostname string // "" means no match expected
	}{
		{
			name: "resolved via label (1:1 NodeGroup)",
			vm: summary.SummaryVmInfo{
				Name:  "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
				Label: map[string]string{"sourceMachineIds": "ec268ed7-821e-9d73-e79f-961262161624"},
			},
			wantHostname: "ip-10-0-1-30",
		},
		{
			name: "resolved via label (NLB backend NodeGroup, 2nd member)",
			vm: summary.SummaryVmInfo{
				Name:  "my-ng-influxdb-back-2",
				Label: map[string]string{"sourceMachineIds": "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf"},
			},
			wantHostname: "ip-10-0-1-138",
		},
		{
			name:         "no label - falls back to substring match (infra-with-nlb naming)",
			vm:           summary.SummaryVmInfo{Name: "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1"},
			wantHostname: "ip-10-0-1-30",
		},
		{
			name:         "no label - falls back to substring match (infra naming)",
			vm:           summary.SummaryVmInfo{Name: "my01-vm-ec288dd0-c6fa-8a49-2f60-bc898311febf-1"},
			wantHostname: "ip-10-0-1-138",
		},
		{
			name:         "role-named VM with no label and no embedded machine ID",
			vm:           summary.SummaryVmInfo{Name: "my-ng-influxdb-back-1"},
			wantHostname: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			node := findSourceServer(c.vm, sourceDetails)
			if c.wantHostname == "" {
				if node != nil {
					t.Errorf("findSourceServer(%+v) = %+v, want no match", c.vm, node)
				}
				return
			}

			if node == nil {
				t.Fatalf("findSourceServer(%+v) = nil, want hostname %q", c.vm, c.wantHostname)
			}
			if node.Hostname != c.wantHostname {
				t.Errorf("findSourceServer(%+v).Hostname = %q, want %q", c.vm, node.Hostname, c.wantHostname)
			}
		})
	}
}

func TestFindSourceServerNilDetails(t *testing.T) {
	vm := summary.SummaryVmInfo{Name: "my-ng-ec268ed7-821e-9d73-e79f-961262161624-1"}
	if node := findSourceServer(vm, nil); node != nil {
		t.Errorf("findSourceServer with nil sourceDetails = %+v, want nil", node)
	}
}
