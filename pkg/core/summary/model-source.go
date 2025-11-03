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

// Package summary provides source infrastructure summary data models
package summary

import (
	"time"
)

// SourceInfraSummary represents the comprehensive source infrastructure summary (on-premise)
type SourceInfraSummary struct {
	SummaryMetadata   SourceSummaryMetadata          `json:"summaryMetadata"`
	Overview          SourceSummaryInfraOverview     `json:"overview"`
	ComputeResources  SourceSummaryComputeResources  `json:"computeResources"`
	NetworkResources  SourceSummaryNetworkResources  `json:"networkResources"`
	StorageResources  SourceSummaryStorageResources  `json:"storageResources"`
	SecurityResources SourceSummarySecurityResources `json:"securityResources"`
}

// SourceSummaryMetadata contains source infrastructure summary generation metadata
type SourceSummaryMetadata struct {
	GeneratedAt    time.Time `json:"generatedAt" example:"2025-10-31T14:30:00Z"`
	InfraName      string    `json:"infraName" example:"datacenter-seoul-1"`
	SummaryVersion string    `json:"summaryVersion" example:"1.0"`
}

// SourceSummaryInfraOverview provides high-level overview of the source infrastructure
type SourceSummaryInfraOverview struct {
	InfraName        string `json:"infraName" example:"datacenter-seoul-1"`
	TotalServerCount int    `json:"totalServerCount" example:"5"`
	TotalCPUCores    int    `json:"totalCpuCores" example:"40"`
	TotalMemoryGB    int    `json:"totalMemoryGb" example:"160"`
	TotalDiskGB      int    `json:"totalDiskGb" example:"2000"`
	TotalNetworks    int    `json:"totalNetworks" example:"3"`
}

// SourceSummaryComputeResources contains source server information
type SourceSummaryComputeResources struct {
	Servers []SourceServerInfo `json:"servers"`
}

// SourceServerInfo represents source server information
type SourceServerInfo struct {
	Hostname     string                  `json:"hostname" example:"web-server-01"`
	CPU          SourceCPUInfo           `json:"cpu"`
	Memory       SourceMemoryInfo        `json:"memory"`
	Disk         SourceDiskInfo          `json:"disk"`
	OS           SourceOSInfo            `json:"os"`
	RootDiskType string                  `json:"rootDiskType" example:"SSD"`
	RootDiskSize int                     `json:"rootDiskSize" example:"100"`
	Network      SourceNetworkBrief      `json:"network"`
	Interfaces   []SourceInterfaceInfo   `json:"interfaces,omitempty"`
	RoutingTable []SourceRoutingTableRow `json:"routingTable,omitempty"`
}

// SourceCPUInfo represents CPU information
type SourceCPUInfo struct {
	Cores        int     `json:"cores" example:"8"`
	Model        string  `json:"model" example:"Intel(R) Xeon(R) CPU E5-2680 v4 @ 2.40GHz"`
	MaxSpeed     float64 `json:"maxSpeed" example:"2.4"`
	Vendor       string  `json:"vendor" example:"GenuineIntel"`
	CPUs         int     `json:"cpus" example:"1"`
	Threads      int     `json:"threads" example:"16"`
	Architecture string  `json:"architecture" example:"x86_64"`
	Utilization  float64 `json:"utilization,omitempty" example:"45.2"`
}

// SourceMemoryInfo represents memory information
type SourceMemoryInfo struct {
	Type        string  `json:"type" example:"DDR4"`
	TotalGB     int     `json:"totalGb" example:"32"`
	Available   int     `json:"available,omitempty" example:"20"`
	Used        int     `json:"used,omitempty" example:"12"`
	Utilization float64 `json:"utilization,omitempty" example:"37.5"`
}

// SourceDiskInfo represents disk information
type SourceDiskInfo struct {
	Label       string  `json:"label" example:"/dev/sda"`
	Type        string  `json:"type" example:"SSD"`
	TotalGB     int     `json:"totalGb" example:"500"`
	Used        int     `json:"used,omitempty" example:"300"`
	Available   int     `json:"available,omitempty" example:"200"`
	Utilization float64 `json:"utilization,omitempty" example:"60.0"`
}

// SourceOSInfo represents OS information
type SourceOSInfo struct {
	Name         string `json:"name" example:"Ubuntu"`
	Version      string `json:"version" example:"22.04"`
	Architecture string `json:"architecture" example:"x86_64"`
	Kernel       string `json:"kernel,omitempty" example:"5.15.0-91-generic"`
}

