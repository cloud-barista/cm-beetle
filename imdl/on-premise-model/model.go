package onpremisemodel

// SourceNlbModel is the top-level wrapper received from cm-honeybee Agent.
// A single HAProxy instance can serve multiple services simultaneously,
// so SourceNlbModel holds a list — one SourceNlb per frontend-backend pair.
type SourceNlbModel struct {
	SourceNlbModel []SourceNlb `json:"sourceNlbModel" validate:"required"`
}

// SourceNlb is the normalized, software-independent NLB configuration
// extracted from the source environment by cm-honeybee.
type SourceNlb struct {
	Software    string                 `json:"software"`              // "haproxy"
	Listener    NlbListenerProperty    `json:"listener"`
	Backend     NlbBackendProperty     `json:"backend"`
	HealthCheck NlbHealthCheckProperty `json:"healthCheck,omitempty"`
}

type OnpremiseInfraModel struct {
	OnpremiseInfraModel OnpremInfra `json:"onpremiseInfraModel" validate:"required"`
}

type OnpremInfra struct {
	Network NetworkProperty `json:"network,omitempty"`
	Nodes   []NodeProperty  `json:"nodes" validate:"required"`

	// TODO: Extend to a general ClusterProperty if other orchestrators (e.g., OpenShift, Rancher) are needed.
	K8sCluster *K8sClusterProperty `json:"k8sCluster,omitempty"`

	// NLBs holds on-premise NLB instances (HAProxy-based), one entry per frontend-backend pair.
	// Populated by cm-honeybee when HAProxy is detected on a node.
	// Used exclusively by POST /recommendation/infraWithNlb; ignored by POST /recommendation/infra.
	NLBs []NlbProperty `json:"nlbs,omitempty"`
}
