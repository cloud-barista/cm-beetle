// Package main is the starting point of CM-Beetle K8s Infra Recommendation Test CLI.
//
// Scoped to POST /recommendation/k8sCluster only: it sends each on-premise scenario fixture
// to every configured CSP/region pair and checks the recommendation against declarative
// expectations. Nothing is provisioned, so there is no migration, SSH, or cleanup step —
// the same scoping rationale as cmd/test-cli/multi-infra-recommendation.
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

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/cloud-barista/cm-beetle/pkg/logger"
)

// restyNoopLogger silences all Resty log output (e.g. "Basic Auth in HTTP mode" warnings).
type restyNoopLogger struct{}

func (restyNoopLogger) Errorf(_ string, _ ...interface{}) {}
func (restyNoopLogger) Warnf(_ string, _ ...interface{})  {}
func (restyNoopLogger) Debugf(_ string, _ ...interface{}) {}

// TestConfig holds test configuration.
type TestConfig struct {
	Test struct {
		Set struct {
			Mode string `yaml:"mode"` // parallel or sequential
		} `yaml:"set"`
		Cases     []TestCase `yaml:"cases"`
		Scenarios []Scenario `yaml:"scenarios"`
	} `yaml:"test"`
	Beetle struct {
		Endpoint       string `yaml:"endpoint"`
		AuthConfigFile string `yaml:"authConfigFile"`
	} `yaml:"beetle"`
}

// TestCase is one target CSP/region pair; only entries with Execute: true are used.
type TestCase struct {
	cloudmodel.CloudProperty `yaml:",inline"`
	Name                     string `yaml:"name"`
	Execute                  bool   `yaml:"execute"`
}

// Scenario is one on-premise fixture plus the expectations it must satisfy.
type Scenario struct {
	File    string `yaml:"file"`
	Name    string `yaml:"name"`
	Execute bool   `yaml:"execute"`
	Expect  Expect `yaml:"expect"`
}

// Expect declares what a scenario's response must satisfy. Pointer fields are optional:
// nil means "do not check". Keeping expectations in config avoids hardcoding per-CSP
// values in this CLI, which would defeat the purpose of testing the recommender.
type Expect struct {
	StatusCode     int    `yaml:"statusCode"`     // required; 200 for positive, 4xx/5xx for negative
	NodeGroupCount *int   `yaml:"nodeGroupCount"` // expected len(k8sNodeGroupList)
	TotalNodeSize  *int   `yaml:"totalNodeSize"`  // expected sum of desiredNodeSize across node groups
	VersionPrefix  string `yaml:"versionPrefix"`  // expected prefix of targetK8sCluster.version
}

// AuthConfig holds Beetle API credentials.
type AuthConfig struct {
	BeetleApiUsername string `json:"beetleApiUsername"`
	BeetleApiPassword string `json:"beetleApiPassword"`
}

// CaseResult is one (scenario × CSP/region) execution result.
type CaseResult struct {
	ScenarioName string
	ScenarioFile string
	DisplayName  string
	Csp          string
	Region       string

	StartTime time.Time
	Duration  time.Duration

	RequestURL   string
	StatusCode   int
	ResponseBody string

	Passed  bool
	Checks  []string // human-readable check outcomes, prefixed with ✅ / ❌
	Failure string   // set when the case could not even be executed
}

