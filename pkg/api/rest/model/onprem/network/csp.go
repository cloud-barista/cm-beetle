package network

// Subnet TODO
type Subnet struct {
	Name     string `json:"name"`
	IPv4CIDR string `json:"ipv4_cidr"` // IPv4 Network Address with CIDR Prefix Length
	IPv6CIDR string `json:"ipv6_cidr"` // IPv6 Network Address with CIDR Prefix Length
}

// VPC TODO
type VPC struct {
	ID           string   `json:"id"`
	Region       string   `json:"region"`
	AddressSpace []string `json:"address_space"` // IPv4 CIDR or IPv6 CIDR
	Subnet       []Subnet `json:"subnet"`
	DNSServer    []DNS    `json:"dns_server"`
}

// NLB TODO
type NLB struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Listener      string `json:"listener"`
	TargetGroup   string `json:"target_group"`
	HealthChecker string `json:"health_checker"`
}

// SecurityGroup TODO
type SecurityGroup struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	VnetID       string         `json:"vnet_id"`
	FirewallRule []FirewallRule `json:"firewall_rule"`
}

// CSP TODO
type CSP struct {
	Name          string          `json:"name"`
	VPC           []VPC           `json:"vpc"`
	NLB           []NLB           `json:"nlb"`
	SecurityGroup []SecurityGroup `json:"security_group"`
}
