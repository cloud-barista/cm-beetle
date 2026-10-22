// Package main provides the CM-Beetle GPU Recommendation Test CLI.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
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

// TestConfig defines test configuration.
type TestConfig struct {
	Test struct {
		Set struct {
			Mode string `yaml:"mode"` // sequential or parallel
		} `yaml:"set"`
		Cases []TestCase `yaml:"cases"`
	} `yaml:"test"`
	Beetle struct {
		Endpoint        string `yaml:"endpoint"`
		RequestBodyFile string `yaml:"requestBodyFile"`
		AuthConfigFile  string `yaml:"authConfigFile"`
	} `yaml:"beetle"`
}

// TestCase defines a target CSP and region.
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

// RecommendationResult holds results for a single CSP-Region recommendation call.
type RecommendationResult struct {
	CSP          string
	Region       string
	DisplayName  string
	StartTime    time.Time
	EndTime      time.Time
	Duration     time.Duration
	Success      bool
	StatusCode   int
	ErrorMessage string
	Candidates   []cloudmodel.RecommendedInfra
}

var configFile = flag.String("config", "testconf/test-config.yaml", "Path to test configuration file")

func init() {
	config.Init()
	l := logger.NewLogger(logger.Config{
		LogLevel:  config.Beetle.LogLevel,
		LogWriter: config.Beetle.LogWriter,
	})
	log.Logger = *l
}

func main() {
	flag.Parse()

	cfg, err := loadConfig(*configFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load test configuration")
	}

	resolvePath := func(p string) string {
		if p == "" || filepath.IsAbs(p) {
			return p
		}
		if _, err := os.Stat(p); err == nil {
			return p
		}
		configDir := filepath.Dir(*configFile)
		candidate := filepath.Join(configDir, filepath.Base(p))
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		return filepath.Join(configDir, p)
	}

	onpremInfra, err := loadOnpremInfra(resolvePath(cfg.Beetle.RequestBodyFile))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load source infrastructure JSON")
	}

	authCfg, _ := loadAuthConfig(resolvePath(cfg.Beetle.AuthConfigFile))

	client := resty.New()
	client.SetTimeout(10 * time.Minute)
	client.SetBaseURL(cfg.Beetle.Endpoint)
	if authCfg.BeetleApiUsername != "" && authCfg.BeetleApiPassword != "" {
		client.SetBasicAuth(authCfg.BeetleApiUsername, authCfg.BeetleApiPassword)
	}

	if err := checkBeetleReadiness(client, cfg.Beetle.Endpoint); err != nil {
		log.Fatal().Err(err).Msg("CM-Beetle service is not ready")
	}

	log.Info().Msg("\n" + strings.Repeat("=", 70))
	log.Info().Msg("STARTING CM-BEETLE GPU INFRASTRUCTURE RECOMMENDATION TESTS")
	if len(onpremInfra.Nodes) > 0 {
		n := onpremInfra.Nodes[0]
		gpuInfo := "None"
		if len(n.GPUCards) > 0 {
			gpuInfo = fmt.Sprintf("%dx %s (%.0f GB VRAM)", len(n.GPUCards), n.GPUCards[0].Model, n.GPUCards[0].MemoryTotalGB)
		}
		log.Info().Msgf("Source Workload: %s | %d vCPU | %d GiB RAM | OS: %s | GPU: %s",
			n.Hostname, n.CPU.Cpus*n.CPU.Threads, n.Memory.TotalSize, n.OS.PrettyName, gpuInfo)
	}
	log.Info().Msg(strings.Repeat("=", 70) + "\n")

	results := make([]RecommendationResult, 0)
	var mu sync.Mutex

	isParallel := strings.EqualFold(cfg.Test.Set.Mode, "parallel")
	var wg sync.WaitGroup

	for i, tc := range cfg.Test.Cases {
		if !tc.Execute {
			log.Info().Msgf("[%d/%d] Skipping %s (%s, %s): execute=false",
				i+1, len(cfg.Test.Cases), tc.Name, tc.Csp, tc.Region)
			continue
		}

		if isParallel {
			wg.Add(1)
			go func(idx int, c TestCase) {
				defer wg.Done()
				res := runGpuRecommendationTest(client, cfg.Beetle.Endpoint, c, onpremInfra)
				mu.Lock()
				results = append(results, res)
				mu.Unlock()
			}(i, tc)
		} else {
			res := runGpuRecommendationTest(client, cfg.Beetle.Endpoint, tc, onpremInfra)
			results = append(results, res)
		}
	}

	if isParallel {
		wg.Wait()
	}

	printSummary(results, onpremInfra)
	if err := writeSummaryReport(results); err != nil {
		log.Warn().Err(err).Msg("Failed to write summary markdown report")
	}
}

