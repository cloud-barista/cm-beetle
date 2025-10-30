package analyzer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ========================================
// Core Data Structures for File Metadata Collection
// ========================================

// FileMetadata represents comprehensive metadata for a file or directory
type FileMetadata struct {
	Path          string    `json:"path"`          // Full absolute path
	Name          string    `json:"name"`          // File or directory name
	Size          int64     `json:"size"`          // Size in bytes (0 for directories)
	IsDir         bool      `json:"isDir"`         // True if directory
	Mode          string    `json:"mode"`          // Permission mode (e.g., "0755")
	ModTime       time.Time `json:"modTime"`       // Last modification time
	AccessTime    time.Time `json:"accessTime"`    // Last access time (if available)
	ChangeTime    time.Time `json:"changeTime"`    // Last status change time (if available)
	Owner         string    `json:"owner"`         // Owner user ID
	Group         string    `json:"group"`         // Group ID
	MimeType      string    `json:"mimeType"`      // MIME type (for files)
	Extension     string    `json:"extension"`     // File extension (e.g., ".txt")
	IsSymlink     bool      `json:"isSymlink"`     // True if symbolic link
	SymlinkTarget string    `json:"symlinkTarget"` // Target if symbolic link
	Checksum      string    `json:"checksum"`      // Optional checksum (MD5/SHA256)
}

// DirectoryEntry represents a single directory or file entry for listing
type DirectoryEntry struct {
	Name  string `json:"name"`  // Name of the file or directory
	Path  string `json:"path"`  // Full path
	IsDir bool   `json:"isDir"` // True if directory
	Size  int64  `json:"size"`  // Size in bytes
}

// ScanResult contains the results of a directory scan
type ScanResult struct {
	BaseDir       string           `json:"baseDir"`       // Base directory scanned
	Entries       []DirectoryEntry `json:"entries"`       // Immediate entries (files + dirs)
	TotalFiles    int              `json:"totalFiles"`    // Total file count (recursive)
	TotalDirs     int              `json:"totalDirs"`     // Total directory count (recursive)
	TotalSize     int64            `json:"totalSize"`     // Total size in bytes (recursive)
	ScanTime      time.Time        `json:"scanTime"`      // When the scan was performed
	IncludeSubDir bool             `json:"includeSubDir"` // Whether subdirectories were included
}

// ScanOptions defines options for scanning directories
type ScanOptions struct {
	BaseDir         string   `json:"baseDir"`         // Base directory to scan (defaults to $HOME)
	Recursive       bool     `json:"recursive"`       // Include subdirectories
	FollowSymlinks  bool     `json:"followSymlinks"`  // Follow symbolic links
	IncludeHidden   bool     `json:"includeHidden"`   // Include hidden files (starting with .)
	MaxDepth        int      `json:"maxDepth"`        // Maximum recursion depth (0 = unlimited)
	CollectChecksum bool     `json:"collectChecksum"` // Calculate file checksums (expensive)
	IncludePatterns []string `json:"includePatterns"` // Include patterns (whitelist)
	ExcludePatterns []string `json:"excludePatterns"` // Exclude patterns (blacklist)
}

// FilterOptions defines include/exclude filter settings
type FilterOptions struct {
	IncludePatterns []string `json:"includePatterns"` // Patterns to include (e.g., "*.txt", "data/**")
	ExcludePatterns []string `json:"excludePatterns"` // Patterns to exclude (e.g., "*.log", "temp/*")
}

// MigrationPlan represents the migration plan configuration
type MigrationPlan struct {
	SourceDir     string         `json:"sourceDir"`     // Source directory path
	IncludeSubDir bool           `json:"includeSubDir"` // Include subdirectories
	FilterOptions FilterOptions  `json:"filterOptions"` // Filter settings
	TotalFiles    int            `json:"totalFiles"`    // Number of files to migrate
	TotalSize     int64          `json:"totalSize"`     // Total size to migrate
	FileList      []FileMetadata `json:"fileList"`      // List of files to migrate
	CreatedAt     time.Time      `json:"createdAt"`     // When the plan was created
}

// ========================================
// Directory Listing Functions
// ========================================

// GetDefaultBaseDir returns the default base directory ($HOME for Linux/Unix)
func GetDefaultBaseDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return homeDir, nil
}

