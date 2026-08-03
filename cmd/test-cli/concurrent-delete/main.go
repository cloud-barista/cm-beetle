// Package main tests concurrent infrastructure deletion with rate limiting
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

var (
	configFile = flag.String("config", "testconf/test-config.yaml", "Path to test configuration file")
)

// NoOpLogger suppresses resty client logs (including Basic Auth warnings)
type NoOpLogger struct{}

func (n *NoOpLogger) Errorf(format string, v ...interface{}) {}
func (n *NoOpLogger) Warnf(format string, v ...interface{})  {}
func (n *NoOpLogger) Debugf(format string, v ...interface{}) {}

// truncateMessage truncates long messages to avoid overwhelming console output
func truncateMessage(msg string, maxLen int) string {
	if len(msg) <= maxLen {
		return msg
	}
	return msg[:maxLen] + "... (truncated)"
}

// TestConfig holds test configuration
type TestConfig struct {
	Test struct {
		Cases []TestCase `yaml:"cases"`
	} `yaml:"test"`
	Beetle struct {
		Endpoint        string `yaml:"endpoint"`
		NamespaceID     string `yaml:"namespaceId"`
		RequestBodyFile string `yaml:"requestBodyFile"`
		AuthConfigFile  string `yaml:"authConfigFile"`
	} `yaml:"beetle"`
}

type TestCase struct {
	cloudmodel.CloudProperty `yaml:",inline"`
	Name                     string `yaml:"name"`
	Execute                  bool   `yaml:"execute"`
}

// InfraInfo holds created infrastructure information
type InfraInfo struct {
	CSP         string
	Region      string
	DisplayName string
	InfraID     string
	NameSeed    string
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	BeetleApiUsername    string `json:"beetleApiUsername"`
	BeetleApiPassword    string `json:"beetleApiPassword"`
	TumblebugApiUsername string `json:"tumblebugApiUsername"`
	TumblebugApiPassword string `json:"tumblebugApiPassword"`
	TumblebugEndpoint    string `json:"tumblebugEndpoint"`
}

