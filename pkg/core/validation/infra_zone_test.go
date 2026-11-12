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

package validation

import (
	"context"
	"net/http"
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	fixtures["/ns/system/resources/spec/aws+g6.xlarge"] = fixture{
		status: http.StatusOK,
		body: tbmodel.SpecInfo{
			Id:                 "aws+g6.xlarge",
			CspSpecName:        "g6.xlarge",
			Architecture:       "x86_64",
			AvailabilityStatus: tbmodel.AvailabilityAvailable,
			AvailableZones:     []string{"ap-northeast-2a", "ap-northeast-2c"},
			UnavailableZones:   []string{"ap-northeast-2b", "ap-northeast-2d"},
		},
	}
	fixtures["/ns/system/resources/spec/aws+t3.medium"] = fixture{
		status: http.StatusOK,
		body: tbmodel.SpecInfo{
			Id:                 "aws+t3.medium",
			CspSpecName:        "t3.medium",
			Architecture:       "x86_64",
			AvailabilityStatus: tbmodel.AvailabilityAvailable,
			AvailableZones:     []string{"ap-northeast-2a", "ap-northeast-2b", "ap-northeast-2c", "ap-northeast-2d"},
		},
	}
	fixtures["/ns/system/resources/spec/aws+only-2a"] = fixture{
		status: http.StatusOK,
		body: tbmodel.SpecInfo{
			Id:                 "aws+only-2a",
			CspSpecName:        "g5.xlarge",
			Architecture:       "x86_64",
			AvailabilityStatus: tbmodel.AvailabilityAvailable,
			AvailableZones:     []string{"ap-northeast-2a"},
		},
	}
	fixtures["/ns/system/resources/spec/aws+only-2c"] = fixture{
		status: http.StatusOK,
		body: tbmodel.SpecInfo{
			Id:                 "aws+only-2c",
			CspSpecName:        "g6e.xlarge",
			Architecture:       "x86_64",
			AvailabilityStatus: tbmodel.AvailabilityAvailable,
			AvailableZones:     []string{"ap-northeast-2c"},
		},
	}
	fixtures["/ns/system/resources/spec/aws+soldout"] = fixture{
		status: http.StatusOK,
		body: tbmodel.SpecInfo{
			Id:                 "aws+soldout",
			CspSpecName:        "p4d.24xlarge",
			Architecture:       "x86_64",
			AvailabilityStatus: tbmodel.AvailabilitySoldOut,
		},
	}
	fixtures["/ns/system/resources/image/aws+ami-ubuntu"] = fixture{
		status: http.StatusOK,
		body: tbmodel.ImageInfo{
			Id:             "aws+ami-ubuntu",
			CspImageName:   "ami-ubuntu",
			OSArchitecture: tbmodel.X86_64,
		},
	}
	fixtures["/ns/" + testNsId + "/resources/vNet/vnet-aws-2b"] = fixture{
		status: http.StatusOK,
		body: tbmodel.VNetInfo{
			Id:             "vnet-aws-2b",
			Name:           "vnet-aws-2b",
			ConnectionName: "aws-ap-northeast-2",
			SubnetInfoList: []tbmodel.SubnetInfo{{Name: "subnet-a", Zone: "ap-northeast-2b"}},
		},
	}
	fixtures["/ns/" + testNsId + "/resources/sshKey/sshkey-aws"] = fixture{
		status: http.StatusOK,
		body: tbmodel.SshKeyInfo{
			Id: "sshkey-aws", Name: "sshkey-aws", ConnectionName: "aws-ap-northeast-2",
		},
	}
	fixtures["/ns/" + testNsId + "/resources/securityGroup/sg-aws"] = fixture{
		status: http.StatusOK,
		body: tbmodel.SecurityGroupInfo{
			Id: "sg-aws", Name: "sg-aws", ConnectionName: "aws-ap-northeast-2",
		},
	}
	fixtures["/ns/system/resources/spec/azure+standard-d4s-v5"] = fixture{
		status: http.StatusOK,
		body: tbmodel.SpecInfo{
			Id:           "azure+standard-d4s-v5",
			CspSpecName:  "Standard_D4s_v5",
			Architecture: "x86_64",
		},
	}
	fixtures["/ns/system/resources/image/azure+ubuntu-2204-gen2"] = fixture{
		status: http.StatusOK,
		body: tbmodel.ImageInfo{
			Id:             "azure+ubuntu-2204-gen2",
			CspImageName:   "ubuntu-2204-gen2",
			OSArchitecture: tbmodel.X86_64,
		},
	}
}

