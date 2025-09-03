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

// CreateSshKey creates a new SSH Key in the specified namespace
func (c *TumblebugClient) CreateSshKey(nsId string, reqBody tbmodel.SshKeyReq) (tbmodel.SshKeyInfo, error) {
	log.Debug().Msg("Creating SSH Key")

	emptyRet := tbmodel.SshKeyInfo{}

	method := "POST"
	url := fmt.Sprintf("%s/ns/%s/resources/sshKey", c.restUrl, nsId)

	resBody := tbmodel.SshKeyInfo{}

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
		log.Error().Err(err).Msg("Failed to create SSH Key")
		return emptyRet, err
	}

	log.Debug().Msg("Created SSH Key successfully")
	return resBody, nil
}

// ReadSshKey retrieves information about a specific SSH Key in the specified namespace
func (c *TumblebugClient) ReadSshKey(nsId, sshKeyId string) (tbmodel.SshKeyInfo, error) {
	log.Debug().Msg("Retrieving SSH Key")

	var emptyRet = tbmodel.SshKeyInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/resources/sshKey/%s", c.restUrl, nsId, sshKeyId)

	reqBody := common.NoBody
	resBody := tbmodel.SshKeyInfo{}

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