// TestReport aggregates every case result for the markdown report.
type TestReport struct {
	TestDateTime  time.Time
	BeetleURL     string
	BeetleVersion string
	GitHash       string
	Targets       []TestCase
	Scenarios     []Scenario
	Results       []CaseResult
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

	targets := selectedTargets(cfg)
	if len(targets) == 0 {
		log.Fatal().Msg("At least 1 test case must have execute: true")
	}
	scenarios := selectedScenarios(cfg)
	if len(scenarios) == 0 {
		log.Fatal().Msg("At least 1 scenario must have execute: true")
	}

	client := resty.New().SetTimeout(2 * time.Minute).SetLogger(restyNoopLogger{})

	if err := checkBeetleReadiness(client, cfg.Beetle.Endpoint); err != nil {
		log.Fatal().Err(err).Msg("CM-Beetle readiness check failed")
	}

	auth, err := loadAuthConfig(cfg.Beetle.AuthConfigFile)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load auth config; proceeding without auth")
	} else if auth.BeetleApiUsername != "" {
		client.SetBasicAuth(auth.BeetleApiUsername, auth.BeetleApiPassword)
	}

	report := &TestReport{
		TestDateTime:  time.Now(),
		BeetleURL:     cfg.Beetle.Endpoint,
		BeetleVersion: getBeetleVersion(),
		GitHash:       getGitHash(),
		Targets:       targets,
		Scenarios:     scenarios,
	}

	fmt.Println("=========================================================")
	fmt.Println(" CM-Beetle K8s Infra Recommendation Test CLI")
	fmt.Printf(" %d scenario(s) × %d target(s) = %d case(s), mode: %s\n",
		len(scenarios), len(targets), len(scenarios)*len(targets), executionMode(cfg))
	fmt.Println("=========================================================")

	report.Results = runAllCases(client, cfg, targets, scenarios)

	if err := generateMarkdownReport(report); err != nil {
		log.Warn().Err(err).Msg("Failed to generate markdown report")
	}

	printFinalSummary(report)

	for _, r := range report.Results {
		if !r.Passed {
			os.Exit(1)
		}
	}
}

// runAllCases executes every (target × scenario) pair, honouring the configured mode.
// Parallelism is per target so that scenarios against one CSP stay ordered and readable.
func runAllCases(client *resty.Client, cfg TestConfig, targets []TestCase, scenarios []Scenario) []CaseResult {
	if executionMode(cfg) == "sequential" {
		var all []CaseResult
		for _, t := range targets {
			all = append(all, runTargetCases(client, cfg, t, scenarios)...)
		}
		return all
	}

	// Each goroutine writes only its own slot, so results stay in target order without a mutex.
	var wg sync.WaitGroup
	perTarget := make([][]CaseResult, len(targets))
	for i, t := range targets {
		wg.Add(1)
		go func(idx int, target TestCase) {
			defer wg.Done()
			perTarget[idx] = runTargetCases(client, cfg, target, scenarios)
		}(i, t)
	}
	wg.Wait()

	var all []CaseResult
	for _, res := range perTarget {
		all = append(all, res...)
	}
	return all
}

func runTargetCases(client *resty.Client, cfg TestConfig, target TestCase, scenarios []Scenario) []CaseResult {
	results := make([]CaseResult, 0, len(scenarios))
	for _, sc := range scenarios {
		results = append(results, runCase(client, cfg, target, sc))
	}
	return results
}

// runCase posts one scenario fixture as a raw body. Sending raw bytes (rather than
// marshalling through a Go struct) is required for the negative fixtures, whose shapes do
// not map onto RecommendK8sInfraRequest at all.
func runCase(client *resty.Client, cfg TestConfig, target TestCase, sc Scenario) CaseResult {
	res := CaseResult{
		ScenarioName: sc.Name,
		ScenarioFile: sc.File,
		DisplayName:  target.Name,
		Csp:          target.Csp,
		Region:       target.Region,
		StartTime:    time.Now(),
	}

	body, err := os.ReadFile(sc.File)
	if err != nil {
		res.Duration = time.Since(res.StartTime)
		res.Failure = fmt.Sprintf("failed to read scenario file: %s", err)
		fmt.Printf("❌ [%s] %s — %s\n", target.Name, sc.Name, res.Failure)
		return res
	}

	url := cfg.Beetle.Endpoint + "/beetle/recommendation/k8sCluster"
	res.RequestURL = fmt.Sprintf("%s?desiredProvider=%s&desiredRegion=%s", url, target.Csp, target.Region)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("desiredProvider", target.Csp).
		SetQueryParam("desiredRegion", target.Region).
		SetBody(body).
		Post(url)

	res.Duration = time.Since(res.StartTime)

	if err != nil {
		res.Failure = fmt.Sprintf("request failed: %s", err)
		fmt.Printf("❌ [%s] %s — %s\n", target.Name, sc.Name, res.Failure)
		return res
	}

	res.StatusCode = resp.StatusCode()
	res.ResponseBody = string(resp.Body())

	res.Checks, res.Passed = evaluate(sc.Expect, res.StatusCode, resp.Body())

	icon := "✅"
	if !res.Passed {
		icon = "❌"
	}
	fmt.Printf("%s [%s] %-20s HTTP %d (%v)\n", icon, target.Name, sc.Name,
		res.StatusCode, res.Duration.Truncate(time.Millisecond))
	for _, c := range res.Checks {
		fmt.Printf("      %s\n", c)
	}

	return res
}

