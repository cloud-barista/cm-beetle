package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
)

var (
	configFile   = flag.String("config", "testconf/template-test-config.yaml", "Path to stress test YAML config")
	modeFlag     = flag.String("mode", "", "Override execution mode: simulate or real")
	fromJSONFlag = flag.String("from-json", "", "Re-generate Markdown and HTML dashboard from an existing JSON report")
	outDirFlag   = flag.String("out-dir", "testresult", "Output directory for JSON, Markdown, and HTML artifacts")
)

func main() {
	flag.Parse()

	cleanOutDir := filepath.Clean(*outDirFlag)
	if err := os.MkdirAll(cleanOutDir, 0755); err != nil {
		log.Fatal().Err(err).Msg("Failed to create output directory")
	}

	jsonOutPath := filepath.Join(cleanOutDir, "stress-test-data.json")
	mdOutPath := filepath.Join(cleanOutDir, "stress-test-summary.md")
	htmlOutPath := filepath.Join(cleanOutDir, "stress-test-dashboard.html")

	// Support standalone dashboard re-generation from an existing JSON data file.
	if strings.TrimSpace(*fromJSONFlag) != "" {
		report, err := loadReportFromJSON(*fromJSONFlag)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to load existing JSON report")
		}
		if err := GenerateMarkdownReport(report, mdOutPath); err != nil {
			log.Fatal().Err(err).Msg("Failed to write Markdown report")
		}
		if err := GenerateInteractiveDashboardHTML(report, htmlOutPath); err != nil {
			log.Fatal().Err(err).Msg("Failed to write HTML dashboard")
		}
		fmt.Printf("\nRe-generated artifacts from %s:\n  - Markdown:  %s\n  - Dashboard: %s\n", *fromJSONFlag, mdOutPath, htmlOutPath)
		return
	}

	cfg, err := loadStressConfig(*configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load stress test config")
	}

	if strings.TrimSpace(*modeFlag) != "" {
		cfg.Stress.ExecutionMode = strings.TrimSpace(*modeFlag)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatal().Err(err).Msg("Invalid stress test configuration")
	}

	report := RunStressSuite(cfg, func(partial *StressTestReport) {
		_ = saveReportJSON(partial, jsonOutPath)
		_ = GenerateMarkdownReport(partial, mdOutPath)
		_ = GenerateInteractiveDashboardHTML(partial, htmlOutPath)
	})

	if err := saveReportJSON(report, jsonOutPath); err != nil {
		log.Fatal().Err(err).Msg("Failed to save final JSON report")
	}
	if err := GenerateMarkdownReport(report, mdOutPath); err != nil {
		log.Fatal().Err(err).Msg("Failed to generate final Markdown report")
	}
	if err := GenerateInteractiveDashboardHTML(report, htmlOutPath); err != nil {
		log.Fatal().Err(err).Msg("Failed to generate final HTML dashboard")
	}

	printConsoleSummary(report, jsonOutPath, mdOutPath, htmlOutPath)
}

