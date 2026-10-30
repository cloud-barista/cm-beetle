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
	"maps"
	"slices"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
)

// ComposeName combines a base name and a seed when both are non-empty.
func ComposeName(baseName, seed string) string {
	if baseName == "" || seed == "" {
		return baseName
	}
	return fmt.Sprintf("%s-%s", seed, baseName)
}

// DeepCopyRecommendedInfra returns an independent deep copy of RecommendedInfra.
func DeepCopyRecommendedInfra(infra cloudmodel.RecommendedInfra) cloudmodel.RecommendedInfra {
	result := infra

	result.TargetVNet.SubnetInfoList = slices.Clone(infra.TargetVNet.SubnetInfoList)

	result.TargetInfra.Label = maps.Clone(infra.TargetInfra.Label)
	result.TargetInfra.NodeGroups = slices.Clone(infra.TargetInfra.NodeGroups)
	for i, ng := range infra.TargetInfra.NodeGroups {
		result.TargetInfra.NodeGroups[i].Label = maps.Clone(ng.Label)
		result.TargetInfra.NodeGroups[i].SubnetIds = slices.Clone(ng.SubnetIds)
		result.TargetInfra.NodeGroups[i].SecurityGroupIds = slices.Clone(ng.SecurityGroupIds)
		result.TargetInfra.NodeGroups[i].DataDiskIds = slices.Clone(ng.DataDiskIds)
	}
	result.TargetInfra.PostCommands = slices.Clone(infra.TargetInfra.PostCommands)
	for i, pc := range infra.TargetInfra.PostCommands {
		result.TargetInfra.PostCommands[i].Command = slices.Clone(pc.Command)
	}

	result.TargetSpecList = slices.Clone(infra.TargetSpecList)
	for i, sp := range infra.TargetSpecList {
		result.TargetSpecList[i].AssociatedObjectList = slices.Clone(sp.AssociatedObjectList)
		result.TargetSpecList[i].Details = slices.Clone(sp.Details)
		result.TargetSpecList[i].AvailableZones = slices.Clone(sp.AvailableZones)
		result.TargetSpecList[i].UnavailableZones = slices.Clone(sp.UnavailableZones)
		if sp.AvailabilityCheckedAt != nil {
			t := *sp.AvailabilityCheckedAt
			result.TargetSpecList[i].AvailabilityCheckedAt = &t
		}
		if sp.AvailabilityRefreshAfter != nil {
			t := *sp.AvailabilityRefreshAfter
			result.TargetSpecList[i].AvailabilityRefreshAfter = &t
		}
	}

	result.TargetOsImageList = slices.Clone(infra.TargetOsImageList)
	for i, img := range infra.TargetOsImageList {
		result.TargetOsImageList[i].RegionList = slices.Clone(img.RegionList)
		result.TargetOsImageList[i].Details = slices.Clone(img.Details)
		result.TargetOsImageList[i].CommandHistory = slices.Clone(img.CommandHistory)
	}

	result.TargetSecurityGroupList = slices.Clone(infra.TargetSecurityGroupList)
	for i, sg := range infra.TargetSecurityGroupList {
		if sg.FirewallRules != nil {
			rules := slices.Clone(*sg.FirewallRules)
			result.TargetSecurityGroupList[i].FirewallRules = &rules
		}
	}

	result.TargetNlbList = slices.Clone(infra.TargetNlbList)

	result.TargetK8sCluster.SubnetIds = slices.Clone(infra.TargetK8sCluster.SubnetIds)
	result.TargetK8sCluster.SecurityGroupIds = slices.Clone(infra.TargetK8sCluster.SecurityGroupIds)
	result.TargetK8sCluster.Label = maps.Clone(infra.TargetK8sCluster.Label)
	result.TargetK8sCluster.K8sNodeGroupList = slices.Clone(infra.TargetK8sCluster.K8sNodeGroupList)
	for i, kng := range infra.TargetK8sCluster.K8sNodeGroupList {
		result.TargetK8sCluster.K8sNodeGroupList[i].Label = maps.Clone(kng.Label)
	}
	if infra.TargetK8sCluster.Storage != nil {
		st := *infra.TargetK8sCluster.Storage
		st.StorageTypes = slices.Clone(st.StorageTypes)
		result.TargetK8sCluster.Storage = &st
	}

	return result
}

