package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
)

func TestLoadTemplateConfigAndDefaults(t *testing.T) {
	cfg, err := loadConfig("testconf/template-test-config.yaml")
	if err != nil {
		t.Fatalf("expected template config to load cleanly, got: %v", err)
	}
	if cfg.Test.ExecutionMode != "seq" {
		t.Errorf("expected default executionMode 'seq', got %q", cfg.Test.ExecutionMode)
	}
	if cfg.Source.NodeGroupName != "nvidiat4" {
		t.Errorf("expected nodeGroupName 'nvidiat4', got %q", cfg.Source.NodeGroupName)
	}
	if cfg.Source.SpecID != "aws+ap-northeast-2+g4dn.xlarge" {
		t.Errorf("expected specId 'aws+ap-northeast-2+g4dn.xlarge', got %q", cfg.Source.SpecID)
	}
	if cfg.Source.ImageID != "ami-00ffffe72a332c865" {
		t.Errorf("expected imageId 'ami-00ffffe72a332c865', got %q", cfg.Source.ImageID)
	}
	if cfg.AIService.ModelName != "gemma4:e2b" {
		t.Errorf("expected modelName 'gemma4:e2b', got %q", cfg.AIService.ModelName)
	}
	if cfg.AIService.TransxStrategy != "relay" {
		t.Errorf("expected default transxStrategy 'relay', got %q", cfg.AIService.TransxStrategy)
	}
	if !cfg.AIService.UseDataEncryption {
		t.Errorf("expected useDataEncryption=true in template config")
	}
	targets := selectedTargets(cfg)
	if len(targets) == 0 {
		t.Fatalf("expected at least 1 selected target in template config")
	}
	if targets[0].Csp != "aws" || targets[0].Region != "ap-northeast-2" {
		t.Errorf("expected aws/ap-northeast-2 target, got %s/%s", targets[0].Csp, targets[0].Region)
	}

	failedInfra := tbmodel.InfraInfo{
		Id:     "mc-daisy",
		Status: "Failed",
		Node:   []tbmodel.NodeInfo{{Id: "nvidial4-1", Status: "Failed", PublicIP: ""}},
	}
	if isInfraHealthyAndRunning(failedInfra) {
		t.Errorf("expected failed infra without PublicIP to be reported unhealthy")
	}
	runningInfra := tbmodel.InfraInfo{
		Id:     "mc-daisy",
		Status: "Running",
		Node:   []tbmodel.NodeInfo{{Id: "nvidial4-1", Status: "Running", PublicIP: "3.35.10.20"}},
	}
	if !isInfraHealthyAndRunning(runningInfra) {
		t.Errorf("expected running infra with PublicIP to be reported healthy")
	}
}

func TestIsPortCoveredByRules(t *testing.T) {
	rules := []tbmodel.FirewallRuleInfo{
		{Direction: "inbound", Protocol: "TCP", Port: "1-65535", CIDR: "0.0.0.0/0"},
	}
	if !isPortCoveredByRules(rules, 3000) || !isPortCoveredByRules(rules, 11434) {
		t.Errorf("expected 1-65535 rule to cover ports 3000 and 11434")
	}
}

func TestEnrichNodeFromInspectionOutput(t *testing.T) {
	sampleOutput := `===GPU_CSV===
0, GPU-25d17822-abcd-1234-ef01-987654321000, NVIDIA L4, 570.195.03, 00000000:31:00.0, 23034
===GPU_BANNER===
Thu Oct 14 02:15:00 2026
+-----------------------------------------------------------------------------------------+
| NVIDIA-SMI 570.195.03             Driver Version: 570.195.03     CUDA Version: 12.8     |
|-----------------------------------------+------------------------+----------------------+
===SYS_INFO===
HOSTNAME=ip-10-0-1-25
VCPUS=4
MEM_GIB=15
DISK_GB=97
OS_ID=ubuntu
OS_VER=22.04
OS_PRETTY=Ubuntu 22.04.5 LTS`

	var node onpremmodel.NodeProperty
	enrichNodeFromInspectionOutput(&node, sampleOutput)

	if node.Hostname != "ip-10-0-1-25" {
		t.Errorf("expected hostname 'ip-10-0-1-25', got %q", node.Hostname)
	}
	if node.CPU.Threads != 4 || node.CPU.Cores != 2 {
		t.Errorf("expected 4 threads / 2 cores, got %d / %d", node.CPU.Threads, node.CPU.Cores)
	}
	if node.Memory.TotalSize != 16 {
		t.Errorf("expected normalized memory 16 GiB, got %d", node.Memory.TotalSize)
	}
	if len(node.GPUCards) != 1 {
		t.Fatalf("expected 1 GPU card parsed, got %d", len(node.GPUCards))
	}
	card := node.GPUCards[0]
	if card.Model != "NVIDIA L4" || card.MemoryTotalGB != 24 || card.CudaVersion != "12.8" {
		t.Errorf("unexpected GPU card parsed: %+v", card)
	}
}

