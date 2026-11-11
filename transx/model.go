package transx

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ============================================================================
// Storage Types: What kind of storage is being used
// ============================================================================

const (
	// StorageTypeFilesystem represents local or remote filesystem storage.
	StorageTypeFilesystem = "filesystem"

	// StorageTypeObjectStorage represents S3-compatible object storage.
	StorageTypeObjectStorage = "objectstorage"
)

// ============================================================================
// Access Types: How to access the storage
// ============================================================================

// Filesystem access types
const (
	// AccessTypeLocal represents local filesystem access (no network).
	AccessTypeLocal = "local"

	// AccessTypeSSH represents remote filesystem access via SSH/rsync.
	AccessTypeSSH = "ssh"
)

// Object Storage access types
const (
	// AccessTypeMinio represents direct S3 SDK access using minio-go.
	AccessTypeMinio = "minio"

	// AccessTypeSpider represents access via CB-Spider Object Storage API.
	AccessTypeSpider = "spider"

	// AccessTypeTumblebug represents access via CB-Tumblebug Object Storage API.
	AccessTypeTumblebug = "tumblebug"
)

// ============================================================================
// Transfer Strategies
// ============================================================================

const (
	// StrategyAuto automatically selects the optimal route (SSH -> SSH: Push -> Pull -> Relay).
	StrategyAuto = "auto"

	// StrategyDirect enforces direct transfer without local staging (SSH -> SSH: Push -> Pull).
	StrategyDirect = "direct"

	// StrategyRelay enforces a 2-step relay via local staging (/tmp/transx-staging/<subdir>).
	StrategyRelay = "relay"

	// StrategySSHAgentFwdPush enforces SSH agent forwarding Push only (Source -> Destination).
	StrategySSHAgentFwdPush = "ssh-agent-fwd-push"

	// StrategySSHAgentFwdPull enforces SSH agent forwarding Pull only (Destination <- Source).
	StrategySSHAgentFwdPull = "ssh-agent-fwd-pull"
)

// ============================================================================
// Pipeline and Step Names (for consistent naming)
// ============================================================================

const (
	PipelineFilesystemTransfer    = "filesystem-transfer"
	PipelineObjectStorageTransfer = "objectstorage-transfer"
	PipelineCrossStorageTransfer  = "cross-storage-transfer"

	StepRsyncTransfer   = "rsync-transfer"
	StepPullToStaging   = "pull-to-staging"
	StepPushFromStaging = "push-from-staging"
	StepDownloadFromS3  = "download-from-s3"
	StepUploadToS3      = "upload-to-s3"
	StepRsyncFromServer = "rsync-from-server"
	StepRsyncToServer   = "rsync-to-server"
)

// ============================================================================
// Staging Configuration
// ============================================================================

const (
	// DefaultStagingPath is the default local staging directory for relay transfers.
	DefaultStagingPath = "/tmp/transx-staging"
)

// ============================================================================
// Filter Options
// ============================================================================

// FilterOption defines file filtering options for transfers.
type FilterOption struct {
	Include []string `json:"include,omitempty"` // Patterns to include (e.g., "*.txt", "data/**")
	Exclude []string `json:"exclude,omitempty"` // Patterns to exclude (e.g., "*.log", "temp/**")
}

// ============================================================================
// Data Migration Model
// ============================================================================

// DataMigrationModel defines a single data migration task.
type DataMigrationModel struct {
	Source      DataLocation `json:"source" validate:"required"`
	Destination DataLocation `json:"destination" validate:"required"`

	// Strategy determines how the transfer is orchestrated (auto, direct, relay, ssh-agent-fwd-push, ssh-agent-fwd-pull).
	Strategy string `json:"strategy,omitempty" default:"auto" validate:"omitempty,oneof=auto direct relay ssh-agent-fwd-push ssh-agent-fwd-pull"`

	// EncryptionKeyID indicates that sensitive fields are encrypted with the specified one-time key.
	EncryptionKeyID string `json:"encryptionKeyId,omitempty"`
}

// IsEncrypted returns true if the model has encrypted sensitive fields.
func (m DataMigrationModel) IsEncrypted() bool {
	return m.EncryptionKeyID != ""
}

// ============================================================================
// Data Location (Unified Structure)
// ============================================================================

