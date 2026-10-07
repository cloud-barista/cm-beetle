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

// Package summary provides target K8s infrastructure summary data models
package summary

import (
	"time"
)

// TargetK8sInfraSummary represents the comprehensive summary of a provisioned K8s infrastructure.
type TargetK8sInfraSummary struct {
	SummaryMetadata   TargetK8sSummaryMetadata `json:"summaryMetadata"`
	Overview          TargetK8sInfraOverview   `json:"overview"`
	NetworkResources  SummaryNetworkResources  `json:"networkResources"`
	SecurityResources SummarySecurityResources `json:"securityResources"`
	NodeGroupList     []SummaryNodeGroupInfo   `json:"nodeGroupList"`
	ComputeOverview   K8sComputeOverview       `json:"computeOverview"`
	CostEstimation    SummaryCostEstimation    `json:"costEstimation"`
}

// TargetK8sSummaryMetadata contains metadata for target K8s infrastructure summary generation.
type TargetK8sSummaryMetadata struct {
	GeneratedAt    time.Time `json:"generatedAt" example:"2026-10-07T15:00:00Z"`
	Namespace      string    `json:"namespace" example:"mig01"`
	ClusterId      string    `json:"clusterId" example:"mig-k8s-cluster"`
	ClusterName    string    `json:"clusterName" example:"mig-k8s-cluster"`
	SummaryVersion string    `json:"summaryVersion" example:"1.0"`
}

// TargetK8sInfraOverview provides a high-level overview of the target K8s infrastructure.
type TargetK8sInfraOverview struct {
	ClusterName          string `json:"clusterName" example:"mig-k8s-cluster"`
	Status               string `json:"status" example:"Active"`
	TargetCloud          string `json:"targetCloud" example:"AWS"`
	TargetRegion         string `json:"targetRegion" example:"ap-northeast-2"`
	ConnectionName       string `json:"connectionName" example:"aws-ap-northeast-2"`
	K8sVersion           string `json:"k8sVersion" example:"1.34"`
	CspClusterId         string `json:"cspClusterId,omitempty" example:"arn:aws:eks:..."`
	TotalNodeGroupCount  int    `json:"totalNodeGroupCount" example:"2"`
	TotalWorkerNodeCount int    `json:"totalWorkerNodeCount" example:"4"`
	ActiveNodeCount      int    `json:"activeNodeCount" example:"4"`
}

// SummaryNodeGroupInfo represents detailed information about a single K8s node group.
type SummaryNodeGroupInfo struct {
	Name            string               `json:"name" example:"workers1"`
	Status          string               `json:"status" example:"Active"`
	SpecId          string               `json:"specId" example:"c5a.xlarge"`
	Vcpu            int                  `json:"vcpu" example:"4"`
	MemoryGib       float64              `json:"memoryGib" example:"8.0"`
	RootDiskType    string               `json:"rootDiskType" example:"default"`
	RootDiskSizeGB  int                  `json:"rootDiskSizeGb" example:"50"`
	OnAutoScaling   string               `json:"onAutoScaling" example:"false"`
	DesiredNodeSize int                  `json:"desiredNodeSize" example:"2"`
	MinNodeSize     int                  `json:"minNodeSize" example:"0"`
	MaxNodeSize     int                  `json:"maxNodeSize" example:"0"`
	Nodes           []SummaryK8sNodeInfo `json:"nodes,omitempty"`
	HourlyCostUSD   float64              `json:"hourlyCostUsd" example:"0.308"`
	MonthlyCostUSD  float64              `json:"monthlyCostUsd" example:"221.76"`
}

// SummaryK8sNodeInfo represents an individual worker node inside a node group.
type SummaryK8sNodeInfo struct {
	CspResourceName string `json:"cspResourceName,omitempty" example:"i-0abcd1234ef567890"`
	CspResourceId   string `json:"cspResourceId,omitempty" example:"csp-06eb41e14121c550a"`
}

// K8sComputeOverview aggregates compute resources across all node groups.
type K8sComputeOverview struct {
	TotalVcpus     int     `json:"totalVcpus" example:"16"`
	TotalMemoryGib float64 `json:"totalMemoryGib" example:"32.0"`
	TotalStorageGB int     `json:"totalStorageGb" example:"200"`
}