func TestBuildTransxMigrationModelAndEphemeralKeys(t *testing.T) {
	keys := &ephemeralSshKeyPair{
		SourceKeyID:      "src-key-1",
		SourceUsername:   "cb-user",
		SourcePrivateKey: "-----BEGIN RSA PRIVATE KEY-----\nSRC\n-----END RSA PRIVATE KEY-----",
		TargetKeyID:      "tgt-key-1",
		TargetUsername:   "ubuntu",
		TargetPrivateKey: "-----BEGIN RSA PRIVATE KEY-----\nTGT\n-----END RSA PRIVATE KEY-----",
	}

	aiCfg := AIServiceConfig{
		RootDir:              "/opt/aimx/ai-service",
		TransferModelWeights: true,
		TransxStrategy:       "relay",
		Filter: TransxFilterConfig{
			Include: []string{"docker-compose.yml"},
			Exclude: []string{"open-webui/uploads/**"},
		},
	}
	dmFull := buildTransxMigrationModel("1.2.3.4", "5.6.7.8", keys, aiCfg)
	if dmFull.Source.Path != "/opt/aimx/ai-service/" {
		t.Errorf("expected trailing slash on source path, got %q", dmFull.Source.Path)
	}
	if dmFull.Strategy != "relay" {
		t.Errorf("expected strategy 'relay', got %q", dmFull.Strategy)
	}
	if dmFull.Destination.Filesystem.SSH.Username != "ubuntu" {
		t.Errorf("expected target username 'ubuntu' from ephemeral key pair, got %q", dmFull.Destination.Filesystem.SSH.Username)
	}
	if dmFull.Source.Filter == nil || len(dmFull.Source.Filter.Include) != 1 {
		t.Fatalf("expected 1 include pattern, got %+v", dmFull.Source.Filter)
	}
	for _, ex := range dmFull.Source.Filter.Exclude {
		if strings.HasPrefix(ex, "ollama/models/blobs") {
			t.Errorf("did not expect model blob exclusion when TransferModelWeights=true, got %v", dmFull.Source.Filter.Exclude)
		}
	}

	aiCfg.TransferModelWeights = false
	dmLight := buildTransxMigrationModel("1.2.3.4", "5.6.7.8", keys, aiCfg)
	foundBlobExclude := false
	for _, ex := range dmLight.Source.Filter.Exclude {
		if ex == "ollama/models/blobs*" {
			foundBlobExclude = true
		}
	}
	if !foundBlobExclude {
		t.Errorf("expected 'ollama/models/blobs*' in exclude filter when TransferModelWeights=false, got %v", dmLight.Source.Filter.Exclude)
	}

	keys.Clear()
	if keys.SourcePrivateKey != "" || keys.TargetPrivateKey != "" {
		t.Errorf("expected ephemeral SSH keys to be wiped after Clear()")
	}
}

