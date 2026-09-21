package recommendation

import (
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/rs/zerolog/log"
)

const (
	minSanitizedCpuSockets uint32 = 1
	minSanitizedCpuCores   uint32 = 1
	minSanitizedCpuThreads uint32 = 1
	minSanitizedMemoryGiB  uint64 = 1
	minSanitizedDiskGB     uint64 = 1
)

// SanitizeSourceInfra creates a sanitized copy of srcInfra with valid node specifications.
func SanitizeSourceInfra(srcInfra onpremmodel.OnpremInfra) onpremmodel.OnpremInfra {
	if len(srcInfra.Nodes) == 0 {
		return srcInfra
	}

	sanitizedNodes := make([]onpremmodel.NodeProperty, len(srcInfra.Nodes))
	for i, node := range srcInfra.Nodes {
		sanitizedNodes[i] = SanitizeSourceNode(node)
	}

	srcInfra.Nodes = sanitizedNodes
	return srcInfra
}

// SanitizeSourceNode mitigates missing or zero-valued hardware fields in a source node.
func SanitizeSourceNode(node onpremmodel.NodeProperty) onpremmodel.NodeProperty {
	sanitized := node

	// Cross-restore missing threads from cores, or fallback to minimum.
	if node.CPU.Threads == 0 {
		if node.CPU.Cores > 0 {
			sanitized.CPU.Threads = node.CPU.Cores
			log.Warn().
				Str("hostname", node.Hostname).
				Str("machineId", node.MachineId).
				Uint32("rawThreads", node.CPU.Threads).
				Uint32("restoredThreads", sanitized.CPU.Threads).
				Msg("Source node CPU threads is 0; restored from cores for recommendation")
		} else {
			sanitized.CPU.Threads = minSanitizedCpuThreads
			log.Warn().
				Str("hostname", node.Hostname).
				Str("machineId", node.MachineId).
				Uint32("rawThreads", node.CPU.Threads).
				Uint32("sanitizedThreads", sanitized.CPU.Threads).
				Msg("Source node CPU threads is 0; defaulting to 1 thread for recommendation")
		}
	}

	// Cross-restore missing cores from threads, or fallback to minimum.
	if node.CPU.Cores == 0 {
		if sanitized.CPU.Threads > 0 {
			sanitized.CPU.Cores = sanitized.CPU.Threads
			log.Warn().
				Str("hostname", node.Hostname).
				Str("machineId", node.MachineId).
				Uint32("rawCores", node.CPU.Cores).
				Uint32("restoredCores", sanitized.CPU.Cores).
				Msg("Source node CPU cores is 0; restored from threads for recommendation")
		} else {
			sanitized.CPU.Cores = minSanitizedCpuCores
			log.Warn().
				Str("hostname", node.Hostname).
				Str("machineId", node.MachineId).
				Uint32("rawCores", node.CPU.Cores).
				Uint32("sanitizedCores", sanitized.CPU.Cores).
				Msg("Source node CPU cores is 0; defaulting to 1 core for recommendation")
		}
	}

	// Mitigate missing or zero CPU socket count.
	if node.CPU.Cpus == 0 {
		sanitized.CPU.Cpus = minSanitizedCpuSockets
		log.Warn().
			Str("hostname", node.Hostname).
			Str("machineId", node.MachineId).
			Uint32("rawCpus", node.CPU.Cpus).
			Uint32("sanitizedCpus", sanitized.CPU.Cpus).
			Msg("Source node CPU cpus is 0; defaulting to 1 socket for recommendation")
	}

	// Mitigate missing or truncated memory size (e.g., cm-honeybee#75 integer division).
	if node.Memory.TotalSize == 0 {
		sanitized.Memory.TotalSize = minSanitizedMemoryGiB
		log.Warn().
			Str("hostname", node.Hostname).
			Str("machineId", node.MachineId).
			Uint64("rawMemoryTotalSize", node.Memory.TotalSize).
			Uint64("sanitizedMemoryTotalSize", sanitized.Memory.TotalSize).
			Msg("Source node memory totalSize is 0; defaulting to 1 GiB for recommendation")
	}

	// Mitigate missing or zero root disk size.
	if node.RootDisk.TotalSize == 0 {
		sanitized.RootDisk.TotalSize = minSanitizedDiskGB
		log.Warn().
			Str("hostname", node.Hostname).
			Str("machineId", node.MachineId).
			Uint64("rawRootDiskTotalSize", node.RootDisk.TotalSize).
			Uint64("sanitizedRootDiskTotalSize", sanitized.RootDisk.TotalSize).
			Msg("Source node rootDisk totalSize is 0; defaulting to 1 GB for recommendation")
	}

	return sanitized
}
