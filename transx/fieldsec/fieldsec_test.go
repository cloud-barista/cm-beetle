package fieldsec

import (
	"testing"
	"time"
)

func TestKeyPairGeneration(t *testing.T) {
	store := NewKeyStore()
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
}

func TestPublicKeyBundle(t *testing.T) {
	store := NewKeyStore()
	kp, _ := store.GenerateKeyPair(5 * time.Minute)

	bundle, err := kp.ExportPublicBundle()
	if err != nil {
		t.Fatalf("ExportPublicBundle failed: %v", err)
	}

	if bundle.KeyID != kp.ID {
		t.Errorf("KeyID mismatch: got %s, want %s", bundle.KeyID, kp.ID)
	}
	if bundle.Algorithm != CryptoAlgorithm {
		t.Errorf("Algorithm mismatch: got %s, want %s", bundle.Algorithm, CryptoAlgorithm)
	}

	// Parse back
	pubKey, err := ParsePublicKeyBundle(bundle)
	if err != nil {
		t.Fatalf("ParsePublicKeyBundle failed: %v", err)
	}
	if pubKey == nil {
		t.Error("Parsed public key should not be nil")
	}
}

func TestEncryptDecryptValue(t *testing.T) {
	store := NewKeyStore()
	kp, _ := store.GenerateKeyPair(5 * time.Minute)

	testCases := []struct {
		name      string
		plaintext string
	}{
		{"short text", "Hello, World!"},
		{"medium text", "This is a medium length text for testing encryption."},
		{"long text", generateLongText(3000)},
		{"empty text", ""},
		{"special characters", "!@#$%^&*()_+{}|:<>?\"'\\n\\t"},
		{"unicode", "안녕하세요 世界 🌍"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := encryptValue(tc.plaintext, kp.PublicKey)
			if err != nil {
				t.Fatalf("encryptValue failed: %v", err)
			}

			if tc.plaintext == "" {
				if encrypted != "" {
					t.Error("Empty plaintext should return empty encrypted string")
				}
				return
			}

			if encrypted == tc.plaintext {
				t.Error("Encrypted text should be different from plaintext")
			}

			decrypted, err := decryptValue(encrypted, kp.PrivateKey)
			if err != nil {
				t.Fatalf("decryptValue failed: %v", err)
			}

			if decrypted != tc.plaintext {
				t.Errorf("Decrypted text mismatch: got %q, want %q", decrypted, tc.plaintext)
			}
		})
	}
}

func TestEncryptDecryptStruct(t *testing.T) {
	store := NewKeyStore()
	kp, _ := store.GenerateKeyPair(5 * time.Minute)

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	type Config struct {
		Host        string      `json:"host"`
		Credentials Credentials `json:"credentials"`
		KeyID       string      `json:"keyId,omitempty"`
	}

	original := Config{
		Host: "localhost",
		Credentials: Credentials{
			Username: "admin",
			Password: "secret123",
		},
	}

	fields := []string{"credentials.password"}

	// Encrypt
	result, err := Encrypt(original, fields, kp.PublicKey, kp.ID, "keyId")
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	encrypted := result.(Config)

	// Verify encryption
	if encrypted.KeyID != kp.ID {
		t.Errorf("KeyID not set: got %s, want %s", encrypted.KeyID, kp.ID)
	}
	if encrypted.Credentials.Password == original.Credentials.Password {
		t.Error("Password should be encrypted")
	}
	if encrypted.Credentials.Username != original.Credentials.Username {
		t.Error("Username should not be modified")
	}

	// Decrypt
	result, err = Decrypt(encrypted, fields, kp.PrivateKey, "keyId")
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	decrypted := result.(Config)

	// Verify decryption
	if decrypted.Credentials.Password != original.Credentials.Password {
		t.Errorf("Password mismatch: got %s, want %s", decrypted.Credentials.Password, original.Credentials.Password)
	}
	if decrypted.KeyID != "" {
		t.Error("KeyID should be cleared after decryption")
	}
}

func TestDecryptWithStore(t *testing.T) {
	store := NewKeyStore()
	kp, _ := store.GenerateKeyPair(5 * time.Minute)

	type Data struct {
		Secret string `json:"secret"`
		KeyID  string `json:"keyId,omitempty"`
	}

	original := Data{Secret: "my-secret"}
	fields := []string{"secret"}

	// Encrypt
	result, _ := Encrypt(original, fields, kp.PublicKey, kp.ID, "keyId")
	encrypted := result.(Data)

	// Decrypt with store
	result, err := DecryptWithStore(encrypted, fields, store, kp.ID, "keyId")
	if err != nil {
		t.Fatalf("DecryptWithStore failed: %v", err)
	}

	decrypted := result.(Data)
	if decrypted.Secret != original.Secret {
		t.Errorf("Secret mismatch: got %s, want %s", decrypted.Secret, original.Secret)
	}

	// Key should be deleted (one-time use)
	_, ok := store.Get(kp.ID)
	if ok {
		t.Error("Key should be deleted after decryption")
	}
}

func TestKeyExpiration(t *testing.T) {
	store := NewKeyStore()
	kp, _ := store.GenerateKeyPair(1 * time.Millisecond)

	time.Sleep(10 * time.Millisecond)

	if !kp.IsExpired() {
		t.Error("Key should be expired")
	}

	type Data struct {
		Secret string `json:"secret"`
		KeyID  string `json:"keyId,omitempty"`
	}

	original := Data{Secret: "test"}
	fields := []string{"secret"}

	result, _ := Encrypt(original, fields, kp.PublicKey, kp.ID, "keyId")
	encrypted := result.(Data)

	_, err := DecryptWithKeyPair(encrypted, fields, kp, kp.ID, "keyId")
	if err != ErrKeyExpired {
		t.Errorf("Expected ErrKeyExpired, got: %v", err)
	}
}

func TestMissingField(t *testing.T) {
	store := NewKeyStore()
	kp, _ := store.GenerateKeyPair(5 * time.Minute)

	type Data struct {
		Name string `json:"name"`
	}

	original := Data{Name: "test"}
	fields := []string{"nonexistent.field"}

	// Should not fail, just skip missing fields
	result, err := Encrypt(original, fields, kp.PublicKey, kp.ID, "")
	if err != nil {
		t.Fatalf("Encrypt should not fail for missing fields: %v", err)
	}

	encrypted := result.(Data)
	if encrypted.Name != original.Name {
		t.Error("Existing fields should not be modified")
	}
}

func TestPathOperations(t *testing.T) {
	m := map[string]any{
		"level1": map[string]any{
			"level2": map[string]any{
				"value": "secret",
			},
		},
	}

	// Get
	val, found := getFieldByPath(m, "level1.level2.value")
	if !found {
		t.Error("Field should be found")
	}
	if val != "secret" {
		t.Errorf("Value mismatch: got %s, want secret", val)
	}

	// Set
	ok := setFieldByPath(m, "level1.level2.value", "encrypted")
	if !ok {
		t.Error("Set should succeed")
	}

	val, _ = getFieldByPath(m, "level1.level2.value")
	if val != "encrypted" {
		t.Errorf("Value should be updated: got %s, want encrypted", val)
	}

	// Missing path
	_, found = getFieldByPath(m, "nonexistent.path")
	if found {
		t.Error("Missing path should return false")
	}
}

func generateLongText(length int) string {
	base := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range length {
		result[i] = base[i%len(base)]
	}
	return string(result)
}
