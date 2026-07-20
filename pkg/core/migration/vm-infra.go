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
	"fmt"
	"strconv"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
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

// CreateVMInfraWithDefaults Create a VM infrastructure with defaults for the computing infra migration
func CreateVMInfraWithDefaults(nsId string, infraModel *cloudmodel.InfraDynamicReq) (cloudmodel.VmInfraInfo, error) {
	log.Info().Msg("Creating VM infrastructure with defaults")

	// Convert the request model from 'cloudmodel.InfraDynamicReq' to 'tbmodel.InfraDynamicReq'
	infraModelConverted, err := modelconv.ConvertWithValidation[cloudmodel.InfraDynamicReq, tbmodel.InfraDynamicReq](*infraModel)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure model (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	vmInfraInfo, err := tbclient.NewSession().CreateInfraDynamic(nsId, infraModelConverted)
	if err != nil {
		log.Error().Err(err).Msgf("failed to migrate the infrastructure (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	// Convert the response model from 'tbmodel.InfraInfo' to 'cloudmodel.VmInfraInfo'
	convertedVmInfraInfo, err := modelconv.ConvertWithValidation[tbmodel.InfraInfo, cloudmodel.VmInfraInfo](vmInfraInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	log.Info().Msgf("VM infrastructure created successfully (nsId: %s, infraName: %s)", nsId, convertedVmInfraInfo.Name)

	return convertedVmInfraInfo, nil
}

// CreateInfra creates a VM infrastructure for the computing infra migration by creating fresh resources (useExisting=false)
func CreateInfra(nsId string, targetInfraModel *cloudmodel.RecommendedInfra) (cloudmodel.VmInfraInfo, error) {
	log.Info().Msg("Creating VM infrastructure")

	emptyRet := cloudmodel.VmInfraInfo{}

	/*
	 * [Input] Receive and validate the target infrastructure model
	 */

	err := validateTargeInfraModel(nsId, targetInfraModel)
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
	 * [Process] Create a VM infrastructure
	 */
	// 1. Check if the namespace exists
	log.Debug().Msgf("Checking if the namespace exists (nsId: %s)", nsId)
	_, err = tbclient.NewSession().ReadNamespace(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to read the namespace (nsId: %s)", nsId)
		return emptyRet, err
	}

	log.Debug().Msgf("Checking if the Infra (%s) exists in the namespace (%s)", targetInfraModel.TargetInfra.Name, nsId)
	tempInfraInfo, err := tbclient.NewSession().ReadInfra(nsId, targetInfraModel.TargetInfra.Name)
	if tempInfraInfo.Id != "" {
		log.Error().Err(err).Msgf("the Infra already exist (nsId: %s, infraName: %s)", nsId, targetInfraModel.TargetInfra.Name)
		return emptyRet, err
	}

	// 2. Create a VM specification (vmSpec)
	// * Skip: No need to regenerate vmSpec in namespace

	// 3. Create a VM OS image (vmOsImage)
	// * Skip: No need to regenerate vmOsImage in namespace

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
			log.Error().Err(err).Msgf("failed to convert SSH key request (nsId: %s)", nsId)
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

	// 7. Create a VM infrastructure (i.e., Infra)
	// Get multi-cloud infrastructure (Infra) request body from the input infraModel
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

	log.Info().Msgf("VM infrastructure created successfully (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)
	return temp, nil
}

// CreateInfraWithExisting creates a VM infrastructure by reusing/ensuring existing resources (useExisting=true)
func CreateInfraWithExisting(nsId string, targetInfraModel *cloudmodel.RecommendedInfra) (cloudmodel.VmInfraInfo, error) {
	log.Info().Msg("Creating VM infrastructure with existing resources")
	emptyRet := cloudmodel.VmInfraInfo{}

	/*
	 * [Input] Receive and validate the target infrastructure model
	 */
	err := validateTargeInfraModelWithExisting(nsId, targetInfraModel)
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
	 * [Process] Create a VM infrastructure
	 */
	// 1. Check if the namespace exists
	log.Debug().Msgf("Checking if the namespace exists (nsId: %s)", nsId)
	_, err = tbclient.NewSession().ReadNamespace(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to read the namespace (nsId: %s)", nsId)
		return emptyRet, err
	}

	log.Debug().Msgf("Checking if the Infra (%s) exists in the namespace (%s)", targetInfraModel.TargetInfra.Name, nsId)
	tempInfraInfo, err := tbclient.NewSession().ReadInfra(nsId, targetInfraModel.TargetInfra.Name)
	if tempInfraInfo.Id != "" {
		log.Error().Err(err).Msgf("the Infra already exist (nsId: %s, infraName: %s)", nsId, targetInfraModel.TargetInfra.Name)
		return emptyRet, err
	}

	// 2. Create a VM specification (vmSpec)
	// * Skip: No need to regenerate vmSpec in namespace

	// 3. Create a VM OS image (vmOsImage)
	// * Skip: No need to regenerate vmOsImage in namespace

	// 4. Use/Create virtual networks (vNet, Subnets)
	netReqs := deriveNetworkIds(targetInfraModel.TargetInfra.NodeGroups)
	for _, netReq := range netReqs {
		err = useOrCreateNetwork(nsId, netReq, targetInfraModel.TargetVNet)
		if err != nil {
			log.Error().Err(err).Msgf("failed to use or create virtual network %s (nsId: %s)", netReq.VNetId, nsId)
			return emptyRet, err
		}
	}

	// 5. Use/Create SSH key pairs (sshKey)
	sshKeyReqs := deriveSshKeyIds(targetInfraModel.TargetInfra.NodeGroups)
	for _, sshKeyReq := range sshKeyReqs {
		err = useOrCreateSshKey(nsId, sshKeyReq, targetInfraModel.TargetSshKey)
		if err != nil {
			log.Error().Err(err).Msgf("failed to use or create SSH key %s (nsId: %s)", sshKeyReq.SshKeyId, nsId)
			return emptyRet, err
		}
	}

	// 6. Use/Create security groups (sg)
	sgReqs := deriveSecurityGroupIds(targetInfraModel.TargetInfra.NodeGroups)
	for _, sgReq := range sgReqs {
		err = useOrCreateSecurityGroup(nsId, sgReq, targetInfraModel.TargetSecurityGroupList)
		if err != nil {
			log.Error().Err(err).Msgf("failed to use or create security group %s (nsId: %s)", sgReq.SecurityGroupId, nsId)
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
	log.Info().Msgf("VM infrastructure created successfully (nsId: %s, infraName: %s)", nsId, infraInfoConverted.Name)
	return temp, nil
}

// List all migrated VM infrastructures
func ListAllVMInfraInfo(nsId string) (cloudmodel.InfraInfoList, error) {
	log.Info().Msg("Listing all migrated VM infrastructures")

	var emptyRet cloudmodel.InfraInfoList
	// var infraInfoList cloudmodel.InfraInfoList

	// Initialize Tumblebug session
	tbSess := tbclient.NewSession()

	infraInfoList, err := tbSess.ReadAllInfra(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to retrieve all migrated VM infrastructures (nsId: %s)", nsId)
		return emptyRet, err
	}

	// Convert the response model from 'tbclient.TbInfraInfoList' to 'cloudmodel.InfraInfoList'
	convertedVmInfraInfoList, err := modelconv.ConvertWithValidation[tbclient.TbInfraInfoList, cloudmodel.InfraInfoList](infraInfoList)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info list (nsId: %s)", nsId)
		return emptyRet, err
	}

	log.Info().Msgf("Retrieved all migrated VM infrastructures (nsId: %s, count: %d) successfully", nsId, len(convertedVmInfraInfoList.Infra))
	return convertedVmInfraInfoList, nil
}

// Get all migrated VM infrastructures
func ListVMInfraIDs(nsId string, option string) (cloudmodel.IdList, error) {
	log.Info().Msg("Listing all migrated VM infrastructure IDs")

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

	log.Info().Msgf("Retrieved all migrated VM infrastructure IDs (nsId: %s, count: %d) successfully", nsId, len(idList.IdList))
	return idList, nil
}

// Get the migrated VM infrastructure
func GetVMInfra(nsId, infraId string) (cloudmodel.InfraInfo, error) {
	log.Info().Msgf("Retrieving the migrated VM infrastructure (nsId: %s, infraId: %s)", nsId, infraId)

	// Initialize Tumblebug session
	tbSess := tbclient.NewSession()
	vmInfraInfo, err := tbSess.ReadInfra(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get the infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return cloudmodel.InfraInfo{}, err
	}

	// Convert the response model from 'tbmodel.InfraInfo' to 'cloudmodel.InfraInfo'
	convertedVmInfraInfo, err := modelconv.ConvertWithValidation[tbmodel.InfraInfo, cloudmodel.InfraInfo](vmInfraInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return cloudmodel.InfraInfo{}, err
	}

	log.Info().Msgf("Retrieved the migrated VM infrastructure (nsId: %s, infraId: %s) successfully", nsId, infraId)
	return convertedVmInfraInfo, nil
}

// Delete the migrated VM infrastructure
func DeleteVMInfra(nsId, infraId, option string) (common.SimpleMsg, error) {
	log.Info().Msg("Deleting the migrated VM infrastructure")

	// Initialize Tumblebug session
	// tbSess := tbclient.NewSession()

	// 1. Read Infra info
	infraInfo, err := tbclient.NewSession().ReadInfra(nsId, infraId)
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
	log.Debug().Msgf("Sleeping for 3 seconds to ensure Infra is deleted (nsId: %s)", nsId)
	time.Sleep(3 * time.Second)

	//3. Delete security groups
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
	log.Debug().Msgf("Sleeping for 3 seconds to ensure security groups are deleted (nsId: %s)", nsId)
	time.Sleep(3 * time.Second)

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
	log.Debug().Msgf("Sleeping for 3 seconds to ensure SSH keys are deleted (nsId: %s)", nsId)
	time.Sleep(3 * time.Second)

	// 5. Delete vNets
	// Collect unique vNet IDs from all Nodes
	vNetIdMap := make(map[string]struct{})
	for _, node := range infraInfo.Node {
		vNetIdMap[node.VNetId] = struct{}{}
	}
	log.Debug().Msgf("Deleting VNets (nsId: %s, vNets: %v)", nsId, vNetIdMap)

	// Delete all vNet

	const vNetDeleteMaxRetries = 10
	const vNetDeleteRetryInterval = 10 * time.Second

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

	// Sleep for a while to ensure all resources are deleted
	log.Debug().Msgf("Sleeping for 3 seconds to ensure VNets are deleted (nsId: %s)", nsId)
	time.Sleep(3 * time.Second)

	// 6. Delete shared resources
	idList, err = tbclient.NewSession().DeleteSharedResources(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to delete shared resources (nsId: %s, infraId: %s)", nsId, infraId)
		return common.SimpleMsg{}, err
	}
	log.Debug().Msgf("Shared resources deleted (nsId: %s, infraId: %s, IdList: %s)", nsId, infraId, idList.IdList)

	/*
	 * [Output] Return the result
	 */

	ret := common.SimpleMsg{
		Message: fmt.Sprintf("Successfully deleted the infrastructure and resources (nsId: %s, infraId: %s)", nsId, infraId),
	}
	log.Info().Msgf("Successfully deleted the infrastructure and resources (nsId: %s, infraId: %s)", nsId, infraId)
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

// NetworkRequirement represents the virtual network and subnets required by the NodeGroups
type NetworkRequirement struct {
	VNetId         string
	SubnetIds      []string
	ConnectionName string
}

// deriveNetworkIds groups and extracts virtual network requirements from NodeGroups
func deriveNetworkIds(nodeGroups []cloudmodel.CreateNodeGroupReq) []NetworkRequirement {
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

	var reqs []NetworkRequirement
	for _, vNetId := range orderedVNets {
		reqs = append(reqs, NetworkRequirement{
			VNetId:         vNetId,
			SubnetIds:      vNetSubnets[vNetId],
			ConnectionName: vNetConnection[vNetId],
		})
	}
	return reqs
}

// useOrCreateNetwork checks if VNet and required subnets exist, and creates them from the creation request if missing
func useOrCreateNetwork(nsId string, netReq NetworkRequirement, vNetCreationReq cloudmodel.VNetReq) error {
	vNetInfo, err := tbclient.NewSession().ReadVNet(nsId, netReq.VNetId)
	vNetExists := (err == nil && vNetInfo.Id != "")
	allSubnetsExist := true

	if vNetExists {
		existingSubnets := make(map[string]bool)
		for _, sub := range vNetInfo.SubnetInfoList {
			existingSubnets[sub.Name] = true
		}
		for _, reqSubnet := range netReq.SubnetIds {
			if !existingSubnets[reqSubnet] {
				allSubnetsExist = false
				log.Warn().Msgf("subnet %s is missing in existing vNet %s", reqSubnet, netReq.VNetId)
				break
			}
		}
	}

	if vNetExists && allSubnetsExist {
		log.Info().Msgf("vNet %s and all required subnets already exist. CM-Beetle will reuse it.", netReq.VNetId)
		return nil
	}

	if vNetCreationReq.CidrBlock == "" {
		return fmt.Errorf("vNet %s (or its subnets) does not exist, and VNet creation request is missing or invalid", netReq.VNetId)
	}

	vNetReq := vNetCreationReq
	vNetReq.Name = netReq.VNetId
	if netReq.ConnectionName != "" {
		vNetReq.ConnectionName = netReq.ConnectionName
	}

	var newSubnetList []cloudmodel.SubnetReq
	for idx, subnetName := range netReq.SubnetIds {
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

// SshKeyRequirement represents the SSH key required by the NodeGroups
type SshKeyRequirement struct {
	SshKeyId       string
	ConnectionName string
}

// deriveSshKeyIds extracts unique SSH key requirements from NodeGroups
func deriveSshKeyIds(nodeGroups []cloudmodel.CreateNodeGroupReq) []SshKeyRequirement {
	var reqs []SshKeyRequirement
	seenSshKeys := make(map[string]bool)
	for _, ng := range nodeGroups {
		if ng.SshKeyId == "" || seenSshKeys[ng.SshKeyId] {
			continue
		}
		seenSshKeys[ng.SshKeyId] = true
		reqs = append(reqs, SshKeyRequirement{
			SshKeyId:       ng.SshKeyId,
			ConnectionName: ng.ConnectionName,
		})
	}
	return reqs
}

// useOrCreateSshKey checks if SSH key exists, and creates it from the creation request if missing
func useOrCreateSshKey(nsId string, sshKeyReq SshKeyRequirement, sshKeyCreationReq cloudmodel.SshKeyReq) error {
	sshKeyInfo, err := tbclient.NewSession().ReadSshKey(nsId, sshKeyReq.SshKeyId)
	if err == nil && sshKeyInfo.Id != "" {
		log.Info().Msgf("SSH key %s already exists. CM-Beetle will reuse it.", sshKeyReq.SshKeyId)
		return nil
	}

	if sshKeyCreationReq.Name == "" {
		return fmt.Errorf("SSH key %s does not exist, and SSH key creation request is missing or invalid", sshKeyReq.SshKeyId)
	}

	req := sshKeyCreationReq
	req.Name = sshKeyReq.SshKeyId
	if sshKeyReq.ConnectionName != "" {
		req.ConnectionName = sshKeyReq.ConnectionName
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

// SecurityGroupRequirement represents the security group required by the NodeGroups
type SecurityGroupRequirement struct {
	SecurityGroupId string
	VNetId          string
	ConnectionName  string
}

// deriveSecurityGroupIds extracts unique security group requirements from NodeGroups
func deriveSecurityGroupIds(nodeGroups []cloudmodel.CreateNodeGroupReq) []SecurityGroupRequirement {
	var reqs []SecurityGroupRequirement
	seenSgs := make(map[string]bool)
	for _, ng := range nodeGroups {
		for _, sgId := range ng.SecurityGroupIds {
			if sgId == "" || seenSgs[sgId] {
				continue
			}
			seenSgs[sgId] = true
			reqs = append(reqs, SecurityGroupRequirement{
				SecurityGroupId: sgId,
				VNetId:          ng.VNetId,
				ConnectionName:  ng.ConnectionName,
			})
		}
	}
	return reqs
}

// useOrCreateSecurityGroup checks if security group exists, and creates it from the creation request list if missing
func useOrCreateSecurityGroup(nsId string, sgReq SecurityGroupRequirement, sgCreationReqList []cloudmodel.SecurityGroupReq) error {
	sgInfo, err := tbclient.NewSession().ReadSecurityGroup(nsId, sgReq.SecurityGroupId)
	if err == nil && sgInfo.Id != "" {
		log.Info().Msgf("Security group %s already exists. CM-Beetle will reuse it.", sgReq.SecurityGroupId)
		return nil
	}

	var sgCreationReq cloudmodel.SecurityGroupReq
	found := false
	for _, sg := range sgCreationReqList {
		if sg.Name == sgReq.SecurityGroupId {
			sgCreationReq = sg
			found = true
			break
		}
	}
	if !found {
		sgCreationReq = cloudmodel.SecurityGroupReq{Name: sgReq.SecurityGroupId}
	}

	if sgCreationReq.ConnectionName == "" && sgReq.ConnectionName != "" {
		sgCreationReq.ConnectionName = sgReq.ConnectionName
	}
	if sgCreationReq.VNetId == "" && sgReq.VNetId != "" {
		sgCreationReq.VNetId = sgReq.VNetId
	}

	if sgCreationReq.ConnectionName == "" || sgCreationReq.VNetId == "" {
		return fmt.Errorf("security group %s does not exist, and required ConnectionName or VNetId is missing", sgReq.SecurityGroupId)
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

// validateTargeInfraModel validates the target infrastructure model for fresh creation (useExisting=false)
func validateTargeInfraModel(nsId string, targetVmInfraModel *cloudmodel.RecommendedInfra) error {
	// * 1. Validate that name fields are not empty
	if targetVmInfraModel == nil {
		log.Error().Msgf("target infrastructure model is nil (nsId: %s)", nsId)
		return fmt.Errorf("target infrastructure model is nil")
	}
	if targetVmInfraModel.TargetInfra.Name == "" { // MCI name
		log.Error().Msgf("target VM infrastructure name is empty (nsId: %s)", nsId)
		return fmt.Errorf("target VM infrastructure name is empty")
	}
	if targetVmInfraModel.TargetVNet.Name == "" {
		log.Error().Msgf("target VNet name is empty (nsId: %s)", nsId)
		return fmt.Errorf("target VNet name is empty")
	}
	if targetVmInfraModel.TargetSshKey.Name == "" {
		log.Error().Msgf("target SSH key name is empty (nsId: %s)", nsId)
		return fmt.Errorf("target SSH key name is empty")
	}
	for _, sg := range targetVmInfraModel.TargetSecurityGroupList {
		if sg.Name == "" {
			log.Error().Msgf("target security group name is empty (nsId: %s)", nsId)
			return fmt.Errorf("target security group name is empty")
		}
	}

	// * 2. Validate that the names or IDs are matched in the model
	// Check if each Node's vNetId matches the target VNet name
	for _, nodegroup := range targetVmInfraModel.TargetInfra.NodeGroups {
		if nodegroup.VNetId != targetVmInfraModel.TargetVNet.Name {
			log.Error().Msgf("target VM infrastructure vNetId (%s) does not match target VNet name (%s)",
				nodegroup.VNetId, targetVmInfraModel.TargetVNet.Name)
			return fmt.Errorf("target VM infrastructure vNetId (%s) does not match target VNet name (%s)",
				nodegroup.VNetId, targetVmInfraModel.TargetVNet.Name)
		}
	}

	// Check if each Node's SshKeyId matches the target SSH key name
	for _, nodegroup := range targetVmInfraModel.TargetInfra.NodeGroups {
		if nodegroup.SshKeyId != targetVmInfraModel.TargetSshKey.Name {
			log.Error().Msgf("target VM infrastructure SshKeyId (%s) does not match target SSH key name (%s)",
				nodegroup.SshKeyId, targetVmInfraModel.TargetSshKey.Name)
			return fmt.Errorf("target VM infrastructure SshKeyId (%s) does not match target SSH key name (%s)",
				nodegroup.SshKeyId, targetVmInfraModel.TargetSshKey.Name)
		}
	}

	// Check if each Node's spec and image are valid and compatible
	for _, nodegroup := range targetVmInfraModel.TargetInfra.NodeGroups {
		specId := strings.TrimSpace(nodegroup.SpecId)
		imageId := strings.TrimSpace(nodegroup.ImageId)
		connectionName := strings.TrimSpace(nodegroup.ConnectionName)

		// 1. Validate SpecId is not empty
		if specId == "" || specId == "empty" {
			err := fmt.Errorf("invalid SpecId '%s' in nodegroup '%s'", specId, nodegroup.Name)
			log.Error().Err(err).Msgf("required SpecId (current SpecId is '%s' for nodegroup '%s')", specId, nodegroup.Name)
			return err
		}

		// 2. Validate ImageId is not empty
		if imageId == "" || imageId == "empty" {
			err := fmt.Errorf("invalid ImageId '%s' in nodegroup '%s'", imageId, nodegroup.Name)
			log.Error().Err(err).Msgf("required ImageId (current ImageId is '%s' for nodegroup '%s')", imageId, nodegroup.Name)
			return err
		}

		// 3. Validate ConnectionName is not empty to extract CSP information
		if connectionName == "" {
			err := fmt.Errorf("invalid ConnectionName '%s' in nodegroup '%s'", connectionName, nodegroup.Name)
			log.Error().Err(err).Msgf("required ConnectionName (current ConnectionName is '%s' for nodegroup '%s')", connectionName, nodegroup.Name)
			return err
		}

		// 4. Extract CSP information from ConnectionName (format: "csp-region")
		connectionParts := strings.Split(connectionName, "-")
		if len(connectionParts) < 2 {
			err := fmt.Errorf("invalid connection name format '%s' in nodegroup '%s'", connectionName, nodegroup.Name)
			log.Error().Err(err).Msgf("invalid connection name format '%s' for nodegroup '%s', expected format: 'csp-region'",
				connectionName, nodegroup.Name)
			return err
		}
		csp := connectionParts[0]

		// 5. Retrieve SpecInfo
		specInfo, err := tbclient.NewSession().ReadVmSpec("system", specId)
		if err != nil {
			log.Error().Err(err).Msgf("failed to read VM spec (nsId: %s, vmSpecId: %s)", nsId, specId)
			return fmt.Errorf("failed to read VM spec (nsId: %s, vmSpecId: %s): %w", nsId, specId, err)
		}

		// 6. Retrieve ImageInfo
		// Note - current imageId format: csp+cspImageName (e.g., alibaba+ubuntu_22_04_x64_20G_alibase_20250722.vhd)
		// ref: https://github.com/cloud-barista/cb-tumblebug/pull/2130#issuecomment-3243624048
		// TODO: ImageId should be updated later as Tumblebug's/ns/{nsId}/resources/image/{imageId}` API changes.
		imageKey := imageId
		if !strings.Contains(imageKey, "+") {
			// If ImageId doesn't contain '+', assume it needs CSP prefix
			imageKey = fmt.Sprintf("%s+%s", csp, imageId)
		}

		imageInfo, err := tbclient.NewSession().ReadVmOsImage("system", imageKey)
		if err != nil {
			log.Error().Err(err).Msgf("failed to read VM OS image (nsId: %s, vmOsImageKey: %s)", nsId, imageKey)
			return fmt.Errorf("failed to read VM OS image (nsId: %s, vmOsImageKey: %s): %w", nsId, imageKey, err)
		}

		// 7. Convert models to cloudmodel format for compatibility check
		specInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.SpecInfo, cloudmodel.SpecInfo](specInfo)
		if err != nil {
			log.Error().Err(err).Msgf("failed to convert spec info for compatibility check (specId: %s)", specId)
			return fmt.Errorf("failed to convert spec info for compatibility check (specId: %s): %w", specId, err)
		}

		imageInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.ImageInfo, cloudmodel.ImageInfo](imageInfo)
		if err != nil {
			log.Error().Err(err).Msgf("failed to convert image info for compatibility check (imageId: %s)", imageId)
			return fmt.Errorf("failed to convert image info for compatibility check (imageId: %s): %w", imageId, err)
		}

		// 8. Check compatibility between spec and image
		isCompatible := recommendation.CheckSpecImageCompatibility(csp, specInfoConverted, imageInfoConverted)
		if !isCompatible {
			log.Error().Msgf("VM spec '%s' and image '%s' are incompatible for CSP '%s' in nodegroup '%s'",
				specId, imageId, csp, nodegroup.Name)
			return fmt.Errorf("VM spec '%s' and image '%s' are incompatible for CSP '%s' in nodegroup '%s'",
				specId, imageId, csp, nodegroup.Name)
		}

		log.Debug().Msgf("VM spec '%s' and image '%s' are compatible for CSP '%s' in nodegroup '%s'",
			specId, imageId, csp, nodegroup.Name)
	}

	// Check if each security group name is contained in the target security group list
	for _, nodegroup := range targetVmInfraModel.TargetInfra.NodeGroups {
		found := false
		for _, sgId := range nodegroup.SecurityGroupIds {
			// Check if the security group name matches any target security group name
			for _, targetSg := range targetVmInfraModel.TargetSecurityGroupList {
				if sgId == targetSg.Name {
					found = true
					break
				}
			}
			if !found {
				log.Error().Msgf("target VM infrastructure security group name (%s) does not match any target security group name in the list",
					sgId)
				return fmt.Errorf("target VM infrastructure security group name (%s) does not match any target security group name in the list",
					sgId)
			}
		}
	}

	// * 3. Validate that the vNet, VM specs, VM OS images, and security groups exist
	// Validate that the vNet exists by the vNet name
	// For VNet, it's normal if the resource doesn't exist
	vNetInfo, err := tbclient.NewSession().ReadVNet(nsId, targetVmInfraModel.TargetVNet.Name)
	if err != nil {
		log.Debug().Msgf("the vNet not found (nsId: %s, vNet.Name: %s), which is normal case", nsId, targetVmInfraModel.TargetVNet.Name)
	}
	if vNetInfo.Id != "" {
		log.Error().Msgf("the vNet already exists (nsId: %s, vNetInfo.Id: %s)", nsId, vNetInfo.Id)
		return fmt.Errorf("the vNet already exists (nsId: %s, vNetInfo.Id: %s)", nsId, vNetInfo.Id)
	}

	// Validate that the SSH key exists by the SSH key name
	// For SSH Key, it's normal if the resource doesn't exist
	sshKeyInfo, err := tbclient.NewSession().ReadSshKey(nsId, targetVmInfraModel.TargetSshKey.Name)
	if err != nil {
		log.Debug().Err(err).Msgf("SSH key not found (nsId: %s, sshKey.Name: %s), which is normal case", nsId, targetVmInfraModel.TargetSshKey.Name)
	}
	if sshKeyInfo.Id != "" {
		log.Error().Msgf("the SSH key already exists (nsId: %s, sshKey.Id: %s)", nsId, sshKeyInfo.Id)
		return fmt.Errorf("the SSH key already exists (nsId: %s, sshKey.Id: %s)", nsId, sshKeyInfo.Id)
	}

	// Note: VM specs and VM OS images validation is now handled above in the spec-image compatibility check loop
	// This provides better validation by checking both existence and compatibility together

	// Validate that the security groups exist by the security group name
	// For Security Groups, it's normal if the resources don't exist
	for _, sg := range targetVmInfraModel.TargetSecurityGroupList {
		sgInfo, err := tbclient.NewSession().ReadSecurityGroup(nsId, sg.Name)
		if err != nil {
			log.Debug().Msgf("the security group not found (nsId: %s, sg.Name: %s), which is normal case", nsId, sg.Name)
		}
		if sgInfo.Id != "" {
			log.Error().Msgf("the security group already exists (nsId: %s, sgInfo.Id: %s)", nsId, sgInfo.Id)
			return fmt.Errorf("the security group already exists (nsId: %s, sgInfo.Id: %s)", nsId, sgInfo.Id)
		}
	}

	return nil
}

// validateTargeInfraModelWithExisting validates the target infrastructure model for resource reuse (useExisting=true)
func validateTargeInfraModelWithExisting(nsId string, targetVmInfraModel *cloudmodel.RecommendedInfra) error {
	// * 1. Validate that name fields are not empty
	if targetVmInfraModel == nil {
		log.Error().Msgf("target infrastructure model is nil (nsId: %s)", nsId)
		return fmt.Errorf("target infrastructure model is nil")
	}
	if targetVmInfraModel.TargetInfra.Name == "" { // MCI name
		log.Error().Msgf("target VM infrastructure name is empty (nsId: %s)", nsId)
		return fmt.Errorf("target VM infrastructure name is empty")
	}

	// Verify required ID fields for NodeGroups in useExisting mode
	for _, nodegroup := range targetVmInfraModel.TargetInfra.NodeGroups {
		if nodegroup.VNetId == "" {
			log.Error().Msgf("VNet ID is empty for nodegroup %s in useExisting mode (nsId: %s)", nodegroup.Name, nsId)
			return fmt.Errorf("VNet ID is empty for nodegroup %s in useExisting mode", nodegroup.Name)
		}
		if nodegroup.SshKeyId == "" {
			log.Error().Msgf("SSH key ID is empty for nodegroup %s in useExisting mode (nsId: %s)", nodegroup.Name, nsId)
			return fmt.Errorf("SSH key ID is empty for nodegroup %s in useExisting mode", nodegroup.Name)
		}
		if len(nodegroup.SecurityGroupIds) == 0 {
			log.Error().Msgf("Security group IDs list is empty for nodegroup %s in useExisting mode (nsId: %s)", nodegroup.Name, nsId)
			return fmt.Errorf("Security group IDs list is empty for nodegroup %s in useExisting mode", nodegroup.Name)
		}
	}

	// Check if each Node's spec and image are valid and compatible (inlined spec/image compatibility check)
	for _, nodegroup := range targetVmInfraModel.TargetInfra.NodeGroups {
		specId := strings.TrimSpace(nodegroup.SpecId)
		imageId := strings.TrimSpace(nodegroup.ImageId)
		connectionName := strings.TrimSpace(nodegroup.ConnectionName)

		if specId == "" || specId == "empty" {
			err := fmt.Errorf("invalid SpecId '%s' in nodegroup '%s'", specId, nodegroup.Name)
			return err
		}
		if imageId == "" || imageId == "empty" {
			err := fmt.Errorf("invalid ImageId '%s' in nodegroup '%s'", imageId, nodegroup.Name)
			return err
		}
		if connectionName == "" {
			err := fmt.Errorf("invalid ConnectionName '%s' in nodegroup '%s'", connectionName, nodegroup.Name)
			return err
		}

		connectionParts := strings.Split(connectionName, "-")
		if len(connectionParts) < 2 {
			err := fmt.Errorf("invalid connection name format '%s' in nodegroup '%s'", connectionName, nodegroup.Name)
			return err
		}
		csp := connectionParts[0]

		specInfo, err := tbclient.NewSession().ReadVmSpec("system", specId)
		if err != nil {
			return fmt.Errorf("failed to read VM spec (nsId: %s, vmSpecId: %s): %w", nsId, specId, err)
		}

		imageKey := imageId
		if !strings.Contains(imageKey, "+") {
			imageKey = fmt.Sprintf("%s+%s", csp, imageId)
		}
		imageInfo, err := tbclient.NewSession().ReadVmOsImage("system", imageKey)
		if err != nil {
			return fmt.Errorf("failed to read VM OS image (nsId: %s, vmOsImageKey: %s): %w", nsId, imageKey, err)
		}

		specInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.SpecInfo, cloudmodel.SpecInfo](specInfo)
		if err != nil {
			return err
		}
		imageInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.ImageInfo, cloudmodel.ImageInfo](imageInfo)
		if err != nil {
			return err
		}

		isCompatible := recommendation.CheckSpecImageCompatibility(csp, specInfoConverted, imageInfoConverted)
		if !isCompatible {
			return fmt.Errorf("VM spec '%s' and image '%s' are incompatible for CSP '%s' in nodegroup '%s'",
				specId, imageId, csp, nodegroup.Name)
		}
	}

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
