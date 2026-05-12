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

// Package migration is to provision target multi-cloud infra for migration
package migration

import (
	"fmt"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// Helpers
// ============================================================================

// GenerateConnectionName builds and validates a connection name from csp and region.
func GenerateConnectionName(csp, region string) (string, error) {
	connectionName := strings.ToLower(fmt.Sprintf("%s-%s", csp, region))

	_, err := tbclient.NewSession().GetConnConfig(connectionName)
	if err != nil {
		log.Error().Err(err).Str("connectionName", connectionName).Msg("Failed to get connection config")
		return "", err
	}

	return connectionName, nil
}

// toMigratedObjectStorageInfo converts a TB ObjectStorageInfo to the migration model representation.
func toMigratedObjectStorageInfo(src tbmodel.ObjectStorageInfo) MigratedObjectStorageInfo {
	return MigratedObjectStorageInfo{
		Id:             src.Id,
		Name:           src.Name,
		Status:         src.Status,
		Description:    src.Description,
		ConnectionName: src.ConnectionName,
		CreationDate:   src.CreationDate,
	}
}

// ============================================================================
// Core functions
// ============================================================================

// CreateObjectStorage migrates object storages to the target cloud.
// It applies late-binding via the seed parameter, creates each bucket, then configures
// versioning and CORS according to CSP support information.
func CreateObjectStorage(nsId string, req storagemodel.RecommendedObjectStorage, seed string) error {
	log.Info().
		Str("nsId", nsId).
		Str("csp", req.TargetCloud.Csp).
		Str("region", req.TargetCloud.Region).
		Int("targetBuckets", len(req.TargetObjectStorages)).
		Msg("Starting object storage migration")

	// Validate connection
	connName, err := GenerateConnectionName(req.TargetCloud.Csp, req.TargetCloud.Region)
	if err != nil {
		return fmt.Errorf("invalid cloud configuration (%s %s): %w", req.TargetCloud.Csp, req.TargetCloud.Region, err)
	}

	// Apply NameSeed (Late Binding) from migration query param
	if seed != "" {
		for i := range req.TargetObjectStorages {
			req.TargetObjectStorages[i].BucketName = common.ComposeName(req.TargetObjectStorages[i].BucketName, seed)
		}
	}

	// Create each bucket
	for i, target := range req.TargetObjectStorages {
		log.Debug().
			Int("index", i+1).
			Int("total", len(req.TargetObjectStorages)).
			Str("sourceBucket", target.SourceBucketName).
			Str("targetBucket", target.BucketName).
			Msg("Creating object storage")

		createReq := tbmodel.ObjectStorageCreateRequest{
			BucketName:     target.BucketName,
			ConnectionName: connName,
			Description:    "Created by CM-Beetle",
		}
		if _, err := tbclient.NewSession().CreateObjectStorage(nsId, createReq); err != nil {
			log.Error().Err(err).Str("bucketName", target.BucketName).Msg("Failed to create object storage")
			return fmt.Errorf("failed to create object storage '%s': %w", target.BucketName, err)
		}

		log.Info().
			Str("sourceBucket", target.SourceBucketName).
			Str("targetBucket", target.BucketName).
			Msg("Object storage created")
	}

	// Configure versioning and CORS
	// Re-check support because the user may have modified the recommendation before calling migration
	supportResp, supportErr := tbclient.NewSession().GetObjectStorageSupport(req.TargetCloud.Csp)
	if supportErr != nil {
		log.Warn().Err(supportErr).Str("csp", req.TargetCloud.Csp).
			Msg("Failed to fetch CSP support info; skipping versioning/CORS configuration")
	} else {
		support, hasSupportInfo := supportResp.Supports[req.TargetCloud.Csp]
		corsSupported := !hasSupportInfo || support.Cors
		versioningSupported := !hasSupportInfo || support.Versioning

		tbSess := tbclient.NewSession()

		for _, target := range req.TargetObjectStorages {
			bucketId := target.BucketName

			// Configure versioning
			if target.VersioningEnabled {
				if !versioningSupported {
					log.Warn().
						Str("bucket", bucketId).
						Str("csp", req.TargetCloud.Csp).
						Msg("Versioning requested but not supported by CSP; skipping")
				} else {
					verReq := tbmodel.ObjectStorageSetVersioningRequest{Status: "Enabled"}
					if err := tbSess.SetObjectStorageVersioning(nsId, bucketId, verReq); err != nil {
						log.Error().Err(err).Str("bucket", bucketId).Msg("Failed to enable versioning; continuing")
					} else {
						log.Info().Str("bucket", bucketId).Msg("Versioning enabled")
					}
				}
			}

			// Configure CORS
			if target.CORSEnabled && len(target.CORSRule) > 0 {
				if !corsSupported {
					log.Warn().
						Str("bucket", bucketId).
						Str("csp", req.TargetCloud.Csp).
						Msg("CORS requested but not supported by CSP; skipping")
				} else {
					corsRules := make([]tbmodel.CorsRule, 0, len(target.CORSRule))
					for _, r := range target.CORSRule {
						corsRules = append(corsRules, tbmodel.CorsRule{
							AllowedOrigin: r.AllowedOrigin,
							AllowedMethod: r.AllowedMethod,
							AllowedHeader: r.AllowedHeader,
							ExposeHeader:  r.ExposeHeader,
							MaxAgeSeconds: r.MaxAgeSeconds,
						})
					}
					corsReq := tbmodel.ObjectStorageSetCorsRequest{CorsRule: corsRules}
					if err := tbSess.SetObjectStorageCORS(nsId, bucketId, corsReq); err != nil {
						log.Error().Err(err).Str("bucket", bucketId).Msg("Failed to set CORS; continuing")
					} else {
						log.Info().Str("bucket", bucketId).Msg("CORS configured")
					}
				}
			}
		}
	}

	log.Info().
		Str("nsId", nsId).
		Str("csp", req.TargetCloud.Csp).
		Str("region", req.TargetCloud.Region).
		Int("totalBuckets", len(req.TargetObjectStorages)).
		Msg("Object storage migration completed")

	return nil
}

// ListObjectStorages returns all migrated object storages in the namespace.
func ListObjectStorages(nsId string) (MigratedObjectStorageListResponse, error) {
	log.Info().Str("nsId", nsId).Msg("Listing object storages")

	result, err := tbclient.NewSession().ListObjectStorages(nsId, "", "", "")
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Msg("Failed to list object storages")
		return MigratedObjectStorageListResponse{}, err
	}

	infos := make([]MigratedObjectStorageInfo, 0, len(result.ObjectStorage))
	for _, item := range result.ObjectStorage {
		infos = append(infos, toMigratedObjectStorageInfo(item))
	}

	resp := MigratedObjectStorageListResponse{ObjectStorages: infos}

	log.Info().Str("nsId", nsId).Int("count", len(infos)).Msg("Object storages listed")
	return resp, nil
}

