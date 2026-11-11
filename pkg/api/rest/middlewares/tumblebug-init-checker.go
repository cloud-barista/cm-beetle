package middlewares

import (
	"context"
	"net/http"
	"sync/atomic"
	"time"

	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var isTbInitialized atomic.Bool

// TumblebugInitChecker verifies that CB-Tumblebug is reachable, ready, and initialized before serving API requests.
func TumblebugInitChecker(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Debug().Msg("[Middleware] Check and mark Tumblebug initialization status")

		if !isTbInitialized.Load() {
			readyCtx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
			tbRes, err := tbclient.NewSession().SetContext(readyCtx).CheckReadyz()
			cancel()

			if err != nil {
				log.Warn().Err(err).Msg("Tumblebug is not responding or returned an error.")
				res := common.SimpleMsg{Message: "Tumblebug is not responding or returned an error. Please check Tumblebug's status."}
				return c.JSON(http.StatusServiceUnavailable, res)
			}

			if !tbRes.Ready {
				log.Warn().Msg("Tumblebug is not ready.")
				res := common.SimpleMsg{Message: "Tumblebug is not ready. Please wait until Tumblebug is ready."}
				return c.JSON(http.StatusServiceUnavailable, res)
			}

			if !tbRes.Initialized {
				log.Warn().Msg("Tumblebug needs to be initialized.")
				res := common.SimpleMsg{Message: "Tumblebug needs to be initialized. See https://github.com/cloud-barista/cb-tumblebug?tab=readme-ov-file#3-initialize-cb-tumblebug-to-configure-multi-cloud-info"}
				return c.JSON(http.StatusServiceUnavailable, res)
			}

			isTbInitialized.Store(true)
			log.Info().Msg("Tumblebug is ready and initialized.")
		}

		return next(c)
	}
}
