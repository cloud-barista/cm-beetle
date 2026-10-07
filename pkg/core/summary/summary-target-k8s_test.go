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
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/stretchr/testify/assert"
)

// TestGenerateMarkdownK8sInfraSummary verifies markdown rendering for K8s infrastructure summary.
func TestGenerateMarkdownK8sInfraSummary(t *testing.T) {
	summary := &TargetK8sInfraSummary{
		SummaryMetadata: TargetK8sSummaryMetadata{
			GeneratedAt:    time.Date(2026, 10, 7, 15, 0, 0, 0, time.UTC),
			Namespace:      "mig01",
			ClusterId:      "test-cluster",
			ClusterName:    "test-cluster",
			SummaryVersion: "1.0",
		},
		Overview: TargetK8sInfraOverview{
			ClusterName:          "test-cluster",
			Status:               "Active",
			TargetCloud:          "AWS",
			TargetRegion:         "ap-northeast-2",
			ConnectionName:       "aws-ap-northeast-2",
			K8sVersion:           "1.34",
			CspClusterId:         "arn:aws:eks:ap-northeast-2:1234567890:cluster/test-cluster",
			TotalNodeGroupCount:  1,
			TotalWorkerNodeCount: 2,
			ActiveNodeCount:      2,
		},
		ComputeOverview: K8sComputeOverview{
			TotalVcpus:     8,
			TotalMemoryGib: 16.0,
			TotalStorageGB: 100,
		},
		NodeGroupList: []SummaryNodeGroupInfo{
			{
				Name:            "workers1",
				Status:          "Active",
				SpecId:          "c5a.xlarge",
				Vcpu:            4,
				MemoryGib:       8.0,
				RootDiskType:    "gp3",
				RootDiskSizeGB:  50,
				OnAutoScaling:   "false",
				DesiredNodeSize: 2,
				MinNodeSize:     0,
				MaxNodeSize:     0,
				HourlyCostUSD:   0.308,
				MonthlyCostUSD:  221.76,
				Nodes: []SummaryK8sNodeInfo{
					{
						CspResourceName: "i-0123456789abcdef0",
						CspResourceId:   "csp-node-1",
					},
					{
						CspResourceName: "i-0123456789abcdef1",
						CspResourceId:   "csp-node-2",
					},
				},
			},
		},
		NetworkResources: SummaryNetworkResources{
			VNets: []SummaryVNetInfo{
				{
					Name:        "mig-vnet",
					CspVNetId:   "vpc-123456",
					CidrBlock:   "10.0.0.0/16",
					SubnetCount: 2,
				},
			},
		},
		SecurityResources: SummarySecurityResources{
			SecurityGroups: []SummarySecurityGroupInfo{
				{
					Name:               "mig-sg",
					CspSecurityGroupId: "sg-123456",
					RuleCount:          1,
				},
			},
		},
		CostEstimation: SummaryCostEstimation{
			Currency:          "USD",
			TotalCostPerHour:  0.308,
			TotalCostPerDay:   7.392,
			TotalCostPerMonth: 221.76,
		},
	}

	md := GenerateMarkdownK8sInfraSummary(summary)

	assert.Contains(t, md, "# Target Cloud Kubernetes Infrastructure Summary")
	assert.Contains(t, md, "test-cluster")
	assert.Contains(t, md, "aws-ap-northeast-2")
	assert.Contains(t, md, "1.34")
	assert.Contains(t, md, "workers1")
	assert.Contains(t, md, "c5a.xlarge")
	assert.Contains(t, md, "i-0123456789abcdef0")
	assert.Contains(t, md, "$221.76")
}

// TestBuildTargetK8sInfraOverview tests overview aggregation logic.
func TestBuildTargetK8sInfraOverview(t *testing.T) {
	clusterInfo := &tbmodel.K8sClusterInfo{
		Name:           "k8s-prod",
		Status:         "Active",
		ConnectionName: "gcp-asia-northeast3",
		Version:        "1.34",
		CspResourceId:  "gke-cluster-123",
		ConnectionConfig: tbmodel.ConnConfig{
			ProviderName: "gcp",
			RegionDetail: tbmodel.RegionDetail{
				RegionName: "asia-northeast3",
			},
		},
		K8sNodeGroupList: []tbmodel.K8sNodeGroupInfo{
			{
				Name:            "pool-1",
				DesiredNodeSize: 3,
			},
		},
	}

	nodeGroups := []SummaryNodeGroupInfo{
		{
			Name:            "pool-1",
			DesiredNodeSize: 3,
			Status:          "Active",
			Nodes: []SummaryK8sNodeInfo{
				{CspResourceName: "node-1"},
				{CspResourceName: "node-2"},
				{CspResourceName: "node-3"},
			},
		},
	}

	overview := buildTargetK8sInfraOverview(clusterInfo, nodeGroups)

	assert.Equal(t, "k8s-prod", overview.ClusterName)
	assert.Equal(t, "GCP", overview.TargetCloud)
	assert.Equal(t, "asia-northeast3", overview.TargetRegion)
	assert.Equal(t, "1.34", overview.K8sVersion)
	assert.Equal(t, 1, overview.TotalNodeGroupCount)
	assert.Equal(t, 3, overview.TotalWorkerNodeCount)
	assert.Equal(t, 3, overview.ActiveNodeCount)
}

// TestBuildK8sCostEstimation tests cost aggregation logic for node groups.
func TestBuildK8sCostEstimation(t *testing.T) {
	clusterInfo := &tbmodel.K8sClusterInfo{
		ConnectionName: "aws-ap-northeast-2",
	}

	nodeGroups := []SummaryNodeGroupInfo{
		{
			Name:            "workers",
			SpecId:          "c5a.xlarge",
			DesiredNodeSize: 2,
			HourlyCostUSD:   0.30,
			MonthlyCostUSD:  216.0,
		},
	}

	cost := buildK8sCostEstimation(clusterInfo, 0.30, nodeGroups)

	assert.Equal(t, "USD", cost.Currency)
	assert.Equal(t, float32(0.30), cost.TotalCostPerHour)
	assert.Equal(t, float32(216.0), cost.TotalCostPerMonth)
	assert.Len(t, cost.ByVm, 1)
	assert.True(t, strings.Contains(cost.ByVm[0].VmName, "nodegroup:workers"))
}

// TestExtractUniqueK8sSshKeyIds verifies unique and sorted SSH Key extraction.
func TestExtractUniqueK8sSshKeyIds(t *testing.T) {
	nodeGroups := []tbmodel.K8sNodeGroupInfo{
		{Name: "ng1", SshKeyId: "key-b"},
		{Name: "ng2", SshKeyId: "key-a"},
		{Name: "ng3", SshKeyId: "key-b"},
		{Name: "ng4", SshKeyId: ""},
	}
	keys := extractUniqueK8sSshKeyIds(nodeGroups)
	assert.Equal(t, []string{"key-a", "key-b"}, keys)
}

