package software

type RPM struct {
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	Release   string   `json:"release"`
	Arch      string   `json:"arch"`
	SourceRpm string   `json:"sourceRpm"`
	Size      int      `json:"size"`
	License   string   `json:"license"`
	Vendor    string   `json:"vendor"`
	Summary   string   `json:"summary"`
	Requires  []string `json:"requires"`
}
