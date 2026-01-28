package fieldsec

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

// CryptoAlgorithm is the algorithm identifier for the hybrid encryption scheme.
//
// This uses a combination of two algorithms:
//   - RSA-OAEP-256: Asymmetric encryption for key exchange (encrypts the AES key)
//   - AES-256-GCM: Symmetric encryption for data (encrypts the actual sensitive data)
//
// Why hybrid encryption?
//   - RSA alone can only encrypt up to 214 bytes (for RSA-2048)
//   - AES alone requires secure key sharing between sender and receiver
//   - Hybrid approach: AES encrypts unlimited data (fast), RSA securely transmits the AES key
//
// AES-256-GCM provides authenticated encryption, ensuring both confidentiality and integrity.
const CryptoAlgorithm = "RSA-OAEP-256+AES-256-GCM"

// KeyPair holds an RSA key pair with metadata.
type KeyPair struct {
	ID         string
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
	ExpiresAt  time.Time
}

// IsExpired returns true if the key pair has expired.
func (kp *KeyPair) IsExpired() bool {
	return time.Now().After(kp.ExpiresAt)
}

// PublicKeyBundle contains the public key data for transmission to clients.
type PublicKeyBundle struct {
	KeyID     string `json:"keyId"`
	Algorithm string `json:"algorithm"`
	PublicKey string `json:"publicKey"` // PEM encoded
	ExpiresAt string `json:"expiresAt"` // RFC3339 format
}

// ExportPublicBundle creates a bundle containing the public key for client use.
func (kp *KeyPair) ExportPublicBundle() (PublicKeyBundle, error) {
	pubASN1, err := x509.MarshalPKIXPublicKey(kp.PublicKey)
	if err != nil {
		return PublicKeyBundle{}, fmt.Errorf("marshal public key: %w", err)
	}

	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	return PublicKeyBundle{
		KeyID:     kp.ID,
		Algorithm: CryptoAlgorithm,
		PublicKey: string(pubPEM),
		ExpiresAt: kp.ExpiresAt.Format(time.RFC3339),
	}, nil
}

// ParsePublicKeyBundle parses a public key bundle and returns the RSA public key.
func ParsePublicKeyBundle(bundle PublicKeyBundle) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(bundle.PublicKey))
	if block == nil {
		return nil, ErrInvalidPublicKey
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidPublicKey, err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("%w: not an RSA key", ErrInvalidPublicKey)
	}

	return rsaPub, nil
}

// KeyStore manages encryption key pairs with automatic expiration.
type KeyStore struct {
	mu   sync.RWMutex
	keys map[string]*KeyPair
}

// NewKeyStore creates a new key store.
func NewKeyStore() *KeyStore {
	return &KeyStore{
		keys: make(map[string]*KeyPair),
	}
}

// GenerateKeyPair creates a new RSA-2048 key pair with the given validity duration.
func (ks *KeyStore) GenerateKeyPair(validity time.Duration) (*KeyPair, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("generate RSA key: %w", err)
	}

	kp := &KeyPair{
		ID:         uuid.New().String(),
		PublicKey:  &privateKey.PublicKey,
		PrivateKey: privateKey,
		ExpiresAt:  time.Now().Add(validity),
	}

	ks.mu.Lock()
	ks.keys[kp.ID] = kp
	ks.mu.Unlock()

	return kp, nil
}

// Get retrieves a key pair by ID.
func (ks *KeyStore) Get(id string) (*KeyPair, bool) {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	kp, ok := ks.keys[id]
	if !ok {
		return nil, false
	}

	return kp, true
}

// Delete removes a key pair from the store.
func (ks *KeyStore) Delete(id string) {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	delete(ks.keys, id)
}

// DeleteExpired removes all expired keys from the store.
func (ks *KeyStore) DeleteExpired() int {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	deleted := 0
	now := time.Now()
	for id, kp := range ks.keys {
		if now.After(kp.ExpiresAt) {
			delete(ks.keys, id)
			deleted++
		}
	}

	return deleted
}

// StartCleanupRoutine starts a background goroutine that periodically removes expired keys.
func (ks *KeyStore) StartCleanupRoutine(interval time.Duration, stop <-chan struct{}) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				ks.DeleteExpired()
			case <-stop:
				return
			}
		}
	}()
}

// encryptWithRSA encrypts data using RSA-OAEP with SHA-256.
func encryptWithRSA(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, plaintext, nil)
}

// decryptWithRSA decrypts data using RSA-OAEP with SHA-256.
func decryptWithRSA(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
}
