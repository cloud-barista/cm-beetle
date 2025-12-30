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

	"github.com/rs/zerolog/log"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// * [Note]
// * This Tumblebug Client is used to interact with the CB-Tumblebug API.
// * The Client contains the Tumblebug APIs required for computing infrastructure migration.
// * Other APIs can be added as needed.

// CreateVNet creates a new Virtual Network (VNet) in the specified namespace
func (s *Session) CreateVNet(nsId string, reqBody tbmodel.VNetReq) (tbmodel.VNetInfo, error) {
	log.Debug().Msg("Creating Virtual Network")

	emptyRet := tbmodel.VNetInfo{}

	url := fmt.Sprintf("/ns/%s/resources/vNet", nsId)

	resBody := tbmodel.VNetInfo{}

	resp, err := s.
		SetBody(reqBody).
		SetResult(&resBody).
		Post(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create VNet")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msg("Created VNet successfully")
	return resBody, nil
}

// ReadVNet retrieves information about a specific Virtual Network (VNet) in the specified namespace
func (s *Session) ReadVNet(nsId, vNetId string) (tbmodel.VNetInfo, error) {
	log.Debug().Msg("Retrieving Virtual Network")

	var emptyRet = tbmodel.VNetInfo{}

	url := fmt.Sprintf("/ns/%s/resources/vNet/%s", nsId, vNetId)

	resBody := tbmodel.VNetInfo{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve VNet")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved VNet (vNetId: %s) successfully", resBody.Id)
	return resBody, nil
}

func (s *Session) DeleteVNet(nsId, vNetId, action string) (tbmodel.SimpleMsg, error) {
	log.Debug().Msg("Deleting Virtual Network")

	emptyRet := tbmodel.SimpleMsg{}

	url := fmt.Sprintf("/ns/%s/resources/vNet/%s", nsId, vNetId)
	if action != "" {
		url += fmt.Sprintf("?action=%s", action)
	}

	resBody := tbmodel.SimpleMsg{}

	resp, err := s.
		SetResult(&resBody).
		Delete(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to delete VNet")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Deleted VNet (vNetId: %s) successfully", vNetId)
	return resBody, nil
}
