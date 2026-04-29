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
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
)

// ComposeName combines a base name and a seed.
// If seed is empty, it returns the base name.
func ComposeName(baseName, seed string) string {
	if seed == "" {
		return baseName
	}
	// If baseName already starts with the seed, return it as is (idempotency)
	if strings.HasPrefix(baseName, seed+"-") {
		return baseName
	}
	return fmt.Sprintf("%s-%s", seed, baseName)
}

// ApplyNameSeed applies the late binding naming strategy.
// It returns a NEW RecommendedVmInfra object with prefixed names,
// while keeping the original seed in the NameSeed field.
func ApplyNameSeed(infra cloudmodel.RecommendedInfra) cloudmodel.RecommendedInfra {
	if infra.NameSeed == "" {
		return infra
	}

	seed := infra.NameSeed
	result := infra // Copy by value

	// 1. Update VNet and Subnets
	result.TargetVNet.Name = ComposeName(infra.TargetVNet.Name, seed)
	for i, subnet := range infra.TargetVNet.SubnetInfoList {
		result.TargetVNet.SubnetInfoList[i].Name = ComposeName(subnet.Name, seed)
	}

	// 2. Update SSH Key
	result.TargetSshKey.Name = ComposeName(infra.TargetSshKey.Name, seed)

	// 3. Update Security Groups
	for i, sg := range infra.TargetSecurityGroupList {
		result.TargetSecurityGroupList[i].Name = ComposeName(sg.Name, seed)
		// Update reference to VNetId if it's a relative name
		if sg.VNetId == infra.TargetVNet.Name {
			result.TargetSecurityGroupList[i].VNetId = result.TargetVNet.Name
		}
	}

	// 4. Update Infra and NodeGroups
	result.TargetInfra.Name = ComposeName(infra.TargetInfra.Name, seed)
	for i, ng := range infra.TargetInfra.NodeGroups {
		result.TargetInfra.NodeGroups[i].Name = ComposeName(ng.Name, seed)
		result.TargetInfra.NodeGroups[i].VNetId = ComposeName(ng.VNetId, seed)
		result.TargetInfra.NodeGroups[i].SubnetId = ComposeName(ng.SubnetId, seed)
		result.TargetInfra.NodeGroups[i].SshKeyId = ComposeName(ng.SshKeyId, seed)

		// Update SecurityGroupIds
		newSgIds := make([]string, len(ng.SecurityGroupIds))
		for j, sgId := range ng.SecurityGroupIds {
			newSgIds[j] = ComposeName(sgId, seed)
		}
		result.TargetInfra.NodeGroups[i].SecurityGroupIds = newSgIds
	}

	return result
}

// ResourceType constants follow the cb-tumblebug naming convention.
// Ref: cb-tumblebug/src/core/model/common.go
const (
	ResourceTypeVNet          = "vNet"
	ResourceTypeSubnet        = "subnet"
	ResourceTypeSshKey        = "sshKey"
	ResourceTypeSecurityGroup = "securityGroup"
	ResourceTypeInfra         = "infra"
)

// PropagateNameChange updates a specific resource's name in the model and propagates
// the change to all child/dependent resources that reference it.
// (e.g., If VNet is renamed, all SecurityGroup.VNetId and SubGroup.VNetId are updated)
//
// Parameters:
//   - infra: The recommendation model to update (will be copied, not mutated)
//   - resourceType: One of "vNet", "subnet", "sshKey", "securityGroup", "mci"
//   - oldName: The current name of the resource to rename
//   - newName: The new name to assign to the resource
func PropagateNameChange(infra cloudmodel.RecommendedInfra, resourceType, oldName, newName string) cloudmodel.RecommendedInfra {
	result := infra // Copy by value

	switch resourceType {
	case ResourceTypeVNet:
		// 1. Rename the VNet itself
		if result.TargetVNet.Name == oldName {
			result.TargetVNet.Name = newName
		}
		// 2. Propagate to SecurityGroups
		for i := range result.TargetSecurityGroupList {
			if result.TargetSecurityGroupList[i].VNetId == oldName {
				result.TargetSecurityGroupList[i].VNetId = newName
			}
		}
		// 3. Propagate to NodeGroups
		for i := range result.TargetInfra.NodeGroups {
			if result.TargetInfra.NodeGroups[i].VNetId == oldName {
				result.TargetInfra.NodeGroups[i].VNetId = newName
			}
		}

	case ResourceTypeSubnet:
		// 1. Rename the Subnet itself
		for i := range result.TargetVNet.SubnetInfoList {
			if result.TargetVNet.SubnetInfoList[i].Name == oldName {
				result.TargetVNet.SubnetInfoList[i].Name = newName
			}
		}
		// 2. Propagate to NodeGroups
		for i := range result.TargetInfra.NodeGroups {
			if result.TargetInfra.NodeGroups[i].SubnetId == oldName {
				result.TargetInfra.NodeGroups[i].SubnetId = newName
			}
		}

	case ResourceTypeSshKey:
		// 1. Rename the SSH Key itself
		if result.TargetSshKey.Name == oldName {
			result.TargetSshKey.Name = newName
		}
		// 2. Propagate to NodeGroups
		for i := range result.TargetInfra.NodeGroups {
			if result.TargetInfra.NodeGroups[i].SshKeyId == oldName {
				result.TargetInfra.NodeGroups[i].SshKeyId = newName
			}
		}

	case ResourceTypeSecurityGroup:
		// 1. Rename the SecurityGroup itself
		for i := range result.TargetSecurityGroupList {
			if result.TargetSecurityGroupList[i].Name == oldName {
				result.TargetSecurityGroupList[i].Name = newName
			}
		}
		// 2. Propagate to SubGroup SecurityGroupIds
		for i, ng := range result.TargetInfra.NodeGroups {
			for j, sgId := range ng.SecurityGroupIds {
				if sgId == oldName {
					result.TargetInfra.NodeGroups[i].SecurityGroupIds[j] = newName
				}
			}
		}

	case ResourceTypeInfra:
		// Rename the Infra itself (no child dependencies on Infra name)
		if result.TargetInfra.Name == oldName {
			result.TargetInfra.Name = newName
		}
	}

	return result
}

