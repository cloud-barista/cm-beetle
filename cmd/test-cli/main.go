// Package main is the starting point of CM-Beetle Test CLI
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/ssh"

	// Import Beetle's existing packages
	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	tbclient "github.com/cloud-barista/cm-beetle/pkg/client/tumblebug"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
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
	BeetleURL             string                     `json:"beetleUrl"`
	NamespaceID           string                     `json:"namespaceId"`
	DesiredCspRegionPairs []cloudmodel.CloudProperty `json:"desiredCspRegionPairs"`
	RequestBodyFile       string                     `json:"requestBodyFile"`
	AuthConfigFile        string                     `json:"authConfigFile,omitempty"`
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
	TestName     string                 `json:"testName"`
	StartTime    time.Time              `json:"startTime"`
	EndTime      time.Time              `json:"endTime"`
	Duration     time.Duration          `json:"duration"`
	Success      bool                   `json:"success"`
	Skipped      bool                   `json:"skipped"` // True if test was skipped
	StatusCode   int                    `json:"statusCode"`
	Response     map[string]interface{} `json:"response"`
	Error        string                 `json:"error,omitempty"`
	ErrorMessage string                 `json:"errorMessage,omitempty"` // Human-readable error message
	ErrorDetails string                 `json:"errorDetails,omitempty"` // Additional error details
	RequestURL   string                 `json:"requestUrl,omitempty"`   // Request URL for debugging
	RequestBody  interface{}            `json:"requestBody,omitempty"`  // Request body for debugging
}

// TestSuite holds all test results
type TestSuite struct {
	Config          TestConfig             `json:"config"`
	Results         []TestResults          `json:"results"`
	CspResults      map[string]TestResults `json:"cspResults"` // Results per CSP-Region pair
	TotalTests      int                    `json:"totalTests"`
	TotalCspPairs   int                    `json:"totalCspPairs"`
	PassedTests     int                    `json:"passedTests"`
	FailedTests     int                    `json:"failedTests"`
	SkippedTests    int                    `json:"skippedTests"`
	PassedCspPairs  int                    `json:"passedCspPairs"`
	FailedCspPairs  int                    `json:"failedCspPairs"`
	SkippedCspPairs int                    `json:"skippedCspPairs"`
	OverallTime     time.Duration          `json:"overallTime"`
}

