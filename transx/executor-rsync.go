package transx

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// ============================================================================
// Transfer Mode
// ============================================================================

// TransferMode defines how rsync transfer is executed.
type TransferMode string

const (
	// TransferModePull pulls data from remote source to local.
	// Direction: remote-source → local
	TransferModePull TransferMode = "pull"

	// TransferModePush pushes data from local to remote destination.
	// Direction: local → remote-destination
	TransferModePush TransferMode = "push"

	// TransferModeAgentForward uses SSH Agent Forwarding to execute rsync
	// on the source server, transferring directly to destination.
	// Direction: remote-source → remote-destination (via source server)
	TransferModeAgentForward TransferMode = "agent-forward"
)

// ============================================================================
// RsyncExecutor
// ============================================================================

// RsyncExecutor implements Executor using rsync for file transfers.
// Supports three transfer modes: Pull, Push, and Agent Forwarding.
type RsyncExecutor struct {
	Mode             TransferMode // Transfer mode (pull, push, agent-forward)
	DeleteExtraneous bool         // --delete: remove extraneous files from destination
	DryRun           bool         // --dry-run: perform trial run without changes
	Verbose          bool         // -v: increase verbosity
	AdditionalArgs   []string     // Additional rsync arguments

	tempKeyFile string // Temporary key file path (for PrivateKey content)
}

// NewRsyncExecutor creates a new RsyncExecutor with automatically determined transfer mode:
//   - SSH → SSH: AgentForward
//   - SSH → Local: Pull
//   - Local → SSH: Push
//   - Local → Local: not supported (returns error)
func NewRsyncExecutor(src, dst DataLocation) (*RsyncExecutor, error) {
	mode, err := determineTransferMode(src, dst)
	if err != nil {
		return nil, err
	}

	exec := &RsyncExecutor{
		Mode:             mode,
		DeleteExtraneous: false,
		DryRun:           false,
		Verbose:          false,
	}

	// Apply SSH options from source
	if src.Filesystem != nil && src.Filesystem.SSH != nil {
		cfg := src.Filesystem.SSH
		exec.DeleteExtraneous = cfg.Delete
		exec.Verbose = cfg.Verbose
		exec.DryRun = cfg.DryRun
	}

	// Merge SSH options from destination (additive)
	if dst.Filesystem != nil && dst.Filesystem.SSH != nil {
		cfg := dst.Filesystem.SSH
		exec.DeleteExtraneous = exec.DeleteExtraneous || cfg.Delete
		exec.Verbose = exec.Verbose || cfg.Verbose
		exec.DryRun = exec.DryRun || cfg.DryRun
	}

	return exec, nil
}

// determineTransferMode selects transfer mode based on endpoint types.
func determineTransferMode(src, dst DataLocation) (TransferMode, error) {
	srcIsRemote := src.Filesystem != nil && src.Filesystem.SSH != nil
	dstIsRemote := dst.Filesystem != nil && dst.Filesystem.SSH != nil

	switch {
	case srcIsRemote && dstIsRemote:
		return TransferModeAgentForward, nil
	case srcIsRemote && !dstIsRemote:
		return TransferModePull, nil
	case !srcIsRemote && dstIsRemote:
		return TransferModePush, nil
	default:
		return "", fmt.Errorf("local-to-local transfer not supported: at least one endpoint must be remote (SSH)")
	}
}

// Execute performs rsync transfer from source to destination.
func (e *RsyncExecutor) Execute(source, destination DataLocation) error {
	switch e.Mode {
	case TransferModePull:
		return e.executePull(source, destination)
	case TransferModePush:
		return e.executePush(source, destination)
	case TransferModeAgentForward:
		return e.executeAgentForward(source, destination)
	default:
		return fmt.Errorf("unknown transfer mode: %s", e.Mode)
	}
}

// ============================================================================
// Pull Mode: Remote Source → Local
// ============================================================================

// executePull pulls data from remote source to local destination.
func (e *RsyncExecutor) executePull(source, destination DataLocation) error {
	// Cleanup temporary key file after execution
	defer e.cleanupTempKeyFile()

	args := e.buildLocalRsyncArgs(source, destination)

	cmd := exec.Command("rsync", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("rsync pull failed: %w\nOutput: %s", err, string(output))
	}

	return nil
}

