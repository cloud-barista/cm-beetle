package cloudmodel

// RecommendedK8sCluster is the recommendation output for a K8s cluster migration
// using the K8sClusterDynamic API path (TB auto-manages vNet/SG/SSHKey).
type RecommendedK8sCluster struct {
	// Cluster is the request body for POST /ns/{nsId}/k8sClusterDynamic.
	// It includes the cluster configuration and the first worker node group.
	Cluster K8sClusterDynamicReq `json:"cluster"`

	// AdditionalNodeGroups holds extra worker node groups to be added after cluster creation
	// via POST /ns/{nsId}/k8sCluster/{id}/k8sNodeGroupDynamic (required for AWS).
	AdditionalNodeGroups []K8sNodeGroupDynamicReq `json:"additionalNodeGroups,omitempty"`
}
