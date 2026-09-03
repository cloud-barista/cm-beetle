package onpremisemodel

// NodeProperty represents a node in the on-premise infrastructure.
// The node includes a physical server, virtual machine (VM), and Kubernetes node.
type NodeProperty struct {
	Hostname      string                     `json:"hostname"`
	MachineId     string                     `json:"machineId"`                              // Unique identifier for the node (e.g., UUID)
	Role          string                     `json:"role,omitempty" example:"control-plane"` // Node role (e.g., "control-plane", "worker", "standalone")
	CPU           CpuProperty                `json:"cpu"`
	Memory        MemoryProperty             `json:"memory"`
	RootDisk      DiskProperty               `json:"rootDisk"`
	DataDisks     []DiskProperty             `json:"dataDisks,omitempty"`
	Interfaces    []NetworkInterfaceProperty `json:"interfaces"`
	RoutingTable  []RouteProperty            `json:"routingTable"`
	FirewallTable []FirewallRuleProperty     `json:"firewallTable,omitempty"`
	OS            OsProperty                 `json:"os"`
	GPU           *GpuProperty               `json:"gpu,omitempty"` // GPU accelerator hardware information (optional)
}

type CpuProperty struct {
	Architecture string  `json:"architecture" example:"x86_64"`
	Cpus         uint32  `json:"cpus" validate:"required" example:"2"`     // Number of physical CPUs (sockets)
	Cores        uint32  `json:"cores" validate:"required" example:"18"`   // Number of physical cores per CPU
	Threads      uint32  `json:"threads" validate:"required" example:"36"` // Number of logical CPUs (threads) per CPU with hyper-threading enabled
	MaxSpeed     float32 `json:"maxSpeed,omitempty" example:"3.6"`         // Maximum speed in GHz
	Vendor       string  `json:"vendor,omitempty" example:"GenuineIntel"`
	Model        string  `json:"model,omitempty" example:"Intel(R) Xeon(R) Gold 6140 CPU @ 2.30GHz"`
}

type MemoryProperty struct {
	Type      string `json:"type" validate:"required" example:"DDR4"`
	TotalSize uint64 `json:"totalSize" validate:"required" example:"128"` // Unit GiB
	Available uint64 `json:"available,omitempty"`                         // Unit GiB
	Used      uint64 `json:"used,omitempty"`                              // Unit GiB
	// TODO: Add or update fields
}

type DiskProperty struct { // note: reference command `df -h`
	Label     string `json:"label" validate:"required"`
	Type      string `json:"type" validate:"required" example:"SSD"`       // SSD, HDD
	TotalSize uint64 `json:"totalSize" validate:"required" example:"1024"` // Unit GB
	Available uint64 `json:"available,omitempty"`                          // Unit GB
	Used      uint64 `json:"used,omitempty"`                               // Unit GB
}

type NetworkInterfaceProperty struct { // note: reference command `ifconfig`
	Name           string   `json:"name,omitempty" validate:"required"` // Interface name (e.g., eth0, ens01, enp0s3)
	MacAddress     string   `json:"macAddress,omitempty"`               // MAC address
	IPv4CidrBlocks []string `json:"ipv4CidrBlocks,omitempty"`           // IPv4 address with prefix length (e.g., 192.168.0.21/24), instead of inet addr, Bcast, and Mask
	IPv6CidrBlocks []string `json:"ipv6CidrBlocks,omitempty"`           // IPv6 address with prefix length (e.g., "2001:db8::1/64")
	Mtu            int      `json:"mtu,omitempty"`                      // Maximum Transmission Unit (MTU) in bytes
	State          string   `json:"state,omitempty"`                    // Interface state (e.g., UP, DOWN)
	// TODO: Add or update fields (e.g., )
}

type RouteProperty struct { // note: reference command `ip route`
	Destination string `json:"destination,omitempty"` // Destination network, expressed in CIDR format
	Gateway     string `json:"gateway,omitempty"`     // Gateway address to which packets are forwarded
	Interface   string `json:"interface,omitempty"`   // Network interface associated with the route
	Metric      int    `json:"metric,omitempty"`      // Metric value indicating the priority of the route
	Protocol    string `json:"protocol,omitempty"`    // Protocol used to set the route (e.g., kernel, static)
	Scope       string `json:"scope,omitempty"`       // Scope of the route (e.g., global, link, host)
	Source      string `json:"source,omitempty"`      // Optionally stores the source address (used for policy-based routing)
	LinkState   string `json:"linkState,omitempty"`   // Link state of the route (e.g., UP, DOWN)
}

