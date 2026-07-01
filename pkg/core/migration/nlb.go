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
	"encoding/json"
	"fmt"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// Core functions
// ============================================================================

// CreateNlbs migrates NLBs to the target cloud infra.
// Each TargetNlb in req.NLBs is sent to Tumblebug as a separate NLBReq.
// All NLBs are attempted; individual failures are recorded in the result.
// If useExisting is true (default) and an NLB with the same nodeGroupId already exists, it is reused.
// (Tumblebug sets NLB Id = nodeGroupId, so GetNlb(nodeGroupId) is used for the existence check.)
func CreateNlbs(nsId, infraId string, req cloudmodel.RecommendedNlb, useExisting bool) (cloudmodel.MigratedNlbResult, error) {
	log.Info().
		Str("nsId", nsId).
		Str("infraId", infraId).
		Int("count", len(req.TargetNlbList)).
		Msg("Starting NLB migration")

	var createdList []cloudmodel.NLBInfo
	var errs []string

	for i, target := range req.TargetNlbList {
		log.Debug().
			Int("index", i+1).
			Int("total", len(req.TargetNlbList)).
			Str("type", target.Type).
			Str("listenerPort", target.Listener.Port).
			Str("nodeGroupId", target.TargetGroup.NodeGroupId).
			Msg("Creating NLB")

		// Validate nodeGroupId
		if target.TargetGroup.NodeGroupId == "" {
			msg := fmt.Sprintf("NLB[%d]: targetGroup.nodeGroupId is empty — skipping", i)
			log.Warn().Msg(msg)
			errs = append(errs, msg)
			continue
		}

		// If useExisting is true, check if an NLB with the same nodeGroupId already exists and reuse it.
		// Tumblebug sets NLB Id = nodeGroupId, so GetNlb(nodeGroupId) is the correct existence check.
		if useExisting {
			existingInfo, getErr := tbclient.NewSession().GetNlb(nsId, infraId, target.TargetGroup.NodeGroupId)
			if getErr == nil && existingInfo.Id != "" {
				log.Info().
					Str("nlbId", existingInfo.Id).
					Str("listenerPort", target.Listener.Port).
					Msg("NLB already exists. CM-Beetle will reuse it.")
				createdList = append(createdList, toNLBInfo(existingInfo))
				continue
			}
		}

		// NLB does not exist (or useExisting=false) — create it
		tbReq := toTumblebugNLBReq(target)
		info, err := tbclient.NewSession().CreateNlb(nsId, infraId, tbReq)
		if err != nil {
			msg := fmt.Sprintf("NLB[%d] (listenerPort=%s, nodeGroupId=%s): %v",
				i, target.Listener.Port, target.TargetGroup.NodeGroupId, err)
			log.Error().Err(err).Msg("Failed to create NLB")
			errs = append(errs, msg)
			continue
		}

		createdList = append(createdList, toNLBInfo(info))
		log.Info().
			Str("nlbId", info.Id).
			Str("listenerPort", target.Listener.Port).
			Msg("NLB created successfully")
	}

	// Determine overall status
	status := "created"
	switch {
	case len(createdList) == 0:
		status = "failed"
	case len(errs) > 0:
		status = "partial"
	}

	desc := fmt.Sprintf("%d NLB(s) created successfully", len(createdList))
	if len(errs) > 0 {
		desc += fmt.Sprintf("; %d failed", len(errs))
	}

	result := cloudmodel.MigratedNlbResult{
		Status:      status,
		Description: desc,
		NlbList:     createdList,
	}

	if status == "failed" {
		return result, fmt.Errorf("all NLB migrations failed: %s", strings.Join(errs, "; "))
	}

	log.Info().
		Str("nsId", nsId).
		Str("infraId", infraId).
		Int("created", len(createdList)).
		Int("failed", len(errs)).
		Msg("NLB migration completed")

	return result, nil
}

// ListNlbs returns all NLBs in the specified infra.
func ListNlbs(nsId, infraId string) ([]cloudmodel.NLBInfo, error) {
	log.Info().Str("nsId", nsId).Str("infraId", infraId).Msg("Listing NLBs")

	resp, err := tbclient.NewSession().ListNlbs(nsId, infraId)
	if err != nil {
		log.Error().Err(err).Str("nsId", nsId).Str("infraId", infraId).Msg("Failed to list NLBs")
		return nil, err
	}

	infos := make([]cloudmodel.NLBInfo, 0, len(resp.NLB))
	for _, item := range resp.NLB {
		infos = append(infos, toNLBInfo(item))
	}

	log.Info().Str("nsId", nsId).Str("infraId", infraId).Int("count", len(infos)).Msg("NLBs listed")
	return infos, nil
}

