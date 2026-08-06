// Package main is the starting point of CM-Beetle Multi-Infra Recommendation Test CLI.
//
// Scoped to exactly one thing: sending a real multi-target recommendation request to
// POST /recommendation/multiInfra and /recommendation/multiInfraWithNlb, and verifying the
// response accurately reflects the request — one result item per requested target, in request
// order. Modeled on cmd/test-cli/infra-with-nlb's structure, but without its migration/
// validation/SSH/cleanup steps, since neither multi-target endpoint provisions anything.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
)

// TestConfig holds test configuration.
type TestConfig struct {
	Test struct {
		Cases []TestCase `yaml:"cases"`
	} `yaml:"test"`
	Beetle struct {
		Endpoint               string `yaml:"endpoint"`
		RequestBodyFile        string `yaml:"requestBodyFile"`
		RequestBodyFileWithNlb string `yaml:"requestBodyFileWithNlb"`
		AuthConfigFile         string `yaml:"authConfigFile"`
	} `yaml:"beetle"`
}

// TestCase is one candidate target; only entries with Execute: true are used.
type TestCase struct {
	cloudmodel.CloudProperty `yaml:",inline"`
	Name                     string `yaml:"name"`
	Execute                  bool   `yaml:"execute"`
}

// AuthConfig holds Beetle API credentials.
type AuthConfig struct {
	BeetleApiUsername string `json:"beetleApiUsername"`
	BeetleApiPassword string `json:"beetleApiPassword"`
}

// TestResults holds one API call's execution result.
type TestResults struct {
	TestName     string        `json:"testName"`
	StartTime    time.Time     `json:"startTime"`
	EndTime      time.Time     `json:"endTime"`
	Duration     time.Duration `json:"duration"`
	Success      bool          `json:"success"`
	StatusCode   int           `json:"statusCode"`
	ErrorMessage string        `json:"errorMessage,omitempty"`
	RequestURL   string        `json:"requestUrl,omitempty"`
}

// AccuracyCheck records whether a response's targetCloud order/coverage matched the request.
type AccuracyCheck struct {
	Passed  bool
	Details []string
}

// MultiInfraTestReport captures request/response and accuracy checks for both endpoints,
// making it possible to write out a single markdown report at the end.
type MultiInfraTestReport struct {
	TestDate     string
	TestTime     string
	TestDateTime time.Time
	BeetleURL    string
	Targets      []cloudmodel.CloudProperty

	Request  controller.RecommendMultiInfraRequest
	Response *model.ApiResponse[[]cloudmodel.RecommendedInfra]
	Result   TestResults
	Accuracy AccuracyCheck

	RequestWithNlb  controller.RecommendMultiInfraWithNlbRequest
	ResponseWithNlb *model.ApiResponse[[]cloudmodel.RecommendedInfra]
	ResultWithNlb   TestResults
	AccuracyWithNlb AccuracyCheck
}

var configFile = flag.String("config", "testconf/test-config.yaml", "Path to config file")

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

