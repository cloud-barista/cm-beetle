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

// Package summary provides source infrastructure summary generation logic
package summary

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
)

// GenerateSourceInfraSummary generates a comprehensive source infrastructure summary from on-premise data
func GenerateSourceInfraSummary(infraName string, infraData onpremmodel.OnpremInfra) (*SourceInfraSummary, error) {
	log.Info().Msgf("Generating source infrastructure summary (infraName: %s)", infraName)

	// Step 1: Build summary metadata
	metadata := SourceSummaryMetadata{
		GeneratedAt:    time.Now(),
		InfraName:      infraName,
		SummaryVersion: "1.0",
	}

	// Step 2: Build infrastructure overview
	overview := buildSourceInfraOverview(infraName, &infraData)

	// Step 3: Collect compute resources
	computeResources := collectSourceComputeResources(&infraData)

	// Step 4: Collect network resources
	networkResources := collectSourceNetworkResources(&infraData)

	// Step 5: Calculate storage resources
	storageResources := calculateSourceStorageResources(&infraData)

	// Step 6: Collect security resources
	securityResources := collectSourceSecurityResources(&infraData)

	// Step 7: Assemble final summary
	sourceSummary := &SourceInfraSummary{
		SummaryMetadata:   metadata,
		Overview:          overview,
		ComputeResources:  computeResources,
		NetworkResources:  networkResources,
		StorageResources:  storageResources,
		SecurityResources: securityResources,
	}

	log.Info().Msgf("Successfully generated source infrastructure summary: %s", infraName)
	return sourceSummary, nil
}

// buildSourceInfraOverview builds the overview section from on-premise infrastructure data
func buildSourceInfraOverview(infraName string, infraData *onpremmodel.OnpremInfra) SourceSummaryInfraOverview {
	totalServerCount := len(infraData.Servers)
	totalCPUCores := 0
	totalMemoryGB := 0
	totalDiskGB := 0

	// Count unique networks from IPv4Networks CIDRs and DefaultGateways
	networkCount := len(infraData.Network.IPv4Networks.CidrBlocks)
	if networkCount == 0 {
		networkCount = len(infraData.Network.IPv4Networks.DefaultGateways)
	}

	// Calculate totals from servers
	for _, server := range infraData.Servers {
		// CPU cores (uint32 to int conversion)
		totalCPUCores += int(server.CPU.Cores)

		// Memory in GB
		totalMemoryGB += int(server.Memory.TotalSize)

		// Root disk size
		totalDiskGB += int(server.RootDisk.TotalSize)

		// Data disks
		for _, disk := range server.DataDisks {
			totalDiskGB += int(disk.TotalSize)
		}
	}

	return SourceSummaryInfraOverview{
		InfraName:        infraName,
		TotalServerCount: totalServerCount,
		TotalCPUCores:    totalCPUCores,
		TotalMemoryGB:    totalMemoryGB,
		TotalDiskGB:      totalDiskGB,
		TotalNetworks:    networkCount,
	}
}

