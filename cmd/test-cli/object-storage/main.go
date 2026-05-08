// Package main is the starting point of CM-Beetle Object Storage Migration Test CLI
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
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
)

// ============================================================================
// Configuration & Type Definitions
// ============================================================================

// TestConfig holds test configuration loaded from YAML.
type TestConfig struct {
	Test struct {
		Set struct {
			Mode string `yaml:"mode"` // parallel or sequential
		} `yaml:"set"`
		Cases []TestCase `yaml:"cases"`
	} `yaml:"test"`
	Beetle struct {
		Endpoint        string `yaml:"endpoint"`
		NamespaceID     string `yaml:"namespaceId"`
		RequestBodyFile string `yaml:"requestBodyFile"`
		AuthConfigFile  string `yaml:"authConfigFile"`
	} `yaml:"beetle"`
}

// TestCase holds a single CSP-Region pair to test.
type TestCase struct {
	Csp     string `yaml:"csp"`
	Region  string `yaml:"region"`
	Name    string `yaml:"name"`
	Execute bool   `yaml:"execute"`
}

// AuthConfig holds basic auth credentials.
type AuthConfig struct {
	BasicAuthUsername string `json:"basicAuthUsername"`
	BasicAuthPassword string `json:"basicAuthPassword"`
}

// OsRequest holds the source object storage information loaded from JSON.
type OsRequest struct {
	NameSeed             string                             `json:"nameSeed"`
	SourceObjectStorages []storagemodel.SourceObjectStorage `json:"sourceObjectStorages"`
}

// TestResults holds the result of a single API test.
type TestResults struct {
	TestName   string        `json:"testName"`
	StartTime  time.Time     `json:"startTime"`
	EndTime    time.Time     `json:"endTime"`
	Duration   time.Duration `json:"duration"`
	Success    bool          `json:"success"`
	Skipped    bool          `json:"skipped"`
	StatusCode int           `json:"statusCode"`
	Response   interface{}   `json:"response,omitempty"`
	Error      string        `json:"error,omitempty"`
	RequestURL string        `json:"requestUrl,omitempty"`
}

// TestSuite aggregates results across all CSP-Region pairs.
// OSTestReport holds all results for a single CSP-Region pair report.
type OSTestReport struct {
	CSP                    string
	Region                 string
	DisplayName            string
	TestDate               string
	TestTime               string
	TestDateTime           time.Time
	BeetleURL              string
	NamespaceID            string
	NameSeed               string
	OsRequest              OsRequest
	RecommendationResponse interface{}
	MigrationResponse      interface{}
	ListResponse           interface{}
	GetResponse            interface{}
	TestResults            []TestResults
	Summary                TestResults
}

type TestSuite struct {
	TotalTests      int
	TotalCspPairs   int
	PassedTests     int
	FailedTests     int
	SkippedTests    int
	PassedCspPairs  int
	FailedCspPairs  int
	SkippedCspPairs int
	OverallTime     time.Duration
	CspResults      map[string]bool
	mu              sync.Mutex
}

// objectStorageInfo is a local type for parsing TB list/get response fields.
type objectStorageInfo struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	ConnectionName string `json:"connectionName"`
}

// objectStorageListData is a local type for parsing the list response body.
type objectStorageListData struct {
	ObjectStorage []objectStorageInfo `json:"objectStorage"`
}

// ============================================================================
// CLI Flags & Initialization
// ============================================================================

var (
	configFile = flag.String("config", "testconf/test-config.yaml", "Path to YAML config file")
)

func init() {
	config.Init()
	l := logger.NewLogger(logger.Config{
		LogLevel:    config.Beetle.LogLevel,
		LogWriter:   config.Beetle.LogWriter,
		LogFilePath: config.Beetle.LogFile.Path,
		MaxSize:     config.Beetle.LogFile.MaxSize,
		MaxBackups:  config.Beetle.LogFile.MaxBackups,
		MaxAge:      config.Beetle.LogFile.MaxAge,
		Compress:    config.Beetle.LogFile.Compress,
	})
	log.Logger = *l
}

// ============================================================================
// Main
// ============================================================================