// DataLocation defines any data location with separated storage type and access method.
type DataLocation struct {
	// StorageType: What kind of storage
	// "filesystem": Local or remote filesystem
	// "objectstorage": S3-compatible object storage
	StorageType string `json:"storageType" validate:"required,oneof=filesystem objectstorage"`

	// Path to the data
	// For Filesystem: File path (e.g., "/data", "/home/user/data")
	// For ObjectStorage: Bucket/Key (e.g., "my-bucket/my-key")
	Path string `json:"path" validate:"required"`

	// Access configuration (one of the following based on StorageType)
	Filesystem    *FilesystemAccess    `json:"filesystem,omitempty"`    // For storageType="filesystem"
	ObjectStorage *ObjectStorageAccess `json:"objectStorage,omitempty"` // For storageType="objectstorage"

	// Filter defines file filtering options
	Filter *FilterOption `json:"filter,omitempty"`

	// Hooks for pre/post processing on source and destination endpoints
	PreCmd  string `json:"preCmd,omitempty"`  // Command to run before transfer
	PostCmd string `json:"postCmd,omitempty"` // Command to run after transfer
}

// MaxHookOutputBytes caps each captured hook output to keep API responses bounded.
const MaxHookOutputBytes = 8192

// HookOutputs holds truncated combined stdout/stderr captured from pre/post commands.
type HookOutputs struct {
	SourcePreCmd       string `json:"sourcePreCmd,omitempty"`
	DestinationPreCmd  string `json:"destinationPreCmd,omitempty"`
	SourcePostCmd      string `json:"sourcePostCmd,omitempty"`
	DestinationPostCmd string `json:"destinationPostCmd,omitempty"`
}

// IsEmpty returns true when no hook command produced output.
func (h HookOutputs) IsEmpty() bool {
	return h.SourcePreCmd == "" &&
		h.DestinationPreCmd == "" &&
		h.SourcePostCmd == "" &&
		h.DestinationPostCmd == ""
}

// MigrationResult summarizes a completed or failed data migration run.
type MigrationResult struct {
	Message           string       `json:"message"`
	ElapsedTime       string       `json:"elapsedTime"`
	RequestedStrategy string       `json:"requestedStrategy,omitempty"`
	EffectiveStrategy string       `json:"effectiveStrategy,omitempty"`
	HookOutputs       *HookOutputs `json:"hookOutputs,omitempty"`
	SourceRecovered   bool         `json:"sourceRecovered,omitempty"`
}

// PreserveOnError reports whether partial hook outputs or recovery state should be kept on error.
func (r MigrationResult) PreserveOnError() bool {
	return r.HookOutputs != nil || r.SourceRecovered
}

// ============================================================================
// Filesystem Access Configuration
// ============================================================================

// FilesystemAccess defines how to access filesystem storage.
type FilesystemAccess struct {
	// AccessType: How to access the filesystem
	// "local": Local filesystem (no network)
	// "ssh": Remote filesystem via SSH
	AccessType string `json:"accessType" validate:"required,oneof=local ssh"`

	// SSH configuration (required when accessType="ssh")
	SSH *SSHConfig `json:"ssh,omitempty"`
}

// SSHConfig defines SSH connection details and rsync options.
type SSHConfig struct {
	// Connection details
	Host           string `json:"host" validate:"required"`
	Port           int    `json:"port,omitempty" default:"22"`
	Username       string `json:"username" validate:"required"`
	ConnectTimeout int    `json:"connectTimeout,omitempty" default:"30"`

	// PEM-encoded private key content (preferred for injected secrets)
	PrivateKey string `json:"privateKey,omitempty"`

	// Path to private key file (legacy, prefer PrivateKey)
	PrivateKeyPath string `json:"privateKeyPath,omitempty"`

	// Use SSH agent for authentication (supports agent forwarding)
	UseAgent bool `json:"useAgent,omitempty"`

	// Rsync options
	Archive    bool `json:"archive,omitempty" default:"true"`
	Compress   bool `json:"compress,omitempty" default:"true"`
	Partial    bool `json:"partial,omitempty" default:"true"`
	MaxRetries int  `json:"maxRetries,omitempty" default:"3"`
	Delete     bool `json:"delete,omitempty"`
	Verbose    bool `json:"verbose,omitempty"`
	DryRun     bool `json:"dryRun,omitempty"`

	explicitOptions bool
}

