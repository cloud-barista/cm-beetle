package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cloud-barista/cm-beetle/transx"
)

func main() {
	var configFile string
	var verbose bool
	var backupOnly bool
	var transferOnly bool
	var restoreOnly bool

	// Setting up command-line flags
	flag.StringVar(&configFile, "config", "direct-mode-config.json", "Migration configuration JSON file path")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging")
	flag.BoolVar(&backupOnly, "backup", false, "Run only the backup step")
	flag.BoolVar(&transferOnly, "transfer", false, "Run only the transfer step")
	flag.BoolVar(&restoreOnly, "restore", false, "Run only the restore step")
	flag.Parse()

	// Record start time (for performance measurement)
	startTime := time.Now()

	// Check configuration file path
	if !filepath.IsAbs(configFile) {
		// Convert relative path to absolute path
		workingDir, err := os.Getwd()
		if err == nil {
			configFile = filepath.Join(workingDir, configFile)
		}
	}

	// Read JSON file
	jsonData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to read config file %s: %v", configFile, err)
	}

	// Parse JSON data
	var dmm transx.DataMigrationModel
	err = json.Unmarshal(jsonData, &dmm)
	if err != nil {
		log.Fatalf("Failed to parse config JSON: %v", err)
	}

	// Validate migration configuration file
	err = transx.Validate(dmm)
	if err != nil {
		log.Fatalf("Invalid migration configuration: %v", err)
	}

	// Display transfer scenario
	srcStorage := dmm.Source.StorageType
	dstStorage := dmm.Destination.StorageType
	strategy := dmm.Strategy
	if strategy == "" {
		strategy = "auto"
	}

	fmt.Printf("Transfer: %s -> %s (strategy: %s)\n", srcStorage, dstStorage, strategy)
	fmt.Printf("Source path: %s\n", dmm.Source.Path)
	fmt.Printf("Destination path: %s\n", dmm.Destination.Path)

	// Display SSH connection info if applicable
	if dmm.Source.Filesystem != nil && dmm.Source.Filesystem.AccessType == transx.AccessTypeSSH && dmm.Source.Filesystem.SSH != nil {
		ssh := dmm.Source.Filesystem.SSH
		fmt.Printf("Source SSH: %s@%s:%d\n", ssh.Username, ssh.Host, ssh.Port)
	}
	if dmm.Destination.Filesystem != nil && dmm.Destination.Filesystem.AccessType == transx.AccessTypeSSH && dmm.Destination.Filesystem.SSH != nil {
		ssh := dmm.Destination.Filesystem.SSH
		fmt.Printf("Destination SSH: %s@%s:%d\n", ssh.Username, ssh.Host, ssh.Port)
	}

	// Expand tilde (~) in SSH private key paths
	expandSSHKeyPath(&dmm)

	// Display commands (in verbose mode)
	if verbose {
		if dmm.Source.PreCmd != "" {
			fmt.Printf("Pre-command: %s\n", dmm.Source.PreCmd)
		}
		if dmm.Destination.PostCmd != "" {
			fmt.Printf("Post-command: %s\n", dmm.Destination.PostCmd)
		}
	}

	// Execute individual steps if specified, otherwise run complete migration
	if backupOnly {
		fmt.Println("Running backup step only...")
		if err := transx.Backup(dmm); err != nil {
			log.Fatalf("Backup failed: %v", err)
		}
		fmt.Println("Backup completed successfully!")
	} else if transferOnly {
		fmt.Println("Running transfer step only...")
		if err := transx.Transfer(dmm); err != nil {
			log.Fatalf("Transfer failed: %v", err)
		}
		fmt.Println("Transfer completed successfully!")
	} else if restoreOnly {
		fmt.Println("Running restore step only...")
		if err := transx.Restore(dmm); err != nil {
			log.Fatalf("Restore failed: %v", err)
		}
		fmt.Println("Restore completed successfully!")
	} else {
		// Execute the complete data migration workflow
		if err := transx.MigrateData(dmm); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
	}

	// Display summary information
	totalTime := time.Since(startTime)
	fmt.Println("\n=== Migration Summary ===")

	// Build source/destination display strings
	srcDisplay := dmm.Source.Path
	destDisplay := dmm.Destination.Path

	if dmm.Source.Filesystem != nil && dmm.Source.Filesystem.SSH != nil {
		ssh := dmm.Source.Filesystem.SSH
		srcDisplay = fmt.Sprintf("%s@%s:%s", ssh.Username, ssh.Host, dmm.Source.Path)
	}
	if dmm.Destination.Filesystem != nil && dmm.Destination.Filesystem.SSH != nil {
		ssh := dmm.Destination.Filesystem.SSH
		destDisplay = fmt.Sprintf("%s@%s:%s", ssh.Username, ssh.Host, dmm.Destination.Path)
	}

	fmt.Printf("Source: %s\n", srcDisplay)
	fmt.Printf("Destination: %s\n", destDisplay)
	fmt.Printf("Total migration time: %s\n", totalTime)
}

// expandSSHKeyPath expands tilde (~) and relative paths (./) in SSH private key paths
func expandSSHKeyPath(dmm *transx.DataMigrationModel) {
	expandPath := func(path string) string {
		if path == "" {
			return path
		}
		// Expand ~ to home directory
		if strings.HasPrefix(path, "~/") {
			homeDir, _ := os.UserHomeDir()
			return filepath.Join(homeDir, path[2:])
		}
		// Convert relative path to absolute path
		if strings.HasPrefix(path, "./") || !filepath.IsAbs(path) {
			workingDir, err := os.Getwd()
			if err == nil {
				return filepath.Join(workingDir, path)
			}
		}
		return path
	}

	if dmm.Source.Filesystem != nil && dmm.Source.Filesystem.SSH != nil {
		dmm.Source.Filesystem.SSH.PrivateKeyPath = expandPath(dmm.Source.Filesystem.SSH.PrivateKeyPath)
	}
	if dmm.Destination.Filesystem != nil && dmm.Destination.Filesystem.SSH != nil {
		// Don't expand paths that are absolute (e.g., /root/.ssh/id_rsa in container)
		path := dmm.Destination.Filesystem.SSH.PrivateKeyPath
		if !filepath.IsAbs(path) {
			dmm.Destination.Filesystem.SSH.PrivateKeyPath = expandPath(path)
		}
	}
}
