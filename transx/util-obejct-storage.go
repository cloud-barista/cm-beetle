package transx

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// PresignedUrlInfo represents a presigned URL result from Object Storage API
type PresignedUrlInfo struct {
	PresignedURL string `xml:"PresignedURL"`
	Expires      int    `xml:"Expires"`
	Method       string `xml:"Method"`
}

// BucketInfo represents bucket listing information
type BucketInfo struct {
	Name        string       `xml:"Name"`
	Prefix      string       `xml:"Prefix"`
	Marker      string       `xml:"Marker"`
	MaxKeys     int          `xml:"MaxKeys"`
	IsTruncated bool         `xml:"IsTruncated"`
	Contents    []ObjectInfo `xml:"Contents"`
}

// ObjectInfo represents an object in the bucket
type ObjectInfo struct {
	Key          string `xml:"Key"`
	LastModified string `xml:"LastModified"`
	ETag         string `xml:"ETag"`
	Size         int64  `xml:"Size"`
	StorageClass string `xml:"StorageClass"`
}

// uploadFileToObjectStorage uploads a single file to Object Storage using presigned URL or SDK
func uploadFileToObjectStorage(localFilePath, objectPath string, destEndpoint EndpointDetails, transferOptions *TransferOptions) error {
	// Check if minio client is enabled
	if transferOptions.ObjectStorageOptions != nil && transferOptions.ObjectStorageOptions.Client == ObjectStorageClientMinio {
		return uploadFileToObjectStorageWithMinioSDK(localFilePath, objectPath, destEndpoint, transferOptions)
	}

	// Use presigned URL API (spider client - original implementation)
	// Generate presigned URL for upload
	apiEndpoint := destEndpoint.GetEndpoint()

	presignedURL, err := generatePresignedURL(apiEndpoint, "upload", objectPath, transferOptions.ObjectStorageOptions)
	if err != nil {
		return fmt.Errorf("failed to generate presigned upload URL: %w", err)
	}

	// Set default options if not provided
	options := transferOptions.ObjectStorageOptions
	if options == nil {
		options = &ObjectStorageOption{
			Timeout:    300,
			MaxRetries: 3,
			UseSSL:     false,
		}
	}

	file, err := os.Open(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to open local file: %w", err)
	}
	defer file.Close()

	// Get file info for content length
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	req, err := http.NewRequest("PUT", presignedURL, file)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.ContentLength = fileInfo.Size()

	client := &http.Client{
		Timeout: time.Duration(options.Timeout) * time.Second,
	}

	// Retry logic
	maxRetries := options.MaxRetries
	if maxRetries == 0 {
		maxRetries = 1
	}

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		// Reset file position for retries
		if attempt > 0 {
			file.Seek(0, 0)
		}

		resp, err := client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("HTTP request failed (attempt %d/%d): %w", attempt+1, maxRetries, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil // Success
		}

		lastErr = fmt.Errorf("upload failed with status %d (attempt %d/%d)", resp.StatusCode, attempt+1, maxRetries)
	}

	return lastErr
}

