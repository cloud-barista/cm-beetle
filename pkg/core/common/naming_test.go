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

package common

import (
	"fmt"
	"sync"
	"testing"
	"time"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sampleRecommendedInfra() cloudmodel.RecommendedInfra {
	return cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{
			Name: "mig-vnet-01",
			SubnetInfoList: []cloudmodel.SubnetReq{
				{Name: "mig-subnet-01"},
			},
		},
		TargetSshKey: cloudmodel.SshKeyReq{
			Name: "mig-sshkey-01",
		},
		TargetSecurityGroupList: []cloudmodel.SecurityGroupReq{
			{Name: "mig-sg-01", VNetId: "mig-vnet-01"},
		},
		TargetInfra: cloudmodel.InfraReq{
			Name: "infra101",
			NodeGroups: []cloudmodel.CreateNodeGroupReq{
				{
					Name:             "ng-influxdb-back",
					VNetId:           "mig-vnet-01",
					SubnetId:         "mig-subnet-01",
					SshKeyId:         "mig-sshkey-01",
					SecurityGroupIds: []string{"mig-sg-01"},
				},
			},
		},
	}
}

func TestApplyNameSeedDoesNotMutateInput(t *testing.T) {
	orig := sampleRecommendedInfra()

	// Apply seed once as in pre-flight validation, then again as in migration.
	first := ApplyNameSeed(orig, "my01")
	second := ApplyNameSeed(orig, "my01")

	// Verify original base names remain intact.
	assert.Equal(t, "mig-vnet-01", orig.TargetVNet.Name)
	assert.Equal(t, "mig-subnet-01", orig.TargetVNet.SubnetInfoList[0].Name)
	assert.Equal(t, "mig-sshkey-01", orig.TargetSshKey.Name)
	assert.Equal(t, "mig-sg-01", orig.TargetSecurityGroupList[0].Name)
	assert.Equal(t, "mig-vnet-01", orig.TargetSecurityGroupList[0].VNetId)
	assert.Equal(t, "infra101", orig.TargetInfra.Name)
	assert.Equal(t, "ng-influxdb-back", orig.TargetInfra.NodeGroups[0].Name)
	assert.Equal(t, "mig-vnet-01", orig.TargetInfra.NodeGroups[0].VNetId)
	assert.Equal(t, "mig-subnet-01", orig.TargetInfra.NodeGroups[0].SubnetId)
	assert.Equal(t, "mig-sshkey-01", orig.TargetInfra.NodeGroups[0].SshKeyId)
	assert.Equal(t, []string{"mig-sg-01"}, orig.TargetInfra.NodeGroups[0].SecurityGroupIds)

	// Verify both seeded copies match and pass composed name validation.
	assert.Equal(t, first, second)
	ok, detail := ValidateComposedNames(second)
	require.True(t, ok, "expected valid composed names, got: %s", detail)
}

func TestPropagateNameChangeDoesNotMutateInput(t *testing.T) {
	orig := sampleRecommendedInfra()

	updated := PropagateNameChange(orig, ResourceTypeVNet, "mig-vnet-01", "custom-vnet")
	updated = PropagateNameChange(updated, ResourceTypeSecurityGroup, "mig-sg-01", "custom-sg")

	// Verify original model slices were not mutated.
	assert.Equal(t, "mig-vnet-01", orig.TargetSecurityGroupList[0].VNetId)
	assert.Equal(t, "mig-sg-01", orig.TargetSecurityGroupList[0].Name)
	assert.Equal(t, "mig-vnet-01", orig.TargetInfra.NodeGroups[0].VNetId)
	assert.Equal(t, []string{"mig-sg-01"}, orig.TargetInfra.NodeGroups[0].SecurityGroupIds)

	// Verify updated model passes referential integrity.
	ok, detail := ValidateReferentialIntegrity(updated)
	require.True(t, ok, "expected valid referential integrity, got: %s", detail)
}

