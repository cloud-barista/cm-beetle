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

	"github.com/labstack/echo/v4"
)

// TestStreamingResponse godoc
// @ID TestStreamingResponse
// @Summary Test streaming response (JSON Lines)
// @Description Returns multiple JSON objects as newline-delimited JSON (JSON Lines format)
// @Tags [Test] Utilities
// @Accept json
// @Produce json
// @Success 200 {object} map[string]any "Multiple JSON objects"
// @Router /test/streaming [get]
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
