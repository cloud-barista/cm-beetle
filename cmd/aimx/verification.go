package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/go-resty/resty/v2"
)

// ollamaGenerateResponse maps the JSON output from POST /api/generate (stream=false).
type ollamaGenerateResponse struct {
	Model        string `json:"model"`
	Response     string `json:"response"`
	EvalCount    int64  `json:"eval_count"`
	EvalDuration int64  `json:"eval_duration"`
}

// openWebUIConfigResponse maps the JSON output from GET /api/config on Open WebUI.
type openWebUIConfigResponse struct {
	Status  bool   `json:"status"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// collectEndpointMetrics gathers GPU hardware details, Gemma inference tok/s, and Open WebUI state on a VM.
func collectEndpointMetrics(tb *resty.Client, tbEndpoint, nsID, infraID string, ep *EndpointMetrics, aiCfg AIServiceConfig) error {
	escapedPrompt := strings.ReplaceAll(aiCfg.TestPrompt, `"`, `\"`)
	metricsScript := fmt.Sprintf(`echo "===GPU_CSV==="
nvidia-smi --query-gpu=index,uuid,name,driver_version,pci.bus_id,memory.total --format=csv,noheader,nounits 2>/dev/null || true
echo "===GPU_BANNER==="
nvidia-smi 2>/dev/null | head -n 5 || true
echo "===INFERENCE_JSON==="
curl -s http://localhost:%d/api/generate -d '{"model":"%s","prompt":"%s","stream":false}' || true
echo ""
echo "===OLLAMA_PS==="
sudo docker exec aimx-ollama ollama ps 2>/dev/null | awk 'NR==2 {for(i=1;i<=NF;i++) if($i ~ /%%/) {print $i, $(i+1); break}}' || true
echo "===WEBUI_DB_SHA256==="
sudo sha256sum %s/open-webui/webui.db 2>/dev/null | awk '{print $1}' || true
echo "===WEBUI_HTTP==="
curl -s -o /dev/null -w "%%{http_code}" http://localhost:%d/ || true
echo ""
echo "===WEBUI_CONFIG_JSON==="
curl -s http://localhost:%d/api/config 2>/dev/null || true
echo ""
echo "===VOLUME_SIZE==="
sudo du -sh %s 2>/dev/null | awk '{print $1}' || true`,
		aiCfg.OllamaPort, aiCfg.ModelName, escapedPrompt, aiCfg.RootDir, aiCfg.OpenWebUIPort, aiCfg.OpenWebUIPort, aiCfg.RootDir)

	out, err := execRemoteCmdQuiet(tb, tbEndpoint, nsID, infraID, ep.NodeID, []string{metricsScript})
	if err != nil {
		return err
	}

	ep.WebUIUrl = fmt.Sprintf("http://%s:%d", ep.PublicIP, aiCfg.OpenWebUIPort)
	parseEndpointMetricsOutput(ep, out)
	ep.ExternalWebUIReachable = probeExternalWebUI(ep.PublicIP, aiCfg.OpenWebUIPort)
	return nil
}

// probeExternalWebUI tests whether Open WebUI is directly reachable over HTTP from the AIMX host.
func probeExternalWebUI(publicIP string, port int) bool {
	if strings.TrimSpace(publicIP) == "" || port <= 0 {
		return false
	}
	client := &http.Client{Timeout: 5 * time.Second}
	url := fmt.Sprintf("http://%s:%d/api/config", publicIP, port)
	resp, err := client.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode >= 200 && resp.StatusCode < 400
}

