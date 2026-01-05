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

// Package controller contains the handlers for REST API
package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// TestAuth godoc
// @ID TestAuth
// @Summary Test authentication
// @Description Checks if the request is authenticated and returns the auth type
// @Tags [Test] Utilities
// @Accept json
// @Produce json
// @Security BasicAuth
// @Security BearerAuth
// @Success 200 {object} model.ApiResponse[map[string]string] "Successfully authenticated"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters"
// @Router /test/auth [get]
// @x-order 2
func TestAuth(c echo.Context) error {
	log.Trace().Msg("TestAuth called")

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		log.Debug().Msg("Missing Authorization header")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Authorization header required"))
	}

	authType := "None"
	if strings.HasPrefix(authHeader, "Basic ") {
		authType = "Basic"
	} else if strings.HasPrefix(authHeader, "Bearer ") {
		authType = "Bearer"
	} else {
		log.Debug().Msg("Invalid Authorization header format")
		return c.JSON(http.StatusBadRequest, model.ErrorResponse("Invalid authorization format", "Must use 'Basic' or 'Bearer' authentication"))
	}

	log.Info().Msgf("Authentication check successful (Type: %s)", authType)
	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(map[string]string{"authType": authType}, "Authenticated as "+authType))
}

// TestTracing godoc
// @ID TestTracing
// @Description Tests distributed tracing by calling Tumblebug's readyz endpoint.
// @Description
// @Description [Note]
// @Description - The 'traceparent' header (W3C Trace Context) is propagated to Tumblebug for distributed tracing.
// @Description - The 'X-Request-Id' header is propagated for log correlation.
// @Description - Use this API to verify that tracing works correctly between Beetle and Tumblebug.
// @Tags [Test] Utilities
// @Accept  json
// @Produce  json
// @Param traceparent header string false "W3C Trace Context (e.g., 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01). Used for distributed tracing."
// @Param X-Request-Id header string false "Unique request ID. Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[string] "Tumblebug service is ready"
// @Failure 503 {object} model.ApiResponse[any] "Tumblebug service unavailable"
// @Router /test/tracing [get]
func TestTracing(c echo.Context) error {

	ctx := c.Request().Context()                    // Get context (contains OTel Trace info if middleware is active)
	log.Ctx(ctx).Info().Msg("RestGetReadyz called") // Log ctx to trace

	// [OpenTelemetry & Fallback]
	// Use SetTraceInfo(ctx) to propagate trace context.
	// It handles both OTel (future) and manual propagation (current fallback).
	session := tbclient.NewSession().SetTraceInfo(ctx)

	// Propagate X-Request-Id for log correlation
	if reqID := c.Response().Header().Get(echo.HeaderXRequestID); reqID != "" {
		session.SetHeader(echo.HeaderXRequestID, reqID)
	}

	_, ret, err := session.IsReady()
	if err != nil {
		log.Err(err).Msg("Failed to call Tumblebug readyz")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Tumblebug service unavailable"))
	}

	return c.JSON(http.StatusOK, model.SimpleSuccessResponse(*ret))
}

// TestStreamingResponse godoc
// @ID TestStreamingResponse
// @Summary Test streaming response (JSON Lines)
// @Description Returns multiple JSON objects as newline-delimited JSON (JSON Lines format)
// @Tags [Test] Utilities
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any "Multiple JSON objects"
// @Router /test/streaming [get]
// @x-order 1
func TestStreamingResponse(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	// Use json.Encoder which automatically appends newline after each Encode()
	enc := json.NewEncoder(c.Response())

	// Simulate multiple JSON responses (streaming/JSON Lines format)
	responses := []map[string]any{
		{"id": 1, "status": "processing", "message": "Starting migration..."},
		{"id": 2, "status": "processing", "message": "Copying data..."},
		{"id": 3, "status": "completed", "message": "Migration completed successfully"},
	}

	for _, resp := range responses {
		if err := enc.Encode(resp); err != nil {
			return err
		}
		c.Response().Flush()
	}

	return nil
}