var (
	configFile         = flag.String("config", "testdata/config-multi-csp-and-region-pair.json", "Path to config file")
	verbose            = flag.Bool("verbose", false, "Enable verbose output")
	skipRemainingTests bool
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

// extractErrorDetails extracts meaningful error information from error responses
func extractErrorDetails(err error, statusCode int) (string, string) {
	if err == nil {
		// Handle HTTP errors without Go error
		if statusCode >= 400 {
			return fmt.Sprintf("HTTP %d error", statusCode), fmt.Sprintf("HTTP %d", statusCode)
		}
		return "", ""
	}

	errorStr := err.Error()
	errorMessage := errorStr
	errorDetails := ""

	// Try to parse JSON error response if it looks like JSON
	if strings.Contains(errorStr, "{") && strings.Contains(errorStr, "}") {
		jsonStart := strings.Index(errorStr, "{")
		jsonEnd := strings.LastIndex(errorStr, "}") + 1
		if jsonStart >= 0 && jsonEnd > jsonStart {
			jsonStr := errorStr[jsonStart:jsonEnd]
			var errorResponse map[string]interface{}
			if err := json.Unmarshal([]byte(jsonStr), &errorResponse); err == nil {
				if message, ok := errorResponse["message"].(string); ok {
					errorMessage = message
				} else if details, ok := errorResponse["error"].(string); ok {
					errorMessage = details
				}
				if details, ok := errorResponse["details"].(string); ok {
					errorDetails = details
				}
			}
		}
	}

	// Add HTTP status context if available
	if statusCode > 0 && errorDetails == "" {
		errorDetails = fmt.Sprintf("HTTP %d", statusCode)
	} else if statusCode > 0 {
		errorDetails = fmt.Sprintf("HTTP %d: %s", statusCode, errorDetails)
	}

	return errorMessage, errorDetails
}

// populateErrorInfo populates error information in TestResults
func populateErrorInfo(result *TestResults, err error, statusCode int, requestURL string, requestBody interface{}) {
	result.Success = false
	result.StatusCode = statusCode
	result.RequestURL = requestURL
	result.RequestBody = requestBody

	// Handle case where err might be nil but we have HTTP error status
	if err != nil {
		result.Error = err.Error()
	} else {
		// Create a generic error message for HTTP errors without Go error
		result.Error = fmt.Sprintf("HTTP %d error", statusCode)
	}

	errorMessage, errorDetails := extractErrorDetails(err, statusCode)
	result.ErrorMessage = errorMessage
	result.ErrorDetails = errorDetails

	// Ensure Response is initialized for backward compatibility
	if result.Response == nil {
		result.Response = make(map[string]interface{})
	}

	// Add error info to response for backward compatibility
	if err != nil {
		result.Response["error"] = err.Error()
	} else {
		result.Response["error"] = fmt.Sprintf("HTTP %d error", statusCode)
	}
	if errorMessage != "" {
		result.Response["message"] = errorMessage
	}
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
		cspRegionPairs = []cloudmodel.CloudProperty{
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
		TotalTests:    7, // Total number of API tests per CSP-Region pair
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
	client.SetTimeout(45 * time.Minute) // Increased timeout to 45 minutes for all operations

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

	// Output source infrastructure summary before starting tests
	log.Info().Msgf("%s", "\n"+strings.Repeat("=", 60)+"\n")
	log.Info().Msg("SOURCE INFRASTRUCTURE SUMMARY")
	log.Info().Msgf("%s", strings.Repeat("=", 60)+"\n")

	sourceSummaryResult := runSourceSummaryTest(client, config, onpremInfraModel)
	if !sourceSummaryResult.Success {
		log.Warn().Msg("Failed to generate source infrastructure summary, but continuing with tests...")
	}

	// Test each CSP-Region pair sequentially
	for pairIndex, cspPair := range cspRegionPairs {
		// Reset skip flag for each CSP-Region pair
		skipRemainingTests = false

		displayName := fmt.Sprintf("%s-%s", cspPair.Csp, cspPair.Region)

		// Add delay between CSP-Region pairs (except for the first one)
		if pairIndex > 0 {
			animatedSleep(10*time.Second, fmt.Sprintf("Waiting for a while before testing %s", displayName))
		}

		log.Info().Msgf("%s", "\n"+strings.Repeat("=", 60)+"\n")
		log.Info().Msgf("Testing CSP-Region Pair %d/%d: %s", pairIndex+1, len(cspRegionPairs), displayName)
		log.Info().Str("csp", cspPair.Csp).Str("region", cspPair.Region).Msg("Starting CSP-Region pair test")
		log.Info().Msgf("%s", strings.Repeat("=", 60)+"\n")

		pairStartTime := time.Now()
		pairPassed := 0
		pairFailed := 0
		pairSkipped := 0

		// Create RecommendVmInfraRequest for this CSP-Region pair
		recommendRequest := controller.RecommendVmInfraRequest{
			DesiredCspAndRegionPair: cloudmodel.CloudProperty{
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

		/*
		 * Test 1: POST /beetle/recommendation/mci
		 */
		recommendedVmInfra, result1 := runRecommendationTest(client, config, cspPair, recommendRequest, displayName)
		suite.Results = append(suite.Results, result1)
		cspReport.TestResults = append(cspReport.TestResults, result1)

		if !result1.Success {
			pairFailed++
			// Count remaining tests as skipped
			pairSkipped = suite.TotalTests - 1 // All tests except the first one that failed
			skipRemainingTests = true
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
					"passedTests":  pairPassed,
					"failedTests":  pairFailed,
					"skippedTests": pairSkipped,
					"totalTests":   suite.TotalTests,
				},
			}

			// Update suite counters
			suite.PassedTests += pairPassed
			suite.FailedTests += pairFailed
			suite.SkippedTests += pairSkipped
			suite.FailedCspPairs++

			// Print pair summary
			log.Info().Msgf("\n--- Summary for %s ---", displayName)
			log.Info().Int("passed", pairPassed).Int("total", suite.TotalTests).Msgf("Tests Passed: %d/%d", pairPassed, suite.TotalTests)
			log.Info().Int("failed", pairFailed).Int("total", suite.TotalTests).Msgf("Tests Failed: %d/%d", pairFailed, suite.TotalTests)
			log.Info().Dur("duration", pairDuration).Msgf("Duration: %v", pairDuration)
			log.Warn().Msg("Status: ❌ RECOMMENDATION TEST FAILED - REMAINING TESTS SKIPPED")

		}

		// Test 1 succeeded, continue with processing
		if result1.Success {
			if structuredResponse, err := convertMapToRecommendVmInfraResponse(result1.Response); err == nil {
				cspReport.RecommendationResponse = structuredResponse
				recommendedVmInfra = *structuredResponse
			} else {
				log.Warn().Err(err).Msg("Failed to convert recommendation response")
			}
			pairPassed++
		} else if result1.Skipped {
			pairSkipped++
		}

		/*
		 * Test 2: POST /beetle/migration/ns/{nsId}/mci
		 */
		var result2 TestResults
		if skipRemainingTests {
			// Create a skipped test result
			result2 = TestResults{
				TestName:     "Test 2: POST /beetle/migration/ns/{nsId}/mci",
				StartTime:    time.Now(),
				EndTime:      time.Now(),
				Duration:     0,
				Success:      false,
				Skipped:      true,
				StatusCode:   0,
				Response:     map[string]interface{}{},
				Error:        "Test skipped due to previous test failure",
				ErrorMessage: "Test skipped due to previous test failure",
				RequestURL:   fmt.Sprintf("%s/beetle/migration/ns/%s/mci", config.BeetleURL, config.NamespaceID),
			}
			log.Warn().Msgf("Test 2 skipped for %s due to previous test failure", displayName)
		} else {
			// Convert RecommendVmInfraResponse to MigrateInfraRequest
			migrationRequest := controller.MigrateInfraRequest(recommendedVmInfra)
			result2 = runMigrationTest(client, config, migrationRequest, displayName)
		}

		suite.Results = append(suite.Results, result2)
		cspReport.TestResults = append(cspReport.TestResults, result2)

		// Handle Test 2 result consistently
		if !result2.Success && !skipRemainingTests {
			// Test 2 was actually executed and failed
			skipRemainingTests = true
			log.Error().Str("csp", cspPair.Csp).Str("region", cspPair.Region).Msg("Test 2 (Migration) failed. Skipping remaining tests for this CSP-Region pair.")
		}

		// Count test results consistently
		if result2.Skipped {
			pairSkipped++
		} else if !result2.Success {
			pairFailed++
		} else {
			pairPassed++
		}

		// Test 2 succeeded, continue with processing
		if result2.Success {
			if structuredResponse, err := convertMapToMigrateInfraResponse(result2.Response); err == nil {
				cspReport.MigrationResponse = structuredResponse
			} else {
				log.Warn().Err(err).Msg("Failed to convert migration response")
			}

			// Extract mciId from response
			if id, ok := result2.Response["id"].(string); ok && id != "" {
				mciId = id
			} else if name, ok := result2.Response["name"].(string); ok && name != "" {
				mciId = name // Use name as fallback
			}
		}

		/*
		 * Test 3: GET /beetle/migration/ns/{nsId}/mci
		 */
		var result3 TestResults
		if skipRemainingTests {
			// Create a skipped test result
			result3 = TestResults{
				TestName:     "Test 3: GET /beetle/migration/ns/{nsId}/mci",
				StartTime:    time.Now(),
				EndTime:      time.Now(),
				Duration:     0,
				Success:      false,
				Skipped:      true,
				StatusCode:   0,
				Response:     map[string]interface{}{},
				Error:        "Test skipped due to previous test failure",
				ErrorMessage: "Test skipped due to previous test failure",
				RequestURL:   fmt.Sprintf("%s/beetle/migration/ns/%s/mci", config.BeetleURL, config.NamespaceID),
			}
			log.Warn().Msgf("Test 3 skipped for %s due to previous test failure", displayName)
		} else {
			result3 = runListMciTest(client, config, displayName)
			if result3.Success {
				if structuredResponse, err := convertMapToMciInfoList(result3.Response); err == nil {
					cspReport.ListMCIResponse = structuredResponse
				} else {
					log.Warn().Err(err).Msg("Failed to convert list MCI response")
				}
			}
		}

		suite.Results = append(suite.Results, result3)
		cspReport.TestResults = append(cspReport.TestResults, result3)

		if result3.Skipped {
			pairSkipped++
		} else if !result3.Success {
			pairFailed++
		} else {
			pairPassed++
		}

		/*
		 * Test 4: GET /beetle/migration/ns/{nsId}/mci?option=id
		 */
		var result4 TestResults
		if skipRemainingTests {
			// Create a skipped test result
			result4 = TestResults{
				TestName:     "Test 4: GET /beetle/migration/ns/{nsId}/mci?option=id",
				StartTime:    time.Now(),
				EndTime:      time.Now(),
				Duration:     0,
				Success:      false,
				Skipped:      true,
				StatusCode:   0,
				Response:     map[string]interface{}{},
				Error:        "Test skipped due to previous test failure",
				ErrorMessage: "Test skipped due to previous test failure",
				RequestURL:   fmt.Sprintf("%s/beetle/migration/ns/%s/mci?option=id", config.BeetleURL, config.NamespaceID),
			}
			log.Warn().Msgf("Test 4 skipped for %s due to previous test failure", displayName)
		} else {
			result4 = runListMciIdsTest(client, config, displayName)
			if result4.Success {
				if structuredResponse, err := convertMapToIdList(result4.Response); err == nil {
					cspReport.ListMCIIDsResponse = structuredResponse
				} else {
					log.Warn().Err(err).Msg("Failed to convert list MCI IDs response")
				}
			}
		}

		suite.Results = append(suite.Results, result4)
		cspReport.TestResults = append(cspReport.TestResults, result4)

		if result4.Skipped {
			pairSkipped++
		} else if !result4.Success {
			pairFailed++
		} else {
			pairPassed++
		}

		/*
		 * Test 5: GET /beetle/migration/ns/{nsId}/mci/{mciId}
		 */
		var result5 TestResults
		if skipRemainingTests {
			// Create a skipped test result
			result5 = TestResults{
				TestName:     "Test 5: GET /beetle/migration/ns/{nsId}/mci/{mciId}",
				StartTime:    time.Now(),
				EndTime:      time.Now(),
				Duration:     0,
				Success:      false,
				Skipped:      true,
				StatusCode:   0,
				Response:     map[string]interface{}{},
				Error:        "Test skipped due to previous test failure",
				ErrorMessage: "Test skipped due to previous test failure",
				RequestURL:   fmt.Sprintf("%s/beetle/migration/ns/%s/mci/%s", config.BeetleURL, config.NamespaceID, mciId),
			}
			log.Warn().Msgf("Test 5 skipped for %s due to previous test failure", displayName)
		} else {
			result5, _ = runGetMciTest(client, config, mciId, displayName)
			if result5.Success {
				if structuredResponse, err := convertMapToVmInfraInfo(result5.Response); err == nil {
					cspReport.GetMCIResponse = structuredResponse
				} else {
					log.Warn().Err(err).Msg("Failed to convert get MCI response")
				}
			}
		}

		suite.Results = append(suite.Results, result5)
		cspReport.TestResults = append(cspReport.TestResults, result5)

		if result5.Skipped {
			pairSkipped++
		} else if !result5.Success {
			pairFailed++
		} else {
			pairPassed++
		}

		/*
		 * Test 6: Remote Command to check accessibility of migrated VM
		 */
		var result6 TestResults
		if skipRemainingTests {
			// Create a skipped test result
			result6 = TestResults{
				TestName:     "Test 6: Remote Command Accessibility Check",
				StartTime:    time.Now(),
				EndTime:      time.Now(),
				Duration:     0,
				Success:      false,
				Skipped:      true,
				StatusCode:   0,
				Response:     map[string]interface{}{},
				Error:        "Test skipped due to previous test failure",
				ErrorMessage: "Test skipped due to previous test failure",
				RequestURL:   "N/A (SSH command)",
			}
			log.Warn().Msgf("Test 6 skipped for %s due to previous test failure", displayName)
		} else {
			result6 = runRemoteCommandTest(client, config, mciId, displayName)
		}

		suite.Results = append(suite.Results, result6)
		cspReport.TestResults = append(cspReport.TestResults, result6)

		if result6.Skipped {
			pairSkipped++
		} else if !result6.Success {
			pairFailed++
		} else {
			pairPassed++
		}

		/*
		 * Target Infrastructure Summary: GET /beetle/summary/target/ns/{nsId}/mci/{mciId}
		 */
		if !skipRemainingTests {
			targetSummaryResult := runTargetSummaryTest(client, config, mciId, cspPair.Csp, displayName)
			if !targetSummaryResult.Success {
				log.Warn().Str("csp", displayName).Msg("Failed to generate target infrastructure summary")
			}
		} else {
			log.Info().Msgf("Target summary skipped for %s due to previous test failure", displayName)
		}

		/*
		 * Migration Report: POST /beetle/report/migration/ns/{nsId}/mci/{mciId}
		 */
		if !skipRemainingTests {
			migrationReportResult := runMigrationReportTest(client, config, onpremInfraModel, mciId, cspPair.Csp, displayName)
			if !migrationReportResult.Success {
				log.Warn().Str("csp", displayName).Msg("Failed to generate migration report")
			}
		} else {
			log.Info().Msgf("Migration report skipped for %s due to previous test failure", displayName)
		}

		/*
		 * Test 7: DELETE /beetle/migration/ns/{nsId}/mci/{mciId}
		 */
		var result7 TestResults
		if skipRemainingTests {
			// Create a skipped test result
			result7 = TestResults{
				TestName:     "Test 7: DELETE /beetle/migration/ns/{nsId}/mci/{mciId}",
				StartTime:    time.Now(),
				EndTime:      time.Now(),
				Duration:     0,
				Success:      false,
				Skipped:      true,
				StatusCode:   0,
				Response:     map[string]interface{}{},
				Error:        "Test skipped due to previous test failure",
				ErrorMessage: "Test skipped due to previous test failure",
				RequestURL:   fmt.Sprintf("%s/beetle/migration/ns/%s/mci/%s", config.BeetleURL, config.NamespaceID, mciId),
			}
			log.Warn().Msgf("Test 7 skipped for %s due to previous test failure", displayName)
		} else {
			result7, _ = runDeleteMciTest(client, config, mciId, displayName)
			if result7.Success {
				cspReport.DeleteMCIResponse = result7.Response
			}
		}

		suite.Results = append(suite.Results, result7)
		cspReport.TestResults = append(cspReport.TestResults, result7)

		if result7.Skipped {
			pairSkipped++
		} else if !result7.Success {
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
				"csp":          cspPair.Csp,
				"region":       cspPair.Region,
				"passedTests":  pairPassed,
				"failedTests":  pairFailed,
				"skippedTests": pairSkipped,
				"totalTests":   suite.TotalTests,
			},
		}

		// Update suite counters
		suite.PassedTests += pairPassed
		suite.FailedTests += pairFailed
		suite.SkippedTests += pairSkipped
		if pairFailed == 0 {
			suite.PassedCspPairs++
		} else {
			suite.FailedCspPairs++
		}

		// Print pair summary
		log.Info().Msgf("\n--- Summary for %s ---", displayName)
		log.Info().Int("passed", pairPassed).Int("total", suite.TotalTests).Msgf("Tests Passed: %d/%d", pairPassed, suite.TotalTests)
		log.Info().Int("failed", pairFailed).Int("total", suite.TotalTests).Msgf("Tests Failed: %d/%d", pairFailed, suite.TotalTests)
		if pairSkipped > 0 {
			log.Info().Int("skipped", pairSkipped).Int("total", suite.TotalTests).Msgf("Tests Skipped: %d/%d", pairSkipped, suite.TotalTests)
		}
		log.Info().Dur("duration", pairDuration).Msgf("Duration: %v", pairDuration)
		if pairFailed == 0 {
			log.Info().Msg("Status: ✅ ALL TESTS PASSED")
		} else {
			log.Warn().Msg("Status: ❌ SOME TESTS FAILED")
		}
	}

	suite.OverallTime = time.Since(startTime)

	// Print overall summary
	log.Info().Msgf("%s", "\n"+strings.Repeat("=", 60)+"\n")
	log.Info().Msg("OVERALL TEST SUMMARY")
	log.Info().Msgf("%s", strings.Repeat("=", 60)+"\n")
	log.Info().Int("total", suite.TotalCspPairs).Msg("Total CSP-Region Pairs")
	log.Info().Int("successful", suite.PassedCspPairs).Msg("Successful Pairs")
	log.Info().Int("failed", suite.FailedCspPairs).Msg("Failed Pairs")
	log.Info().Int("totalTests", suite.TotalTests*suite.TotalCspPairs).Msg("Total Tests")
	log.Info().Int("passed", suite.PassedTests).Msg("Passed Tests")
	log.Info().Int("failed", suite.FailedTests).Msg("Failed Tests")
	if suite.SkippedTests > 0 {
		log.Info().Int("skipped", suite.SkippedTests).Msg("Skipped Tests")
	}
	log.Info().Dur("overallTime", suite.OverallTime).Msgf("Overall Time: %v", suite.OverallTime)

	log.Info().Msg("\nPer CSP-Region Results:")
	for name, result := range suite.CspResults {
		status := "✅"
		if !result.Success {
			status = "❌"
		}
		passedTests := result.Response["passedTests"].(int)
		totalTests := result.Response["totalTests"].(int)
		skippedTests := 0
		if val, ok := result.Response["skippedTests"].(int); ok {
			skippedTests = val
		}

		if skippedTests > 0 {
			log.Info().
				Str("status", status).
				Str("csp", name).
				Int("passed", passedTests).
				Int("skipped", skippedTests).
				Int("total", totalTests).
				Dur("duration", result.Duration).
				Msgf("%s %s - %d/%d tests passed, %d skipped (Duration: %v)",
					status, name, passedTests, totalTests, skippedTests, result.Duration)
		} else {
			log.Info().
				Str("status", status).
				Str("csp", name).
				Int("passed", passedTests).
				Int("total", totalTests).
				Dur("duration", result.Duration).
				Msgf("%s %s - %d/%d tests passed (Duration: %v)",
					status, name, passedTests, totalTests, result.Duration)
		}
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
func runRecommendationTest(client *resty.Client, config TestConfig, cspPair cloudmodel.CloudProperty, recommendRequest controller.RecommendVmInfraRequest, displayName string) (controller.RecommendVmInfraResponse, TestResults) {
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
		populateErrorInfo(&result, err, 0, url, recommendRequest)
		fmt.Printf("❌ Test 1 failed: %s\n", result.ErrorMessage)
		log.Error().Err(err).Str("url", url).Msg("Recommendation test failed")
		return controller.RecommendVmInfraResponse{}, result
	}

	result.Success = true
	result.StatusCode = 200
	result.RequestURL = url
	result.RequestBody = recommendRequest

	// Convert struct response to map for TestResults compatibility
	responseMap := make(map[string]interface{})
	jsonBytes, _ := json.Marshal(response)
	json.Unmarshal(jsonBytes, &responseMap)
	result.Response = responseMap
	fmt.Println("✅ Test 1 passed")
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

	var response controller.MigrateInfraResponse
	var mciId string

	// Setup cleanup defer function
	defer func() {
		// Check if migration failed and MCI ID is available for cleanup
		if strings.Contains(strings.ToLower(response.Status), "failed") && mciId != "" {
			cleanupMci(client, config, mciId, displayName)
		}
	}()

	// Log API call details
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/mci", config.BeetleURL, config.NamespaceID)
	log.Debug().Msgf("API Request URL: %s", url)

	// Log request body
	log.Debug().Msgf("API Request Body: %+v", migrationRequestBody)

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

	// Extract MCI ID from response for potential cleanup
	if response.Id != "" {
		mciId = response.Id
	} else if response.Name != "" {
		mciId = response.Name
	}

	// Log response body
	log.Debug().Msgf("API Response Body: %+v", response)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	if err != nil {
		result.Success = false
		result.Error = err.Error()
		result.StatusCode = 0
		fmt.Printf("❌ Test 2 failed: %s\n", result.Error)
	} else if strings.Contains(strings.ToLower(response.Status), "failed") {
		result.Success = false
		result.Error = fmt.Errorf("failed to migrate infra (MCI status: %s)", response.Status).Error()
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
			Skipped:  true,
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

// cleanupMci performs cleanup by deleting MCI when migration fails
func cleanupMci(client *resty.Client, config TestConfig, mciId, displayName string) {
	if mciId == "" {
		log.Warn().Msg("Cleanup skipped: No MCI ID available")
		return
	}

	log.Info().Str("mciId", mciId).Str("csp", displayName).Msg("Starting cleanup: Deleting failed MCI")

	// Wait before API call for stability
	time.Sleep(2 * time.Second)

	// Log API call details
	url := fmt.Sprintf("%s/beetle/migration/ns/%s/mci/%s?option=terminate", config.BeetleURL, config.NamespaceID, mciId)
	log.Debug().Msgf("Cleanup API Request URL: %s", url)

	var response map[string]interface{}

	// Make HTTP request directly with resty to capture full error response
	resp, err := client.R().
		SetResult(&response).
		SetError(&response).
		Delete(url)

	// Log response details
	log.Debug().Msgf("Cleanup HTTP Status: %s", resp.Status())
	log.Debug().Msgf("Cleanup Response Body: %+v", response)

	if err != nil || resp.StatusCode() >= 400 {
		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode()
		}
		log.Warn().Err(err).Str("url", url).Int("statusCode", statusCode).Msg("Cleanup failed: Unable to delete MCI")
	} else {
		log.Info().Str("mciId", mciId).Str("csp", displayName).Msg("Cleanup completed: MCI successfully deleted")
	}
}

// runDeleteMciTest performs Test 7: DELETE /beetle/migration/ns/{nsId}/mci/{mciId}
func runDeleteMciTest(client *resty.Client, config TestConfig, mciId, displayName string) (TestResults, bool) {
	if mciId == "" {
		fmt.Println("⚠️  Test 7 skipped: No MCI ID available")
		return TestResults{
			TestName: fmt.Sprintf("DELETE /beetle/migration/ns/{nsId}/mci/{mciId} (%s)", displayName),
			Success:  false,
			Skipped:  true,
			Error:    "MCI ID not available from previous tests",
		}, false
	}

	fmt.Printf("\n--- Test 7: DELETE /beetle/migration/ns/%s/mci/%s?option=terminate ---\n", config.NamespaceID, mciId)

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

	// Make HTTP request directly with resty to capture full error response
	resp, err := client.R().
		SetResult(&response).
		SetError(&response). // This captures error response body
		Delete(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	// Log response details
	log.Debug().Msgf("HTTP Status: %s", resp.Status())
	log.Debug().Msgf("Response Body: %+v", response)

	if err != nil || resp.StatusCode() >= 400 {
		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode()
		}

		// Create a more specific error if we only have HTTP error without Go error
		var finalErr error = err
		if err == nil && resp.StatusCode() >= 400 {
			// Create error from HTTP response body if available
			if len(response) > 0 {
				if respBytes, jsonErr := json.Marshal(response); jsonErr == nil {
					finalErr = fmt.Errorf("HTTP %d: %s", statusCode, string(respBytes))
				} else {
					finalErr = fmt.Errorf("HTTP %d error", statusCode)
				}
			} else {
				finalErr = fmt.Errorf("HTTP %d error", statusCode)
			}
		}

		populateErrorInfo(&result, finalErr, statusCode, url, nil)

		// If we have error response, include it in the result
		if len(response) > 0 {
			result.Response = response
		}

		fmt.Printf("❌ Test 6 failed: %s\n", result.ErrorMessage)
		log.Error().Err(err).Str("url", url).Int("statusCode", statusCode).Msg("Delete MCI test failed")
		return result, false
	}

	result.Success = true
	result.StatusCode = 200
	result.RequestURL = url
	result.RequestBody = nil
	result.Response = response
	fmt.Println("✅ Test 7 passed")
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
	sb.WriteString(fmt.Sprintf("- CM-Beetle: v0.4.1 (%s)\n", getGitHash()))
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

	// ========================================================================
	// First Section: Summary Information
	// ========================================================================

	// Test result section
	sb.WriteString(fmt.Sprintf("## Test result for %s\n\n", strings.ToUpper(report.CSP)))

	// Test Results Summary
	sb.WriteString("### Test Results Summary\n\n")
	sb.WriteString("| Test | Endpoint | Status | Duration | Details |\n")
	sb.WriteString("|------|----------|--------|----------|----------|\n")

	for i, result := range report.TestResults {
		status := "✅ **PASS**"
		if result.Skipped {
			status = "⏭️ **SKIP**"
		} else if !result.Success {
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
			endpoint = "Remote Command Accessibility Check" // Test 6: SSH connectivity check
		case 6:
			endpoint = fmt.Sprintf("`DELETE /beetle/migration/ns/%s/mci/{{mciId}}`", report.NamespaceID) // Test 7: Delete MCI
		default:
			endpoint = "" // Handle any additional tests
		}

		duration := result.Duration.Truncate(time.Millisecond)
		details := "Pass"
		if result.Skipped {
			details = "Skip"
		} else if !result.Success {
			details = "Fail"
		}

		sb.WriteString(fmt.Sprintf("| %d | %s | %s | %v | %s |\n",
			i+1, endpoint, status, duration, details))
	}

	sb.WriteString("\n")

	// Overall Summary
	passedCount := 0
	skippedCount := 0
	for _, result := range report.TestResults {
		if result.Success {
			passedCount++
		} else if result.Skipped {
			skippedCount++
		}
	}

	if skippedCount > 0 {
		sb.WriteString(fmt.Sprintf("**Overall Result**: %d/%d tests passed, %d skipped", passedCount, len(report.TestResults), skippedCount))
	} else {
		sb.WriteString(fmt.Sprintf("**Overall Result**: %d/%d tests passed", passedCount, len(report.TestResults)))
	}
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

	// ========================================================================
	// Second Section: Detailed Test Case Information
	// ========================================================================

	sb.WriteString("---\n\n")
	sb.WriteString("## Detailed Test Case Results\n\n")
	sb.WriteString("> [!INFO]\n")
	sb.WriteString("> This section provides detailed information for each test case, including API request information and response details.\n\n")

	// Test Case 1: Recommendation
	sb.WriteString("### Test Case 1: Recommend a target model for computing infra\n\n")

	// API request information
	sb.WriteString("#### 1.1 API Request Information\n\n")
	sb.WriteString("- **API Endpoint**: `POST /beetle/recommendation/mci`\n")
	sb.WriteString("- **Purpose**: Get infrastructure recommendations for migration\n")
	sb.WriteString("- **Required Parameters**: `desiredCsp` and `desiredRegion` in request body\n\n")

	sb.WriteString("**Request Body**:\n\n")
	sb.WriteString("<details>\n")
	sb.WriteString("  <summary> <ins>Click to see the request body </ins> </summary>\n\n")
	sb.WriteString("```json\n")
	reqJSON, _ := json.MarshalIndent(report.RecommendationRequest, "", "  ")
	sb.WriteString(string(reqJSON))
	sb.WriteString("\n```\n\n")
	sb.WriteString("</details>\n\n")

	// API response information
	sb.WriteString("#### 1.2 API Response Information\n\n")
	if report.RecommendationResponse != nil {
		sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
		sb.WriteString("- **Response**: Infrastructure recommendation generated successfully\n\n")
		sb.WriteString("**Response Body**:\n\n")
		sb.WriteString("<details>\n")
		sb.WriteString("  <summary> <ins>Click to see the response body</ins> </summary>\n\n")
		sb.WriteString("```json\n")
		respJSON, _ := json.MarshalIndent(report.RecommendationResponse, "", "  ")
		sb.WriteString(string(respJSON))
		sb.WriteString("\n```\n\n")
		sb.WriteString("</details>\n\n")
	} else {
		sb.WriteString("- **Status**: ❌ **FAILED**\n")
		sb.WriteString("- **Error**: No response received\n\n")
		// Add detailed error information if available
		if len(report.TestResults) > 0 {
			result := report.TestResults[0] // First test is recommendation
			if result.ErrorMessage != "" {
				sb.WriteString("**Error Message**:\n\n```\n")
				sb.WriteString(result.ErrorMessage)
				sb.WriteString("\n```\n\n")
			}
			if result.ErrorDetails != "" {
				sb.WriteString(fmt.Sprintf("**Error Details**: %s\n\n", result.ErrorDetails))
			}
			if result.RequestURL != "" {
				sb.WriteString(fmt.Sprintf("**Request URL**: `%s`\n\n", result.RequestURL))
			}
		}
	}

	// Test Case 2: Migration
	sb.WriteString("### Test Case 2: Migrate the computing infra as defined in the target model\n\n")

	// API request information
	sb.WriteString("#### 2.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `POST /beetle/migration/ns/%s/mci`\n", report.NamespaceID))
	sb.WriteString("- **Purpose**: Create and migrate infrastructure based on recommendation\n")
	sb.WriteString(fmt.Sprintf("- **Namespace ID**: `%s`\n", report.NamespaceID))
	sb.WriteString("- **Request Body**: Uses the response from the previous recommendation step\n\n")

	// API response information
	sb.WriteString("#### 2.2 API Response Information\n\n")
	if report.MigrationResponse != nil {
		sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
		sb.WriteString("- **Response**: Infrastructure migration completed successfully\n\n")
		sb.WriteString("**Response Body**:\n\n")
		sb.WriteString("<details>\n")
		sb.WriteString("  <summary> <ins>Click to see the response body </ins> </summary>\n\n")
		sb.WriteString("```json\n")
		migJSON, _ := json.MarshalIndent(report.MigrationResponse, "", "  ")
		sb.WriteString(string(migJSON))
		sb.WriteString("\n```\n\n")
		sb.WriteString("</details>\n\n")
	} else {
		sb.WriteString("- **Status**: ❌ **FAILED**\n")
		sb.WriteString("- **Error**: Migration failed\n\n")
		// Add detailed error information if available
		if len(report.TestResults) > 1 {
			result := report.TestResults[1] // Second test is migration
			if result.ErrorMessage != "" {
				sb.WriteString("**Error Message**:\n\n```\n")
				sb.WriteString(result.ErrorMessage)
				sb.WriteString("\n```\n\n")
			}
			if result.ErrorDetails != "" {
				sb.WriteString(fmt.Sprintf("**Error Details**: %s\n\n", result.ErrorDetails))
			}
			if result.RequestURL != "" {
				sb.WriteString(fmt.Sprintf("**Request URL**: `%s`\n\n", result.RequestURL))
			}
		}
	}

	// Test Case 3: List MCIs
	sb.WriteString("### Test Case 3: Get a list of MCIs\n\n")

	// API request information
	sb.WriteString("#### 3.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `GET /beetle/migration/ns/%s/mci`\n", report.NamespaceID))
	sb.WriteString("- **Purpose**: Retrieve all Multi-Cloud Infrastructure instances\n")
	sb.WriteString(fmt.Sprintf("- **Namespace ID**: `%s`\n", report.NamespaceID))
	sb.WriteString("- **Request Body**: None (GET request)\n\n")

	// API response information
	sb.WriteString("#### 3.2 API Response Information\n\n")
	if report.ListMCIResponse != nil {
		sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
		sb.WriteString("- **Response**: MCI list retrieved successfully\n\n")
		sb.WriteString("**Response Body**:\n\n")
		sb.WriteString("```json\n")
		listJSON, _ := json.MarshalIndent(report.ListMCIResponse, "", "  ")
		sb.WriteString(string(listJSON))
		sb.WriteString("\n```\n\n")
	} else {
		sb.WriteString("- **Status**: ❌ **FAILED**\n")
		sb.WriteString("- **Error**: No response received\n\n")
		// Add detailed error information if available
		if len(report.TestResults) > 2 {
			result := report.TestResults[2] // Third test is list MCIs
			if result.ErrorMessage != "" {
				sb.WriteString("**Error Message**:\n\n```\n")
				sb.WriteString(result.ErrorMessage)
				sb.WriteString("\n```\n\n")
			}
			if result.ErrorDetails != "" {
				sb.WriteString(fmt.Sprintf("**Error Details**: %s\n\n", result.ErrorDetails))
			}
		}
	}

	// Test Case 4: List MCI IDs
	sb.WriteString("### Test Case 4: Get a list of MCI IDs\n\n")

	// API request information
	sb.WriteString("#### 4.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `GET /beetle/migration/ns/%s/mci?option=id`\n", report.NamespaceID))
	sb.WriteString("- **Purpose**: Retrieve MCI IDs only (lightweight response)\n")
	sb.WriteString(fmt.Sprintf("- **Namespace ID**: `%s`\n", report.NamespaceID))
	sb.WriteString("- **Query Parameter**: `option=id`\n")
	sb.WriteString("- **Request Body**: None (GET request)\n\n")

	// API response information
	sb.WriteString("#### 4.2 API Response Information\n\n")
	if report.ListMCIIDsResponse != nil {
		sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
		sb.WriteString("- **Response**: MCI IDs retrieved successfully\n\n")
		sb.WriteString("**Response Body**:\n\n")
		sb.WriteString("```json\n")
		idsJSON, _ := json.MarshalIndent(report.ListMCIIDsResponse, "", "  ")
		sb.WriteString(string(idsJSON))
		sb.WriteString("\n```\n\n")
	} else {
		sb.WriteString("- **Status**: ❌ **FAILED**\n")
		sb.WriteString("- **Error**: No response received\n\n")
		// Add detailed error information if available
		if len(report.TestResults) > 3 {
			result := report.TestResults[3] // Fourth test is list MCI IDs
			if result.ErrorMessage != "" {
				sb.WriteString("**Error Message**:\n\n```\n")
				sb.WriteString(result.ErrorMessage)
				sb.WriteString("\n```\n\n")
			}
			if result.ErrorDetails != "" {
				sb.WriteString(fmt.Sprintf("**Error Details**: %s\n\n", result.ErrorDetails))
			}
		}
	}

	// Test Case 5: Get specific MCI (if available)
	if report.GetMCIResponse != nil {
		sb.WriteString("### Test Case 5: Get a specific MCI\n\n")

		// API request information
		sb.WriteString("#### 5.1 API Request Information\n\n")
		sb.WriteString(fmt.Sprintf("- **API Endpoint**: `GET /beetle/migration/ns/%s/mci/{{mciId}}`\n", report.NamespaceID))
		sb.WriteString("- **Purpose**: Retrieve detailed information for a specific MCI\n")
		sb.WriteString(fmt.Sprintf("- **Namespace ID**: `%s`\n", report.NamespaceID))
		sb.WriteString("- **Path Parameter**: `{{mciId}}` - The specific MCI identifier\n")
		sb.WriteString("- **Request Body**: None (GET request)\n\n")

		// API response information
		sb.WriteString("#### 5.2 API Response Information\n\n")
		sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
		sb.WriteString("- **Response**: MCI details retrieved successfully\n\n")
		sb.WriteString("**Response Body**:\n\n")
		sb.WriteString("<details>\n")
		sb.WriteString("  <summary> <ins>Click to see the response body </ins> </summary>\n\n")
		sb.WriteString("```json\n")
		getMCIJSON, _ := json.MarshalIndent(report.GetMCIResponse, "", "  ")
		sb.WriteString(string(getMCIJSON))
		sb.WriteString("\n```\n\n")
		sb.WriteString("</details>\n\n")
	}

	// Test Case 6: Remote Command Accessibility Check
	if len(report.TestResults) > 5 && report.TestResults[5].TestName != "" {
		sb.WriteString("### Test Case 6: Remote Command Accessibility Check\n\n")

		// API request information
		sb.WriteString("#### 6.1 Test Information\n\n")
		sb.WriteString("- **Test Type**: SSH Connectivity Test for All VMs\n")
		sb.WriteString("- **Purpose**: Verify that all migrated VMs are accessible via SSH\n")
		sb.WriteString("- **Method**: Extract public IP and SSH key from MCI access info for each VM, then execute remote command\n")
		sb.WriteString("- **Command Executed**: `uname -a` (to verify system information)\n")
		sb.WriteString("- **Authentication**: SSH key-based authentication\n")
		sb.WriteString("- **Scope**: Tests all VMs across all subgroups in the MCI\n\n")

		// Test result information
		sb.WriteString("#### 6.2 Test Result Information\n\n")
		remoteResult := report.TestResults[5] // 6th test is remote command
		if remoteResult.Success {
			// Success case - show detailed results for all VMs
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString("- **Result**: All VMs are accessible via SSH\n\n")

			if response, ok := remoteResult.Response["overallStatus"].(map[string]interface{}); ok {
				if message, exists := response["message"].(string); exists {
					sb.WriteString(fmt.Sprintf("**Summary**: %s\n\n", message))
				}
			}

			if totalVMs, ok := remoteResult.Response["totalVMs"].(int); ok {
				if successfulTests, ok := remoteResult.Response["successfulTests"].(int); ok {
					if failedTests, ok := remoteResult.Response["failedTests"].(int); ok {
						sb.WriteString("**Test Statistics**:\n\n")
						sb.WriteString(fmt.Sprintf("- Total VMs: %d\n", totalVMs))
						sb.WriteString(fmt.Sprintf("- Successful Tests: %d\n", successfulTests))
						sb.WriteString(fmt.Sprintf("- Failed Tests: %d\n\n", failedTests))
					}
				}
			}

			// Show VM-specific results
			if vmResults, ok := remoteResult.Response["vmResults"].([]interface{}); ok && len(vmResults) > 0 {
				sb.WriteString("**Per-VM Test Results**:\n\n")
				for _, vmResult := range vmResults {
					if vmMap, ok := vmResult.(map[string]interface{}); ok {
						if vmId, ok := vmMap["vmId"].(string); ok {
							sb.WriteString(fmt.Sprintf("- **VM ID**: `%s`", vmId))
							if publicIP, ok := vmMap["publicIP"].(string); ok {
								sb.WriteString(fmt.Sprintf(" (IP: `%s`)", publicIP))
							}
							if subGroup, ok := vmMap["subGroup"].(string); ok {
								sb.WriteString(fmt.Sprintf(" - SubGroup: `%s`", subGroup))
							}
							if status, ok := vmMap["status"].(string); ok {
								if status == "success" {
									sb.WriteString(" - ✅ **SUCCESS**")
									if command, ok := vmMap["command"].(string); ok && command != "" {
										sb.WriteString(fmt.Sprintf(" (Command: `%s`)", command))
									}
									if output, ok := vmMap["output"].(string); ok && output != "" {
										sb.WriteString(fmt.Sprintf(" → Output: `%s`", strings.TrimSpace(output)))
									}
								} else {
									sb.WriteString(" - ❌ **FAILED**")
									if errorMsg, ok := vmMap["error"].(string); ok {
										sb.WriteString(fmt.Sprintf(" (Error: %s)", errorMsg))
									}
								}
							}
							sb.WriteString("\n")
						}
					}
				}
				sb.WriteString("\n")
			}

			sb.WriteString("**Complete Test Details**:\n\n")
			sb.WriteString("<details>\n")
			sb.WriteString("  <summary> <ins>Click to see detailed test results </ins> </summary>\n\n")
			sb.WriteString("```json\n")
			remoteJSON, _ := json.MarshalIndent(remoteResult.Response, "", "  ")
			sb.WriteString(string(remoteJSON))
			sb.WriteString("\n```\n\n")
			sb.WriteString("</details>\n\n")
		} else if remoteResult.Skipped {
			// Skipped case
			sb.WriteString("- **Status**: ⏭️ **SKIPPED**\n")
			sb.WriteString(fmt.Sprintf("- **Reason**: %s\n\n", remoteResult.Error))
		} else {
			// Failure case - check if it's partial failure or complete failure
			if totalVMs, ok := remoteResult.Response["totalVMs"].(int); ok {
				if successfulTests, ok := remoteResult.Response["successfulTests"].(int); ok {
					if failedTests, ok := remoteResult.Response["failedTests"].(int); ok {
						if successfulTests > 0 {
							// Partial failure
							sb.WriteString("- **Status**: ⚠️ **PARTIAL SUCCESS**\n")
							sb.WriteString("- **Result**: Some VMs failed SSH connectivity test\n\n")
							sb.WriteString("**Test Statistics**:\n\n")
							sb.WriteString(fmt.Sprintf("- Total VMs: %d\n", totalVMs))
							sb.WriteString(fmt.Sprintf("- Successful Tests: %d\n", successfulTests))
							sb.WriteString(fmt.Sprintf("- Failed Tests: %d\n\n", failedTests))
						} else {
							// Complete failure
							sb.WriteString("- **Status**: ❌ **FAILED**\n")
							sb.WriteString("- **Error**: All SSH connectivity tests failed\n\n")
						}
					}
				}
			} else {
				// Complete failure
				sb.WriteString("- **Status**: ❌ **FAILED**\n")
				sb.WriteString("- **Error**: SSH connectivity test failed\n\n")
			}

			if remoteResult.ErrorMessage != "" {
				sb.WriteString("**Error Message**:\n\n```\n")
				sb.WriteString(remoteResult.ErrorMessage)
				sb.WriteString("\n```\n\n")
			} else if remoteResult.Error != "" {
				sb.WriteString("**Error**:\n\n```\n")
				sb.WriteString(remoteResult.Error)
				sb.WriteString("\n```\n\n")
			}

			if remoteResult.ErrorDetails != "" {
				sb.WriteString(fmt.Sprintf("**Error Details**: %s\n\n", remoteResult.ErrorDetails))
			}

			// Show failed VM details if available
			if vmResults, ok := remoteResult.Response["vmResults"].([]map[string]interface{}); ok && len(vmResults) > 0 {
				sb.WriteString("**Detailed VM Analysis**:\n\n")
				for _, vmResult := range vmResults {
					if vmId, ok := vmResult["vmId"].(string); ok {
						sb.WriteString(fmt.Sprintf("- **VM ID**: `%s`", vmId))
						if publicIP, ok := vmResult["publicIP"].(string); ok {
							sb.WriteString(fmt.Sprintf(" (IP: `%s`)", publicIP))
						}
						if subGroup, ok := vmResult["subGroup"].(string); ok {
							sb.WriteString(fmt.Sprintf(" - SubGroup: `%s`", subGroup))
						}

						if status, ok := vmResult["status"].(string); ok {
							if status == "success" {
								sb.WriteString(" - ✅ **SUCCESS**")
								if output, ok := vmResult["output"].(string); ok && output != "" {
									sb.WriteString(fmt.Sprintf(" (Output: `%s`)", strings.TrimSpace(output)))
								}
							} else {
								sb.WriteString(" - ❌ **FAILED**")

								// Show network test result if available
								if networkTest, ok := vmResult["networkTest"].(map[string]interface{}); ok {
									if reachable, ok := networkTest["reachable"].(bool); ok {
										if reachable {
											if latency, ok := networkTest["latency"].(string); ok {
												sb.WriteString(fmt.Sprintf(" | Network: ✅ (Latency: %s)", latency))
											} else {
												sb.WriteString(" | Network: ✅")
											}
										} else {
											if netErr, ok := networkTest["error"].(string); ok {
												sb.WriteString(fmt.Sprintf(" | Network: ❌ (%s)", netErr))
											} else {
												sb.WriteString(" | Network: ❌")
											}
										}
									}
								}

								// Show diagnosis if available
								if diagnosis, ok := vmResult["diagnosis"].(string); ok {
									sb.WriteString(fmt.Sprintf("\n  - **Diagnosis**: %s", diagnosis))
								}

								// Show error details
								if errorMsg, ok := vmResult["error"].(string); ok {
									sb.WriteString(fmt.Sprintf("\n  - **Error**: %s", errorMsg))
								}
							}
						}
						sb.WriteString("\n")
					}
				}
				sb.WriteString("\n")

				// Add troubleshooting guide for failed VMs
				hasFailures := false
				for _, vmResult := range vmResults {
					if status, ok := vmResult["status"].(string); ok && status == "failed" {
						hasFailures = true
						break
					}
				}

				if hasFailures {
					sb.WriteString("**Troubleshooting Guide**:\n\n")
					sb.WriteString("1. **Network Issues**: If network test fails, check security group rules for port 22\n")
					sb.WriteString("2. **SSH Service**: If network test passes but SSH fails, check if SSH service is running on VM\n")
					sb.WriteString("3. **VM Status**: Verify that VM is fully started and not in a transitional state\n")
					sb.WriteString("4. **Authentication**: Ensure SSH key matches the key used during VM creation\n")
					sb.WriteString("5. **User Account**: Try different usernames (ubuntu, ec2-user, admin, root)\n\n")
				}
			}
		}
	}

	// Test Case 7: Delete MCI (always show if test was attempted)
	if len(report.TestResults) > 6 && report.TestResults[6].TestName != "" {
		sb.WriteString("### Test Case 7: Delete the migrated computing infra\n\n")

		// API request information
		sb.WriteString("#### 7.1 API Request Information\n\n")
		sb.WriteString(fmt.Sprintf("- **API Endpoint**: `DELETE /beetle/migration/ns/%s/mci/{{mciId}}`\n", report.NamespaceID))
		sb.WriteString("- **Purpose**: Delete the migrated infrastructure and clean up resources\n")
		sb.WriteString(fmt.Sprintf("- **Namespace ID**: `%s`\n", report.NamespaceID))
		sb.WriteString("- **Path Parameter**: `{{mciId}}` - The MCI identifier to delete\n")
		sb.WriteString("- **Query Parameter**: `option=terminate` (terminates all resources)\n")
		sb.WriteString("- **Request Body**: None (DELETE request)\n\n")

		// API response information
		sb.WriteString("#### 7.2 API Response Information\n\n")
		deleteResult := report.TestResults[6] // 7th test is delete
		if deleteResult.Success && len(report.DeleteMCIResponse) > 0 {
			// Success case - show response
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString("- **Response**: Infrastructure deletion completed successfully\n\n")
			sb.WriteString("**Response Body**:\n\n")
			sb.WriteString("```json\n")
			delJSON, _ := json.MarshalIndent(report.DeleteMCIResponse, "", "  ")
			sb.WriteString(string(delJSON))
			sb.WriteString("\n```\n\n")
		} else if deleteResult.Skipped {
			// Skipped case
			sb.WriteString("- **Status**: ⏭️ **SKIPPED**\n")
			sb.WriteString(fmt.Sprintf("- **Reason**: %s\n\n", deleteResult.Error))
		} else if !deleteResult.Success {
			// Failure case - show error message
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			sb.WriteString("- **Error**: Infrastructure deletion failed\n\n")

			if deleteResult.ErrorMessage != "" {
				sb.WriteString("**Error Message**:\n\n```\n")
				sb.WriteString(deleteResult.ErrorMessage)
				sb.WriteString("\n```\n\n")
			} else if deleteResult.Error != "" {
				sb.WriteString("**Error**:\n\n```\n")
				sb.WriteString(deleteResult.Error)
				sb.WriteString("\n```\n\n")
			}

			if deleteResult.ErrorDetails != "" {
				sb.WriteString(fmt.Sprintf("**Error Details**: %s\n\n", deleteResult.ErrorDetails))
			}

			if deleteResult.RequestURL != "" {
				sb.WriteString(fmt.Sprintf("**Request URL**: `%s`\n\n", deleteResult.RequestURL))
			}

			// Show error response if available
			if len(deleteResult.Response) > 0 {
				sb.WriteString("**Response Body**:\n\n")
				sb.WriteString("```json\n")
				errJSON, _ := json.MarshalIndent(deleteResult.Response, "", "  ")
				sb.WriteString(string(errJSON))
				sb.WriteString("\n```\n\n")
			}
		}
	}

	return sb.String()
}

