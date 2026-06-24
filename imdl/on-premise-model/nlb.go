package onpremisemodel

// ============================================================================
// On-premise NLB property types
// Identical JSON tags to nlb-model.NlbProperty — no cross-import needed.
// ============================================================================

// NlbProperty represents a single NLB instance on the on-premise environment
// (one entry per HAProxy frontend-backend pair).
type NlbProperty struct {
	HostMachineId string                 `json:"hostMachineId,omitempty"` // MachineId of the node running HAProxy
	Software      string                 `json:"software"`                // "haproxy"
	Listener      NlbListenerProperty    `json:"listener"`
	Backend       NlbBackendProperty     `json:"backend"`
	HealthCheck   NlbHealthCheckProperty `json:"healthCheck,omitempty"`
}

// NlbListenerProperty captures the frontend listener of the source NLB.
type NlbListenerProperty struct {
	BindAddress string `json:"bindAddress"` // "*" = all interfaces (→ PUBLIC), specific IP = INTERNAL
	Port        int    `json:"port"`        // Listener port (1–65535)
	Protocol    string `json:"protocol"`    // "tcp" | "udp"
}

// NlbBackendProperty captures the backend configuration of the source NLB.
type NlbBackendProperty struct {
	Name     string              `json:"name"`     // Backend section name
	Balance  string              `json:"balance"`  // "roundrobin" | "leastconn" | "source" (note only)
	Protocol string              `json:"protocol"` // "tcp" | "http"
	Servers  []NlbServerProperty `json:"servers"`
}

// NlbServerProperty captures a single backend server in the source NLB.
type NlbServerProperty struct {
	Name   string `json:"name"`
	IP     string `json:"ip"`               // Server IP; used for IP correlation at recommendation time
	Port   int    `json:"port"`             // Server port
	Weight int    `json:"weight,omitempty"` // Traffic weight (reference only)
}

// NlbHealthCheckProperty captures health check settings of the source NLB.
type NlbHealthCheckProperty struct {
	Enabled   bool   `json:"enabled"`
	Protocol  string `json:"protocol,omitempty"`  // "tcp" | "http"
	Port      int    `json:"port,omitempty"`      // 0 = same as server port
	Interval  int    `json:"interval,omitempty"`  // seconds; default 10
	Timeout   int    `json:"timeout,omitempty"`   // seconds; default 10
	Threshold int    `json:"threshold,omitempty"` // default 3
}
