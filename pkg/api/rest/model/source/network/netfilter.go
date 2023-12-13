package network

type Chain struct {
	ChainName string   `json:"chain_name"`
	Rules     []string `json:"rules"`
}

type Table struct {
	TableName string  `json:"table_name"`
	Chains    []Chain `json:"chains"`
}

type Netfilter struct {
	IPv4Tables []Table `json:"ipv4_tables"`
	IPv6Tables []Table `json:"ipv6_tables"`
}