// evaluate applies the scenario's declared expectations plus the structural checks that
// every successful recommendation must satisfy.
func evaluate(exp Expect, statusCode int, body []byte) (checks []string, passed bool) {
	passed = true

	if statusCode == exp.StatusCode {
		checks = append(checks, fmt.Sprintf("✅ status code %d as expected", statusCode))
	} else {
		passed = false
		checks = append(checks, fmt.Sprintf("❌ status code %d, want %d", statusCode, exp.StatusCode))
		// Keep going: the body may still explain the mismatch in the report.
	}

	// Nothing further to assert for negative scenarios.
	if exp.StatusCode != 200 {
		return checks, passed
	}
	if statusCode != 200 {
		return checks, passed
	}

	var apiResp model.ApiResponse[cloudmodel.RecommendedInfra]
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return append(checks, fmt.Sprintf("❌ failed to parse response: %s", err)), false
	}
	if !apiResp.Success {
		return append(checks, fmt.Sprintf("❌ response success=false: %s", apiResp.Error)), false
	}

	cluster := apiResp.Data.TargetK8sCluster

	if cluster.Name == "" {
		passed = false
		checks = append(checks, "❌ targetK8sCluster.name is empty")
	}
	if cluster.Version == "" {
		passed = false
		checks = append(checks, "❌ targetK8sCluster.version is empty")
	} else {
		checks = append(checks, fmt.Sprintf("✅ version: %s", cluster.Version))
	}

	if exp.VersionPrefix != "" {
		if strings.HasPrefix(cluster.Version, exp.VersionPrefix) {
			checks = append(checks, fmt.Sprintf("✅ version has prefix %q", exp.VersionPrefix))
		} else {
			passed = false
			checks = append(checks, fmt.Sprintf("❌ version %q lacks prefix %q", cluster.Version, exp.VersionPrefix))
		}
	}

	groups := cluster.K8sNodeGroupList
	if exp.NodeGroupCount != nil {
		if len(groups) == *exp.NodeGroupCount {
			checks = append(checks, fmt.Sprintf("✅ node group count: %d", len(groups)))
		} else {
			passed = false
			checks = append(checks, fmt.Sprintf("❌ node group count %d, want %d", len(groups), *exp.NodeGroupCount))
		}
	} else {
		checks = append(checks, fmt.Sprintf("ℹ️  node group count: %d", len(groups)))
	}

	total := 0
	for i, ng := range groups {
		total += ng.DesiredNodeSize
		if ng.SpecId == "" {
			passed = false
			checks = append(checks, fmt.Sprintf("❌ node group[%d] %q has empty specId", i, ng.Name))
		}
		if ng.DesiredNodeSize < 1 {
			passed = false
			checks = append(checks, fmt.Sprintf("❌ node group[%d] %q desiredNodeSize=%d (< 1)", i, ng.Name, ng.DesiredNodeSize))
		}
		checks = append(checks, fmt.Sprintf("ℹ️  node group[%d] %q spec=%s image=%s nodes=%d",
			i, ng.Name, ng.SpecId, emptyAsDash(ng.ImageId), ng.DesiredNodeSize))
	}

	if exp.TotalNodeSize != nil {
		if total == *exp.TotalNodeSize {
			checks = append(checks, fmt.Sprintf("✅ total node size: %d", total))
		} else {
			passed = false
			checks = append(checks, fmt.Sprintf("❌ total node size %d, want %d", total, *exp.TotalNodeSize))
		}
	}

	return checks, passed
}

func emptyAsDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}

func executionMode(cfg TestConfig) string {
	if strings.EqualFold(cfg.Test.Set.Mode, "sequential") {
		return "sequential"
	}
	return "parallel"
}