// ApplyNameSeed applies the late-binding naming strategy using the given seed.
// It returns a new deep-copied RecommendedInfra object with all resource names prefixed by seed.
func ApplyNameSeed(infra cloudmodel.RecommendedInfra, seed string) cloudmodel.RecommendedInfra {
	result := DeepCopyRecommendedInfra(infra)
	if seed == "" {
		return result
	}

	// 1. Update VNet and Subnets
	result.TargetVNet.Name = ComposeName(infra.TargetVNet.Name, seed)
	for i, subnet := range result.TargetVNet.SubnetInfoList {
		result.TargetVNet.SubnetInfoList[i].Name = ComposeName(subnet.Name, seed)
	}

	// 2. Update SSH Key
	result.TargetSshKey.Name = ComposeName(infra.TargetSshKey.Name, seed)

	// 3. Update Security Groups
	for i, sg := range result.TargetSecurityGroupList {
		result.TargetSecurityGroupList[i].Name = ComposeName(sg.Name, seed)
		// Update reference to VNetId if it matches the base VNet name
		if infra.TargetVNet.Name != "" && sg.VNetId == infra.TargetVNet.Name {
			result.TargetSecurityGroupList[i].VNetId = result.TargetVNet.Name
		}
	}

	// 4. Update K8s cluster and its references (node group names remain scoped to the cluster).
	if infra.TargetK8sCluster.Name != "" {
		result.TargetK8sCluster.Name = ComposeName(infra.TargetK8sCluster.Name, seed)
		result.TargetK8sCluster.VNetId = ComposeName(infra.TargetK8sCluster.VNetId, seed)

		for i, subnetId := range result.TargetK8sCluster.SubnetIds {
			result.TargetK8sCluster.SubnetIds[i] = ComposeName(subnetId, seed)
		}
		for i, sgId := range result.TargetK8sCluster.SecurityGroupIds {
			result.TargetK8sCluster.SecurityGroupIds[i] = ComposeName(sgId, seed)
		}
		for i, ng := range result.TargetK8sCluster.K8sNodeGroupList {
			result.TargetK8sCluster.K8sNodeGroupList[i].SshKeyId = ComposeName(ng.SshKeyId, seed)
		}
	}

	// 5. Update Infra, NodeGroups, and PostCommands
	result.TargetInfra.Name = ComposeName(infra.TargetInfra.Name, seed)
	for i, ng := range result.TargetInfra.NodeGroups {
		result.TargetInfra.NodeGroups[i].Name = ComposeName(ng.Name, seed)
		result.TargetInfra.NodeGroups[i].VNetId = ComposeName(ng.VNetId, seed)
		result.TargetInfra.NodeGroups[i].SubnetId = ComposeName(ng.SubnetId, seed)
		for j, subnetId := range ng.SubnetIds {
			result.TargetInfra.NodeGroups[i].SubnetIds[j] = ComposeName(subnetId, seed)
		}
		result.TargetInfra.NodeGroups[i].SshKeyId = ComposeName(ng.SshKeyId, seed)
		for j, sgId := range ng.SecurityGroupIds {
			result.TargetInfra.NodeGroups[i].SecurityGroupIds[j] = ComposeName(sgId, seed)
		}
	}
	for i, pc := range result.TargetInfra.PostCommands {
		result.TargetInfra.PostCommands[i].NodeGroupId = ComposeName(pc.NodeGroupId, seed)
	}

	// 6. Update NLB target group NodeGroupId references
	for i, nlb := range result.TargetNlbList {
		result.TargetNlbList[i].TargetGroup.NodeGroupId = ComposeName(nlb.TargetGroup.NodeGroupId, seed)
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
// (e.g., If VNet is renamed, all SecurityGroup.VNetId and NodeGroup.VNetId are updated)
//
// Parameters:
//   - infra: The recommendation model to update (will be copied, not mutated)
//   - resourceType: One of "vNet", "subnet", "sshKey", "securityGroup", "infra"
//   - oldName: The current name of the resource to rename
//   - newName: The new name to assign to the resource
func PropagateNameChange(infra cloudmodel.RecommendedInfra, resourceType, oldName, newName string) cloudmodel.RecommendedInfra {
	result := DeepCopyRecommendedInfra(infra)
	if oldName == "" {
		return result
	}

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
		// 3. Propagate to NodeGroups and K8s cluster
		for i := range result.TargetInfra.NodeGroups {
			if result.TargetInfra.NodeGroups[i].VNetId == oldName {
				result.TargetInfra.NodeGroups[i].VNetId = newName
			}
		}
		if result.TargetK8sCluster.VNetId == oldName {
			result.TargetK8sCluster.VNetId = newName
		}

	case ResourceTypeSubnet:
		// 1. Rename the Subnet itself
		for i := range result.TargetVNet.SubnetInfoList {
			if result.TargetVNet.SubnetInfoList[i].Name == oldName {
				result.TargetVNet.SubnetInfoList[i].Name = newName
			}
		}
		// 2. Propagate to NodeGroups and K8s cluster
		for i, ng := range result.TargetInfra.NodeGroups {
			if ng.SubnetId == oldName {
				result.TargetInfra.NodeGroups[i].SubnetId = newName
			}
			for j, subnetId := range ng.SubnetIds {
				if subnetId == oldName {
					result.TargetInfra.NodeGroups[i].SubnetIds[j] = newName
				}
			}
		}
		for i, subnetId := range result.TargetK8sCluster.SubnetIds {
			if subnetId == oldName {
				result.TargetK8sCluster.SubnetIds[i] = newName
			}
		}

	case ResourceTypeSshKey:
		// 1. Rename the SSH Key itself
		if result.TargetSshKey.Name == oldName {
			result.TargetSshKey.Name = newName
		}
		// 2. Propagate to NodeGroups and K8s NodeGroups
		for i := range result.TargetInfra.NodeGroups {
			if result.TargetInfra.NodeGroups[i].SshKeyId == oldName {
				result.TargetInfra.NodeGroups[i].SshKeyId = newName
			}
		}
		for i := range result.TargetK8sCluster.K8sNodeGroupList {
			if result.TargetK8sCluster.K8sNodeGroupList[i].SshKeyId == oldName {
				result.TargetK8sCluster.K8sNodeGroupList[i].SshKeyId = newName
			}
		}

	case ResourceTypeSecurityGroup:
		// 1. Rename the SecurityGroup itself
		for i := range result.TargetSecurityGroupList {
			if result.TargetSecurityGroupList[i].Name == oldName {
				result.TargetSecurityGroupList[i].Name = newName
			}
		}
		// 2. Propagate to NodeGroup and K8s SecurityGroupIds
		for i, ng := range result.TargetInfra.NodeGroups {
			for j, sgId := range ng.SecurityGroupIds {
				if sgId == oldName {
					result.TargetInfra.NodeGroups[i].SecurityGroupIds[j] = newName
				}
			}
		}
		for i, sgId := range result.TargetK8sCluster.SecurityGroupIds {
			if sgId == oldName {
				result.TargetK8sCluster.SecurityGroupIds[i] = newName
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

// IsValidNameSeed checks if the nameSeed query param is acceptable.
// Rules: 1-20 characters, alphanumeric and hyphens only, starts with alphanumeric.
// An empty string is valid (means no seed is applied).
func IsValidNameSeed(seed string) (bool, string) {
	if seed == "" {
		return true, ""
	}
	if len(seed) > 20 {
		return false, "nameSeed must be 20 characters or fewer"
	}
	for i, r := range seed {
		if i == 0 && !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')) {
			return false, "nameSeed must start with an alphanumeric character"
		}
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-') {
			return false, "nameSeed contains invalid characters (only alphanumeric and hyphens allowed)"
		}
	}
	return true, ""
}

// ValidateComposedNames checks all resource names and referential integrity.
// This function expects names to already have NameSeed applied (if any).
func ValidateComposedNames(infra cloudmodel.RecommendedInfra) (bool, string) {
	// 1. Validate name formats (names are validated as-is; seed should be applied before calling this)
	// Check VNet
	if ok, detail := IsValidName(infra.TargetVNet.Name); !ok {
		return false, fmt.Sprintf("VNet name [%s]: %s", infra.TargetVNet.Name, detail)
	}

	// Check Subnets
	for _, subnet := range infra.TargetVNet.SubnetInfoList {
		if ok, detail := IsValidName(subnet.Name); !ok {
			return false, fmt.Sprintf("Subnet name [%s]: %s", subnet.Name, detail)
		}
	}

	// Check SSH Key
	if ok, detail := IsValidName(infra.TargetSshKey.Name); !ok {
		return false, fmt.Sprintf("SSH Key name [%s]: %s", infra.TargetSshKey.Name, detail)
	}

	// Check Security Groups
	for _, sg := range infra.TargetSecurityGroupList {
		if ok, detail := IsValidName(sg.Name); !ok {
			return false, fmt.Sprintf("Security Group name [%s]: %s", sg.Name, detail)
		}
	}

	// Check Infra and NodeGroups
	if ok, detail := IsValidName(infra.TargetInfra.Name); !ok {
		return false, fmt.Sprintf("Infra name [%s]: %s", infra.TargetInfra.Name, detail)
	}
	for _, ng := range infra.TargetInfra.NodeGroups {
		if ok, detail := IsValidName(ng.Name); !ok {
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
	nodeGroups := make(map[string]bool, len(infra.TargetInfra.NodeGroups))
	for _, ng := range infra.TargetInfra.NodeGroups {
		nodeGroups[ng.Name] = true
		if ng.VNetId != vnetName {
			return false, fmt.Sprintf("NodeGroup [%s] refers to non-existent VNet [%s]", ng.Name, ng.VNetId)
		}
		if !subnets[ng.SubnetId] {
			return false, fmt.Sprintf("NodeGroup [%s] refers to non-existent Subnet [%s]", ng.Name, ng.SubnetId)
		}
		for _, subnetId := range ng.SubnetIds {
			if !subnets[subnetId] {
				return false, fmt.Sprintf("NodeGroup [%s] refers to non-existent Subnet [%s] in subnetIds", ng.Name, subnetId)
			}
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

	// Check optional PostCommand and NLB NodeGroupId references
	for _, pc := range infra.TargetInfra.PostCommands {
		if pc.NodeGroupId != "" && !nodeGroups[pc.NodeGroupId] {
			return false, fmt.Sprintf("PostCommand refers to non-existent NodeGroup [%s]", pc.NodeGroupId)
		}
	}
	for _, nlb := range infra.TargetNlbList {
		if nlb.TargetGroup.NodeGroupId != "" && len(nodeGroups) > 0 && !nodeGroups[nlb.TargetGroup.NodeGroupId] {
			return false, fmt.Sprintf("NLB refers to non-existent NodeGroup [%s]", nlb.TargetGroup.NodeGroupId)
		}
	}

	return true, ""
}
