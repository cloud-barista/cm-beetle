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

// Package summary provides target cloud NLB summary collection logic
package summary

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
)

// collectNlbResources retrieves and converts active NLBs for the specified target infrastructure
func collectNlbResources(nsId, infraId string) ([]SummaryNlbInfo, error) {
	nlbListResp, err := tbclient.NewSession().ListNlbs(nsId, infraId)
	if err != nil {
		if isNlbNotFoundError(err) {
			return []SummaryNlbInfo{}, nil
		}
		return nil, fmt.Errorf("failed to list NLBs for infra %s: %w", infraId, err)
	}

	nlbs := make([]SummaryNlbInfo, 0, len(nlbListResp.NLB))
	for _, item := range nlbListResp.NLB {
		// Skip NLB records that have a pending deletion tombstone.
		if item.DeletionRequestedAt != "" {
			continue
		}

		name := item.Name
		if name == "" {
			name = item.Id
		}

		createdTimeStr := ""
		if !item.CreatedTime.IsZero() {
			createdTimeStr = item.CreatedTime.Format(time.RFC3339)
		}

		nlbs = append(nlbs, SummaryNlbInfo{
			Id:              item.Id,
			Name:            name,
			CspResourceName: item.CspResourceName,
			CspResourceId:   item.CspResourceId,
			ConnectionName:  item.ConnectionName,
			Type:            item.Type,
			Scope:           item.Scope,
			Status:          item.Status,
			Description:     item.Description,
			Listener: SummaryNlbListenerInfo{
				Protocol: item.Listener.Protocol,
				Port:     item.Listener.Port,
				IP:       item.Listener.IP,
				DnsName:  item.Listener.DNSName,
			},
			TargetGroup: SummaryNlbTargetGroupInfo{
				Protocol:    item.TargetGroup.Protocol,
				Port:        item.TargetGroup.Port,
				NodeGroupId: item.TargetGroup.NodeGroupId,
				Nodes:       item.TargetGroup.Nodes,
			},
			HealthChecker: SummaryNlbHealthCheckerInfo{
				Protocol:  item.HealthChecker.Protocol,
				Port:      item.HealthChecker.Port,
				Interval:  item.HealthChecker.Interval,
				Timeout:   item.HealthChecker.Timeout,
				Threshold: item.HealthChecker.Threshold,
			},
			CreatedTime: createdTimeStr,
		})
	}

	// Sort NLBs deterministically by name.
	sort.Slice(nlbs, func(i, j int) bool {
		return nlbs[i].Name < nlbs[j].Name
	})

	return nlbs, nil
}

// isNlbNotFoundError checks whether an error indicates no NLB resources exist
func isNlbNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	errLower := strings.ToLower(err.Error())
	return strings.Contains(errLower, "404") ||
		strings.Contains(errLower, "not found") ||
		strings.Contains(errLower, "does not exist")
}