type FirewallRuleProperty struct { // note: reference command `iptables -v -t filter -S`
	SrcCIDR   string `json:"srcCIDR,omitempty"`   // e.g., "123.123.123.123/32", "123.123.123.123/24", "0.0.0.0/0"
	SrcPorts  string `json:"srcPorts,omitempty"`  // e.g., "80", "80,443", "1024-65535", "*" (for all ports)
	DstCIDR   string `json:"dstCIDR,omitempty"`   // e.g., "123.123.123.123/32", "123.123.123.123/24", "0.0.0.0/0"
	DstPorts  string `json:"dstPorts,omitempty"`  // e.g., "80", "80,443", "1024-65535", "*" (for all ports)
	Protocol  string `json:"protocol,omitempty"`  // e.g., "TCP", "UDP", "ICMP", "*" (for all protocol)
	Direction string `json:"direction,omitempty"` // e.g., inbound, outbound
	Action    string `json:"action,omitempty"`    // e.g., allow, deny
}

type OsProperty struct { // note: reference command `cat /etc/os-release`
	PrettyName      string `json:"prettyName" validate:"required" example:"Ubuntu 22.04.3 LTS"` // Pretty name
	Version         string `json:"version,omitempty" example:"22.04.3 LTS (Jammy Jellyfish)"`   // Full version string
	Name            string `json:"name,omitempty" example:"Ubuntu"`
	VersionID       string `json:"versionId,omitempty" example:"22.04"`
	VersionCodename string `json:"versionCodename,omitempty" example:"jammy"`
	ID              string `json:"id,omitempty" example:"ubuntu"`
	IDLike          string `json:"idLike,omitempty" example:"debian"`
}

// GpuProperty represents GPU and accelerator hardware information of an on-premise node.
type GpuProperty struct {
	Count         uint32      `json:"count" validate:"required" example:"1"`           // Number of physical GPU devices/chips
	Vendor        string      `json:"vendor,omitempty" example:"NVIDIA"`               // GPU Vendor/Manufacturer (e.g., "NVIDIA", "AMD", "Intel")
	Model         string      `json:"model,omitempty" example:"NVIDIA A100-PCIE-40GB"` // Primary GPU model name (e.g., "Tesla T4", "NVIDIA A100-PCIE-40GB", "GeForce RTX 4090")
	Type          string      `json:"type,omitempty" example:"GPU"`                    // Accelerator type: "GPU", "NPU", "TPU" (defaults to "GPU")
	TotalMemoryGB float32     `json:"totalMemoryGB,omitempty" example:"40"`            // Total VRAM across all devices in GB
	DriverVersion string      `json:"driverVersion,omitempty" example:"535.129.03"`    // Installed GPU driver version
	CudaVersion   string      `json:"cudaVersion,omitempty" example:"12.2"`            // Supported/Installed CUDA version (e.g., "12.2", "12.4")
	Architecture  string      `json:"architecture,omitempty" example:"Ampere"`         // GPU Microarchitecture (e.g., "Ampere", "Hopper", "Ada Lovelace", "Turing", "Volta")
	Details       []GpuDetail `json:"details,omitempty"`                               // Detailed information per individual physical GPU device
}

// GpuDetail represents detailed hardware attributes of an individual physical GPU device.
type GpuDetail struct {
	Index       uint32  `json:"index" example:"0"`                                         // Device index (e.g., 0, 1)
	Uuid        string  `json:"uuid,omitempty" example:"GPU-12345678-abcd-ef01-2345-..."`  // Unique device UUID from driver (e.g., NVML GPU UUID)
	Model       string  `json:"model,omitempty" example:"NVIDIA A100-PCIE-40GB"`           // Specific model for this device
	PciBusId    string  `json:"pciBusId,omitempty" example:"0000:01:00.0"`                 // PCIe Bus identifier (e.g., "0000:01:00.0")
	MemoryTotal float32 `json:"memoryTotal,omitempty" example:"40"`                        // Memory capacity in GB
	MemoryFree  float32 `json:"memoryFree,omitempty" example:"38"`                         // Available/Free memory in GB
	MemoryUsed  float32 `json:"memoryUsed,omitempty" example:"2"`                          // Used memory in GB
}

