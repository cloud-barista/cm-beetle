package controller

import (
	"fmt"
	"net/http"

	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// Request Models
// ============================================================================

// RecommendObjectStorageRequest represents a request for object storage migration recommendations
type RecommendObjectStorageRequest struct {
	NameSeed             string                             `json:"nameSeed,omitempty" example:"my"` // Base string for bucket name prefix (e.g., 'my' -> 'my-os-01'); applied at migration time
	DesiredCloud         storagemodel.CloudProperty         `json:"desiredCloud" validate:"required"`
	SourceObjectStorages []storagemodel.SourceObjectStorage `json:"sourceObjectStorages" validate:"required,min=1"`
}

// RecommendObjectStorage godoc
// @ID RecommendObjectStorage
// @Summary Recommend an object storage for cloud migration
// @Description Recommend an appropriate object storage for cloud migration
// @Description
// @Description [Note] `desiredCsp` and `desiredRegion` are required.
// @Description - `desiredCsp` and `desiredRegion` can set on the query parameter or the request body.
// @Description
// @Description - If desiredCsp and desiredRegion are set on request body, the values in the query parameter will be ignored.
// @Description
// @Description [Note] The recommended bucket name uses a default pattern (`mig-bucket-01`, `mig-bucket-02`, ...).
// @Description - Bucket names must be globally unique across all accounts in the target cloud provider.
// @Description - CB-Tumblebug internally generates a uid and uses it as the actual bucket name in the cloud.
// @Description - The `bucketName` field in the recommendation result represents the intended name, not the final cloud resource name.
// @Description
// @Description [Note] `nameSeed` enables dynamic naming via **Late Binding**.
// @Description - Set `nameSeed` (e.g., `my`) to prefix bucket names at migration time: `my-os-01`.
// @Description - The recommendation result stores base names only; the prefix is applied when `MigrateObjectStorage` is called.
// @Description
// @Tags [Recommendation] Managed middleware (preview)
// @Accept json
// @Produce	json
// @Param request body RecommendObjectStorageRequest true "Specify the your object storage to be migrated"
// @Param desiredCsp query string false "CSP (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,tencent,ibm,openstack,ncp,nhn,kt) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[storagemodel.RecommendedObjectStorage] "Successfully recommended object storage"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Router /recommendation/middleware/objectStorage [post]
func RecommendObjectStorage(c echo.Context) error {

	// [Input]
	// Extract request body
	var req RecommendObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// Get csp and region from query params (higher priority)
	desiredCsp := c.QueryParam("desiredCsp")
	desiredRegion := c.QueryParam("desiredRegion")

	// Fallback to request body if query params are not provided
	if desiredCsp == "" {
		desiredCsp = req.DesiredCloud.Csp
	}
	if desiredRegion == "" {
		desiredRegion = req.DesiredCloud.Region
	}

	// Validate required parameters
	if desiredCsp == "" || desiredRegion == "" {
		log.Warn().Msg("desiredCsp and desiredRegion are required (via query params or request body)")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP and region required"))
	}

	// Validate source object storages
	if len(req.SourceObjectStorages) == 0 {
		log.Warn().Msg("At least one source object storage must be provided")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one source bucket required"))
	}

	log.Info().
		Str("desiredCsp", desiredCsp).
		Str("region", desiredRegion).
		Int("sourceBuckets", len(req.SourceObjectStorages)).
		Msg("Processing object storage recommendation request")

	// [Process]
	// Fetch CSP feature support from Tumblebug
	supportResp, err := tbclient.NewSession().GetObjectStorageSupport(desiredCsp)
	if err != nil {
		log.Warn().Err(err).Str("csp", desiredCsp).Msg("Failed to fetch CSP object storage support info; proceeding without support check")
	}

	support, hasSupportInfo := supportResp.Supports[desiredCsp]
	if !hasSupportInfo {
		log.Warn().Str("csp", desiredCsp).Msg("No support info found for CSP; all features assumed supported")
	}

	corsSupported := !hasSupportInfo || support.Cors
	versioningSupported := !hasSupportInfo || support.Versioning

	log.Debug().
		Str("csp", desiredCsp).
		Bool("corsSupported", corsSupported).
		Bool("versioningSupported", versioningSupported).
		Msg("CSP object storage feature support")

	var warnings []string
	targetObjectStorages := make([]storagemodel.TargetObjectStorage, 0, len(req.SourceObjectStorages))

	for i, source := range req.SourceObjectStorages {
		targetBucketName := fmt.Sprintf("os-%02d", i+1)

		versioningEnabled := source.VersioningEnabled
		corsEnabled := source.CORSEnabled
		corsRule := source.CORSRule

		// Adjust based on CSP support and emit warnings
		if source.VersioningEnabled && !versioningSupported {
			versioningEnabled = false
			msg := fmt.Sprintf("Bucket '%s': versioning disabled (not supported on %s)", source.BucketName, desiredCsp)
			warnings = append(warnings, msg)
			log.Info().Msg(msg)
		}
		if source.CORSEnabled && !corsSupported {
			corsEnabled = false
			corsRule = nil
			msg := fmt.Sprintf("Bucket '%s': CORS disabled (not supported on %s)", source.BucketName, desiredCsp)
			warnings = append(warnings, msg)
			log.Info().Msg(msg)
		}

		target := storagemodel.TargetObjectStorage{
			SourceBucketName: source.BucketName,
			BucketName:       targetBucketName,
			BucketSpecProperty: storagemodel.BucketSpecProperty{
				VersioningEnabled: versioningEnabled,
				CORSEnabled:       corsEnabled,
				CORSRule:          corsRule,
			},
		}

		targetObjectStorages = append(targetObjectStorages, target)

		log.Debug().
			Str("sourceBucket", source.BucketName).
			Str("targetBucket", targetBucketName).
			Bool("versioning", versioningEnabled).
			Bool("cors", corsEnabled).
			Msg("Generated target object storage recommendation")
	}

	// Determine overall status
	status := "success"
	if len(warnings) > 0 {
		status = "partial"
	}

	// Build response
	objectStorageInfo := storagemodel.RecommendedObjectStorage{
		NameSeed:             req.NameSeed,
		Status:               status,
		Description:          fmt.Sprintf("Successfully recommended %d object storage configuration(s)", len(targetObjectStorages)),
		Warnings:             warnings,
		TargetCloud:          storagemodel.CloudProperty{Csp: desiredCsp, Region: desiredRegion},
		TargetObjectStorages: targetObjectStorages,
	}

	log.Info().
		Str("desiredCsp", desiredCsp).
		Str("region", desiredRegion).
		Int("targetBuckets", len(targetObjectStorages)).
		Msg("Object storage recommendation completed successfully")

	// [Output]
	successMsg := fmt.Sprintf("Recommended %d object storage(s) for %s %s",
		len(targetObjectStorages), desiredCsp, desiredRegion)
	res := model.SuccessResponseWithMessage(objectStorageInfo, successMsg)

	return c.JSON(http.StatusOK, res)
}
