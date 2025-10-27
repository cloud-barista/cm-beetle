package transx

import "fmt"

// Custom Error Types for Error-Only Approach

// MigrationError represents an error during the migration process
type MigrationError struct {
	Stage string // "backup", "transfer", or "restore"
	Err   error
}

func (e *MigrationError) Error() string {
	return fmt.Sprintf("migration failed at %s stage: %v", e.Stage, e.Err)
}

func (e *MigrationError) Unwrap() error {
	return e.Err
}

// OperationError provides detailed context about transx operation failures
// This unified error type handles backup, restore, and transfer operations
type OperationError struct {
	Operation   string            // "backup", "restore", "transfer"
	Method      string            // transfer method (for transfer operations)
	Source      string            // source path/endpoint
	Destination string            // destination path/endpoint
	Command     string            // executed command (for backup/restore)
	Output      string            // command output (for backup/restore)
	IsRelayMode bool              // relay mode flag (for transfer)
	Context     map[string]string // additional context information
	Err         error             // underlying error
}

func (e *OperationError) Error() string {
	switch e.Operation {
	case "backup":
		return fmt.Sprintf("backup operation failed for source '%s': %v", e.Source, e.Err)
	case "restore":
		return fmt.Sprintf("restore operation failed for destination '%s': %v", e.Destination, e.Err)
	case "transfer":
		modeStr := "direct"
		if e.IsRelayMode {
			modeStr = "relay"
		}
		return fmt.Sprintf("transfer failed (%s, %s mode): %s → %s: %v", e.Method, modeStr, e.Source, e.Destination, e.Err)
	default:
		return fmt.Sprintf("operation '%s' failed: %v", e.Operation, e.Err)
	}
}

func (e *OperationError) Unwrap() error {
	return e.Err
}

// GetOutput returns the command output for debugging (applicable to backup/restore operations)
func (e *OperationError) GetOutput() string {
	return e.Output
}

// GetMethod returns the transfer method (applicable to transfer operations)
func (e *OperationError) GetMethod() string {
	return e.Method
}

// IsOperation checks if the error is for a specific operation type
func (e *OperationError) IsOperation(operation string) bool {
	return e.Operation == operation
}
