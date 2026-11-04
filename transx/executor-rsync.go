package transx

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

// ============================================================================
// PreFlightError
// ============================================================================

// PreFlightError indicates an inter-server (Primary -> Peer) SSH reachability failure during pre-flight verification.
type PreFlightError struct {
	Mode TransferMode
	Err  error
}

// Error returns the formatted pre-flight error message.
func (e *PreFlightError) Error() string {
	return fmt.Sprintf("pre-flight check failed for mode %s: %v", e.Mode, e.Err)
}

// Unwrap returns the underlying error.
func (e *PreFlightError) Unwrap() error {
	return e.Err
}

// ============================================================================
// RsyncExecutor
// ============================================================================

// RsyncExecutor implements Executor using rsync for file transfers.
type RsyncExecutor struct {
	Mode             TransferMode
	ResolvedMode     TransferMode
	Archive          bool
	Compress         bool
	Partial          bool
	MaxRetries       int
	DeleteExtraneous bool
	DryRun           bool
	Verbose          bool
	AdditionalArgs   []string

	tempKeyFile string
}

// NewRsyncExecutor creates a new RsyncExecutor with automatically determined transfer mode.
func NewRsyncExecutor(src, dst DataLocation) (*RsyncExecutor, error) {
	mode, err := determineTransferMode(src, dst)
	if err != nil {
		return nil, err
	}
	return NewRsyncExecutorFromLocations(src, dst, mode), nil
}

// NewRsyncExecutorWithMode creates a new RsyncExecutor with an explicitly specified transfer mode.
func NewRsyncExecutorWithMode(src, dst DataLocation, mode TransferMode) (*RsyncExecutor, error) {
	srcIsRemote := src.Filesystem != nil && src.Filesystem.SSH != nil
	dstIsRemote := dst.Filesystem != nil && dst.Filesystem.SSH != nil

	switch mode {
	case TransferModeSSHAgentFwdAuto, TransferModeSSHAgentFwdPush, TransferModeSSHAgentFwdPull:
		if !srcIsRemote || !dstIsRemote {
			return nil, fmt.Errorf("transfer mode %q requires both source and destination to be remote SSH endpoints", mode)
		}
	case TransferModePull:
		if !srcIsRemote || dstIsRemote {
			return nil, fmt.Errorf("transfer mode %q requires remote SSH source and local destination", mode)
		}
	case TransferModePush:
		if srcIsRemote {
			return nil, fmt.Errorf("transfer mode %q requires local source", mode)
		}
	default:
		return nil, fmt.Errorf("unsupported transfer mode: %s", mode)
	}

	return NewRsyncExecutorFromLocations(src, dst, mode), nil
}

// NewRsyncExecutorFromLocations initializes RsyncExecutor and merges SSH/rsync options from source and destination.
func NewRsyncExecutorFromLocations(src, dst DataLocation, mode TransferMode) *RsyncExecutor {
	exec := &RsyncExecutor{
		Mode:             mode,
		ResolvedMode:     mode,
		Archive:          true,
		Compress:         true,
		Partial:          true,
		MaxRetries:       3,
		DeleteExtraneous: false,
		DryRun:           false,
		Verbose:          false,
	}

	appliedFirst := false
	for _, loc := range []DataLocation{src, dst} {
		if loc.Filesystem == nil || loc.Filesystem.SSH == nil {
			continue
		}
		cfg := loc.Filesystem.SSH
		exec.DeleteExtraneous = exec.DeleteExtraneous || cfg.Delete
		exec.Verbose = exec.Verbose || cfg.Verbose
		exec.DryRun = exec.DryRun || cfg.DryRun

		if cfg.explicitOptions {
			if !appliedFirst {
				exec.Archive = cfg.Archive
				exec.Compress = cfg.Compress
				exec.Partial = cfg.Partial
				exec.MaxRetries = cfg.MaxRetries
				appliedFirst = true
			} else {
				exec.Archive = exec.Archive && cfg.Archive
				exec.Compress = exec.Compress && cfg.Compress
				exec.Partial = exec.Partial && cfg.Partial
				if cfg.MaxRetries != 3 {
					exec.MaxRetries = cfg.MaxRetries
				}
			}
		} else if cfg.MaxRetries > 0 {
			exec.MaxRetries = cfg.MaxRetries
		}
	}

	return exec
}

