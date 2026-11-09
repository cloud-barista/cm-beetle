package main

import (
	"fmt"
	"math/rand"
	"strings"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/imdl/lineage"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
)

type rolePreset struct {
	role          string
	subnetOctet   int
	firewallTable []onpremmodel.FirewallRuleProperty
}

type osPreset struct {
	prettyName string
	version    string
	name       string
	versionID  string
	codename   string
	id         string
	idLike     string
}

var defaultRolePresets = []rolePreset{
	{
		role:        "web",
		subnetOctet: 1,
		firewallTable: []onpremmodel.FirewallRuleProperty{
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "22", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "80", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "443", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "outbound", DstCIDR: "0.0.0.0/0", DstPorts: "*", Protocol: "*", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
		},
	},
	{
		role:        "was",
		subnetOctet: 1,
		firewallTable: []onpremmodel.FirewallRuleProperty{
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "22", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "8080", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "8443", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "outbound", DstCIDR: "0.0.0.0/0", DstPorts: "*", Protocol: "*", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
		},
	},
	{
		role:        "db",
		subnetOctet: 2,
		firewallTable: []onpremmodel.FirewallRuleProperty{
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "22", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "3306", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "5432", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "outbound", DstCIDR: "0.0.0.0/0", DstPorts: "*", Protocol: "*", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
		},
	},
	{
		role:        "worker",
		subnetOctet: 2,
		firewallTable: []onpremmodel.FirewallRuleProperty{
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "22", Protocol: "tcp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "inbound", DstCIDR: "0.0.0.0/0", DstPorts: "*", Protocol: "icmp", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
			{Action: "allow", Direction: "outbound", DstCIDR: "0.0.0.0/0", DstPorts: "*", Protocol: "*", SrcCIDR: "0.0.0.0/0", SrcPorts: "*"},
		},
	},
}

var defaultOSPresets = []osPreset{
	{
		prettyName: "Ubuntu 22.04.3 LTS",
		version:    "22.04.3 LTS (Jammy Jellyfish)",
		name:       "Ubuntu",
		versionID:  "22.04",
		codename:   "jammy",
		id:         "ubuntu",
		idLike:     "debian",
	},
	{
		prettyName: "Ubuntu 20.04.6 LTS",
		version:    "20.04.6 LTS (Focal Fossa)",
		name:       "Ubuntu",
		versionID:  "20.04",
		codename:   "focal",
		id:         "ubuntu",
		idLike:     "debian",
	},
}

// GenerateRandomOnpremInfra creates a synthetic OnpremInfra model with nodeCount nodes.
func GenerateRandomOnpremInfra(rng *rand.Rand, nodeCount int, cfg GeneratorConfig) onpremmodel.OnpremInfra {
	nodes := make([]onpremmodel.NodeProperty, 0, nodeCount)
	gateways := make([]onpremmodel.GatewayProperty, 0, nodeCount)

	for i := 0; i < nodeCount; i++ {
		preset := defaultRolePresets[i%len(defaultRolePresets)]
		osInfo := defaultOSPresets[rng.Intn(len(defaultOSPresets))]

		vcpus := randomStepInt(rng, cfg.MinVCPU, cfg.MaxVCPU, 1)
		memGB := randomStepInt(rng, cfg.MinMemoryGB, cfg.MaxMemoryGB, 2)
		diskGB := randomStepInt(rng, cfg.MinDiskGB, cfg.MaxDiskGB, 10)

		cores := uint32(max(1, vcpus/2))
		threads := uint32(max(1, vcpus))
		machineID := fmt.Sprintf("node-%04x-%04x-%04x-%04d", rng.Intn(0xffff), rng.Intn(0xffff), rng.Intn(0xffff), i+1)
		hostname := fmt.Sprintf("onprem-%s-%03d", preset.role, i+1)

		hostOctet := (i % 240) + 10
		ipAddr := fmt.Sprintf("10.0.%d.%d", preset.subnetOctet, hostOctet)
		cidr := fmt.Sprintf("%s/24", ipAddr)
		gwIP := fmt.Sprintf("10.0.%d.1", preset.subnetOctet)
		macAddr := fmt.Sprintf("02:42:0a:00:%02x:%02x", preset.subnetOctet, hostOctet)

		fwCopy := make([]onpremmodel.FirewallRuleProperty, len(preset.firewallTable))
		copy(fwCopy, preset.firewallTable)

		node := onpremmodel.NodeProperty{
			Hostname:  hostname,
			MachineId: machineID,
			Role:      preset.role,
			CPU: onpremmodel.CpuProperty{
				Architecture: "x86_64",
				Cpus:         1,
				Cores:        cores,
				Threads:      threads,
				MaxSpeed:     2.5,
				Vendor:       "GenuineIntel",
				Model:        "Intel(R) Xeon(R) Platinum 8259CL CPU @ 2.50GHz",
			},
			Memory: onpremmodel.MemoryProperty{
				Type:      "DDR4",
				TotalSize: uint64(memGB),
				Used:      uint64(max(1, memGB/2)),
				Available: uint64(max(1, memGB-memGB/2)),
			},
			RootDisk: onpremmodel.DiskProperty{
				Label:     "/dev/sda1",
				Type:      "SSD",
				TotalSize: uint64(diskGB),
				Used:      uint64(max(5, diskGB/3)),
				Available: uint64(max(5, diskGB-diskGB/3)),
			},
			Interfaces: []onpremmodel.NetworkInterfaceProperty{
				{
					Name:           "ens5",
					MacAddress:     macAddr,
					IPv4CidrBlocks: []string{cidr},
					Mtu:            1500,
					State:          "UP",
				},
			},
			RoutingTable: []onpremmodel.RouteProperty{
				{
					Destination: "0.0.0.0/0",
					Gateway:     gwIP,
					Interface:   "ens5",
					Metric:      100,
					Protocol:    "dhcp",
					Scope:       "global",
					LinkState:   "UP",
				},
			},
			FirewallTable: fwCopy,
			OS: onpremmodel.OsProperty{
				PrettyName:      osInfo.prettyName,
				Version:         osInfo.version,
				Name:            osInfo.name,
				VersionID:       osInfo.versionID,
				VersionCodename: osInfo.codename,
				ID:              osInfo.id,
				IDLike:          osInfo.idLike,
			},
		}

		nodes = append(nodes, node)
		gateways = append(gateways, onpremmodel.GatewayProperty{
			InterfaceName: "ens5",
			IP:            gwIP,
			MachineId:     machineID,
		})
	}

	return onpremmodel.OnpremInfra{
		Network: onpremmodel.NetworkProperty{
			IPv4Networks: onpremmodel.NetworkDetail{
				CidrBlocks:      []string{"10.0.1.0/24", "10.0.2.0/24"},
				DefaultGateways: gateways,
			},
		},
		Nodes: nodes,
	}
}

