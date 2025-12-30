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

// ReadVmSpec retrieves information about a specific VM Spec in the specified namespace
func (s *Session) ReadVmSpec(nsId, vmSpecId string) (tbmodel.SpecInfo, error) {
	log.Debug().Msg("Retrieving VM Spec")

	var emptyRet = tbmodel.SpecInfo{}

	url := fmt.Sprintf("/ns/%s/resources/spec/%s", nsId, vmSpecId)

	resBody := tbmodel.SpecInfo{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve VM Spec")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved VM Spec (vmSpecId: %s) successfully", resBody.Id)
	return resBody, nil
}
