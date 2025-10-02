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

// Package migration is to privision targat multi-cloud infra for migration
package migration

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"

	// cloudmodel "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/cloud/infra"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/modelconv"
	"github.com/rs/zerolog/log"
)

//"log"

//csv file handling

// REST API (echo)

// "github.com/cloud-barista/cm-beetle/pkg/core/mcir"

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

// const labelAutoGen string = "AutoGen"

// DefaultSystemLabel is const for string to specify the Default System Label
const DefaultSystemLabel string = "Managed by CM-Beetle"

// CreateVMInfraWithDefaults Create a VM infrastructure with defaults for the computing infra migration
func CreateVMInfraWithDefaults(nsId string, infraModel *cloudmodel.MciDynamicReq) (cloudmodel.VmInfraInfo, error) {
	log.Info().Msg("Creating VM infrastructure with defaults")

	// Initialize Tumblebug client
	tbApiConfig := tbclient.ApiConfig{
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
		RestUrl:  config.Tumblebug.RestUrl,
	}
	tbCli := tbclient.NewClient(tbApiConfig)
	// Convert the request model from 'cloudmodel.MciDynamicReq' to 'tbmodel.MciDynamicReq'
	infraModelConverted, err := modelconv.ConvertWithValidation[cloudmodel.MciDynamicReq, tbmodel.MciDynamicReq](*infraModel)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure model (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	vmInfraInfo, err := tbCli.CreateMciDynamic(nsId, infraModelConverted)
	if err != nil {
		log.Error().Err(err).Msgf("failed to migrate the infrastructure (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	// Convert the response model from 'tbmodel.MciInfo' to 'cloudmodel.VmInfraInfo'
	convertedVmInfraInfo, err := modelconv.ConvertWithValidation[tbmodel.MciInfo, cloudmodel.VmInfraInfo](vmInfraInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	log.Info().Msgf("VM infrastructure created successfully (nsId: %s, mciName: %s)", nsId, convertedVmInfraInfo.MciInfo.Name)

	return convertedVmInfraInfo, nil
}

// CreateVMInfra creates a VM infrastructure for the computing infra migration
func CreateVMInfra(nsId string, targetInfraModel *cloudmodel.RecommendedVmInfra) (cloudmodel.VmInfraInfo, error) {
	log.Info().Msg("Creating VM infrastructure")

	emptyRet := cloudmodel.VmInfraInfo{}

	/*
	 * [Input] Receive and validate the target infrastructure model
	 */

	err := validateTargeVmtInfraModel(nsId, targetInfraModel)
	if err != nil {
		log.Error().Err(err).Msgf("failed to validate the target infrastructure model (nsId: %s)", nsId)
		return emptyRet, err
	}
	log.Info().Msgf("the target infrastructure model is valid (nsId: %s)", nsId)

	// Initialize Tumblebug client
	tbApiConfig := tbclient.ApiConfig{
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
		RestUrl:  config.Tumblebug.RestUrl,
	}
	tbCli := tbclient.NewClient(tbApiConfig)

	/*
	 * [Process] Create a VM infrastructure
	 */
	// 1. Check if the namespace exists
	log.Debug().Msgf("Checking if the namespace exists (nsId: %s)", nsId)
	_, err = tbCli.ReadNamespace(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to read the namespace (nsId: %s)", nsId)
		return emptyRet, err
	}

	log.Debug().Msgf("Checking if the MCI (%s) exists in the namespace (%s)", targetInfraModel.TargetVmInfra.Name, nsId)
	tempMciInfo, err := tbCli.ReadMci(nsId, targetInfraModel.TargetVmInfra.Name)
	if tempMciInfo.Id != "" {
		log.Error().Err(err).Msgf("the MCI already exist (nsId: %s, mciName: %s)", nsId, targetInfraModel.TargetVmInfra.Name)
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

	vNetInfo, err := tbCli.CreateVNet(nsId, tbVNetReq)
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

	sshKeyInfo, err := tbCli.CreateSshKey(nsId, tbSshKeyReq)
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

		// Create security group
		log.Debug().Msgf("Creating a security group (nsId: %s, sgReq.sgName: %s, sgReq.VNetId: %s, vNetInfo.vNetId: %s)",
			nsId, sgReq.Name, sgReq.VNetId, vNetInfo.Id)

		// Convert model from 'cloudmodel.SecurityGroupReq' to 'tbmodel.SecurityGroupReq'
		tbSgReq, err := modelconv.ConvertWithValidation[cloudmodel.SecurityGroupReq, tbmodel.SecurityGroupReq](sgReq)
		if err != nil {
			log.Error().Err(err).Msgf("failed to convert SSH key request (nsId: %s)", nsId)
			return emptyRet, err
		}

		sgInfo, err := tbCli.CreateSecurityGroup(nsId, tbSgReq, "")
		if err != nil {
			log.Error().Err(err).Msgf("failed to create the security group (nsId: %s)", nsId)
			return emptyRet, err
		}
		log.Debug().Msgf("security group created: %s", sgInfo.Id)

		sgInfoList = append(sgInfoList, sgInfo)
	}
	log.Debug().Msgf("sgInfoList length: %d", len(sgInfoList))
	log.Debug().Msgf("sgInfoList: %+v", sgInfoList)

	// 7. Create a VM infrastructure (i.e., MCI)
	// Get multi-cloud infrastructure (MCI) request body from the input infraModel
	mciReq := targetInfraModel.TargetVmInfra
	log.Debug().Msgf("Creating a multi-cloud infrastructure (nsId: %s, mciName: %s)", nsId, mciReq.Name)
	log.Debug().Msgf("mciReq: %+v", mciReq)

	// Convert model from 'cloudmodel.SecurityGroupReq' to 'tbmodel.SecurityGroupReq'
	tbMciReq, err := modelconv.ConvertWithValidation[cloudmodel.MciReq, tbmodel.MciReq](mciReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the MCI request (nsId: %s)", nsId)
		return emptyRet, err
	}
	log.Debug().Msgf("tbMciReq: %+v", tbMciReq)

	// Set post-command for stable mci provisioning if a user didn't set it
	// If a user already set it, use it as is
	if len(tbMciReq.PostCommand.Command) == 0 {
		log.Debug().Msgf("Setting default post-command `uname -a` for stable MCI provisioning (nsId: %s)", nsId)

		commands := []string{
			"uname -a",
		}
		username := "cb-user"

		tbMciReq.PostCommand = tbmodel.MciCmdReq{
			UserName: username,
			Command:  commands,
		}
	}

	// Create multi-cloud infrastructure
	mciInfo, err := tbCli.CreateMci(nsId, tbMciReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create the multi-cloud infrastructure (nsId: %s)", nsId)
		return emptyRet, err
	}
	log.Debug().Msgf("multi-cloud infrastructure created: %s", mciInfo.Id)

	/*
	 * [Output] Return the created multi-cloud infrastructure info
	 */

	// Convert the response model from 'tbmodel.MciInfo' to 'cloudmodel.MciInfo'
	mciInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.MciInfo, cloudmodel.MciInfo](mciInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s)", nsId)
		return emptyRet, err
	}
	var temp cloudmodel.VmInfraInfo
	temp.MciInfo = mciInfoConverted

	log.Info().Msgf("VM infrastructure created successfully (nsId: %s, mciName: %s)", nsId, mciInfoConverted.Name)
	return temp, nil
}

// List all migrated VM infrastructures
func ListAllVMInfraInfo(nsId string) (cloudmodel.MciInfoList, error) {
	log.Info().Msg("Listing all migrated VM infrastructures")

	var emptyRet cloudmodel.MciInfoList
	// var mciInfoList cloudmodel.MciInfoList

	// Initialize Tumblebug client
	tbApiConfig := tbclient.ApiConfig{
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
		RestUrl:  config.Tumblebug.RestUrl,
	}
	tbCli := tbclient.NewClient(tbApiConfig)

	mciInfoList, err := tbCli.ReadAllMci(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to retrieve all migrated VM infrastructures (nsId: %s)", nsId)
		return emptyRet, err
	}

	// Convert the response model from 'tbclient.TbMciInfoList' to 'cloudmodel.MciInfoList'
	convertedVmInfraInfoList, err := modelconv.ConvertWithValidation[tbclient.TbMciInfoList, cloudmodel.MciInfoList](mciInfoList)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info list (nsId: %s)", nsId)
		return emptyRet, err
	}

	log.Info().Msgf("Retrieved all migrated VM infrastructures (nsId: %s, count: %d) successfully", nsId, len(convertedVmInfraInfoList.Mci))
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

	// Initialize Tumblebug client
	tbApiConfig := tbclient.ApiConfig{
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
		RestUrl:  config.Tumblebug.RestUrl,
	}
	tbCli := tbclient.NewClient(tbApiConfig)
	mciIdList, err := tbCli.ReadMciIDs(nsId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get the infrastructure IDs (nsId: %s)", nsId)
		return emptyRet, err
	}

	// Return the result
	idList.IdList = append(idList.IdList, mciIdList.IdList...)

	log.Info().Msgf("Retrieved all migrated VM infrastructure IDs (nsId: %s, count: %d) successfully", nsId, len(idList.IdList))
	return idList, nil
}

// Get the migrated VM infrastructure
func GetVMInfra(nsId, infraId string) (cloudmodel.MciInfo, error) {
	log.Info().Msgf("Retrieving the migrated VM infrastructure (nsId: %s, infraId: %s)", nsId, infraId)

	// Initialize Tumblebug client
	tbApiConfig := tbclient.ApiConfig{
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
		RestUrl:  config.Tumblebug.RestUrl,
	}
	tbCli := tbclient.NewClient(tbApiConfig)
	vmInfaInfo, err := tbCli.ReadMci(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to get the infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return cloudmodel.MciInfo{}, err
	}

	// Convert the response model from 'tbmodel.MciInfo' to 'cloudmodel.MciInfo'
	convertedVmInfaInfo, err := modelconv.ConvertWithValidation[tbmodel.MciInfo, cloudmodel.MciInfo](vmInfaInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return cloudmodel.MciInfo{}, err
	}

	log.Info().Msgf("Retrieved the migrated VM infrastructure (nsId: %s, infraId: %s) successfully", nsId, infraId)
	return convertedVmInfaInfo, nil
}

// Delete the migrated VM infrastructure
func DeleteVMInfra(nsId, infraId, option string) (common.SimpleMsg, error) {
	log.Info().Msg("Deleting the migrated VM infrastructure")

	// Initialize Tumblebug client
	apiConfig := tbclient.ApiConfig{
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
		RestUrl:  config.Tumblebug.RestUrl,
	}
	tbCli := tbclient.NewClient(apiConfig)

	// 1. Read MCI info
	mciInfo, err := tbCli.ReadMci(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Msgf("failed to read the infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return common.SimpleMsg{}, err
	}

	// 2. Delete MCI
	idList, err := tbCli.DeleteMci(nsId, infraId, option)
	if err != nil {
		log.Error().Err(err).Msgf("failed to delete the infrastructure (nsId: %s, infraId: %s)", nsId, infraId)
		return common.SimpleMsg{}, err
	}
	log.Debug().Msgf("MCI deleted (nsId: %s, infraId: %s, IdList: %s)", nsId, infraId, idList.IdList)

	// Sleep for a while to ensure previous deletions are completed
	log.Debug().Msgf("Sleeping for 3 seconds to ensure MCI is deleted (nsId: %s)", nsId)
	time.Sleep(3 * time.Second)

	//3. Delete security groups
	// Collect unique security group IDs from all VMs
	sgIdMap := make(map[string]struct{})
	for _, vm := range mciInfo.Vm {
		for _, sgId := range vm.SecurityGroupIds {
			sgIdMap[sgId] = struct{}{}
		}
	}
	log.Debug().Msgf("Deleting security groups (nsId: %s, SGs: %v)", nsId, sgIdMap)

	// Delete all security groups
	for sgId := range sgIdMap {
		msg, err := tbCli.DeleteSecurityGroup(nsId, sgId)
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
	// Collect unique SSH Key IDs from all VMs
	sshKeyIdMap := make(map[string]struct{})
	for _, vm := range mciInfo.Vm {
		sshKeyIdMap[vm.SshKeyId] = struct{}{}
	}
	log.Debug().Msgf("Deleting SSH keys (nsId: %s, sshKeys: %v)", nsId, sshKeyIdMap)

	// Delete all SSH Key
	for sshKeyId := range sshKeyIdMap {
		// Delete SSH Key
		log.Debug().Msgf("Deleting SSH key (nsId: %s, sshKeyId: %s)", nsId, sshKeyId)
		msg, err := tbCli.DeleteSshKey(nsId, sshKeyId)
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
	// Collect unique vNet IDs from all VMs
	vNetIdMap := make(map[string]struct{})
	for _, vm := range mciInfo.Vm {
		vNetIdMap[vm.VNetId] = struct{}{}
	}
	log.Debug().Msgf("Deleting VNets (nsId: %s, vNets: %v)", nsId, vNetIdMap)

	// Delete all vNet
	for vNetId := range vNetIdMap {
		log.Debug().Msgf("Deleting VNet (nsId: %s, vNetId: %s, action: %s)", nsId, vNetId, "withsubnets")
		msg, err := tbCli.DeleteVNet(nsId, vNetId, "withsubnets")
		if err != nil {
			log.Error().Err(err).Msgf("failed to delete VNet (nsId: %s, vNetId: %s)", nsId, vNetId)
			// Continue deleting other resources even if this fails
		} else {
			log.Debug().Msgf("VNet deleted (nsId: %s, vNetId: %s, msg:%s)", nsId, vNetId, msg)
		}
	}

	// Sleep for a while to ensure all resources are deleted
	log.Debug().Msgf("Sleeping for 3 seconds to ensure VNets are deleted (nsId: %s)", nsId)
	time.Sleep(3 * time.Second)

	// 6. Delete shared resources
	idList, err = tbCli.DeleteSharedResources(nsId)
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

func validateTargeVmtInfraModel(nsId string, targetVmInfraModel *cloudmodel.RecommendedVmInfra) error {

	// * 1. Validate that name fieleds are not empty
	if targetVmInfraModel == nil {
		log.Error().Msgf("target infrastructure model is nil (nsId: %s)", nsId)
		return fmt.Errorf("target infrastructure model is nil")
	}
	if targetVmInfraModel.TargetVmInfra.Name == "" { // MCI name
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
	// Check if the each VM's vNetId matches the target VNet name
	for _, subgroup := range targetVmInfraModel.TargetVmInfra.SubGroups {
		if subgroup.VNetId != targetVmInfraModel.TargetVNet.Name {
			log.Error().Msgf("target VM infrastructure vNetId (%s) does not match target VNet name (%s)",
				subgroup.VNetId, targetVmInfraModel.TargetVNet.Name)
			return fmt.Errorf("target VM infrastructure vNetId (%s) does not match target VNet name (%s)",
				subgroup.VNetId, targetVmInfraModel.TargetVNet.Name)
		}
	}

	// Check if each VM's SshKeyId matches the target SSH key name
	for _, subgroup := range targetVmInfraModel.TargetVmInfra.SubGroups {
		if subgroup.SshKeyId != targetVmInfraModel.TargetSshKey.Name {
			log.Error().Msgf("target VM infrastructure SshKeyId (%s) does not match target SSH key name (%s)",
				subgroup.SshKeyId, targetVmInfraModel.TargetSshKey.Name)
			return fmt.Errorf("target VM infrastructure SshKeyId (%s) does not match target SSH key name (%s)",
				subgroup.SshKeyId, targetVmInfraModel.TargetSshKey.Name)
		}
	}

	// Check if each VM's spec is contained in the target VM spec list
	for _, subgroup := range targetVmInfraModel.TargetVmInfra.SubGroups {
		found := false
		for _, vmSpec := range targetVmInfraModel.TargetVmSpecList {
			if subgroup.SpecId == vmSpec.Id {
				found = true
				break
			}
		}
		if !found {
			log.Error().Msgf("VM spec '%s' not found in target spec list", subgroup.SpecId)
			return fmt.Errorf("VM spec '%s' not found in target spec list", subgroup.SpecId)
		}
	}

	// Check if each VM's OS image is contained in the target VM OS image list
	for _, subgroup := range targetVmInfraModel.TargetVmInfra.SubGroups {
		found := false
		for _, vmOsImage := range targetVmInfraModel.TargetVmOsImageList {
			if subgroup.ImageId == vmOsImage.Id {
				found = true
				break
			}
		}
		if !found {
			log.Error().Msgf("VM OS image '%s' not found in target image list", subgroup.ImageId)
			return fmt.Errorf("VM OS image '%s' not found in target image list", subgroup.ImageId)
		}
	}

	// Check if each security group name is contained in the target security group list
	for _, subgroup := range targetVmInfraModel.TargetVmInfra.SubGroups {
		found := false
		for _, sgId := range subgroup.SecurityGroupIds {
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
	// Initialize Tumblebug client
	apiConfig := tbclient.ApiConfig{
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
		RestUrl:  config.Tumblebug.RestUrl,
	}
	tbCli := tbclient.NewClient(apiConfig)

	// Validate that the vNet exists by the vNet name
	// For VNet, it's normal if the resource doesn't exist
	vNetInfo, err := tbCli.ReadVNet(nsId, targetVmInfraModel.TargetVNet.Name)
	if err != nil {
		log.Debug().Msgf("the vNet not found (nsId: %s, vNet.Name: %s), which is normal case", nsId, targetVmInfraModel.TargetVNet.Name)
	}
	if vNetInfo.Id != "" {
		log.Error().Msgf("the vNet already exists (nsId: %s, vNetInfo.Id: %s)", nsId, vNetInfo.Id)
		return fmt.Errorf("the vNet already exists (nsId: %s, vNetInfo.Id: %s)", nsId, vNetInfo.Id)
	}

	// Validate that the SSH key exists by the SSH key name
	// For SSH Key, it's normal if the resource doesn't exist
	sshKeyInfo, err := tbCli.ReadSshKey(nsId, targetVmInfraModel.TargetSshKey.Name)
	if err != nil {
		log.Debug().Err(err).Msgf("SSH key not found (nsId: %s, sshKey.Name: %s), which is normal case", nsId, targetVmInfraModel.TargetSshKey.Name)
	}
	if sshKeyInfo.Id != "" {
		log.Error().Msgf("the SSH key already exists (nsId: %s, sshKey.Id: %s)", nsId, sshKeyInfo.Id)
		return fmt.Errorf("the SSH key already exists (nsId: %s, sshKey.Id: %s)", nsId, sshKeyInfo.Id)
	}

	// Validate that the VM specs exist by the VM spec ID
	for _, vmSpec := range targetVmInfraModel.TargetVmSpecList {
		_, err := tbCli.ReadVmSpec("system", vmSpec.Id)
		if err != nil {
			log.Error().Err(err).Msgf("failed to read VM spec (nsId: %s, vmSpecId: %s)", nsId, vmSpec.Id)
			return fmt.Errorf("failed to read VM spec (nsId: %s, vmSpecId: %s): %w", nsId, vmSpec.Id, err)
		}
	}

	// Validate that the VM OS images exist by the VM OS image ID
	// * Note - current imageId format: csp+cspImageName (e.g., alibaba+ubuntu_22_04_x64_20G_alibase_20250722.vhd)
	// ref: https://github.com/cloud-barista/cb-tumblebug/pull/2130#issuecomment-3243624048
	// TODO: ImageId should be updated later as Tumblebug's/ns/{nsId}/resources/image/{imageId}` API changes.
	for _, vmOsImage := range targetVmInfraModel.TargetVmOsImageList {
		imageKey := fmt.Sprintf("%s+%s", vmOsImage.ProviderName, vmOsImage.Id)

		_, err := tbCli.ReadVmOsImage("system", imageKey)
		if err != nil {
			log.Error().Err(err).Msgf("failed to read VM OS image (nsId: %s, vmOsImageKey: %s)", nsId, imageKey)
			return fmt.Errorf("failed to read VM OS image (nsId: %s, vmOsImageKey: %s): %w", nsId, imageKey, err)
		}
	}

	// Validate that the security groups exist by the security group name
	// For Security Groups, it's normal if the resources don't exist
	for _, sg := range targetVmInfraModel.TargetSecurityGroupList {
		sgInfo, err := tbCli.ReadSecurityGroup(nsId, sg.Name)
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