func main() {
	flag.Parse()

	cfg, err := loadConfig(*configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	sourceInfra, _, err := loadSourceInfraModel(cfg.Beetle.RequestBodyFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load source infra model")
	}
	sourceInfraWithNlb, _, err := loadSourceInfraModel(cfg.Beetle.RequestBodyFileWithNlb)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load NLB-aware source infra model")
	}

	var targets []cloudmodel.CloudProperty
	for _, c := range cfg.Test.Cases {
		if c.Execute {
			targets = append(targets, c.CloudProperty)
		}
	}
	if len(targets) < 2 {
		log.Fatal().Msg("At least 2 test cases must have execute: true")
	}
	if len(targets) > 10 {
		targets = targets[:10]
	}

	client := resty.New()
	client.SetTimeout(2 * time.Minute)

	if err := checkBeetleReadiness(client, cfg.Beetle.Endpoint); err != nil {
		log.Fatal().Err(err).Msg("CM-Beetle readiness check failed")
	}

	authConfig, err := loadAuthConfig(cfg.Beetle.AuthConfigFile)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load auth config; proceeding without auth")
	} else if authConfig.BeetleApiUsername != "" {
		client.SetBasicAuth(authConfig.BeetleApiUsername, authConfig.BeetleApiPassword)
	}

	now := time.Now()
	report := &MultiInfraTestReport{
		TestDate:     now.Format("2006-01-02"),
		TestTime:     now.Format("15:04:05"),
		TestDateTime: now,
		BeetleURL:    cfg.Beetle.Endpoint,
		Targets:      targets,
	}

	fmt.Println("\n--- Test 1: POST /beetle/recommendation/multiInfra ---")
	report.Request = controller.RecommendMultiInfraRequest{
		DesiredCspAndRegionPairs: targets,
		SourceInfra:              sourceInfra,
	}
	report.Response, report.Result = runMultiInfraRecommendationTest(
		client, cfg.Beetle.Endpoint, "/beetle/recommendation/multiInfra", report.Request)
	if report.Result.Success {
		report.Accuracy = checkInputOutputAccuracy(targets, report.Response)
		printAccuracy("Test 1", report.Accuracy)
	}

	fmt.Println("\n--- Test 2: POST /beetle/recommendation/multiInfraWithNlb ---")
	report.RequestWithNlb = controller.RecommendMultiInfraWithNlbRequest{
		DesiredCspAndRegionPairs: targets,
		SourceInfra:              sourceInfraWithNlb,
	}
	report.ResponseWithNlb, report.ResultWithNlb = runMultiInfraRecommendationTest(
		client, cfg.Beetle.Endpoint, "/beetle/recommendation/multiInfraWithNlb",
		controller.RecommendMultiInfraRequest(report.RequestWithNlb))
	if report.ResultWithNlb.Success {
		report.AccuracyWithNlb = checkInputOutputAccuracy(targets, report.ResponseWithNlb)
		printAccuracy("Test 2", report.AccuracyWithNlb)
	}

	if err := generateMarkdownReport(report); err != nil {
		log.Warn().Err(err).Msg("Failed to generate markdown report")
	}

	allPassed := report.Result.Success && report.Accuracy.Passed &&
		report.ResultWithNlb.Success && report.AccuracyWithNlb.Passed

	fmt.Println("\n=========================================================")
	if allPassed {
		fmt.Println(" ✅ All recommendation input/output checks passed.")
	} else {
		fmt.Println(" ❌ Some recommendation input/output checks failed. See testresult/ for details.")
	}
	fmt.Println("=========================================================")

	if !allPassed {
		os.Exit(1)
	}
}

// runMultiInfraRecommendationTest sends the request body as-is to the given path and returns
// the parsed response alongside execution metadata; both request and response are expected to
// use the shared cloudmodel.RecommendedInfra shape regardless of which endpoint is called.
func runMultiInfraRecommendationTest(client *resty.Client, beetleURL, path string, reqBody controller.RecommendMultiInfraRequest) (*model.ApiResponse[[]cloudmodel.RecommendedInfra], TestResults) {
	url := beetleURL + path
	result := TestResults{TestName: fmt.Sprintf("POST %s", path), StartTime: time.Now(), RequestURL: url}

	var apiResponse model.ApiResponse[[]cloudmodel.RecommendedInfra]
	err := common.ExecuteHttpRequest(client, "POST", url, nil, true, &reqBody, &apiResponse, 0)

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	if err != nil {
		result.ErrorMessage = err.Error()
		fmt.Printf("❌ %s failed: %s\n", result.TestName, err)
		return nil, result
	}
	if !apiResponse.Success {
		result.ErrorMessage = apiResponse.Error
		fmt.Printf("❌ %s failed: %s\n", result.TestName, apiResponse.Error)
		return &apiResponse, result
	}

	result.Success = true
	result.StatusCode = 200
	fmt.Printf("✅ %s passed (%d target(s), %v)\n", result.TestName, len(apiResponse.Data), result.Duration.Truncate(time.Millisecond))
	return &apiResponse, result
}

// checkInputOutputAccuracy verifies the response has exactly one item per requested target,
// in request order, matched via targetCloud — the contract both endpoints must uphold.
func checkInputOutputAccuracy(targets []cloudmodel.CloudProperty, resp *model.ApiResponse[[]cloudmodel.RecommendedInfra]) AccuracyCheck {
	check := AccuracyCheck{Passed: true}

	if resp == nil {
		check.Passed = false
		check.Details = append(check.Details, "no response to check")
		return check
	}
	if len(resp.Data) != len(targets) {
		check.Passed = false
		check.Details = append(check.Details, fmt.Sprintf("expected %d result item(s), got %d", len(targets), len(resp.Data)))
		return check
	}
	for i, want := range targets {
		got := resp.Data[i].TargetCloud
		if strings.EqualFold(got.Csp, want.Csp) && strings.EqualFold(got.Region, want.Region) {
			check.Details = append(check.Details, fmt.Sprintf("result[%d] targetCloud=%s/%s status=%s — matches request", i, got.Csp, got.Region, resp.Data[i].Status))
		} else {
			check.Passed = false
			check.Details = append(check.Details, fmt.Sprintf("result[%d] targetCloud=%s/%s, want %s/%s", i, got.Csp, got.Region, want.Csp, want.Region))
		}
	}
	return check
}

