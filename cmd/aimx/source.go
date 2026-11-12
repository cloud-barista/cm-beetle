package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/go-resty/resty/v2"
)

//go:embed scripts/ensure-gpu-runtime.sh
var ensureGpuRuntimeScript string

// cudaRegex extracts the CUDA version reported in the standard nvidia-smi banner.
var cudaRegex = regexp.MustCompile(`CUDA Version:\s*([0-9.]+)`)

// setupSourceEnvironment provisions or reuses the source GPU VM and seeds Ollama + the configured AI web service.
func setupSourceEnvironment(tb *resty.Client, cfg TestConfig, auth AuthConfig, state *MigrationState) error {
	if cfg.Source.isExternalSSH() {
		return verifyExternalSourceEnvironment(tb, cfg, auth, state)
	}

	start := time.Now()
	appName := cfg.AIService.appDisplayName()
	appPort := cfg.AIService.activeAppPort()

	fmt.Println("\n================================================================================")
	fmt.Printf("🚀 [Stage 1] Source GPU VM Provisioning & AI Service (Ollama + %s) Setup\n", appName)
	fmt.Println("================================================================================")

	infraInfo, created, err := ensureSourceGpuInfra(tb, cfg, auth)
	if err != nil {
		return err
	}
	if created {
		waitWithCountdown(cfg.Timing.VmBootStabilizeSeconds, "Waiting for newly provisioned Source GPU VM boot & NVIDIA driver init")
	}

	metrics, sgIDs, err := extractNodeEndpointInfo(infraInfo, cfg.AIService)
	if err != nil {
		return err
	}
	fmt.Printf("📌 Source GPU Node: id=%s, ip=%s, spec=%s, image=%s\n",
		metrics.NodeID, metrics.PublicIP, metrics.SpecID, metrics.ImageID)

	if cfg.Source.OpenServicePorts && len(sgIDs) > 0 {
		openSecurityGroupServicePorts(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, sgIDs, cfg.AIService)
	}

	if err := pollVmSshReady(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.Timing); err != nil {
		return err
	}

	if err := deployOllamaContainerFirst(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.AIService, cfg.Timing); err != nil {
		return err
	}

	if err := pullOllamaModels(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.AIService, "Source"); err != nil {
		return err
	}

	fmt.Printf("🐳 Starting %s container on Source GPU VM (:%d)...\n", appName, appPort)
	upCmd := buildDockerComposeUpWithRetryCmd(cfg.AIService.RootDir, "")
	if _, err := execRemoteCmd(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, []string{upCmd}); err != nil {
		return fmt.Errorf("failed to start %s container on source VM: %w", appName, err)
	}

	waitWithCountdown(cfg.Timing.ServiceInitWaitSeconds, fmt.Sprintf("Waiting for Source Ollama & %s containers to initialize", appName))
	if err := pollServiceEndpointsReady(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.AIService, cfg.Timing); err != nil {
		return err
	}

	if cfg.AIService.isKotaemon() {
		ensureKotaemonLangdetectCache(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.AIService)
		if err := ensureKotaemonOllamaDefaults(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.AIService, cfg.Timing); err != nil {
			fmt.Printf("⚠️  Kotaemon auto-default check warning: %v (verify 'Resources' tab in browser if needed)\n", err)
		}
	}

	warmupOllamaModels(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.AIService)

	fmt.Println("⚡ Running baseline GPU inference & collecting Source metrics...")
	if err := collectEndpointMetrics(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, &metrics, cfg.AIService); err != nil {
		return fmt.Errorf("failed to collect baseline source metrics: %w", err)
	}

	metrics.WebUIUrl = fmt.Sprintf("http://%s:%d", metrics.PublicIP, appPort)
	state.Source = metrics
	state.StageDurations["setup-source"] = time.Since(start).Round(time.Second).String()
	if err := saveState(resolveConfigPath(cfg.ConfigDir, cfg.Beetle.StateFile), state); err != nil {
		return err
	}

	fmt.Printf("✅ [Stage 1 Complete] Source GPU='%s' (%.0fGB), Models='%s', Speed=%.2f tok/s\n",
		metrics.GpuModel, metrics.GpuVramGB, cfg.AIService.modelSummary(), metrics.EvalTokensPerSec)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("🌐 [Source %s Ready] Browser Link : %s\n", appName, metrics.WebUIUrl)
	fmt.Printf("🌐 [Source Ollama API Ready] API Endpoint : http://%s:%d/api/tags\n", metrics.PublicIP, cfg.AIService.OllamaPort)
	if cfg.AIService.isKotaemon() {
		fmt.Printf("🔐 [Kotaemon Login] Username: %s | Password: %s\n", cfg.AIService.AdminUsername, cfg.AIService.AdminPassword)
		fmt.Println("📄 [Next Demo Action before Migration]:")
		fmt.Println("   1) Open the Source Kotaemon link above and sign in.")
		fmt.Println("   2) Upload Cloud-Barista conference PDF slides in the 'File Index' tab.")
		fmt.Println("   3) Ask 2-3 questions to verify slide citations with yellow highlights on the right panel.")
		fmt.Println("   4) Run '-mode recommend' -> '-mode migrate' -> '-mode verify' (or migrate via Beetle UX Lab UI).")
	}
	fmt.Println("--------------------------------------------------------------------------------")
	return nil
}

