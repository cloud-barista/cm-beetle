package network

type _NetworkSubsystem struct {
	NetworkInterfaces []NIC     `json:"network_interfaces"`
	Routes            []Route   `json:"routes"`
	Netfilter         Netfilter `json:"netfilter"`
	Bonding           []Bonding `json:"bonding"`
}

type VirtualNetwork struct {
	OVS        []OVSBridge `json:"ovs"`
	LibvirtNet LibvirtNet  `json:"libvirt_net"`
}

type Network struct {
	NetworkSubsystem _NetworkSubsystem `json:"network_subsystem"`
	VirtualNetwork   VirtualNetwork    `json:"virtual_network"`
}