func TestApplyNameSeedConcurrentSafety(t *testing.T) {
	shared := sampleRecommendedInfra()
	shared.TargetK8sCluster = cloudmodel.K8sClusterReq{
		Name:             "k8s-cluster-01",
		VNetId:           "mig-vnet-01",
		SubnetIds:        []string{"mig-subnet-01"},
		SecurityGroupIds: []string{"mig-sg-01"},
		K8sNodeGroupList: []cloudmodel.K8sNodeGroupReq{
			{Name: "workers1", SshKeyId: "mig-sshkey-01"},
		},
	}

	const workers = 32
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(idx int) {
			defer wg.Done()
			seed := fmt.Sprintf("user%02d", idx)
			seeded := ApplyNameSeed(shared, seed)
			ok, detail := ValidateComposedNames(seeded)
			assert.True(t, ok, detail)
			assert.Equal(t, seed+"-infra101", seeded.TargetInfra.Name)
			assert.Equal(t, seed+"-ng-influxdb-back", seeded.TargetInfra.NodeGroups[0].Name)
			assert.Equal(t, seed+"-k8s-cluster-01", seeded.TargetK8sCluster.Name)

			renamed := PropagateNameChange(shared, ResourceTypeVNet, "mig-vnet-01", seed+"-vnet")
			ok, detail = ValidateReferentialIntegrity(renamed)
			assert.True(t, ok, detail)
		}(i)
	}
	wg.Wait()

	// Ensure shared model remains untouched after concurrent calls.
	assert.Equal(t, "mig-vnet-01", shared.TargetVNet.Name)
	assert.Equal(t, "ng-influxdb-back", shared.TargetInfra.NodeGroups[0].Name)
	assert.Equal(t, "mig-sshkey-01", shared.TargetK8sCluster.K8sNodeGroupList[0].SshKeyId)
}