// downloadFileFromObjectStorage downloads a single file from Object Storage using presigned URL or SDK
func downloadFileFromObjectStorage(localFilePath, objectPath string, sourceEndpoint EndpointDetails, transferOptions *TransferOptions) error {
	// Check if minio client is enabled
	if transferOptions.ObjectStorageOptions != nil && transferOptions.ObjectStorageOptions.Client == ObjectStorageClientMinio {
		return downloadFileFromObjectStorageWithMinioSDK(localFilePath, objectPath, sourceEndpoint, transferOptions)
	}

	// Use presigned URL API (spider client - original implementation)
	// Generate presigned URL for download
	apiEndpoint := sourceEndpoint.GetEndpoint()

	presignedURL, err := generatePresignedURL(apiEndpoint, "download", objectPath, transferOptions.ObjectStorageOptions)
	if err != nil {
		return fmt.Errorf("failed to generate presigned download URL: %w", err)
	}

	// Set default options if not provided
	options := transferOptions.ObjectStorageOptions
	if options == nil {
		options = &ObjectStorageOption{
			Timeout:    300,
			MaxRetries: 3,
			UseSSL:     false,
		}
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(localFilePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	client := &http.Client{
		Timeout: time.Duration(options.Timeout) * time.Second,
	}

	// Retry logic
	maxRetries := options.MaxRetries
	if maxRetries == 0 {
		maxRetries = 1
	}

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, err := http.NewRequest("GET", presignedURL, nil)
		if err != nil {
			return fmt.Errorf("failed to create HTTP request: %w", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("HTTP request failed (attempt %d/%d): %w", attempt+1, maxRetries, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			lastErr = fmt.Errorf("download failed with status %d (attempt %d/%d)", resp.StatusCode, attempt+1, maxRetries)
			continue
		}

		// Create output file
		file, err := os.Create(localFilePath)
		if err != nil {
			lastErr = fmt.Errorf("failed to create local file: %w", err)
			continue
		}
		defer file.Close()

		// Copy data from response to file
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			lastErr = fmt.Errorf("failed to write file data: %w", err)
			continue
		}

		return nil // Success
	}

	return lastErr
}

// generatePresignedURL generates a presigned URL for Object Storage operations
func generatePresignedURL(apiEndpoint, operation, objectPath string, options *ObjectStorageOption) (string, error) {
	if options == nil {
		return "", fmt.Errorf("ObjectStorageOption is required")
	}

	// Build presigned URL request URL for Object Storage API
	// Format: GET /spider/s3/presigned/{operation}/{bucket-name}/{object-key}?ConnectionName={conn}&expires={seconds}
	presignedAPIURL := fmt.Sprintf("%s/presigned/%s/%s", apiEndpoint, operation, objectPath)

	// Add query parameters
	params := fmt.Sprintf("ConnectionName=%s", options.AccessKeyId)
	if options.ExpiresIn > 0 {
		params += fmt.Sprintf("&expires=%d", options.ExpiresIn)
	} else {
		params += "&expires=3600" // Default 1 hour
	}

	presignedAPIURL += "?" + params

	// Make HTTP GET request to Object Storage API to get the actual presigned URL
	client := &http.Client{
		Timeout: time.Duration(options.Timeout) * time.Second,
	}
	if options.Timeout == 0 {
		client.Timeout = 300 * time.Second // 5 minutes default
	}

	resp, err := client.Get(presignedAPIURL)
	if err != nil {
		return "", fmt.Errorf("failed to request presigned URL from Object Storage API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("object Storage API returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse XML response to extract presigned URL
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Fix XML parsing issue: CB-Spider returns unescaped & characters in URLs
	// Replace unescaped & with &amp; for proper XML parsing
	bodyString := string(bodyBytes)

	// This handles CB-Spider's improper XML encoding of query parameters in presigned URLs
	// Simple approach: replace all & with &amp; first, then fix double-escaping
	bodyString = strings.ReplaceAll(bodyString, "&", "&amp;")
	// Fix double-escaping of common XML entities
	bodyString = strings.ReplaceAll(bodyString, "&amp;amp;", "&amp;")
	bodyString = strings.ReplaceAll(bodyString, "&amp;lt;", "&lt;")
	bodyString = strings.ReplaceAll(bodyString, "&amp;gt;", "&gt;")
	bodyString = strings.ReplaceAll(bodyString, "&amp;quot;", "&quot;")
	bodyString = strings.ReplaceAll(bodyString, "&amp;apos;", "&apos;")

	fixedBodyBytes := []byte(bodyString)

	// Parse XML response using struct
	var presignedInfo PresignedUrlInfo
	if err := xml.Unmarshal(fixedBodyBytes, &presignedInfo); err != nil {
		return "", fmt.Errorf("failed to parse presigned URL XML response: %w", err)
	}

	if presignedInfo.PresignedURL == "" {
		return "", fmt.Errorf("empty presigned URL in response")
	}

	// Decode HTML entities (e.g., &amp; -> &) that may be present in XML
	decodedURL := html.UnescapeString(presignedInfo.PresignedURL)

	return decodedURL, nil
}

// checkBucketExists checks if the bucket exists using HEAD request or SDK
func checkBucketExists(endpoint EndpointDetails, options *ObjectStorageOption) error {
	// Check if minio client is enabled
	if options != nil && options.Client == ObjectStorageClientMinio {
		return checkBucketExistsWithMinioSDK(endpoint, options)
	}

	// Use presigned URL API (spider client - original implementation)
	if options == nil {
		return fmt.Errorf("ObjectStorageOption is required")
	}

	if options.AccessKeyId == "" {
		return fmt.Errorf("accessKeyId is required for bucket check")
	}

	bucket, _, err := endpoint.GetBucketAndObjectKey()
	if err != nil {
		return fmt.Errorf("failed to parse bucket name: %w", err)
	}

	apiBase := endpoint.GetEndpoint()
	if apiBase == "" {
		return fmt.Errorf("object Storage API endpoint is required")
	}

	// Build bucket check URL
	// Format: HEAD /spider/s3/{bucket-name}?ConnectionName={conn}
	bucketCheckURL := fmt.Sprintf("%s/%s?ConnectionName=%s", apiBase, bucket, options.AccessKeyId)

	// Create HTTP client
	client := &http.Client{
		Timeout: time.Duration(options.Timeout) * time.Second,
	}
	if options.Timeout == 0 {
		client.Timeout = 300 * time.Second // 5 minutes default
	}

	// Make HEAD request to check bucket existence
	req, err := http.NewRequest("HEAD", bucketCheckURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create bucket check request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to check bucket existence: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return nil // Bucket exists
	case 404:
		return fmt.Errorf("bucket '%s' does not exist", bucket)
	default:
		return fmt.Errorf("bucket check failed with status %d", resp.StatusCode)
	}
}

// listBucketObjects lists objects in a bucket with optional prefix filtering using presigned URL API or SDK
func listBucketObjects(endpoint EndpointDetails, prefix string, options *ObjectStorageOption) ([]ObjectInfo, error) {
	// Check if minio client is enabled
	if options != nil && options.Client == ObjectStorageClientMinio {
		return listBucketObjectsWithMinioSDK(endpoint, prefix, options)
	}

	// Use presigned URL API (spider client - original implementation)
	if options == nil {
		return nil, fmt.Errorf("ObjectStorageOption is required")
	}

	if options.AccessKeyId == "" {
		return nil, fmt.Errorf("accessKeyId is required for listing objects")
	}

	bucket, _, err := endpoint.GetBucketAndObjectKey()
	if err != nil {
		return nil, fmt.Errorf("failed to parse bucket name: %w", err)
	}

	apiBase := endpoint.GetEndpoint()
	if apiBase == "" {
		return nil, fmt.Errorf("object Storage API endpoint is required")
	}

	// Build bucket listing URL
	// Format: GET /spider/s3/{bucket-name}?ConnectionName={conn}&prefix={prefix}
	listURL := fmt.Sprintf("%s/%s?ConnectionName=%s", apiBase, bucket, options.AccessKeyId)
	if prefix != "" {
		listURL += "&prefix=" + prefix
	}

	// Create HTTP client
	client := &http.Client{
		Timeout: time.Duration(options.Timeout) * time.Second,
	}
	if options.Timeout == 0 {
		client.Timeout = 300 * time.Second // 5 minutes default
	}

	// Make GET request to list objects
	resp, err := client.Get(listURL)
	if err != nil {
		return nil, fmt.Errorf("failed to list bucket objects: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bucket listing failed with status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse XML response using struct
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse XML response using struct
	var bucketInfo BucketInfo
	if err := xml.Unmarshal(bodyBytes, &bucketInfo); err != nil {
		return nil, fmt.Errorf("failed to parse bucket listing XML response: %w", err)
	}

	return bucketInfo.Contents, nil
}

// parseObjectListXML parses the XML response from bucket listing

// ========================================
// MinIO SDK-based Object Storage Functions
// ========================================

// createMinioClient creates a MinIO client for S3-compatible storage
func createMinioClient(endpoint string, options *ObjectStorageOption) (*minio.Client, error) {
	if options == nil {
		return nil, fmt.Errorf("ObjectStorageOption is required")
	}

	if options.Client != ObjectStorageClientMinio {
		return nil, fmt.Errorf("minio client is not enabled")
	}

	if options.SecretAccessKey == "" {
		return nil, fmt.Errorf("secretAccessKey is required for MinIO SDK mode")
	}

	// Remove protocol prefix if present (MinIO SDK doesn't expect it)
	endpoint = strings.TrimPrefix(endpoint, "http://")
	endpoint = strings.TrimPrefix(endpoint, "https://")

	// Default to SSL if not specified
	useSSL := options.UseSSL

	// Initialize MinIO client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(options.AccessKeyId, options.SecretAccessKey, ""),
		Secure: useSSL,
		Region: options.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	return minioClient, nil
}

// uploadFileToObjectStorageWithMinioSDK uploads a file using MinIO SDK
func uploadFileToObjectStorageWithMinioSDK(localFilePath, objectPath string, destEndpoint EndpointDetails, transferOptions *TransferOptions) error {
	options := transferOptions.ObjectStorageOptions
	if options == nil {
		return fmt.Errorf("ObjectStorageOption is required")
	}

	// Create MinIO client
	endpoint := destEndpoint.GetEndpoint()
	minioClient, err := createMinioClient(endpoint, options)
	if err != nil {
		return fmt.Errorf("failed to create MinIO client: %w", err)
	}

	// Parse bucket and object key
	bucket, objectKey, err := destEndpoint.GetBucketAndObjectKey()
	if err != nil {
		return fmt.Errorf("failed to parse bucket and object key: %w", err)
	}

	// If objectPath is provided, use it as the object key
	if objectPath != "" {
		parts := strings.SplitN(objectPath, "/", 2)
		if len(parts) == 2 {
			bucket = parts[0]
			objectKey = parts[1]
		}
	}

	// Create context with timeout
	ctx := context.Background()
	if options.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(options.Timeout)*time.Second)
		defer cancel()
	}

	// Upload file with retry logic
	maxRetries := options.MaxRetries
	if maxRetries == 0 {
		maxRetries = 1
	}

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		_, err = minioClient.FPutObject(ctx, bucket, objectKey, localFilePath, minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		})
		if err == nil {
			return nil // Success
		}
		lastErr = fmt.Errorf("upload failed (attempt %d/%d): %w", attempt+1, maxRetries, err)
	}

	return lastErr
}

// downloadFileFromObjectStorageWithMinioSDK downloads a file using MinIO SDK
func downloadFileFromObjectStorageWithMinioSDK(localFilePath, objectPath string, sourceEndpoint EndpointDetails, transferOptions *TransferOptions) error {
	options := transferOptions.ObjectStorageOptions
	if options == nil {
		return fmt.Errorf("ObjectStorageOption is required")
	}

	// Create MinIO client
	endpoint := sourceEndpoint.GetEndpoint()
	minioClient, err := createMinioClient(endpoint, options)
	if err != nil {
		return fmt.Errorf("failed to create MinIO client: %w", err)
	}

	// Parse bucket and object key
	bucket, objectKey, err := sourceEndpoint.GetBucketAndObjectKey()
	if err != nil {
		return fmt.Errorf("failed to parse bucket and object key: %w", err)
	}

	// If objectPath is provided, use it as the object key
	if objectPath != "" {
		parts := strings.SplitN(objectPath, "/", 2)
		if len(parts) == 2 {
			bucket = parts[0]
			objectKey = parts[1]
		}
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(localFilePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create context with timeout
	ctx := context.Background()
	if options.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(options.Timeout)*time.Second)
		defer cancel()
	}

	// Download file with retry logic
	maxRetries := options.MaxRetries
	if maxRetries == 0 {
		maxRetries = 1
	}

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		err = minioClient.FGetObject(ctx, bucket, objectKey, localFilePath, minio.GetObjectOptions{})
		if err == nil {
			return nil // Success
		}
		lastErr = fmt.Errorf("download failed (attempt %d/%d): %w", attempt+1, maxRetries, err)
	}

	return lastErr
}

// checkBucketExistsWithMinioSDK checks if a bucket exists using MinIO SDK
func checkBucketExistsWithMinioSDK(endpoint EndpointDetails, options *ObjectStorageOption) error {
	if options == nil {
		return fmt.Errorf("ObjectStorageTransferOption is required")
	}

	// Create MinIO client
	apiEndpoint := endpoint.GetEndpoint()
	minioClient, err := createMinioClient(apiEndpoint, options)
	if err != nil {
		return fmt.Errorf("failed to create MinIO client: %w", err)
	}

	// Parse bucket name
	bucket, _, err := endpoint.GetBucketAndObjectKey()
	if err != nil {
		return fmt.Errorf("failed to parse bucket name: %w", err)
	}

	// Create context with timeout
	ctx := context.Background()
	if options.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(options.Timeout)*time.Second)
		defer cancel()
	}

	// Check if bucket exists
	exists, err := minioClient.BucketExists(ctx, bucket)
	if err != nil {
		return fmt.Errorf("failed to check bucket existence: %w", err)
	}

	if !exists {
		return fmt.Errorf("bucket '%s' does not exist", bucket)
	}

	return nil
}