func printAccuracy(label string, check AccuracyCheck) {
	status := "✅"
	if !check.Passed {
		status = "❌"
	}
	fmt.Printf("%s %s input/output accuracy check\n", status, label)
	for _, d := range check.Details {
		fmt.Printf("      - %s\n", d)
	}
}

// checkBeetleReadiness checks if CM-Beetle is ready using GET /beetle/readyz.
func checkBeetleReadiness(client *resty.Client, beetleURL string) error {
	fmt.Println("🔍 Checking CM-Beetle readiness...")
	url := beetleURL + "/beetle/readyz"

	var response map[string]interface{}
	var emptyBody interface{} = common.NoBody
	err := common.ExecuteHttpRequest(client, "GET", url, nil, common.SetUseBody(emptyBody), &emptyBody, &response, 0)
	if err != nil {
		return fmt.Errorf("CM-Beetle readiness check failed: %w", err)
	}
	if message, ok := response["message"].(string); ok && strings.Contains(message, "NOT ready") {
		return fmt.Errorf("CM-Beetle is not ready: %s", message)
	}
	fmt.Println("✅ CM-Beetle is ready!")
	return nil
}

func loadConfig(configPath string) (TestConfig, error) {
	var cfg TestConfig
	file, err := os.Open(configPath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func loadAuthConfig(authConfigPath string) (AuthConfig, error) {
	var auth AuthConfig
	file, err := os.Open(authConfigPath)
	if err != nil {
		return auth, err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&auth); err != nil {
		return auth, fmt.Errorf("failed to decode auth config: %w", err)
	}
	return auth, nil
}

// loadSourceInfraModel reads {"nameSeed": ..., "sourceInfra": {...}} — same shape as
// cmd/test-cli/infra-with-nlb's fixture files.
func loadSourceInfraModel(requestBodyPath string) (onpremmodel.OnpremInfra, string, error) {
	var tempRequest struct {
		NameSeed    string                  `json:"nameSeed"`
		SourceInfra onpremmodel.OnpremInfra `json:"sourceInfra"`
	}
	file, err := os.Open(requestBodyPath)
	if err != nil {
		return tempRequest.SourceInfra, tempRequest.NameSeed, fmt.Errorf("failed to open request body file: %w", err)
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&tempRequest); err != nil {
		return tempRequest.SourceInfra, tempRequest.NameSeed, fmt.Errorf("failed to decode source infra model: %w", err)
	}
	return tempRequest.SourceInfra, tempRequest.NameSeed, nil
}

// generateMarkdownReport writes the request/response bodies and accuracy checks for both
// endpoints to testresult/multi-infra-recommendation-test-report.md.
func generateMarkdownReport(report *MultiInfraTestReport) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	testResultDir := filepath.Join(cwd, "testresult")
	if err := os.MkdirAll(testResultDir, 0755); err != nil {
		return err
	}

	content := generateMarkdownContent(report)
	path := filepath.Join(testResultDir, "multi-infra-recommendation-test-report.md")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}
	log.Info().Str("file", path).Msg("✅ Test report generated and saved")
	return nil
}

func generateMarkdownContent(report *MultiInfraTestReport) string {
	var sb strings.Builder

	sb.WriteString("# CM-Beetle Multi-Infra Recommendation Test Results\n\n")
	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> Verifies that `POST /recommendation/multiInfra` and `POST /recommendation/multiInfraWithNlb`\n")
	sb.WriteString("> return exactly one result item per requested target, in request order.\n\n")

	sb.WriteString("## Environment\n\n")
	sb.WriteString(fmt.Sprintf("- CM-Beetle URL: %s\n", report.BeetleURL))
	sb.WriteString(fmt.Sprintf("- Test Date: %s %s\n", report.TestDate, report.TestTime))
	sb.WriteString(fmt.Sprintf("- Target Count: %d\n", len(report.Targets)))
	var targetList []string
	for _, t := range report.Targets {
		targetList = append(targetList, fmt.Sprintf("%s/%s", t.Csp, t.Region))
	}
	sb.WriteString(fmt.Sprintf("- Targets: %s\n\n", strings.Join(targetList, ", ")))

	writeSummarySection(&sb, report)

	writeEndpointSection(&sb, 1, "Test 1: Recommend multiple target infrastructures",
		"/beetle/recommendation/multiInfra", "Get one best-match cloud infrastructure recommendation per requested CSP/region pair",
		report.Request, report.Response, report.Result, report.Accuracy)
	writeEndpointSection(&sb, 2, "Test 2: Recommend multiple target infrastructures (with NLB)",
		"/beetle/recommendation/multiInfraWithNlb", "Get one best-match, NLB-aware cloud infrastructure recommendation per requested CSP/region pair",
		report.RequestWithNlb, report.ResponseWithNlb, report.ResultWithNlb, report.AccuracyWithNlb)

	return sb.String()
}

