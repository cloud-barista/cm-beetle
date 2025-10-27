package transx

import (
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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

// uploadFileToObjectStorage uploads a single file to Object Storage using presigned URL
func uploadFileToObjectStorage(localFilePath, objectPath string, destEndpoint EndpointDetails, transferOptions *TransferOptions) error {
	// Generate presigned URL for upload
	apiEndpoint := destEndpoint.GetEndpoint()

	presignedURL, err := generatePresignedURL(apiEndpoint, "upload", objectPath, transferOptions.ObjectStorageTransferOptions)
	if err != nil {
		return fmt.Errorf("failed to generate presigned upload URL: %w", err)
	}

	// Set default options if not provided
	options := transferOptions.ObjectStorageTransferOptions
	if options == nil {
		options = &ObjectStorageTransferOption{
			Timeout:    300,
			MaxRetries: 3,
			VerifySSL:  false,
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

// downloadFileFromObjectStorage downloads a single file from Object Storage using presigned URL
func downloadFileFromObjectStorage(localFilePath, objectPath string, sourceEndpoint EndpointDetails, transferOptions *TransferOptions) error {

	// Generate presigned URL for download
	apiEndpoint := sourceEndpoint.GetEndpoint()

	presignedURL, err := generatePresignedURL(apiEndpoint, "download", objectPath, transferOptions.ObjectStorageTransferOptions)
	if err != nil {
		return fmt.Errorf("failed to generate presigned download URL: %w", err)
	}

	// Set default options if not provided
	options := transferOptions.ObjectStorageTransferOptions
	if options == nil {
		options = &ObjectStorageTransferOption{
			Timeout:    300,
			MaxRetries: 3,
			VerifySSL:  false,
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
func generatePresignedURL(apiEndpoint, operation, objectPath string, options *ObjectStorageTransferOption) (string, error) {
	if options == nil {
		return "", fmt.Errorf("ObjectStorageTransferOption is required")
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

// checkBucketExists checks if the bucket exists using HEAD request
func checkBucketExists(endpoint EndpointDetails, options *ObjectStorageTransferOption) error {
	if options == nil {
		return fmt.Errorf("ObjectStorageTransferOption is required")
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

// listBucketObjects lists objects in a bucket with optional prefix filtering
func listBucketObjects(endpoint EndpointDetails, prefix string, options *ObjectStorageTransferOption) ([]ObjectInfo, error) {
	if options == nil {
		return nil, fmt.Errorf("ObjectStorageTransferOption is required")
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