func TestDeepCopyRecommendedInfra(t *testing.T) {
	checkedAt := time.Date(2026, 10, 30, 12, 0, 0, 0, time.UTC)
	rules := []cloudmodel.FirewallRuleReq{
		{Ports: "22", Protocol: "TCP", Direction: "inbound", CIDR: "0.0.0.0/0"},
	}
	orig := sampleRecommendedInfra()
	orig.TargetInfra.Label = map[string]string{"env": "prod"}
	orig.TargetInfra.NodeGroups[0].Label = map[string]string{"role": "db"}
	orig.TargetInfra.NodeGroups[0].SubnetIds = []string{"mig-subnet-01"}
	orig.TargetInfra.NodeGroups[0].DataDiskIds = []string{"disk-01"}
	orig.TargetInfra.PostCommands = []cloudmodel.PostCommandReq{
		{InfraCmdReq: cloudmodel.InfraCmdReq{Command: []string{"echo ok"}}},
	}
	orig.TargetSecurityGroupList[0].FirewallRules = &rules
	orig.TargetSpecList = []cloudmodel.SpecInfo{
		{
			Id:                    "spec-01",
			AssociatedObjectList:  []string{"obj-01"},
			Details:               []cloudmodel.KeyValue{{Key: "k", Value: "v"}},
			AvailableZones:        []string{"zone-a"},
			AvailabilityCheckedAt: &checkedAt,
		},
	}
	orig.TargetOsImageList = []cloudmodel.ImageInfo{
		{
			Id:             "img-01",
			RegionList:     []string{"ap-northeast-2"},
			Details:        []cloudmodel.KeyValue{{Key: "os", Value: "ubuntu"}},
			CommandHistory: []cloudmodel.ImageSourceCommandHistory{{Index: 1, CommandExecuted: "uname -a"}},
		},
	}
	orig.TargetNlbList = []cloudmodel.NlbReq{
		{Type: "PUBLIC", TargetGroup: cloudmodel.NlbTargetGroupReq{NodeGroupId: "ng-influxdb-back"}},
	}
	orig.TargetK8sCluster = cloudmodel.K8sClusterReq{
		Name:             "k8s-01",
		SubnetIds:        []string{"mig-subnet-01"},
		SecurityGroupIds: []string{"mig-sg-01"},
		Label:            map[string]string{"cluster": "main"},
		K8sNodeGroupList: []cloudmodel.K8sNodeGroupReq{
			{Name: "workers1", Label: map[string]string{"tier": "worker"}},
		},
		Storage: &cloudmodel.K8sClusterStorageReq{StorageTypes: []string{"ssd-general"}},
	}

	copied := DeepCopyRecommendedInfra(orig)
	assert.Equal(t, orig, copied)

	// Mutate nested slices, maps, and pointers in copied and verify orig stays untouched.
	copied.TargetInfra.Label["env"] = "dev"
	copied.TargetInfra.NodeGroups[0].Label["role"] = "web"
	copied.TargetInfra.NodeGroups[0].SubnetIds[0] = "mutated"
	copied.TargetInfra.NodeGroups[0].DataDiskIds[0] = "mutated"
	copied.TargetInfra.PostCommands[0].Command[0] = "mutated"
	(*copied.TargetSecurityGroupList[0].FirewallRules)[0].Ports = "8080"
	copied.TargetSpecList[0].AvailableZones[0] = "zone-b"
	*copied.TargetSpecList[0].AvailabilityCheckedAt = checkedAt.Add(time.Hour)
	copied.TargetOsImageList[0].RegionList[0] = "us-east-1"
	copied.TargetOsImageList[0].CommandHistory[0].CommandExecuted = "whoami"
	copied.TargetNlbList[0].Type = "INTERNAL"
	copied.TargetK8sCluster.Label["cluster"] = "mutated"
	copied.TargetK8sCluster.K8sNodeGroupList[0].Label["tier"] = "mutated"
	copied.TargetK8sCluster.Storage.StorageTypes[0] = "hdd"

	assert.Equal(t, "prod", orig.TargetInfra.Label["env"])
	assert.Equal(t, "db", orig.TargetInfra.NodeGroups[0].Label["role"])
	assert.Equal(t, "mig-subnet-01", orig.TargetInfra.NodeGroups[0].SubnetIds[0])
	assert.Equal(t, "disk-01", orig.TargetInfra.NodeGroups[0].DataDiskIds[0])
	assert.Equal(t, "echo ok", orig.TargetInfra.PostCommands[0].Command[0])
	assert.Equal(t, "22", (*orig.TargetSecurityGroupList[0].FirewallRules)[0].Ports)
	assert.Equal(t, "zone-a", orig.TargetSpecList[0].AvailableZones[0])
	assert.Equal(t, checkedAt, *orig.TargetSpecList[0].AvailabilityCheckedAt)
	assert.Equal(t, "ap-northeast-2", orig.TargetOsImageList[0].RegionList[0])
	assert.Equal(t, "uname -a", orig.TargetOsImageList[0].CommandHistory[0].CommandExecuted)
	assert.Equal(t, "PUBLIC", orig.TargetNlbList[0].Type)
	assert.Equal(t, "main", orig.TargetK8sCluster.Label["cluster"])
	assert.Equal(t, "worker", orig.TargetK8sCluster.K8sNodeGroupList[0].Label["tier"])
	assert.Equal(t, "ssd-general", orig.TargetK8sCluster.Storage.StorageTypes[0])

	// Empty seed should also return an isolated deep copy.
	emptySeedCopy := ApplyNameSeed(orig, "")
	emptySeedCopy.TargetVNet.SubnetInfoList[0].Name = "mutated-subnet"
	assert.Equal(t, "mig-subnet-01", orig.TargetVNet.SubnetInfoList[0].Name)
}