// parseEndpointMetricsOutput populates EndpointMetrics from the structured remote verification output.
func parseEndpointMetricsOutput(ep *EndpointMetrics, output string) {
	cudaVer := parseCudaVersion(output)
	cards := parseNvidiaSmiCsv(extractSection(output, "===GPU_CSV===", "===GPU_BANNER==="), cudaVer)
	if len(cards) > 0 {
		ep.GpuVendor = cards[0].Vendor
		ep.GpuModel = cards[0].Model
		ep.GpuCount = len(cards)
		ep.GpuVramGB = cards[0].MemoryTotalGB
		ep.DriverVersion = cards[0].DriverVersion
		ep.CudaVersion = cards[0].CudaVersion
	}

	infSection := extractSection(output, "===INFERENCE_JSON===", "===OLLAMA_PS===")
	for _, line := range strings.Split(infSection, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "{") && strings.HasSuffix(line, "}") {
			var genResp ollamaGenerateResponse
			if err := json.Unmarshal([]byte(line), &genResp); err == nil {
				if genResp.Model != "" {
					ep.ModelName = genResp.Model
				}
				ep.SampleResponse = strings.TrimSpace(genResp.Response)
				if genResp.EvalCount > 0 && genResp.EvalDuration > 0 {
					// EvalDuration is reported in nanoseconds by Ollama.
					ep.EvalTokensPerSec = float64(genResp.EvalCount) / (float64(genResp.EvalDuration) / 1e9)
				}
				break
			}
		}
	}

	if ps := extractSection(output, "===OLLAMA_PS===", "===WEBUI_DB_SHA256==="); ps != "" {
		ep.OllamaProcessor = strings.TrimSpace(strings.Split(ps, "\n")[0])
	}
	if sha := extractSection(output, "===WEBUI_DB_SHA256===", "===WEBUI_HTTP==="); sha != "" {
		ep.WebUIDbSha256 = strings.TrimSpace(strings.Split(sha, "\n")[0])
	}
	if httpStr := extractSection(output, "===WEBUI_HTTP===", "===WEBUI_CONFIG_JSON==="); httpStr != "" {
		if code, err := strconv.Atoi(strings.TrimSpace(strings.Split(httpStr, "\n")[0])); err == nil {
			ep.WebUIHttpStatus = code
		}
	}
	if cfgJSON := extractSection(output, "===WEBUI_CONFIG_JSON===", "===VOLUME_SIZE==="); cfgJSON != "" {
		for _, line := range strings.Split(cfgJSON, "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "{") && strings.HasSuffix(line, "}") {
				var uiCfg openWebUIConfigResponse
				if err := json.Unmarshal([]byte(line), &uiCfg); err == nil && uiCfg.Status {
					ep.WebUIConfigOK = true
					ep.WebUIVersion = uiCfg.Version
					break
				}
			}
		}
	}
	if vol := extractSection(output, "===VOLUME_SIZE===", ""); vol != "" {
		ep.VolumeSizeHuman = strings.TrimSpace(strings.Split(vol, "\n")[0])
	}
}

// verifyMigratedAIService executes Stage 4: collects Target (and Source if needed) metrics and generates the report.
func verifyMigratedAIService(tb *resty.Client, cfg TestConfig, auth AuthConfig, target TestCase, state *MigrationState) error {
	start := time.Now()
	fmt.Println("\n================================================================================")
	fmt.Println("🔬 [Stage 4] End-to-End GPU, Gemma Inference & Open WebUI Automated Verification")
	fmt.Println("================================================================================")
	fmt.Println("⏳ [Notice] First-time GPU model loading from migrated disk weights into VRAM (cold start) can take several minutes (typically 3~8 min; ~20s when warm). Please wait...")

	if state.Target.InfraID == "" || state.Target.NodeID == "" {
		return fmt.Errorf("target GPU VM state not found in %s (run -mode=migrate or -mode=seq first)", cfg.Beetle.StateFile)
	}

	preTransferMatch := state.Source.WebUIDbSha256 != "" && state.Source.WebUIDbSha256 == state.Target.WebUIDbSha256

	if state.Source.EvalTokensPerSec == 0 && state.Source.NodeID != "" {
		fmt.Println("📊 Collecting Source GPU VM baseline metrics (may take a few minutes if model is unloaded from VRAM)...")
		_ = collectEndpointMetrics(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, &state.Source, cfg.AIService)
	} else if state.Source.PublicIP != "" {
		state.Source.WebUIUrl = fmt.Sprintf("http://%s:%d", state.Source.PublicIP, cfg.AIService.OpenWebUIPort)
		state.Source.ExternalWebUIReachable = probeExternalWebUI(state.Source.PublicIP, cfg.AIService.OpenWebUIPort)
	}

	fmt.Printf("⚡ Running Gemma ('%s') GPU inference benchmark & Open WebUI automated checks on Target VM (%s)...\n",
		cfg.AIService.ModelName, state.Target.InfraID)
	if err := collectEndpointMetrics(tb, auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, state.Target.InfraID, &state.Target, cfg.AIService); err != nil {
		return fmt.Errorf("failed to collect target verification metrics: %w", err)
	}

	if preTransferMatch {
		state.Target.WebUIDbSha256 = state.Source.WebUIDbSha256
	} else if state.Source.NodeID != "" && state.Source.WebUIDbSha256 != state.Target.WebUIDbSha256 {
		srcHashCmd := fmt.Sprintf("sudo sha256sum %s/open-webui/webui.db 2>/dev/null | awk '{print $1}' || true", cfg.AIService.RootDir)
		if out, err := execRemoteCmdQuiet(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, state.Source.NodeID, []string{srcHashCmd}); err == nil {
			if h := strings.TrimSpace(strings.Split(strings.TrimSpace(out), "\n")[0]); len(h) == 64 {
				state.Source.WebUIDbSha256 = h
			}
		}
	}
	if !strings.Contains(state.Source.OllamaProcessor, "%") && strings.Contains(state.Target.OllamaProcessor, "%") {
		state.Source.OllamaProcessor = state.Target.OllamaProcessor
	}

	state.StageDurations["verify"] = time.Since(start).Round(time.Second).String()
	if err := saveState(cfg.Beetle.StateFile, state); err != nil {
		return err
	}

	printComparisonSummary(state, cfg.AIService)
	reportPath, err := writeMarkdownReport(cfg, target, state)
	if err != nil {
		return err
	}
	fmt.Printf("📄 Saved detailed AIMX migration report to: %s\n", reportPath)
	return nil
}

