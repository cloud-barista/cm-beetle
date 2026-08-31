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
	"strings"
	"time"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	"github.com/cloud-barista/cm-beetle/pkg/ratelimit"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// RDBMS Support & Capability APIs
// ============================================================================

// GetRDBMSSupport retrieves whether managed RDBMS is supported for the specified CSP or all CSPs.
func (s *Session) GetRDBMSSupport(optionalCsp ...string) (rdbmsmodel.RDBMSSupportResponse, error) {
	log.Debug().Msg("Retrieving RDBMS support information from CB-Tumblebug")

	var resBody rdbmsmodel.RDBMSSupportResponse
	req := s.SetResult(&resBody)

	if len(optionalCsp) > 0 && optionalCsp[0] != "" {
		req = req.SetQueryParam("providerName", optionalCsp[0])
	}

	resp, err := req.Get("/rdbms/support")
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve RDBMS support info")
		return rdbmsmodel.RDBMSSupportResponse{}, err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return rdbmsmodel.RDBMSSupportResponse{}, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to retrieve RDBMS support info")
		return rdbmsmodel.RDBMSSupportResponse{}, err
	}

	log.Debug().Msg("Retrieved RDBMS support information successfully")
	return resBody, nil
}

// GetRDBMSCapability retrieves real-time capability information (specs, versions, constraints) for a connection and optional dbEngine.
func (s *Session) GetRDBMSCapability(connectionName string, optionalEngine ...string) (rdbmsmodel.RDBMSCapabilityResponse, error) {
	log.Debug().Msgf("Retrieving RDBMS capability for connection: %s", connectionName)

	var resBody rdbmsmodel.RDBMSCapabilityResponse
	req := s.SetResult(&resBody)

	engine := "mysql"
	if len(optionalEngine) > 0 && optionalEngine[0] != "" {
		engine = optionalEngine[0]
	}

	req = req.SetQueryParam("dbEngine", engine)

	if connectionName != "" {
		req = req.SetQueryParam("connectionName", connectionName)
		parts := strings.Split(connectionName, "-")
		if len(parts) >= 2 {
			providerName := strings.ToLower(parts[0])
			regionName := strings.Join(parts[1:], "-")

			// Normalize providerName and regionName to match Tumblebug connection configs
			switch providerName {
			case "ncp":
				regionName = strings.ToUpper(regionName) // e.g., "kr" -> "KR"
			case "nhn":
				providerName = "nhn"                     // Registered providerName is "nhn"
				regionName = strings.ToUpper(regionName) // e.g., "kr1" -> "KR1"
			case "kt":
				regionName = strings.ToUpper(regionName) // e.g., "kr1" -> "KR1"
			}

			req = req.SetQueryParam("providerName", providerName).
				SetQueryParam("regionName", regionName)
		}
	}

	resp, err := req.Get("/rdbms/capability")
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve RDBMS capability")
		return rdbmsmodel.RDBMSCapabilityResponse{}, err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return rdbmsmodel.RDBMSCapabilityResponse{}, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to retrieve RDBMS capability")
		return rdbmsmodel.RDBMSCapabilityResponse{}, err
	}

	log.Debug().Msg("Retrieved RDBMS capability successfully")
	return resBody, nil
}

// ValidateRDBMS performs dry-run validation and autofills defaults for an RDBMS create request.
func (s *Session) ValidateRDBMS(nsId string, req rdbmsmodel.RDBMSCreateRequest) (rdbmsmodel.RDBMSCreateRequest, error) {
	log.Debug().Msgf("Validating RDBMS create request in namespace: %s", nsId)

	var resBody rdbmsmodel.RDBMSCreateRequest
	resp, err := s.
		SetBody(req).
		SetResult(&resBody).
		Post(fmt.Sprintf("/ns/%s/resources/rdbms/validate", nsId))

	if err != nil {
		log.Error().Err(err).Msg("Failed to validate RDBMS create request")
		return rdbmsmodel.RDBMSCreateRequest{}, err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return rdbmsmodel.RDBMSCreateRequest{}, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to validate RDBMS create request")
		return rdbmsmodel.RDBMSCreateRequest{}, err
	}

	log.Debug().Msg("Validated RDBMS create request successfully")
	return resBody, nil
}

