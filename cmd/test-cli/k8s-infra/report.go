package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

// printBanner prints target header banner.
func printBanner(t TestCase) {
	fmt.Printf("\n=========================================================\n")
	fmt.Printf(" Target: %s (%s / %s)\n", t.Name, t.Csp, t.Region)
	fmt.Printf("=========================================================\n")
}

// printStepStart prints initial step announcement.
func printStepStart(target string, number int, name string) {
	fmt.Printf("▶  [%s] Step %d: %s ...\n", target, number, name)
}

// printStep reports finished step result and duration.
func printStep(target string, res StepResult) {
	icon := "✅"
	switch {
	case res.Skipped:
		icon = "⏭️"
	case !res.Success:
		icon = "❌"
	}
	fmt.Printf("%s [%s] Step %d: %s — done in %v\n", icon, target, res.Number, res.Name,
		res.Duration.Truncate(time.Millisecond))
	for _, n := range res.Notes {
		fmt.Printf("      [%s] %s\n", target, n)
	}
	if res.Error != "" {
		fmt.Printf("      [%s] error: %s\n", target, res.Error)
	}
}

// progressf prints an in-flight progress line.
func progressf(target, format string, a ...interface{}) {
	fmt.Printf("      [%s] %s\n", target, fmt.Sprintf(format, a...))
}

// truncate truncates string with ellipsis if longer than n.
func truncate(s string, n int) string {
	s = strings.TrimSpace(s)
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// countSteps counts passed and total steps in a report.
func countSteps(r *CSPTestReport) (passed, total int) {
	for _, s := range r.Steps {
		total++
		if s.Success {
			passed++
		}
	}
	return passed, total
}

// getGitHash returns short git commit hash.
func getGitHash() string {
	out, err := exec.Command("git", "rev-parse", "--short", "HEAD").Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(out))
}

// getBeetleVersion returns beetle release tag or commit hash.
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

// maskSensitiveInfo redacts cloud subscription IDs and emails from report.
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

// testResultDir returns or creates testresult directory.
func testResultDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(cwd, "testresult")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

// generateCSPReport writes Markdown report for one CSP target.
func generateCSPReport(report *CSPTestReport) error {
	dir, err := testResultDir()
	if err != nil {
		return err
	}
	path := filepath.Join(dir, fmt.Sprintf("k8s-infra-migration-test-results-%s.md", strings.ToLower(report.CSP)))
	if err := os.WriteFile(path, []byte(maskSensitiveInfo(buildCSPMarkdown(report))), 0644); err != nil {
		return err
	}
	log.Info().Str("file", path).Msg("✅ CSP test report generated and saved")
	return nil
}

// buildCSPMarkdown builds Markdown text for one CSP target.
func buildCSPMarkdown(report *CSPTestReport) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# CM-Beetle K8s Infra Migration Test Results — %s\n\n", report.DisplayName))
	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> Full lifecycle against a real CSP: recommend → validate → migrate → list → get → workload → delete → residual.\n\n")

	sb.WriteString("## Environment\n\n")
	sb.WriteString(fmt.Sprintf("- CSP / Region: %s / %s\n", report.CSP, report.Region))
	sb.WriteString(fmt.Sprintf("- CM-Beetle URL: %s\n", report.BeetleURL))
	sb.WriteString(fmt.Sprintf("- CM-Beetle Version: %s\n", report.BeetleVersion))
	sb.WriteString(fmt.Sprintf("- Git Commit: %s\n", report.GitHash))
	sb.WriteString(fmt.Sprintf("- Namespace: %s\n", report.NamespaceID))
	sb.WriteString(fmt.Sprintf("- Test Date: %s\n", report.TestDateTime.Format("2006-01-02 15:04:05 MST")))
	if report.ClusterID != "" {
		sb.WriteString(fmt.Sprintf("- Cluster ID: %s\n", report.ClusterID))
	}
	sb.WriteString("\n")

	passed, total := countSteps(report)
	sb.WriteString("## Test Results Summary\n\n")
	sb.WriteString("| Step | Description | Status | Duration |\n")
	sb.WriteString("|------|-------------|--------|----------|\n")
	var totalDuration time.Duration
	for _, s := range report.Steps {
		status := "✅ **PASS**"
		switch {
		case s.Skipped:
			status = "⏭️ **SKIP**"
		case !s.Success:
			status = "❌ **FAIL**"
		}
		sb.WriteString(fmt.Sprintf("| %d | %s | %s | %v |\n", s.Number, s.Name, status, s.Duration.Truncate(time.Millisecond)))
		totalDuration += s.Duration
	}
	sb.WriteString(fmt.Sprintf("\n**Overall Result**: %d/%d steps passed", passed, total))
	if passed == total {
		sb.WriteString(" ✅\n\n")
	} else {
		sb.WriteString(" ❌\n\n")
	}
	sb.WriteString(fmt.Sprintf("**Total Duration**: %v\n\n---\n\n", totalDuration.Truncate(time.Second)))

	sb.WriteString("## Step Details\n\n")
	for _, s := range report.Steps {
		sb.WriteString(fmt.Sprintf("### Step %d — %s\n\n", s.Number, s.Name))
		sb.WriteString(fmt.Sprintf("- **Duration**: %v\n", s.Duration.Truncate(time.Millisecond)))
		if s.StatusCode != 0 {
			sb.WriteString(fmt.Sprintf("- **Status Code**: %d\n", s.StatusCode))
		}
		if s.Error != "" {
			sb.WriteString(fmt.Sprintf("- **Error**: %s\n", s.Error))
		}
		sb.WriteString("\n")
		for _, n := range s.Notes {
			sb.WriteString(fmt.Sprintf("- %s\n", n))
		}
		sb.WriteString("\n")
	}

	if report.Recommendation != nil {
		sb.WriteString("## Recommendation (input to migration)\n\n")
		sb.WriteString("<details>\n  <summary> <ins>Click to see the recommendation</ins> </summary>\n\n```json\n")
		b, _ := json.MarshalIndent(report.Recommendation, "", "  ")
		sb.WriteString(string(b))
		sb.WriteString("\n```\n\n</details>\n\n")
	}
	if report.ClusterInfo != nil {
		sb.WriteString("## Created Cluster\n\n")
		sb.WriteString("<details>\n  <summary> <ins>Click to see the cluster info</ins> </summary>\n\n```json\n")
		b, _ := json.MarshalIndent(report.ClusterInfo, "", "  ")
		sb.WriteString(string(b))
		sb.WriteString("\n```\n\n</details>\n\n")
	}

	return sb.String()
}