// runGpuRecommendationTest executes POST /beetle/recommendation/infra for a target.
func runGpuRecommendationTest(client *resty.Client, endpoint string, tc TestCase, infra onpremmodel.OnpremInfra) RecommendationResult {
	displayName := tc.Name
	if displayName == "" {
		displayName = fmt.Sprintf("%s-%s", tc.Csp, tc.Region)
	}

	log.Info().Msgf("--> Testing GPU recommendation for: %s (%s, %s)", displayName, tc.Csp, tc.Region)

	reqBody := controller.RecommendInfraRequest{
		DesiredCspAndRegionPair: cloudmodel.CloudProperty{
			Csp:    tc.Csp,
			Region: tc.Region,
		},
		OnpremiseInfraModel: infra,
	}

	res := RecommendationResult{
		CSP:         tc.Csp,
		Region:      tc.Region,
		DisplayName: displayName,
		StartTime:   time.Now(),
	}

	url := fmt.Sprintf("%s/beetle/recommendation/infra?desiredCsp=%s&desiredRegion=%s&limit=3",
		endpoint, tc.Csp, tc.Region)

	var apiResp model.ApiResponse[[]cloudmodel.RecommendedInfra]
	var errResp map[string]interface{}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetResult(&apiResp).
		SetError(&errResp).
		Post(url)

	res.EndTime = time.Now()
	res.Duration = res.EndTime.Sub(res.StartTime)

	if err != nil {
		res.Success = false
		res.ErrorMessage = err.Error()
		log.Error().Err(err).Msgf("Request failed for %s", displayName)
		return res
	}

	res.StatusCode = resp.StatusCode()
	if resp.IsError() {
		res.Success = false
		res.ErrorMessage = fmt.Sprintf("HTTP %d: %s", resp.StatusCode(), resp.String())
		log.Error().Msgf("Failed with status %d for %s", resp.StatusCode(), displayName)
		return res
	}

	res.Success = true
	res.Candidates = apiResp.Data

	log.Info().Msgf("Received %d candidate(s) for %s in %v", len(res.Candidates), displayName, res.Duration)
	for idx, cand := range res.Candidates {
		log.Info().Msgf("  Candidate #%d [%s]: %s", idx+1, cand.Status, cand.Description)
		for _, ng := range cand.TargetInfra.NodeGroups {
			log.Info().Msgf("    NodeGroup [%s]: Spec=%s Image=%s | %s",
				ng.Name, ng.SpecId, ng.ImageId, ng.Description)
		}
	}

	return res
}

func printSummary(results []RecommendationResult, infra onpremmodel.OnpremInfra) {
	fmt.Printf("\n%s\n", strings.Repeat("=", 80))
	fmt.Printf("GPU RECOMMENDATION TEST RESULTS SUMMARY\n")
	if len(infra.Nodes) > 0 {
		n := infra.Nodes[0]
		gpuInfo := "None"
		if len(n.GPUCards) > 0 {
			gpuInfo = fmt.Sprintf("%dx %s (%.0f GB VRAM)", len(n.GPUCards), n.GPUCards[0].Model, n.GPUCards[0].MemoryTotalGB)
		}
		fmt.Printf("Source Workload: Host '%s' (%d vCPU, %d GiB RAM, %s, %s)\n",
			n.Hostname, n.CPU.Cpus*n.CPU.Threads, n.Memory.TotalSize, n.OS.PrettyName, gpuInfo)
	}
	fmt.Printf("%s\n", strings.Repeat("=", 80))

	passed := 0
	failed := 0

	for _, r := range results {
		status := "PASS"
		if !r.Success {
			status = "FAIL"
			failed++
		} else {
			passed++
		}

		candDesc := "No candidates"
		if len(r.Candidates) > 0 {
			candDesc = r.Candidates[0].Description
		}

		fmt.Printf("[%s] %-18s (%s/%s) | Duration: %7v\n", status, r.DisplayName, r.CSP, r.Region, r.Duration.Round(time.Millisecond))
		if r.Success {
			fmt.Printf("       Top Candidate: %s\n", candDesc)
			if len(r.Candidates) > 0 && len(r.Candidates[0].TargetInfra.NodeGroups) > 0 {
				ng := r.Candidates[0].TargetInfra.NodeGroups[0]
				fmt.Printf("       Node Detail  : %s\n", ng.Description)
			}
		} else {
			fmt.Printf("       Error        : %s\n", r.ErrorMessage)
		}
		fmt.Println()
	}

	fmt.Printf("%s\n", strings.Repeat("-", 80))
	fmt.Printf("Total Targets: %d  |  Passed: %d  |  Failed: %d\n", len(results), passed, failed)
	fmt.Printf("%s\n\n", strings.Repeat("=", 80))
}

