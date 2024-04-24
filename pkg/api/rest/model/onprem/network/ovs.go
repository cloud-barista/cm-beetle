package network

type OVSBridge struct {
	AutoAttach          any      `json:"auto_attach"`
	Controller          any      `json:"controller"`
	DatapathID          string   `json:"datapath_id"`
	DatapathType        string   `json:"datapath_type"`
	DatapathVersion     string   `json:"datapath_version"`
	ExternalIds         []any    `json:"external_ids"`
	FailMode            any      `json:"fail_mode"`
	FloodVlans          any      `json:"flood_vlans"`
	FlowTables          []any    `json:"flow_tables"`
	Ipfix               any      `json:"ipfix"`
	McastSnoopingEnable bool     `json:"mcast_snooping_enable"`
	Mirrors             any      `json:"mirrors"`
	Name                string   `json:"name"`
	Netflow             any      `json:"netflow"`
	OtherConfig         []any    `json:"other_config"`
	Ports               []string `json:"ports"`
	Protocols           any      `json:"protocols"`
	RstpEnable          bool     `json:"rstp_enable"`
	RstpStatus          []any    `json:"rstp_status"`
	Sflow               any      `json:"sflow"`
	Status              []any    `json:"status"`
	StpEnable           bool     `json:"stp_enable"`
	UUID                string   `json:"uuid"`
	Version             string   `json:"version"`
}