// determineTransferMode selects the default transfer mode based on endpoint types.
func determineTransferMode(src, dst DataLocation) (TransferMode, error) {
	srcIsRemote := src.Filesystem != nil && src.Filesystem.SSH != nil
	dstIsRemote := dst.Filesystem != nil && dst.Filesystem.SSH != nil

	switch {
	case srcIsRemote && dstIsRemote:
		return TransferModeSSHAgentFwdAuto, nil
	case srcIsRemote && !dstIsRemote:
		return TransferModePull, nil
	case !srcIsRemote && dstIsRemote:
		return TransferModePush, nil
	default:
		return TransferModePush, nil
	}
}

// Execute performs rsync transfer from source to destination according to Mode.
func (e *RsyncExecutor) Execute(source, destination DataLocation) error {
	switch e.Mode {
	case TransferModePull:
		return e.executePull(source, destination)
	case TransferModePush:
		return e.executePush(source, destination)
	case TransferModeSSHAgentFwdAuto:
		return e.executeSSHAgentFwdAuto(source, destination)
	case TransferModeSSHAgentFwdPush:
		return e.executeSSHAgentFwdPush(source, destination)
	case TransferModeSSHAgentFwdPull:
		return e.executeSSHAgentFwdPull(source, destination)
	default:
		return fmt.Errorf("unknown transfer mode: %s", e.Mode)
	}
}

// runWithRetry executes fn and retries up to MaxRetries times on transient transfer failures (never on PreFlightError).
func (e *RsyncExecutor) runWithRetry(fn func() error) error {
	maxAttempts := e.MaxRetries
	if maxAttempts < 1 {
		maxAttempts = 1
	}

	var lastErr error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		lastErr = fn()
		if lastErr == nil {
			return nil
		}
		var preFlightErr *PreFlightError
		if errors.As(lastErr, &preFlightErr) {
			return lastErr
		}
		if attempt < maxAttempts {
			time.Sleep(time.Duration(attempt) * 3 * time.Second)
		}
	}
	return fmt.Errorf("transfer failed after %d attempt(s): %w", maxAttempts, lastErr)
}

// ============================================================================
// Pull Mode: Remote Source -> Local
// ============================================================================

// executePull pulls data from remote source to local destination.
func (e *RsyncExecutor) executePull(source, destination DataLocation) error {
	defer e.cleanupTempKeyFile()
	e.ResolvedMode = TransferModePull
	e.ensureLocalDirectory(destination)

	args := e.buildLocalRsyncArgs(source, destination)
	return e.runWithRetry(func() error {
		cmd := exec.Command("rsync", args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("rsync pull failed: %w\nOutput: %s", err, string(output))
		}
		return nil
	})
}

// ============================================================================
// Push Mode: Local -> Remote Destination (or Local -> Local)
// ============================================================================

// executePush pushes data from local source to remote or local destination.
func (e *RsyncExecutor) executePush(source, destination DataLocation) error {
	defer e.cleanupTempKeyFile()
	e.ResolvedMode = TransferModePush

	if destination.IsLocal() {
		e.ensureLocalDirectory(destination)
	} else {
		e.ensureRemoteDirectory(destination)
	}

	args := e.buildLocalRsyncArgs(source, destination)
	return e.runWithRetry(func() error {
		cmd := exec.Command("rsync", args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("rsync push failed: %w\nOutput: %s", err, string(output))
		}
		return nil
	})
}

// ============================================================================
// SSH Agent Forwarding Modes: Remote Source <-> Remote Destination
// ============================================================================

// executeSSHAgentFwdAuto tries ssh-agent-fwd-push first and falls back to ssh-agent-fwd-pull on PreFlightError.
func (e *RsyncExecutor) executeSSHAgentFwdAuto(source, destination DataLocation) error {
	pushErr := e.executeSSHAgentFwdPush(source, destination)
	if pushErr == nil {
		e.ResolvedMode = TransferModeSSHAgentFwdPush
		return nil
	}

	var preFlightErr *PreFlightError
	if !errors.As(pushErr, &preFlightErr) {
		return pushErr
	}

	pullErr := e.executeSSHAgentFwdPull(source, destination)
	if pullErr == nil {
		e.ResolvedMode = TransferModeSSHAgentFwdPull
		return nil
	}

	var pullPreFlightErr *PreFlightError
	if errors.As(pullErr, &pullPreFlightErr) {
		return &PreFlightError{
			Mode: TransferModeSSHAgentFwdAuto,
			Err:  fmt.Errorf("both push (%v) and pull (%v) pre-flight checks failed", pushErr, pullErr),
		}
	}
	return pullErr
}

