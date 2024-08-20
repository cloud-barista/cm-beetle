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

// Package main is the starting point of CM-Beetle
package main

import (
	"flag"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/fsnotify/fsnotify"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"

	restServer "github.com/cloud-barista/cm-beetle/pkg/api/rest"
)

// NoOpLogger is an implementation of resty.Logger that discards all logs.
type NoOpLogger struct{}

func (n *NoOpLogger) Errorf(format string, v ...interface{}) {}
func (n *NoOpLogger) Warnf(format string, v ...interface{})  {}
func (n *NoOpLogger) Debugf(format string, v ...interface{}) {}

func init() {

	common.SystemReady = false

	// Initialize the configuration from "config.yaml" file or environment variables
	config.Init()

	// Initialize the logger
	logger := logger.NewLogger(logger.Config{
		LogLevel:    viper.GetString("beetle.loglevel"),
		LogWriter:   viper.GetString("beetle.logwriter"),
		LogFilePath: viper.GetString("beetle.logfile.path"),
		MaxSize:     viper.GetInt("beetle.logfile.maxsize"),
		MaxBackups:  viper.GetInt("beetle.logfile.maxbackups"),
		MaxAge:      viper.GetInt("beetle.logfile.maxage"),
		Compress:    viper.GetBool("beetle.logfile.compress"),
	})

	// Set the global logger
	log.Logger = *logger

	// Check Tumblebug readiness
	tumblebugEp := viper.GetString("beetle.tumblebug.endpoint")
	url := tumblebugEp + "/tumblebug/readyz"
	isReady, err := checkReadiness(url)

	if err != nil || !isReady {
		log.Fatal().Err(err).Msg("Tumblebug is not ready. Exiting...")
	}

	log.Info().Msg("Tumblebug is ready. Initializing Beetle...")

}

func checkReadiness(url string) (bool, error) {
	// Create a new resty client
	client := resty.New()

	// Disable Resty default logging by setting a no-op logger
	client.SetLogger(&NoOpLogger{})

	// Set for retries
	retryMaxAttempts := 20
	retryWaitTime := 3 * time.Second
	retryMaxWaitTime := 80 * time.Second
	// Configure retries
	client.
		// Set retry count to non zero to enable retries
		SetRetryCount(retryMaxAttempts).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(retryWaitTime).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(retryMaxWaitTime).
		// SetRetryAfter sets callback to calculate wait time between retries.
		// Default (nil) implies exponential backoff with jitter
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			attempt := resp.Request.Attempt // Current attempt number
			maxAttempts := retryMaxAttempts // Maximum attempt number

			log.Info().Msgf("check readiness by %s. Attempt %d/%d.",
				resp.Request.URL, attempt, maxAttempts)

			// Always retry after the calculated wait time
			return retryWaitTime, nil
		})

	resp, err := client.R().Get(url)

	if err != nil || resp.IsError() {
		return false, err
	}

	return true, nil
}

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

func main() {

	log.Info().Msg("CM-Beetle server is starting...")

	// Set the default port number "8056" for the REST API server to listen on
	port := flag.String("port", "8056", "port number for the restapiserver to listen to")
	flag.Parse()

	// Validate port
	if portInt, err := strconv.Atoi(*port); err != nil || portInt < 1 || portInt > 65535 {
		log.Fatal().Msgf("%s is not a valid port number. Please retry with a valid port number (ex: -port=[1-65535]).", *port)
	}
	log.Debug().Msgf("port number: %s", *port)

	// common.SpiderRestUrl = common.NVL(os.Getenv("BEETLE_SPIDER_REST_URL"), "http://localhost:1024/spider")
	// common.DragonflyRestUrl = common.NVL(os.Getenv("BEETLE_DRAGONFLY_REST_URL"), "http://localhost:9090/dragonfly")
	common.TumblebugRestUrl = common.NVL(os.Getenv("BEETLE_TUMBLEBUG_REST_URL"), "http://localhost:1323/tumblebug")
	common.DBUrl = common.NVL(os.Getenv("BEETLE_SQLITE_URL"), "localhost:3306")
	common.DBDatabase = common.NVL(os.Getenv("BEETLE_SQLITE_DATABASE"), "cm_beetle")
	common.DBUser = common.NVL(os.Getenv("BEETLE_SQLITE_USER"), "cm_beetle")
	common.DBPassword = common.NVL(os.Getenv("BEETLE_SQLITE_PASSWORD"), "cm_beetle")
	common.AutocontrolDurationMs = common.NVL(os.Getenv("BEETLE_AUTOCONTROL_DURATION_MS"), "10000")

	// load the latest configuration from DB (if exist)
	// fmt.Println("")
	// fmt.Println("[Update system environment]")
	// common.UpdateGlobalVariable(common.StrDragonflyRestUrl)
	// common.UpdateGlobalVariable(common.StrSpiderRestUrl)
	// common.UpdateGlobalVariable(common.StrAutocontrolDurationMs)

	// load config
	//masterConfigInfos = confighandler.GetMasterConfigInfos()

	//Setup database (meta_db/dat/cmbeetle.s3db)
	log.Info().Msg("setting SQL Database")
	err := os.MkdirAll("./meta_db/dat/", os.ModePerm)
	if err != nil {
		log.Error().Err(err).Msg("error creating directory")
	}
	log.Debug().Msgf("database file path: %s", "./meta_db/dat/cmbeetle.s3db")

	// Watch config file changes
	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Debug().Str("file", e.Name).Msg("config file changed")
			err := viper.ReadInConfig()
			if err != nil { // Handle errors reading the config file
				log.Fatal().Err(err).Msg("fatal error in config file")
			}
			err = viper.Unmarshal(&common.RuntimeConf)
			if err != nil {
				log.Panic().Err(err).Msg("error unmarshaling runtime configuration")
			}
		})
	}()

	// Launch API servers (REST)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// Start REST Server
	go func() {
		restServer.RunServer(*port)
		wg.Done()
	}()

	wg.Wait()
}