// GetNlb returns details of a specific NLB.
func GetNlb(nsId, infraId, nlbId string) (cloudmodel.NLBInfo, error) {
	log.Info().Str("nsId", nsId).Str("infraId", infraId).Str("nlbId", nlbId).Msg("Getting NLB")

	info, err := tbclient.NewSession().GetNlb(nsId, infraId, nlbId)
	if err != nil {
		log.Error().Err(err).Str("nlbId", nlbId).Msg("Failed to get NLB")
		return cloudmodel.NLBInfo{}, err
	}

	log.Info().Str("nlbId", nlbId).Msg("NLB retrieved")
	return toNLBInfo(info), nil
}

// GetNlbHealth performs a live health check on NLB targets via CSP.
// Unlike GetNlb (which returns cached state), this call reaches out to the CSP to get current health status.
func GetNlbHealth(nsId, infraId, nlbId string) (cloudmodel.NLBInfo, error) {
	log.Info().Str("nsId", nsId).Str("infraId", infraId).Str("nlbId", nlbId).Msg("Getting NLB health")

	info, err := tbclient.NewSession().GetNlbHealth(nsId, infraId, nlbId)
	if err != nil {
		log.Error().Err(err).Str("nlbId", nlbId).Msg("Failed to get NLB health")
		return cloudmodel.NLBInfo{}, err
	}

	log.Info().Str("nlbId", nlbId).Str("status", info.Status).Msg("NLB health retrieved")
	return toNLBInfo(info), nil
}

// DeleteNlb deletes a specific NLB. Treats 404 as already deleted (idempotent).
func DeleteNlb(nsId, infraId, nlbId string) error {
	log.Info().Str("nsId", nsId).Str("infraId", infraId).Str("nlbId", nlbId).Msg("Deleting NLB")

	err := tbclient.NewSession().DeleteNlb(nsId, infraId, nlbId)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			log.Info().Str("nlbId", nlbId).Msg("NLB not found; treating as already deleted")
			return nil
		}
		log.Error().Err(err).Str("nlbId", nlbId).Msg("Failed to delete NLB")
		return fmt.Errorf("failed to delete NLB '%s': %w", nlbId, err)
	}

	log.Info().Str("nlbId", nlbId).Msg("NLB deleted")

	// Some CSPs complete NLB deletion asynchronously; dependent resources (e.g. VNet/subnets)
	// may remain locked until the CSP finishes internal cleanup. Wait before returning.
	const nlbDeleteSettleTime = 15 * time.Second
	log.Debug().Str("nlbId", nlbId).Msgf("Waiting %s for CSP-side NLB resource cleanup", nlbDeleteSettleTime)
	time.Sleep(nlbDeleteSettleTime)

	return nil
}

// ============================================================================
// Converters
// ============================================================================

// toTumblebugNLBReq converts a cloudmodel.NlbReq to Tumblebug's NLBReq.
func toTumblebugNLBReq(t cloudmodel.NlbReq) tbmodel.NLBReq {
	return tbmodel.NLBReq{
		CspResourceId: t.CspResourceId,
		Description:   t.Description,
		Type:          t.Type,
		Scope:         t.Scope,
		Listener: tbmodel.NLBListenerReq{
			Protocol: t.Listener.Protocol,
			Port:     t.Listener.Port,
		},
		TargetGroup: tbmodel.NLBTargetGroupReq{
			Protocol:    t.TargetGroup.Protocol,
			Port:        t.TargetGroup.Port,
			NodeGroupId: t.TargetGroup.NodeGroupId,
		},
		HealthChecker: tbmodel.NLBHealthCheckerReq{
			Interval:  t.HealthChecker.Interval,
			Threshold: t.HealthChecker.Threshold,
			Timeout:   t.HealthChecker.Timeout,
		},
	}
}

// toNLBInfo converts a Tumblebug NLBInfo to cloudmodel.NLBInfo via JSON round-trip.
// Both types are structurally equivalent, so this avoids manual field mapping
// that would silently drift as either type evolves.
func toNLBInfo(src tbmodel.NLBInfo) cloudmodel.NLBInfo {
	data, err := json.Marshal(src)
	if err != nil {
		log.Warn().Err(err).Str("nlbId", src.Id).Msg("toNLBInfo: marshal failed")
		return cloudmodel.NLBInfo{}
	}
	var result cloudmodel.NLBInfo
	if err := json.Unmarshal(data, &result); err != nil {
		log.Warn().Err(err).Str("nlbId", src.Id).Msg("toNLBInfo: unmarshal failed")
	}
	return result
}
