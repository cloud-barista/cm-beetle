package middlewares

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

// ResponseBodyDump returns a middleware that captures and processes response bodies.
//
// This middleware uses Echo's BodyDump, which wraps the response writer to capture
// the response body. The Handler function is called AFTER the response is complete,
// because the response body and HTTP status code are only available once the handler
// has finished writing the response.
//
// Flow:
//  1. Request arrives → request-id middleware sets status to "Handling"
//  2. Echo handler processes the request and writes response
//  3. BodyDump Handler is called (this middleware) → status updated to "Success" or "Error"
func ResponseBodyDump() echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			// Skip request tracking for API documentation paths
			if strings.HasPrefix(c.Path(), "/beetle/api") {
				return true
			}
			return false
		},
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			// log.Debug().Msg("Start - BodyDump() middleware")

			// Get the request ID
			reqID := c.Request().Header.Get(echo.HeaderXRequestID)
			// log.Debug().Msgf("(BodyDump middleware) Request ID: %s", reqID)

			// Get the content type
			contentType := c.Response().Header().Get(echo.HeaderContentType)
			//log.Trace().Msgf("contentType: %s", contentType)

			// log.Debug().Msgf("Request body: %s", string(reqBody))
			// log.Debug().Msgf("Response body: %s", string(resBody))

			// Load the request details by ID
			details, ok := common.GetRequest(reqID)
			if !ok {
				log.Error().Msg("Request ID not found")
				return
			}
			//log.Trace().Msg("OK, common.GetRequest(reqID)")
			details.EndTime = time.Now()

			details.Status = common.RequestStatusSuccess
			if c.Response().Status >= 400 {
				details.Status = common.RequestStatusError
			}

			// Set "X-Request-Id" in response header
			c.Response().Header().Set(echo.HeaderXRequestID, reqID)

			// Process response body based on content type
		processResponse:
			switch contentType {
			case echo.MIMEApplicationJSON:

				// Split the response body by newlines to handle JSON lines (streaming response)
				resBodies := bytes.Split(resBody, []byte("\n"))

				// Remove trailing empty element if present (caused by trailing newline)
				if len(resBodies) > 0 && len(resBodies[len(resBodies)-1]) == 0 {
					resBodies = resBodies[:len(resBodies)-1]
				}

				// Check if we have any valid JSON content
				if len(resBodies) == 0 || len(resBodies[0]) == 0 {
					log.Error().Msg("No valid response JSON found")
					break processResponse
				}

				// Extract error message from JSON response for failed requests
				if c.Response().Status >= 400 {
					// Unmarshal the last response (error info is in the last object)
					var lastResData map[string]any
					if err := json.Unmarshal(resBodies[len(resBodies)-1], &lastResData); err != nil {
						log.Error().Err(err).Msg("Error while unmarshaling error response")
						break processResponse
					}

					if message, exists := lastResData["message"]; exists && message != nil {
						if msgStr, ok := message.(string); ok {
							details.ErrorResponse = msgStr
						} else {
							details.ErrorResponse = "Error response message is not a string"
						}
					} else {
						details.ErrorResponse = "No error message found"
					}
					break processResponse
				}

				// Store the response data for successful requests
				var responseData []any
				for _, resBody := range resBodies {
					var obj any
					if err := json.Unmarshal(resBody, &obj); err != nil {
						log.Error().Err(err).Msg("Error unmarshalling JSON line")
						continue
					}
					responseData = append(responseData, obj)
				}

				// Store as single object if only one, otherwise as array
				if len(responseData) == 1 {
					details.ResponseData = responseData[0]
				} else {
					details.ResponseData = responseData
				}
			}

			// Store details of the request (always executed for all content types)
			if err := common.SetRequest(reqID, details); err != nil {
				log.Error().Err(err).Msg("Failed to store request details")
			}
			// log.Debug().Msg("End - BodyDump() middleware")
		},
	})
}
