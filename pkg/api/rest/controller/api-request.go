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

package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// CheckHTTPVersion godoc
// @ID CheckHTTPVersion
// @Summary Check HTTP version of incoming request
// @Description Checks and returns the HTTP protocol version of the incoming request.
// @Description
// @Description [Note]
// @Description - The X-Request-Id header value (auto-generated if not provided) is propagated to Tumblebug when Beetle calls its APIs for distributed tracing.
// @Tags [Admin] API Request Management
// @Accept  json
// @Produce  json
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[string] "HTTP protocol version (e.g., HTTP/1.1, HTTP/2.0)"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error"
// @Router /httpVersion [get]
func CheckHTTPVersion(c echo.Context) error {
	// Access the *http.Request object from the echo.Context
	req := c.Request()

	// Determine the HTTP protocol version of the request
	return c.JSON(http.StatusOK, model.SimpleSuccessResponse(req.Proto))
}

// RestGetRequest godoc
// @ID GetRequest
// @Summary Get request details
// @Description Retrieves the details of a specific API request tracked by Beetle, including correlated outbound CB-Tumblebug sub-requests.
// @Description
// @Description [Note]
// @Description - Request tracking is managed independently by Beetle using `reqId` (from `X-Request-Id`).
// @Description - Outbound state-changing and recommendation calls to CB-Tumblebug are recorded in `subRequests[]` with `subReqId` (`{reqId}-{seq}`) and `tumblebugSource` (`cm-beetle:{reqId}`).
// @Description - Use `brief=true` during high-frequency polling to omit large request/response bodies while still receiving `status`, `subRequestSummary`, and `subRequests`.
// @Description - Set `withTumblebug=true` to enrich each `subRequests[]` entry with live `upstreamStatus` and `upstreamProgress` from CB-Tumblebug `GET /request/{subReqId}`.
// @Description
// @Description [Status Values]
// @Description - Handling: Request is currently being processed
// @Description - Success: Request completed successfully
// @Description - Error: Request failed with an error
// @Description
// @Description [Retry Information]
// @Description - retry.retryable: Whether the failed request can be retried (present only for retriable errors)
// @Description - retry.retryAfter: Suggested retry delay in seconds
// @Description - retry.retryReason: Human-readable reason for retry requirement (e.g., "Rate Limit Exceeded")
// @Tags [Admin] API Request Management
// @Accept  json
// @Produce  json
// @Param reqId path string true "Request ID (from X-Request-Id header of a previous Beetle API call)"
// @Param brief query bool false "Omit requestInfo.body and responseData for lightweight polling (default: false)" default(false)
// @Param withSubRequests query bool false "Include outbound CB-Tumblebug subRequests list (default: true)" default(true)
// @Param withTumblebug query bool false "Enrich subRequests with live upstream status/progress from CB-Tumblebug (default: false)" default(false)
// @Success 200 {object} model.ApiResponse[common.RequestDetails]
// @Failure 404 {object} model.ApiResponse[any]
// @Failure 500 {object} model.ApiResponse[any]
// @Router /request/{reqId} [get]
func RestGetRequest(c echo.Context) error {
	reqId := c.Param("reqId")

	details, ok := common.GetRequest(reqId)
	if !ok {
		return c.JSON(http.StatusNotFound, model.SimpleErrorResponse("Request ID not found"))
	}

	brief := c.QueryParam("brief") == "true"
	withSubRequests := c.QueryParam("withSubRequests") != "false"
	withTumblebug := c.QueryParam("withTumblebug") == "true"

	if withTumblebug && withSubRequests && len(details.SubRequests) > 0 {
		enrichCtx, cancel := context.WithTimeout(c.Request().Context(), 3*time.Second)
		defer cancel()

		for i := range details.SubRequests {
			subID := details.SubRequests[i].SubReqId
			if subID == "" {
				continue
			}
			tbReq, err := tbclient.NewSession().SetContext(enrichCtx).ReadRequest(subID)
			if err != nil {
				continue
			}
			details.SubRequests[i].UpstreamStatus = tbReq.Status
			details.SubRequests[i].UpstreamError = tbReq.ErrorResponse
			if strings.EqualFold(tbReq.Status, common.RequestStatusHandling) && tbReq.ResponseData != nil {
				details.SubRequests[i].UpstreamProgress = tbReq.ResponseData
			}
		}
	}

	common.ApplyRequestProjection(&details, brief, withSubRequests)
	return c.JSON(http.StatusOK, model.SuccessResponse(details))
}

