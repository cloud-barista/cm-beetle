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

			// Dump the response body if content type is "application/json"
			if contentType == echo.MIMEApplicationJSON {
				// Load the request details by ID
				details, ok := common.GetRequest(reqID)
				if !ok {
					log.Error().Msg("Request ID not found")
					return
				}
				//log.Trace().Msg("OK, common.GetRequest(reqID)")
				details.EndTime = time.Now()

				// Set "X-Request-Id" in response header
				c.Response().Header().Set(echo.HeaderXRequestID, reqID)

				// Split the response body by newlines to handle multiple JSON objects (i.e., streaming response)
				parts := bytes.Split(resBody, []byte("\n"))
				if len(parts) == 0 {
					log.Error().Msg("Response body is empty")
					return
				}
				responseJsonLines := parts[:len(parts)-1]

				// Check if responseJsonLines has any content
				if len(responseJsonLines) == 0 {
					log.Error().Msg("No valid response JSON lines found")
					return
				}

				// Unmarshal the latest response body
				latestResponse := responseJsonLines[len(responseJsonLines)-1]
				var resData any
				if err := json.Unmarshal(latestResponse, &resData); err != nil {
					log.Error().Err(err).Msg("Error while unmarshaling response body")
					return
				}

				// Check and store error response
				// 1XX: Information responses
				// 2XX: Successful responses (200 OK, 201 Created, 202 Accepted, 204 No Content)
				// 3XX: Redirection messages
				// 4XX: Client error responses (400 Bad Request, 401 Unauthorized, 404 Not Found, 408 Request Timeout)
				// 5XX: Server error responses (500 Internal Server Error, 501 Not Implemented, 503 Service Unavailable)
				details.Status = common.RequestStatusCompleted
				if c.Response().Status >= 400 {
					details.Status = common.RequestStatusFailed
					if data, ok := resData.(map[string]any); ok {
						if message, exists := data["message"]; exists && message != nil {
							if msgStr, ok := message.(string); ok {
								details.ErrorResponse = msgStr
							} else {
								details.ErrorResponse = "Error response message is not a string"
							}
						} else {
							details.ErrorResponse = "No error message found"
						}
					}
				}

				// Store the response data
				if len(responseJsonLines) > 1 {
					// handle streaming response
					// convert JSON lines to JSON array
					var responseJsonArray []any
					for _, jsonLine := range responseJsonLines {
						var obj any
						err := json.Unmarshal(jsonLine, &obj)
						if err != nil {
							log.Error().Err(err).Msg("error unmarshalling JSON line")
							continue
						}
						responseJsonArray = append(responseJsonArray, obj)
					}
					details.ResponseData = responseJsonArray
				} else {
					// single response
					// type casting is required
					switch data := resData.(type) {
					case map[string]any:
						details.ResponseData = data
					case []any:
						details.ResponseData = data
					case string:
						details.ResponseData = data
					default:
						log.Error().Msgf("unexpected response data type (%T)", data)
					}
				}

				// Store details of the request
				if err := common.SetRequest(reqID, details); err != nil {
					log.Error().Err(err).Msg("Failed to store request details")
				}
			}
			// log.Debug().Msg("End - BodyDump() middleware")
		},
	})
}
