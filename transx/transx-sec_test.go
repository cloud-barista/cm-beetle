package transx

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/cloud-barista/cm-beetle/transx/fieldsec"
)

func TestKeyStore(t *testing.T) {
	store := NewKeyStore()

	// Generate and store a key
	kp, err := store.GenerateKeyPair(5 * time.Minute)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	if kp.ID == "" {
		t.Error("KeyPair ID should not be empty")
	}
	if kp.PublicKey == nil {
		t.Error("PublicKey should not be nil")
	}
	if kp.PrivateKey == nil {
		t.Error("PrivateKey should not be nil")
	}
	if kp.IsExpired() {
		t.Error("Newly created key should not be expired")
	}

	// Retrieve the key
	retrieved, ok := store.Get(kp.ID)
	if !ok {
		t.Error("Key should be found in store")
	}
	if retrieved.ID != kp.ID {
		t.Errorf("Retrieved key ID mismatch: got %s, want %s", retrieved.ID, kp.ID)
	}

	// Delete the key
	store.Delete(kp.ID)
	_, ok = store.Get(kp.ID)
	if ok {
		t.Error("Key should not be found after delete")
	}
}

func TestExportAndParsePublicBundle(t *testing.T) {
	store := NewKeyStore()
	kp, err := store.GenerateKeyPair(5 * time.Minute)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	bundle, err := kp.ExportPublicBundle()
	if err != nil {
		t.Fatalf("ExportPublicBundle failed: %v", err)
	}

	if bundle.KeyID != kp.ID {
		t.Errorf("KeyID mismatch: got %s, want %s", bundle.KeyID, kp.ID)
	}
	if bundle.Algorithm != fieldsec.CryptoAlgorithm {
		t.Errorf("Algorithm mismatch: got %s, want %s", bundle.Algorithm, fieldsec.CryptoAlgorithm)
	}

	// Parse the bundle back
	pubKey, err := ParsePublicKeyBundle(bundle)
	if err != nil {
		t.Fatalf("ParsePublicKeyBundle failed: %v", err)
	}
	if pubKey == nil {
		t.Error("Parsed public key should not be nil")
	}
}

func TestEncryptDecryptModel(t *testing.T) {
	// Generate key
	store := NewKeyStore()
	kp, err := store.GenerateKeyPair(5 * time.Minute)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Create a model with sensitive data
	model := DataMigrationModel{
		Source: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        "/data/source",
			Filesystem: &FilesystemAccess{
				AccessType: AccessTypeSSH,
				SSH: &SSHConfig{
					Host:       "192.168.1.100",
					Port:       22,
					Username:   "ubuntu",
					PrivateKey: "-----BEGIN RSA PRIVATE KEY-----\nMIIE...(mock key content)...\n-----END RSA PRIVATE KEY-----",
				},
			},
		},
		Destination: DataLocation{
			StorageType: StorageTypeObjectStorage,
			Path:        "my-bucket/backup",
			ObjectStorage: &ObjectStorageAccess{
				AccessType: AccessTypeMinio,
				Minio: &S3MinioConfig{
					Endpoint:        "s3.amazonaws.com",
					AccessKeyId:     "AKIAIOSFODNN7EXAMPLE",
					SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
					Region:          "us-east-1",
				},
			},
		},
	}

	// Encrypt
	encModel, err := EncryptModel(model, kp.PublicKey, kp.ID)
	if err != nil {
		t.Fatalf("EncryptModel failed: %v", err)
	}

	// Verify encryption metadata
	if !encModel.IsEncrypted() {
		t.Error("Model should be marked as encrypted")
	}
	if encModel.EncryptionKeyID != kp.ID {
		t.Errorf("KeyID mismatch: got %s, want %s", encModel.EncryptionKeyID, kp.ID)
	}

	// Verify fields are encrypted (should be different from original)
	if encModel.Source.Filesystem.SSH.PrivateKey == model.Source.Filesystem.SSH.PrivateKey {
		t.Error("SSH PrivateKey should be encrypted")
	}
	if encModel.Destination.ObjectStorage.Minio.AccessKeyId == model.Destination.ObjectStorage.Minio.AccessKeyId {
		t.Error("Minio AccessKeyId should be encrypted")
	}
	if encModel.Destination.ObjectStorage.Minio.SecretAccessKey == model.Destination.ObjectStorage.Minio.SecretAccessKey {
		t.Error("Minio SecretAccessKey should be encrypted")
	}

	// Decrypt
	decModel, err := DecryptModelWith(encModel, kp)
	if err != nil {
		t.Fatalf("DecryptModelWith failed: %v", err)
	}

	// Verify decrypted values match original
	if decModel.Source.Filesystem.SSH.PrivateKey != model.Source.Filesystem.SSH.PrivateKey {
		t.Error("Decrypted SSH PrivateKey should match original")
	}
	if decModel.Destination.ObjectStorage.Minio.AccessKeyId != model.Destination.ObjectStorage.Minio.AccessKeyId {
		t.Error("Decrypted Minio AccessKeyId should match original")
	}
	if decModel.Destination.ObjectStorage.Minio.SecretAccessKey != model.Destination.ObjectStorage.Minio.SecretAccessKey {
		t.Error("Decrypted Minio SecretAccessKey should match original")
	}
}

