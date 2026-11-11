package transx

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestMigrateDataWithResultSymmetricHooksOrder(t *testing.T) {
	if _, err := exec.LookPath("rsync"); err != nil {
		t.Skip("rsync binary not available in test environment")
	}

	tmpDir := t.TempDir()
	srcDir := filepath.Join(tmpDir, "src")
	dstDir := filepath.Join(tmpDir, "dst")
	orderFile := filepath.Join(tmpDir, "order.log")

	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatalf("failed to create srcDir: %v", err)
	}
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		t.Fatalf("failed to create dstDir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(srcDir, "webui.db"), []byte("ai-stack-state-v1"), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	dmm := DataMigrationModel{
		Source: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        srcDir,
			Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
			PreCmd:      "printf '1-src-pre\\n' >> " + orderFile + " && printf 'QUIESCED_DB_SHA256=abc123\\n'",
			PostCmd:     "printf '4-src-post\\n' >> " + orderFile + " && printf 'SOURCE_UNQUIESCED=true\\n'",
		},
		Destination: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        dstDir,
			Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
			PreCmd:      "printf '2-dst-pre\\n' >> " + orderFile + " && printf 'TARGET_READY=true\\n'",
			PostCmd:     "test -f " + filepath.Join(dstDir, "webui.db") + " && printf '5-dst-post\\n' >> " + orderFile + " && printf 'TARGET_DB_SHA256=abc123\\n'",
		},
		Strategy: StrategyAuto,
	}

	result, err := MigrateDataWithResult(dmm)
	if err != nil {
		t.Fatalf("MigrateDataWithResult failed: %v", err)
	}

	orderBytes, err := os.ReadFile(orderFile)
	if err != nil {
		t.Fatalf("failed to read order log: %v", err)
	}
	gotOrder := strings.TrimSpace(string(orderBytes))
	wantOrder := "1-src-pre\n2-dst-pre\n4-src-post\n5-dst-post"
	if gotOrder != wantOrder {
		t.Fatalf("unexpected hook execution order:\ngot:\n%s\nwant:\n%s", gotOrder, wantOrder)
	}

	if result.HookOutputs == nil {
		t.Fatalf("expected non-nil HookOutputs")
	}
	if result.HookOutputs.SourcePreCmd != "QUIESCED_DB_SHA256=abc123" {
		t.Errorf("unexpected SourcePreCmd output: %q", result.HookOutputs.SourcePreCmd)
	}
	if result.HookOutputs.DestinationPreCmd != "TARGET_READY=true" {
		t.Errorf("unexpected DestinationPreCmd output: %q", result.HookOutputs.DestinationPreCmd)
	}
	if result.HookOutputs.SourcePostCmd != "SOURCE_UNQUIESCED=true" {
		t.Errorf("unexpected SourcePostCmd output: %q", result.HookOutputs.SourcePostCmd)
	}
	if result.HookOutputs.DestinationPostCmd != "TARGET_DB_SHA256=abc123" {
		t.Errorf("unexpected DestinationPostCmd output: %q", result.HookOutputs.DestinationPostCmd)
	}
	if result.SourceRecovered {
		t.Errorf("expected SourceRecovered=false on normal success")
	}
	if result.RequestedStrategy != StrategyAuto || result.EffectiveStrategy == "" {
		t.Errorf("unexpected strategies: requested=%q effective=%q", result.RequestedStrategy, result.EffectiveStrategy)
	}
}

