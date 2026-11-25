package controller

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cloud-barista/cm-beetle/transx"
	"github.com/stretchr/testify/assert"
)

func TestDecryptErrorMessage(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{name: "Direct ErrKeyNotFound", err: transx.ErrKeyNotFound, want: "Encryption key not found or already used"},
		{name: "Wrapped ErrKeyNotFound", err: fmt.Errorf("lookup: %w", transx.ErrKeyNotFound), want: "Encryption key not found or already used"},
		{name: "Direct ErrKeyExpired", err: transx.ErrKeyExpired, want: "Encryption key has expired"},
		{name: "Wrapped ErrKeyExpired", err: fmt.Errorf("expired: %w", transx.ErrKeyExpired), want: "Encryption key has expired"},
		{name: "Direct ErrKeyMismatch", err: transx.ErrKeyMismatch, want: "Encryption key ID mismatch"},
		{name: "Wrapped ErrKeyMismatch", err: fmt.Errorf("%w: expected a, got b", transx.ErrKeyMismatch), want: "Encryption key ID mismatch"},
		{name: "Other error", err: errors.New("bad ciphertext"), want: "Decryption failed"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, decryptErrorMessage(tt.err))
		})
	}
}
