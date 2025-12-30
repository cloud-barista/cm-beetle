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

// CreateNamespace creates a new namespace in Tumblebug
func (s *Session) CreateNamespace(nsReq tbmodel.NsReq) (tbmodel.NsInfo, error) {
	log.Debug().Msg("Creating new namespace")

	var resBody tbmodel.NsInfo
	resp, err := s.
		SetBody(nsReq).
		SetResult(&resBody).
		Post("/ns")

	if err != nil {
		log.Error().Err(err).Msg("Failed to create namespace")
		return tbmodel.NsInfo{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to create namespace")
		return tbmodel.NsInfo{}, err
	}

	log.Debug().Msgf("Namespace (nsId: %s) created successfully", resBody.Id)
	return resBody, nil
}

// ReadNamespace retrieves information about a specific namespace
func (s *Session) ReadNamespace(nsId string) (tbmodel.NsInfo, error) {
	log.Debug().Msg("Retrieving namespace information")

	var resBody tbmodel.NsInfo
	resp, err := s.
		SetResult(&resBody).
		Get(fmt.Sprintf("/ns/%s", nsId))

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve namespace")
		return tbmodel.NsInfo{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to retrieve namespace")
		return tbmodel.NsInfo{}, err
	}

	log.Debug().Msgf("Retrieved namespace (nsId: %s) successfully", resBody.Id)
	return resBody, nil
}
