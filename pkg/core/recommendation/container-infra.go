package recommendation

import (
	"fmt"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/rs/zerolog/log"
)

// Honeybee v0.4.0 Kubernetes models
type KubernetesInfoList struct {
	Servers []Kubernetes `json:"servers" validate:"required"`
}

type Kubernetes struct {
	NodeCount NodeCount              `json:"node_count"`
	Nodes     []Node                 `json:"nodes"`
	Workloads map[string]interface{} `json:"workloads"`
}

type NodeCount struct {
	Total        int `json:"total"`
	ControlPlane int `json:"control_plane"`
	Worker       int `json:"worker"`
}

type NodeType string

const (
	NodeTypeControlPlane NodeType = "control-plane"
	NodeTypeWorker       NodeType = "worker"
)

type Node struct {
	Type      NodeType    `json:"type"`
	Name      interface{} `json:"name,omitempty"`
	Labels    interface{} `json:"labels,omitempty"`
	Addresses interface{} `json:"addresses,omitempty"`
	NodeSpec  NodeSpec    `json:"node_spec,omitempty"`
	NodeInfo  interface{} `json:"node_info,omitempty"`
}

type NodeSpec struct {
	CPU              int `json:"cpu"`               // cores
	Memory           int `json:"memory"`            // MiB
	EphemeralStorage int `json:"ephemeral_storage"` // MiB
}

type Helm struct {
	Repo    []Repo    `json:"repo"`
	Release []Release `json:"release"`
}

type Repo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Release struct {
	Name             string    `json:"name"`
	Namespace        string    `json:"namespace"`
	Revision         int       `json:"revision"`
	Updated          time.Time `json:"updated"`
	Status           string    `json:"status"`
	ChartNameVersion string    `json:"chart"`
	AppVersion       string    `json:"app_version"`
}

// Beetle defines structs

// NodeGroupInfo holds aggregated information for a node group
type NodeGroupInfo struct {
	Count       int
	TotalCPU    int
	TotalMemory int // in MiB
}

// RecommendK8sCluster recommends K8s control plane configuration based on honeybee source cluster data
func RecommendK8sCluster(provider, region string, k8sInfoList KubernetesInfoList) (tbmodel.K8sClusterDynamicReq, error) {
	var result tbmodel.K8sClusterDynamicReq

	// TODO: Use the following code when implementing actual API calls
	// client := resty.New()
	// client.SetBaseURL(config.Tumblebug.RestUrl)
	// client.SetBasicAuth(config.Tumblebug.API.Username, config.Tumblebug.API.Password)

	// Collect control plane node information
	// cpInfo := &NodeGroupInfo{Count: 0, TotalCPU: 0, TotalMemory: 0}

	log.Info().Int("totalServers", len(k8sInfoList.Servers)).Msg("Processing K8s servers for control plane recommendation")

	// Count kubernetes clusters to check if kubernetes cluster exists
	k8sCount := 0
	for i, k8s := range k8sInfoList.Servers {
		if k8s.NodeCount.Total == 0 {
			log.Warn().
				Int("serverIndex", i).
				Msg("Server has zero total nodes, skipping")
			continue
		}
		k8sCount += 1
	}

	// Validate kubernetes cluster existence
	if k8sCount == 0 {
		log.Warn().Msg("No kubernetes clusters found in the source cluster, skipping recommendation")
		return result, fmt.Errorf("no kubernetes clusters found in the source cluster")
	}

	// Note: for the time being, use the 1st control plane in the list
	k8sClusterInfo := Kubernetes{}
	for _, k8s := range k8sInfoList.Servers {
		if k8s.NodeCount.Total == 0 {
			continue
		}

		k8sClusterInfo = k8s
		break
	}

	log.Debug().Msgf("Selected kubernetes cluster information: %+v", k8sClusterInfo)

	// TODO: Implement spec recommendation logic based on actual resource requirements
	// For now, using hardcoded spec for testing

	k8sControlPlaneNodeCount := k8sClusterInfo.NodeCount.ControlPlane

	// Build K8sClusterDynamicReq
	recommendedControlPlane := tbmodel.K8sClusterDynamicReq{
		Name:            "recommended-k8s-cluster",
		NodeGroupName:   "recommended-worker-nodegroup",
		ConnectionName:  fmt.Sprintf("%s-%s", provider, region),
		SpecId:          "aws+ap-northeast-2+t3a.large", // TODO: Search and assign a recommended spec ID
		ImageId:         "default",                      // TODO: Search and assign a recommended image ID
		Version:         "1.29",
		DesiredNodeSize: fmt.Sprintf("%d", k8sControlPlaneNodeCount),
		MinNodeSize:     "1",
		MaxNodeSize:     fmt.Sprintf("%d", k8sControlPlaneNodeCount),
		OnAutoScaling:   "false",
		RootDiskType:    "default",
		RootDiskSize:    "default",
		Description:     fmt.Sprintf("A recommended control plane with %d nodes", k8sControlPlaneNodeCount),
	}

	log.Info().
		Str("clusterName", recommendedControlPlane.Name).
		Str("specId", recommendedControlPlane.SpecId).
		Str("desiredNodeSize", recommendedControlPlane.DesiredNodeSize).
		Msg("K8s control plane recommendation completed")

	return recommendedControlPlane, nil
}

