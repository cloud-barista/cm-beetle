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
func CreateVMInfraWithDefaults(nsId string, infraModel *cloudmodel.TbMciDynamicReq) (cloudmodel.VmInfraInfo, error) {

	// Initialize Tumblebug client
	tbApiConfig := tbclient.ApiConfig{
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
		RestUrl:  config.Tumblebug.RestUrl,
	}
	tbCli := tbclient.NewClient(tbApiConfig)
	// Convert the request model from 'cloudmodel.TbMciDynamicReq' to 'tbmodel.TbMciDynamicReq'
	infraModelConverted, err := modelconv.ConvertWithValidation[cloudmodel.TbMciDynamicReq, tbmodel.TbMciDynamicReq](*infraModel)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure model (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	vmInfraInfo, err := tbCli.CreateMciDynamic(nsId, infraModelConverted)
	if err != nil {
		log.Error().Err(err).Msgf("failed to migrate the infrastructure (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	// Convert the response model from 'tbmodel.TbMciInfo' to 'cloudmodel.VmInfraInfo'
	convertedVmInfraInfo, err := modelconv.ConvertWithValidation[tbmodel.TbMciInfo, cloudmodel.VmInfraInfo](vmInfraInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s)", nsId)
		return cloudmodel.VmInfraInfo{}, err
	}

	return convertedVmInfraInfo, nil
}

// CreateVMInfra creates a VM infrastructure for the computing infra migration
func CreateVMInfra(nsId string, targetInfraModel *cloudmodel.RecommendedVmInfra) (cloudmodel.VmInfraInfo, error) {

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

	// 2. Create a VM specification (vmSpec)
	// * Skip: No need to regenerate vmSpec in namespace

	// 3. Create a VM OS image (vmOsImage)
	// * Skip: No need to regenerate vmOsImage in namespace

	// 4. Create a virtual network (vNet)
	// Get vNet request body from the input infraModel
	vNetReq := targetInfraModel.TargetVNet
	log.Debug().Msgf("Creating a vNet (nsId: %s, vNetName: %s)", nsId, vNetReq.Name)
	log.Debug().Msgf("vNetReq: %+v", vNetReq)

	// Convert model from 'cloudmodel.TbVNetReq' to 'tbmodel.TbVNetReq'
	tbVNetReq, err := modelconv.ConvertWithValidation[cloudmodel.TbVNetReq, tbmodel.TbVNetReq](vNetReq)
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

	// Convert model from 'cloudmodel.TbSshKeyReq' to 'tbmodel.TbSshKeyReq'
	tbSshKeyReq, err := modelconv.ConvertWithValidation[cloudmodel.TbSshKeyReq, tbmodel.TbSshKeyReq](sshKeyReq)
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

	sgInfoList := []tbmodel.TbSecurityGroupInfo{}
	for _, sgReq := range sgReqList {
		// Create security group
		log.Debug().Msgf("Creating a security group (nsId: %s, sgReq.sgName: %s, sgReq.VNetId: %s, vNetInfo.vNetId: %s)",
			nsId, sgReq.Name, sgReq.VNetId, vNetInfo.Id)

		// Convert model from 'cloudmodel.TbSecurityGroupReq' to 'tbmodel.TbSecurityGroupReq'
		tbSgReq, err := modelconv.ConvertWithValidation[cloudmodel.TbSecurityGroupReq, tbmodel.TbSecurityGroupReq](sgReq)
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

	// Convert model from 'cloudmodel.TbSecurityGroupReq' to 'tbmodel.TbSecurityGroupReq'
	tbMciReq, err := modelconv.ConvertWithValidation[cloudmodel.TbMciReq, tbmodel.TbMciReq](mciReq)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert SSH key request (nsId: %s)", nsId)
		return emptyRet, err
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

	// Convert the response model from 'tbmodel.TbMciInfo' to 'cloudmodel.TbMciInfo'
	mciInfoConverted, err := modelconv.ConvertWithValidation[tbmodel.TbMciInfo, cloudmodel.TbMciInfo](mciInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s)", nsId)
		return emptyRet, err
	}
	var temp cloudmodel.VmInfraInfo
	temp.TbMciInfo = mciInfoConverted

	// return emptyRet, fmt.Errorf("CreateVMInfra is not implemented yet")
	return temp, nil
}

// List all migrated VM infrastructures
func ListAllVMInfraInfo(nsId string) (cloudmodel.MciInfoList, error) {

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

	return convertedVmInfraInfoList, nil
}

// Get all migrated VM infrastructures
func ListVMInfraIDs(nsId string, option string) (cloudmodel.IdList, error) {

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

	return idList, nil
}

// Get the migrated VM infrastructure
func GetVMInfra(nsId, infraId string) (cloudmodel.TbMciInfo, error) {

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
		return cloudmodel.TbMciInfo{}, err
	}

	// Convert the response model from 'tbmodel.TbMciInfo' to 'cloudmodel.TbMciInfo'
	convertedVmInfaInfo, err := modelconv.ConvertWithValidation[tbmodel.TbMciInfo, cloudmodel.TbMciInfo](vmInfaInfo)
	if err != nil {
		log.Error().Err(err).Msgf("failed to convert the multi-cloud infrastructure info (nsId: %s, infraId: %s)", nsId, infraId)
		return cloudmodel.TbMciInfo{}, err
	}

	return convertedVmInfaInfo, nil
}

// Delete the migrated VM infrastructure
func DeleteVMInfra(nsId, infraId, option string) (common.SimpleMsg, error) {

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

	//3. Delete security groups
	// Collect unique security group IDs from all VMs
	sgIdMap := make(map[string]struct{})
	for _, vm := range mciInfo.Vm {
		for _, sgId := range vm.SecurityGroupIds {
			sgIdMap[sgId] = struct{}{}
		}
	}
	// Delete all security groups
	for sgId := range sgIdMap {
		msg, err := tbCli.DeleteSecurityGroup(nsId, sgId)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to delete security group (nsId: %s, sgId: %s)", nsId, sgId)
			// Continue deleting other resources even if this fails
		}
		log.Debug().Msgf("Security group deleted (nsId: %s, sgId: %s, msg: %s)", nsId, sgId, msg)
	}

	// 4. Delete SSH Key
	// Collect unique SSH Key IDs from all VMs
	sshKeyIdMap := make(map[string]struct{})
	for _, vm := range mciInfo.Vm {
		sshKeyIdMap[vm.SshKeyId] = struct{}{}
	}
	// Delete all SSH Key
	for sshKeyId := range sshKeyIdMap {
		// Delete SSH Key
		log.Debug().Msgf("Deleting SSH key (nsId: %s, sshKeyId: %s)", nsId, sshKeyId)
		msg, err := tbCli.DeleteSshKey(nsId, sshKeyId)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to delete SSH key (nsId: %s, sshKeyId: %s)", nsId, sshKeyId)
			// Continue deleting other resources even if this fails
		}
		log.Debug().Msgf("SSH key deleted (nsId: %s, sshKeyId: %s, msg: %s)", nsId, sshKeyId, msg)
	}

	// 5. Delete vNets
	// Collect unique vNet IDs from all VMs
	vNetIdMap := make(map[string]struct{})
	for _, vm := range mciInfo.Vm {
		vNetIdMap[vm.VNetId] = struct{}{}
	}
	// Delete all vNet
	for vNetId := range vNetIdMap {
		log.Debug().Msgf("Deleting VNet (nsId: %s, vNetId: %s, action: %s)", nsId, vNetId, "withsubnets")
		msg, err := tbCli.DeleteVNet(nsId, vNetId, "withsubnets")
		if err != nil {
			log.Warn().Err(err).Msgf("failed to delete VNet (nsId: %s, vNetId: %s)", nsId, vNetId)
			// Continue deleting other resources even if this fails
		}
		log.Debug().Msgf("VNet deleted (nsId: %s, vNetId: %s, msg:%s)", nsId, vNetId, msg)
	}

	// Sleep for a while to ensure all resources are deleted
	time.Sleep(15 * time.Second)

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

	return ret, nil
}

