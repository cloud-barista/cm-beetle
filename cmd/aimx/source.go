package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
	"github.com/go-resty/resty/v2"
)

// cudaRegex extracts the CUDA version reported in the standard nvidia-smi banner.
var cudaRegex = regexp.MustCompile(`CUDA Version:\s*([0-9.]+)`)

// setupSourceEnvironment provisions or reuses the source GPU VM and seeds Ollama (Gemma) + Open WebUI.
func setupSourceEnvironment(tb *resty.Client, cfg TestConfig, auth AuthConfig, state *MigrationState) error {
	start := time.Now()
	fmt.Println("\n================================================================================")
	fmt.Println("🚀 [Stage 1] Source GPU VM Provisioning & AI Service (Gemma + Open WebUI) Setup")
	fmt.Println("================================================================================")

	infraInfo, created, err := ensureSourceGpuInfra(tb, cfg, auth)
	if err != nil {
		return err
	}
	if created {
		waitWithCountdown(cfg.Timing.VmBootStabilizeSeconds, "Waiting for newly provisioned Source GPU VM boot & NVIDIA driver init")
	}

	metrics, sgIDs, err := extractNodeEndpointInfo(infraInfo, cfg.AIService.ModelName)
	if err != nil {
		return err
	}
	fmt.Printf("📌 Source GPU Node: id=%s, ip=%s, spec=%s, image=%s\n",
		metrics.NodeID, metrics.PublicIP, metrics.SpecID, metrics.ImageID)

	if cfg.Source.OpenServicePorts && len(sgIDs) > 0 {
		openSecurityGroupServicePorts(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, sgIDs, cfg.AIService)
	}

	if _, err := pollGpuVmReady(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.Timing); err != nil {
		return err
	}

	if err := deployDockerComposeStack(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.AIService); err != nil {
		return err
	}

	waitWithCountdown(cfg.Timing.ServiceInitWaitSeconds, "Waiting for Source Ollama & Open WebUI containers to initialize")
	if err := pollServiceEndpointsReady(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, cfg.AIService.OllamaPort, cfg.AIService.OpenWebUIPort, cfg.Timing); err != nil {
		return err
	}

	fmt.Printf("📥 Pulling Gemma model '%s' on Source GPU VM...\n", cfg.AIService.ModelName)
	pullCmd := fmt.Sprintf("sudo docker exec aimx-ollama ollama pull %s", cfg.AIService.ModelName)
	if _, err := execRemoteCmd(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, metrics.NodeID, []string{pullCmd}); err != nil {
		return fmt.Errorf("failed to pull Gemma model %s on source VM: %w", cfg.AIService.ModelName, err)
	}

	fmt.Println("⚡ Running baseline Gemma GPU inference & collecting Source metrics...")
	if err := collectEndpointMetrics(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, &metrics, cfg.AIService); err != nil {
		return fmt.Errorf("failed to collect baseline source metrics: %w", err)
	}

	metrics.WebUIUrl = fmt.Sprintf("http://%s:%d", metrics.PublicIP, cfg.AIService.OpenWebUIPort)
	state.Source = metrics
	state.StageDurations["setup-source"] = time.Since(start).Round(time.Second).String()
	if err := saveState(cfg.Beetle.StateFile, state); err != nil {
		return err
	}

	fmt.Printf("✅ [Stage 1 Complete] Source GPU='%s' (%.0fGB), Gemma='%s', Speed=%.2f tok/s\n",
		metrics.GpuModel, metrics.GpuVramGB, metrics.ModelName, metrics.EvalTokensPerSec)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("🌐 [Source Open WebUI Ready] Browser Link : %s\n", metrics.WebUIUrl)
	fmt.Printf("🌐 [Source Ollama API Ready] API Endpoint : http://%s:%d/api/tags\n", metrics.PublicIP, cfg.AIService.OllamaPort)
	fmt.Println("--------------------------------------------------------------------------------")
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
func extractNodeEndpointInfo(infraInfo tbmodel.InfraInfo, modelName string) (EndpointMetrics, []string, error) {
	if len(infraInfo.Node) == 0 {
		return EndpointMetrics{}, nil, fmt.Errorf("infra '%s' has no nodes", infraInfo.Id)
	}
	n := infraInfo.Node[0]
	return EndpointMetrics{
		InfraID:     infraInfo.Id,
		NodeID:      n.Id,
		PublicIP:    n.PublicIP,
		SshKeyID:    n.SshKeyId,
		SshUsername: n.NodeUserName,
		Csp:         n.ConnectionConfig.ProviderName,
		Region:      n.Region.Region,
		SpecID:      n.SpecId,
		ImageID:     n.ImageId,
		ModelName:   modelName,
	}, n.SecurityGroupIds, nil
}

// openSecurityGroupServicePorts ensures Ollama and Open WebUI TCP ports are open on the Security Group via Tumblebug.
func openSecurityGroupServicePorts(tb *resty.Client, tbEndpoint, nsID string, sgIDs []string, aiCfg AIServiceConfig) {
	baseURL := strings.TrimRight(tbEndpoint, "/")
	requiredPorts := []int{aiCfg.OpenWebUIPort, aiCfg.OllamaPort}

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

// buildDockerComposeYAML generates the Docker Compose file for Ollama (GPU) and Open WebUI.
func buildDockerComposeYAML(aiCfg AIServiceConfig) string {
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
    volumes:
      - ./ollama:/root/.ollama
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
`, aiCfg.OllamaPort, aiCfg.OpenWebUIPort)
}

// ensureDockerAndNvidiaRuntime verifies Docker Compose and NVIDIA Container Toolkit on a target/source GPU VM.
func ensureDockerAndNvidiaRuntime(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string) error {
	checkAndFixScript := `set -e
if ! command -v docker >/dev/null 2>&1; then
  curl -fsSL https://get.docker.com | sudo sh
fi
if ! sudo docker compose version >/dev/null 2>&1; then
  sudo apt-get -o DPkg::Lock::Timeout=120 update -qq && sudo apt-get -o DPkg::Lock::Timeout=120 install -y -qq docker-compose-v2 || sudo apt-get -o DPkg::Lock::Timeout=120 install -y -qq docker-compose-plugin
fi
if command -v nvidia-ctk >/dev/null 2>&1; then
  if ! grep -q "nvidia" /etc/docker/daemon.json 2>/dev/null; then
    sudo nvidia-ctk runtime configure --runtime=docker >/dev/null 2>&1 || true
    sudo systemctl restart docker
  fi
fi
echo "RUNTIME_READY"`

	_, err := execRemoteCmd(tb, tbEndpoint, nsID, infraID, nodeID, []string{checkAndFixScript})
	if err != nil {
		return fmt.Errorf("failed to verify Docker/NVIDIA container runtime on %s: %w", infraID, err)
	}
	return nil
}

// deployDockerComposeStack creates the AI service directory structure, writes docker-compose.yml, and starts containers.
func deployDockerComposeStack(tb *resty.Client, tbEndpoint, nsID, infraID, nodeID string, aiCfg AIServiceConfig) error {
	if err := ensureDockerAndNvidiaRuntime(tb, tbEndpoint, nsID, infraID, nodeID); err != nil {
		return err
	}

	composeContent := buildDockerComposeYAML(aiCfg)
	deployCmd := fmt.Sprintf(`set -e
sudo mkdir -p %s/ollama %s/open-webui
sudo chown -R $(id -un):$(id -gn) %s
cat << 'EOF' > %s/docker-compose.yml
%sEOF
cd %s && sudo docker compose up -d`,
		aiCfg.RootDir, aiCfg.RootDir, aiCfg.RootDir, aiCfg.RootDir, composeContent, aiCfg.RootDir)

	fmt.Printf("🐳 Deploying Ollama + Open WebUI Docker Compose stack in '%s'...\n", aiCfg.RootDir)
	_, err := execRemoteCmd(tb, tbEndpoint, nsID, infraID, nodeID, []string{deployCmd})
	if err != nil {
		return fmt.Errorf("failed to deploy Docker Compose stack on %s: %w", infraID, err)
	}
	return nil
}

// discoverSourceInfra performs hybrid inspection of the live source GPU VM and updates the base OnpremInfra model.
func discoverSourceInfra(tb *resty.Client, cfg TestConfig, auth AuthConfig, state *MigrationState) (onpremmodel.OnpremInfra, error) {
	var baseInfra onpremmodel.OnpremInfra
	rawBytes, err := os.ReadFile(cfg.Beetle.BaseInfraFile)
	if err != nil {
		return baseInfra, fmt.Errorf("failed to read base infra file %s: %w", cfg.Beetle.BaseInfraFile, err)
	}
	if err := json.Unmarshal(rawBytes, &baseInfra); err != nil {
		return baseInfra, fmt.Errorf("failed to unmarshal base infra JSON %s: %w", cfg.Beetle.BaseInfraFile, err)
	}
	if len(baseInfra.Nodes) == 0 {
		return baseInfra, fmt.Errorf("base infra file %s must contain at least 1 node", cfg.Beetle.BaseInfraFile)
	}

	if err := enrichNodeFromTumblebugMetadata(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, state, &baseInfra); err != nil {
		return baseInfra, err
	}

	fmt.Printf("🔎 Inspecting live hardware & GPU specs on Source VM '%s' via CB-Tumblebug Remote Command...\n", cfg.Source.InfraID)
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

	out, err := execRemoteCmdQuiet(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, state.Source.NodeID, []string{inspectScript})
	if err != nil {
		return baseInfra, fmt.Errorf("failed to inspect live hardware & GPU specs on source VM '%s': %w", cfg.Source.InfraID, err)
	}
	enrichNodeFromInspectionOutput(&baseInfra.Nodes[0], out)

	if discBytes, mErr := json.MarshalIndent(baseInfra, "", "  "); mErr == nil {
		_ = os.MkdirAll("testresult", 0o755)
		_ = os.WriteFile("testresult/source-infra-discovered.json", discBytes, 0o644)
	}

	state.DiscoveredOnpremInfra = baseInfra
	if err := saveState(cfg.Beetle.StateFile, state); err != nil {
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