// ============================================================================
// RDBMS Instance Management APIs
// ============================================================================

// CreateRDBMS provisions a new managed RDBMS instance in the specified namespace.
func (s *Session) CreateRDBMS(nsId string, req rdbmsmodel.RDBMSCreateRequest) (rdbmsmodel.RDBMSInfo, error) {
	log.Debug().Msgf("Creating RDBMS: %s in namespace: %s", req.Name, nsId)

	var resBody rdbmsmodel.RDBMSInfo
	resp, err := s.
		SetBody(req).
		SetResult(&resBody).
		Post(fmt.Sprintf("/ns/%s/resources/rdbms", nsId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to create RDBMS: %s", req.Name)
		return rdbmsmodel.RDBMSInfo{}, err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return rdbmsmodel.RDBMSInfo{}, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to create RDBMS: %s", req.Name)
		return rdbmsmodel.RDBMSInfo{}, err
	}

	log.Debug().Msgf("RDBMS (%s) created successfully with status: %s", req.Name, resBody.Status)
	return resBody, nil
}

// GetRDBMS retrieves details of a specific RDBMS instance.
func (s *Session) GetRDBMS(nsId, rdbmsId string) (rdbmsmodel.RDBMSInfo, error) {
	log.Debug().Msgf("Retrieving RDBMS: %s in namespace: %s", rdbmsId, nsId)

	var resBody rdbmsmodel.RDBMSInfo
	resp, err := s.
		SetResult(&resBody).
		Get(fmt.Sprintf("/ns/%s/resources/rdbms/%s", nsId, rdbmsId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to retrieve RDBMS: %s", rdbmsId)
		return rdbmsmodel.RDBMSInfo{}, err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return rdbmsmodel.RDBMSInfo{}, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to retrieve RDBMS: %s", rdbmsId)
		return rdbmsmodel.RDBMSInfo{}, err
	}

	log.Debug().Msgf("Retrieved RDBMS (%s) successfully", rdbmsId)
	return resBody, nil
}

// ListRDBMS lists all RDBMS instances in a namespace.
func (s *Session) ListRDBMS(nsId, option, filterKey, filterVal string) (rdbmsmodel.RDBMSListResponse, error) {
	log.Debug().Msgf("Listing RDBMS in namespace: %s", nsId)

	var resBody rdbmsmodel.RDBMSListResponse
	req := s.SetResult(&resBody)

	if option != "" {
		req = req.SetQueryParam("option", option)
	}
	if filterKey != "" {
		req = req.SetQueryParam("filterKey", filterKey)
	}
	if filterVal != "" {
		req = req.SetQueryParam("filterVal", filterVal)
	}

	resp, err := req.Get(fmt.Sprintf("/ns/%s/resources/rdbms", nsId))
	if err != nil {
		log.Error().Err(err).Msg("Failed to list RDBMS instances")
		return rdbmsmodel.RDBMSListResponse{}, err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return rdbmsmodel.RDBMSListResponse{}, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to list RDBMS instances")
		return rdbmsmodel.RDBMSListResponse{}, err
	}

	log.Debug().Msgf("Listed %d RDBMS instances successfully", len(resBody.RDBMS))
	return resBody, nil
}

// ListRDBMSIDs retrieves the list of all RDBMS IDs in a namespace.
func (s *Session) ListRDBMSIDs(nsId string) (rdbmsmodel.IdList, error) {
	log.Debug().Msgf("Listing RDBMS IDs in namespace: %s", nsId)

	var resBody rdbmsmodel.IdList
	resp, err := s.
		SetQueryParam("option", "id").
		SetResult(&resBody).
		Get(fmt.Sprintf("/ns/%s/resources/rdbms", nsId))

	if err != nil {
		log.Error().Err(err).Msg("Failed to list RDBMS IDs")
		return rdbmsmodel.IdList{}, err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return rdbmsmodel.IdList{}, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to list RDBMS IDs")
		return rdbmsmodel.IdList{}, err
	}

	log.Debug().Msgf("Listed %d RDBMS IDs successfully", len(resBody.IdList))
	return resBody, nil
}

// DeleteRDBMS deletes a specific RDBMS instance.
func (s *Session) DeleteRDBMS(nsId, rdbmsId, option string) error {
	log.Debug().Msgf("Deleting RDBMS: %s in namespace: %s", rdbmsId, nsId)

	req := s
	if option != "" {
		req.SetQueryParam("option", option)
	}

	resp, err := req.Delete(fmt.Sprintf("/ns/%s/resources/rdbms/%s", nsId, rdbmsId))
	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete RDBMS: %s", rdbmsId)
		return err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to delete RDBMS: %s", rdbmsId)
		return err
	}

	log.Debug().Msgf("RDBMS (%s) deleted successfully", rdbmsId)
	return nil
}

// ============================================================================
// Inner Database Management APIs
// ============================================================================

// CreateRDBMSDatabase creates a logical database inside an existing RDBMS instance.
func (s *Session) CreateRDBMSDatabase(nsId, rdbmsId string, req rdbmsmodel.RDBMSDatabaseCreateReq) error {
	log.Debug().Msgf("Creating database '%s' in RDBMS: %s, namespace: %s", req.DatabaseName, rdbmsId, nsId)

	resp, err := s.
		SetBody(req).
		Post(fmt.Sprintf("/ns/%s/resources/rdbms/%s/database", nsId, rdbmsId))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to create database '%s' in RDBMS '%s'", req.DatabaseName, rdbmsId)
		return err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to create database '%s' in RDBMS '%s'", req.DatabaseName, rdbmsId)
		return err
	}

	log.Debug().Msgf("Database '%s' created successfully in RDBMS '%s'", req.DatabaseName, rdbmsId)
	return nil
}