// verifyExternalSourceEnvironment inspects and validates an existing external SSH source GPU VM (e.g., mc-terrarium ai-stack-testbed).
func verifyExternalSourceEnvironment(tb *resty.Client, cfg TestConfig, auth AuthConfig, state *MigrationState) error {
	start := time.Now()
	appName := cfg.AIService.appDisplayName()
	appPort := cfg.AIService.activeAppPort()

	fmt.Println("\n================================================================================")
	fmt.Printf("🔎 [Stage 1] External Source GPU VM Verification (%s + Ollama)\n", appName)
	fmt.Println("================================================================================")

	if err := enrichNodeFromExternalSource(cfg, state, nil); err != nil {
		return err
	}
	fmt.Printf("📌 External Source GPU Node: node=%s, ip=%s, csp=%s/%s, spec=%s\n",
		state.Source.NodeID, state.Source.PublicIP, state.Source.Csp, state.Source.Region, state.Source.SpecID)

	checkCmd := fmt.Sprintf(
		"nvidia-smi --query-gpu=name,memory.total,driver_version --format=csv,noheader && curl -sf http://localhost:%d/api/tags >/dev/null && code=$(curl -s -o /dev/null -w '%%{http_code}' http://localhost:%d/) && [ \"$code\" -ge 200 ] && [ \"$code\" -lt 400 ] && echo \"EXTERNAL_SOURCE_READY:$code\"",
		cfg.AIService.OllamaPort, appPort,
	)
	out, err := execExternalSshCmd(cfg, checkCmd, false)
	if err != nil || !strings.Contains(out, "EXTERNAL_SOURCE_READY:") {
		return fmt.Errorf("external source GPU VM '%s' (%s) is not ready: err=%v, out=%s", state.Source.NodeID, state.Source.PublicIP, err, out)
	}
	fmt.Printf("✅ External Source GPU & AI stack healthy: %s\n", strings.Split(out, "\n")[0])

	if cfg.AIService.isKotaemon() {
		_, _ = execExternalSshCmd(cfg, buildLangdetectCacheScript(cfg.AIService.RootDir), false)
	}

	fmt.Println("⚡ Running baseline GPU inference & collecting External Source metrics...")
	if err := collectSourceEndpointMetrics(tb, cfg, auth, &state.Source); err != nil {
		return fmt.Errorf("failed to collect external source baseline metrics: %w", err)
	}

	state.Source.WebUIUrl = fmt.Sprintf("http://%s:%d", state.Source.PublicIP, appPort)
	state.StageDurations["setup-source"] = time.Since(start).Round(time.Second).String()
	if err := saveState(resolveConfigPath(cfg.ConfigDir, cfg.Beetle.StateFile), state); err != nil {
		return err
	}

	fmt.Printf("✅ [Stage 1 Complete] External Source GPU='%s' (%.0fGB), Models='%s', Speed=%.2f tok/s\n",
		state.Source.GpuModel, state.Source.GpuVramGB, cfg.AIService.modelSummary(), state.Source.EvalTokensPerSec)
	return nil
}

// isInfraHealthyAndRunning returns true if the Infra has at least one Running node with an assigned PublicIP.
func isInfraHealthyAndRunning(infra tbmodel.InfraInfo) bool {
	if infra.Id == "" || len(infra.Node) == 0 {
		return false
	}
	for _, n := range infra.Node {
		if strings.EqualFold(n.Status, "Running") && strings.TrimSpace(n.PublicIP) != "" {
			return true
		}
	}
	return false
}

// cleanupFailedInfra terminates an incomplete or failed Infra and releases unused default resources.
func cleanupFailedInfra(tb *resty.Client, baseURL, nsID, infraID string) error {
	if strings.TrimSpace(infraID) == "" {
		return nil
	}
	refineURL := fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s?action=refine", baseURL, nsID, infraID)
	_, _ = tb.R().Get(refineURL)

	delURL := fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s", baseURL, nsID, infraID)
	termResp, err := tb.R().SetQueryParam("option", "terminate").Delete(delURL)
	if err != nil || termResp.IsError() {
		return fmt.Errorf("failed to terminate incomplete infra '%s': err=%v, status=%d, body=%s",
			infraID, err, termResp.StatusCode(), termResp.String())
	}
	time.Sleep(3 * time.Second)

	defResURL := fmt.Sprintf("%s/tumblebug/ns/%s/defaultResources", baseURL, nsID)
	_, _ = tb.R().Delete(defResURL)
	return nil
}

// ensureSourceGpuInfra retrieves an existing source Infra or provisions one directly via POST /tumblebug/ns/{nsId}/infraDynamic.
func ensureSourceGpuInfra(tb *resty.Client, cfg TestConfig, auth AuthConfig) (tbmodel.InfraInfo, bool, error) {
	baseURL := strings.TrimRight(auth.TumblebugEndpoint, "/")
	getURL := fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s", baseURL, cfg.Source.NamespaceID, cfg.Source.InfraID)

	var existing tbmodel.InfraInfo
	resp, err := tb.R().SetResult(&existing).Get(getURL)
	if err == nil && !resp.IsError() && existing.Id != "" {
		if isInfraHealthyAndRunning(existing) {
			fmt.Printf("✅ Found existing running Source GPU Infra '%s' in namespace '%s' (%d node(s), ip=%s)\n",
				existing.Id, cfg.Source.NamespaceID, len(existing.Node), existing.Node[0].PublicIP)
			return existing, false, nil
		}
		fmt.Printf("⚠️  Found incomplete/failed Infra '%s' (status=%s); cleaning it up before provisioning...\n",
			existing.Id, existing.Status)
		if err := cleanupFailedInfra(tb, baseURL, cfg.Source.NamespaceID, cfg.Source.InfraID); err != nil {
			return tbmodel.InfraInfo{}, false, err
		}
	}

	if !cfg.Source.AutoProvisionIfMissing {
		return tbmodel.InfraInfo{}, false, fmt.Errorf("source infra '%s' not found in namespace '%s' and autoProvisionIfMissing is false",
			cfg.Source.InfraID, cfg.Source.NamespaceID)
	}

	dynReq := tbmodel.InfraDynamicReq{
		Name: cfg.Source.InfraID,
		NodeGroups: []tbmodel.CreateNodeGroupDynamicReq{
			{
				Name:          cfg.Source.NodeGroupName,
				NodeGroupSize: 1,
				SpecId:        cfg.Source.SpecID,
				ImageId:       cfg.Source.ImageID,
				RootDiskType:  cfg.Source.RootDiskType,
				RootDiskSize:  cfg.Source.RootDiskSize,
				Label:         map[string]string{"accelerator": "gpu"},
			},
		},
	}

	fmt.Printf("🔧 Source GPU NodeGroup '%s' (spec=%s, image=%s, disk=%s %dGB)\n",
		cfg.Source.NodeGroupName, cfg.Source.SpecID, cfg.Source.ImageID, cfg.Source.RootDiskType, cfg.Source.RootDiskSize)

	fmt.Printf("🚀 Provisioning Source GPU Infra '%s' via POST /tumblebug/ns/%s/infraDynamic...\n",
		cfg.Source.InfraID, cfg.Source.NamespaceID)
	createURL := fmt.Sprintf("%s/tumblebug/ns/%s/infraDynamic", baseURL, cfg.Source.NamespaceID)
	var created tbmodel.InfraInfo
	createResp, err := tb.R().SetBody(&dynReq).SetResult(&created).Post(createURL)
	if err != nil || createResp.IsError() || !isInfraHealthyAndRunning(created) {
		return tbmodel.InfraInfo{}, false, fmt.Errorf("failed to create source GPU infra '%s': err=%v, status=%d, body=%s",
			cfg.Source.InfraID, err, createResp.StatusCode(), createResp.String())
	}

	fmt.Printf("✅ Source GPU Infra '%s' provisioned successfully (spec=%s, ip=%s)!\n",
		created.Id, cfg.Source.SpecID, created.Node[0].PublicIP)
	return created, true, nil
}

