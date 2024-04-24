package software

type Software struct {
	DEB    []DEB  `json:"deb"`
	RPM    []RPM  `json:"rpm"`
	Docker Docker `json:"docker"`
	Podman Podman `json:"podman"`
}
