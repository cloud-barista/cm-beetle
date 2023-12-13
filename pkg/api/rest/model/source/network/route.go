package network

type Route struct {
	Destination string `json:"destination"`
	Netmask     string `json:"netmask"`
	NextHop     string `json:"next_hop"`
}