// UnmarshalJSON pre-populates default SSH and rsync options before decoding JSON.
func (c *SSHConfig) UnmarshalJSON(data []byte) error {
	type sshConfigAlias SSHConfig
	alias := sshConfigAlias{
		Port:            22,
		ConnectTimeout:  30,
		Archive:         true,
		Compress:        true,
		Partial:         true,
		MaxRetries:      3,
		explicitOptions: true,
	}
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	*c = SSHConfig(alias)
	return nil
}

// ============================================================================
// Object Storage Access Configuration
// ============================================================================

// ObjectStorageAccess defines how to access object storage.
type ObjectStorageAccess struct {
	// AccessType: How to access object storage
	// "minio": Direct S3 SDK access using minio-go
	// "spider": Via CB-Spider Object Storage API
	// "tumblebug": Via CB-Tumblebug Object Storage API
	AccessType string `json:"accessType" validate:"required,oneof=minio spider tumblebug"`

	// Provider-specific configurations (one required based on accessType)
	Minio     *S3MinioConfig   `json:"minio,omitempty"`     // For accessType="minio"
	Spider    *SpiderConfig    `json:"spider,omitempty"`    // For accessType="spider"
	Tumblebug *TumblebugConfig `json:"tumblebug,omitempty"` // For accessType="tumblebug"
}

// S3MinioConfig defines S3 SDK configuration using minio-go.
type S3MinioConfig struct {
	Endpoint        string `json:"endpoint" validate:"required"`
	AccessKeyId     string `json:"accessKeyId" validate:"required"`
	SecretAccessKey string `json:"secretAccessKey" validate:"required"`
	Region          string `json:"region,omitempty" default:"us-east-1"`
	UseSSL          bool   `json:"useSSL,omitempty" default:"true"`
}

// SpiderConfig defines CB-Spider Object Storage API configuration.
// Endpoint should include /spider prefix (e.g., "http://localhost:1024/spider").
type SpiderConfig struct {
	Endpoint       string      `json:"endpoint" validate:"required"`       // CB-Spider API base URL (e.g., "http://localhost:1024/spider")
	ConnectionName string      `json:"connectionName" validate:"required"` // CB-Spider connection name (e.g., "aws-connection")
	Expires        int         `json:"expires,omitempty" default:"3600"`   // Presigned URL expiration in seconds
	Auth           *AuthConfig `json:"auth,omitempty"`                     // Optional authentication configuration
}

// TumblebugConfig defines CB-Tumblebug Object Storage API configuration.
// Endpoint should include /tumblebug prefix (e.g., "http://localhost:1323/tumblebug").
type TumblebugConfig struct {
	Endpoint string      `json:"endpoint" validate:"required"`     // CB-Tumblebug API base URL (e.g., "http://localhost:1323/tumblebug")
	NsId     string      `json:"nsId" validate:"required"`         // Namespace ID for multi-tenancy
	OsId     string      `json:"osId" validate:"required"`         // Object Storage ID
	Expires  int         `json:"expires,omitempty" default:"3600"` // Presigned URL expiration in seconds
	Auth     *AuthConfig `json:"auth,omitempty"`                   // Optional authentication configuration
}

// ============================================================================
// Authentication Configuration
// ============================================================================

// Auth types
const (
	// AuthTypeBasic represents HTTP Basic Authentication.
	AuthTypeBasic = "basic"

	// AuthTypeJWT represents JWT (JSON Web Token) Authentication.
	AuthTypeJWT = "jwt"

	// AuthTypeAPIKey represents API Key Authentication.
	// AuthTypeAPIKey = "apikey" // TODO: Not tested yet

	// AuthTypeOAuth represents OAuth 2.0 Authentication.
	// AuthTypeOAuth = "oauth" // TODO: Not tested yet
)

// AuthConfig defines authentication configuration.
// Use authType to specify the authentication method.
type AuthConfig struct {
	AuthType string           `json:"authType" validate:"required"` // "basic", "jwt" ("apikey", "oauth" not tested yet)
	Basic    *BasicAuthConfig `json:"basic,omitempty"`              // For authType="basic"
	JWT      *JWTAuthConfig   `json:"jwt,omitempty"`                // For authType="jwt"
	// APIKey   *APIKeyConfig    `json:"apiKey,omitempty"`             // For authType="apikey" (TODO: Not tested yet)
	// OAuth    *OAuthConfig     `json:"oauth,omitempty"`              // For authType="oauth" (TODO: Not tested yet)
}

