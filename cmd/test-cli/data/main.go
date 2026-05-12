// Package main is the starting point of CM-Beetle Data Migration Test CLI
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
	"github.com/cloud-barista/cm-beetle/transx"
)

// ============================================================================
// Configuration & Type Definitions
// ============================================================================

// TestConfig holds the full test configuration from test-config.yaml.
type TestConfig struct {
	Test struct {
		Set struct {
			Mode string `yaml:"mode"` // parallel or sequential
		} `yaml:"set"`
		Source struct {
			Csp      string `yaml:"csp"`      // CSP for source Object Storage (e.g., aws)
			Region   string `yaml:"region"`   // Region for source Object Storage (e.g., ap-northeast-2)
			NameSeed string `yaml:"nameSeed"` // Prefix for source OS naming (default: target nameSeed + "s")
		} `yaml:"source"`
		Upload struct {
			Filter struct {
				Include []string `yaml:"include"` // File patterns to include (empty = all)
				Exclude []string `yaml:"exclude"` // File patterns to exclude
			} `yaml:"filter"`
		} `yaml:"upload"`
		Cases []TestCase `yaml:"cases"`
	} `yaml:"test"`

	Beetle struct {
		Endpoint          string `yaml:"endpoint"`
		NamespaceID       string `yaml:"namespaceId"`
		AuthConfigFile    string `yaml:"authConfigFile"`
		OsCreationReqFile string `yaml:"osCreationReqFile"` // Path to JSON file with bucket specs for direct OS creation
	} `yaml:"beetle"`

	Tumblebug struct {
		Endpoint       string `yaml:"endpoint"`       // Used by test CLI (host) to call Tumblebug directly
		ServerEndpoint string `yaml:"serverEndpoint"` // Used inside the request body sent to CM-Beetle server (Docker network)
		Auth           struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"auth"`
	} `yaml:"tumblebug"`

	TestData struct {
		BaseDir string `yaml:"baseDir"` // Path to store generated dummy data
		Cleanup bool   `yaml:"cleanup"` // Delete dummy data after test
	} `yaml:"testData"`

	Migration struct {
		NameSeed string `yaml:"nameSeed"` // Seed for target object storage naming
		Poll     struct {
			IntervalSec int `yaml:"intervalSec"` // Polling interval in seconds
			TimeoutSec  int `yaml:"timeoutSec"`  // Polling timeout in seconds
		} `yaml:"poll"`
	} `yaml:"migration"`
}

// TestCase defines a single target CSP-Region pair to test.
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

// OsCreationReq holds the bucket specifications loaded from the OS creation request JSON file.
type OsCreationReq struct {
	TargetObjectStorages []storagemodel.TargetObjectStorage `json:"targetObjectStorages"`
}

// ============================================================================
// Encryption Key Response Type
// ============================================================================

// EncryptionKeyResponse wraps the response from GET /beetle/migration/data/encryptionKey.
type EncryptionKeyResponse struct {
	Success bool `json:"success"`
	Data    struct {
		KeyID     string `json:"keyId"`
		Algorithm string `json:"algorithm"`
		PublicKey string `json:"publicKey"`
		ExpiresAt string `json:"expiresAt"`
	} `json:"data"`
}

// ============================================================================
// Async API Response Types
// ============================================================================

// AsyncJobResponseWrapper wraps the 202 response from POST /beetle/migration/data.
type AsyncJobResponseWrapper struct {
	Success bool `json:"success"`
	Data    struct {
		ReqID     string `json:"reqId"`
		Status    string `json:"status"`
		StatusURL string `json:"statusUrl"`
	} `json:"data"`
}

// RequestStatusWrapper wraps the GET /beetle/request/{reqId} response.
type RequestStatusWrapper struct {
	Success bool `json:"success"`
	Data    struct {
		Status        string `json:"status"`
		ErrorResponse string `json:"errorResponse"`
	} `json:"data"`
}

// ============================================================================
// Test Result & Report Types
// ============================================================================

// TestResults holds the result of a single test step.
type TestResults struct {
	TestName     string        `json:"testName"`
	StartTime    time.Time     `json:"startTime"`
	EndTime      time.Time     `json:"endTime"`
	Duration     time.Duration `json:"duration"`
	Success      bool          `json:"success"`
	Skipped      bool          `json:"skipped"`
	StatusCode   int           `json:"statusCode"`
	Response     interface{}   `json:"response,omitempty"`
	Error        string        `json:"error,omitempty"`
	RequestURL   string        `json:"requestUrl,omitempty"`
	RequestBody  interface{}   `json:"requestBody,omitempty"`  // Sanitized request body (passwords masked)
	ResponseBody interface{}   `json:"responseBody,omitempty"` // Parsed response body
	PollNote     string        `json:"pollNote,omitempty"`     // Describes async polling performed after initial response
}

// VerificationResult holds object-level comparison between source and target buckets.
type VerificationResult struct {
	SourceObjectCount int
	TargetObjectCount int
	MatchedCount      int
	MissingInTarget   []string // keys present in source but absent from target
	ExtraInTarget     []string // keys present in target but absent from source
	Matched           bool     // true if source and target contents are identical
	SrcKeys           []string // all source object keys
	DstKeys           []string // all target object keys
}

// DataTestReport holds all results for a single CSP-Region pair.
type DataTestReport struct {
	CSP          string
	SourceCSP    string // CSP of the source object storage
	SourceRegion string // Region of the source object storage
	Region       string
	DisplayName  string
	TestDate     string
	TestTime     string
	TestDateTime time.Time
	BeetleURL    string
	NamespaceID  string
	NameSeed     string
	SourceOsId   string
	TargetOsId   string
	TestResults  []TestResults
	Summary      TestResults
	VerifyResult *VerificationResult
}

// TestSuite aggregates results across all CSP-Region pairs.
type TestSuite struct {
	TotalCspPairs   int
	PassedCspPairs  int
	FailedCspPairs  int
	SkippedCspPairs int
	OverallTime     time.Duration
	CspResults      map[string]bool
	mu              sync.Mutex
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
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	cfg, err := loadConfig(*configFile)
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	authConfig, err := loadAuthConfig(cfg.Beetle.AuthConfigFile)
	if err != nil {
		return fmt.Errorf("load auth config: %w", err)
	}

	osCreationReq, err := loadOsCreationReq(cfg.Beetle.OsCreationReqFile)
	if err != nil {
		return fmt.Errorf("load OS creation request: %w", err)
	}

	// Check CM-Beetle readiness before starting tests.
	probeClient := resty.New().SetTimeout(10 * time.Second)
	if err := checkReadiness(probeClient, cfg.Beetle.Endpoint); err != nil {
		return fmt.Errorf("CM-Beetle not ready: %w", err)
	}

	// Generate dummy test data once (before running per-case tests).
	if err := generateDummyData(cfg); err != nil {
		return fmt.Errorf("generate dummy data: %w", err)
	}

	// Create shared HTTP client for pre/post-flight steps.
	mainClient := resty.New()
	mainClient.SetTimeout(10 * time.Minute)
	mainClient.SetBaseURL(cfg.Beetle.Endpoint)
	if authConfig.BasicAuthUsername != "" {
		mainClient.SetBasicAuth(authConfig.BasicAuthUsername, authConfig.BasicAuthPassword)
	}

	// Pre-flight: create source Object Storage and upload data (shared across all test cases).
	sourceNameSeed := cfg.Test.Source.NameSeed
	if sourceNameSeed == "" {
		sourceNameSeed = cfg.Migration.NameSeed + "s"
	}
	sourceTc := TestCase{Csp: cfg.Test.Source.Csp, Region: cfg.Test.Source.Region, Name: "source"}
	sourceOsId, preFlightErr := runPreFlightSteps(mainClient, cfg, authConfig, osCreationReq, sourceNameSeed, sourceTc)

	// defer post-flight cleanup: always runs regardless of how run() exits
	// (normal completion, pre-flight error, signal, or panic recovery).
	defer func() {
		if sourceOsId != "" {
			log.Info().Msgf("%s", "\n"+strings.Repeat("=", 60))
			log.Info().Msg("Post-flight: deleting shared source Object Storage")
			log.Info().Msgf("%s", strings.Repeat("=", 60))
			runDeleteOSStep(mainClient, cfg, sourceOsId, "source",
				"Post-flight: DELETE /migration/middleware/ns/{nsId}/objectStorage/{sourceOsId} (cleanup)",
				"empty")
		}
		if cfg.TestData.Cleanup {
			if err := os.RemoveAll(cfg.TestData.BaseDir); err != nil {
				log.Warn().Err(err).Msg("Failed to remove dummy data directory")
			} else {
				log.Info().Str("dir", cfg.TestData.BaseDir).Msg("Dummy data cleaned up")
			}
		}
	}()

	if preFlightErr != nil {
		log.Error().Err(preFlightErr).Msg("Pre-flight steps failed")
		return fmt.Errorf("pre-flight: %w", preFlightErr)
	}

	// Signal handler: on SIGINT/SIGTERM, cancel remaining test cases.
	// The deferred cleanup above will still execute after run() returns.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)
	go func() {
		select {
		case sig := <-sigCh:
			log.Warn().Str("signal", sig.String()).Msg("Interrupted — skipping remaining tests, cleanup will run")
			cancel()
		case <-ctx.Done():
		}
	}()

	suite := &TestSuite{
		TotalCspPairs: len(cfg.Test.Cases),
		CspResults:    make(map[string]bool),
	}

	startTime := time.Now()
	isParallel := cfg.Test.Set.Mode == "parallel"
	var wg sync.WaitGroup

	for i, tc := range cfg.Test.Cases {
		if ctx.Err() != nil {
			log.Warn().Msg("Interrupted — skipping remaining test cases")
			break
		}

		if !tc.Execute {
			log.Info().Msgf("[%d/%d] Skipping %s (execute=false)", i+1, len(cfg.Test.Cases), displayName(tc))
			suite.mu.Lock()
			suite.SkippedCspPairs++
			suite.mu.Unlock()
			continue
		}

		// Append index suffix to nameSeed in parallel mode to avoid name collisions.
		caseNameSeed := cfg.Migration.NameSeed
		if isParallel && len(cfg.Test.Cases) > 1 {
			caseNameSeed = fmt.Sprintf("%s%02d", cfg.Migration.NameSeed, i+1)
		}

		if isParallel {
			wg.Add(1)
			go func(idx int, tc TestCase, seed string) {
				defer wg.Done()
				runTestCase(idx, tc, cfg, authConfig, osCreationReq, sourceOsId, seed, suite)
			}(i, tc, caseNameSeed)
		} else {
			runTestCase(i, tc, cfg, authConfig, osCreationReq, sourceOsId, caseNameSeed, suite)
		}
	}

	if isParallel {
		wg.Wait()
	}

	suite.OverallTime = time.Since(startTime)
	printFinalSummary(suite)

	if suite.FailedCspPairs > 0 {
		return fmt.Errorf("%d CSP-Region pair(s) failed", suite.FailedCspPairs)
	}
	return nil
}

// ============================================================================
// Pre/Post-Flight Steps (shared across all test cases)
// ============================================================================

// runPreFlightSteps creates the shared source Object Storage and uploads dummy data to it.
// This runs once in main() before all per-target test cases begin.
// Returns the created source OS ID, or an error if any step fails.
func runPreFlightSteps(client *resty.Client, cfg TestConfig, authConfig AuthConfig, osCreationReq OsCreationReq, sourceNameSeed string, sourceTc TestCase) (string, error) {
	log.Info().Msgf("%s", "\n"+strings.Repeat("=", 60))
	log.Info().Msg("Pre-flight: creating shared source Object Storage and uploading dummy data")
	log.Info().Msgf("%s", strings.Repeat("=", 60))

	osId, r1 := runCreateOSFromFileStep(client, cfg, osCreationReq, sourceNameSeed, sourceTc, "source",
		"Pre-flight: POST /migration/middleware/ns/{nsId}/objectStorage")
	if !r1.Success {
		return "", fmt.Errorf("create source Object Storage: %s", r1.Error)
	}

	dummyDataAbsPath, err := filepath.Abs(cfg.TestData.BaseDir)
	if err != nil {
		return osId, fmt.Errorf("resolve dummy data path: %w", err)
	}

	r2 := runUploadDummyDataStep(client, cfg, authConfig, osId, dummyDataAbsPath, "source",
		"Pre-flight: POST /migration/data (upload: local → source OS, encrypted, with filter, async)")
	if !r2.Success {
		// Clean up the created source OS before returning the error.
		log.Warn().Str("sourceOsId", osId).Msg("Upload failed — deleting source Object Storage for cleanup")
		runDeleteOSStep(client, cfg, osId, "source",
			"Pre-flight cleanup: DELETE /migration/middleware/ns/{nsId}/objectStorage",
			"empty")
		return "", fmt.Errorf("upload data to source Object Storage: %s", r2.Error)
	}

	log.Info().Str("sourceOsId", osId).Msg("Pre-flight complete: source Object Storage ready with uploaded data")
	return osId, nil
}

// ============================================================================
// Test Execution Per CSP-Region Pair
// ============================================================================

// runTestCase executes the full data migration test for one CSP-Region target.
// Steps (per target):
//  1. Create target Object Storage (from JSON file)
//  2. Migrate data: source OS → target OS (encrypted, async with polling)
//  3. Delete target Object Storage (cleanup)
//  4. Generate markdown report
//
// Source Object Storage creation and data upload are performed once in main()
// (pre-flight) and shared across all test cases. Source OS deletion is performed
// once in main() (post-flight) after all test cases complete.
func runTestCase(idx int, tc TestCase, cfg TestConfig, authConfig AuthConfig, osCreationReq OsCreationReq, sourceOsId, nameSeed string, suite *TestSuite) {
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
	report := &DataTestReport{
		CSP:          tc.Csp,
		SourceCSP:    cfg.Test.Source.Csp,
		SourceRegion: cfg.Test.Source.Region,
		Region:       tc.Region,
		DisplayName:  name,
		TestDate:     pairStartTime.Format("January 2, 2006"),
		TestTime:     pairStartTime.Format("15:04:05 MST"),
		TestDateTime: pairStartTime,
		BeetleURL:    cfg.Beetle.Endpoint,
		NamespaceID:  cfg.Beetle.NamespaceID,
		NameSeed:     nameSeed,
		SourceOsId:   sourceOsId,
		TestResults:  make([]TestResults, 0),
	}

	record := func(result TestResults) {
		suite.mu.Lock()
		defer suite.mu.Unlock()
		report.TestResults = append(report.TestResults, result)
		if !result.Success && !result.Skipped {
			pairFailed++
		}
	}

	stopTesting := false
	var targetOsId string

	// Step 1: Create target Object Storage (directly from JSON file).
	if !stopTesting {
		osId, r1 := runCreateOSFromFileStep(client, cfg, osCreationReq, nameSeed, tc, name,
			"Step 1: POST /migration/middleware/ns/{nsId}/objectStorage (target)")
		targetOsId = osId
		report.TargetOsId = targetOsId
		record(r1)
		if !r1.Success {
			stopTesting = true
		}
	} else {
		record(skippedResult("Step 1: Create target Object Storage", name))
	}

	// Step 2: Migrate data: source OS → target OS (encrypted, async).
	if !stopTesting && sourceOsId != "" && targetOsId != "" {
		r2 := runMigrateDataStep(client, cfg, authConfig, sourceOsId, targetOsId, name,
			"Step 2: POST /migration/data (migrate: source OS → target OS, encrypted, async)")
		record(r2)
		if !r2.Success {
			stopTesting = true
		}
	} else {
		record(skippedResult("Step 2: Migrate data: source OS → target OS", name))
	}

	// Step 3: Verify migrated data (compare source and target object lists).
	if !stopTesting && sourceOsId != "" && targetOsId != "" {
		vr, r3 := runVerifyDataStep(cfg, authConfig, sourceOsId, targetOsId, name,
			"Step 3: Verify migrated data (compare source and target object lists)")
		report.VerifyResult = &vr
		record(r3)
		// Verification failure is noted but does not block cleanup (Step 4).
	} else {
		record(skippedResult("Step 3: Verify migrated data", name))
	}

	// Step 4: Delete target Object Storage (cleanup — always runs).
	if targetOsId != "" {
		r4 := runDeleteOSStep(client, cfg, targetOsId, name,
			"Step 4: DELETE /migration/middleware/ns/{nsId}/objectStorage/{targetOsId} (cleanup)",
			"empty") // empty the bucket first, then delete
		record(r4)
	} else {
		record(skippedResult("Step 4: Delete target Object Storage (cleanup)", name))
	}

	_ = stopTesting

	pairDuration := time.Since(pairStartTime)
	report.Summary = TestResults{
		TestName:  fmt.Sprintf("CSP-Region Pair: %s", name),
		StartTime: pairStartTime,
		EndTime:   time.Now(),
		Duration:  pairDuration,
		Success:   pairFailed == 0,
	}

	if err := generateMarkdownReport(report); err != nil {
		log.Warn().Err(err).Str("target", name).Msg("Failed to generate markdown report")
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

	log.Info().Msgf("--- %s: %d of 4 migration steps passed (duration: %v) ---", name, 4-pairFailed, pairDuration)
}

// ============================================================================
// Individual Test Step Functions
// ============================================================================

// fetchEncryptionKey requests a one-time RSA public key bundle from
// GET /beetle/migration/data/encryptionKey and returns it for use with
// transx.ParsePublicKeyBundle and transx.EncryptModel.
// The key is only consumed (invalidated) when the server decrypts a request with it.
func fetchEncryptionKey(client *resty.Client, cfg TestConfig) (transx.PublicKeyBundle, error) {
	url := fmt.Sprintf("%s/beetle/migration/data/encryptionKey", cfg.Beetle.Endpoint)
	var keyResp EncryptionKeyResponse
	resp, err := client.R().SetResult(&keyResp).Get(url)
	if err != nil {
		return transx.PublicKeyBundle{}, fmt.Errorf("get encryption key: %w", err)
	}
	if resp.IsError() {
		return transx.PublicKeyBundle{}, fmt.Errorf("get encryption key HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
	}
	return transx.PublicKeyBundle{
		KeyID:     keyResp.Data.KeyID,
		Algorithm: keyResp.Data.Algorithm,
		PublicKey: keyResp.Data.PublicKey,
		ExpiresAt: keyResp.Data.ExpiresAt,
	}, nil
}

// runCreateOSFromFileStep builds a RecommendedObjectStorage from osCreationReq and calls POST
// /beetle/migration/middleware/ns/{nsId}/objectStorage directly (no recommendation step).
// Returns the computed osId and the test result.
func runCreateOSFromFileStep(
	client *resty.Client, cfg TestConfig,
	osCreationReq OsCreationReq, nameSeed string, tc TestCase, name, stepLabel string,
) (string, TestResults) {
	testLabel := stepLabel
	log.Info().Msgf("\n--- %s (%s) ---", testLabel, name)

	// Allow previous operations to settle.
	time.Sleep(3 * time.Second)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	if len(osCreationReq.TargetObjectStorages) == 0 {
		result.Error = "no targetObjectStorages defined in os creation request file"
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		log.Error().Msgf("❌ %s failed: %s", testLabel, result.Error)
		return "", result
	}

	reqBody := controller.MigrateObjectStorageRequest{
		RecommendedObjectStorage: storagemodel.RecommendedObjectStorage{
			Status:               "recommended",
			Description:          "Direct creation (no recommendation step)",
			TargetCloud:          storagemodel.CloudProperty{Csp: tc.Csp, Region: tc.Region},
			TargetObjectStorages: osCreationReq.TargetObjectStorages,
		},
	}

	osId := common.ComposeName(osCreationReq.TargetObjectStorages[0].BucketName, nameSeed)

	url := fmt.Sprintf("%s/beetle/migration/middleware/ns/%s/objectStorage",
		cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID)
	result.RequestURL = url
	result.RequestBody = reqBody

	resp, err := client.R().
		SetBody(reqBody).
		SetQueryParam("nameSeed", nameSeed).
		Post(url)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.StatusCode = resp.StatusCode()

	// Capture response body (both success and error cases).
	if len(resp.Body()) > 0 {
		var respData interface{}
		if jsonErr := json.Unmarshal(resp.Body(), &respData); jsonErr == nil {
			result.ResponseBody = respData
		}
	}

	if err != nil {
		result.Error = err.Error()
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return "", result
	}
	// 409 Conflict means the OS already exists — treat as success (idempotent).
	if resp.StatusCode() == 409 {
		result.Success = true
		log.Info().Msgf("✅ %s passed: OS already exists, reusing (osId: %s)", testLabel, osId)
		return osId, result
	}
	if resp.IsError() {
		result.Error = fmt.Sprintf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
		log.Error().Msgf("❌ %s failed: HTTP %d — %s", testLabel, resp.StatusCode(), string(resp.Body()))
		return "", result
	}

	result.Success = true
	log.Info().Msgf("✅ %s passed: OS created (osId: %s)", testLabel, osId)
	return osId, result
}

// runUploadDummyDataStep uploads local dummy data directly to the source OS using transx.Transfer().
// This runs in the test CLI process (not via CM-Beetle API) because the CM-Beetle API intentionally
// rejects local filesystem access for security reasons (server-side filesystem protection).
// Applies file filter from cfg.Upload.Filter.
func runUploadDummyDataStep(
	client *resty.Client, cfg TestConfig, authConfig AuthConfig,
	sourceOsId, dummyDataPath, name, stepLabel string,
) TestResults {
	testLabel := stepLabel
	log.Info().Msgf("\n--- %s (%s) ---", testLabel, name)

	// Allow source OS creation to settle before uploading.
	time.Sleep(3 * time.Second)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	tbAuth := buildTumblebugAuth(cfg, authConfig)

	var filter *transx.FilterOption
	if len(cfg.Test.Upload.Filter.Include) > 0 || len(cfg.Test.Upload.Filter.Exclude) > 0 {
		filter = &transx.FilterOption{
			Include: cfg.Test.Upload.Filter.Include,
			Exclude: cfg.Test.Upload.Filter.Exclude,
		}
	}

	model := transx.DataMigrationModel{
		Source: transx.DataLocation{
			StorageType: transx.StorageTypeFilesystem,
			Path:        dummyDataPath,
			Filesystem:  &transx.FilesystemAccess{AccessType: transx.AccessTypeLocal},
			Filter:      filter,
		},
		Destination: transx.DataLocation{
			StorageType: transx.StorageTypeObjectStorage,
			Path:        sourceOsId,
			ObjectStorage: &transx.ObjectStorageAccess{
				AccessType: transx.AccessTypeTumblebug,
				Tumblebug: &transx.TumblebugConfig{
					Endpoint: cfg.Tumblebug.Endpoint,
					NsId:     cfg.Beetle.NamespaceID,
					OsId:     sourceOsId,
					Auth:     tbAuth,
				},
			},
		},
		Strategy: transx.StrategyAuto,
	}

	log.Info().
		Str("source", dummyDataPath).
		Str("destination", sourceOsId).
		Msg("Uploading local dummy data directly to source Object Storage via transx.Transfer()")

	if err := transx.Transfer(model); err != nil {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = err.Error()
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return result
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.Success = true
	result.Response = map[string]string{"source": dummyDataPath, "destination": sourceOsId}
	log.Info().Dur("duration", result.Duration).Msgf("✅ %s passed: upload completed successfully", testLabel)
	return result
}

// runMigrateDataStep calls POST /beetle/migration/data to migrate data between two Object Storages.
// Source is a Tumblebug Object Storage (source OS → target OS). Fetches a one-time encryption key
// and encrypts sensitive fields (Tumblebug auth passwords) before sending.
// Runs asynchronously and polls for completion.
func runMigrateDataStep(
	client *resty.Client, cfg TestConfig, authConfig AuthConfig,
	sourceOsId, targetOsId, name, stepLabel string,
) TestResults {
	testLabel := stepLabel
	log.Info().Msgf("\n--- %s (%s) ---", testLabel, name)

	// Allow target OS creation to settle before migrating.
	time.Sleep(3 * time.Second)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	tbAuth := buildTumblebugAuth(cfg, authConfig)
	// serverTBEndpoint: CM-Beetle server uses this to reach Tumblebug (Docker network).
	// Falls back to Endpoint when ServerEndpoint is not set.
	serverTBEndpoint := cfg.Tumblebug.ServerEndpoint
	if serverTBEndpoint == "" {
		serverTBEndpoint = cfg.Tumblebug.Endpoint
	}

	model := transx.DataMigrationModel{
		Source: transx.DataLocation{
			StorageType: transx.StorageTypeObjectStorage,
			Path:        sourceOsId,
			ObjectStorage: &transx.ObjectStorageAccess{
				AccessType: transx.AccessTypeTumblebug,
				Tumblebug: &transx.TumblebugConfig{
					Endpoint: serverTBEndpoint,
					NsId:     cfg.Beetle.NamespaceID,
					OsId:     sourceOsId,
					Auth:     tbAuth,
				},
			},
		},
		Destination: transx.DataLocation{
			StorageType: transx.StorageTypeObjectStorage,
			Path:        targetOsId,
			ObjectStorage: &transx.ObjectStorageAccess{
				AccessType: transx.AccessTypeTumblebug,
				Tumblebug: &transx.TumblebugConfig{
					Endpoint: serverTBEndpoint,
					NsId:     cfg.Beetle.NamespaceID,
					OsId:     targetOsId,
					Auth:     tbAuth,
				},
			},
		},
		Strategy: transx.StrategyAuto,
	}

	// Build sanitized request body for reporting (passwords masked).
	type sanitizedAuth struct {
		AuthType string `json:"authType,omitempty"`
		Username string `json:"username,omitempty"`
		Password string `json:"password"` // always masked
	}
	type sanitizedTB struct {
		Endpoint string         `json:"endpoint"`
		NsId     string         `json:"nsId"`
		OsId     string         `json:"osId"`
		Auth     *sanitizedAuth `json:"auth,omitempty"`
	}
	type sanitizedOS struct {
		AccessType string       `json:"accessType"`
		Tumblebug  *sanitizedTB `json:"tumblebug,omitempty"`
	}
	type sanitizedLoc struct {
		StorageType   string       `json:"storageType"`
		Path          string       `json:"path"`
		ObjectStorage *sanitizedOS `json:"objectStorage,omitempty"`
	}
	type sanitizedMigReq struct {
		Source      sanitizedLoc `json:"source"`
		Destination sanitizedLoc `json:"destination"`
		Strategy    string       `json:"strategy"`
		Note        string       `json:"_note"` // informational
	}
	buildSanitizedLoc := func(loc transx.DataLocation) sanitizedLoc {
		s := sanitizedLoc{StorageType: loc.StorageType, Path: loc.Path}
		if loc.ObjectStorage != nil {
			os := &sanitizedOS{AccessType: loc.ObjectStorage.AccessType}
			if loc.ObjectStorage.Tumblebug != nil {
				tb := loc.ObjectStorage.Tumblebug
				sa := &sanitizedTB{Endpoint: tb.Endpoint, NsId: tb.NsId, OsId: tb.OsId}
				if tb.Auth != nil && tb.Auth.Basic != nil {
					sa.Auth = &sanitizedAuth{
						AuthType: tb.Auth.AuthType,
						Username: tb.Auth.Basic.Username,
						Password: "***",
					}
				}
				os.Tumblebug = sa
			}
			s.ObjectStorage = os
		}
		return s
	}
	result.RequestBody = sanitizedMigReq{
		Source:      buildSanitizedLoc(model.Source),
		Destination: buildSanitizedLoc(model.Destination),
		Strategy:    model.Strategy,
		Note:        "Sensitive fields (passwords) are masked with *** in this report",
	}

	// Obtain a one-time encryption key and encrypt sensitive fields before sending.
	bundle, err := fetchEncryptionKey(client, cfg)
	if err != nil {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = fmt.Sprintf("failed to get encryption key: %s", err)
		log.Error().Err(err).Msgf("❌ %s failed: cannot get encryption key", testLabel)
		return result
	}
	pubKey, err := transx.ParsePublicKeyBundle(bundle)
	if err != nil {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = fmt.Sprintf("failed to parse encryption key: %s", err)
		log.Error().Err(err).Msgf("❌ %s failed: cannot parse encryption key", testLabel)
		return result
	}
	encModel, err := transx.EncryptModel(model, pubKey, bundle.KeyID)
	if err != nil {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = fmt.Sprintf("failed to encrypt request: %s", err)
		log.Error().Err(err).Msgf("❌ %s failed: cannot encrypt request", testLabel)
		return result
	}
	log.Info().Str("keyId", bundle.KeyID).Msg("Request encrypted for migration")

	url := fmt.Sprintf("%s/beetle/migration/data", cfg.Beetle.Endpoint)
	result.RequestURL = url

	var asyncResp AsyncJobResponseWrapper
	resp, err := client.R().
		SetBody(encModel).
		SetResult(&asyncResp).
		Post(url)

	result.StatusCode = resp.StatusCode()

	// Capture initial API response body.
	if len(resp.Body()) > 0 {
		var respData interface{}
		if jsonErr := json.Unmarshal(resp.Body(), &respData); jsonErr == nil {
			result.ResponseBody = respData
		}
	}

	if err != nil {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = err.Error()
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return result
	}
	if resp.StatusCode() != 202 {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = fmt.Sprintf("HTTP %d: %s", resp.StatusCode(), string(resp.Body()))
		log.Error().Msgf("❌ %s failed: expected 202, got HTTP %d — %s", testLabel, resp.StatusCode(), string(resp.Body()))
		return result
	}

	reqId := asyncResp.Data.ReqID
	if reqId == "" {
		// Try extracting reqId from X-Request-Id header as fallback.
		reqId = resp.Header().Get("X-Request-Id")
	}
	if reqId == "" {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = "no reqId in migration response"
		log.Error().Msgf("❌ %s failed: no reqId returned", testLabel)
		return result
	}

	log.Info().Str("reqId", reqId).Msgf("Migration started, polling for completion...")

	// Poll for migration completion.
	intervalSec := cfg.Migration.Poll.IntervalSec
	if intervalSec <= 0 {
		intervalSec = 10
	}
	timeoutSec := cfg.Migration.Poll.TimeoutSec
	if timeoutSec <= 0 {
		timeoutSec = 600
	}

	statusURL := fmt.Sprintf("%s/beetle/request/%s", cfg.Beetle.Endpoint, reqId)
	deadline := time.Now().Add(time.Duration(timeoutSec) * time.Second)
	finalStatus := ""
	finalErrMsg := ""

	for time.Now().Before(deadline) {
		time.Sleep(time.Duration(intervalSec) * time.Second)

		var statusResp RequestStatusWrapper
		statusR, statusErr := client.R().
			SetResult(&statusResp).
			Get(statusURL)

		if statusErr != nil {
			log.Warn().Err(statusErr).Msg("Failed to poll migration status")
			continue
		}
		if statusR.IsError() {
			log.Warn().Msgf("Status check HTTP %d", statusR.StatusCode())
			continue
		}

		finalStatus = statusResp.Data.Status
		finalErrMsg = statusResp.Data.ErrorResponse
		log.Info().Str("status", finalStatus).Msgf("Migration status: %s", finalStatus)

		if finalStatus == "Success" || finalStatus == "Error" {
			break
		}
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	switch finalStatus {
	case "Success":
		result.Success = true
		result.Response = map[string]string{"reqId": reqId, "status": "Success"}
		result.PollNote = fmt.Sprintf(
			"Initial response was `202 Accepted` (status: `Handling`). "+
				"The test CLI polled `GET /beetle/request/%s` every %ds (timeout: %ds) "+
				"until migration completed with status `Success`. Step 3 was then executed.",
			reqId, intervalSec, timeoutSec)
		log.Info().Str("reqId", reqId).Dur("duration", result.Duration).
			Msgf("✅ %s passed: migration completed successfully", testLabel)
	case "Error":
		result.Error = fmt.Sprintf("Migration failed: %s", finalErrMsg)
		result.PollNote = fmt.Sprintf(
			"Initial response was `202 Accepted` (status: `Handling`). "+
				"The test CLI polled `GET /beetle/request/%s` every %ds (timeout: %ds). "+
				"Migration ended with status `Error`.",
			reqId, intervalSec, timeoutSec)
		log.Error().Str("reqId", reqId).Msgf("❌ %s failed: %s", testLabel, finalErrMsg)
	default:
		result.Error = fmt.Sprintf("Migration timed out after %ds (last status: %s)", timeoutSec, finalStatus)
		result.PollNote = fmt.Sprintf(
			"Initial response was `202 Accepted` (status: `Handling`). "+
				"The test CLI polled `GET /beetle/request/%s` every %ds but migration did not complete within %ds.",
			reqId, intervalSec, timeoutSec)
		log.Error().Str("reqId", reqId).Msgf("❌ %s timed out", testLabel)
	}

	return result
}

// runVerifyDataStep lists objects in both source and target buckets via the CM-Beetle API
// (GET /migration/middleware/ns/{nsId}/objectStorage/{osId}/object) and compares them.
// It verifies that all objects migrated from source are present in target and no unexpected
// objects were added. Runs after migration and before target OS deletion.
func runVerifyDataStep(
	cfg TestConfig, authConfig AuthConfig,
	sourceOsId, targetOsId, name, stepLabel string,
) (VerificationResult, TestResults) {
	testLabel := stepLabel
	log.Info().Msgf("\n--- %s (%s) ---", testLabel, name)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}
	var verifyResult VerificationResult

	client := resty.New().SetTimeout(2 * time.Minute)

	// listObjects calls the CM-Beetle list-objects API and returns object keys.
	type objectEntry struct {
		Key string `json:"key"`
	}
	type listData struct {
		OsId    string        `json:"osId"`
		Count   int           `json:"count"`
		Objects []objectEntry `json:"objects"`
	}
	type listResp struct {
		Success bool     `json:"success"`
		Data    listData `json:"data"`
		Message string   `json:"message"`
		Error   string   `json:"error"`
	}

	listObjects := func(osId string) ([]string, error) {
		url := fmt.Sprintf("%s/beetle/migration/middleware/ns/%s/objectStorage/%s/object",
			cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, osId)

		var resp listResp
		httpResp, err := client.R().
			SetBasicAuth(authConfig.BasicAuthUsername, authConfig.BasicAuthPassword).
			SetResult(&resp).
			Get(url)
		if err != nil {
			return nil, fmt.Errorf("request failed: %w", err)
		}
		if httpResp.IsError() {
			return nil, fmt.Errorf("API error %s: %s", httpResp.Status(), httpResp.String())
		}
		if !resp.Success {
			return nil, fmt.Errorf("API returned failure: %s", resp.Error)
		}

		keys := make([]string, 0, len(resp.Data.Objects))
		for _, o := range resp.Data.Objects {
			keys = append(keys, o.Key)
		}
		return keys, nil
	}

	// List source bucket objects.
	srcKeys, err := listObjects(sourceOsId)
	if err != nil {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = fmt.Sprintf("failed to list source objects: %s", err)
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return verifyResult, result
	}

	// List target bucket objects.
	dstKeys, err := listObjects(targetOsId)
	if err != nil {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = fmt.Sprintf("failed to list target objects: %s", err)
		log.Error().Err(err).Msgf("❌ %s failed", testLabel)
		return verifyResult, result
	}

	// Capture the raw object lists as response body for the report.
	type objList struct {
		OsId    string   `json:"osId"`
		Count   int      `json:"count"`
		Objects []string `json:"objects"`
	}
	result.ResponseBody = map[string]interface{}{
		"source": objList{OsId: sourceOsId, Count: len(srcKeys), Objects: srcKeys},
		"target": objList{OsId: targetOsId, Count: len(dstKeys), Objects: dstKeys},
	}

	// If source is empty, the upload step may have failed silently.
	if len(srcKeys) == 0 {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)
		result.Error = "source bucket has 0 objects — pre-flight upload may have failed"
		log.Error().Str("sourceOsId", sourceOsId).Msgf("❌ %s failed: %s", testLabel, result.Error)
		return verifyResult, result
	}

	// Compare object keys.
	srcKeySet := make(map[string]bool, len(srcKeys))
	for _, k := range srcKeys {
		srcKeySet[k] = true
	}
	dstKeySet := make(map[string]bool, len(dstKeys))
	for _, k := range dstKeys {
		dstKeySet[k] = true
	}

	var missing, extra []string
	for k := range srcKeySet {
		if !dstKeySet[k] {
			missing = append(missing, k)
		}
	}
	for k := range dstKeySet {
		if !srcKeySet[k] {
			extra = append(extra, k)
		}
	}
	sort.Strings(missing)
	sort.Strings(extra)

	matchedCount := len(srcKeys) - len(missing)
	if matchedCount < 0 {
		matchedCount = 0
	}

	verifyResult = VerificationResult{
		SourceObjectCount: len(srcKeys),
		TargetObjectCount: len(dstKeys),
		MatchedCount:      matchedCount,
		MissingInTarget:   missing,
		ExtraInTarget:     extra,
		Matched:           len(missing) == 0 && len(extra) == 0,
		SrcKeys:           srcKeys,
		DstKeys:           dstKeys,
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.Success = verifyResult.Matched

	if verifyResult.Matched {
		log.Info().
			Int("objects", verifyResult.SourceObjectCount).
			Dur("duration", result.Duration).
			Msgf("✅ %s passed: all %d object(s) verified in target", testLabel, verifyResult.SourceObjectCount)
	} else {
		result.Error = fmt.Sprintf("%d missing in target, %d extra in target", len(missing), len(extra))
		log.Error().
			Int("source", verifyResult.SourceObjectCount).
			Int("target", verifyResult.TargetObjectCount).
			Int("missing", len(missing)).
			Int("extra", len(extra)).
			Msgf("❌ %s failed: data mismatch between source and target", testLabel)
	}

	return verifyResult, result
}

// runDeleteOSStep deletes an Object Storage in three phases:
//  1. Call DELETE with option=empty exactly once — empties the bucket.
//  2. Call DELETE (no option) up to maxRetries — standard deletion of the now-empty bucket.
//  3. Fallback: if phase 2 fails entirely, call DELETE with option=force up to maxRetries.
//
// This always runs as cleanup regardless of whether earlier steps failed.
func runDeleteOSStep(client *resty.Client, cfg TestConfig, osId, name, stepLabel, _ string) TestResults {
	testLabel := stepLabel
	log.Info().Msgf("\n--- %s (%s, osId=%s) ---", testLabel, name, osId)

	result := TestResults{
		TestName:  fmt.Sprintf("%s (%s)", testLabel, name),
		StartTime: time.Now(),
	}

	baseURL := fmt.Sprintf("%s/beetle/migration/middleware/ns/%s/objectStorage/%s",
		cfg.Beetle.Endpoint, cfg.Beetle.NamespaceID, osId)

	// doDelete performs a single DELETE call with the given option ("" = no query param).
	// Returns (success, statusCode).
	doDelete := func(opt string) (bool, int) {
		url := baseURL
		if opt != "" {
			url += "?option=" + opt
		}
		resp, err := client.R().Delete(url)
		code := resp.StatusCode()
		if err == nil && (code < 400 || code == 404) {
			return true, code
		}
		log.Warn().Err(err).Int("statusCode", code).Str("option", opt).
			Msgf("DELETE failed for '%s'", osId)
		return false, code
	}

	const maxRetries = 3
	const retryDelay = 10 * time.Second

	// Phase 1: option=empty, exactly once.
	log.Info().Str("osId", osId).Msg("Phase 1: DELETE ?option=empty (once)")
	doDelete("empty") // result ignored; proceed to standard delete regardless

	// Phase 2: standard DELETE (no option), up to maxRetries.
	log.Info().Str("osId", osId).Msg("Phase 2: DELETE (standard)")
	deleted := false
	for attempt := 1; attempt <= maxRetries; attempt++ {
		if attempt > 1 {
			log.Info().Str("osId", osId).Int("attempt", attempt).Msg("Retrying standard deletion...")
			time.Sleep(retryDelay)
		}
		if ok, code := doDelete(""); ok {
			deleted = true
			result.StatusCode = code
			result.RequestURL = baseURL
			break
		}
	}

	// Phase 3: fallback to option=force, up to maxRetries.
	if !deleted {
		log.Warn().Str("osId", osId).Msg("Phase 2 failed — fallback: DELETE ?option=force")
		for attempt := 1; attempt <= maxRetries; attempt++ {
			if attempt > 1 {
				log.Info().Str("osId", osId).Int("attempt", attempt).Msg("Retrying force deletion...")
				time.Sleep(retryDelay)
			}
			if ok, code := doDelete("force"); ok {
				deleted = true
				result.StatusCode = code
				result.RequestURL = baseURL + "?option=force"
				break
			}
		}
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	if deleted {
		result.Success = true
		log.Info().Msgf("✅ %s passed: deleted '%s'", testLabel, osId)
	} else {
		result.Error = fmt.Sprintf("Failed to delete '%s': all phases exhausted (empty → standard → force)", osId)
		result.StatusCode = 500
		log.Error().Msgf("❌ %s failed: cleanup incomplete for '%s'", testLabel, osId)
	}
	return result
}

// ============================================================================
// Dummy Data Generation
// ============================================================================

// generateDummyData creates test files in the configured base directory.
// Generates files of various sizes and types to simulate on-premise data.
func generateDummyData(cfg TestConfig) error {
	baseDir := cfg.TestData.BaseDir
	if baseDir == "" {
		baseDir = "./dummydata"
	}

	// Check if dummy data already exists.
	if info, err := os.Stat(baseDir); err == nil && info.IsDir() {
		entries, _ := os.ReadDir(baseDir)
		if len(entries) > 0 {
			log.Info().Str("dir", baseDir).Msg("Dummy data already exists, skipping generation")
			return nil
		}
	}

	log.Info().Str("dir", baseDir).Msg("Generating dummy test data...")
	startTime := time.Now()

	type fileSpec struct {
		subdir string
		count  int
		minKB  int
		maxKB  int
		ext    string
	}

	specs := []fileSpec{
		{subdir: "small", count: 10, minKB: 1, maxKB: 10, ext: ".txt"},
		{subdir: "small", count: 5, minKB: 10, maxKB: 50, ext: ".md"},
		{subdir: "medium", count: 5, minKB: 100, maxKB: 500, ext: ".json"},
		{subdir: "medium", count: 3, minKB: 500, maxKB: 1024, ext: ".csv"},
		{subdir: "large", count: 2, minKB: 2048, maxKB: 5120, ext: ".bin"},
		{subdir: "docs", count: 3, minKB: 5, maxKB: 50, ext: ".txt"},
		{subdir: "docs", count: 2, minKB: 10, maxKB: 100, ext: ".md"},
		{subdir: "nested/level1/level2", count: 2, minKB: 1, maxKB: 10, ext: ".txt"},
		{subdir: "nested/level1/level3", count: 2, minKB: 1, maxKB: 10, ext: ".txt"},
	}

	totalFiles := 0
	for _, spec := range specs {
		dir := filepath.Join(baseDir, spec.subdir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		for i := 0; i < spec.count; i++ {
			sizeKB := spec.minKB + rand.Intn(spec.maxKB-spec.minKB+1)
			content := make([]byte, sizeKB*1024)
			rand.Read(content) //nolint:gosec // dummy data generation, not security-sensitive

			// For text-based files, use readable ASCII content.
			if spec.ext == ".txt" || spec.ext == ".md" || spec.ext == ".json" || spec.ext == ".csv" {
				content = generateTextContent(sizeKB, spec.ext)
			}

			fileName := fmt.Sprintf("file%03d%s", i+1, spec.ext)
			filePath := filepath.Join(dir, fileName)

			if err := os.WriteFile(filePath, content, 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %w", filePath, err)
			}
			totalFiles++
		}
	}

	log.Info().Int("files", totalFiles).Str("dir", baseDir).
		Dur("elapsed", time.Since(startTime)).Msg("Dummy data generated")
	return nil
}

// generateTextContent creates readable text content for a given size and extension.
func generateTextContent(sizeKB int, ext string) []byte {
	const lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. " +
		"Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. " +
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris.\n"

	targetSize := sizeKB * 1024
	var buf []byte
	switch ext {
	case ".json":
		entry := `{"id":%d,"name":"item-%d","value":%d,"active":true}` + "\n"
		for len(buf) < targetSize {
			n := len(buf)/len(entry) + 1
			buf = append(buf, []byte(fmt.Sprintf(entry, n, n, rand.Intn(1000)))...)
		}
	case ".csv":
		buf = append(buf, []byte("id,name,value,timestamp\n")...)
		row := "item-%d,data-%d,%d,%s\n"
		for len(buf) < targetSize {
			n := len(buf)/30 + 1
			buf = append(buf, []byte(fmt.Sprintf(row, n, n, rand.Intn(1000), time.Now().Format("2006-01-02")))...)
		}
	default:
		for len(buf) < targetSize {
			buf = append(buf, []byte(lorem)...)
		}
	}
	return buf[:targetSize]
}

// ============================================================================
// Helper Functions
// ============================================================================

// buildTumblebugAuth constructs the auth request for Tumblebug access.
func buildTumblebugAuth(cfg TestConfig, authConfig AuthConfig) *transx.AuthConfig {
	username := cfg.Tumblebug.Auth.Username
	password := cfg.Tumblebug.Auth.Password

	// Fall back to beetle auth if tumblebug auth is not configured.
	if username == "" && authConfig.BasicAuthUsername != "" {
		username = authConfig.BasicAuthUsername
		password = authConfig.BasicAuthPassword
	}
	if username == "" {
		return nil
	}
	return &transx.AuthConfig{
		AuthType: transx.AuthTypeBasic,
		Basic:    &transx.BasicAuthConfig{Username: username, Password: password},
	}
}

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
		Error:     "skipped due to previous step failure",
	}
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
}

// ============================================================================
// Markdown Report Generation
// ============================================================================

// generateMarkdownReport writes a per-target test result to testresult/.
func generateMarkdownReport(report *DataTestReport) error {
	execDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	testResultDir := filepath.Join(execDir, "testresult")
	if err := os.MkdirAll(testResultDir, 0755); err != nil {
		return fmt.Errorf("failed to create testresult directory: %w", err)
	}

	filename := filepath.Join(testResultDir,
		fmt.Sprintf("data-test-results-%s-to-%s.md",
			strings.ToLower(report.SourceCSP),
			strings.ToLower(report.CSP)))

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

// maskSensitiveInfo redacts sensitive data from report content.
func maskSensitiveInfo(content string) string {
	// 1. Mask Azure Subscription IDs.
	reSub := regexp.MustCompile(`(?i)/subscriptions/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)
	content = reSub.ReplaceAllString(content, "/subscriptions/AZURE_SUBSCRIPTION_ID")

	// 2. Mask GCP Project IDs.
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

	// 3. Mask email addresses.
	reEmail := regexp.MustCompile(`[a-zA-Z0-9+_.\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}`)
	content = reEmail.ReplaceAllString(content, "MASKED_EMAIL")

	return content
}