// GetObjectStorage returns details of a specific migrated object storage.
func GetObjectStorage(nsId, osId string) (MigratedObjectStorageInfo, error) {
	log.Info().Str("nsId", nsId).Str("osId", osId).Msg("Getting object storage")

	result, err := tbclient.NewSession().GetObjectStorage(nsId, osId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to get object storage")
		return MigratedObjectStorageInfo{}, err
	}

	log.Info().Str("nsId", nsId).Str("osId", osId).Msg("Object storage retrieved")
	return toMigratedObjectStorageInfo(result), nil
}

// ExistObjectStorage checks whether a specific object storage exists.
func ExistObjectStorage(nsId, osId string) (bool, error) {
	log.Info().Str("nsId", nsId).Str("osId", osId).Msg("Checking object storage existence")

	exists, err := tbclient.NewSession().ExistObjectStorage(nsId, osId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to check object storage existence")
		return false, err
	}

	log.Info().Str("nsId", nsId).Str("osId", osId).Bool("exists", exists).Msg("Object storage existence checked")
	return exists, nil
}

// DeleteObjectStorage deletes a specific object storage. Treats 404 as already deleted (idempotent).
// option controls deletion behavior: "" (standard), "empty" (empty first), "force" (force with contents), "reconcile" (metadata only).
func DeleteObjectStorage(nsId, osId, option string) error {
	log.Info().Str("nsId", nsId).Str("osId", osId).Str("option", option).Msg("Deleting object storage")

	err := tbclient.NewSession().DeleteObjectStorage(nsId, osId, option)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			log.Info().Str("nsId", nsId).Str("osId", osId).Msg("Object storage not found; treating as already deleted")
			return nil
		}
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to delete object storage")
		return fmt.Errorf("failed to delete object storage '%s': %w", osId, err)
	}

	log.Info().Str("nsId", nsId).Str("osId", osId).Str("option", option).Msg("Object storage deleted")
	return nil
}

// ListObjectStorageObjects returns the list of objects stored in a specific object storage bucket.
func ListObjectStorageObjects(nsId, osId string) (StorageObjectListResponse, error) {
	log.Info().Str("nsId", nsId).Str("osId", osId).Msg("Listing objects in object storage")

	result, err := tbclient.NewSession().ListObjectStorageObjects(nsId, osId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to list objects in object storage")
		return StorageObjectListResponse{}, err
	}

	objects := make([]StorageObjectInfo, 0, len(result.Objects))
	for _, obj := range result.Objects {
		objects = append(objects, StorageObjectInfo{
			Key:          obj.Key,
			Size:         obj.Size,
			LastModified: obj.LastModified,
			ETag:         obj.ETag,
			StorageClass: obj.StorageClass,
		})
	}

	resp := StorageObjectListResponse{
		OsId:    osId,
		Count:   len(objects),
		Objects: objects,
	}

	log.Info().Str("nsId", nsId).Str("osId", osId).Int("count", resp.Count).Msg("Objects listed")
	return resp, nil
}

// GetStorageObject retrieves metadata of a specific object from an object storage bucket.
func GetStorageObject(nsId, osId, objectKey string) (StorageObjectMetadata, error) {
	log.Info().Str("nsId", nsId).Str("osId", osId).Str("objectKey", objectKey).Msg("Getting object metadata")

	obj, err := tbclient.NewSession().GetStorageObject(nsId, osId, objectKey)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Str("objectKey", objectKey).Msg("Failed to get object metadata")
		return StorageObjectMetadata{}, err
	}

	metadata := StorageObjectMetadata{
		Key:          obj.Key,
		Size:         obj.Size,
		LastModified: obj.LastModified,
		ETag:         obj.ETag,
		StorageClass: obj.StorageClass,
	}

	log.Info().Str("nsId", nsId).Str("osId", osId).Str("objectKey", objectKey).Msg("Object metadata retrieved")
	return metadata, nil
}

// DeleteStorageObject deletes a specific object from an object storage bucket.
func DeleteStorageObject(nsId, osId, objectKey string) error {
	log.Info().Str("nsId", nsId).Str("osId", osId).Str("objectKey", objectKey).Msg("Deleting object from object storage")

	if err := tbclient.NewSession().DeleteStorageObject(nsId, osId, objectKey); err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Str("objectKey", objectKey).Msg("Failed to delete object")
		return err
	}

	log.Info().Str("nsId", nsId).Str("osId", osId).Str("objectKey", objectKey).Msg("Object deleted")
	return nil
}
