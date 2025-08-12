// Package main is the starting point of CM-Beetle Test CLI
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"

	// Import Beetle's existing packages
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/core/recommendation"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"
)

// CSPTestReport holds test results for a specific CSP
type CSPTestReport struct {
	CSP                    string
	Region                 string
	DisplayName            string
	TestDate               string
	TestTime               string
	TestDateTime           time.Time
	BeetleURL              string
	NamespaceID            string
	OnpremiseInfraModel    onpremmodel.OnpremInfra
	RecommendationRequest  controller.RecommendVmInfraRequest
	RecommendationResponse *controller.RecommendVmInfraResponse
	MigrationResponse      *controller.MigrateInfraResponse
	ListMCIResponse        *cloudmodel.MciInfoList
	ListMCIIDsResponse     *cloudmodel.IdList
	GetMCIResponse         *cloudmodel.VmInfraInfo
	DeleteMCIResponse      map[string]interface{} // Simple response
	TestResults            []TestResults
	Summary                TestResults
}

// TestConfig holds test configuration
type TestConfig struct {
	BeetleURL             string                         `json:"beetleUrl"`
	NamespaceID           string                         `json:"namespaceId"`
	DesiredCspRegionPairs []recommendation.CspRegionPair `json:"desiredCspRegionPairs"`
	RequestBodyFile       string                         `json:"requestBodyFile"`
	AuthConfigFile        string                         `json:"authConfigFile,omitempty"`
	// Legacy fields for backward compatibility
	DesiredCSP    string `json:"desiredCsp,omitempty"`
	DesiredRegion string `json:"desiredRegion,omitempty"`
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	BasicAuthUsername string `json:"basicAuthUsername"`
	BasicAuthPassword string `json:"basicAuthPassword"`
}

// TestResults holds test execution results
type TestResults struct {
	TestName   string                 `json:"testName"`
	StartTime  time.Time              `json:"startTime"`
	EndTime    time.Time              `json:"endTime"`
	Duration   time.Duration          `json:"duration"`
	Success    bool                   `json:"success"`
	StatusCode int                    `json:"statusCode"`
	Response   map[string]interface{} `json:"response"`
	Error      string                 `json:"error,omitempty"`
}

// TestSuite holds all test results
type TestSuite struct {
	Config         TestConfig             `json:"config"`
	Results        []TestResults          `json:"results"`
	CspResults     map[string]TestResults `json:"cspResults"` // Results per CSP-Region pair
	TotalTests     int                    `json:"totalTests"`
	TotalCspPairs  int                    `json:"totalCspPairs"`
	PassedTests    int                    `json:"passedTests"`
	FailedTests    int                    `json:"failedTests"`
	PassedCspPairs int                    `json:"passedCspPairs"`
	FailedCspPairs int                    `json:"failedCspPairs"`
	OverallTime    time.Duration          `json:"overallTime"`
}

var (
	configFile = flag.String("config", "testdata/config-multi-csp-and-region-pair.json", "Path to config file")
	verbose    = flag.Bool("verbose", false, "Enable verbose output")
)

func init() {
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
}

