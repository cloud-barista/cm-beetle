package network

type Host struct {
	NetworkInterface []NIC          `json:"network_interface"`
	DNS              DNS            `json:"dns"`
	Route            []Route        `json:"route"`
	FirewallRule     []FirewallRule `json:"firewall_rule"`
}

type Network struct {
	Host Host `json:"host"`
	CSP  CSP  `json:"csp"`
}
