package controller

import (
	"fmt"
	"net/http"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func GenerateConnectionName(csp, region string) (string, error) {

	connectionName := fmt.Sprintf("%s-%s", csp, region)

	tbApiConfig := tbclient.ApiConfig{
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
		RestUrl:  config.Tumblebug.RestUrl,
	}
	tbCli := tbclient.NewClient(tbApiConfig)
	_, err := tbCli.GetConnConfig(connectionName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get connection config")
		return "", err
	}

	return connectionName, nil
}

// ============================================================================
// Migration Request/Response Models
// ============================================================================

// MigrateObjectStorageRequest represents a request for object storage migration
type MigrateObjectStorageRequest struct {
	ObjectStorageInfo
}

// MigrateObjectStorage godoc
// @ID MigrateObjectStorage
// @Summary Migrate object storages to cloud
// @Description Migrate object storages to cloud based on recommendation results
// @Description
// @Description [Note] This API creates object storages (buckets) in the target cloud.
// @Description - Input should be the output from RecommendObjectStorage API
// @Description - Connection name format: `{csp}-{region}` (e.g., aws-ap-northeast-2)
// @Description
// @Description [Note]
// @Description * Examples(test result): https://github.com/cloud-barista/cm-beetle/blob/main/docs/test-results-data-migration.md
// @Description
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce	json
// @Param request body MigrateObjectStorageRequest true "Object storage migration request (use RecommendObjectStorage response)"
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 "OK - Object storages created successfully"
// @Failure 400 {object} common.SimpleMsg "Invalid request"
// @Failure 500 {object} common.SimpleMsg "Internal server error"
// @Router /migration/middleware/objectStorage [post]
func MigrateObjectStorage(c echo.Context) error {
	// [Input]
	// Extract request body
	var req MigrateObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: fmt.Sprintf("Invalid request format: %v", err),
		})
	}

	// Validate required parameters
	if req.TargetCloud.Csp == "" || req.TargetCloud.Region == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "targetCloud.csp and targetCloud.region are required",
		})
	}

	// Validate target object storages
	if len(req.TargetObjectStorages) == 0 {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "At least one target object storage must be provided",
		})
	}

	// Generate and validate connection name
	connName, err := GenerateConnectionName(req.TargetCloud.Csp, req.TargetCloud.Region)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate or validate connection name")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: fmt.Sprintf("Invalid connection configuration for %s-%s: %v",
				req.TargetCloud.Csp, req.TargetCloud.Region, err),
		})
	}

	log.Info().
		Str("csp", req.TargetCloud.Csp).
		Str("region", req.TargetCloud.Region).
		Str("connName", connName).
		Int("targetBuckets", len(req.TargetObjectStorages)).
		Msg("Starting object storage migration")

	// [Process]
	// Initialize Tumblebug client
	tbClient := tbclient.NewDefaultClient()

	// Create each object storage (bucket)
	for i, target := range req.TargetObjectStorages {
		log.Debug().
			Int("index", i+1).
			Int("total", len(req.TargetObjectStorages)).
			Str("sourceBucket", target.SourceBucketName).
			Str("targetBucket", target.BucketName).
			Msg("Creating object storage")

		// Create object storage (bucket)
		err := tbClient.CreateObjectStorage(target.BucketName, connName)
		if err != nil {
			log.Error().
				Err(err).
				Str("bucketName", target.BucketName).
				Msg("Failed to create object storage")
			return c.JSON(http.StatusInternalServerError, common.SimpleMsg{
				Message: fmt.Sprintf("Failed to create object storage '%s': %v", target.BucketName, err),
			})
		}

		log.Info().
			Str("sourceBucket", target.SourceBucketName).
			Str("targetBucket", target.BucketName).
			Msg("Successfully created object storage")
	}

	// TODO: Configure versioning for each bucket if versioningEnabled is true
	// - Iterate through req.TargetObjectStorages
	// - For each bucket where target.VersioningEnabled == true:
	//   - Call Tumblebug API to enable versioning on the bucket
	//   - Handle errors appropriately (decide whether to fail or continue)

	// TODO: Configure CORS settings for each bucket if corsEnabled is true
	// - Iterate through req.TargetObjectStorages
	// - For each bucket where target.CORSEnabled == true:
	//   - Extract CORS rules from target.CORSRules
	//   - Call Tumblebug API to configure CORS on the bucket
	//   - Pass allowedOrigins, allowedMethods, allowedHeaders, exposeHeaders, maxAgeSeconds
	//   - Handle errors appropriately (decide whether to fail or continue)

	log.Info().
		Str("csp", req.TargetCloud.Csp).
		Str("region", req.TargetCloud.Region).
		Int("totalBuckets", len(req.TargetObjectStorages)).
		Msg("Object storage migration completed successfully")

	return c.NoContent(http.StatusOK)
}