// SourceNetworkBrief represents brief network information for a server
type SourceNetworkBrief struct {
	IPAddress   string `json:"ipAddress" example:"192.168.1.10"`
	PublicIP    string `json:"publicIp,omitempty" example:"203.0.113.45"`
	MacAddress  string `json:"macAddress,omitempty" example:"00:1a:2b:3c:4d:5e"`
	NetworkName string `json:"networkName,omitempty" example:"office-lan"`
}

// SourceSummaryNetworkResources contains source network information
type SourceSummaryNetworkResources struct {
	Networks []SourceNetworkInfo `json:"networks"`
}

// SourceNetworkInfo represents source network information
type SourceNetworkInfo struct {
	NetworkName string   `json:"networkName" example:"office-lan"`
	NetworkCIDR string   `json:"networkCidr" example:"192.168.1.0/24"`
	Gateway     string   `json:"gateway,omitempty" example:"192.168.1.1"`
	DNSServers  []string `json:"dnsServers,omitempty" example:"8.8.8.8,8.8.4.4"`
	ServerCount int      `json:"serverCount" example:"5"`
	Description string   `json:"description,omitempty" example:"Office LAN network"`
}

// SourceSummaryStorageResources contains storage breakdown information
type SourceSummaryStorageResources struct {
	TotalGB     int                     `json:"totalGb" example:"2000"`
	UsedGB      int                     `json:"usedGb" example:"1200"`
	AvailableGB int                     `json:"availableGb" example:"800"`
	ByType      []SourceStorageByType   `json:"byType"`
	ByServer    []SourceStorageByServer `json:"byServer"`
}

// SourceStorageByType represents storage breakdown by type (SSD, HDD, etc.)
type SourceStorageByType struct {
	Type        string `json:"type" example:"SSD"`
	TotalGB     int    `json:"totalGb" example:"1000"`
	ServerCount int    `json:"serverCount" example:"3"`
}

// SourceStorageByServer represents storage breakdown by server
type SourceStorageByServer struct {
	Hostname string `json:"hostname" example:"web-server-01"`
	TotalGB  int    `json:"totalGb" example:"500"`
	Type     string `json:"type" example:"SSD"`
}

// SourceInterfaceInfo represents network interface information
type SourceInterfaceInfo struct {
	Name           string   `json:"name" example:"eth0"`
	IPAddress      string   `json:"ipAddress,omitempty" example:"192.168.1.10"`
	IPv4CidrBlocks []string `json:"ipv4CidrBlocks,omitempty" example:"192.168.1.10/24"`
	IPv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty" example:"fe80::1/64"`
	MacAddress     string   `json:"macAddress,omitempty" example:"00:1a:2b:3c:4d:5e"`
	MTU            int      `json:"mtu,omitempty" example:"1500"`
	State          string   `json:"state,omitempty" example:"up"`
}

// SourceRoutingTableRow represents a routing table entry
type SourceRoutingTableRow struct {
	Destination string `json:"destination" example:"0.0.0.0/0"`
	Gateway     string `json:"gateway,omitempty" example:"192.168.1.1"`
	Interface   string `json:"interface" example:"eth0"`
	Metric      int    `json:"metric,omitempty" example:"100"`
	Protocol    string `json:"protocol,omitempty" example:"kernel"`
}

// SourceSummarySecurityResources contains source security information
type SourceSummarySecurityResources struct {
	ServerFirewalls []SourceServerFirewall `json:"serverFirewalls"`
}

// SourceServerFirewall represents firewall configuration for a server
type SourceServerFirewall struct {
	Hostname      string               `json:"hostname" example:"web-server-01"`
	FirewallRules []SourceFirewallRule `json:"firewallRules"`
}

// SourceFirewallRule represents a firewall rule
type SourceFirewallRule struct {
	SrcCIDR   string `json:"srcCidr" example:"0.0.0.0/0"`
	SrcPorts  string `json:"srcPorts" example:"*"`
	DstCIDR   string `json:"dstCidr" example:"0.0.0.0/0"`
	DstPorts  string `json:"dstPorts" example:"22,80,443"`
	Protocol  string `json:"protocol" example:"tcp"`
	Direction string `json:"direction" example:"inbound"`
	Action    string `json:"action" example:"allow"`
}
