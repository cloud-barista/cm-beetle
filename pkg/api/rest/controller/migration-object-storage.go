/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package controller has handlers and their request/response bodies for migration APIs
package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/migration"
	"github.com/cloud-barista/cm-beetle/pkg/ratelimit"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// mapObjectStorageUpstreamStatus maps upstream Tumblebug error messages to appropriate HTTP status codes.
func mapObjectStorageUpstreamStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}
	msg := strings.ToLower(err.Error())
	switch {
	case strings.Contains(msg, "invalid cloud configuration"),
		strings.Contains(msg, "api error: 400"),
		strings.Contains(msg, "bad request"):
		return http.StatusBadRequest
	case strings.Contains(msg, "api error: 409"),
		strings.Contains(msg, "already exists"),
		strings.Contains(msg, "conflict"),
		strings.Contains(msg, "bucketnotempty"),
		strings.Contains(msg, "not empty"):
		return http.StatusConflict
	case strings.Contains(msg, "api error: 404"),
		strings.Contains(msg, "not found"),
		strings.Contains(msg, "does not exist"):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

// isValidObjectStorageDeleteOption checks whether option is in the allowed set ("", "empty", "force").
func isValidObjectStorageDeleteOption(option string) bool {
	return option == "" || option == "empty" || option == "force"
}

// ============================================================================
// Request Models
// ============================================================================

// MigrateObjectStorageRequest represents a request for object storage migration
type MigrateObjectStorageRequest struct {
	storagemodel.RecommendedObjectStorage
}

// ============================================================================
// Object Storage Migration API
// ============================================================================

// MigrateObjectStorage godoc
// @ID MigrateObjectStorage
// @Summary Migrate object storages to cloud
// @Description Migrate object storages (buckets) to the target cloud based on recommendation results.
// @Description
// @Description [Provisioning & Partial Success Behavior]
// @Description - Each bucket in `targetObjectStorages` is attempted independently (input should be the `RecommendObjectStorage` output).
// @Description - Connection name is derived automatically from `targetCloud.csp` and `targetCloud.region`.
// @Description - If at least one bucket is created, returns `201 Created` with `MigratedObjectStorageResult` (`status: "created"` when all buckets and Versioning/CORS settings succeed, or `status: "partial"` with `warnings` if any bucket or feature configuration was skipped or failed).
// @Description - If all target buckets fail (for example, `409 Conflict` when a bucket with the same `osId`/`name` already exists in the namespace), the corresponding error status code is returned.
// @Description
// @Description [Late-Binding Naming (`nameSeed`)]
// @Description - If `nameSeed` query param is set (e.g., `?nameSeed=my`), bucket names are prefixed at migration time (`my-os-01`).
// @Description - If `nameSeed` is omitted, bucket names from `targetObjectStorages` are used as-is.
// @Description
// @Description [Day-2 Operations Scope]
// @Description - CM-Beetle covers target bucket provisioning, initial Versioning/CORS setup, and post-migration verification.
// @Description - For post-creation configuration updates (CORS, Versioning), presigned URLs, or metadata reconcile/prune operations, use CB-Tumblebug (`/tumblebug/ns/{nsId}/resources/objectStorage/...`).
// @Description
// @Description [Async Mode]
// @Description - Send header `Prefer: respond-async` to execute asynchronously (`202 Accepted` with `reqId`, then poll `GET /request/{reqId}`).
// @Tags [Migration] Managed Object Storage
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param nameSeed query string false "Optional prefix for bucket names (e.g., 'my' -> 'my-os-01'). Applied at migration time."
// @Param request body MigrateObjectStorageRequest true "Object storage migration request (use RecommendObjectStorage response)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param Prefer header string false "Set to 'respond-async' to run this migration asynchronously (RFC 7240)" Enums(respond-async)
// @Success 201 {object} model.ApiResponse[migration.MigratedObjectStorageResult] "Created - Object storages provisioned (status: created or partial)"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Migration started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters or cloud configuration"
// @Failure 404 {object} model.ApiResponse[any] "Target namespace not found"
// @Failure 409 {object} model.ApiResponse[any] "Object storage with the same name already exists in the namespace"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during object storage creation"
// @Failure 503 {object} model.ApiResponse[any] "Too many requests - retry after the given time"
// @Header 503 {string} Retry-After "Seconds until client should retry"
// @Router /migration/middleware/ns/{nsId}/objectStorage [post]
func MigrateObjectStorage(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	var req MigrateObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("Failed to bind request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if req.TargetCloud.Csp == "" || req.TargetCloud.Region == "" {
		log.Warn().Msg("CSP and region are required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP and region required"))
	}

	if len(req.TargetObjectStorages) == 0 {
		log.Warn().Msg("At least one target object storage must be provided")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one target bucket required"))
	}

	nameSeed := c.QueryParam("nameSeed")
	if ok, detail := common.IsValidNameSeed(nameSeed); !ok {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid nameSeed: "+detail))
	}

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (migration.MigratedObjectStorageResult, error) {
			return migration.CreateObjectStorage(nsId, req.RecommendedObjectStorage, nameSeed)
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
			"Object storage migration started. Use GET /request/{reqId} to check status."))
	}

	result, err := migration.CreateObjectStorage(nsId, req.RecommendedObjectStorage, nameSeed)
	if err != nil {
		if retryAfter, ok := ratelimit.RetryAfter(err); ok {
			c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many requests to the underlying infrastructure provider; retry after the given time"))
		}
		log.Error().Err(err).Msg("Object storage migration failed")
		return c.JSON(mapObjectStorageUpstreamStatus(err), model.SimpleErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, model.SuccessResponseWithMessage(result, result.Description))
}