// ============================================================================
// Push Mode: Local → Remote Destination
// ============================================================================

// executePush pushes data from local source to remote destination.
func (e *RsyncExecutor) executePush(source, destination DataLocation) error {
	// Cleanup temporary key file after execution
	defer e.cleanupTempKeyFile()

	args := e.buildLocalRsyncArgs(source, destination)

	cmd := exec.Command("rsync", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("rsync push failed: %w\nOutput: %s", err, string(output))
	}

	return nil
}

// ============================================================================
// Agent Forward Mode: Remote Source → Remote Destination
// ============================================================================

// executeAgentForward uses SSH Agent Forwarding to execute rsync on source server.
// The source server runs rsync to transfer data directly to destination server.
func (e *RsyncExecutor) executeAgentForward(source, destination DataLocation) error {
	if source.Filesystem == nil || source.Filesystem.SSH == nil {
		return fmt.Errorf("source SSH config is required for agent-forward mode")
	}
	if destination.Filesystem == nil || destination.Filesystem.SSH == nil {
		return fmt.Errorf("destination SSH config is required for agent-forward mode")
	}

	srcSSH := source.Filesystem.SSH
	dstSSH := destination.Filesystem.SSH

	// Load private key data
	keyData, err := e.loadPrivateKeyData(srcSSH)
	if err != nil {
		return fmt.Errorf("failed to load source private key data: %w", err)
	}

	// Create SSH signer
	signer, err := ssh.ParsePrivateKey(keyData)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	// Parse raw private key for agent (supports RSA, ECDSA, Ed25519)
	rawKey, err := ssh.ParseRawPrivateKey(keyData)
	if err != nil {
		return fmt.Errorf("failed to parse raw private key: %w", err)
	}

	// Create in-memory SSH agent with the private key
	keyring := agent.NewKeyring()
	if err := keyring.Add(agent.AddedKey{PrivateKey: rawKey}); err != nil {
		return fmt.Errorf("failed to add key to agent: %w", err)
	}

	// SSH client config for source server
	config := &ssh.ClientConfig{
		User: srcSSH.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to source server
	srcAddr := e.formatSSHAddress(srcSSH)
	client, err := ssh.Dial("tcp", srcAddr, config)
	if err != nil {
		return fmt.Errorf("failed to connect to source server: %w", err)
	}
	defer client.Close()

	// Setup agent forwarding to allow source server to authenticate to destination
	if err := agent.ForwardToAgent(client, keyring); err != nil {
		return fmt.Errorf("failed to setup agent forwarding: %w", err)
	}

	// Create session
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create SSH session: %w", err)
	}
	defer session.Close()

	// Request Agent Forwarding for the session
	if err := agent.RequestAgentForwarding(session); err != nil {
		return fmt.Errorf("failed to request agent forwarding: %w", err)
	}

	// Build rsync command to run on source server
	rsyncCmd := e.buildRemoteRsyncCommand(source, destination, dstSSH)

	// Execute rsync on source server
	output, err := session.CombinedOutput(rsyncCmd)
	if err != nil {
		return fmt.Errorf("rsync agent-forward failed: %w\nOutput: %s", err, string(output))
	}

	return nil
}

// loadPrivateKeyData returns private key bytes from PrivateKey content or PrivateKeyPath file.
func (e *RsyncExecutor) loadPrivateKeyData(cfg *SSHConfig) ([]byte, error) {
	if cfg.PrivateKey != "" {
		return []byte(normalizePrivateKey(cfg.PrivateKey)), nil
	}
	if cfg.PrivateKeyPath != "" {
		return os.ReadFile(cfg.PrivateKeyPath)
	}
	return nil, fmt.Errorf("no private key provided")
}

// formatSSHAddress formats SSH address as host:port.
func (e *RsyncExecutor) formatSSHAddress(sshCfg *SSHConfig) string {
	host := sshCfg.Host
	port := sshCfg.Port

	// Check if host already contains port
	if strings.Contains(host, ":") {
		return host
	}

	if port == 0 {
		port = 22
	}

	return net.JoinHostPort(host, fmt.Sprintf("%d", port))
}

