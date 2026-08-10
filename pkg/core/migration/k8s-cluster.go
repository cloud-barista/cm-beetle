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

// Package migration contains the K8s cluster migration logic.
package migration

import (
	"fmt"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	"github.com/rs/zerolog/log"
)

// k8sPrereqs tracks the prerequisite resources (VNet, SshKey, SecurityGroups) that were
// freshly created by a single CreateK8sCluster call. Only resources actually created in
// this run are tracked; resources that already existed (idempotency case) are not included
// so that rollback does not delete resources owned by other workloads.
type k8sPrereqs struct {
	nsId     string
	vNetId   string
	sshKeyId string
	sgIds    []string
}

// rollback deletes the tracked resources in reverse creation order (SGs → SshKey → VNet).
// Errors during rollback are logged as warnings but do not prevent subsequent deletions.
func (p *k8sPrereqs) rollback() {
	for i := len(p.sgIds) - 1; i >= 0; i-- {
		if _, err := tbclient.NewSession().DeleteSecurityGroup(p.nsId, p.sgIds[i]); err != nil {
			log.Warn().Err(err).Str("sgId", p.sgIds[i]).Msg("Rollback: failed to delete SecurityGroup")
		} else {
			log.Info().Str("sgId", p.sgIds[i]).Msg("Rollback: SecurityGroup deleted")
		}
	}
	if p.sshKeyId != "" {
		if _, err := tbclient.NewSession().DeleteSshKey(p.nsId, p.sshKeyId); err != nil {
			log.Warn().Err(err).Str("sshKeyId", p.sshKeyId).Msg("Rollback: failed to delete SshKey")
		} else {
			log.Info().Str("sshKeyId", p.sshKeyId).Msg("Rollback: SshKey deleted")
		}
	}
	if p.vNetId != "" {
		if _, err := tbclient.NewSession().DeleteVNet(p.nsId, p.vNetId, "withSubnets"); err != nil {
			log.Warn().Err(err).Str("vNetId", p.vNetId).Msg("Rollback: failed to delete VNet")
		} else {
			log.Info().Str("vNetId", p.vNetId).Msg("Rollback: VNet deleted")
		}
	}
}

// cspRejectsCombinedIngressEgress lists CSPs that reject a SecurityGroup create request
// containing both ingress and egress rules in a single call (Tencent Cloud returns
// InvalidParameter.Coexist). For these, the SG is created with inbound rules first and the
// outbound rules are added via a separate call, so the full recommended ruleset is applied
// (rather than silently dropping the outbound rules).
var cspRejectsCombinedIngressEgress = map[string]bool{"tencent": true}

// createSecurityGroup creates a fresh SecurityGroup and tracks it for rollback. For CSPs that
// reject combined ingress/egress in one request, it creates with the inbound rules first, then
// adds the outbound rules via a separate call. It returns the create error unchanged so callers
// can still detect the "already exists" idempotency case.
func (p *k8sPrereqs) createSecurityGroup(provider string, sgReq tbmodel.SecurityGroupReq) (tbmodel.SecurityGroupInfo, error) {
	if !cspRejectsCombinedIngressEgress[strings.ToLower(provider)] {
		info, err := tbclient.NewSession().CreateSecurityGroup(p.nsId, sgReq, "")
		if err != nil {
			return info, err
		}
		p.sgIds = append(p.sgIds, info.Id) // track only freshly created SGs
		return info, nil
	}

	// Split path: create with inbound rules, then add outbound rules separately.
	inbound, outbound := splitFirewallRulesByDirection(sgReq.FirewallRules)
	sgReq.FirewallRules = &inbound
	info, err := tbclient.NewSession().CreateSecurityGroup(p.nsId, sgReq, "")
	if err != nil {
		return info, err
	}
	p.sgIds = append(p.sgIds, info.Id) // track before adding outbound so rollback covers it
	if len(outbound) > 0 {
		if _, err := tbclient.NewSession().AddFirewallRules(p.nsId, info.Id, outbound); err != nil {
			return info, fmt.Errorf("created SecurityGroup %s with inbound rules but failed to add outbound rules: %w", info.Id, err)
		}
	}
	return info, nil
}