// SummarizeOnpremResources returns the total vCPUs, memory (GiB), and root disk (GB) across all nodes.
func SummarizeOnpremResources(infra onpremmodel.OnpremInfra) (int, int, int) {
	totalVCPU := 0
	totalMem := 0
	totalDisk := 0
	for _, n := range infra.Nodes {
		vcpu := int(n.CPU.Cpus * n.CPU.Threads)
		if vcpu <= 0 {
			vcpu = int(n.CPU.Cores)
		}
		totalVCPU += max(1, vcpu)
		totalMem += int(n.Memory.TotalSize)
		totalDisk += int(n.RootDisk.TotalSize)
	}
	return totalVCPU, totalMem, totalDisk
}

// ConsolidateIdenticalNodeGroups merges NodeGroups that share identical spec, image, disk, and security groups.
func ConsolidateIdenticalNodeGroups(rec cloudmodel.RecommendedInfra) cloudmodel.RecommendedInfra {
	if len(rec.TargetInfra.NodeGroups) <= 1 {
		return rec
	}

	type groupKey struct {
		specID       string
		imageID      string
		cspImageName string
		diskType     string
		diskSize     int
		subnetID     string
		sgKey        string
	}

	order := make([]groupKey, 0)
	merged := make(map[groupKey]cloudmodel.CreateNodeGroupReq)
	sourceIDs := make(map[groupKey][]string)
	oldNamesByKey := make(map[groupKey][]string)

	for _, ng := range rec.TargetInfra.NodeGroups {
		key := groupKey{
			specID:       ng.SpecId,
			imageID:      ng.ImageId,
			cspImageName: ng.CspImageName,
			diskType:     ng.RootDiskType,
			diskSize:     ng.RootDiskSize,
			subnetID:     ng.SubnetId,
			sgKey:        strings.Join(ng.SecurityGroupIds, ","),
		}

		size := ng.NodeGroupSize
		if size <= 0 {
			size = 1
		}

		existing, ok := merged[key]
		if !ok {
			ngCopy := ng
			ngCopy.NodeGroupSize = size
			merged[key] = ngCopy
			order = append(order, key)
		} else {
			existing.NodeGroupSize += size
			merged[key] = existing
		}
		if ng.Name != "" {
			oldNamesByKey[key] = append(oldNamesByKey[key], ng.Name)
		}

		if ng.Label != nil {
			if srcID, has := ng.Label[lineage.LabelSourceMachineIds]; has && srcID != "" {
				sourceIDs[key] = append(sourceIDs[key], srcID)
			}
		}
	}

	consolidated := make([]cloudmodel.CreateNodeGroupReq, 0, len(order))
	renamedMap := make(map[string]string)
	for i, key := range order {
		ng := merged[key]
		newName := fmt.Sprintf("ng-tier-%02d", i+1)
		for _, oldName := range oldNamesByKey[key] {
			renamedMap[oldName] = newName
		}
		ng.Name = newName
		if len(sourceIDs[key]) > 0 {
			if ng.Label == nil {
				ng.Label = make(map[string]string)
			}
			ng.Label[lineage.LabelSourceMachineIds] = strings.Join(sourceIDs[key], ",")
		}
		consolidated = append(consolidated, ng)
	}

	for i, pc := range rec.TargetInfra.PostCommands {
		if newName, ok := renamedMap[pc.NodeGroupId]; ok {
			rec.TargetInfra.PostCommands[i].NodeGroupId = newName
		}
	}
	for i, nlb := range rec.TargetNlbList {
		if newName, ok := renamedMap[nlb.TargetGroup.NodeGroupId]; ok {
			rec.TargetNlbList[i].TargetGroup.NodeGroupId = newName
		}
	}

	rec.TargetInfra.NodeGroups = consolidated
	return rec
}

func randomStepInt(rng *rand.Rand, minVal, maxVal, step int) int {
	if maxVal <= minVal || step <= 0 {
		return minVal
	}
	steps := (maxVal - minVal) / step
	if steps <= 0 {
		return minVal
	}
	return minVal + rng.Intn(steps+1)*step
}
