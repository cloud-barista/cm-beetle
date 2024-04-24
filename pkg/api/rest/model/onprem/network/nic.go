package network

type NIC struct {
	Interface string   `json:"interface"`
	Address   []string `json:"address"`
	Gateway   []string `json:"gateway"`
	Route     []Route  `json:"route"`
	MAC       string   `json:"mac"`
	MTU       int      `json:"mtu"`
}
