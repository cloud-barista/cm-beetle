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
	"unicode"

	rdbmsmodel "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// GetRDBMSSupport retrieves CSP support matrix for managed RDBMS via CB-Tumblebug.
func GetRDBMSSupport(providerName string) (rdbmsmodel.RDBMSSupportResponse, error) {
	log.Debug().Str("providerName", providerName).Msg("Fetching RDBMS support matrix via CB-Tumblebug")
	return tbclient.NewSession().GetRDBMSSupport(providerName)
}

// GetRDBMSCapability retrieves real-time capability information for a connection via CB-Tumblebug.
func GetRDBMSCapability(connectionName string) (rdbmsmodel.RDBMSCapabilityResponse, error) {
	log.Debug().Str("connectionName", connectionName).Msg("Fetching RDBMS capability via CB-Tumblebug")
	return tbclient.NewSession().GetRDBMSCapability(connectionName)
}

// ValidateRDBMS performs dry-run validation and default autofill for an RDBMS create request.
// It executes a two-stage validation pipeline:
//  1. Beetle Deep Validation: Referential & cross-connection integrity checks (e.g. verifying VNet/SG connection matches target RDBMS connection)
//  2. Tumblebug Infra Validation: Dry-run check against CSP capability & API constraints via CB-Tumblebug.
func ValidateRDBMS(nsId string, req rdbmsmodel.RDBMSCreateRequest) (rdbmsmodel.RDBMSCreateRequest, error) {
	nsId = strings.TrimSpace(nsId)
	if nsId == "" {
		return req, fmt.Errorf("nsId is required")
	}
	targetConn := strings.TrimSpace(req.ConnectionName)
	if targetConn == "" {
		return req, fmt.Errorf("connectionName is required")
	}

	tbSess := tbclient.NewSession()

	// 1. Referential & Cross-connection check for VNet
	if req.VNetId != "" {
		vNetInfo, err := tbSess.ReadVNet(nsId, req.VNetId)
		if err == nil && vNetInfo.Id != "" {
			if vNetInfo.ConnectionName != "" && !strings.EqualFold(vNetInfo.ConnectionName, targetConn) {
				return req, fmt.Errorf("cross-connection mismatch: VNet '%s' belongs to connection '%s', but RDBMS target connection is '%s'",
					req.VNetId, vNetInfo.ConnectionName, targetConn)
			}
		}
	}

	// 2. Referential & Cross-connection check for Security Groups
	if len(req.SecurityGroupIds) > 0 {
		for _, sgId := range req.SecurityGroupIds {
			if sgId == "" {
				continue
			}
			sgInfo, err := tbSess.ReadSecurityGroup(nsId, sgId)
			if err == nil && sgInfo.Id != "" {
				if sgInfo.ConnectionName != "" && !strings.EqualFold(sgInfo.ConnectionName, targetConn) {
					return req, fmt.Errorf("cross-connection mismatch: SecurityGroup '%s' belongs to connection '%s', but RDBMS target connection is '%s'",
						sgId, sgInfo.ConnectionName, targetConn)
				}
			}
		}
	}

	// 3. NHN Cloud specific flag validation
	if req.NHNDBSGToAllowAllInbound {
		if !strings.HasPrefix(strings.ToLower(targetConn), "nhn") {
			return req, fmt.Errorf("nhnDBSGToAllowAllInbound is only supported for NHN Cloud")
		}
		if !req.PublicAccess {
			return req, fmt.Errorf("nhnDBSGToAllowAllInbound requires publicAccess=true")
		}
	}

	// 4. Pass to Tumblebug for live CSP capability strict dry-run validation (autoFillDefaults=false)
	req.AutoFillDefaults = false
	log.Info().Str("nsId", nsId).Str("rdbmsName", req.Name).Str("connection", targetConn).Msg("Validating RDBMS create request via CB-Tumblebug (strict mode: autoFillDefaults=false)")
	return tbSess.ValidateRDBMS(nsId, req)
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
		if targetEngine == "mariadb" && !isMariaDBSupported(support, hasSupport) {
			targetEngine = "mysql"
			warning := fmt.Sprintf("MariaDB is not supported on CSP '%s'. Recommended MySQL as fallback for instance '%s'.", desiredCsp, src.InstanceName)
			warnings = append(warnings, warning)
			log.Warn().Msg(warning)
		}

		// Fetch engine-specific live capability from Tumblebug
		connName := fmt.Sprintf("%s-%s", desiredCsp, desiredRegion)
		capaResp, capErr := tbclient.NewSession().GetRDBMSCapability(connName, targetEngine)
		if capErr != nil {
			log.Warn().Err(capErr).Str("connectionName", connName).Str("engine", targetEngine).Msg("Failed to fetch live RDBMS capability from Tumblebug")
		}
		capa := capaResp.Supports

		// Select Engine Version with robust fallback using live supportedVersions
		targetVersion := selectEngineVersion(targetEngine, src.EngineVersion, capa.SupportedVersions, &warnings, src.InstanceName)

		// 2. DB Spec Recommendation with capacity-aware best fit using live instanceSpecs / dbInstanceSpecOptions / DBMSRequirements
		targetSpec := selectDBInstanceSpecFromCapability(src.Vcpu, src.MemoryMb, targetEngine, capa)

		// 3. Storage Type Recommendation using live Notes.StorageTypes and StorageTypeOptions
		targetStorageType, selectedNote := selectStorageTypeFromCapability(src.StorageType, capa, &warnings, src.InstanceName)

		// Determine storage size boundaries using live capability info
		minStorage := capa.StorageSizeRange.Min
		maxStorage := capa.StorageSizeRange.Max
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

		// 4. Admin Credentials Default from Capability
		targetAdminUser := "cbuser"
		if capa.AdminUserNameRequirement != nil && capa.AdminUserNameRequirement.FixedValue != "" {
			targetAdminUser = capa.AdminUserNameRequirement.FixedValue
		} else if capa.AdminUserNameRequirement != nil && len(capa.AdminUserNameRequirement.ReservedValues) > 0 {
			for _, reserved := range capa.AdminUserNameRequirement.ReservedValues {
				if strings.EqualFold(targetAdminUser, reserved) {
					targetAdminUser = "dbadmin"
					break
				}
			}
		}

		// 5. Network & CSP-specific warnings
		if desiredCsp == "ncp" && src.PublicAccess {
			warning := fmt.Sprintf("NCP Cloud DB does not provide external public IP by default; instance '%s' will be created within private VPC.", src.InstanceName)
			warnings = append(warnings, warning)
		}
		targetNHNDBSG := src.NHNDBSGToAllowAllInbound
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

		targetBackupDays := src.BackupRetentionDays
		if strings.EqualFold(desiredCsp, "ibm") {
			targetBackupDays = 0 // IBM Cloud Databases does not support setting BackupRetentionDays during provisioning
		}

		targetInst := rdbmsmodel.TargetRDBMSInstance{
			SourceInstanceName:       src.InstanceName,
			SourceMachineId:          src.MachineId,
			RDBMSName:                targetName,
			DBEngine:                 targetEngine,
			DBEngineVersion:          targetVersion,
			DBInstanceSpec:           targetSpec,
			StorageType:              targetStorageType,
			StorageSize:              targetStorageSize,
			Iops:                     targetIops,
			AdminUserName:            targetAdminUser,
			HighAvailability:         src.HighAvailability,
			BackupRetentionDays:      targetBackupDays,
			PublicAccess:             src.PublicAccess,
			NHNDBSGToAllowAllInbound: targetNHNDBSG,
			Databases:                targetDatabases,
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
// selectEngineVersion selects the best-matching target database engine version from capability.
// It handles standard semver ("8.0", "8.0.32") as well as CSP-specific codes (e.g. NHN "MYSQL_V8032", Tencent "8.0").
func selectEngineVersion(engine, requestedVersion string, supported []string, warnings *[]string, instName string) string {
	requestedVersion = strings.TrimSpace(requestedVersion)
	if len(supported) == 0 {
		if requestedVersion != "" {
			return requestedVersion
		}
		if strings.EqualFold(engine, "mariadb") {
			return "10.6"
		}
		return "8.0"
	}

	// 1. Exact match (case-insensitive)
	for _, v := range supported {
		if strings.EqualFold(v, requestedVersion) {
			return v
		}
	}

	// 2. Direct substring / prefix / suffix match
	if requestedVersion != "" {
		for _, v := range supported {
			if strings.EqualFold(v, requestedVersion) ||
				strings.HasPrefix(v, requestedVersion) ||
				strings.HasPrefix(requestedVersion, v) ||
				strings.Contains(strings.ToLower(v), strings.ToLower(requestedVersion)) {
				return v
			}
		}
	}

	// 3. CSP Version Code Normalization match (e.g. "8.0" -> "80" matching "MYSQL_V8032", "5.7" -> "57" matching "MYSQL_V5744")
	if requestedVersion != "" {
		reqDigits := extractVersionDigits(requestedVersion)
		if len(reqDigits) >= 2 {
			for _, v := range supported {
				vDigits := extractVersionDigits(v)
				if strings.HasPrefix(vDigits, reqDigits) || strings.Contains(vDigits, reqDigits) {
					return v
				}
			}
		}
	}

	// 4. Preferred default version (e.g., 8.0 for MySQL, 10.6 for MariaDB)
	preferred := "8.0"
	preferredDigits := "80"
	if strings.EqualFold(engine, "mariadb") {
		preferred = "10.6"
		preferredDigits = "106"
	}

	for _, v := range supported {
		if strings.Contains(strings.ToLower(v), preferred) || strings.Contains(extractVersionDigits(v), preferredDigits) {
			return v
		}
	}

	// Fallback to first supported option with informational warning
	if requestedVersion != "" && !strings.EqualFold(supported[0], requestedVersion) {
		msg := fmt.Sprintf("Requested %s version '%s' could not be strictly matched. Selected supported version '%s' for instance '%s'.",
			engine, requestedVersion, supported[0], instName)
		*warnings = append(*warnings, msg)
	}

	return supported[0]
}

// extractVersionDigits extracts consecutive digit sequences for version matching (e.g. "8.0.35" -> "8035", "MYSQL_V8032" -> "8032")
func extractVersionDigits(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// selectStorageTypeFromCapability selects the optimal storage type and guidance note dynamically from Tumblebug capability.
func selectStorageTypeFromCapability(req string, capa rdbmsmodel.RDBMSMetaInfo, warnings *[]string, instName string) (string, *rdbmsmodel.StorageTypeNote) {
	// If the CSP does not support storage type selection (e.g. Azure, IBM, NCP), storageType must be omitted
	if !capa.SupportsStorageTypeSelection {
		return "", nil
	}

	// Filter out non-actionable placeholder options (e.g. "NA")
	usableOptions := make([]string, 0, len(capa.StorageTypeOptions))
	for _, opt := range capa.StorageTypeOptions {
		if !strings.EqualFold(opt, "NA") && strings.TrimSpace(opt) != "" {
			usableOptions = append(usableOptions, opt)
		}
	}
	if len(usableOptions) == 0 {
		return "", nil
	}

	// Collect storage notes map from notes or storageTypeGuidance
	notesMap := make(map[string]rdbmsmodel.StorageTypeNote)
	if capa.Notes != nil && len(capa.Notes.StorageTypes) > 0 {
		for _, note := range capa.Notes.StorageTypes {
			notesMap[strings.ToLower(note.StorageType)] = note
		}
	}
	if len(capa.StorageTypeGuidance) > 0 {
		for k, v := range capa.StorageTypeGuidance {
			notesMap[strings.ToLower(k)] = v
		}
	}

	// 1. Check if user requested a specific type and if it is supported in capability
	if req != "" && !strings.EqualFold(req, "NA") {
		reqLower := strings.ToLower(req)
		// Check against usable capability options
		for _, opt := range usableOptions {
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
	}

	// 2. Select recommended storage from notes if flagged recommended
	for _, note := range notesMap {
		if note.Recommended && !strings.EqualFold(note.StorageType, "NA") {
			return note.StorageType, &note
		}
	}

	// 3. Select defaultStorageType if provided by capability
	if capa.DefaultStorageType != "" && !strings.EqualFold(capa.DefaultStorageType, "NA") {
		if note, ok := notesMap[strings.ToLower(capa.DefaultStorageType)]; ok {
			return capa.DefaultStorageType, &note
		}
		return capa.DefaultStorageType, nil
	}

	// 4. Fallback to first usable storage option from capability
	first := usableOptions[0]
	if note, ok := notesMap[strings.ToLower(first)]; ok {
		return first, &note
	}
	return first, nil
}

// selectDBInstanceSpecFromCapability selects the best-fitting instance spec dynamically from Tumblebug capability using resource fit.
func selectDBInstanceSpecFromCapability(vcpu, memoryMb int, engine string, capa rdbmsmodel.RDBMSMetaInfo) string {
	engineLower := strings.ToLower(engine)

	// 1. If detailed DBInstanceSpecs list is available, evaluate requirements
	if len(capa.DBInstanceSpecs) > 0 {
		var bestSpec string
		var minDiff int64 = 1<<62 - 1

		for _, spec := range capa.DBInstanceSpecs {
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

		// If no spec strictly satisfies the requirements, check ReferenceDBInstanceSpec in DBMSRequirements
		if req, ok := capa.DBMSRequirements[engineLower]; ok && req.ReferenceDBInstanceSpec != "" {
			return req.ReferenceDBInstanceSpec
		}

		// Fallback to the first spec from capability
		return capa.DBInstanceSpecs[0].Name
	}

	// 2. Check engine-specific ReferenceDBInstanceSpec in DBMSRequirements from Tumblebug capability
	if req, ok := capa.DBMSRequirements[engineLower]; ok && req.ReferenceDBInstanceSpec != "" {
		return req.ReferenceDBInstanceSpec
	}

	// 3. If string options list is provided from Tumblebug capability, return the first valid option
	for _, opt := range capa.DBInstanceSpecOptions {
		if strings.TrimSpace(opt) != "" {
			return opt
		}
	}

	return ""
}

// isMariaDBSupported checks if the target CSP supports MariaDB using Tumblebug support info.
func isMariaDBSupported(support rdbmsmodel.RDBMSCSPSupportInfo, hasSupport bool) bool {
	if hasSupport && len(support.SupportedDBEngines) > 0 {
		for _, eng := range support.SupportedDBEngines {
			if strings.EqualFold(eng, "mariadb") {
				return true
			}
		}
	}
	return false
}