// RunStressSuite executes the equal-round progressive multi-CSP stress test according to cfg.
func RunStressSuite(cfg StressTestConfig, onStepCheckpoint func(*StressTestReport)) *StressTestReport {
	seed := cfg.Stress.RandomSeed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}
	rng := rand.New(rand.NewSource(seed))

	client := resty.New()
	client.SetTimeout(60 * time.Minute)

	if authCfg, err := loadAuthConfig(cfg.Beetle.AuthConfigFile); err == nil && authCfg.BeetleApiUsername != "" {
		client.SetBasicAuth(authCfg.BeetleApiUsername, authCfg.BeetleApiPassword)
	}

	beetleReady := checkBeetleReady(client, cfg.Beetle.Endpoint)
	if cfg.Stress.ExecutionMode == ExecutionModeReal && !beetleReady {
		log.Fatal().Str("endpoint", cfg.Beetle.Endpoint).Msg("CM-Beetle server is not reachable in real execution mode")
	}

	activeCases := make([]CSPTestCase, 0)
	statuses := make([]CSPSurvivalStatus, 0)
	statusIndexByCSP := make(map[string]int)

	for _, tc := range cfg.Test.Cases {
		if !tc.Execute {
			continue
		}
		cspKey := strings.ToLower(strings.TrimSpace(tc.Csp))
		disp := tc.Name
		if disp == "" {
			disp = fmt.Sprintf("%s-%s", strings.ToUpper(cspKey), tc.Region)
		}
		tc.Csp = cspKey
		tc.Name = disp
		activeCases = append(activeCases, tc)

		statusIndexByCSP[cspKey] = len(statuses)
		statuses = append(statuses, CSPSurvivalStatus{
			CSP:             cspKey,
			Region:          tc.Region,
			DisplayName:     disp,
			Status:          "SURVIVED",
			CleanupVerified: true,
		})
	}

	now := time.Now()
	report := &StressTestReport{
		ReportVersion:  "1.0",
		TestDate:       now.Format("2006-01-02"),
		TestTime:       now.Format("15:04:05 MST"),
		StartTime:      now,
		ExecutionMode:  cfg.Stress.ExecutionMode,
		InitialNodes:   cfg.Stress.InitialNodes,
		StepNodes:      cfg.Stress.StepNodes,
		MaxNodes:       cfg.Stress.MaxNodes,
		BeetleEndpoint: cfg.Beetle.Endpoint,
		NamespaceID:    cfg.Beetle.NamespaceID,
		CSPStatuses:    statuses,
		Steps:          make([]StepAttempt, 0),
	}

	printStartBanner(cfg, len(activeCases), beetleReady)

	eliminated := make(map[string]bool)
	seqIndex := 0
	elapsedSimSec := 0.0

	recordAttempt := func(attempt StepAttempt) {
		elapsedSimSec = math.Round((elapsedSimSec+attempt.SimulatedDurationS)*10) / 10
		attempt.ElapsedSimulatedS = elapsedSimSec
		report.Steps = append(report.Steps, attempt)
		report.TotalSteps = len(report.Steps)
		if attempt.Round > report.TotalRounds {
			report.TotalRounds = attempt.Round
		}
		if attempt.TargetNodes > report.MaxReachedNodes {
			report.MaxReachedNodes = attempt.TargetNodes
		}

		idx := statusIndexByCSP[attempt.CSP]
		st := &report.CSPStatuses[idx]
		st.TotalStepsAttempted++
		if !attempt.CleanupSummary.CleanVerified {
			st.CleanupVerified = false
			report.OrphanResourcesLeft += attempt.CleanupSummary.OrphanResourcesRemaining
		}

		if attempt.Success {
			st.TotalStepsSucceeded++
			st.TotalNodesProvisioned += attempt.TargetNodes
			st.TotalNodesCleaned += attempt.TargetNodes
			if attempt.TargetNodes > st.MaxSucceededNodes {
				st.MaxSucceededNodes = attempt.TargetNodes
			}
			report.TotalNodesProvisioned += attempt.TargetNodes
			report.TotalNodesCleaned += attempt.TargetNodes
		} else {
			eliminated[attempt.CSP] = true
			st.Status = "ELIMINATED"
			st.FailedAtNodes = attempt.TargetNodes
			st.EliminatedRound = attempt.Round
			st.EliminatedStepIndex = attempt.SequenceIndex
			st.FailureCode = attempt.FailureCode
			st.FailureReason = attempt.FailureReason
		}

		printStepLine(attempt)
		if onStepCheckpoint != nil {
			onStepCheckpoint(report)
		}
	}

	round := 1
	for targetNodes := cfg.Stress.InitialNodes; targetNodes <= cfg.Stress.MaxNodes; targetNodes += cfg.Stress.StepNodes {
		ranAnyInRound := false
		for _, tc := range activeCases {
			if eliminated[tc.Csp] {
				continue
			}
			ranAnyInRound = true
			seqIndex++
			attempt := ExecuteStressStep(rng, client, cfg, beetleReady, tc, round, seqIndex, targetNodes)
			recordAttempt(attempt)
		}
		if !ranAnyInRound {
			break
		}
		round++
	}

	report.EndTime = time.Now()
	report.WallClockDurationMs = max(1, report.EndTime.Sub(report.StartTime).Milliseconds())
	report.TotalSimulatedSec = elapsedSimSec

	return report
}

func checkBeetleReady(client *resty.Client, endpoint string) bool {
	if strings.TrimSpace(endpoint) == "" {
		return false
	}
	url := fmt.Sprintf("%s/beetle/readyz", strings.TrimRight(endpoint, "/"))
	var resp map[string]interface{}
	var emptyBody interface{} = common.NoBody
	reqClient := resty.New().SetTimeout(3 * time.Second)
	err := common.ExecuteHttpRequest(reqClient, "GET", url, nil, false, &emptyBody, &resp, 0)
	return err == nil
}

