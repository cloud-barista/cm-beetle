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

// Package migration is to provision target infra for migration
package migration

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/cloud-barista/cm-beetle/pkg/core/validation"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	"github.com/rs/zerolog/log"
)

const (
	// ActionCreate is const for Create
	ActionCreate string = "Create"

	// ActionTerminate is const for Terminate
	ActionTerminate string = "Terminate"

	// ActionSuspend is const for Suspend
	ActionSuspend string = "Suspend"

	// ActionResume is const for Resume
	ActionResume string = "Resume"

	// ActionReboot is const for Reboot
	ActionReboot string = "Reboot"

	// ActionRefine is const for Refine
	ActionRefine string = "Refine"

	// ActionComplete is const for Complete
	ActionComplete string = "None"
)
const (
	// StatusRunning is const for Running
	StatusRunning string = "Running"

	// StatusSuspended is const for Suspended
	StatusSuspended string = "Suspended"

	// StatusFailed is const for Failed
	StatusFailed string = "Failed"

	// StatusTerminated is const for Terminated
	StatusTerminated string = "Terminated"

	// StatusCreating is const for Creating
	StatusCreating string = "Creating"

	// StatusSuspending is const for Suspending
	StatusSuspending string = "Suspending"

	// StatusResuming is const for Resuming
	StatusResuming string = "Resuming"

	// StatusRebooting is const for Rebooting
	StatusRebooting string = "Rebooting"

	// StatusTerminating is const for Terminating
	StatusTerminating string = "Terminating"

	// StatusUndefined is const for Undefined
	StatusUndefined string = "Undefined"

	// StatusComplete is const for Complete
	StatusComplete string = "None"
)

// DefaultSystemLabel is const for string to specify the Default System Label
const DefaultSystemLabel string = "Managed by CM-Beetle"

