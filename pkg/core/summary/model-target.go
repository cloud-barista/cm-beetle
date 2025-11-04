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

// Package summary provides target infrastructure summary data models
package summary

import (
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
)

// TargetInfraSummary represents the comprehensive target infrastructure summary
// Note: This is a summary-specific struct, not modifying existing tbmodel structs
type TargetInfraSummary struct {
	SummaryMetadata   TargetSummaryMetadata    `json:"summaryMetadata"`
	Overview          TargetInfraOverview      `json:"overview"`
	NetworkResources  SummaryNetworkResources  `json:"networkResources"`
	SecurityResources SummarySecurityResources `json:"securityResources"`
	ComputeResources  SummaryComputeResources  `json:"computeResources"`
	CostEstimation    SummaryCostEstimation    `json:"costEstimation"`
}

// TargetSummaryMetadata contains target infrastructure summary generation metadata
type TargetSummaryMetadata struct {
	GeneratedAt    time.Time `json:"generatedAt" example:"2025-10-31T14:30:00Z"`
	Namespace      string    `json:"namespace" example:"mig01"`
	MciId          string    `json:"mciId" example:"mmci01"`
	MciName        string    `json:"mciName" example:"my-migrated-infrastructure"`
	SummaryVersion string    `json:"summaryVersion" example:"1.0"`
}

// TargetInfraOverview provides high-level overview of the target infrastructure
type TargetInfraOverview struct {
	MciName         string            `json:"mciName" example:"mmci01"`
	MciDescription  string            `json:"mciDescription" example:"Migrated infrastructure"`
	Status          string            `json:"status" example:"Running"`
	TargetCloud     string            `json:"targetCloud" example:"AWS"`
	TargetRegion    string            `json:"targetRegion" example:"ap-northeast-2"`
	TotalVmCount    int               `json:"totalVmCount" example:"3"`
	RunningVmCount  int               `json:"runningVmCount" example:"3"`
	StoppedVmCount  int               `json:"stoppedVmCount" example:"0"`
	Label           map[string]string `json:"label,omitempty"`
	InstallMonAgent string            `json:"installMonAgent" example:"no"`
}

// SummaryNetworkResources contains VNet and Subnet information
type SummaryNetworkResources struct {
	VNets []SummaryVNetInfo `json:"vnets"`
}

// SummaryVNetInfo represents VNet information for summary (Name only, no ID duplication)
type SummaryVNetInfo struct {
	Name           string              `json:"name" example:"mig-vnet-01"`
	CspVNetId      string              `json:"cspVnetId" example:"vpc-06ea213ee81b3e1c4"`
	CidrBlock      string              `json:"cidrBlock" example:"192.168.0.0/16"`
	Region         string              `json:"region" example:"ap-northeast-2"`
	Subnets        []SummarySubnetInfo `json:"subnets"`
	SubnetCount    int                 `json:"subnetCount" example:"2"`
	ConnectionName string              `json:"connectionName" example:"aws-ap-northeast-2"`
}

// SummarySubnetInfo represents Subnet information for summary
type SummarySubnetInfo struct {
	Name        string `json:"name" example:"mig-subnet-01"`
	CspSubnetId string `json:"cspSubnetId" example:"subnet-047dfd6ca50d6791d"`
	CidrBlock   string `json:"cidrBlock" example:"192.168.110.0/24"`
	Zone        string `json:"zone,omitempty" example:"ap-northeast-2a"`
}

// SummarySecurityResources contains SSH Key and Security Group information
type SummarySecurityResources struct {
	SshKeys        []SummarySshKeyInfo        `json:"sshKeys"`
	SecurityGroups []SummarySecurityGroupInfo `json:"securityGroups"`
}

// SummarySshKeyInfo represents SSH Key information for summary
type SummarySshKeyInfo struct {
	Name        string `json:"name" example:"mig-sshkey-01"`
	CspSshKeyId string `json:"cspSshKeyId" example:"d3vkftmqjs728ptbetpg"`
	Username    string `json:"username" example:"cb-user"`
	PublicKey   string `json:"publicKey,omitempty"` // Truncated for security
	Fingerprint string `json:"fingerprint,omitempty" example:"1a:2b:3c:4d:..."`
}

// SummarySecurityGroupInfo represents Security Group information for summary
type SummarySecurityGroupInfo struct {
	Name               string                `json:"name" example:"mig-sg-01"`
	CspSecurityGroupId string                `json:"cspSecurityGroupId" example:"sg-065ead8c271abf7a3"`
	VNetName           string                `json:"vnetName" example:"mig-vnet-01"`
	Rules              []SummaryFirewallRule `json:"rules"`
	RuleCount          int                   `json:"ruleCount" example:"5"`
}