func main() {
	flag.Parse()

	// Setup logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"})

	log.Info().Msg("========================================")
	log.Info().Msg("Concurrent Delete Test (5 Infra)")
	log.Info().Msg("========================================")

	// Load configuration
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	// Load authentication config
	authConfig, err := loadAuthConfig(config.Beetle.AuthConfigFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load auth config")
	}

	// Validate prerequisites
	log.Info().Msg("Checking prerequisites...")
	if err := validatePrerequisites(config); err != nil {
		log.Fatal().Err(err).Msg("Prerequisites check failed")
	}
	log.Info().Msg("✅ Prerequisites validated")

	// Load onpremise infrastructure model
	onpremInfraModel, _, err := loadOnpremInfraModel(config.Beetle.RequestBodyFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load onpremise infra model")
	}

	// Filter enabled test cases (limit to 5)
	var enabledCases []TestCase
	for _, tc := range config.Test.Cases {
		if tc.Execute {
			enabledCases = append(enabledCases, tc)
			if len(enabledCases) >= 5 {
				break
			}
		}
	}

	if len(enabledCases) == 0 {
		log.Fatal().Msg("No test cases enabled in config")
	}

	log.Info().Msgf("Selected %d CSP(s) for testing", len(enabledCases))
	for i, tc := range enabledCases {
		log.Info().Msgf("  [%d] %s (%s, %s)", i+1, tc.Name, tc.Csp, tc.Region)
	}

	// Phase 1: Recommend and Migrate
	log.Info().Msg("")
	log.Info().Msg("Phase 1: Recommendation & Migration")
	log.Info().Msg("----------------------------------------")

	var infraList []InfraInfo
	var mu sync.Mutex

	for i, tc := range enabledCases {
		caseNameSeed := fmt.Sprintf("cd%02d", i+1)

		log.Info().Msgf("[%d/%d] Processing %s with NameSeed: %s",
			i+1, len(enabledCases), tc.Name, caseNameSeed)

		// Recommendation
		infraID, err := recommendAndMigrate(config, tc, onpremInfraModel, caseNameSeed, authConfig)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to create infrastructure for %s", tc.Name)
			continue
		}

		mu.Lock()
		infraList = append(infraList, InfraInfo{
			CSP:         tc.Csp,
			Region:      tc.Region,
			DisplayName: tc.Name,
			InfraID:     infraID,
			NameSeed:    caseNameSeed,
		})
		mu.Unlock()

		log.Info().Msgf("✅ Created: %s (infraId: %s)", tc.Name, infraID)
	}

	if len(infraList) == 0 {
		log.Fatal().Msg("No infrastructure created successfully")
	}

	if len(infraList) < 2 {
		log.Warn().Msgf("Only %d infrastructure created (expected 5). Concurrent deletion test may not be effective", len(infraList))
		log.Warn().Msg("Proceeding anyway...")
	}

	// Phase 2: Wait for stabilization
	waitTime := 10 * time.Second
	log.Info().Msg("")
	log.Info().Msgf("Phase 2: Waiting %v for infrastructure stabilization...", waitTime)
	time.Sleep(waitTime)

	// Phase 3: Concurrent Delete
	log.Info().Msg("")
	log.Info().Msg("Phase 3: Concurrent Deletion")
	log.Info().Msg("----------------------------------------")
	log.Info().Msgf("Deleting %d infrastructure(s) concurrently...", len(infraList))

	var wg sync.WaitGroup
	startTime := time.Now()

	for i, infra := range infraList {
		wg.Add(1)
		go func(idx int, info InfraInfo) {
			defer wg.Done()

			deleteStart := time.Now()
			err := deleteInfra(config, info.InfraID, authConfig)
			duration := time.Since(deleteStart)

			if err != nil {
				log.Error().
					Str("infraId", info.InfraID).
					Str("csp", info.CSP).
					Dur("duration", duration).
					Err(err).
					Msgf("❌ [%d] Failed to delete %s", idx+1, info.DisplayName)
			} else {
				log.Info().
					Str("infraId", info.InfraID).
					Str("csp", info.CSP).
					Dur("duration", duration).
					Msgf("✅ [%d] Deleted %s", idx+1, info.DisplayName)
			}
		}(i, infra)
	}

	wg.Wait()
	totalDuration := time.Since(startTime)

	// Results
	log.Info().Msg("")
	log.Info().Msg("========================================")
	log.Info().Msg("Test Completed")
	log.Info().Msg("========================================")
	log.Info().Msgf("Total infrastructure: %d", len(infraList))
	log.Info().Msgf("Concurrent deletion time: %v", totalDuration)
	log.Info().Msgf("Expected time (with rate limiting): ~%v",
		time.Duration(len(infraList)*600)*time.Millisecond)
	log.Info().Msg("")
	log.Info().Msg("Check logs for rate limiter messages:")
	log.Info().Msg("  - 'Entered delete rate limiter queue'")
	log.Info().Msg("  - 'Rate limiting ReadInfra call'")
	log.Info().Msg("  - 'Rate limiter speeding up/slowing down'")
}

func loadConfig(filename string) (TestConfig, error) {
	var config TestConfig
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	return config, err
}