// ============================================================================
// Object Storage Management APIs
// ============================================================================

// ListObjectStorages godoc
// @ID ListObjectStorages
// @Summary List object storages (buckets)
// @Description Get the list of all object storages (buckets) in the namespace
// @Tags [Migration] Managed Object Storage
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param option query string false "Option to filter list (e.g., 'id')" Enums(id)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[storagemodel.ObjectStorageListResponse] "Successfully retrieved object storage list (default)"
// @Success 200 {object} model.ApiResponse[storagemodel.IdList] "Successfully retrieved object storage ID list (option=id)"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during list operation"
// @Failure 503 {object} model.ApiResponse[any] "Too many requests - retry after the given time"
// @Header 503 {string} Retry-After "Seconds until client should retry"
// @Router /migration/middleware/ns/{nsId}/objectStorage [get]
func ListObjectStorages(c echo.Context) error {
	nsId := c.Param("nsId")
	if nsId == "" {
		log.Warn().Msg("nsId is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("nsId required"))
	}

	option := c.QueryParam("option")
	if option != "" && option != "id" {
		err := fmt.Errorf("invalid request, the option (option: %s) is invalid", option)
		log.Warn().Msg(err.Error())
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	if option == "id" {
		idList, err := migration.ListObjectStorageIDs(nsId)
		if err != nil {
			if retryAfter, ok := ratelimit.RetryAfter(err); ok {
				c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
				return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
					"Too many requests to the underlying infrastructure provider; retry after the given time"))
			}
			log.Error().Err(err).Str("nsId", nsId).Msg("Failed to list object storage IDs")
			return c.JSON(mapObjectStorageUpstreamStatus(err), model.SimpleErrorResponse(fmt.Sprintf("Failed to list object storage IDs: %v", err)))
		}
		return c.JSON(http.StatusOK, model.SuccessResponse(idList))
	}

	result, err := migration.ListObjectStorages(nsId)
	if err != nil {
		if retryAfter, ok := ratelimit.RetryAfter(err); ok {
			c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many requests to the underlying infrastructure provider; retry after the given time"))
		}
		log.Error().Err(err).Str("nsId", nsId).Msg("Failed to list object storages")
		return c.JSON(mapObjectStorageUpstreamStatus(err), model.SimpleErrorResponse(fmt.Sprintf("Failed to list object storages: %v", err)))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}

// GetObjectStorage godoc
// @ID GetObjectStorage
// @Summary Get object storage (bucket) details
// @Description Get details of a specific object storage (bucket)
// @Tags [Migration] Managed Object Storage
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[storagemodel.ObjectStorageInfo] "Successfully retrieved object storage details"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during get operation"
// @Failure 503 {object} model.ApiResponse[any] "Too many requests - retry after the given time"
// @Header 503 {string} Retry-After "Seconds until client should retry"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId} [get]
func GetObjectStorage(c echo.Context) error {
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

	result, err := migration.GetObjectStorage(nsId, osId)
	if err != nil {
		if retryAfter, ok := ratelimit.RetryAfter(err); ok {
			c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many requests to the underlying infrastructure provider; retry after the given time"))
		}
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to get object storage")
		status := mapObjectStorageUpstreamStatus(err)
		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object storage '%s' not found", osId)))
		}
		return c.JSON(status, model.SimpleErrorResponse(fmt.Sprintf("Failed to get object storage '%s': %v", osId, err)))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(result))
}

