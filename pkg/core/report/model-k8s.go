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
	"github.com/cloud-barista/cm-beetle/pkg/core/summary"
)

// K8sMigrationReport represents a source-to-target comparative migration analysis report for K8s.
type K8sMigrationReport struct {
	Metadata           ReportMetadata                  `json:"metadata"`
	ExecutiveSummary   K8sExecutiveSummary             `json:"executiveSummary"`
	VersionAnalysis    K8sVersionAnalysis              `json:"versionAnalysis"`
	NodeGroupMappings  []WorkerNodeGroupMapping        `json:"nodeGroupMappings"`
	SizingAnalysis     []WorkerSizingItem              `json:"sizingAnalysis"`
	ExcludedWorkers    []ExcludedWorkerItem            `json:"excludedWorkers,omitempty"`
	ResourceComparison K8sResourceComparison           `json:"resourceComparison"`
	CostSummary        CostSummary                     `json:"costSummary"`
	Recommendations    []Recommendation                `json:"recommendations"`
	SourceDetails      *summary.SourceInfraSummary     `json:"sourceDetails"`
	TargetDetails      *summary.TargetK8sInfraSummary  `json:"targetDetails"`
}

// K8sExecutiveSummary provides an executive summary for K8s migration.
type K8sExecutiveSummary struct {
	MigrationStatus      string  `json:"migrationStatus" example:"Completed"`
	SourceClusterName    string  `json:"sourceClusterName" example:"on-prem-k8s"`
	SourceTotalNodes     int     `json:"sourceTotalNodes" example:"3"`
	SourceWorkerNodes    int     `json:"sourceWorkerNodes" example:"2"`
	TargetCloud          string  `json:"targetCloud" example:"AWS"`
	TargetRegion         string  `json:"targetRegion" example:"ap-northeast-2"`
	TargetClusterName    string  `json:"targetClusterName" example:"mig-k8s-cluster"`
	TargetNodeGroupCount int     `json:"targetNodeGroupCount" example:"1"`
	TargetTotalNodes     int     `json:"targetTotalNodes" example:"2"`
	MonthlyCostUSD       float64 `json:"monthlyCostUsd" example:"221.76"`
}

// K8sVersionAnalysis details version changes between source and target clusters.
type K8sVersionAnalysis struct {
	SourceVersion string `json:"sourceVersion" example:"1.28.0"`
	TargetVersion string `json:"targetVersion" example:"1.34"`
	UpgradeStatus string `json:"upgradeStatus" example:"Upgraded (Compatible)"`
	Notes         string `json:"notes,omitempty" example:"Upgraded to target CSP supported minimum stable version"`
}

// WorkerNodeGroupMapping maps on-premise worker nodes to target cloud node groups.
type WorkerNodeGroupMapping struct {
	NodeGroupName   string              `json:"nodeGroupName" example:"workers1"`
	TargetSpecId    string              `json:"targetSpecId" example:"c5a.xlarge"`
	TargetNodeCount int                 `json:"targetNodeCount" example:"2"`
	AssignedWorkers []SourceServerBrief `json:"assignedWorkers"`
}

// WorkerSizingItem explains the sizing rationale for a worker node.
type WorkerSizingItem struct {
	WorkerHostname string  `json:"workerHostname" example:"k8s-worker-01"`
	SourceVcpu     int     `json:"sourceVcpu" example:"1"`
	SourceMemoryGb float64 `json:"sourceMemoryGb" example:"2.0"`
	TargetSpecId   string  `json:"targetSpecId" example:"c5a.xlarge"`
	TargetVcpu     int     `json:"targetVcpu" example:"4"`
	TargetMemoryGb float64 `json:"targetMemoryGb" example:"8.0"`
	SizingReason   string  `json:"sizingReason" example:"Upscaled to minimum viable K8s worker floor (2vCPU / 4GiB)"`
}

// ExcludedWorkerItem records workers excluded from migration and the explicit reason.
type ExcludedWorkerItem struct {
	Hostname string `json:"hostname" example:"k8s-master-01"`
	Role     string `json:"role" example:"control-plane"`
	Reason   string `json:"reason" example:"Control-plane abstracted into CSP managed control plane"`
}

// K8sResourceComparison computes delta between source on-premise and target cloud resources.
type K8sResourceComparison struct {
	WorkerCountDelta ResourceChange `json:"workerCountDelta"`
	VcpuDelta        ResourceChange `json:"vcpuDelta"`
	MemoryGibDelta   ResourceChange `json:"memoryGibDelta"`
	StorageGbDelta   ResourceChange `json:"storageGbDelta"`
}

