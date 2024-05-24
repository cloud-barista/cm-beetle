package network

type FirewallRule struct {
	Priority  uint   `json:"priority"` // Lower has higher priority
	Src       string `json:"src"`
	Dst       string `json:"dst"`
	SrcPorts  string `json:"src_ports"`
	DstPorts  string `json:"dst_ports"`
	Protocol  string `json:"protocol"`  // TCP, UDP, ICMP
	Direction string `json:"direction"` // inbound, outbound
	Action    string `json:"action"`    // allow, deny
}