func validateTargeVmtInfraModel(nsId string, targetVmInfraModel *cloudmodel.RecommendedVmInfra) error {

	// * 1. Validate that name fieleds are not empty
	if targetVmInfraModel == nil {
		return fmt.Errorf("target infrastructure model is nil")
	}
	if targetVmInfraModel.TargetVmInfra.Name == "" { // MCI name
		return fmt.Errorf("target VM infrastructure name is empty")
	}
	if targetVmInfraModel.TargetVNet.Name == "" {
		return fmt.Errorf("target VNet name is empty")
	}
	if targetVmInfraModel.TargetSshKey.Name == "" {
		return fmt.Errorf("target SSH key name is empty")
	}
	for _, sg := range targetVmInfraModel.TargetSecurityGroupList {
		if sg.Name == "" {
			return fmt.Errorf("target security group name is empty")
		}
	}

	// * 2. Validate that the names or IDs are matched in the model
	// Check if the each VM's vNetId matches the target VNet name
	for _, vm := range targetVmInfraModel.TargetVmInfra.Vm {
		if vm.VNetId != targetVmInfraModel.TargetVNet.Name {
			return fmt.Errorf("target VM infrastructure vNetId (%s) does not match target VNet name (%s)",
				vm.VNetId, targetVmInfraModel.TargetVNet.Name)
		}
	}

	// Check if each VM's SshKeyId matches the target SSH key name
	for _, vm := range targetVmInfraModel.TargetVmInfra.Vm {
		if vm.SshKeyId != targetVmInfraModel.TargetSshKey.Name {
			return fmt.Errorf("target VM infrastructure SshKeyId (%s) does not match target SSH key name (%s)",
				vm.SshKeyId, targetVmInfraModel.TargetSshKey.Name)
		}
	}

	// Check if each VM's spec is contained in the target VM spec list
	for _, vm := range targetVmInfraModel.TargetVmInfra.Vm {
		found := false
		for _, vmSpec := range targetVmInfraModel.TargetVmSpecList {
			if vm.SpecId == vmSpec.Id {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("target VM infrastructure vmSpecId (%s) does not match any target VM spec ID in the list",
				vm.SpecId)
		}
	}

	// Check if each VM's OS image is contained in the target VM OS image list
	for _, vm := range targetVmInfraModel.TargetVmInfra.Vm {
		found := false
		for _, vmOsImage := range targetVmInfraModel.TargetVmOsImageList {
			if vm.ImageId == vmOsImage.Id {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("target VM infrastructure vmOsImageId (%s) does not match any target VM OS image ID in the list",
				vm.ImageId)
		}
	}

	// Check if each security group name is contained in the target security group list
	for _, vm := range targetVmInfraModel.TargetVmInfra.Vm {
		found := false
		for _, sgId := range vm.SecurityGroupIds {
			// Check if the security group name matches any target security group name
			for _, targetSg := range targetVmInfraModel.TargetSecurityGroupList {
				if sgId == targetSg.Name {
					found = true
					break
				}
			}
			if !found {
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
	for _, vmOsImage := range targetVmInfraModel.TargetVmOsImageList {
		_, err := tbCli.ReadVmOsImage("system", vmOsImage.Id)
		if err != nil {
			log.Error().Err(err).Msgf("failed to read VM OS image (nsId: %s, vmOsImageId: %s)", nsId, vmOsImage.Id)
			return fmt.Errorf("failed to read VM OS image (nsId: %s, vmOsImageId: %s): %w", nsId, vmOsImage.Id, err)
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
