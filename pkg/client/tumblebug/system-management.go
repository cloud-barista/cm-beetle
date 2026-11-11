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
	"net/http"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/ratelimit"
	"github.com/rs/zerolog/log"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// * [Note]
// * This Tumblebug Client is used to interact with the CB-Tumblebug API.
// * The Client contains the Tumblebug APIs required for computing infrastructure migration.
// * Other APIs can be added as needed.

// IsReady checks if Tumblebug is ready
func (s *Session) IsReady() (bool, *string, error) {
	log.Debug().Msg("Checking Tumblebug readiness")

	url := "/readyz"

	resBody := tbmodel.SimpleMsg{}

	resp, err := s.
		SetResult(&resBody).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("Tumblebug readiness check failed")
		return false, nil, err
	}
	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return false, nil, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		return false, nil, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msg("Tumblebug is ready")
	return true, &resBody.Message, nil
}

// CheckReadyz queries Tumblebug GET /readyz and returns readiness and initialization status.
func (s *Session) CheckReadyz() (tbmodel.ReadyzResponse, error) {
	var resBody tbmodel.ReadyzResponse
	resp, err := s.
		NoCache().
		SetResult(&resBody).
		Get("/readyz")
	if err != nil {
		return resBody, err
	}
	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return resBody, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		return resBody, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return resBody, nil
}

// TumblebugRequestSummary represents upstream tracking status returned by Tumblebug GET /request/{reqId}.
type TumblebugRequestSummary struct {
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
	Status        string    `json:"status"`
	ResponseData  any       `json:"responseData,omitempty"`
	ErrorResponse string    `json:"errorResponse,omitempty"`
}

// ReadRequest retrieves upstream request tracking details from Tumblebug GET /request/{reqId}.
func (s *Session) ReadRequest(reqId string) (TumblebugRequestSummary, error) {
	var resBody TumblebugRequestSummary
	if reqId == "" {
		return resBody, fmt.Errorf("request ID is required")
	}

	url := fmt.Sprintf("/request/%s", reqId)
	resp, err := s.
		SetCacheDuration(VeryShortDuration).
		SetResult(&resBody).
		Get(url)
	if err != nil {
		return resBody, err
	}
	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return resBody, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		return resBody, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return resBody, nil
}
