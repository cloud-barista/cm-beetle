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
	"net/http"
	"strings"

	model "github.com/cloud-barista/cm-beetle/pkg/api/rest/model/beetle"
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
// @Success 200 {object} model.ApiResponse[map[string]string] "Authentication successful"
// @Failure 400 {object} model.ApiResponse[any] "Invalid Request"
// @Router /test/auth [get]
// @x-order 2
func TestAuth(c echo.Context) error {
	log.Trace().Msg("TestAuth called")

	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		log.Debug().Msg("Missing Authorization header")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Missing Authorization header"))
	}

	authType := "None"
	if strings.HasPrefix(authHeader, "Basic ") {
		authType = "Basic"
	} else if strings.HasPrefix(authHeader, "Bearer ") {
		authType = "Bearer"
	} else {
		log.Debug().Msg("Invalid Authorization header format")
		return c.JSON(http.StatusBadRequest, model.ErrorResponse("Invalid Authorization header format", "Authorization header must start with 'Basic ' or 'Bearer '"))
	}

	log.Info().Msgf("Authentication check successful (Type: %s)", authType)
	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(map[string]string{"authType": authType}, "Authentication successful"))
}
