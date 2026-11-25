package logger

import (
	"io"

	"github.com/cloud-barista/cm-beetle/pkg/redact"
	"github.com/rs/zerolog"
)

// MaskingWriter intercepts the log output stream and masks sensitive data.
type MaskingWriter struct {
	w io.Writer
}

// NewMaskingWriter wraps an io.Writer with sensitive data masking.
func NewMaskingWriter(w io.Writer) *MaskingWriter {
	return &MaskingWriter{w: w}
}

// Write intercepts the log bytes, masks any sensitive information, and forwards to the underlying writer.
// Returns len(p) on success to satisfy the io.Writer contract regardless of masked byte length differences.
func (mw *MaskingWriter) Write(p []byte) (n int, err error) {
	masked := redact.Bytes(p)
	if _, err := mw.w.Write(masked); err != nil {
		return 0, err
	}
	return len(p), nil
}

// WriteLevel implements zerolog.LevelWriter for level-aware writers.
func (mw *MaskingWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	masked := redact.Bytes(p)
	if lw, ok := mw.w.(zerolog.LevelWriter); ok {
		if _, err := lw.WriteLevel(level, masked); err != nil {
			return 0, err
		}
	} else {
		if _, err := mw.w.Write(masked); err != nil {
			return 0, err
		}
	}
	return len(p), nil
}
