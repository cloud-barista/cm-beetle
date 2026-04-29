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
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// * [Note]
// * This Tumblebug Client is used to interact with the CB-Tumblebug API.
// * The Client contains the Tumblebug APIs required for computing infrastructure migration.
// * Other APIs can be added as needed.

// CreateInfra creates a new Infra (Infrastructure) in the specified namespace
func (s *Session) CreateInfra(nsId string, reqBody tbmodel.InfraReq) (tbmodel.InfraInfo, error) {
	log.Debug().Msg("Creating Infra")

	emptyRet := tbmodel.InfraInfo{}
	url := fmt.Sprintf("/ns/%s/infra", nsId)
	resBody := tbmodel.InfraInfo{}

	resp, err := s.
		SetBody(&reqBody).
		SetResult(&resBody).
		Post(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create Infra")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msg("Created Infra successfully")
	return resBody, nil
}

// CreateInfraDynamic creates a new Infra (Infrastructure) with defaults in the specified namespace
func (s *Session) CreateInfraDynamic(nsId string, reqBody tbmodel.InfraDynamicReq) (tbmodel.InfraInfo, error) {
	log.Debug().Msg("Creating Infra with defaults")

	emptyRet := tbmodel.InfraInfo{}

	// Increased timeout to 45 minutes for all operations
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Minute)
	defer cancel()

	url := fmt.Sprintf("/ns/%s/infraDynamic", nsId)
	resBody := tbmodel.InfraInfo{}

	resp, err := s.
		SetContext(ctx).
		SetBody(&reqBody).
		SetResult(&resBody).
		Post(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create Infra with defaults")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msg("Created Infra with defaults successfully")
	return resBody, nil

}

type TbInfraInfoList struct {
	Infra []tbmodel.InfraInfo `json:"infra"`
}

// ReadAllInfra retrieves all Infras (Infrastructure) in the specified namespace
func (s *Session) ReadAllInfra(nsId string) (TbInfraInfoList, error) {
	log.Debug().Msg("Retrieving all Infras")
	var emptyRet = TbInfraInfoList{}

	url := fmt.Sprintf("/ns/%s/infra", nsId)
	resBody := TbInfraInfoList{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve all Infras")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved all Infras (count: %d) successfully", len(resBody.Infra))
	return resBody, nil
}

// ReadInfra retrieves information about a specific Infra (Infrastructure) in the specified namespace
func (s *Session) ReadInfra(nsId, infraId string) (tbmodel.InfraInfo, error) {
	log.Debug().Msg("Retrieving Infra")

	var emptyRet = tbmodel.InfraInfo{}

	url := fmt.Sprintf("/ns/%s/infra/%s", nsId, infraId)
	resBody := tbmodel.InfraInfo{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve Infra")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved Infra (infraId: %s) successfully", resBody.Id)
	return resBody, nil
}

func (s *Session) ReadInfraAccessInfo(nsId, infraId, option, accessInfoOption string) (tbmodel.InfraAccessInfo, error) {
	log.Debug().Msg("Retrieving Infra Access Info")

	var emptyRet tbmodel.InfraAccessInfo

	url := fmt.Sprintf("/ns/%s/infra/%s", nsId, infraId)
	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
		if accessInfoOption != "" {
			url += fmt.Sprintf("&accessInfoOption=%s", accessInfoOption)
		}
	}

	resBody := tbmodel.InfraAccessInfo{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve Infra Access Info")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved Infra Access Info (infraId: %s) successfully", infraId)
	return resBody, nil
}

func (s *Session) ReadInfraIDs(nsId string) (tbmodel.IdList, error) {
	log.Debug().Msg("Retrieving Infra IDs")

	emptyRet := tbmodel.IdList{}

	url := fmt.Sprintf("/ns/%s/infra", nsId)

	option := "id" // Use 'ids' option to retrieve only IDs
	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
	}

	resBody := tbmodel.IdList{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve Infra IDs")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved Infra IDs (count: %d) successfully", len(resBody.IdList))
	return resBody, nil

}

// DeleteInfra deletes a specific Infra (Infrastructure) in the specified namespace
func (s *Session) DeleteInfra(nsId, infraId, option string) (tbmodel.IdList, error) {
	log.Debug().Msg("Deleting Infra")

	// Increased timeout to 45 minutes for all operations
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Minute)
	defer cancel()

	emptyRet := tbmodel.IdList{}

	url := fmt.Sprintf("/ns/%s/infra/%s", nsId, infraId)

	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
	}

	resBody := tbmodel.IdList{}

	resp, err := s.
		SetContext(ctx).
		SetResult(&resBody).
		Delete(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to delete Infra")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Deleted Infra (infraId: %s) successfully", infraId)
	return resBody, nil
}

// InfraRecommendSpec finds appropriate VM specs by filtering and prioritzing.
func (s *Session) InfraRecommendSpec(planToSearchProperVm string) ([]tbmodel.SpecInfo, error) {
	log.Debug().Msg("Infra Recommend Spec")

	var vmSpecInfoList = []tbmodel.SpecInfo{}
	var emptyRet = []tbmodel.SpecInfo{}

	// Lookup VM specs
	url := fmt.Sprintf("/recommendSpec")

	// Request body
	reqRecommVm := new(tbmodel.RecommendSpecReq)
	err := json.Unmarshal([]byte(planToSearchProperVm), reqRecommVm)
	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyRet, err
	}
	// log.Trace().Msgf("deployment plan for the VM recommendation: %+v", reqRecommVm)

	// Response body
	resp, err := s.
		SetBody(reqRecommVm).
		SetResult(&vmSpecInfoList).
		Post(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to recommend VM specs")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Found VM specs (count: %d) successfully", len(vmSpecInfoList))
	return vmSpecInfoList, nil
}
