package infra

// MountedInformation TODO
type MountedInformation struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Filesystem  string `json:"filesystem"`
	Option      string `json:"option"`
}

// MountPoint TODO
type MountPoint struct {
	MountedInformation []MountedInformation `json:"mounted_information"`
}

// Storage TODO
type Storage struct {
	MountPoint MountPoint `json:"mount_point"`
}