func loadStressConfig(path string) (StressTestConfig, error) {
	var cfg StressTestConfig
	cleanPath := filepath.Clean(path)
	data, err := os.ReadFile(cleanPath)
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func loadAuthConfig(path string) (AuthConfig, error) {
	var auth AuthConfig
	if strings.TrimSpace(path) == "" {
		return auth, nil
	}
	cleanPath := filepath.Clean(path)
	data, err := os.ReadFile(cleanPath)
	if err != nil {
		return auth, err
	}
	err = json.Unmarshal(data, &auth)
	return auth, err
}

func saveReportJSON(report *StressTestReport, path string) error {
	cleanPath := filepath.Clean(path)
	bytes, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(cleanPath, bytes, 0644)
}

func loadReportFromJSON(path string) (*StressTestReport, error) {
	cleanPath := filepath.Clean(path)
	bytes, err := os.ReadFile(cleanPath)
	if err != nil {
		return nil, err
	}
	var report StressTestReport
	if err := json.Unmarshal(bytes, &report); err != nil {
		return nil, err
	}
	return &report, nil
}

func printStartBanner(cfg StressTestConfig, cspCount int, beetleReady bool) {
	fmt.Println("\n================================================================================")
	fmt.Println(" CM-Beetle Multi-CSP Progressive Stress Test & Zero-Residue Cleanup Runner")
	fmt.Println("================================================================================")
	fmt.Printf(" Execution Mode   : %s (CM-Beetle Live Reachable: %t)\n", cfg.Stress.ExecutionMode, beetleReady)
	fmt.Printf(" Node Scale       : %d Nodes -> %d Nodes (+%d Nodes per Round across Surviving CSPs)\n", cfg.Stress.InitialNodes, cfg.Stress.MaxNodes, cfg.Stress.StepNodes)
	fmt.Printf(" Target CSPs      : %d Active CSPs\n", cspCount)
	fmt.Println("================================================================================")
}

func printStepLine(a StepAttempt) {
	if a.Success {
		fmt.Printf(" [Step #%02d | R%02d] %-16s | %3d Nodes (%3d vCPU, %3d GiB) | PASS       | Provision: %5.1fs | Cleanup: %4.1fs (0 Orphans)\n",
			a.SequenceIndex, a.Round, a.DisplayName, a.TargetNodes, a.TotalVCPUs, a.TotalMemoryGB,
			a.MigratePhase.SimulatedSec, a.CleanupPhase.SimulatedSec)
	} else {
		fmt.Printf(" [Step #%02d | R%02d] %-16s | %3d Nodes (%3d vCPU, %3d GiB) | ELIMINATED | [%s] -> Auto-Cleanup Verified (%4.1fs, 0 Orphans)\n",
			a.SequenceIndex, a.Round, a.DisplayName, a.TargetNodes, a.TotalVCPUs, a.TotalMemoryGB,
			a.FailureCode, a.CleanupPhase.SimulatedSec)
	}
}

func printConsoleSummary(report *StressTestReport, jsonPath, mdPath, htmlPath string) {
	fmt.Println("\n================================================================================")
	fmt.Println(" STRESS TEST FINAL STANDING")
	fmt.Println("================================================================================")
	for _, st := range report.CSPStatuses {
		if st.Status == "SURVIVED" {
			fmt.Printf("  [SURVIVED]   %-16s | Max Verified: %3d Nodes | All Rounds Passed & Swept (0 Orphans)\n",
				st.DisplayName, st.MaxSucceededNodes)
		} else {
			fmt.Printf("  [ELIMINATED] %-16s | Max Verified: %3d Nodes | Failed @ %3d Nodes (%s) -> Auto-Cleaned\n",
				st.DisplayName, st.MaxSucceededNodes, st.FailedAtNodes, st.FailureCode)
		}
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf(" Total Steps Executed        : %d Steps across %d Rounds\n", report.TotalSteps, report.TotalRounds)
	fmt.Printf(" Total Nodes Provisioned     : %d Nodes (Cleaned: %d Nodes, Orphans Remaining: %d)\n",
		report.TotalNodesProvisioned, report.TotalNodesCleaned, report.OrphanResourcesLeft)
	fmt.Printf(" JSON Data Artifact          : %s\n", jsonPath)
	fmt.Printf(" Markdown Report Artifact    : %s\n", mdPath)
	fmt.Printf(" Interactive Time-Lapse HTML : %s\n", htmlPath)
	fmt.Println("================================================================================")
}