// executeSSHAgentFwdPush connects to source with SSH agent forwarding and pushes directly to destination (src -> dst).
func (e *RsyncExecutor) executeSSHAgentFwdPush(source, destination DataLocation) error {
	if source.Filesystem == nil || source.Filesystem.SSH == nil {
		return fmt.Errorf("source SSH config is required for %s mode", TransferModeSSHAgentFwdPush)
	}
	if destination.Filesystem == nil || destination.Filesystem.SSH == nil {
		return fmt.Errorf("destination SSH config is required for %s mode", TransferModeSSHAgentFwdPush)
	}

	srcSSH := source.Filesystem.SSH
	dstSSH := destination.Filesystem.SSH

	if err := e.verifyAgentFwdPreFlight(srcSSH, dstSSH, TransferModeSSHAgentFwdPush); err != nil {
		return err
	}

	e.ensureRemoteDirectory(destination)
	rsyncCmd := e.buildRemoteRsyncCommand(source, destination, dstSSH)

	if err := e.runWithRetry(func() error {
		return e.runRemoteSSHWithForwardedAgent(srcSSH, dstSSH, rsyncCmd, TransferModeSSHAgentFwdPush)
	}); err != nil {
		return err
	}

	e.ResolvedMode = TransferModeSSHAgentFwdPush
	return nil
}

// executeSSHAgentFwdPull connects to destination with SSH agent forwarding and pulls directly from source (dst <- src).
func (e *RsyncExecutor) executeSSHAgentFwdPull(source, destination DataLocation) error {
	if source.Filesystem == nil || source.Filesystem.SSH == nil {
		return fmt.Errorf("source SSH config is required for %s mode", TransferModeSSHAgentFwdPull)
	}
	if destination.Filesystem == nil || destination.Filesystem.SSH == nil {
		return fmt.Errorf("destination SSH config is required for %s mode", TransferModeSSHAgentFwdPull)
	}

	srcSSH := source.Filesystem.SSH
	dstSSH := destination.Filesystem.SSH

	if err := e.verifyAgentFwdPreFlight(dstSSH, srcSSH, TransferModeSSHAgentFwdPull); err != nil {
		return err
	}

	e.ensureRemoteDirectory(destination)
	rsyncCmd := e.buildRemoteRsyncPullCommand(source, destination, srcSSH)

	if err := e.runWithRetry(func() error {
		return e.runRemoteSSHWithForwardedAgent(dstSSH, srcSSH, rsyncCmd, TransferModeSSHAgentFwdPull)
	}); err != nil {
		return err
	}

	e.ResolvedMode = TransferModeSSHAgentFwdPull
	return nil
}

// verifyAgentFwdPreFlight checks primary SSH connectivity, rsync installation, and Primary -> Peer SSH reachability.
func (e *RsyncExecutor) verifyAgentFwdPreFlight(primarySSH, peerSSH *SSHConfig, mode TransferMode) error {
	peerUser := peerSSH.Username
	if peerUser == "" {
		peerUser = "root"
	}
	peerHost := peerSSH.Host
	if idx := strings.Index(peerHost, ":"); idx != -1 {
		peerHost = peerHost[:idx]
	}

	sshOpts := e.buildRemoteSSHOptions(peerSSH)
	checkCmd := fmt.Sprintf(
		"command -v rsync >/dev/null 2>&1 || exit 10; %s -o BatchMode=yes -o ConnectTimeout=5 %s@%s 'command -v rsync >/dev/null 2>&1 || exit 11'",
		sshOpts, peerUser, peerHost,
	)

	err := e.runRemoteSSHWithForwardedAgent(primarySSH, peerSSH, checkCmd, mode)
	if err == nil {
		return nil
	}

	var exitErr *ssh.ExitError
	if errors.As(err, &exitErr) {
		switch exitErr.ExitStatus() {
		case 10:
			return fmt.Errorf("rsync is not installed on primary server (%s): %w", primarySSH.Host, err)
		case 11:
			return fmt.Errorf("rsync is not installed on peer server (%s): %w", peerSSH.Host, err)
		default:
			return &PreFlightError{Mode: mode, Err: err}
		}
	}

	return err
}

