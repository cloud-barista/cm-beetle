package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"

	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
)

// caseOutcome separates expected behavior from status code conformity.
type caseOutcome int

const (
	outcomeAsExpected caseOutcome = iota
	outcomeDeviation
	outcomeUnexpected
)

func (o caseOutcome) icon() string {
	switch o {
	case outcomeAsExpected:
		return "✅"
	case outcomeDeviation:
		return "⚠️ "
	default:
		return "❌"
	}
}

func (o caseOutcome) label() string {
	switch o {
	case outcomeAsExpected:
		return "as expected"
	case outcomeDeviation:
		return "as expected, non-conforming status code"
	default:
		return "UNEXPECTED"
	}
}

// CaseResult holds result for one scenario and CSP-region pair.
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

	Outcome caseOutcome
	Passed  bool
	Checks  []string
	Failure string
}

// RecommendationReport aggregates recommendation results for markdown generation.
type RecommendationReport struct {
	TestDateTime  time.Time
	BeetleURL     string
	BeetleVersion string
	GitHash       string
	Targets       []TestCase
	Scenarios     []Scenario
	Results       []CaseResult
}

// runRecommendationSuite executes the recommendation scenario test suite.
func runRecommendationSuite(client *resty.Client, cfg TestConfig, auth AuthConfig, targets []TestCase, scenarios []Scenario) *RecommendationReport {
	report := &RecommendationReport{
		TestDateTime:  time.Now(),
		BeetleURL:     cfg.Beetle.Endpoint,
		BeetleVersion: getBeetleVersion(),
		GitHash:       getGitHash(),
		Targets:       targets,
		Scenarios:     scenarios,
	}

	fmt.Println("=========================================================")
	fmt.Println(" CM-Beetle K8s Infra Recommendation Test Suite")
	fmt.Printf(" %d scenario(s) × %d target(s) = %d case(s), concurrency: %s\n",
		len(scenarios), len(targets), len(scenarios)*len(targets), cfg.Test.Set.Mode)
	fmt.Println("=========================================================")

	report.Results = runAllRecommendationCases(client, cfg, targets, scenarios)
	return report
}

// runAllRecommendationCases executes all target and scenario pairs.
func runAllRecommendationCases(client *resty.Client, cfg TestConfig, targets []TestCase, scenarios []Scenario) []CaseResult {
	if cfg.Test.Set.Mode == "sequential" {
		var all []CaseResult
		for _, t := range targets {
			all = append(all, runTargetCases(client, cfg, t, scenarios)...)
		}
		return all
	}

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

// runTargetCases executes all scenarios sequentially for a single target.
func runTargetCases(client *resty.Client, cfg TestConfig, target TestCase, scenarios []Scenario) []CaseResult {
	results := make([]CaseResult, 0, len(scenarios))
	for i, sc := range scenarios {
		if i > 0 {
			time.Sleep(150 * time.Millisecond)
		}
		results = append(results, runCase(client, cfg, target, sc))
	}
	return results
}

// runCase executes one scenario fixture against one CSP target.
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

	res.Checks, res.Outcome = evaluate(sc.Expect, res.StatusCode, resp.Body())
	res.Passed = res.Outcome != outcomeUnexpected

	fmt.Printf("%s [%s] %-22s HTTP %d — %s (%v)\n", res.Outcome.icon(), target.Name, sc.Name,
		res.StatusCode, res.Outcome.label(), res.Duration.Truncate(time.Millisecond))
	for _, c := range res.Checks {
		fmt.Printf("      %s\n", c)
	}

	return res
}

// wantsAcceptance checks if expectation expects HTTP 200.
func (e Expect) wantsAcceptance() bool { return e.StatusCode == 200 }

// accepted reports whether HTTP status code is in 2xx range.
func accepted(statusCode int) bool { return statusCode >= 200 && statusCode < 300 }

