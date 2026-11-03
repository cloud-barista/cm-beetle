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

// Package report provides NLB section markdown formatting for migration reports
package report

import (
	"fmt"
	"strings"
)

// writeNlbMigrationStatus writes the Network Load Balancers section when NLB mappings exist
func writeNlbMigrationStatus(md *strings.Builder, report *MigrationReport) {
	if report == nil || len(report.NlbMappings) == 0 {
		return
	}

	md.WriteString("## ⚖️ Network Load Balancers (NLBs)\n\n")
	md.WriteString(fmt.Sprintf("**Summary:** %d of %d load balancer(s) migrated from on-premise software load balancers to cloud-managed NLBs.\n\n",
		report.ExecutiveSummary.MigratedNLBs, report.ExecutiveSummary.TotalNLBs))

	for _, mapping := range report.NlbMappings {
		src := mapping.SourceNlb
		targetTitle := "Not Migrated"
		if mapping.TargetNlb != nil {
			targetTitle = mapping.TargetNlb.Name
		}

		md.WriteString(fmt.Sprintf("### %d. %s (%s) → %s\n\n",
			mapping.MappingID, src.BackendName, src.Software, targetTitle))

		hostLabel := src.HostHostname
		if hostLabel == "" {
			hostLabel = src.HostMachineId
		}
		srcLbVal := src.Software
		if hostLabel != "" {
			srcLbVal = fmt.Sprintf("%s (Host: %s)", src.Software, hostLabel)
		}

		bindAddr := src.BindAddress
		if bindAddr == "" {
			bindAddr = "*"
		}
		srcListenerVal := fmt.Sprintf("%s:%d (%s)", bindAddr, src.ListenerPort, strings.ToUpper(src.ListenerProto))
		srcBackendVal := fmt.Sprintf("%s (%s:%d, %s)", src.BackendName, strings.ToUpper(src.BackendProto), src.BackendPort, src.Balance)

		srcHcVal := "Disabled"
		if src.HealthCheck.Enabled {
			srcHcVal = fmt.Sprintf("Enabled (Interval: %ds, Timeout: %ds, Threshold: %d)",
				src.HealthCheck.Interval, src.HealthCheck.Timeout, src.HealthCheck.Threshold)
		}

		targetLbVal := "-"
		targetTypeVal := "-"
		targetListenerVal := "-"
		targetBackendVal := "-"
		targetHcVal := "-"
		lbNote := "Not migrated"
		typeNote := "Not migrated"
		listenerNote := "Not migrated"
		backendNote := "Not migrated"
		hcNote := "Not migrated"

		if tNlb := mapping.TargetNlb; tNlb != nil {
			cspIdent := tNlb.CspResourceName
			if cspIdent == "" {
				cspIdent = tNlb.CspResourceId
			}
			if cspIdent != "" {
				targetLbVal = fmt.Sprintf("%s (CSP ID: %s)", tNlb.Name, cspIdent)
			} else {
				targetLbVal = tNlb.Name
			}
			lbNote = "Replaced by managed NLB"

			targetTypeVal = fmt.Sprintf("%s / %s", tNlb.Type, tNlb.Scope)
			typeNote = "Cloud-native NLB"

			targetListenerVal = fmt.Sprintf("%s:%s", strings.ToUpper(tNlb.Listener.Protocol), tNlb.Listener.Port)
			var listenerNotes []string
			if tNlb.Listener.DnsName != "" {
				listenerNotes = append(listenerNotes, fmt.Sprintf("DNS: %s", tNlb.Listener.DnsName))
			}
			if tNlb.Listener.IP != "" {
				listenerNotes = append(listenerNotes, fmt.Sprintf("IP: %s", tNlb.Listener.IP))
			}
			if tNlb.Listener.Port != "" && tNlb.Listener.Port != fmt.Sprintf("%d", src.ListenerPort) {
				listenerNotes = append(listenerNotes, fmt.Sprintf("Port adjusted (%d -> %s) for CSP constraint", src.ListenerPort, tNlb.Listener.Port))
			}
			if len(listenerNotes) > 0 {
				listenerNote = strings.Join(listenerNotes, "; ")
			} else {
				listenerNote = "Preserved"
			}

			targetBackendVal = fmt.Sprintf("%s (%s:%s, %d nodes)",
				tNlb.TargetGroup.NodeGroupId, strings.ToUpper(tNlb.TargetGroup.Protocol), tNlb.TargetGroup.Port, len(tNlb.TargetGroup.Nodes))
			backendNote = "Target Group attached to NodeGroup"

			timeoutStr := fmt.Sprintf("%ds", tNlb.HealthChecker.Timeout)
			if tNlb.HealthChecker.Timeout < 0 {
				timeoutStr = "CSP Default"
			}
			hcProto := strings.ToUpper(tNlb.HealthChecker.Protocol)
			if hcProto == "" {
				hcProto = strings.ToUpper(tNlb.TargetGroup.Protocol)
			}
			hcPort := tNlb.HealthChecker.Port
			if hcPort == "" {
				hcPort = tNlb.TargetGroup.Port
			}
			targetHcVal = fmt.Sprintf("%s:%s (Interval: %ds, Timeout: %s, Threshold: %d)",
				hcProto, hcPort, tNlb.HealthChecker.Interval, timeoutStr, tNlb.HealthChecker.Threshold)

			if tNlb.HealthChecker.Timeout < 0 {
				hcNote = "Timeout uses CSP default"
			} else if src.HealthCheck.Enabled && tNlb.HealthChecker.Timeout > 0 && tNlb.HealthChecker.Timeout != int(src.HealthCheck.Timeout) {
				hcNote = fmt.Sprintf("Timeout adjusted (%ds -> %ds) for CSP constraint", src.HealthCheck.Timeout, tNlb.HealthChecker.Timeout)
			} else {
				hcNote = "Preserved"
			}
		}

		md.WriteString("**Configuration Comparison:**\n\n")
		md.WriteString("| Property | Source (On-Premise) | Target (Cloud NLB) | Note |\n")
		md.WriteString("|----------|---------------------|--------------------|------|\n")
		md.WriteString(fmt.Sprintf("| **Load Balancer / Name** | %s | %s | %s |\n", srcLbVal, targetLbVal, lbNote))
		md.WriteString(fmt.Sprintf("| **Type / Scope** | Software LB | %s | %s |\n", targetTypeVal, typeNote))
		md.WriteString(fmt.Sprintf("| **Frontend Listener** | %s | %s | %s |\n", srcListenerVal, targetListenerVal, listenerNote))
		md.WriteString(fmt.Sprintf("| **Backend / Target Group** | %s | %s | %s |\n", srcBackendVal, targetBackendVal, backendNote))
		md.WriteString(fmt.Sprintf("| **Health Check** | %s | %s | %s |\n\n", srcHcVal, targetHcVal, hcNote))

		if len(mapping.BackendNodeMappings) > 0 {
			md.WriteString("**Backend Node Mappings:**\n\n")
			md.WriteString("| No. | Source Backend Server | Source Endpoint | Target Node | Target CSP VM ID | Target Private IP | Target Public IP | Target Port | Status |\n")
			md.WriteString("|-----|-----------------------|-----------------|-------------|------------------|-------------------|------------------|-------------|--------|\n")

			for i, bm := range mapping.BackendNodeMappings {
				srcServerDisplay := bm.SourceServerName
				if bm.SourceHostname != "" {
					srcServerDisplay = fmt.Sprintf("%s (%s)", bm.SourceServerName, bm.SourceHostname)
				}
				targetNode := bm.TargetVmName
				if targetNode == "" {
					targetNode = "-"
				}
				targetCspID := bm.TargetCspVmID
				if targetCspID == "" {
					targetCspID = "-"
				}
				privIP := bm.TargetPrivateIP
				if privIP == "" {
					privIP = "-"
				}
				pubIP := bm.TargetPublicIP
				if pubIP == "" {
					pubIP = "-"
				}
				portStr := bm.TargetPort
				if portStr == "" {
					portStr = "-"
				}
				statusStr := bm.Status
				if statusStr == "" {
					statusStr = "-"
				}

				md.WriteString(fmt.Sprintf("| %d | %s | %s:%d | %s | %s | %s | %s | %s | %s |\n",
					i+1, srcServerDisplay, bm.SourceIP, bm.SourcePort, targetNode, targetCspID, privIP, pubIP, portStr, statusStr))
			}
			md.WriteString("\n")
		}

		if len(mapping.Notes) > 0 {
			md.WriteString("**Migration Notes:**\n")
			for _, note := range mapping.Notes {
				md.WriteString(fmt.Sprintf("- %s\n", note))
			}
			md.WriteString("\n")
		}
	}

	md.WriteString("---\n\n")
}
