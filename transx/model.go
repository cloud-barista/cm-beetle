package transx

import (
	"fmt"
	"strings"
)

// Transfer method constants
const (
	TransferMethodRsync         = "rsync"
	TransferMethodObjectStorage = "object-storage"
)

// Object Storage client constants
const (
	ObjectStorageClientSpider    = "spider"    // CB-Spider presigned URL API (default)
	ObjectStorageClientMinio     = "minio"     // MinIO SDK for direct S3-compatible access
	ObjectStorageClientTumblebug = "tumblebug" // CB-Tumblebug Object Storage API
)

// FilterOption defines file filtering options for both rsync and Object Storage transfers.
// Supports glob patterns with rsync-like filtering behavior:
// 1. If include patterns are specified, only matching files are included
// 2. Exclude patterns are applied after include patterns (exclude takes priority)
// 3. Supports ** for recursive directory matching
type FilterOption struct {
	Include []string `json:"include,omitempty"` // Patterns to include (e.g., "*.txt", "data/**")
	Exclude []string `json:"exclude,omitempty"` // Patterns to exclude (e.g., "*.log", "temp/**", ".git/**")
}

// DataMigrationModel defines a single data migration task supporting multiple protocols.
type DataMigrationModel struct {
	Source                     EndpointDetails  `json:"source" validate:"required"`                     // Source endpoint configuration
	SourceTransferOptions      *TransferOptions `json:"sourceTransferOptions" validate:"required"`      // Source-specific transfer options
	Destination                EndpointDetails  `json:"destination" validate:"required"`                // Destination endpoint configuration
	DestinationTransferOptions *TransferOptions `json:"destinationTransferOptions" validate:"required"` // Destination-specific transfer options
}

// EndpointDetails defines the source/destination endpoint for data transfer and backup/restore operations.
// Simple unified structure supporting SSH-based rsync, Object Storage API endpoints, and local filesystem transfers.
type EndpointDetails struct {
	// Endpoint configuration (auto-detects protocol based on provided fields)
	Endpoint string `json:"endpoint,omitempty"` // SSH host/IP or Object Storage API endpoint (e.g., "server.com", "http://localhost:1024/spider/s3")
	Port     int    `json:"port,omitempty"`     // Port for SSH host/IP (default: 22) or Object Storage API endpoint (default: 1024)

	// Data location (required)
	DataPath string `json:"dataPath" validate:"required"` // Local path, remote path, or Object Storage bucket path (e.g., "/data", "bucket/object-key")

	// Command execution
	BackupCmd  string `json:"backupCmd,omitempty"`  // Backup command string to be executed on this endpoint
	RestoreCmd string `json:"restoreCmd,omitempty"` // Restore command string to be executed on this endpoint
}

// TransferOptions defines options for various data transfer methods.
type TransferOptions struct {
	// Transfer method specification (required)
	Method string `json:"method" validate:"required"` // Transfer method: "rsync", "object-storage"

	// Rsync-specific options
	RsyncOptions *RsyncOption `json:"rsyncOptions,omitempty"`

	// Object Storage-specific options (CB-Spider, AWS S3, etc.)
	ObjectStorageOptions *ObjectStorageOption `json:"objectStorageOptions,omitempty"`
}

// RsyncOption defines rsync-specific transfer options and SSH connection options.
type RsyncOption struct {
	// SSH connection & authentication options (integrated)
	Username          string `json:"username,omitempty"`          // SSH username
	SSHPrivateKeyPath string `json:"sshPrivateKeyPath,omitempty"` // SSH private key path

	// InsecureSkipHostKeyVerification, if true, relaxes host key checking for SSH connections.
	// Adds "-o StrictHostKeyChecking=accept-new -o UserKnownHostsFile=/dev/null" options.
	// Warning: This can be a security risk and should only be used in trusted environments.
	InsecureSkipHostKeyVerification bool `json:"insecureSkipHostKeyVerification,omitempty" default:"false"`
	ConnectTimeout                  int  `json:"connectTimeout,omitempty" default:"30"` // SSH connection timeout in seconds

	// Transfer behavior options
	Verbose  bool `json:"verbose,omitempty" default:"false"`  // Enable verbose logging
	DryRun   bool `json:"dryRun,omitempty" default:"false"`   // Perform a trial run with no changes made
	Progress bool `json:"progress,omitempty" default:"false"` // Show progress during transfer

	// Rsync-specific options
	Compress  bool          `json:"compress,omitempty" default:"true"`   // -z, --compress: Compress file data during the transfer
	Archive   bool          `json:"archive,omitempty" default:"true"`    // -a, --archive: Archive mode; equals -rlptgoD (no -H,-A,-X)
	Delete    bool          `json:"delete,omitempty" default:"false"`    // --delete: Delete extraneous files from dest dirs
	RsyncPath string        `json:"rsyncPath,omitempty" default:"rsync"` // Path to the rsync executable (if empty, uses system PATH)
	Filter    *FilterOption `json:"filter,omitempty"`                    // File filtering options (include/exclude patterns) - use nested structure for better organization

	// TransferDirContentsOnly, if true, adds a trailing slash to source paths
	// to transfer only the contents of the directory and not the directory itself.
	TransferDirContentsOnly bool `json:"transferDirContentsOnly,omitempty" default:"false"`
}

