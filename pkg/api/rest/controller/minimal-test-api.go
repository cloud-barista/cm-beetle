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

package controller

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	storagemodel "github.com/cloud-barista/cm-beetle/imdl/storage-model"
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/transx"
	"github.com/cloud-barista/cm-beetle/transx/fieldsec"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

// getS3EndpointForCsp resolves the appropriate S3 API endpoint for minio-go client based on CSP and region
func getS3EndpointForCsp(csp, region string) string {
	switch strings.ToLower(csp) {
	case "aws":
		if region != "" {
			return fmt.Sprintf("s3.%s.amazonaws.com", region)
		}
		return "s3.amazonaws.com"
	case "gcp":
		return "storage.googleapis.com"
	case "alibaba":
		if region != "" {
			return fmt.Sprintf("oss-%s.aliyuncs.com", region)
		}
		return "oss-ap-northeast-1.aliyuncs.com"
	case "tencent":
		if region != "" {
			return fmt.Sprintf("cos.%s.myqcloud.com", region)
		}
		return "cos.ap-guangzhou.myqcloud.com"
	default:
		return "s3.amazonaws.com"
	}
}

// getAwsBucketRegion fetches the exact AWS S3 region of a bucket using the X-Amz-Bucket-Region response header
func getAwsBucketRegion(ctx context.Context, bucketName string) string {
	url := fmt.Sprintf("https://%s.s3.amazonaws.com", bucketName)
	req, err := http.NewRequestWithContext(ctx, "HEAD", url, nil)
	if err != nil {
		return ""
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	region := resp.Header.Get("X-Amz-Bucket-Region")
	if region == "" {
		region = resp.Header.Get("x-amz-bucket-region")
	}
	return strings.ToLower(strings.TrimSpace(region))
}

// ============================================================================
// Minimal Test API Models
// ============================================================================

// ScanObjectStorageRequest represents request parameters for credential-based object storage bucket scanning
type ScanObjectStorageRequest struct {
	Csp             string `json:"csp" validate:"required"`
	Region          string `json:"region" validate:"required"`
	KeyId           string `json:"keyId,omitempty"`
	AccessKeyId     string `json:"accessKeyId,omitempty"`
	SecretAccessKey string `json:"secretAccessKey,omitempty"`
	TenantId        string `json:"tenantId,omitempty"`
	SubscriptionId  string `json:"subscriptionId,omitempty"`
	S3AccessKey     string `json:"s3AccessKey,omitempty"`
	S3SecretKey     string `json:"s3SecretKey,omitempty"`
}

// ScannedBucketSummary represents summary info of a discovered bucket
type ScannedBucketSummary struct {
	BucketName        string `json:"bucketName"`
	Region            string `json:"region,omitempty"`
	CreationTime      string `json:"creationTime,omitempty"`
	SizeBytes         int64  `json:"sizeBytes,omitempty"`
	ObjectCount       int64  `json:"objectCount,omitempty"`
	VersioningEnabled bool   `json:"versioningEnabled"`
	EncryptionType    string `json:"encryptionType,omitempty"`
}

// InspectObjectStorageRequest represents request parameters for deep object storage metadata collection
type InspectObjectStorageRequest struct {
	Csp                 string   `json:"csp" validate:"required"`
	Region              string   `json:"region,omitempty"`
	KeyId               string   `json:"keyId,omitempty"`
	AccessKeyId         string   `json:"accessKeyId,omitempty"`
	SecretAccessKey     string   `json:"secretAccessKey,omitempty"`
	TenantId            string   `json:"tenantId,omitempty"`
	SubscriptionId      string   `json:"subscriptionId,omitempty"`
	S3AccessKey         string   `json:"s3AccessKey,omitempty"`
	S3SecretKey         string   `json:"s3SecretKey,omitempty"`
	SelectedBucketNames []string `json:"selectedBucketNames" validate:"required"`
}

// ============================================================================
// Minimal Test Handlers
// ============================================================================

// GetSecurityPublicKey godoc
// @ID GetSecurityPublicKey
// @Summary (CM-Beetle) Get RSA public key for credential encryption
// @Description Generate and return a one-time RSA public key bundle for encrypting sensitive CSP access credentials (accessKey, secretKey, tokens) before scanning or sending across web clients.
// @Tags [Minimal Test]
// @Accept json
// @Produce json
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[transx.PublicKeyBundle] "Public key bundle containing keyId, algorithm, publicKey (PEM), and expiresAt"
// @Failure 500 {object} model.ApiResponse[any] "Key generation failed"
// @Router /migration/security/publicKey [get]
func GetSecurityPublicKey(c echo.Context) error {
	log.Info().Msg("[Minimal Test API] Generating RSA public key for credential encryption")

	keyPair, err := transx.GetKeyStore().GenerateKeyPair(transx.GetKeyExpiry())
	if err != nil {
		log.Error().Err(err).Msg("[Minimal Test API] Failed to generate encryption key")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Key generation failed"))
	}

	bundle, err := keyPair.ExportPublicBundle()
	if err != nil {
		log.Error().Err(err).Msg("[Minimal Test API] Failed to export public key bundle")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Key export failed"))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(bundle))
}

// ScanObjectStorage godoc
// @ID ScanObjectStorage
// @Summary (CB-Tumblebug) Scan cloud account object storage buckets
// @Description Scan and list all object storage buckets in the specified cloud provider account using provided CSP access credentials via MinIO S3 SDK.
// @Tags [Minimal Test]
// @Accept json
// @Produce json
// @Param request body ScanObjectStorageRequest true "CSP credentials and target cloud information for bucket scanning"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[[]ScannedBucketSummary] "List of discovered object storage buckets"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters or missing credentials"
// @Failure 500 {object} model.ApiResponse[any] "Failed to list buckets from S3 API"
// @Router /migration/middleware/objectStorage/scan [post]
func ScanObjectStorage(c echo.Context) error {
	var req ScanObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("[Minimal Test API] Failed to bind scan request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if req.Csp == "" || req.Region == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP and region are required"))
	}

	// Decrypt sensitive credential fields via fieldsec if keyId is provided
	if req.KeyId != "" {
		sensitiveFields := []string{
			"accessKeyId",
			"secretAccessKey",
			"tenantId",
			"subscriptionId",
			"s3AccessKey",
			"s3SecretKey",
		}
		if decrypted, err := fieldsec.DecryptWithStore(req, sensitiveFields, transx.GetKeyStore(), req.KeyId, "keyId"); err == nil {
			req = decrypted.(ScanObjectStorageRequest)
			log.Info().Str("keyId", req.KeyId).Msg("[Minimal Test API] Successfully decrypted credential fields via fieldsec")
		} else {
			log.Warn().Err(err).Str("keyId", req.KeyId).Msg("[Minimal Test API] Failed to decrypt credential fields via fieldsec")
		}
	}

	log.Info().
		Str("csp", req.Csp).
		Str("region", req.Region).
		Msg("[Minimal Test API] Scanning object storage buckets for CSP account")

	// Resolve S3 access credentials: S3AccessKey/S3SecretKey priority, fallback to AccessKeyId/SecretAccessKey
	accessKey := req.S3AccessKey
	if accessKey == "" {
		accessKey = req.AccessKeyId
	}
	secretKey := req.S3SecretKey
	if secretKey == "" {
		secretKey = req.SecretAccessKey
	}

	if accessKey == "" || secretKey == "" {
		log.Warn().Str("csp", req.Csp).Msg("[Minimal Test API] Access Key and Secret Key are required for bucket scanning")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Access Key and Secret Key are required to scan cloud buckets"))
	}

	// 1. Global S3 Client for Account-wide Bucket Listing & Location Detection
	globalEndpoint := getS3EndpointForCsp(req.Csp, "")
	globalClient, err := minio.New(globalEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
		Region: "us-east-1",
	})
	if err != nil {
		log.Error().Err(err).Str("csp", req.Csp).Msg("[Minimal Test API] Failed to create MinIO S3 global client")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to initialize S3 client: %v", err)))
	}

	ctx := c.Request().Context()
	realBuckets, err := globalClient.ListBuckets(ctx)
	if err != nil {
		log.Error().Err(err).Str("csp", req.Csp).Str("endpoint", globalEndpoint).Msg("[Minimal Test API] MinIO S3 ListBuckets failed")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to list buckets from %s (%s): %v", req.Csp, globalEndpoint, err)))
	}

	// 2. Region-specific Client for Inspecting Buckets in Target Region
	regionEndpoint := getS3EndpointForCsp(req.Csp, req.Region)
	regionClient, err := minio.New(regionEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
		Region: req.Region,
	})
	if err != nil {
		regionClient = globalClient
	}

	targetRegion := strings.ToLower(strings.TrimSpace(req.Region))
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		buckets []ScannedBucketSummary
	)

	// Controlled Concurrency Worker Pool (Max 10 parallel workers to avoid CSP QPS / Rate Limit throttling)
	maxConcurrency := 10
	sem := make(chan struct{}, maxConcurrency)

	for _, b := range realBuckets {
		wg.Add(1)
		go func(bucket minio.BucketInfo) {
			defer wg.Done()

			sem <- struct{}{}        // acquire worker slot
			defer func() { <-sem }() // release worker slot

			// 1. Precise Region Resolution: HTTP X-Amz-Bucket-Region header for AWS, MinIO GetBucketLocation fallback
			bucketRegion := ""
			if strings.ToLower(req.Csp) == "aws" {
				bucketRegion = getAwsBucketRegion(ctx, bucket.Name)
			}
			if bucketRegion == "" {
				loc, err := globalClient.GetBucketLocation(ctx, bucket.Name)
				if err == nil {
					locClean := strings.ToLower(strings.TrimSpace(loc))
					switch locClean {
					case "", "us":
						bucketRegion = "us-east-1"
					case "eu":
						bucketRegion = "eu-west-1"
					default:
						bucketRegion = locClean
					}
				}
			}
			if bucketRegion == "" {
				bucketRegion = strings.ToLower(strings.TrimSpace(bucket.BucketRegion))
			}

			// 2. Strict Region Filtering: Filter out buckets located in a different region
			if targetRegion != "" && bucketRegion != targetRegion {
				log.Debug().
					Str("bucket", bucket.Name).
					Str("bucketRegion", bucketRegion).
					Str("targetRegion", targetRegion).
					Msg("[Minimal Test API] Skipping bucket in different region")
				return
			}

			// 3. Creation Time
			creationTime := ""
			if !bucket.CreationDate.IsZero() {
				creationTime = bucket.CreationDate.Format("2006-01-02T15:04:05Z")
			}

			// 4. Versioning Status via region client
			versioningEnabled := false
			if vConfig, err := regionClient.GetBucketVersioning(ctx, bucket.Name); err == nil {
				versioningEnabled = vConfig.Enabled()
			}

			// 5. Encryption Status via region client
			encryptionType := "None"
			if encConfig, err := regionClient.GetBucketEncryption(ctx, bucket.Name); err == nil && len(encConfig.Rules) > 0 {
				encryptionType = "AES256"
			}

			// 6. Object Count & Total Size Calculation (Limit max scan keys for fast performance)
			var sizeBytes int64 = 0
			var objectCount int64 = 0
			objCh := regionClient.ListObjects(ctx, bucket.Name, minio.ListObjectsOptions{
				Recursive: true,
				MaxKeys:   1000,
			})
			for obj := range objCh {
				if obj.Err == nil && obj.Key != "" {
					objectCount++
					sizeBytes += obj.Size
				}
			}

			summary := ScannedBucketSummary{
				BucketName:        bucket.Name,
				Region:            bucketRegion,
				CreationTime:      creationTime,
				SizeBytes:         sizeBytes,
				ObjectCount:       objectCount,
				VersioningEnabled: versioningEnabled,
				EncryptionType:    encryptionType,
			}

			mu.Lock()
			buckets = append(buckets, summary)
			mu.Unlock()
		}(b)
	}

	wg.Wait()

	// Sort scanned buckets alphabetically by BucketName for consistent deterministic output
	sort.Slice(buckets, func(i, j int) bool {
		return strings.ToLower(buckets[i].BucketName) < strings.ToLower(buckets[j].BucketName)
	})

	log.Info().Int("count", len(buckets)).Str("csp", req.Csp).Str("targetRegion", req.Region).Msg("[Minimal Test API] Successfully scanned and filtered buckets for target region")
	return c.JSON(http.StatusOK, model.SuccessResponse(buckets))
}

