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

// Package report provides NLB migration matching and analysis logic
package report

import (
	"fmt"
	"strings"

	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
)

// buildNlbMigrationMappings correlates source load balancers with target cloud NLBs using lineage labels
func buildNlbMigrationMappings(sourceSummary *summary.SourceInfraSummary, targetSummary *summary.TargetInfraSummary) []NlbMigrationMapping {
	if sourceSummary == nil || len(sourceSummary.NetworkResources.NLBs) == 0 {
		return nil
	}

	vmByMachineID := make(map[string]*summary.SummaryVmInfo)
	nodeGroupToBackend := make(map[string]string)
	nodeGroupToMachineIDs := make(map[string]map[string]bool)

	if targetSummary != nil {
		for i := range targetSummary.ComputeResources.Vms {
			vm := &targetSummary.ComputeResources.Vms[i]
			machineID := resolveSourceMachineID(*vm)
			if machineID != "" {
				vmByMachineID[machineID] = vm
			}

			ngID := vm.NodeGroupId
			if ngID == "" {
				ngID = inferNodeGroupIDFromVmName(vm.Name)
			}
			if ngID == "" {
				continue
			}

			if nodeGroupToMachineIDs[ngID] == nil {
				nodeGroupToMachineIDs[ngID] = make(map[string]bool)
			}
			if machineID != "" {
				nodeGroupToMachineIDs[ngID][machineID] = true
			}

			if vm.Label != nil {
				if backend := strings.TrimSpace(vm.Label[lineage.LabelNlbBackend]); backend != "" {
					nodeGroupToBackend[ngID] = backend
				}
				if rawIDs := vm.Label[lineage.LabelSourceMachineIds]; rawIDs != "" {
					for _, id := range strings.Split(rawIDs, ",") {
						trimmed := strings.TrimSpace(id)
						if trimmed != "" {
							nodeGroupToMachineIDs[ngID][trimmed] = true
						}
					}
				}
			}
		}
	}

	usedTargetIdx := make(map[int]bool)
	mappings := make([]NlbMigrationMapping, 0, len(sourceSummary.NetworkResources.NLBs))

	for idx, srcNlb := range sourceSummary.NetworkResources.NLBs {
		var matchedTarget *summary.SummaryNlbInfo
		if targetSummary != nil && len(targetSummary.NetworkResources.NLBs) > 0 {
			matchedIdx := matchTargetNlbIndex(srcNlb, targetSummary.NetworkResources.NLBs, nodeGroupToBackend, nodeGroupToMachineIDs, usedTargetIdx)
			if matchedIdx >= 0 {
				usedTargetIdx[matchedIdx] = true
				targetCopy := targetSummary.NetworkResources.NLBs[matchedIdx]
				matchedTarget = &targetCopy
			}
		}

		hostVmName := ""
		if srcNlb.HostMachineId != "" {
			if hostVM := findTargetVmByMachineID(srcNlb.HostMachineId, vmByMachineID, targetSummary); hostVM != nil {
				hostVmName = hostVM.Name
			}
		}

		backendMappings := make([]NlbBackendNodeMapping, 0, len(srcNlb.BackendServers))
		for _, srv := range srcNlb.BackendServers {
			targetPort := fmt.Sprintf("%d", srv.Port)
			if matchedTarget != nil && matchedTarget.TargetGroup.Port != "" {
				targetPort = matchedTarget.TargetGroup.Port
			}

			nodeMapping := NlbBackendNodeMapping{
				SourceServerName: srv.Name,
				SourceIP:         srv.IP,
				SourcePort:       srv.Port,
				SourceHostname:   srv.Hostname,
				SourceMachineID:  srv.MachineId,
				TargetPort:       targetPort,
			}

			if targetVM := findTargetVmByMachineID(srv.MachineId, vmByMachineID, targetSummary); targetVM != nil {
				nodeMapping.TargetVmName = targetVM.Name
				nodeMapping.TargetCspVmID = targetVM.CspVmId
				nodeMapping.TargetPrivateIP = targetVM.Misc.PrivateIp
				nodeMapping.TargetPublicIP = targetVM.Misc.PublicIp
				nodeMapping.Status = targetVM.Status
			}

			backendMappings = append(backendMappings, nodeMapping)
		}

		status := "Not Migrated"
		if matchedTarget != nil {
			status = "Migrated"
		}

		targetCloud := "Cloud"
		if targetSummary != nil && targetSummary.Overview.TargetCloud != "" {
			targetCloud = targetSummary.Overview.TargetCloud
		}

		notes := buildNlbMigrationNotes(srcNlb, matchedTarget, hostVmName, targetCloud)

		mappings = append(mappings, NlbMigrationMapping{
			MappingID:            idx + 1,
			SourceNlb:            srcNlb,
			TargetNlb:            matchedTarget,
			HostNodeTargetVmName: hostVmName,
			BackendNodeMappings:  backendMappings,
			MigrationStatus:      status,
			Notes:                notes,
		})
	}

	return mappings
}

