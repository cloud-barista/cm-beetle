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

// AsyncJobResponse represents the response for asynchronous job submission.
// Returned with HTTP 202 Accepted status for long-running operations.
type AsyncJobResponse struct {
	// ReqID is the unique identifier for tracking the request status.
	// This is the same as the X-Request-Id header value.
	ReqID string `json:"reqId" example:"1706500000000000000"`
	// Status is the current status of the request (Handling, Success, Error).
	Status string `json:"status" example:"Handling"`
	// StatusURL is the relative URL to check the request status.
	StatusURL string `json:"statusUrl" example:"/beetle/request/1706500000000000000"`
}