// ExistObjectStorage godoc
// @ID ExistObjectStorage
// @Summary Check object storage (bucket) existence
// @Description Check if a specific object storage (bucket) exists
// @Description
// @Description [Note]
// @Description - Returns 200 OK if the bucket exists, 404 Not Found if it does not exist
// @Tags [Migration] Managed Object Storage
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 "OK - Object storage exists"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during existence check"
// @Failure 503 {object} model.ApiResponse[any] "Too many requests - retry after the given time"
// @Header 503 {string} Retry-After "Seconds until client should retry"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId} [head]
func ExistObjectStorage(c echo.Context) error {
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

	exists, err := migration.ExistObjectStorage(nsId, osId)
	if err != nil {
		if retryAfter, ok := ratelimit.RetryAfter(err); ok {
			c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many requests to the underlying infrastructure provider; retry after the given time"))
		}
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to check object storage existence")
		status := mapObjectStorageUpstreamStatus(err)
		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object storage '%s' not found", osId)))
		}
		return c.JSON(status, model.SimpleErrorResponse(fmt.Sprintf("Failed to check object storage '%s' existence: %v", osId, err)))
	}

	if !exists {
		return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object storage '%s' not found", osId)))
	}

	return c.NoContent(http.StatusOK)
}

// DeleteObjectStorage godoc
// @ID DeleteObjectStorage
// @Summary Delete or empty object storage (bucket)
// @Description Delete a specific object storage (bucket) or empty all objects inside it.
// @Description
// @Description By default this API runs synchronously. Send header `Prefer: respond-async` to run it
// @Description asynchronously (RFC 7240). Check progress via `GET /request/{reqId}`
// @Description (status flow: Handling -> Success / Error). Only the `"respond-async"` token is recognized.
// @Description
// @Description Behavior is controlled by the `option` query parameter (mutually exclusive):
// @Description - (omitted): Standard delete (fails if the bucket is not empty; returns 204 if already deleted).
// @Description - `empty`: Empty all objects in the bucket while keeping the bucket itself (returns 404 if the bucket does not exist).
// @Description - `force`: Empty all objects and force-delete the bucket and its metadata (passed to Spider as `force=true`; returns 204 if already deleted).
// @Tags [Migration] Managed Object Storage
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param option query string false "Delete or empty option" Enums(empty, force)
// @Param Prefer header string false "Set to 'respond-async' to run this operation asynchronously (RFC 7240)" Enums(respond-async)
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 204 "Object storage deleted (or emptied when option=empty) successfully"
// @Success 202 {object} model.ApiResponse[model.AsyncJobResponse] "Operation started asynchronously - use GET /request/{reqId} to check status"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters or unsupported option"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found (when option=empty)"
// @Failure 409 {object} model.ApiResponse[any] "Bucket is not empty (use option=force to delete with contents)"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during deletion"
// @Failure 503 {object} model.ApiResponse[any] "Too many requests - retry after the given time"
// @Header 503 {string} Retry-After "Seconds until client should retry"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId} [delete]
func DeleteObjectStorage(c echo.Context) error {
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

	option := c.QueryParam("option")
	if !isValidObjectStorageDeleteOption(option) {
		errMsg := fmt.Sprintf("Invalid option '%s': allowed values are 'empty' or 'force'", option)
		log.Warn().Str("nsId", nsId).Str("osId", osId).Str("option", option).Msg(errMsg)
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(errMsg))
	}

	if preferRespondAsync(c) {
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		started := common.RunAsync(reqID, func() (map[string]any, error) {
			if err := migration.DeleteObjectStorage(nsId, osId, option); err != nil {
				return nil, err
			}
			if option == "empty" {
				return map[string]any{"message": fmt.Sprintf("Object storage '%s' emptied successfully (bucket preserved)", osId)}, nil
			}
			return map[string]any{"message": fmt.Sprintf("Object storage '%s' deleted successfully", osId)}, nil
		})
		if !started {
			c.Response().Header().Set("Retry-After", "5")
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many async jobs in progress; retry shortly, or retry without Prefer: respond-async"))
		}
		c.Response().Header().Set("Preference-Applied", "respond-async")
		asyncMsg := fmt.Sprintf("Object storage '%s' deletion started. Use GET /request/{reqId} to check status.", osId)
		if option == "empty" {
			asyncMsg = fmt.Sprintf("Object storage '%s' empty task started. Use GET /request/{reqId} to check status.", osId)
		}
		return c.JSON(http.StatusAccepted, model.SuccessResponseWithMessage(
			model.AsyncJobResponse{
				ReqID:     reqID,
				Status:    common.RequestStatusHandling,
				StatusURL: fmt.Sprintf("/beetle/request/%s", reqID),
			},
			asyncMsg))
	}

	if err := migration.DeleteObjectStorage(nsId, osId, option); err != nil {
		if retryAfter, ok := ratelimit.RetryAfter(err); ok {
			c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many requests to the underlying infrastructure provider; retry after the given time"))
		}
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Str("option", option).Msg("Failed to delete object storage")
		status := mapObjectStorageUpstreamStatus(err)
		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object storage '%s' not found", osId)))
		}
		return c.JSON(status, model.SimpleErrorResponse(err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}

// ListObjectStorageObjects godoc
// @ID ListObjectStorageObjects
// @Summary List objects in an object storage bucket
// @Description List objects stored in a specific object storage bucket for post-migration verification (proxies Tumblebug GET /ns/{nsId}/resources/objectStorage/{osId}/object).
// @Tags [Migration] Managed Object Storage
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[migration.StorageObjectListResponse] "Successfully retrieved object list"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object storage not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during list operation"
// @Failure 503 {object} model.ApiResponse[any] "Too many requests - retry after the given time"
// @Header 503 {string} Retry-After "Seconds until client should retry"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId}/object [get]
func ListObjectStorageObjects(c echo.Context) error {
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

	result, err := migration.ListObjectStorageObjects(nsId, osId)
	if err != nil {
		if retryAfter, ok := ratelimit.RetryAfter(err); ok {
			c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many requests to the underlying infrastructure provider; retry after the given time"))
		}
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Msg("Failed to list objects in object storage")
		status := mapObjectStorageUpstreamStatus(err)
		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object storage '%s' not found", osId)))
		}
		return c.JSON(status, model.SimpleErrorResponse(fmt.Sprintf("Failed to list objects in '%s': %v", osId, err)))
	}

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(result, fmt.Sprintf("Listed %d object(s) in '%s'", result.Count, osId)))
}