func awsZoneTarget(specId, subnetZone string) *cloudmodel.RecommendedInfra {
	return &cloudmodel.RecommendedInfra{
		TargetCloud: cloudmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"},
		TargetInfra: cloudmodel.InfraReq{
			Name: "infra-aws-new",
			NodeGroups: []cloudmodel.CreateNodeGroupReq{
				{
					Name:             "ng-gpu",
					ConnectionName:   "aws-ap-northeast-2",
					SpecId:           specId,
					ImageId:          "ami-ubuntu",
					VNetId:           "vnet-aws-new",
					SubnetId:         "subnet-a",
					SecurityGroupIds: []string{"sg-aws-new"},
					SshKeyId:         "sshkey-aws-new",
				},
			},
		},
		TargetVNet: cloudmodel.VNetReq{
			Name:           "vnet-aws-new",
			ConnectionName: "aws-ap-northeast-2",
			CidrBlock:      "10.0.0.0/16",
			SubnetInfoList: []cloudmodel.SubnetReq{
				{Name: "subnet-a", IPv4_CIDR: "10.0.1.0/24", Zone: subnetZone},
			},
		},
		TargetSshKey: cloudmodel.SshKeyReq{
			Name:           "sshkey-aws-new",
			ConnectionName: "aws-ap-northeast-2",
		},
		TargetSecurityGroupList: []cloudmodel.SecurityGroupReq{
			{Name: "sg-aws-new", ConnectionName: "aws-ap-northeast-2", VNetId: "vnet-aws-new"},
		},
	}
}

