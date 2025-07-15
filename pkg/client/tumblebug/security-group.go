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

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/rs/zerolog/log"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// * [Note]
// * This Tumblebug Client is used to interact with the CB-Tumblebug API.
// * The Client contains the Tumblebug APIs required for computing infrastructure migration.
// * Other APIs can be added as needed.

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
