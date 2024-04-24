package infra

type OS struct {
	Name         string `json:"name"`
	Vendor       string `json:"vendor"`
	Version      string `json:"version"`
	Release      string `json:"release"`
	Architecture string `json:"architecture"`
}

type Kernel struct {
	Release      string `json:"release"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
}

type Node struct {
	Hostname   string `json:"hostname"`
	Hypervisor string `json:"hypervisor"`
	Machineid  string `json:"machineid"`
	Timezone   string `json:"timezone"`
}

type System struct {
	OS     OS     `json:"os"`
	Kernel Kernel `json:"kernel"`
	Node   Node   `json:"node"`
}

type CPU struct {
	Vendor  string `json:"vendor"`
	Model   string `json:"model"`
	Speed   uint   `json:"speed"`   // MHz
	Cache   uint   `json:"cache"`   // KB
	Cpus    uint   `json:"cpus"`    // ea
	Cores   uint   `json:"cores"`   // ea
	Threads uint   `json:"threads"` // ea
}

type Memory struct {
	Type  string `json:"type"`
	Speed uint   `json:"speed"` // MHz
	Size  uint   `json:"size"`  // MB
}

type Storage struct {
	Name   string `json:"name"`
	Driver string `json:"driver"`
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
	Serial string `json:"serial"`
	Size   uint   `json:"size"` // GB
}

type ComputeResource struct {
	CPU     CPU       `json:"cpu"`
	Memory  Memory    `json:"memory"`
	Storage []Storage `json:"storage"`
}

type Compute struct {
	OS              System          `json:"os"`
	ComputeResource ComputeResource `json:"compute_resource"`
}