// runSourceSummaryTest performs Source Summary: POST /beetle/summary/source
func runSourceSummaryTest(client *resty.Client, config TestConfig, onpremInfraModel onpremmodel.OnpremInfra) TestResults {
	result := TestResults{
		TestName:  "Source Summary: POST /beetle/summary/source",
		StartTime: time.Now(),
	}

	// Wait before API call for stability with animation
	animatedSleep(2*time.Second, "Preparing source infrastructure summary...")

	// Log API call details
	log.Info().Msg("\n--- Source Infrastructure Summary ---")

	// Prepare request body
	requestBody := map[string]interface{}{
		"onpremiseInfraModel": onpremInfraModel,
	}

	// Log request
	log.Debug().Interface("request", requestBody).Msg("Source summary request body")

	// Make API call with markdown format
	url := fmt.Sprintf("%s/beetle/summary/source?format=md", config.BeetleURL)
	result.RequestURL = url
	result.RequestBody = requestBody

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = resp.StatusCode()

	if err != nil {
		errorMsg, errorDetails := extractErrorDetails(err, resp.StatusCode())
		populateErrorInfo(&result, err, resp.StatusCode(), url, requestBody)
		log.Error().
			Err(err).
			Int("statusCode", resp.StatusCode()).
			Str("errorMessage", errorMsg).
			Str("errorDetails", errorDetails).
			Msg("Source summary failed")
		result.Success = false
		return result
	}

	if resp.StatusCode() != http.StatusOK {
		err := fmt.Errorf("unexpected status code: %d", resp.StatusCode())
		errorMsg, errorDetails := extractErrorDetails(err, resp.StatusCode())
		populateErrorInfo(&result, err, resp.StatusCode(), url, requestBody)
		log.Error().
			Int("statusCode", resp.StatusCode()).
			Str("errorMessage", errorMsg).
			Str("errorDetails", errorDetails).
			Msg("Source summary failed")
		result.Success = false
		return result
	}

	// Get markdown content
	markdownContent := string(resp.Body())

	// Log the markdown summary (it's already formatted)
	log.Info().Msg("\n" + markdownContent)

	// Save markdown to file
	cwd, err := os.Getwd()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get current working directory")
		cwd = "."
	}

	testResultDir := filepath.Join(cwd, "testresult")
	if err := os.MkdirAll(testResultDir, 0755); err != nil {
		log.Error().Err(err).Str("path", testResultDir).Msg("Failed to create testresult directory")
		result.Success = false
		result.Error = fmt.Sprintf("Failed to create directory: %v", err)
		return result
	}

	filename := "beetle-summary-source.md"
	filepath := filepath.Join(testResultDir, filename)

	if err := os.WriteFile(filepath, []byte(markdownContent), 0644); err != nil {
		log.Error().Err(err).Str("file", filepath).Msg("Failed to write source summary file")
		result.Success = false
		result.Error = fmt.Sprintf("Failed to write file: %v", err)
		return result
	}

	log.Info().Str("file", filepath).Msg("✅ Source infrastructure summary saved to file")

	result.Success = true
	result.Response = map[string]interface{}{
		"markdown": markdownContent,
		"file":     filepath,
	}

	return result
}

