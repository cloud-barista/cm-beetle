package onpremisemodel

// NetworkProerty represents a network for on-premise infrastructure.
// In other perspective, it can be a network for servers and/or a collection of networks extracted from a host.//
// * [Important] Information in the IPv4Networks list should be as non-duplicated as possible.
type NetworkProperty struct { // note: reference command `ip route`, `netstat -rn`, and `lshw -c network`
	IPv4Networks NetworkDetail `json:"ipv4Networks,omitempty"`
	IPv6Networks NetworkDetail `json:"ipv6Networks,omitempty"` // TBD
	// TODO: Add or update fields
}

// NetworkDetail represents a collection of the default route interfaces extracted from each host.
// Note: A network admin/operator "manually" inputs CIDR blocks (e.g., 10.0.0.0/16) of their networks.
// Note: The default gateways are "extracted" from each host and is used to estimate the upper layer address space of the network.
type NetworkDetail struct {
	CidrBlocks      []string          `json:"cidrBlocks,omitempty"`
	DefaultGateways []GatewayProperty `json:"defaultGateways,omitempty"`
}

type GatewayProperty struct {
	IP            string `json:"ip,omitempty"`            // IP address of the gateway
	InterfaceName string `json:"interfaceName,omitempty"` // Name of the network interface associated with the gateway
	MachineId     string `json:"machineId,omitempty"`     // Unique identifier for the machine (e.g., UUID)
}