// runRemoteSSHWithForwardedAgent connects to primarySSH with an in-memory agent holding both keys and runs remoteCmd.
func (e *RsyncExecutor) runRemoteSSHWithForwardedAgent(primarySSH, peerSSH *SSHConfig, remoteCmd string, mode TransferMode) error {
	primaryKeyData, err := e.loadPrivateKeyData(primarySSH)
	if err != nil {
		return fmt.Errorf("failed to load primary (%s) private key data: %w", primarySSH.Host, err)
	}

	signer, err := ssh.ParsePrivateKey(primaryKeyData)
	if err != nil {
		return fmt.Errorf("failed to parse primary private key: %w", err)
	}

	primaryRawKey, err := ssh.ParseRawPrivateKey(primaryKeyData)
	if err != nil {
		return fmt.Errorf("failed to parse primary raw private key: %w", err)
	}

	keyring := agent.NewKeyring()
	if err := keyring.Add(agent.AddedKey{PrivateKey: primaryRawKey}); err != nil {
		return fmt.Errorf("failed to add primary key to agent: %w", err)
	}

	if peerSSH.PrivateKey != "" || peerSSH.PrivateKeyPath != "" {
		peerKeyData, err := e.loadPrivateKeyData(peerSSH)
		if err != nil {
			return fmt.Errorf("failed to load peer (%s) private key data: %w", peerSSH.Host, err)
		}
		peerRawKey, err := ssh.ParseRawPrivateKey(peerKeyData)
		if err != nil {
			return fmt.Errorf("failed to parse peer raw private key: %w", err)
		}
		if err := keyring.Add(agent.AddedKey{PrivateKey: peerRawKey}); err != nil {
			return fmt.Errorf("failed to add peer key to agent: %w", err)
		}
	}

	timeout := time.Duration(primarySSH.ConnectTimeout) * time.Second
	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	config := &ssh.ClientConfig{
		User:            primarySSH.Username,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
	}

	primaryAddr := e.formatSSHAddress(primarySSH)
	client, err := ssh.Dial("tcp", primaryAddr, config)
	if err != nil {
		return fmt.Errorf("failed to connect to server (%s): %w", primarySSH.Host, err)
	}
	defer client.Close()

	if err := agent.ForwardToAgent(client, keyring); err != nil {
		return fmt.Errorf("failed to setup agent forwarding on %s: %w", primarySSH.Host, err)
	}

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create SSH session on %s: %w", primarySSH.Host, err)
	}
	defer session.Close()

	if err := agent.RequestAgentForwarding(session); err != nil {
		return fmt.Errorf("failed to request agent forwarding on %s: %w", primarySSH.Host, err)
	}

	output, err := session.CombinedOutput(remoteCmd)
	if err != nil {
		return fmt.Errorf("remote command (%s) failed on %s: %w\nOutput: %s", mode, primarySSH.Host, err, string(output))
	}
	return nil
}

// ensureLocalDirectory creates the local destination directory if it does not exist yet.
func (e *RsyncExecutor) ensureLocalDirectory(destination DataLocation) {
	if e.DryRun || !destination.IsLocal() {
		return
	}
	targetDir := strings.TrimSpace(destination.Path)
	if targetDir == "" {
		return
	}
	if !isDirectoryPath(targetDir) {
		targetDir = filepath.Dir(targetDir)
	}
	_ = os.MkdirAll(targetDir, 0755)
}

// ensureRemoteDirectory creates the destination directory via SSH if it does not exist yet.
func (e *RsyncExecutor) ensureRemoteDirectory(destination DataLocation) {
	if e.DryRun || destination.Filesystem == nil || destination.Filesystem.SSH == nil {
		return
	}
	targetDir := strings.TrimSpace(destination.Path)
	if targetDir == "" {
		return
	}
	if !isDirectoryPath(targetDir) {
		targetDir = filepath.Dir(targetDir)
	}
	escaped := strings.ReplaceAll(targetDir, "'", "'\\''")
	cmd := fmt.Sprintf("mkdir -p '%s' 2>/dev/null || (sudo -n mkdir -p '%s' && sudo -n chown $(id -un):$(id -gn) '%s') || true", escaped, escaped, escaped)
	_, _ = executeSSHCommand(cmd, destination.Filesystem.SSH)
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

	if strings.Contains(host, ":") {
		return host
	}

	if port == 0 {
		port = 22
	}

	return net.JoinHostPort(host, fmt.Sprintf("%d", port))
}

