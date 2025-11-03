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

// Package summary provides source infrastructure summary markdown generation
package summary

import (
	"fmt"
	"strings"
)

// GenerateSourceMarkdownSummary converts SourceInfraSummary to markdown format
func GenerateSourceMarkdownSummary(summary *SourceInfraSummary) string {
	var md strings.Builder

	// Title and Metadata
	md.WriteString("# Source Infrastructure Summary\n\n")
	md.WriteString(fmt.Sprintf("**Generated At:** %s\n\n", summary.SummaryMetadata.GeneratedAt.Format("2006-01-02 15:04:05")))
	md.WriteString(fmt.Sprintf("**Infrastructure Name:** %s\n\n", summary.SummaryMetadata.InfraName))
	md.WriteString(fmt.Sprintf("**Summary Version:** %s\n\n", summary.SummaryMetadata.SummaryVersion))
	md.WriteString("---\n\n")

	// Overview Section
	md.WriteString("## Overview\n\n")
	md.WriteString(generateSourceOverviewMarkdown(&summary.Overview))
	md.WriteString("\n")

	// Compute Resources Section
	md.WriteString("## Compute Resources\n\n")
	md.WriteString(generateSourceComputeResourcesMarkdown(&summary.ComputeResources))
	md.WriteString("\n")

	// Network Resources Section
	md.WriteString("## Network Resources\n\n")
	md.WriteString(generateSourceNetworkResourcesMarkdown(&summary.NetworkResources))
	md.WriteString(generateSourceServerNetworkDetailsMarkdown(&summary.ComputeResources))
	md.WriteString("\n")

	// Security Resources Section
	md.WriteString("## Security Resources\n\n")
	md.WriteString(generateSourceSecurityResourcesMarkdown(&summary.SecurityResources))
	md.WriteString("\n")

	// Storage Resources Section
	md.WriteString("## Storage Resources\n\n")
	md.WriteString(generateSourceStorageResourcesMarkdown(&summary.StorageResources))
	md.WriteString("\n")

	return md.String()
}

// generateSourceOverviewMarkdown generates markdown for source infrastructure overview
func generateSourceOverviewMarkdown(overview *SourceSummaryInfraOverview) string {
	var md strings.Builder

	md.WriteString("| Metric | Value |\n")
	md.WriteString("|--------|-------|\n")
	md.WriteString(fmt.Sprintf("| Infrastructure Name | %s |\n", overview.InfraName))
	md.WriteString(fmt.Sprintf("| Total Servers | %d |\n", overview.TotalServerCount))
	md.WriteString(fmt.Sprintf("| Total CPU Cores | %d |\n", overview.TotalCPUCores))
	md.WriteString(fmt.Sprintf("| Total Memory (GB) | %d |\n", overview.TotalMemoryGB))
	md.WriteString(fmt.Sprintf("| Total Disk (GB) | %d |\n", overview.TotalDiskGB))
	md.WriteString(fmt.Sprintf("| Total Networks | %d |\n", overview.TotalNetworks))

	return md.String()
}