// extractNodeEndpointInfo extracts connection and resource identifiers from the first node of a Tumblebug InfraInfo.
func extractNodeEndpointInfo(infraInfo tbmodel.InfraInfo, aiCfg AIServiceConfig) (EndpointMetrics, []string, error) {
	if len(infraInfo.Node) == 0 {
		return EndpointMetrics{}, nil, fmt.Errorf("infra '%s' has no nodes", infraInfo.Id)
	}
	n := infraInfo.Node[0]
	return EndpointMetrics{
		InfraID:        infraInfo.Id,
		NodeID:         n.Id,
		PublicIP:       n.PublicIP,
		SshKeyID:       n.SshKeyId,
		SshUsername:    n.NodeUserName,
		Csp:            n.ConnectionConfig.ProviderName,
		Region:         n.Region.Region,
		SpecID:         n.SpecId,
		ImageID:        n.ImageId,
		AppType:        aiCfg.AppType,
		ModelName:      aiCfg.ModelName,
		EmbeddingModel: aiCfg.EmbeddingModel,
	}, n.SecurityGroupIds, nil
}

// openSecurityGroupServicePorts ensures Ollama and active web app TCP ports are open on the Security Group via Tumblebug.
func openSecurityGroupServicePorts(tb *resty.Client, tbEndpoint, nsID string, sgIDs []string, aiCfg AIServiceConfig) {
	baseURL := strings.TrimRight(tbEndpoint, "/")
	requiredPorts := []int{aiCfg.activeAppPort(), aiCfg.OllamaPort}

	for _, sgID := range sgIDs {
		sgURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/securityGroup/%s", baseURL, nsID, sgID)
		var sgInfo tbmodel.SecurityGroupInfo
		_, _ = tb.R().SetResult(&sgInfo).Get(sgURL)

		var missingRules []tbmodel.FirewallRuleReq
		for _, port := range requiredPorts {
			if port <= 0 || isPortCoveredByRules(sgInfo.FirewallRules, port) {
				continue
			}
			missingRules = append(missingRules, tbmodel.FirewallRuleReq{
				Ports:     strconv.Itoa(port),
				Protocol:  "TCP",
				Direction: "inbound",
				CIDR:      "0.0.0.0/0",
			})
		}
		if len(missingRules) == 0 {
			continue
		}

		rulesURL := fmt.Sprintf("%s/tumblebug/ns/%s/resources/securityGroup/%s/rules", baseURL, nsID, sgID)
		updateReq := tbmodel.SecurityGroupUpdateReq{FirewallRules: missingRules}
		_, _ = tb.R().SetHeader("Content-Type", "application/json").SetBody(&updateReq).Post(rulesURL)
	}
}

// isPortCoveredByRules returns true if an inbound TCP/ALL rule already covers targetPort (single port or range).
func isPortCoveredByRules(rules []tbmodel.FirewallRuleInfo, targetPort int) bool {
	for _, r := range rules {
		if !strings.EqualFold(r.Direction, "inbound") {
			continue
		}
		proto := strings.ToUpper(strings.TrimSpace(r.Protocol))
		if proto != "TCP" && proto != "ALL" && proto != "*" {
			continue
		}
		portSpec := strings.TrimSpace(r.Port)
		if portSpec == "*" || portSpec == "ALL" || portSpec == "1-65535" || portSpec == "0-65535" {
			return true
		}
		if strings.Contains(portSpec, "-") {
			parts := strings.SplitN(portSpec, "-", 2)
			startP, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
			endP, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
			if err1 == nil && err2 == nil && startP <= targetPort && targetPort <= endP {
				return true
			}
		} else if p, err := strconv.Atoi(portSpec); err == nil && p == targetPort {
			return true
		}
	}
	return false
}