func main() {
	flag.Parse()

	// Load test configuration
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	// Load auth configuration if specified
	authConfig, err := loadAuthConfig(config.AuthConfigFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load auth config")
	}

	// Handle legacy configuration format
	cspRegionPairs := config.DesiredCspRegionPairs
	if len(cspRegionPairs) == 0 && config.DesiredCSP != "" && config.DesiredRegion != "" {
		// Use legacy fields if new format is not provided
		cspRegionPairs = []recommendation.CspRegionPair{
			{
				Csp:    config.DesiredCSP,
				Region: config.DesiredRegion,
			},
		}
	}

	if len(cspRegionPairs) == 0 {
		log.Fatal().Msg("No CSP-Region pairs configured")
	}

	if *verbose {
		fmt.Printf("Starting Beetle API test with %d CSP-Region pairs\n", len(cspRegionPairs))
		for i, pair := range cspRegionPairs {
			displayName := fmt.Sprintf("%s-%s", pair.Csp, pair.Region)
			fmt.Printf("  %d. %s (%s, %s)\n", i+1, displayName, pair.Csp, pair.Region)
		}
	}

	// Initialize test suite
	suite := &TestSuite{
		Config:        config,
		Results:       make([]TestResults, 0),
		CspResults:    make(map[string]TestResults),
		TotalTests:    6, // Total number of API tests per CSP-Region pair
		TotalCspPairs: len(cspRegionPairs),
	}

	startTime := time.Now()

	// Load request body from JSON file
	onpremInfraModel, err := loadOnpremInfraModel(config.RequestBodyFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load onprem infra model")
	}

	// Initialize HTTP client
	client := resty.New()
	client.SetTimeout(20 * time.Minute) // Increased timeout to 20 minutes for all operations

	// Check CM-Beetle readiness first
	if err := checkBeetleReadiness(client, config.BeetleURL); err != nil {
		log.Fatal().Err(err).Msg("CM-Beetle readiness check failed")
	}

	// Set Basic Auth if configured
	if authConfig.BasicAuthUsername != "" && authConfig.BasicAuthPassword != "" {
		log.Info().Msg("🔐 Setting up Basic Authentication...")
		client.SetBasicAuth(authConfig.BasicAuthUsername, authConfig.BasicAuthPassword)
		log.Info().Msg("✅ Basic Auth configured")
	}

	// Test each CSP-Region pair sequentially
	for pairIndex, cspPair := range cspRegionPairs {
		displayName := fmt.Sprintf("%s-%s", cspPair.Csp, cspPair.Region)

		// Add delay between CSP-Region pairs (except for the first one)
		if pairIndex > 0 {
			animatedSleep(10*time.Second, fmt.Sprintf("Waiting for a while before testing %s", displayName))
		}

		log.Info().Msgf("\n" + strings.Repeat("=", 60) + "\n")
		log.Info().Msgf("Testing CSP-Region Pair %d/%d: %s", pairIndex+1, len(cspRegionPairs), displayName)
		log.Info().Str("csp", cspPair.Csp).Str("region", cspPair.Region).Msg("Starting CSP-Region pair test")
		log.Info().Msgf(strings.Repeat("=", 60) + "\n")

		pairStartTime := time.Now()
		pairPassed := 0
		pairFailed := 0

		// Create RecommendVmInfraRequest for this CSP-Region pair
		recommendRequest := controller.RecommendVmInfraRequest{
			DesiredCspAndRegionPair: recommendation.CspRegionPair{
				Csp:    cspPair.Csp,
				Region: cspPair.Region,
			},
			OnpremiseInfraModel: onpremInfraModel,
		}

		// Initialize CSP test report
		cspReport := &CSPTestReport{
			CSP:                   cspPair.Csp,
			Region:                cspPair.Region,
			DisplayName:           displayName,
			TestDate:              pairStartTime.Format("January 2, 2006"),
			TestTime:              pairStartTime.Format("15:04:05 MST"),
			TestDateTime:          pairStartTime,
			BeetleURL:             config.BeetleURL,
			NamespaceID:           config.NamespaceID,
			OnpremiseInfraModel:   onpremInfraModel,
			RecommendationRequest: recommendRequest,
			TestResults:           make([]TestResults, 0),
		}

		var mciId string // Will be extracted from migration response

		// Test 1: POST /beetle/recommendation/mci
		recommendedVmInfra, result1 := runRecommendationTest(client, config, cspPair, recommendRequest, displayName)
		suite.Results = append(suite.Results, result1)
		cspReport.TestResults = append(cspReport.TestResults, result1)

		if !result1.Success {
			pairFailed++
			log.Error().Str("csp", cspPair.Csp).Str("region", cspPair.Region).Msg("Test 1 (Recommendation) failed. Skipping remaining tests for this CSP-Region pair.")

			// Skip to summary for this pair since Test 1 failed
			pairDuration := time.Since(pairStartTime)

			// Complete CSP report summary
			cspReport.Summary = TestResults{
				TestName:  fmt.Sprintf("CSP-Region Pair: %s", displayName),
				Success:   false,
				Error:     "Test 1 (Recommendation) failed - remaining tests skipped",
				StartTime: pairStartTime,
				EndTime:   time.Now(),
				Duration:  pairDuration,
			}

			// Store the result per CSP
			suite.CspResults[displayName] = TestResults{
				TestName:  fmt.Sprintf("CSP-Region Pair: %s", displayName),
				Success:   false,
				Error:     "Test 1 (Recommendation) failed",
				StartTime: pairStartTime,
				EndTime:   time.Now(),
				Duration:  pairDuration,
				Response: map[string]interface{}{
					"passedTests": pairPassed,
					"failedTests": pairFailed,
					"totalTests":  suite.TotalTests,
				},
			}

			// Update suite counters
			suite.PassedTests += pairPassed
			suite.FailedTests += pairFailed
			suite.FailedCspPairs++

			// Print pair summary
			log.Info().Msgf("\n--- Summary for %s ---", displayName)
			log.Info().Int("passed", pairPassed).Int("total", suite.TotalTests).Msgf("Tests Passed: %d/%d", pairPassed, suite.TotalTests)
			log.Info().Int("failed", pairFailed).Int("total", suite.TotalTests).Msgf("Tests Failed: %d/%d", pairFailed, suite.TotalTests)
			log.Info().Dur("duration", pairDuration).Msgf("Duration: %v", pairDuration)
			log.Warn().Msg("Status: ❌ RECOMMENDATION TEST FAILED - REMAINING TESTS SKIPPED")

			continue // Skip to next CSP-Region pair
		}

		// Test 1 succeeded, continue with processing
		if structuredResponse, err := convertMapToRecommendVmInfraResponse(result1.Response); err == nil {
			cspReport.RecommendationResponse = structuredResponse
		} else {
			log.Warn().Err(err).Msg("Failed to convert recommendation response")
		}

		pairPassed++

		// Test 2: POST /beetle/migration/ns/{nsId}/mci
		// Convert RecommendVmInfraResponse to MigrateInfraRequest
		migrationRequest := controller.MigrateInfraRequest(recommendedVmInfra)
		result2 := runMigrationTest(client, config, migrationRequest, displayName)
		suite.Results = append(suite.Results, result2)
		cspReport.TestResults = append(cspReport.TestResults, result2)

		if !result2.Success {
			pairFailed++
			log.Error().Str("csp", cspPair.Csp).Str("region", cspPair.Region).Msg("Test 2 (Migration) failed. Skipping remaining tests for this CSP-Region pair.")

			// Skip to summary for this pair since Test 2 failed
			pairDuration := time.Since(pairStartTime)

			// Complete CSP report summary
			cspReport.Summary = TestResults{
				TestName:  fmt.Sprintf("CSP-Region Pair: %s", displayName),
				Success:   false,
				Error:     "Test 2 (Migration) failed - remaining tests skipped",
				StartTime: pairStartTime,
				EndTime:   time.Now(),
				Duration:  pairDuration,
			}

			// Store the result per CSP
			suite.CspResults[displayName] = TestResults{
				TestName:  fmt.Sprintf("CSP-Region Pair: %s", displayName),
				Success:   false,
				Error:     "Test 2 (Migration) failed",
				StartTime: pairStartTime,
				EndTime:   time.Now(),
				Duration:  pairDuration,
				Response: map[string]interface{}{
					"passedTests": pairPassed,
					"failedTests": pairFailed,
					"totalTests":  suite.TotalTests,
				},
			}

			// Update suite counters
			suite.PassedTests += pairPassed
			suite.FailedTests += pairFailed
			suite.FailedCspPairs++

			// Print pair summary
			log.Info().Msgf("\n--- Summary for %s ---", displayName)
			log.Info().Int("passed", pairPassed).Int("total", suite.TotalTests).Msgf("Tests Passed: %d/%d", pairPassed, suite.TotalTests)
			log.Info().Int("failed", pairFailed).Int("total", suite.TotalTests).Msgf("Tests Failed: %d/%d", pairFailed, suite.TotalTests)
			log.Info().Dur("duration", pairDuration).Msgf("Duration: %v", pairDuration)
			log.Warn().Msg("Status: ❌ MIGRATION TEST FAILED - REMAINING TESTS SKIPPED")

			continue // Skip to next CSP-Region pair
		}

		// Test 2 succeeded, continue with processing
		if structuredResponse, err := convertMapToMigrateInfraResponse(result2.Response); err == nil {
			cspReport.MigrationResponse = structuredResponse
		} else {
			log.Warn().Err(err).Msg("Failed to convert migration response")
		}

		pairPassed++
		// Extract mciId from response
		if id, ok := result2.Response["id"].(string); ok && id != "" {
			mciId = id
		} else if name, ok := result2.Response["name"].(string); ok && name != "" {
			mciId = name // Use name as fallback
		}

		// Test 3: GET /beetle/migration/ns/{nsId}/mci
		result3 := runListMciTest(client, config, displayName)
		suite.Results = append(suite.Results, result3)
		cspReport.TestResults = append(cspReport.TestResults, result3)
		if result3.Success {
			if structuredResponse, err := convertMapToMciInfoList(result3.Response); err == nil {
				cspReport.ListMCIResponse = structuredResponse
			} else {
				log.Warn().Err(err).Msg("Failed to convert list MCI response")
			}
		}

		if !result3.Success {
			pairFailed++
		} else {
			pairPassed++
		}

		// Test 4: GET /beetle/migration/ns/{nsId}/mci?option=id
		result4 := runListMciIdsTest(client, config, displayName)
		suite.Results = append(suite.Results, result4)
		cspReport.TestResults = append(cspReport.TestResults, result4)
		if result4.Success {
			if structuredResponse, err := convertMapToIdList(result4.Response); err == nil {
				cspReport.ListMCIIDsResponse = structuredResponse
			} else {
				log.Warn().Err(err).Msg("Failed to convert list MCI IDs response")
			}
		}

		if !result4.Success {
			pairFailed++
		} else {
			pairPassed++
		}

		// Test 5: GET /beetle/migration/ns/{nsId}/mci/{mciId}
		result5, executed5 := runGetMciTest(client, config, mciId, displayName)
		suite.Results = append(suite.Results, result5)
		cspReport.TestResults = append(cspReport.TestResults, result5)
		if result5.Success {
			if structuredResponse, err := convertMapToVmInfraInfo(result5.Response); err == nil {
				cspReport.GetMCIResponse = structuredResponse
			} else {
				log.Warn().Err(err).Msg("Failed to convert get MCI response")
			}
		}

		if !executed5 || !result5.Success {
			pairFailed++
		} else {
			pairPassed++
		}

		// Test 6: DELETE /beetle/migration/ns/{nsId}/mci/{mciId}
		result6, executed6 := runDeleteMciTest(client, config, mciId, displayName)
		suite.Results = append(suite.Results, result6)
		cspReport.TestResults = append(cspReport.TestResults, result6)
		if result6.Success {
			cspReport.DeleteMCIResponse = result6.Response
		}

		if !executed6 || !result6.Success {
			pairFailed++
		} else {
			pairPassed++
		}

		pairDuration := time.Since(pairStartTime)

		// Complete CSP report summary
		cspReport.Summary = TestResults{
			TestName:  fmt.Sprintf("CSP-Region Pair: %s", displayName),
			StartTime: pairStartTime,
			EndTime:   time.Now(),
			Duration:  pairDuration,
			Success:   pairFailed == 0,
			Response: map[string]interface{}{
				"csp":         cspPair.Csp,
				"region":      cspPair.Region,
				"passedTests": pairPassed,
				"failedTests": pairFailed,
				"totalTests":  suite.TotalTests,
			},
		}

		// Generate markdown report for this CSP
		if err := generateMarkdownReport(cspReport); err != nil {
			log.Warn().Str("csp", displayName).Err(err).Msg("Failed to generate markdown report")
		}

		// Store CSP-Region pair summary
		suite.CspResults[displayName] = TestResults{
			TestName:  fmt.Sprintf("CSP-Region Pair: %s", displayName),
			StartTime: pairStartTime,
			EndTime:   time.Now(),
			Duration:  pairDuration,
			Success:   pairFailed == 0,
			Response: map[string]interface{}{
				"csp":         cspPair.Csp,
				"region":      cspPair.Region,
				"passedTests": pairPassed,
				"failedTests": pairFailed,
				"totalTests":  suite.TotalTests,
			},
		}

		// Update suite counters
		suite.PassedTests += pairPassed
		suite.FailedTests += pairFailed
		if pairFailed == 0 {
			suite.PassedCspPairs++
		} else {
			suite.FailedCspPairs++
		}

		// Print pair summary
		log.Info().Msgf("\n--- Summary for %s ---", displayName)
		log.Info().Int("passed", pairPassed).Int("total", suite.TotalTests).Msgf("Tests Passed: %d/%d", pairPassed, suite.TotalTests)
		log.Info().Int("failed", pairFailed).Int("total", suite.TotalTests).Msgf("Tests Failed: %d/%d", pairFailed, suite.TotalTests)
		log.Info().Dur("duration", pairDuration).Msgf("Duration: %v", pairDuration)
		if pairFailed == 0 {
			log.Info().Msg("Status: ✅ ALL TESTS PASSED")
		} else {
			log.Warn().Msg("Status: ❌ SOME TESTS FAILED")
		}
	}

	suite.OverallTime = time.Since(startTime)

	// Print overall summary
	log.Info().Msgf("\n" + strings.Repeat("=", 60) + "\n")
	log.Info().Msg("OVERALL TEST SUMMARY")
	log.Info().Msgf(strings.Repeat("=", 60) + "\n")
	log.Info().Int("total", suite.TotalCspPairs).Msg("Total CSP-Region Pairs")
	log.Info().Int("successful", suite.PassedCspPairs).Msg("Successful Pairs")
	log.Info().Int("failed", suite.FailedCspPairs).Msg("Failed Pairs")
	log.Info().Int("totalTests", suite.TotalTests*suite.TotalCspPairs).Msg("Total Tests")
	log.Info().Int("passed", suite.PassedTests).Msg("Passed Tests")
	log.Info().Int("failed", suite.FailedTests).Msg("Failed Tests")
	log.Info().Dur("overallTime", suite.OverallTime).Msgf("Overall Time: %v", suite.OverallTime)

	log.Info().Msg("\nPer CSP-Region Results:")
	for name, result := range suite.CspResults {
		status := "✅"
		if !result.Success {
			status = "❌"
		}
		log.Info().
			Str("status", status).
			Str("csp", name).
			Int("passed", result.Response["passedTests"].(int)).
			Int("total", result.Response["totalTests"].(int)).
			Dur("duration", result.Duration).
			Msgf("%s %s - %d/%d tests passed (Duration: %v)",
				status, name, result.Response["passedTests"], result.Response["totalTests"], result.Duration)
	}

	// Note: JSON results saving has been disabled as requested
	// Results are available in markdown format in testresult directory

	// Exit with error code if any CSP pairs failed
	if suite.FailedCspPairs > 0 {
		os.Exit(1)
	}
}

