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

// Package report provides migration report data models
package report

import (
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
)

// MigrationReport represents a comprehensive migration analysis report
type MigrationReport struct {
	Metadata          ReportMetadata              `json:"metadata"`
	ExecutiveSummary  ExecutiveSummary            `json:"executiveSummary"`
	MigrationMappings []SourceTargetMapping       `json:"migrationMappings"`
	NetworkAnalysis   NetworkMigrationAnalysis    `json:"networkAnalysis"`
	SecurityAnalysis  SecurityMigrationAnalysis   `json:"securityAnalysis"`
	CostSummary       CostSummary                 `json:"costSummary"`
	Recommendations   []Recommendation            `json:"recommendations"`
	SourceDetails     *summary.SourceInfraSummary `json:"sourceDetails"`
	TargetDetails     *summary.TargetInfraSummary `json:"targetDetails"`
}

// ReportMetadata contains report generation metadata
type ReportMetadata struct {
	GeneratedAt   time.Time `json:"generatedAt" example:"2025-11-04T10:30:00Z"`
	MigrationID   string    `json:"migrationId" example:"mig01/mmci01"`
	Namespace     string    `json:"namespace" example:"mig01"`
	MciID         string    `json:"mciId" example:"mmci01"`
	ReportVersion string    `json:"reportVersion" example:"1.0"`
}

// ExecutiveSummary provides high-level migration overview
type ExecutiveSummary struct {
	MigrationStatus string  `json:"migrationStatus" example:"Completed"`
	TotalServers    int     `json:"totalServers" example:"2"`
	MigratedServers int     `json:"migratedServers" example:"2"`
	FailedServers   int     `json:"failedServers" example:"0"`
	TargetCloud     string  `json:"targetCloud" example:"AWS"`
	TargetRegion    string  `json:"targetRegion" example:"ap-northeast-2"`
	MonthlyCostUSD  float64 `json:"monthlyCostUsd" example:"382.46"`
}

// SourceTargetMapping represents a mapping between source server and target VM
type SourceTargetMapping struct {
	MappingID       int                    `json:"mappingId" example:"1"`
	SourceServer    SourceServerBrief      `json:"sourceServer"`
	TargetVM        TargetVMBrief          `json:"targetVM"`
	ResourceChanges ResourceChangeAnalysis `json:"resourceChanges"`
	MigrationStatus string                 `json:"migrationStatus" example:"Success"`
	CostPerMonth    float64                `json:"costPerMonth" example:"247.68"`
}

// SourceServerBrief contains brief source server information
type SourceServerBrief struct {
	Hostname      string `json:"hostname" example:"cm-nfs"`
	MachineID     string `json:"machineId" example:"0036e4b9-c8b4-e811-906e-000ffee02d5c"`
	CPUModel      string `json:"cpuModel" example:"Intel(R) Xeon(R) CPU E5-2680 v4 @ 2.40GHz"`
	CPUs          int    `json:"cpus" example:"1"`
	CPUThreads    int    `json:"cpuThreads" example:"16"`
	MemoryGB      int    `json:"memoryGb" example:"16"`
	DiskGB        int    `json:"diskGb" example:"1093"`
	DiskType      string `json:"diskType" example:"HDD"`
	OSName        string `json:"osName" example:"Ubuntu 22.04"`
	PrimaryIP     string `json:"primaryIp" example:"172.29.0.102"`
	FirewallRules int    `json:"firewallRules" example:"3"`
}

// TargetVMBrief contains brief target VM information
type TargetVMBrief struct {
	InstanceName    string   `json:"instanceName" example:"migrated-0036e4b9-c8b4-e811-906e-000ffee02d5c-1"`
	InstanceID      string   `json:"instanceId" example:"i-093b4b7c722ed9ec5"`
	SourceMachineID string   `json:"sourceMachineId" example:"0036e4b9-c8b4-e811-906e-000ffee02d5c"`
	Status          string   `json:"status" example:"Running"`
	SpecName        string   `json:"specName" example:"c5a.2xlarge"`
	VCPUs           int      `json:"vcpus" example:"8"`
	MemoryGB        float32  `json:"memoryGb" example:"16.0"`
	RootDiskGB      int      `json:"rootDiskGb" example:"50"`
	RootDiskType    string   `json:"rootDiskType" example:"gp2"`
	PublicIP        string   `json:"publicIp" example:"43.201.85.138"`
	PrivateIP       string   `json:"privateIp" example:"192.168.110.167"`
	SecurityGroups  []string `json:"securityGroups" example:"mig-sg-02"`
}

// ResourceChangeAnalysis contains detailed resource change analysis
type ResourceChangeAnalysis struct {
	CPUChange      ResourceChange `json:"cpuChange"`
	MemoryChange   ResourceChange `json:"memoryChange"`
	StorageChange  ResourceChange `json:"storageChange"`
	NetworkChange  NetworkChange  `json:"networkChange"`
	SecurityChange SecurityChange `json:"securityChange"`
}

