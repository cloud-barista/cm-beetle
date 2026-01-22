package transx

import (
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
	// StrategyAuto automatically selects the best transfer method.
	StrategyAuto = "auto"

	// StrategyDirect forces direct transfer (e.g., SSH agent forwarding).
	StrategyDirect = "direct"

	// StrategyRelay forces relay via local machine.
	StrategyRelay = "relay"
)

// ============================================================================
// Pipeline and Step Names (for consistent naming)
// ============================================================================

const (
	PipelineFilesystemTransfer    = "filesystem-transfer"
	PipelineObjectStorageTransfer = "objectstorage-transfer"
	PipelineCrossStorageTransfer  = "cross-storage-transfer"

	StepRsyncTransfer   = "rsync-transfer"
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

	// Strategy determines how the transfer is orchestrated.
	// "auto": Automatically select best method.
	// "direct": Force direct transfer (e.g., SSH agent forwarding).
	// "relay": Force relay via local machine.
	Strategy string `json:"strategy,omitempty" default:"auto" validate:"omitempty,oneof=auto direct relay"`
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

	// Hooks for pre/post processing
	PreCmd  string `json:"preCmd,omitempty"`  // Command to run before transfer (source only)
	PostCmd string `json:"postCmd,omitempty"` // Command to run after transfer (destination only)
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

	// Authentication (priority: PrivateKey > PrivateKeyPath > Agent > none)
	// At least one authentication method should be available.
	//
	// PrivateKey: PEM-encoded private key content (preferred for injected secrets).
	// In JSON, use single line with \n for newlines:
	//   "privateKey": "-----BEGIN RSA PRIVATE KEY-----\nMIIE...\n-----END RSA PRIVATE KEY-----"
	PrivateKey     string `json:"privateKey,omitempty"`
	PrivateKeyPath string `json:"privateKeyPath,omitempty"` // Path to private key file (legacy, prefer PrivateKey)
	UseAgent       bool   `json:"useAgent,omitempty"`       // Use SSH agent for authentication (supports agent forwarding)

	// Rsync options
	Archive  bool `json:"archive,omitempty" default:"true"`
	Compress bool `json:"compress,omitempty" default:"true"`
	Delete   bool `json:"delete,omitempty"`
	Verbose  bool `json:"verbose,omitempty"`
	DryRun   bool `json:"dryRun,omitempty"`
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
type SpiderConfig struct {
	Endpoint       string `json:"endpoint" validate:"required"`
	ConnectionName string `json:"connectionName" validate:"required"`
	Expires        int    `json:"expires,omitempty" default:"3600"`
}

// TumblebugConfig defines CB-Tumblebug Object Storage API configuration.
type TumblebugConfig struct {
	Endpoint string `json:"endpoint" validate:"required"`
	NsId     string `json:"nsId" validate:"required"`
	OsId     string `json:"osId" validate:"required"`
	Expires  int    `json:"expires,omitempty" default:"3600"`
}

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
