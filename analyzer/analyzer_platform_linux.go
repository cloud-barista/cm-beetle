// +build linux

package analyzer

import (
	"fmt"
	"os"
	"syscall"
)

// extractPlatformMetadata extracts Linux-specific metadata
func extractPlatformMetadata(info os.FileInfo, metadata *FileMetadata) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		// Extract Unix/Linux specific metadata
		metadata.Owner = fmt.Sprintf("%d", stat.Uid)
		metadata.Group = fmt.Sprintf("%d", stat.Gid)
		// Note: Access time and change time can be extracted from stat.Atim and stat.Ctim
		// but require platform-specific handling
	}
}
