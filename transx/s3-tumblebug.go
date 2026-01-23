package transx

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
	username string
	password string
	jwtToken string
	// apiKey       string // TODO: Not tested yet
	// apiKeyHeader string // TODO: Not tested yet
	// oauthToken   string // TODO: Not tested yet
	// oauthType    string // TODO: Not tested yet
}

// NewTumblebugProvider creates a new TumblebugProvider.
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

	// Extract auth credentials from config
	var username, password, jwtToken string
	if config.Auth != nil {
		if config.Auth.AuthType == AuthTypeBasic && config.Auth.Basic != nil {
			username = config.Auth.Basic.Username
			password = config.Auth.Basic.Password
		} else if config.Auth.AuthType == AuthTypeJWT && config.Auth.JWT != nil {
			jwtToken = config.Auth.JWT.Token
		}
	}

	p := &TumblebugProvider{
		endpoint: strings.TrimSuffix(config.Endpoint, "/"),
		nsId:     config.NsId,
		osId:     config.OsId,
		expires:  expires,
		username: username,
		password: password,
		jwtToken: jwtToken,
	}

	return p, nil
}

// GeneratePresignedURL generates a presigned URL via CB-Tumblebug API.
// Uses the new endpoint: GET /ns/{nsId}/resources/objectStorage/{osId}/object/{objectKey}
// Query parameters:
//   - operation: "upload" or "download"
//   - expires: expiration time in seconds (default: 3600)
func (p *TumblebugProvider) GeneratePresignedURL(action, key string) (string, error) {
	// URL encode the object key to handle special characters and paths
	encodedKey := url.PathEscape(key)

	// GET /ns/{nsId}/resources/objectStorage/{osId}/object/{objectKey}?operation=xxx&expires=xxx
	apiURL := fmt.Sprintf("%s/ns/%s/resources/objectStorage/%s/object/%s?operation=%s&expires=%d",
		p.endpoint, p.nsId, p.osId, encodedKey, action, p.expires)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Apply authentication
	if p.username != "" && p.password != "" {
		req.SetBasicAuth(p.username, p.password)
	} else if p.jwtToken != "" {
		req.Header.Set("Authorization", "Bearer "+p.jwtToken)
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
	// GET /ns/{nsId}/resources/objectStorage/{osId}
	// Note: Tumblebug API returns all objects, client-side filtering applied
	apiURL := fmt.Sprintf("%s/ns/%s/resources/objectStorage/%s",
		p.endpoint, p.nsId, p.osId)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Apply authentication
	if p.username != "" && p.password != "" {
		req.SetBasicAuth(p.username, p.password)
	} else if p.jwtToken != "" {
		req.Header.Set("Authorization", "Bearer "+p.jwtToken)
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

	// Filter objects by prefix (client-side filtering)
	var objects []ObjectInfo
	for _, obj := range response.Contents {
		if prefix == "" || strings.HasPrefix(obj.Key, prefix) {
			objects = append(objects, ObjectInfo{
				Key:          obj.Key,
				Size:         obj.Size,
				LastModified: obj.LastModified,
				ETag:         obj.ETag,
			})
		}
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
	Contents []tumblebugObjectEntry `json:"contents"`
}

type tumblebugObjectEntry struct {
	Key          string `json:"key"`
	Size         int64  `json:"size"`
	LastModified string `json:"lastModified"`
	ETag         string `json:"eTag"`
	StorageClass string `json:"storageClass,omitempty"`
}