// splitFirewallRulesByDirection separates a rule set into inbound and outbound rules.
func splitFirewallRulesByDirection(rules *[]tbmodel.FirewallRuleReq) (inbound, outbound []tbmodel.FirewallRuleReq) {
	if rules == nil {
		return nil, nil
	}
	for _, r := range *rules {
		if strings.EqualFold(r.Direction, "outbound") {
			outbound = append(outbound, r)
		} else {
			inbound = append(inbound, r)
		}
	}
	return inbound, outbound
}

// pollK8sClusterActive waits until the cluster reaches Active status, polling every 15 seconds
// up to maxAttempts times. It returns the latest cluster info regardless of the final status.
// If the cluster does not become Active within the timeout, a non-nil error is returned together
// with the last known cluster info so the caller can decide how to handle the partial state.
func pollK8sClusterActive(nsId, clusterId string, maxAttempts int) (tbmodel.K8sClusterInfo, error) {
	var clusterInfo tbmodel.K8sClusterInfo
	for i := 0; i < maxAttempts; i++ {
		updated, err := tbclient.NewSession().ReadK8sCluster(nsId, clusterId)
		if err != nil {
			log.Warn().Err(err).Int("attempt", i+1).Msg("Failed to poll cluster status; retrying")
		} else {
			clusterInfo = updated
			if string(clusterInfo.Status) == "Active" {
				log.Info().Str("clusterId", clusterId).Msg("K8s cluster is Active")
				return clusterInfo, nil
			}
			log.Debug().Str("status", string(clusterInfo.Status)).Int("attempt", i+1).Msg("Cluster not yet Active")
		}
		time.Sleep(15 * time.Second)
	}
	return clusterInfo, fmt.Errorf("K8s cluster %s did not reach Active state after %d attempts (last status: %s)",
		clusterId, maxAttempts, string(clusterInfo.Status))
}

// minViableWorkerVcpu / minViableWorkerMemGiB mirror the recommender's floor: managed K8s
// worker nodes below 2 vCPU / 4 GiB cannot reliably host the kubelet + system daemonsets, and
// Azure AKS rejects them for the mandatory system node pool.
const (
	minViableWorkerVcpu   = 2
	minViableWorkerMemGiB = 4.0
)

// validateNodeGroupSpecs rejects node groups whose worker spec is below the minimum viable
// size. It resolves each spec via Tumblebug so the check works on any request (not only ones
// produced by this service's recommender).
func validateNodeGroupSpecs(nodeGroups []cloudmodel.K8sNodeGroupReq) error {
	for _, ng := range nodeGroups {
		spec, err := tbclient.NewSession().ReadVmSpec("system", ng.SpecId)
		if err != nil {
			return fmt.Errorf("failed to look up spec %q for node group %q: %w", ng.SpecId, ng.Name, err)
		}
		if spec.VCPU < minViableWorkerVcpu || spec.MemoryGiB < minViableWorkerMemGiB {
			return fmt.Errorf("node group %q spec %q (%dvCPU/%.0fGiB) is below the minimum viable K8s worker spec (%dvCPU/%dGiB)",
				ng.Name, ng.SpecId, spec.VCPU, spec.MemoryGiB, minViableWorkerVcpu, int(minViableWorkerMemGiB))
		}
	}
	return nil
}

