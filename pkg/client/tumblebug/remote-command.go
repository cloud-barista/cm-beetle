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

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/cloud-barista/cm-beetle/pkg/ratelimit"
	"github.com/rs/zerolog/log"
)

// RemoteCommandToInfra sends a remote command to the nodes of the specified Infra via
// Tumblebug's remote command API (POST /ns/{nsId}/cmd/infra/{infraId}).
// Pass nodeGroupId and/or nodeId to target a specific nodeGroup or node; pass "" to target all nodes.
func (s *Session) RemoteCommandToInfra(nsId, infraId, nodeGroupId, nodeId string, reqBody tbmodel.InfraCmdReq) (tbmodel.InfraSshCmdResultForAPI, error) {
	log.Debug().Msgf("Sending remote command to Infra (nsId: %s, infraId: %s, nodeGroupId: %s, nodeId: %s)", nsId, infraId, nodeGroupId, nodeId)

	emptyRet := tbmodel.InfraSshCmdResultForAPI{}
	resBody := tbmodel.InfraSshCmdResultForAPI{}

	req := s.SetBody(&reqBody).SetResult(&resBody)
	if nodeGroupId != "" {
		req = req.SetQueryParam("nodeGroupId", nodeGroupId)
	}
	if nodeId != "" {
		req = req.SetQueryParam("nodeId", nodeId)
	}

	resp, err := req.Post(fmt.Sprintf("/ns/%s/cmd/infra/%s", nsId, infraId))
	if err != nil {
		log.Error().Err(err).Msg("Failed to send remote command to Infra")
		return emptyRet, err
	}
	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return emptyRet, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		return emptyRet, fmt.Errorf("API request failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	log.Debug().Msgf("Sent remote command to Infra (infraId: %s) successfully", infraId)
	return resBody, nil
}
