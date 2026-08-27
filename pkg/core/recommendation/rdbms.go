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

// Package recommendation provides logic to recommend optimal target cloud resources
package recommendation

import (
	"fmt"
	"strings"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// ValidateRDBMS performs dry-run validation and default autofill for an RDBMS create request via CB-Tumblebug.
func ValidateRDBMS(nsId string, req rdbmsmodel.RDBMSCreateRequest) (rdbmsmodel.RDBMSCreateRequest, error) {
	if nsId == "" {
		nsId = "default"
	}
	log.Info().Str("nsId", nsId).Str("rdbmsName", req.Name).Msg("Validating RDBMS create request via CB-Tumblebug")
	return tbclient.NewSession().ValidateRDBMS(nsId, req)
}

// ValidateRecommendedRDBMS validates all recommended RDBMS instances against target cloud constraints.
func ValidateRecommendedRDBMS(nsId string, req rdbmsmodel.RecommendedRDBMS) ([]rdbmsmodel.RDBMSCreateRequest, error) {
	if nsId == "" {
		nsId = "default"
	}

	connName := strings.ToLower(fmt.Sprintf("%s-%s", req.TargetCloud.Csp, req.TargetCloud.Region))
	if req.TargetCloud.Csp == "" || req.TargetCloud.Region == "" {
		return nil, fmt.Errorf("target cloud CSP and region are required")
	}

	validatedList := make([]rdbmsmodel.RDBMSCreateRequest, 0, len(req.TargetRDBMSInstances))
	tbSess := tbclient.NewSession()

	for _, inst := range req.TargetRDBMSInstances {
		adminUser := inst.AdminUserName
		if adminUser == "" {
			adminUser = "cbuser"
		}
		adminPass := inst.AdminUserPassword
		if adminPass == "" {
			adminPass = "BeetleRdbms1234!"
		}

		createReq := rdbmsmodel.RDBMSCreateRequest{
			Name:                inst.RDBMSName,
			ConnectionName:      connName,
			VNetId:              inst.VNetId,
			SubnetIds:           inst.SubnetIds,
			SecurityGroupIds:    inst.SecurityGroupIds,
			DBEngine:            inst.DBEngine,
			DBEngineVersion:     inst.DBEngineVersion,
			DBInstanceSpec:      inst.DBInstanceSpec,
			StorageType:         inst.StorageType,
			StorageSize:         inst.StorageSize,
			Iops:                inst.Iops,
			AdminUserName:       adminUser,
			AdminUserPassword:   adminPass,
			HighAvailability:    inst.HighAvailability,
			BackupRetentionDays: inst.BackupRetentionDays,
			PublicAccess:        inst.PublicAccess,
			DeletionProtection:  inst.DeletionProtection,
			AutoFillDefaults:    true,
		}

		validated, vErr := tbSess.ValidateRDBMS(nsId, createReq)
		if vErr != nil {
			log.Warn().Err(vErr).Str("rdbmsName", inst.RDBMSName).Msg("Validation via Tumblebug failed; retaining local recommendation")
			validatedList = append(validatedList, createReq)
		} else {
			validatedList = append(validatedList, validated)
		}
	}

	return validatedList, nil
}

// RecommendRDBMS recommends optimal managed RDBMS instances for target cloud migration.
func RecommendRDBMS(desiredCsp, desiredRegion string, sources []rdbmsmodel.SourceRDBMSProperty) (rdbmsmodel.RecommendedRDBMS, error) {
	desiredCsp = strings.ToLower(strings.TrimSpace(desiredCsp))
	desiredRegion = strings.ToLower(strings.TrimSpace(desiredRegion))

	log.Info().
		Str("csp", desiredCsp).
		Str("region", desiredRegion).
		Int("sourceCount", len(sources)).
		Msg("Starting RDBMS recommendation")

	warnings := make([]string, 0)

	// Fetch CSP support info from Tumblebug
	supportResp, err := tbclient.NewSession().GetRDBMSSupport(desiredCsp)
	if err != nil {
		log.Warn().Err(err).Str("csp", desiredCsp).Msg("Failed to fetch CSP RDBMS support info from Tumblebug")
	}

	support, hasSupport := supportResp.Supports[desiredCsp]
	if hasSupport && !support.Supported {
		return rdbmsmodel.RecommendedRDBMS{
			Status:      "failed",
			Description: fmt.Sprintf("Managed RDBMS is not supported on CSP '%s'", desiredCsp),
			TargetCloud: rdbmsmodel.CloudProperty{Csp: desiredCsp, Region: desiredRegion},
		}, fmt.Errorf("managed RDBMS is not supported on CSP '%s'", desiredCsp)
	}

	// Fetch real-time capability if connectionName is constructible
	connName := fmt.Sprintf("%s-%s", desiredCsp, desiredRegion)
	capabilityResp, capErr := tbclient.NewSession().GetRDBMSCapability(connName)
	if capErr != nil {
		log.Warn().Err(capErr).Str("connectionName", connName).Msg("Failed to fetch live RDBMS capability from Tumblebug")
	}
	cap := capabilityResp.Supports

	// Recommend each target RDBMS instance
	targetInstances := make([]rdbmsmodel.TargetRDBMSInstance, 0, len(sources))

	for i, src := range sources {
		instNum := i + 1
		targetName := fmt.Sprintf("mig-rdbms-%02d", instNum)

		// 1. Engine & Version Recommendation
		targetEngine := strings.ToLower(strings.TrimSpace(src.Engine))
		if targetEngine == "" {
			targetEngine = "mysql"
		}

		// Check MariaDB support
		if targetEngine == "mariadb" && !isMariaDBSupported(desiredCsp, support, hasSupport) {
			targetEngine = "mysql"
			warning := fmt.Sprintf("MariaDB is not supported on CSP '%s'. Recommended MySQL as fallback for instance '%s'.", desiredCsp, src.InstanceName)
			warnings = append(warnings, warning)
			log.Warn().Msg(warning)
		}

		// Select Engine Version
		targetVersion := src.EngineVersion
		if len(cap.SupportedVersions) > 0 {
			matched := false
			for _, v := range cap.SupportedVersions {
				if strings.HasPrefix(v, src.EngineVersion) || strings.HasPrefix(src.EngineVersion, v) {
					targetVersion = v
					matched = true
					break
				}
			}
			if !matched {
				// Use the default or first supported version
				targetVersion = cap.SupportedVersions[0]
			}
		}
		if targetVersion == "" {
			targetVersion = "8.0"
		}

		// 2. DB Spec Recommendation
		targetSpec := selectDBSpec(desiredCsp, targetEngine, src.Vcpu, src.MemoryMb, cap)

		// 3. Storage Type & Size Recommendation
		targetStorageType := cap.DefaultStorageType
		if targetStorageType == "" {
			targetStorageType = selectDefaultStorageType(desiredCsp)
		}

		targetStorageSize := src.StorageSizeGb
		minStorage := cap.StorageSizeRange.Min
		if minStorage == 0 {
			minStorage = selectDefaultMinStorage(desiredCsp, targetStorageType)
		}
		if targetStorageSize < minStorage {
			warning := fmt.Sprintf("Adjusted storage size for instance '%s' from %dGB to minimum %dGB required by %s (%s).",
				src.InstanceName, targetStorageSize, minStorage, desiredCsp, targetStorageType)
			warnings = append(warnings, warning)
			targetStorageSize = minStorage
		}

		targetIops := ""
		if src.Iops > 0 {
			targetIops = fmt.Sprintf("%d", src.Iops)
		} else if desiredCsp == "aws" && targetStorageType == "gp3" {
			targetIops = "3000"
		}

		// 4. Admin Credentials Default
		targetAdminUser := "cbuser"
		if cap.AdminUserNameRequirement != nil && cap.AdminUserNameRequirement.FixedValue != "" {
			targetAdminUser = cap.AdminUserNameRequirement.FixedValue
		} else if desiredCsp == "ibm" {
			targetAdminUser = "admin"
		} else if desiredCsp == "tencent" {
			targetAdminUser = "root"
		}

		// 5. Network & CSP-specific warnings
		if desiredCsp == "ncp" && src.PublicAccess {
			warning := fmt.Sprintf("NCP Cloud DB does not provide external public IP by default; instance '%s' will be created within private VPC.", src.InstanceName)
			warnings = append(warnings, warning)
		}
		if desiredCsp == "nhn" && src.PublicAccess {
			warning := fmt.Sprintf("NHN Cloud blocks port 3306 by default; ensure DB Security Group inbound permit rule is enabled for instance '%s'.", src.InstanceName)
			warnings = append(warnings, warning)
		}
		if src.HighAvailability && desiredCsp == "aws" {
			warning := fmt.Sprintf("High availability (Multi-AZ) on AWS requires Subnets in at least two distinct Availability Zones for instance '%s'.", src.InstanceName)
			warnings = append(warnings, warning)
		}

		// 6. Map Inner Databases
		targetDatabases := make([]rdbmsmodel.TargetDatabase, 0, len(src.Databases))
		for _, db := range src.Databases {
			targetDatabases = append(targetDatabases, rdbmsmodel.TargetDatabase{
				DatabaseName: db.DatabaseName,
				CharacterSet: db.CharacterSet,
			})
		}

		targetInst := rdbmsmodel.TargetRDBMSInstance{
			SourceInstanceName:  src.InstanceName,
			SourceMachineId:     src.MachineId,
			RDBMSName:           targetName,
			DBEngine:            targetEngine,
			DBEngineVersion:     targetVersion,
			DBInstanceSpec:      targetSpec,
			StorageType:         targetStorageType,
			StorageSize:         targetStorageSize,
			Iops:                targetIops,
			AdminUserName:       targetAdminUser,
			HighAvailability:    src.HighAvailability,
			BackupRetentionDays: src.BackupRetentionDays,
			PublicAccess:        src.PublicAccess,
			Databases:           targetDatabases,
		}

		targetInstances = append(targetInstances, targetInst)
	}

	result := rdbmsmodel.RecommendedRDBMS{
		Status:               "recommended",
		Description:          fmt.Sprintf("Successfully recommended %d managed RDBMS configuration(s) for %s (%s)", len(targetInstances), desiredCsp, desiredRegion),
		Warnings:             warnings,
		TargetCloud:          rdbmsmodel.CloudProperty{Csp: desiredCsp, Region: desiredRegion},
		TargetRDBMSInstances: targetInstances,
	}

	log.Info().
		Int("targetInstances", len(targetInstances)).
		Int("warnings", len(warnings)).
		Msg("RDBMS recommendation completed")

	return result, nil
}

// selectDBSpec selects the optimal DB instance spec for the target CSP and required compute resources.
func selectDBSpec(csp, engine string, vcpu, memoryMb int, cap rdbmsmodel.RDBMSMetaInfo) string {
	// If capability provides a reference or list of options, try matching
	if len(cap.DBInstanceSpecOptions) > 0 {
		// If capability has options, pick a suitable one (default to first/reference)
		return cap.DBInstanceSpecOptions[0]
	}

	// Standard default specs per CSP
	switch csp {
	case "aws":
		if vcpu >= 4 || memoryMb >= 16384 {
			return "db.m6i.xlarge"
		} else if vcpu >= 2 || memoryMb >= 8192 {
			return "db.t3.large"
		}
		return "db.t3.medium"
	case "azure":
		if vcpu >= 4 {
			return "Standard_D4ds_v4"
		} else if vcpu >= 2 {
			return "Standard_D2ds_v4"
		}
		return "Standard_B1ms"
	case "gcp":
		if vcpu >= 4 {
			return "db-custom-4-16384"
		}
		return "db-custom-2-8192"
	case "alibaba":
		if vcpu >= 4 {
			return "mysql.n4.xlarge.1"
		}
		return "mysql.n4.large.1"
	case "tencent":
		return "8000"
	case "ibm":
		return "multitenant"
	case "openstack":
		return "m1.small"
	case "ncp":
		return "SVR.VDBAS.AMD.STAND.C002.M008.NET.SSD.B050.G003"
	case "nhn":
		return "m2.c2m4"
	default:
		return "db.t3.medium"
	}
}

// selectDefaultStorageType returns standard storage type per CSP.
func selectDefaultStorageType(csp string) string {
	switch csp {
	case "aws":
		return "gp3"
	case "gcp":
		return "PD_SSD"
	case "alibaba":
		return "cloud_essd"
	case "tencent":
		return "local_ssd"
	case "nhn":
		return "General SSD"
	default:
		return "gp3"
	}
}

// selectDefaultMinStorage returns minimum storage size in GB per CSP and storage type.
func selectDefaultMinStorage(csp, storageType string) int {
	switch csp {
	case "aws":
		if storageType == "gp3" {
			return 100
		}
		return 20
	case "tencent":
		return 50
	case "alibaba", "azure", "gcp", "nhn", "openstack":
		return 20
	case "ibm":
		return 30
	default:
		return 20
	}
}

// isMariaDBSupported checks if the target CSP supports MariaDB using live support info or static reference.
func isMariaDBSupported(csp string, support rdbmsmodel.RDBMSCSPSupportInfo, hasSupport bool) bool {
	if hasSupport && len(support.SupportedDBEngines) > 0 {
		for _, eng := range support.SupportedDBEngines {
			if strings.EqualFold(eng, "mariadb") {
				return true
			}
		}
		return false
	}

	// Static reference: only AWS, Alibaba, NHN, and OpenStack support MariaDB
	switch csp {
	case "aws", "alibaba", "nhn", "openstack":
		return true
	default:
		return false
	}
}