// BasicAuthConfig defines HTTP Basic Authentication credentials.
type BasicAuthConfig struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// JWTAuthConfig defines JWT authentication configuration.
type JWTAuthConfig struct {
	Token string `json:"token" validate:"required"`
}

// APIKeyConfig defines API Key authentication configuration.
// TODO: Not tested yet
// type APIKeyConfig struct {
// 	Key        string `json:"key" validate:"required"`
// 	HeaderName string `json:"headerName,omitempty" default:"X-API-Key"` // Header name for API key
// }

// OAuthConfig defines OAuth 2.0 authentication configuration.
// TODO: Not tested yet
// type OAuthConfig struct {
// 	AccessToken string `json:"accessToken" validate:"required"`
// 	TokenType   string `json:"tokenType,omitempty" default:"Bearer"` // e.g., "Bearer"
// }

// ============================================================================
// Validation
// ============================================================================

// Validate checks if DataMigrationModel satisfies requirements.
func Validate(dmm DataMigrationModel) error {
	if err := validateLocation(dmm.Source, "source"); err != nil {
		return err
	}
	if err := validateLocation(dmm.Destination, "destination"); err != nil {
		return err
	}
	if err := validateStrategy(dmm); err != nil {
		return err
	}
	return nil
}

// validateStrategy verifies that the requested strategy is compatible with the source and destination endpoints.
func validateStrategy(dmm DataMigrationModel) error {
	strategy := strings.TrimSpace(dmm.Strategy)
	if strategy == "" {
		strategy = StrategyAuto
	}

	switch strategy {
	case StrategyAuto, StrategyDirect, StrategyRelay, StrategySSHAgentFwdPush, StrategySSHAgentFwdPull:
	default:
		return fmt.Errorf("strategy: unsupported strategy %q (allowed: %s, %s, %s, %s, %s)",
			dmm.Strategy, StrategyAuto, StrategyDirect, StrategyRelay, StrategySSHAgentFwdPush, StrategySSHAgentFwdPull)
	}

	srcIsSSH := dmm.Source.IsFilesystem() && dmm.Source.Filesystem != nil && dmm.Source.Filesystem.AccessType == AccessTypeSSH
	dstIsSSH := dmm.Destination.IsFilesystem() && dmm.Destination.Filesystem != nil && dmm.Destination.Filesystem.AccessType == AccessTypeSSH
	srcIsLocal := dmm.Source.IsLocal()
	dstIsLocal := dmm.Destination.IsLocal()
	srcIsOS := dmm.Source.IsObjectStorage()
	dstIsOS := dmm.Destination.IsObjectStorage()

	if strategy == StrategySSHAgentFwdPush || strategy == StrategySSHAgentFwdPull {
		if !srcIsSSH || !dstIsSSH {
			return fmt.Errorf("strategy %q is only supported when both source and destination are remote SSH filesystems", strategy)
		}
		return nil
	}

	if srcIsSSH && dstIsSSH {
		return nil
	}

	if (srcIsLocal && dstIsSSH) || (srcIsSSH && dstIsLocal) || (srcIsLocal && dstIsLocal) {
		if strategy == StrategyRelay {
			return fmt.Errorf("strategy %q is not applicable when a local filesystem endpoint is involved (use %q or %q)", strategy, StrategyAuto, StrategyDirect)
		}
		return nil
	}

	if (srcIsLocal && dstIsOS) || (srcIsOS && dstIsLocal) {
		if strategy == StrategyRelay {
			return fmt.Errorf("strategy %q is not applicable for local-to-objectstorage or objectstorage-to-local transfer (use %q or %q)", strategy, StrategyAuto, StrategyDirect)
		}
		return nil
	}

	if (srcIsOS && dstIsOS) || (srcIsSSH && dstIsOS) || (srcIsOS && dstIsSSH) {
		if strategy == StrategyDirect {
			return fmt.Errorf("strategy %q is not supported for remote transfers involving object storage (local staging via %q or %q is required)", strategy, StrategyAuto, StrategyRelay)
		}
		return nil
	}

	return nil
}

