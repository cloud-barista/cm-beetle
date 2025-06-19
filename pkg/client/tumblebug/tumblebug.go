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

// Package tbclient provides client functions to interact with CB-Tumblebug API
package tbclient

import (
	"fmt"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// SimpleMessage represents a simple message response
// type SimpleMessage struct {
// 	Message string `json:"message"`
// }

// TumblebugClient is the structure for Tumblebug client
type TumblebugClient struct {
	client   *resty.Client
	restUrl  string
	username string
	password string
}

// ApiConfig holds the configuration for Tumblebug API
type ApiConfig struct {
	RestUrl  string
	Username string
	Password string
}

// NewClient creates a new Tumblebug client
func NewClient(apiConfig ApiConfig) *TumblebugClient {
	client := resty.New()
	client.SetBasicAuth(apiConfig.Username, apiConfig.Password)

	return &TumblebugClient{
		client:   client,
		restUrl:  apiConfig.RestUrl,
		username: apiConfig.Username,
		password: apiConfig.Password,
	}
}

// NewDefaultClient creates a new Tumblebug client using global config
func NewDefaultClient() *TumblebugClient {
	client := resty.New()
	client.SetBasicAuth(config.Tumblebug.API.Username, config.Tumblebug.API.Password)

	return &TumblebugClient{
		client:   client,
		restUrl:  config.Tumblebug.RestUrl,
		username: config.Tumblebug.API.Username,
		password: config.Tumblebug.API.Password,
	}
}

// IsReady checks if Tumblebug is ready
func (c *TumblebugClient) IsReady() (bool, error) {
	log.Debug().Msg("Checking Tumblebug readiness")

	method := "GET"
	url := fmt.Sprintf("%s/readyz", c.restUrl)

	reqBody := common.NoBody
	resBody := tbmodel.SimpleMsg{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Tumblebug readiness check failed")
		return false, err
	}

	log.Debug().Msg("Tumblebug is ready")
	return true, nil
}

// CreateNamespace creates a new namespace in Tumblebug
func (c *TumblebugClient) CreateNamespace(nsReq tbmodel.NsReq) (tbmodel.NsInfo, error) {
	log.Debug().Msg("Creating new namespace")

	emptyRet := tbmodel.NsInfo{}

	method := "POST"
	url := fmt.Sprintf("%s/ns", c.restUrl)

	reqBody := nsReq
	resBody := tbmodel.NsInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(reqBody),
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create namespace")
		return emptyRet, err
	}

	log.Debug().Msgf("Namespace (nsId: %s) created successfully", resBody.Id)
	return resBody, nil
}

// ReadNamespace retrieves information about a specific namespace
func (c *TumblebugClient) ReadNamespace(nsId string) (tbmodel.NsInfo, error) {
	log.Debug().Msg("Retrieving namespace information")

	var emptyRet = tbmodel.NsInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s", c.restUrl, nsId)

	reqBody := common.NoBody
	resBody := tbmodel.NsInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve namespace")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved namespace (nsId: %s) successfully", resBody.Id)
	return resBody, nil
}

// CreateVNet creates a new Virtual Network (VNet) in the specified namespace
func (c *TumblebugClient) CreateVNet(nsId string, reqBody tbmodel.TbVNetReq) (tbmodel.TbVNetInfo, error) {
	log.Debug().Msg("Creating Virtual Network")

	emptyRet := tbmodel.TbVNetInfo{}

	method := "POST"
	url := fmt.Sprintf("%s/ns/%s/resources/vNet", c.restUrl, nsId)

	resBody := tbmodel.TbVNetInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(reqBody),
		&reqBody,
		&resBody,
		common.ShortDuration,
	)
	if err != nil {
		log.Error().Msg("Failed to create VNet")
		return emptyRet, err
	}

	log.Debug().Msg("Created VNet successfully")
	return resBody, nil
}

// ReadVNet retrieves information about a specific Virtual Network (VNet) in the specified namespace
func (c *TumblebugClient) ReadVNet(nsId, vNetId string) (tbmodel.TbVNetInfo, error) {
	log.Debug().Msg("Retrieving Virtual Network")

	var emptyRet = tbmodel.TbVNetInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/resources/vNet/%s", c.restUrl, nsId, vNetId)

	reqBody := common.NoBody
	resBody := tbmodel.TbVNetInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve VNet")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved VNet (vNetId: %s) successfully", resBody.Id)
	return resBody, nil
}