// ListDirectory lists immediate files and subdirectories in the given path
// This is the main function for UI-driven directory browsing
func ListDirectory(path string) (*ScanResult, error) {
	// If path is empty, use default base directory
	if path == "" {
		defaultDir, err := GetDefaultBaseDir()
		if err != nil {
			return nil, err
		}
		path = defaultDir
	}

	// Verify path exists and is a directory
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to access path %s: %w", path, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("path %s is not a directory", path)
	}

	// Read directory entries
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", path, err)
	}

	// Build result
	result := &ScanResult{
		BaseDir:       path,
		Entries:       make([]DirectoryEntry, 0, len(entries)),
		ScanTime:      time.Now(),
		IncludeSubDir: false,
	}

	for _, entry := range entries {
		entryInfo, err := entry.Info()
		if err != nil {
			continue // Skip entries we can't read
		}

		fullPath := filepath.Join(path, entry.Name())
		dirEntry := DirectoryEntry{
			Name:  entry.Name(),
			Path:  fullPath,
			IsDir: entry.IsDir(),
			Size:  entryInfo.Size(),
		}

		result.Entries = append(result.Entries, dirEntry)

		// Update statistics
		if entry.IsDir() {
			result.TotalDirs++
		} else {
			result.TotalFiles++
			result.TotalSize += entryInfo.Size()
		}
	}

	return result, nil
}

// ScanDirectory performs a comprehensive scan of a directory with optional recursion
// This is used when user confirms migration target and needs full file list
func ScanDirectory(options ScanOptions) (*ScanResult, error) {
	// Set default base directory if not specified
	if options.BaseDir == "" {
		defaultDir, err := GetDefaultBaseDir()
		if err != nil {
			return nil, err
		}
		options.BaseDir = defaultDir
	}

	// Verify base directory exists
	info, err := os.Stat(options.BaseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to access base directory %s: %w", options.BaseDir, err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("base path %s is not a directory", options.BaseDir)
	}

	result := &ScanResult{
		BaseDir:       options.BaseDir,
		Entries:       make([]DirectoryEntry, 0),
		ScanTime:      time.Now(),
		IncludeSubDir: options.Recursive,
	}

	// Scan directory
	if options.Recursive {
		err = scanRecursive(options.BaseDir, options, result, 0)
	} else {
		err = scanSingleLevel(options.BaseDir, options, result)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

// scanSingleLevel scans only the immediate directory level
func scanSingleLevel(dirPath string, options ScanOptions, result *ScanResult) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	for _, entry := range entries {
		// Skip hidden files if not included
		if !options.IncludeHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		fullPath := filepath.Join(dirPath, entry.Name())

		// Apply filters
		if !shouldIncludePath(fullPath, options.BaseDir, options.IncludePatterns, options.ExcludePatterns) {
			continue
		}

		entryInfo, err := entry.Info()
		if err != nil {
			continue // Skip entries we can't read
		}

		dirEntry := DirectoryEntry{
			Name:  entry.Name(),
			Path:  fullPath,
			IsDir: entry.IsDir(),
			Size:  entryInfo.Size(),
		}

		result.Entries = append(result.Entries, dirEntry)

		// Update statistics
		if entry.IsDir() {
			result.TotalDirs++
		} else {
			result.TotalFiles++
			result.TotalSize += entryInfo.Size()
		}
	}

	return nil
}

// scanRecursive performs recursive directory scanning
func scanRecursive(dirPath string, options ScanOptions, result *ScanResult, depth int) error {
	// Check max depth
	if options.MaxDepth > 0 && depth > options.MaxDepth {
		return nil
	}

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Skip directories we can't access
			return nil
		}

		// Skip the base directory itself
		if path == dirPath {
			return nil
		}

		// Skip hidden files if not included
		if !options.IncludeHidden && strings.HasPrefix(d.Name(), ".") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Handle symlinks
		if d.Type()&os.ModeSymlink != 0 && !options.FollowSymlinks {
			return nil
		}

		// Apply filters
		if !shouldIncludePath(path, options.BaseDir, options.IncludePatterns, options.ExcludePatterns) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Get file info
		info, err := d.Info()
		if err != nil {
			return nil // Skip entries we can't read
		}

		dirEntry := DirectoryEntry{
			Name:  d.Name(),
			Path:  path,
			IsDir: d.IsDir(),
			Size:  info.Size(),
		}

		result.Entries = append(result.Entries, dirEntry)

		// Update statistics
		if d.IsDir() {
			result.TotalDirs++
		} else {
			result.TotalFiles++
			result.TotalSize += info.Size()
		}

		return nil
	})

	return err
}

// ========================================
// Filter Functions (Based on transx pattern matching)
// ========================================

// shouldIncludePath determines if a path should be included based on filter patterns
// This follows rsync-like filtering logic:
// 1. If include patterns exist, path must match at least one
// 2. Exclude patterns are then applied to remove unwanted paths
// 3. If no patterns specified, all paths are included
func shouldIncludePath(fullPath, baseDir string, includePatterns, excludePatterns []string) bool {
	// Get relative path from base directory
	relPath, err := filepath.Rel(baseDir, fullPath)
	if err != nil {
		// If we can't get relative path, use full path
		relPath = fullPath
	}

	// Normalize path separators for consistent matching
	relPath = filepath.ToSlash(relPath)

	// Step 1: Apply include patterns (whitelist)
	if len(includePatterns) > 0 {
		included := false
		for _, pattern := range includePatterns {
			if matchPattern(relPath, pattern) {
				included = true
				break
			}
		}
		if !included {
			return false // Path doesn't match any include pattern
		}
	}

	// Step 2: Apply exclude patterns (blacklist)
	for _, pattern := range excludePatterns {
		if matchPattern(relPath, pattern) {
			return false // Path matches an exclude pattern
		}
	}

	return true // Path passed all filters
}