// runTargetSummaryTest performs Target Summary: GET /beetle/summary/target/ns/{nsId}/mci/{mciId}
func runTargetSummaryTest(client *resty.Client, config TestConfig, mciId, cspName, displayName string) TestResults {
	result := TestResults{
		TestName:  "Target Summary: GET /beetle/summary/target/ns/{nsId}/mci/{mciId}",
		StartTime: time.Now(),
	}

	// Wait before API call for stability with animation
	animatedSleep(2*time.Second, "Preparing target infrastructure summary...")

	// Log API call details
	log.Info().Msgf("\n--- Target Infrastructure Summary for %s ---", displayName)

	// Make API call with markdown format
	url := fmt.Sprintf("%s/beetle/summary/target/ns/%s/mci/%s?format=md",
		config.BeetleURL, config.NamespaceID, mciId)
	result.RequestURL = url

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = resp.StatusCode()

	if err != nil {
		errorMsg, errorDetails := extractErrorDetails(err, resp.StatusCode())
		populateErrorInfo(&result, err, resp.StatusCode(), url, nil)
		log.Error().
			Err(err).
			Int("statusCode", resp.StatusCode()).
			Str("errorMessage", errorMsg).
			Str("errorDetails", errorDetails).
			Msg("Target summary failed")
		result.Success = false
		return result
	}

	if resp.StatusCode() != http.StatusOK {
		err := fmt.Errorf("unexpected status code: %d", resp.StatusCode())
		errorMsg, errorDetails := extractErrorDetails(err, resp.StatusCode())
		populateErrorInfo(&result, err, resp.StatusCode(), url, nil)
		log.Error().
			Int("statusCode", resp.StatusCode()).
			Str("errorMessage", errorMsg).
			Str("errorDetails", errorDetails).
			Msg("Target summary failed")
		result.Success = false
		return result
	}

	// Get markdown content
	markdownContent := string(resp.Body())

	// Save markdown to file
	cwd, err := os.Getwd()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get current working directory")
		cwd = "."
	}

	testResultDir := filepath.Join(cwd, "testresult")
	if err := os.MkdirAll(testResultDir, 0755); err != nil {
		log.Error().Err(err).Str("path", testResultDir).Msg("Failed to create testresult directory")
		result.Success = false
		result.Error = fmt.Sprintf("Failed to create directory: %v", err)
		return result
	}

	filename := fmt.Sprintf("beetle-summary-target-%s.md", cspName)
	filepath := filepath.Join(testResultDir, filename)

	if err := os.WriteFile(filepath, []byte(markdownContent), 0644); err != nil {
		log.Error().Err(err).Str("file", filepath).Msg("Failed to write target summary file")
		result.Success = false
		result.Error = fmt.Sprintf("Failed to write file: %v", err)
		return result
	}

	log.Info().Str("file", filepath).Msg("✅ Target infrastructure summary saved to file")

	result.Success = true
	result.Response = map[string]interface{}{
		"markdown": markdownContent,
		"file":     filepath,
	}

	return result
}

