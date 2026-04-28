package onpremisemodel

type OnpremiseInfraModel struct {
	OnpremiseInfraModel OnpremInfra `json:"onpremiseInfraModel" validate:"required"`
}

type OnpremInfra struct {
	Network NetworkProperty  `json:"network,omitempty"`
	Servers []ServerProperty `json:"servers" validate:"required"`
	// TODO: Add other fields
	// Example: FirewallDevice FirewallDeviceProperty `json:"firewallDevice,omitempty"`
}