func main() {
	flag.Parse()

	cfg, err := loadConfig(*configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	authConfig, err := loadAuthConfig(cfg.Beetle.AuthConfigFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load auth config")
	}

	osRequest, err := loadOsRequest(cfg.Beetle.RequestBodyFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load object storage request")
	}

	baseNameSeed := osRequest.NameSeed
	log.Info().Str("nameSeed", baseNameSeed).Msg("Loaded NameSeed from request file")

	// Check CM-Beetle readiness before starting tests
	probeClient := resty.New().SetTimeout(10 * time.Second)
	if err := checkReadiness(probeClient, cfg.Beetle.Endpoint); err != nil {
		log.Fatal().Err(err).Msg("CM-Beetle readiness check failed")
	}

	suite := &TestSuite{
		TotalTests:    6,
		TotalCspPairs: len(cfg.Test.Cases),
		CspResults:    make(map[string]bool),
	}

	startTime := time.Now()
	isParallel := cfg.Test.Set.Mode == "parallel"
	var wg sync.WaitGroup

	for i, tc := range cfg.Test.Cases {
		if !tc.Execute {
			log.Info().Msgf("[%d/%d] Skipping %s (execute=false)", i+1, len(cfg.Test.Cases), displayName(tc))
			suite.mu.Lock()
			suite.SkippedCspPairs++
			suite.mu.Unlock()
			continue
		}

		// Append index suffix to nameSeed in parallel mode to avoid name collisions.
		caseNameSeed := baseNameSeed
		if isParallel && len(cfg.Test.Cases) > 1 {
			caseNameSeed = fmt.Sprintf("%s%02d", baseNameSeed, i+1)
		}

		if isParallel {
			wg.Add(1)
			go func(idx int, tc TestCase, seed string) {
				defer wg.Done()
				runTestCase(idx, tc, cfg, osRequest, seed, suite, authConfig)
			}(i, tc, caseNameSeed)
		} else {
			runTestCase(i, tc, cfg, osRequest, caseNameSeed, suite, authConfig)
		}
	}

	if isParallel {
		wg.Wait()
	}

	suite.OverallTime = time.Since(startTime)
	printFinalSummary(suite)
}

// ============================================================================
// Test Execution Per CSP-Region Pair
// ============================================================================

// runTestCase executes the full 6-step object storage test for one CSP-Region pair.
func runTestCase(idx int, tc TestCase, cfg TestConfig, osRequest OsRequest, nameSeed string, suite *TestSuite, authConfig AuthConfig) {
	name := displayName(tc)
	log.Info().Msgf("%s", "\n"+strings.Repeat("=", 60))
	log.Info().Msgf("[%d/%d] Testing %s (%s %s) — NameSeed: %q", idx+1, suite.TotalCspPairs, name, tc.Csp, tc.Region, nameSeed)
	log.Info().Msgf("%s", strings.Repeat("=", 60))

	client := resty.New()
	client.SetTimeout(10 * time.Minute)
	client.SetBaseURL(cfg.Beetle.Endpoint)
	if authConfig.BasicAuthUsername != "" {
		client.SetBasicAuth(authConfig.BasicAuthUsername, authConfig.BasicAuthPassword)
	}

	pairStartTime := time.Now()
	pairFailed := 0
	report := &OSTestReport{
		CSP:          tc.Csp,
		Region:       tc.Region,
		DisplayName:  name,
		TestDate:     pairStartTime.Format("January 2, 2006"),
		TestTime:     pairStartTime.Format("15:04:05 MST"),
		TestDateTime: pairStartTime,
		BeetleURL:    cfg.Beetle.Endpoint,
		NamespaceID:  cfg.Beetle.NamespaceID,
		NameSeed:     nameSeed,
		OsRequest:    osRequest,
		TestResults:  make([]TestResults, 0),
	}

	record := func(result TestResults) {
		suite.mu.Lock()
		defer suite.mu.Unlock()
		report.TestResults = append(report.TestResults, result)
		if result.Success {
			suite.PassedTests++
		} else if result.Skipped {
			suite.SkippedTests++
		} else {
			suite.FailedTests++
			pairFailed++
		}
	}

	stopTesting := false
	var recommendation storagemodel.RecommendedObjectStorage

	// Test 1: Recommend object storage
	rec, result1 := runRecommendTest(client, cfg, tc, osRequest, nameSeed, name)
	record(result1)
	if result1.Success {
		recommendation = rec
		report.RecommendationResponse = rec
	} else {
		stopTesting = true
	}

	// Compute expected osIds using Late Binding: ComposeName(baseName, nameSeed) = "my-os-01"
	var osIds []string
	for _, target := range recommendation.TargetObjectStorages {
		osIds = append(osIds, common.ComposeName(target.BucketName, nameSeed))
	}

	// Test 2: Migrate (create buckets)
	if !stopTesting {
		result2 := runMigrateTest(client, cfg, recommendation, name)
		record(result2)
		if result2.Success {
			report.MigrationResponse = result2.Response
		} else {
			stopTesting = true
		}
	} else {
		record(skippedResult("Test 2: Migrate", name))
	}

	// Test 3: List object storages
	if !stopTesting {
		result3 := runListTest(client, cfg, name)
		record(result3)
		if result3.Success {
			report.ListResponse = result3.Response
		} else {
			stopTesting = true
		}
	} else {
		record(skippedResult("Test 3: List", name))
	}

	// Test 4: Exist check on first bucket (HEAD)
	if !stopTesting && len(osIds) > 0 {
		result4 := runExistTest(client, cfg, osIds[0], name)
		record(result4)
	} else {
		record(skippedResult("Test 4: Exist", name))
	}

	// Test 5: Get first bucket
	if !stopTesting && len(osIds) > 0 {
		result5 := runGetTest(client, cfg, osIds[0], name)
		record(result5)
		if result5.Success {
			report.GetResponse = result5.Response
		}
	} else {
		record(skippedResult("Test 5: Get", name))
	}

	// Test 6: Delete cleanup — always attempt regardless of previous failures
	if len(osIds) > 0 {
		result6 := runDeleteAllTest(client, cfg, osIds, name)
		record(result6)
	} else {
		record(skippedResult("Test 6: Delete (cleanup)", name))
	}

	pairDuration := time.Since(pairStartTime)
	report.Summary = TestResults{
		TestName:  fmt.Sprintf("CSP-Region Pair: %s", name),
		StartTime: pairStartTime,
		EndTime:   time.Now(),
		Duration:  pairDuration,
		Success:   pairFailed == 0,
	}

	if err := generateMarkdownReport(report); err != nil {
		log.Warn().Err(err).Str("csp", name).Msg("Failed to generate markdown report")
	}

	suite.mu.Lock()
	if pairFailed == 0 {
		suite.PassedCspPairs++
		suite.CspResults[name] = true
	} else {
		suite.FailedCspPairs++
		suite.CspResults[name] = false
	}
	suite.mu.Unlock()

	log.Info().Msgf("--- %s: %d of %d tests passed (duration: %v) ---", name, 6-pairFailed, 6, pairDuration)
}

// ============================================================================
// Individual Test Functions
// ============================================================================

// runRecommendTest calls POST /beetle/recommendation/middleware/objectStorage.
func runRecommendTest(
	client *resty.Client, cfg TestConfig, tc TestCase,
	osRequest OsRequest, nameSeed, name string,
) (storagemodel.RecommendedObjectStorage, TestResults) {
	const testLabel = "Test 1: POST /recommendation/middleware/objectStorage"
	log.Info().Msgf("\n--- %s (%s) ---", testLabel, name)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	reqBody := controller.RecommendObjectStorageRequest{
		NameSeed: nameSeed,
		DesiredCloud: storagemodel.CloudProperty{
			Csp:    tc.Csp,
			Region: tc.Region,
		},
		SourceObjectStorages: osRequest.SourceObjectStorages,
	}

	url := fmt.Sprintf("%s/beetle/recommendation/middleware/objectStorage", cfg.Beetle.Endpoint)
	result.RequestURL = url

	var apiResp model.ApiResponse[storagemodel.RecommendedObjectStorage]
	resp, err := client.R().
		SetBody(reqBody).
		SetResult(&apiResp).
		Post(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = resp.StatusCode()

	if err != nil {
		result.Error = err.Error()
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return storagemodel.RecommendedObjectStorage{}, result
	}
	if resp.IsError() {
		result.Error = fmt.Sprintf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
		log.Error().Msgf("❌ %s failed: HTTP %d — %s", testLabel, resp.StatusCode(), string(resp.Body()))
		return storagemodel.RecommendedObjectStorage{}, result
	}

	result.Success = true
	result.Response = toMap(apiResp)
	log.Info().Msgf("✅ %s passed: recommended %d bucket(s) for %s %s",
		testLabel, len(apiResp.Data.TargetObjectStorages), tc.Csp, tc.Region)
	return apiResp.Data, result
}

// runMigrateTest calls POST /beetle/migration/middleware/ns/{nsId}/objectStorage.
func runMigrateTest(
	client *resty.Client, cfg TestConfig,
	recommendation storagemodel.RecommendedObjectStorage, name string,
) TestResults {
	const testLabel = "Test 2: POST /migration/middleware/ns/{nsId}/objectStorage"
	log.Info().Msgf("\n--- %s (%s) ---", testLabel, name)

	// Allow previous operations to settle.
	time.Sleep(3 * time.Second)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	reqBody := controller.MigrateObjectStorageRequest{
		RecommendedObjectStorage: recommendation,
	}

	url := fmt.Sprintf("%s/beetle/migration/middleware/ns/%s/objectStorage",
		cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID)
	result.RequestURL = url

	resp, err := client.R().
		SetBody(reqBody).
		Post(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = resp.StatusCode()

	if err != nil {
		result.Error = err.Error()
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return result
	}
	if resp.IsError() {
		result.Error = fmt.Sprintf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
		log.Error().Msgf("❌ %s failed: HTTP %d — %s", testLabel, resp.StatusCode(), string(resp.Body()))
		return result
	}

	result.Success = true
	log.Info().Msgf("✅ %s passed: %d bucket(s) created", testLabel, len(recommendation.TargetObjectStorages))
	return result
}

// runListTest calls GET /beetle/migration/middleware/ns/{nsId}/objectStorage.
func runListTest(client *resty.Client, cfg TestConfig, name string) TestResults {
	const testLabel = "Test 3: GET /migration/middleware/ns/{nsId}/objectStorage"
	log.Info().Msgf("\n--- %s (%s) ---", testLabel, name)

	time.Sleep(2 * time.Second)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	url := fmt.Sprintf("%s/beetle/migration/middleware/ns/%s/objectStorage",
		cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID)
	result.RequestURL = url

	var apiResp model.ApiResponse[objectStorageListData]
	resp, err := client.R().
		SetResult(&apiResp).
		Get(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = resp.StatusCode()

	if err != nil {
		result.Error = err.Error()
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return result
	}
	if resp.IsError() {
		result.Error = fmt.Sprintf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
		log.Error().Msgf("❌ %s failed: HTTP %d — %s", testLabel, resp.StatusCode(), string(resp.Body()))
		return result
	}

	result.Success = true
	result.Response = toMap(apiResp)
	log.Info().Msgf("✅ %s passed: %d object storage(s) listed", testLabel, len(apiResp.Data.ObjectStorage))
	return result
}

// runExistTest calls HEAD /beetle/migration/middleware/ns/{nsId}/objectStorage/{osId}.
func runExistTest(client *resty.Client, cfg TestConfig, osId, name string) TestResults {
	const testLabel = "Test 4: HEAD /migration/middleware/ns/{nsId}/objectStorage/{osId}"
	log.Info().Msgf("\n--- %s (%s, osId=%s) ---", testLabel, name, osId)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	url := fmt.Sprintf("%s/beetle/migration/middleware/ns/%s/objectStorage/%s",
		cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, osId)
	result.RequestURL = url

	resp, err := client.R().Head(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = resp.StatusCode()

	if err != nil {
		result.Error = err.Error()
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return result
	}
	if resp.IsError() {
		result.Error = fmt.Sprintf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
		log.Error().Msgf("❌ %s failed: HTTP %d — %s", testLabel, resp.StatusCode(), string(resp.Body()))
		return result
	}

	result.Success = true
	log.Info().Msgf("✅ %s passed: '%s' exists (HTTP %d)", testLabel, osId, resp.StatusCode())
	return result
}

// runGetTest calls GET /beetle/migration/middleware/ns/{nsId}/objectStorage/{osId}.
func runGetTest(client *resty.Client, cfg TestConfig, osId, name string) TestResults {
	const testLabel = "Test 5: GET /migration/middleware/ns/{nsId}/objectStorage/{osId}"
	log.Info().Msgf("\n--- %s (%s, osId=%s) ---", testLabel, name, osId)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	url := fmt.Sprintf("%s/beetle/migration/middleware/ns/%s/objectStorage/%s",
		cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, osId)
	result.RequestURL = url

	var apiResp map[string]interface{}
	resp, err := client.R().
		SetResult(&apiResp).
		Get(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = resp.StatusCode()

	if err != nil {
		result.Error = err.Error()
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return result
	}
	if resp.IsError() {
		result.Error = fmt.Sprintf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
		log.Error().Msgf("❌ %s failed: HTTP %d — %s", testLabel, resp.StatusCode(), string(resp.Body()))
		return result
	}

	result.Success = true
	result.Response = apiResp
	log.Info().Msgf("✅ %s passed: retrieved '%s'", testLabel, osId)
	return result
}

// runDeleteAllTest deletes all osIds as cleanup (Test 6).
// This always runs regardless of whether earlier tests failed.
func runDeleteAllTest(client *resty.Client, cfg TestConfig, osIds []string, name string) TestResults {
	const testLabel = "Test 6: DELETE object storages (cleanup)"
	log.Info().Msgf("\n--- %s (%s) ---", testLabel, name)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	allSucceeded := true
	for _, osId := range osIds {
		url := fmt.Sprintf("%s/beetle/migration/middleware/ns/%s/objectStorage/%s",
			cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, osId)

		const maxRetries = 3
		deleted := false
		for attempt := 1; attempt <= maxRetries; attempt++ {
			if attempt > 1 {
				log.Info().Str("osId", osId).Int("attempt", attempt).Msg("Retrying deletion...")
				time.Sleep(10 * time.Second)
			}

			resp, err := client.R().Delete(url)
			statusCode := resp.StatusCode()

			if err == nil && (statusCode < 400 || statusCode == 404) {
				log.Info().Msgf("Deleted '%s' (HTTP %d)", osId, statusCode)
				deleted = true
				break
			}

			log.Warn().Err(err).Int("statusCode", statusCode).Int("attempt", attempt).
				Msgf("Deletion attempt failed for '%s'", osId)
		}

		if !deleted {
			allSucceeded = false
			log.Error().Msgf("Failed to delete '%s' after %d attempts", osId, maxRetries)
		}
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = 204

	if allSucceeded {
		result.Success = true
		log.Info().Msgf("✅ %s passed: deleted %d bucket(s)", testLabel, len(osIds))
	} else {
		result.Error = "One or more buckets failed to delete"
		log.Error().Msgf("❌ %s: cleanup incomplete", testLabel)
	}
	return result
}

// ============================================================================
// Helper Functions
// ============================================================================

// checkReadiness verifies CM-Beetle is ready to serve requests.
func checkReadiness(client *resty.Client, endpoint string) error {
	url := fmt.Sprintf("%s/beetle/readyz", endpoint)
	var resp map[string]interface{}
	r, err := client.R().SetResult(&resp).Get(url)
	if err != nil {
		return fmt.Errorf("readiness check error: %w", err)
	}
	if r.IsError() {
		return fmt.Errorf("readiness check failed (HTTP %d)", r.StatusCode())
	}
	if msg, ok := resp["message"].(string); ok && strings.Contains(msg, "NOT ready") {
		return fmt.Errorf("CM-Beetle not ready: %s", msg)
	}
	log.Info().Msg("CM-Beetle readiness check passed")
	return nil
}

// displayName returns a human-readable name for a test case.
func displayName(tc TestCase) string {
	if tc.Name != "" {
		return tc.Name
	}
	return fmt.Sprintf("%s-%s", tc.Csp, tc.Region)
}

// skippedResult returns a pre-filled TestResults indicating the test was skipped.
func skippedResult(testLabel, name string) TestResults {
	now := time.Now()
	return TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: now,
		EndTime:   now,
		Success:   false,
		Skipped:   true,
		Error:     "skipped due to previous test failure",
	}
}

// toMap marshals any value to map[string]interface{} for TestResults.Response.
func toMap(v interface{}) map[string]interface{} {
	b, _ := json.Marshal(v)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	return m
}

// printFinalSummary prints overall test results to the log.
func printFinalSummary(suite *TestSuite) {
	log.Info().Msgf("%s", "\n"+strings.Repeat("=", 60))
	log.Info().Msg("OVERALL TEST SUMMARY")
	log.Info().Msgf("%s", strings.Repeat("=", 60))
	log.Info().Int("total", suite.TotalCspPairs).Msg("Total CSP-Region Pairs")
	log.Info().Int("passed", suite.PassedCspPairs).Msg("Passed Pairs")
	log.Info().Int("failed", suite.FailedCspPairs).Msg("Failed Pairs")
	log.Info().Int("skipped", suite.SkippedCspPairs).Msg("Skipped Pairs")
	log.Info().Dur("overallTime", suite.OverallTime).Msgf("Overall Time: %v", suite.OverallTime)

	log.Info().Msg("\nPer CSP-Region Results:")
	for name, ok := range suite.CspResults {
		status := "✅"
		if !ok {
			status = "❌"
		}
		log.Info().Msgf("  %s %s", status, name)
	}

	if suite.FailedCspPairs > 0 {
		os.Exit(1)
	}
}

// generateMarkdownReport writes a per-CSP test result to testresult/.
func generateMarkdownReport(report *OSTestReport) error {
	execDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	testResultDir := filepath.Join(execDir, "testresult")
	if err := os.MkdirAll(testResultDir, 0755); err != nil {
		return fmt.Errorf("failed to create testresult directory: %w", err)
	}

	filename := filepath.Join(testResultDir,
		fmt.Sprintf("os-test-results-%s.md", strings.ToLower(report.CSP)))

	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create report file: %w", err)
	}
	defer f.Close()

	content := generateMarkdownContent(report)
	content = maskSensitiveInfo(content)

	if _, err := f.WriteString(content); err != nil {
		return fmt.Errorf("failed to write report: %w", err)
	}

	fmt.Printf("📝 Report saved: %s\n", filename)
	return nil
}

// maskSensitiveInfo redacts sensitive data from the content
func maskSensitiveInfo(content string) string {
	// 1. Mask Azure Subscription IDs
	// Pattern: /subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	reSub := regexp.MustCompile(`(?i)/subscriptions/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)
	content = reSub.ReplaceAllString(content, "/subscriptions/AZURE_SUBSCRIPTION_ID")

	// 2. Mask GCP Project IDs in common URL patterns
	// Pattern: projects/project-id-123/zones/...
	reGCP := regexp.MustCompile(`projects/([a-z0-9\-]+)/`)
	content = reGCP.ReplaceAllStringFunc(content, func(match string) string {
		parts := strings.Split(match, "/")
		if len(parts) >= 2 {
			projectId := parts[1]
			if projectId == "compute" || projectId == "v1" {
				return match
			}
			return "projects/GCP_PROJECT_ID/"
		}
		return match
	})

	// 3. Mask email addresses (often contains project-id or user-id)
	reEmail := regexp.MustCompile(`[a-zA-Z0-9+_.\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}`)
	content = reEmail.ReplaceAllString(content, "MASKED_EMAIL")

	return content
}

func generateMarkdownContent(report *OSTestReport) string {
	var sb strings.Builder

	// ========================================================================
	// Header
	// ========================================================================
	sb.WriteString(fmt.Sprintf("# CM-Beetle test results for %s (object storage)\n\n", strings.ToUpper(report.CSP)))
	sb.WriteString("> [!NOTE]\n")
	sb.WriteString(fmt.Sprintf("> This document presents comprehensive test results for CM-Beetle object storage integration with %s.\n\n", strings.ToUpper(report.CSP)))

	// ========================================================================
	// Environment and scenario
	// ========================================================================
	sb.WriteString("## Environment and scenario\n\n")
	sb.WriteString("### Environment\n\n")
	sb.WriteString(fmt.Sprintf("- CM-Beetle: %s\n", getBeetleVersion()))
	sb.WriteString(fmt.Sprintf("- CB-Tumblebug: v%s\n", getVersionFromDockerCompose("cb-tumblebug")))
	sb.WriteString(fmt.Sprintf("- Target CSP: %s\n", strings.ToUpper(report.CSP)))
	sb.WriteString(fmt.Sprintf("- Target Region: %s\n", report.Region))
	sb.WriteString(fmt.Sprintf("- CM-Beetle URL: %s\n", report.BeetleURL))
	sb.WriteString(fmt.Sprintf("- Namespace: %s\n", report.NamespaceID))
	sb.WriteString(fmt.Sprintf("- Name Seed: %s\n", report.NameSeed))
	sb.WriteString("- Test CLI: Custom automated testing tool\n")
	sb.WriteString(fmt.Sprintf("- Test Date: %s\n", report.TestDate))
	sb.WriteString(fmt.Sprintf("- Test Time: %s\n", report.TestTime))
	sb.WriteString(fmt.Sprintf("- Test Execution: %s\n\n", report.TestDateTime.Format("2006-01-02 15:04:05 MST")))

	sb.WriteString("### Scenario\n\n")
	sb.WriteString("1. Recommend target object storage (buckets) via Beetle\n")
	sb.WriteString("1. Migrate (create) object storages via Beetle\n")
	sb.WriteString("1. List all object storages via Beetle\n")
	sb.WriteString("1. Check existence of first bucket (HEAD) via Beetle\n")
	sb.WriteString("1. Get first bucket details via Beetle\n")
	sb.WriteString("1. Delete all buckets (cleanup) via Beetle\n\n")

	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> Some long request/response bodies are in the collapsible section for better readability.\n\n")

	// ========================================================================
	// Test Results Summary
	// ========================================================================
	sb.WriteString(fmt.Sprintf("## Test result for %s\n\n", strings.ToUpper(report.CSP)))
	sb.WriteString("### Test Results Summary\n\n")
	sb.WriteString("| Test | Step (Endpoint / Description) | Status | Duration | Details |\n")
	sb.WriteString("|------|-------------------------------|--------|----------|----------|\n")

	endpoints := []string{
		"`POST /beetle/recommendation/middleware/objectStorage`",
		fmt.Sprintf("`POST /beetle/migration/middleware/ns/%s/objectStorage`", report.NamespaceID),
		fmt.Sprintf("`GET /beetle/migration/middleware/ns/%s/objectStorage`", report.NamespaceID),
		fmt.Sprintf("`HEAD /beetle/migration/middleware/ns/%s/objectStorage/{{osId}}`", report.NamespaceID),
		fmt.Sprintf("`GET /beetle/migration/middleware/ns/%s/objectStorage/{{osId}}`", report.NamespaceID),
		fmt.Sprintf("`DELETE /beetle/migration/middleware/ns/%s/objectStorage/{{osId}}`", report.NamespaceID),
	}

	passedCount, skippedCount := 0, 0
	for i, r := range report.TestResults {
		status := "✅ **PASS**"
		details := "Pass"
		if r.Skipped {
			status = "⏭️ **SKIP**"
			details = "Skip"
			skippedCount++
		} else if !r.Success {
			status = "❌ **FAIL**"
			details = "Fail"
		} else {
			passedCount++
		}
		endpoint := ""
		if i < len(endpoints) {
			endpoint = endpoints[i]
		}
		sb.WriteString(fmt.Sprintf("| %d | %s | %s | %v | %s |\n",
			i+1, endpoint, status, r.Duration.Truncate(time.Millisecond), details))
	}
	sb.WriteString("\n")

	totalTests := len(report.TestResults)
	if skippedCount > 0 {
		sb.WriteString(fmt.Sprintf("**Overall Result**: %d/%d tests passed, %d skipped", passedCount, totalTests, skippedCount))
	} else {
		sb.WriteString(fmt.Sprintf("**Overall Result**: %d/%d tests passed", passedCount, totalTests))
	}
	if passedCount == totalTests-skippedCount && skippedCount < totalTests {
		sb.WriteString(" ✅\n\n")
	} else if passedCount == totalTests {
		sb.WriteString(" ✅\n\n")
	} else {
		sb.WriteString(" ❌\n\n")
	}

	sb.WriteString(fmt.Sprintf("**Total Duration**: %v\n\n", report.Summary.Duration))
	sb.WriteString(fmt.Sprintf("*Test executed on %s at %s (%s) using CM-Beetle automated test CLI*\n\n",
		report.TestDate,
		report.TestTime,
		report.TestDateTime.Format("2006-01-02 15:04:05 MST")))

	// ========================================================================
	// Detailed Test Case Results
	// ========================================================================
	sb.WriteString("---\n\n")
	sb.WriteString("## Detailed Test Case Results\n\n")
	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> This section provides detailed information for each test case, including API request information and response details.\n\n")

	// Test Case 1: Recommendation
	sb.WriteString("### Test Case 1: Recommend target object storage\n\n")
	sb.WriteString("#### 1.1 API Request Information\n\n")
	sb.WriteString("- **API Endpoint**: `POST /beetle/recommendation/middleware/objectStorage`\n")
	sb.WriteString("- **Purpose**: Get object storage recommendations for migration\n")
	sb.WriteString(fmt.Sprintf("- **Target CSP**: `%s`\n", report.CSP))
	sb.WriteString(fmt.Sprintf("- **Target Region**: `%s`\n\n", report.Region))
	sb.WriteString("**Request Body**:\n\n")
	sb.WriteString("<details>\n")
	sb.WriteString("  <summary> <ins>Click to see the request body</ins> </summary>\n\n")
	sb.WriteString("```json\n")
	reqJSON, _ := json.MarshalIndent(report.OsRequest, "", "  ")
	sb.WriteString(string(reqJSON))
	sb.WriteString("\n```\n\n")
	sb.WriteString("</details>\n\n")

	sb.WriteString("#### 1.2 API Response Information\n\n")
	if len(report.TestResults) > 0 {
		r1 := report.TestResults[0]
		if r1.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString("- **Response**: Object storage recommendation generated successfully\n\n")
			if report.RecommendationResponse != nil {
				sb.WriteString("**Response Body**:\n\n")
				sb.WriteString("<details>\n")
				sb.WriteString("  <summary> <ins>Click to see the response body</ins> </summary>\n\n")
				sb.WriteString("```json\n")
				respJSON, _ := json.MarshalIndent(report.RecommendationResponse, "", "  ")
				sb.WriteString(string(respJSON))
				sb.WriteString("\n```\n\n")
				sb.WriteString("</details>\n\n")
			}
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if r1.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n\n", r1.Error))
			}
		}
	}

	// Test Case 2: Migration
	sb.WriteString("### Test Case 2: Migrate (create) object storages\n\n")
	sb.WriteString("#### 2.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `POST /beetle/migration/middleware/ns/%s/objectStorage`\n", report.NamespaceID))
	sb.WriteString("- **Purpose**: Create object storages (buckets) in the target cloud\n")
	sb.WriteString(fmt.Sprintf("- **Namespace ID**: `%s`\n\n", report.NamespaceID))
	if report.RecommendationResponse != nil {
		sb.WriteString("**Request Body** (recommendation result):\n\n")
		sb.WriteString("<details>\n")
		sb.WriteString("  <summary> <ins>Click to see the request body</ins> </summary>\n\n")
		sb.WriteString("```json\n")
		migReqJSON, _ := json.MarshalIndent(report.RecommendationResponse, "", "  ")
		sb.WriteString(string(migReqJSON))
		sb.WriteString("\n```\n\n")
		sb.WriteString("</details>\n\n")
	}
	sb.WriteString("#### 2.2 API Response Information\n\n")
	if len(report.TestResults) > 1 {
		r2 := report.TestResults[1]
		if r2.Skipped {
			sb.WriteString("- **Status**: ⏭️ **SKIPPED** (previous test failed)\n\n")
		} else if r2.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString("- **Response**: Object storages created successfully\n\n")
			if report.MigrationResponse != nil {
				sb.WriteString("**Response Body**:\n\n")
				sb.WriteString("<details>\n")
				sb.WriteString("  <summary> <ins>Click to see the response body</ins> </summary>\n\n")
				sb.WriteString("```json\n")
				migJSON, _ := json.MarshalIndent(report.MigrationResponse, "", "  ")
				sb.WriteString(string(migJSON))
				sb.WriteString("\n```\n\n")
				sb.WriteString("</details>\n\n")
			}
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if r2.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n\n", r2.Error))
			}
		}
	}

	// Test Case 3: List
	sb.WriteString("### Test Case 3: List object storages\n\n")
	sb.WriteString("#### 3.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `GET /beetle/migration/middleware/ns/%s/objectStorage`\n", report.NamespaceID))
	sb.WriteString("- **Purpose**: Retrieve all object storages in the namespace\n")
	sb.WriteString("- **Request Body**: None (GET request)\n\n")
	sb.WriteString("#### 3.2 API Response Information\n\n")
	if len(report.TestResults) > 2 {
		r3 := report.TestResults[2]
		if r3.Skipped {
			sb.WriteString("- **Status**: ⏭️ **SKIPPED** (previous test failed)\n\n")
		} else if r3.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString("- **Response**: Object storage list retrieved successfully\n\n")
			if report.ListResponse != nil {
				sb.WriteString("**Response Body**:\n\n")
				sb.WriteString("```json\n")
				listJSON, _ := json.MarshalIndent(report.ListResponse, "", "  ")
				sb.WriteString(string(listJSON))
				sb.WriteString("\n```\n\n")
			}
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if r3.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n\n", r3.Error))
			}
		}
	}

	// Test Case 4: Exist (HEAD)
	sb.WriteString("### Test Case 4: Check existence of first bucket\n\n")
	sb.WriteString("#### 4.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `HEAD /beetle/migration/middleware/ns/%s/objectStorage/{{osId}}`\n", report.NamespaceID))
	sb.WriteString("- **Purpose**: Verify the first migrated bucket exists\n")
	sb.WriteString("- **Request Body**: None (HEAD request)\n\n")
	sb.WriteString("#### 4.2 API Response Information\n\n")
	if len(report.TestResults) > 3 {
		r4 := report.TestResults[3]
		if r4.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString(fmt.Sprintf("- **HTTP Status**: %d\n\n", r4.StatusCode))
		} else if r4.Skipped {
			sb.WriteString("- **Status**: ⏭️ **SKIPPED** (previous test failed)\n\n")
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if r4.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n\n", r4.Error))
			}
		}
	}

	// Test Case 5: Get
	sb.WriteString("### Test Case 5: Get first bucket details\n\n")
	sb.WriteString("#### 5.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `GET /beetle/migration/middleware/ns/%s/objectStorage/{{osId}}`\n", report.NamespaceID))
	sb.WriteString("- **Purpose**: Retrieve detailed information for the first migrated bucket\n")
	sb.WriteString("- **Request Body**: None (GET request)\n\n")
	sb.WriteString("#### 5.2 API Response Information\n\n")
	if len(report.TestResults) > 4 {
		r5 := report.TestResults[4]
		if r5.Skipped {
			sb.WriteString("- **Status**: ⏭️ **SKIPPED** (previous test failed)\n\n")
		} else if r5.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString("- **Response**: Bucket details retrieved successfully\n\n")
			if report.GetResponse != nil {
				sb.WriteString("**Response Body**:\n\n")
				sb.WriteString("<details>\n")
				sb.WriteString("  <summary> <ins>Click to see the response body</ins> </summary>\n\n")
				sb.WriteString("```json\n")
				getJSON, _ := json.MarshalIndent(report.GetResponse, "", "  ")
				sb.WriteString(string(getJSON))
				sb.WriteString("\n```\n\n")
				sb.WriteString("</details>\n\n")
			}
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if r5.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n\n", r5.Error))
			}
		}
	}

	// Test Case 6: Delete
	sb.WriteString("### Test Case 6: Delete all buckets (cleanup)\n\n")
	sb.WriteString("#### 6.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `DELETE /beetle/migration/middleware/ns/%s/objectStorage/{{osId}}`\n", report.NamespaceID))
	sb.WriteString("- **Purpose**: Delete all migrated buckets as cleanup\n")
	sb.WriteString("- **Note**: Always runs regardless of previous test failures\n\n")
	sb.WriteString("#### 6.2 API Response Information\n\n")
	if len(report.TestResults) > 5 {
		r6 := report.TestResults[5]
		if r6.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString("- **Response**: All buckets deleted successfully\n\n")
		} else if r6.Skipped {
			sb.WriteString("- **Status**: ⏭️ **SKIPPED** (no buckets to delete)\n\n")
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if r6.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n\n", r6.Error))
			}
		}
	}

	return sb.String()
}

