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

// Package migration provides logic to provision target multi-cloud infra and managed middleware
package migration

import (
	"fmt"
	"strings"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// Core RDBMS Migration Functions
// ============================================================================

// CreateRDBMS migrates managed RDBMS instances to the target cloud.
// It applies late-binding via the seed parameter, provisions each DB instance via CB-Tumblebug,
// and optionally creates inner logical databases if specified in the recommendation.
func CreateRDBMS(nsId string, req rdbmsmodel.RecommendedRDBMS, seed string) error {
	log.Info().
		Str("nsId", nsId).
		Str("csp", req.TargetCloud.Csp).
		Str("region", req.TargetCloud.Region).
		Int("targetInstances", len(req.TargetRDBMSInstances)).
		Msg("Starting managed RDBMS migration")

	// Validate connection
	connName, err := GenerateConnectionName(req.TargetCloud.Csp, req.TargetCloud.Region)
	if err != nil {
		return fmt.Errorf("invalid cloud configuration (%s %s): %w", req.TargetCloud.Csp, req.TargetCloud.Region, err)
	}

	// Apply NameSeed (Late Binding) from migration query param
	if seed != "" {
		for i := range req.TargetRDBMSInstances {
			req.TargetRDBMSInstances[i].RDBMSName = common.ComposeName(req.TargetRDBMSInstances[i].RDBMSName, seed)
		}
	}

	tbSess := tbclient.NewSession()

	// Provision each RDBMS instance
	for i, target := range req.TargetRDBMSInstances {
		log.Info().
			Int("index", i+1).
			Int("total", len(req.TargetRDBMSInstances)).
			Str("sourceInstance", target.SourceInstanceName).
			Str("targetRDBMS", target.RDBMSName).
			Str("engine", target.DBEngine).
			Msg("Provisioning managed RDBMS instance")

		adminUser := target.AdminUserName
		if adminUser == "" {
			adminUser = "cbuser"
		}
		adminPass := target.AdminUserPassword
		if adminPass == "" {
			adminPass = "BeetleRdbms1234!" // default temporary password if omitted
		}

		backupDays := target.BackupRetentionDays
		if strings.HasPrefix(strings.ToLower(connName), "ibm") {
			backupDays = 0 // IBM Cloud Databases does not support setting BackupRetentionDays during provisioning
		}

		createReq := rdbmsmodel.RDBMSCreateRequest{
			Name:                     target.RDBMSName,
			ConnectionName:           connName,
			VNetId:                   target.VNetId,
			SubnetIds:                target.SubnetIds,
			SecurityGroupIds:         target.SecurityGroupIds,
			DBEngine:                 target.DBEngine,
			DBEngineVersion:          target.DBEngineVersion,
			DBInstanceSpec:           target.DBInstanceSpec,
			StorageType:              target.StorageType,
			StorageSize:              target.StorageSize,
			Iops:                     target.Iops,
			AdminUserName:            adminUser,
			AdminUserPassword:        adminPass,
			HighAvailability:         target.HighAvailability,
			BackupRetentionDays:      backupDays,
			PublicAccess:             target.PublicAccess,
			NHNDBSGToAllowAllInbound: target.NHNDBSGToAllowAllInbound,
			DeletionProtection:       target.DeletionProtection,
			Description:              fmt.Sprintf("Migrated by CM-Beetle from source instance %s", target.SourceInstanceName),
			AutoFillDefaults:         true,
		}

		createdInfo, err := tbSess.CreateRDBMS(nsId, createReq)
		if err != nil {
			log.Error().Err(err).Str("rdbmsName", target.RDBMSName).Msg("Failed to provision managed RDBMS")
			return fmt.Errorf("failed to provision managed RDBMS '%s': %w", target.RDBMSName, err)
		}

		log.Info().
			Str("rdbmsName", target.RDBMSName).
			Str("status", createdInfo.Status).
			Str("endpoint", createdInfo.Endpoint).
			Msg("Managed RDBMS instance provisioned successfully")

		// Create specified inner databases if any
		if len(target.Databases) > 0 {
			log.Info().
				Str("rdbmsName", target.RDBMSName).
				Int("databasesCount", len(target.Databases)).
				Msg("Creating inner logical databases")

			for _, db := range target.Databases {
				dbReq := rdbmsmodel.RDBMSDatabaseCreateReq{
					DatabaseName:      db.DatabaseName,
					AdminUserPassword: adminPass,
				}
				if dbErr := tbSess.CreateRDBMSDatabase(nsId, target.RDBMSName, dbReq); dbErr != nil {
					log.Warn().
						Err(dbErr).
						Str("rdbmsName", target.RDBMSName).
						Str("databaseName", db.DatabaseName).
						Msg("Failed to create inner database; continuing with next database")
				} else {
					log.Info().
						Str("rdbmsName", target.RDBMSName).
						Str("databaseName", db.DatabaseName).
						Msg("Inner database created successfully")
				}
			}
		}
	}

	log.Info().
		Str("nsId", nsId).
		Str("csp", req.TargetCloud.Csp).
		Str("region", req.TargetCloud.Region).
		Int("totalInstances", len(req.TargetRDBMSInstances)).
		Msg("Managed RDBMS migration completed")

	return nil
}

// ListRDBMS returns all migrated RDBMS instances in the namespace.
func ListRDBMS(nsId string) (rdbmsmodel.RDBMSListResponse, error) {
	log.Info().Str("nsId", nsId).Msg("Listing managed RDBMS instances")

	result, err := tbclient.NewSession().ListRDBMS(nsId, "", "", "")
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Msg("Failed to list managed RDBMS instances")
		return rdbmsmodel.RDBMSListResponse{}, err
	}

	log.Info().Str("nsId", nsId).Int("count", len(result.RDBMS)).Msg("Managed RDBMS instances listed")
	return result, nil
}