// matchPattern checks if a file path matches a pattern
// Supports glob patterns including ** for recursive directory matching
// Pattern examples:
//   - "*.log" matches any .log file
//   - "data/*" matches files directly in data/ directory
//   - "data/**" matches all files recursively under data/
//   - "data/**/*.json" matches all .json files recursively under data/
func matchPattern(path, pattern string) bool {
	// Handle ** (double asterisk) for recursive directory matching
	if strings.Contains(pattern, "**") {
		return matchPatternWithDoubleAsterisk(path, pattern)
	}

	// Direct match using filepath.Match
	matched, err := filepath.Match(pattern, path)
	if err == nil && matched {
		return true
	}

	// Try matching with path prefix for directory patterns
	if strings.Contains(pattern, "/") {
		matched, err := filepath.Match(pattern, path)
		if err == nil && matched {
			return true
		}
	}

	// Try matching basename for simple patterns (e.g., "*.log")
	if !strings.Contains(pattern, "/") {
		basename := filepath.Base(path)
		matched, err := filepath.Match(pattern, basename)
		if err == nil && matched {
			return true
		}
	}

	return false
}

// matchPatternWithDoubleAsterisk handles patterns with ** for recursive matching
func matchPatternWithDoubleAsterisk(path, pattern string) bool {
	parts := strings.Split(pattern, "**")

	if len(parts) == 1 {
		matched, _ := filepath.Match(pattern, path)
		return matched
	}

	if len(parts) == 2 {
		prefix := strings.TrimSuffix(parts[0], "/")
		suffix := strings.TrimPrefix(parts[1], "/")

		// Case 1: "data/**" - matches everything under data/
		if prefix != "" && suffix == "" {
			if strings.ContainsAny(prefix, "*?[") {
				return matchPrefixWithGlob(path, prefix)
			}
			return strings.HasPrefix(path, prefix+"/") || path == prefix
		}

		// Case 2: "**/*.json" - matches all .json files anywhere
		if prefix == "" && suffix != "" {
			if matched, _ := filepath.Match(suffix, filepath.Base(path)); matched {
				return true
			}
			if strings.HasSuffix(path, suffix) {
				return true
			}
			pathParts := strings.Split(path, "/")
			for i := range pathParts {
				subPath := strings.Join(pathParts[i:], "/")
				if matched, _ := filepath.Match(suffix, subPath); matched {
					return true
				}
			}
			return false
		}

		// Case 3: "data/**/file.txt" - matches with both prefix and suffix
		if prefix != "" && suffix != "" {
			if strings.ContainsAny(prefix, "*?[") {
				return matchPrefixSuffixWithGlob(path, prefix, suffix)
			}
			if !strings.HasPrefix(path, prefix+"/") && path != prefix {
				return false
			}
			if suffix == "" {
				return true
			}
			if strings.HasSuffix(path, suffix) {
				return true
			}
			pathAfterPrefix := strings.TrimPrefix(path, prefix+"/")
			pathParts := strings.Split(pathAfterPrefix, "/")
			for i := range pathParts {
				subPath := strings.Join(pathParts[i:], "/")
				if matched, _ := filepath.Match(suffix, subPath); matched {
					return true
				}
			}
			return false
		}
	}

	return false
}