// checkBeetleReadiness checks if CM-Beetle is ready using GET /beetle/readyz
func checkBeetleReadiness(client *resty.Client, beetleURL string) error {
	fmt.Println("\n🔍 Checking CM-Beetle readiness...")

	url := fmt.Sprintf("%s/beetle/readyz", beetleURL)

	var response map[string]interface{}
	var emptyBody interface{} = common.NoBody

	err := common.ExecuteHttpRequest(
		client,
		"GET",
		url,
		nil, // no custom headers for readiness check
		common.SetUseBody(emptyBody),
		&emptyBody,
		&response,
		0, // no cache
	)

	if err != nil {
		return fmt.Errorf("CM-Beetle readiness check failed: %v", err)
	}

	// Check if the response indicates ready status
	if message, ok := response["message"].(string); ok {
		if strings.Contains(message, "NOT ready") {
			return fmt.Errorf("CM-Beetle is not ready: %s", message)
		}
		log.Info().Str("message", message).Msg("CM-Beetle readiness check")
	} else {
		log.Info().Msg("✅ CM-Beetle is ready!")
	}

	return nil
}

// convertMapToRecommendVmInfraResponse converts map[string]interface{} to RecommendVmInfraResponse
func convertMapToRecommendVmInfraResponse(responseMap map[string]interface{}) (*controller.RecommendVmInfraResponse, error) {
	jsonBytes, err := json.Marshal(responseMap)
	if err != nil {
		return nil, err
	}

	var response controller.RecommendVmInfraResponse
	if err := json.Unmarshal(jsonBytes, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// convertMapToMigrateInfraResponse converts map[string]interface{} to MigrateInfraResponse
func convertMapToMigrateInfraResponse(responseMap map[string]interface{}) (*controller.MigrateInfraResponse, error) {
	jsonBytes, err := json.Marshal(responseMap)
	if err != nil {
		return nil, err
	}

	var response controller.MigrateInfraResponse
	if err := json.Unmarshal(jsonBytes, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// convertMapToMciInfoList converts map[string]interface{} to MciInfoList
func convertMapToMciInfoList(responseMap map[string]interface{}) (*cloudmodel.MciInfoList, error) {
	jsonBytes, err := json.Marshal(responseMap)
	if err != nil {
		return nil, err
	}

	var response cloudmodel.MciInfoList
	if err := json.Unmarshal(jsonBytes, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// convertMapToIdList converts map[string]interface{} to IdList
func convertMapToIdList(responseMap map[string]interface{}) (*cloudmodel.IdList, error) {
	jsonBytes, err := json.Marshal(responseMap)
	if err != nil {
		return nil, err
	}

	var response cloudmodel.IdList
	if err := json.Unmarshal(jsonBytes, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// convertMapToVmInfraInfo converts map[string]interface{} to VmInfraInfo
func convertMapToVmInfraInfo(responseMap map[string]interface{}) (*cloudmodel.VmInfraInfo, error) {
	jsonBytes, err := json.Marshal(responseMap)
	if err != nil {
		return nil, err
	}

	var response cloudmodel.VmInfraInfo
	if err := json.Unmarshal(jsonBytes, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

// runRecommendationTest performs Test 1: POST /beetle/recommendation/mci
func runRecommendationTest(client *resty.Client, config TestConfig, cspPair recommendation.CspRegionPair, recommendRequest controller.RecommendVmInfraRequest, displayName string) (controller.RecommendVmInfraResponse, TestResults) {
	log.Info().Msg("\n--- Test 1: POST /beetle/recommendation/mci ---")

	// Wait before API call for stability with animation
	animatedSleep(5*time.Second, "Waiting for a while for the previous task to be completed safely")

	result := TestResults{
		TestName:  fmt.Sprintf("POST /beetle/recommendation/mci (%s)", displayName),
		StartTime: time.Now(),
	}

	// Log API call details
	url := fmt.Sprintf("%s/beetle/recommendation/mci?desiredCsp=%s&desiredRegion=%s", config.BeetleURL, cspPair.Csp, cspPair.Region)
	log.Debug().Msgf("API Request URL: %s", url)

	// Log request body
	log.Debug().Msgf("API Request Body: %+v", recommendRequest)

	var response controller.RecommendVmInfraResponse
	err := common.ExecuteHttpRequest(
		client,
		"POST",
		url,
		nil,  // no custom headers
		true, // use body
		&recommendRequest,
		&response,
		0, // no cache duration
	)

	// Log response body
	log.Debug().Msgf("API Response Body: %+v", response)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		result.StatusCode = 0
		fmt.Printf("❌ Test 1 failed: %s\n", result.Error)
	} else {
		result.Success = true
		result.StatusCode = 200
		// Convert struct response to map for TestResults compatibility
		responseMap := make(map[string]interface{})
		jsonBytes, _ := json.Marshal(response)
		json.Unmarshal(jsonBytes, &responseMap)
		result.Response = responseMap
		fmt.Println("✅ Test 1 passed")
	}

	return response, result
}

// runMigrationTest performs Test 2: POST /beetle/migration/ns/{nsId}/mci
func runMigrationTest(client *resty.Client, config TestConfig, migrationRequestBody controller.MigrateInfraRequest, displayName string) TestResults {
	fmt.Printf("\n--- Test 2: POST /beetle/migration/ns/%s/mci ---\n", config.NamespaceID)

	// Wait before API call for stability (migration needs more time) with spinner
	animatedSleep(5*time.Second, "Waiting for a while for the previous task to be completed safely")

	result := TestResults{
		TestName:  fmt.Sprintf("POST /beetle/migration/ns/{nsId}/mci (%s)", displayName),
		StartTime: time.Now(),
	}

	// Log API call details
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/mci", config.BeetleURL, config.NamespaceID)
	log.Debug().Msgf("API Request URL: %s", url)

	// Log request body
	log.Debug().Msgf("API Request Body: %+v", migrationRequestBody)

	var response controller.MigrateInfraResponse
	err := common.ExecuteHttpRequest(
		client,
		"POST",
		url,
		nil,
		true,
		&migrationRequestBody,
		&response,
		0,
	)

	// Log response body
	log.Debug().Msgf("API Response Body: %+v", response)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		result.StatusCode = 0
		fmt.Printf("❌ Test 2 failed: %s\n", result.Error)
	} else {
		result.Success = true
		result.StatusCode = 200
		// Convert struct response to map for TestResults compatibility
		responseMap := make(map[string]interface{})
		jsonBytes, _ := json.Marshal(response)
		json.Unmarshal(jsonBytes, &responseMap)
		result.Response = responseMap
		fmt.Println("✅ Test 2 passed")
	}

	return result
}

// runListMciTest performs Test 3: GET /beetle/migration/ns/{nsId}/mci
func runListMciTest(client *resty.Client, config TestConfig, displayName string) TestResults {
	fmt.Printf("\n--- Test 3: GET /beetle/migration/ns/%s/mci ---\n", config.NamespaceID)

	// Wait before API call for stability with animation
	animatedSleep(5*time.Second, "Waiting for a while for the previous task to be completed safely")

	result := TestResults{
		TestName:  fmt.Sprintf("GET /beetle/migration/ns/{nsId}/mci (%s)", displayName),
		StartTime: time.Now(),
	}

	// Log API call details
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/mci", config.BeetleURL, config.NamespaceID)
	log.Debug().Msgf("API Request URL: %s", url)
	log.Debug().Msg("API Request Body: none")

	var response cloudmodel.MciInfoList
	var emptyBody interface{} = common.NoBody
	err := common.ExecuteHttpRequest(
		client,
		"GET",
		url,
		nil,
		common.SetUseBody(emptyBody),
		&emptyBody,
		&response,
		0,
	)

	// Log response body
	log.Debug().Msgf("API Response Body: %+v", response)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		result.StatusCode = 0
		fmt.Printf("❌ Test 3 failed: %s\n", result.Error)
	} else {
		result.Success = true
		result.StatusCode = 200
		// Convert struct response to map for TestResults compatibility
		responseMap := make(map[string]interface{})
		jsonBytes, _ := json.Marshal(response)
		json.Unmarshal(jsonBytes, &responseMap)
		result.Response = responseMap
		fmt.Println("✅ Test 3 passed")
	}

	return result
}

// runListMciIdsTest performs Test 4: GET /beetle/migration/ns/{nsId}/mci?option=id
func runListMciIdsTest(client *resty.Client, config TestConfig, displayName string) TestResults {
	fmt.Printf("\n--- Test 4: GET /beetle/migration/ns/%s/mci?option=id ---\n", config.NamespaceID)

	// Wait before API call for stability with animation
	animatedSleep(5*time.Second, "Waiting for a while for the previous task to be completed safely")

	result := TestResults{
		TestName:  fmt.Sprintf("GET /beetle/migration/ns/{nsId}/mci?option=id (%s)", displayName),
		StartTime: time.Now(),
	}

	// Log API call details
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/mci?option=id", config.BeetleURL, config.NamespaceID)
	log.Debug().Msgf("API Request URL: %s", url)
	log.Debug().Msg("API Request Body: none")

	var response cloudmodel.IdList
	var emptyBody interface{} = common.NoBody
	err := common.ExecuteHttpRequest(
		client,
		"GET",
		url,
		nil,
		common.SetUseBody(emptyBody),
		&emptyBody,
		&response,
		0,
	)

	// Log response body
	log.Debug().Msgf("API Response Body: %+v", response)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		result.StatusCode = 0
		fmt.Printf("❌ Test 4 failed: %s\n", result.Error)
	} else {
		result.Success = true
		result.StatusCode = 200
		// Convert struct response to map for TestResults compatibility
		responseMap := make(map[string]interface{})
		jsonBytes, _ := json.Marshal(response)
		json.Unmarshal(jsonBytes, &responseMap)
		result.Response = responseMap
		fmt.Println("✅ Test 4 passed")
	}

	return result
}

// runGetMciTest performs Test 5: GET /beetle/migration/ns/{nsId}/mci/{mciId}
func runGetMciTest(client *resty.Client, config TestConfig, mciId, displayName string) (TestResults, bool) {
	if mciId == "" {
		fmt.Println("⚠️  Test 5 skipped: No MCI ID available")
		return TestResults{
			TestName: fmt.Sprintf("GET /beetle/migration/ns/{nsId}/mci/{mciId} (%s)", displayName),
			Success:  false,
			Error:    "MCI ID not available from previous tests",
		}, false
	}

	fmt.Printf("\n--- Test 5: GET /beetle/migration/ns/%s/mci/%s ---\n", config.NamespaceID, mciId)

	// Wait before API call for stability with animation
	animatedSleep(5*time.Second, "Waiting for a while for the previous task to be completed safely")

	result := TestResults{
		TestName:  fmt.Sprintf("GET /beetle/migration/ns/{nsId}/mci/{mciId} (%s)", displayName),
		StartTime: time.Now(),
	}

	// Log API call details
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/mci/%s", config.BeetleURL, config.NamespaceID, mciId)
	log.Debug().Msgf("API Request URL: %s", url)
	log.Debug().Msg("API Request Body: none")

	var response cloudmodel.VmInfraInfo
	var emptyBody interface{} = common.NoBody
	err := common.ExecuteHttpRequest(
		client,
		"GET",
		url,
		nil,
		common.SetUseBody(emptyBody),
		&emptyBody,
		&response,
		0,
	)

	// Log response body
	log.Debug().Msgf("API Response Body: %+v", response)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		result.StatusCode = 0
		fmt.Printf("❌ Test 5 failed: %s\n", result.Error)
	} else {
		result.Success = true
		result.StatusCode = 200
		// Convert struct response to map for TestResults compatibility
		responseMap := make(map[string]interface{})
		jsonBytes, _ := json.Marshal(response)
		json.Unmarshal(jsonBytes, &responseMap)
		result.Response = responseMap
		fmt.Println("✅ Test 5 passed")
	}

	return result, true
}

// runDeleteMciTest performs Test 6: DELETE /beetle/migration/ns/{nsId}/mci/{mciId}
func runDeleteMciTest(client *resty.Client, config TestConfig, mciId, displayName string) (TestResults, bool) {
	if mciId == "" {
		fmt.Println("⚠️  Test 6 skipped: No MCI ID available")
		return TestResults{
			TestName: fmt.Sprintf("DELETE /beetle/migration/ns/{nsId}/mci/{mciId} (%s)", displayName),
			Success:  false,
			Error:    "MCI ID not available from previous tests",
		}, false
	}

	fmt.Printf("\n--- Test 6: DELETE /beetle/migration/ns/%s/mci/%s?option=terminate ---\n", config.NamespaceID, mciId)

	// Wait before API call for stability with animation
	animatedSleep(5*time.Second, "Waiting for a while for the previous task to be completed safely")

	result := TestResults{
		TestName:  fmt.Sprintf("DELETE /beetle/migration/ns/{nsId}/mci/{mciId} (%s)", displayName),
		StartTime: time.Now(),
	}

	// Log API call details
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/mci/%s?option=terminate", config.BeetleURL, config.NamespaceID, mciId)
	log.Debug().Msgf("API Request URL: %s", url)
	log.Debug().Msg("API Request Body: none")

	var response map[string]interface{}
	var emptyBody interface{} = common.NoBody
	err := common.ExecuteHttpRequest(
		client,
		"DELETE",
		url,
		nil,
		common.SetUseBody(emptyBody),
		&emptyBody,
		&response,
		0,
	)

	// Log response body
	if response != nil {
		log.Debug().Msgf("API Response Body: %+v", response)
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		result.StatusCode = 0
		fmt.Printf("❌ Test 6 failed: %s\n", result.Error)
	} else {
		result.Success = true
		result.StatusCode = 200
		result.Response = response
		fmt.Println("✅ Test 6 passed")
	}

	return result, true
}

// checkBeetleReadiness checks if CM-Beetle is ready using GET /beetle/readyz

func loadConfig(configPath string) (TestConfig, error) {
	var config TestConfig

	file, err := os.Open(configPath)
	if err != nil {
		return config, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, fmt.Errorf("failed to decode config: %w", err)
	}

	return config, nil
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

func loadOnpremInfraModel(requestBodyPath string) (onpremmodel.OnpremInfra, error) {
	var infraModel onpremmodel.OnpremInfra

	file, err := os.Open(requestBodyPath)
	if err != nil {
		return infraModel, fmt.Errorf("failed to open request body file: %w", err)
	}
	defer file.Close()

	// First load into a temporary structure to extract onpremiseInfraModel
	var tempRequest struct {
		OnpremiseInfraModel onpremmodel.OnpremInfra `json:"onpremiseInfraModel"`
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tempRequest); err != nil {
		return infraModel, fmt.Errorf("failed to decode onprem infra model: %w", err)
	}

	return tempRequest.OnpremiseInfraModel, nil
}

// getGitHash returns the current git commit hash
func getGitHash() string {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

// getVersionFromDockerCompose extracts version from docker-compose.yaml for a given service
func getVersionFromDockerCompose(serviceName string) string {
	dockerComposePath := "../../deployments/docker-compose/docker-compose.yaml"

	// Try absolute path first
	if _, err := os.Stat(dockerComposePath); os.IsNotExist(err) {
		// Try relative to repo root
		dockerComposePath = "deployments/docker-compose/docker-compose.yaml"
	}

	content, err := os.ReadFile(dockerComposePath)
	if err != nil {
		return "unknown"
	}

	// Pattern to match: cloudbaristaorg/serviceName:version
	pattern := fmt.Sprintf(`cloudbaristaorg/%s:([\d\.]+)`, serviceName)
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(string(content))

	if len(matches) >= 2 {
		return matches[1]
	}
	return "unknown"
}

// getCmModelVersionFromGoMod extracts cm-model version from go.mod
func getCmModelVersionFromGoMod() string {
	goModPath := "../../go.mod"

	// Try absolute path first
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		// Try relative to repo root
		goModPath = "go.mod"
	}

	content, err := os.ReadFile(goModPath)
	if err != nil {
		return "unknown"
	}

	// Pattern to match: github.com/cloud-barista/cm-model version
	pattern := `github\.com/cloud-barista/cm-model\s+(v[\d\.]+)`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(string(content))

	if len(matches) >= 2 {
		return matches[1]
	}
	return "unknown"
}

// generateMarkdownReport generates a markdown report for a specific CSP
func generateMarkdownReport(report *CSPTestReport) error {
	// Get the absolute path to the testresult directory
	execDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	// Create testresult directory relative to current working directory
	testResultDir := filepath.Join(execDir, "testresult")
	if err := os.MkdirAll(testResultDir, 0755); err != nil {
		return fmt.Errorf("failed to create testresult directory: %w", err)
	}

	filename := filepath.Join(testResultDir, fmt.Sprintf("beetle-test-results-%s.md",
		strings.ToLower(report.CSP)))

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create markdown file: %w", err)
	}
	defer file.Close()

	content := generateMarkdownContent(report)

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("failed to write markdown content: %w", err)
	}

	fmt.Printf("📝 Markdown report saved to: %s\n", filename)
	return nil
}

// generateMarkdownContent creates the markdown content based on the original document format
func generateMarkdownContent(report *CSPTestReport) string {
	var sb strings.Builder

	// Header
	sb.WriteString(fmt.Sprintf("# CM-Beetle test results for %s\n\n", strings.ToUpper(report.CSP)))

	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> This document presents comprehensive test results for CM-Beetle integration with ")
	sb.WriteString(fmt.Sprintf("%s cloud infrastructure.\n\n", strings.ToUpper(report.CSP)))

	// Environment and scenario
	sb.WriteString("## Environment and scenario\n\n")
	sb.WriteString("### Environment\n\n")
	sb.WriteString(fmt.Sprintf("- CM-Beetle: v0.4.0 (%s)\n", getGitHash()))
	sb.WriteString(fmt.Sprintf("- cm-model: %s\n", getCmModelVersionFromGoMod()))
	sb.WriteString(fmt.Sprintf("- CB-Tumblebug: v%s\n", getVersionFromDockerCompose("cb-tumblebug")))
	sb.WriteString(fmt.Sprintf("- CB-Spider: v%s\n", getVersionFromDockerCompose("cb-spider")))
	sb.WriteString(fmt.Sprintf("- CB-MapUI: v%s\n", getVersionFromDockerCompose("cb-mapui")))
	sb.WriteString(fmt.Sprintf("- Target CSP: %s\n", strings.ToUpper(report.CSP)))
	sb.WriteString(fmt.Sprintf("- Target Region: %s\n", report.Region))
	sb.WriteString(fmt.Sprintf("- CM-Beetle URL: %s\n", report.BeetleURL))
	sb.WriteString(fmt.Sprintf("- Namespace: %s\n", report.NamespaceID))
	sb.WriteString("- Test CLI: Custom automated testing tool\n")
	sb.WriteString(fmt.Sprintf("- Test Date: %s\n", report.TestDate))
	sb.WriteString(fmt.Sprintf("- Test Time: %s\n", report.TestTime))
	sb.WriteString(fmt.Sprintf("- Test Execution: %s\n\n", report.TestDateTime.Format("2006-01-02 15:04:05 MST")))

	sb.WriteString("### Scenario\n\n")
	sb.WriteString("1. Recommend a target model for computing infra via Beetle\n")
	sb.WriteString("1. Migrate the computing infra as defined in the target model via Beetle\n")
	sb.WriteString("1. List all MCIs via Beetle\n")
	sb.WriteString("1. List MCI IDs via Beetle\n")
	sb.WriteString("1. Get specific MCI details via Beetle\n")
	sb.WriteString("1. Delete the migrated computing infra via Beetle\n\n")

	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> Some long request/response bodies are in the collapsible section for better readability.\n\n")

	// Test result section
	sb.WriteString(fmt.Sprintf("## Test result for %s\n\n", strings.ToUpper(report.CSP)))

	// Test Results Summary
	sb.WriteString("### Test Results Summary\n\n")
	sb.WriteString("| Test | Endpoint | Status | Duration | Details |\n")
	sb.WriteString("|------|----------|--------|----------|----------|\n")

	for i, result := range report.TestResults {
		status := "✅ **PASS**"
		if !result.Success {
			status = "❌ **FAIL**"
		}

		var endpoint string
		switch i {
		case 0:
			endpoint = "`POST /beetle/recommendation/mci`"
		case 1:
			endpoint = fmt.Sprintf("`POST /beetle/migration/ns/%s/mci`", report.NamespaceID)
		case 2:
			endpoint = fmt.Sprintf("`GET /beetle/migration/ns/%s/mci`", report.NamespaceID)
		case 3:
			endpoint = fmt.Sprintf("`GET /beetle/migration/ns/%s/mci?option=id`", report.NamespaceID)
		case 4:
			endpoint = fmt.Sprintf("`GET /beetle/migration/ns/%s/mci/{{mciId}}`", report.NamespaceID)
		case 5:
			endpoint = fmt.Sprintf("`DELETE /beetle/migration/ns/%s/mci/{{mciId}}`", report.NamespaceID)
		}

		duration := result.Duration.Truncate(time.Millisecond)
		details := "Success"
		if !result.Success {
			if errorMsg, ok := result.Response["error"].(string); ok {
				details = errorMsg
			} else if msg, ok := result.Response["message"].(string); ok {
				details = msg
			} else {
				details = fmt.Sprintf("HTTP %d", result.StatusCode)
			}
		}

		sb.WriteString(fmt.Sprintf("| %d | %s | %s | %v | %s |\n",
			i+1, endpoint, status, duration, details))
	}

	sb.WriteString("\n")

	// Overall Summary
	passedCount := 0
	for _, result := range report.TestResults {
		if result.Success {
			passedCount++
		}
	}

	sb.WriteString(fmt.Sprintf("**Overall Result**: %d/%d tests passed", passedCount, len(report.TestResults)))
	if passedCount == len(report.TestResults) {
		sb.WriteString(" ✅\n\n")
	} else {
		sb.WriteString(" ❌\n\n")
	}

	sb.WriteString(fmt.Sprintf("**Total Duration**: %v\n\n", report.Summary.Duration))

	// Test time
	sb.WriteString(fmt.Sprintf("*Test executed on %s at %s (%s) using CM-Beetle automated test CLI*\n\n",
		report.TestDate,
		report.TestTime,
		report.TestDateTime.Format("2006-01-02 15:04:05 MST")))

	// Test 1: Recommendation
	sb.WriteString("### Recommend a target model for computing infra\n\n")
	sb.WriteString("> [!Note]\n")
	sb.WriteString("> `desiredCsp` and `desiredRegion` are required in the request body.\n\n")

	sb.WriteString("- API: `POST /beetle/recommendation/mci`\n")
	sb.WriteString("- Request body:\n\n")
	sb.WriteString("<details>\n")
	sb.WriteString("  <summary> <ins>Click to see the request body </ins> </summary>\n\n")
	sb.WriteString("```json\n")
	reqJSON, _ := json.MarshalIndent(report.RecommendationRequest, "", "  ")
	sb.WriteString(string(reqJSON))
	sb.WriteString("\n```\n\n")
	sb.WriteString("</details>\n\n")

	sb.WriteString("- Response body:\n\n")
	if report.RecommendationResponse != nil {
		sb.WriteString("<details>\n")
		sb.WriteString("  <summary> <ins>Click to see the response body</ins> </summary>\n\n")
		sb.WriteString("```json\n")
		respJSON, _ := json.MarshalIndent(report.RecommendationResponse, "", "  ")
		sb.WriteString(string(respJSON))
		sb.WriteString("\n```\n\n")
		sb.WriteString("</details>\n\n")
	} else {
		sb.WriteString("❌ **Test Failed**: No response received\n\n")
	}

	// Test 2: Migration
	sb.WriteString("### Migrate the computing infra as defined in the target model\n\n")
	sb.WriteString(fmt.Sprintf("- API: `POST /beetle/migration/ns/%s/mci`\n", report.NamespaceID))
	sb.WriteString(fmt.Sprintf("- nsId: `%s`\n", report.NamespaceID))
	sb.WriteString("- Request body: **same as the response from the previous step**\n")
	sb.WriteString("- Response body:\n\n")

	if report.MigrationResponse != nil {
		sb.WriteString("<details>\n")
		sb.WriteString("  <summary> <ins>Click to see the response body </ins> </summary>\n\n")
		sb.WriteString("```json\n")
		migJSON, _ := json.MarshalIndent(report.MigrationResponse, "", "  ")
		sb.WriteString(string(migJSON))
		sb.WriteString("\n```\n\n")
		sb.WriteString("</details>\n\n")
	} else {
		sb.WriteString("❌ **Test Failed**: Migration failed\n\n")
	}

	// Test 3: List MCIs
	sb.WriteString("### Get a list of MCIs\n\n")
	sb.WriteString(fmt.Sprintf("- API: `GET /beetle/migration/ns/%s/mci`\n", report.NamespaceID))
	sb.WriteString(fmt.Sprintf("- nsId: `%s`\n", report.NamespaceID))
	sb.WriteString("- Request body: None\n")
	sb.WriteString("- Response body:\n\n")

	if report.ListMCIResponse != nil {
		sb.WriteString("```json\n")
		listJSON, _ := json.MarshalIndent(report.ListMCIResponse, "", "  ")
		sb.WriteString(string(listJSON))
		sb.WriteString("\n```\n\n")
	} else {
		sb.WriteString("❌ **Test Failed**: No response received\n\n")
	}

	// Test 4: List MCI IDs
	sb.WriteString("### Get a list of MCI IDs\n\n")
	sb.WriteString(fmt.Sprintf("- API: `GET /beetle/migration/ns/%s/mci?option=id`\n", report.NamespaceID))
	sb.WriteString(fmt.Sprintf("- nsId: `%s`\n", report.NamespaceID))
	sb.WriteString("- Request body: None\n")
	sb.WriteString("- Response body:\n\n")

	if report.ListMCIIDsResponse != nil {
		sb.WriteString("```json\n")
		idsJSON, _ := json.MarshalIndent(report.ListMCIIDsResponse, "", "  ")
		sb.WriteString(string(idsJSON))
		sb.WriteString("\n```\n\n")
	} else {
		sb.WriteString("❌ **Test Failed**: No response received\n\n")
	}

	// Test 5: Get specific MCI (if available)
	if report.GetMCIResponse != nil {
		sb.WriteString("### Get a specific MCI\n\n")
		sb.WriteString(fmt.Sprintf("- API: `GET /beetle/migration/ns/%s/mci/{{mciId}}`\n", report.NamespaceID))
		sb.WriteString(fmt.Sprintf("- nsId: `%s`\n", report.NamespaceID))
		sb.WriteString("- Request body: None\n")
		sb.WriteString("- Response body:\n\n")
		sb.WriteString("<details>\n")
		sb.WriteString("  <summary> <ins>Click to see the response body </ins> </summary>\n\n")
		sb.WriteString("```json\n")
		getMCIJSON, _ := json.MarshalIndent(report.GetMCIResponse, "", "  ")
		sb.WriteString(string(getMCIJSON))
		sb.WriteString("\n```\n\n")
		sb.WriteString("</details>\n\n")
	}

	// Test 6: Delete MCI (if available)
	if len(report.DeleteMCIResponse) > 0 {
		sb.WriteString("### Delete the migrated computing infra\n\n")
		sb.WriteString(fmt.Sprintf("- API: `DELETE /beetle/migration/ns/%s/mci/{{mciId}}`\n", report.NamespaceID))
		sb.WriteString(fmt.Sprintf("- nsId: `%s`\n", report.NamespaceID))
		sb.WriteString("- Request body: None\n")
		sb.WriteString("- Response body:\n\n")
		sb.WriteString("```json\n")
		delJSON, _ := json.MarshalIndent(report.DeleteMCIResponse, "", "  ")
		sb.WriteString(string(delJSON))
		sb.WriteString("\n```\n\n")
	}

	return sb.String()
}

// Animation functions
func animatedSleep(duration time.Duration, message string) {
	spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	end := time.Now().Add(duration)
	i := 0

	fmt.Printf("\r%s %s", spinner[i%len(spinner)], message)

	for time.Now().Before(end) {
		time.Sleep(100 * time.Millisecond)
		i++
		fmt.Printf("\r%s %s", spinner[i%len(spinner)], message)
	}
	fmt.Printf("\r✅ %s - Complete!    \n", message)
}
