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

// ReadRegionInfo reads region information in a specific provider
func (c *TumblebugClient) ReadRegionInfo(providerName string, regionName string) (tbmodel.RegionDetail, error) {
	log.Debug().Msg("Read Region Info")

	emptyRet := tbmodel.RegionDetail{}

	method := "GET"
	url := fmt.Sprintf("%s/provider/%s/region/%s", c.restUrl, providerName, regionName)

	// Request body
	tbReqt := common.NoBody
	tbResp := tbmodel.RegionDetail{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(tbReqt),
		&tbReqt,
		&tbResp,
		common.VeryShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved region (regionId: %s) successfully", tbResp.RegionId)
	return tbResp, nil
}
