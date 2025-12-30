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

// CreateMci creates a new MCI (Multi-Cloud Image) in the specified namespace
func (s *Session) CreateMci(nsId string, reqBody tbmodel.MciReq) (tbmodel.MciInfo, error) {
	log.Debug().Msg("Creating MCI")

	emptyRet := tbmodel.MciInfo{}
	url := fmt.Sprintf("/ns/%s/mci", nsId)
	resBody := tbmodel.MciInfo{}

	resp, err := s.
		SetBody(&reqBody).
		SetResult(&resBody).
		Post(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create MCI")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msg("Created MCI successfully")
	return resBody, nil
}

// CreateMciDynamic creates a new MCI (Multi-Cloud Image) with defaults in the specified namespace
func (s *Session) CreateMciDynamic(nsId string, reqBody tbmodel.MciDynamicReq) (tbmodel.MciInfo, error) {
	log.Debug().Msg("Creating MCI with defaults")

	emptyRet := tbmodel.MciInfo{}

	// Increased timeout to 45 minutes for all operations
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Minute)
	defer cancel()

	url := fmt.Sprintf("/ns/%s/mciDynamic", nsId)
	resBody := tbmodel.MciInfo{}

	resp, err := s.
		SetContext(ctx).
		SetBody(&reqBody).
		SetResult(&resBody).
		Post(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to create MCI with defaults")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msg("Created MCI with defaults successfully")
	return resBody, nil

}

type TbMciInfoList struct {
	Mci []tbmodel.MciInfo `json:"mci"`
}

// ReadAllMci retrieves all MCIs (Multi-Cloud Images) in the specified namespace
func (s *Session) ReadAllMci(nsId string) (TbMciInfoList, error) {
	log.Debug().Msg("Retrieving all MCIs")
	var emptyRet = TbMciInfoList{}

	url := fmt.Sprintf("/ns/%s/mci", nsId)
	resBody := TbMciInfoList{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve all MCIs")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved all MCIs (count: %d) successfully", len(resBody.Mci))
	return resBody, nil
}

// ReadMci retrieves information about a specific MCI (Multi-Cloud Image) in the specified namespace
func (s *Session) ReadMci(nsId, mciId string) (tbmodel.MciInfo, error) {
	log.Debug().Msg("Retrieving MCI")

	var emptyRet = tbmodel.MciInfo{}

	url := fmt.Sprintf("/ns/%s/mci/%s", nsId, mciId)
	resBody := tbmodel.MciInfo{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve MCI")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved MCI (mciId: %s) successfully", resBody.Id)
	return resBody, nil
}

func (s *Session) ReadMciAccessInfo(nsId, mciId, option, accessInfoOption string) (tbmodel.MciAccessInfo, error) {
	log.Debug().Msg("Retrieving MCI Access Info")

	var emptyRet tbmodel.MciAccessInfo

	url := fmt.Sprintf("/ns/%s/mci/%s", nsId, mciId)
	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
		if accessInfoOption != "" {
			url += fmt.Sprintf("&accessInfoOption=%s", accessInfoOption)
		}
	}

	resBody := tbmodel.MciAccessInfo{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve MCI Access Info")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved MCI Access Info (mciId: %s) successfully", mciId)
	return resBody, nil
}

func (s *Session) ReadMciIDs(nsId string) (tbmodel.IdList, error) {
	log.Debug().Msg("Retrieving MCI IDs")

	emptyRet := tbmodel.IdList{}

	url := fmt.Sprintf("/ns/%s/mci", nsId)

	option := "id" // Use 'ids' option to retrieve only IDs
	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
	}

	resBody := tbmodel.IdList{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve MCI IDs")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved MCI IDs (count: %d) successfully", len(resBody.IdList))
	return resBody, nil

}

// DeleteMci deletes a specific MCI (Multi-Cloud Image) in the specified namespace
func (s *Session) DeleteMci(nsId, mciId, option string) (tbmodel.IdList, error) {
	log.Debug().Msg("Deleting MCI")

	// Increased timeout to 45 minutes for all operations
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Minute)
	defer cancel()

	emptyRet := tbmodel.IdList{}

	url := fmt.Sprintf("/ns/%s/mci/%s", nsId, mciId)

	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
	}

	resBody := tbmodel.IdList{}

	resp, err := s.
		SetContext(ctx).
		SetResult(&resBody).
		Delete(url)

	if err != nil {
		log.Error().Err(err).Msg("Failed to delete MCI")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Deleted MCI (mciId: %s) successfully", mciId)
	return resBody, nil
}

// MciRecommendSpec finds appropriate VM specs by filtering and prioritzing.
func (s *Session) MciRecommendSpec(planToSearchProperVm string) ([]tbmodel.SpecInfo, error) {
	log.Debug().Msg("MCI Recommend Spec")

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
