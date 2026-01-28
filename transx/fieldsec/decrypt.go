package fieldsec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// decryptValue decrypts a hybrid-encrypted string value.
func decryptValue(encrypted string, privateKey *rsa.PrivateKey) (string, error) {
	if encrypted == "" {
		return "", nil
	}

	// 1. Decode base64 JSON
	jsonBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", fmt.Errorf("decode base64: %w", err)
	}

	var payload encryptedPayload
	if err := json.Unmarshal(jsonBytes, &payload); err != nil {
		return "", fmt.Errorf("unmarshal payload: %w", err)
	}

	// 2. Decrypt AES key with RSA
	aesKey, err := decryptWithRSA(payload.EncryptedKey, privateKey)
	if err != nil {
		return "", fmt.Errorf("decrypt AES key: %w", err)
	}

	// 3. Decrypt data with AES-GCM
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", fmt.Errorf("create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("create GCM: %w", err)
	}

	plaintext, err := gcm.Open(nil, payload.Nonce, payload.Ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decrypt data: %w", err)
	}

	return string(plaintext), nil
}

// Decrypt decrypts specified fields in a struct using hybrid decryption.
// The model can be any struct that is JSON-serializable.
// Fields are specified as dot-separated JSON paths (e.g., "source.ssh.privateKey").
// Returns a new struct of the same type with decrypted fields.
//
// Parameters:
//   - model: any JSON-serializable struct with encrypted fields
//   - fields: list of dot-separated JSON paths to decrypt
//   - privateKey: RSA private key for decryption
//   - keyIDField: JSON path where the keyID is stored (will be cleared after decryption)
//
// Example:
//
//	fields := []string{"source.ssh.privateKey", "destination.minio.secretKey"}
//	decrypted, err := Decrypt(model, fields, privKey, "encryptionKeyId")
func Decrypt(model any, fields []string, privateKey *rsa.PrivateKey, keyIDField string) (any, error) {
	// 1. Convert struct to JSON
	data, err := json.Marshal(model)
	if err != nil {
		return nil, fmt.Errorf("marshal model: %w", err)
	}

	// 2. Parse as map for field manipulation
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("unmarshal to map: %w", err)
	}

	// 3. Clear key ID field if specified
	if keyIDField != "" {
		m[keyIDField] = ""
	}

	// 4. Decrypt each field
	for _, path := range fields {
		value, found := getFieldByPath(m, path)
		if !found || value == "" {
			continue // Skip missing or empty fields
		}

		decrypted, err := decryptValue(value, privateKey)
		if err != nil {
			return nil, fmt.Errorf("decrypt field %s: %w", path, err)
		}

		if !setFieldByPath(m, path, decrypted) {
			return nil, fmt.Errorf("%w: %s", ErrInvalidPath, path)
		}
	}

	// 5. Convert back to original type
	result, err := mapToStruct(m, model)
	if err != nil {
		return nil, fmt.Errorf("convert to struct: %w", err)
	}

	return result, nil
}

// DecryptWithKeyPair validates the key pair and decrypts the model.
// It checks for key expiration and ID mismatch before decryption.
//
// Parameters:
//   - model: any JSON-serializable struct with encrypted fields
//   - fields: list of dot-separated JSON paths to decrypt
//   - keyPair: the key pair containing the private key
//   - expectedKeyID: the key ID that was used for encryption
//   - keyIDField: JSON path where the keyID is stored
func DecryptWithKeyPair(model any, fields []string, keyPair *KeyPair, expectedKeyID, keyIDField string) (any, error) {
	if keyPair == nil {
		return nil, ErrKeyNotFound
	}

	if keyPair.IsExpired() {
		return nil, ErrKeyExpired
	}

	if keyPair.ID != expectedKeyID {
		return nil, fmt.Errorf("%w: expected %s, got %s", ErrKeyMismatch, expectedKeyID, keyPair.ID)
	}

	return Decrypt(model, fields, keyPair.PrivateKey, keyIDField)
}

// DecryptWithStore retrieves the key from the store and decrypts the model.
// After successful decryption, the key is automatically deleted (one-time use).
//
// Parameters:
//   - model: any JSON-serializable struct with encrypted fields
//   - fields: list of dot-separated JSON paths to decrypt
//   - store: the key store containing the private key
//   - keyID: the key ID used for encryption
//   - keyIDField: JSON path where the keyID is stored
func DecryptWithStore(model any, fields []string, store *KeyStore, keyID, keyIDField string) (any, error) {
	keyPair, ok := store.Get(keyID)
	if !ok {
		return nil, ErrKeyNotFound
	}

	result, err := DecryptWithKeyPair(model, fields, keyPair, keyID, keyIDField)
	if err != nil {
		return nil, err
	}

	// Delete the key after successful decryption (one-time use)
	store.Delete(keyID)

	return result, nil
}