// ============================================================================
// Object Storage Management APIs
// ============================================================================

// ListObjectStorages godoc
// @ID ListObjectStorages
// @Summary List object storages (buckets)
// @Description Get the list of all object storages (buckets) in the specified cloud service provider and region
// @Description
// @Description [Note] Connection name format: `{csp}-{region}` (e.g., aws-ap-northeast-2)
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param csp query string true "Cloud service provider" Enums(aws,alibaba) default(aws)
// @Param region query string true "Cloud region" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} tbclient.ListAllMyBucketsResult "List of object storages"
// @Failure 400 {object} common.SimpleMsg "Invalid request"
// @Failure 500 {object} common.SimpleMsg "Internal server error"
// @Router /migration/middleware/objectStorage [get]
func ListObjectStorages(c echo.Context) error {
	// Get csp and region from query params
	csp := c.QueryParam("csp")
	region := c.QueryParam("region")

	// Validate required parameters
	if csp == "" || region == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "csp and region query parameters are required",
		})
	}

	// Generate and validate connection name
	connName, err := GenerateConnectionName(csp, region)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate or validate connection name")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: fmt.Sprintf("Invalid connection configuration for %s-%s: %v", csp, region, err),
		})
	}

	log.Info().
		Str("csp", csp).
		Str("region", region).
		Str("connName", connName).
		Msg("Listing object storages")

	// Initialize Tumblebug client
	tbClient := tbclient.NewDefaultClient()

	// List object storages
	result, err := tbClient.ListObjectStorages(connName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list object storages")
		return c.JSON(http.StatusInternalServerError, common.SimpleMsg{
			Message: fmt.Sprintf("Failed to list object storages: %v", err),
		})
	}

	log.Info().
		Str("csp", csp).
		Str("region", region).
		Int("bucketCount", len(result.Buckets.Bucket)).
		Msg("Successfully listed object storages")

	return c.JSON(http.StatusOK, result)
}

// GetObjectStorage godoc
// @ID GetObjectStorage
// @Summary Get object storage (bucket) details
// @Description Get details of a specific object storage (bucket)
// @Description
// @Description [Note] Connection name format: `{csp}-{region}` (e.g., aws-ap-northeast-2)
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param objectStorageName path string true "Object Storage Name (bucket name)"
// @Param csp query string true "Cloud service provider" Enums(aws,alibaba) default(aws)
// @Param region query string true "Cloud region" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 {object} tbclient.ListBucketResult "Object storage details"
// @Failure 400 {object} common.SimpleMsg "Invalid request"
// @Failure 404 {object} common.SimpleMsg "Object storage not found"
// @Failure 500 {object} common.SimpleMsg "Internal server error"
// @Router /migration/middleware/objectStorage/{objectStorageName} [get]
func GetObjectStorage(c echo.Context) error {
	// Get path parameter
	objectStorageName := c.Param("objectStorageName")
	if objectStorageName == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "objectStorageName is required",
		})
	}

	// Get csp and region from query params
	csp := c.QueryParam("csp")
	region := c.QueryParam("region")

	// Validate required parameters
	if csp == "" || region == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "csp and region query parameters are required",
		})
	}

	// Generate and validate connection name
	connName, err := GenerateConnectionName(csp, region)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate or validate connection name")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: fmt.Sprintf("Invalid connection configuration for %s-%s: %v", csp, region, err),
		})
	}

	log.Info().
		Str("csp", csp).
		Str("region", region).
		Str("connName", connName).
		Str("objectStorageName", objectStorageName).
		Msg("Getting object storage details")

	// Initialize Tumblebug client
	tbClient := tbclient.NewDefaultClient()

	// Get object storage details
	result, err := tbClient.GetObjectStorage(objectStorageName, connName)
	if err != nil {
		log.Error().Err(err).
			Str("objectStorageName", objectStorageName).
			Msg("Failed to get object storage")
		return c.JSON(http.StatusInternalServerError, common.SimpleMsg{
			Message: fmt.Sprintf("Failed to get object storage '%s': %v", objectStorageName, err),
		})
	}

	log.Info().
		Str("csp", csp).
		Str("region", region).
		Str("objectStorageName", objectStorageName).
		Msg("Successfully retrieved object storage details")

	return c.JSON(http.StatusOK, result)
}