func TestMigrateDataWithResultSourceRecoveryOnFailure(t *testing.T) {
	tmpDir := t.TempDir()
	srcDir := filepath.Join(tmpDir, "missing-src-for-transfer")
	dstDir := filepath.Join(tmpDir, "dst")
	recoveryMarker := filepath.Join(tmpDir, "recovered.marker")

	if err := os.MkdirAll(dstDir, 0755); err != nil {
		t.Fatalf("failed to create dstDir: %v", err)
	}

	// Case 1: Destination.PreCmd fails after Source.PreCmd quiesced the source.
	dmmDestPreFail := DataMigrationModel{
		Source: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        srcDir,
			Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
			PreCmd:      "printf 'src-stopped'",
			PostCmd:     "printf 'recovered-after-dest-pre' > " + recoveryMarker + " && printf 'src-restarted'",
		},
		Destination: DataLocation{
			StorageType: StorageTypeFilesystem,
			Path:        dstDir,
			Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
			PreCmd:      "printf 'dest-pre-boom' && exit 1",
		},
		Strategy: StrategyAuto,
	}

	res1, err := MigrateDataWithResult(dmmDestPreFail)
	if err == nil {
		t.Fatalf("expected error when Destination.PreCmd fails")
	}
	var migErr *MigrationError
	if !errors.As(err, &migErr) || migErr.Stage != StageBackup {
		t.Fatalf("expected MigrationError at stage %q, got %v", StageBackup, err)
	}
	if !res1.SourceRecovered {
		t.Errorf("expected SourceRecovered=true after Destination.PreCmd failure")
	}
	if res1.HookOutputs == nil || res1.HookOutputs.SourcePostCmd != "src-restarted" {
		t.Errorf("expected SourcePostCmd output to be captured during recovery, got %+v", res1.HookOutputs)
	}
	markerBytes, readErr := os.ReadFile(recoveryMarker)
	if readErr != nil || string(markerBytes) != "recovered-after-dest-pre" {
		t.Fatalf("expected recovery marker to be written, got err=%v content=%q", readErr, string(markerBytes))
	}

	// Case 2: Transfer fails after Source.PreCmd and Destination.PreCmd succeeded.
	if _, lookErr := exec.LookPath("rsync"); lookErr == nil {
		_ = os.Remove(recoveryMarker)
		dmmTransferFail := DataMigrationModel{
			Source: DataLocation{
				StorageType: StorageTypeFilesystem,
				Path:        filepath.Join(tmpDir, "nonexistent-source-dir"),
				Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
				PreCmd:      "printf 'src-quiesced'",
				PostCmd:     "printf 'recovered-after-transfer' > " + recoveryMarker + " && printf 'src-unquiesced'",
			},
			Destination: DataLocation{
				StorageType: StorageTypeFilesystem,
				Path:        dstDir,
				Filesystem:  &FilesystemAccess{AccessType: AccessTypeLocal},
				PostCmd:     "printf 'should-not-run'",
			},
			Strategy: StrategyAuto,
		}

		res2, err2 := MigrateDataWithResult(dmmTransferFail)
		if err2 == nil {
			t.Fatalf("expected transfer error for nonexistent source path")
		}
		if !errors.As(err2, &migErr) || migErr.Stage != StageTransfer {
			t.Fatalf("expected MigrationError at stage %q, got %v", StageTransfer, err2)
		}
		if !res2.SourceRecovered {
			t.Errorf("expected SourceRecovered=true after transfer failure")
		}
		if res2.HookOutputs == nil || res2.HookOutputs.DestinationPostCmd != "" {
			t.Errorf("expected DestinationPostCmd not to run after transfer failure, got %+v", res2.HookOutputs)
		}
	}
}

func TestTruncateHookOutputBoundsAndUTF8(t *testing.T) {
	raw := strings.Repeat("가나다라마바사아자차카타파하", 400)
	got := truncateHookOutput([]byte(raw))
	if len(got) > MaxHookOutputBytes {
		t.Fatalf("expected truncated output <= %d bytes, got %d", MaxHookOutputBytes, len(got))
	}
	if !utf8.ValidString(got) {
		t.Fatalf("expected truncated output to remain valid UTF-8")
	}
	if !strings.HasSuffix(got, "\n...(truncated)") {
		t.Fatalf("expected truncation suffix, got tail %q", got[len(got)-20:])
	}
}
