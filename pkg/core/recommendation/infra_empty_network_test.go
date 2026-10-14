package recommendation

import (
	"testing"

	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/stretchr/testify/assert"
)

func TestRecommendVNet_EmptyNetworkInfo(t *testing.T) {
	srcInfra := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{
			{MachineId: "node-01"},
		},
	}

	csps := []struct {
		csp    string
		region string
	}{
		{"aws", "ap-northeast-2"},
		{"gcp", "asia-northeast3"},
		{"azure", "koreacentral"},
		{"ncp", "kr"},
	}

	for _, tc := range csps {
		t.Run(tc.csp, func(t *testing.T) {
			vnets, err := RecommendVNet(tc.csp, tc.region, srcInfra)
			assert.NoError(t, err)
			assert.NotEmpty(t, vnets)
			if len(vnets) > 0 {
				assert.Equal(t, DefaultTargetVNetCidr, vnets[0].CidrBlock)
				assert.NotEmpty(t, vnets[0].SubnetInfoList)
				assert.Equal(t, DefaultTargetSubnetCidr, vnets[0].SubnetInfoList[0].IPv4_CIDR)
				assert.Contains(t, vnets[0].Description, "Default target virtual network")
			}
		})
	}
}

func TestRecommendInfraCandidates_NoPanicOnEmptyNetwork(t *testing.T) {
	srcInfra := onpremmodel.OnpremInfra{
		Nodes: []onpremmodel.NodeProperty{
			{
				Hostname:  "test-node",
				MachineId: "test-node-01",
				CPU: onpremmodel.CpuProperty{
					Cpus:    2,
					Cores:   2,
					Threads: 4,
				},
				Memory: onpremmodel.MemoryProperty{
					TotalSize: 8,
				},
				RootDisk: onpremmodel.DiskProperty{
					TotalSize: 50,
				},
			},
		},
	}

	assert.NotPanics(t, func() {
		_, _ = RecommendInfraCandidates("aws", "ap-northeast-2", srcInfra, 1, 0.0)
	})
}
