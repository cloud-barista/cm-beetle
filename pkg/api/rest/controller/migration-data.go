/*
Copyright 2019 The Cloud-Barista Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package controller has handlers and their request/response bodies for migration APIs
package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/transx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// Encryption Key API
// ============================================================================

// GetDataMigrationEncryptionKey godoc
// @ID GetDataMigrationEncryptionKey
// @Summary Get encryption public key for secure data migration
// @Description Generate and return a one-time RSA public key for encrypting sensitive fields in DataMigrationModel.
// @Description
// @Description [Encryption Workflow]
// @Description 1. Client calls this API to get a public key bundle
// @Description 2. Client encrypts sensitive fields using the public key
// @Description 3. Client sends encrypted model to POST /migration/data
// @Description 4. Server decrypts using the stored private key (auto-deleted after use)
// @Description
// @Description [Note]
// @Description * **One-time key**: Automatically invalidated after first decryption use
// @Description * **Key validity**: 30 minutes from generation (configurable in server)
// @Description * **Encrypted fields**: SSH privateKey, S3 credentials, auth passwords/tokens
// @Description * **Algorithm**: RSA-OAEP-256 (key exchange) + AES-256-GCM (data encryption)
// @Description
// @Description [Client Example]
// @Description See: https://github.com/cloud-barista/cm-beetle/blob/main/transx/README.md#usage-encrypted-transmission-recommended-for-production
// @Description
// @Tags [Migration] Data (incubating)
// @Accept json
// @Produce json
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Success 200 {object} model.ApiResponse[transx.PublicKeyBundle] "Public key bundle containing keyId, algorithm, publicKey (PEM), and expiresAt"
// @Failure 500 {object} model.ApiResponse[any] "Key generation failed"
// @Router /migration/data/encryptionKey [get]
func GetDataMigrationEncryptionKey(c echo.Context) error {
	log.Info().Msg("Generating encryption key for data migration")

	// Generate a one-time use key
	keyPair, err := transx.GetKeyStore().GenerateKeyPair(transx.GetKeyExpiry())
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate encryption key")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Key generation failed"))
	}

	// Export public key bundle for client
	bundle, err := keyPair.ExportPublicBundle()
	if err != nil {
		log.Error().Err(err).Msg("Failed to export public key bundle")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Key export failed"))
	}

	log.Info().
		Str("keyId", bundle.KeyID).
		Str("expiresAt", bundle.ExpiresAt).
		Msg("Encryption key generated successfully")

	return c.JSON(http.StatusOK, model.SuccessResponse(bundle))
}

// MigrateData godoc
// @ID MigrateData
// @Summary Migrate data from source to target
// @Description Migrate data from source to target. Supports both plaintext and encrypted requests.
// @Description
// @Description [Endpoint Requirements]
// @Description * Both source and destination must be remote endpoints (SSH or object storage)
// @Description * Local filesystem access is not allowed for security reasons
// @Description
// @Description [Transfer Options]
// @Description * Strategy: auto (default), direct, relay
// @Description * SSH: Supports PrivateKey content or PrivateKeyPath
// @Description
// @Description [Encryption Support]
// @Description * To encrypt sensitive fields, first call GET /migration/data/encryptionKey
// @Description * Encrypted requests include `encryptionKeyId` field
// @Description * Server automatically detects and decrypts encrypted requests
// @Description
// @Description [Examples]
// @Description * Test results: https://github.com/cloud-barista/cm-beetle/blob/main/docs/test-results-data-migration.md
// @Description
// @Tags [Migration] Data (incubating)
// @Accept  json
// @Produce  json
// @Param X-Request-Id header string false "Unique request ID (auto-generated if not provided). Used for tracking request status and correlating logs."
// @Param reqBody body transx.DataMigrationModel true "Data migration request (supports plaintext or encrypted with encryptionKeyId)"
// @Success 200 {object} model.ApiResponse[string] "Data migrated successfully"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters or decryption failed"
// @Failure 500 {object} model.ApiResponse[any] "Internal server error during data migration"
// @Router /migration/data [post]
func MigrateData(c echo.Context) error {

	req := new(transx.DataMigrationModel)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("failed to bind the request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// Decrypt if encrypted request
	if req.IsEncrypted() {
		log.Info().Str("keyId", req.EncryptionKeyID).Msg("Decrypting encrypted request")

		decryptedModel, err := transx.DecryptModel(*req)
		if err != nil {
			log.Error().Err(err).Str("keyId", req.EncryptionKeyID).Msg("Failed to decrypt request")

			switch err {
			case transx.ErrKeyNotFound:
				return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Encryption key not found or already used"))
			case transx.ErrKeyExpired:
				return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Encryption key has expired"))
			case transx.ErrKeyMismatch:
				return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Encryption key ID mismatch"))
			default:
				return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Decryption failed"))
			}
		}

		log.Info().Msg("Request decrypted successfully")
		req = &decryptedModel
	}

	err := transx.Validate(*req)
	if err != nil {
		log.Error().Err(err).Msg("invalid request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// Security check: Prevent access to local filesystem
	// API users must not access the server's local filesystem
	if req.Source.IsLocal() {
		log.Warn().Msg("rejected: source uses local filesystem")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Local filesystem access not allowed for source; use SSH or object storage"))
	}
	if req.Destination.IsLocal() {
		log.Warn().Msg("rejected: destination uses local filesystem")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Local filesystem access not allowed for destination; use SSH or object storage"))
	}

	log.Info().
		Str("sourceType", req.Source.StorageType).
		Str("sourcePath", req.Source.Path).
		Str("destType", req.Destination.StorageType).
		Str("destPath", req.Destination.Path).
		Str("strategy", req.Strategy).
		Msg("Starting data migration")

	// Start time measurement
	startTime := time.Now()

	// Execute migration
	err = transx.Transfer(*req)

	// Calculate elapsed time
	elapsedTime := time.Since(startTime)

	if err != nil {
		log.Error().Err(err).Dur("elapsedTime", elapsedTime).Msg("failed to migrate data")
		errorMsg := fmt.Sprintf("Data migration failed: %v (%s)", err, elapsedTime.Round(time.Millisecond))
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(errorMsg))
	}

	log.Info().Dur("elapsedTime", elapsedTime).Msg("Data migration completed successfully")
	successMsg := fmt.Sprintf("Data migrated successfully (%s)", elapsedTime.Round(time.Millisecond))
	return c.JSON(http.StatusOK, model.SimpleSuccessResponse(successMsg))
}

// ============================================================================
// Encryption Test APIs (for development/testing only)
// ============================================================================

// TestEncryptData godoc
// @ID TestEncryptData
// @Summary [TEST ONLY] Encrypt sensitive fields in data migration model
// @Description **⚠️ FOR TESTING ONLY**: In production, encryption MUST be performed client-side.
// @Description This API is provided solely for testing the encryption workflow without implementing client-side encryption.
// @Description
// @Description Receives a plaintext DataMigrationModel and returns an encrypted version.
// @Description The server generates a new key pair, encrypts sensitive fields, and returns the encrypted model.
// @Description
// @Description [Security Warning]
// @Description * Sending plaintext credentials to server defeats the purpose of encryption
// @Description * Use this API only for development/testing environments
// @Description * In production, use client-side encryption with GET /migration/data/encryptionKey
// @Description
// @Description [Test Workflow]
// @Description 1. Call GET /migration/data/encryptionKey to get a public key bundle
// @Description 2. Call this API with the public key bundle and plaintext model
// @Description 3. Receive encrypted model with keyId
// @Description 4. Call POST /migration/data/test/decrypt to verify decryption
// @Description
// @Tags [Migration] Data (incubating)
// @Accept json
// @Produce json
// @Param X-Request-Id header string false "Unique request ID"
// @Param reqBody body model.EncryptionTestRequest true "Public key bundle and plaintext model"
// @Success 200 {object} model.ApiResponse[transx.DataMigrationModel] "Encrypted model with encryptionKeyId set"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request"
// @Failure 500 {object} model.ApiResponse[any] "Encryption failed"
// @Router /migration/data/test/encrypt [post]
func TestEncryptData(c echo.Context) error {
	req := new(model.EncryptionTestRequest)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("failed to bind the request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// Validate public key bundle
	if req.PublicKeyBundle.KeyID == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("publicKeyBundle.keyId is required"))
	}
	if req.PublicKeyBundle.PublicKey == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("publicKeyBundle.publicKey is required"))
	}

	// Check if model is already encrypted
	if req.Model.IsEncrypted() {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Model is already encrypted"))
	}

	log.Warn().Msg("[TEST API] Encrypting model server-side - this should be done client-side in production")

	// Parse the public key from the bundle
	publicKey, err := transx.ParsePublicKeyBundle(req.PublicKeyBundle)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse public key")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid public key format"))
	}

	// Encrypt the model using the provided public key
	encryptedModel, err := transx.EncryptModel(req.Model, publicKey, req.PublicKeyBundle.KeyID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encrypt model")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Encryption failed: "+err.Error()))
	}

	log.Info().
		Str("keyId", req.PublicKeyBundle.KeyID).
		Msg("[TEST API] Encryption test completed")

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(encryptedModel, "Encryption successful. Use this encrypted model with POST /migration/data or POST /migration/data/test/decrypt"))
}

// TestDecryptData godoc
// @ID TestDecryptData
// @Summary [TEST ONLY] Test decryption of encrypted data migration model
// @Description **⚠️ FOR TESTING ONLY**: This API tests server-side decryption without executing migration.
// @Description Use this to verify the encryption workflow before calling the actual migration API.
// @Description
// @Description Receives an encrypted DataMigrationModel, decrypts it, and returns verification results.
// @Description This API does NOT execute actual data migration.
// @Description
// @Description [Test Workflow]
// @Description 1. Call GET /migration/data/encryptionKey to get a public key
// @Description 2. Encrypt your model using the public key (client-side), or use POST /migration/data/test/encrypt
// @Description 3. Call this API with the encrypted model
// @Description 4. Verify the decryption result shows expected fields
// @Description
// @Description [Note]
// @Description * The encryption key is consumed (deleted) after this test
// @Description * You need to generate a new key for actual migration
// @Description
// @Tags [Migration] Data (incubating)
// @Accept json
// @Produce json
// @Param X-Request-Id header string false "Unique request ID"
// @Param reqBody body transx.DataMigrationModel true "Encrypted data migration model"
// @Success 200 {object} model.ApiResponse[transx.DataMigrationModel] "Decrypted model (sensitive fields restored)"
// @Failure 400 {object} model.ApiResponse[any] "Decryption failed or model not encrypted"
// @Router /migration/data/test/decrypt [post]
func TestDecryptData(c echo.Context) error {
	req := new(transx.DataMigrationModel)
	if err := c.Bind(req); err != nil {
		log.Error().Err(err).Msg("failed to bind the request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	// Check if model is encrypted
	if !req.IsEncrypted() {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Model is not encrypted (encryptionKeyId is empty)"))
	}

	log.Info().Str("keyId", req.EncryptionKeyID).Msg("Testing decryption")

	// Attempt to decrypt
	decryptedModel, err := transx.DecryptModel(*req)
	if err != nil {
		log.Error().Err(err).Str("keyId", req.EncryptionKeyID).Msg("Decryption test failed")

		switch err {
		case transx.ErrKeyNotFound:
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Encryption key not found or already used"))
		case transx.ErrKeyExpired:
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Encryption key has expired"))
		case transx.ErrKeyMismatch:
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Encryption key ID mismatch"))
		default:
			return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Decryption failed: "+err.Error()))
		}
	}

	log.Info().
		Str("keyId", req.EncryptionKeyID).
		Msg("Decryption test passed")

	return c.JSON(http.StatusOK, model.SuccessResponseWithMessage(decryptedModel, "Decryption successful"))
}
