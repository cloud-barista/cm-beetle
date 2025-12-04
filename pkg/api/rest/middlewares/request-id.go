package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/labstack/echo/v4"
)

func RequestIdAndDetailsIssuer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// log.Debug().Msg("Start - Request ID middleware")

		// Skip request tracking for API documentation paths
		if strings.HasPrefix(c.Path(), "/beetle/api") {
			return next(c)
		}

		// Make X-Request-Id visible to all handlers
		c.Response().Header().Set("Access-Control-Expose-Headers", echo.HeaderXRequestID)

		// Get or generate Request ID
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		if reqID == "" {
			reqID = fmt.Sprintf("%d", time.Now().UnixNano())
			c.Request().Header.Set(echo.HeaderXRequestID, reqID)
		}

		//log.Trace().Msgf("(Request ID middleware) Request ID: %s", reqID)
		if common.HasRequest(reqID) {
			return echo.NewHTTPError(http.StatusConflict,
				fmt.Sprintf("the X-Request-Id '%s' is already in use; "+
					"use DELETE /beetle/request/{reqId} to remove it, or omit the header to auto-generate a new one", reqID))
		}

		// Set "X-Request-Id" in response header
		c.Response().Header().Set(echo.HeaderXRequestID, reqID)

		details := common.RequestDetails{
			StartTime:   time.Now(),
			Status:      common.RequestStatusInProgress,
			RequestInfo: common.ExtractRequestInfo(c.Request()),
		}
		if err := common.SetRequest(reqID, details); err != nil {
			return fmt.Errorf("failed to store request details: %w", err)
		}

		// log.Debug().Msg("End - Request ID middleware")

		return next(c)
	}
}

// NOTE - This is an example of how to use the RequestID middleware from echo
// func RequestIdIssuer() echo.MiddlewareFunc {
// 	return middleware.RequestIDWithConfig(middleware.RequestIDConfig{
// 		Generator: func() string {
// 			return fmt.Sprintf("%d", time.Now().UnixNano())
// 		},
// 		TargetHeader: echo.HeaderXRequestID,
// 	})
// }