// generateSourceComputeResourcesMarkdown generates markdown for source compute resources
func generateSourceComputeResourcesMarkdown(resources *SourceSummaryComputeResources) string {
	var md strings.Builder

	md.WriteString(fmt.Sprintf("### Servers (%d)\n\n", len(resources.Servers)))

	for i, server := range resources.Servers {
		md.WriteString(fmt.Sprintf("#### %d. %s\n\n", i+1, server.Hostname))

		// Server details table
		md.WriteString("| Component | Details |\n")
		md.WriteString("|-----------|----------|\n")

		// CPU with detailed info
		md.WriteString(fmt.Sprintf("| CPU | %s |\n", server.CPU.Model))
		md.WriteString(fmt.Sprintf("| **CPU CPUs** | %d |\n", server.CPU.CPUs))
		md.WriteString(fmt.Sprintf("| CPU Cores | %d |\n", server.CPU.Cores))
		md.WriteString(fmt.Sprintf("| **CPU Threads** | %d |\n", server.CPU.Threads))
		if server.CPU.MaxSpeed > 0 {
			md.WriteString(fmt.Sprintf("| CPU Speed | %.2f GHz |\n", server.CPU.MaxSpeed))
		}
		if server.CPU.Architecture != "" {
			md.WriteString(fmt.Sprintf("| Architecture | %s |\n", server.CPU.Architecture))
		}

		// Memory
		md.WriteString(fmt.Sprintf("| **Memory** | %d GB", server.Memory.TotalGB))
		if server.Memory.Type != "" {
			md.WriteString(fmt.Sprintf(" (%s)", server.Memory.Type))
		}
		md.WriteString(" |\n")

		// Disk
		if server.Disk.TotalGB > 0 {
			md.WriteString(fmt.Sprintf("| **Disk** | %d GB (%s) |\n", server.Disk.TotalGB, server.Disk.Type))
		}

		// OS
		md.WriteString(fmt.Sprintf("| **OS** | %s %s |\n", server.OS.Name, server.OS.Version))

		// Network brief
		if server.Network.IPAddress != "" {
			md.WriteString(fmt.Sprintf("| **Primary IP** | %s |\n", server.Network.IPAddress))
		}

		md.WriteString("\n")
	}

	return md.String()
}

// generateSourceNetworkResourcesMarkdown generates markdown for source network resources
func generateSourceNetworkResourcesMarkdown(resources *SourceSummaryNetworkResources) string {
	var md strings.Builder

	md.WriteString(fmt.Sprintf("### Networks (%d)\n\n", len(resources.Networks)))

	if len(resources.Networks) == 0 {
		md.WriteString("*No networks defined*\n\n")
		return md.String()
	}

	for i, network := range resources.Networks {
		md.WriteString(fmt.Sprintf("#### %d. %s\n\n", i+1, network.NetworkName))

		md.WriteString("| Property | Value |\n")
		md.WriteString("|----------|-------|\n")
		md.WriteString(fmt.Sprintf("| Network CIDR | %s |\n", network.NetworkCIDR))

		if network.Gateway != "" {
			md.WriteString(fmt.Sprintf("| Gateway | %s |\n", network.Gateway))
		}

		if len(network.DNSServers) > 0 {
			md.WriteString(fmt.Sprintf("| DNS Servers | %s |\n", strings.Join(network.DNSServers, ", ")))
		}

		md.WriteString(fmt.Sprintf("| Connected Servers | %d |\n", network.ServerCount))

		if network.Description != "" {
			md.WriteString(fmt.Sprintf("| Description | %s |\n", network.Description))
		}

		md.WriteString("\n")
	}

	return md.String()
}

// generateSourceServerNetworkDetailsMarkdown generates markdown for server network interfaces and routing
func generateSourceServerNetworkDetailsMarkdown(resources *SourceSummaryComputeResources) string {
	var md strings.Builder

	md.WriteString(fmt.Sprintf("### Network Details by Server (%d servers)\n\n", len(resources.Servers)))

	for i, server := range resources.Servers {
		md.WriteString(fmt.Sprintf("#### %d. %s\n\n", i+1, server.Hostname))

		// Network Interfaces - Only show active interfaces with IP
		if len(server.Interfaces) > 0 {
			activeInterfaces := []SourceInterfaceInfo{}
			for _, iface := range server.Interfaces {
				// Show only interfaces with IP addresses or in 'up' state
				if len(iface.IPv4CidrBlocks) > 0 || iface.State == "up" {
					activeInterfaces = append(activeInterfaces, iface)
				}
			}

			if len(activeInterfaces) > 0 {
				md.WriteString("**Active Interfaces:**\n\n")
				md.WriteString("| Interface | IP Address | State |\n")
				md.WriteString("|-----------|------------|-------|\n")

				for _, iface := range activeInterfaces {
					ipAddr := strings.Join(iface.IPv4CidrBlocks, ", ")
					if ipAddr == "" {
						ipAddr = "-"
					}
					state := iface.State
					if state == "" {
						state = "-"
					}

					md.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
						iface.Name, ipAddr, state))
				}
				md.WriteString("\n")
			}
		}

		// Routing Table - Only show IPv4 default route and main routes
		if len(server.RoutingTable) > 0 {
			mainRoutes := []SourceRoutingTableRow{}
			for _, route := range server.RoutingTable {
				// Show only IPv4 routes (exclude IPv6) and important routes
				if !strings.Contains(route.Destination, ":") {
					// Include default route (0.0.0.0/0) and /24 or smaller networks
					if route.Destination == "0.0.0.0/0" ||
						(strings.Contains(route.Destination, "/") &&
							!strings.HasPrefix(route.Destination, "fe80") &&
							!strings.HasPrefix(route.Destination, "ff00")) {
						mainRoutes = append(mainRoutes, route)
					}
				}
			}

			if len(mainRoutes) > 0 {
				md.WriteString("**Main Routes:**\n\n")
				md.WriteString("| Destination | Gateway | Interface |\n")
				md.WriteString("|-------------|---------|-----------|\n")

				for _, route := range mainRoutes {
					gateway := route.Gateway
					if gateway == "" {
						gateway = "on-link"
					}

					md.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
						route.Destination, gateway, route.Interface))
				}
				md.WriteString("\n")
			}
		}
	}

	return md.String()
}