// generateMarkdownContent builds the full markdown report string.
func generateMarkdownContent(report *DataTestReport) string {
	var sb strings.Builder

	// ========================================================================
	// Header
	// ========================================================================
	sb.WriteString(fmt.Sprintf("# CM-Beetle test results for %s %s (data migration)\n\n",
		strings.ToUpper(report.CSP), report.Region))
	sb.WriteString("> [!NOTE]\n")
	sb.WriteString(fmt.Sprintf("> This document presents test results for CM-Beetle data migration to %s (%s).\n\n",
		strings.ToUpper(report.CSP), report.Region))

	// ========================================================================
	// Environment and scenario
	// ========================================================================
	sb.WriteString("## Environment and scenario\n\n")
	sb.WriteString("### Environment\n\n")
	sb.WriteString(fmt.Sprintf("- CM-Beetle: %s\n", getBeetleVersion()))
	sb.WriteString(fmt.Sprintf("- CB-Tumblebug: v%s\n", getVersionFromDockerCompose("cb-tumblebug")))
	sb.WriteString(fmt.Sprintf("- Source CSP: %s\n", strings.ToUpper(report.SourceCSP)))
	sb.WriteString(fmt.Sprintf("- Source Region: %s\n", report.SourceRegion))
	sb.WriteString(fmt.Sprintf("- Target CSP: %s\n", strings.ToUpper(report.CSP)))
	sb.WriteString(fmt.Sprintf("- Target Region: %s\n", report.Region))
	sb.WriteString(fmt.Sprintf("- Source OS ID: %s\n", report.SourceOsId))
	sb.WriteString(fmt.Sprintf("- Target OS ID: %s\n", report.TargetOsId))
	sb.WriteString(fmt.Sprintf("- CM-Beetle URL: %s\n", report.BeetleURL))
	sb.WriteString(fmt.Sprintf("- Namespace: %s\n", report.NamespaceID))
	sb.WriteString(fmt.Sprintf("- Name Seed: %s\n", report.NameSeed))
	sb.WriteString("- Test CLI: CM-Beetle data migration automated test CLI\n")
	sb.WriteString(fmt.Sprintf("- Test Date: %s\n", report.TestDate))
	sb.WriteString(fmt.Sprintf("- Test Time: %s\n", report.TestTime))
	sb.WriteString(fmt.Sprintf("- Test Execution: %s\n\n", report.TestDateTime.Format("2006-01-02 15:04:05 MST")))

	sb.WriteString("### Scenario\n\n")
	sb.WriteString("**Pre-flight (shared, runs once before all target tests):**\n\n")
	sb.WriteString("1. Create source object storage via CM-Beetle\n")
	sb.WriteString("1. Upload dummy data to source object storage (local → source OS, encrypted, with filter, async)\n\n")
	sb.WriteString("**Per target CSP-Region (Steps 1–4 below):**\n\n")
	sb.WriteString("1. Create target object storage via CM-Beetle\n")
	sb.WriteString("1. Migrate data: source OS → target OS (encrypted, async)\n")
	sb.WriteString("1. Verify migrated data: compare source and target object lists\n")
	sb.WriteString("1. Delete target object storage (cleanup) via CM-Beetle\n\n")
	sb.WriteString("**Post-flight (shared, runs once after all target tests):**\n\n")
	sb.WriteString("1. Delete source object storage (cleanup) via CM-Beetle\n\n")

	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> Some long request/response bodies are in the collapsible section for better readability.\n\n")

	// ========================================================================
	// Test Results Summary
	// ========================================================================
	sb.WriteString(fmt.Sprintf("## Test result for %s %s\n\n", strings.ToUpper(report.CSP), report.Region))
	sb.WriteString("### Test Results Summary\n\n")
	sb.WriteString("| Step | Endpoint / Description | Status | Duration | Details |\n")
	sb.WriteString("|------|------------------------|--------|----------|---------|\n")

	endpoints := []string{
		fmt.Sprintf("`POST /beetle/migration/middleware/ns/%s/objectStorage` (target)", report.NamespaceID),
		"`POST /beetle/migration/data` (migrate: source OS → target OS, encrypted, async)",
		"Verify migrated data (compare source and target object lists)",
		fmt.Sprintf("`DELETE /beetle/migration/middleware/ns/%s/objectStorage/{targetOsId}` (cleanup)", report.NamespaceID),
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
		sb.WriteString(fmt.Sprintf("**Overall Result**: %d/%d steps passed, %d skipped", passedCount, totalTests, skippedCount))
	} else {
		sb.WriteString(fmt.Sprintf("**Overall Result**: %d/%d steps passed", passedCount, totalTests))
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
	sb.WriteString("## Detailed Test Results\n\n")
	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> This section provides detailed information for each test step, including API request and response details.\n\n")

	// Step 1: Create target OS
	sb.WriteString("### Step 1: Create target object storage\n\n")
	sb.WriteString("#### 1.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `POST /beetle/migration/middleware/ns/%s/objectStorage`\n", report.NamespaceID))
	sb.WriteString("- **Purpose**: Create target object storage bucket directly from JSON spec\n")
	sb.WriteString(fmt.Sprintf("- **Target CSP**: `%s`\n", report.CSP))
	sb.WriteString(fmt.Sprintf("- **Target Region**: `%s`\n", report.Region))
	sb.WriteString(fmt.Sprintf("- **Target OS ID**: `%s`\n", report.TargetOsId))
	sb.WriteString(fmt.Sprintf("- **Source OS ID**: `%s` (shared pre-flight)\n\n", report.SourceOsId))
	if len(report.TestResults) > 0 && report.TestResults[0].RequestBody != nil {
		sb.WriteString(detailsBlock("Request Body", report.TestResults[0].RequestBody))
	}
	sb.WriteString("#### 1.2 API Response Information\n\n")
	if len(report.TestResults) > 0 {
		r := report.TestResults[0]
		if r.Skipped {
			sb.WriteString("- **Status**: ⏭️ **SKIPPED** (previous step failed)\n\n")
		} else if r.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS** (HTTP 201 Created)\n\n")
			if r.ResponseBody != nil {
				sb.WriteString(detailsBlock("Response Body", r.ResponseBody))
			}
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if r.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n", r.Error))
			}
			if r.ResponseBody != nil {
				sb.WriteString(detailsBlock("Response Body", r.ResponseBody))
			} else {
				sb.WriteString("\n")
			}
		}
	}

	// Step 2: Migrate data
	sb.WriteString("### Step 2: Migrate data (source OS → target OS, encrypted)\n\n")
	sb.WriteString("#### 2.1 API Request Information\n\n")
	sb.WriteString("- **API Endpoint**: `POST /beetle/migration/data`\n")
	sb.WriteString("- **Purpose**: Migrate data from source object storage to target object storage\n")
	sb.WriteString("- **Transfer Direction**: Source Object Storage (Tumblebug) → Target Object Storage (Tumblebug)\n")
	sb.WriteString("- **Encryption**: Tumblebug auth credentials encrypted with one-time RSA public key\n")
	sb.WriteString("- **Strategy**: auto\n\n")
	if len(report.TestResults) > 1 && report.TestResults[1].RequestBody != nil {
		sb.WriteString(detailsBlock("Request Body (sanitized — passwords masked)", report.TestResults[1].RequestBody))
	}
	sb.WriteString("#### 2.2 API Response Information\n\n")
	if len(report.TestResults) > 1 {
		r := report.TestResults[1]
		if r.Skipped {
			sb.WriteString("- **Status**: ⏭️ **SKIPPED** (previous step failed)\n\n")
		} else if r.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString(fmt.Sprintf("- **Duration**: %v\n", r.Duration.Truncate(time.Millisecond)))
			if r.PollNote != "" {
				sb.WriteString(fmt.Sprintf("- **Note**: %s\n", r.PollNote))
			}
			if r.ResponseBody != nil {
				sb.WriteString(detailsBlock("Initial Response Body (202 Accepted)", r.ResponseBody))
			} else {
				sb.WriteString("\n")
			}
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if r.PollNote != "" {
				sb.WriteString(fmt.Sprintf("- **Note**: %s\n", r.PollNote))
			}
			if r.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n", r.Error))
			}
			if r.ResponseBody != nil {
				sb.WriteString(detailsBlock("Response Body", r.ResponseBody))
			} else {
				sb.WriteString("\n")
			}
		}
	}

	// Step 3: Verify migrated data
	sb.WriteString("### Step 3: Verify migrated data\n\n")
	sb.WriteString("#### 3.1 Action\n\n")
	sb.WriteString("- **Purpose**: Compare source and target bucket object lists to confirm migration completeness\n")
	sb.WriteString("- **Method**: List all objects in source and target via CM-Beetle API, then compare keys\n\n")
	if len(report.TestResults) > 2 && report.TestResults[2].ResponseBody != nil {
		sb.WriteString(detailsBlock("Object Lists (source and target)", report.TestResults[2].ResponseBody))
	}
	if report.VerifyResult != nil && len(report.VerifyResult.SrcKeys) > 0 {
		sb.WriteString(treeBlock(
			fmt.Sprintf("Source data tree (%d objects)", len(report.VerifyResult.SrcKeys)),
			buildObjectTree(report.VerifyResult.SrcKeys),
		))
	}
	if report.VerifyResult != nil && len(report.VerifyResult.DstKeys) > 0 {
		sb.WriteString(treeBlock(
			fmt.Sprintf("Target data tree (%d objects)", len(report.VerifyResult.DstKeys)),
			buildObjectTree(report.VerifyResult.DstKeys),
		))
	}
	sb.WriteString("#### 3.2 Result\n\n")
	if len(report.TestResults) > 2 {
		r := report.TestResults[2]
		if r.Skipped {
			sb.WriteString("- **Status**: ⏭️ **SKIPPED** (migration step failed)\n\n")
		} else if r.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			if report.VerifyResult != nil {
				vr := report.VerifyResult
				sb.WriteString(fmt.Sprintf("- **Source objects**: %d\n", vr.SourceObjectCount))
				sb.WriteString(fmt.Sprintf("- **Target objects**: %d\n", vr.TargetObjectCount))
				sb.WriteString(fmt.Sprintf("- **Matched**: %d/%d ✅\n", vr.MatchedCount, vr.SourceObjectCount))
			}
			sb.WriteString(fmt.Sprintf("- **Duration**: %v\n\n", r.Duration.Truncate(time.Millisecond)))
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if report.VerifyResult != nil {
				vr := report.VerifyResult
				sb.WriteString(fmt.Sprintf("- **Source objects**: %d\n", vr.SourceObjectCount))
				sb.WriteString(fmt.Sprintf("- **Target objects**: %d\n", vr.TargetObjectCount))
				sb.WriteString(fmt.Sprintf("- **Matched**: %d/%d\n", vr.MatchedCount, vr.SourceObjectCount))
				if len(vr.MissingInTarget) > 0 {
					sb.WriteString(fmt.Sprintf("- **Missing in target** (%d):\n", len(vr.MissingInTarget)))
					for _, k := range vr.MissingInTarget {
						sb.WriteString(fmt.Sprintf("  - `%s`\n", k))
					}
				}
				if len(vr.ExtraInTarget) > 0 {
					sb.WriteString(fmt.Sprintf("- **Extra in target** (%d):\n", len(vr.ExtraInTarget)))
					for _, k := range vr.ExtraInTarget {
						sb.WriteString(fmt.Sprintf("  - `%s`\n", k))
					}
				}
			}
			if r.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n", r.Error))
			}
			sb.WriteString("\n")
		}
	}

	// Step 4: Delete target OS
	sb.WriteString("### Step 4: Delete target object storage (cleanup)\n\n")
	sb.WriteString("#### 4.1 API Request Information\n\n")
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `DELETE /beetle/migration/middleware/ns/%s/objectStorage/%s`\n",
		report.NamespaceID, report.TargetOsId))
	sb.WriteString("- **Purpose**: Delete target object storage bucket as cleanup\n")
	sb.WriteString("- **Note**: Always runs regardless of previous step failures\n\n")
	sb.WriteString("#### 4.2 API Response Information\n\n")
	if len(report.TestResults) > 3 {
		r := report.TestResults[3]
		if r.Skipped {
			sb.WriteString("- **Status**: ⏭️ **SKIPPED** (no target OS to delete)\n\n")
		} else if r.Success {
			sb.WriteString("- **Status**: ✅ **SUCCESS**\n")
			sb.WriteString("- **Response**: Target object storage deleted successfully\n\n")
		} else {
			sb.WriteString("- **Status**: ❌ **FAILED**\n")
			if r.Error != "" {
				sb.WriteString(fmt.Sprintf("- **Error**: %s\n\n", r.Error))
			}
		}
	}

	return sb.String()
}
func detailsBlock(summary string, v interface{}) string {
	if v == nil {
		return ""
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return fmt.Sprintf("<details>\n<summary>%s</summary>\n\n```json\n%s\n```\n\n</details>\n\n", summary, string(b))
}

// treeBlock wraps pre-formatted tree text in a collapsible HTML details block with a plain code fence.
func treeBlock(summary, content string) string {
	return fmt.Sprintf("<details>\n<summary>%s</summary>\n\n```\n%s```\n\n</details>\n\n", summary, content)
}

// buildObjectTree renders a list of object keys as an ASCII tree (like the `tree` command).
func buildObjectTree(keys []string) string {
	if len(keys) == 0 {
		return "(empty)\n"
	}

	sorted := make([]string, len(keys))
	copy(sorted, keys)
	sort.Strings(sorted)

	type treeNode struct {
		children map[string]*treeNode
	}
	newNode := func() *treeNode { return &treeNode{children: make(map[string]*treeNode)} }

	root := newNode()
	for _, key := range sorted {
		node := root
		for _, part := range strings.Split(key, "/") {
			if _, ok := node.children[part]; !ok {
				node.children[part] = newNode()
			}
			node = node.children[part]
		}
	}

	var sb strings.Builder
	sb.WriteString(".\n")

	var render func(node *treeNode, prefix string)
	render = func(node *treeNode, prefix string) {
		names := make([]string, 0, len(node.children))
		for name := range node.children {
			names = append(names, name)
		}
		sort.Strings(names)
		for i, name := range names {
			isLast := i == len(names)-1
			connector := "├── "
			childPrefix := prefix + "│   "
			if isLast {
				connector = "└── "
				childPrefix = prefix + "    "
			}
			sb.WriteString(prefix + connector + name + "\n")
			render(node.children[name], childPrefix)
		}
	}
	render(root, "")
	return sb.String()
}

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
// Config & Auth Loaders
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

func loadOsCreationReq(path string) (OsCreationReq, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return OsCreationReq{}, fmt.Errorf("read %s: %w", path, err)
	}
	var req OsCreationReq
	if err := json.Unmarshal(data, &req); err != nil {
		return OsCreationReq{}, fmt.Errorf("parse %s: %w", path, err)
	}
	return req, nil
}
