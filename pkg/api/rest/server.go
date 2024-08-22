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
	"bytes"
	"context"
	"io"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	rest_common "github.com/cloud-barista/cm-beetle/pkg/api/rest/common"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/middlewares"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/spf13/viper"

	"crypto/subtle"
	"fmt"
	"os"

	"net/http"
	"net/url"

	// REST API (echo)
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// echo-swagger middleware
	_ "github.com/cloud-barista/cm-beetle/api"
	echoSwagger "github.com/swaggo/echo-swagger"

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

 ██████╗ ███████╗███████╗████████╗██╗     ███████╗
 ██╔══██╗██╔════╝██╔════╝╚══██╔══╝██║     ██╔════╝
 ██████╔╝█████╗  █████╗     ██║   ██║     █████╗  
 ██╔══██╗██╔══╝  ██╔══╝     ██║   ██║     ██╔══╝  
 ██████╔╝███████╗███████╗   ██║   ███████╗███████╗
 ╚═════╝ ╚══════╝╚══════╝   ╚═╝   ╚══════╝╚══════╝

 Computing Infrastructure Migration Technology
 ________________________________________________`
)

// RunServer func start Rest API server
func RunServer(port string) {

	log.Info().Msg("CM-Beetle REST API server is starting...")

	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger()) // default logger middleware in echo

	APILogSkipPatterns := [][]string{
		{"/beetle/api"},
		{"/tumblebug/api"},
		// {"/mci", "option=status"},
	}

	// Custom logger middleware with zerolog
	e.Use(middlewares.Zerologger(APILogSkipPatterns))

	e.Use(middleware.Recover())
	// limit the application to 20 requests/sec using the default in-memory store
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.HideBanner = true
	//e.colorer.Printf(banner, e.colorer.Red("v"+Version), e.colorer.Blue(website))

	allowedOrigins := viper.GetString("beetle.api.allow.origins")
	if allowedOrigins == "" {
		log.Fatal().Msg("allow_ORIGINS env variable for CORS is " + allowedOrigins +
			". Please provide a proper value and source setup.env again. EXITING...")
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{allowedOrigins},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Conditions to prevent abnormal operation due to typos (e.g., ture, falss, etc.)
	enableAuth := viper.GetString("beetle.api.auth.enabled") == "true"

	apiUser := viper.GetString("beetle.api.username")
	apiPass := viper.GetString("beetle.api.password")

	if enableAuth {
		e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
			// Skip authentication for some routes that do not require authentication
			Skipper: func(c echo.Context) bool {
				if c.Path() == "/beetle/readyz" ||
					c.Path() == "/beetle/httpVersion" ||
					strings.HasPrefix(c.Path(), "/tumblebug") {
					log.Debug().Msgf("Skip authentication for %s", c.Path())
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

	// The router group for Tumblebug wrapper, which has /tumblebug as prefix
	gTumblebug := e.Group("/tumblebug")

	// Set the target server for the proxy
	target := viper.GetString("beetle.tumblebug.endpoint")
	url, err := url.Parse(target)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// proxy middleware to forward the specified requests to the target server
	gTumblebug.Use(middlewares.Proxy(middlewares.ProxyConfig{
		URL: url,
		Rewrite: map[string]string{
			"/*": "/$1",
		},
		ModifyResponse: func(res *http.Response) error {
			resBytes, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}

			resBody := strings.TrimSuffix(string(resBytes), "\n")
			log.Debug().Msgf("[Proxy] response from %s", res.Request.URL)
			log.Trace().Msgf("[Proxy] response body: %s", resBody)

			res.Body = io.NopCloser(bytes.NewReader(resBytes))
			return nil
		},
	}))

	// Beetle API group which has /beetle as prefix
	gBeetle := e.Group("/beetle")

	// Swagger API docs
	swaggerRedirect := func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/beetle/api/index.html")
	}
	gBeetle.GET("/api", swaggerRedirect)
	gBeetle.GET("/api/", swaggerRedirect)

	gBeetle.GET("/api/*", echoSwagger.WrapHandler)

	// System management APIs
	gBeetle.GET("/readyz", rest_common.RestGetReadyz)
	gBeetle.GET("/httpVersion", rest_common.RestCheckHTTPVersion)

	// Namespace API group
	// gNamespace := gBeetle.Group("/ns")
	// gNamespace.POST("", controller.RestPostNs)
	// gNamespace.GET("", controller.RestGetAllNs)
	// gNamespace.GET("/:nsId", controller.RestGetNs)
	// gNamespace.DELETE("/:nsId", controller.RestDeleteNs)

	// Recommendation API group
	gRecommendation := gBeetle.Group("/recommendation")
	gRecommendation.POST("/mci", controller.RecommendInfra)

	// Migration API group
	gMigration := gBeetle.Group("/migration")
	gMigration.POST("/ns/:nsId/mci", controller.MigrateInfra)
	gMigration.GET("/ns/:nsId/mci/:mciId", controller.GetInfra)
	gMigration.DELETE("/ns/:nsId/mci/:mciId", controller.DeleteInfra)
	// gMigration.POST("/ns/:nsId/mci/network", controller.MigrateInfra)
	// gMigration.POST("/ns/:nsId/mci/storage", controller.MigrateInfra)
	// gMigration.POST("/ns/:nsId/mci/instance", controller.MigrateInfra)

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

	// Sample API group (for developers to add new API)
	gSample := gBeetle.Group("/sample")
	gSample.GET("/users", controller.GetUsers)
	gSample.GET("/users/:id", controller.GetUser)
	gSample.POST("/users", controller.CreateUser)
	gSample.PUT("/users/:id", controller.UpdateUser)
	gSample.DELETE("/users/:id", controller.DeleteUser)

	// Start API server
	selfEndpoint := viper.GetString("beetle.self.endpoint")
	apidashboard := " http://" + selfEndpoint + "/beetle/api"

	if enableAuth {
		fmt.Println(" Access to API dashboard" + " (username: " + apiUser + " / password: " + apiPass + ")")
	}
	fmt.Printf(noticeColor, apidashboard)
	fmt.Println("\n ")

	// A context for graceful shutdown (It is based on the signal package)selfEndpoint := os.Getenv("BEETLE_SELF_ENDPOINT")
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

		log.Info().Msg("Stopping CM-Beetle REST API server")
		ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("Error when graceful shutting down CM-Beetle API server")
			e.Logger.Panic(err)
		}
	}(&wg)

	port = fmt.Sprintf(":%s", port)
	common.SystemReady = true
	if err := e.Start(port); err != nil && err != http.ErrServerClosed {
		log.Error().Err(err).Msg("Error when starting CM-Beetle API server")
		e.Logger.Panic("Shuttig down the server: ", err)
	}

	log.Info().Msg("CM-Beetle REST API server is started.")

	wg.Wait()
}