// buildRemoteRsyncCommand constructs the rsync command executed on the source server.
func (e *RsyncExecutor) buildRemoteRsyncCommand(source, destination DataLocation, dstSSH *SSHConfig) string {
	var parts []string
	parts = append(parts, "rsync", "-avz")

	if e.DeleteExtraneous {
		parts = append(parts, "--delete")
	}
	if e.DryRun {
		parts = append(parts, "--dry-run")
	}
	if e.Verbose {
		parts = append(parts, "-v")
	}

	// SSH options for destination connection
	sshOpts := e.buildRemoteSSHOptions(dstSSH)
	if sshOpts != "" {
		parts = append(parts, "-e", fmt.Sprintf("'%s'", sshOpts))
	}

	// Filter arguments
	if source.Filter != nil {
		for _, pattern := range source.Filter.Include {
			parts = append(parts, fmt.Sprintf("--include='%s'", pattern))
		}
		for _, pattern := range source.Filter.Exclude {
			parts = append(parts, fmt.Sprintf("--exclude='%s'", pattern))
		}
	}

	// Additional arguments
	parts = append(parts, e.AdditionalArgs...)

	// Source path (local on source server)
	srcPath := source.Path
	if !strings.HasSuffix(srcPath, "/") && isDirectoryPath(srcPath) {
		srcPath += "/"
	}
	parts = append(parts, srcPath)

	// Destination path (remote from source server's perspective)
	dstUser := dstSSH.Username
	if dstUser == "" {
		dstUser = "root"
	}
	dstHost := dstSSH.Host
	if idx := strings.Index(dstHost, ":"); idx != -1 {
		dstHost = dstHost[:idx]
	}
	dstPath := destination.Path
	parts = append(parts, fmt.Sprintf("%s@%s:%s", dstUser, dstHost, dstPath))

	return strings.Join(parts, " ")
}

// buildRemoteSSHOptions constructs SSH options for rsync's -e flag.
// No identity file is needed; the forwarded agent handles authentication.
func (e *RsyncExecutor) buildRemoteSSHOptions(dstSSH *SSHConfig) string {
	parts := []string{"ssh"}

	// Port
	port := dstSSH.Port
	if port == 0 {
		if idx := strings.Index(dstSSH.Host, ":"); idx != -1 {
			fmt.Sscanf(dstSSH.Host[idx+1:], "%d", &port)
		}
	}
	if port > 0 && port != 22 {
		parts = append(parts, fmt.Sprintf("-p %d", port))
	}

	// [Note] No -i option needed: SSH Agent Forwarding provides authentication
	// The forwarded agent from the host will authenticate to destination

	// Disable strict host key checking
	parts = append(parts, "-o StrictHostKeyChecking=no")
	parts = append(parts, "-o UserKnownHostsFile=/dev/null")

	return strings.Join(parts, " ")
}

// ============================================================================
// Common Helper Functions
// ============================================================================

// buildLocalRsyncArgs constructs rsync command arguments for local execution.
func (e *RsyncExecutor) buildLocalRsyncArgs(source, destination DataLocation) []string {
	args := []string{"-avz"} // archive, verbose, compress

	if e.DeleteExtraneous {
		args = append(args, "--delete")
	}
	if e.DryRun {
		args = append(args, "--dry-run")
	}
	if e.Verbose {
		args = append(args, "-v")
	}

	// SSH options
	sshCmd := e.buildSSHCommand(source, destination)
	if sshCmd != "" {
		args = append(args, "-e", sshCmd)
	}

	// Filter options
	args = append(args, e.buildFilterArgs(source)...)
	args = append(args, e.buildFilterArgs(destination)...)

	// Additional arguments
	args = append(args, e.AdditionalArgs...)

	// Source and destination paths
	args = append(args, e.buildPath(source))
	args = append(args, e.buildPath(destination))

	return args
}

// buildPath returns the rsync path (user@host:path for SSH, or local path).
func (e *RsyncExecutor) buildPath(loc DataLocation) string {
	path := loc.Path

	// Ensure trailing slash for directories
	if !strings.HasSuffix(path, "/") && isDirectoryPath(path) {
		path += "/"
	}

	// Remote path (SSH)
	if loc.Filesystem != nil && loc.Filesystem.SSH != nil {
		cfg := loc.Filesystem.SSH
		user := cfg.Username
		if user == "" {
			user = "root"
		}

		host := cfg.Host
		if idx := strings.Index(host, ":"); idx != -1 {
			host = host[:idx] // Port is handled in -e option
		}
		return fmt.Sprintf("%s@%s:%s", user, host, path)
	}

	// Local path
	return path
}

