package main

import (
	"fmt"
	"strings"
	"time"

	tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/transx"
	"github.com/go-resty/resty/v2"
)

// ephemeralSshKeyPair holds Source and Target SSH keys strictly in memory during Stage 3 transfer.
type ephemeralSshKeyPair struct {
	SourceKeyID      string
	SourceUsername   string
	SourcePrivateKey string
	TargetKeyID      string
	TargetUsername   string
	TargetPrivateKey string
}

// Clear wipes the in-memory PEM private keys immediately after Stage 3 completes.
func (k *ephemeralSshKeyPair) Clear() {
	if k == nil {
		return
	}
	k.SourcePrivateKey = ""
	k.TargetPrivateKey = ""
}

// migrateInfraAndAIService executes Stage 3: Target GPU VM provisioning + transx AI model/app migration + Target stack startup.
func migrateInfraAndAIService(beetle, tb *resty.Client, cfg TestConfig, auth AuthConfig, target TestCase, state *MigrationState) error {
	stageStart := time.Now()
	fmt.Println("\n================================================================================")
	fmt.Printf("🚚 [Stage 3] Target GPU Infra Migration & AI Model/Application Migration (%s / %s)\n", target.Csp, target.Region)
	fmt.Println("================================================================================")

	if state.Source.PublicIP == "" || state.Source.SshKeyID == "" {
		if err := refreshSourceEndpointFromTumblebug(tb, cfg, auth, state); err != nil {
			return err
		}
	}

	if _, err := provisionTargetGpuInfra(beetle, tb, cfg, auth, target, state); err != nil {
		return err
	}

	fmt.Println("\n🔑 [Stage 3-2] Retrieving Source & Target SSH keys from CB-Tumblebug (ephemeral in-memory management)...")
	ephemeralKeys, err := acquireEphemeralSshKeys(tb, auth.TumblebugEndpoint, cfg, state)
	if err != nil {
		return err
	}
	defer ephemeralKeys.Clear()

	fmt.Println("⏸️  [Stage 3-3] Quiescing Source AI containers, syncing SQLite DB & normalizing permissions for transx...")
	quiesceCmd := fmt.Sprintf(`set -e
cd %s && sudo docker compose stop && sync
sudo chown -R %s:%s %s
sudo chmod -R u+rwX %s
echo "===QUIESCED_DB_SHA256==="
sha256sum %s/open-webui/webui.db 2>/dev/null | awk '{print $1}' || true`,
		cfg.AIService.RootDir, ephemeralKeys.SourceUsername, ephemeralKeys.SourceUsername, cfg.AIService.RootDir, cfg.AIService.RootDir, cfg.AIService.RootDir)
	quiesceOut, err := execRemoteCmd(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, state.Source.NodeID, []string{quiesceCmd})
	if err != nil {
		return fmt.Errorf("failed to quiesce source AI containers: %w", err)
	}
	if latestSha := strings.TrimSpace(extractSection(quiesceOut, "===QUIESCED_DB_SHA256===", "")); latestSha != "" {
		state.Source.WebUIDbSha256 = strings.Split(latestSha, "\n")[0]
	}

	waitWithCountdown(cfg.Timing.DbFlushWaitSeconds, "Waiting for Open WebUI SQLite WAL & Ollama disk buffers to flush")

	fmt.Println("🐳 [Stage 3-4] Ensuring Docker Compose & NVIDIA Container Runtime on Target GPU VM...")
	if err := ensureDockerAndNvidiaRuntime(tb, auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, state.Target.InfraID, state.Target.NodeID); err != nil {
		return err
	}
	prepTargetDirCmd := fmt.Sprintf("sudo mkdir -p %s && sudo chown -R %s:%s %s",
		cfg.AIService.RootDir, ephemeralKeys.TargetUsername, ephemeralKeys.TargetUsername, cfg.AIService.RootDir)
	if _, err := execRemoteCmd(tb, auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, state.Target.InfraID, state.Target.NodeID, []string{prepTargetDirCmd}); err != nil {
		return fmt.Errorf("failed to prepare target AI service directory: %w", err)
	}

	txModel := buildTransxMigrationModel(state.Source.PublicIP, state.Target.PublicIP, ephemeralKeys, cfg.AIService)
	fmt.Printf("📦 [Stage 3-5] Transferring AI Service stack (%s) via CM-Beetle transx API (strategy=%s, transferModelWeights=%v, encrypt=%v)...\n",
		cfg.AIService.RootDir, txModel.Strategy, cfg.AIService.TransferModelWeights, cfg.AIService.UseDataEncryption)
	if txModel.Source.Filter != nil {
		fmt.Printf("   🔎 Active transx Filter — include=%v, exclude=%v\n",
			txModel.Source.Filter.Include, txModel.Source.Filter.Exclude)
	}

	txStart := time.Now()
	if err := executeTransxMigration(beetle, cfg.Beetle.Endpoint, txModel, cfg.AIService.UseDataEncryption); err != nil {
		return err
	}
	state.DataTransferDuration = time.Since(txStart).Round(time.Second).String()
	state.DataTransferStrategy = txModel.Strategy
	fmt.Printf("✅ transx Data Migration finished in %s!\n", state.DataTransferDuration)

	fmt.Println("▶️  [Stage 3-6] Verifying transferred SQLite DB hash & starting Migrated Ollama (Gemma) + Open WebUI stack on Target GPU VM...")
	startTargetCmd := fmt.Sprintf(`set -e
echo "===TRANSFERRED_DB_SHA256==="
sha256sum %s/open-webui/webui.db 2>/dev/null | awk '{print $1}' || true
cd %s && sudo docker compose up -d`, cfg.AIService.RootDir, cfg.AIService.RootDir)
	startOut, err := execRemoteCmd(tb, auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, state.Target.InfraID, state.Target.NodeID, []string{startTargetCmd})
	if err != nil {
		return fmt.Errorf("failed to start migrated AI service stack on target VM: %w", err)
	}
	if transferredSha := strings.TrimSpace(extractSection(startOut, "===TRANSFERRED_DB_SHA256===", "")); transferredSha != "" {
		state.Target.WebUIDbSha256 = strings.Split(transferredSha, "\n")[0]
	}

	// Resume source containers so both Source and Target Open WebUI remain online for side-by-side browser check.
	resumeSourceCmd := fmt.Sprintf("cd %s && sudo docker compose up -d", cfg.AIService.RootDir)
	_, _ = execRemoteCmd(tb, auth.TumblebugEndpoint, cfg.Source.NamespaceID, cfg.Source.InfraID, state.Source.NodeID, []string{resumeSourceCmd})

	waitWithCountdown(cfg.Timing.TargetServiceReadySeconds, "Waiting for Target GPU Ollama & Open WebUI services to initialize")
	if err := pollServiceEndpointsReady(tb, auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, state.Target.InfraID, state.Target.NodeID, cfg.AIService.OllamaPort, cfg.AIService.OpenWebUIPort, cfg.Timing); err != nil {
		return err
	}

	if !cfg.AIService.TransferModelWeights {
		fmt.Printf("📥 Pulling Gemma model '%s' on Target GPU VM (transferModelWeights=false)...\n", cfg.AIService.ModelName)
		pullCmd := fmt.Sprintf("sudo docker exec aimx-ollama ollama pull %s", cfg.AIService.ModelName)
		if _, err := execRemoteCmd(tb, auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, state.Target.InfraID, state.Target.NodeID, []string{pullCmd}); err != nil {
			return fmt.Errorf("failed to pull Gemma model on target VM: %w", err)
		}
	}

	fmt.Printf("🔥 [Stage 3-7] Pre-warming migrated model '%s' into Target GPU VRAM (cold load may take 3~5 min)...\n", cfg.AIService.ModelName)
	warmupCmd := fmt.Sprintf(`curl -s http://localhost:%d/api/generate -d '{"model":"%s","prompt":"Hello","stream":false,"keep_alive":"30m"}' >/dev/null && sudo docker exec aimx-ollama ollama ps`,
		cfg.AIService.OllamaPort, cfg.AIService.ModelName)
	_, _ = execRemoteCmd(tb, auth.TumblebugEndpoint, cfg.Beetle.NamespaceID, state.Target.InfraID, state.Target.NodeID, []string{warmupCmd})

	state.Source.WebUIUrl = fmt.Sprintf("http://%s:%d", state.Source.PublicIP, cfg.AIService.OpenWebUIPort)
	state.Target.WebUIUrl = fmt.Sprintf("http://%s:%d", state.Target.PublicIP, cfg.AIService.OpenWebUIPort)
	state.StageDurations["migrate"] = time.Since(stageStart).Round(time.Second).String()
	if err := saveState(cfg.Beetle.StateFile, state); err != nil {
		return err
	}

	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("🌐 [Source Open WebUI Link] %s\n", state.Source.WebUIUrl)
	fmt.Printf("🌐 [Target Open WebUI Link] %s  (Migrated!)\n", state.Target.WebUIUrl)
	fmt.Println("--------------------------------------------------------------------------------")
	return nil
}

