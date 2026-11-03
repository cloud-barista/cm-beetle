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

	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
)

func sampleSourceInfraSummaryWithNLB() *summary.SourceInfraSummary {
	return &summary.SourceInfraSummary{
		Overview: summary.SourceSummaryInfraOverview{
			InfraName:        "infra-3-servers",
			TotalServerCount: 3,
			TotalNLBs:        1,
		},
		ComputeResources: summary.SourceSummaryComputeResources{
			Servers: []summary.SourceServerInfo{
				{
					Hostname:  "ip-10-0-1-30",
					MachineId: "ec268ed7-821e-9d73-e79f-961262161624",
					Network:   summary.SourceNetworkBrief{IPAddress: "10.0.1.30/24"},
				},
				{
					Hostname:  "ip-10-0-1-221",
					MachineId: "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
					Network:   summary.SourceNetworkBrief{IPAddress: "10.0.1.221/24"},
				},
				{
					Hostname:  "ip-10-0-1-138",
					MachineId: "ec288dd0-c6fa-8a49-2f60-bc898311febf",
					Network:   summary.SourceNetworkBrief{IPAddress: "10.0.1.138/24"},
				},
			},
		},
		NetworkResources: summary.SourceSummaryNetworkResources{
			NLBs: []summary.SourceNlbInfo{
				{
					Software:      "haproxy",
					HostMachineId: "ec268ed7-821e-9d73-e79f-961262161624",
					HostHostname:  "ip-10-0-1-30",
					BindAddress:   "*",
					ListenerPort:  9999,
					ListenerProto: "tcp",
					BackendName:   "influxdb_back",
					BackendProto:  "tcp",
					BackendPort:   8086,
					Balance:       "roundrobin",
					BackendServers: []summary.SourceNlbBackendServerInfo{
						{
							Name:      "influx1",
							IP:        "10.0.1.221",
							Port:      8086,
							Weight:    1,
							MachineId: "ec2d32b5-98fb-5a96-7913-d3db1ec18932",
							Hostname:  "ip-10-0-1-221",
						},
						{
							Name:      "influx2",
							IP:        "10.0.1.138",
							Port:      8086,
							Weight:    1,
							MachineId: "ec288dd0-c6fa-8a49-2f60-bc898311febf",
							Hostname:  "ip-10-0-1-138",
						},
					},
					HealthCheck: summary.SourceNlbHealthCheckInfo{
						Enabled:   true,
						Protocol:  "tcp",
						Port:      8086,
						Interval:  10,
						Timeout:   10,
						Threshold: 3,
					},
				},
			},
		},
		SecurityResources: summary.SourceSummarySecurityResources{
			ServerFirewalls: []summary.SourceServerFirewall{
				{
					Hostname: "ip-10-0-1-30",
					FirewallRules: []summary.SourceFirewallRule{
						{Direction: "inbound", Protocol: "tcp", DstPorts: "9999", SrcCIDR: "0.0.0.0/0", Action: "allow"},
						{Direction: "inbound", Protocol: "tcp", DstPorts: "22", SrcCIDR: "0.0.0.0/0", Action: "allow"},
					},
				},
				{
					Hostname: "ip-10-0-1-221",
					FirewallRules: []summary.SourceFirewallRule{
						{Direction: "inbound", Protocol: "tcp", DstPorts: "22", SrcCIDR: "0.0.0.0/0", Action: "allow"},
						{Direction: "inbound", Protocol: "tcp", DstPorts: "8086", SrcCIDR: "10.0.0.0/16", Action: "allow"},
					},
				},
				{
					Hostname: "ip-10-0-1-138",
					FirewallRules: []summary.SourceFirewallRule{
						{Direction: "inbound", Protocol: "tcp", DstPorts: "22", SrcCIDR: "0.0.0.0/0", Action: "allow"},
						{Direction: "inbound", Protocol: "tcp", DstPorts: "8086", SrcCIDR: "10.0.0.0/16", Action: "allow"},
					},
				},
			},
		},
	}
}

