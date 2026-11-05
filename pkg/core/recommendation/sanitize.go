package recommendation

import (
	"fmt"
	"strings"

	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/rs/zerolog/log"
)

const (
	minSanitizedCpuSockets uint32 = 1
	minSanitizedCpuCores   uint32 = 1
	minSanitizedCpuThreads uint32 = 1
	minSanitizedMemoryGiB  uint64 = 1
	minSanitizedDiskGB     uint64 = 1
	shortMachineIdLen      int    = 8
	maxHostSegmentLen      int    = 22
)

// splitDecomposedMachineId extracts the physical MachineId and GPU model suffix from a decomposed ID.
// Example: "ec289605-...-primary-a100" -> ("ec289605-...", "a100").
func splitDecomposedMachineId(machineId string) (string, string) {
	if idx := strings.LastIndex(machineId, "-primary-"); idx > 0 {
		return machineId[:idx], sanitizeName(machineId[idx+len("-primary-"):])
	}
	if idx := strings.LastIndex(machineId, "-gpu-"); idx > 0 {
		return machineId[:idx], sanitizeName(machineId[idx+len("-gpu-"):])
	}
	return machineId, ""
}

// cleanAlphanumeric strips non-alphanumeric characters and returns a lowercase string.
func cleanAlphanumeric(s string) string {
	var b strings.Builder
	for _, r := range strings.ToLower(s) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// buildShortMachineIdMap computes a unique 8-character identifier per physical MachineId.
// Example: "ec289605-629d-eddf-d4df-0f4c172d84df" -> "172d84df" (last 8 chars).
func buildShortMachineIdMap(nodes []onpremmodel.NodeProperty) map[string]string {
	result := make(map[string]string, len(nodes))
	used := make(map[string]bool, len(nodes))

	for _, node := range nodes {
		origId, _ := splitDecomposedMachineId(node.MachineId)
		if _, exists := result[origId]; exists {
			continue
		}
		cleaned := cleanAlphanumeric(origId)
		if cleaned == "" {
			cleaned = "node0000"
		}

		candidate := cleaned
		if len(cleaned) > shortMachineIdLen {
			candidate = cleaned[len(cleaned)-shortMachineIdLen:]
		}

		if used[candidate] && len(cleaned) > shortMachineIdLen {
			head := cleaned[:shortMachineIdLen]
			if !used[head] {
				candidate = head
			}
		}

		base := candidate
		seq := 2
		for used[candidate] {
			suffix := fmt.Sprintf("%d", seq)
			if len(base)+len(suffix) > shortMachineIdLen {
				candidate = base[:shortMachineIdLen-len(suffix)] + suffix
			} else {
				candidate = base + suffix
			}
			seq++
		}

		used[candidate] = true
		result[origId] = candidate
	}
	return result
}

// hasAsciiLetter reports whether s contains at least one lowercase ASCII letter.
func hasAsciiLetter(s string) bool {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			return true
		}
	}
	return false
}

// extractHostnameLabel extracts the short hostname label and sanitizes it for NodeGroup naming.
// Example: "ai-gpu-01.corp.local" -> "ai-gpu-01".
func extractHostnameLabel(hostname string) string {
	trimmed := strings.TrimSpace(strings.ToLower(hostname))
	if trimmed == "" {
		return ""
	}
	if strings.Contains(trimmed, ".") && hasAsciiLetter(trimmed) {
		parts := strings.Split(trimmed, ".")
		if parts[0] != "" {
			trimmed = parts[0]
		}
	}
	sanitized := sanitizeName(trimmed)
	for _, prefix := range []string{"ng-", "vm-", "migrated-"} {
		if strings.HasPrefix(sanitized, prefix) && len(sanitized) > len(prefix) {
			sanitized = strings.TrimPrefix(sanitized, prefix)
		}
	}
	if len(sanitized) > maxHostSegmentLen {
		sanitized = strings.Trim(sanitized[:maxHostSegmentLen], "-")
	}
	return sanitized
}

// buildNodeGroupNameAndLineage returns a concise, unique NodeGroup name and its physical MachineId.
// Standalone (1:1): "ng-{hostname}-{shortId8}", Decomposed GPU (1:N): "ng-{hostname}-{gpuModel}".
func buildNodeGroupNameAndLineage(
	node onpremmodel.NodeProperty,
	shortIds map[string]string,
	usedNames map[string]bool,
) (string, string) {
	origMachineId, gpuModel := splitDecomposedMachineId(node.MachineId)
	shortId := shortIds[origMachineId]
	if shortId == "" {
		cleaned := cleanAlphanumeric(origMachineId)
		if len(cleaned) > shortMachineIdLen {
			shortId = cleaned[len(cleaned)-shortMachineIdLen:]
		} else if cleaned != "" {
			shortId = cleaned
		} else {
			shortId = sanitizeName(origMachineId)
		}
	}

	host := extractHostnameLabel(node.Hostname)
	sanitizedOrigId := sanitizeName(origMachineId)

	var baseName string
	if gpuModel != "" {
		if host != "" && host != sanitizedOrigId {
			baseName = fmt.Sprintf("ng-%s-%s", host, gpuModel)
		} else {
			baseName = fmt.Sprintf("ng-%s-%s", shortId, gpuModel)
		}
		if usedNames[baseName] && !strings.Contains(baseName, shortId) {
			baseName = fmt.Sprintf("ng-%s-%s-%s", host, gpuModel, shortId)
		}
	} else {
		if host == "" || host == shortId || host == sanitizedOrigId {
			baseName = fmt.Sprintf("ng-%s", shortId)
		} else {
			baseName = fmt.Sprintf("ng-%s-%s", host, shortId)
		}
	}

	finalName := sanitizeName(baseName)
	if len(finalName) < 3 {
		finalName = "ng-" + finalName
	}
	seq := 2
	candidate := finalName
	for usedNames[candidate] {
		candidate = fmt.Sprintf("%s-%d", finalName, seq)
		seq++
	}
	usedNames[candidate] = true
	return candidate, origMachineId
}

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