// refreshSourceEndpointFromTumblebug populates state.Source connection metadata if `migrate` is called standalone.
func refreshSourceEndpointFromTumblebug(tb *resty.Client, cfg TestConfig, auth AuthConfig, state *MigrationState) error {
	baseURL := strings.TrimRight(auth.TumblebugEndpoint, "/")
	getURL := fmt.Sprintf("%s/tumblebug/ns/%s/infra/%s", baseURL, cfg.Source.NamespaceID, cfg.Source.InfraID)
	var infraInfo tbmodel.InfraInfo
	resp, err := tb.R().SetResult(&infraInfo).Get(getURL)
	if err != nil || resp.IsError() {
		return fmt.Errorf("failed to fetch source infra '%s' from Tumblebug: err=%v, status=%d", cfg.Source.InfraID, err, resp.StatusCode())
	}
	metrics, _, err := extractNodeEndpointInfo(infraInfo, cfg.AIService.ModelName)
	if err != nil {
		return err
	}
	metrics.WebUIUrl = fmt.Sprintf("http://%s:%d", metrics.PublicIP, cfg.AIService.OpenWebUIPort)
	state.Source = metrics
	return nil
}

// acquireEphemeralSshKeys retrieves both Source and Target SSH keys from CB-Tumblebug into an ephemeral struct.
func acquireEphemeralSshKeys(tb *resty.Client, tbEndpoint string, cfg TestConfig, state *MigrationState) (*ephemeralSshKeyPair, error) {
	srcKey, srcUser, err := fetchSshPrivateKey(tb, tbEndpoint, cfg.Source.NamespaceID, state.Source.SshKeyID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch source SSH key '%s': %w", state.Source.SshKeyID, err)
	}
	tgtKey, tgtUser, err := fetchSshPrivateKey(tb, tbEndpoint, cfg.Beetle.NamespaceID, state.Target.SshKeyID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch target SSH key '%s': %w", state.Target.SshKeyID, err)
	}

	resolvedSrcUser := firstNonEmpty(state.Source.SshUsername, srcUser)
	resolvedTgtUser := firstNonEmpty(state.Target.SshUsername, tgtUser)
	if resolvedSrcUser == "" || resolvedTgtUser == "" {
		return nil, fmt.Errorf("failed to resolve SSH username from Tumblebug for source (%q) or target (%q)", resolvedSrcUser, resolvedTgtUser)
	}

	return &ephemeralSshKeyPair{
		SourceKeyID:      state.Source.SshKeyID,
		SourceUsername:   resolvedSrcUser,
		SourcePrivateKey: srcKey,
		TargetKeyID:      state.Target.SshKeyID,
		TargetUsername:   resolvedTgtUser,
		TargetPrivateKey: tgtKey,
	}, nil
}

