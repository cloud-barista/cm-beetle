package controller

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// CORSRule represents CORS rule configuration
type CORSRule struct {
	AllowedOrigins []string `json:"allowedOrigins" validate:"required"` // Allowed origins (e.g., ["*"], ["https://example.com"])
	AllowedMethods []string `json:"allowedMethods" validate:"required"` // Allowed HTTP methods (e.g., ["GET", "PUT", "POST"])
	AllowedHeaders []string `json:"allowedHeaders,omitempty"`           // Allowed headers (e.g., ["*"])
	ExposeHeaders  []string `json:"exposeHeaders,omitempty"`            // Headers to expose (e.g., ["ETag"])
	MaxAgeSeconds  int      `json:"maxAgeSeconds,omitempty"`            // Preflight request cache time in seconds
}

// ============================================================================
// Request Models
// ============================================================================

// RecommendObjectStorageRequest represents a request for object storage migration recommendations
type RecommendObjectStorageRequest struct {
	DesiredCloud         cloudmodel.CloudProperty      `json:"desiredCloud" validate:"required"`
	SourceObjectStorages []SourceObjectStorageProperty `json:"sourceObjectStorages" validate:"required,min=1"`
}

// SourceObjectStorageProperty represents source object storage properties from on-premise environment
type SourceObjectStorageProperty struct {
	// Basic identification
	BucketName string `json:"bucketName" validate:"required"` // Actual bucket name

	// Feature settings
	VersioningEnabled bool       `json:"versioningEnabled,omitempty"` // Whether versioning is enabled
	CORSEnabled       bool       `json:"corsEnabled,omitempty"`       // Whether CORS is enabled
	CORSRules         []CORSRule `json:"corsRules,omitempty"`         // CORS rules configuration

	// Capacity information (for cost estimation and recommendations)
	TotalSizeBytes int64 `json:"totalSizeBytes,omitempty"` // Total storage size in bytes
	ObjectCount    int64 `json:"objectCount,omitempty"`    // Total number of objects

	// Access pattern (critical for storage class selection)
	AccessFrequency string `json:"accessFrequency,omitempty"` // "frequent", "infrequent", or "archive"

	// Security settings
	EncryptionEnabled bool `json:"encryptionEnabled,omitempty"` // Whether encryption is enabled
	IsPublic          bool `json:"isPublic,omitempty"`          // Whether bucket has public access

	// Metadata
	Tags         map[string]string `json:"tags,omitempty"`         // Bucket tags
	CreationDate string            `json:"creationDate,omitempty"` // Creation date (RFC3339)
}

// ============================================================================
// Response Models
// ============================================================================

// RecommendObjectStorageResponse represents object storage recommendation response

type RecommendObjectStorageResponse struct {
	ObjectStorageInfo
}

type ObjectStorageInfo struct {
	Status               string                        `json:"status"`
	Description          string                        `json:"description"`
	TargetCloud          cloudmodel.CloudProperty      `json:"targetCloud"`
	TargetObjectStorages []TargetObjectStorageProperty `json:"targetObjectStorages"`
}

// TargetObjectStorageProperty represents recommended target object storage configuration
type TargetObjectStorageProperty struct {
	SourceBucketName  string     `json:"sourceBucketName"`    // Source bucket name for referencing
	BucketName        string     `json:"bucketName"`          // Recommended target bucket name with random suffix by xid pkg
	VersioningEnabled bool       `json:"versioningEnabled"`   // Whether to enable versioning
	CORSEnabled       bool       `json:"corsEnabled"`         // Whether CORS is configured
	CORSRules         []CORSRule `json:"corsRules,omitempty"` // CORS rules configuration
}