// ObjectStorageOption defines Object Storage transfer options with client-specific configurations.
// Supports three clients: MinIO (direct S3 access), Spider (CB-Spider API), and Tumblebug (CB-Tumblebug API).
// Only one client configuration should be provided based on the Client field.
type ObjectStorageOption struct {
	// Client selection (required)
	Client string `json:"client" validate:"required,oneof=spider minio tumblebug"` // Object storage client: "spider", "minio", or "tumblebug"

	// Client-specific configurations (provide only one based on Client field)
	MinIOConfig     *MinIOConfig     `json:"minioConfig,omitempty"`     // MinIO/S3 client configuration (for Client="minio")
	SpiderConfig    *SpiderConfig    `json:"spiderConfig,omitempty"`    // CB-Spider client configuration (for Client="spider")
	TumblebugConfig *TumblebugConfig `json:"tumblebugConfig,omitempty"` // CB-Tumblebug client configuration (for Client="tumblebug")

	// Common HTTP request configuration (optional)
	Timeout    int `json:"timeout,omitempty" default:"300"`  // HTTP request timeout in seconds (default: 300)
	MaxRetries int `json:"maxRetries,omitempty" default:"3"` // Maximum number of retry attempts (default: 3)

	// File filtering options (applied after listing objects, before upload/download)
	Filter *FilterOption `json:"filter,omitempty"` // File filtering options (include/exclude patterns)
}

// MinIOConfig defines MinIO/S3-compatible storage client configuration.
// Used for direct access to AWS S3, MinIO, or other S3-compatible storage services.
type MinIOConfig struct {
	Endpoint        string `json:"endpoint" validate:"required"`         // S3-compatible storage endpoint (e.g., "s3.amazonaws.com", "play.min.io:9000")
	AccessKeyId     string `json:"accessKeyId" validate:"required"`      // AWS Access Key ID
	SecretAccessKey string `json:"secretAccessKey" validate:"required"`  // AWS Secret Access Key
	Region          string `json:"region,omitempty" default:"us-east-1"` // AWS region (default: "us-east-1")
	UseSSL          bool   `json:"useSSL,omitempty" default:"true"`      // Use HTTPS for connections (default: true)
}

// SpiderConfig defines CB-Spider Object Storage API client configuration.
// Uses CB-Spider's presigned URL API for object storage operations.
type SpiderConfig struct {
	Endpoint       string `json:"endpoint" validate:"required"`       // CB-Spider API endpoint (e.g., "http://localhost:1024/spider/s3")
	ConnectionName string `json:"connectionName" validate:"required"` // CB-Spider connection name (e.g., "aws-config01", "aws-ap-northeast-2")
	Expires        int    `json:"expires,omitempty" default:"3600"`   // Presigned URL expiration time in seconds (default: 3600)
}

// TumblebugConfig defines CB-Tumblebug Object Storage API client configuration.
// Uses CB-Tumblebug's namespace-aware Object Storage API.
type TumblebugConfig struct {
	Endpoint string `json:"endpoint" validate:"required"`     // CB-Tumblebug API endpoint (e.g., "http://localhost:1323/tumblebug/objectstorage")
	NsId     string `json:"nsId" validate:"required"`         // Namespace ID for multi-tenancy support
	OsId     string `json:"osId" validate:"required"`         // Object Storage ID configured in Tumblebug
	Expires  int    `json:"expires,omitempty" default:"3600"` // Presigned URL expiration time in seconds (default: 3600)
}