// RecommendK8sNodeGroup recommends K8s worker node group configuration based on honeybee source cluster data
func RecommendK8sNodeGroup(provider, region string, k8sInfoList KubernetesInfoList) (tbmodel.K8sNodeGroupReq, error) {
	var emptyRes = tbmodel.K8sNodeGroupReq{}

	// TODO: Use the following code when implementing actual API calls
	// client := resty.New()
	// client.SetBaseURL(config.Tumblebug.RestUrl)
	// client.SetBasicAuth(config.Tumblebug.API.Username, config.Tumblebug.API.Password)

	// Collect control plane node information
	// cpInfo := &NodeGroupInfo{Count: 0, TotalCPU: 0, TotalMemory: 0}

	log.Info().Int("totalServers", len(k8sInfoList.Servers)).Msg("Processing K8s servers for control plane recommendation")

	// Count kubernetes clusters to check if kubernetes cluster exists
	k8sCount := 0
	for i, k8s := range k8sInfoList.Servers {
		if k8s.NodeCount.Total == 0 {
			log.Warn().
				Int("serverIndex", i).
				Msg("Server has zero total nodes, skipping")
			continue
		}
		k8sCount += 1
	}

	// Validate kubernetes cluster existence
	if k8sCount == 0 {
		log.Warn().Msg("No kubernetes clusters found in the source cluster, skipping recommendation")
		return emptyRes, fmt.Errorf("no kubernetes clusters found in the source cluster")
	}

	// Note: for the time being, use the 1st kubernetes cluster in the list
	k8sClusterInfo := Kubernetes{}
	for _, k8s := range k8sInfoList.Servers {
		if k8s.NodeCount.Total == 0 {
			continue
		}

		k8sClusterInfo = k8s
		break
	}

	log.Debug().Msgf("Selected kubernetes cluster information: %+v", k8sClusterInfo)

	// TODO: Implement spec recommendation logic based on actual resource requirements
	// For now, using hardcoded spec for testing

	k8sNodeGroupCount := k8sClusterInfo.NodeCount.Worker

	// TODO: Implement spec recommendation logic based on actual resource requirements
	// For now, using hardcoded spec for testing

	// Build K8sNodeGroupReq
	recommendedNodeGroup := tbmodel.K8sNodeGroupReq{
		Name:            "recommended-worker-nodegroup",
		SpecId:          "aws+ap-northeast-2+t3a.xlarge", // Hardcoded for testing
		ImageId:         "default",
		DesiredNodeSize: fmt.Sprintf("%d", k8sNodeGroupCount),
		MinNodeSize:     "1",
		MaxNodeSize:     fmt.Sprintf("%d", k8sNodeGroupCount),
		OnAutoScaling:   "true",
		RootDiskType:    "default",
		RootDiskSize:    "default",
		Description:     fmt.Sprintf("A recommended node group with %d nodes", k8sNodeGroupCount),
	}

	log.Info().
		Str("nodeGroupName", recommendedNodeGroup.Name).
		Str("specId", recommendedNodeGroup.SpecId).
		Str("desiredNodeSize", recommendedNodeGroup.DesiredNodeSize).
		Msg("K8s worker node group recommendation completed")

	return recommendedNodeGroup, nil
}
