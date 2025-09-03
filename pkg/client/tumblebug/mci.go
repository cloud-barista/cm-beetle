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
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/rs/zerolog/log"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// * [Note]
// * This Tumblebug Client is used to interact with the CB-Tumblebug API.
// * The Client contains the Tumblebug APIs required for computing infrastructure migration.
// * Other APIs can be added as needed.

// CreateMci creates a new MCI (Multi-Cloud Image) in the specified namespace
func (c *TumblebugClient) CreateMci(nsId string, reqBody tbmodel.MciReq) (tbmodel.MciInfo, error) {
	log.Debug().Msg("Creating MCI")

	emptyRet := tbmodel.MciInfo{}

	method := "POST"
	url := fmt.Sprintf("%s/ns/%s/mci", c.restUrl, nsId)

	resBody := tbmodel.MciInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(reqBody),
		&reqBody,
		&resBody,
		common.ShortDuration,
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create MCI")
		return emptyRet, err
	}

	log.Debug().Msg("Created MCI successfully")
	return resBody, nil
}

// CreateMciDynamic creates a new MCI (Multi-Cloud Image) with defaults in the specified namespace
func (c *TumblebugClient) CreateMciDynamic(nsId string, reqBody tbmodel.MciDynamicReq) (tbmodel.MciInfo, error) {
	log.Debug().Msg("Creating MCI with defaults")

	emptyRet := tbmodel.MciInfo{}

	c.client.SetTimeout(45 * time.Minute) // Increased timeout to 45 minutes for all operations

	method := "POST"
	url := fmt.Sprintf("%s/ns/%s/mciDynamic", c.restUrl, nsId)

	resBody := tbmodel.MciInfo{}

	err := common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(reqBody),
		&reqBody,
		&resBody,
		common.ShortDuration,
	)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create MCI with defaults")
		return emptyRet, err
	}

	log.Debug().Msg("Created MCI with defaults successfully")
	return resBody, nil

}

type TbMciInfoList struct {
	Mci []tbmodel.MciInfo `json:"mci"`
}

// ReadAllMci retrieves all MCIs (Multi-Cloud Images) in the specified namespace
func (c *TumblebugClient) ReadAllMci(nsId string) (TbMciInfoList, error) {
	log.Debug().Msg("Retrieving all MCIs")
	var emptyRet = TbMciInfoList{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/mci", c.restUrl, nsId)
	reqBody := common.NoBody
	resBody := TbMciInfoList{}
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
		log.Error().Err(err).Msg("Failed to retrieve all MCIs")
		return emptyRet, err
	}
	log.Debug().Msgf("Retrieved all MCIs (count: %d) successfully", len(resBody.Mci))
	return resBody, nil
}

// ReadMci retrieves information about a specific MCI (Multi-Cloud Image) in the specified namespace
func (c *TumblebugClient) ReadMci(nsId, mciId string) (tbmodel.MciInfo, error) {
	log.Debug().Msg("Retrieving MCI")

	var emptyRet = tbmodel.MciInfo{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/mci/%s", c.restUrl, nsId, mciId)

	reqBody := common.NoBody
	resBody := tbmodel.MciInfo{}

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
		log.Error().Err(err).Msg("Failed to retrieve MCI")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved MCI (mciId: %s) successfully", resBody.Id)
	return resBody, nil
}

func (c *TumblebugClient) ReadMciIDs(nsId string) (tbmodel.IdList, error) {
	log.Debug().Msg("Retrieving MCI IDs")

	emptyRet := tbmodel.IdList{}

	method := "GET"
	url := fmt.Sprintf("%s/ns/%s/mci", c.restUrl, nsId)

	option := "id" // Use 'ids' option to retrieve only IDs
	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
	}

	reqBody := common.NoBody
	resBody := tbmodel.IdList{}

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
		log.Error().Err(err).Msg("Failed to retrieve MCI IDs")
		return emptyRet, err
	}

	log.Debug().Msgf("Retrieved MCI IDs (count: %d) successfully", len(resBody.IdList))
	return resBody, nil

}

// DeleteMci deletes a specific MCI (Multi-Cloud Image) in the specified namespace
func (c *TumblebugClient) DeleteMci(nsId, mciId, option string) (tbmodel.IdList, error) {
	log.Debug().Msg("Deleting MCI")

	c.client.SetTimeout(45 * time.Minute) // Increased timeout to 45 minutes for all operations

	emptyRet := tbmodel.IdList{}

	method := "DELETE"
	url := fmt.Sprintf("%s/ns/%s/mci/%s", c.restUrl, nsId, mciId)

	if option != "" {
		url += fmt.Sprintf("?option=%s", option)
	}

	reqBody := common.NoBody
	resBody := tbmodel.IdList{}

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
		log.Error().Err(err).Msg("Failed to delete MCI")
		return emptyRet, err
	}

	log.Debug().Msgf("Deleted MCI (mciId: %s) successfully", mciId)
	return resBody, nil
}

// MciRecommendSpec finds appropriate VM specs by filtering and prioritzing.
func (c *TumblebugClient) MciRecommendSpec(planToSearchProperVm string) ([]tbmodel.SpecInfo, error) {
	log.Debug().Msg("MCI Recommend Spec")

	var vmSpecInfoList = []tbmodel.SpecInfo{}
	var emptyRet = []tbmodel.SpecInfo{}

	// Lookup VM specs
	method := "POST"
	url := fmt.Sprintf("%s/recommendSpec", c.restUrl)

	// Request body
	reqRecommVm := new(tbmodel.RecommendSpecReq)
	err := json.Unmarshal([]byte(planToSearchProperVm), reqRecommVm)
	if err != nil {
		log.Error().Err(err).Msg("")
		return emptyRet, err
	}
	// log.Trace().Msgf("deployment plan for the VM recommendation: %+v", reqRecommVm)

	// Response body
	err = common.ExecuteHttpRequest(
		c.client,
		method,
		url,
		nil,
		common.SetUseBody(*reqRecommVm),
		reqRecommVm,
		&vmSpecInfoList,
		common.VeryShortDuration,
	)

	if err != nil {
		log.Error().Err(err).Msg("Failed to recommend VM specs")
		return emptyRet, err
	}

	log.Debug().Msgf("Found VM specs (count: %d) successfully", len(vmSpecInfoList))
	return vmSpecInfoList, nil
}
