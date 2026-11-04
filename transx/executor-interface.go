// Package transx provides transfer executors for data migration.
// Supports:
//   - Rsync: Local/remote filesystem transfers with SSH support
//   - S3: Object Storage transfers using presigned URLs
package transx

// Executor defines the interface for transfer operations.
type Executor interface {
	// Execute performs the transfer from source to destination.
	Execute(source, destination DataLocation) error
}

// ============================================================================
// Transfer Mode (Generalized Execution Direction)
// ============================================================================

// TransferMode defines the execution direction of a transfer step across Rsync and S3 executors.
type TransferMode string

const (
	// TransferModePull pulls data from a remote endpoint to local filesystem (remote-ssh -> local, or object-storage -> local).
	TransferModePull TransferMode = "pull"

	// TransferModePush pushes data from local filesystem to a destination (local -> remote-ssh, local -> object-storage, or local -> local).
	TransferModePush TransferMode = "push"

	// TransferModeSSHAgentFwdAuto tries ssh-agent-fwd-push first and falls back to ssh-agent-fwd-pull on pre-flight failure.
	TransferModeSSHAgentFwdAuto TransferMode = "ssh-agent-fwd-auto"

	// TransferModeSSHAgentFwdPush connects to source with SSH agent forwarding and pushes directly to destination (src -> dst).
	TransferModeSSHAgentFwdPush TransferMode = "ssh-agent-fwd-push"

	// TransferModeSSHAgentFwdPull connects to destination with SSH agent forwarding and pulls directly from source (dst <- src).
	TransferModeSSHAgentFwdPull TransferMode = "ssh-agent-fwd-pull"
)