func TestDecryptModelWith(t *testing.T) {
	store := NewKeyStore()

	// Generate key in store
	kp, err := store.GenerateKeyPair(5 * time.Minute)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Create and encrypt model
	model := DataMigrationModel{
		Source: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        "/data",
			Filesystem: &FilesystemAccess{
				AccessType: AccessTypeSSH,
				SSH: &SSHConfig{
					Host:       "192.168.1.100",
					Username:   "ubuntu",
					PrivateKey: "secret-key-content",
				},
			},
		},
		Destination: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        "/backup",
			Filesystem: &FilesystemAccess{
				AccessType: AccessTypeLocal,
			},
		},
	}

	encModel, err := EncryptModel(model, kp.PublicKey, kp.ID)
	if err != nil {
		t.Fatalf("EncryptModel failed: %v", err)
	}

	// Decrypt with keypair
	decModel, err := DecryptModelWith(encModel, kp)
	if err != nil {
		t.Fatalf("DecryptModelWith failed: %v", err)
	}

	// Verify decryption
	if decModel.Source.Filesystem.SSH.PrivateKey != model.Source.Filesystem.SSH.PrivateKey {
		t.Error("Decrypted value mismatch")
	}
}

func TestExpiredKey(t *testing.T) {
	// Create an already expired key
	store := NewKeyStore()
	kp, err := store.GenerateKeyPair(1 * time.Millisecond)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Wait for expiration
	time.Sleep(10 * time.Millisecond)

	if !kp.IsExpired() {
		t.Error("Key should be expired")
	}

	// Try to decrypt with expired key
	model := DataMigrationModel{
		Source: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        "/data",
			Filesystem: &FilesystemAccess{
				AccessType: AccessTypeSSH,
				SSH: &SSHConfig{
					Host:       "test",
					Username:   "test",
					PrivateKey: "secret",
				},
			},
		},
		Destination: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        "/backup",
			Filesystem: &FilesystemAccess{
				AccessType: AccessTypeLocal,
			},
		},
	}

	encModel, _ := EncryptModel(model, kp.PublicKey, kp.ID)

	_, err = DecryptModelWith(encModel, kp)
	if err != ErrKeyExpired {
		t.Errorf("Expected ErrKeyExpired, got: %v", err)
	}
}

