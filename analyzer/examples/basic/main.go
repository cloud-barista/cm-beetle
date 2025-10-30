package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cloud-barista/cm-beetle/analyzer"
)

func main() {
	fmt.Println("=== Analyzer - File Data Migration Example ===")

	// Example 1: List default directory
	fmt.Println("1. Listing default home directory...")
	result, err := analyzer.ListDirectory("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Base Directory: %s\n", result.BaseDir)
	fmt.Printf("Total Files: %d, Total Directories: %d\n", result.TotalFiles, result.TotalDirs)
	fmt.Printf("First 5 entries:\n")
	for i, entry := range result.Entries {
		if i >= 5 {
			break
		}
		entryType := "File"
		if entry.IsDir {
			entryType = "Dir "
		}
		fmt.Printf("  [%s] %s (Size: %d bytes)\n", entryType, entry.Name, entry.Size)
	}

	// Example 2: List specific directory
	fmt.Println("\n2. Listing /tmp directory...")
	tmpResult, err := analyzer.ListDirectory("/tmp")
	if err != nil {
		log.Printf("Error listing /tmp: %v\n", err)
	} else {
		fmt.Printf("Total entries in /tmp: %d\n", len(tmpResult.Entries))
	}

	// Example 3: Scan directory recursively with filters
	fmt.Println("\n3. Scanning home directory recursively (max depth 2, exclude hidden files)...")
	homeDir, _ := analyzer.GetDefaultBaseDir()
	scanOptions := analyzer.ScanOptions{
		BaseDir:       homeDir,
		Recursive:     true,
		MaxDepth:      2,
		IncludeHidden: false,
		ExcludePatterns: []string{
			"*.log",
			".git/**",
			".cache/**",
			"node_modules/**",
		},
	}
	scanResult, err := analyzer.ScanDirectory(scanOptions)
	if err != nil {
		log.Printf("Error scanning: %v\n", err)
	} else {
		fmt.Printf("Total Files: %d, Total Directories: %d, Total Size: %d bytes\n",
			scanResult.TotalFiles, scanResult.TotalDirs, scanResult.TotalSize)
	}

	// Example 4: Create migration plan with filters
	fmt.Println("\n4. Creating migration plan...")
	filters := analyzer.FilterOptions{
		IncludePatterns: []string{
			"*.txt",
			"*.md",
			"docs/**",
		},
		ExcludePatterns: []string{
			"*.tmp",
			"*.log",
			".git/**",
		},
	}

	plan, err := analyzer.CreateMigrationPlan(homeDir+"/Documents", true, filters)
	if err != nil {
		log.Printf("Error creating migration plan: %v\n", err)
	} else {
		fmt.Printf("Migration Plan Created:\n")
		fmt.Printf("  Source: %s\n", plan.SourceDir)
		fmt.Printf("  Total Files: %d\n", plan.TotalFiles)
		fmt.Printf("  Total Size: %d bytes\n", plan.TotalSize)
		fmt.Printf("  Include Subdirectories: %v\n", plan.IncludeSubDir)
		fmt.Printf("  Created At: %s\n", plan.CreatedAt.Format("2006-01-02 15:04:05"))

		// Show first 5 files
		if len(plan.FileList) > 0 {
			fmt.Println("  Sample files:")
			for i, file := range plan.FileList {
				if i >= 5 {
					break
				}
				fmt.Printf("    - %s (%d bytes, modified: %s)\n",
					file.Name, file.Size, file.ModTime.Format("2006-01-02"))
			}
		}
	}

	// Example 5: Extract metadata for a specific file
	fmt.Println("\n5. Extracting metadata for a specific file...")
	testFile := "/etc/hosts"
	metadata, err := analyzer.ExtractFileMetadata(testFile, false)
	if err != nil {
		log.Printf("Error extracting metadata: %v\n", err)
	} else {
		jsonData, _ := json.MarshalIndent(metadata, "", "  ")
		fmt.Printf("Metadata for %s:\n%s\n", testFile, string(jsonData))
	}

	// Example 6: Get directory statistics
	fmt.Println("\n6. Getting directory statistics...")
	fileCount, dirCount, totalSize, err := analyzer.GetDirectoryStatistics("/var/log", false)
	if err != nil {
		log.Printf("Error getting statistics: %v\n", err)
	} else {
		fmt.Printf("/var/log statistics (non-recursive):\n")
		fmt.Printf("  Files: %d, Directories: %d, Total Size: %d bytes\n",
			fileCount, dirCount, totalSize)
	}

	fmt.Println("\n=== Example completed ===")
}