func TestParseEndpointMetricsAndWriteReport(t *testing.T) {
	sampleOut := `===GPU_CSV===
0, GPU-25d17822-abcd, NVIDIA L4, 570.195.03, 0000:31:00.0, 23034
===GPU_BANNER===
| NVIDIA-SMI 570.195.03             Driver Version: 570.195.03     CUDA Version: 12.8     |
===INFERENCE_JSON===
{"model":"gemma2:2b","response":"GPUs provide massive parallel tensor cores that accelerate matrix multiplications.","eval_count":85,"eval_duration":1000000000}
===OLLAMA_PS===
100% GPU
===WEBUI_DB_SHA256===
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
===WEBUI_HTTP===
200
===WEBUI_CONFIG_JSON===
{"status":true,"name":"Open WebUI","version":"0.5.10"}
===VOLUME_SIZE===
1.8G`

	ep := EndpointMetrics{
		InfraID:  "mc-daisy",
		NodeID:   "nvidial4-1",
		PublicIP: "3.35.10.20",
		Csp:      "aws",
		Region:   "ap-northeast-2",
		SpecID:   "aws+ap-northeast-2+g6.xlarge",
		ImageID:  "ami-00ffffe72a332c865",
		WebUIUrl: "http://3.35.10.20:3000",
	}
	parseEndpointMetricsOutput(&ep, sampleOut)

	if ep.GpuModel != "NVIDIA L4" || ep.GpuVramGB != 24 {
		t.Errorf("expected NVIDIA L4 24GB, got %s %.0fGB", ep.GpuModel, ep.GpuVramGB)
	}
	if ep.EvalTokensPerSec != 85.0 {
		t.Errorf("expected 85.0 tok/s, got %.2f", ep.EvalTokensPerSec)
	}
	if ep.OllamaProcessor != "100% GPU" {
		t.Errorf("expected '100%% GPU', got %q", ep.OllamaProcessor)
	}
	if ep.WebUIHttpStatus != 200 || !ep.WebUIConfigOK || ep.WebUIVersion != "0.5.10" {
		t.Errorf("expected HTTP 200 and WebUIConfigOK v0.5.10, got status=%d ok=%v ver=%q",
			ep.WebUIHttpStatus, ep.WebUIConfigOK, ep.WebUIVersion)
	}

	tmpDir := t.TempDir()
	cfg, _ := loadConfig("testconf/template-test-config.yaml")
	cfg.Beetle.StateFile = filepath.Join(tmpDir, ".aimx-state.json")

	state := &MigrationState{
		Source:               ep,
		Target:               ep,
		DataTransferDuration: "42s",
		DataTransferStrategy: "relay",
		StageDurations:       map[string]string{"setup-source": "2m10s", "migrate": "3m15s"},
	}
	reportPath, err := writeMarkdownReport(cfg, TestCase{Name: "aws-ap-northeast-2"}, state)
	if err != nil {
		t.Fatalf("failed to write markdown report: %v", err)
	}
	content, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read generated report: %v", err)
	}
	if !strings.Contains(string(content), "85.00 tokens/sec") ||
		!strings.Contains(string(content), "Verified Identical") ||
		!strings.Contains(string(content), "http://3.35.10.20:3000") {
		t.Errorf("generated report missing expected metrics:\n%s", string(content))
	}
}

func TestConsumeCommandSSEStream(t *testing.T) {
	sseSuccess := `: connected to stream for xRequestId=req-123

data: {"type":"CommandStatus","nodeId":"nvidiat4-1","commandIndex":1,"status":{"status":"Handling"}}

data: {"type":"CommandLog","nodeId":"nvidiat4-1","commandIndex":1,"log":{"stream":"stdout","line":"pulling manifest","lineNumber":1}}

data: {"type":"CommandLog","nodeId":"nvidiat4-1","commandIndex":1,"log":{"stream":"stdout","line":"success","lineNumber":2}}

data: {"type":"CommandDone","summary":{"totalNodes":1,"completedNodes":1,"failedNodes":0,"elapsedSeconds":12}}

`
	out, err := consumeCommandSSEStream(strings.NewReader(sseSuccess), "mc-daisy", false)
	if err != nil {
		t.Fatalf("expected SSE stream success, got error: %v", err)
	}
	if out != "pulling manifest\nsuccess" {
		t.Errorf("unexpected stdout from SSE stream: %q", out)
	}

	sseFail := `: connected to stream for xRequestId=req-456

data: {"type":"CommandLog","nodeId":"nvidiat4-1","commandIndex":1,"log":{"stream":"stderr","line":"permission denied","lineNumber":1}}

data: {"type":"CommandStatus","nodeId":"nvidiat4-1","commandIndex":1,"status":{"status":"Failed","errorMessage":"exit status 1"}}

data: {"type":"CommandDone","summary":{"totalNodes":1,"completedNodes":0,"failedNodes":1,"elapsedSeconds":2}}

`
	_, failErr := consumeCommandSSEStream(strings.NewReader(sseFail), "mc-daisy", false)
	if failErr == nil || !strings.Contains(failErr.Error(), "exit status 1") || !strings.Contains(failErr.Error(), "permission denied") {
		t.Errorf("expected SSE failure error containing exit status and stderr, got: %v", failErr)
	}

	rawAnsi := "verifying sha256 digest ⠙ \x1b[K\x1b[?25h\x1b[A\x1b[A\x1b[1Gpulling 6dd7ac24fe13: 100% ▕████████▏ 3.5 GB/3.5 GB\x1b[K"
	cleaned := sanitizeStreamLogLine(rawAnsi)
	if cleaned != "pulling 6dd7ac24fe13: 100% ▕████████▏ 3.5 GB/3.5 GB" {
		t.Errorf("unexpected sanitized ANSI line: %q", cleaned)
	}
	if !isTransientProgressLine(cleaned) {
		t.Errorf("expected progress line to be detected as transient: %q", cleaned)
	}
	if isTransientProgressLine("success") {
		t.Errorf("expected 'success' to be a permanent milestone line")
	}
}


