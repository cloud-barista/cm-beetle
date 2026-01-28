// Package fieldsec provides field-level encryption for Go structs using hybrid encryption.
package fieldsec

import "errors"

// Encryption-related errors.
var (
	// ErrKeyNotFound indicates the requested key does not exist in the store.
	ErrKeyNotFound = errors.New("encryption key not found")

	// ErrKeyExpired indicates the key has expired and cannot be used.
	ErrKeyExpired = errors.New("encryption key expired")

	// ErrKeyMismatch indicates the provided key ID does not match the expected key.
	ErrKeyMismatch = errors.New("encryption key ID mismatch")

	// ErrDecryptionFailed indicates the decryption operation failed.
	ErrDecryptionFailed = errors.New("decryption failed")

	// ErrInvalidPublicKey indicates the public key format is invalid.
	ErrInvalidPublicKey = errors.New("invalid public key format")

	// ErrInvalidPath indicates the JSON path format is invalid.
	ErrInvalidPath = errors.New("invalid field path")

	// ErrFieldNotFound indicates the field was not found at the given path.
	ErrFieldNotFound = errors.New("field not found at path")
)
