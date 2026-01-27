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

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/rs/zerolog/log"
)

// * [Note]
// * This Tumblebug Client is used to interact with the CB-Tumblebug API.
// * The Client contains the Tumblebug APIs required for object storage management.
// *
// * All object storage models are imported from cb-tumblebug packages
// * to ensure consistency and reuse of Tumblebug data structures.

// ObjectStorageFeatureSupport represents feature support information for object storage
// TODO: Replace with tbmodel.ObjectStorageFeatureSupport when available in the Tumblebug version
type ObjectStorageFeatureSupport struct {
	CORS       bool `json:"cors" example:"true"`
	Versioning bool `json:"versioning" example:"true"`
}

// ObjectStorageSupportResponse represents CSP support information for object storage features
// TODO: Replace with tbmodel.ObjectStorageSupportResponse when available in the Tumblebug version
type ObjectStorageSupportResponse struct {
	ResourceType string                                 `json:"resourceType" example:"objectStorage"`
	Supports     map[string]ObjectStorageFeatureSupport `json:"supports"`
}

// ObjectStorageVersioning represents versioning configuration for object storage
// TODO: Replace with tbmodel.ObjectStorageVersioning when available in the Tumblebug version
type ObjectStorageVersioning struct {
	Status string `json:"status" example:"Enabled"` // Enabled, Suspended
}

// LocationConstraint represents the location constraint of a bucket
// TODO: Replace with tbmodel.LocationConstraint when available in the Tumblebug version
type LocationConstraint struct {
	Location string `json:"location" example:"ap-northeast-2"`
}

// ObjectStorageListResponse represents the list response for object storages
// TODO: Replace with tbmodel.ObjectStorageListResponse when available in the Tumblebug version
// Note: The old ObjectStorageListResponse (with Owner/Buckets) has been renamed to ObjectStorageListBucketsResponse
//
//	for backward compatibility. This is the new structure that returns an array of ObjectStorageInfo.
type ObjectStorageListResponse struct {
	ObjectStorage []tbmodel.ObjectStorageInfo `json:"objectStorage"`
}

// ============================================================================
// Object Storage Management APIs
// ============================================================================

