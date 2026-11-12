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

package recommendation

import (
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsZonalAndRegionalSubnetCSP(t *testing.T) {
	zonalCSPs := []string{"aws", "AWS", "alibaba", "tencent", "ibm", "ncp", "ncpvpc", "kt", "ktvpc"}
	for _, provider := range zonalCSPs {
		assert.Truef(t, IsZonalSubnetCSP(provider), "expected %s to be classified as zonal-subnet CSP", provider)
		assert.Falsef(t, IsRegionalSubnetWarningCSP(provider), "expected %s not to trigger regional subnet warning", provider)
	}

	regionalCSPs := []string{"azure", "gcp", "nhn", "openstack", "ktclassic"}
	for _, provider := range regionalCSPs {
		assert.Falsef(t, IsZonalSubnetCSP(provider), "expected %s to be classified as regional-subnet CSP", provider)
	}

	assert.True(t, IsRegionalSubnetWarningCSP("azure"))
	assert.True(t, IsRegionalSubnetWarningCSP("gcp"))
	assert.False(t, IsRegionalSubnetWarningCSP("nhn"))
}

func TestFilterAvailableSpecsAndNoFallback(t *testing.T) {
	mixedSpecs := []cloudmodel.SpecInfo{
		{Id: "aws+g6.xlarge", AvailabilityStatus: cloudmodel.AvailabilityAvailable, AvailableZones: []string{"ap-northeast-2a"}},
		{Id: "aws+g5.xlarge", AvailabilityStatus: cloudmodel.AvailabilityUnsupported},
		{Id: "aws+p4d.24xlarge", AvailabilityStatus: cloudmodel.AvailabilityQuotaExceeded},
		{Id: "aws+g4dn.xlarge", AvailabilityStatus: cloudmodel.AvailabilitySoldOut},
		{Id: "aws+g6e.xlarge", AvailabilityStatus: cloudmodel.AvailabilityUnverified},
		{Id: "aws+broken.xlarge", AvailabilityStatus: cloudmodel.AvailabilityAvailable, UnavailableZones: []string{"ap-northeast-2a", "ap-northeast-2b"}},
	}

	filtered := filterAvailableSpecs("aws", mixedSpecs)
	require.Len(t, filtered, 2)
	assert.Equal(t, "aws+g6.xlarge", filtered[0].Id)
	assert.Equal(t, "aws+g6e.xlarge", filtered[1].Id)

	// Verify No Fallback rule: when every candidate is SoldOut, return an empty slice.
	allSoldOutTb := []tbmodel.SpecInfo{
		{Id: "aws+g6.xlarge", AvailabilityStatus: tbmodel.AvailabilitySoldOut},
		{Id: "aws+g5.xlarge", AvailabilityStatus: tbmodel.AvailabilitySoldOut},
	}
	assert.Empty(t, filterAvailableTbSpecs("aws", allSoldOutTb))
}

func TestAvailabilityPriorityInCpuAndGpuSorting(t *testing.T) {
	cpuSpecs := []cloudmodel.SpecInfo{
		{
			Id: "aws+m6i.large-unverified", CspSpecName: "m6i.large",
			VCPU: 2, MemoryGiB: 8, CostPerHour: 0.08,
			AvailabilityStatus: cloudmodel.AvailabilityUnverified,
		},
		{
			Id: "aws+m6i.large-available", CspSpecName: "m6i.large",
			VCPU: 2, MemoryGiB: 8, CostPerHour: 0.10,
			AvailabilityStatus: cloudmodel.AvailabilityAvailable,
		},
	}
	sortByProximityWithCost(cpuSpecs, 2, 8, "aws", "intel")
	assert.Equal(t, "aws+m6i.large-available", cpuSpecs[0].Id)

	gpuNode := onpremmodel.NodeProperty{
		MachineId: "gpu-01",
		CPU:       onpremmodel.CpuProperty{Cpus: 1, Threads: 4},
		Memory:    onpremmodel.MemoryProperty{TotalSize: 16},
		GPUCards: []onpremmodel.GpuCardProperty{
			{Vendor: "NVIDIA", Model: "NVIDIA L4", MemoryTotalGB: 24},
		},
	}
	gpuSpecs := []cloudmodel.SpecInfo{
		{
			Id: "aws+g6.xlarge-unverified", CspSpecName: "g6.xlarge",
			VCPU: 4, MemoryGiB: 16, AcceleratorCount: 1, AcceleratorMemoryGB: 24, AcceleratorModel: "NVIDIA L4",
			CostPerHour: 0.70, AvailabilityStatus: cloudmodel.AvailabilityUnverified,
		},
		{
			Id: "aws+g6.xlarge-available", CspSpecName: "g6.xlarge",
			VCPU: 4, MemoryGiB: 16, AcceleratorCount: 1, AcceleratorMemoryGB: 24, AcceleratorModel: "NVIDIA L4",
			CostPerHour: 0.80, AvailabilityStatus: cloudmodel.AvailabilityAvailable,
		},
	}
	sortGpuByProximityWithCost(gpuSpecs, gpuNode, "aws")
	assert.Equal(t, "aws+g6.xlarge-available", gpuSpecs[0].Id)
}

func TestAlignCandidateSubnetsWithSpecZones(t *testing.T) {
	// Case 1: Zonal CSP selects deterministic common AvailableZone without mutating shared skeleton slice.
	skeleton := cloudmodel.RecommendedInfra{
		TargetCloud: cloudmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"},
		TargetInfra: cloudmodel.InfraReq{Name: "gpu-infra-01"},
		TargetVNet: cloudmodel.VNetReq{
			Name:           "vnet-01",
			ConnectionName: "aws-ap-northeast-2",
			CidrBlock:      "10.0.0.0/16",
			SubnetInfoList: []cloudmodel.SubnetReq{{Name: "subnet-01", IPv4_CIDR: "10.0.1.0/24"}},
		},
		TargetSshKey: cloudmodel.SshKeyReq{Name: "sshkey-01", ConnectionName: "aws-ap-northeast-2"},
		TargetSecurityGroupList: []cloudmodel.SecurityGroupReq{
			{Name: "sg-01", ConnectionName: "aws-ap-northeast-2", VNetId: "vnet-01"},
		},
	}

	cand1 := skeleton
	cand1.TargetInfra.NodeGroups = []cloudmodel.CreateNodeGroupReq{
		{
			Name: "ng-cpu", ConnectionName: "aws-ap-northeast-2", SpecId: "aws+t3.medium", ImageId: "ami-1",
			VNetId: "vnet-01", SubnetId: "subnet-01", SshKeyId: "sshkey-01", SecurityGroupIds: []string{"sg-01"},
		},
		{
			Name: "ng-gpu", ConnectionName: "aws-ap-northeast-2", SpecId: "aws+g6.xlarge", ImageId: "ami-2",
			VNetId: "vnet-01", SubnetId: "subnet-01", SshKeyId: "sshkey-01", SecurityGroupIds: []string{"sg-01"},
		},
	}
	cand1.TargetSpecList = []cloudmodel.SpecInfo{
		{Id: "aws+t3.medium", AvailableZones: []string{"ap-northeast-2b", "ap-northeast-2c", "ap-northeast-2a"}},
		{Id: "aws+g6.xlarge", AvailableZones: []string{"ap-northeast-2c", "ap-northeast-2a"}, UnavailableZones: []string{"ap-northeast-2b"}},
	}

	alignCandidateSubnetsWithSpecZones(&cand1)
	require.Len(t, cand1.TargetVNet.SubnetInfoList, 1)
	assert.Equal(t, "ap-northeast-2a", cand1.TargetVNet.SubnetInfoList[0].Zone)
	assert.Equal(t, "", skeleton.TargetVNet.SubnetInfoList[0].Zone, "shared skeleton slice must remain unmodified")

	ok, detail := common.ValidateComposedNames(cand1)
	assert.Truef(t, ok, "expected composed names to remain valid: %s", detail)
	seeded := common.ApplyNameSeed(cand1, "demo")
	okSeeded, detailSeeded := common.ValidateComposedNames(seeded)
	assert.Truef(t, okSeeded, "expected seeded names to remain valid: %s", detailSeeded)
	assert.Equal(t, "ap-northeast-2a", seeded.TargetVNet.SubnetInfoList[0].Zone)

	// Case 2: Disjoint AvailableZones across NodeGroups leave Zone unset (No Fallback).
	disjoint := skeleton
	disjoint.TargetInfra.NodeGroups = []cloudmodel.CreateNodeGroupReq{
		{
			Name: "ng-1", ConnectionName: "aws-ap-northeast-2", SpecId: "aws+spec-a", ImageId: "ami-1",
			VNetId: "vnet-01", SubnetId: "subnet-01", SshKeyId: "sshkey-01", SecurityGroupIds: []string{"sg-01"},
		},
		{
			Name: "ng-2", ConnectionName: "aws-ap-northeast-2", SpecId: "aws+spec-c", ImageId: "ami-2",
			VNetId: "vnet-01", SubnetId: "subnet-01", SshKeyId: "sshkey-01", SecurityGroupIds: []string{"sg-01"},
		},
	}
	disjoint.TargetSpecList = []cloudmodel.SpecInfo{
		{Id: "aws+spec-a", AvailableZones: []string{"ap-northeast-2a"}},
		{Id: "aws+spec-c", AvailableZones: []string{"ap-northeast-2c"}},
	}
	alignCandidateSubnetsWithSpecZones(&disjoint)
	assert.Equal(t, "", disjoint.TargetVNet.SubnetInfoList[0].Zone)

	// Case 3: Regional CSPs (azure, gcp, nhn) always preserve empty Subnet.Zone.
	for _, regionalCSP := range []string{"azure", "gcp", "nhn"} {
		reg := skeleton
		reg.TargetCloud.Csp = regionalCSP
		reg.TargetVNet.SubnetInfoList = []cloudmodel.SubnetReq{
			{Name: "subnet-01", IPv4_CIDR: "10.0.1.0/24", Zone: "accidental-zone-1"},
		}
		reg.TargetInfra.NodeGroups = []cloudmodel.CreateNodeGroupReq{
			{Name: "ng-1", ConnectionName: regionalCSP + "-region1", SpecId: regionalCSP + "+spec-1"},
		}
		reg.TargetSpecList = []cloudmodel.SpecInfo{
			{Id: regionalCSP + "+spec-1", AvailableZones: []string{"zone-1", "zone-2"}},
		}
		alignCandidateSubnetsWithSpecZones(&reg)
		assert.Equal(t, "", reg.TargetVNet.SubnetInfoList[0].Zone)
	}
}
