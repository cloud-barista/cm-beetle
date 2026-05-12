package onpremisemodel

// K8sClusterProperty describes an on-premise Kubernetes cluster
// that may be migrated together with its nodes.
type K8sClusterProperty struct {
	Name    string `json:"name" validate:"required"`    // [Required] Cluster name
	Version string `json:"version" validate:"required"` // [Required] Kubernetes version (e.g., "1.29.3")

	PodCIDR       string `json:"podCIDR,omitempty"`       // [Reference] Pod network CIDR (e.g., "10.244.0.0/16")
	ServiceCIDR   string `json:"serviceCIDR,omitempty"`   // [Reference] Service network CIDR (e.g., "10.96.0.0/12")
	CNIPlugin     string `json:"cniPlugin,omitempty"`     // [Reference] CNI plugin name (e.g., "calico", "flannel", "cilium")
	NodePortRange string `json:"nodePortRange,omitempty"` // [Reference] NodePort range (e.g., "30000-32767")
}