// CreateInfraWithDefaults Create an infrastructure with defaults for the computing infra migration
func CreateInfraWithDefaults(nsId string, infraModel *cloudmodel.InfraDynamicReq) (cloudmodel.VmInfraInfo, error) {
	log.Info().Msg("Creating an infrastructure with defaults")

	// Convert the request model from 'cloudmodel.InfraDynamicReq' to 'tbmodel.InfraDynamicReq'
	infraModelConverted, err := modelconv.ConvertWithValidation[cloudmodel.InfraDynamicReq, tbmodel.InfraDynamicReq](*infraModel)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure model (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	infraInfo, err := tbclient.NewSession().CreateInfraDynamic(nsId, infraModelConverted)
	if err != nil {
		log.Error().Err(err).Msgf("failed to migrate the infrastructure (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	// Convert the response model from 'tbmodel.InfraInfo' to 'cloudmodel.VmInfraInfo'
	convertedInfraInfo, err := modelconv.ConvertWithValidation[tbmodel.InfraInfo, cloudmodel.VmInfraInfo](infraInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	log.Info().Msgf("Infrastructure created successfully (nsId: %s, infraName: %s)", nsId, convertedInfraInfo.Name)

	return convertedInfraInfo, nil
}

// CreateInfra creates an infrastructure for the computing infra migration by creating fresh resources (useExisting=false)
func CreateInfra(nsId string, targetInfraModel *cloudmodel.RecommendedInfra) (cloudmodel.VmInfraInfo, error) {
	log.Info().Msg("Creating an infrastructure")

	emptyRet := cloudmodel.VmInfraInfo{}

	/*
	 * [Input] Receive and validate the target infrastructure model
	 */

	err := validation.ValidateTargetInfra(nsId, targetInfraModel, false).Err()
	if err != nil {
		log.Error().Err(err).Msgf("failed to validate the target infrastructure model (nsId: %s)", nsId)
		return emptyRet, err
	}
	log.Info().Msgf("the target infrastructure model is valid (nsId: %s)", nsId)

	// Preflight: resolve the latest CSP image and confirm available system disk per nodegroup.
	err = preflightCheckCspProvisioning(nsId, targetInfraModel.TargetInfra.NodeGroups)
	if err != nil {
		log.Error().Err(err).Msgf("failed to run preflight check for CSP provisioning (nsId: %s)", nsId)
		return emptyRet, err
	}

	// Initialize Tumblebug session
	// tbSess := tbclient.NewSession()

	/*
	 * [Process] Create an infrastructure
	 */
	// 1. Check if the namespace exists
	log.Debug().Msgf("Checking if the namespace exists (nsId: %s)", nsId)
	_, err = tbclient.NewSession().ReadNamespace(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to read the namespace (nsId: %s)", nsId)
		return emptyRet, err
	}

	// 2. Create a node specification (Spec)
	// * Skip: No need to regenerate node spec in namespace

	// 3. Create a node image (OS)
	// * Skip: No need to regenerate node image (OS) in namespace

	// 4. Create a virtual network (vNet)
	// Get vNet request body from the input infraModel
	vNetReq := targetInfraModel.TargetVNet
	log.Debug().Msgf("Creating a vNet (nsId: %s, vNetName: %s)", nsId, vNetReq.Name)
	log.Debug().Msgf("vNetReq: %+v", vNetReq)

	// Convert model from 'cloudmodel.VNetReq' to 'tbmodel.VNetReq'
	tbVNetReq, err := modelconv.ConvertWithValidation[cloudmodel.VNetReq, tbmodel.VNetReq](vNetReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert vNet request (nsId: %s)", nsId)
		return emptyRet, err
	}

	vNetInfo, err := tbclient.NewSession().CreateVNet(nsId, tbVNetReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create the vNet (nsId: %s)", nsId)
		return emptyRet, err
	}

	log.Debug().Msgf("vNet created: %s", vNetInfo.Id)
	// * Note: "vNetInfo.Id" should be used if any of the following steps require vNetId.

	// 5. Create a SSH key pair (sshKey)
	sshKeyReq := targetInfraModel.TargetSshKey
	log.Debug().Msgf("Creating a SSH key (nsId: %s, sshKeyName: %s)", nsId, sshKeyReq.Name)
	log.Debug().Msgf("sshKeyReq: %+v", sshKeyReq)

	// Convert model from 'cloudmodel.SshKeyReq' to 'tbmodel.SshKeyReq'
	tbSshKeyReq, err := modelconv.ConvertWithValidation[cloudmodel.SshKeyReq, tbmodel.SshKeyReq](sshKeyReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert SSH key request (nsId: %s)", nsId)
		return emptyRet, err
	}

	sshKeyInfo, err := tbclient.NewSession().CreateSshKey(nsId, tbSshKeyReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create the SSH key (nsId: %s)", nsId)
		return emptyRet, err
	}
	log.Debug().Msgf("SSH key created: %s", sshKeyInfo.Id)

	// 6. Create a security group (sg)
	// Get security group request body from the input infraModel
	sgReqList := targetInfraModel.TargetSecurityGroupList
	log.Debug().Msgf("Creating security groups (nsId: %s, sgCount: %d)", nsId, len(sgReqList))
	log.Debug().Msgf("sgReqList: %+v", sgReqList)

	sgInfoList := []tbmodel.SecurityGroupInfo{}
	for _, sgReq := range sgReqList {

		// Check if SSH access rule exists and add if missing
		sgReq = checkAndSupportSSHAccessRule(sgReq)

		// Deduplicate firewall rules before sending to Tumblebug
		if sgReq.FirewallRules != nil {
			originalCount := len(*sgReq.FirewallRules)
			dedupedRules := recommendation.DeduplicateFirewallRules(*sgReq.FirewallRules)
			sgReq.FirewallRules = &dedupedRules
			if originalCount != len(dedupedRules) {
				log.Warn().Msgf("Removed %d duplicate firewall rule(s) for SG '%s' (original: %d, deduplicated: %d)",
					originalCount-len(dedupedRules), sgReq.Name, originalCount, len(dedupedRules))
			}
		}

		// Create security group
		log.Debug().Msgf("Creating a security group (nsId: %s, sgReq.sgName: %s, sgReq.VNetId: %s, vNetInfo.vNetId: %s)",
			nsId, sgReq.Name, sgReq.VNetId, vNetInfo.Id)

		// Convert model from 'cloudmodel.SecurityGroupReq' to 'tbmodel.SecurityGroupReq'
		tbSgReq, err := modelconv.ConvertWithValidation[cloudmodel.SecurityGroupReq, tbmodel.SecurityGroupReq](sgReq)
		if err != nil {
			log.Error().Err(err).Msgf("failed to convert security group request (nsId: %s)", nsId)
			return emptyRet, err
		}

		sgInfo, err := tbclient.NewSession().CreateSecurityGroup(nsId, tbSgReq, "")
		if err != nil {
			log.Error().Err(err).Msgf("failed to create the security group (nsId: %s)", nsId)
			return emptyRet, err
		}
		log.Debug().Msgf("security group created: %s", sgInfo.Id)

		sgInfoList = append(sgInfoList, sgInfo)
	}
	log.Debug().Msgf("sgInfoList length: %d", len(sgInfoList))
	log.Debug().Msgf("sgInfoList: %+v", sgInfoList)

	// 7. Create an infrastructure (infra)
	// Get infrastructure (Infra) request body from the input infraModel
	infraReq := targetInfraModel.TargetInfra
	log.Debug().Msgf("Creating a multi-cloud infrastructure (nsId: %s, infraName: %s)", nsId, infraReq.Name)
	log.Debug().Msgf("infraReq: %+v", infraReq)

	// Convert model from 'cloudmodel.InfraReq' to 'tbmodel.InfraReq'
	tbInfraReq, err := modelconv.ConvertWithValidation[cloudmodel.InfraReq, tbmodel.InfraReq](infraReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the Infra request (nsId: %s)", nsId)
		return emptyRet, err
	}
	log.Debug().Msgf("tbInfraReq: %+v", tbInfraReq)

	// Set post-command for stable infra provisioning if a user didn't set it
	// If a user already set it, use it as is
	if len(tbInfraReq.PostCommand.Command) == 0 {
		log.Debug().Msgf("Setting default post-command `uname -a` for stable Infra provisioning (nsId: %s)", nsId)

		commands := []string{
			"uname -a",
		}
		username := "cb-user"

		tbInfraReq.PostCommand = tbmodel.InfraCmdReq{
			UserName: username,
			Command:  commands,
		}
	}

	// Create multi-cloud infrastructure
	infraInfo, err := tbclient.NewSession().CreateInfra(nsId, tbInfraReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create the multi-cloud infrastructure (nsId: %s)", nsId)

		// TODO: Consider implementing resource rollback in case of failure at this step (e.g., delete created vNet, SSH key, security groups)
		// ! But first, be cautious about the rollback since it may cause unintended consequences if not implemented properly (e.g., deleting resources that are shared with other infrastructures or used by other applications)
		// ? Second, consider the trade-off between keeping the failed infrastructure for troubleshooting and rolling back the created resources, which is more beneficial for users in case of failure at this step

		return emptyRet, err
	}
	log.Debug().Msgf("multi-cloud infrastructure created: %s", infraInfo.Id)

	/*
	 * [Output] Return the created multi-cloud infrastructure info
	 */

	// Convert the response model from 'tbmodel.InfraInfo' to 'cloudmodel.InfraInfo'
	infraInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.InfraInfo, cloudmodel.InfraInfo](infraInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s)", nsId)
		return emptyRet, err
	}
	var temp cloudmodel.VmInfraInfo
	temp.InfraInfo = infraInfoConverted

	log.Debug().Msgf("Stabilizing newly created infrastructure (nsId: %s, infraName: %s)...", nsId, infraInfoConverted.Name)
	time.Sleep(5 * time.Second)

	// Option A: Check SSH readiness for IBM Cloud
	if strings.ToLower(targetInfraModel.TargetCloud.Csp) == "ibm" {
		log.Info().Msgf("IBM Cloud detected - performing SSH readiness check (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)

		const maxWaitTime = 3 * time.Minute
		const checkInterval = 10 * time.Second

		if _, err := CheckSSHReadinessWithDetails(nsId, infraInfoConverted.Id, maxWaitTime, checkInterval); err != nil {
			log.Warn().Err(err).Msgf("SSH readiness check incomplete - Nodes may need additional time (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)
			// Warning only - migration itself succeeded
		} else {
			log.Info().Msgf("SSH readiness confirmed - Nodes are accessible (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)
		}
	}

	log.Info().Msgf("Infrastructure created successfully (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)
	return temp, nil
}

// CreateInfraWithExisting creates an infrastructure by reusing/ensuring existing resources (useExisting=true)
func CreateInfraWithExisting(nsId string, targetInfraModel *cloudmodel.RecommendedInfra) (cloudmodel.VmInfraInfo, error) {
	log.Info().Msg("Creating infrastructure with existing resources")
	emptyRet := cloudmodel.VmInfraInfo{}

	/*
	 * [Input] Receive and validate the target infrastructure model
	 */
	err := validation.ValidateTargetInfra(nsId, targetInfraModel, true).Err()
	if err != nil {
		log.Error().Err(err).Msgf("failed to validate the target infrastructure model (nsId: %s)", nsId)
		return emptyRet, err
	}
	log.Info().Msgf("the target infrastructure model is valid (nsId: %s)", nsId)

	// Preflight: resolve the latest CSP image and confirm available system disk per nodegroup.
	err = preflightCheckCspProvisioning(nsId, targetInfraModel.TargetInfra.NodeGroups)
	if err != nil {
		log.Error().Err(err).Msgf("failed to run preflight check for CSP provisioning (nsId: %s)", nsId)
		return emptyRet, err
	}

	/*
	 * [Process] Create an infrastructure
	 */
	// 1. Check if the namespace exists
	log.Debug().Msgf("Checking if the namespace exists (nsId: %s)", nsId)
	_, err = tbclient.NewSession().ReadNamespace(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to read the namespace (nsId: %s)", nsId)
		return emptyRet, err
	}

	// 2. Create a node specification (spec)
	// * Skip: No need to regenerate node spec in namespace

	// 3. Create a node image (OS)
	// * Skip: No need to regenerate node image (OS) in namespace

	// 4. Use/Create virtual networks (vNet, Subnets)
	netRequirements := validation.DeriveNetworkRequirements(targetInfraModel.TargetInfra.NodeGroups)
	for _, netRequirement := range netRequirements {
		err = useOrCreateNetwork(nsId, netRequirement, targetInfraModel.TargetVNet)
		if err != nil {
			log.Error().Err(err).Msgf("failed to use or create virtual network %s (nsId: %s)", netRequirement.VNetId, nsId)
			return emptyRet, err
		}
	}

	// 5. Use/Create SSH key pairs (sshKey)
	sshKeyRequirements := validation.DeriveSshKeyRequirements(targetInfraModel.TargetInfra.NodeGroups)
	for _, sshKeyRequirement := range sshKeyRequirements {
		err = useOrCreateSshKey(nsId, sshKeyRequirement, targetInfraModel.TargetSshKey)
		if err != nil {
			log.Error().Err(err).Msgf("failed to use or create SSH key %s (nsId: %s)", sshKeyRequirement.SshKeyId, nsId)
			return emptyRet, err
		}
	}

	// 6. Use/Create security groups (sg)
	sgRequirements := validation.DeriveSecurityGroupRequirements(targetInfraModel.TargetInfra.NodeGroups)
	for _, sgRequirement := range sgRequirements {
		err = useOrCreateSecurityGroup(nsId, sgRequirement, targetInfraModel.TargetSecurityGroupList)
		if err != nil {
			log.Error().Err(err).Msgf("failed to use or create security group %s (nsId: %s)", sgRequirement.SecurityGroupId, nsId)
			return emptyRet, err
		}
	}

	// 7. Create an infrastructure (Infra)
	infraReq := targetInfraModel.TargetInfra
	log.Debug().Msgf("Creating an infrastructure (nsId: %s, infraName: %s)", nsId, infraReq.Name)
	tbInfraReq, err := modelconv.ConvertWithValidation[cloudmodel.InfraReq, tbmodel.InfraReq](infraReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the Infra request (nsId: %s)", nsId)
		return emptyRet, err
	}

	if len(tbInfraReq.PostCommand.Command) == 0 {
		log.Debug().Msgf("Setting default post-command `uname -a` for stable Infra provisioning (nsId: %s)", nsId)
		tbInfraReq.PostCommand = tbmodel.InfraCmdReq{
			UserName: "cb-user",
			Command:  []string{"uname -a"},
		}
	}

	infraInfo, err := tbclient.NewSession().CreateInfra(nsId, tbInfraReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create the infrastructure (nsId: %s)", nsId)
		return emptyRet, err
	}
	log.Debug().Msgf("infrastructure created: %s", infraInfo.Id)

	infraInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.InfraInfo, cloudmodel.InfraInfo](infraInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s)", nsId)
		return emptyRet, err
	}

	var temp cloudmodel.VmInfraInfo
	temp.InfraInfo = infraInfoConverted
	log.Debug().Msgf("Stabilizing newly created infrastructure (nsId: %s, infraName: %s)...", nsId, infraInfoConverted.Name)
	time.Sleep(5 * time.Second)

	// Option A: Check SSH readiness for IBM Cloud
	if strings.ToLower(targetInfraModel.TargetCloud.Csp) == "ibm" {
		log.Info().Msgf("IBM Cloud detected - performing SSH readiness check (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)

		const maxWaitTime = 3 * time.Minute
		const checkInterval = 10 * time.Second

		if _, err := CheckSSHReadinessWithDetails(nsId, infraInfoConverted.Id, maxWaitTime, checkInterval); err != nil {
			log.Warn().Err(err).Msgf("SSH readiness check incomplete - Nodes may need additional time (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)
			// Warning only - migration itself succeeded
		} else {
			log.Info().Msgf("SSH readiness confirmed - Nodes are accessible (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)
		}
	}

	log.Info().Msgf("Infrastructure created successfully (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)
	return temp, nil
}

// List all migrated infrastructures
func ListAllInfraInfo(nsId string) (cloudmodel.InfraInfoList, error) {
	log.Info().Msg("Listing all migrated infrastructures")

	var emptyRet cloudmodel.InfraInfoList
	// var infraInfoList cloudmodel.InfraInfoList

	// Initialize Tumblebug session
	tbSess := tbclient.NewSession()

	infraInfoList, err := tbSess.ReadAllInfra(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to retrieve all migrated infrastructures (nsId: %s)", nsId)
		return emptyRet, err
	}

	// Convert the response model from 'tbclient.TbInfraInfoList' to 'cloudmodel.InfraInfoList'
	convertedInfraInfoList, err := modelconv.ConvertWithValidation[tbclient.TbInfraInfoList, cloudmodel.InfraInfoList](infraInfoList)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info list (nsId: %s)", nsId)
		return emptyRet, err
	}

	log.Info().Msgf("Retrieved all migrated infrastructures (nsId: %s, count: %d) successfully", nsId, len(convertedInfraInfoList.Infra))
	return convertedInfraInfoList, nil
}

// Get all migrated infrastructures
func ListInfraIDs(nsId string, option string) (cloudmodel.IdList, error) {
	log.Info().Msg("Listing all migrated infrastructure IDs")

	var emptyRet cloudmodel.IdList
	var idList cloudmodel.IdList
	idList.IdList = make([]string, 0)

	/*
	 * Validate the input
	 */

	if option != "id" {
		log.Error().Msgf("invalid option: %s", option)
		return emptyRet, fmt.Errorf("invalid option: %s", option)
	}

	// Initialize Tumblebug session
	tbSess := tbclient.NewSession()
	infraIdList, err := tbSess.ReadInfraIDs(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get the infrastructure IDs (nsId: %s)", nsId)
		return emptyRet, err
	}

	// Return the result
	idList.IdList = append(idList.IdList, infraIdList.IdList...)

	log.Info().Msgf("Retrieved all migrated infrastructure IDs (nsId: %s, count: %d) successfully", nsId, len(idList.IdList))
	return idList, nil
}

// Get the migrated infrastructure
func GetInfra(nsId, infraId string) (cloudmodel.InfraInfo, error) {
	log.Info().Msgf("Retrieving the migrated infrastructure (nsId: %s, infraId: %s)", nsId, infraId)

	// Initialize Tumblebug session
	tbSess := tbclient.NewSession()
	infraInfo, err := tbSess.ReadInfra(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get the infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return cloudmodel.InfraInfo{}, err
	}

	// Convert the response model from 'tbmodel.InfraInfo' to 'cloudmodel.InfraInfo'
	convertedInfraInfo, err := modelconv.ConvertWithValidation[tbmodel.InfraInfo, cloudmodel.InfraInfo](infraInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return cloudmodel.InfraInfo{}, err
	}

	log.Info().Msgf("Retrieved the migrated infrastructure (nsId: %s, infraId: %s) successfully", nsId, infraId)
	return convertedInfraInfo, nil
}

// Delete the migrated infrastructure
func DeleteInfra(nsId, infraId, option string) (common.SimpleMsg, error) {
	log.Info().Msgf("Deleting the migrated infrastructure (nsId: %s, infraId: %s)", nsId, infraId)

	// 1. Read Infra info. This call is paced by the client-side TB rate limiter
	// (pkg/client/tumblebug); deletion tolerates a longer wait than interactive reads.
	readCtx, cancelRead := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelRead()
	infraInfo, err := tbclient.NewSession().SetContext(readCtx).ReadInfra(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to read the infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return common.SimpleMsg{}, err
	}

	// 2. Delete Infra
	idList, err := tbclient.NewSession().DeleteInfra(nsId, infraId, option)
	if err != nil {
		log.Error().Err(err).Msgf("failed to delete the infrastructure (nsId: %s, infraId: %s)", nsId, infraId)
		return common.SimpleMsg{}, err
	}
	log.Debug().Msgf("Infra deleted (nsId: %s, infraId: %s, IdList: %s)", nsId, infraId, idList.IdList)

	// Sleep for a while to ensure previous deletions are completed
	log.Debug().Msgf("Waiting for CSP to complete Infra deletion (nsId: %s)", nsId)
	time.Sleep(2 * time.Second)

	// 3. Delete security groups
	// Collect unique security group IDs from all Nodes
	sgIdMap := make(map[string]struct{})
	for _, node := range infraInfo.Node {
		for _, sgId := range node.SecurityGroupIds {
			sgIdMap[sgId] = struct{}{}
		}
	}
	log.Debug().Msgf("Deleting security groups (nsId: %s, SGs: %v)", nsId, sgIdMap)

	// Delete all security groups
	for sgId := range sgIdMap {
		msg, err := tbclient.NewSession().DeleteSecurityGroup(nsId, sgId)
		if err != nil {
			log.Error().Err(err).Msgf("failed to delete security group (nsId: %s, sgId: %s)", nsId, sgId)
			// Continue deleting other resources even if this fails
		} else {
			log.Debug().Msgf("Security group deleted (nsId: %s, sgId: %s, msg: %s)", nsId, sgId, msg)
		}
	}

	// Sleep for a while to ensure previous deletions are completed
	log.Debug().Msgf("Waiting for CSP to complete security group deletion (nsId: %s)", nsId)
	time.Sleep(2 * time.Second)

	// 4. Delete SSH Key
	// Collect unique SSH Key IDs from all Nodes
	sshKeyIdMap := make(map[string]struct{})
	for _, node := range infraInfo.Node {
		sshKeyIdMap[node.SshKeyId] = struct{}{}
	}
	log.Debug().Msgf("Deleting SSH keys (nsId: %s, sshKeys: %v)", nsId, sshKeyIdMap)

	// Delete all SSH Key
	for sshKeyId := range sshKeyIdMap {
		// Delete SSH Key
		log.Debug().Msgf("Deleting SSH key (nsId: %s, sshKeyId: %s)", nsId, sshKeyId)
		msg, err := tbclient.NewSession().DeleteSshKey(nsId, sshKeyId)
		if err != nil {
			log.Error().Err(err).Msgf("failed to delete SSH key (nsId: %s, sshKeyId: %s)", nsId, sshKeyId)
			// Continue deleting other resources even if this fails
		} else {
			log.Debug().Msgf("SSH key deleted (nsId: %s, sshKeyId: %s, msg: %s)", nsId, sshKeyId, msg)
		}
	}

	// Sleep for a while to ensure previous deletions are completed
	log.Debug().Msgf("Waiting for CSP to complete SSH key deletion (nsId: %s)", nsId)
	time.Sleep(2 * time.Second)

	// 5. Delete vNets
	// Collect unique vNet IDs from all Nodes
	vNetIdMap := make(map[string]struct{})
	for _, node := range infraInfo.Node {
		vNetIdMap[node.VNetId] = struct{}{}
	}
	log.Debug().Msgf("Deleting VNets (nsId: %s, vNets: %v)", nsId, vNetIdMap)

	// Delete all vNet

	const vNetDeleteMaxRetries = 5
	const vNetDeleteRetryInterval = 6 * time.Second

	for vNetId := range vNetIdMap {
		var deleteErr error
		for attempt := 1; attempt <= vNetDeleteMaxRetries; attempt++ {
			log.Debug().Msgf("Deleting VNet (nsId: %s, vNetId: %s, attempt: %d/%d)",
				nsId, vNetId, attempt, vNetDeleteMaxRetries)
			msg, err := tbclient.NewSession().DeleteVNet(nsId, vNetId, "withsubnets")
			if err == nil {
				log.Debug().Msgf("VNet deleted (nsId: %s, vNetId: %s, msg: %s)", nsId, vNetId, msg)
				deleteErr = nil
				break
			}
			deleteErr = err
			if attempt < vNetDeleteMaxRetries {
				log.Warn().Err(err).Msgf("VNet deletion failed (nsId: %s, vNetId: %s, attempt: %d/%d) — "+
					"CSP may still be releasing subnet dependencies. Retrying in %s...",
					nsId, vNetId, attempt, vNetDeleteMaxRetries, vNetDeleteRetryInterval)
				time.Sleep(vNetDeleteRetryInterval)
			}
		}
		if deleteErr != nil {
			log.Error().Err(deleteErr).Msgf("failed to delete VNet after %d attempts (nsId: %s, vNetId: %s)",
				vNetDeleteMaxRetries, nsId, vNetId)
		}
	}

	log.Debug().Msgf("VNet deletion completed (nsId: %s)", nsId)

	/*
	 * [Output] Return the result
	 */

	ret := common.SimpleMsg{
		Message: fmt.Sprintf("Infrastructure and resources deleted successfully (nsId: %s, infraId: %s)", nsId, infraId),
	}
	log.Info().Msgf("Infrastructure deletion completed (nsId: %s, infraId: %s)", nsId, infraId)
	return ret, nil
}

// preflightCheckCspProvisioning resolves the latest CSP image and confirms available system disk per nodegroup
func preflightCheckCspProvisioning(nsId string, nodeGroups []cloudmodel.CreateNodeGroupReq) error {
	log.Info().Msgf("running preflight check for all nodegroups (nsId: %s)", nsId)
	for i := range nodeGroups {
		ng := &nodeGroups[i]
		precheck, reviewErr := recommendation.PreflightCheckCspProvisioning(
			ng.SpecId, ng.ImageId, ng.CspImageName, ng.RootDiskType,
		)
		if reviewErr != nil {
			log.Warn().Err(reviewErr).Msgf("preflight check failed for nodegroup %s (specId: %s, imageId: %s); proceeding with cached image",
				ng.Name, ng.SpecId, ng.ImageId)
			continue
		}
		if !precheck.IsAvailable {
			return fmt.Errorf("image %s is not available for nodegroup %s (specId: %s); aborting migration",
				ng.ImageId, ng.Name, ng.SpecId)
		}
		if precheck.ResolvedCspImageName != ng.CspImageName {
			log.Info().Msgf("nodegroup %s: CspImageName resolved from %q to %q", ng.Name, ng.CspImageName, precheck.ResolvedCspImageName)
			ng.CspImageName = precheck.ResolvedCspImageName
		}
		if precheck.SuggestedSystemDisk != "" && ng.RootDiskType != precheck.SuggestedSystemDisk {
			log.Info().Msgf("nodegroup %s: RootDiskType updated from %q to suggested %q", ng.Name, ng.RootDiskType, precheck.SuggestedSystemDisk)
			ng.RootDiskType = precheck.SuggestedSystemDisk
		}
	}
	log.Info().Msgf("spec-image pair preflight check passed (nsId: %s)", nsId)
	return nil
}

// useOrCreateNetwork checks if VNet and required subnets exist, and creates them from the creation request if missing
func useOrCreateNetwork(nsId string, netRequirement validation.NetworkRequirement, vNetCreationReq cloudmodel.VNetReq) error {
	needsCreate, issue := validation.CheckNetworkAvailability(nsId, netRequirement, vNetCreationReq)
	if issue != nil {
		return fmt.Errorf("%s", issue.Message)
	}
	if !needsCreate {
		log.Info().Msgf("vNet %s and all required subnets already exist. CM-Beetle will reuse it.", netRequirement.VNetId)
		return nil
	}

	vNetReq := vNetCreationReq
	vNetReq.Name = netRequirement.VNetId
	if netRequirement.ConnectionName != "" {
		vNetReq.ConnectionName = netRequirement.ConnectionName
	}

	var newSubnetList []cloudmodel.SubnetReq
	for idx, subnetName := range netRequirement.SubnetIds {
		var subReq cloudmodel.SubnetReq
		if idx < len(vNetCreationReq.SubnetInfoList) {
			subReq = vNetCreationReq.SubnetInfoList[idx]
		} else if len(vNetCreationReq.SubnetInfoList) > 0 {
			subReq = vNetCreationReq.SubnetInfoList[0]
		}
		subReq.Name = subnetName
		newSubnetList = append(newSubnetList, subReq)
	}
	if len(newSubnetList) > 0 {
		vNetReq.SubnetInfoList = newSubnetList
	}

	log.Debug().Msgf("Creating a vNet (nsId: %s, vNetName: %s)", nsId, vNetReq.Name)
	tbVNetReq, err := modelconv.ConvertWithValidation[cloudmodel.VNetReq, tbmodel.VNetReq](vNetReq)
	if err != nil {
		return err
	}

	_, err = tbclient.NewSession().CreateVNet(nsId, tbVNetReq)
	if err != nil {
		return err
	}

	log.Debug().Msgf("vNet created: %s", vNetReq.Name)
	return nil
}

// useOrCreateSshKey checks if SSH key exists, and creates it from the creation request if missing
func useOrCreateSshKey(nsId string, sshKeyRequirement validation.SshKeyRequirement, sshKeyCreationReq cloudmodel.SshKeyReq) error {
	needsCreate, issue := validation.CheckSshKeyAvailability(nsId, sshKeyRequirement, sshKeyCreationReq)
	if issue != nil {
		return fmt.Errorf("%s", issue.Message)
	}
	if !needsCreate {
		log.Info().Msgf("SSH key %s already exists. CM-Beetle will reuse it.", sshKeyRequirement.SshKeyId)
		return nil
	}

	req := sshKeyCreationReq
	req.Name = sshKeyRequirement.SshKeyId
	if sshKeyRequirement.ConnectionName != "" {
		req.ConnectionName = sshKeyRequirement.ConnectionName
	}

	log.Debug().Msgf("Creating a SSH key (nsId: %s, sshKeyName: %s)", nsId, req.Name)
	tbSshKeyReq, err := modelconv.ConvertWithValidation[cloudmodel.SshKeyReq, tbmodel.SshKeyReq](req)
	if err != nil {
		return err
	}

	_, err = tbclient.NewSession().CreateSshKey(nsId, tbSshKeyReq)
	if err != nil {
		return err
	}
	log.Debug().Msgf("SSH key created: %s", req.Name)
	return nil
}

// useOrCreateSecurityGroup checks if security group exists, and creates it from the creation request list if missing
func useOrCreateSecurityGroup(nsId string, sgRequirement validation.SecurityGroupRequirement, sgCreationReqList []cloudmodel.SecurityGroupReq) error {
	needsCreate, issue := validation.CheckSecurityGroupAvailability(nsId, sgRequirement, sgCreationReqList)
	if issue != nil {
		return fmt.Errorf("%s", issue.Message)
	}
	if !needsCreate {
		log.Info().Msgf("Security group %s already exists. CM-Beetle will reuse it.", sgRequirement.SecurityGroupId)
		return nil
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

	if sgCreationReq.ConnectionName == "" && sgRequirement.ConnectionName != "" {
		sgCreationReq.ConnectionName = sgRequirement.ConnectionName
	}
	if sgCreationReq.VNetId == "" && sgRequirement.VNetId != "" {
		sgCreationReq.VNetId = sgRequirement.VNetId
	}

	sgCreationReq = checkAndSupportSSHAccessRule(sgCreationReq)

	// Deduplicate firewall rules before sending to Tumblebug
	if sgCreationReq.FirewallRules != nil {
		originalCount := len(*sgCreationReq.FirewallRules)
		dedupedRules := recommendation.DeduplicateFirewallRules(*sgCreationReq.FirewallRules)
		sgCreationReq.FirewallRules = &dedupedRules
		if originalCount != len(dedupedRules) {
			log.Warn().Msgf("Removed %d duplicate firewall rule(s) for SG '%s' during fallback creation (original: %d, deduplicated: %d)",
				originalCount-len(dedupedRules), sgCreationReq.Name, originalCount, len(dedupedRules))
		}
	}

	log.Debug().Msgf("Creating a security group (nsId: %s, sgName: %s, VNetId: %s)", nsId, sgCreationReq.Name, sgCreationReq.VNetId)
	tbSgReq, err := modelconv.ConvertWithValidation[cloudmodel.SecurityGroupReq, tbmodel.SecurityGroupReq](sgCreationReq)
	if err != nil {
		return err
	}

	_, err = tbclient.NewSession().CreateSecurityGroup(nsId, tbSgReq, "")
	if err != nil {
		return err
	}
	log.Debug().Msgf("security group created: %s", sgCreationReq.Name)
	return nil
}

// checkAndSupportSSHAccessRule checks if SSH access rule exists in the security group and adds it if missing
// This function provides SSH connectivity support during migration phase
func checkAndSupportSSHAccessRule(sgReq cloudmodel.SecurityGroupReq) cloudmodel.SecurityGroupReq {
	// Check if FirewallRules is nil
	if sgReq.FirewallRules == nil {
		log.Warn().Msgf("Security group '%s' has no firewall rules defined, adding SSH access rule for remote management", sgReq.Name)

		sshRule := cloudmodel.FirewallRuleReq{
			Direction: "inbound",
			Protocol:  "tcp",
			CIDR:      "0.0.0.0/0",
			Ports:     "22",
		}

		rules := []cloudmodel.FirewallRuleReq{sshRule}
		sgReq.FirewallRules = &rules

		return sgReq
	}

	// Check if SSH rule exists in the firewall rules
	hasSSHRule := containsSSHRuleInMigration(*sgReq.FirewallRules)

	if !hasSSHRule {
		log.Warn().Msgf("Security group '%s' does not have SSH access rule from 0.0.0.0/0, adding SSH access rule for remote management", sgReq.Name)

		sshRule := cloudmodel.FirewallRuleReq{
			Direction: "inbound",
			Protocol:  "tcp",
			CIDR:      "0.0.0.0/0",
			Ports:     "22",
		}

		// Add SSH rule to existing rules
		*sgReq.FirewallRules = append(*sgReq.FirewallRules, sshRule)
	} else {
		log.Debug().Msgf("Security group '%s' already has SSH access rule from 0.0.0.0/0", sgReq.Name)
	}

	return sgReq
}

// containsSSHRuleInMigration checks if the security group rules contain an SSH access rule from 0.0.0.0/0
// This function is specifically used during migration phase
func containsSSHRuleInMigration(rules []cloudmodel.FirewallRuleReq) bool {
	for _, rule := range rules {
		// Must be inbound TCP rule from 0.0.0.0/0
		if rule.Direction != "inbound" || (rule.Protocol != "tcp" && rule.Protocol != "TCP") {
			continue
		}

		// Must allow access from anywhere (0.0.0.0/0)
		if rule.CIDR != "0.0.0.0/0" {
			continue
		}

		// Check if port 22 is covered by this rule
		if isSSHPortCoveredInMigration(rule.Ports) {
			log.Debug().Msgf("SSH rule found during migration: protocol=%s, direction=%s, ports=%s, cidr=%s",
				rule.Protocol, rule.Direction, rule.Ports, rule.CIDR)
			return true
		}
	}
	return false
}

// isSSHPortCoveredInMigration checks if port 22 is covered by the given port specification
// Handles three port formats: single port (22), comma-separated ports (22,23,24), port range (22-24)
func isSSHPortCoveredInMigration(portSpec string) bool {
	if portSpec == "" {
		return false
	}

	portSpec = strings.TrimSpace(portSpec)

	// Case 1: Single port (22)
	if !strings.Contains(portSpec, ",") && !strings.Contains(portSpec, "-") {
		return portSpec == "22"
	}

	// Case 2: Comma-separated ports (22,23,24)
	if strings.Contains(portSpec, ",") {
		ports := strings.Split(portSpec, ",")
		for _, port := range ports {
			port = strings.TrimSpace(port)
			if port == "22" {
				log.Debug().Msgf("SSH port 22 found in comma-separated ports during migration: %s", portSpec)
				return true
			}
		}
		return false
	}

	// Case 3: Port range (22-24)
	if strings.Contains(portSpec, "-") {
		parts := strings.Split(portSpec, "-")
		if len(parts) != 2 {
			log.Warn().Msgf("Invalid port range format during migration: %s", portSpec)
			return false
		}

		startPort, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		endPort, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

		if err1 != nil || err2 != nil {
			log.Warn().Msgf("Invalid port range format - non-numeric values during migration: %s", portSpec)
			return false
		}

		if startPort <= 22 && 22 <= endPort {
			log.Debug().Msgf("SSH port 22 found in port range during migration: %s", portSpec)
			return true
		}
	}

	return false
}

// NodeSSHStatusDetail represents the SSH readiness status of a node (internal use)
type NodeSSHStatusDetail struct {
	ID        string
	Name      string
	PublicIP  string
	PrivateIP string
	Username  string
	Status    string
	SSHReady  bool
	SSHPort   int
	Error     string
}

// CheckSSHReadinessWithDetails checks if all Nodes in the infrastructure are SSH-accessible
// and returns detailed status for each node. It works for any CSP; "Running" status doesn't mean
// cloud-init (SSH user setup) is done (IBM Cloud VPC: up to ~3 min in testing), so this polls
// until ready or maxWaitTime elapses.
//
// Parameters:
//   - nsId: Namespace ID
//   - infraId: Infrastructure ID
//   - maxWaitTime: Maximum time to wait (e.g., 3*time.Minute)
//   - checkInterval: Interval between checks (e.g., 10*time.Second)
//
// Returns:
//   - []NodeSSHStatusDetail: Detailed status for each node
//   - error: nil if all Nodes are SSH-ready, or error describing the issue
func CheckSSHReadinessWithDetails(nsId string, infraId string, maxWaitTime time.Duration, checkInterval time.Duration) ([]NodeSSHStatusDetail, error) {
	log.Info().Msgf("Starting SSH readiness check (nsId: %s, infraId: %s, maxWait: %v, interval: %v)",
		nsId, infraId, maxWaitTime, checkInterval)

	deadline := time.Now().Add(maxWaitTime)
	attempt := 0
	maxAttempts := int(maxWaitTime / checkInterval)
	var nodeStatusList []NodeSSHStatusDetail

	// Nodes already confirmed SSH-ready are never probed again, so Tumblebug doesn't
	// keep re-running a remote command against Nodes that have already responded.
	confirmedReady := make(map[string]bool)

	for time.Now().Before(deadline) {
		attempt++
		log.Debug().Msgf("SSH readiness check attempt %d/%d (nsId: %s, infraId: %s)",
			attempt, maxAttempts, nsId, infraId)

		// Get current infra status from Tumblebug
		infraInfo, err := tbclient.NewSession().ReadInfra(nsId, infraId)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to read infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
			return nil, fmt.Errorf("failed to read infrastructure info: %w", err)
		}

		// Build status list for each node
		nodeStatusList = make([]NodeSSHStatusDetail, 0, len(infraInfo.Node))
		readyNodes := 0

		for _, node := range infraInfo.Node {
			nodeStatus := NodeSSHStatusDetail{
				ID:        node.Id,
				Name:      node.Name,
				PublicIP:  node.PublicIP,
				PrivateIP: node.PrivateIP,
				Username:  node.NodeUserName,
				Status:    node.Status,
				SSHPort:   22,
			}

			switch {
			case strings.ToLower(node.Status) != "running":
				nodeStatus.Error = fmt.Sprintf("Node not running yet (status: %s)", node.Status)
			case confirmedReady[node.Id]:
				nodeStatus.SSHReady = true
				readyNodes++
			case probeNodeSSHReachability(nsId, infraId, node.Id):
				nodeStatus.SSHReady = true
				confirmedReady[node.Id] = true
				readyNodes++
			default:
				nodeStatus.Error = "SSH not reachable yet via Tumblebug remote command"
			}

			nodeStatusList = append(nodeStatusList, nodeStatus)
		}

		// Check if all Nodes are ready
		totalNodes := len(infraInfo.Node)
		if totalNodes > 0 && readyNodes == totalNodes {
			log.Info().Msgf("All Nodes are SSH-ready (%d/%d) (nsId: %s, infraId: %s)",
				readyNodes, totalNodes, nsId, infraId)
			return nodeStatusList, nil
		}

		log.Debug().Msgf("SSH readiness: %d/%d Nodes ready (nsId: %s, infraId: %s)",
			readyNodes, totalNodes, nsId, infraId)

		// Wait before next attempt (unless this is the last attempt)
		if attempt < maxAttempts && time.Now().Add(checkInterval).Before(deadline) {
			time.Sleep(checkInterval)
		}
	}

	return nodeStatusList, fmt.Errorf("SSH readiness timeout after %v: not all Nodes became accessible", maxWaitTime)
}

// probeNodeSSHReachability checks SSH reachability for a single Node by running a lightweight
// command on it via Tumblebug's remote command API, scoped to that Node only (nodeId), instead of
// connecting to the node directly or targeting the whole Infra (which would make Tumblebug re-run
// the command against every Node, including ones already confirmed ready or not running yet).
func probeNodeSSHReachability(nsId, infraId, nodeId string) bool {
	cmdReq := tbmodel.InfraCmdReq{Command: []string{"echo ready"}, TimeoutMinutes: 1}

	result, err := tbclient.NewSession().RemoteCommandToInfra(nsId, infraId, "", nodeId, cmdReq)
	if err != nil {
		log.Debug().Err(err).Msgf("SSH reachability probe failed (nsId: %s, infraId: %s, nodeId: %s)", nsId, infraId, nodeId)
		return false
	}

	for _, r := range result.Results {
		if r.NodeId == nodeId && r.Error == "" {
			return true
		}
	}
	return false
}
