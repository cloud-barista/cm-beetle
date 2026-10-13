package config

import (
	"os"
	"testing"
)

func TestLogMaskingEnabled_Config(t *testing.T) {
	// 1. Test environment variable parsing
	t.Setenv("BEETLE_LOG_MASKING_ENABLED", "false")
	Init()

	if Beetle.LogMaskingEnabled == nil {
		t.Fatal("expected Beetle.LogMaskingEnabled to be non-nil when set via env")
	}
	if *Beetle.LogMaskingEnabled != false {
		t.Errorf("expected Beetle.LogMaskingEnabled to be false, got %v", *Beetle.LogMaskingEnabled)
	}

	// 2. Test when set to true
	t.Setenv("BEETLE_LOG_MASKING_ENABLED", "true")
	Init()

	if Beetle.LogMaskingEnabled == nil {
		t.Fatal("expected Beetle.LogMaskingEnabled to be non-nil when set via env")
	}
	if *Beetle.LogMaskingEnabled != true {
		t.Errorf("expected Beetle.LogMaskingEnabled to be true, got %v", *Beetle.LogMaskingEnabled)
	}

	// 3. Test default when unset
	os.Unsetenv("BEETLE_LOG_MASKING_ENABLED")
	Init()
	// When unset, Beetle.LogMaskingEnabled should either be nil (handled as true by NewLogger) or true from config.yaml
	if Beetle.LogMaskingEnabled != nil && *Beetle.LogMaskingEnabled != true {
		t.Errorf("expected Beetle.LogMaskingEnabled to be true or nil when unset, got %v", *Beetle.LogMaskingEnabled)
	}
}
