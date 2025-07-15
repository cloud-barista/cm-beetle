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

// IsReady checks if Tumblebug is ready
func (c *TumblebugClient) IsReady() (bool, error) {
	log.Debug().Msg("Checking Tumblebug readiness")

	method := "GET"
	url := fmt.Sprintf("%s/readyz", c.restUrl)

	reqBody := common.NoBody
	resBody := tbmodel.SimpleMsg{}

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
		log.Error().Err(err).Msg("Tumblebug readiness check failed")
		return false, err
	}

	log.Debug().Msg("Tumblebug is ready")
	return true, nil
}