// ListObjectStorages retrieves the list of all object storages (buckets) in a namespace
func (s *Session) ListObjectStorages(nsId string, option string, filterKey string, filterVal string) (ObjectStorageListResponse, error) {
	log.Debug().Msgf("Listing object storages in namespace: %s", nsId)

	var resBody ObjectStorageListResponse
	req := s.SetResult(&resBody)

	// Add optional query parameters
	if option != "" {
		req = req.SetQueryParam("option", option)
	}
	if filterKey != "" {
		req = req.SetQueryParam("filterKey", filterKey)
	}
	if filterVal != "" {
		req = req.SetQueryParam("filterVal", filterVal)
	}

	resp, err := req.Get(fmt.Sprintf("/ns/%s/resources/objectStorage", nsId))

	if err != nil {
		log.Error().Err(err).Msg("Failed to list object storages")
		return ObjectStorageListResponse{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to list object storages")
		return ObjectStorageListResponse{}, err
	}

	log.Debug().Msgf("Listed %d object storages successfully", len(resBody.ObjectStorage))
	return resBody, nil
}

// CreateObjectStorage creates a new object storage (bucket)
func (s *Session) CreateObjectStorage(nsId string, req tbmodel.ObjectStorageCreateRequest) (tbmodel.ObjectStorageInfo, error) {
	log.Debug().Msgf("Creating object storage: %s in namespace: %s", req.BucketName, nsId)

	var resBody tbmodel.ObjectStorageInfo
	resp, err := s.
		SetBody(req).
		SetResult(&resBody).
		Put(fmt.Sprintf("/ns/%s/resources/objectStorage", nsId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to create object storage: %s", req.BucketName)
		return tbmodel.ObjectStorageInfo{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to create object storage: %s", req.BucketName)
		return tbmodel.ObjectStorageInfo{}, err
	}

	log.Debug().Msgf("Object storage (%s) created successfully", req.BucketName)
	return resBody, nil
}

// GetObjectStorage retrieves details of an object storage (bucket)
func (s *Session) GetObjectStorage(nsId string, osId string) (tbmodel.ObjectStorageInfo, error) {
	log.Debug().Msgf("Retrieving object storage: %s in namespace: %s", osId, nsId)

	var resBody tbmodel.ObjectStorageInfo
	resp, err := s.
		SetResult(&resBody).
		Get(fmt.Sprintf("/ns/%s/resources/objectStorage/%s", nsId, osId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to retrieve object storage: %s", osId)
		return tbmodel.ObjectStorageInfo{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to retrieve object storage: %s", osId)
		return tbmodel.ObjectStorageInfo{}, err
	}

	log.Debug().Msgf("Retrieved object storage (%s) successfully", osId)
	return resBody, nil
}

// ExistObjectStorage checks the existence of an object storage (bucket)
func (s *Session) ExistObjectStorage(nsId string, osId string) (bool, error) {
	log.Debug().Msgf("Checking existence of object storage: %s in namespace: %s", osId, nsId)

	resp, err := s.Head(fmt.Sprintf("/ns/%s/resources/objectStorage/%s", nsId, osId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to check existence of object storage: %s", osId)
		return false, err
	}

	// HTTP Status OK is 200
	exists := resp.StatusCode() == 200
	log.Debug().Msgf("Object storage (%s) exists: %v", osId, exists)
	return exists, nil
}

// DeleteObjectStorage deletes an object storage (bucket)
func (s *Session) DeleteObjectStorage(nsId string, osId string) error {
	log.Debug().Msgf("Deleting object storage: %s in namespace: %s", osId, nsId)

	resp, err := s.Delete(fmt.Sprintf("/ns/%s/resources/objectStorage/%s", nsId, osId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete object storage: %s", osId)
		return err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to delete object storage: %s", osId)
		return err
	}

	log.Debug().Msgf("Object storage (%s) deleted successfully", osId)
	return nil
}

// ============================================================================
// CORS Management APIs
// ============================================================================

// GetObjectStorageCORS retrieves the CORS configuration of an object storage (bucket)
func (s *Session) GetObjectStorageCORS(nsId string, osId string) (tbmodel.ObjectStorageGetCorsResponse, error) {
	log.Debug().Msgf("Retrieving CORS configuration for object storage: %s in namespace: %s", osId, nsId)

	var resBody tbmodel.ObjectStorageGetCorsResponse
	resp, err := s.
		SetResult(&resBody).
		Get(fmt.Sprintf("/ns/%s/resources/objectStorage/%s/cors", nsId, osId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to retrieve CORS configuration for object storage: %s", osId)
		return tbmodel.ObjectStorageGetCorsResponse{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to retrieve CORS configuration for object storage: %s", osId)
		return tbmodel.ObjectStorageGetCorsResponse{}, err
	}

	log.Debug().Msgf("Retrieved CORS configuration for object storage (%s) successfully", osId)
	return resBody, nil
}

// SetObjectStorageCORS sets the CORS configuration of an object storage (bucket)
func (s *Session) SetObjectStorageCORS(nsId string, osId string, req tbmodel.ObjectStorageSetCorsRequest) error {
	log.Debug().Msgf("Setting CORS configuration for object storage: %s in namespace: %s", osId, nsId)

	resp, err := s.
		SetBody(req).
		Put(fmt.Sprintf("/ns/%s/resources/objectStorage/%s/cors", nsId, osId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to set CORS configuration for object storage: %s", osId)
		return err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to set CORS configuration for object storage: %s", osId)
		return err
	}

	log.Debug().Msgf("Set CORS configuration for object storage (%s) successfully", osId)
	return nil
}

// DeleteObjectStorageCORS deletes the CORS configuration of an object storage (bucket)
func (s *Session) DeleteObjectStorageCORS(nsId string, osId string) error {
	log.Debug().Msgf("Deleting CORS configuration for object storage: %s in namespace: %s", osId, nsId)

	resp, err := s.Delete(fmt.Sprintf("/ns/%s/resources/objectStorage/%s/cors", nsId, osId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete CORS configuration for object storage: %s", osId)
		return err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to delete CORS configuration for object storage: %s", osId)
		return err
	}

	log.Debug().Msgf("Deleted CORS configuration for object storage (%s) successfully", osId)
	return nil
}

// ============================================================================
// Versioning Management APIs
// ============================================================================

// GetObjectStorageVersioning retrieves the versioning configuration of an object storage (bucket)
func (s *Session) GetObjectStorageVersioning(nsId string, osId string) (ObjectStorageVersioning, error) {
	log.Debug().Msgf("Retrieving versioning configuration for object storage: %s in namespace: %s", osId, nsId)

	var resBody ObjectStorageVersioning
	resp, err := s.
		SetResult(&resBody).
		Get(fmt.Sprintf("/ns/%s/resources/objectStorage/%s/versioning", nsId, osId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to retrieve versioning configuration for object storage: %s", osId)
		return ObjectStorageVersioning{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to retrieve versioning configuration for object storage: %s", osId)
		return ObjectStorageVersioning{}, err
	}

	log.Debug().Msgf("Retrieved versioning configuration for object storage (%s) successfully", osId)
	return resBody, nil
}

// SetObjectStorageVersioning sets the versioning configuration of an object storage (bucket)
func (s *Session) SetObjectStorageVersioning(nsId string, osId string, req tbmodel.ObjectStorageSetVersioningRequest) error {
	log.Debug().Msgf("Setting versioning configuration for object storage: %s in namespace: %s", osId, nsId)

	resp, err := s.
		SetBody(req).
		Put(fmt.Sprintf("/ns/%s/resources/objectStorage/%s/versioning", nsId, osId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to set versioning configuration for object storage: %s", osId)
		return err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to set versioning configuration for object storage: %s", osId)
		return err
	}

	log.Debug().Msgf("Set versioning configuration for object storage (%s) successfully", osId)
	return nil
}

// ============================================================================
// Location APIs
// ============================================================================

// GetObjectStorageLocation retrieves the location of an object storage (bucket)
func (s *Session) GetObjectStorageLocation(nsId string, osId string) (LocationConstraint, error) {
	log.Debug().Msgf("Retrieving location of object storage: %s in namespace: %s", osId, nsId)

	var resBody LocationConstraint
	resp, err := s.
		SetResult(&resBody).
		Get(fmt.Sprintf("/ns/%s/resources/objectStorage/%s/location", nsId, osId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to retrieve location of object storage: %s", osId)
		return LocationConstraint{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to retrieve location of object storage: %s", osId)
		return LocationConstraint{}, err
	}

	log.Debug().Msgf("Retrieved location of object storage (%s) successfully", osId)
	return resBody, nil
}

// ============================================================================
// Object Storage Support Information APIs
// ============================================================================

// GetObjectStorageSupport retrieves CSP support information for object storage features
// If cspType is empty, returns support information for all CSPs
func (s *Session) GetObjectStorageSupport(cspType string) (ObjectStorageSupportResponse, error) {
	log.Debug().Msgf("Retrieving object storage support information for CSP: %s", cspType)

	var resBody ObjectStorageSupportResponse
	req := s.SetResult(&resBody)

	// Add optional CSP type query parameter
	if cspType != "" {
		req = req.SetQueryParam("cspType", cspType)
	}

	resp, err := req.Get("/objectStorage/support")

	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve object storage support information")
		return ObjectStorageSupportResponse{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to retrieve object storage support information")
		return ObjectStorageSupportResponse{}, err
	}

	log.Debug().Msgf("Retrieved object storage support information successfully")
	return resBody, nil
}