// buildDockerComposeYAML generates the Docker Compose configuration for Ollama and Kotaemon or Open WebUI.
func buildDockerComposeYAML(aiCfg AIServiceConfig) string {
	appPort := aiCfg.activeAppPort()
	if aiCfg.isKotaemon() {
		return fmt.Sprintf(`services:
  ollama:
    image: ollama/ollama:latest
    container_name: aimx-ollama
    restart: unless-stopped
    ports:
      - "%d:11434"
    environment:
      - OLLAMA_LOAD_TIMEOUT=15m
      - OLLAMA_KEEP_ALIVE=30m
      - OLLAMA_MAX_LOADED_MODELS=2
      - OLLAMA_NUM_PARALLEL=2
      - OLLAMA_CONTEXT_LENGTH=16384
    volumes:
      - ./ollama:/root/.ollama
      - ./ollama/.nv:/root/.nv
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: all
              capabilities: [gpu]

  kotaemon:
    # Full image (v0.12.0) with built-in PDF viewer, citation highlighting, and advanced document/table parsers.
    image: ghcr.io/cinnamon/kotaemon:main-full
    container_name: aimx-kotaemon
    restart: unless-stopped
    ports:
      - "%d:7860"
    environment:
      - GRADIO_SERVER_NAME=0.0.0.0
      - GRADIO_SERVER_PORT=7860
      - KH_APP_VERSION=v0.12.0
      - FTLANG_CACHE=/app/ktem_app_data/fasttext-langdetect
      - KH_OLLAMA_URL=http://ollama:11434/v1/
      - LOCAL_MODEL=%s
      - LOCAL_MODEL_EMBEDDINGS=%s
      - KH_FEATURE_USER_MANAGEMENT=true
      - KH_FEATURE_USER_MANAGEMENT_ADMIN=%s
      - KH_FEATURE_USER_MANAGEMENT_PASSWORD=%s
    volumes:
      - ./ktem_app_data:/app/ktem_app_data
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: all
              capabilities: [gpu]
    depends_on:
      - ollama
`, aiCfg.OllamaPort, appPort, aiCfg.ModelName, aiCfg.EmbeddingModel, aiCfg.AdminUsername, aiCfg.AdminPassword)
	}

	return fmt.Sprintf(`services:
  ollama:
    image: ollama/ollama:latest
    container_name: aimx-ollama
    restart: unless-stopped
    ports:
      - "%d:11434"
    environment:
      - OLLAMA_LOAD_TIMEOUT=15m
      - OLLAMA_KEEP_ALIVE=30m
      - OLLAMA_CONTEXT_LENGTH=16384
    volumes:
      - ./ollama:/root/.ollama
      - ./ollama/.nv:/root/.nv
    deploy:
      resources:
        reservations:
          devices:
            - driver: nvidia
              count: all
              capabilities: [gpu]

  open-webui:
    image: ghcr.io/open-webui/open-webui:main
    container_name: aimx-open-webui
    restart: unless-stopped
    ports:
      - "%d:8080"
    environment:
      - OLLAMA_BASE_URL=http://ollama:11434
      - WEBUI_AUTH=False
    volumes:
      - ./open-webui:/app/backend/data
    depends_on:
      - ollama
`, aiCfg.OllamaPort, appPort)
}

// buildDockerComposeUpWithRetryCmd builds a 3-attempt shell command for starting Docker Compose services reliably.
func buildDockerComposeUpWithRetryCmd(rootDir, serviceName string) string {
	svcArg := strings.TrimSpace(serviceName)
	upTarget := "sudo docker compose up -d"
	if svcArg != "" {
		upTarget = "sudo docker compose up -d " + svcArg
	}
	return fmt.Sprintf(`cd %s
for attempt in 1 2 3; do
  if %s; then
    break
  fi
  if [ "$attempt" -eq 3 ]; then
    exit 1
  fi
  echo "Retrying docker compose up (attempt $((attempt+1))/3)..."
  sudo systemctl start containerd docker 2>/dev/null || true
  sleep 10
done`, rootDir, upTarget)
}

// buildPrepareMigrationHelperScript generates the helper script stored in <rootDir>/prepare-for-migration.sh.
func buildPrepareMigrationHelperScript(rootDir, dbPath string) string {
	return fmt.Sprintf(`#!/bin/bash
set -e
DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$DIR"
sudo docker compose stop && sync
sleep 3
sudo chown -R "$(id -un):$(id -gn)" "$DIR"
sudo chmod -R u+rwX,g+rX,o+rX "$DIR"
echo "QUIESCED_DB_SHA256=$(sha256sum "$DIR/%s" 2>/dev/null | awk '{print $1}')"
sudo docker compose up -d
`, dbPath)
}

// buildLangdetectCacheScript pre-populates Kotaemon's fast-langdetect model cache from the container image.
func buildLangdetectCacheScript(rootDir string) string {
	cacheDir := strings.TrimRight(rootDir, "/") + "/ktem_app_data/fasttext-langdetect"
	return fmt.Sprintf(`sudo mkdir -p "%s"
if [ ! -s "%s/lid.176.bin" ]; then
  sudo docker cp aimx-kotaemon:/app/.venv/lib/python3.10/site-packages/fast_langdetect/resources/lid.176.ftz "%s/lid.176.bin" 2>/dev/null || true
  sudo chown -R "$(id -un):$(id -gn)" "%s" 2>/dev/null || true
fi`, cacheDir, cacheDir, cacheDir, cacheDir)
}

// ensureDockerAndNvidiaRuntime verifies NVIDIA GPU driver, Docker Compose, and NVIDIA Container Toolkit across CSPs.
func ensureDockerAndNvidiaRuntime(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, timing TimingConfig) error {
	script := strings.ReplaceAll(ensureGpuRuntimeScript, "\r\n", "\n")
	out, err := execRemoteCmd(tb, tbEndpoint, nsID, infraID, nodeID, []string{script})
	if err != nil {
		return fmt.Errorf("failed to verify Docker/NVIDIA container runtime on %s: %w", infraID, err)
	}
	if strings.Contains(out, "GPU_REBOOT_NEEDED") {
		fmt.Printf("🔄 Newly installed NVIDIA kernel module on '%s' requires reboot; rebooting VM...\n", infraID)
		_, _ = execRemoteCmdQuiet(tb, tbEndpoint, nsID, infraID, nodeID, []string{"sudo nohup sh -c 'sleep 2 && reboot' >/dev/null 2>&1 &"})
		waitWithCountdown(30, fmt.Sprintf("Waiting for '%s' reboot to activate NVIDIA GPU driver", infraID))
		if _, rErr := pollGpuVmReady(tb, tbEndpoint, nsID, infraID, nodeID, timing); rErr != nil {
			return rErr
		}
		_, _ = execRemoteCmdQuiet(tb, tbEndpoint, nsID, infraID, nodeID, []string{"sudo systemctl start containerd docker"})
	}
	return nil
}