func (c *TumblebugClient) DeleteVNet(nsId, vNetId, action string) (tbmodel.SimpleMsg, error) {
	log.Debug().Msg("Deleting Virtual Network")

	emptyRet := tbmodel.SimpleMsg{}

	method := "DELETE"
	url := fmt.Sprintf("%s/ns/%s/resources/vNet/%s", c.restUrl, nsId, vNetId)

	if action != "" {
		url += fmt.Sprintf("?action=%s", action)
	}

	reqBody := common.NoBody
	resBody := tbmodel.SimpleMsg{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to delete VNet")
		return emptyRet, err
	}

	log.Debug().Msgf("Deleted VNet (vNetId: %s) successfully", vNetId)
	return resBody, nil
}

// CreateSecurityGroup creates a new Security Group in the specified namespace
func (c *TumblebugClient) CreateSecurityGroup(nsId string, reqBody tbmodel.TbSecurityGroupReq, option string) (tbmodel.TbSecurityGroupInfo, error) {
	log.Debug().Msg("Creating Security Group")

	var emptyRet = tbmodel.TbSecurityGroupInfo{}

	method := "POST"
	url := fmt.Sprintf("%s/ns/%s/resources/securityGroup", c.restUrl, nsId)

	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
	}

	var resBody tbmodel.TbSecurityGroupInfo
	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(reqBody),
		&reqBody,
		&resBody,
		common.ShortDuration,
	)
	if err != nil {
		log.Error().Msg("Failed to create Security Group")
		return emptyRet, err
	}

	log.Debug().Msg("Created Security Group successfully")
	return resBody, nil
}

// ReadSecurityGroup retrieves information about a specific Security Group in the specified namespace
func (c *TumblebugClient) ReadSecurityGroup(nsId, securityGroupId string) (tbmodel.TbSecurityGroupInfo, error) {
	log.Debug().Msg("Retrieving Security Group")

	var emptyRet = tbmodel.TbSecurityGroupInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/resources/securityGroup/%s", c.restUrl, nsId, securityGroupId)
	// /ns/{nsId}/resources/securityGroup/{securityGroupId}

	reqBody := common.NoBody
	resBody := tbmodel.TbSecurityGroupInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve Security Group")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved Security Group (securityGroupId: %s) successfully", resBody.Id)
	return resBody, nil
}

// DeleteSecurityGroup deletes a specific Security Group in the specified namespace
func (c *TumblebugClient) DeleteSecurityGroup(nsId, securityGroupId string) (tbmodel.SimpleMsg, error) {
	log.Debug().Msg("Deleting Security Group")

	emptyRet := tbmodel.SimpleMsg{}

	method := "DELETE"
	url := fmt.Sprintf("%s/ns/%s/resources/securityGroup/%s", c.restUrl, nsId, securityGroupId)

	reqBody := common.NoBody
	resBody := tbmodel.SimpleMsg{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to delete Security Group")
		return emptyRet, err
	}

	log.Debug().Msgf("Deleted Security Group (securityGroupId: %s) successfully", securityGroupId)
	return resBody, nil
}

// ReadVmSpec retrieves information about a specific VM Spec in the specified namespace
func (c *TumblebugClient) ReadVmSpec(nsId, vmSpecId string) (tbmodel.TbSpecInfo, error) {
	log.Debug().Msg("Retrieving VM Spec")

	var emptyRet = tbmodel.TbSpecInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/resources/spec/%s", c.restUrl, nsId, vmSpecId)

	reqBody := common.NoBody
	resBody := tbmodel.TbSpecInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve VM Spec")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved VM Spec (vmSpecId: %s) successfully", resBody.Id)
	return resBody, nil
}

// ReadVmOsImage retrieves information about a specific VM OS Image in the specified namespace
func (c *TumblebugClient) ReadVmOsImage(nsId, vmOsImageId string) (tbmodel.TbImageInfo, error) {
	log.Debug().Msg("Retrieving VM OS Image")

	var emptyRet = tbmodel.TbImageInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/resources/image/%s", c.restUrl, nsId, vmOsImageId)

	reqBody := common.NoBody
	resBody := tbmodel.TbImageInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve VM OS Image")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved VM OS Image (vmOsImageId: %s) successfully", resBody.Id)
	return resBody, nil
}