func validatePrerequisites(config TestConfig) error {
	client := resty.New().SetTimeout(5 * time.Second)
	client.SetLogger(&NoOpLogger{}) // Suppress warnings

	// Check 1: Beetle server readiness
	readyzURL := fmt.Sprintf("%s/beetle/readyz", config.Beetle.Endpoint)
	resp, err := client.R().Get(readyzURL)
	if err != nil || resp.StatusCode() != 200 {
		return fmt.Errorf("Beetle server not ready at %s (ensure 'make run' is running)", config.Beetle.Endpoint)
	}

	// Check message in response body
	var respBody map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &respBody); err == nil {
		if message, ok := respBody["message"].(string); ok {
			if strings.Contains(message, "NOT ready") {
				return fmt.Errorf("Beetle server not ready: %s", message)
			}
		}
	}

	// Check 2: Namespace exists
	nsURL := fmt.Sprintf("%s/tumblebug/ns/%s", config.Beetle.Endpoint, config.Beetle.NamespaceID)
	resp, err = client.R().Get(nsURL)
	if err != nil || resp.StatusCode() == 404 {
		return fmt.Errorf("namespace '%s' does not exist. Create it first:\n  curl -X POST %s/tumblebug/ns -d '{\"name\":\"%s\"}'",
			config.Beetle.NamespaceID, config.Beetle.Endpoint, config.Beetle.NamespaceID)
	}

	// Check 3: Request body file exists
	if _, err := os.Stat(config.Beetle.RequestBodyFile); os.IsNotExist(err) {
		return fmt.Errorf("request body file not found: %s", config.Beetle.RequestBodyFile)
	}

	return nil
}