// IsValidName checks if the name complies with general cloud resource naming rules.
// Rules: 3-63 characters, alphanumeric and hyphens, starts with alphanumeric.
func IsValidName(name string) (bool, string) {
	if len(name) < 3 || len(name) > 63 {
		return false, "name length must be between 3 and 63 characters"
	}
	// Simplified regex check without importing regexp if possible,
	// but for robustness, let's just use basic logic for now.
	for i, r := range name {
		if i == 0 && !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false, "name must start with an alphanumeric character"
		}
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-') {
			return false, "name contains invalid characters (only alphanumeric and hyphens allowed)"
		}
	}
	return true, ""
}

// ValidateComposedNames checks all resource names and referential integrity.
func ValidateComposedNames(infra cloudmodel.RecommendedInfra) (bool, string) {
	seed := infra.NameSeed

	// 1. Validate name formats
	// Check VNet
	if ok, detail := IsValidName(ComposeName(infra.TargetVNet.Name, seed)); !ok {
		return false, fmt.Sprintf("VNet name [%s]: %s", infra.TargetVNet.Name, detail)
	}

	// Check Subnets
	for _, subnet := range infra.TargetVNet.SubnetInfoList {
		if ok, detail := IsValidName(ComposeName(subnet.Name, seed)); !ok {
			return false, fmt.Sprintf("Subnet name [%s]: %s", subnet.Name, detail)
		}
	}

	// Check SSH Key
	if ok, detail := IsValidName(ComposeName(infra.TargetSshKey.Name, seed)); !ok {
		return false, fmt.Sprintf("SSH Key name [%s]: %s", infra.TargetSshKey.Name, detail)
	}

	// Check Security Groups
	for _, sg := range infra.TargetSecurityGroupList {
		if ok, detail := IsValidName(ComposeName(sg.Name, seed)); !ok {
			return false, fmt.Sprintf("Security Group name [%s]: %s", sg.Name, detail)
		}
	}

	// Check Infra and NodeGroups
	if ok, detail := IsValidName(ComposeName(infra.TargetInfra.Name, seed)); !ok {
		return false, fmt.Sprintf("Infra name [%s]: %s", infra.TargetInfra.Name, detail)
	}
	for _, ng := range infra.TargetInfra.NodeGroups {
		if ok, detail := IsValidName(ComposeName(ng.Name, seed)); !ok {
			return false, fmt.Sprintf("NodeGroup/VM name [%s]: %s", ng.Name, detail)
		}
	}

	// 2. Validate Referential Integrity
	if ok, detail := ValidateReferentialIntegrity(infra); !ok {
		return false, detail
	}

	return true, ""
}

// ValidateReferentialIntegrity verifies that all internal references (IDs)
// in the model point to resources that exist within the same model.
func ValidateReferentialIntegrity(infra cloudmodel.RecommendedInfra) (bool, string) {
	vnetName := infra.TargetVNet.Name
	sshKeyName := infra.TargetSshKey.Name

	// Map subnets for quick lookup
	subnets := make(map[string]bool)
	for _, s := range infra.TargetVNet.SubnetInfoList {
		subnets[s.Name] = true
	}

	// Map security groups for quick lookup
	sgs := make(map[string]bool)
	for _, sg := range infra.TargetSecurityGroupList {
		sgs[sg.Name] = true
		// Check SG -> VNet reference
		if sg.VNetId != vnetName {
			return false, fmt.Sprintf("Security Group [%s] refers to non-existent VNet [%s]", sg.Name, sg.VNetId)
		}
	}

	// Check Infra NodeGroups references
	for _, ng := range infra.TargetInfra.NodeGroups {
		if ng.VNetId != vnetName {
			return false, fmt.Sprintf("NodeGroup [%s] refers to non-existent VNet [%s]", ng.Name, ng.VNetId)
		}
		if !subnets[ng.SubnetId] {
			return false, fmt.Sprintf("NodeGroup [%s] refers to non-existent Subnet [%s]", ng.Name, ng.SubnetId)
		}
		if ng.SshKeyId != sshKeyName {
			return false, fmt.Sprintf("NodeGroup [%s] refers to non-existent SSH Key [%s]", ng.Name, ng.SshKeyId)
		}
		for _, sgId := range ng.SecurityGroupIds {
			if !sgs[sgId] {
				return false, fmt.Sprintf("NodeGroup [%s] refers to non-existent Security Group [%s]", ng.Name, sgId)
			}
		}
	}

	return true, ""
}