// listBucketObjectsWithMinioSDK lists objects in a bucket using MinIO SDK
func listBucketObjectsWithMinioSDK(endpoint EndpointDetails, prefix string, options *ObjectStorageOption) ([]ObjectInfo, error) {
	if options == nil {
		return nil, fmt.Errorf("ObjectStorageTransferOption is required")
	}

	// Create MinIO client
	apiEndpoint := endpoint.GetEndpoint()
	minioClient, err := createMinioClient(apiEndpoint, options)
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	// Parse bucket name
	bucket, _, err := endpoint.GetBucketAndObjectKey()
	if err != nil {
		return nil, fmt.Errorf("failed to parse bucket name: %w", err)
	}

	// Create context with timeout
	ctx := context.Background()
	if options.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(options.Timeout)*time.Second)
		defer cancel()
	}

	// List objects recursively
	objectCh := minioClient.ListObjects(ctx, bucket, minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true,
	})

	var objects []ObjectInfo
	for object := range objectCh {
		if object.Err != nil {
			return nil, fmt.Errorf("error listing objects: %w", object.Err)
		}

		// Filter out directory markers (placeholder objects created by some S3 clients)
		// AWS S3/MinIO doesn't have native directory support - it's a flat key-value store.
		// Some tools (AWS Console, s3cmd, etc.) create empty objects with keys ending in "/"
		// to simulate directory structure. These should be excluded from file transfers.
		// Example: "resources/" (Size: 0) is a directory marker, not an actual file.
		if object.Size == 0 && strings.HasSuffix(object.Key, "/") {
			continue // Skip directory markers
		}

		objects = append(objects, ObjectInfo{
			Key:          object.Key,
			LastModified: object.LastModified.Format(time.RFC3339),
			ETag:         object.ETag,
			Size:         object.Size,
			StorageClass: object.StorageClass,
		})
	}

	return objects, nil
}