func writeSummaryReport(results []RecommendationResult) error {
	dir := "testresult"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	reportPath := filepath.Join(dir, "gpu-recommendation-test-report.md")
	f, err := os.Create(reportPath)
	if err != nil {
		return err
	}
	defer f.Close()

	var sb strings.Builder
	sb.WriteString("# CM-Beetle GPU Recommendation & Match Rate Test Report\n\n")
	sb.WriteString(fmt.Sprintf("- **Generated At**: %s\n", time.Now().Format("2006-01-02 15:04:05 MST")))
	sb.WriteString("- **Source Workload**: 1x NVIDIA Tesla T4 (16GB VRAM, 4 vCPU, 16 GiB RAM, Ubuntu 22.04 LTS)\n\n")

	sb.WriteString("## Target CSP Summary\n\n")
	sb.WriteString("| Status | Display Name | CSP | Region | Duration | Candidates | Top Candidate Status |\n")
	sb.WriteString("|:------:|:-------------|:----|:-------|:---------|:----------:|:---------------------|\n")

	for _, r := range results {
		status := "✅ PASS"
		candCount := len(r.Candidates)
		topStatus := "N/A"
		if !r.Success {
			status = "❌ FAIL"
		} else if candCount > 0 {
			topStatus = r.Candidates[0].Status
		}

		sb.WriteString(fmt.Sprintf("| %s | %s | `%s` | `%s` | %v | %d | `%s` |\n",
			status, r.DisplayName, r.CSP, r.Region, r.Duration.Round(time.Millisecond), candCount, topStatus))
	}

	sb.WriteString("\n## Detailed Recommendation Results\n\n")
	for idx, r := range results {
		sb.WriteString(fmt.Sprintf("### %d. %s (`%s` / `%s`)\n\n", idx+1, r.DisplayName, r.CSP, r.Region))
		if !r.Success {
			sb.WriteString(fmt.Sprintf("> [!WARNING]\n> **Error**: %s\n\n", r.ErrorMessage))
			continue
		}

		if len(r.Candidates) == 0 {
			sb.WriteString("> No recommendation candidates returned.\n\n")
			continue
		}

		for cIdx, cand := range r.Candidates {
			sb.WriteString(fmt.Sprintf("#### Candidate #%d (`%s`)\n\n", cIdx+1, cand.Status))
			sb.WriteString(fmt.Sprintf("- **Description**: %s\n", cand.Description))
			sb.WriteString(fmt.Sprintf("- **Target VNet**: `%s` (Subnets: %d)\n", cand.TargetVNet.Name, len(cand.TargetVNet.SubnetInfoList)))
			sb.WriteString("- **Recommended Nodes**:\n")
			for _, ng := range cand.TargetInfra.NodeGroups {
				sb.WriteString(fmt.Sprintf("  - **Node**: `%s` | **Spec**: `%s` | **Image**: `%s`\n", ng.Name, ng.SpecId, ng.ImageId))
				sb.WriteString(fmt.Sprintf("    - *Detail*: %s\n", ng.Description))
			}
			sb.WriteString("\n")
		}
	}

	_, err = f.WriteString(sb.String())
	return err
}

func loadConfig(path string) (TestConfig, error) {
	var cfg TestConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}

func loadAuthConfig(path string) (AuthConfig, error) {
	var cfg AuthConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func loadOnpremInfra(path string) (onpremmodel.OnpremInfra, error) {
	var infra onpremmodel.OnpremInfra
	data, err := os.ReadFile(path)
	if err != nil {
		return infra, err
	}

	var wrapped struct {
		OnpremiseInfraModel onpremmodel.OnpremInfra `json:"onpremiseInfraModel"`
	}
	if err := json.Unmarshal(data, &wrapped); err == nil && len(wrapped.OnpremiseInfraModel.Nodes) > 0 {
		return wrapped.OnpremiseInfraModel, nil
	}

	err = json.Unmarshal(data, &infra)
	return infra, err
}

func checkBeetleReadiness(client *resty.Client, endpoint string) error {
	url := fmt.Sprintf("%s/beetle/readyz", endpoint)
	var resp map[string]interface{}
	var empty interface{} = common.NoBody

	return common.ExecuteHttpRequest(
		client, "GET", url, nil, common.SetUseBody(empty), &empty, &resp, 0,
	)
}
