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

package summary

import (
	"strings"
	"testing"

	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
)

func TestIsMainInterface(t *testing.T) {
	cases := []struct {
		name string
		want bool
	}{
		{"lo", true},
		{"eth0", true},
		{"eno1", true},
		{"ens5", true},   // systemd PCI-slot predictable name (e.g. AWS/KVM guests)
		{"enp0s3", true}, // systemd PCI-ID predictable name
		{"enx001122334455", true},
		{"br-ex", true},
		{"tap0", false},
		{"veth1234", false},
		{"docker0", false},
		{"", false},
	}

	for _, c := range cases {
		if got := isMainInterface(c.name); got != c.want {
			t.Errorf("isMainInterface(%q) = %v, want %v", c.name, got, c.want)
		}
	}
}

func TestGenerateSourceInfraSummaryWithNLB(t *testing.T) {
	infra := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{
			{
				Hostname:  "ip-10-0-1-30",
				MachineId: "ec268ed7-821e-9d73-e79f-961262161624",
				Interfaces: []onpremmodel.NetworkInterfaceProperty{
					{Name: "ens5", IPv4CidrBlocks: []string{"10.0.1.30/24"}, State: "up"},
				},
			},
			{
				Hostname:  "ip-10-0-1-221",
				MachineId: "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
				Interfaces: []onpremmodel.NetworkInterfaceProperty{
					{Name: "ens5", IPv4CidrBlocks: []string{"10.0.1.221/24"}, State: "up"},
				},
			},
			{
				Hostname:  "ip-10-0-1-138",
				MachineId: "ec288dd0-c6fa-8a49-2f60-bc898311febf",
				Interfaces: []onpremmodel.NetworkInterfaceProperty{
					{Name: "ens5", IPv4CidrBlocks: []string{"10.0.1.138/24"}, State: "up"},
				},
			},
		},
		NLBs: []onpremmodel.NlbProperty{
			{
				HostMachineId: "ec268ed7-821e-9d73-e79f-961262161624",
				Software:      "haproxy",
				Listener: onpremmodel.NlbListenerProperty{
					BindAddress: "*",
					Port:        9999,
					Protocol:    "tcp",
				},
				Backend: onpremmodel.NlbBackendProperty{
					Name:     "influxdb_back",
					Balance:  "roundrobin",
					Protocol: "tcp",
					Servers: []onpremmodel.NlbServerProperty{
						{Name: "influx1", IP: "10.0.1.221", Port: 8086, Weight: 1},
						{Name: "influx2", IP: "10.0.1.138", Port: 8086, Weight: 1},
					},
				},
				HealthCheck: onpremmodel.NlbHealthCheckProperty{
					Enabled:   true,
					Interval:  10,
					Timeout:   10,
					Threshold: 3,
				},
			},
		},
	}

	srcSummary, err := GenerateSourceInfraSummary("test-nlb-infra", infra)
	if err != nil {
		t.Fatalf("GenerateSourceInfraSummary failed: %v", err)
	}

	if srcSummary.Overview.TotalNLBs != 1 {
		t.Fatalf("Overview.TotalNLBs = %d, want 1", srcSummary.Overview.TotalNLBs)
	}
	if len(srcSummary.NetworkResources.NLBs) != 1 {
		t.Fatalf("len(NetworkResources.NLBs) = %d, want 1", len(srcSummary.NetworkResources.NLBs))
	}

	nlb := srcSummary.NetworkResources.NLBs[0]
	if nlb.HostHostname != "ip-10-0-1-30" {
		t.Errorf("HostHostname = %q, want %q", nlb.HostHostname, "ip-10-0-1-30")
	}
	if nlb.BackendPort != 8086 {
		t.Errorf("BackendPort = %d, want 8086", nlb.BackendPort)
	}
	if len(nlb.BackendServers) != 2 {
		t.Fatalf("len(BackendServers) = %d, want 2", len(nlb.BackendServers))
	}
	if nlb.BackendServers[0].Hostname != "ip-10-0-1-221" || nlb.BackendServers[0].MachineId != "ec2d32b5-98fb-5a96-7913-d3db1ec18932" {
		t.Errorf("BackendServers[0] mapping = (%q, %q), want (ip-10-0-1-221, ec2d32b5-98fb-5a96-7913-d3db1ec18932)",
			nlb.BackendServers[0].Hostname, nlb.BackendServers[0].MachineId)
	}

	md := GenerateSourceMarkdownSummary(srcSummary)
	for _, wantSub := range []string{
		"| Total Load Balancers | 1 |",
		"### Network Load Balancers (1)",
		"#### 1. influxdb_back (haproxy on ip-10-0-1-30)",
		"| Frontend Listener | *:9999 (TCP) |",
		"**Backend Servers:**",
		"| influx1 | 10.0.1.221:8086 | 1 | ip-10-0-1-221 | ec2d32b5-98fb-5a96-7913-d3db1ec18932 |",
	} {
		if !strings.Contains(md, wantSub) {
			t.Errorf("source summary markdown missing %q\n--- got ---\n%s", wantSub, md)
		}
	}
}