// collectSourceComputeResources collects server information from on-premise data
func collectSourceComputeResources(infraData *onpremmodel.OnpremInfra) SourceSummaryComputeResources {
	var servers []SourceServerInfo

	for _, server := range infraData.Servers {
		// CPU information
		cpuInfo := SourceCPUInfo{
			Cores:        int(server.CPU.Cores),
			Model:        server.CPU.Model,
			MaxSpeed:     float64(server.CPU.MaxSpeed),
			Vendor:       server.CPU.Vendor,
			CPUs:         int(server.CPU.Cpus),
			Threads:      int(server.CPU.Threads),
			Architecture: server.CPU.Architecture,
		}

		// Memory information
		memoryInfo := SourceMemoryInfo{
			Type:    server.Memory.Type,
			TotalGB: int(server.Memory.TotalSize),
		}

		// Root disk information
		diskInfo := SourceDiskInfo{
			Label:   server.RootDisk.Label,
			Type:    server.RootDisk.Type,
			TotalGB: int(server.RootDisk.TotalSize),
		}

		// OS information
		osInfo := SourceOSInfo{
			Name:         server.OS.Name,
			Version:      server.OS.VersionID,
			Architecture: server.CPU.Architecture,
		}

		// Network brief (use first active network interface if available)
		networkBrief := SourceNetworkBrief{}
		for _, iface := range server.Interfaces {
			// Skip loopback and virtual interfaces for brief
			if iface.Name != "lo" && len(iface.IPv4CidrBlocks) > 0 {
				networkBrief.IPAddress = iface.IPv4CidrBlocks[0]
				networkBrief.MacAddress = iface.MacAddress
				networkBrief.NetworkName = iface.Name
				break
			}
		}

		// Collect main network interfaces (lo, eth0, eno*, br-ex, etc.)
		var interfaces []SourceInterfaceInfo
		for _, iface := range server.Interfaces {
			// Filter main interfaces
			if isMainInterface(iface.Name) {
				ipAddress := ""
				if len(iface.IPv4CidrBlocks) > 0 {
					ipAddress = iface.IPv4CidrBlocks[0]
				}

				interfaces = append(interfaces, SourceInterfaceInfo{
					Name:           iface.Name,
					IPAddress:      ipAddress,
					IPv4CidrBlocks: iface.IPv4CidrBlocks,
					IPv6CidrBlocks: iface.IPv6CidrBlocks,
					MacAddress:     iface.MacAddress,
					MTU:            int(iface.Mtu),
					State:          iface.State,
				})
			}
		}

		// Collect routing table entries for main interfaces
		var routingTable []SourceRoutingTableRow
		for _, route := range server.RoutingTable {
			// Include routes for main interfaces
			if isMainInterface(route.Interface) {
				routingTable = append(routingTable, SourceRoutingTableRow{
					Destination: route.Destination,
					Gateway:     route.Gateway,
					Interface:   route.Interface,
					Metric:      int(route.Metric),
					Protocol:    route.Protocol,
				})
			}
		}

		serverInfo := SourceServerInfo{
			Hostname:     server.Hostname,
			CPU:          cpuInfo,
			Memory:       memoryInfo,
			Disk:         diskInfo,
			OS:           osInfo,
			RootDiskType: server.RootDisk.Type,
			RootDiskSize: int(server.RootDisk.TotalSize),
			Network:      networkBrief,
			Interfaces:   interfaces,
			RoutingTable: routingTable,
		}

		servers = append(servers, serverInfo)
	}

	return SourceSummaryComputeResources{
		Servers: servers,
	}
}

// isMainInterface checks if the interface is a main/primary interface
func isMainInterface(name string) bool {
	// Main interfaces typically include: lo, eth*, eno*, enp*, br-ex, etc.
	// Exclude tap*, veth*, and other virtual interfaces
	if name == "" {
		return false
	}

	mainPrefixes := []string{"lo", "eth", "eno", "enp", "br-"}
	for _, prefix := range mainPrefixes {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}

	return false
}

// collectSourceNetworkResources collects deduplicated network information from on-premise data
func collectSourceNetworkResources(infraData *onpremmodel.OnpremInfra) SourceSummaryNetworkResources {
	var networks []SourceNetworkInfo

	// Use a map to deduplicate networks by CIDR
	networkMap := make(map[string]*SourceNetworkInfo)

	// Collect networks from infraData.Network.IPv4Networks.CidrBlocks
	if len(infraData.Network.IPv4Networks.CidrBlocks) > 0 {
		for i, cidr := range infraData.Network.IPv4Networks.CidrBlocks {
			if _, exists := networkMap[cidr]; !exists {
				// Find corresponding gateway
				gateway := ""
				if i < len(infraData.Network.IPv4Networks.DefaultGateways) {
					gateway = infraData.Network.IPv4Networks.DefaultGateways[i].IP
				}

				networkMap[cidr] = &SourceNetworkInfo{
					NetworkName: fmt.Sprintf("network-%d", len(networkMap)+1),
					NetworkCIDR: cidr,
					Gateway:     gateway,
					DNSServers:  []string{}, // DNS info not available in current model
					ServerCount: 0,
					Description: fmt.Sprintf("Network %s", cidr),
				}
			}
		}
	}

	// Collect additional networks from DefaultGateways if no CIDRBlocks
	if len(infraData.Network.IPv4Networks.CidrBlocks) == 0 && len(infraData.Network.IPv4Networks.DefaultGateways) > 0 {
		for _, gw := range infraData.Network.IPv4Networks.DefaultGateways {
			// Estimate CIDR from gateway (e.g., 192.168.1.1 -> 192.168.1.0/24)
			cidr := estimateCIDRFromGateway(gw.IP)

			if _, exists := networkMap[cidr]; !exists {
				networkMap[cidr] = &SourceNetworkInfo{
					NetworkName: fmt.Sprintf("network-%d", len(networkMap)+1),
					NetworkCIDR: cidr,
					Gateway:     gw.IP,
					DNSServers:  []string{}, // DNS info not available in current model
					ServerCount: 0,
					Description: fmt.Sprintf("Network with gateway %s", gw.IP),
				}
			}
		}
	}

	// Count servers per network by matching their interface IPs to network CIDRs
	for cidr := range networkMap {
		serverCount := 0
		for _, server := range infraData.Servers {
			serverInNetwork := false
			for _, iface := range server.Interfaces {
				for _, ipCidr := range iface.IPv4CidrBlocks {
					// Simple check: if IP CIDR contains network CIDR prefix
					if strings.Contains(ipCidr, strings.Split(cidr, "/")[0][:strings.LastIndex(strings.Split(cidr, "/")[0], ".")]) {
						serverInNetwork = true
						break
					}
				}
				if serverInNetwork {
					break
				}
			}
			if serverInNetwork {
				serverCount++
			}
		}
		networkMap[cidr].ServerCount = serverCount
	}

	// Convert map to slice
	for _, network := range networkMap {
		networks = append(networks, *network)
	}

	return SourceSummaryNetworkResources{
		Networks: networks,
	}
}