func TestBuildNlbMigrationMappingsAndMarkdown_AWS(t *testing.T) {
	srcSummary := sampleSourceInfraSummaryWithNLB()
	targetSummary := &summary.TargetInfraSummary{
		Overview: summary.TargetInfraOverview{
			InfraName:      "my01-infra101",
			TargetCloud:    "AWS",
			TargetRegion:   "ap-northeast-2",
			TotalVmCount:   3,
			RunningVmCount: 3,
			TotalNlbCount:  1,
		},
		ComputeResources: summary.SummaryComputeResources{
			Vms: []summary.SummaryVmInfo{
				{
					Id:          "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
					Name:        "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1",
					NodeGroupId: "my01-ng-ec268ed7-821e-9d73-e79f-961262161624",
					CspVmId:     "i-0e9a8dff79587097e",
					Status:      "Running",
					Misc: summary.SummaryVmMiscInfo{
						PrivateIp:      "10.0.1.201",
						PublicIp:       "54.180.141.244",
						SecurityGroups: []string{"my01-mig-sg-02"},
					},
					Label: map[string]string{
						lineage.LabelSourceMachineIds: "ec268ed7-821e-9d73-e79f-961262161624",
					},
				},
				{
					Id:          "my01-ng-influxdb-back-1",
					Name:        "my01-ng-influxdb-back-1",
					NodeGroupId: "my01-ng-influxdb-back",
					CspVmId:     "i-0dee049fb9412465b",
					Status:      "Running",
					Misc: summary.SummaryVmMiscInfo{
						PrivateIp:      "10.0.1.221",
						PublicIp:       "13.125.79.237",
						SecurityGroups: []string{"my01-mig-sg-01"},
					},
					Label: map[string]string{
						lineage.LabelSourceMachineIds: "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
						lineage.LabelNlbBackend:       "influxdb_back",
					},
				},
				{
					Id:          "my01-ng-influxdb-back-2",
					Name:        "my01-ng-influxdb-back-2",
					NodeGroupId: "my01-ng-influxdb-back",
					CspVmId:     "i-0cf1d4327934156dd",
					Status:      "Running",
					Misc: summary.SummaryVmMiscInfo{
						PrivateIp:      "10.0.1.253",
						PublicIp:       "3.35.54.160",
						SecurityGroups: []string{"my01-mig-sg-01"},
					},
					Label: map[string]string{
						lineage.LabelSourceMachineIds: "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
						lineage.LabelNlbBackend:       "influxdb_back",
					},
				},
			},
		},
		NetworkResources: summary.SummaryNetworkResources{
			NLBs: []summary.SummaryNlbInfo{
				{
					Id:              "my01-ng-influxdb-back",
					Name:            "my01-ng-influxdb-back",
					CspResourceName: "tbsd5l543sjssm4r416n",
					ConnectionName:  "aws-ap-northeast-2",
					Type:            "PUBLIC",
					Scope:           "REGION",
					Listener: summary.SummaryNlbListenerInfo{
						Protocol: "TCP",
						Port:     "9999",
						DnsName:  "tbsd5l543sjssm4r416n-ef49.elb.ap-northeast-2.amazonaws.com",
					},
					TargetGroup: summary.SummaryNlbTargetGroupInfo{
						Protocol:    "TCP",
						Port:        "8086",
						NodeGroupId: "my01-ng-influxdb-back",
						Nodes:       []string{"my01-ng-influxdb-back-1", "my01-ng-influxdb-back-2"},
					},
					HealthChecker: summary.SummaryNlbHealthCheckerInfo{
						Protocol:  "TCP",
						Port:      "8086",
						Interval:  10,
						Timeout:   10,
						Threshold: 3,
					},
				},
			},
		},
		SecurityResources: summary.SummarySecurityResources{
			SecurityGroups: []summary.SummarySecurityGroupInfo{
				{
					Name:               "my01-mig-sg-01",
					CspSecurityGroupId: "sg-0704eaa2a1108babc",
					VNetName:           "my01-mig-vnet-01",
					RuleCount:          3,
					Rules: []summary.SummaryFirewallRule{
						{Direction: "inbound", Protocol: "TCP", FromPort: "9999", ToPort: "9999", Cidr: "0.0.0.0/0"},
						{Direction: "inbound", Protocol: "TCP", FromPort: "8086", ToPort: "8086", Cidr: "10.0.0.0/16"},
						{Direction: "inbound", Protocol: "TCP", FromPort: "8086", ToPort: "8086", Cidr: "0.0.0.0/0"},
					},
				},
			},
		},
	}

	nlbMappings := buildNlbMigrationMappings(srcSummary, targetSummary)
	if len(nlbMappings) != 1 {
		t.Fatalf("len(nlbMappings) = %d, want 1", len(nlbMappings))
	}
	if nlbMappings[0].MigrationStatus != "Migrated" {
		t.Errorf("MigrationStatus = %q, want Migrated", nlbMappings[0].MigrationStatus)
	}
	if nlbMappings[0].HostNodeTargetVmName != "my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1" {
		t.Errorf("HostNodeTargetVmName = %q, want my01-ng-ec268ed7-821e-9d73-e79f-961262161624-1", nlbMappings[0].HostNodeTargetVmName)
	}

	netAnalysis := buildNetworkAnalysis(srcSummary, targetSummary)
	if len(netAnalysis.IPMappings) != 3 {
		t.Errorf("len(netAnalysis.IPMappings) = %d, want 3 (including merged NLB backend VMs)", len(netAnalysis.IPMappings))
	}

	execSummary := buildExecutiveSummary(srcSummary, targetSummary)
	nlbStatus, totalNLBs, migratedNLBs, failedNLBs := computeNlbMigrationStatus(nlbMappings)
	execSummary.NlbMigrationStatus = nlbStatus
	execSummary.TotalNLBs = totalNLBs
	execSummary.MigratedNLBs = migratedNLBs
	execSummary.FailedNLBs = failedNLBs

	report := &MigrationReport{
		Metadata:         ReportMetadata{Namespace: "mig01", InfraID: "my01-infra101"},
		ExecutiveSummary: execSummary,
		NlbMappings:      nlbMappings,
		NetworkAnalysis:  netAnalysis,
		SourceDetails:    srcSummary,
		TargetDetails:    targetSummary,
	}

	md := GenerateMigrationReportMarkdown(report)
	for _, wantSub := range []string{
		"**NLB Migration Status:** Completed",
		"| 8 | **Network Load Balancer** | 1 | ✅ Created | my01-ng-influxdb-back (TCP:9999 -> TCP:8086) |",
		"## ⚖️ Network Load Balancers (NLBs)",
		"### 1. influxdb_back (haproxy) → my01-ng-influxdb-back",
		"**Configuration Comparison:**",
		"**Backend Node Mappings:**",
		"| 1 | influx1 (ip-10-0-1-221) | 10.0.1.221:8086 | my01-ng-influxdb-back-1 | i-0dee049fb9412465b | 10.0.1.221 | 13.125.79.237 | 8086 | Running |",
		"| 1 | inbound | TCP | 9999 | 0.0.0.0/0 | - | Created by system |",
		"| 2 | inbound | TCP | 8086 | 10.0.0.0/16 | inbound tcp 8086 from 10.0.0.0/16 | Migrated from source |",
		"| 3 | inbound | TCP | 8086 | 0.0.0.0/0 | - | Created by system |",
	} {
		if !strings.Contains(md, wantSub) {
			t.Errorf("migration report markdown missing %q\n--- got ---\n%s", wantSub, md)
		}
	}
}