// Validate checks if the fields of DataMigrationModel satisfy basic requirements for transfer tasks.
func Validate(dmm DataMigrationModel) error {
	sourceRsyncPath := dmm.Source.GetRsyncPath(nil)    // Basic validation without specific options
	destRsyncPath := dmm.Destination.GetRsyncPath(nil) // Basic validation without specific options

	if strings.TrimSpace(sourceRsyncPath) == "" || strings.TrimSpace(dmm.Source.DataPath) == "" {
		return fmt.Errorf("source path must be provided for transfer task")
	}
	if strings.TrimSpace(destRsyncPath) == "" || strings.TrimSpace(dmm.Destination.DataPath) == "" {
		return fmt.Errorf("destination path must be provided for transfer task")
	}

	// Validate SSH port for source if it's a remote endpoint
	if dmm.Source.IsRemote() {
		sourcePort := dmm.Source.GetPort()
		if sourcePort != 0 && (sourcePort < 1 || sourcePort > 65535) {
			return fmt.Errorf("source SSH port %d is out of valid range (1-65535)", sourcePort)
		}
		if strings.TrimSpace(dmm.Source.GetEndpoint()) == "" {
			return fmt.Errorf("source HostIP must be provided for remote transfer task")
		}
	}
	// Validate SSH port for destination if it's a remote endpoint
	if dmm.Destination.IsRemote() {
		destPort := dmm.Destination.GetPort()
		if destPort != 0 && (destPort < 1 || destPort > 65535) {
			return fmt.Errorf("destination SSH port %d is out of valid range (1-65535)", destPort)
		}
		if strings.TrimSpace(dmm.Destination.GetEndpoint()) == "" {
			return fmt.Errorf("destination HostIP must be provided for remote transfer task")
		}
	}

	// Validate Object Storage configuration
	if dmm.SourceTransferOptions != nil &&
		dmm.SourceTransferOptions.Method == TransferMethodObjectStorage &&
		dmm.SourceTransferOptions.ObjectStorageOptions != nil {
		if err := validateObjectStorageOptions(dmm.SourceTransferOptions.ObjectStorageOptions, "source"); err != nil {
			return err
		}
	}

	if dmm.DestinationTransferOptions != nil &&
		dmm.DestinationTransferOptions.Method == TransferMethodObjectStorage &&
		dmm.DestinationTransferOptions.ObjectStorageOptions != nil {
		if err := validateObjectStorageOptions(dmm.DestinationTransferOptions.ObjectStorageOptions, "destination"); err != nil {
			return err
		}
	}

	return nil
}

// validateObjectStorageOptions validates Object Storage transfer options
func validateObjectStorageOptions(options *ObjectStorageOption, context string) error {
	if strings.TrimSpace(options.Client) == "" {
		return fmt.Errorf("%s Object Storage client must be specified (spider, minio, or tumblebug)", context)
	}

	// Validate client-specific configuration
	switch options.Client {
	case ObjectStorageClientMinio:
		if options.MinIOConfig == nil {
			return fmt.Errorf("%s MinIO client requires minioConfig (provide: endpoint, accessKeyId, secretAccessKey)", context)
		}
		if strings.TrimSpace(options.MinIOConfig.Endpoint) == "" {
			return fmt.Errorf("%s MinIO endpoint must be provided (e.g., \"s3.amazonaws.com\", \"play.min.io:9000\")", context)
		}
		if strings.TrimSpace(options.MinIOConfig.AccessKeyId) == "" {
			return fmt.Errorf("%s MinIO accessKeyId must be provided", context)
		}
		if strings.TrimSpace(options.MinIOConfig.SecretAccessKey) == "" {
			return fmt.Errorf("%s MinIO secretAccessKey must be provided", context)
		}

	case ObjectStorageClientSpider:
		if options.SpiderConfig == nil {
			return fmt.Errorf("%s Spider client requires spiderConfig (provide: endpoint, connectionName)", context)
		}
		if strings.TrimSpace(options.SpiderConfig.Endpoint) == "" {
			return fmt.Errorf("%s Spider endpoint must be provided (e.g., \"http://localhost:1024/spider/s3\")", context)
		}
		if strings.TrimSpace(options.SpiderConfig.ConnectionName) == "" {
			return fmt.Errorf("%s Spider connectionName must be provided (e.g., \"aws-config01\", \"aws-ap-northeast-2\")", context)
		}

	case ObjectStorageClientTumblebug:
		if options.TumblebugConfig == nil {
			return fmt.Errorf("%s Tumblebug client requires tumblebugConfig (provide: endpoint, nsId)", context)
		}
		if strings.TrimSpace(options.TumblebugConfig.Endpoint) == "" {
			return fmt.Errorf("%s Tumblebug endpoint must be provided (e.g., \"http://localhost:1323/tumblebug\")", context)
		}
		if strings.TrimSpace(options.TumblebugConfig.NsId) == "" {
			return fmt.Errorf("%s Tumblebug nsId (namespace ID) must be provided", context)
		}

	default:
		return fmt.Errorf("%s unsupported Object Storage client: %s (must be spider, minio, or tumblebug)", context, options.Client)
	}

	return nil
}