func TestKeyStoreCleanupExpired(t *testing.T) {
	store := NewKeyStore()

	// Create expired keys
	for range 3 {
		kp, _ := store.GenerateKeyPair(1 * time.Millisecond)
		_ = kp // keep reference
	}

	// Create valid key
	validKp, _ := store.GenerateKeyPair(5 * time.Minute)

	// Wait for expiration
	time.Sleep(10 * time.Millisecond)

	// Cleanup expired
	deleted := store.DeleteExpired()
	if deleted != 3 {
		t.Errorf("Expected 3 deleted, got %d", deleted)
	}

	// Valid key should still exist
	_, ok := store.Get(validKp.ID)
	if !ok {
		t.Error("Valid key should still exist")
	}
}

func TestNonEncryptedModelPassthrough(t *testing.T) {
	store := NewKeyStore()
	kp, _ := store.GenerateKeyPair(5 * time.Minute)

	// Model without encryption
	model := DataMigrationModel{
		Source: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        "/data",
		},
		Destination: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        "/backup",
		},
	}

	// DecryptModelWith should pass through non-encrypted model
	result, err := DecryptModelWith(model, kp)
	if err != nil {
		t.Fatalf("DecryptModelWith should not fail for non-encrypted model: %v", err)
	}
	if result.Source.Path != model.Source.Path {
		t.Error("Non-encrypted model should pass through unchanged")
	}
}

func TestLongTextEncryption(t *testing.T) {
	store := NewKeyStore()
	kp, err := store.GenerateKeyPair(5 * time.Minute)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Create model with long SSH key (simulating real RSA key ~3KB)
	longKey := generateLongText(3000)

	model := DataMigrationModel{
		Source: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        "/data",
			Filesystem: &FilesystemAccess{
				AccessType: AccessTypeSSH,
				SSH: &SSHConfig{
					Host:       "192.168.1.100",
					Username:   "ubuntu",
					PrivateKey: longKey,
				},
			},
		},
		Destination: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        "/backup",
			Filesystem: &FilesystemAccess{
				AccessType: AccessTypeLocal,
			},
		},
	}

	// Encrypt and decrypt
	encModel, err := EncryptModel(model, kp.PublicKey, kp.ID)
	if err != nil {
		t.Fatalf("EncryptModel failed: %v", err)
	}

	decModel, err := DecryptModelWith(encModel, kp)
	if err != nil {
		t.Fatalf("DecryptModelWith failed: %v", err)
	}

	if decModel.Source.Filesystem.SSH.PrivateKey != longKey {
		t.Error("Long text should be preserved after encrypt/decrypt")
	}
}

// generateLongText generates a text of approximately the given length.
func generateLongText(length int) string {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range length {
		result[i] = base[i%len(base)]
	}
	return string(result)
}

// mockStepExecutor simulates step execution for pipeline fallback and staging cleanup tests.
type mockStepExecutor struct {
	err error
}

// Execute returns the configured error for the mock executor.
func (m *mockStepExecutor) Execute(source, destination DataLocation) error {
	return m.err
}