// RecommendObjectStorage godoc
// @ID RecommendObjectStorage
// @Summary Recommend an object storage for cloud migration
// @Description Recommend an appropriate object storage for cloud migration
// @Description
// @Description [Note] `desiredProvider` and `desiredRegion` are required.
// @Description - `desiredProvider` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredProvider and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Description
// @Description [Warning] the recommended bucket name may be globally unique.
// @Description - Beetle supports adding a suffix based on the existing bucket name to ensure uniqueness.
// @Description - Suppose that the existing bucket name is unique enough.
// @Description - Generate a suffix based on the existing bucket name.
// @Description - e.g., "my-bucket" -> SHA256 hash -> base64 URL-safe encoding (6 bytes) -> lowercase -> "my-bucket-{suffix}"
// @Tags [Recommendation] Managed middleware (preview)
// @Accept json
// @Produce	json
// @Param request body RecommendObjectStorageRequest true "Specify the your object storage to be migrated"
// @Param desiredProvider query string false "Provider (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,ncp) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} RecommendObjectStorageResponse "The result of recommended object storage"
// @Failure 400 {object} common.SimpleMsg "Invalid request"
// @Failure 500 {object} common.SimpleMsg "Internal server error"
// @Router /recommendation/middleware/objectStorage [post]
func RecommendObjectStorage(c echo.Context) error {
	// [Input]
	// Extract request body
	var req RecommendObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: fmt.Sprintf("Invalid request format: %v", err),
		})
	}

	// Get provider and region from query params (higher priority)
	desiredProvider := c.QueryParam("desiredProvider")
	desiredRegion := c.QueryParam("desiredRegion")

	// Fallback to request body if query params are not provided
	if desiredProvider == "" {
		desiredProvider = req.DesiredCloud.Csp
	}
	if desiredRegion == "" {
		desiredRegion = req.DesiredCloud.Region
	}

	// Validate required parameters
	if desiredProvider == "" || desiredRegion == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "desiredProvider and desiredRegion are required (via query params or request body)",
		})
	}

	// Validate source object storages
	if len(req.SourceObjectStorages) == 0 {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "At least one source object storage must be provided",
		})
	}

	log.Info().
		Str("provider", desiredProvider).
		Str("region", desiredRegion).
		Int("sourceBuckets", len(req.SourceObjectStorages)).
		Msg("Processing object storage recommendation request")

	// [Process]
	// Prepare target object storages based on source configurations
	targetObjectStorages := make([]TargetObjectStorageProperty, 0, len(req.SourceObjectStorages))

	for _, source := range req.SourceObjectStorages {
		// Generate unique bucket name with random suffix
		// Bucket names must be globally unique across all AWS/cloud accounts
		// * Suppose that the existing bucket name is unique enough to generate a new unique name
		uniqueSuffix := createShortSuffix(source.BucketName)
		targetBucketName := fmt.Sprintf("%s-%s", source.BucketName, uniqueSuffix)

		target := TargetObjectStorageProperty{
			SourceBucketName:  source.BucketName,
			BucketName:        targetBucketName,
			VersioningEnabled: source.VersioningEnabled,
			CORSEnabled:       source.CORSEnabled,
			CORSRules:         source.CORSRules,
		}

		targetObjectStorages = append(targetObjectStorages, target)

		log.Debug().
			Str("sourceBucket", source.BucketName).
			Str("targetBucket", targetBucketName).
			Bool("versioning", source.VersioningEnabled).
			Bool("cors", source.CORSEnabled).
			Msg("Generated target object storage recommendation")
	}

	// Build response
	response := RecommendObjectStorageResponse{
		ObjectStorageInfo: ObjectStorageInfo{
			Status:      "success",
			Description: fmt.Sprintf("Successfully recommended %d object storage configuration(s)", len(targetObjectStorages)),
			TargetCloud: cloudmodel.CloudProperty{
				Csp:    desiredProvider,
				Region: desiredRegion,
			},
			TargetObjectStorages: targetObjectStorages,
		},
	}

	log.Info().
		Str("provider", desiredProvider).
		Str("region", desiredRegion).
		Int("targetBuckets", len(targetObjectStorages)).
		Msg("Object storage recommendation completed successfully")

	return c.JSON(http.StatusOK, response)
}

func createShortSuffix(existingBucketName string) string {

	// 1. SHA256 hash (returns 32 bytes)
	hashBytes := sha256.Sum256([]byte(existingBucketName))

	// 2. Base64 URL-Safe encoding (Padding '=' will be removed)
	// 6 bytes (48 bits) are enough to create an 8-character string.
	encoded := base64.URLEncoding.EncodeToString(hashBytes[:6])
	suffix := strings.TrimRight(encoded, "=")

	// 3. Convert to lowercase for S3 bucket name compatibility (S3 only allows lowercase)
	suffix = strings.ToLower(suffix)

	return suffix
}