// deployOllamaContainerFirst prepares the host directories, writes docker-compose.yml, and starts Ollama before model pulls.
func deployOllamaContainerFirst(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, aiCfg AIServiceConfig, timing TimingConfig) error {
	if err := ensureDockerAndNvidiaRuntime(tb, tbEndpoint, nsID, infraID, nodeID, timing); err != nil {
		return err
	}

	composeContent := buildDockerComposeYAML(aiCfg)
	prepHelper := buildPrepareMigrationHelperScript(aiCfg.RootDir, aiCfg.relativeDbPath())
	subDir := aiCfg.appDataSubdir()
	extraDirs := fmt.Sprintf("%s/ollama/.nv", aiCfg.RootDir)
	staleContainer := "aimx-kotaemon"
	if aiCfg.isKotaemon() {
		extraDirs = fmt.Sprintf("%s/ollama/.nv %s/ktem_app_data/fasttext-langdetect", aiCfg.RootDir, aiCfg.RootDir)
		staleContainer = "aimx-open-webui"
	}

	deployCmd := fmt.Sprintf(`set -e
sudo mkdir -p %s/ollama %s/%s %s
if [ "%s" != "/opt/ai-stack" ] && [ ! -e /opt/ai-stack ]; then
  sudo ln -s "%s" /opt/ai-stack || true
fi
sudo chown -R $(id -un):$(id -gn) %s
sudo docker rm -f %s 2>/dev/null || true
cat << 'EOF' > %s/docker-compose.yml
%sEOF
cat << 'EOF' > %s/prepare-for-migration.sh
%sEOF
chmod +x %s/prepare-for-migration.sh
%s
for i in $(seq 1 24); do
  if curl -sf http://localhost:%d/api/tags >/dev/null 2>&1; then
    break
  fi
  sleep 3
done`,
		aiCfg.RootDir, aiCfg.RootDir, subDir, extraDirs,
		aiCfg.RootDir, aiCfg.RootDir,
		aiCfg.RootDir,
		staleContainer,
		aiCfg.RootDir, composeContent,
		aiCfg.RootDir, prepHelper,
		aiCfg.RootDir,
		buildDockerComposeUpWithRetryCmd(aiCfg.RootDir, "ollama"),
		aiCfg.OllamaPort)

	fmt.Printf("🐳 Preparing '%s' (%s) & starting Ollama container...\n", aiCfg.RootDir, aiCfg.appDisplayName())
	if _, err := execRemoteCmd(tb, tbEndpoint, nsID, infraID, nodeID, []string{deployCmd}); err != nil {
		return fmt.Errorf("failed to deploy Docker Compose config and Ollama on %s: %w", infraID, err)
	}
	return nil
}

// buildOllamaPullWithRetryCmd builds a 3-attempt shell command for pulling an Ollama model reliably.
func buildOllamaPullWithRetryCmd(rootDir, modelTag string) string {
	return fmt.Sprintf(`for attempt in 1 2 3; do
  if sudo docker exec aimx-ollama ollama pull %s; then
    break
  fi
  if [ "$attempt" -eq 3 ]; then
    exit 1
  fi
  echo "Retrying ollama pull %s (attempt $((attempt+1))/3)..."
  sudo systemctl start containerd docker 2>/dev/null || true
  cd %s && sudo docker compose up -d ollama
  sleep 10
done`, modelTag, modelTag, rootDir)
}

// pullOllamaModels pulls the primary generation model and optional embedding model into aimx-ollama.
func pullOllamaModels(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, aiCfg AIServiceConfig, roleLabel string) error {
	fmt.Printf("📥 Pulling LLM '%s' on %s GPU VM...\n", aiCfg.ModelName, roleLabel)
	pullGenCmd := buildOllamaPullWithRetryCmd(aiCfg.RootDir, aiCfg.ModelName)
	if _, err := execRemoteCmd(tb, tbEndpoint, nsID, infraID, nodeID, []string{pullGenCmd}); err != nil {
		return fmt.Errorf("failed to pull LLM model %s on %s VM: %w", aiCfg.ModelName, strings.ToLower(roleLabel), err)
	}

	if emb := strings.TrimSpace(aiCfg.EmbeddingModel); emb != "" {
		fmt.Printf("📥 Pulling multilingual embedding model '%s' on %s GPU VM...\n", emb, roleLabel)
		pullEmbCmd := buildOllamaPullWithRetryCmd(aiCfg.RootDir, emb)
		if _, err := execRemoteCmd(tb, tbEndpoint, nsID, infraID, nodeID, []string{pullEmbCmd}); err != nil {
			return fmt.Errorf("failed to pull embedding model %s on %s VM: %w", emb, strings.ToLower(roleLabel), err)
		}
	}
	return nil
}

// warmupOllamaModels preloads the generation and embedding models into GPU VRAM with a 30m keep-alive window.
func warmupOllamaModels(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, aiCfg AIServiceConfig) {
	genBody, _ := json.Marshal(map[string]any{
		"model":      aiCfg.ModelName,
		"prompt":     "Hello!",
		"options":    map[string]any{"num_predict": 32},
		"stream":     false,
		"keep_alive": "30m",
	})
	var cmds []string
	cmds = append(cmds, fmt.Sprintf(`curl -sf --max-time 300 http://localhost:%d/api/generate -H "Content-Type: application/json" -d '%s' >/dev/null || true`,
		aiCfg.OllamaPort, string(genBody)))
	if emb := strings.TrimSpace(aiCfg.EmbeddingModel); emb != "" {
		embBody, _ := json.Marshal(map[string]any{
			"model":      emb,
			"input":      "Cloud-Barista conference presentation RAG migration",
			"keep_alive": "30m",
		})
		cmds = append(cmds, fmt.Sprintf(`curl -sf --max-time 300 http://localhost:%d/api/embed -H "Content-Type: application/json" -d '%s' >/dev/null || true`,
			aiCfg.OllamaPort, string(embBody)))
	}
	cmds = append(cmds, "sudo docker exec aimx-ollama ollama ps || true")
	_, _ = execRemoteCmd(tb, tbEndpoint, nsID, infraID, nodeID, []string{strings.Join(cmds, " && ")})
}

// ensureKotaemonLangdetectCache copies lid.176.ftz into the host volume so offline CJK citation matching never stalls.
func ensureKotaemonLangdetectCache(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, aiCfg AIServiceConfig) {
	_, _ = execRemoteCmdQuiet(tb, tbEndpoint, nsID, infraID, nodeID, []string{buildLangdetectCacheScript(aiCfg.RootDir)})
}

