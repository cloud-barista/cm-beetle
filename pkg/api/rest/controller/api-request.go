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
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
	"github.com/go-resty/resty/v2"
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
// @Param X-Request-Id header string false "Custom request ID (auto-generated if not provided). Used as a trace ID for log correlation across Beetle and Tumblebug."
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /httpVersion [get]
func CheckHTTPVersion(c echo.Context) error {
	// Access the *http.Request object from the echo.Context
	req := c.Request()

	// Determine the HTTP protocol version of the request
	return c.JSON(http.StatusOK, common.SimpleMsg{Message: req.Proto})
}

// TestTracing godoc
// @ID TestTracing
// @Summary Test tracing to Tumblebug
// @Description Tests distributed tracing by calling Tumblebug's readyz endpoint with the X-Request-Id header.
// @Description
// @Description [Note]
// @Description - The X-Request-Id header value (auto-generated if not provided) is propagated to Tumblebug for distributed tracing.
// @Description - Use this API to verify that tracing works correctly between Beetle and Tumblebug.
// @Tags [Admin] API Request Management
// @Accept  json
// @Produce  json
// @Param X-Request-Id header string false "Custom request ID (auto-generated if not provided). Used as a trace ID for log correlation across Beetle and Tumblebug."
// @Success 200 {object} common.SimpleMsg
// @Failure 503 {object} common.SimpleMsg
// @Router /test/tracing [get]
func TestTracing(c echo.Context) error {

	ctx := c.Request().Context()                    // Get context
	log.Ctx(ctx).Info().Msg("RestGetReadyz called") // Log ctx to trace

	// Initialize resty client with basic auth
	client := resty.New()
	apiUser := config.Tumblebug.API.Username
	apiPass := config.Tumblebug.API.Password
	client.SetBasicAuth(apiUser, apiPass)

	// set tumblebug rest url
	epTumblebug := config.Tumblebug.RestUrl

	// Search and set a target VM spec
	method := "GET"
	url := fmt.Sprintf("%s/readyz", epTumblebug)

	// Headers
	headers := map[string]string{
		echo.HeaderXRequestID: ctx.Value(logger.TraceIdKey).(string),
	}

	// Request body
	tbReqt := common.NoBody

	// Response body
	tbResp := model.SimpleMsg{}

	err := common.ExecuteHttpRequest(
		client,
		method,
		url,
		headers,
		common.SetUseBody(tbReqt),
		&tbReqt,
		&tbResp,
		common.VeryShortDuration,
	)

	if err != nil {
		log.Err(err).Msg("")
		return c.JSON(http.StatusInternalServerError, common.SimpleMsg{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, common.SimpleMsg{Message: tbResp.Message})
}

// RestGetRequest godoc
// @ID GetRequest
// @Summary Get request details
// @Description Retrieves the details of a specific API request tracked by Beetle.
// @Description
// @Description [Note]
// @Description - Request tracking is managed independently by Beetle (not shared with Tumblebug).
// @Description - The reqId corresponds to the X-Request-Id header value from a previous API call.
// @Description - Do NOT call Tumblebug's /request/{reqId} API with this reqId; each system manages its own request tracking.
// @Description
// @Description [Status Values]
// @Description - received: Request has been received and registered (for async operations)
// @Description - in-progress: Request is currently being processed
// @Description - completed: Request completed successfully
// @Description - failed: Request failed with an error
// @Tags [Admin] API Request Management
// @Accept  json
// @Produce  json
// @Param reqId path string true "Request ID (from X-Request-Id header of a previous Beetle API call)"
// @Success 200 {object} common.RequestDetails
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /request/{reqId} [get]
func RestGetRequest(c echo.Context) error {
	reqId := c.Param("reqId")

	if details, ok := common.GetRequest(reqId); ok {
		return c.JSON(http.StatusOK, details)
	}

	return c.JSON(http.StatusNotFound, common.SimpleMsg{Message: "Request ID not found"})
}

// RestGetAllRequests godoc
// @ID GetAllRequests
// @Summary Get all requests
// @Description Retrieves all API requests tracked by Beetle with optional filters.
// @Description
// @Description [Note]
// @Description - Request tracking is managed independently by Beetle (not shared with Tumblebug).
// @Description - This API only returns requests made to Beetle, not to Tumblebug.
// @Description
// @Description [Status Values]
// @Description - received: Request has been received and registered (for async operations)
// @Description - in-progress: Request is currently being processed
// @Description - completed: Request completed successfully
// @Description - failed: Request failed with an error
// @Tags [Admin] API Request Management
// @Accept  json
// @Produce  json
// @Param status query string false "Filter by request status" Enums(received, in-progress, completed, failed) default()
// @Param method query string false "Filter by HTTP method (GET, POST, PUT, DELETE, etc.)" Enums(GET, POST, PUT, DELETE) default()
// @Param url query string false "Filter by request URL"
// @Param time query string false "Filter by time in minutes from now (to get recent requests)"
// @Param savefile query string false "Option to save the results to a file (set 'true' to activate)" Enums(true,false) default(false)
// @Success 200 {object} map[string][]common.RequestDetails
// @Router /requests [get]
func RestGetAllRequests(c echo.Context) error {
	// Filter parameters
	statusFilter := c.QueryParam("status")
	methodFilter := c.QueryParam("method")
	urlFilter := c.QueryParam("url")
	timeFilter := c.QueryParam("time") // in minutes

	// Build filter
	var filter *common.RequestFilter
	if statusFilter != "" || methodFilter != "" || urlFilter != "" || timeFilter != "" {
		filter = &common.RequestFilter{
			Status: statusFilter,
			Method: methodFilter,
			URL:    urlFilter,
		}
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
			return c.JSON(http.StatusInternalServerError, common.SimpleMsg{Message: "Failed to create log file"})
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
	return c.JSON(http.StatusOK, map[string][]common.RequestDetails{"requests": allRequests})
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
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Router /request/{reqId} [delete]
func RestDeleteRequest(c echo.Context) error {
	reqId := c.Param("reqId")

	if common.HasRequest(reqId) {
		common.RemoveRequest(reqId)
		return c.JSON(http.StatusOK, common.SimpleMsg{Message: "Request deleted successfully"})
	}

	return c.JSON(http.StatusNotFound, common.SimpleMsg{Message: "Request ID not found"})
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
// @Success 200 {object} common.SimpleMsg
// @Router /requests [delete]
func RestDeleteAllRequests(c echo.Context) error {
	common.RemoveAllRequests()
	return c.JSON(http.StatusOK, common.SimpleMsg{Message: "All requests deleted successfully"})
}