func selectedTargets(cfg TestConfig) []TestCase {
	var out []TestCase
	for _, c := range cfg.Test.Cases {
		if c.Execute {
			out = append(out, c)
		}
	}
	return out
}

func selectedScenarios(cfg TestConfig) []Scenario {
	var out []Scenario
	for _, s := range cfg.Test.Scenarios {
		if s.Execute {
			out = append(out, s)
		}
	}
	return out
}

func printFinalSummary(report *TestReport) {
	passed := 0
	for _, r := range report.Results {
		if r.Passed {
			passed++
		}
	}
	fmt.Println("\n=========================================================")
	fmt.Println(" OVERALL TEST SUMMARY")
	fmt.Println("=========================================================")
	fmt.Printf(" Total Cases : %d\n", len(report.Results))
	fmt.Printf(" Passed      : %d\n", passed)
	fmt.Printf(" Failed      : %d\n", len(report.Results)-passed)
	if passed == len(report.Results) {
		fmt.Println(" ✅ All recommendation checks passed.")
	} else {
		fmt.Println(" ❌ Some checks failed. See testresult/ for details.")
	}
	fmt.Println("=========================================================")
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

// getGitHash returns the current git commit hash.
func getGitHash() string {
	out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

// getBeetleVersion returns the CM-Beetle version from git tags or commit hash.
// Only beetle release tags (v[0-9]*) are matched to avoid picking up imdl/*, transx/* tags.
func getBeetleVersion() string {
	commitHash := getGitHash()

	if exactTag, err := exec.Command("git", "describe", "--tags", "--exact-match", "--match", "v[0-9]*").Output(); err == nil {
		return strings.TrimSpace(string(exactTag))
	}
	if tag, err := exec.Command("git", "describe", "--tags", "--abbrev=0", "--match", "v[0-9]*").Output(); err == nil {
		if tagStr := strings.TrimSpace(string(tag)); tagStr != "" {
			if commitHash != "unknown" {
				return fmt.Sprintf("%s+ (%s)", tagStr, commitHash)
			}
			return tagStr + "+"
		}
	}
	if commitHash != "unknown" {
		return fmt.Sprintf("main (%s)", commitHash)
	}
	return "main (unknown)"
}

// maskSensitiveInfo removes CSP account identifiers from report content.
func maskSensitiveInfo(content string) string {
	reSub := regexp.MustCompile(`(?i)/subscriptions/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)
	content = reSub.ReplaceAllString(content, "/subscriptions/AZURE_SUBSCRIPTION_ID")

	reGCP := regexp.MustCompile(`projects/([a-z0-9\-]+)/`)
	content = reGCP.ReplaceAllStringFunc(content, func(match string) string {
		parts := strings.Split(match, "/")
		if len(parts) >= 2 && (parts[1] == "compute" || parts[1] == "v1") {
			return match
		}
		return "projects/GCP_PROJECT_ID/"
	})

	reEmail := regexp.MustCompile(`[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}`)
	return reEmail.ReplaceAllString(content, "MASKED_EMAIL")
}

// generateMarkdownReport writes the full result matrix and per-case details to
// testresult/k8s-infra-recommendation-test-report.md.
func generateMarkdownReport(report *TestReport) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := filepath.Join(cwd, "testresult")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	path := filepath.Join(dir, "k8s-infra-recommendation-test-report.md")
	if err := os.WriteFile(path, []byte(maskSensitiveInfo(buildMarkdown(report))), 0644); err != nil {
		return err
	}
	log.Info().Str("file", path).Msg("✅ Test report generated and saved")
	return nil
}

func buildMarkdown(report *TestReport) string {
	var sb strings.Builder

	sb.WriteString("# CM-Beetle K8s Infra Recommendation Test Results\n\n")
	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> Verifies `POST /recommendation/k8sCluster` against on-premise scenario fixtures.\n")
	sb.WriteString("> No resources are provisioned by this test.\n\n")

	sb.WriteString("## Environment\n\n")
	sb.WriteString(fmt.Sprintf("- CM-Beetle URL: %s\n", report.BeetleURL))
	sb.WriteString(fmt.Sprintf("- CM-Beetle Version: %s\n", report.BeetleVersion))
	sb.WriteString(fmt.Sprintf("- Git Commit: %s\n", report.GitHash))
	sb.WriteString(fmt.Sprintf("- Test Date: %s\n", report.TestDateTime.Format("2006-01-02 15:04:05 MST")))

	var targetList []string
	for _, t := range report.Targets {
		targetList = append(targetList, fmt.Sprintf("%s/%s", t.Csp, t.Region))
	}
	sb.WriteString(fmt.Sprintf("- Targets (%d): %s\n", len(report.Targets), strings.Join(targetList, ", ")))
	sb.WriteString(fmt.Sprintf("- Scenarios: %d\n\n", len(report.Scenarios)))

	writeSummaryMatrix(&sb, report)
	writeCaseDetails(&sb, report)

	return sb.String()
}

// writeSummaryMatrix renders a scenario × target pass/fail grid so a regression in one CSP
// or one scenario is visible at a glance.
func writeSummaryMatrix(sb *strings.Builder, report *TestReport) {
	sb.WriteString("## Summary Matrix\n\n")

	sb.WriteString("| Scenario |")
	for _, t := range report.Targets {
		sb.WriteString(fmt.Sprintf(" %s |", t.Name))
	}
	sb.WriteString("\n|---|")
	for range report.Targets {
		sb.WriteString("---|")
	}
	sb.WriteString("\n")

	index := make(map[string]CaseResult, len(report.Results))
	for _, r := range report.Results {
		index[r.ScenarioName+"|"+r.DisplayName] = r
	}

	passed := 0
	for _, sc := range report.Scenarios {
		sb.WriteString(fmt.Sprintf("| `%s` |", sc.Name))
		for _, t := range report.Targets {
			r, ok := index[sc.Name+"|"+t.Name]
			switch {
			case !ok:
				sb.WriteString(" — |")
			case r.Passed:
				passed++
				sb.WriteString(fmt.Sprintf(" ✅ %d |", r.StatusCode))
			default:
				sb.WriteString(fmt.Sprintf(" ❌ %d |", r.StatusCode))
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString(fmt.Sprintf("\n**Overall Result**: %d/%d cases passed", passed, len(report.Results)))
	if passed == len(report.Results) {
		sb.WriteString(" ✅\n\n")
	} else {
		sb.WriteString(" ❌\n\n")
	}
	sb.WriteString("---\n\n")
}

func writeCaseDetails(sb *strings.Builder, report *TestReport) {
	sb.WriteString("## Case Details\n\n")

	for _, r := range report.Results {
		status := "✅ PASS"
		if !r.Passed {
			status = "❌ FAIL"
		}
		sb.WriteString(fmt.Sprintf("### %s — %s (%s)\n\n", r.ScenarioName, r.DisplayName, status))
		sb.WriteString(fmt.Sprintf("- **Fixture**: `%s`\n", r.ScenarioFile))
		sb.WriteString(fmt.Sprintf("- **Request**: `POST %s`\n", r.RequestURL))
		sb.WriteString(fmt.Sprintf("- **Status Code**: %d\n", r.StatusCode))
		sb.WriteString(fmt.Sprintf("- **Duration**: %v\n", r.Duration.Truncate(time.Millisecond)))
		if r.Failure != "" {
			sb.WriteString(fmt.Sprintf("- **Failure**: %s\n", r.Failure))
		}
		sb.WriteString("\n")

		if len(r.Checks) > 0 {
			sb.WriteString("**Checks**:\n\n")
			for _, c := range r.Checks {
				sb.WriteString(fmt.Sprintf("- %s\n", c))
			}
			sb.WriteString("\n")
		}

		if r.ResponseBody != "" {
			sb.WriteString("<details>\n  <summary> <ins>Click to see the response body</ins> </summary>\n\n```json\n")
			sb.WriteString(prettyJSON(r.ResponseBody))
			sb.WriteString("\n```\n\n</details>\n\n")
		}
		sb.WriteString("---\n\n")
	}
}

func prettyJSON(raw string) string {
	var v interface{}
	if err := json.Unmarshal([]byte(raw), &v); err != nil {
		return raw
	}
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return raw
	}
	return string(out)
}