func TestComposeNameAndEmptyFieldsPreserved(t *testing.T) {
	assert.Equal(t, "", ComposeName("", "blue"))
	assert.Equal(t, "infra101", ComposeName("infra101", ""))
	assert.Equal(t, "blue-infra101", ComposeName("infra101", "blue"))

	// K8s recommendation model leaves TargetInfra.Name, cluster VNetId, SG VNetId, and NG SshKeyId empty.
	k8sModel := cloudmodel.RecommendedInfra{
		TargetVNet: cloudmodel.VNetReq{
			Name:           "k8s-vpc",
			SubnetInfoList: []cloudmodel.SubnetReq{{Name: "k8s-subnet-01"}},
		},
		TargetSshKey: cloudmodel.SshKeyReq{Name: "k8s-sshkey"},
		TargetSecurityGroupList: []cloudmodel.SecurityGroupReq{
			{Name: "k8s-sg-01", VNetId: ""},
		},
		TargetK8sCluster: cloudmodel.K8sClusterReq{
			Name:    "on-prem-k8s-cluster",
			VNetId:  "",
			K8sNodeGroupList: []cloudmodel.K8sNodeGroupReq{
				{Name: "workers1", SshKeyId: ""},
			},
		},
	}

	seeded := ApplyNameSeed(k8sModel, "blue")
	assert.Equal(t, "blue-k8s-vpc", seeded.TargetVNet.Name)
	assert.Equal(t, "blue-k8s-subnet-01", seeded.TargetVNet.SubnetInfoList[0].Name)
	assert.Equal(t, "blue-k8s-sshkey", seeded.TargetSshKey.Name)
	assert.Equal(t, "blue-k8s-sg-01", seeded.TargetSecurityGroupList[0].Name)
	assert.Equal(t, "", seeded.TargetSecurityGroupList[0].VNetId)
	assert.Equal(t, "", seeded.TargetInfra.Name)
	assert.Equal(t, "blue-on-prem-k8s-cluster", seeded.TargetK8sCluster.Name)
	assert.Equal(t, "", seeded.TargetK8sCluster.VNetId)
	assert.Equal(t, "", seeded.TargetK8sCluster.K8sNodeGroupList[0].SshKeyId)
}

func TestOptionalReferencesSeededAndPropagated(t *testing.T) {
	infra := sampleRecommendedInfra()
	infra.TargetVNet.SubnetInfoList = append(infra.TargetVNet.SubnetInfoList, cloudmodel.SubnetReq{Name: "mig-subnet-02"})
	infra.TargetInfra.NodeGroups[0].SubnetIds = []string{"mig-subnet-01", "mig-subnet-02"}
	infra.TargetInfra.PostCommands = []cloudmodel.PostCommandReq{
		{NodeGroupId: "ng-influxdb-back", InfraCmdReq: cloudmodel.InfraCmdReq{Command: []string{"echo targeted"}}},
		{NodeGroupId: "", InfraCmdReq: cloudmodel.InfraCmdReq{Command: []string{"echo all"}}},
	}
	infra.TargetNlbList = []cloudmodel.NlbReq{
		{Type: "PUBLIC", TargetGroup: cloudmodel.NlbTargetGroupReq{NodeGroupId: "ng-influxdb-back"}},
	}

	seeded := ApplyNameSeed(infra, "prod")
	assert.Equal(t, []string{"prod-mig-subnet-01", "prod-mig-subnet-02"}, seeded.TargetInfra.NodeGroups[0].SubnetIds)
	assert.Equal(t, "prod-ng-influxdb-back", seeded.TargetInfra.PostCommands[0].NodeGroupId)
	assert.Equal(t, "", seeded.TargetInfra.PostCommands[1].NodeGroupId)
	assert.Equal(t, "prod-ng-influxdb-back", seeded.TargetNlbList[0].TargetGroup.NodeGroupId)
	ok, detail := ValidateComposedNames(seeded)
	require.True(t, ok, detail)

	renamed := PropagateNameChange(infra, ResourceTypeSubnet, "mig-subnet-02", "custom-subnet-02")
	assert.Equal(t, []string{"mig-subnet-01", "custom-subnet-02"}, renamed.TargetInfra.NodeGroups[0].SubnetIds)
	ok, detail = ValidateReferentialIntegrity(renamed)
	require.True(t, ok, detail)

	// Broken optional references must be rejected by referential integrity check.
	brokenNlb := DeepCopyRecommendedInfra(infra)
	brokenNlb.TargetNlbList[0].TargetGroup.NodeGroupId = "missing-ng"
	ok, _ = ValidateReferentialIntegrity(brokenNlb)
	assert.False(t, ok)

	brokenPostCmd := DeepCopyRecommendedInfra(infra)
	brokenPostCmd.TargetInfra.PostCommands[0].NodeGroupId = "missing-ng"
	ok, _ = ValidateReferentialIntegrity(brokenPostCmd)
	assert.False(t, ok)
}



