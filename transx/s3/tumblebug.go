package s3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TumblebugProvider implements Provider for CB-Tumblebug Object Storage API.
// Based on CB-Tumblebug swagger.yaml [Infra Resource] Object Storage Management endpoints.
type TumblebugProvider struct {
	endpoint string
	nsId     string
	osId     string
	expires  int
}

// NewTumblebugProvider creates a new TumblebugProvider from TumblebugConfig.
func NewTumblebugProvider(config *TumblebugConfig) (*TumblebugProvider, error) {
	if config == nil {
		return nil, fmt.Errorf("tumblebug config is required")
	}
	if strings.TrimSpace(config.Endpoint) == "" {
		return nil, fmt.Errorf("tumblebug endpoint is required")
	}
	if strings.TrimSpace(config.NsId) == "" {
		return nil, fmt.Errorf("tumblebug nsId is required")
	}
	if strings.TrimSpace(config.OsId) == "" {
		return nil, fmt.Errorf("tumblebug osId is required")
	}

	expires := config.Expires
	if expires <= 0 {
		expires = 3600 // Default 1 hour
	}

	return &TumblebugProvider{
		endpoint: strings.TrimSuffix(config.Endpoint, "/"),
		nsId:     config.NsId,
		osId:     config.OsId,
		expires:  expires,
	}, nil
}

// GeneratePresignedURL generates a presigned URL via CB-Tumblebug API.
// Uses the new endpoint: GET /ns/{nsId}/resources/objectStorage/{osId}/object/{objectKey}
// Query parameters:
//   - operation: "upload" or "download"
//   - expires: expiration time in seconds (default: 3600)
func (p *TumblebugProvider) GeneratePresignedURL(action, key string) (string, error) {
	// URL encode the object key to handle special characters and paths
	encodedKey := url.PathEscape(key)

	// GET /tumblebug/ns/{nsId}/resources/objectStorage/{osId}/object/{objectKey}?operation=xxx&expires=xxx
	apiURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/objectStorage/%s/object/%s?operation=%s&expires=%d",
		p.endpoint, p.nsId, p.osId, encodedKey, action, p.expires)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("tumblebug API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("tumblebug API returned status %d: %s", resp.StatusCode, string(body))
	}

	var response tumblebugPresignedURLResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if response.PresignedURL == "" {
		return "", fmt.Errorf("empty presigned URL in response")
	}

	return response.PresignedURL, nil
}

// ListObjects lists objects via CB-Tumblebug API.
// Uses GET /ns/{nsId}/resources/objectStorage/{osId} to list objects in bucket.
func (p *TumblebugProvider) ListObjects(prefix string) ([]ObjectInfo, error) {
	// GET /tumblebug/ns/{nsId}/resources/objectStorage/{osId}
	apiURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/objectStorage/%s",
		p.endpoint, p.nsId, p.osId)

	if prefix != "" {
		apiURL += "?prefix=" + url.QueryEscape(prefix)
	}

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("tumblebug API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("tumblebug API returned status %d: %s", resp.StatusCode, string(body))
	}

	var response tumblebugListResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	var objects []ObjectInfo
	for _, obj := range response.Objects {
		objects = append(objects, ObjectInfo{
			Key:          obj.Key,
			Size:         obj.Size,
			LastModified: obj.LastModified,
			ETag:         obj.ETag,
		})
	}

	return objects, nil
}

// GetBucket returns the osId as the bucket identifier.
func (p *TumblebugProvider) GetBucket() string {
	return p.osId
}

// Response types for Tumblebug API
type tumblebugPresignedURLResponse struct {
	PresignedURL string `json:"presignedUrl"`
	ExpiresAt    string `json:"expiresAt,omitempty"`
}

type tumblebugListResponse struct {
	Objects []tumblebugObjectEntry `json:"objects"`
}

type tumblebugObjectEntry struct {
	Key          string `json:"key"`
	Size         int64  `json:"size"`
	LastModified string `json:"lastModified"`
	ETag         string `json:"etag"`
}