// generateSummaryReport writes summary Markdown report across all targets.
func generateSummaryReport(reports []*CSPTestReport) error {
	dir, err := testResultDir()
	if err != nil {
		return err
	}

	var sb strings.Builder
	sb.WriteString("# CM-Beetle K8s Infra Migration Test Summary\n\n")
	if len(reports) > 0 {
		sb.WriteString(fmt.Sprintf("- CM-Beetle Version: %s\n", reports[0].BeetleVersion))
		sb.WriteString(fmt.Sprintf("- Test Date: %s\n\n", reports[0].TestDateTime.Format("2006-01-02 15:04:05 MST")))
	}

	sb.WriteString("| Target | CSP / Region | Steps Passed | Cluster ID | Result |\n")
	sb.WriteString("|--------|--------------|--------------|------------|--------|\n")
	for _, r := range reports {
		passed, total := countSteps(r)
		result := "✅ PASS"
		if passed != total {
			result = "❌ FAIL"
		}
		clusterID := r.ClusterID
		if clusterID == "" {
			clusterID = "—"
		}
		sb.WriteString(fmt.Sprintf("| %s | %s / %s | %d/%d | `%s` | %s |\n",
			r.DisplayName, r.CSP, r.Region, passed, total, clusterID, result))
	}
	sb.WriteString("\n## Per-CSP Reports\n\n")
	for _, r := range reports {
		sb.WriteString(fmt.Sprintf("- [%s](k8s-infra-migration-test-results-%s.md)\n", r.DisplayName, strings.ToLower(r.CSP)))
	}

	path := filepath.Join(dir, "k8s-infra-migration-test-summary-all.md")
	if err := os.WriteFile(path, []byte(maskSensitiveInfo(sb.String())), 0644); err != nil {
		return err
	}
	log.Info().Str("file", path).Msg("✅ Summary report generated and saved")
	return nil
}

// printMigrationFinalSummary prints final console summary for migration.
func printMigrationFinalSummary(reports []*CSPTestReport) bool {
	fmt.Println("\n=========================================================")
	fmt.Println(" OVERALL TEST SUMMARY")
	fmt.Println("=========================================================")

	allPassed := true
	for _, r := range reports {
		passed, total := countSteps(r)
		status := "✅"
		if passed != total {
			status = "❌"
			allPassed = false
		}
		fmt.Printf(" %s %-22s %d/%d steps passed\n", status, r.DisplayName, passed, total)
	}
	fmt.Println("=========================================================")
	return allPassed
}

// countOutcomes tallies recommendation judgements.
func countOutcomes(results []CaseResult) (exact, deviation, unexpected int) {
	for _, r := range results {
		switch r.Outcome {
		case outcomeAsExpected:
			exact++
		case outcomeDeviation:
			deviation++
		default:
			unexpected++
		}
	}
	return exact, deviation, unexpected
}