// ExistObjectStorage godoc
// @ID ExistObjectStorage
// @Summary Check object storage (bucket) existence
// @Description Check if a specific object storage (bucket) exists
// @Description
// @Description [Note]
// @Description - Connection name format: `{csp}-{region}` (e.g., aws-ap-northeast-2)
// @Description - Returns 200 OK if the bucket exists, 404 Not Found if it doesn't exist
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param objectStorageName path string true "Object Storage Name (bucket name)"
// @Param csp query string true "Cloud service provider" Enums(aws,alibaba) default(aws)
// @Param region query string true "Cloud region" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 "OK - Object storage exists"
// @Failure 400 {object} common.SimpleMsg "Invalid request"
// @Failure 404 {object} common.SimpleMsg "Object storage not found"
// @Failure 500 {object} common.SimpleMsg "Internal server error"
// @Router /migration/middleware/objectStorage/{objectStorageName} [head]
func ExistObjectStorage(c echo.Context) error {
	// Get path parameter
	objectStorageName := c.Param("objectStorageName")
	if objectStorageName == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "objectStorageName is required",
		})
	}

	// Get csp and region from query params
	csp := c.QueryParam("csp")
	region := c.QueryParam("region")

	// Validate required parameters
	if csp == "" || region == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "csp and region query parameters are required",
		})
	}

	// Generate and validate connection name
	connName, err := GenerateConnectionName(csp, region)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate or validate connection name")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: fmt.Sprintf("Invalid connection configuration for %s-%s: %v", csp, region, err),
		})
	}

	log.Info().
		Str("csp", csp).
		Str("region", region).
		Str("connName", connName).
		Str("objectStorageName", objectStorageName).
		Msg("Checking object storage existence")

	// Initialize Tumblebug client
	tbClient := tbclient.NewDefaultClient()

	// Check object storage existence
	exists, err := tbClient.ExistObjectStorage(objectStorageName, connName)
	if err != nil {
		log.Error().Err(err).
			Str("objectStorageName", objectStorageName).
			Msg("Failed to check object storage existence")
		return c.JSON(http.StatusInternalServerError, common.SimpleMsg{
			Message: fmt.Sprintf("Failed to check object storage '%s' existence: %v", objectStorageName, err),
		})
	}

	log.Info().
		Str("csp", csp).
		Str("region", region).
		Str("objectStorageName", objectStorageName).
		Bool("exists", exists).
		Msg("Successfully checked object storage existence")

	if exists {
		return c.NoContent(http.StatusOK)
	}

	return c.JSON(http.StatusNotFound, common.SimpleMsg{
		Message: fmt.Sprintf("Object storage '%s' not found", objectStorageName),
	})
}

// DeleteObjectStorage godoc
// @ID DeleteObjectStorage
// @Summary Delete object storage (bucket)
// @Description Delete a specific object storage (bucket)
// @Description
// @Description [Note]
// @Description - Connection name format: `{csp}-{region}` (e.g., aws-ap-northeast-2)
// @Description - The bucket must be empty before deletion
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param objectStorageName path string true "Object Storage Name (bucket name)"
// @Param csp query string true "Cloud service provider" Enums(aws,alibaba) default(aws)
// @Param region query string true "Cloud region" default(ap-northeast-2)
// @Param X-Request-Id header string false "Custom request ID (NOTE: It will be used as a trace ID.)"
// @Success 200 "OK - Object storage deleted successfully"
// @Failure 400 {object} common.SimpleMsg "Invalid request"
// @Failure 404 {object} common.SimpleMsg "Object storage not found"
// @Failure 500 {object} common.SimpleMsg "Internal server error"
// @Router /migration/middleware/objectStorage/{objectStorageName} [delete]
func DeleteObjectStorage(c echo.Context) error {
	// Get path parameter
	objectStorageName := c.Param("objectStorageName")
	if objectStorageName == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "objectStorageName is required",
		})
	}

	// Get csp and region from query params
	csp := c.QueryParam("csp")
	region := c.QueryParam("region")

	// Validate required parameters
	if csp == "" || region == "" {
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: "csp and region query parameters are required",
		})
	}

	// Generate and validate connection name
	connName, err := GenerateConnectionName(csp, region)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate or validate connection name")
		return c.JSON(http.StatusBadRequest, common.SimpleMsg{
			Message: fmt.Sprintf("Invalid connection configuration for %s-%s: %v", csp, region, err),
		})
	}

	log.Info().
		Str("csp", csp).
		Str("region", region).
		Str("connName", connName).
		Str("objectStorageName", objectStorageName).
		Msg("Deleting object storage")

	// Initialize Tumblebug client
	tbClient := tbclient.NewDefaultClient()

	// Delete object storage
	err = tbClient.DeleteObjectStorage(objectStorageName, connName)
	if err != nil {
		log.Error().Err(err).
			Str("objectStorageName", objectStorageName).
			Msg("Failed to delete object storage")
		return c.JSON(http.StatusInternalServerError, common.SimpleMsg{
			Message: fmt.Sprintf("Failed to delete object storage '%s': %v", objectStorageName, err),
		})
	}

	log.Info().
		Str("csp", csp).
		Str("region", region).
		Str("objectStorageName", objectStorageName).
		Msg("Successfully deleted object storage")

	return c.NoContent(http.StatusOK)
}
