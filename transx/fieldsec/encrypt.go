package fieldsec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
)

// encryptedPayload represents a hybrid-encrypted value.
// Stored as base64-encoded JSON in the model field.
type encryptedPayload struct {
	Version      int    `json:"v"`  // Schema version for future compatibility
	EncryptedKey []byte `json:"ek"` // AES key encrypted with RSA-OAEP
	Nonce        []byte `json:"n"`  // AES-GCM nonce (12 bytes)
	Ciphertext   []byte `json:"ct"` // AES-GCM encrypted data
}

// encryptValue encrypts a single string value using hybrid encryption.
// Uses AES-256-GCM for data and RSA-OAEP for key exchange.
func encryptValue(plaintext string, publicKey *rsa.PublicKey) (string, error) {
	if plaintext == "" {
		return "", nil
	}

	// 1. Generate random AES-256 key (32 bytes)
	aesKey := make([]byte, 32)
	if _, err := rand.Read(aesKey); err != nil {
		return "", fmt.Errorf("generate AES key: %w", err)
	}

	// 2. Encrypt data with AES-GCM
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", fmt.Errorf("create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("create GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("generate nonce: %w", err)
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	// 3. Encrypt AES key with RSA-OAEP
	encryptedKey, err := encryptWithRSA(aesKey, publicKey)
	if err != nil {
		return "", fmt.Errorf("encrypt AES key: %w", err)
	}

	// 4. Package everything
	payload := encryptedPayload{
		Version:      1,
		EncryptedKey: encryptedKey,
		Nonce:        nonce,
		Ciphertext:   ciphertext,
	}

	// 5. Encode as base64 JSON
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}

	return base64.StdEncoding.EncodeToString(jsonBytes), nil
}

// Encrypt encrypts specified fields in a struct using hybrid encryption.
// The model can be any struct that is JSON-serializable.
// Fields are specified as dot-separated JSON paths (e.g., "source.ssh.privateKey").
// Returns a new struct of the same type with encrypted fields.
//
// Parameters:
//   - model: any JSON-serializable struct
//   - fields: list of dot-separated JSON paths to encrypt
//   - publicKey: RSA public key for encryption
//   - keyID: identifier for the key (stored in the result for decryption)
//   - keyIDField: JSON path where the keyID should be stored (e.g., "encryptionKeyId")
//
// Example:
//
//	fields := []string{"source.ssh.privateKey", "destination.minio.secretKey"}
//	encrypted, err := Encrypt(model, fields, pubKey, "key-123", "encryptionKeyId")
func Encrypt(model any, fields []string, publicKey *rsa.PublicKey, keyID, keyIDField string) (any, error) {
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

	// 3. Set key ID if field is specified
	if keyIDField != "" && keyID != "" {
		m[keyIDField] = keyID
	}

	// 4. Encrypt each field
	for _, path := range fields {
		value, found := getFieldByPath(m, path)
		if !found || value == "" {
			continue // Skip missing or empty fields
		}

		encrypted, err := encryptValue(value, publicKey)
		if err != nil {
			return nil, fmt.Errorf("encrypt field %s: %w", path, err)
		}

		if !setFieldByPath(m, path, encrypted) {
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

// mapToStruct converts a map back to the original struct type.
func mapToStruct(m map[string]any, original any) (any, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	// Create a new instance of the same type
	resultPtr := reflect.New(reflect.TypeOf(original))
	if err := json.Unmarshal(data, resultPtr.Interface()); err != nil {
		return nil, err
	}

	return resultPtr.Elem().Interface(), nil
}
