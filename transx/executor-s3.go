package transx

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// S3Executor implements Executor for S3 object storage transfers.
// Uses presigned URLs for authentication-free upload/download.
type S3Executor struct {
	Provider S3Provider // S3 provider for generating presigned URLs
}

// NewS3Executor creates a new S3Executor with the given provider.
func NewS3Executor(provider S3Provider) *S3Executor {
	return &S3Executor{
		Provider: provider,
	}
}

// Execute performs S3 transfer from source to destination.
func (e *S3Executor) Execute(source, destination DataLocation) error {
	// Determine transfer direction based on StorageType
	srcIsS3 := source.IsObjectStorage()
	dstIsS3 := destination.IsObjectStorage()

	switch {
	case !srcIsS3 && dstIsS3:
		// Filesystem -> S3 (upload)
		return e.upload(source.Path, destination.Path, source.Filter)
	case srcIsS3 && !dstIsS3:
		// S3 -> Filesystem (download)
		return e.download(source.Path, destination.Path, source.Filter)
	case srcIsS3 && dstIsS3:
		// S3 -> S3 (not implemented yet)
		return fmt.Errorf("S3 to S3 transfer not implemented")
	default:
		return fmt.Errorf("invalid transfer: both source and destination are filesystem")
	}
}

// upload transfers local files to S3.
func (e *S3Executor) upload(localPath, s3Path string, filter *FilterOption) error {
	files, err := e.listLocalFiles(localPath, filter)
	if err != nil {
		return fmt.Errorf("failed to list local files: %w", err)
	}

	// Parse bucket and key from s3Path (e.g., "bucket-name/prefix/")
	_, keyPrefix := ParseBucketAndKey(s3Path)

	basePath := localPath
	if !isDirectory(localPath) {
		basePath = filepath.Dir(localPath)
	}

	for _, file := range files {
		relPath, err := filepath.Rel(basePath, file)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// Use only the key prefix, not the full path including bucket name
		s3Key := filepath.Join(keyPrefix, relPath)
		s3Key = strings.ReplaceAll(s3Key, "\\", "/") // Normalize to forward slashes

		if err := e.uploadFile(file, s3Key); err != nil {
			return fmt.Errorf("failed to upload %s: %w", file, err)
		}
	}

	return nil
}

// download transfers S3 objects to local filesystem.
func (e *S3Executor) download(s3Path, localPath string, filter *FilterOption) error {
	// Parse bucket and key from s3Path (e.g., "bucket-name/prefix/")
	_, keyPrefix := ParseBucketAndKey(s3Path)

	objects, err := e.Provider.ListObjects(keyPrefix)
	if err != nil {
		return fmt.Errorf("failed to list S3 objects: %w", err)
	}

	for _, obj := range objects {
		if filter != nil && !e.matchesFilter(obj.Key, filter) {
			continue
		}

		relPath := strings.TrimPrefix(obj.Key, keyPrefix)
		relPath = strings.TrimPrefix(relPath, "/")
		localFile := filepath.Join(localPath, relPath)

		// Create parent directories
		if err := os.MkdirAll(filepath.Dir(localFile), 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		if err := e.downloadFile(obj.Key, localFile); err != nil {
			return fmt.Errorf("failed to download %s: %w", obj.Key, err)
		}
	}

	return nil
}

// uploadFile uploads a single file using presigned URL.
func (e *S3Executor) uploadFile(localPath, s3Key string) error {
	presignedURL, err := e.Provider.GeneratePresignedURL("upload", s3Key)
	if err != nil {
		return fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	file, err := os.Open(localPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat file: %w", err)
	}

	req, err := http.NewRequest(http.MethodPut, presignedURL, file)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.ContentLength = stat.Size()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("upload request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}

// downloadFile downloads a single file using presigned URL.
func (e *S3Executor) downloadFile(s3Key, localPath string) error {
	presignedURL, err := e.Provider.GeneratePresignedURL("download", s3Key)
	if err != nil {
		return fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	resp, err := http.Get(presignedURL)
	if err != nil {
		return fmt.Errorf("download request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("download failed with status %d: %s", resp.StatusCode, string(body))
	}

	file, err := os.Create(localPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// listLocalFiles returns a list of files matching the filter.
func (e *S3Executor) listLocalFiles(path string, filter *FilterOption) ([]string, error) {
	var files []string

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return []string{path}, nil
	}

	err = filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filter != nil && !e.matchesFilter(filePath, filter) {
			return nil
		}
		files = append(files, filePath)
		return nil
	})

	return files, err
}

// matchesFilter checks if a path matches the filter patterns.
func (e *S3Executor) matchesFilter(path string, filter *FilterOption) bool {
	if filter == nil {
		return true
	}

	// Check exclude patterns first
	for _, pattern := range filter.Exclude {
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return false
		}
	}

	// If include patterns exist, path must match at least one
	if len(filter.Include) > 0 {
		for _, pattern := range filter.Include {
			if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
				return true
			}
		}
		return false
	}

	return true
}

// isDirectory checks if the path is a directory.
func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