// CreateSshKey creates a new SSH Key in the specified namespace
func (c *TumblebugClient) CreateSshKey(nsId string, reqBody tbmodel.TbSshKeyReq) (tbmodel.TbSshKeyInfo, error) {
	log.Debug().Msg("Creating SSH Key")

	emptyRet := tbmodel.TbSshKeyInfo{}

	method := "POST"
	url := fmt.Sprintf("%s/ns/%s/resources/sshKey", c.restUrl, nsId)

	resBody := tbmodel.TbSshKeyInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(reqBody),
		&reqBody,
		&resBody,
		common.ShortDuration,
	)
	if err != nil {
		log.Error().Msg("Failed to create SSH Key")
		return emptyRet, err
	}

	log.Debug().Msg("Created SSH Key successfully")
	return resBody, nil
}

// ReadSshKey retrieves information about a specific SSH Key in the specified namespace
func (c *TumblebugClient) ReadSshKey(nsId, sshKeyId string) (tbmodel.TbSshKeyInfo, error) {
	log.Debug().Msg("Retrieving SSH Key")

	var emptyRet = tbmodel.TbSshKeyInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/resources/sshKey/%s", c.restUrl, nsId, sshKeyId)

	reqBody := common.NoBody
	resBody := tbmodel.TbSshKeyInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve SSH Key")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved SSH Key (sshKeyId: %s) successfully", resBody.Id)
	return resBody, nil
}

func (c *TumblebugClient) DeleteSshKey(nsId, sshKeyId string) (tbmodel.SimpleMsg, error) {
	log.Debug().Msg("Deleting SSH Key")

	emptyRet := tbmodel.SimpleMsg{}

	method := "DELETE"
	url := fmt.Sprintf("%s/ns/%s/resources/sshKey/%s", c.restUrl, nsId, sshKeyId)

	reqBody := common.NoBody
	resBody := tbmodel.SimpleMsg{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to delete SSH Key")
		return emptyRet, err
	}

	log.Debug().Msgf("Deleted SSH Key (sshKeyId: %s) successfully", sshKeyId)
	return resBody, nil
}

// CreateMci creates a new MCI (Multi-Cloud Image) in the specified namespace
func (c *TumblebugClient) CreateMci(nsId string, reqBody tbmodel.TbMciReq) (tbmodel.TbMciInfo, error) {
	log.Debug().Msg("Creating MCI")

	emptyRet := tbmodel.TbMciInfo{}

	method := "POST"
	url := fmt.Sprintf("%s/ns/%s/mci", c.restUrl, nsId)

	resBody := tbmodel.TbMciInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(reqBody),
		&reqBody,
		&resBody,
		common.ShortDuration,
	)
	if err != nil {
		log.Error().Msg("Failed to create MCI")
		return emptyRet, err
	}

	log.Debug().Msg("Created MCI successfully")
	return resBody, nil
}

// ReadMci retrieves information about a specific MCI (Multi-Cloud Image) in the specified namespace
func (c *TumblebugClient) ReadMci(nsId, mciId string) (tbmodel.TbMciInfo, error) {
	log.Debug().Msg("Retrieving MCI")

	var emptyRet = tbmodel.TbMciInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/mci/%s", c.restUrl, nsId, mciId)

	reqBody := common.NoBody
	resBody := tbmodel.TbMciInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve MCI")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved MCI (mciId: %s) successfully", resBody.Id)
	return resBody, nil
}

// DeleteMci deletes a specific MCI (Multi-Cloud Image) in the specified namespace
func (c *TumblebugClient) DeleteMci(nsId, mciId, option string) (tbmodel.IdList, error) {
	log.Debug().Msg("Deleting MCI")

	emptyRet := tbmodel.IdList{}

	method := "DELETE"
	url := fmt.Sprintf("%s/ns/%s/mci/%s", c.restUrl, nsId, mciId)

	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
	}

	reqBody := common.NoBody
	resBody := tbmodel.IdList{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to delete MCI")
		return emptyRet, err
	}

	log.Debug().Msgf("Deleted MCI (mciId: %s) successfully", mciId)
	return resBody, nil
}

func (c *TumblebugClient) DeleteSharedResources(nsId string) (tbmodel.IdList, error) {
	log.Debug().Msg("Deleting shared resources in namespace")

	emptyRet := tbmodel.IdList{}

	method := "DELETE"
	url := fmt.Sprintf("%s/ns/%s/sharedResources", c.restUrl, nsId)

	reqBody := common.NoBody
	resBody := tbmodel.IdList{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		false,
		&reqBody,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to delete shared resources")
		return emptyRet, err
	}

	log.Debug().Msgf("Deleted shared resources in namespace (nsId: %s) successfully", nsId)
	return resBody, nil
}