// ========================================
// Object Filtering Functions
// ========================================

// filterObjectList filters a list of objects based on include/exclude patterns
// Filtering logic follows rsync-like pattern matching:
// 1. If include patterns are specified, only matching objects are included
// 2. Exclude patterns are then applied to remove unwanted objects
// 3. If no patterns are specified, all objects are included
func filterObjectList(objects []ObjectInfo, exclude, include []string) []ObjectInfo {
	if len(include) == 0 && len(exclude) == 0 {
		return objects // No filtering needed
	}

	filtered := []ObjectInfo{}
	for _, obj := range objects {
		if shouldTransferObject(obj.Key, exclude, include) {
			filtered = append(filtered, obj)
		}
	}
	return filtered
}

// shouldTransferObject determines if an object should be transferred based on include/exclude patterns
// Pattern matching uses filepath.Match syntax:
// - "*" matches any sequence of non-separator characters
// - "?" matches any single non-separator character
// - "[...]" matches any character in the brackets
// - "**" (double asterisk) is NOT supported by filepath.Match
//
// Examples:
//   - "*.log" matches "app.log", "error.log"
//   - "temp/*" matches "temp/file.txt", "temp/data.json"
//   - "data/*.json" matches "data/config.json"
func shouldTransferObject(objectKey string, exclude, include []string) bool {
	// Step 1: Apply include patterns (whitelist)
	// If include patterns exist, object must match at least one
	if len(include) > 0 {
		included := false
		for _, pattern := range include {
			if matchPattern(objectKey, pattern) {
				included = true
				break
			}
		}
		if !included {
			return false // Object doesn't match any include pattern
		}
	}

	// Step 2: Apply exclude patterns (blacklist)
	// If object matches any exclude pattern, it's filtered out
	for _, pattern := range exclude {
		if matchPattern(objectKey, pattern) {
			return false // Object matches an exclude pattern
		}
	}

	return true // Object passed all filters
}