// generateSourceStorageResourcesMarkdown generates markdown for source storage resources
func generateSourceStorageResourcesMarkdown(resources *SourceSummaryStorageResources) string {
	var md strings.Builder

	// Storage by server (먼저 출력)
	if len(resources.ByServer) > 0 {
		md.WriteString(fmt.Sprintf("### Storage by Server (%d servers)\n\n", len(resources.ByServer)))
		md.WriteString("| Hostname | RootDisk (GB) | Type |\n")
		md.WriteString("|----------|---------------|------|\n")

		for _, storage := range resources.ByServer {
			md.WriteString(fmt.Sprintf("| %s | %d | %s |\n",
				storage.Hostname, storage.TotalGB, storage.Type))
		}
		md.WriteString("\n")
	}

	// Storage by type (나중에 출력)
	if len(resources.ByType) > 0 {
		md.WriteString("### Storage by Type\n\n")
		md.WriteString("| Type | Total (GB) | Servers |\n")
		md.WriteString("|------|------------|----------|\n")

		for _, storage := range resources.ByType {
			md.WriteString(fmt.Sprintf("| %s | %d | %d |\n",
				storage.Type, storage.TotalGB, storage.ServerCount))
		}
	}

	return md.String()
}

// generateSourceSecurityResourcesMarkdown generates markdown for source security resources
func generateSourceSecurityResourcesMarkdown(resources *SourceSummarySecurityResources) string {
	var md strings.Builder

	if len(resources.ServerFirewalls) == 0 {
		md.WriteString("*No firewall rules configured*\n\n")
		return md.String()
	}

	md.WriteString(fmt.Sprintf("### Firewall Rules by Server (%d servers)\n\n", len(resources.ServerFirewalls)))

	for i, serverFw := range resources.ServerFirewalls {
		md.WriteString(fmt.Sprintf("#### %d. %s\n\n", i+1, serverFw.Hostname))

		if len(serverFw.FirewallRules) == 0 {
			md.WriteString("*No firewall rules*\n\n")
			continue
		}

		md.WriteString(fmt.Sprintf("**Firewall Rules:** (%d rules)\n\n", len(serverFw.FirewallRules)))
		md.WriteString("| Direction | Protocol | Source | Src Ports | Destination | Dst Ports | Action |\n")
		md.WriteString("|-----------|----------|--------|-----------|-------------|-----------|--------|\n")

		for _, rule := range serverFw.FirewallRules {
			md.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s |\n",
				rule.Direction,
				rule.Protocol,
				rule.SrcCIDR,
				rule.SrcPorts,
				rule.DstCIDR,
				rule.DstPorts,
				rule.Action))
		}

		md.WriteString("\n")
	}

	return md.String()
}