// runMigrationReportTest performs Migration Report: POST /beetle/report/migration/ns/{nsId}/mci/{mciId}
func runMigrationReportTest(client *resty.Client, config TestConfig, onpremInfraModel onpremmodel.OnpremInfra, mciId, cspName, displayName string) TestResults {
	result := TestResults{
		TestName:  "Migration Report: POST /beetle/report/migration/ns/{nsId}/mci/{mciId}",
		StartTime: time.Now(),
	}

	// Wait before API call for stability with animation
	animatedSleep(2*time.Second, "Preparing migration report...")

	// Log API call details
	log.Info().Msgf("\n--- Migration Report for %s ---", displayName)

	// Prepare request body
	requestBody := map[string]interface{}{
		"onpremiseInfraModel": onpremInfraModel,
	}

	// Make API call
	url := fmt.Sprintf("%s/beetle/report/migration/ns/%s/mci/%s",
		config.BeetleURL, config.NamespaceID, mciId)
	result.RequestURL = url

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		Post(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = resp.StatusCode()

	if err != nil {
		errorMsg, errorDetails := extractErrorDetails(err, resp.StatusCode())
		populateErrorInfo(&result, err, resp.StatusCode(), url, nil)
		log.Error().
			Err(err).
			Int("statusCode", resp.StatusCode()).
			Str("errorMessage", errorMsg).
			Str("errorDetails", errorDetails).
			Msg("Migration report failed")
		result.Success = false
		return result
	}

	if resp.StatusCode() != http.StatusOK {
		err := fmt.Errorf("unexpected status code: %d", resp.StatusCode())
		errorMsg, errorDetails := extractErrorDetails(err, resp.StatusCode())
		populateErrorInfo(&result, err, resp.StatusCode(), url, nil)
		log.Error().
			Int("statusCode", resp.StatusCode()).
			Str("errorMessage", errorMsg).
			Str("errorDetails", errorDetails).
			Msg("Migration report failed")
		result.Success = false
		return result
	}

	// Get markdown content
	markdownContent := string(resp.Body())

	// Save markdown to file
	cwd, err := os.Getwd()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get current working directory")
		cwd = "."
	}

	testResultDir := filepath.Join(cwd, "testresult")
	if err := os.MkdirAll(testResultDir, 0755); err != nil {
		log.Error().Err(err).Str("path", testResultDir).Msg("Failed to create testresult directory")
		result.Success = false
		result.Error = fmt.Sprintf("Failed to create directory: %v", err)
		return result
	}

	filename := fmt.Sprintf("beetle-report-mig-source-to-%s.md", cspName)
	filepath := filepath.Join(testResultDir, filename)

	if err := os.WriteFile(filepath, []byte(markdownContent), 0644); err != nil {
		log.Error().Err(err).Str("file", filepath).Msg("Failed to write migration report file")
		result.Success = false
		result.Error = fmt.Sprintf("Failed to write file: %v", err)
		return result
	}

	log.Info().Str("file", filepath).Msg("✅ Migration report saved to file")

	result.Success = true
	result.Response = map[string]interface{}{
		"markdown": markdownContent,
		"file":     filepath,
	}

	return result
}

