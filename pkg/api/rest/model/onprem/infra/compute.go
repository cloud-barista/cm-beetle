package infra

type OS struct {
	Name         string `json:"name"`
	Vendor       string `json:"vendor" validate:"required"`
	Version      string `json:"version"`
	Release      string `json:"release" validate:"required"`
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
	OS     OS     `json:"os" validate:"required"`
	Kernel Kernel `json:"kernel"`
	Node   Node   `json:"node"`
}

type CPU struct {
	Vendor   string `json:"vendor"`
	Model    string `json:"model"`
	MaxSpeed uint   `json:"max_speed"`                   // MHz
	Cache    uint   `json:"cache"`                       // KB
	Cpus     uint   `json:"cpus" validate:"required"`    // ea
	Cores    uint   `json:"cores" validate:"required"`   // ea
	Threads  uint   `json:"threads" validate:"required"` // ea
}

type Memory struct {
	Type  string `json:"type"`
	Speed uint   `json:"speed"`                    // MHz
	Size  uint   `json:"size" validate:"required"` // MB
}

type Disk struct {
	Label string `json:"label"`
	Type  string `json:"type"`
	Size  uint   `json:"size"` // GB
}

type ComputeResource struct {
	CPU      CPU    `json:"cpu" validate:"required"`
	Memory   Memory `json:"memory" validate:"required"`
	RootDisk Disk   `json:"root_disk"`
	DataDisk []Disk `json:"data_disk"`
}

// Keypair TODO
type Keypair struct {
	Name       string `json:"name"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// Connection TODO
type Connection struct {
	Keypair Keypair `json:"keypair"`
}

type Compute struct {
	OS              System          `json:"os" validate:"required"`
	ComputeResource ComputeResource `json:"compute_resource" validate:"required"`
	Connection      []Connection    `json:"connection"`
}