// matchTargetNlbIndex finds the best matching target NLB index using 2-stage deterministic lineage matching
func matchTargetNlbIndex(
	srcNlb summary.SourceNlbInfo,
	targetNlbs []summary.SummaryNlbInfo,
	nodeGroupToBackend map[string]string,
	nodeGroupToMachineIDs map[string]map[string]bool,
	usedTargetIdx map[int]bool,
) int {
	srcBackendPort := fmt.Sprintf("%d", srcNlb.BackendPort)
	srcListenerPort := fmt.Sprintf("%d", srcNlb.ListenerPort)

	// Stage 1: Exact match via nlbBackend label on the target NLB's NodeGroup.
	firstBackendMatch := -1
	for i, tNlb := range targetNlbs {
		if usedTargetIdx[i] {
			continue
		}
		ngID := tNlb.TargetGroup.NodeGroupId
		if ngID != "" && nodeGroupToBackend[ngID] == srcNlb.BackendName {
			if firstBackendMatch < 0 {
				firstBackendMatch = i
			}
			if tNlb.TargetGroup.Port == srcBackendPort || tNlb.Listener.Port == srcListenerPort {
				return i
			}
		}
	}
	if firstBackendMatch >= 0 {
		return firstBackendMatch
	}

	// Stage 2: Match via cm-source-machine-ids intersection and backend port.
	for i, tNlb := range targetNlbs {
		if usedTargetIdx[i] {
			continue
		}
		if tNlb.TargetGroup.Port != srcBackendPort {
			continue
		}
		ngID := tNlb.TargetGroup.NodeGroupId
		machineSet := nodeGroupToMachineIDs[ngID]
		if len(machineSet) == 0 {
			continue
		}
		for _, srv := range srcNlb.BackendServers {
			if srv.MachineId != "" && machineSet[srv.MachineId] {
				return i
			}
		}
	}

	return -1
}

// buildNlbMigrationNotes generates architectural and CSP-specific adjustment notes for an NLB mapping
func buildNlbMigrationNotes(srcNlb summary.SourceNlbInfo, targetNlb *summary.SummaryNlbInfo, hostVmName, targetCloud string) []string {
	var notes []string

	hostDisplay := srcNlb.HostHostname
	if hostDisplay == "" {
		hostDisplay = srcNlb.HostMachineId
	}
	if hostDisplay != "" {
		if hostVmName != "" {
			notes = append(notes, fmt.Sprintf("On-premise %s on %s (migrated to %s) is replaced by cloud-managed NLB.", srcNlb.Software, hostDisplay, hostVmName))
		} else {
			notes = append(notes, fmt.Sprintf("On-premise %s on %s is replaced by cloud-managed NLB.", srcNlb.Software, hostDisplay))
		}
	}

	if targetNlb == nil {
		notes = append(notes, "No corresponding cloud NLB was found in the target infrastructure.")
		return notes
	}

	srcListenerPort := fmt.Sprintf("%d", srcNlb.ListenerPort)
	if targetNlb.Listener.Port != "" && targetNlb.Listener.Port != srcListenerPort {
		notes = append(notes, fmt.Sprintf("Listener port was adjusted from %s to %s due to %s forwarding rule port constraints.", srcListenerPort, targetNlb.Listener.Port, targetCloud))
	}

	if targetNlb.HealthChecker.Timeout < 0 {
		notes = append(notes, fmt.Sprintf("Health check timeout uses CSP default setting on %s.", targetCloud))
	} else if srcNlb.HealthCheck.Enabled && targetNlb.HealthChecker.Timeout > 0 && targetNlb.HealthChecker.Timeout != int(srcNlb.HealthCheck.Timeout) {
		notes = append(notes, fmt.Sprintf("Health check timeout was adjusted from %ds to %ds to satisfy %s health check constraints (timeout < interval).", srcNlb.HealthCheck.Timeout, targetNlb.HealthChecker.Timeout, targetCloud))
	}

	return notes
}

// computeNlbMigrationStatus calculates aggregate NLB migration status and counts for ExecutiveSummary
func computeNlbMigrationStatus(mappings []NlbMigrationMapping) (status string, total, migrated, failed int) {
	total = len(mappings)
	if total == 0 {
		return "", 0, 0, 0
	}

	for _, m := range mappings {
		if m.MigrationStatus == "Migrated" {
			migrated++
		}
	}
	failed = total - migrated

	switch {
	case migrated == total:
		status = "Completed"
	case migrated > 0:
		status = "Partially Completed"
	default:
		status = "Not Migrated"
	}

	return status, total, migrated, failed
}

// findTargetVmByMachineID resolves a target VM by source MachineId using label lookup and name fallback
func findTargetVmByMachineID(machineID string, vmByMachineID map[string]*summary.SummaryVmInfo, targetSummary *summary.TargetInfraSummary) *summary.SummaryVmInfo {
	if machineID == "" {
		return nil
	}
	if vm, ok := vmByMachineID[machineID]; ok {
		return vm
	}
	if targetSummary == nil {
		return nil
	}
	for i := range targetSummary.ComputeResources.Vms {
		if strings.Contains(targetSummary.ComputeResources.Vms[i].Name, machineID) {
			return &targetSummary.ComputeResources.Vms[i]
		}
	}
	return nil
}

// inferNodeGroupIDFromVmName strips a trailing "-<digits>" suffix from a VM name when NodeGroupId is unset
func inferNodeGroupIDFromVmName(vmName string) string {
	idx := strings.LastIndexByte(vmName, '-')
	if idx <= 0 || idx == len(vmName)-1 {
		return ""
	}
	suffix := vmName[idx+1:]
	for i := 0; i < len(suffix); i++ {
		if suffix[i] < '0' || suffix[i] > '9' {
			return ""
		}
	}
	return vmName[:idx]
}
