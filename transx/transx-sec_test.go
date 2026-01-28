package transx

import (
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
	decModel, err := DecryptModel(encModel, kp)
	if err != nil {
		t.Fatalf("DecryptModel failed: %v", err)
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

func TestDecryptModelWithStore(t *testing.T) {
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

	// Decrypt with store (should auto-delete key)
	decModel, err := DecryptModelWithStore(encModel, store)
	if err != nil {
		t.Fatalf("DecryptModelWithStore failed: %v", err)
	}

	// Verify decryption
	if decModel.Source.Filesystem.SSH.PrivateKey != model.Source.Filesystem.SSH.PrivateKey {
		t.Error("Decrypted value mismatch")
	}

	// Verify key is deleted (one-time use)
	_, ok := store.Get(kp.ID)
	if ok {
		t.Error("Key should be deleted after decryption (one-time use)")
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

	_, err = DecryptModel(encModel, kp)
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

	// DecryptModel should pass through non-encrypted model
	result, err := DecryptModel(model, kp)
	if err != nil {
		t.Fatalf("DecryptModel should not fail for non-encrypted model: %v", err)
	}
	if result.Source.Path != model.Source.Path {
		t.Error("Non-encrypted model should pass through unchanged")
	}

	// DecryptModelWithStore should also pass through
	result, err = DecryptModelWithStore(model, store)
	if err != nil {
		t.Fatalf("DecryptModelWithStore should not fail for non-encrypted model: %v", err)
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

	decModel, err := DecryptModel(encModel, kp)
	if err != nil {
		t.Fatalf("DecryptModel failed: %v", err)
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
