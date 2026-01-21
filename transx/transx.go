package transx

import (
	"fmt"
	"strings"
)

// Transfer runs the transfer command to transfer data as defined by the given DataMigrationModel.
// * Note: This is the core operation that always executes when called.
func Transfer(dmm DataMigrationModel) error {
	if err := Validate(dmm); err != nil {
		return fmt.Errorf("data migration model validation failed: %w", err)
	}

	// Check if we're operating in relay mode (both source and destination are remote)
	isRelayMode := dmm.IsRelayMode()

	// Handle different transfer scenarios
	if isRelayMode {
		// Relay mode: both endpoints are remote
		return performRelayTransfer(dmm)
	} else {
		// Direct mode: at least one endpoint is local
		return performDirectTransfer(dmm)
	}
}

// Backup executes the BackupCmd defined in the source EndpointDetails of the DataMigrationModel.
// * Note: This is an optional operation that runs only if BackupCmd is configured.
func Backup(dmm DataMigrationModel) error {
	// Use source endpoint for backup operations
	source := dmm.Source
	if strings.TrimSpace(source.BackupCmd) == "" {
		return fmt.Errorf("backup command is not defined for source")
	}

	// Get transfer options for SSH configuration
	transferOptions := dmm.SourceTransferOptions

	output, err := executeCommand(source.BackupCmd, source, transferOptions)
	if err != nil {
		// Build detailed source path for error context
		sourcePath := source.DataPath
		if source.IsRemote() {
			var username string
			if transferOptions != nil && transferOptions.RsyncOptions != nil {
				username = transferOptions.RsyncOptions.Username
			}
			if strings.TrimSpace(username) != "" {
				sourcePath = fmt.Sprintf("%s@%s:%s", username, source.GetEndpoint(), source.DataPath)
			} else {
				sourcePath = fmt.Sprintf("%s:%s", source.GetEndpoint(), source.DataPath)
			}
		}

		// Use inline error creation
		return &OperationError{
			Operation: "backup",
			Source:    sourcePath,
			Command:   source.BackupCmd,
			Output:    string(output),
			Err:       fmt.Errorf("command execution failed: %w", err),
		}
	}

	return nil
}

// Restore executes the RestoreCmd defined in the destination EndpointDetails of the DataMigrationModel.
// * Note: This is an optional operation that runs only if RestoreCmd is configured.
func Restore(dmm DataMigrationModel) error {
	// Use destination endpoint for restore operations
	destination := dmm.Destination
	if strings.TrimSpace(destination.RestoreCmd) == "" {
		return fmt.Errorf("restore command is not defined for destination")
	}

	// Get transfer options for SSH configuration
	transferOptions := dmm.DestinationTransferOptions

	output, err := executeCommand(destination.RestoreCmd, destination, transferOptions)
	if err != nil {
		// Build detailed destination path for error context
		destinationPath := destination.DataPath
		if destination.IsRemote() {
			var username string
			if transferOptions != nil && transferOptions.RsyncOptions != nil {
				username = transferOptions.RsyncOptions.Username
			}
			if strings.TrimSpace(username) != "" {
				destinationPath = fmt.Sprintf("%s@%s:%s", username, destination.GetEndpoint(), destination.DataPath)
			} else {
				destinationPath = fmt.Sprintf("%s:%s", destination.GetEndpoint(), destination.DataPath)
			}
		}

		// Use inline error creation
		return &OperationError{
			Operation:   "restore",
			Destination: destinationPath,
			Command:     destination.RestoreCmd,
			Output:      string(output),
			Err:         fmt.Errorf("command execution failed: %w", err),
		}
	}

	return nil
}

// MigrateData manages the complete data migration workflow:
// 1. If Source.BackupCmd is available, perform Backup
// 2. Always perform Transfer
// 3. If Destination.RestoreCmd is available, perform Restore
// This provides a simple one-call approach to handle the entire data migration pipeline.
// * Note: Transfer is the optional operation that runs BackUp, Transfer, and Restore sequentially.
func MigrateData(dmm DataMigrationModel) error {
	// Step 1: Check and perform backup if BackupCmd is defined
	if strings.TrimSpace(dmm.Source.BackupCmd) != "" {
		if err := Backup(dmm); err != nil {
			return &MigrationError{
				Stage: "backup",
				Err:   fmt.Errorf("backup operation failed: %w", err),
			}
		}
	}

	// Step 2: Always perform the data transfer (core functionality)
	if err := Transfer(dmm); err != nil {
		return &MigrationError{
			Stage: "transfer",
			Err:   fmt.Errorf("data transfer failed: %w", err),
		}
	}

	// Step 3: Check and perform restore if RestoreCmd is defined
	if strings.TrimSpace(dmm.Destination.RestoreCmd) != "" {
		if err := Restore(dmm); err != nil {
			return &MigrationError{
				Stage: "restore",
				Err:   fmt.Errorf("restore operation failed: %w", err),
			}
		}
	}

	return nil
}
