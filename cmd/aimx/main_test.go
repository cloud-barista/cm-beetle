package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
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
	if cfg.AIService.AppType != "open-webui" {
		t.Errorf("expected appType 'open-webui', got %q", cfg.AIService.AppType)
	}
	if cfg.AIService.relativeDbPath() != "open-webui/webui.db" {
		t.Errorf("expected relativeDbPath 'open-webui/webui.db', got %q", cfg.AIService.relativeDbPath())
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
	if cfg.AIService.TransxStrategy != "auto" {
		t.Errorf("expected default transxStrategy 'auto', got %q", cfg.AIService.TransxStrategy)
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

func TestLoadKotaemonTemplateAndNoFallbackValidation(t *testing.T) {
	cfg, err := loadConfig("testconf/template-kotaemon-config.yaml")
	if err != nil {
		t.Fatalf("expected kotaemon template config to load cleanly, got: %v", err)
	}
	if !cfg.AIService.isKotaemon() {
		t.Fatalf("expected isKotaemon()=true, got appType=%q", cfg.AIService.AppType)
	}
	if cfg.AIService.EmbeddingModel != "bge-m3" {
		t.Errorf("expected embeddingModel 'bge-m3', got %q", cfg.AIService.EmbeddingModel)
	}
	if cfg.AIService.activeAppPort() != 7860 {
		t.Errorf("expected activeAppPort 7860, got %d", cfg.AIService.activeAppPort())
	}
	if cfg.AIService.relativeDbPath() != "ktem_app_data/user_data/sql.db" {
		t.Errorf("expected Kotaemon DB path 'ktem_app_data/user_data/sql.db', got %q", cfg.AIService.relativeDbPath())
	}
	if cfg.AIService.dbFileName() != "sql.db" {
		t.Errorf("expected dbFileName 'sql.db', got %q", cfg.AIService.dbFileName())
	}

	missingAppType := cfg
	missingAppType.AIService.AppType = ""
	if err := validateConfig(missingAppType); err == nil {
		t.Errorf("expected error when aiService.appType is empty (No Fallback rule)")
	}

	unknownAppType := cfg
	unknownAppType.AIService.AppType = "unknown-rag"
	if err := validateConfig(unknownAppType); err == nil {
		t.Errorf("expected error when aiService.appType is unsupported")
	}

	missingEmb := cfg
	missingEmb.AIService.EmbeddingModel = ""
	if err := validateConfig(missingEmb); err == nil {
		t.Errorf("expected error when aiService.embeddingModel is missing for kotaemon")
	}

	missingAdmin := cfg
	missingAdmin.AIService.AdminUsername = ""
	if err := validateConfig(missingAdmin); err == nil {
		t.Errorf("expected error when aiService.adminUsername is missing for kotaemon")
	}
}

func TestBuildDockerComposeYAMLKotaemonAndOpenWebUI(t *testing.T) {
	cfg, err := loadConfig("testconf/template-kotaemon-config.yaml")
	if err != nil {
		t.Fatalf("unexpected error loading kotaemon template: %v", err)
	}
	compose := buildDockerComposeYAML(cfg.AIService)
	expectedSnippets := []string{
		"image: ghcr.io/cinnamon/kotaemon:main-full",
		"container_name: aimx-kotaemon",
		"\"7860:7860\"",
		"KH_OLLAMA_URL=http://ollama:11434/v1/",
		"LOCAL_MODEL=gemma4:e2b",
		"LOCAL_MODEL_EMBEDDINGS=bge-m3",
		"OLLAMA_MAX_LOADED_MODELS=2",
		"OLLAMA_CONTEXT_LENGTH=16384",
		"./ollama/.nv:/root/.nv",
		"KH_APP_VERSION=v0.12.0",
		"FTLANG_CACHE=/app/ktem_app_data/fasttext-langdetect",
		"./ktem_app_data:/app/ktem_app_data",
	}
	for _, s := range expectedSnippets {
		if !strings.Contains(compose, s) {
			t.Errorf("expected Kotaemon compose YAML to contain %q, got:\n%s", s, compose)
		}
	}

	owCfg, err := loadConfig("testconf/template-test-config.yaml")
	if err != nil {
		t.Fatalf("unexpected error loading open-webui template: %v", err)
	}
	owCompose := buildDockerComposeYAML(owCfg.AIService)
	if !strings.Contains(owCompose, "ghcr.io/open-webui/open-webui:main") || !strings.Contains(owCompose, "\"3000:8080\"") {
		t.Errorf("unexpected Open WebUI compose YAML:\n%s", owCompose)
	}
}

func TestIsPortCoveredByRules(t *testing.T) {
	rules := []tbmodel.FirewallRuleInfo{
		{Direction: "inbound", Protocol: "TCP", Port: "1-65535", CIDR: "0.0.0.0/0"},
	}
	if !isPortCoveredByRules(rules, 7860) || !isPortCoveredByRules(rules, 11434) {
		t.Errorf("expected 1-65535 rule to cover ports 7860 and 11434")
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

func TestBuildTransxMigrationModelAndKotaemonProtection(t *testing.T) {
	keys := &ephemeralSshKeyPair{
		SourceKeyID:      "src-key-1",
		SourceUsername:   "ubuntu",
		SourcePrivateKey: "-----BEGIN RSA PRIVATE KEY-----\nSRC\n-----END RSA PRIVATE KEY-----",
		TargetKeyID:      "tgt-key-1",
		TargetUsername:   "cb-user",
		TargetPrivateKey: "-----BEGIN RSA PRIVATE KEY-----\nTGT\n-----END RSA PRIVATE KEY-----",
	}

	aiCfg := AIServiceConfig{
		AppType:              "kotaemon",
		RootDir:              "/opt/aimx/ai-service",
		TransferModelWeights: true,
		TransxStrategy:       "relay",
		Filter: TransxFilterConfig{
			Include: []string{"docker-compose.yml"},
			Exclude: []string{"*.log", "*.db-wal", "ktem_app_data/user_data/**", "ktem_app_data/chunks_cache_dir/**", "ktem_app_data/fasttext-langdetect/**", "ktem_app_data/huggingface/**"},
		},
	}
	dmFull := buildTransxMigrationModel("1.2.3.4", "5.6.7.8", keys, aiCfg)
	if dmFull.Source.Path != "/opt/aimx/ai-service/" {
		t.Errorf("expected trailing slash on source path, got %q", dmFull.Source.Path)
	}
	if dmFull.Destination.Filesystem.SSH.Username != "cb-user" {
		t.Errorf("expected target username 'cb-user', got %q", dmFull.Destination.Filesystem.SSH.Username)
	}
	for _, ex := range dmFull.Source.Filter.Exclude {
		if strings.Contains(ex, "db-wal") || strings.Contains(ex, "ktem_app_data") {
			t.Errorf("protected Kotaemon/WAL pattern must not be excluded, got %v", dmFull.Source.Filter.Exclude)
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

func TestParseEndpointMetricsAndWriteReportBothApps(t *testing.T) {
	sampleOut := `===GPU_CSV===
0, GPU-25d17822-abcd, NVIDIA L4, 570.195.03, 0000:31:00.0, 23034
===GPU_BANNER===
| NVIDIA-SMI 570.195.03             Driver Version: 570.195.03     CUDA Version: 12.8     |
===INFERENCE_JSON===
{"model":"gemma4:e2b","response":"GPUs provide massive parallel tensor cores that accelerate matrix multiplications.","eval_count":85,"eval_duration":1000000000}
===OLLAMA_PS===
100% GPU
===WEBUI_DB_SHA256===
e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
===WEBUI_HTTP===
200
===WEBUI_CONFIG_JSON===
{"status":true,"name":"Open WebUI","version":"0.5.10"}
===VOLUME_SIZE===
6.1G`

	ktCfg, err := loadConfig("testconf/template-kotaemon-config.yaml")
	if err != nil {
		t.Fatalf("failed to load kotaemon template: %v", err)
	}
	epKt := EndpointMetrics{
		InfraID:  "mc-daisy",
		NodeID:   "nvidiat4-1",
		PublicIP: "3.35.10.20",
		Csp:      "aws",
		Region:   "ap-northeast-2",
		SpecID:   "aws+ap-northeast-2+g4dn.xlarge",
		ImageID:  "ami-00ffffe72a332c865",
		WebUIUrl: "http://3.35.10.20:7860",
	}
	parseEndpointMetricsOutput(&epKt, sampleOut, ktCfg.AIService)
	if !epKt.WebUIConfigOK || epKt.WebUIVersion != "main-full" {
		t.Errorf("expected Kotaemon HTTP 200 to set WebUIConfigOK=true and version 'main-full', got ok=%v ver=%q",
			epKt.WebUIConfigOK, epKt.WebUIVersion)
	}

	tmpDir := t.TempDir()
	ktCfg.Beetle.StateFile = filepath.Join(tmpDir, ".aimx-state.json")
	state := &MigrationState{
		Source:               epKt,
		Target:               epKt,
		DataTransferDuration: "52s",
		DataTransferStrategy: "auto",
		StageDurations:       map[string]string{"setup-source": "2m40s", "migrate": "3m10s"},
	}
	reportPath, err := writeMarkdownReport(ktCfg, TestCase{Name: "aws-ap-northeast-2"}, state)
	if err != nil {
		t.Fatalf("failed to write kotaemon markdown report: %v", err)
	}
	contentBytes, err := os.ReadFile(reportPath)
	if err != nil {
		t.Fatalf("failed to read generated report: %v", err)
	}
	content := string(contentBytes)
	for _, expected := range []string{"Kotaemon (`:7860`)", "gemma4:e2b + bge-m3", "sql.db", "Verified Identical", "http://3.35.10.20:7860"} {
		if !strings.Contains(content, expected) {
			t.Errorf("kotaemon report missing %q:\n%s", expected, content)
		}
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

func TestApplyTargetSubnetZone(t *testing.T) {
	baseRec := func() cloudmodel.RecommendedInfra {
		return cloudmodel.RecommendedInfra{
			TargetVNet: cloudmodel.VNetReq{
				SubnetInfoList: []cloudmodel.SubnetReq{{Name: "subnet-01", IPv4_CIDR: "10.0.1.0/24"}},
			},
			TargetInfra: cloudmodel.InfraReq{
				NodeGroups: []cloudmodel.CreateNodeGroupReq{{Name: "ng-01", SpecId: "aws+ap-northeast-2+g5.xlarge"}},
			},
			TargetSpecList: []cloudmodel.SpecInfo{
				{
					Id:               "aws+ap-northeast-2+g5.xlarge",
					AvailableZones:   []string{"ap-northeast-2c", "ap-northeast-2a"},
					UnavailableZones: []string{"ap-northeast-2b", "ap-northeast-2d"},
				},
			},
		}
	}

	backendRec := baseRec()
	backendRec.TargetVNet.SubnetInfoList[0].Zone = "ap-northeast-2a"
	gotBackend, err := applyTargetSubnetZone(&backendRec, TestCase{CloudProperty: cloudmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"}})
	if err != nil || gotBackend != "ap-northeast-2a" || backendRec.TargetVNet.SubnetInfoList[0].Zone != "ap-northeast-2a" {
		t.Errorf("expected backend recommended zone 'ap-northeast-2a' to be preserved, got %q (err=%v)", gotBackend, err)
	}

	awsRec := baseRec()
	gotAws, err := applyTargetSubnetZone(&awsRec, TestCase{CloudProperty: cloudmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"}})
	if err != nil || gotAws != "ap-northeast-2c" || awsRec.TargetVNet.SubnetInfoList[0].Zone != "ap-northeast-2c" {
		t.Errorf("expected zonal CSP aws to pick AvailableZones[0] 'ap-northeast-2c', got %q / %q (err=%v)",
			gotAws, awsRec.TargetVNet.SubnetInfoList[0].Zone, err)
	}

	overrideRec := baseRec()
	gotOverride, err := applyTargetSubnetZone(&overrideRec, TestCase{
		CloudProperty: cloudmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"},
		PreferredZone: "ap-northeast-2a",
	})
	if err != nil || gotOverride != "ap-northeast-2a" || overrideRec.TargetVNet.SubnetInfoList[0].Zone != "ap-northeast-2a" {
		t.Errorf("expected PreferredZone 'ap-northeast-2a' to take precedence, got %q (err=%v)", gotOverride, err)
	}

	invalidRec := baseRec()
	if _, err := applyTargetSubnetZone(&invalidRec, TestCase{
		CloudProperty: cloudmodel.CloudProperty{Csp: "aws", Region: "ap-northeast-2"},
		PreferredZone: "ap-northeast-2b",
	}); err == nil {
		t.Errorf("expected invalid PreferredZone 'ap-northeast-2b' to return an error")
	}

	for _, regionalCSP := range []string{"azure", "gcp", "nhn"} {
		regRec := baseRec()
		gotReg, err := applyTargetSubnetZone(&regRec, TestCase{CloudProperty: cloudmodel.CloudProperty{Csp: regionalCSP, Region: "korea"}})
		if err != nil || gotReg != "" || regRec.TargetVNet.SubnetInfoList[0].Zone != "" {
			t.Errorf("expected regional CSP %s to keep SubnetReq.Zone empty, got %q (err=%v)", regionalCSP, regRec.TargetVNet.SubnetInfoList[0].Zone, err)
		}
	}

	regPinned := baseRec()
	if _, err := applyTargetSubnetZone(&regPinned, TestCase{
		CloudProperty: cloudmodel.CloudProperty{Csp: "azure", Region: "koreacentral"},
		PreferredZone: "koreacentral-1",
	}); err == nil {
		t.Errorf("expected PreferredZone on regional CSP azure to return an error")
	}
}

func TestExternalSshConfigAndConnectionsParsing(t *testing.T) {
	cfg, err := loadConfig("testconf/template-ai-stack-testbed-config.yaml")
	if err != nil {
		t.Fatalf("expected ai-stack-testbed template config to load cleanly, got: %v", err)
	}
	if !cfg.Source.isExternalSSH() {
		t.Fatalf("expected isExternalSSH()=true for ai-stack-testbed template")
	}
	targets := selectedTargets(cfg)
	if len(targets) != 1 || targets[0].Csp != "gcp" || targets[0].Region != "asia-northeast3" {
		t.Fatalf("expected primary target gcp/asia-northeast3, got %+v", targets)
	}

	missingNode := cfg
	missingNode.Source.NodeName = ""
	if err := validateConfig(missingNode); err == nil {
		t.Errorf("expected error when source.nodeName is empty with connectionsFile")
	}

	bothSet := cfg
	bothSet.Source.SshHost = "10.0.0.10"
	if err := validateConfig(bothSet); err == nil {
		t.Errorf("expected error when both connectionsFile and sshHost are set")
	}

	directNoPort := cfg
	directNoPort.Source.ConnectionsFile = ""
	directNoPort.Source.SshHost = "10.0.0.10"
	directNoPort.Source.SshUsername = "ubuntu"
	directNoPort.Source.SshPrivateKeyPath = "/tmp/key"
	directNoPort.Source.SshPort = 0
	if err := validateConfig(directNoPort); err == nil {
		t.Errorf("expected error when source.sshPort <= 0 in direct sshHost mode")
	}

	tbAutoCfg, err := loadConfig("testconf/template-kotaemon-config.yaml")
	if err != nil {
		t.Fatalf("unexpected error loading kotaemon template: %v", err)
	}
	tbAutoCfg.AIService.ModelName = "auto"
	if err := validateConfig(tbAutoCfg); err == nil {
		t.Errorf("expected error when aiService.modelName is 'auto' in CB-Tumblebug source mode")
	}

	tmpDir := t.TempDir()
	connPath := filepath.Join(tmpDir, "source_connections_public.json")
	mockJSON := `[{"name":"ai-gpu-01","ip_address":"203.0.113.10","ssh_port":2222,"user":"ubuntu","private_key":"mock-key-material"}]`
	if err := os.WriteFile(connPath, []byte(mockJSON), 0o600); err != nil {
		t.Fatalf("failed to write mock connections JSON: %v", err)
	}
	cfg.Source.ConnectionsFile = connPath
	access, err := resolveExternalSourceAccess(cfg)
	if err != nil {
		t.Fatalf("expected resolveExternalSourceAccess to succeed, got: %v", err)
	}
	if access.Host != "203.0.113.10" || access.Port != 2222 || access.Username != "ubuntu" {
		t.Errorf("unexpected externalSourceAccess parsed: host=%s port=%d user=%s", access.Host, access.Port, access.Username)
	}
	access.Clear()
	if access.PrivateKey != "" {
		t.Errorf("expected Clear() to wipe in-memory key material")
	}

	sampleCompose := `services:
  kotaemon:
    environment:
      - LOCAL_MODEL=gemma4:e4b
      - LOCAL_MODEL_EMBEDDINGS=bge-m3
`
	genModel, embModel := parseComposeModelEnv(sampleCompose)
	if genModel != "gemma4:e4b" || embModel != "bge-m3" {
		t.Errorf("expected parsed models gemma4:e4b + bge-m3, got %q + %q", genModel, embModel)
	}
}

func TestKotaemonDefaultsPythonScript(t *testing.T) {
	pyBin, err := exec.LookPath("python3")
	if err != nil {
		t.Skip("python3 not installed; skipping SQLite defaults script execution test")
	}

	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "sql.db")
	seedPy := `import json, sqlite3, sys
conn = sqlite3.connect(sys.argv[1])
cur = conn.cursor()
cur.execute('CREATE TABLE llm_table (name TEXT PRIMARY KEY, "default" INTEGER)')
cur.execute('INSERT INTO llm_table VALUES ("openai", 1), ("ollama", 0)')
cur.execute('CREATE TABLE embedding (name TEXT PRIMARY KEY, "default" INTEGER)')
cur.execute('INSERT INTO embedding VALUES ("openai", 1), ("ollama", 0)')
cur.execute('CREATE TABLE user (id TEXT PRIMARY KEY, username TEXT)')
cur.execute('INSERT INTO user VALUES ("uid-admin-1", "admin")')
cur.execute('CREATE TABLE settings (id TEXT PRIMARY KEY, user TEXT, setting TEXT)')
cur.execute('INSERT INTO settings VALUES ("sid-1", "uid-admin-1", ?)', (json.dumps({"custom.keep": "yes", "reasoning.lang": "en"}),))
conn.commit()
conn.close()
`
	if out, sErr := exec.Command(pyBin, "-c", seedPy, dbPath).CombinedOutput(); sErr != nil {
		t.Fatalf("failed to seed mock Kotaemon sql.db: %v (%s)", sErr, string(out))
	}

	script := buildKotaemonDefaultsPythonScript()
	cmd1 := exec.Command(pyBin, "-", dbPath, "admin")
	cmd1.Stdin = strings.NewReader(script)
	out1, err := cmd1.CombinedOutput()
	if err != nil || !strings.Contains(string(out1), "DEFAULT_UPDATED") {
		t.Fatalf("expected first run to output DEFAULT_UPDATED, err=%v, out=%s", err, string(out1))
	}

	cmd2 := exec.Command(pyBin, "-", dbPath, "admin")
	cmd2.Stdin = strings.NewReader(script)
	out2, err := cmd2.CombinedOutput()
	if err != nil || !strings.Contains(string(out2), "DEFAULT_UNCHANGED") {
		t.Fatalf("expected second run to be idempotent (DEFAULT_UNCHANGED), err=%v, out=%s", err, string(out2))
	}

	verifyPy := `import json, sqlite3, sys
conn = sqlite3.connect(sys.argv[1])
cur = conn.cursor()
llm_def = cur.execute('SELECT "default" FROM llm_table WHERE name="ollama"').fetchone()[0]
emb_def = cur.execute('SELECT "default" FROM embedding WHERE name="ollama"').fetchone()[0]
raw_setting = cur.execute('SELECT setting FROM settings WHERE user="uid-admin-1"').fetchone()[0]
conn.close()
data = json.loads(raw_setting)
data["_llm_def"] = llm_def
data["_emb_def"] = emb_def
print(json.dumps(data))
`
	verifyOut, err := exec.Command(pyBin, "-c", verifyPy, dbPath).CombinedOutput()
	if err != nil {
		t.Fatalf("failed to query updated sql.db: %v (%s)", err, string(verifyOut))
	}
	var got map[string]any
	if err := json.Unmarshal(verifyOut, &got); err != nil {
		t.Fatalf("failed to decode settings JSON: %v", err)
	}
	if got["_llm_def"] != float64(1) || got["_emb_def"] != float64(1) {
		t.Errorf("expected ollama default=1 in llm_table and embedding, got %v", got)
	}
	if got["custom.keep"] != "yes" {
		t.Errorf("expected existing custom setting key to be preserved, got %v", got["custom.keep"])
	}
	if got["index.options.1.reader_mode"] != "docling" ||
		got["index.options.1.use_reranking"] != false ||
		got["index.options.1.use_llm_reranking"] != false ||
		got["reasoning.lang"] != "ko" ||
		got["reasoning.options.simple.highlight_citation"] != "inline" ||
		got["reasoning.options.simple.create_mindmap"] != true {
		t.Errorf("unexpected Kotaemon RAG settings persisted: %+v", got)
	}
}
