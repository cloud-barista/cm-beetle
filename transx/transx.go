package transx

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// ============================================================================
// Main Entry Points
// ============================================================================

// Transfer runs the data transfer as defined by the given DataMigrationModel.
// It automatically selects the appropriate transfer strategy based on source/destination types.
func Transfer(dmm DataMigrationModel) error {
	if err := Validate(dmm); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Plan the transfer pipeline
	pipeline, err := Plan(dmm)
	if err != nil {
		return fmt.Errorf("planning failed: %w", err)
	}

	// Execute the pipeline
	return pipeline.Execute()
}

// MigrateData manages the complete data migration workflow:
// 1. If Source.PreCmd is defined, perform pre-processing (e.g., backup)
// 2. Always perform Transfer
// 3. If Destination.PostCmd is defined, perform post-processing (e.g., restore)
func MigrateData(dmm DataMigrationModel) error {
	// Step 1: Pre-processing (optional, e.g., backup)
	if strings.TrimSpace(dmm.Source.PreCmd) != "" {
		if err := executePreCommand(dmm.Source); err != nil {
			return &MigrationError{Stage: StageBackup, Err: err}
		}
	}

	// Step 2: Transfer (core)
	if err := Transfer(dmm); err != nil {
		return &MigrationError{Stage: StageTransfer, Err: err}
	}

	// Step 3: Post-processing (optional, e.g., restore)
	if strings.TrimSpace(dmm.Destination.PostCmd) != "" {
		if err := executePostCommand(dmm.Destination); err != nil {
			return &MigrationError{Stage: StageRestore, Err: err}
		}
	}

	return nil
}

// ============================================================================
// Command Execution
// ============================================================================

// executePreCommand executes the PreCmd defined in the source DataLocation.
func executePreCommand(src DataLocation) error {
	if strings.TrimSpace(src.PreCmd) == "" {
		return fmt.Errorf("pre-command not defined")
	}

	output, err := executeCommand(src.PreCmd, src)
	if err != nil {
		return &OperationError{
			Operation: OperationPreCmd,
			Source:    buildLocationPath(src),
			Command:   src.PreCmd,
			Output:    string(output),
			Err:       err,
		}
	}
	return nil
}

// executePostCommand executes the PostCmd defined in the destination DataLocation.
func executePostCommand(dest DataLocation) error {
	if strings.TrimSpace(dest.PostCmd) == "" {
		return fmt.Errorf("post-command not defined")
	}

	output, err := executeCommand(dest.PostCmd, dest)
	if err != nil {
		return &OperationError{
			Operation:   OperationPostCmd,
			Destination: buildLocationPath(dest),
			Command:     dest.PostCmd,
			Output:      string(output),
			Err:         err,
		}
	}
	return nil
}

// ============================================================================
// Helper Functions
// ============================================================================

// buildLocationPath returns a human-readable path representation for error messages.
func buildLocationPath(loc DataLocation) string {
	switch loc.StorageType {
	case StorageTypeFilesystem:
		if loc.Filesystem != nil && loc.Filesystem.AccessType == AccessTypeSSH && loc.Filesystem.SSH != nil {
			ssh := loc.Filesystem.SSH
			if ssh.Username != "" {
				return fmt.Sprintf("%s@%s:%s", ssh.Username, ssh.Host, loc.Path)
			}
			return fmt.Sprintf("%s:%s", ssh.Host, loc.Path)
		}
		return loc.Path

	case StorageTypeObjectStorage:
		return buildObjectStoragePath(loc)

	default:
		return loc.Path
	}
}

// buildObjectStoragePath returns a readable path for object storage locations.
// Each access type has a prefix for easy identification in error messages.
func buildObjectStoragePath(loc DataLocation) string {
	if loc.ObjectStorage == nil {
		return loc.Path
	}

	switch loc.ObjectStorage.AccessType {
	case AccessTypeMinio:
		// minio: endpoint/bucket/key
		if loc.ObjectStorage.Minio != nil {
			return fmt.Sprintf("minio: %s/%s", loc.ObjectStorage.Minio.Endpoint, loc.Path)
		}
		return fmt.Sprintf("minio: %s", loc.Path)

	case AccessTypeSpider:
		// spider: [connectionName] endpoint/path
		if loc.ObjectStorage.Spider != nil {
			cfg := loc.ObjectStorage.Spider
			return fmt.Sprintf("spider: [%s] %s/%s", cfg.ConnectionName, cfg.Endpoint, loc.Path)
		}
		return fmt.Sprintf("spider: %s", loc.Path)

	case AccessTypeTumblebug:
		// tumblebug: endpoint/ns/{nsId}/os/{osId}/path
		if loc.ObjectStorage.Tumblebug != nil {
			cfg := loc.ObjectStorage.Tumblebug
			return fmt.Sprintf("tumblebug: %s/ns/%s/os/%s/%s", cfg.Endpoint, cfg.NsId, cfg.OsId, loc.Path)
		}
		return fmt.Sprintf("tumblebug: %s", loc.Path)
	}

	return loc.Path
}