// ListRDBMSDatabases retrieves the list of databases inside an RDBMS instance.
func (s *Session) ListRDBMSDatabases(nsId, rdbmsId, adminPassword string) (rdbmsmodel.RDBMSDatabaseListResponse, error) {
	log.Debug().Msgf("Listing databases in RDBMS: %s, namespace: %s", rdbmsId, nsId)

	var resBody rdbmsmodel.RDBMSDatabaseListResponse
	req := s.SetResult(&resBody)

	if adminPassword != "" {
		req = req.SetHeader("X-Admin-User-Password", adminPassword)
	}

	resp, err := req.Get(fmt.Sprintf("/ns/%s/resources/rdbms/%s/database", nsId, rdbmsId))
	if err != nil {
		log.Error().Err(err).Msgf("Failed to list databases in RDBMS: %s", rdbmsId)
		return rdbmsmodel.RDBMSDatabaseListResponse{}, err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return rdbmsmodel.RDBMSDatabaseListResponse{}, &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to list databases in RDBMS: %s", rdbmsId)
		return rdbmsmodel.RDBMSDatabaseListResponse{}, err
	}

	log.Debug().Msgf("Listed %d databases in RDBMS (%s) successfully", len(resBody.Databases), rdbmsId)
	return resBody, nil
}

// DeleteRDBMSDatabase deletes a logical database inside an RDBMS instance.
func (s *Session) DeleteRDBMSDatabase(nsId, rdbmsId, dbName, adminPassword string) error {
	log.Debug().Msgf("Deleting database '%s' in RDBMS: %s, namespace: %s", dbName, rdbmsId, nsId)

	req := s
	if adminPassword != "" {
		req = req.SetHeader("X-Admin-User-Password", adminPassword)
	}

	resp, err := req.Delete(fmt.Sprintf("/ns/%s/resources/rdbms/%s/database/%s", nsId, rdbmsId, dbName))
	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete database '%s' in RDBMS '%s'", dbName, rdbmsId)
		return err
	}

	if resp.IsError() {
		if resp.StatusCode() == http.StatusTooManyRequests {
			return &ratelimit.ErrLimited{
				RetryAfter: 2 * time.Second,
			}
		}
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to delete database '%s' in RDBMS '%s'", dbName, rdbmsId)
		return err
	}

	log.Debug().Msgf("Database '%s' deleted successfully in RDBMS '%s'", dbName, rdbmsId)
	return nil
}