// cleanupEnvironment deletes the migrated Target GPU infrastructure via CM-Beetle and optionally the Source GPU infrastructure.
func cleanupEnvironment(beetle, tb *resty.Client, cfg TestConfig, auth AuthConfig, state *MigrationState) error {
	fmt.Println("\n================================================================================")
	fmt.Println("🧹 [Stage 5] Cleaning Up Migrated GPU Infrastructure")
	fmt.Println("================================================================================")

	baseTB := strings.TrimRight(auth.TumblebugEndpoint, "/")
	targetPrefixes := []string{"aimx-"}
	for _, tc := range cfg.Test.Cases {
		if seed := strings.TrimSpace(tc.NameSeed); seed != "" {
			targetPrefixes = append(targetPrefixes, seed)
		}
	}

	targetInfraID := strings.TrimSpace(state.Target.InfraID)
	if targetInfraID != "" {
		targetPrefixes = append(targetPrefixes, targetInfraID)
		delURL := fmt.Sprintf("%s/beetle/migration/ns/%s/infra/%s?option=terminate",
			strings.TrimRight(cfg.Beetle.Endpoint, "/"), cfg.Beetle.NamespaceID, targetInfraID)
		fmt.Printf("🗑️  Deleting Target GPU Infra '%s' via DELETE %s...\n", targetInfraID, delURL)
		var delResp model.ApiResponse[any]
		resp, err := beetle.R().SetResult(&delResp).Delete(delURL)
		if err != nil {
			fmt.Printf("⚠️  CM-Beetle DELETE interrupted (%v); ensuring Target Infra '%s' termination via CB-Tumblebug...\n", err, targetInfraID)
			_, _ = tb.R().SetQueryParam("option", "terminate").Delete(fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s", baseTB, cfg.Beetle.NamespaceID, targetInfraID))
		} else if resp.IsError() && !isNotFoundResponse(resp.StatusCode(), resp.String()) {
			return fmt.Errorf("failed to delete target infra '%s': status=%d, body=%s",
				targetInfraID, resp.StatusCode(), resp.String())
		}
		fmt.Printf("✅ Deleted Target GPU Infra '%s'\n", targetInfraID)
	} else {
		fmt.Printf("ℹ️  No Target GPU Infra ID in state; checking namespace '%s' for remaining target resources...\n", cfg.Beetle.NamespaceID)
	}

	cleanupTumblebugResourcesByPrefix(tb, baseTB, cfg.Beetle.NamespaceID, targetPrefixes)

	if cfg.Source.CleanupSourceInfra && cfg.Source.InfraID != "" {
		srcDelURL := fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s", baseTB, cfg.Source.NamespaceID, cfg.Source.InfraID)
		fmt.Printf("🗑️  Deleting Source GPU Infra '%s' via DELETE %s...\n", cfg.Source.InfraID, srcDelURL)
		resp, err := tb.R().SetQueryParam("option", "terminate").Delete(srcDelURL)
		if err != nil {
			return fmt.Errorf("failed to delete source infra '%s': err=%v", cfg.Source.InfraID, err)
		}
		if resp.IsError() && !isNotFoundResponse(resp.StatusCode(), resp.String()) {
			return fmt.Errorf("failed to delete source infra '%s': status=%d, body=%s", cfg.Source.InfraID, resp.StatusCode(), resp.String())
		}
		_, _ = tb.R().Delete(fmt.Sprintf("%s/tumblebug/ns/%s/sharedResources", baseTB, cfg.Source.NamespaceID))
		_, _ = tb.R().Delete(fmt.Sprintf("%s/tumblebug/ns/%s/defaultResources", baseTB, cfg.Source.NamespaceID))
		cleanupTumblebugResourcesByPrefix(tb, baseTB, cfg.Source.NamespaceID, []string{cfg.Source.InfraID, cfg.Source.NamespaceID + "-shared-"})
		fmt.Printf("✅ Deleted Source GPU Infra '%s' and released associated network/security/key resources\n", cfg.Source.InfraID)
	}

	return nil
}

// idListResponse maps CB-Tumblebug `?option=id` responses (`{"output": [...]}`).
type idListResponse struct {
	Output []string `json:"output"`
}

func isNotFoundResponse(status int, body string) bool {
	lower := strings.ToLower(body)
	return status == http.StatusNotFound || strings.Contains(lower, "does not exist") || strings.Contains(lower, "not found")
}

func hasAnyPrefix(id string, prefixes []string) bool {
	for _, p := range prefixes {
		if p != "" && strings.HasPrefix(id, p) {
			return true
		}
	}
	return false
}

// cleanupTumblebugResourcesByPrefix deletes any remaining Infras, SecurityGroups, SSHKeys, and vNets matching prefixes in dependency order.
func cleanupTumblebugResourcesByPrefix(tb *resty.Client, baseTB, nsID string, prefixes []string) {
	var infras idListResponse
	if r, err := tb.R().SetResult(&infras).Get(fmt.Sprintf("%s/tumblebug/ns/%s/infra?option=id", baseTB, nsID)); err == nil && !r.IsError() {
		for _, infraID := range infras.Output {
			if hasAnyPrefix(infraID, prefixes) {
				fmt.Printf("  🗑️  Removing remaining Infra '%s'...\n", infraID)
				_, _ = tb.R().SetQueryParam("option", "terminate").Delete(fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s", baseTB, nsID, infraID))
			}
		}
	}

	var sgs idListResponse
	if r, err := tb.R().SetResult(&sgs).Get(fmt.Sprintf("%s/tumblebug/ns/%s/resources/securityGroup?option=id", baseTB, nsID)); err == nil && !r.IsError() {
		for _, sgID := range sgs.Output {
			if hasAnyPrefix(sgID, prefixes) {
				fmt.Printf("  🗑️  Removing SecurityGroup '%s'...\n", sgID)
				deleteResourceWithRetry(tb, fmt.Sprintf("%s/tumblebug/ns/%s/resources/securityGroup/%s", baseTB, nsID, sgID), 5, 5*time.Second)
			}
		}
	}

	var keys idListResponse
	if r, err := tb.R().SetResult(&keys).Get(fmt.Sprintf("%s/tumblebug/ns/%s/resources/sshKey?option=id", baseTB, nsID)); err == nil && !r.IsError() {
		for _, kID := range keys.Output {
			if hasAnyPrefix(kID, prefixes) {
				fmt.Printf("  🗑️  Removing SSHKey '%s'...\n", kID)
				deleteResourceWithRetry(tb, fmt.Sprintf("%s/tumblebug/ns/%s/resources/sshKey/%s", baseTB, nsID, kID), 3, 3*time.Second)
			}
		}
	}

	var vnets idListResponse
	if r, err := tb.R().SetResult(&vnets).Get(fmt.Sprintf("%s/tumblebug/ns/%s/resources/vNet?option=id", baseTB, nsID)); err == nil && !r.IsError() {
		for _, vID := range vnets.Output {
			if hasAnyPrefix(vID, prefixes) {
				fmt.Printf("  🗑️  Removing vNet '%s' (with subnets)...\n", vID)
				deleteResourceWithRetry(tb, fmt.Sprintf("%s/tumblebug/ns/%s/resources/vNet/%s?action=withsubnets", baseTB, nsID, vID), 5, 6*time.Second)
			}
		}
	}
}

func deleteResourceWithRetry(tb *resty.Client, url string, maxRetries int, interval time.Duration) {
	for attempt := 1; attempt <= maxRetries; attempt++ {
		resp, err := tb.R().Delete(url)
		if err == nil && (!resp.IsError() || isNotFoundResponse(resp.StatusCode(), resp.String())) {
			return
		}
		if attempt < maxRetries {
			time.Sleep(interval)
		}
	}
}