func validateLocation(loc DataLocation, context string) error {
	if strings.TrimSpace(loc.Path) == "" {
		return fmt.Errorf("%s: path is required", context)
	}

	switch loc.StorageType {
	case StorageTypeFilesystem:
		return validateFilesystemAccess(loc.Filesystem, context)

	case StorageTypeObjectStorage:
		return validateObjectStorageAccess(loc.ObjectStorage, context)

	default:
		return fmt.Errorf("%s: unsupported storage type: %s", context, loc.StorageType)
	}
}

func validateFilesystemAccess(fs *FilesystemAccess, context string) error {
	if fs == nil {
		return fmt.Errorf("%s: filesystem access config required", context)
	}

	switch fs.AccessType {
	case AccessTypeLocal:
		// No additional config needed
		return nil

	case AccessTypeSSH:
		if fs.SSH == nil {
			return fmt.Errorf("%s: SSH config required for ssh access", context)
		}
		if strings.TrimSpace(fs.SSH.Host) == "" {
			return fmt.Errorf("%s: SSH host is required", context)
		}
		if strings.TrimSpace(fs.SSH.Username) == "" {
			return fmt.Errorf("%s: SSH username is required", context)
		}
		if fs.SSH.Port < 0 || fs.SSH.Port > 65535 {
			return fmt.Errorf("%s: SSH port out of range", context)
		}
		return nil

	default:
		return fmt.Errorf("%s: unsupported filesystem access type: %s", context, fs.AccessType)
	}
}

func validateObjectStorageAccess(os *ObjectStorageAccess, context string) error {
	if os == nil {
		return fmt.Errorf("%s: object storage access config required", context)
	}

	switch os.AccessType {
	case AccessTypeMinio:
		if os.Minio == nil {
			return fmt.Errorf("%s: minio S3 config required", context)
		}
		if strings.TrimSpace(os.Minio.Endpoint) == "" {
			return fmt.Errorf("%s: S3 endpoint is required", context)
		}
		if strings.TrimSpace(os.Minio.AccessKeyId) == "" || strings.TrimSpace(os.Minio.SecretAccessKey) == "" {
			return fmt.Errorf("%s: S3 credentials are required", context)
		}
		return nil

	case AccessTypeSpider:
		if os.Spider == nil {
			return fmt.Errorf("%s: Spider config required", context)
		}
		if strings.TrimSpace(os.Spider.Endpoint) == "" {
			return fmt.Errorf("%s: Spider endpoint is required", context)
		}
		if strings.TrimSpace(os.Spider.ConnectionName) == "" {
			return fmt.Errorf("%s: Spider connection name is required", context)
		}
		return nil

	case AccessTypeTumblebug:
		if os.Tumblebug == nil {
			return fmt.Errorf("%s: Tumblebug config required", context)
		}
		if strings.TrimSpace(os.Tumblebug.Endpoint) == "" {
			return fmt.Errorf("%s: Tumblebug endpoint is required", context)
		}
		if strings.TrimSpace(os.Tumblebug.NsId) == "" {
			return fmt.Errorf("%s: Tumblebug namespace ID is required", context)
		}
		if strings.TrimSpace(os.Tumblebug.OsId) == "" {
			return fmt.Errorf("%s: Tumblebug object storage ID is required", context)
		}
		return nil

	default:
		return fmt.Errorf("%s: unsupported object storage access type: %s", context, os.AccessType)
	}
}

// ============================================================================
// Helper Functions
// ============================================================================

// IsFilesystem returns true if the location uses filesystem storage.
func (loc DataLocation) IsFilesystem() bool {
	return loc.StorageType == StorageTypeFilesystem
}

// IsObjectStorage returns true if the location uses object storage.
func (loc DataLocation) IsObjectStorage() bool {
	return loc.StorageType == StorageTypeObjectStorage
}

// IsLocal returns true if the location is local filesystem.
func (loc DataLocation) IsLocal() bool {
	return loc.IsFilesystem() && loc.Filesystem != nil && loc.Filesystem.AccessType == AccessTypeLocal
}

// IsRemote returns true if the location requires network access.
func (loc DataLocation) IsRemote() bool {
	if loc.IsObjectStorage() {
		return true
	}
	return loc.IsFilesystem() && loc.Filesystem != nil && loc.Filesystem.AccessType == AccessTypeSSH
}

// NeedsLocalStaging returns true if this location needs local staging for relay.
func (loc DataLocation) NeedsLocalStaging() bool {
	return loc.IsRemote()
}
