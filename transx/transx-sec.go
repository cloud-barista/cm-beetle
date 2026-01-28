package transx

import (
	"crypto/rsa"
	"sync"
	"time"

	"github.com/cloud-barista/cm-beetle/transx/fieldsec"
)

// ============================================================================
// Singleton KeyStore Management
// ============================================================================

var (
	keyStore     *KeyStore
	keyStoreOnce sync.Once
	keyExpiry    time.Duration
)

// InitKeyStore initializes the singleton KeyStore and starts the cleanup routine.
// This should be called once from main() during server startup.
// Thread-safe: uses sync.Once to ensure single initialization.
//
// Parameters:
//   - keyExpiryDuration: duration after which generated keys expire (e.g., 30*time.Minute)
//   - cleanupInterval: interval for background cleanup of expired keys (e.g., 10*time.Minute)
func InitKeyStore(keyExpiryDuration, cleanupInterval time.Duration) {
	keyStoreOnce.Do(func() {
		keyStore = NewKeyStore()
		keyExpiry = keyExpiryDuration
		stopCleanup := make(chan struct{})
		keyStore.StartCleanupRoutine(cleanupInterval, stopCleanup)
	})
}

// GetKeyStore returns the singleton KeyStore instance.
// Panics if InitKeyStore() was not called.
func GetKeyStore() *KeyStore {
	if keyStore == nil {
		panic("transx: KeyStore not initialized. Call InitKeyStore() first.")
	}
	return keyStore
}

// GetKeyExpiry returns the configured key expiry duration.
// Panics if InitKeyStore() was not called.
func GetKeyExpiry() time.Duration {
	if keyStore == nil {
		panic("transx: KeyStore not initialized. Call InitKeyStore() first.")
	}
	return keyExpiry
}

// ============================================================================
// Sensitive Fields for DataMigrationModel
// ============================================================================
//
// The following fields are encrypted/decrypted by EncryptModel/DecryptModel:
//
//   SSH Private Keys:
//     - source.filesystem.ssh.privateKey
//     - destination.filesystem.ssh.privateKey
//
//   S3/Minio Credentials:
//     - source.objectStorage.minio.accessKeyId
//     - source.objectStorage.minio.secretAccessKey
//     - destination.objectStorage.minio.accessKeyId
//     - destination.objectStorage.minio.secretAccessKey
//
//   Spider Authentication:
//     - source.objectStorage.spider.auth.basic.password
//     - source.objectStorage.spider.auth.jwt.token
//     - destination.objectStorage.spider.auth.basic.password
//     - destination.objectStorage.spider.auth.jwt.token
//
//   Tumblebug Authentication:
//     - source.objectStorage.tumblebug.auth.basic.password
//     - source.objectStorage.tumblebug.auth.jwt.token
//     - destination.objectStorage.tumblebug.auth.basic.password
//     - destination.objectStorage.tumblebug.auth.jwt.token
//
// ============================================================================

// sensitiveFields defines the JSON paths of fields that contain sensitive data.
var sensitiveFields = []string{
	// SSH private keys
	"source.filesystem.ssh.privateKey",
	"destination.filesystem.ssh.privateKey",

	// S3/Minio credentials
	"source.objectStorage.minio.accessKeyId",
	"source.objectStorage.minio.secretAccessKey",
	"destination.objectStorage.minio.accessKeyId",
	"destination.objectStorage.minio.secretAccessKey",

	// Spider authentication
	"source.objectStorage.spider.auth.basic.password",
	"source.objectStorage.spider.auth.jwt.token",
	"destination.objectStorage.spider.auth.basic.password",
	"destination.objectStorage.spider.auth.jwt.token",

	// Tumblebug authentication
	"source.objectStorage.tumblebug.auth.basic.password",
	"source.objectStorage.tumblebug.auth.jwt.token",
	"destination.objectStorage.tumblebug.auth.basic.password",
	"destination.objectStorage.tumblebug.auth.jwt.token",
}

// keyIDField is the JSON field name where the encryption key ID is stored.
const keyIDField = "encryptionKeyId"

// Re-export types from fieldsec for convenience
type (
	KeyPair         = fieldsec.KeyPair
	KeyStore        = fieldsec.KeyStore
	PublicKeyBundle = fieldsec.PublicKeyBundle
)

// Re-export functions from fieldsec
var (
	NewKeyStore          = fieldsec.NewKeyStore
	ParsePublicKeyBundle = fieldsec.ParsePublicKeyBundle
)

// EncryptModel encrypts all sensitive fields in DataMigrationModel.
// Uses hybrid encryption (AES-256-GCM + RSA-OAEP) for fields of any size.
//
// Parameters:
//   - model: the DataMigrationModel to encrypt
//   - publicKey: RSA public key for encryption
//   - keyID: identifier for the key (for server-side lookup)
//
// Returns a new DataMigrationModel with encrypted fields and EncryptionKeyID set.
func EncryptModel(model DataMigrationModel, publicKey *rsa.PublicKey, keyID string) (DataMigrationModel, error) {
	result, err := fieldsec.Encrypt(model, sensitiveFields, publicKey, keyID, keyIDField)
	if err != nil {
		return DataMigrationModel{}, err
	}

	return result.(DataMigrationModel), nil
}

// DecryptModelWith decrypts all sensitive fields in DataMigrationModel using the provided key pair.
// If the model is not encrypted (EncryptionKeyID is empty), returns as-is.
// Use this for testing or when managing keys externally.
//
// Parameters:
//   - model: the DataMigrationModel to decrypt
//   - keyPair: the key pair containing the private key
//
// Returns a new DataMigrationModel with decrypted fields and EncryptionKeyID cleared.
func DecryptModelWith(model DataMigrationModel, keyPair *KeyPair) (DataMigrationModel, error) {
	if !model.IsEncrypted() {
		return model, nil
	}

	result, err := fieldsec.DecryptWithKeyPair(
		model,
		sensitiveFields,
		keyPair,
		model.EncryptionKeyID,
		keyIDField,
	)
	if err != nil {
		return DataMigrationModel{}, err
	}

	return result.(DataMigrationModel), nil
}

// DecryptModel decrypts all sensitive fields in DataMigrationModel using the singleton KeyStore.
// After successful decryption, the key is automatically deleted (one-time use).
// Panics if InitKeyStore() was not called.
//
// Parameters:
//   - model: the DataMigrationModel to decrypt
//
// Returns a new DataMigrationModel with decrypted fields and EncryptionKeyID cleared.
func DecryptModel(model DataMigrationModel) (DataMigrationModel, error) {
	if !model.IsEncrypted() {
		return model, nil
	}

	result, err := fieldsec.DecryptWithStore(
		model,
		sensitiveFields,
		GetKeyStore(),
		model.EncryptionKeyID,
		keyIDField,
	)
	if err != nil {
		return DataMigrationModel{}, err
	}

	return result.(DataMigrationModel), nil
}