// buildKotaemonDefaultsPythonScript builds the SQLite initialization script for Ollama defaults and Korean/Docling RAG settings.
func buildKotaemonDefaultsPythonScript() string {
	return `import json, os, sqlite3, sys, uuid
db = sys.argv[1]
admin_user = sys.argv[2].strip().lower()
changed = False
if os.path.exists(db):
    conn = sqlite3.connect(db, timeout=10)
    cur = conn.cursor()
    existing = {r[0] for r in cur.execute("SELECT name FROM sqlite_master WHERE type='table'").fetchall()}
    candidates = ["llm_table", "llm", "embedding", "resources__llm", "resources__embedding"]
    for t in sorted(existing):
        tl = t.lower()
        if ("llm" in tl or "embedding" in tl) and t not in candidates:
            candidates.append(t)
    for t in candidates:
        if t not in existing or not t.replace("_", "").isalnum():
            continue
        cols = [c[1] for c in cur.execute(f'PRAGMA table_info("{t}")').fetchall()]
        if "name" in cols and "default" in cols:
            row = cur.execute(f'SELECT "default" FROM "{t}" WHERE name = ?', ("ollama",)).fetchone()
            if row is not None and row[0] != 1:
                cur.execute(f'UPDATE "{t}" SET "default" = 0')
                cur.execute(f'UPDATE "{t}" SET "default" = 1 WHERE name = ?', ("ollama",))
                changed = True
    if "user" in existing and "settings" in existing and admin_user:
        admin_row = cur.execute("SELECT id FROM user WHERE lower(username) = ?", (admin_user,)).fetchone()
        if admin_row:
            uid = admin_row[0]
            desired_settings = {
                "index.options.1.reader_mode": "docling",
                "index.options.1.use_reranking": False,
                "index.options.1.use_llm_reranking": False,
                "reasoning.use": "simple",
                "reasoning.lang": "ko",
                "reasoning.options.simple.highlight_citation": "inline",
                "reasoning.options.simple.create_mindmap": True,
                "reasoning.options.complex.highlight_citation": "inline",
                "reasoning.options.complex.create_mindmap": True,
            }
            s_row = cur.execute("SELECT id, setting FROM settings WHERE user = ?", (uid,)).fetchone()
            if s_row:
                sid, raw = s_row
                current = json.loads(raw) if isinstance(raw, str) and raw else {}
                if any(current.get(k) != v for k, v in desired_settings.items()):
                    current.update(desired_settings)
                    cur.execute("UPDATE settings SET setting = ? WHERE id = ?", (json.dumps(current), sid))
                    changed = True
            else:
                sid = uuid.uuid4().hex
                cur.execute(
                    "INSERT INTO settings (id, user, setting) VALUES (?, ?, ?)",
                    (sid, uid, json.dumps(desired_settings)),
                )
                changed = True
    if changed:
        conn.commit()
    conn.close()
print("DEFAULT_UPDATED" if changed else "DEFAULT_UNCHANGED")
`
}

// ensureKotaemonOllamaDefaults activates 'ollama' defaults and Korean/Docling RAG settings in Kotaemon's SQLite database.
func ensureKotaemonOllamaDefaults(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, aiCfg AIServiceConfig, timing TimingConfig) error {
	escapedAdmin := strings.ReplaceAll(strings.TrimSpace(aiCfg.AdminUsername), `'`, `'\''`)
	pyCmd := fmt.Sprintf(`sudo docker exec -i aimx-kotaemon python3 - /app/ktem_app_data/user_data/sql.db '%s' << 'PY_EOF'
%sPY_EOF`, escapedAdmin, buildKotaemonDefaultsPythonScript())
	out, err := execRemoteCmdQuiet(tb, tbEndpoint, nsID, infraID, nodeID, []string{pyCmd})
	if err != nil {
		return err
	}
	if strings.Contains(out, "DEFAULT_UPDATED") {
		fmt.Println("🔧 Activated 'ollama' defaults & Korean RAG settings (Docling / inline citation / mindmap) in Kotaemon sql.db; restarting container...")
		if _, rErr := execRemoteCmdQuiet(tb, tbEndpoint, nsID, infraID, nodeID, []string{"sudo docker restart aimx-kotaemon"}); rErr != nil {
			return rErr
		}
		time.Sleep(10 * time.Second)
		return pollServiceEndpointsReady(tb, tbEndpoint, nsID, infraID, nodeID, aiCfg, timing)
	}
	return nil
}