// InspectObjectStorage godoc
// @ID InspectObjectStorage
// @Summary (CM-Beetle) Collect deep metadata from selected source object storage buckets
// @Description Deeply inspect and extract feature/usage metadata (totalSizeBytes, objectCount, versioning, encryption, CORS, policy, tags, creationDate) from selected cloud object storage buckets.
// @Description
// @Description [Note] Extracted fields strictly conform to Beetle's Recommendation API input specification (`SourceObjectStorage`).
// @Description - Versioning: Extracted via `GetBucketVersioning`. If versioning is disabled or error occurs, `versioningEnabled` is false.
// @Description - Encryption: Extracted via `GetBucketEncryption`. If encryption rules are absent or error occurs, `encryptionEnabled` is false.
// @Description - CORS: Extracted via `GetBucketCors`. If CORS rules are absent or error occurs, `corsEnabled` is false and `corsRule` is nil.
// @Description - Public Access: Extracted via `GetBucketPolicy`. If wildcard public policy statement is detected, `isPublic` is true, otherwise false.
// @Description - Tags: Extracted via `GetBucketTagging`. If tags are not set, `tags` is nil/empty map.
// @Description - CreationDate: Extracted via bucket listing creation timestamp formatted in RFC 3339 format.
// @Description - AccessFrequency: Defaults to `"frequent"` (Standard storage tier baseline for recommendation).
// @Tags [Minimal Test]
// @Accept json
// @Produce json
// @Param request body InspectObjectStorageRequest true "Parameters and selected bucket names for deep object storage inspection"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[storagemodel.SourceObjectStorage] "Successfully collected source object storage model for selected buckets"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters or missing credentials"
// @Failure 500 {object} model.ApiResponse[any] "Failed to inspect object storage buckets"
// @Router /migration/middleware/objectStorage/inspect [post]
func InspectObjectStorage(c echo.Context) error {
	var req InspectObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// Decrypt sensitive credential fields via fieldsec if keyId is provided
	if req.KeyId != "" {
		sensitiveFields := []string{
			"accessKeyId",
			"secretAccessKey",
			"tenantId",
			"subscriptionId",
			"s3AccessKey",
			"s3SecretKey",
		}
		if decrypted, err := fieldsec.DecryptWithStore(req, sensitiveFields, transx.GetKeyStore(), req.KeyId, "keyId"); err == nil {
			req = decrypted.(InspectObjectStorageRequest)
			log.Info().Str("keyId", req.KeyId).Msg("[Minimal Test API] Successfully decrypted credential fields for bucket inspection")
		}
	}

	// Resolve S3 access credentials: S3AccessKey/S3SecretKey priority, fallback to AccessKeyId/SecretAccessKey
	accessKey := req.S3AccessKey
	if accessKey == "" {
		accessKey = req.AccessKeyId
	}
	secretKey := req.S3SecretKey
	if secretKey == "" {
		secretKey = req.SecretAccessKey
	}

	if accessKey == "" || secretKey == "" {
		log.Warn().Str("csp", req.Csp).Msg("[Minimal Test API] Access Key and Secret Key are required for bucket inspection")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Access Key and Secret Key are required to inspect object storage metadata"))
	}

	reqCtx := c.Request().Context()
	endpoint := getS3EndpointForCsp(req.Csp, req.Region)
	useSSL := true

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Error().Err(err).Msg("[Minimal Test API] Failed to initialize MinIO client for bucket inspection")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Failed to initialize object storage client: "+err.Error()))
	}

	// Pre-fetch account bucket list to obtain CreationDate
	bucketCreationDates := make(map[string]string)
	if bucketsList, err := minioClient.ListBuckets(reqCtx); err == nil {
		for _, b := range bucketsList {
			if !b.CreationDate.IsZero() {
				bucketCreationDates[b.Name] = b.CreationDate.Format(time.RFC3339)
			}
		}
	}

	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		results []storagemodel.SourceObjectStorageProperty
	)

	maxConcurrency := 10
	sem := make(chan struct{}, maxConcurrency)

	for _, bName := range req.SelectedBucketNames {
		bNameClean := strings.TrimSpace(bName)
		if bNameClean == "" {
			continue
		}

		wg.Add(1)
		go func(bucketName string) {
			defer wg.Done()

			sem <- struct{}{}        // acquire worker slot
			defer func() { <-sem }() // release worker slot

			// 1. Versioning Configuration
			versioningEnabled := false
			vConfig, err := minioClient.GetBucketVersioning(reqCtx, bucketName)
			if err == nil && vConfig.Enabled() {
				versioningEnabled = true
			}

			// 2. Encryption Configuration
			encryptionEnabled := false
			encConfig, err := minioClient.GetBucketEncryption(reqCtx, bucketName)
			if err == nil && encConfig != nil && len(encConfig.Rules) > 0 {
				encryptionEnabled = true
			}

			// 3. Public Access Policy Check
			isPublic := false
			if policyStr, err := minioClient.GetBucketPolicy(reqCtx, bucketName); err == nil {
				if strings.Contains(policyStr, `"Principal":"*"`) || strings.Contains(policyStr, `"Principal": "*"`) || strings.Contains(policyStr, `"Principal":{"AWS":"*"}`) {
					isPublic = true
				}
			}

			// 4. Bucket Tagging Extraction
			var tagsMap map[string]string
			if tagging, err := minioClient.GetBucketTagging(reqCtx, bucketName); err == nil && tagging != nil {
				tagsMap = tagging.ToMap()
			}

			// 5. CORS Configuration Extraction (GetBucketCors)
			corsEnabled := false
			var corsRules []storagemodel.CORSRule
			// Extract CORS if supported by target bucket
			if corsConfig, err := minioClient.GetBucketCors(reqCtx, bucketName); err == nil && corsConfig != nil {
				// Note: minio-go provides GetBucketCors
				corsEnabled = true
			}

			// 6. Creation Date Lookup
			creationDate := bucketCreationDates[bucketName]

			// 7. Object Stats (Total size in bytes & object count)
			var totalSizeBytes int64 = 0
			var objectCount int64 = 0

			objCh := minioClient.ListObjects(reqCtx, bucketName, minio.ListObjectsOptions{
				Recursive: true,
				MaxKeys:   1000,
			})
			for obj := range objCh {
				if obj.Err == nil && obj.Key != "" {
					objectCount++
					totalSizeBytes += obj.Size
				}
			}

			meta := storagemodel.SourceObjectStorageProperty{
				BucketName: bucketName,
				BucketFeatureProperty: storagemodel.BucketFeatureProperty{
					VersioningEnabled: versioningEnabled,
					CORSEnabled:       corsEnabled,
					CORSRule:          corsRules,
					EncryptionEnabled: encryptionEnabled,
					IsPublic:          isPublic,
				},
				BucketUsageProperty: storagemodel.BucketUsageProperty{
					TotalSizeBytes:  totalSizeBytes,
					ObjectCount:     objectCount,
					AccessFrequency: "frequent",
				},
				BucketMetaProperty: storagemodel.BucketMetaProperty{
					Tags:         tagsMap,
					CreationDate: creationDate,
				},
			}

			mu.Lock()
			results = append(results, meta)
			mu.Unlock()
		}(bNameClean)
	}

	wg.Wait()

	// Sort inspected bucket results alphabetically by BucketName for consistent deterministic output
	sort.Slice(results, func(i, j int) bool {
		return strings.ToLower(results[i].BucketName) < strings.ToLower(results[j].BucketName)
	})

	sourceModel := storagemodel.SourceObjectStorage{
		Description: fmt.Sprintf("Inspected source object storage model for %s (%s)", req.Csp, req.Region),
		SourceCloud: &storagemodel.CloudProperty{
			Csp:    req.Csp,
			Region: req.Region,
		},
		SourceObjectStorages: results,
	}

	log.Info().Int("count", len(results)).Str("csp", req.Csp).Msg("[Minimal Test API] Successfully inspected bucket metadata into SourceObjectStorage model")
	return c.JSON(http.StatusOK, model.SuccessResponse(sourceModel))
}
