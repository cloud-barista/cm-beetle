// Package transx provides S3-compatible Object Storage providers.
// Supported providers:
//   - Direct: AWS S3, MinIO, and S3-compatible storage (via minio-go SDK)
//   - Spider: via CB-Spider Object Storage API
//   - Tumblebug: via CB-Tumblebug Object Storage API
package transx

import (
	"strings"
)

// S3Provider defines the interface for S3-compatible object storage operations.
type S3Provider interface {
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

// MinioConfig defines configuration for S3-compatible storage access using minio-go SDK.
// Supports: AWS S3, MinIO, Ceph, DigitalOcean Spaces, etc.
//
// MinioConfig is defined here as it's S3-specific.
// SpiderConfig and TumblebugConfig are defined in model.go as they're shared with the main transx package.
type MinioConfig struct {
	Endpoint        string `json:"endpoint" validate:"required"`
	AccessKeyId     string `json:"accessKeyId" validate:"required"`
	SecretAccessKey string `json:"secretAccessKey" validate:"required"`
	Region          string `json:"region,omitempty" default:"us-east-1"`
	UseSSL          bool   `json:"useSSL,omitempty" default:"true"`
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
