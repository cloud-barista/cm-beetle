package onpremisemodel

// NodeProperty represents a node in the on-premise infrastructure.
// The node includes a physical server, virtual machine (VM), and Kubernetes node.
type NodeProperty struct {
	Hostname       string                     `json:"hostname"`
	MachineId      string                     `json:"machineId"`                              // Unique identifier for the node (e.g., UUID)
	Role           string                     `json:"role,omitempty" example:"control-plane"` // Node role (e.g., "control-plane", "worker", "standalone")
	CPU            CpuProperty                `json:"cpu"`
	Memory         MemoryProperty             `json:"memory"`
	RootDisk       DiskProperty               `json:"rootDisk"`
	DataDisks      []DiskProperty             `json:"dataDisks,omitempty"`
	Interfaces     []NetworkInterfaceProperty `json:"interfaces"`
	RoutingTable   []RouteProperty            `json:"routingTable"`
	FirewallTable  []FirewallRuleProperty     `json:"firewallTable,omitempty"`
	ListeningPorts []ListeningPortProperty    `json:"listeningPorts,omitempty"` // Listening TCP/UDP sockets on the node
	OS             OsProperty                 `json:"os"`
	GPUCards       []GpuCardProperty          `json:"gpuCards,omitempty"` // Physical GPU cards installed on the node
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

// GpuCardProperty represents a physical GPU card or accelerator module installed in a node.
type GpuCardProperty struct {
	DriverIndex      string  `json:"driverIndex,omitempty" example:"0"`                        // Driver device index (e.g., "0", "card0", "card10")
	Uuid             string  `json:"uuid,omitempty" example:"GPU-12345678-abcd-ef01-2345-..."` // Unique device UUID from driver (e.g., NVML GPU UUID)
	Vendor           string  `json:"vendor,omitempty" example:"NVIDIA"`                        // GPU Vendor (e.g., "NVIDIA", "AMD", "Intel")
	Model            string  `json:"model,omitempty" example:"NVIDIA A100-PCIE-40GB"`          // Specific model for this card/chip
	Type             string  `json:"type,omitempty" example:"GPU"`                             // Accelerator type: "GPU", "NPU", "TPU" (defaults to "GPU")
	Architecture     string  `json:"architecture,omitempty" example:"Ampere"`                  // GPU Microarchitecture (e.g., "Ampere", "Turing", "CDNA 4")
	DriverVersion    string  `json:"driverVersion,omitempty" example:"535.129.03"`             // Installed driver version for this device
	CudaVersion      string  `json:"cudaVersion,omitempty" example:"12.2"`                     // CUDA or compute API version (e.g., "12.2", "ROCm 6.2")
	Slot             string  `json:"slot,omitempty" example:"PCIe Slot 1"`                     // Physical PCIe slot label if available (optional)
	PciBusId         string  `json:"pciBusId,omitempty" example:"0000:01:00.0"`                // PCIe Bus identifier (e.g., "0000:01:00.0")
	ECCEnabled       bool    `json:"eccEnabled,omitempty" example:"true"`                      // Whether Error-Correcting Code (ECC) memory protection is enabled
	MemoryTotalGB    float32 `json:"memoryTotalGB,omitempty" example:"40"`                     // Total physical VRAM capacity in GB (Total + Reserved)
	MemoryReservedGB float32 `json:"memoryReservedGB,omitempty" example:"2"`                   // VRAM reserved for ECC parity and system overhead in GB
	MemoryFreeGB     float32 `json:"memoryFreeGB,omitempty" example:"38"`                      // Available/Free memory in GB
	MemoryUsedGB     float32 `json:"memoryUsedGB,omitempty" example:"2"`                       // Used memory in GB
}

// ListeningPortProperty represents a listening TCP socket or bound UDP socket on a node.
type ListeningPortProperty struct {
	Protocol    string `json:"protocol" example:"tcp"`                  // Transport protocol (e.g., "tcp", "udp")
	BindAddress string `json:"bindAddress,omitempty" example:"0.0.0.0"` // Local bind address (e.g., "0.0.0.0", "::", "127.0.0.1")
	Port        int    `json:"port" example:"8080"`                     // Listening port number
	ProcessName string `json:"processName,omitempty" example:"nginx"`   // Owning process name if available
}
