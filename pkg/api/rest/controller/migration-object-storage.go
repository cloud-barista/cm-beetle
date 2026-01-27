package controller

import (
	"fmt"
	"net/http"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func GenerateConnectionName(csp, region string) (string, error) {

	connectionName := fmt.Sprintf("%s-%s", csp, region)

	tbSess := tbclient.NewSession()
	_, err := tbSess.GetConnConfig(connectionName)
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
// @Description [Note]
// @Description - This API creates object storages (buckets) in the target cloud within the specified namespace
// @Description - Input should be the output from RecommendObjectStorage API
// @Description - Connection name is automatically generated from CSP and region in the request body
// @Description
// @Description [Examples]
// @Description * Test results: https://github.com/cloud-barista/cm-beetle/blob/main/docs/test-results-data-migration.md
// @Description
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce	json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param request body MigrateObjectStorageRequest true "Object storage migration request (use RecommendObjectStorage response)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 201 "Created - Object storages created successfully"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during object storage creation"
// @Router /migration/middleware/ns/{nsId}/objectStorage [post]
func MigrateObjectStorage(c echo.Context) error {
	// Get path parameter
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	// [Input]
	// Extract request body
	var req MigrateObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// Validate required parameters
	if req.TargetCloud.Csp == "" || req.TargetCloud.Region == "" {
		log.Warn().Msg("CSP and region are required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP and region required"))
	}

	// Validate target object storages
	if len(req.TargetObjectStorages) == 0 {
		log.Warn().Msg("At least one target object storage must be provided")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one target bucket required"))
	}

	// Generate and validate connection name
	connName, err := GenerateConnectionName(req.TargetCloud.Csp, req.TargetCloud.Region)
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate or validate connection name")
		errorMsg := fmt.Sprintf("Invalid cloud configuration: %s %s", req.TargetCloud.Csp, req.TargetCloud.Region)
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(errorMsg))
	}

	log.Info().
		Str("csp", req.TargetCloud.Csp).
		Str("region", req.TargetCloud.Region).
		Str("connName", connName).
		Int("targetBuckets", len(req.TargetObjectStorages)).
		Msg("Starting object storage migration")

	// [Process]
	// Initialize Tumblebug session
	// tbSess := tbclient.NewSession()

	// Create each object storage (bucket)
	for i, target := range req.TargetObjectStorages {
		log.Debug().
			Int("index", i+1).
			Int("total", len(req.TargetObjectStorages)).
			Str("sourceBucket", target.SourceBucketName).
			Str("targetBucket", target.BucketName).
			Msg("Creating object storage")

		// Create object storage (bucket)
		req := tbmodel.ObjectStorageCreateRequest{
			BucketName:     target.BucketName,
			ConnectionName: connName,
			Description:    "Created by CM-Beetle",
		}
		_, err := tbclient.NewSession().CreateObjectStorage(nsId, req)
		if err != nil {
			log.Error().
				Err(err).
				Str("bucketName", target.BucketName).
				Msg("Failed to create object storage")
			errorMsg := fmt.Sprintf("Failed to create obejct storage '%s'", target.BucketName)
			return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(errorMsg))
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

	return c.NoContent(http.StatusCreated)
}

// ============================================================================
// Object Storage Management APIs
// ============================================================================

// ListObjectStorages godoc
// @ID ListObjectStorages
// @Summary List object storages (buckets)
// @Description Get the list of all object storages (buckets) in the namespace
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[tbclient.ObjectStorageListResponse] "Successfully retrieved object storage list"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during list operation"
// @Router /migration/middleware/ns/{nsId}/objectStorage [get]
func ListObjectStorages(c echo.Context) error {
	// Get path parameter
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	log.Info().
		Str("nsId", nsId).
		Msg("Listing object storages")

	// Initialize Tumblebug session
	tbSess := tbclient.NewSession()

	// List object storages
	result, err := tbSess.ListObjectStorages(nsId, "", "", "")
	if err != nil {
		errorMsg := fmt.Sprintf("Failed to list object storages: %v", err)
		log.Error().Err(err).Msg(errorMsg)
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(errorMsg))
	}

	log.Info().
		Str("nsId", nsId).
		Int("bucketCount", len(result.ObjectStorage)).
		Msg("Successfully listed object storages")

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}