// buildBaseRsyncFlags constructs the common rsync flags based on RsyncExecutor options.
func (e *RsyncExecutor) buildBaseRsyncFlags() []string {
	var flags []string
	if e.Archive {
		flags = append(flags, "-a")
	} else {
		flags = append(flags, "-r")
	}
	if e.Compress {
		flags = append(flags, "-z")
	}
	if e.Partial {
		flags = append(flags, "--partial", "--append-verify")
	}
	flags = append(flags, "--timeout=180")
	if e.DeleteExtraneous {
		flags = append(flags, "--delete")
	}
	if e.DryRun {
		flags = append(flags, "--dry-run")
	}
	if e.Verbose {
		flags = append(flags, "-v")
	}
	return flags
}

// buildRemoteRsyncCommand constructs the rsync command executed on the source server (ssh-agent-fwd-push: src -> dst).
func (e *RsyncExecutor) buildRemoteRsyncCommand(source, destination DataLocation, dstSSH *SSHConfig) string {
	parts := []string{"rsync"}
	parts = append(parts, e.buildBaseRsyncFlags()...)

	sshOpts := e.buildRemoteSSHOptions(dstSSH)
	if sshOpts != "" {
		parts = append(parts, "-e", fmt.Sprintf("'%s'", sshOpts))
	}

	if source.Filter != nil {
		for _, pattern := range source.Filter.Include {
			parts = append(parts, fmt.Sprintf("--include='%s'", pattern))
		}
		for _, pattern := range source.Filter.Exclude {
			parts = append(parts, fmt.Sprintf("--exclude='%s'", pattern))
		}
	}

	parts = append(parts, e.AdditionalArgs...)

	srcPath := source.Path
	if !strings.HasSuffix(srcPath, "/") && isDirectoryPath(srcPath) {
		srcPath += "/"
	}
	parts = append(parts, srcPath)

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

// buildRemoteRsyncPullCommand constructs the rsync command executed on the destination server (ssh-agent-fwd-pull: dst <- src).
func (e *RsyncExecutor) buildRemoteRsyncPullCommand(source, destination DataLocation, srcSSH *SSHConfig) string {
	parts := []string{"rsync"}
	parts = append(parts, e.buildBaseRsyncFlags()...)

	sshOpts := e.buildRemoteSSHOptions(srcSSH)
	if sshOpts != "" {
		parts = append(parts, "-e", fmt.Sprintf("'%s'", sshOpts))
	}

	if source.Filter != nil {
		for _, pattern := range source.Filter.Include {
			parts = append(parts, fmt.Sprintf("--include='%s'", pattern))
		}
		for _, pattern := range source.Filter.Exclude {
			parts = append(parts, fmt.Sprintf("--exclude='%s'", pattern))
		}
	}

	parts = append(parts, e.AdditionalArgs...)

	srcUser := srcSSH.Username
	if srcUser == "" {
		srcUser = "root"
	}
	srcHost := srcSSH.Host
	if idx := strings.Index(srcHost, ":"); idx != -1 {
		srcHost = srcHost[:idx]
	}
	srcPath := source.Path
	if !strings.HasSuffix(srcPath, "/") && isDirectoryPath(srcPath) {
		srcPath += "/"
	}
	parts = append(parts, fmt.Sprintf("%s@%s:%s", srcUser, srcHost, srcPath))
	parts = append(parts, destination.Path)

	return strings.Join(parts, " ")
}

// buildRemoteSSHOptions constructs SSH options for rsync's -e flag when using forwarded agent authentication.
func (e *RsyncExecutor) buildRemoteSSHOptions(peerSSH *SSHConfig) string {
	parts := []string{"ssh"}

	port := peerSSH.Port
	if port == 0 {
		if idx := strings.Index(peerSSH.Host, ":"); idx != -1 {
			fmt.Sscanf(peerSSH.Host[idx+1:], "%d", &port)
		}
	}
	if port > 0 && port != 22 {
		parts = append(parts, fmt.Sprintf("-p %d", port))
	}

	parts = append(parts, "-o StrictHostKeyChecking=no")
	parts = append(parts, "-o UserKnownHostsFile=/dev/null")

	return strings.Join(parts, " ")
}

// ============================================================================
// Common Helper Functions
// ============================================================================

// buildLocalRsyncArgs constructs rsync command arguments for local execution.
func (e *RsyncExecutor) buildLocalRsyncArgs(source, destination DataLocation) []string {
	args := e.buildBaseRsyncFlags()

	sshCmd := e.buildSSHCommand(source, destination)
	if sshCmd != "" {
		args = append(args, "-e", sshCmd)
	}

	args = append(args, e.buildFilterArgs(source)...)
	args = append(args, e.buildFilterArgs(destination)...)
	args = append(args, e.AdditionalArgs...)
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