func TestValidateTargetInfra_SubnetZoneAlignment(t *testing.T) {
	ctx := context.Background()

	t.Run("zonal CSP with valid AvailableZone passes", func(t *testing.T) {
		res := ValidateTargetInfra(ctx, testNsId, awsZoneTarget("aws+g6.xlarge", "ap-northeast-2a"), false)
		assert.True(t, res.Valid)
		assert.Empty(t, res.Issues)
	})

	t.Run("zonal CSP with unavailable zone fails with CodeSubnetZoneUnavailable", func(t *testing.T) {
		res := ValidateTargetInfra(ctx, testNsId, awsZoneTarget("aws+g6.xlarge", "ap-northeast-2b"), false)
		assert.False(t, res.Valid)
		require.Len(t, res.Issues, 1)
		assert.Equal(t, CodeSubnetZoneUnavailable, res.Issues[0].Code)
		assert.Equal(t, SeverityError, res.Issues[0].Severity)
		assert.Equal(t, "targetVNet.subnetInfoList[0].zone", res.Issues[0].Path)
	})

	t.Run("zonal CSP with empty zone fails when spec has UnavailableZones", func(t *testing.T) {
		res := ValidateTargetInfra(ctx, testNsId, awsZoneTarget("aws+g6.xlarge", ""), false)
		assert.False(t, res.Valid)
		require.Len(t, res.Issues, 1)
		assert.Equal(t, CodeSubnetZoneUnavailable, res.Issues[0].Code)
		assert.Equal(t, "targetVNet.subnetInfoList[0].zone", res.Issues[0].Path)
	})

	t.Run("zonal CSP with empty zone passes when spec is offered in all zones", func(t *testing.T) {
		res := ValidateTargetInfra(ctx, testNsId, awsZoneTarget("aws+t3.medium", ""), false)
		assert.True(t, res.Valid)
		assert.Empty(t, res.Issues)
	})

	t.Run("zonal CSP with SoldOut spec fails with CodeSubnetZoneUnavailable", func(t *testing.T) {
		res := ValidateTargetInfra(ctx, testNsId, awsZoneTarget("aws+soldout", "ap-northeast-2a"), false)
		assert.False(t, res.Valid)
		require.Len(t, res.Issues, 1)
		assert.Equal(t, CodeSubnetZoneUnavailable, res.Issues[0].Code)
		assert.Equal(t, "targetInfra.nodeGroups[0].specId", res.Issues[0].Path)
	})

	t.Run("zonal CSP with disjoint multi-NodeGroup zones on unpinned subnet fails", func(t *testing.T) {
		target := awsZoneTarget("aws+only-2a", "")
		ng2 := target.TargetInfra.NodeGroups[0]
		ng2.Name = "ng-gpu-2"
		ng2.SpecId = "aws+only-2c"
		target.TargetInfra.NodeGroups = append(target.TargetInfra.NodeGroups, ng2)

		res := ValidateTargetInfra(ctx, testNsId, target, false)
		assert.False(t, res.Valid)
		require.Len(t, res.Issues, 1)
		assert.Equal(t, CodeSubnetZoneUnavailable, res.Issues[0].Code)
		assert.Equal(t, "targetVNet.subnetInfoList[0].zone", res.Issues[0].Path)
	})

	t.Run("useExisting mode detects unavailable zone on existing Tumblebug VNet", func(t *testing.T) {
		target := awsZoneTarget("aws+g6.xlarge", "")
		target.TargetInfra.Name = "infra-aws-existing"
		target.TargetVNet.Name = "vnet-aws-2b"
		target.TargetInfra.NodeGroups[0].VNetId = "vnet-aws-2b"
		target.TargetSecurityGroupList[0].VNetId = "vnet-aws-2b"
		target.TargetSecurityGroupList[0].Name = "sg-aws"
		target.TargetInfra.NodeGroups[0].SecurityGroupIds = []string{"sg-aws"}
		target.TargetSshKey.Name = "sshkey-aws"
		target.TargetInfra.NodeGroups[0].SshKeyId = "sshkey-aws"

		res := ValidateTargetInfra(ctx, testNsId, target, true)
		assert.False(t, res.Valid)
		require.Len(t, res.Issues, 1)
		assert.Equal(t, CodeSubnetZoneUnavailable, res.Issues[0].Code)
		assert.Equal(t, "targetInfra.nodeGroups[0].subnetId", res.Issues[0].Path)
	})

	t.Run("TargetSpecList availability metadata enriches unenriched ReadVmSpec response", func(t *testing.T) {
		fixtures["/ns/system/resources/spec/aws+unenriched"] = fixture{
			status: http.StatusOK,
			body: tbmodel.SpecInfo{
				Id:           "aws+unenriched",
				CspSpecName:  "g5.2xlarge",
				Architecture: "x86_64",
			},
		}
		target := awsZoneTarget("aws+unenriched", "ap-northeast-2d")
		target.TargetSpecList = []cloudmodel.SpecInfo{
			{
				Id:                 "aws+unenriched",
				CspSpecName:        "g5.2xlarge",
				AvailabilityStatus: cloudmodel.AvailabilityAvailable,
				AvailableZones:     []string{"ap-northeast-2a", "ap-northeast-2c"},
				UnavailableZones:   []string{"ap-northeast-2b", "ap-northeast-2d"},
			},
		}

		res := ValidateTargetInfra(ctx, testNsId, target, false)
		assert.False(t, res.Valid)
		require.Len(t, res.Issues, 1)
		assert.Equal(t, CodeSubnetZoneUnavailable, res.Issues[0].Code)
		assert.Equal(t, "targetVNet.subnetInfoList[0].zone", res.Issues[0].Path)
	})

	t.Run("regional CSP with pinned Subnet.Zone returns non-blocking warning", func(t *testing.T) {
		target := &cloudmodel.RecommendedInfra{
			TargetCloud: cloudmodel.CloudProperty{Csp: "azure", Region: "koreacentral"},
			TargetInfra: cloudmodel.InfraReq{
				Name: "infra-azure-new",
				NodeGroups: []cloudmodel.CreateNodeGroupReq{
					{
						Name:             "ng-azure",
						ConnectionName:   "azure-koreacentral",
						SpecId:           "azure+standard-d4s-v5",
						ImageId:          "ubuntu-2204-gen2",
						VNetId:           "vnet-az-new",
						SubnetId:         "subnet-a",
						SecurityGroupIds: []string{"sg-az-new"},
						SshKeyId:         "sshkey-az-new",
					},
				},
			},
			TargetVNet: cloudmodel.VNetReq{
				Name:           "vnet-az-new",
				ConnectionName: "azure-koreacentral",
				CidrBlock:      "10.0.0.0/16",
				SubnetInfoList: []cloudmodel.SubnetReq{
					{Name: "subnet-a", IPv4_CIDR: "10.0.1.0/24", Zone: "koreacentral-1"},
				},
			},
			TargetSshKey: cloudmodel.SshKeyReq{
				Name:           "sshkey-az-new",
				ConnectionName: "azure-koreacentral",
			},
			TargetSecurityGroupList: []cloudmodel.SecurityGroupReq{
				{Name: "sg-az-new", ConnectionName: "azure-koreacentral", VNetId: "vnet-az-new"},
			},
		}

		res := ValidateTargetInfra(ctx, testNsId, target, false)
		assert.True(t, res.Valid)
		require.Len(t, res.Issues, 1)
		assert.Equal(t, CodeRegionalSubnetZonePinned, res.Issues[0].Code)
		assert.Equal(t, SeverityWarning, res.Issues[0].Severity)
		assert.Equal(t, "targetVNet.subnetInfoList[0].zone", res.Issues[0].Path)
	})
}