// matchPrefixWithGlob matches path against a prefix pattern with wildcards
func matchPrefixWithGlob(path, prefix string) bool {
	pathParts := strings.Split(path, "/")
	prefixParts := strings.Split(prefix, "/")

	if len(pathParts) < len(prefixParts) {
		return false
	}

	for i := 0; i <= len(pathParts)-len(prefixParts); i++ {
		matched := true
		for j, prefixPart := range prefixParts {
			if partMatched, _ := filepath.Match(prefixPart, pathParts[i+j]); !partMatched {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}

	return false
}

// matchPrefixSuffixWithGlob matches path against prefix and suffix patterns
func matchPrefixSuffixWithGlob(path, prefix, suffix string) bool {
	pathParts := strings.Split(path, "/")
	prefixParts := strings.Split(prefix, "/")

	if len(pathParts) < len(prefixParts) {
		return false
	}

	for i := 0; i <= len(pathParts)-len(prefixParts); i++ {
		matched := true
		for j, prefixPart := range prefixParts {
			if partMatched, _ := filepath.Match(prefixPart, pathParts[i+j]); !partMatched {
				matched = false
				break
			}
		}
		if matched {
			// Prefix matched, now check suffix
			if suffix == "" {
				return true
			}
			remainingPath := strings.Join(pathParts[i+len(prefixParts):], "/")
			if strings.HasSuffix(remainingPath, suffix) {
				return true
			}
			if matched, _ := filepath.Match(suffix, filepath.Base(remainingPath)); matched {
				return true
			}
		}
	}

	return false
}

// ========================================
// Metadata Extraction Functions
// ========================================

// ExtractFileMetadata extracts comprehensive metadata for a file or directory
func ExtractFileMetadata(path string, collectChecksum bool) (*FileMetadata, error) {
	// Get file info
	info, err := os.Lstat(path) // Use Lstat to handle symlinks properly
	if err != nil {
		return nil, fmt.Errorf("failed to get file info for %s: %w", path, err)
	}

	metadata := &FileMetadata{
		Path:      path,
		Name:      filepath.Base(path),
		Size:      info.Size(),
		IsDir:     info.IsDir(),
		Mode:      info.Mode().String(),
		ModTime:   info.ModTime(),
		Extension: filepath.Ext(path),
	}

	// Check if symlink
	if info.Mode()&os.ModeSymlink != 0 {
		metadata.IsSymlink = true
		target, err := os.Readlink(path)
		if err == nil {
			metadata.SymlinkTarget = target
		}
	}

	// Extract additional metadata (platform-specific)
	extractPlatformMetadata(info, metadata)

	// Calculate checksum if requested and it's a regular file
	if collectChecksum && !info.IsDir() && !metadata.IsSymlink {
		// Note: Checksum calculation can be added here
		// For now, leaving it empty as it can be expensive
		// metadata.Checksum = calculateChecksum(path)
	}

	return metadata, nil
}

// CollectFileList collects detailed metadata for all files in a directory
// This is used when creating a migration plan
func CollectFileList(options ScanOptions) ([]FileMetadata, error) {
	scanResult, err := ScanDirectory(options)
	if err != nil {
		return nil, err
	}

	fileList := make([]FileMetadata, 0, scanResult.TotalFiles)

	for _, entry := range scanResult.Entries {
		// Skip directories if we only want files
		if entry.IsDir {
			continue
		}

		metadata, err := ExtractFileMetadata(entry.Path, options.CollectChecksum)
		if err != nil {
			// Log error but continue with other files
			continue
		}

		fileList = append(fileList, *metadata)
	}

	return fileList, nil
}

// CreateMigrationPlan creates a migration plan based on user's selection
func CreateMigrationPlan(sourceDir string, includeSubDir bool, filters FilterOptions) (*MigrationPlan, error) {
	// Prepare scan options
	options := ScanOptions{
		BaseDir:         sourceDir,
		Recursive:       includeSubDir,
		IncludeHidden:   false,
		FollowSymlinks:  false,
		MaxDepth:        0,     // Unlimited
		CollectChecksum: false, // Don't collect checksums during initial scan
		IncludePatterns: filters.IncludePatterns,
		ExcludePatterns: filters.ExcludePatterns,
	}

	// Collect file list
	fileList, err := CollectFileList(options)
	if err != nil {
		return nil, err
	}

	// Calculate totals
	totalSize := int64(0)
	for _, file := range fileList {
		totalSize += file.Size
	}

	plan := &MigrationPlan{
		SourceDir:     sourceDir,
		IncludeSubDir: includeSubDir,
		FilterOptions: filters,
		TotalFiles:    len(fileList),
		TotalSize:     totalSize,
		FileList:      fileList,
		CreatedAt:     time.Now(),
	}

	return plan, nil
}

// GetDirectoryStatistics calculates statistics for a directory
func GetDirectoryStatistics(path string, recursive bool) (fileCount int, dirCount int, totalSize int64, err error) {
	if !recursive {
		// Single level statistics
		entries, err := os.ReadDir(path)
		if err != nil {
			return 0, 0, 0, err
		}

		for _, entry := range entries {
			if entry.IsDir() {
				dirCount++
			} else {
				fileCount++
				info, err := entry.Info()
				if err == nil {
					totalSize += info.Size()
				}
			}
		}
	} else {
		// Recursive statistics
		err = filepath.WalkDir(path, func(p string, d fs.DirEntry, walkErr error) error {
			if walkErr != nil {
				return nil // Skip errors
			}
			if p == path {
				return nil // Skip base directory
			}
			if d.IsDir() {
				dirCount++
			} else {
				fileCount++
				info, err := d.Info()
				if err == nil {
					totalSize += info.Size()
				}
			}
			return nil
		})
	}

	return fileCount, dirCount, totalSize, err
}