// executeCommand executes a command either locally or remotely via SSH.
func executeCommand(command string, loc DataLocation) ([]byte, error) {
	// Object storage doesn't support command execution
	if loc.IsObjectStorage() {
		return nil, fmt.Errorf("command execution not supported for object storage")
	}

	// Filesystem: check if local or SSH
	if loc.Filesystem == nil {
		return nil, fmt.Errorf("filesystem access config required")
	}

	switch loc.Filesystem.AccessType {
	case AccessTypeLocal:
		// Local execution
		cmd := exec.Command("sh", "-c", command)
		return cmd.CombinedOutput()

	case AccessTypeSSH:
		// Remote execution via SSH
		if loc.Filesystem.SSH == nil {
			return nil, fmt.Errorf("SSH config required for remote command execution")
		}
		return executeSSHCommand(command, loc.Filesystem.SSH)

	default:
		return nil, fmt.Errorf("command execution not supported for access type: %s", loc.Filesystem.AccessType)
	}
}

// executeSSHCommand executes a command on a remote server via SSH.
// Authentication priority: PrivateKey > PrivateKeyPath > SSH Agent
func executeSSHCommand(command string, cfg *SSHConfig) ([]byte, error) {
	authMethods, err := buildSSHAuthMethods(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to build SSH auth methods: %w", err)
	}

	if len(authMethods) == 0 {
		return nil, fmt.Errorf("no SSH authentication method available")
	}

	// Build SSH client config
	timeout := time.Duration(cfg.ConnectTimeout) * time.Second
	if timeout == 0 {
		timeout = 30 * time.Second
	}

	sshConfig := &ssh.ClientConfig{
		User:            cfg.Username,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // TODO: Consider proper host key verification
		Timeout:         timeout,
	}

	// Connect to SSH server
	port := cfg.Port
	if port == 0 {
		port = 22
	}
	addr := fmt.Sprintf("%s:%d", cfg.Host, port)

	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", addr, err)
	}
	defer client.Close()

	// Create session and run command
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create SSH session: %w", err)
	}
	defer session.Close()

	return session.CombinedOutput(command)
}

// buildSSHAuthMethods builds SSH authentication methods from config.
// Priority: PrivateKey (string) > PrivateKeyPath (file) > SSH Agent
func buildSSHAuthMethods(cfg *SSHConfig) ([]ssh.AuthMethod, error) {
	var authMethods []ssh.AuthMethod

	// 1. Try PrivateKey string (preferred for injected secrets)
	if strings.TrimSpace(cfg.PrivateKey) != "" {
		// Normalize escaped newlines to actual newlines
		// This handles keys from environment variables or YAML where \n is literal
		keyContent := normalizePrivateKey(cfg.PrivateKey)

		signer, err := ssh.ParsePrivateKey([]byte(keyContent))
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %w", err)
		}
		authMethods = append(authMethods, ssh.PublicKeys(signer))
	}

	// 2. Try PrivateKeyPath (legacy file-based approach)
	if len(authMethods) == 0 && strings.TrimSpace(cfg.PrivateKeyPath) != "" {
		keyBytes, err := os.ReadFile(cfg.PrivateKeyPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read private key file: %w", err)
		}
		signer, err := ssh.ParsePrivateKey(keyBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key file: %w", err)
		}
		authMethods = append(authMethods, ssh.PublicKeys(signer))
	}

	// 3. Try SSH Agent (if explicitly requested or no other auth available)
	if cfg.UseAgent || len(authMethods) == 0 {
		if agentAuth := getSSHAgentAuth(); agentAuth != nil {
			authMethods = append(authMethods, agentAuth)
		}
	}

	return authMethods, nil
}

// getSSHAgentAuth returns SSH agent authentication method if available.
func getSSHAgentAuth() ssh.AuthMethod {
	socket := os.Getenv("SSH_AUTH_SOCK")
	if socket == "" {
		return nil
	}

	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil
	}
	// Note: This connection is not explicitly closed, but it's acceptable
	// as the agent client will handle it during the SSH session lifecycle.

	agentClient := agent.NewClient(conn)
	return ssh.PublicKeysCallback(agentClient.Signers)
}

// normalizePrivateKey converts escaped newlines to actual newlines.
// This handles keys from environment variables or YAML where \n is a literal string.
// Supports both formats:
// - Multi-line (already has real newlines): passed through as-is
// - Single-line with escaped \n: converted to real newlines
func normalizePrivateKey(key string) string {
	// If the key already contains real newlines, return as-is
	if strings.Contains(key, "\n") && !strings.Contains(key, "\\n") {
		return key
	}

	// Convert literal \n to actual newlines
	return strings.ReplaceAll(key, "\\n", "\n")
}

// ============================================================================
// Deprecated: Backward Compatibility Functions
// ============================================================================

// Backup executes the PreCmd defined in the source DataLocation.
// Deprecated: Use executePreCommand directly or set Source.PreCmd and call MigrateData.
func Backup(dmm DataMigrationModel) error {
	return executePreCommand(dmm.Source)
}

// Restore executes the PostCmd defined in the destination DataLocation.
// Deprecated: Use executePostCommand directly or set Destination.PostCmd and call MigrateData.
func Restore(dmm DataMigrationModel) error {
	return executePostCommand(dmm.Destination)
}
