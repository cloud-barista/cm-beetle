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

// ReadRegionInfo reads region information in a specific provider
func (s *Session) ReadRegionInfo(providerName string, regionName string) (tbmodel.RegionDetail, error) {
	log.Debug().Msg("Read Region Info")

	emptyRet := tbmodel.RegionDetail{}

	url := fmt.Sprintf("/provider/%s/region/%s", providerName, regionName)

	// Request body
	tbResp := tbmodel.RegionDetail{}

	resp, err := s.
		SetResult(&tbResp).
		Get(url)

	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved region (regionId: %s) successfully", tbResp.RegionId)
	return tbResp, nil
}

// ReadAvailableK8sVersion fetches available K8s versions for the given provider and region.
// providerName and regionName must be lowercase (e.g., "aws", "ap-northeast-2").
func (s *Session) ReadAvailableK8sVersion(providerName, regionName string) ([]tbmodel.K8sClusterVersionDetailAvailable, error) {
	log.Debug().Str("providerName", providerName).Str("regionName", regionName).Msg("Reading available K8s versions")

	emptyRet := []tbmodel.K8sClusterVersionDetailAvailable{}

	var resBody []tbmodel.K8sClusterVersionDetailAvailable

	resp, err := s.
		SetQueryParam("providerName", providerName).
		SetQueryParam("regionName", regionName).
		SetResult(&resBody).
		Get("/availableK8sVersion")

	if err != nil {
		log.Error().Err(err).Msg("Failed to read available K8s versions")
		return emptyRet, err
	}
	if resp.IsError() {
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Retrieved %d available K8s versions for %s/%s", len(resBody), providerName, regionName)
	return resBody, nil
}