// ListRDBMSIDs returns all migrated RDBMS IDs in the namespace.
func ListRDBMSIDs(nsId string) (rdbmsmodel.IdList, error) {
	log.Info().Str("nsId", nsId).Msg("Listing managed RDBMS IDs")

	result, err := tbclient.NewSession().ListRDBMSIDs(nsId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Msg("Failed to list managed RDBMS IDs")
		return rdbmsmodel.IdList{}, err
	}

	return result, nil
}

// GetRDBMS returns details of a specific migrated RDBMS instance.
func GetRDBMS(nsId, rdbmsId string) (rdbmsmodel.RDBMSInfo, error) {
	log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Getting managed RDBMS details")

	result, err := tbclient.NewSession().GetRDBMS(nsId, rdbmsId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Failed to get managed RDBMS details")
		return rdbmsmodel.RDBMSInfo{}, err
	}

	return result, nil
}

// DeleteRDBMS deletes a specific managed RDBMS instance. Treats 404 as already deleted (idempotent).
func DeleteRDBMS(nsId, rdbmsId, option string) error {
	log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Str("option", option).Msg("Deleting managed RDBMS instance")

	err := tbclient.NewSession().DeleteRDBMS(nsId, rdbmsId, option)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Managed RDBMS not found; treating as already deleted")
			return nil
		}
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Failed to delete managed RDBMS instance")
		return fmt.Errorf("failed to delete managed RDBMS '%s': %w", rdbmsId, err)
	}

	log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Managed RDBMS instance deleted successfully")
	return nil
}

// CreateRDBMSDatabase creates a logical database in a specific RDBMS instance.
func CreateRDBMSDatabase(nsId, rdbmsId string, req rdbmsmodel.RDBMSDatabaseCreateReq) error {
	log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Str("databaseName", req.DatabaseName).Msg("Creating logical database")

	err := tbclient.NewSession().CreateRDBMSDatabase(nsId, rdbmsId, req)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Str("databaseName", req.DatabaseName).Msg("Failed to create logical database")
		return err
	}

	log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Str("databaseName", req.DatabaseName).Msg("Logical database created")
	return nil
}

// ListRDBMSDatabases returns logical databases in a specific RDBMS instance.
func ListRDBMSDatabases(nsId, rdbmsId, adminPassword string) (rdbmsmodel.RDBMSDatabaseListResponse, error) {
	log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Listing logical databases")

	result, err := tbclient.NewSession().ListRDBMSDatabases(nsId, rdbmsId, adminPassword)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Msg("Failed to list logical databases")
		return rdbmsmodel.RDBMSDatabaseListResponse{}, err
	}

	log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Int("count", len(result.Databases)).Msg("Logical databases listed")
	return result, nil
}

// DeleteRDBMSDatabase deletes a logical database in a specific RDBMS instance.
func DeleteRDBMSDatabase(nsId, rdbmsId, dbName, adminPassword string) error {
	log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Str("dbName", dbName).Msg("Deleting logical database")

	err := tbclient.NewSession().DeleteRDBMSDatabase(nsId, rdbmsId, dbName, adminPassword)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("rdbmsId", rdbmsId).Str("dbName", dbName).Msg("Failed to delete logical database")
		return err
	}

	log.Info().Str("nsId", nsId).Str("rdbmsId", rdbmsId).Str("dbName", dbName).Msg("Logical database deleted")
	return nil
}
