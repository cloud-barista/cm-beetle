/*
Copyright 2024 The Cloud-Barista Authors.
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
	"regexp"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/csp"
	"github.com/rs/zerolog/log"
)

const (
	minViableWorkerVcpu   = 2
	minViableWorkerMemGiB = 4.0
)

// ValidateTargetK8sInfra checks whether target K8s infra model is valid and can be provisioned.
func ValidateTargetK8sInfra(nsId string, target *cloudmodel.RecommendedInfra) ValidationResult {
	if target == nil {
		return newResult([]ValidationIssue{{
			Code:     CodeRequiredFieldMissing,
			Severity: SeverityError,
			Path:     "target",
			Message:  "target infrastructure model is nil",
		}})
	}

	cluster := target.TargetK8sCluster
	if cluster.Name == "" {
		return newResult([]ValidationIssue{{
			Code:     CodeRequiredFieldMissing,
			Severity: SeverityError,
			Path:     "targetK8sCluster.name",
			Message:  "K8s cluster name is required",
		}})
	}

	var issues []ValidationIssue

	// Validate cluster name format
	if ok, detail := common.IsValidName(cluster.Name); !ok {
		issues = append(issues, ValidationIssue{
			Code:     CodeReferentialIntegrity,
			Severity: SeverityError,
			Path:     "targetK8sCluster.name",
			Message:  fmt.Sprintf("K8s cluster name [%s]: %s", cluster.Name, detail),
		})
	}

	// Validate companion resource names if specified for fresh creation
	if target.TargetVNet.Name != "" {
		if ok, detail := common.IsValidName(target.TargetVNet.Name); !ok {
			issues = append(issues, ValidationIssue{
				Code:     CodeReferentialIntegrity,
				Severity: SeverityError,
				Path:     "targetVNet.name",
				Message:  fmt.Sprintf("VNet name [%s]: %s", target.TargetVNet.Name, detail),
			})
		}
		for i, subnet := range target.TargetVNet.SubnetInfoList {
			if ok, detail := common.IsValidName(subnet.Name); !ok {
				issues = append(issues, ValidationIssue{
					Code:     CodeReferentialIntegrity,
					Severity: SeverityError,
					Path:     fmt.Sprintf("targetVNet.subnetInfoList[%d].name", i),
					Message:  fmt.Sprintf("Subnet name [%s]: %s", subnet.Name, detail),
				})
			}
		}
	}

	if target.TargetSshKey.Name != "" {
		if ok, detail := common.IsValidName(target.TargetSshKey.Name); !ok {
			issues = append(issues, ValidationIssue{
				Code:     CodeReferentialIntegrity,
				Severity: SeverityError,
				Path:     "targetSshKey.name",
				Message:  fmt.Sprintf("SSH key name [%s]: %s", target.TargetSshKey.Name, detail),
			})
		}
	}

	for i, sg := range target.TargetSecurityGroupList {
		if sg.Name != "" {
			if ok, detail := common.IsValidName(sg.Name); !ok {
				issues = append(issues, ValidationIssue{
					Code:     CodeReferentialIntegrity,
					Severity: SeverityError,
					Path:     fmt.Sprintf("targetSecurityGroupList[%d].name", i),
					Message:  fmt.Sprintf("SecurityGroup name [%s]: %s", sg.Name, detail),
				})
			}
		}
	}

	connectionName := strings.TrimSpace(cluster.ConnectionName)
	if connectionName == "" {
		issues = append(issues, ValidationIssue{
			Code:     CodeRequiredFieldMissing,
			Severity: SeverityError,
			Path:     "targetK8sCluster.connectionName",
			Message:  "connectionName is required",
		})
		return newResult(issues)
	}

	connParts := strings.Split(connectionName, "-")
	if len(connParts) < 2 {
		issues = append(issues, ValidationIssue{
			Code:     CodeInvalidConnectionName,
			Severity: SeverityError,
			Path:     "targetK8sCluster.connectionName",
			Message:  fmt.Sprintf("invalid connection name format %q, expected 'csp-region'", connectionName),
		})
		return newResult(issues)
	}
	provider := strings.ToLower(connParts[0])
	region := strings.ToLower(strings.Join(connParts[1:], "-"))

	// Resolve provider and region from Tumblebug connection configuration if available
	connConfig, connErr := tbclient.NewSession().GetConnConfig(connectionName)
	if connErr == nil {
		if connConfig.ProviderName != "" {
			provider = strings.ToLower(connConfig.ProviderName)
		}
		if connConfig.RegionDetail.RegionName != "" {
			region = strings.ToLower(connConfig.RegionDetail.RegionName)
		}
	}

	// Check if K8s cluster already exists in the namespace
	clusterInfo, err := tbclient.NewSession().ReadK8sCluster(nsId, cluster.Name)
	if err == nil && clusterInfo.Id != "" {
		issues = append(issues, ValidationIssue{
			Code:     CodeResourceAlreadyExists,
			Severity: SeverityError,
			Path:     "targetK8sCluster.name",
			Message:  fmt.Sprintf("K8s cluster '%s' already exists in namespace '%s'", cluster.Name, nsId),
		})
	}

	// Validate K8s version compatibility
	if cluster.Version != "" {
		availableVersions, err := tbclient.NewSession().GetAvailableK8sVersions(provider, region)
		if err == nil && len(availableVersions) > 0 {
			versionMatch := false
			for _, av := range availableVersions {
				if av.Name == cluster.Version || av.Id == cluster.Version || strings.HasPrefix(av.Id, cluster.Version) {
					versionMatch = true
					break
				}
			}
			if !versionMatch {
				var validNames []string
				for _, av := range availableVersions {
					validNames = append(validNames, av.Name)
				}
				issues = append(issues, ValidationIssue{
					Code:     CodeUnsupportedK8sVersion,
					Severity: SeverityError,
					Path:     "targetK8sCluster.version",
					Message: fmt.Sprintf("K8s version '%s' is not supported for provider '%s' in region '%s'; available versions: [%s]",
						cluster.Version, provider, region, strings.Join(validNames, ", ")),
				})
			}
		}
	}

	// Validate CSP K8s profile constraints
	profile, profileErr := tbclient.NewSession().GetK8sClusterProfile(provider, region)
	if profileErr == nil {
		// Check subnet count requirement
		subnetCount := len(cluster.SubnetIds)
		if subnetCount == 0 && target.TargetVNet.Name != "" {
			subnetCount = len(target.TargetVNet.SubnetInfoList)
		}
		if profile.RequiredSubnetCount > 0 && subnetCount < profile.RequiredSubnetCount {
			issues = append(issues, ValidationIssue{
				Code:     CodeInsufficientSubnets,
				Severity: SeverityError,
				Path:     "targetK8sCluster.subnetIds",
				Message: fmt.Sprintf("provider '%s' requires at least %d subnet(s), but %d provided",
					provider, profile.RequiredSubnetCount, subnetCount),
			})
		}

		// Check required initial node group
		if profile.NodeGroupsOnCreation && len(cluster.K8sNodeGroupList) == 0 {
			issues = append(issues, ValidationIssue{
				Code:     CodeMissingRequiredNodeGroup,
				Severity: SeverityError,
				Path:     "targetK8sCluster.k8sNodeGroupList",
				Message:  fmt.Sprintf("provider '%s' requires at least one node group upon cluster creation", provider),
			})
		}

		var ngNameRegex *regexp.Regexp
		if profile.NodeGroupNamingRule != "" {
			var err error
			ngNameRegex, err = regexp.Compile(profile.NodeGroupNamingRule)
			if err != nil {
				log.Warn().Err(err).Str("rule", profile.NodeGroupNamingRule).Msg("Failed to compile nodeGroupNamingRule regex")
			}
		}

		var specNameRegex *regexp.Regexp
		if profile.NodeSpecNamingRule != "" {
			var err error
			specNameRegex, err = regexp.Compile(profile.NodeSpecNamingRule)
			if err != nil {
				log.Warn().Err(err).Str("rule", profile.NodeSpecNamingRule).Msg("Failed to compile nodeSpecNamingRule regex")
			}
		}

		// Validate each node group
		for i, ng := range cluster.K8sNodeGroupList {
			path := fmt.Sprintf("targetK8sCluster.k8sNodeGroupList[%d]", i)

			// Validate node group name requirement and naming rule
			if profile.RequireNodeGroupName && strings.TrimSpace(ng.Name) == "" {
				issues = append(issues, ValidationIssue{
					Code:     CodeRequiredFieldMissing,
					Severity: SeverityError,
					Path:     path + ".name",
					Message:  fmt.Sprintf("node group name is required for provider '%s' (naming rule: %s)", provider, profile.NodeGroupNamingRule),
				})
			} else if ng.Name != "" && ngNameRegex != nil {
				if !ngNameRegex.MatchString(ng.Name) {
					issues = append(issues, ValidationIssue{
						Code:     CodeInvalidNamingRule,
						Severity: SeverityError,
						Path:     path + ".name",
						Message:  fmt.Sprintf("node group name %q does not match naming rule %q for provider '%s'", ng.Name, profile.NodeGroupNamingRule, provider),
					})
				}
			}

			// Validate autoscaling sizing constraints
			isAutoScaling := strings.EqualFold(ng.OnAutoScaling, "true") || ng.OnAutoScaling == "1"
			if !isAutoScaling {
				// Azure rejects minNodeSize > 0 when autoscaling is disabled
				if provider == "azure" && ng.MinNodeSize > 0 {
					issues = append(issues, ValidationIssue{
						Code:     CodeInvalidNodeGroupSize,
						Severity: SeverityError,
						Path:     path + ".minNodeSize",
						Message:  fmt.Sprintf("Azure AKS requires minNodeSize to be 0 when onAutoScaling is false, got %d", ng.MinNodeSize),
					})
				}
				// Fixed size node groups require positive desired count
				if (provider == csp.AWS || provider == csp.Tencent || provider == csp.IBM) && ng.DesiredNodeSize < 1 {
					issues = append(issues, ValidationIssue{
						Code:     CodeInvalidNodeGroupSize,
						Severity: SeverityError,
						Path:     path + ".desiredNodeSize",
						Message:  fmt.Sprintf("provider '%s' requires desiredNodeSize >= 1, got %d", provider, ng.DesiredNodeSize),
					})
				}
			} else {
				// Validate min <= desired <= max relationship for autoscaling
				if ng.MinNodeSize < 0 || ng.DesiredNodeSize < ng.MinNodeSize || ng.MaxNodeSize < ng.DesiredNodeSize {
					issues = append(issues, ValidationIssue{
						Code:     CodeInvalidNodeGroupSize,
						Severity: SeverityError,
						Path:     path,
						Message: fmt.Sprintf("invalid autoscaling sizes for node group %q: must satisfy minNodeSize (%d) <= desiredNodeSize (%d) <= maxNodeSize (%d)",
							ng.Name, ng.MinNodeSize, ng.DesiredNodeSize, ng.MaxNodeSize),
					})
				}
			}

			// Validate root disk size minimum
			if profile.RootDiskSizeMinGB > 0 && ng.RootDiskSize > 0 && ng.RootDiskSize < profile.RootDiskSizeMinGB {
				issues = append(issues, ValidationIssue{
					Code:     CodeInvalidRootDiskSize,
					Severity: SeverityError,
					Path:     path + ".rootDiskSize",
					Message: fmt.Sprintf("root disk size %d GB is below minimum required size %d GB for provider '%s' in region '%s'",
						ng.RootDiskSize, profile.RootDiskSizeMinGB, provider, region),
				})
			}

			// Validate spec size viability and spec naming rule
			if ng.SpecId != "" {
				specInfo, err := tbclient.NewSession().ReadVmSpec("system", ng.SpecId)
				if err == nil {
					if specInfo.VCPU < minViableWorkerVcpu || specInfo.MemoryGiB < minViableWorkerMemGiB {
						issues = append(issues, ValidationIssue{
							Code:     CodeInvalidNodeGroupSize,
							Severity: SeverityError,
							Path:     path + ".specId",
							Message: fmt.Sprintf("node group %q spec %q (%dvCPU/%.0fGiB) is below the minimum viable K8s worker spec (%dvCPU/%dGiB)",
								ng.Name, ng.SpecId, specInfo.VCPU, specInfo.MemoryGiB, minViableWorkerVcpu, int(minViableWorkerMemGiB)),
						})
					}
					if specNameRegex != nil && specInfo.CspSpecName != "" && !specNameRegex.MatchString(specInfo.CspSpecName) {
						issues = append(issues, ValidationIssue{
							Code:     CodeInvalidNamingRule,
							Severity: SeverityError,
							Path:     path + ".specId",
							Message: fmt.Sprintf("spec %q (CspSpecName: %q) does not match nodeSpecNamingRule %q for provider '%s'",
								ng.SpecId, specInfo.CspSpecName, profile.NodeSpecNamingRule, provider),
						})
					}
				}
			}
		}
	}

	// Validate fresh creation companion resources do not collide
	if target.TargetVNet.Name != "" {
		if vnetInfo, err := tbclient.NewSession().ReadVNet(nsId, target.TargetVNet.Name); err == nil && vnetInfo.Id != "" {
			issues = append(issues, ValidationIssue{
				Code:     CodeResourceAlreadyExists,
				Severity: SeverityError,
				Path:     "targetVNet.name",
				Message:  fmt.Sprintf("VNet '%s' already exists in namespace '%s'", target.TargetVNet.Name, nsId),
			})
		}
	}

	if target.TargetSshKey.Name != "" {
		if sshKeyInfo, err := tbclient.NewSession().ReadSshKey(nsId, target.TargetSshKey.Name); err == nil && sshKeyInfo.Id != "" {
			issues = append(issues, ValidationIssue{
				Code:     CodeResourceAlreadyExists,
				Severity: SeverityError,
				Path:     "targetSshKey.name",
				Message:  fmt.Sprintf("SSH key '%s' already exists in namespace '%s'", target.TargetSshKey.Name, nsId),
			})
		}
	}

	for i, sg := range target.TargetSecurityGroupList {
		if sg.Name == "" {
			continue
		}
		if sgInfo, err := tbclient.NewSession().ReadSecurityGroup(nsId, sg.Name); err == nil && sgInfo.Id != "" {
			issues = append(issues, ValidationIssue{
				Code:     CodeResourceAlreadyExists,
				Severity: SeverityError,
				Path:     fmt.Sprintf("targetSecurityGroupList[%d].name", i),
				Message:  fmt.Sprintf("security group '%s' already exists in namespace '%s'", sg.Name, nsId),
			})
		}
	}

	return newResult(issues)
}

