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

// CreateVNet creates a new Virtual Network (VNet) in the specified namespace
func (c *TumblebugClient) CreateVNet(nsId string, reqBody tbmodel.VNetReq) (tbmodel.VNetInfo, error) {
	log.Debug().Msg("Creating Virtual Network")

	emptyRet := tbmodel.VNetInfo{}

	method := "POST"
	url := fmt.Sprintf("%s/ns/%s/resources/vNet", c.restUrl, nsId)

	resBody := tbmodel.VNetInfo{}

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
		log.Error().Err(err).Msg("Failed to create VNet")
		return emptyRet, err
	}

	log.Debug().Msg("Created VNet successfully")
	return resBody, nil
}

// ReadVNet retrieves information about a specific Virtual Network (VNet) in the specified namespace
func (c *TumblebugClient) ReadVNet(nsId, vNetId string) (tbmodel.VNetInfo, error) {
	log.Debug().Msg("Retrieving Virtual Network")

	var emptyRet = tbmodel.VNetInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/resources/vNet/%s", c.restUrl, nsId, vNetId)

	reqBody := common.NoBody
	resBody := tbmodel.VNetInfo{}

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
