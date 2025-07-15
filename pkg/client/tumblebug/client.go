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
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/go-resty/resty/v2"
)

// * [Note]
// * This Tumblebug Client is used to interact with the CB-Tumblebug API.
// * The Client contains the Tumblebug APIs required for computing infrastructure migration.
// * Other APIs can be added as needed.

// TumblebugClient is the structure for Tumblebug client
type TumblebugClient struct {
	client   *resty.Client
	restUrl  string
	username string
	password string
}

// ApiConfig holds the configuration for Tumblebug API
type ApiConfig struct {
	RestUrl  string
	Username string
	Password string
}

// NewClient creates a new Tumblebug client
func NewClient(apiConfig ApiConfig) *TumblebugClient {
	client := resty.New()
	client.SetBasicAuth(apiConfig.Username, apiConfig.Password)

	return &TumblebugClient{
		client:   client,
		restUrl:  apiConfig.RestUrl,
		username: apiConfig.Username,
		password: apiConfig.Password,
	}
}

// NewDefaultClient creates a new Tumblebug client using global config
func NewDefaultClient() *TumblebugClient {
	client := resty.New()
	client.SetBasicAuth(config.Tumblebug.API.Username, config.Tumblebug.API.Password)

	return &TumblebugClient{
		client:   client,
		restUrl:  config.Tumblebug.RestUrl,
		username: config.Tumblebug.API.Username,
		password: config.Tumblebug.API.Password,
	}
}