// printRecommendationFinalSummary prints final console summary for recommendation suite.
func printRecommendationFinalSummary(report *RecommendationReport) bool {
	exact, deviation, unexpected := countOutcomes(report.Results)
	total := len(report.Results)

	fmt.Println("\n=========================================================")
	fmt.Println(" OVERALL RECOMMENDATION TEST SUMMARY")
	fmt.Println("=========================================================")
	fmt.Printf(" Total cases            : %d\n", total)
	fmt.Printf(" Behaved as expected    : %d\n", exact+deviation)
	fmt.Printf("   ├ fully conforming   : %d\n", exact)
	fmt.Printf("   └ status code differs: %d\n", deviation)
	fmt.Printf(" Unexpected behaviour   : %d\n", unexpected)
	fmt.Println("=========================================================")

	switch {
	case unexpected > 0:
		fmt.Printf(" ❌ %d case(s) did not behave as expected. See testresult/ for details.\n", unexpected)
	case deviation > 0:
		fmt.Printf(" ✅ All %d case(s) behaved as expected.\n", total)
		fmt.Printf(" ⚠️  %d of them returned a status code different from declared\n", deviation)
	default:
		fmt.Printf(" ✅ All %d case(s) behaved as expected with conforming status codes.\n", total)
	}
	fmt.Println("=========================================================")
	return unexpected == 0
}

// generateRecommendationMarkdownReport writes recommendation result report.
func generateRecommendationMarkdownReport(report *RecommendationReport) error {
	dir, err := testResultDir()
	if err != nil {
		return err
	}
	path := filepath.Join(dir, "k8s-infra-recommendation-test-report.md")
	if err := os.WriteFile(path, []byte(maskSensitiveInfo(buildRecommendationMarkdown(report))), 0644); err != nil {
		return err
	}
	log.Info().Str("file", path).Msg("✅ Recommendation test report generated and saved")
	return nil
}

// buildRecommendationMarkdown builds Markdown text for recommendation suite.
func buildRecommendationMarkdown(report *RecommendationReport) string {
	var sb strings.Builder

	sb.WriteString("# CM-Beetle K8s Infra Recommendation Test Results\n\n")
	sb.WriteString("> [!NOTE]\n")
	sb.WriteString("> Verifies `POST /recommendation/k8sCluster` against on-premise scenario fixtures.\n")
	sb.WriteString("> No cloud resources are provisioned by this test.\n\n")

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

// writeSummaryMatrix renders a scenario × target matrix.
func writeSummaryMatrix(sb *strings.Builder, report *RecommendationReport) {
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

	for _, sc := range report.Scenarios {
		sb.WriteString(fmt.Sprintf("| `%s` |", sc.Name))
		for _, t := range report.Targets {
			r, ok := index[sc.Name+"|"+t.Name]
			if !ok {
				sb.WriteString(" — |")
				continue
			}
			sb.WriteString(fmt.Sprintf(" %s %d |", r.Outcome.icon(), r.StatusCode))
		}
		sb.WriteString("\n")
	}

	exact, deviation, unexpected := countOutcomes(report.Results)
	total := len(report.Results)

	sb.WriteString("\n| Legend | Meaning |\n|---|---|\n")
	sb.WriteString("| ✅ | Behaved as expected, with the status code the API declares |\n")
	sb.WriteString("| ⚠️ | Behaved as expected (input correctly accepted or rejected), but status code differs |\n")
	sb.WriteString("| ❌ | Did not behave as expected |\n\n")

	sb.WriteString(fmt.Sprintf("**Behaved as expected**: %d/%d", exact+deviation, total))
	if unexpected == 0 {
		sb.WriteString(" ✅\n\n")
	} else {
		sb.WriteString(" ❌\n\n")
	}
	sb.WriteString(fmt.Sprintf("- Fully conforming: %d\n", exact))
	sb.WriteString(fmt.Sprintf("- Status code differs: %d\n", deviation))
	sb.WriteString(fmt.Sprintf("- Unexpected behaviour: %d\n\n", unexpected))
	sb.WriteString("---\n\n")
}

// writeCaseDetails renders details for each case result.
func writeCaseDetails(sb *strings.Builder, report *RecommendationReport) {
	sb.WriteString("## Case Details\n\n")

	for _, r := range report.Results {
		sb.WriteString(fmt.Sprintf("### %s — %s (%s %s)\n\n",
			r.ScenarioName, r.DisplayName, r.Outcome.icon(), r.Outcome.label()))
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

// prettyJSON formats raw JSON string with indentation.
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
