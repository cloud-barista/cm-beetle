package s3

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinioProvider implements Provider using minio-go SDK.
// Supports: AWS S3, MinIO, Ceph, DigitalOcean Spaces, and other S3-compatible services.
type MinioProvider struct {
	client   *minio.Client
	bucket   string
	endpoint string
	useSSL   bool
}

// NewMinioProvider creates a new MinioProvider from MinioConfig.
func NewMinioProvider(config *MinioConfig, bucket string) (*MinioProvider, error) {
	if config == nil {
		return nil, fmt.Errorf("minio config is required")
	}

	useSSL := config.UseSSL
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyId, config.SecretAccessKey, ""),
		Secure: useSSL,
		Region: config.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %w", err)
	}

	return &MinioProvider{
		client:   client,
		bucket:   bucket,
		endpoint: config.Endpoint,
		useSSL:   useSSL,
	}, nil
}

// GeneratePresignedURL generates a presigned URL for S3 operations.
func (p *MinioProvider) GeneratePresignedURL(action, key string) (string, error) {
	ctx := context.Background()
	expires := 1 * time.Hour // Default expiration

	switch action {
	case "upload":
		url, err := p.client.PresignedPutObject(ctx, p.bucket, key, expires)
		if err != nil {
			return "", fmt.Errorf("failed to generate presigned PUT URL: %w", err)
		}
		return url.String(), nil

	case "download":
		url, err := p.client.PresignedGetObject(ctx, p.bucket, key, expires, nil)
		if err != nil {
			return "", fmt.Errorf("failed to generate presigned GET URL: %w", err)
		}
		return url.String(), nil

	default:
		return "", fmt.Errorf("unsupported action: %s (use 'upload' or 'download')", action)
	}
}

// ListObjects lists objects in the bucket with the given prefix.
func (p *MinioProvider) ListObjects(prefix string) ([]ObjectInfo, error) {
	ctx := context.Background()
	var objects []ObjectInfo

	objectCh := p.client.ListObjects(ctx, p.bucket, minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true,
	})

	for obj := range objectCh {
		if obj.Err != nil {
			return nil, fmt.Errorf("error listing objects: %w", obj.Err)
		}
		objects = append(objects, ObjectInfo{
			Key:          obj.Key,
			Size:         obj.Size,
			LastModified: obj.LastModified.Format(time.RFC3339),
			ETag:         obj.ETag,
		})
	}

	return objects, nil
}

// GetBucket returns the bucket name.
func (p *MinioProvider) GetBucket() string {
	return p.bucket
}

// UploadFile uploads a local file to S3.
func (p *MinioProvider) UploadFile(localPath, key string) error {
	ctx := context.Background()
	_, err := p.client.FPutObject(ctx, p.bucket, key, localPath, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}

// DownloadFile downloads a file from S3 to local path.
func (p *MinioProvider) DownloadFile(key, localPath string) error {
	ctx := context.Background()
	err := p.client.FGetObject(ctx, p.bucket, key, localPath, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	return nil
}