// SummaryFirewallRule represents a firewall rule for summary
type SummaryFirewallRule struct {
	Direction string `json:"direction" example:"inbound"`
	Protocol  string `json:"protocol" example:"tcp"`
	FromPort  string `json:"fromPort" example:"22"`
	ToPort    string `json:"toPort" example:"22"`
	Cidr      string `json:"cidr" example:"0.0.0.0/0"`
}

// SummaryComputeResources contains VM, Spec, and Image information
type SummaryComputeResources struct {
	Specs  []SummarySpecInfoWithUsage  `json:"specs"`
	Images []SummaryImageInfoWithUsage `json:"images"`
	Vms    []SummaryVmInfo             `json:"vms"`
}

// SummarySpecInfoWithUsage extends tbmodel.SpecInfo with usage count
type SummarySpecInfoWithUsage struct {
	tbmodel.SpecInfo     // Embed full SpecInfo from CB-Tumblebug
	UsageCount       int `json:"usageCount" example:"2"` // Number of VMs using this spec
}

// SummaryImageInfoWithUsage extends tbmodel.ImageInfo with usage count
type SummaryImageInfoWithUsage struct {
	tbmodel.ImageInfo     // Embed full ImageInfo from CB-Tumblebug
	UsageCount        int `json:"usageCount" example:"3"` // Number of VMs using this image
}

// SummaryVmInfo represents VM information for summary with restructured format
type SummaryVmInfo struct {
	Name    string             `json:"name" example:"migrated-server-1"`
	CspVmId string             `json:"cspVmId" example:"i-0a1b2c3d4e5f6g7h8"`
	Status  string             `json:"status" example:"Running"`
	Spec    SummaryVmSpecInfo  `json:"spec"`
	Image   SummaryVmImageInfo `json:"image"`
	Misc    SummaryVmMiscInfo  `json:"misc"`
	Region  string             `json:"region" example:"ap-northeast-2"`
	Zone    string             `json:"zone,omitempty" example:"ap-northeast-2a"`
}

// SummaryVmSpecInfo represents VM Spec summary embedded in VM info
type SummaryVmSpecInfo struct {
	Name         string  `json:"name" example:"t3a.xlarge"`
	VCpus        int     `json:"vcpus" example:"4"`
	MemoryGiB    float32 `json:"memoryGiB" example:"16"`
	Architecture string  `json:"architecture,omitempty" example:"x86_64"`
}

// SummaryVmImageInfo represents VM Image summary embedded in VM info
type SummaryVmImageInfo struct {
	Name         string `json:"name" example:"ubuntu22.04"`
	Id           string `json:"id,omitempty" example:"ami-010be25c3775061c9"`
	OsType       string `json:"osType,omitempty" example:"Ubuntu 22.04"`
	Distribution string `json:"distribution" example:"Ubuntu"`
	OsVersion    string `json:"osVersion" example:"22.04"`
}

// SummaryVmMiscInfo contains network and security details for VM
type SummaryVmMiscInfo struct {
	VNet           string   `json:"vnet" example:"mig-vnet-01"`
	Subnet         string   `json:"subnet" example:"mig-subnet-01"`
	PublicIp       string   `json:"publicIp,omitempty" example:"43.201.59.126"`
	PrivateIp      string   `json:"privateIp" example:"192.168.110.10"`
	SecurityGroups []string `json:"securityGroups" example:"mig-sg-01"`
	SshKey         string   `json:"sshKey" example:"mig-sshkey-01"`
	ConnectionName string   `json:"connectionName" example:"aws-ap-northeast-2"`
}

// SummaryCostEstimation provides cost analysis
type SummaryCostEstimation struct {
	Currency          string                `json:"currency" example:"USD"`
	TotalCostPerHour  float32               `json:"totalCostPerHour" example:"0.4512"`
	TotalCostPerDay   float32               `json:"totalCostPerDay" example:"10.83"`
	TotalCostPerMonth float32               `json:"totalCostPerMonth" example:"324.86"`
	ByRegion          []SummaryCostByRegion `json:"byRegion"`
	ByVm              []SummaryCostByVm     `json:"byVm"`
}

// SummaryCostByRegion represents cost breakdown by region
type SummaryCostByRegion struct {
	Csp          string  `json:"csp" example:"AWS"`
	Region       string  `json:"region" example:"ap-northeast-2"`
	VmCount      int     `json:"vmCount" example:"3"`
	CostPerHour  float32 `json:"costPerHour" example:"0.4512"`
	CostPerMonth float32 `json:"costPerMonth" example:"324.86"`
}

// SummaryCostByVm represents cost breakdown by individual VM
type SummaryCostByVm struct {
	VmName       string  `json:"vmName" example:"migrated-server-1"`
	SpecName     string  `json:"specName" example:"t3a.xlarge"`
	CostPerHour  float32 `json:"costPerHour" example:"0.1504"`
	CostPerMonth float32 `json:"costPerMonth" example:"108.29"`
}