// runRemoteCommandTest performs Test 6: Remote Command to check accessibility of migrated VM
func runRemoteCommandTest(client *resty.Client, config TestConfig, mciId, displayName string) TestResults {
	log.Info().Msg("\n--- Test 6: Remote Command Accessibility Check ---")

	// Wait before test for stability with animation
	animatedSleep(5*time.Second, "Waiting before VM accessibility test")

	result := TestResults{
		TestName:   fmt.Sprintf("Test 6: Remote Command Accessibility Check (%s)", displayName),
		StartTime:  time.Now(),
		RequestURL: "N/A (SSH command)",
	}

	defer func() {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
	}()

	// Step 1: Get MCI Access Info
	log.Info().Msg("Step 1: Getting MCI access information...")

	tbClient := tbclient.NewDefaultClient()
	accessInfo, err := tbClient.ReadMciAccessInfo(config.NamespaceID, mciId, "accessinfo", "showSshKey")
	if err != nil {
		populateErrorInfo(&result, err, 0, "ReadMciAccessInfo", nil)
		log.Error().Err(err).Msg("Failed to get MCI access info")
		return result
	}

	log.Debug().Msgf("Access info retrieved successfully")

	// Step 2: Extract all VM information from all subgroups
	if len(accessInfo.MciSubGroupAccessInfo) == 0 {
		err := fmt.Errorf("no subgroups found in MCI access info")
		populateErrorInfo(&result, err, 0, "Extract VM Info", nil)
		log.Error().Err(err).Msg("No VM subgroups found")
		return result
	}

	// Collect all VMs from all subgroups
	var allVMs []struct {
		SubGroupId string
		VmInfo     interface{}
	}

	totalVMs := 0
	for _, subGroup := range accessInfo.MciSubGroupAccessInfo {
		for _, vmInfo := range subGroup.MciVmAccessInfo {
			allVMs = append(allVMs, struct {
				SubGroupId string
				VmInfo     interface{}
			}{
				SubGroupId: subGroup.SubGroupId,
				VmInfo:     vmInfo,
			})
			totalVMs++
		}
	}

	if totalVMs == 0 {
		err := fmt.Errorf("no VMs found in any subgroup")
		populateErrorInfo(&result, err, 0, "Extract VM Info", nil)
		log.Error().Err(err).Msg("No VMs found in any subgroup")
		return result
	}

	log.Info().Msgf("Step 2: Found %d VMs across %d subgroups for testing", totalVMs, len(accessInfo.MciSubGroupAccessInfo))

	// Step 3: Perform SSH connectivity test for all VMs
	log.Info().Msg("Step 3: Testing SSH connectivity for all VMs...")

	vmTestResults := make([]map[string]interface{}, 0)
	successfulTests := 0
	failedTests := 0

	for i, vm := range allVMs {
		// Type assertion to get VM info
		vmInfo, ok := vm.VmInfo.(tbmodel.MciVmAccessInfo)
		if !ok {
			log.Error().Msgf("Failed to cast VM info for VM %d", i+1)
			failedTests++
			continue
		}

		publicIP := vmInfo.PublicIP
		privateKey := vmInfo.PrivateKey
		vmId := vmInfo.VmId
		userName := vmInfo.VmUserName

		if userName == "" {
			log.Warn().Str("vmId", vmId).Msg("No username provided in access info, defaulting to 'cb-user'")
			userName = "cb-user" // Default user in this platform
		}

		log.Info().Msgf("Testing VM %d/%d: %s (IP: %s, SubGroup: %s)", i+1, totalVMs, vmId, publicIP, vm.SubGroupId)

		vmResult := map[string]interface{}{
			"vmId":      vmId,
			"publicIP":  publicIP,
			"subGroup":  vm.SubGroupId,
			"userName":  userName,
			"testOrder": i + 1,
		}

		if publicIP == "" {
			vmResult["status"] = "failed"
			vmResult["error"] = "no public IP available"
			vmResult["sshTest"] = "skipped"
			failedTests++
			log.Warn().Str("vmId", vmId).Msg("⚠️ Skipping SSH test - no public IP available")
		} else if privateKey == "" {
			vmResult["status"] = "failed"
			vmResult["error"] = "no private key available"
			vmResult["sshTest"] = "skipped"
			failedTests++
			log.Warn().Str("vmId", vmId).Msg("⚠️ Skipping SSH test - no private key available")
		} else {
			// Debug: Log private key info (safely)
			keyPreview := privateKey
			if len(keyPreview) > 100 {
				keyPreview = keyPreview[:50] + "..." + keyPreview[len(keyPreview)-50:]
			}
			log.Debug().Str("vmId", vmId).Str("keyPreview", keyPreview).Bool("hasLiteralNewlines", strings.Contains(privateKey, "\\n")).Msg("Private key info")

			// Determine username (priority: vmInfo.VmUserName > default "cb-user")
			sshUserName := userName

			// Perform SSH connectivity test with retry logic
			log.Info().Str("vmId", vmId).Str("ip", publicIP).Str("user", sshUserName).Msg("🔍 Testing SSH connectivity...")

			const maxRetries = 3
			const retryDelay = 3 * time.Second
			var sshResult string
			var lastErr error

			for attempt := 1; attempt <= maxRetries; attempt++ {
				if attempt > 1 {
					log.Info().Str("vmId", vmId).Int("attempt", attempt).Int("maxRetries", maxRetries).Msg("🔄 Retrying SSH connection...")
					time.Sleep(retryDelay)
				}

				sshResult, lastErr = testSSHConnectivity(publicIP, privateKey, sshUserName)
				if lastErr == nil {
					// Success
					vmResult["status"] = "success"
					vmResult["sshTest"] = "successful"
					vmResult["command"] = "uname -a"
					vmResult["output"] = sshResult
					vmResult["attempts"] = attempt
					successfulTests++
					if attempt > 1 {
						log.Info().Str("vmId", vmId).Str("ip", publicIP).Int("attempt", attempt).Msg("✅ SSH connectivity test passed after retry")
					} else {
						log.Info().Str("vmId", vmId).Str("ip", publicIP).Msg("✅ SSH connectivity test passed")
					}
					break
				} else {
					log.Warn().Err(lastErr).Str("vmId", vmId).Str("ip", publicIP).Int("attempt", attempt).Int("maxRetries", maxRetries).Msg("⚠️ SSH connection attempt failed")
				}
			}

			if lastErr != nil {
				vmResult["status"] = "failed"
				vmResult["error"] = lastErr.Error()
				vmResult["sshTest"] = "failed"
				vmResult["attempts"] = maxRetries
				failedTests++
				log.Error().Err(lastErr).Str("vmId", vmId).Str("ip", publicIP).Int("maxRetries", maxRetries).Msg("❌ SSH connectivity test failed after all retries")
			}
		}

		vmTestResults = append(vmTestResults, vmResult)
	}

	// Determine overall test result
	overallSuccess := failedTests == 0

	// Success or partial success
	result.Success = overallSuccess
	result.StatusCode = 200
	result.Response = map[string]interface{}{
		"totalVMs":        totalVMs,
		"successfulTests": successfulTests,
		"failedTests":     failedTests,
		"overallStatus": map[string]interface{}{
			"success": overallSuccess,
			"message": fmt.Sprintf("%d/%d VMs accessible via SSH", successfulTests, totalVMs),
		},
		"vmResults": vmTestResults,
	}

	if overallSuccess {
		log.Info().Int("successful", successfulTests).Int("total", totalVMs).Msg("✅ All VMs passed SSH connectivity test")
		fmt.Println("✅ Test 6 passed")
	} else {
		log.Warn().Int("successful", successfulTests).Int("failed", failedTests).Int("total", totalVMs).Msg("⚠️ Some VMs failed SSH connectivity test")
		fmt.Printf("❌ Test 6 failed: %d/%d VMs failed SSH connectivity\n", failedTests, totalVMs)
	}

	return result
}

// testSSHConnectivity tests SSH connection to a VM and runs a simple command
func testSSHConnectivity(host, privateKey, username string) (string, error) {
	// Clean and normalize private key format
	// Replace literal \n with actual newlines
	cleanedKey := strings.ReplaceAll(privateKey, "\\n", "\n")

	// Ensure proper PEM format
	if !strings.HasPrefix(cleanedKey, "-----BEGIN") {
		return "", fmt.Errorf("invalid private key format: missing PEM header")
	}

	// Parse private key
	signer, err := ssh.ParsePrivateKey([]byte(cleanedKey))
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %v", err)
	}

	// SSH client config with reasonable timeout
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Note: For testing only
		Timeout:         45 * time.Second,            // Reasonable timeout
	}

	// Connect to SSH
	address := net.JoinHostPort(host, "22")
	conn, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return "", fmt.Errorf("failed to connect via SSH: %v", err)
	}
	defer conn.Close()

	// Create session
	session, err := conn.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create SSH session: %v", err)
	}
	defer session.Close()

	// Run command
	output, err := session.CombinedOutput("uname -a")
	if err != nil {
		return "", fmt.Errorf("failed to run command: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
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