// matchPattern checks if a file path matches a pattern
// Supports glob patterns including ** for recursive directory matching
// Pattern examples:
//   - "*.log" matches any .log file in any directory
//   - "data/*" matches files directly in data/ directory
//   - "data/**" matches all files recursively under data/
//   - "data/**/*.json" matches all .json files recursively under data/
func matchPattern(path, pattern string) bool {
	// Handle ** (double asterisk) for recursive directory matching
	if strings.Contains(pattern, "**") {
		return matchPatternWithDoubleAsterisk(path, pattern)
	}

	// Direct match attempt using filepath.Match
	matched, err := filepath.Match(pattern, path)
	if err == nil && matched {
		return true
	}

	// Try matching with path prefix for directory patterns
	// e.g., "logs/*" should match "logs/app.log"
	if strings.Contains(pattern, "/") {
		matched, err := filepath.Match(pattern, path)
		if err == nil && matched {
			return true
		}
	}

	// Try matching basename for simple patterns
	// e.g., "*.log" should match "subdir/app.log"
	if !strings.Contains(pattern, "/") {
		basename := filepath.Base(path)
		matched, err := filepath.Match(pattern, basename)
		if err == nil && matched {
			return true
		}
	}

	return false
}

// matchPatternWithDoubleAsterisk handles patterns with ** for recursive matching
// Examples:
//   - "data/**" matches "data/file.txt", "data/sub/file.txt", "data/sub/deep/file.txt"
//   - "data/**/*.json" matches all .json files under data/ recursively
//   - "**/test/**" matches any path containing /test/ directory
//   - "*/raw/2025/**" matches "data/raw/2025/file.txt", "backup/raw/2025/sub/file.txt"
func matchPatternWithDoubleAsterisk(path, pattern string) bool {
	// Split pattern by ** to handle different cases
	parts := strings.Split(pattern, "**")

	if len(parts) == 1 {
		// No ** found, shouldn't happen but fallback to regular match
		matched, _ := filepath.Match(pattern, path)
		return matched
	}

	// Handle different ** patterns
	if len(parts) == 2 {
		prefix := parts[0]
		suffix := parts[1]

		// Remove trailing / from prefix and leading / from suffix
		prefix = strings.TrimSuffix(prefix, "/")
		suffix = strings.TrimPrefix(suffix, "/")

		// Case 1: "data/**" - matches everything under data/
		if prefix != "" && suffix == "" {
			// Check if prefix contains wildcards
			if strings.ContainsAny(prefix, "*?[") {
				// Use glob matching for prefix
				return matchPrefixWithGlob(path, prefix)
			}
			return strings.HasPrefix(path, prefix+"/") || path == prefix
		}

		// Case 2: "**/*.json" - matches all .json files anywhere
		if prefix == "" && suffix != "" {
			// Extract the basename pattern
			matched, _ := filepath.Match(suffix, filepath.Base(path))
			if matched {
				return true
			}
			// Try matching the full suffix against path
			if strings.HasSuffix(path, suffix) {
				return true
			}
			// Try pattern matching on path segments
			pathParts := strings.Split(path, "/")
			for i := range pathParts {
				subPath := strings.Join(pathParts[i:], "/")
				matched, _ := filepath.Match(suffix, subPath)
				if matched {
					return true
				}
			}
			return false
		}

		// Case 3: "data/**/file.txt" or "*/raw/2025/**" - matches with prefix pattern
		if prefix != "" && suffix != "" {
			// Check if prefix contains wildcards
			if strings.ContainsAny(prefix, "*?[") {
				// Match prefix with glob pattern, then check suffix
				return matchPrefixSuffixWithGlob(path, prefix, suffix)
			}

			// Simple prefix without wildcards
			if !strings.HasPrefix(path, prefix+"/") && path != prefix {
				return false
			}
			// Check if suffix matches
			if suffix == "" {
				return true
			}
			if strings.HasSuffix(path, suffix) {
				return true
			}
			// Try matching suffix pattern
			pathAfterPrefix := strings.TrimPrefix(path, prefix+"/")
			pathParts := strings.Split(pathAfterPrefix, "/")
			for i := range pathParts {
				subPath := strings.Join(pathParts[i:], "/")
				matched, _ := filepath.Match(suffix, subPath)
				if matched {
					return true
				}
			}
			return false
		}
	}

	// Handle multiple ** in pattern (rare case)
	// For now, use simple contains check
	return strings.Contains(path, strings.ReplaceAll(pattern, "**", ""))
}