// writeSummarySection writes an overview table of both endpoint tests, mirroring
// cmd/test-cli/infra-with-nlb's "Test Results Summary" section.
func writeSummarySection(sb *strings.Builder, report *MultiInfraTestReport) {
	sb.WriteString("### Test Results Summary\n\n")
	sb.WriteString("| Test | Step (Endpoint / Description) | Status | Duration | Details |\n")
	sb.WriteString("|------|-------------------------------|--------|----------|----------|\n")

	rows := []struct {
		endpoint string
		result   TestResults
		accuracy AccuracyCheck
	}{
		{"/beetle/recommendation/multiInfra", report.Result, report.Accuracy},
		{"/beetle/recommendation/multiInfraWithNlb", report.ResultWithNlb, report.AccuracyWithNlb},
	}

	passed := 0
	var totalDuration time.Duration
	for i, r := range rows {
		status := "✅ **PASS**"
		details := "Pass"
		if !r.result.Success {
			status = "❌ **FAIL**"
			details = r.result.ErrorMessage
		} else if !r.accuracy.Passed {
			status = "❌ **FAIL**"
			details = "Input/output accuracy check failed"
		} else {
			passed++
		}
		sb.WriteString(fmt.Sprintf("| %d | `POST %s` | %s | %v | %s |\n",
			i+1, r.endpoint, status, r.result.Duration.Truncate(time.Millisecond), details))
		totalDuration += r.result.Duration
	}

	sb.WriteString(fmt.Sprintf("\n**Overall Result**: %d/%d tests passed", passed, len(rows)))
	if passed == len(rows) {
		sb.WriteString(" ✅\n\n")
	} else {
		sb.WriteString(" ❌\n\n")
	}
	sb.WriteString(fmt.Sprintf("**Total Duration**: %v\n\n", totalDuration))
	sb.WriteString(fmt.Sprintf("*Test executed on %s at %s using CM-Beetle automated test CLI*\n\n---\n\n",
		report.TestDateTime.Format("January 2, 2006"), report.TestDateTime.Format("15:04:05 MST")))
}

func writeEndpointSection(sb *strings.Builder, testNum int, title, endpoint, purpose string, reqBody interface{}, resp *model.ApiResponse[[]cloudmodel.RecommendedInfra], result TestResults, accuracy AccuracyCheck) {
	sb.WriteString(fmt.Sprintf("## %s\n\n", title))

	sb.WriteString(fmt.Sprintf("#### %d.1 API Request Information\n\n", testNum))
	sb.WriteString(fmt.Sprintf("- **API Endpoint**: `POST %s`\n", endpoint))
	sb.WriteString(fmt.Sprintf("- **Purpose**: %s\n\n", purpose))

	sb.WriteString("**Request Body**:\n\n<details>\n  <summary> <ins>Click to see the request body</ins> </summary>\n\n```json\n")
	reqJSON, _ := json.MarshalIndent(reqBody, "", "  ")
	sb.WriteString(string(reqJSON))
	sb.WriteString("\n```\n\n</details>\n\n")

	sb.WriteString(fmt.Sprintf("#### %d.2 API Response Information\n\n", testNum))
	status := "✅ **PASS**"
	if !result.Success {
		status = "❌ **FAIL**"
	}
	sb.WriteString(fmt.Sprintf("- **Status**: %s\n", status))
	sb.WriteString(fmt.Sprintf("- **Duration**: %v\n", result.Duration.Truncate(time.Millisecond)))
	if result.ErrorMessage != "" {
		sb.WriteString(fmt.Sprintf("- **Error**: %s\n", result.ErrorMessage))
	}
	sb.WriteString("\n")

	sb.WriteString("**Response Body**:\n\n<details>\n  <summary> <ins>Click to see the response body</ins> </summary>\n\n```json\n")
	if resp != nil {
		respJSON, _ := json.MarshalIndent(resp, "", "  ")
		sb.WriteString(string(respJSON))
	} else {
		sb.WriteString("(no response)")
	}
	sb.WriteString("\n```\n\n</details>\n\n")

	sb.WriteString("**Input/Output Accuracy Check**:\n\n")
	accStatus := "✅ PASS"
	if !accuracy.Passed {
		accStatus = "❌ FAIL"
	}
	sb.WriteString(fmt.Sprintf("- Overall: %s\n", accStatus))
	for _, d := range accuracy.Details {
		sb.WriteString(fmt.Sprintf("- %s\n", d))
	}
	sb.WriteString("\n---\n\n")
}