func TestBuildNlbMigrationMappings_CspConstraints_GCP_Azure_IBM(t *testing.T) {
	srcSummary := sampleSourceInfraSummaryWithNLB()

	// Test Stage 2 matching (without nlbBackend label) + GCP port adjustment + Azure -1 timeout + IBM timeout adjustment.
	gcpTarget := &summary.TargetInfraSummary{
		Overview: summary.TargetInfraOverview{TargetCloud: "GCP"},
		ComputeResources: summary.SummaryComputeResources{
			Vms: []summary.SummaryVmInfo{
				{
					Name:        "my01-ng-influxdb-back-1",
					NodeGroupId: "my01-ng-influxdb-back",
					Label: map[string]string{
						lineage.LabelSourceMachineIds: "ec2d32b5-98fb-5a96-7913-d3db1ec18932,ec288dd0-c6fa-8a49-2f60-bc898311febf",
					},
				},
			},
		},
		NetworkResources: summary.SummaryNetworkResources{
			NLBs: []summary.SummaryNlbInfo{
				{
					Name:           "my01-ng-influxdb-back",
					Type:           "PUBLIC",
					Scope:          "REGION",
					Listener:       summary.SummaryNlbListenerInfo{Protocol: "TCP", Port: "8086", IP: "34.64.10.20"},
					TargetGroup:    summary.SummaryNlbTargetGroupInfo{Protocol: "TCP", Port: "8086", NodeGroupId: "my01-ng-influxdb-back", Nodes: []string{"my01-ng-influxdb-back-1"}},
					HealthChecker:  summary.SummaryNlbHealthCheckerInfo{Protocol: "TCP", Port: "8086", Interval: 10, Timeout: 10, Threshold: 3},
					ConnectionName: "gcp-asia-northeast3",
				},
			},
		},
	}

	gcpMappings := buildNlbMigrationMappings(srcSummary, gcpTarget)
	if len(gcpMappings) != 1 || gcpMappings[0].MigrationStatus != "Migrated" {
		t.Fatalf("GCP Stage-2 NLB matching failed: %+v", gcpMappings)
	}
	notesJoined := strings.Join(gcpMappings[0].Notes, " ")
	if !strings.Contains(notesJoined, "Listener port was adjusted from 9999 to 8086") {
		t.Errorf("expected GCP port adjustment note, got: %v", gcpMappings[0].Notes)
	}

	azureTarget := *gcpTarget
	azureTarget.Overview.TargetCloud = "AZURE"
	azureTarget.NetworkResources.NLBs = []summary.SummaryNlbInfo{
		{
			Name:          "my01-ng-influxdb-back",
			Type:          "PUBLIC",
			Scope:         "REGION",
			Listener:      summary.SummaryNlbListenerInfo{Protocol: "TCP", Port: "9999", IP: "20.196.1.1"},
			TargetGroup:   summary.SummaryNlbTargetGroupInfo{Protocol: "TCP", Port: "8086", NodeGroupId: "my01-ng-influxdb-back"},
			HealthChecker: summary.SummaryNlbHealthCheckerInfo{Protocol: "TCP", Port: "8086", Interval: 10, Timeout: -1, Threshold: 3},
		},
	}
	azureMappings := buildNlbMigrationMappings(srcSummary, &azureTarget)
	if !strings.Contains(strings.Join(azureMappings[0].Notes, " "), "Health check timeout uses CSP default setting on AZURE") {
		t.Errorf("expected Azure CSP default timeout note, got: %v", azureMappings[0].Notes)
	}

	ibmTarget := *gcpTarget
	ibmTarget.Overview.TargetCloud = "IBM"
	ibmTarget.NetworkResources.NLBs = []summary.SummaryNlbInfo{
		{
			Name:          "my01-ng-influxdb-back",
			Type:          "PUBLIC",
			Scope:         "REGION",
			Listener:      summary.SummaryNlbListenerInfo{Protocol: "TCP", Port: "9999", DnsName: "ibm-nlb.lb.appdomain.cloud"},
			TargetGroup:   summary.SummaryNlbTargetGroupInfo{Protocol: "TCP", Port: "8086", NodeGroupId: "my01-ng-influxdb-back"},
			HealthChecker: summary.SummaryNlbHealthCheckerInfo{Protocol: "TCP", Port: "8086", Interval: 10, Timeout: 5, Threshold: 3},
		},
	}
	ibmMappings := buildNlbMigrationMappings(srcSummary, &ibmTarget)
	if !strings.Contains(strings.Join(ibmMappings[0].Notes, " "), "Health check timeout was adjusted from 10s to 5s") {
		t.Errorf("expected IBM timeout adjustment note, got: %v", ibmMappings[0].Notes)
	}
}

