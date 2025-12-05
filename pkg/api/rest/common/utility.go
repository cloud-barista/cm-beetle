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

// Package common is to handle REST API for common funcitonalities
package common

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type TbConnectionName struct {
	ConnectionName string `json:"connectionName"`
}

type Existence struct {
	Exists bool `json:"exists"`
}

func SendExistence(c echo.Context, httpCode int, existence bool) error {
	return c.JSON(httpCode, Existence{Exists: existence})
}

type Status struct {
	Message string `json:"message"`
}

func SendMessage(c echo.Context, httpCode int, msg string) error {
	return c.JSON(httpCode, Status{Message: msg})
}

func Send(c echo.Context, httpCode int, json interface{}) error {
	return c.JSON(httpCode, json)
}

func Validate(c echo.Context, params []string) error {
	var err error
	for _, name := range params {
		err = validate.Var(c.Param(name), "required")
		if err != nil {
			return err
		}
	}
	return nil
}

type SimpleMessage struct {
	common.SimpleMsg
}

// GetReadyz godoc
// @ID GetReadyz
// @Summary Check Beetle is ready
// @Description Check Beetle is ready
// @Tags [Admin] System management
// @Accept  json
// @Produce  json
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} SimpleMessage
// @Failure 503 {object} SimpleMessage
// @Router /readyz [get]
func GetReadyz(c echo.Context) error {

	ctx := c.Request().Context()                    // Get context
	log.Ctx(ctx).Info().Msg("RestGetReadyz called") // Log ctx to trace

	message := SimpleMessage{}
	message.Message = "CM-Beetle is ready"
	if !common.SystemReady {
		message.Message = "CM-Beetle is NOT ready"
		return c.JSON(http.StatusServiceUnavailable, &message)
	}
	return c.JSON(http.StatusOK, &message)
}