// discoverSourceInfra performs hybrid inspection of the live source GPU VM and updates the base OnpremInfra model.
func discoverSourceInfra(tb *resty.Client, cfg TestConfig, auth AuthConfig, state *MigrationState) (onpremmodel.OnpremInfra, error) {
	var baseInfra onpremmodel.OnpremInfra
	baseInfraPath := resolveConfigPath(cfg.ConfigDir, cfg.Beetle.BaseInfraFile)
	rawBytes, err := os.ReadFile(baseInfraPath)
	if err != nil {
		return baseInfra, fmt.Errorf("failed to read base infra file %s: %w", baseInfraPath, err)
	}
	if err := json.Unmarshal(rawBytes, &baseInfra); err != nil {
		return baseInfra, fmt.Errorf("failed to unmarshal base infra JSON %s: %w", baseInfraPath, err)
	}
	if len(baseInfra.Nodes) == 0 {
		return baseInfra, fmt.Errorf("base infra file %s must contain at least 1 node", baseInfraPath)
	}

	if cfg.Source.isExternalSSH() {
		if err := enrichNodeFromExternalSource(cfg, state, &baseInfra); err != nil {
			return baseInfra, err
		}
		fmt.Printf("🔎 Inspecting live hardware & GPU specs on External Source VM '%s' (%s) via SSH...\n",
			state.Source.NodeID, state.Source.PublicIP)
	} else {
		if err := enrichNodeFromTumblebugMetadata(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, state, &baseInfra); err != nil {
			return baseInfra, err
		}
		fmt.Printf("🔎 Inspecting live hardware & GPU specs on Source VM '%s' via CB-Tumblebug Remote Command...\n", cfg.Source.InfraID)
	}

	inspectScript := `echo "===GPU_CSV==="
nvidia-smi --query-gpu=index,uuid,name,driver_version,pci.bus_id,memory.total --format=csv,noheader,nounits 2>/dev/null
echo "===GPU_BANNER==="
nvidia-smi 2>/dev/null | head -n 5
echo "===SYS_INFO==="
echo "HOSTNAME=$(hostname)"
echo "ARCH=$(uname -m)"
echo "VCPUS=$(nproc)"
echo "MEM_GIB=$(free -g | awk '/^Mem:/{print $2}')"
echo "MEM_USED_GIB=$(free -g | awk '/^Mem:/{print $3}')"
echo "MEM_AVAIL_GIB=$(free -g | awk '/^Mem:/{print $7}')"
echo "DISK_GB=$(df -BG / | awk 'NR==2 {gsub("G","",$2); print $2}')"
echo "DISK_USED_GB=$(df -BG / | awk 'NR==2 {gsub("G","",$3); print $3}')"
echo "DISK_AVAIL_GB=$(df -BG / | awk 'NR==2 {gsub("G","",$4); print $4}')"
if [ -f /etc/os-release ]; then
  . /etc/os-release
  echo "OS_ID=$ID"
  echo "OS_ID_LIKE=$ID_LIKE"
  echo "OS_NAME=$NAME"
  echo "OS_VER=$VERSION_ID"
  echo "OS_VER_FULL=$VERSION"
  echo "OS_CODENAME=$VERSION_CODENAME"
  echo "OS_PRETTY=$PRETTY_NAME"
fi`

	out, err := runSourceCmd(tb, cfg, auth, state.Source.NodeID, []string{inspectScript}, false)
	if err != nil {
		return baseInfra, fmt.Errorf("failed to inspect live hardware & GPU specs on source VM '%s': %w", cfg.Source.displayLabel(), err)
	}
	enrichNodeFromInspectionOutput(&baseInfra.Nodes[0], out)
	if len(baseInfra.Nodes[0].GPUCards) > 0 {
		c := baseInfra.Nodes[0].GPUCards[0]
		state.Source.GpuVendor = c.Vendor
		state.Source.GpuModel = c.Model
		state.Source.GpuCount = len(baseInfra.Nodes[0].GPUCards)
		state.Source.GpuVramGB = c.MemoryTotalGB
		state.Source.DriverVersion = c.DriverVersion
		state.Source.CudaVersion = c.CudaVersion
	}

	statePath := resolveConfigPath(cfg.ConfigDir, cfg.Beetle.StateFile)
	outDir := filepath.Dir(statePath)
	if outDir == "" || outDir == "." {
		outDir = "testresult"
	}
	if discBytes, mErr := json.MarshalIndent(baseInfra, "", "  "); mErr == nil {
		_ = os.MkdirAll(outDir, 0o755)
		_ = os.WriteFile(filepath.Join(outDir, "source-infra-discovered.json"), discBytes, 0o644)
	}

	state.DiscoveredOnpremInfra = baseInfra
	if err := saveState(statePath, state); err != nil {
		return baseInfra, err
	}
	return baseInfra, nil
}