func TestGenerateSourceInfraSummaryWithoutNLB(t *testing.T) {
	infra := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{
			{Hostname: "web-01", MachineId: "m-01"},
		},
	}

	srcSummary, err := GenerateSourceInfraSummary("plain-infra", infra)
	if err != nil {
		t.Fatalf("GenerateSourceInfraSummary failed: %v", err)
	}
	if srcSummary.Overview.TotalNLBs != 0 {
		t.Errorf("TotalNLBs = %d, want 0", srcSummary.Overview.TotalNLBs)
	}

	md := GenerateSourceMarkdownSummary(srcSummary)
	if strings.Contains(md, "Total Load Balancers") || strings.Contains(md, "Network Load Balancers") {
		t.Errorf("non-NLB source summary should not render NLB sections:\n%s", md)
	}
}

func TestGenerateTargetMarkdownSummaryWithAndWithoutNLB(t *testing.T) {
	plainSummary := &TargetInfraSummary{
		Overview: TargetInfraOverview{
			InfraName:      "plain-infra",
			TotalVmCount:   1,
			RunningVmCount: 1,
			TotalNlbCount:  0,
		},
	}
	plainMD := GenerateMarkdownSummary(plainSummary)
	if strings.Contains(plainMD, "Total NLBs") || strings.Contains(plainMD, "Network Load Balancers (NLB)") {
		t.Errorf("non-NLB target summary should not render NLB sections:\n%s", plainMD)
	}

	nlbSummary := &TargetInfraSummary{
		Overview: TargetInfraOverview{
			InfraName:      "my01-infra101",
			TotalVmCount:   3,
			RunningVmCount: 3,
			TotalNlbCount:  1,
		},
		ComputeResources: SummaryComputeResources{
			Vms: []SummaryVmInfo{
				{
					Id:          "my01-ng-influxdb-back-1",
					Name:        "my01-ng-influxdb-back-1",
					NodeGroupId: "my01-ng-influxdb-back",
					CspVmId:     "i-0dee049fb9412465b",
					Status:      "Running",
					Misc: SummaryVmMiscInfo{
						PrivateIp: "10.0.1.221",
						PublicIp:  "13.125.79.237",
					},
				},
			},
		},
		NetworkResources: SummaryNetworkResources{
			NLBs: []SummaryNlbInfo{
				{
					Id:              "my01-ng-influxdb-back",
					Name:            "my01-ng-influxdb-back",
					CspResourceName: "tbsd5l543sjssm4r416n",
					CspResourceId:   "arn:aws:elasticloadbalancing:ap-northeast-2:123:loadbalancer/net/tbsd5l543sjssm4r416n/ef49",
					ConnectionName:  "aws-ap-northeast-2",
					Type:            "PUBLIC",
					Scope:           "REGION",
					Listener: SummaryNlbListenerInfo{
						Protocol: "TCP",
						Port:     "9999",
						DnsName:  "tbsd5l543sjssm4r416n-ef49.elb.ap-northeast-2.amazonaws.com",
					},
					TargetGroup: SummaryNlbTargetGroupInfo{
						Protocol:    "TCP",
						Port:        "8086",
						NodeGroupId: "my01-ng-influxdb-back",
						Nodes:       []string{"my01-ng-influxdb-back-1"},
					},
					HealthChecker: SummaryNlbHealthCheckerInfo{
						Protocol:  "TCP",
						Port:      "8086",
						Interval:  10,
						Timeout:   10,
						Threshold: 3,
					},
				},
			},
		},
	}

	nlbMD := GenerateMarkdownSummary(nlbSummary)
	for _, wantSub := range []string{
		"| **Total NLBs** | 1 |",
		"### Network Load Balancers (NLB)",
		"#### NLB: my01-ng-influxdb-back",
		"**Target Group Nodes:**",
		"| my01-ng-influxdb-back-1 | i-0dee049fb9412465b | 10.0.1.221 | 13.125.79.237 | Running |",
	} {
		if !strings.Contains(nlbMD, wantSub) {
			t.Errorf("target NLB summary markdown missing %q\n--- got ---\n%s", wantSub, nlbMD)
		}
	}
}
