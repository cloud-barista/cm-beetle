package logger

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestMaskingWriter_Write(t *testing.T) {
	buf := &bytes.Buffer{}
	mw := NewMaskingWriter(buf)

	input := []byte(`{"message":"user login","password":"secretPassword"}`)
	n, err := mw.Write(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != len(input) {
		t.Errorf("Write() returned n = %d, want len(p) = %d", n, len(input))
	}

	output := buf.String()
	if strings.Contains(output, "secretPassword") {
		t.Errorf("output contains unmasked password: %s", output)
	}
	if !strings.Contains(output, `"password":"********"`) {
		t.Errorf("output does not contain masked password: %s", output)
	}
}

func TestMaskingWriter_ZerologIntegration(t *testing.T) {
	buf := &bytes.Buffer{}
	mw := NewMaskingWriter(buf)
	logger := zerolog.New(mw).With().Logger()

	logger.Info().
		Str("user", "admin").
		Str("MasterUserPassword", "Password123!").
		Msg("Internal Call Start")

	output := buf.String()
	if strings.Contains(output, "Password123!") {
		t.Errorf("zerolog output leaked sensitive data: %s", output)
	}
	if !strings.Contains(output, `"MasterUserPassword":"********"`) {
		t.Errorf("zerolog output missing masked field: %s", output)
	}
}

func TestConfigureWriter_MaskingToggle(t *testing.T) {
	// Masking enabled via configureWriter
	loggerEnabled := configureWriter("stdout", zerolog.DebugLevel, true)
	if loggerEnabled == nil {
		t.Fatal("configureWriter returned nil")
	}

	// Masking disabled via configureWriter
	loggerDisabled := configureWriter("stdout", zerolog.DebugLevel, false)
	if loggerDisabled == nil {
		t.Fatal("configureWriter returned nil")
	}

	// NewLogger with nil (default: enabled)
	lDefault := NewLogger(Config{LogWriter: "stdout"})
	if lDefault == nil {
		t.Fatal("NewLogger returned nil for default")
	}

	// NewLogger with LogMaskingEnabled: true
	lTrue := NewLogger(Config{LogWriter: "stdout", LogMaskingEnabled: BoolPtr(true)})
	if lTrue == nil {
		t.Fatal("NewLogger returned nil for true")
	}

	// NewLogger with LogMaskingEnabled: false
	lFalse := NewLogger(Config{LogWriter: "stdout", LogMaskingEnabled: BoolPtr(false)})
	if lFalse == nil {
		t.Fatal("NewLogger returned nil for false")
	}
}

func BenchmarkZerolog_WithMasking(b *testing.B) {
	mw := NewMaskingWriter(io.Discard)
	logger := zerolog.New(mw)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().Str("nsId", "ns01").Str("nodeId", "node01").Msg("processing request")
	}
}

func BenchmarkZerolog_WithoutMasking(b *testing.B) {
	logger := zerolog.New(io.Discard)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info().Str("nsId", "ns01").Str("nodeId", "node01").Msg("processing request")
	}
}