// getBeetleVersion returns the CM-Beetle version from git tags or commit hash.
func getBeetleVersion() string {
	hash := func() string {
		out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
		if err != nil {
			return "unknown"
		}
		return strings.TrimSpace(string(out))
	}()
	if out, err := exec.Command("git", "describe", "--tags", "--exact-match").Output(); err == nil {
		return strings.TrimSpace(string(out))
	}
	if out, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output(); err == nil {
		tag := strings.TrimSpace(string(out))
		if tag != "" {
			return fmt.Sprintf("%s+ (%s)", tag, hash)
		}
	}
	return fmt.Sprintf("main (%s)", hash)
}

// getVersionFromDockerCompose extracts a service image version from docker-compose.yaml.
func getVersionFromDockerCompose(serviceName string) string {
	for _, p := range []string{
		"../../deployments/docker-compose/docker-compose.yaml",
		"deployments/docker-compose/docker-compose.yaml",
	} {
		content, err := os.ReadFile(p)
		if err != nil {
			continue
		}
		for _, line := range strings.Split(string(content), "\n") {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "#") {
				continue
			}
			if strings.Contains(trimmed, "cloudbaristaorg/"+serviceName+":") {
				parts := strings.SplitN(trimmed, ":", 3)
				if len(parts) == 3 {
					return strings.TrimSpace(parts[2])
				}
			}
		}
	}
	return "unknown"
}

// ============================================================================
// Config & Request Loaders
// ============================================================================

func loadConfig(path string) (TestConfig, error) {
	var cfg TestConfig
	f, err := os.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("open config: %w", err)
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("decode config: %w", err)
	}
	return cfg, nil
}

func loadAuthConfig(path string) (AuthConfig, error) {
	var auth AuthConfig
	if path == "" {
		return auth, nil
	}
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return auth, nil
	}
	if err != nil {
		return auth, fmt.Errorf("open auth config: %w", err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&auth); err != nil {
		return auth, fmt.Errorf("decode auth config: %w", err)
	}
	return auth, nil
}

func loadOsRequest(path string) (OsRequest, error) {
	var req OsRequest
	f, err := os.Open(path)
	if err != nil {
		return req, fmt.Errorf("open request file: %w", err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&req); err != nil {
		return req, fmt.Errorf("decode request file: %w", err)
	}
	return req, nil
}