// CreateK8sCluster provisions a K8s cluster on the target CSP from a recommendation result.
// It follows the same pattern as CreateInfra for VM migration:
//  1. Create VNet (with subnets)
//  2. Create SSH Key
//  3. Create Security Group (sets VNetId from step 1)
//  4. Check CSP's nodeGroupsOnCreation flag
//     5a. nodeGroupsOnCreation=true  (Azure/GCP/NHN/NCP/IBM):
//     Include NodeGroups in K8sClusterReq and create in a single call, then poll for Active.
//     5b. nodeGroupsOnCreation=false (AWS/Alibaba/Tencent):
//     Create cluster first, poll for Active, then add NodeGroups separately.
//
// Both paths poll until the cluster is Active before returning, ensuring a consistent response.
func CreateK8sCluster(nsId string, req *cloudmodel.RecommendedInfra) (tbmodel.K8sClusterInfo, error) {
	log.Info().Str("nsId", nsId).Msg("Creating K8s cluster")

	emptyRet := tbmodel.K8sClusterInfo{}

	// prereqs tracks only resources created fresh in this call for rollback on failure.
	// Resources that already existed (idempotency case) are not tracked.
	prereqs := &k8sPrereqs{nsId: nsId}

	// 1. Verify namespace exists
	_, err := tbclient.NewSession().ReadNamespace(nsId)
	if err != nil {
		return emptyRet, fmt.Errorf("namespace not found (nsId: %s): %w", nsId, err)
	}

	// 1b. Reject node groups below the minimum viable worker spec before creating any resources.
	// The recommender already floors specs; this guards against manually-crafted or stale requests.
	if err := validateNodeGroupSpecs(req.TargetK8sCluster.K8sNodeGroupList); err != nil {
		return emptyRet, err
	}

	// 2. Create VNet
	tbVNetReq, err := modelconv.ConvertWithValidation[cloudmodel.VNetReq, tbmodel.VNetReq](req.TargetVNet)
	if err != nil {
		return emptyRet, fmt.Errorf("failed to convert VNet request: %w", err)
	}

	vNetInfo, err := tbclient.NewSession().CreateVNet(nsId, tbVNetReq)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Info().Str("vNetId", tbVNetReq.Name).Msg("VNet already exists, reading existing resource")
			vNetInfo, err = tbclient.NewSession().ReadVNet(nsId, tbVNetReq.Name)
			if err != nil {
				return emptyRet, fmt.Errorf("VNet already exists but failed to read it: %w", err)
			}
		} else {
			return emptyRet, fmt.Errorf("failed to create VNet: %w", err)
		}
	} else {
		prereqs.vNetId = vNetInfo.Id // track only freshly created VNet
	}
	log.Debug().Str("vNetId", vNetInfo.Id).Msg("VNet ready")

	// Collect subnet IDs from the created VNet
	subnetIds := make([]string, 0, len(vNetInfo.SubnetInfoList))
	for _, sn := range vNetInfo.SubnetInfoList {
		subnetIds = append(subnetIds, sn.Id)
	}

	// 3. Create SSH Key
	tbSshKeyReq, err := modelconv.ConvertWithValidation[cloudmodel.SshKeyReq, tbmodel.SshKeyReq](req.TargetSshKey)
	if err != nil {
		prereqs.rollback()
		return emptyRet, fmt.Errorf("failed to convert SshKey request: %w", err)
	}

	sshKeyInfo, err := tbclient.NewSession().CreateSshKey(nsId, tbSshKeyReq)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			log.Info().Str("sshKeyId", tbSshKeyReq.Name).Msg("SshKey already exists, reading existing resource")
			sshKeyInfo, err = tbclient.NewSession().ReadSshKey(nsId, tbSshKeyReq.Name)
			if err != nil {
				prereqs.rollback()
				return emptyRet, fmt.Errorf("SshKey already exists but failed to read it: %w", err)
			}
		} else {
			prereqs.rollback()
			return emptyRet, fmt.Errorf("failed to create SSH key: %w", err)
		}
	} else {
		prereqs.sshKeyId = sshKeyInfo.Id // track only freshly created SshKey
	}
	log.Debug().Str("sshKeyId", sshKeyInfo.Id).Msg("SSH key ready")

	// 4. Create Security Groups (set VNetId to the actual created VNet ID)
	sgIds := make([]string, 0, len(req.TargetSecurityGroupList))
	for _, sg := range req.TargetSecurityGroupList {
		sg.VNetId = vNetInfo.Id

		tbSgReq, err := modelconv.ConvertWithValidation[cloudmodel.SecurityGroupReq, tbmodel.SecurityGroupReq](sg)
		if err != nil {
			prereqs.rollback()
			return emptyRet, fmt.Errorf("failed to convert SecurityGroup request: %w", err)
		}

		sgInfo, err := prereqs.createSecurityGroup(req.TargetCloud.Csp, tbSgReq)
		if err != nil {
			if strings.Contains(err.Error(), "already exists") {
				log.Info().Str("sgId", sg.Name).Msg("SecurityGroup already exists, reading existing resource")
				sgInfo, err = tbclient.NewSession().ReadSecurityGroup(nsId, sg.Name)
				if err != nil {
					prereqs.rollback()
					return emptyRet, fmt.Errorf("SecurityGroup already exists but failed to read it: %w", err)
				}
			} else {
				prereqs.rollback()
				return emptyRet, fmt.Errorf("failed to create security group %s: %w", sg.Name, err)
			}
		}
		log.Debug().Str("sgId", sgInfo.Id).Msg("Security group ready")
		sgIds = append(sgIds, sgInfo.Id)
	}

	// 5. Check CSP's nodeGroupsOnCreation requirement.
	// Some CSPs (Azure, GCP, NHN, NCP, IBM) require NodeGroups to be included in the
	// cluster creation request. Others (AWS, Alibaba, Tencent) require separate addition.
	provider := req.TargetCloud.Csp
	nodeGroupsOnCreation, err := tbclient.NewSession().CheckK8sNodeGroupsOnCreation(provider)
	if err != nil {
		log.Warn().Err(err).Str("provider", provider).
			Msg("Failed to check nodeGroupsOnCreation; defaulting to separate NodeGroup creation")
		nodeGroupsOnCreation = false
	}
	log.Info().Str("provider", provider).Bool("nodeGroupsOnCreation", nodeGroupsOnCreation).
		Msg("K8s cluster creation strategy determined")

	// Build the NodeGroup requests with actual SSH key ID applied.
	tbNodeGroupReqs := make([]tbmodel.K8sNodeGroupReq, 0, len(req.TargetK8sCluster.K8sNodeGroupList))
	for _, ng := range req.TargetK8sCluster.K8sNodeGroupList {
		ng.SshKeyId = sshKeyInfo.Id
		tbNgReq, err := modelconv.ConvertWithValidation[cloudmodel.K8sNodeGroupReq, tbmodel.K8sNodeGroupReq](ng)
		if err != nil {
			return emptyRet, fmt.Errorf("failed to convert NodeGroup request for %s: %w", ng.Name, err)
		}
		tbNodeGroupReqs = append(tbNodeGroupReqs, tbNgReq)
	}

	// 6. Build K8sClusterReq with actual resource IDs.
	clusterReq := req.TargetK8sCluster
	clusterReq.VNetId = vNetInfo.Id
	clusterReq.SubnetIds = subnetIds
	clusterReq.SecurityGroupIds = sgIds

	if nodeGroupsOnCreation {
		// 6a. Include NodeGroups in the creation request (Azure/GCP/NHN/NCP/IBM).
		clusterReq.K8sNodeGroupList = req.TargetK8sCluster.K8sNodeGroupList
		for i := range clusterReq.K8sNodeGroupList {
			clusterReq.K8sNodeGroupList[i].SshKeyId = sshKeyInfo.Id
		}
	} else {
		// 6b. Create cluster without NodeGroups first (AWS/Alibaba/Tencent).
		clusterReq.K8sNodeGroupList = nil
	}

	tbClusterReq, err := modelconv.ConvertWithValidation[cloudmodel.K8sClusterReq, tbmodel.K8sClusterReq](clusterReq)
	if err != nil {
		return emptyRet, fmt.Errorf("failed to convert K8sCluster request: %w", err)
	}

	// 7. Create K8s cluster.
	clusterInfo, err := tbclient.NewSession().CreateK8sCluster(nsId, tbClusterReq)
	if err != nil {
		prereqs.rollback()
		return emptyRet, fmt.Errorf("failed to create K8s cluster: %w", err)
	}
	log.Info().Str("clusterId", clusterInfo.Id).Str("status", string(clusterInfo.Status)).Msg("K8s cluster created")

	// 8. Poll until Active — applies to both creation strategies for consistent API behaviour.
	// nodeGroupsOnCreation=true  (Azure/GCP): cluster + NodeGroups created together; poll here.
	// nodeGroupsOnCreation=false (AWS):       poll here before adding NodeGroups separately.
	// Generous timeout: IBM IKS takes ~22 min to reach Active even for a small (2-node) cluster,
	// and provisioning time grows with cluster size across all CSPs — so a per-CSP timeout would
	// be false precision. Kept synchronous for now; the async job pattern (see analysis #8) is the
	// eventual fix for holding the connection this long.
	const maxPollAttempts = 160 // 160 × 15 s = 40 min
	if string(clusterInfo.Status) != "Active" {
		log.Info().Str("clusterId", clusterInfo.Id).Str("provider", provider).Msg("Waiting for K8s cluster to become Active")
		var pollErr error
		clusterInfo, pollErr = pollK8sClusterActive(nsId, clusterInfo.Id, maxPollAttempts)
		if pollErr != nil {
			log.Warn().Err(pollErr).Str("clusterId", clusterInfo.Id).Msg("Cluster polling timed out; returning partial info")
			return clusterInfo, pollErr
		}
	}

	// For nodeGroupsOnCreation=true, NodeGroups are already part of the cluster — done.
	if nodeGroupsOnCreation {
		log.Info().Str("clusterId", clusterInfo.Id).Msg("K8s cluster migration completed (NodeGroups included at creation)")
		return clusterInfo, nil
	}

	for _, tbNgReq := range tbNodeGroupReqs {
		updatedClusterInfo, err := tbclient.NewSession().AddK8sNodeGroup(nsId, clusterInfo.Id, tbNgReq)
		if err != nil {
			// K8s cluster already exists; return it so the caller knows which cluster to clean up if needed.
			return clusterInfo, fmt.Errorf("K8s cluster %s created but node group %s failed — delete the cluster if worker nodes are not needed: %w", clusterInfo.Id, tbNgReq.Name, err)
		}
		clusterInfo = updatedClusterInfo
		log.Info().Str("nodeGroup", tbNgReq.Name).Str("status", string(clusterInfo.Status)).Msg("Node group added")
	}

	log.Info().Str("clusterId", clusterInfo.Id).Msg("K8s cluster migration completed")
	return clusterInfo, nil
}

