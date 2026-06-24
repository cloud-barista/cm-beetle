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

// Package tbclient provides client functions to interact with CB-Tumblebug API
package tbclient

import (
	"fmt"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// NLB Management APIs
// ============================================================================

// CreateNlb creates a new NLB in the specified namespace and infra.
func (s *Session) CreateNlb(nsId, infraId string, req tbmodel.NLBReq) (tbmodel.NLBInfo, error) {
	log.Debug().Str("nsId", nsId).Str("infraId", infraId).Msg("Creating NLB")

	var resBody tbmodel.NLBInfo
	resp, err := s.
		SetBody(req).
		SetResult(&resBody).
		Post(fmt.Sprintf("/ns/%s/infra/%s/nlb", nsId, infraId))

	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("infraId", infraId).Msg("Failed to create NLB")
		return tbmodel.NLBInfo{}, err
	}
	if resp.IsError() {
		err := fmt.Errorf("API error %s: %s", resp.Status(), resp.Body())
		log.Error().Err(err).Msg("Failed to create NLB")
		return tbmodel.NLBInfo{}, err
	}

	log.Debug().Str("nsId", nsId).Str("infraId", infraId).Msg("NLB created successfully")
	return resBody, nil
}

// NLBListResponse is the response body for GET /ns/{nsId}/infra/{infraId}/nlb.
type NLBListResponse struct {
	NLB []tbmodel.NLBInfo `json:"nlb"`
}

// ListNlbs retrieves all NLBs in the specified namespace and infra.
func (s *Session) ListNlbs(nsId, infraId string) (NLBListResponse, error) {
	log.Debug().Str("nsId", nsId).Str("infraId", infraId).Msg("Listing NLBs")

	var resBody NLBListResponse
	resp, err := s.
		SetResult(&resBody).
		Get(fmt.Sprintf("/ns/%s/infra/%s/nlb", nsId, infraId))

	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("infraId", infraId).Msg("Failed to list NLBs")
		return NLBListResponse{}, err
	}
	if resp.IsError() {
		err := fmt.Errorf("API error %s: %s", resp.Status(), resp.Body())
		log.Error().Err(err).Msg("Failed to list NLBs")
		return NLBListResponse{}, err
	}

	log.Debug().Str("nsId", nsId).Str("infraId", infraId).Msg("NLBs listed successfully")
	return resBody, nil
}

// GetNlb retrieves a specific NLB.
func (s *Session) GetNlb(nsId, infraId, nlbId string) (tbmodel.NLBInfo, error) {
	log.Debug().Str("nsId", nsId).Str("infraId", infraId).Str("nlbId", nlbId).Msg("Getting NLB")

	var resBody tbmodel.NLBInfo
	resp, err := s.
		SetResult(&resBody).
		Get(fmt.Sprintf("/ns/%s/infra/%s/nlb/%s", nsId, infraId, nlbId))

	if err != nil {
		log.Error().Err(err).Str("nlbId", nlbId).Msg("Failed to get NLB")
		return tbmodel.NLBInfo{}, err
	}
	if resp.IsError() {
		err := fmt.Errorf("API error %s: %s", resp.Status(), resp.Body())
		log.Error().Err(err).Msg("Failed to get NLB")
		return tbmodel.NLBInfo{}, err
	}

	log.Debug().Str("nlbId", nlbId).Msg("NLB retrieved successfully")
	return resBody, nil
}

// DeleteNlb deletes a specific NLB.
func (s *Session) DeleteNlb(nsId, infraId, nlbId string) error {
	log.Debug().Str("nsId", nsId).Str("infraId", infraId).Str("nlbId", nlbId).Msg("Deleting NLB")

	resp, err := s.Delete(fmt.Sprintf("/ns/%s/infra/%s/nlb/%s", nsId, infraId, nlbId))
	if err != nil {
		log.Error().Err(err).Str("nlbId", nlbId).Msg("Failed to delete NLB")
		return err
	}
	if resp.IsError() {
		err := fmt.Errorf("API error %s: %s", resp.Status(), resp.Body())
		log.Error().Err(err).Msg("Failed to delete NLB")
		return err
	}

	log.Debug().Str("nlbId", nlbId).Msg("NLB deleted successfully")
	return nil
}
