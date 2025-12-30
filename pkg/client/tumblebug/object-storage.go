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

	"github.com/rs/zerolog/log"
)

// * [Note]
// * This Tumblebug Client is used to interact with the CB-Tumblebug API.
// * The Client contains the Tumblebug APIs required for object storage management.

// Owner represents the owner information in S3 bucket list response
type Owner struct {
	ID          string `xml:"ID" json:"id" example:"aws-ap-northeast-2"`
	DisplayName string `xml:"DisplayName" json:"displayName" example:"aws-ap-northeast-2"`
}

// Bucket represents a single bucket in S3 bucket list response
type Bucket struct {
	Name         string `xml:"Name" json:"name" example:"spider-test-bucket"`
	CreationDate string `xml:"CreationDate" json:"creationDate" example:"2025-09-04T04:18:06Z"`
}

// Buckets represents the collection of buckets in S3 bucket list response
type Buckets struct {
	Bucket []Bucket `xml:"Bucket" json:"bucket"`
}

// ListAllMyBucketsResult represents the response structure for S3 ListAllMyBuckets operation
type ListAllMyBucketsResult struct {
	Owner   Owner   `xml:"Owner" json:"owner"`
	Buckets Buckets `xml:"Buckets" json:"buckets"`
}

// ListBucketResult represents the response structure for S3 ListBucket operation
type ListBucketResult struct {
	Name        string `xml:"Name" json:"name" example:"spider-test-bucket"`
	Prefix      string `xml:"Prefix" json:"prefix" example:""`
	Marker      string `xml:"Marker" json:"marker" example:""`
	MaxKeys     int    `xml:"MaxKeys" json:"maxKeys" example:"1000"`
	IsTruncated bool   `xml:"IsTruncated" json:"isTruncated" example:"false"`
}

// LocationConstraint represents the location constraint of a bucket
type LocationConstraint struct {
	Location string `xml:",chardata" json:"location" example:"ap-northeast-2"`
}

// ListObjectStorages retrieves the list of all object storages (buckets)
func (s *Session) ListObjectStorages(connName string) (ListAllMyBucketsResult, error) {
	log.Debug().Msg("Listing object storages")

	var resBody ListAllMyBucketsResult
	resp, err := s.
		SetHeader("credential", connName).
		SetResult(&resBody).
		Get("/resources/objectStorage")

	if err != nil {
		log.Error().Err(err).Msg("Failed to list object storages")
		return ListAllMyBucketsResult{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msg("Failed to list object storages")
		return ListAllMyBucketsResult{}, err
	}

	log.Debug().Msgf("Listed %d object storages successfully", len(resBody.Buckets.Bucket))
	return resBody, nil
}

// CreateObjectStorage creates a new object storage (bucket)
func (s *Session) CreateObjectStorage(objectStorageName, connName string) error {
	log.Debug().Msgf("Creating object storage: %s", objectStorageName)

	resp, err := s.
		SetHeader("credential", connName).
		Put(fmt.Sprintf("/resources/objectStorage/%s", objectStorageName))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to create object storage: %s", objectStorageName)
		return err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to create object storage: %s", objectStorageName)
		return err
	}

	log.Debug().Msgf("Object storage (%s) created successfully", objectStorageName)
	return nil
}

// GetObjectStorage retrieves details of an object storage (bucket)
func (s *Session) GetObjectStorage(objectStorageName, connName string) (ListBucketResult, error) {
	log.Debug().Msgf("Retrieving object storage: %s", objectStorageName)

	var resBody ListBucketResult
	resp, err := s.
		SetHeader("credential", connName).
		SetResult(&resBody).
		Get(fmt.Sprintf("/resources/objectStorage/%s", objectStorageName))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to retrieve object storage: %s", objectStorageName)
		return ListBucketResult{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to retrieve object storage: %s", objectStorageName)
		return ListBucketResult{}, err
	}

	log.Debug().Msgf("Retrieved object storage (%s) successfully", objectStorageName)
	return resBody, nil
}

// ExistObjectStorage checks the existence of an object storage (bucket)
func (s *Session) ExistObjectStorage(objectStorageName, connName string) (bool, error) {
	log.Debug().Msgf("Checking existence of object storage: %s", objectStorageName)

	url := fmt.Sprintf("/resources/objectStorage/%s", objectStorageName)

	resp, err := s.
		SetHeader("credential", connName).
		Head(url)

	if err != nil {
		log.Error().Err(err).Msgf("Failed to check existence of object storage: %s", objectStorageName)
		return false, err
	}

	// HTTP Status OK is 200
	exists := resp.StatusCode() == 200
	log.Debug().Msgf("Object storage (%s) exists: %v", objectStorageName, exists)
	return exists, nil
}

// GetObjectStorageLocation retrieves the location of an object storage (bucket)
func (s *Session) GetObjectStorageLocation(objectStorageName, connName string) (LocationConstraint, error) {
	log.Debug().Msgf("Retrieving location of object storage: %s", objectStorageName)

	var resBody LocationConstraint
	resp, err := s.
		SetHeader("credential", connName).
		SetResult(&resBody).
		Get(fmt.Sprintf("/resources/objectStorage/%s/location", objectStorageName))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to retrieve location of object storage: %s", objectStorageName)
		return LocationConstraint{}, err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to retrieve location of object storage: %s", objectStorageName)
		return LocationConstraint{}, err
	}

	log.Debug().Msgf("Retrieved location of object storage (%s) successfully", objectStorageName)
	return resBody, nil
}

// DeleteObjectStorage deletes an object storage (bucket)
func (s *Session) DeleteObjectStorage(objectStorageName, connName string) error {
	log.Debug().Msgf("Deleting object storage: %s", objectStorageName)

	resp, err := s.
		SetHeader("credential", connName).
		Delete(fmt.Sprintf("/resources/objectStorage/%s", objectStorageName))

	if err != nil {
		log.Error().Err(err).Msgf("Failed to delete object storage: %s", objectStorageName)
		return err
	}

	if resp.IsError() {
		err := fmt.Errorf("API Error: %s (Body: %s)", resp.Status(), string(resp.Body()))
		log.Error().Err(err).Msgf("Failed to delete object storage: %s", objectStorageName)
		return err
	}

	log.Debug().Msgf("Object storage (%s) deleted successfully", objectStorageName)
	return nil
}