func loadAuthConfig(authConfigPath string) (AuthConfig, error) {
	var authConfig AuthConfig

	// Return empty auth config if file path is not specified
	if authConfigPath == "" {
		return authConfig, nil
	}

	file, err := os.Open(authConfigPath)
	if err != nil {
		// If auth config file doesn't exist, return empty config (no error)
		if os.IsNotExist(err) {
			return authConfig, nil
		}
		return authConfig, fmt.Errorf("failed to open auth config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&authConfig); err != nil {
		return authConfig, fmt.Errorf("failed to decode auth config: %w", err)
	}

	return authConfig, nil
}

func loadOnpremInfraModel(filename string) (onpremmodel.OnpremInfra, string, error) {
	var infraModel onpremmodel.OnpremInfra
	var nameSeed string

	file, err := os.Open(filename)
	if err != nil {
		return infraModel, nameSeed, fmt.Errorf("failed to open request body file: %w", err)
	}
	defer file.Close()

	var tempRequest struct {
		NameSeed            string                  `json:"nameSeed"`
		OnpremiseInfraModel onpremmodel.OnpremInfra `json:"onpremiseInfraModel"`
	}

	if err := json.NewDecoder(file).Decode(&tempRequest); err != nil {
		return infraModel, nameSeed, fmt.Errorf("failed to decode request: %w", err)
	}

	return tempRequest.OnpremiseInfraModel, tempRequest.NameSeed, nil
}

func recommendAndMigrate(config TestConfig, tc TestCase, onpremInfra onpremmodel.OnpremInfra, nameSeed string, authConfig AuthConfig) (string, error) {
	// Use longer timeout: infrastructure creation can take 10-15 minutes
	client := resty.New().SetTimeout(20 * time.Minute)

	// Suppress resty logs (including Basic Auth warnings)
	client.SetLogger(&NoOpLogger{})

	// Set authentication if provided
	if authConfig.BeetleApiUsername != "" && authConfig.BeetleApiPassword != "" {
		client.SetBasicAuth(authConfig.BeetleApiUsername, authConfig.BeetleApiPassword)
	}

	// Step 1: Recommendation
	recommendReq := controller.RecommendInfraRequest{
		DesiredCspAndRegionPair: cloudmodel.CloudProperty{
			Csp:    tc.Csp,
			Region: tc.Region,
		},
		OnpremiseInfraModel: onpremInfra,
	}

	// Use query parameters like infra test does
	url := fmt.Sprintf("%s/beetle/recommendation/infra?desiredCsp=%s&desiredRegion=%s&limit=1",
		config.Beetle.Endpoint, tc.Csp, tc.Region)

	log.Debug().
		Str("csp", tc.Csp).
		Str("region", tc.Region).
		Int("nodeCount", len(onpremInfra.Nodes)).
		Str("url", url).
		Msg("Sending recommendation request")

	var recommendResp model.ApiResponse[[]cloudmodel.RecommendedInfra]
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(recommendReq).
		SetResult(&recommendResp).
		Post(url)

	if err != nil {
		return "", fmt.Errorf("recommendation request failed: %v", err)
	}

	if resp.StatusCode() != 200 {
		// Read error message from response body
		errorMsg := recommendResp.Error
		if errorMsg == "" {
			errorMsg = truncateMessage(string(resp.Body()), 200)
		}
		log.Error().
			Str("csp", tc.Csp).
			Str("region", tc.Region).
			Int("status", resp.StatusCode()).
			Str("error", truncateMessage(errorMsg, 150)).
			Msg("Recommendation failed")
		return "", fmt.Errorf("recommendation failed (status %d)", resp.StatusCode())
	}

	if !recommendResp.Success || len(recommendResp.Data) == 0 {
		errorMsg := recommendResp.Error
		if errorMsg == "" {
			errorMsg = "no recommendations returned"
		}
		log.Error().
			Str("csp", tc.Csp).
			Str("region", tc.Region).
			Bool("success", recommendResp.Success).
			Int("dataCount", len(recommendResp.Data)).
			Str("error", errorMsg).
			Str("responseBody", truncateMessage(string(resp.Body()), 300)).
			Msg("Recommendation returned empty or failed")
		return "", fmt.Errorf("recommendation failed: %s", errorMsg)
	}

	log.Debug().
		Str("csp", tc.Csp).
		Str("region", tc.Region).
		Int("recommendationCount", len(recommendResp.Data)).
		Msg("Recommendation successful")

	// Step 2: Migration
	migrationReq := controller.MigrateInfraRequest{
		RecommendedInfra: recommendResp.Data[0],
	}

	url = fmt.Sprintf("%s/beetle/migration/ns/%s/infra?nameSeed=%s",
		config.Beetle.Endpoint, config.Beetle.NamespaceID, nameSeed)

	var migrationResp model.ApiResponse[controller.MigrateInfraResponse]
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(migrationReq).
		SetResult(&migrationResp).
		Post(url)

	if err != nil {
		return "", fmt.Errorf("migration request failed: %v", err)
	}

	// Accept both 200 OK and 201 Created as success
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		// Read error message from response body
		errorMsg := migrationResp.Error
		if errorMsg == "" {
			errorMsg = truncateMessage(string(resp.Body()), 200)
		}
		log.Error().
			Str("csp", tc.Csp).
			Str("region", tc.Region).
			Str("nameSeed", nameSeed).
			Int("status", resp.StatusCode()).
			Str("error", truncateMessage(errorMsg, 150)).
			Msg("Migration failed")
		return "", fmt.Errorf("migration failed (status %d)", resp.StatusCode())
	}

	if !migrationResp.Success {
		return "", fmt.Errorf("migration failed: %s", migrationResp.Error)
	}

	// Get actual infraId from migration response
	infraID := migrationResp.Data.Id
	if infraID == "" {
		// Fallback: predict infraId (Beetle applies nameSeed to infrastructure name)
		infraID = common.ComposeName(recommendResp.Data[0].TargetInfra.Name, nameSeed)
	}

	return infraID, nil
}

func deleteInfra(config TestConfig, infraID string, authConfig AuthConfig) error {
	// Use longer timeout: infrastructure deletion can take 10-15 minutes
	client := resty.New().SetTimeout(20 * time.Minute)

	// Suppress resty logs
	client.SetLogger(&NoOpLogger{})

	// Set authentication if provided
	if authConfig.BeetleApiUsername != "" && authConfig.BeetleApiPassword != "" {
		client.SetBasicAuth(authConfig.BeetleApiUsername, authConfig.BeetleApiPassword)
	}

	url := fmt.Sprintf("%s/beetle/migration/ns/%s/infra/%s?option=terminate",
		config.Beetle.Endpoint, config.Beetle.NamespaceID, infraID)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Delete(url)

	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 && resp.StatusCode() != 202 {
		errorMsg := truncateMessage(string(resp.Body()), 200)
		return fmt.Errorf("delete failed (status %d): %s", resp.StatusCode(), errorMsg)
	}

	return nil
}