func TestSSHAgentFwdAndStrategyValidation(t *testing.T) {
	srcSSH := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "/var/lib/mysql/data",
		Filesystem: &FilesystemAccess{
			AccessType: AccessTypeSSH,
			SSH: &SSHConfig{
				Host:     "10.0.1.10",
				Port:     22,
				Username: "ubuntu",
			},
		},
	}
	dstSSH := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "/var/lib/mysql/target",
		Filesystem: &FilesystemAccess{
			AccessType: AccessTypeSSH,
			SSH: &SSHConfig{
				Host:     "203.0.113.50",
				Port:     2222,
				Username: "admin",
			},
		},
	}

	autoPipeline, err := Plan(DataMigrationModel{
		Source:      srcSSH,
		Destination: dstSSH,
		Strategy:    StrategyAuto,
	})
	if err != nil {
		t.Fatalf("Plan(StrategyAuto) failed: %v", err)
	}
	if len(autoPipeline.Steps) != 1 || autoPipeline.Steps[0].Mode != TransferModeSSHAgentFwdAuto {
		t.Errorf("Expected 1 step with mode %s, got %+v", TransferModeSSHAgentFwdAuto, autoPipeline.Steps)
	}
	if autoPipeline.Fallback == nil || len(autoPipeline.Fallback.Steps) != 2 {
		t.Errorf("Expected StrategyAuto to attach 2-step Relay fallback pipeline, got %+v", autoPipeline.Fallback)
	}

	directPipeline, err := Plan(DataMigrationModel{
		Source:      srcSSH,
		Destination: dstSSH,
		Strategy:    StrategyDirect,
	})
	if err != nil {
		t.Fatalf("Plan(StrategyDirect) failed: %v", err)
	}
	if len(directPipeline.Steps) != 1 || directPipeline.Steps[0].Mode != TransferModeSSHAgentFwdAuto {
		t.Errorf("Expected StrategyDirect to use %s, got %+v", TransferModeSSHAgentFwdAuto, directPipeline.Steps)
	}
	if directPipeline.Fallback != nil {
		t.Errorf("Expected StrategyDirect to have nil Fallback, got %+v", directPipeline.Fallback)
	}

	pushPipeline, err := Plan(DataMigrationModel{
		Source:      srcSSH,
		Destination: dstSSH,
		Strategy:    StrategySSHAgentFwdPush,
	})
	if err != nil {
		t.Fatalf("Plan(StrategySSHAgentFwdPush) failed: %v", err)
	}
	if len(pushPipeline.Steps) != 1 || pushPipeline.Steps[0].Mode != TransferModeSSHAgentFwdPush {
		t.Errorf("Expected StrategySSHAgentFwdPush to use %s, got %+v", TransferModeSSHAgentFwdPush, pushPipeline.Steps)
	}

	pullPipeline, err := Plan(DataMigrationModel{
		Source:      srcSSH,
		Destination: dstSSH,
		Strategy:    StrategySSHAgentFwdPull,
	})
	if err != nil {
		t.Fatalf("Plan(StrategySSHAgentFwdPull) failed: %v", err)
	}
	if len(pullPipeline.Steps) != 1 || pullPipeline.Steps[0].Mode != TransferModeSSHAgentFwdPull {
		t.Errorf("Expected StrategySSHAgentFwdPull to use %s, got %+v", TransferModeSSHAgentFwdPull, pullPipeline.Steps)
	}

	relayPipeline, err := Plan(DataMigrationModel{
		Source:      srcSSH,
		Destination: dstSSH,
		Strategy:    StrategyRelay,
	})
	if err != nil {
		t.Fatalf("Plan(StrategyRelay) failed: %v", err)
	}
	if len(relayPipeline.Steps) != 2 || relayPipeline.Steps[0].Mode != TransferModePull || relayPipeline.Steps[1].Mode != TransferModePush {
		t.Errorf("Expected 2-step Relay (pull -> push), got %+v", relayPipeline.Steps)
	}

	localSrc := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "/tmp/local-src",
		Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
	}
	localDst := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "/tmp/local-dst",
		Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
	}
	localPipeline, err := Plan(DataMigrationModel{
		Source:      localSrc,
		Destination: localDst,
		Strategy:    StrategyAuto,
	})
	if err != nil {
		t.Fatalf("Expected local -> local with StrategyAuto to succeed in transx, got: %v", err)
	}
	if len(localPipeline.Steps) != 1 || localPipeline.Steps[0].Mode != TransferModePush {
		t.Errorf("Expected local -> local to plan 1-step push, got %+v", localPipeline.Steps)
	}

	if err := Validate(DataMigrationModel{
		Source:      localSrc,
		Destination: localDst,
		Strategy:    StrategyRelay,
	}); err == nil {
		t.Errorf("Expected local -> local with StrategyRelay to fail validation")
	}

	if err := Validate(DataMigrationModel{
		Source:      localSrc,
		Destination: dstSSH,
		Strategy:    StrategySSHAgentFwdPush,
	}); err == nil {
		t.Errorf("Expected local -> ssh with StrategySSHAgentFwdPush to fail validation")
	}

	execPush, err := NewRsyncExecutorWithMode(srcSSH, dstSSH, TransferModeSSHAgentFwdPush)
	if err != nil {
		t.Fatalf("NewRsyncExecutorWithMode(Push) failed: %v", err)
	}
	pushCmd := execPush.buildRemoteRsyncCommand(srcSSH, dstSSH, dstSSH.Filesystem.SSH)
	if !strings.Contains(pushCmd, "/var/lib/mysql/data/ admin@203.0.113.50:/var/lib/mysql/target") {
		t.Errorf("Unexpected push rsync command: %s", pushCmd)
	}
	if !strings.Contains(pushCmd, "--partial --append-verify --timeout=180") {
		t.Errorf("Expected resilience flags in push rsync command: %s", pushCmd)
	}

	execPull, err := NewRsyncExecutorWithMode(srcSSH, dstSSH, TransferModeSSHAgentFwdPull)
	if err != nil {
		t.Fatalf("NewRsyncExecutorWithMode(Pull) failed: %v", err)
	}
	pullCmd := execPull.buildRemoteRsyncPullCommand(srcSSH, dstSSH, srcSSH.Filesystem.SSH)
	if !strings.Contains(pullCmd, "ubuntu@10.0.1.10:/var/lib/mysql/data/ /var/lib/mysql/target") {
		t.Errorf("Unexpected pull rsync command: %s", pullCmd)
	}
}

