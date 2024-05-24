package software

type DEB struct {
	Package       string   `json:"package"`
	Status        string   `json:"status"`
	Priority      string   `json:"priority"`
	Architecture  string   `json:"architecture"`
	MultiArch     string   `json:"multi_arch"`
	Maintainer    string   `json:"maintainer"`
	Version       string   `json:"version"`
	Section       string   `json:"section"`
	InstalledSize int64    `json:"installed_size"`
	Depends       string   `json:"depends"`
	Conffiles     []string `json:"conffiles"`
	PreDepends    string   `json:"pre_depends"`
	Description   string   `json:"description"`
	Source        string   `json:"source"`
	Homepage      string   `json:"homepage"`
}
