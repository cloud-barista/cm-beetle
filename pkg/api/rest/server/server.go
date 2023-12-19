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

// Package server is to handle REST API
package server

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
	"time"

	rest_common "github.com/cloud-barista/cm-beetle/pkg/api/rest/common"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/route"
	"github.com/spf13/viper"

	"crypto/subtle"
	"fmt"
	"os"

	"net/http"

	// REST API (echo)
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// echo-swagger middleware
	_ "github.com/cloud-barista/cm-beetle/pkg/api/rest/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	// Black import (_) is for running a package's init() function without using its other contents.
	_ "github.com/cloud-barista/cm-beetle/pkg/config"
	_ "github.com/cloud-barista/cm-beetle/pkg/logger"
	"github.com/rs/zerolog/log"
)

//var masterConfigInfos confighandler.MASTERCONFIGTYPE

const (
	infoColor    = "\033[1;34m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
	debugColor   = "\033[0;36m%s\033[0m"
)

const (
	website = " https://github.com/cloud-barista/cm-beetle"
	banner  = `    
                                         
 ██████╗ ███████╗ █████╗ ██████╗ ██╗   ██╗
 ██╔══██╗██╔════╝██╔══██╗██╔══██╗╚██╗ ██╔╝
 ██████╔╝█████╗  ███████║██║  ██║ ╚████╔╝ 
 ██╔══██╗██╔══╝  ██╔══██║██║  ██║  ╚██╔╝  
 ██║  ██║███████╗██║  ██║██████╔╝   ██║   
 ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝╚═════╝    ╚═╝   

 Cloud Migration
 ________________________________________________`
)

// RunServer func start Rest API server

// @title CM-Beetle REST API
// @version latest
// @description CM-Beetle REST API

// @contact.name API Support
// @contact.url http://cloud-barista.github.io
// @contact.email contact-to-cloud-barista@googlegroups.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /beetle

// @securityDefinitions.basic BasicAuth
func RunServer(port string) {

	log.Info().Msg("Setting CM-Beetle REST API server")

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// limit the application to 20 requests/sec using the default in-memory store
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.HideBanner = true
	//e.colorer.Printf(banner, e.colorer.Red("v"+Version), e.colorer.Blue(website))

	allowedOrigins := viper.GetString("api.allow.origins")
	if allowedOrigins == "" {
		log.Fatal().Msg("allow_ORIGINS env variable for CORS is " + allowedOrigins +
			". Please provide a proper value and source setup.env again. EXITING...")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{allowedOrigins},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Conditions to prevent abnormal operation due to typos (e.g., ture, falss, etc.)
	enableAuth := viper.GetString("api.auth.enabled") == "true"

	apiUser := viper.GetString("api.username")
	apiPass := viper.GetString("api.password")

	if enableAuth {
		e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
			// Skip authentication for some routes that do not require authentication
			Skipper: func(c echo.Context) bool {
				if c.Path() == "/beetle/health" ||
					c.Path() == "/beetle/httpVersion" {
					return true
				}
				return false
			},
			Validator: func(username, password string, c echo.Context) (bool, error) {
				// Be careful to use constant time comparison to prevent timing attacks
				if subtle.ConstantTimeCompare([]byte(username), []byte(apiUser)) == 1 &&
					subtle.ConstantTimeCompare([]byte(password), []byte(apiPass)) == 1 {
					return true, nil
				}
				return false, nil
			},
		}))
	}

	fmt.Println("\n \n ")
	fmt.Print(banner)
	fmt.Println("\n ")
	fmt.Println("\n ")
	fmt.Printf(infoColor, website)
	fmt.Println("\n \n ")

	// Route for system management
	e.GET("/beetle/swagger/*", echoSwagger.WrapHandler)
	// e.GET("/beetle/swaggerActive", rest_common.RestGetSwagger)
	e.GET("/beetle/health", rest_common.RestGetHealth)
	e.GET("/beetle/httpVersion", rest_common.RestCheckHTTPVersion)

	// Beetle API group which has /beetle as prefix
	groupBase := e.Group("/beetle")

	// Sample API group (for developers to add new API)
	groupSample := groupBase.Group("/sample")
	route.RegisterSampleRoutes(groupSample)

	// Recommendation API group
	groupRecommendation := groupBase.Group("/recommendation")
	route.RegisterRecommendationRoutes(groupRecommendation)

	// Migration API group
	groupMigration := groupBase.Group("/migration")
	route.RegisterMigrationRoutes(groupMigration)

	// Route
	// e.GET("/beetle/connConfig", rest_common.RestGetConnConfigList)
	// e.GET("/beetle/connConfig/:connConfigName", rest_common.RestGetConnConfig)

	// path specific timeout and ratelimit
	// timeout middleware
	// timeoutConfig := middleware.TimeoutConfig{
	// 	Timeout:      60 * time.Second,
	// 	Skipper:      middleware.DefaultSkipper,
	// 	ErrorMessage: "Error: request time out (60s)",
	// }

	// g.GET("/:nsId/mcis/:mcisId", rest_mcis.RestGetMcis, middleware.TimeoutWithConfig(timeoutConfig),
	// 	middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(2)))
	// g.GET("/:nsId/mcis", rest_mcis.RestGetAllMcis, middleware.TimeoutWithConfig(timeoutConfig),
	// 	middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(2)))

	// g.POST("/:nsId/mcis/:mcisId/vm", rest_mcis.RestPostMcisVm)
	// g.GET("/:nsId/mcis/:mcisId/vm/:vmId", rest_mcis.RestGetMcisVm)
	// g.GET("/:nsId/mcis/:mcisId/subgroup", rest_mcis.RestGetMcisGroupIds)
	// g.GET("/:nsId/mcis/:mcisId/subgroup/:subgroupId", rest_mcis.RestGetMcisGroupVms)
	// g.POST("/:nsId/mcis/:mcisId/subgroup/:subgroupId", rest_mcis.RestPostMcisSubGroupScaleOut)
	// g.DELETE("/:nsId/mcis", rest_mcis.RestDelAllMcis)

	selfEndpoint := viper.GetString("self.endpoint")
	apidashboard := " http://" + selfEndpoint + "/beetle/swagger/index.html"

	if enableAuth {
		fmt.Println(" Access to API dashboard" + " (username: " + apiUser + " / password: " + apiPass + ")")
	}
	fmt.Printf(noticeColor, apidashboard)
	fmt.Println("\n ")

	// A context for graceful shutdown (It is based on the signal package)selfEndpoint := os.Getenv("SELF_ENDPOINT")
	// NOTE -
	// Use os.Interrupt Ctrl+C or Ctrl+Break on Windows
	// Use syscall.KILL for Kill(can't be caught or ignored) (POSIX)
	// Use syscall.SIGTERM for Termination (ANSI)
	// Use syscall.SIGINT for Terminal interrupt (ANSI)
	// Use syscall.SIGQUIT for Terminal quit (POSIX)
	gracefulShutdownContext, stop := signal.NotifyContext(context.TODO(),
		os.Interrupt, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	// Wait graceful shutdown (and then main thread will be finished)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		// Block until a signal is triggered
		<-gracefulShutdownContext.Done()

		fmt.Println("\n[Stop] CM-Beetle REST Server")
		log.Info().Msg("stopping CM-Beetle REST Server")
		ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Panic(err)
		}
	}(&wg)

	log.Info().Msg("starting CM-Beetle REST API server")
	port = fmt.Sprintf(":%s", port)
	if err := e.Start(port); err != nil && err != http.ErrServerClosed {
		e.Logger.Panic("shuttig down the server")
	}

	wg.Wait()
}
