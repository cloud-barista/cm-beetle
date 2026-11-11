package controller

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// Dummy blank identifier to retain tbmodel import for Swagger annotations without logging
var _ = tbmodel.ObjectStorageSupportResponse{}

// GetObjectStorageSupport godoc
// @ID GetObjectStorageSupport
// @Summary Get CSP feature support map for object storage (Proxied to Tumblebug)
// @Description Forward object storage support request to CB-Tumblebug via Proxy (Returns ObjectStorageSupportResponse)
// @Tags [Recommendation] Managed Object Storage
// @Accept json
// @Produce json
// @Param providerName query string false "CSP Type filter (e.g., aws, azure, gcp)"
// @Success 200 {object} tbmodel.ObjectStorageSupportResponse
// @Router /recommendation/middleware/objectStorage/support [get]
func GetObjectStorageSupport(c echo.Context) error {
	providerName := c.QueryParam("providerName")

	result, err := tbclient.NewSession().SetContext(c.Request().Context()).GetObjectStorageSupport(providerName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get object storage support matrix")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to get object storage support matrix: %v", err)))
	}

	return c.JSON(http.StatusOK, result)
}

// ============================================================================
// Request Models
// ============================================================================

// RecommendObjectStorageRequest represents a request for object storage migration recommendations
type RecommendObjectStorageRequest struct {
	DesiredCloud         storagemodel.CloudProperty                 `json:"desiredCloud" validate:"required"`
	SourceObjectStorages []storagemodel.SourceObjectStorageProperty `json:"sourceObjectStorages" validate:"required,min=1"`
}

// RecommendObjectStorage godoc
// @ID RecommendObjectStorage
// @Summary Recommend an object storage for cloud migration
// @Description Recommend an appropriate object storage configuration for cloud migration.
// @Description
// @Description [Note] `desiredCsp` and `desiredRegion` are required.
// @Description - `desiredCsp` and `desiredRegion` can be provided via query parameters or `desiredCloud` in the request body.
// @Description - Query parameters take precedence over `desiredCloud` values in the request body when both are set.
// @Description
// @Description [Note] The recommended bucket name uses the default base pattern (`os-01`, `os-02`, ...).
// @Description - `bucketName` becomes the CB-Tumblebug `osId` (`name`) within the target namespace.
// @Description - CB-Tumblebug internally generates a globally unique `uid` (`cspResourceName`) as the actual cloud bucket name.
// @Description - To apply a deployment prefix via late binding, pass the `nameSeed` query parameter to `POST /migration/middleware/ns/{nsId}/objectStorage?nameSeed=xxx`.
// @Description
// @Tags [Recommendation] Managed Object Storage
// @Accept json
// @Produce	json
// @Param request body RecommendObjectStorageRequest true "Specify the your object storage to be migrated"
// @Param desiredCsp query string false "CSP (e.g., aws, azure, gcp)" Enums(aws,azure,gcp,alibaba,tencent,ibm,openstack,ncp,nhn,kt) default(aws)
// @Param desiredRegion query string false "Region (e.g., ap-northeast-2)" default(ap-northeast-2)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this recommendation asynchronously (RFC 7240)" Enums(respond-async)
// @Success 200 {object} model.ApiResponse[storagemodel.RecommendedObjectStorage] "Successfully recommended object storage"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Recommendation started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during recommendation"
// @Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"
// @Router /recommendation/middleware/objectStorage [post]
func RecommendObjectStorage(c echo.Context) error {

	// Extract request body.
	var req RecommendObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// Read desired CSP and region from query params first, falling back to request body.
	desiredCsp := strings.ToLower(strings.TrimSpace(c.QueryParam("desiredCsp")))
	desiredRegion := strings.TrimSpace(c.QueryParam("desiredRegion"))

	if desiredCsp == "" {
		desiredCsp = strings.ToLower(strings.TrimSpace(req.DesiredCloud.Csp))
	}
	if desiredRegion == "" {
		desiredRegion = strings.TrimSpace(req.DesiredCloud.Region)
	}

	// Validate required target cloud parameters.
	if desiredCsp == "" || desiredRegion == "" {
		log.Warn().Msg("desiredCsp and desiredRegion are required (via query params or request body)")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP and region required"))
	}

	// Validate source object storages.
	if len(req.SourceObjectStorages) == 0 {
		log.Warn().Msg("At least one source object storage must be provided")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one source bucket required"))
	}

	log.Info().
		Str("desiredCsp", desiredCsp).
		Str("region", desiredRegion).
		Int("sourceBuckets", len(req.SourceObjectStorages)).
		Msg("Processing object storage recommendation request")

	sources := req.SourceObjectStorages
	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(c.Request().Context(), reqID, func(ctx context.Context) (storagemodel.RecommendedObjectStorage, error) {
			return buildObjectStorageRecommendation(ctx, desiredCsp, desiredRegion, sources), nil
		})
		if !started {
			c.Response().Header().Set("Retry-After", "5")
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many async jobs in progress; retry shortly, or retry without Prefer: respond-async"))
		}
		c.Response().Header().Set("Preference-Applied", "respond-async")
		return c.JSON(http.StatusAccepted, model.SuccessResponseWithMessage(
			model.AsyncJobResponse{
				ReqID:     reqID,
				Status:    common.RequestStatusHandling,
				StatusURL: fmt.Sprintf("/beetle/request/%s", reqID),
			},
			"Recommendation started. Use GET /request/{reqId} to check status."))
	}

	objectStorageInfo := buildObjectStorageRecommendation(c.Request().Context(), desiredCsp, desiredRegion, sources)

	// [Output]
	successMsg := fmt.Sprintf("Recommended %d object storage(s) for %s %s",
		len(objectStorageInfo.TargetObjectStorages), desiredCsp, desiredRegion)
	res := model.SuccessResponseWithMessage(objectStorageInfo, successMsg)

	return c.JSON(http.StatusOK, res)
}

func buildObjectStorageRecommendation(ctx context.Context, desiredCsp, desiredRegion string, sources []storagemodel.SourceObjectStorageProperty) storagemodel.RecommendedObjectStorage {
	// Fetch CSP feature support from Tumblebug using normalized lowercase CSP name.
	supportResp, err := tbclient.NewSession().SetContext(ctx).GetObjectStorageSupport(desiredCsp)
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
	targetObjectStorages := make([]storagemodel.TargetObjectStorage, 0, len(sources))

	for i, source := range sources {
		targetBucketName := fmt.Sprintf("os-%02d", i+1)

		versioningEnabled := source.VersioningEnabled
		corsEnabled := source.CORSEnabled
		corsRule := source.CORSRule

		// Adjust based on CSP support and emit warnings.
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
	}

	status := "success"
	if len(warnings) > 0 {
		status = "partial"
	}

	log.Info().
		Str("desiredCsp", desiredCsp).
		Str("region", desiredRegion).
		Int("targetBuckets", len(targetObjectStorages)).
		Msg("Object storage recommendation completed successfully")

	return storagemodel.RecommendedObjectStorage{
		Status:               status,
		Description:          fmt.Sprintf("Successfully recommended %d object storage configuration(s)", len(targetObjectStorages)),
		Warnings:             warnings,
		TargetCloud:          storagemodel.CloudProperty{Csp: desiredCsp, Region: desiredRegion},
		TargetObjectStorages: targetObjectStorages,
	}
}
