package network

type NIC struct {
	Interface  string   `json:"interface"`
	Address    []string `json:"address"`
	Gateway    []string `json:"gateway"`
	MACAddress string   `json:"mac_address"`
	MTU        int      `json:"mtu"`
}
