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

// Package model contains the request/response models for REST APIs
package model

import "github.com/cloud-barista/cm-beetle/transx"

// EncryptionTestRequest represents the request body for encryption test API.
// It contains the public key bundle and the plaintext model to be encrypted.
type EncryptionTestRequest struct {
	// PublicKeyBundle contains the public key obtained from GET /migration/data/encryptionKey
	PublicKeyBundle transx.PublicKeyBundle `json:"publicKeyBundle"`
	// Model is the plaintext DataMigrationModel to be encrypted
	Model transx.DataMigrationModel `json:"model"`
}
