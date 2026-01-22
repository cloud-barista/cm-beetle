package s3

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SpiderProvider implements Provider for CB-Spider S3 Object Storage API.
// Based on CB-Spider swagger.yaml [S3 Object Storage Management] endpoints.
type SpiderProvider struct {
	endpoint       string
	connectionName string
	expires        int
	bucket         string
}

// NewSpiderProvider creates a new SpiderProvider from SpiderConfig.
func NewSpiderProvider(config *SpiderConfig, bucket string) (*SpiderProvider, error) {
	if config == nil {
		return nil, fmt.Errorf("spider config is required")
	}
	if strings.TrimSpace(config.Endpoint) == "" {
		return nil, fmt.Errorf("spider endpoint is required")
	}
	if strings.TrimSpace(config.ConnectionName) == "" {
		return nil, fmt.Errorf("spider connectionName is required")
	}

	expires := config.Expires
	if expires <= 0 {
		expires = 3600 // Default 1 hour
	}

	return &SpiderProvider{
		endpoint:       strings.TrimSuffix(config.Endpoint, "/"),
		connectionName: config.ConnectionName,
		expires:        expires,
		bucket:         bucket,
	}, nil
}

// GeneratePresignedURL generates a presigned URL via CB-Spider S3 API.
// Uses the CB-Spider special feature endpoints:
//   - GET /s3/presigned/download/{BucketName}/{ObjectKey} for download
//   - GET /s3/presigned/upload/{BucketName}/{ObjectKey} for upload
func (p *SpiderProvider) GeneratePresignedURL(action, key string) (string, error) {
	var apiURL string

	// URL encode the object key to handle special characters and paths
	encodedKey := url.PathEscape(key)

	switch action {
	case "upload":
		// GET /s3/presigned/upload/{BucketName}/{ObjectKey}
		apiURL = fmt.Sprintf("%s/spider/s3/presigned/upload/%s/%s?ConnectionName=%s&expires=%d",
			p.endpoint, p.bucket, encodedKey, url.QueryEscape(p.connectionName), p.expires)
	case "download":
		// GET /s3/presigned/download/{BucketName}/{ObjectKey}
		apiURL = fmt.Sprintf("%s/spider/s3/presigned/download/%s/%s?ConnectionName=%s&expires=%d",
			p.endpoint, p.bucket, encodedKey, url.QueryEscape(p.connectionName), p.expires)
	default:
		return "", fmt.Errorf("unsupported action: %s (use 'upload' or 'download')", action)
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("spider API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("spider API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	presignedURL := extractPresignedURL(string(body))
	if presignedURL == "" {
		return "", fmt.Errorf("failed to extract presigned URL from response: %s", string(body))
	}

	return presignedURL, nil
}

// ListObjects lists objects via CB-Spider S3 API.
// Uses GET /s3/{BucketName}?ConnectionName=xxx to list objects in bucket.
func (p *SpiderProvider) ListObjects(prefix string) ([]ObjectInfo, error) {
	// GET /s3/{BucketName}?ConnectionName=xxx&prefix=xxx
	apiURL := fmt.Sprintf("%s/spider/s3/%s?ConnectionName=%s",
		p.endpoint, p.bucket, url.QueryEscape(p.connectionName))

	if prefix != "" {
		apiURL += "&prefix=" + url.QueryEscape(prefix)
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("spider API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("spider API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return parseSpiderListResponse(body)
}

// GetBucket returns the bucket name.
func (p *SpiderProvider) GetBucket() string {
	return p.bucket
}

// spiderListResponse represents CB-Spider's S3 list bucket response (XML format).
// Follows standard S3 ListBucketResult format.
type spiderListResponse struct {
	XMLName     xml.Name            `xml:"ListBucketResult"`
	Name        string              `xml:"Name"`
	Prefix      string              `xml:"Prefix"`
	MaxKeys     int                 `xml:"MaxKeys"`
	IsTruncated bool                `xml:"IsTruncated"`
	Contents    []spiderObjectEntry `xml:"Contents"`
}

type spiderObjectEntry struct {
	Key          string `xml:"Key"`
	Size         int64  `xml:"Size"`
	LastModified string `xml:"LastModified"`
	ETag         string `xml:"ETag"`
	StorageClass string `xml:"StorageClass"`
}

// parseSpiderListResponse parses the CB-Spider S3 list objects response.
func parseSpiderListResponse(body []byte) ([]ObjectInfo, error) {
	var result spiderListResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse XML response: %w", err)
	}

	var objects []ObjectInfo
	for _, entry := range result.Contents {
		objects = append(objects, ObjectInfo{
			Key:          entry.Key,
			Size:         entry.Size,
			LastModified: entry.LastModified,
			ETag:         entry.ETag,
		})
	}

	return objects, nil
}

// spiderPresignedURLResponse represents CB-Spider presigned URL response.
// Based on spider.S3PresignedURLXML definition.
type spiderPresignedURLResponse struct {
	XMLName      xml.Name `xml:"PresignedURL" json:"-"`
	PresignedURL string   `xml:"PresignedURL" json:"PresignedURL"`
	Method       string   `xml:"Method" json:"Method"`
	Expires      int      `xml:"Expires" json:"Expires"`
}

// extractPresignedURL extracts the presigned URL from Spider API response.
// Supports both JSON and XML response formats.
func extractPresignedURL(response string) string {
	response = strings.TrimSpace(response)

	// Try JSON parsing first
	if strings.HasPrefix(response, "{") {
		var jsonResp spiderPresignedURLResponse
		if err := json.Unmarshal([]byte(response), &jsonResp); err == nil && jsonResp.PresignedURL != "" {
			return jsonResp.PresignedURL
		}

		// Try alternate JSON field names
		patterns := []string{`"PresignedURL":"`, `"presignedUrl":"`, `"url":"`}
		for _, pattern := range patterns {
			if idx := strings.Index(response, pattern); idx != -1 {
				start := idx + len(pattern)
				end := strings.Index(response[start:], `"`)
				if end != -1 {
					return response[start : start+end]
				}
			}
		}
	}

	// Try XML parsing
	if strings.HasPrefix(response, "<") {
		var xmlResp spiderPresignedURLResponse
		if err := xml.Unmarshal([]byte(response), &xmlResp); err == nil && xmlResp.PresignedURL != "" {
			return xmlResp.PresignedURL
		}
	}

	return ""
}