// matchPrefixWithGlob matches path against a prefix pattern with wildcards
// Example: prefix="*/raw", path="data/raw/file.txt" -> true
func matchPrefixWithGlob(path, prefix string) bool {
	pathParts := strings.Split(path, "/")
	prefixParts := strings.Split(prefix, "/")

	// Path must have at least as many parts as prefix
	if len(pathParts) < len(prefixParts) {
		return false
	}

	// Try to match prefix parts against path parts
	for i := 0; i <= len(pathParts)-len(prefixParts); i++ {
		matched := true
		for j, prefixPart := range prefixParts {
			partMatched, _ := filepath.Match(prefixPart, pathParts[i+j])
			if !partMatched {
				matched = false
				break
			}
		}
		if matched {
			return true
		}
	}

	return false
}

// matchPrefixSuffixWithGlob matches path against prefix and suffix patterns
// Example: prefix="*/raw/2025", suffix="", path="data/raw/2025/file.txt" -> true
func matchPrefixSuffixWithGlob(path, prefix, suffix string) bool {
	pathParts := strings.Split(path, "/")
	prefixParts := strings.Split(prefix, "/")

	// Path must have at least as many parts as prefix
	if len(pathParts) < len(prefixParts) {
		return false
	}

	// Try to match prefix parts against path parts
	for i := 0; i <= len(pathParts)-len(prefixParts); i++ {
		matched := true
		for j, prefixPart := range prefixParts {
			partMatched, _ := filepath.Match(prefixPart, pathParts[i+j])
			if !partMatched {
				matched = false
				break
			}
		}
		if matched {
			// Prefix matched, now check suffix
			if suffix == "" {
				return true // No suffix requirement
			}

			// Get the remaining path after prefix
			remainingParts := pathParts[i+len(prefixParts):]
			if len(remainingParts) == 0 {
				// No remaining parts, check if suffix is empty or matches empty
				return suffix == ""
			}

			remainingPath := strings.Join(remainingParts, "/")

			// Try matching suffix
			if strings.HasSuffix(remainingPath, suffix) {
				return true
			}

			// Try pattern matching on remaining parts
			for k := range remainingParts {
				subPath := strings.Join(remainingParts[k:], "/")
				suffixMatched, _ := filepath.Match(suffix, subPath)
				if suffixMatched {
					return true
				}
			}
		}
	}

	return false
}