func TestSSHConfigUnmarshalJSONDefaultsAndOverrides(t *testing.T) {
	var defaultCfg SSHConfig
	if err := json.Unmarshal([]byte(`{"host":"10.0.0.1","username":"ubuntu","verbose":true}`), &defaultCfg); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if !defaultCfg.Archive || !defaultCfg.Compress || !defaultCfg.Partial || defaultCfg.MaxRetries != 3 || !defaultCfg.Verbose {
		t.Errorf("Expected defaults (Archive=true, Compress=true, Partial=true, MaxRetries=3, Verbose=true), got %+v", defaultCfg)
	}

	var disabledCfg SSHConfig
	if err := json.Unmarshal([]byte(`{"host":"10.0.0.2","username":"ubuntu","archive":false,"compress":false,"partial":false,"maxRetries":1}`), &disabledCfg); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	if disabledCfg.Archive || disabledCfg.Compress || disabledCfg.Partial || disabledCfg.MaxRetries != 1 {
		t.Errorf("Expected explicit false overrides to be preserved, got %+v", disabledCfg)
	}

	srcLoc := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "/data/src",
		Filesystem:  &FilesystemAccess{AccessType: AccessTypeSSH, SSH: &disabledCfg},
	}
	dstLoc := DataLocation{
		StorageType: StorageTypeFilesystem,
		Path:        "/data/dst",
		Filesystem:  &FilesystemAccess{AccessType: AccessTypeSSH, SSH: &defaultCfg},
	}
	exec := NewRsyncExecutorFromLocations(srcLoc, dstLoc, TransferModeSSHAgentFwdAuto)
	if exec.Archive || exec.Compress || exec.Partial || exec.MaxRetries != 1 {
		t.Errorf("Expected RsyncExecutor to honor explicit false/1 options, got Archive=%v Compress=%v Partial=%v MaxRetries=%d",
			exec.Archive, exec.Compress, exec.Partial, exec.MaxRetries)
	}
	flags := strings.Join(exec.buildBaseRsyncFlags(), " ")
	if !strings.Contains(flags, "-r") || strings.Contains(flags, "-a") || strings.Contains(flags, "-z") || strings.Contains(flags, "--partial") {
		t.Errorf("Unexpected base rsync flags for disabled options: %s", flags)
	}
}

