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
	"strconv"
	"strings"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// GetRDBMSSupport retrieves CSP support matrix for managed RDBMS via CB-Tumblebug.
func GetRDBMSSupport(cspType string) (rdbmsmodel.RDBMSSupportResponse, error) {
	log.Debug().Str("cspType", cspType).Msg("Fetching RDBMS support matrix via CB-Tumblebug")
	return tbclient.NewSession().GetRDBMSSupport(cspType)
}

// GetRDBMSCapability retrieves real-time capability information for a connection via CB-Tumblebug.
func GetRDBMSCapability(connectionName string) (rdbmsmodel.RDBMSCapabilityResponse, error) {
	log.Debug().Str("connectionName", connectionName).Msg("Fetching RDBMS capability via CB-Tumblebug")
	return tbclient.NewSession().GetRDBMSCapability(connectionName)
}

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

		// Fetch engine-specific live capability from Tumblebug
		connName := fmt.Sprintf("%s-%s", desiredCsp, desiredRegion)
		capResp, capErr := tbclient.NewSession().GetRDBMSCapability(connName, targetEngine)
		if capErr != nil {
			log.Warn().Err(capErr).Str("connectionName", connName).Str("engine", targetEngine).Msg("Failed to fetch live RDBMS capability from Tumblebug")
		}
		cap := capResp.Supports

		// Select Engine Version with robust fallback using live supportedVersions
		targetVersion := selectEngineVersion(targetEngine, src.EngineVersion, cap.SupportedVersions, &warnings, src.InstanceName)

		// 2. DB Spec Recommendation with capacity-aware best fit using live instanceSpecs / dbInstanceSpecOptions
		targetSpec := selectDBSpecFromCapability(src.Vcpu, src.MemoryMb, cap)

		// 3. Storage Type Recommendation using live Notes.StorageTypes and StorageTypeOptions
		targetStorageType, selectedNote := selectStorageTypeFromCapability(src.StorageType, cap, &warnings, src.InstanceName)

		// Determine storage size boundaries using live capability info
		minStorage := cap.StorageSizeRange.Min
		maxStorage := cap.StorageSizeRange.Max
		if selectedNote != nil {
			if selectedNote.MinSize > 0 {
				minStorage = selectedNote.MinSize
			}
			if selectedNote.MaxSize > 0 {
				maxStorage = selectedNote.MaxSize
			}
		}

		targetStorageSize := src.StorageSizeGb
		if minStorage > 0 && targetStorageSize < minStorage {
			warning := fmt.Sprintf("Adjusted storage size for instance '%s' from %dGB to minimum %dGB required by target cloud (%s).",
				src.InstanceName, targetStorageSize, minStorage, targetStorageType)
			warnings = append(warnings, warning)
			targetStorageSize = minStorage
		}
		if maxStorage > 0 && targetStorageSize > maxStorage {
			warning := fmt.Sprintf("Clamped storage size for instance '%s' from %dGB to maximum %dGB supported by target cloud.",
				src.InstanceName, targetStorageSize, maxStorage)
			warnings = append(warnings, warning)
			targetStorageSize = maxStorage
		}

		// IOPS handling: dynamic validation using live StorageTypeNote constraints
		targetIops := ""
		if selectedNote != nil && selectedNote.RequiresIops {
			assignedIops := 3000
			if selectedNote.IopsRange != nil && selectedNote.IopsRange.Min > 0 {
				assignedIops = selectedNote.IopsRange.Min
			}
			if src.Iops > 0 {
				assignedIops = src.Iops
				if selectedNote.IopsRange != nil {
					if selectedNote.IopsRange.Min > 0 && assignedIops < selectedNote.IopsRange.Min {
						assignedIops = selectedNote.IopsRange.Min
					}
					if selectedNote.IopsRange.Max > 0 && assignedIops > selectedNote.IopsRange.Max {
						assignedIops = selectedNote.IopsRange.Max
					}
				}
			}
			targetIops = fmt.Sprintf("%d", assignedIops)
		}

		// 4. Admin Credentials Default
		targetAdminUser := "cbuser"
		if cap.AdminUserNameRequirement != nil && cap.AdminUserNameRequirement.FixedValue != "" {
			targetAdminUser = cap.AdminUserNameRequirement.FixedValue
		} else if desiredCsp == "ibm" {
			targetAdminUser = "admin"
		} else if desiredCsp == "tencent" {
			targetAdminUser = "root"
		} else if desiredCsp == "azure" || desiredCsp == "alibaba" || desiredCsp == "ncp" {
			targetAdminUser = "dbadmin"
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

// selectEngineVersion chooses the best-matching supported version from capability, or a standard fallback.
func selectEngineVersion(engine, requestedVersion string, supported []string, warnings *[]string, instName string) string {
	requestedVersion = strings.TrimSpace(requestedVersion)
	if len(supported) == 0 {
		if requestedVersion != "" {
			return requestedVersion
		}
		if engine == "mariadb" {
			return "10.6"
		}
		return "8.0"
	}

	// 1. Exact match
	for _, v := range supported {
		if strings.EqualFold(v, requestedVersion) {
			return v
		}
	}

	// 2. Prefix match (e.g. requested "8.0.35" matching "8.0" or requested "8" matching "8.0")
	if requestedVersion != "" {
		for _, v := range supported {
			if strings.HasPrefix(requestedVersion, v) || strings.HasPrefix(v, requestedVersion) {
				return v
			}
		}
		// Warning on version mismatch fallback
		msg := fmt.Sprintf("Requested %s version '%s' is not supported on target cloud. Selected supported version '%s' instead for instance '%s'.",
			engine, requestedVersion, supported[0], instName)
		*warnings = append(*warnings, msg)
	}

	// 3. Preferred default version if available in supported list
	preferred := "8.0"
	if engine == "mariadb" {
		preferred = "10.6"
	}
	for _, v := range supported {
		if strings.HasPrefix(v, preferred) {
			return v
		}
	}

	// Fallback to first supported option
	return supported[0]
}

// selectStorageTypeFromCapability selects the optimal storage type and guidance note dynamically from Tumblebug capability.
func selectStorageTypeFromCapability(requestedType string, cap rdbmsmodel.RDBMSMetaInfo, warnings *[]string, instName string) (string, *rdbmsmodel.StorageTypeNote) {
	req := strings.TrimSpace(requestedType)

	// Collect storage notes map from notes or storageTypeGuidance
	notesMap := make(map[string]rdbmsmodel.StorageTypeNote)
	if cap.Notes != nil && len(cap.Notes.StorageTypes) > 0 {
		for _, note := range cap.Notes.StorageTypes {
			notesMap[strings.ToLower(note.StorageType)] = note
		}
	}
	if len(cap.StorageTypeGuidance) > 0 {
		for k, v := range cap.StorageTypeGuidance {
			notesMap[strings.ToLower(k)] = v
		}
	}

	// 1. Check if user requested a specific type and if it is supported in capability
	if req != "" {
		reqLower := strings.ToLower(req)
		// Check against capability options
		if len(cap.StorageTypeOptions) > 0 {
			for _, opt := range cap.StorageTypeOptions {
				if strings.EqualFold(opt, req) {
					if note, ok := notesMap[reqLower]; ok {
						return opt, &note
					}
					return opt, nil
				}
			}
			msg := fmt.Sprintf("Requested storage type '%s' is not supported on target cloud. Replaced with capability recommended storage for instance '%s'.",
				req, instName)
			*warnings = append(*warnings, msg)
		} else if note, ok := notesMap[reqLower]; ok {
			return note.StorageType, &note
		}
	}

	// 2. Select recommended storage from notes if flagged recommended
	for _, note := range notesMap {
		if note.Recommended {
			return note.StorageType, &note
		}
	}

	// 3. Select defaultStorageType if provided by capability
	if cap.DefaultStorageType != "" {
		if note, ok := notesMap[strings.ToLower(cap.DefaultStorageType)]; ok {
			return cap.DefaultStorageType, &note
		}
		return cap.DefaultStorageType, nil
	}

	// 4. Fallback to first storage option from capability
	if len(cap.StorageTypeOptions) > 0 {
		first := cap.StorageTypeOptions[0]
		if note, ok := notesMap[strings.ToLower(first)]; ok {
			return first, &note
		}
		return first, nil
	}

	return "gp2", nil
}

// selectDBSpecFromCapability selects the best-fitting instance spec dynamically from Tumblebug capability using resource fit.
func selectDBSpecFromCapability(vcpu, memoryMb int, cap rdbmsmodel.RDBMSMetaInfo) string {
	// If detailed DBInstanceSpecs list is available, evaluate requirements
	if len(cap.DBInstanceSpecs) > 0 {
		var bestSpec string
		var minDiff int64 = 1<<62 - 1

		for _, spec := range cap.DBInstanceSpecs {
			specCpu, _ := strconv.Atoi(spec.VCpuCount)
			specMem, _ := strconv.Atoi(spec.MemSizeMiB)
			if specCpu == 0 {
				specCpu = 2
			}
			if specMem == 0 {
				specMem = 4096
			}

			// Evaluate specs that satisfy the required compute resources
			if specCpu >= vcpu && specMem >= memoryMb {
				cpuDiff := int64(specCpu - vcpu)
				memDiff := int64(specMem - memoryMb)
				// Combined resource distance: minimize excess allocation
				diff := (cpuDiff * 10000) + memDiff
				if diff < minDiff {
					minDiff = diff
					bestSpec = spec.Name
				}
			}
		}

		if bestSpec != "" {
			return bestSpec
		}

		// Fallback to the first spec if no spec strictly satisfies the requirements
		return cap.DBInstanceSpecs[0].Name
	}

	// If only string options list is provided, return the first option
	if len(cap.DBInstanceSpecOptions) > 0 {
		return cap.DBInstanceSpecOptions[0]
	}

	return "db.t3.medium"
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

	// Static reference fallback
	switch csp {
	case "aws", "alibaba", "nhn", "openstack":
		return true
	default:
		return false
	}
}