// RestGetAllRequests godoc
// @ID GetAllRequests
// @Summary Get all requests
// @Description Retrieves API requests tracked by Beetle sorted by `startTime` descending (newest first), with optional filters and payload projection.
// @Description
// @Description [Note]
// @Description - Request tracking is managed independently by Beetle.
// @Description - Each entry includes `tumblebugSource` (`cm-beetle:{reqId}`) and `subRequestSummary` for cross-service correlation with CB-Tumblebug (`GET /tumblebug/requests?source=cm-beetle:{reqId}`).
// @Description
// @Description [Status Values]
// @Description - Handling: Request is currently being processed
// @Description - Success: Request completed successfully
// @Description - Error: Request failed with an error
// @Description
// @Description [Retry Information]
// @Description - retry.retryable: Whether the failed request can be retried (present only for retriable errors)
// @Description - retry.retryAfter: Suggested retry delay in seconds
// @Description - retry.retryReason: Human-readable reason for retry requirement (e.g., "Rate Limit Exceeded")
// @Tags [Admin] API Request Management
// @Accept  json
// @Produce  json
// @Param status query string false "Filter by request status" Enums(Handling, Success, Error) default()
// @Param method query string false "Filter by HTTP method (GET, POST, PUT, DELETE, etc.)" Enums(GET, POST, PUT, DELETE) default()
// @Param url query string false "Filter by request URL substring"
// @Param source query string false "Filter by X-Request-Source header or tumblebugSource substring"
// @Param time query string false "Filter by time in minutes from now (to get recent requests)"
// @Param limit query int false "Maximum number of recent requests to return (0 = unlimited)" default(0)
// @Param brief query bool false "Omit requestInfo.body and responseData for lightweight listing (default: false)" default(false)
// @Param savefile query string false "Option to save the results to a file (set 'true' to activate)" Enums(true,false) default(false)
// @Success 200 {object} map[string][]common.RequestDetails
// @Failure 400 {object} model.ApiResponse[any]
// @Router /requests [get]
func RestGetAllRequests(c echo.Context) error {
	statusFilter := c.QueryParam("status")
	methodFilter := c.QueryParam("method")
	urlFilter := c.QueryParam("url")
	sourceFilter := c.QueryParam("source")
	timeFilter := c.QueryParam("time")
	brief := c.QueryParam("brief") == "true"
	limitFilter := c.QueryParam("limit")

	limit := 0
	if limitFilter != "" {
		parsedLimit, err := strconv.Atoi(limitFilter)
		if err != nil || parsedLimit < 0 {
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("limit must be a non-negative integer"))
		}
		limit = parsedLimit
	}

	filter := &common.RequestFilter{
		Status: statusFilter,
		Method: methodFilter,
		URL:    urlFilter,
		Source: sourceFilter,
		Limit:  limit,
		Brief:  brief,
	}
	if timeFilter != "" {
		if minutes, err := strconv.Atoi(timeFilter); err == nil {
			filter.Since = time.Now().Add(-time.Duration(minutes) * time.Minute)
		}
	}

	// Get all requests with filter
	allRequests := common.GetAllRequests(filter)

	// Option to save the filtered results to a file
	if c.QueryParam("savefile") == "true" {
		beetleRoot := config.Beetle.Root
		logPath := filepath.Join(beetleRoot, "log", "request_log_"+time.Now().Format("20060102_150405")+".json")
		file, err := os.Create(logPath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Failed to create log file"))
		}
		defer file.Close()

		// Write filtered results as formatted JSON array
		file.WriteString("[\n")
		for i, detail := range allRequests {
			// Use MarshalIndent for pretty-printed JSON
			jsonLine, err := json.MarshalIndent(detail, "  ", "  ")
			if err != nil {
				log.Error().Err(err).Msg("Failed to marshal request detail")
				continue
			}

			file.Write(jsonLine)

			// Add comma except for the last item
			if i < len(allRequests)-1 {
				file.WriteString(",")
			}
			file.WriteString("\n")
		}
		file.WriteString("]\n")

		log.Info().Msgf("Filtered request log saved to: %s", logPath)

		// Return only the file path when savefile is requested
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":   "Filtered requests saved successfully",
			"file_path": logPath,
			"count":     len(allRequests),
		})
	}

	// Return the filtered requests data when savefile is not requested
	return c.JSON(http.StatusOK, model.SuccessResponse(map[string][]common.RequestDetails{"requests": allRequests}))
}

// RestDeleteRequest godoc
// @ID DeleteRequest
// @Summary Delete a specific request's details
// @Description Deletes the tracking details of a specific API request from Beetle.
// @Description
// @Description [Note]
// @Description - This only removes the request tracking record from Beetle's memory.
// @Description - It does NOT affect any data in Tumblebug or cancel any ongoing operations.
// @Tags [Admin] API Request Management
// @Accept  json
// @Produce  json
// @Param reqId path string true "Request ID to delete (from X-Request-Id header of a previous Beetle API call)"
// @Success 200 {object} model.ApiResponse[any]
// @Failure 404 {object} model.ApiResponse[any]
// @Router /request/{reqId} [delete]
func RestDeleteRequest(c echo.Context) error {
	reqId := c.Param("reqId")

	if common.HasRequest(reqId) {
		common.RemoveRequest(reqId)
		return c.JSON(http.StatusOK, model.SimpleSuccessResponse("Request deleted successfully"))
	}

	return c.JSON(http.StatusNotFound, model.SimpleErrorResponse("Request ID not found"))
}

// RestDeleteAllRequests godoc
// @ID DeleteAllRequests
// @Summary Delete all requests' details
// @Description Deletes all API request tracking records from Beetle.
// @Description
// @Description [Note]
// @Description - This only clears Beetle's request tracking memory.
// @Description - It does NOT affect any data in Tumblebug or cancel any ongoing operations.
// @Tags [Admin] API Request Management
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ApiResponse[any]
// @Router /requests [delete]
func RestDeleteAllRequests(c echo.Context) error {
	common.RemoveAllRequests()
	return c.JSON(http.StatusOK, model.SimpleSuccessResponse("All requests deleted successfully"))
}