// fetchSshPrivateKey retrieves the PEM private key and username from CB-Tumblebug's SSH key resource API.
func fetchSshPrivateKey(tb *resty.Client, tbEndpoint, nsID, sshKeyID string) (string, string, error) {
	url := fmt.Sprintf("%s/tumblebug/ns/%s/resources/sshKey/%s", strings.TrimRight(tbEndpoint, "/"), nsID, sshKeyID)
	var keyInfo tbmodel.SshKeyInfo
	resp, err := tb.R().SetResult(&keyInfo).Get(url)
	if err != nil {
		return "", "", fmt.Errorf("SSH key HTTP error: %w", err)
	}
	if resp.IsError() {
		return "", "", fmt.Errorf("SSH key API failed (status %d): %s", resp.StatusCode(), resp.String())
	}
	if strings.TrimSpace(keyInfo.PrivateKey) == "" {
		return "", "", fmt.Errorf("SSH key '%s' in namespace '%s' returned empty privateKey", sshKeyID, nsID)
	}
	return keyInfo.PrivateKey, firstNonEmpty(keyInfo.VerifiedUsername, keyInfo.Username), nil
}

// buildTransxFilter constructs the transx.FilterOption from configured filter rules and model weight options.
func buildTransxFilter(aiCfg AIServiceConfig) *transx.FilterOption {
	var includes []string
	seenInc := make(map[string]bool)
	for _, p := range aiCfg.Filter.Include {
		p = strings.TrimSpace(p)
		if p != "" && !seenInc[p] {
			seenInc[p] = true
			includes = append(includes, p)
		}
	}

	var excludes []string
	seenExc := make(map[string]bool)
	addExclude := func(patterns ...string) {
		for _, p := range patterns {
			p = strings.TrimSpace(p)
			if p != "" && !seenExc[p] {
				seenExc[p] = true
				excludes = append(excludes, p)
			}
		}
	}

	addExclude(aiCfg.Filter.Exclude...)
	if !aiCfg.TransferModelWeights {
		addExclude("ollama/models/blobs*", "ollama/models/manifests*")
	}

	if len(includes) == 0 && len(excludes) == 0 {
		return nil
	}
	return &transx.FilterOption{
		Include: includes,
		Exclude: excludes,
	}
}

