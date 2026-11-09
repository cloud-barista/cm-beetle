package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
)

// ExecuteTwoStageCleanup deletes the migrated infrastructure and sweeps any residual SG, SSH Key, and VNet resources.
func ExecuteTwoStageCleanup(
	client *resty.Client,
	endpoint string,
	nsID string,
	infraID string,
	seededInfra *cloudmodel.RecommendedInfra,
) (StepPhaseResult, CleanupSummary) {
	startTime := time.Now()
	summary := CleanupSummary{
		Attempted: true,
		Logs:      make([]string, 0),
	}

	phase := StepPhaseResult{
		Phase:     "Cleanup",
		StartTime: startTime,
	}

	orphans := 0

	// Stage 1: Delete the top-level multi-cloud infrastructure (MCI).
	if infraID != "" {
		deleted, statusCode, msg := deleteInfraWithFallback(client, endpoint, nsID, infraID)
		summary.InfraDeleted = deleted
		phase.StatusCode = statusCode
		summary.Logs = append(summary.Logs, msg)
		if !deleted {
			orphans++
		}
	}

	// Stage 2: Sweep child resources directly in reverse dependency order (SG -> SSH Key -> VNet).
	if seededInfra != nil {
		for _, sg := range seededInfra.TargetSecurityGroupList {
			if sg.Name == "" {
				continue
			}
			ok, logMsg := deleteSecurityGroupWithRetry(client, endpoint, nsID, sg.Name)
			summary.Logs = append(summary.Logs, logMsg)
			if ok {
				summary.SecurityGroupsCleaned++
			} else {
				orphans++
			}
		}

		if seededInfra.TargetSshKey.Name != "" {
			ok, logMsg := deleteSSHKey(client, endpoint, nsID, seededInfra.TargetSshKey.Name)
			summary.Logs = append(summary.Logs, logMsg)
			if ok {
				summary.SshKeysCleaned++
			} else {
				orphans++
			}
		}

		if seededInfra.TargetVNet.Name != "" {
			ok, logMsg := deleteVNetWithRetry(client, endpoint, nsID, seededInfra.TargetVNet.Name)
			summary.Logs = append(summary.Logs, logMsg)
			if ok {
				summary.VNetsCleaned++
			} else {
				orphans++
			}
		}
	}

	endTime := time.Now()
	durationMs := endTime.Sub(startTime).Milliseconds()

	summary.OrphanResourcesRemaining = orphans
	summary.CleanVerified = orphans == 0
	summary.DurationMs = durationMs
	summary.SimulatedSec = float64(durationMs) / 1000.0

	phase.EndTime = endTime
	phase.DurationMs = durationMs
	phase.SimulatedSec = summary.SimulatedSec
	phase.Success = summary.CleanVerified
	if summary.CleanVerified {
		phase.Details = fmt.Sprintf("2-stage cleanup verified (Infra=%t, SGs=%d, SSHKeys=%d, VNets=%d, Orphans=0)",
			summary.InfraDeleted, summary.SecurityGroupsCleaned, summary.SshKeysCleaned, summary.VNetsCleaned)
	} else {
		phase.ErrorMessage = fmt.Sprintf("Cleanup finished with %d residual resource(s)", orphans)
		phase.Details = strings.Join(summary.Logs, " | ")
	}

	return phase, summary
}

func deleteInfraWithFallback(client *resty.Client, endpoint, nsID, infraID string) (bool, int, string) {
	const maxRetries = 3
	lastStatus := 0

	for attempt := 1; attempt <= maxRetries; attempt++ {
		if attempt > 1 {
			time.Sleep(time.Duration(attempt*5) * time.Second)
		}

		option := "terminate"
		if attempt == maxRetries {
			option = "force"
		}

		url := fmt.Sprintf("%s/beetle/migration/ns/%s/infra/%s?option=%s", endpoint, nsID, infraID, option)
		resp, err := client.R().Delete(url)
		if resp != nil {
			lastStatus = resp.StatusCode()
		}

		if err == nil && resp != nil && (resp.IsSuccess() || isNotFoundResponse(resp.StatusCode(), resp.String())) {
			return true, resp.StatusCode(), fmt.Sprintf("Infra %s cleaned (option=%s, status=%d)", infraID, option, resp.StatusCode())
		}

		log.Warn().Err(err).Str("infraId", infraID).Int("attempt", attempt).Msg("Infra deletion attempt did not complete cleanly")
	}

	return false, lastStatus, fmt.Sprintf("Infra %s deletion failed after %d attempts", infraID, maxRetries)
}

func deleteSecurityGroupWithRetry(client *resty.Client, endpoint, nsID, sgID string) (bool, string) {
	const maxRetries = 4
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/resources/securityGroup/%s", endpoint, nsID, sgID)

	for attempt := 1; attempt <= maxRetries; attempt++ {
		if attempt > 1 {
			time.Sleep(4 * time.Second)
		}

		resp, err := client.R().Delete(url)
		if err == nil && resp != nil && (resp.IsSuccess() || isNotFoundResponse(resp.StatusCode(), resp.String())) {
			return true, fmt.Sprintf("SecurityGroup %s swept", sgID)
		}
	}

	return false, fmt.Sprintf("SecurityGroup %s sweep failed", sgID)
}

func deleteSSHKey(client *resty.Client, endpoint, nsID, sshKeyID string) (bool, string) {
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/resources/sshKey/%s", endpoint, nsID, sshKeyID)
	resp, err := client.R().Delete(url)
	if err == nil && resp != nil && (resp.IsSuccess() || isNotFoundResponse(resp.StatusCode(), resp.String())) {
		return true, fmt.Sprintf("SSHKey %s swept", sshKeyID)
	}
	return false, fmt.Sprintf("SSHKey %s sweep failed", sshKeyID)
}

func deleteVNetWithRetry(client *resty.Client, endpoint, nsID, vNetID string) (bool, string) {
	const maxRetries = 4

	for attempt := 1; attempt <= maxRetries; attempt++ {
		if attempt > 1 {
			time.Sleep(5 * time.Second)
		}

		action := "withsubnets"
		if attempt == maxRetries {
			action = "force"
		}

		url := fmt.Sprintf("%s/beetle/migration/ns/%s/resources/vNet/%s?action=%s", endpoint, nsID, vNetID, action)
		resp, err := client.R().Delete(url)
		if err == nil && resp != nil && (resp.IsSuccess() || isNotFoundResponse(resp.StatusCode(), resp.String())) {
			return true, fmt.Sprintf("VNet %s swept (action=%s)", vNetID, action)
		}
	}

	return false, fmt.Sprintf("VNet %s sweep failed", vNetID)
}

func isNotFoundResponse(statusCode int, body string) bool {
	if statusCode == http.StatusNotFound {
		return true
	}
	lower := strings.ToLower(body)
	return strings.Contains(lower, "does not exist") || strings.Contains(lower, "not found") || strings.Contains(lower, "already deleted")
}