// enrichNodeFromTumblebugMetadata populates MachineId and SecurityGroup FirewallTable rules from CB-Tumblebug's live Infra and SecurityGroup APIs.
func enrichNodeFromTumblebugMetadata(tb *resty.Client, tbEndpoint, nsID, infraID string, state *MigrationState, baseInfra *onpremmodel.OnpremInfra) error {
	baseURL := strings.TrimRight(tbEndpoint, "/")
	var infraInfo tbmodel.InfraInfo
	resp, err := tb.R().SetResult(&infraInfo).Get(fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s", baseURL, nsID, infraID))
	if err != nil || resp.IsError() {
		return fmt.Errorf("failed to query source infra '%s' from Tumblebug: err=%v, status=%d, body=%s",
			infraID, err, resp.StatusCode(), resp.String())
	}
	if len(infraInfo.Node) == 0 {
		return fmt.Errorf("source infra '%s' has 0 nodes in Tumblebug", infraID)
	}

	tbNode := infraInfo.Node[0]
	state.Source.NodeID = tbNode.Id

	var fwRules []onpremmodel.FirewallRuleProperty
	for _, sgID := range tbNode.SecurityGroupIds {
		var sgInfo tbmodel.SecurityGroupInfo
		sgResp, sgErr := tb.R().SetResult(&sgInfo).Get(fmt.Sprintf("%s/tumblebug/ns/%s/resources/securityGroup/%s", baseURL, nsID, sgID))
		if sgErr != nil || sgResp.IsError() {
			return fmt.Errorf("failed to query securityGroup '%s' for source infra '%s': err=%v, status=%d",
				sgID, infraID, sgErr, sgResp.StatusCode())
		}
		for _, r := range sgInfo.FirewallRules {
			proto := strings.ToLower(strings.TrimSpace(r.Protocol))
			if proto == "all" {
				proto = "*"
			}
			ports := strings.TrimSpace(r.Port)
			if ports == "" || strings.EqualFold(ports, "all") || ports == "1-65535" || ports == "0-65535" {
				ports = "*"
			}
			fwRules = append(fwRules, onpremmodel.FirewallRuleProperty{
				SrcCIDR:   strings.TrimSpace(r.CIDR),
				SrcPorts:  "*",
				DstCIDR:   "0.0.0.0/0",
				DstPorts:  ports,
				Protocol:  proto,
				Direction: strings.ToLower(strings.TrimSpace(r.Direction)),
				Action:    "allow",
			})
		}
	}
	if len(fwRules) > 0 {
		baseInfra.Nodes[0].FirewallTable = fwRules
	}

	baseInfra.Nodes[0].MachineId = state.Source.NodeID
	for i := range baseInfra.Network.IPv4Networks.DefaultGateways {
		baseInfra.Network.IPv4Networks.DefaultGateways[i].MachineId = state.Source.NodeID
	}
	return nil
}

// enrichNodeFromInspectionOutput parses the remote inspection stdout and updates NodeProperty fields in place.
func enrichNodeFromInspectionOutput(node *onpremmodel.NodeProperty, output string) {
	cudaVer := parseCudaVersion(output)
	cards := parseNvidiaSmiCsv(extractSection(output, "===GPU_CSV===", "===GPU_BANNER==="), cudaVer)
	if len(cards) > 0 {
		node.GPUCards = cards
	}

	sysSection := extractSection(output, "===SYS_INFO===", "")
	for _, line := range strings.Split(sysSection, "\n") {
		line = strings.TrimSpace(line)
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key, val := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		switch key {
		case "HOSTNAME":
			if val != "" {
				node.Hostname = val
			}
		case "ARCH":
			if val != "" {
				node.CPU.Architecture = val
			}
		case "VCPUS":
			if v, err := strconv.ParseUint(val, 10, 32); err == nil && v > 0 {
				node.CPU.Cpus = 1
				node.CPU.Threads = uint32(v)
				node.CPU.Cores = uint32(max(1, int(v)/2))
			}
		case "MEM_GIB":
			if m, err := strconv.ParseUint(val, 10, 64); err == nil && m > 0 {
				node.Memory.TotalSize = normalizeStandardMemoryGiB(m)
			}
		case "MEM_USED_GIB":
			if m, err := strconv.ParseUint(val, 10, 64); err == nil {
				node.Memory.Used = m
			}
		case "MEM_AVAIL_GIB":
			if m, err := strconv.ParseUint(val, 10, 64); err == nil {
				node.Memory.Available = m
			}
		case "DISK_GB":
			if d, err := strconv.ParseUint(val, 10, 64); err == nil && d > 0 {
				node.RootDisk.TotalSize = d
			}
		case "DISK_USED_GB":
			if d, err := strconv.ParseUint(val, 10, 64); err == nil {
				node.RootDisk.Used = d
			}
		case "DISK_AVAIL_GB":
			if d, err := strconv.ParseUint(val, 10, 64); err == nil {
				node.RootDisk.Available = d
			}
		case "OS_ID":
			if val != "" {
				node.OS.ID = strings.ToLower(val)
			}
		case "OS_ID_LIKE":
			if val != "" {
				node.OS.IDLike = strings.ToLower(val)
			}
		case "OS_NAME":
			if val != "" {
				node.OS.Name = val
			}
		case "OS_VER":
			if val != "" {
				node.OS.VersionID = val
			}
		case "OS_VER_FULL":
			if val != "" {
				node.OS.Version = val
			}
		case "OS_CODENAME":
			if val != "" {
				node.OS.VersionCodename = strings.ToLower(val)
			}
		case "OS_PRETTY":
			if val != "" {
				node.OS.PrettyName = val
			}
		}
	}
}

// normalizeStandardMemoryGiB rounds Linux `free -g` reported memory up to standard cloud power-of-two GiB sizes.
func normalizeStandardMemoryGiB(rawGiB uint64) uint64 {
	standardSizes := []uint64{4, 8, 16, 32, 64, 128, 256, 512, 1024}
	for _, std := range standardSizes {
		if rawGiB <= std && std-rawGiB <= std/10 {
			return std
		}
	}
	return rawGiB
}

// extractSection extracts text between startMarker and endMarker (or end of string if endMarker is empty).
func extractSection(text, startMarker, endMarker string) string {
	sIdx := strings.Index(text, startMarker)
	if sIdx == -1 {
		return ""
	}
	rest := text[sIdx+len(startMarker):]
	if endMarker != "" {
		if eIdx := strings.Index(rest, endMarker); eIdx != -1 {
			rest = rest[:eIdx]
		}
	}
	return strings.TrimSpace(rest)
}

// parseCudaVersion extracts the CUDA version string from nvidia-smi output.
func parseCudaVersion(text string) string {
	matches := cudaRegex.FindStringSubmatch(text)
	if len(matches) >= 2 {
		return matches[1]
	}
	return ""
}

// parseNvidiaSmiCsv parses `nvidia-smi --query-gpu=index,uuid,name,driver_version,pci.bus_id,memory.total --format=csv,noheader,nounits`.
func parseNvidiaSmiCsv(csvText, cudaVersion string) []onpremmodel.GpuCardProperty {
	var cards []onpremmodel.GpuCardProperty
	for _, line := range strings.Split(strings.TrimSpace(csvText), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Split(line, ",")
		if len(fields) < 6 {
			continue
		}
		idx := strings.TrimSpace(fields[0])
		if _, err := strconv.Atoi(idx); err != nil {
			continue
		}
		uuid := strings.TrimSpace(fields[1])
		modelName := strings.TrimSpace(fields[2])
		driverVer := strings.TrimSpace(fields[3])
		pciBus := strings.TrimSpace(fields[4])
		memMiBStr := strings.TrimSpace(fields[5])

		mib, err := strconv.ParseFloat(memMiBStr, 64)
		if err != nil || mib <= 0 {
			continue
		}
		memGB := snapToStandardVramGB(float32(mib / 1024.0))

		cards = append(cards, onpremmodel.GpuCardProperty{
			DriverIndex:   idx,
			Uuid:          uuid,
			Vendor:        "NVIDIA",
			Model:         modelName,
			Type:          "GPU",
			DriverVersion: driverVer,
			CudaVersion:   cudaVersion,
			PciBusId:      pciBus,
			MemoryTotalGB: memGB,
		})
	}
	return cards
}

// snapToStandardVramGB snaps raw MiB-derived GiB (e.g., 22.49 GiB on 24GB L4) to the standard GPU VRAM tier.
func snapToStandardVramGB(rawGiB float32) float32 {
	tiers := []float32{8, 16, 24, 32, 40, 48, 80, 96, 144, 192, 288}
	for _, tier := range tiers {
		if math.Abs(float64(tier-rawGiB)) <= 2.5 {
			return tier
		}
	}
	return float32(math.Round(float64(rawGiB)))
}
