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
	"strconv"
	"sync"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/lkvstore"
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
		LogLevel:    config.Beetle.LogLevel,
		LogWriter:   config.Beetle.LogWriter,
		LogFilePath: config.Beetle.LogFile.Path,
		MaxSize:     config.Beetle.LogFile.MaxSize,
		MaxBackups:  config.Beetle.LogFile.MaxBackups,
		MaxAge:      config.Beetle.LogFile.MaxAge,
		Compress:    config.Beetle.LogFile.Compress,
	})

	// Set the global logger
	log.Logger = *logger

	// Initialize the local key-value store with the specified file path
	dbFilePath := config.Beetle.LKVStore.Path
	lkvstore.Init(lkvstore.Config{
		DbFilePath: dbFilePath,
	})

	// Check Tumblebug readiness
	apiUrl := config.Tumblebug.RestUrl + "/readyz"
	isReady, err := checkReadiness(apiUrl)

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
// // @termsOfService none

// @contact.name API Support
// @contact.url https://github.com/cloud-barista/cm-beetle/issues/new/choose

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8056
// @BasePath /beetle

// @securityDefinitions.basic BasicAuth

// @externalDocs.description ▶▶▶ CB-Tumblebug REST API (access via Beetle's reverse proxy)
// @externalDocs.url http://localhost:8056/tumblebug/api/index.html

func main() {

	log.Info().Msg("preparing to run CM-Beetle server...")

	// Load the state from the file back into the key-value store
	if err := lkvstore.LoadLkvStore(); err != nil {
		log.Warn().Err(err).Msgf("note - the lkvstore (file) may not exist at the initial startup.")
	}

	log.Info().Msg("successfully load data from the lkvstore (file).")

	defer func() {
		// Save the current state of the key-value store to file
		if err := lkvstore.SaveLkvStore(); err != nil {
			log.Error().Err(err).Msgf("error saving data to the lkvstore (file).")
		}
		log.Info().Msg("successfully save data to the lkvstore (file).")
	}()

	log.Info().Msg("Setting CM-Beetle REST API server...")

	// Set the default port number "8056" for the REST API server to listen on
	port := flag.String("port", "8056", "port number for the restapiserver to listen to")
	flag.Parse()

	// Validate port
	if portInt, err := strconv.Atoi(*port); err != nil || portInt < 1 || portInt > 65535 {
		log.Fatal().Msgf("%s is not a valid port number. Please retry with a valid port number (ex: -port=[1-65535]).", *port)
	}
	log.Debug().Msgf("port number: %s", *port)

	// Watch config file changes
	go func() {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Debug().Str("file", e.Name).Msg("config file changed")
			err := viper.ReadInConfig()
			if err != nil { // Handle errors reading the config file
				log.Fatal().Err(err).Msg("fatal error in config file")
			}
			err = viper.Unmarshal(&config.RuntimeConfig)
			if err != nil {
				log.Panic().Err(err).Msg("error unmarshaling runtime configuration")
			}
			config.Beetle = config.RuntimeConfig.Beetle
			config.Beetle.Tumblebug.RestUrl = config.Beetle.Tumblebug.Endpoint + "/tumblebug"
			config.Tumblebug = config.Beetle.Tumblebug
		})
	}()

	// Create the default namespace
	log.Debug().Msgf("creating the default namespace (%s)", common.DefaulNamespaceId)

	apiConfig := tbclient.ApiConfig{
		RestUrl:  config.Tumblebug.RestUrl,
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
	}

	tbCli := tbclient.NewClient(apiConfig)
	nsInfo, err := tbCli.ReadNamespace(common.DefaulNamespaceId)
	if err != nil {
		log.Debug().Msgf("not found the default namespace (nsId: %s)", common.DefaulNamespaceId)
		nsReq := tbmodel.NsReq{
			Name:        common.DefaulNamespaceId,
			Description: "Default namespace for computing infra migration",
		}
		nsInfo, err = tbCli.CreateNamespace(nsReq)
		if err != nil {
			log.Error().Err(err).Msg("failed to create a namespace")
		}
	}
	log.Info().Msgf("created the default namespace (nsId: %s)", nsInfo.Id)

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
