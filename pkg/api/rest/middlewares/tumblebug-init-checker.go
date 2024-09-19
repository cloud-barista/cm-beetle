package middlewares

import (
	"fmt"
	"net/http"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	tbResource "github.com/cloud-barista/cb-tumblebug/src/api/rest/server/resource"
)

var isTbInitalized bool = false

// TumblebugInitChecker
func TumblebugInitChecker(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if !isTbInitalized {
			// Initialize resty client with basic auth
			client := resty.New()
			apiUser := config.Tumblebug.API.Username
			apiPass := config.Tumblebug.API.Password
			client.SetBasicAuth(apiUser, apiPass)

			// set tumblebug rest url
			epTumblebug := config.Tumblebug.RestUrl

			// Search and set a target VM spec
			method := "GET"
			nsId := "system"
			url := fmt.Sprintf("%s/ns/%s/resources/image", epTumblebug, nsId)

			tbReqt := common.NoBody
			tbResp := tbResource.RestGetAllImageResponse{}

			err := common.ExecuteHttpRequest(
				client,
				method,
				url,
				nil,
				common.SetUseBody(tbReqt),
				&tbReqt,
				&tbResp,
				common.VeryShortDuration,
			)

			if err != nil {
				log.Warn().Err(err).Msg("Tumblebug needs to be initialized.")
				res := common.SimpleMsg{Message: "Tumblebug needs to be initialized. See https://github.com/cloud-barista/cb-tumblebug?tab=readme-ov-file#3-initialize-cb-tumblebug-to-configure-multi-cloud-info"}
				return c.JSON(http.StatusServiceUnavailable, res)
			}

			isTbInitalized = true
		}

		return next(c)
	}
}