// ResourceChange represents a change in a resource
type ResourceChange struct {
	ResourceType string  `json:"resourceType" example:"CPU"`
	SourceValue  string  `json:"sourceValue" example:"2 cores"`
	TargetValue  string  `json:"targetValue" example:"8 vCPU"`
	ChangeType   string  `json:"changeType" example:"Upgrade"` // Upgrade, Downgrade, Same, Added, Removed
	ChangeRatio  float64 `json:"changeRatio" example:"4.0"`    // Multiplier or percentage
	Description  string  `json:"description" example:"+6 cores (4x upgrade)"`
}

// NetworkChange represents network configuration changes
type NetworkChange struct {
	SourceIP        string `json:"sourceIp" example:"172.29.0.102"`
	TargetPrivateIP string `json:"targetPrivateIp" example:"192.168.110.167"`
	TargetPublicIP  string `json:"targetPublicIp" example:"43.201.85.138"`
	ChangeType      string `json:"changeType" example:"Public IP Added"`
	Description     string `json:"description" example:"Added public IP for internet access"`
}

// SecurityChange represents security configuration changes
type SecurityChange struct {
	SourceRules      int    `json:"sourceRules" example:"3"`
	TargetRules      int    `json:"targetRules" example:"6"`
	ConversionStatus string `json:"conversionStatus" example:"Converted"`
	Description      string `json:"description" example:"Converted 3 iptables rules to AWS Security Group"`
}

// NetworkMigrationAnalysis contains network migration analysis
type NetworkMigrationAnalysis struct {
	SourceNetwork SourceNetworkInfo `json:"sourceNetwork"`
	TargetNetwork TargetNetworkInfo `json:"targetNetwork"`
	IPMappings    []IPMapping       `json:"ipMappings"`
	CIDRPreserved bool              `json:"cidrPreserved"`
	Description   string            `json:"description"`
}

// SourceNetworkInfo contains source network information
type SourceNetworkInfo struct {
	CIDR             string `json:"cidr" example:"192.168.110.0/24"`
	Gateway          string `json:"gateway" example:"192.168.110.254"`
	ConnectedServers int    `json:"connectedServers" example:"2"`
}

// TargetNetworkInfo contains target network information
type TargetNetworkInfo struct {
	VNetName   string `json:"vnetName" example:"mig-vnet-01"`
	VNetCIDR   string `json:"vnetCidr" example:"192.168.96.0/19"`
	SubnetCIDR string `json:"subnetCidr" example:"192.168.110.0/24"`
	CSPVNetID  string `json:"cspVnetId" example:"vpc-0200cd398ed7c7f17"`
}

// IPMapping represents source to target IP mapping
type IPMapping struct {
	SourceIP        string `json:"sourceIp" example:"172.29.0.102"`
	SourceHostname  string `json:"sourceHostname" example:"cm-nfs"`
	TargetPrivateIP string `json:"targetPrivateIp" example:"192.168.110.167"`
	TargetPublicIP  string `json:"targetPublicIp" example:"43.201.85.138"`
}

// SecurityMigrationAnalysis contains security migration analysis
type SecurityMigrationAnalysis struct {
	Conversions []SecurityConversion `json:"conversions"`
	Summary     string               `json:"summary"`
}

// SecurityConversion represents a security rule conversion
type SecurityConversion struct {
	SourceHostname string `json:"sourceHostname" example:"cm-nfs"`
	SourceRules    int    `json:"sourceRules" example:"3"`
	TargetSGName   string `json:"targetSgName" example:"mig-sg-02"`
	TargetRules    int    `json:"targetRules" example:"6"`
	ConversionType string `json:"conversionType" example:"iptables to AWS SG"`
	Status         string `json:"status" example:"Converted"`
}

// CostSummary contains cost analysis summary
type CostSummary struct {
	TotalHourlyCost  float64         `json:"totalHourlyCost" example:"0.5312"`
	TotalDailyCost   float64         `json:"totalDailyCost" example:"12.75"`
	TotalMonthlyCost float64         `json:"totalMonthlyCost" example:"382.46"`
	TotalYearlyCost  float64         `json:"totalYearlyCost" example:"4589.52"`
	CostByComponent  []ComponentCost `json:"costByComponent"`
}

// ComponentCost represents cost for a single component
type ComponentCost struct {
	ComponentName  string  `json:"componentName" example:"cm-nfs (migrated)"`
	SpecName       string  `json:"specName" example:"c5a.2xlarge"`
	MonthlyCost    float64 `json:"monthlyCost" example:"247.68"`
	CostPercentage float64 `json:"costPercentage" example:"64.8"`
}

// Recommendation represents a recommendation or action item
type Recommendation struct {
	Category    string   `json:"category" example:"Cost Optimization"`
	Priority    string   `json:"priority" example:"High"` // High, Medium, Low
	Title       string   `json:"title" example:"Consider Reserved Instances"`
	Description string   `json:"description" example:"Reserved Instances can save up to 35% (~$248/month)"`
	ActionItems []string `json:"actionItems,omitempty"`
}