// estimateCIDRFromGateway estimates CIDR block from gateway IP (simple /24 assumption)
func estimateCIDRFromGateway(gateway string) string {
	parts := strings.Split(gateway, ".")
	if len(parts) == 4 {
		return fmt.Sprintf("%s.%s.%s.0/24", parts[0], parts[1], parts[2])
	}
	return gateway + "/24"
}

// calculateSourceStorageResources calculates storage breakdown from on-premise data
func calculateSourceStorageResources(infraData *onpremmodel.OnpremInfra) SourceSummaryStorageResources {
	totalGB := 0
	storageByTypeMap := make(map[string]*SourceStorageByType)
	var storageByServer []SourceStorageByServer

	for _, server := range infraData.Servers {
		serverTotalGB := 0

		// Root disk
		rootDiskSize := int(server.RootDisk.TotalSize)
		rootDiskType := server.RootDisk.Type
		serverTotalGB += rootDiskSize
		totalGB += rootDiskSize

		// Aggregate by type
		if _, exists := storageByTypeMap[rootDiskType]; !exists {
			storageByTypeMap[rootDiskType] = &SourceStorageByType{
				Type:        rootDiskType,
				TotalGB:     0,
				ServerCount: 0,
			}
		}
		storageByTypeMap[rootDiskType].TotalGB += rootDiskSize

		// Data disks
		for _, disk := range server.DataDisks {
			diskSize := int(disk.TotalSize)
			diskType := disk.Type
			serverTotalGB += diskSize
			totalGB += diskSize

			if _, exists := storageByTypeMap[diskType]; !exists {
				storageByTypeMap[diskType] = &SourceStorageByType{
					Type:        diskType,
					TotalGB:     0,
					ServerCount: 0,
				}
			}
			storageByTypeMap[diskType].TotalGB += diskSize
		}

		// Add server storage summary
		if serverTotalGB > 0 {
			storageByServer = append(storageByServer, SourceStorageByServer{
				Hostname: server.Hostname,
				TotalGB:  serverTotalGB,
				Type:     rootDiskType,
			})

			// Count unique servers per disk type
			if entry, exists := storageByTypeMap[rootDiskType]; exists {
				entry.ServerCount++
			}
		}
	}

	// Convert map to slice
	var storageByType []SourceStorageByType
	for _, v := range storageByTypeMap {
		storageByType = append(storageByType, *v)
	}

	return SourceSummaryStorageResources{
		TotalGB:     totalGB,
		UsedGB:      0, // Not available from OnpremInfra
		AvailableGB: 0, // Not available from OnpremInfra
		ByType:      storageByType,
		ByServer:    storageByServer,
	}
}

// collectSourceSecurityResources collects firewall rules from on-premise servers
func collectSourceSecurityResources(infraData *onpremmodel.OnpremInfra) SourceSummarySecurityResources {
	var serverFirewalls []SourceServerFirewall

	for _, server := range infraData.Servers {
		var firewallRules []SourceFirewallRule

		// Collect firewall rules if they exist
		for _, rule := range server.FirewallTable {
			firewallRules = append(firewallRules, SourceFirewallRule{
				SrcCIDR:   rule.SrcCIDR,
				SrcPorts:  rule.SrcPorts,
				DstCIDR:   rule.DstCIDR,
				DstPorts:  rule.DstPorts,
				Protocol:  rule.Protocol,
				Direction: rule.Direction,
				Action:    rule.Action,
			})
		}

		// Add server to list (even if no firewall rules)
		serverFirewalls = append(serverFirewalls, SourceServerFirewall{
			Hostname:      server.Hostname,
			FirewallRules: firewallRules,
		})
	}

	return SourceSummarySecurityResources{
		ServerFirewalls: serverFirewalls,
	}
}