func TestBuildNlbMigrationMappings_NotMigratedAndZeroNlbBackwardCompat(t *testing.T) {
	srcSummary := sampleSourceInfraSummaryWithNLB()
	emptyTarget := &summary.TargetInfraSummary{
		Overview: summary.TargetInfraOverview{
			TargetCloud:  "AWS",
			TargetRegion: "ap-northeast-2",
			TotalVmCount: 3,
		},
	}

	unmigrated := buildNlbMigrationMappings(srcSummary, emptyTarget)
	status, total, migrated, failed := computeNlbMigrationStatus(unmigrated)
	if status != "Not Migrated" || total != 1 || migrated != 0 || failed != 1 {
		t.Fatalf("computeNlbMigrationStatus = (%q, %d, %d, %d), want (Not Migrated, 1, 0, 1)", status, total, migrated, failed)
	}

	report := &MigrationReport{
		Metadata: ReportMetadata{Namespace: "mig01", InfraID: "my01-infra101"},
		ExecutiveSummary: ExecutiveSummary{
			MigrationStatus:    "Completed",
			TotalServers:       3,
			MigratedServers:    3,
			NlbMigrationStatus: status,
			TotalNLBs:          total,
			MigratedNLBs:       migrated,
			FailedNLBs:         failed,
		},
		NlbMappings:   unmigrated,
		SourceDetails: srcSummary,
		TargetDetails: emptyTarget,
	}
	md := GenerateMigrationReportMarkdown(report)
	if !strings.Contains(md, "⚠️ **Unmigrated Load Balancers:** 1") ||
		!strings.Contains(md, "| 8 | **Network Load Balancer** | 0 | ⚠️ Not created | 0 of 1 migrated |") {
		t.Errorf("expected unmigrated NLB indicators in markdown:\n%s", md)
	}

	// Verify 100% backward compatibility when both Source and Target have 0 NLBs.
	plainReport := &MigrationReport{
		Metadata: ReportMetadata{Namespace: "mig01", InfraID: "plain-infra"},
		ExecutiveSummary: ExecutiveSummary{
			MigrationStatus: "Completed",
			TotalServers:    1,
			MigratedServers: 1,
		},
		SourceDetails: &summary.SourceInfraSummary{},
		TargetDetails: emptyTarget,
	}
	plainMD := GenerateMigrationReportMarkdown(plainReport)
	if strings.Contains(plainMD, "Network Load Balancer") || strings.Contains(plainMD, "Load Balancers") {
		t.Errorf("0-NLB migration report should not render any NLB rows or sections:\n%s", plainMD)
	}
}
