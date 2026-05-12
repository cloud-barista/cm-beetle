package onpremisemodel

type OnpremiseInfraModel struct {
	OnpremiseInfraModel OnpremInfra `json:"onpremiseInfraModel" validate:"required"`
}

type OnpremInfra struct {
	Network NetworkProperty `json:"network,omitempty"`
	Nodes   []NodeProperty  `json:"nodes" validate:"required"`

	// TODO: Extend to a general ClusterProperty if other orchestrators (e.g., OpenShift, Rancher) are needed.
	K8sCluster *K8sClusterProperty `json:"k8sCluster,omitempty"`
	// TODO: Add other fields
	// Example: FirewallDevice FirewallDeviceProperty `json:"firewallDevice,omitempty"`
}