// buildTransxMigrationModel constructs the CM-Beetle transx DataMigrationModel for SSH-to-SSH transfer.
func buildTransxMigrationModel(srcIP, tgtIP string, keys *ephemeralSshKeyPair, aiCfg AIServiceConfig) transx.DataMigrationModel {
	srcPath := strings.TrimRight(aiCfg.RootDir, "/") + "/"
	tgtPath := strings.TrimRight(aiCfg.RootDir, "/") + "/"
	strategy := strings.ToLower(strings.TrimSpace(aiCfg.TransxStrategy))

	var srcUser, tgtUser, srcKey, tgtKey string
	if keys != nil {
		srcUser = keys.SourceUsername
		tgtUser = keys.TargetUsername
		srcKey = keys.SourcePrivateKey
		tgtKey = keys.TargetPrivateKey
	}

	return transx.DataMigrationModel{
		Source: transx.DataLocation{
			StorageType: transx.StorageTypeFilesystem,
			Path:        srcPath,
			Filesystem: &transx.FilesystemAccess{
				AccessType: transx.AccessTypeSSH,
				SSH: &transx.SSHConfig{
					Host:       srcIP,
					Port:       22,
					Username:   srcUser,
					PrivateKey: srcKey,
					Archive:    true,
					Compress:   true,
				},
			},
			Filter: buildTransxFilter(aiCfg),
		},
		Destination: transx.DataLocation{
			StorageType: transx.StorageTypeFilesystem,
			Path:        tgtPath,
			Filesystem: &transx.FilesystemAccess{
				AccessType: transx.AccessTypeSSH,
				SSH: &transx.SSHConfig{
					Host:       tgtIP,
					Port:       22,
					Username:   tgtUser,
					PrivateKey: tgtKey,
					Archive:    true,
					Compress:   true,
				},
			},
		},
		Strategy: strategy,
	}
}

// executeTransxMigration encrypts SSH keys with Beetle's 1-time KeyStore public key (if enabled) and calls POST /beetle/migration/data.
func executeTransxMigration(beetle *resty.Client, beetleEndpoint string, dm transx.DataMigrationModel, encrypt bool) error {
	baseURL := strings.TrimRight(beetleEndpoint, "/")

	if encrypt {
		keyURL := fmt.Sprintf("%s/beetle/migration/data/encryptionKey", baseURL)
		var keyResp model.ApiResponse[transx.PublicKeyBundle]
		r, err := beetle.R().SetResult(&keyResp).Get(keyURL)
		if err != nil || r.IsError() || !keyResp.Success {
			return fmt.Errorf("failed to obtain Beetle 1-time encryption key: err=%v, status=%d", err, r.StatusCode())
		}
		pubKey, err := transx.ParsePublicKeyBundle(keyResp.Data)
		if err != nil {
			return fmt.Errorf("failed to parse transx PublicKeyBundle: %w", err)
		}
		dm, err = transx.EncryptModel(dm, pubKey, keyResp.Data.KeyID)
		if err != nil {
			return fmt.Errorf("failed to encrypt transx DataMigrationModel: %w", err)
		}
		fmt.Printf("   🔒 Encrypted Source & Target SSH private keys with Beetle 1-time keyId=%s\n", keyResp.Data.KeyID)
	}

	migURL := fmt.Sprintf("%s/beetle/migration/data", baseURL)
	var migResp model.ApiResponse[any]
	resp, err := beetle.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&dm).
		SetResult(&migResp).
		Post(migURL)
	if err != nil {
		return fmt.Errorf("transx data migration HTTP error: %w", err)
	}
	if resp.IsError() || !migResp.Success {
		return fmt.Errorf("transx data migration failed (status %d): %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// firstNonEmpty returns the first non-blank string from the provided candidates.
func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}


