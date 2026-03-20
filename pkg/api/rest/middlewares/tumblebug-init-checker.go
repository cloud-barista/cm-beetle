package middlewares

import (
	"fmt"
	"net/http"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var isTbInitalized bool = false

// TumblebugInitChecker
func TumblebugInitChecker(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		log.Debug().Msg("[Middleware] Check and mark Tumblebug initialization status ")

		if !isTbInitalized {
			// Initialize resty client with basic auth
			client := resty.New()
			apiUser := config.Tumblebug.API.Username
			apiPass := config.Tumblebug.API.Password
			client.SetBasicAuth(apiUser, apiPass)

			// set tumblebug rest url
			epTumblebug := config.Tumblebug.RestUrl

			method := "GET"
			url := fmt.Sprintf("%s/readyz", epTumblebug)

			tbReq := common.NoBody
			tbRes := tbmodel.ReadyzResponse{}

			err := common.ExecuteHttpRequest(
				client,
				method,
				url,
				nil,
				common.SetUseBody(tbReq),
				&tbReq,
				&tbRes,
				common.VeryShortDuration,
			)

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

			isTbInitalized = true
			log.Info().Msg("Tumblebug is ready and initialized.")
		}

		return next(c)
	}
}