func TestPreFlightFallbackAndStagingCleanup(t *testing.T) {
	rootStaging := createLocalStagingLocation("/")
	if rootStaging.Path != filepath.Join(DefaultStagingPath, "root") {
		t.Errorf("Expected root path staging to be isolated under /tmp/transx-staging/root, got %s", rootStaging.Path)
	}

	stagingLoc := createLocalStagingLocation("/var/data/transx-unit-test-cleanup")
	fallbackPipeline := &Pipeline{
		Name:              PipelineFilesystemTransfer,
		Strategy:          StrategyRelay,
		EffectiveStrategy: StrategyRelay,
		Steps: []Step{
			{
				Name:        StepPullToStaging,
				Mode:        TransferModePull,
				Source:      DataLocation{StorageType: StorageTypeFilesystem, Path: "/remote/src"},
				Destination: stagingLoc,
				Executor:    &mockStepExecutor{err: nil},
			},
			{
				Name:        StepPushFromStaging,
				Mode:        TransferModePush,
				Source:      stagingLoc,
				Destination: DataLocation{StorageType: StorageTypeFilesystem, Path: "/remote/dst"},
				Executor:    &mockStepExecutor{err: nil},
			},
		},
	}

	primaryPipeline := &Pipeline{
		Name:              PipelineFilesystemTransfer,
		Strategy:          StrategyAuto,
		EffectiveStrategy: StrategyAuto,
		Steps: []Step{
			{
				Name:        StepRsyncTransfer,
				Mode:        TransferModeSSHAgentFwdAuto,
				Source:      DataLocation{StorageType: StorageTypeFilesystem, Path: "/remote/src"},
				Destination: DataLocation{StorageType: StorageTypeFilesystem, Path: "/remote/dst"},
				Executor: &mockStepExecutor{
					err: &PreFlightError{Mode: TransferModeSSHAgentFwdAuto, Err: fmt.Errorf("port 22 filtered")},
				},
			},
		},
		Fallback: fallbackPipeline,
	}

	if err := primaryPipeline.Execute(); err != nil {
		t.Fatalf("Expected PreFlightError to trigger Fallback and succeed, got: %v", err)
	}
	if primaryPipeline.EffectiveStrategy != StrategyRelay {
		t.Errorf("Expected EffectiveStrategy to be updated to %q, got %q", StrategyRelay, primaryPipeline.EffectiveStrategy)
	}
	if _, err := os.Stat(stagingLoc.Path); !os.IsNotExist(err) {
		t.Errorf("Expected staging directory %s to be cleaned up on success", stagingLoc.Path)
	}

	failStagingLoc := createLocalStagingLocation("/var/data/transx-unit-test-fail-preserve")
	defer os.RemoveAll(failStagingLoc.Path)

	failRelayPipeline := &Pipeline{
		Name:              PipelineFilesystemTransfer,
		Strategy:          StrategyRelay,
		EffectiveStrategy: StrategyRelay,
		Steps: []Step{
			{
				Name:        StepPullToStaging,
				Mode:        TransferModePull,
				Source:      DataLocation{StorageType: StorageTypeFilesystem, Path: "/remote/src"},
				Destination: failStagingLoc,
				Executor:    &mockStepExecutor{err: nil},
			},
			{
				Name:        StepPushFromStaging,
				Mode:        TransferModePush,
				Source:      failStagingLoc,
				Destination: DataLocation{StorageType: StorageTypeFilesystem, Path: "/remote/dst"},
				Executor:    &mockStepExecutor{err: fmt.Errorf("destination disk full")},
			},
		},
	}
	if err := failRelayPipeline.Execute(); err == nil {
		t.Fatalf("Expected failing relay pipeline to return error")
	}
	if _, err := os.Stat(failStagingLoc.Path); err != nil {
		t.Errorf("Expected staging directory %s to be preserved on failure for debugging, got err: %v", failStagingLoc.Path, err)
	}
}