// ListK8sClusters returns all K8s clusters in the namespace.
func ListK8sClusters(nsId string) ([]tbmodel.K8sClusterInfo, error) {
	list, err := tbclient.NewSession().ReadAllK8sClusters(nsId)
	if err != nil {
		return nil, err
	}
	if list.K8sClusters == nil {
		return []tbmodel.K8sClusterInfo{}, nil
	}
	return list.K8sClusters, nil
}

// GetK8sCluster returns the K8s cluster info for the given ID.
func GetK8sCluster(nsId, clusterId string) (tbmodel.K8sClusterInfo, error) {
	return tbclient.NewSession().ReadK8sCluster(nsId, clusterId)
}

// DeleteK8sCluster deletes the K8s cluster after removing all node groups.
// AWS EKS requires node groups to be deleted before the cluster can be deleted.
func DeleteK8sCluster(nsId, clusterId string) error {
	// Read the cluster to get its node groups.
	clusterInfo, err := tbclient.NewSession().ReadK8sCluster(nsId, clusterId)
	if err != nil {
		return fmt.Errorf("failed to read K8s cluster before deletion: %w", err)
	}

	// Delete all node groups first and wait for each to be removed.
	// EKS rejects cluster deletion while nodegroups are in any state other than deleted.
	for _, ng := range clusterInfo.K8sNodeGroupList {
		log.Info().Str("clusterId", clusterId).Str("nodeGroup", ng.Id).Msg("Deleting node group before cluster")
		if err := tbclient.NewSession().DeleteK8sNodeGroup(nsId, clusterId, ng.Id); err != nil {
			log.Warn().Err(err).Str("nodeGroup", ng.Id).Msg("Node group deletion failed (may already be gone); continuing")
			continue
		}
		// Poll until the nodegroup disappears from the cluster's nodegroup list.
		const maxNgWait = 80 // 80 × 15 s = 20 min
		for i := 0; i < maxNgWait; i++ {
			time.Sleep(15 * time.Second)
			updated, err := tbclient.NewSession().ReadK8sCluster(nsId, clusterId)
			if err != nil {
				break // cluster itself is gone; proceed to deletion
			}
			stillExists := false
			for _, existing := range updated.K8sNodeGroupList {
				if existing.Id == ng.Id {
					stillExists = true
					break
				}
			}
			if !stillExists {
				log.Info().Str("nodeGroup", ng.Id).Int("attempt", i+1).Msg("Node group deleted")
				break
			}
			log.Debug().Str("nodeGroup", ng.Id).Int("attempt", i+1).Msg("Waiting for node group deletion")
		}
	}

	return tbclient.NewSession().DeleteK8sCluster(nsId, clusterId)
}
