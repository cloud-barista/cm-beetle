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
	"fmt"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
)

// NetworkRequirement represents the virtual network and subnets required by NodeGroups.
type NetworkRequirement struct {
	VNetId         string
	SubnetIds      []string
	ConnectionName string
}

// DeriveNetworkRequirements groups and extracts virtual network requirements from NodeGroups.
func DeriveNetworkRequirements(nodeGroups []cloudmodel.CreateNodeGroupReq) []NetworkRequirement {
	vNetSubnets := make(map[string][]string)
	vNetConnection := make(map[string]string)
	var orderedVNets []string
	seenVNets := make(map[string]bool)

	for _, ng := range nodeGroups {
		if ng.VNetId == "" {
			continue
		}
		if !seenVNets[ng.VNetId] {
			seenVNets[ng.VNetId] = true
			orderedVNets = append(orderedVNets, ng.VNetId)
		}
		if ng.ConnectionName != "" && vNetConnection[ng.VNetId] == "" {
			vNetConnection[ng.VNetId] = ng.ConnectionName
		}
		if ng.SubnetId != "" {
			exists := false
			for _, sub := range vNetSubnets[ng.VNetId] {
				if sub == ng.SubnetId {
					exists = true
					break
				}
			}
			if !exists {
				vNetSubnets[ng.VNetId] = append(vNetSubnets[ng.VNetId], ng.SubnetId)
			}
		}
	}

	var requirements []NetworkRequirement
	for _, vNetId := range orderedVNets {
		requirements = append(requirements, NetworkRequirement{
			VNetId:         vNetId,
			SubnetIds:      vNetSubnets[vNetId],
			ConnectionName: vNetConnection[vNetId],
		})
	}
	return requirements
}

// CheckNetworkAvailability reports whether the required vNet and subnets already
// exist, or - if not - whether vNetCreationReq carries enough data to create them.
// It performs reads only; no resource is created or modified.
func CheckNetworkAvailability(nsId string, netRequirement NetworkRequirement, vNetCreationReq cloudmodel.VNetReq) (needsCreate bool, issue *ValidationIssue) {
	vNetInfo, err := tbclient.NewSession().ReadVNet(nsId, netRequirement.VNetId)
	vNetExists := err == nil && vNetInfo.Id != ""

	// Existing by ID is not enough: Tumblebug only guarantees the ID is unique,
	// not that it was provisioned under the requested CSP/region. ConnectionName
	// is Tumblebug's connection identifier, which is fixed to one CSP+region+
	// credential, so a mismatch here means reusing it would silently attach
	// resources to the wrong CSP/region.
	if vNetExists && netRequirement.ConnectionName != "" && vNetInfo.ConnectionName != netRequirement.ConnectionName {
		return false, &ValidationIssue{
			Code:     CodeConnectionMismatch,
			Severity: SeverityError,
			Path:     fmt.Sprintf("targetVNet (id: %s)", netRequirement.VNetId),
			Message: fmt.Sprintf("vNet '%s' already exists in namespace '%s', but under connection '%s', not the requested '%s' - reusing it would place resources in a different CSP/region",
				netRequirement.VNetId, nsId, vNetInfo.ConnectionName, netRequirement.ConnectionName),
		}
	}

	allSubnetsExist := true
	if vNetExists {
		existingSubnets := make(map[string]bool)
		for _, sub := range vNetInfo.SubnetInfoList {
			existingSubnets[sub.Name] = true
		}
		for _, subnetId := range netRequirement.SubnetIds {
			if !existingSubnets[subnetId] {
				allSubnetsExist = false
				break
			}
		}
	}

	if vNetExists && allSubnetsExist {
		return false, nil
	}

	if vNetCreationReq.CidrBlock == "" {
		return true, &ValidationIssue{
			Code:     CodeResourceNotAvailable,
			Severity: SeverityError,
			Path:     fmt.Sprintf("targetVNet (id: %s)", netRequirement.VNetId),
			Message: fmt.Sprintf("vNet '%s' (or one of its required subnets) does not exist, "+
				"and no valid VNet creation data (cidrBlock) was provided to create it", netRequirement.VNetId),
		}
	}
	return true, nil
}

// SshKeyRequirement represents the SSH key required by NodeGroups.
type SshKeyRequirement struct {
	SshKeyId       string
	ConnectionName string
}

// DeriveSshKeyRequirements extracts unique SSH key requirements from NodeGroups.
func DeriveSshKeyRequirements(nodeGroups []cloudmodel.CreateNodeGroupReq) []SshKeyRequirement {
	var requirements []SshKeyRequirement
	seenSshKeys := make(map[string]bool)
	for _, ng := range nodeGroups {
		if ng.SshKeyId == "" || seenSshKeys[ng.SshKeyId] {
			continue
		}
		seenSshKeys[ng.SshKeyId] = true
		requirements = append(requirements, SshKeyRequirement{
			SshKeyId:       ng.SshKeyId,
			ConnectionName: ng.ConnectionName,
		})
	}
	return requirements
}