// GetStorageObject godoc
// @ID GetStorageObject
// @Summary Get metadata of an object in an object storage bucket
// @Description Retrieve metadata of a specific object via HTTP response headers (`ETag`, `Last-Modified`, `Content-Length`, `X-Storage-Class`)
// @Description by proxying Tumblebug `HEAD /ns/{nsId}/resources/objectStorage/{osId}/object/{objectKey}`.
// @Description Note: URL-encode `objectKey` when it contains slashes (e.g., `folder%2Ffile.txt`).
// @Tags [Migration] Managed Object Storage
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param objectKey path string true "Object key (URL-encode slashes if needed)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 "OK - Object metadata returned in response headers (ETag, Last-Modified, Content-Length, X-Storage-Class)"
// @Header 200 {string} ETag "Entity tag (content hash) of the object"
// @Header 200 {string} Last-Modified "Last modified timestamp of the object"
// @Header 200 {string} Content-Length "Object size in bytes (when provided by upstream)"
// @Header 200 {string} X-Storage-Class "Storage class of the object (when provided by upstream)"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Failure 503 {object} model.ApiResponse[any] "Too many requests - retry after the given time"
// @Header 503 {string} Retry-After "Seconds until client should retry"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId}/object/{objectKey} [head]
func GetStorageObject(c echo.Context) error {
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

	objectKey := c.Param("objectKey")
	if objectKey == "" {
		log.Warn().Msg("objectKey is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("objectKey required"))
	}

	result, err := migration.GetStorageObject(nsId, osId, objectKey)
	if err != nil {
		if retryAfter, ok := ratelimit.RetryAfter(err); ok {
			c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many requests to the underlying infrastructure provider; retry after the given time"))
		}
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Str("objectKey", objectKey).Msg("Failed to get object metadata")
		status := mapObjectStorageUpstreamStatus(err)
		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object '%s' not found in '%s'", objectKey, osId)))
		}
		return c.JSON(status, model.SimpleErrorResponse(fmt.Sprintf("Failed to get metadata for '%s': %v", objectKey, err)))
	}

	respHeader := c.Response().Header()
	if result.ETag != "" {
		respHeader.Set("ETag", result.ETag)
	}
	if result.LastModified != "" {
		respHeader.Set("Last-Modified", result.LastModified)
	}
	if result.Size >= 0 {
		respHeader.Set("Content-Length", strconv.FormatInt(result.Size, 10))
	}
	if result.StorageClass != "" {
		respHeader.Set("X-Storage-Class", result.StorageClass)
	}

	return c.NoContent(http.StatusOK)
}

