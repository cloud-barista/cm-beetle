package network

type ActiveAggregatorInfo struct {
	AggregatorID      string `json:"aggregator_id"`
	NumberOfPorts     string `json:"number_of_ports"`
	ActorKey          string `json:"actor_key"`
	PartnerKey        string `json:"partner_key"`
	PartnerMACAddress string `json:"partner_mac_address"`
}

type ADInfo struct {
	LACPActive           string               `json:"lacp_active"`
	LACPRate             string               `json:"lacp_rate"`
	MinLinks             string               `json:"min_links"`
	ADSelect             string               `json:"ad_select"`
	SystemPriority       string               `json:"system_priority"`
	SystemMACAddress     string               `json:"system_mac_address"`
	ActiveAggregatorInfo ActiveAggregatorInfo `json:"active_aggregator_info"`
}

type SlavesList struct {
	Interfaces           []string `json:"interfaces"`
	PrimarySlave         string   `json:"primary_slave"`
	CurrentlyActiveSlave string   `json:"currently_active_slave"`
}

type SlaveInterface struct {
	Name             string `json:"name"`
	MIIStatus        string `json:"mii_status"`
	Speed            string `json:"speed"`
	Duplex           string `json:"duplex"`
	LinkFailureCount string `json:"link_failure_count"`
	PermanentHWAddr  string `json:"permanent_hw_addr"`
	AggregatorID     string `json:"aggregator_id"`
}

type Bonding struct {
	Name               string           `json:"name"`
	BondingMode        string           `json:"bonding_mode"`
	SlavesList         SlavesList       `json:"slaves_list"`
	TransmitHashPolicy string           `json:"transmit_hash_policy"`
	AddrList           []string         `json:"addr_list"`
	MIIStatus          string           `json:"mii_status"`
	MIIPollingInterval string           `json:"mii_polling_interval"`
	ADInfo             ADInfo           `json:"ad_info"`
	SlaveInterface     []SlaveInterface `json:"slave_interface"`
}
