// Package s3 provides S3-compatible Object Storage providers.
// Supported providers:
//   - Direct: AWS S3, MinIO, and S3-compatible storage (via minio-go SDK)
//   - Spider: via CB-Spider Object Storage API
//   - Tumblebug: via CB-Tumblebug Object Storage API
package s3

import (
	"fmt"
	"strings"
)

// Provider defines the interface for S3-compatible object storage operations.
type Provider interface {
	// GeneratePresignedURL generates a presigned URL for upload or download.
	// action: "upload" or "download"
	// key: object key (file path within bucket)
	GeneratePresignedURL(action, key string) (string, error)

	// ListObjects lists objects with the given prefix.
	ListObjects(prefix string) ([]ObjectInfo, error)

	// GetBucket returns the bucket/container name for this provider.
	GetBucket() string
}

// ObjectInfo represents metadata about a storage object.
type ObjectInfo struct {
	Key          string // Object key (path)
	Size         int64  // Size in bytes
	LastModified string // Last modified timestamp
	ETag         string // Entity tag (hash)
}

// Config types for each provider

// MinioConfig defines configuration for S3-compatible storage access using minio-go SDK.
// Supports: AWS S3, MinIO, Ceph, DigitalOcean Spaces, etc.
type MinioConfig struct {
	Endpoint        string `json:"endpoint" validate:"required"`
	AccessKeyId     string `json:"accessKeyId" validate:"required"`
	SecretAccessKey string `json:"secretAccessKey" validate:"required"`
	Region          string `json:"region,omitempty" default:"us-east-1"`
	UseSSL          bool   `json:"useSSL,omitempty" default:"true"`
}

// SpiderConfig defines CB-Spider Object Storage API client configuration.
type SpiderConfig struct {
	Endpoint       string `json:"endpoint" validate:"required"`
	ConnectionName string `json:"connectionName" validate:"required"`
	Expires        int    `json:"expires,omitempty" default:"3600"`
}

// TumblebugConfig defines CB-Tumblebug Object Storage API client configuration.
type TumblebugConfig struct {
	Endpoint string `json:"endpoint" validate:"required"`
	NsId     string `json:"nsId" validate:"required"`
	OsId     string `json:"osId" validate:"required"`
	Expires  int    `json:"expires,omitempty" default:"3600"`
}

// New creates a Provider based on the provider type and configuration.
func New(providerType string, config interface{}, bucket string) (Provider, error) {
	switch providerType {
	case "minio":
		cfg, ok := config.(*MinioConfig)
		if !ok {
			return nil, fmt.Errorf("invalid config type for minio provider")
		}
		return NewMinioProvider(cfg, bucket)

	case "spider":
		cfg, ok := config.(*SpiderConfig)
		if !ok {
			return nil, fmt.Errorf("invalid config type for spider provider")
		}
		return NewSpiderProvider(cfg, bucket)

	case "tumblebug":
		cfg, ok := config.(*TumblebugConfig)
		if !ok {
			return nil, fmt.Errorf("invalid config type for tumblebug provider")
		}
		return NewTumblebugProvider(cfg)

	default:
		return nil, fmt.Errorf("unsupported provider type: %s", providerType)
	}
}

// ParseBucketAndKey parses the path into bucket and key components.
// Path format: "bucket-name/path/to/object" or "bucket-name/"
func ParseBucketAndKey(path string) (bucket, key string) {
	path = strings.TrimPrefix(path, "/")
	parts := strings.SplitN(path, "/", 2)

	if len(parts) == 0 {
		return "", ""
	}

	bucket = parts[0]
	if len(parts) > 1 {
		key = parts[1]
	}

	return bucket, key
}