// DeleteStorageObject godoc
// @ID DeleteStorageObject
// @Summary Delete an object from an object storage bucket
// @Description Delete a specific object from an object storage bucket
// @Description by proxying Tumblebug `DELETE /ns/{nsId}/resources/objectStorage/{osId}/object/{objectKey}`.
// @Description Note: URL-encode `objectKey` when it contains slashes (e.g., `folder%2Ffile.txt`).
// @Tags [Migration] Managed Object Storage
// @Accept json
// @Produce json
// @Param nsId path string true "Namespace ID" default(mig01)
// @Param osId path string true "Object Storage ID (bucket ID)"
// @Param objectKey path string true "Object key (URL-encode slashes if needed)"
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 204 "Object deleted"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Failure 404 {object} model.ApiResponse[any] "Object not found"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Failure 503 {object} model.ApiResponse[any] "Too many requests - retry after the given time"
// @Header 503 {string} Retry-After "Seconds until client should retry"
// @Router /migration/middleware/ns/{nsId}/objectStorage/{osId}/object/{objectKey} [delete]
func DeleteStorageObject(c echo.Context) error {
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

	objectKey := c.Param("objectKey")
	if objectKey == "" {
		log.Warn().Msg("objectKey is required")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("objectKey required"))
	}

	if err := migration.DeleteStorageObject(nsId, osId, objectKey); err != nil {
		if retryAfter, ok := ratelimit.RetryAfter(err); ok {
			c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
			return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
				"Too many requests to the underlying infrastructure provider; retry after the given time"))
		}
		log.Error().Err(err).Str("nsId", nsId).Str("osId", osId).Str("objectKey", objectKey).Msg("Failed to delete object")
		status := mapObjectStorageUpstreamStatus(err)
		if status == http.StatusNotFound {
			return c.JSON(http.StatusNotFound, model.SimpleErrorResponse(fmt.Sprintf("Object '%s' not found in '%s'", objectKey, osId)))
		}
		return c.JSON(status, model.SimpleErrorResponse(fmt.Sprintf("Failed to delete '%s': %v", objectKey, err)))
	}

	return c.NoContent(http.StatusNoContent)
}