// buildSSHCommand constructs the SSH command for rsync's -e option.
func (e *RsyncExecutor) buildSSHCommand(source, destination DataLocation) string {
	var cfg *SSHConfig
	switch e.Mode {
	case TransferModePull:
		// Pull: remote source → local, use source SSH config
		if source.Filesystem != nil {
			cfg = source.Filesystem.SSH
		}
	case TransferModePush:
		// Push: local → remote destination, use destination SSH config
		if destination.Filesystem != nil {
			cfg = destination.Filesystem.SSH
		}
	}

	if cfg == nil {
		return ""
	}

	parts := []string{"ssh"}

	port := cfg.Port
	if port == 0 {
		// Extract port from host if specified
		if idx := strings.Index(cfg.Host, ":"); idx != -1 {
			fmt.Sscanf(cfg.Host[idx+1:], "%d", &port)
		}
	}
	if port > 0 && port != 22 {
		parts = append(parts, fmt.Sprintf("-p %d", port))
	}

	// Priority: PrivateKey > PrivateKeyPath
	if cfg.PrivateKey != "" {
		// Create temporary file for PrivateKey content
		tempPath, err := e.createTempKeyFile(cfg.PrivateKey)
		if err == nil {
			parts = append(parts, fmt.Sprintf("-i %s", tempPath))
		}
	} else if cfg.PrivateKeyPath != "" {
		parts = append(parts, fmt.Sprintf("-i %s", cfg.PrivateKeyPath))
	}

	// Disable strict host key checking for automation
	parts = append(parts, "-o StrictHostKeyChecking=no", "-o UserKnownHostsFile=/dev/null")

	return strings.Join(parts, " ")
}

// buildFilterArgs constructs rsync filter arguments from Filter.
func (e *RsyncExecutor) buildFilterArgs(loc DataLocation) []string {
	if loc.Filter == nil {
		return nil
	}

	var args []string

	for _, pattern := range loc.Filter.Include {
		args = append(args, "--include="+pattern)
	}
	for _, pattern := range loc.Filter.Exclude {
		args = append(args, "--exclude="+pattern)
	}

	return args
}

// isDirectoryPath returns true if path appears to be a directory (no file extension).
func isDirectoryPath(path string) bool {
	// Heuristic: paths without extension are likely directories
	base := path
	if idx := strings.LastIndex(path, "/"); idx != -1 {
		base = path[idx+1:]
	}
	return !strings.Contains(base, ".")
}

// createTempKeyFile writes PrivateKey content to a temporary file with secure permissions.
// Returns the temporary file path. Caller must call cleanupTempKeyFile() after use.
// Uses transx-specific directory under system temp for isolation and easy cleanup.
func (e *RsyncExecutor) createTempKeyFile(privateKey string) (string, error) {
	// Use transx-specific directory under system temp (cross-platform)
	transxTempDir := filepath.Join(os.TempDir(), "transx")

	// Create directory with secure permissions (owner only)
	if err := os.MkdirAll(transxTempDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create transx temp directory: %w", err)
	}

	// Pattern includes PID for debugging; CreateTemp adds random suffix for uniqueness
	pattern := fmt.Sprintf("key-%d-*", os.Getpid())
	tmpFile, err := os.CreateTemp(transxTempDir, pattern)
	if err != nil {
		return "", fmt.Errorf("failed to create temp key file: %w", err)
	}

	// Set secure permissions (owner read-only)
	if err := tmpFile.Chmod(0600); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to set temp key file permissions: %w", err)
	}

	// Write normalized private key content
	normalizedKey := normalizePrivateKey(privateKey)
	if _, err := tmpFile.WriteString(normalizedKey); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to write temp key file: %w", err)
	}

	if err := tmpFile.Close(); err != nil {
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to close temp key file: %w", err)
	}

	e.tempKeyFile = tmpFile.Name()
	return e.tempKeyFile, nil
}

// cleanupTempKeyFile removes the temporary key file if it exists.
func (e *RsyncExecutor) cleanupTempKeyFile() {
	if e.tempKeyFile != "" {
		os.Remove(e.tempKeyFile)
		e.tempKeyFile = ""
	}
}