// evaluate verifies response against declared expectation and structural rules.
func evaluate(exp Expect, statusCode int, body []byte) (checks []string, outcome caseOutcome) {
	outcomeMatches := exp.wantsAcceptance() == accepted(statusCode)

	switch {
	case !outcomeMatches && exp.wantsAcceptance():
		checks = append(checks, fmt.Sprintf("❌ input was rejected (HTTP %d) but a recommendation was expected", statusCode))
		return checks, outcomeUnexpected
	case !outcomeMatches:
		checks = append(checks, fmt.Sprintf("❌ input was accepted (HTTP %d) but rejection was expected", statusCode))
		return checks, outcomeUnexpected
	case statusCode == exp.StatusCode:
		checks = append(checks, fmt.Sprintf("✅ status code %d as expected", statusCode))
		outcome = outcomeAsExpected
	default:
		checks = append(checks,
			fmt.Sprintf("✅ input rejected as expected"),
			fmt.Sprintf("⚠️  status code %d, but the API declares %d for this case", statusCode, exp.StatusCode))
		outcome = outcomeDeviation
	}

	if !exp.wantsAcceptance() {
		return checks, outcome
	}

	var apiResp model.ApiResponse[cloudmodel.RecommendedInfra]
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return append(checks, fmt.Sprintf("❌ failed to parse response: %s", err)), outcomeUnexpected
	}
	if !apiResp.Success {
		return append(checks, fmt.Sprintf("❌ response success=false: %s", apiResp.Error)), outcomeUnexpected
	}

	cluster := apiResp.Data.TargetK8sCluster

	if cluster.Name == "" {
		outcome = outcomeUnexpected
		checks = append(checks, "❌ targetK8sCluster.name is empty")
	}
	if cluster.Version == "" {
		outcome = outcomeUnexpected
		checks = append(checks, "❌ targetK8sCluster.version is empty")
	} else {
		checks = append(checks, fmt.Sprintf("✅ version: %s", cluster.Version))
	}

	if exp.VersionPrefix != "" {
		if strings.HasPrefix(cluster.Version, exp.VersionPrefix) {
			checks = append(checks, fmt.Sprintf("✅ version has prefix %q", exp.VersionPrefix))
		} else {
			outcome = outcomeUnexpected
			checks = append(checks, fmt.Sprintf("❌ version %q lacks prefix %q", cluster.Version, exp.VersionPrefix))
		}
	}

	groups := cluster.K8sNodeGroupList
	if exp.NodeGroupCount != nil {
		if len(groups) == *exp.NodeGroupCount {
			checks = append(checks, fmt.Sprintf("✅ node group count: %d", len(groups)))
		} else {
			outcome = outcomeUnexpected
			checks = append(checks, fmt.Sprintf("❌ node group count %d, want %d", len(groups), *exp.NodeGroupCount))
		}
	} else {
		checks = append(checks, fmt.Sprintf("ℹ️  node group count: %d", len(groups)))
	}

	total := 0
	for i, ng := range groups {
		total += ng.DesiredNodeSize
		if ng.SpecId == "" {
			outcome = outcomeUnexpected
			checks = append(checks, fmt.Sprintf("❌ node group[%d] %q has empty specId", i, ng.Name))
		}
		if ng.DesiredNodeSize < 1 {
			outcome = outcomeUnexpected
			checks = append(checks, fmt.Sprintf("❌ node group[%d] %q desiredNodeSize=%d (< 1)", i, ng.Name, ng.DesiredNodeSize))
		}
		checks = append(checks, fmt.Sprintf("ℹ️  node group[%d] %q spec=%s image=%s nodes=%d",
			i, ng.Name, ng.SpecId, emptyAsDash(ng.ImageId), ng.DesiredNodeSize))
	}

	if exp.TotalNodeSize != nil {
		if total == *exp.TotalNodeSize {
			checks = append(checks, fmt.Sprintf("✅ total node size: %d", total))
		} else {
			outcome = outcomeUnexpected
			checks = append(checks, fmt.Sprintf("❌ total node size %d, want %d", total, *exp.TotalNodeSize))
		}
	}

	return checks, outcome
}

// emptyAsDash returns a dash if the input string is empty.
func emptyAsDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}