// CheckSshKeyAvailability reports whether the required SSH key already exists,
// or - if not - whether sshKeyCreationReq carries enough data to create it.
// It performs reads only; no resource is created or modified.
func CheckSshKeyAvailability(nsId string, sshKeyRequirement SshKeyRequirement, sshKeyCreationReq cloudmodel.SshKeyReq) (needsCreate bool, issue *ValidationIssue) {
	sshKeyInfo, err := tbclient.NewSession().ReadSshKey(nsId, sshKeyRequirement.SshKeyId)
	if err == nil && sshKeyInfo.Id != "" {
		// Existing by ID is not enough: verify it was provisioned under the
		// requested CSP/region (Tumblebug's ConnectionName), not just that an
		// SSH key with this ID happens to exist somewhere in the namespace.
		if sshKeyRequirement.ConnectionName != "" && sshKeyInfo.ConnectionName != sshKeyRequirement.ConnectionName {
			return false, &ValidationIssue{
				Code:     CodeConnectionMismatch,
				Severity: SeverityError,
				Path:     fmt.Sprintf("targetSshKey (id: %s)", sshKeyRequirement.SshKeyId),
				Message: fmt.Sprintf("SSH key '%s' already exists in namespace '%s', but under connection '%s', not the requested '%s' - reusing it would place resources in a different CSP/region",
					sshKeyRequirement.SshKeyId, nsId, sshKeyInfo.ConnectionName, sshKeyRequirement.ConnectionName),
			}
		}
		return false, nil
	}

	if sshKeyCreationReq.Name == "" {
		return true, &ValidationIssue{
			Code:     CodeResourceNotAvailable,
			Severity: SeverityError,
			Path:     fmt.Sprintf("targetSshKey (id: %s)", sshKeyRequirement.SshKeyId),
			Message: fmt.Sprintf("SSH key '%s' does not exist, and no valid SSH key creation data was provided to create it",
				sshKeyRequirement.SshKeyId),
		}
	}
	return true, nil
}

// SecurityGroupRequirement represents the security group required by NodeGroups.
type SecurityGroupRequirement struct {
	SecurityGroupId string
	VNetId          string
	ConnectionName  string
}

// DeriveSecurityGroupRequirements extracts unique security group requirements from NodeGroups.
func DeriveSecurityGroupRequirements(nodeGroups []cloudmodel.CreateNodeGroupReq) []SecurityGroupRequirement {
	var requirements []SecurityGroupRequirement
	seenSgs := make(map[string]bool)
	for _, ng := range nodeGroups {
		for _, sgId := range ng.SecurityGroupIds {
			if sgId == "" || seenSgs[sgId] {
				continue
			}
			seenSgs[sgId] = true
			requirements = append(requirements, SecurityGroupRequirement{
				SecurityGroupId: sgId,
				VNetId:          ng.VNetId,
				ConnectionName:  ng.ConnectionName,
			})
		}
	}
	return requirements
}

// CheckSecurityGroupAvailability reports whether the required security group already
// exists, or - if not - whether sgCreationReqList carries enough data (a matching
// entry with ConnectionName and VNetId resolvable) to create it.
// It performs reads only; no resource is created or modified.
func CheckSecurityGroupAvailability(nsId string, sgRequirement SecurityGroupRequirement, sgCreationReqList []cloudmodel.SecurityGroupReq) (needsCreate bool, issue *ValidationIssue) {
	sgInfo, err := tbclient.NewSession().ReadSecurityGroup(nsId, sgRequirement.SecurityGroupId)
	if err == nil && sgInfo.Id != "" {
		// Existing by ID is not enough: verify it was provisioned under the
		// requested CSP/region (Tumblebug's ConnectionName), not just that a
		// security group with this ID happens to exist somewhere in the namespace.
		if sgRequirement.ConnectionName != "" && sgInfo.ConnectionName != sgRequirement.ConnectionName {
			return false, &ValidationIssue{
				Code:     CodeConnectionMismatch,
				Severity: SeverityError,
				Path:     fmt.Sprintf("targetSecurityGroupList (id: %s)", sgRequirement.SecurityGroupId),
				Message: fmt.Sprintf("security group '%s' already exists in namespace '%s', but under connection '%s', not the requested '%s' - reusing it would place resources in a different CSP/region",
					sgRequirement.SecurityGroupId, nsId, sgInfo.ConnectionName, sgRequirement.ConnectionName),
			}
		}
		return false, nil
	}

	var sgCreationReq cloudmodel.SecurityGroupReq
	found := false
	for _, sg := range sgCreationReqList {
		if sg.Name == sgRequirement.SecurityGroupId {
			sgCreationReq = sg
			found = true
			break
		}
	}
	if !found {
		sgCreationReq = cloudmodel.SecurityGroupReq{Name: sgRequirement.SecurityGroupId}
	}

	connectionName := sgCreationReq.ConnectionName
	if connectionName == "" {
		connectionName = sgRequirement.ConnectionName
	}
	vNetId := sgCreationReq.VNetId
	if vNetId == "" {
		vNetId = sgRequirement.VNetId
	}

	if connectionName == "" || vNetId == "" {
		return true, &ValidationIssue{
			Code:     CodeResourceNotAvailable,
			Severity: SeverityError,
			Path:     fmt.Sprintf("targetSecurityGroupList (id: %s)", sgRequirement.SecurityGroupId),
			Message: fmt.Sprintf("security group '%s' does not exist, and required connectionName or vNetId "+
				"is missing to create it", sgRequirement.SecurityGroupId),
		}
	}
	return true, nil
}