// GetObjectStorage godoc
// @ID GetObjectStorage
// @Summary Get object storage (bucket) details
// @Description Get details of a specific object storage (bucket)
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[tbmodel.ObjectStorageInfo] "Successfully retrieved object storage details"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during get operation"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId} [get]
func GetObjectStorage(c echo.Context) error {
	// Get path parameters
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	osId := c.Param("osId")
	if osId == "" {
		log.Warn().Msg("osId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("osId required"))
	}

	log.Info().
		Str("nsId", nsId).
		Str("osId", osId).
		Msg("Getting object storage details")

	// Initialize Tumblebug session
	tbSess := tbclient.NewSession()

	// Get object storage details
	result, err := tbSess.GetObjectStorage(nsId, osId)
	if err != nil {
		log.Error().Err(err).
			Str("osId", osId).
			Msg("Failed to get object storage")

		// Check if error is due to not found
		if strings.Contains(strings.ToLower(err.Error()), "not found") || strings.Contains(strings.ToLower(err.Error()), "does not exist") {
			errorMsg := fmt.Sprintf("Object storage '%s' not found", osId)
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(errorMsg))
		}

		errorMsg := fmt.Sprintf("Failed to get object storage '%s': %v", osId, err)
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(errorMsg))
	}

	log.Info().
		Str("nsId", nsId).
		Str("osId", osId).
		Msg("Successfully retrieved object storage details")

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}

// ExistObjectStorage godoc
// @ID ExistObjectStorage
// @Summary Check object storage (bucket) existence
// @Description Check if a specific object storage (bucket) exists
// @Description
// @Description [Note]
// @Description - Returns 200 OK if the bucket exists, 404 Not Found if it doesn't exist
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 "OK - Object storage exists"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during existence check"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId} [head]
func ExistObjectStorage(c echo.Context) error {
	// Get path parameters
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	osId := c.Param("osId")
	if osId == "" {
		log.Warn().Msg("osId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("osId required"))
	}

	log.Info().
		Str("nsId", nsId).
		Str("osId", osId).
		Msg("Checking object storage existence")

	// Initialize Tumblebug session
	tbSess := tbclient.NewSession()

	// Check object storage existence
	exists, err := tbSess.ExistObjectStorage(nsId, osId)
	if err != nil {
		log.Error().Err(err).
			Str("osId", osId).
			Msg("Failed to check object storage existence")

		// Check if error is due to not found
		if strings.Contains(strings.ToLower(err.Error()), "not found") || strings.Contains(strings.ToLower(err.Error()), "does not exist") {
			msg := fmt.Sprintf("Object storage '%s' not found", osId)
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(msg))
		}

		errorMsg := fmt.Sprintf("Failed to check object storage '%s' existence: %v", osId, err)
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(errorMsg))
	}

	log.Info().
		Str("nsId", nsId).
		Str("osId", osId).
		Bool("exists", exists).
		Msg("Successfully checked object storage existence")

	if !exists {
		msg := fmt.Sprintf("Object storage '%s' not found", osId)
		log.Info().Msg(msg)
		return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(msg))
	}

	return c.NoContent(http.StatusOK)
}

// DeleteObjectStorage godoc
// @ID DeleteObjectStorage
// @Summary Delete object storage (bucket)
// @Description Delete a specific object storage (bucket)
// @Description
// @Description [Note]
// @Description - The bucket must be empty before deletion
// @Tags [Migration] Managed middleware (preview)
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 204 "Object storage deleted successfully"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during deletion"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId} [delete]
func DeleteObjectStorage(c echo.Context) error {
	// Get path parameters
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	osId := c.Param("osId")
	if osId == "" {
		log.Warn().Msg("osId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("osId required"))
	}

	log.Info().
		Str("nsId", nsId).
		Str("osId", osId).
		Msg("Deleting object storage")

	// Initialize Tumblebug session
	tbSess := tbclient.NewSession()

	// Delete object storage
	err := tbSess.DeleteObjectStorage(nsId, osId)
	if err != nil {
		log.Error().Err(err).
			Str("osId", osId).
			Msg("Failed to delete object storage")

		// Check if error is due to not found
		if strings.Contains(strings.ToLower(err.Error()), "not found") || strings.Contains(strings.ToLower(err.Error()), "does not exist") {
			msg := fmt.Sprintf("Object storage '%s' not found", osId)
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(msg))
		}

		errorMsg := fmt.Sprintf("Failed to delete object storage '%s': %v", osId, err)
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(errorMsg))
	}

	log.Info().
		Str("nsId", nsId).
		Str("osId", osId).
		Msg("Successfully deleted object storage")

	return c.NoContent(http.StatusNoContent)
}
