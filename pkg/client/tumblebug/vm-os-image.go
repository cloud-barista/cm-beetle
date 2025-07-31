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

// ReadVmOsImage retrieves information about a specific VM OS Image in the specified namespace
func (c *TumblebugClient) ReadVmOsImage(nsId, vmOsImageId string) (tbmodel.TbImageInfo, error) {
	log.Debug().Msg("Retrieving VM OS Image")

	var emptyRet = tbmodel.TbImageInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/resources/image/%s", c.restUrl, nsId, vmOsImageId)

	reqBody := common.NoBody
	resBody := tbmodel.TbImageInfo{}

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
		log.Error().Err(err).Msg("Failed to retrieve VM OS Image")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved VM OS Image (vmOsImageId: %s) successfully", resBody.Id)
	return resBody, nil
}

// SearchVmOsImage searches VM OS images
func (c *TumblebugClient) SearchVmOsImage(nsId string, searchImageReq tbmodel.SearchImageRequest) (tbmodel.SearchImageResponse, error) {
	log.Debug().Msg("Search VM OS images")

	var emptyRet = tbmodel.SearchImageResponse{}

	method := "POST"
	url := fmt.Sprintf("%s/ns/%s/resources/searchImage", c.restUrl, nsId)

	// reqBody := common.NoBody
	resBody := tbmodel.SearchImageResponse{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(searchImageReq),
		&searchImageReq,
		&resBody,
		common.ShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to search VM OS Images")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved VM OS Images (count: %d) successfully", resBody.ImageCount)

	return resBody, nil
}
