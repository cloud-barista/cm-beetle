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
	"github.com/cloud-barista/cm-beetle/pkg/csp"
	"github.com/cloud-barista/cm-beetle/transx"
	"github.com/cloud-barista/cm-beetle/transx/fieldsec"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

// isValidCloudRegion checks that a cloud region identifier contains only safe characters.
func isValidCloudRegion(region string) bool {
	if len(region) == 0 || len(region) > 64 {
		return false
	}
	for _, r := range region {
		isAlphaNum := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
		if !isAlphaNum && r != '-' && r != '_' {
			return false
		}
	}
	return true
}

// getS3EndpointForCsp resolves the S3-compatible endpoint for supported CSPs without silent fallback.
func getS3EndpointForCsp(cspName, region string) (string, error) {
	cspKey := strings.ToLower(strings.TrimSpace(cspName))
	regionClean := strings.ToLower(strings.TrimSpace(region))
	if regionClean != "" && !isValidCloudRegion(regionClean) {
		return "", fmt.Errorf("invalid region format '%s'", region)
	}

	switch cspKey {
	case csp.AWS:
		if regionClean != "" {
			return fmt.Sprintf("s3.%s.amazonaws.com", regionClean), nil
		}
		return "s3.amazonaws.com", nil
	case csp.GCP:
		return "storage.googleapis.com", nil
	case csp.Alibaba:
		if regionClean != "" {
			return fmt.Sprintf("oss-%s.aliyuncs.com", regionClean), nil
		}
		return "oss-ap-northeast-1.aliyuncs.com", nil
	case csp.Tencent:
		if regionClean != "" {
			return fmt.Sprintf("cos.%s.myqcloud.com", regionClean), nil
		}
		return "cos.ap-guangzhou.myqcloud.com", nil
	case csp.IBM:
		if regionClean != "" {
			return fmt.Sprintf("s3.%s.cloud-object-storage.appdomain.cloud", regionClean), nil
		}
		return "s3.us.cloud-object-storage.appdomain.cloud", nil
	case csp.NCP:
		if regionClean != "" {
			return fmt.Sprintf("%s.object.ncloudstorage.com", regionClean), nil
		}
		return "kr.object.ncloudstorage.com", nil
	default:
		return "", fmt.Errorf("unsupported CSP '%s' for S3 test helper; supported CSPs: aws, gcp, alibaba, tencent, ibm, ncp", cspName)
	}
}

// getAwsBucketRegion fetches the AWS S3 region of a bucket via the X-Amz-Bucket-Region response header.
func getAwsBucketRegion(ctx context.Context, bucketName string) string {
	url := fmt.Sprintf("https://%s.s3.amazonaws.com", bucketName)
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, url, nil)
	if err != nil {
		return ""
	}
	httpClient := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpClient.Do(req)
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
// Test-Only Client Helper API Models (Non-Production)
// ============================================================================

// ScanObjectStorageRequest represents request parameters for test-only object storage bucket scanning.
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

// ScannedBucketSummary represents summary info of a discovered bucket.
type ScannedBucketSummary struct {
	BucketName        string `json:"bucketName"`
	Region            string `json:"region,omitempty"`
	CreationTime      string `json:"creationTime,omitempty"`
	SizeBytes         int64  `json:"sizeBytes,omitempty"`
	ObjectCount       int64  `json:"objectCount,omitempty"`
	VersioningEnabled bool   `json:"versioningEnabled"`
	EncryptionType    string `json:"encryptionType,omitempty"`
}

// InspectObjectStorageRequest represents request parameters for test-only deep object storage metadata collection.
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
// Test-Only Client Helper Handlers (Non-Production)
// ============================================================================

// GetSecurityPublicKey godoc
// @ID GetSecurityPublicKey
// @Summary [TEST ONLY] Get RSA public key for credential encryption
// @Description [WARNING - NON-PRODUCTION TEST HELPER] This endpoint is provided strictly for local testing, evaluation, and Beetle UX Lab prototyping.
// @Description Do NOT rely on this helper flow in production systems; production clients must discover source buckets client-side and encrypt data migration credentials via `/migration/data/encryptionKey` and the `transx/fieldsec` package.
// @Description
// @Description Generates and returns a one-time RSA public key bundle for encrypting sensitive CSP access credentials before calling the test-only `/scan` or `/inspect` helper endpoints.
// @Tags [Test] Client Helper (Non-Production)
// @Accept json
// @Produce json
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[transx.PublicKeyBundle] "Public key bundle containing keyId, algorithm, publicKey (PEM), and expiresAt"
// @Failure 500 {object} model.ApiResponse[any] "Key generation failed"
// @Router /migration/security/publicKey [get]
func GetSecurityPublicKey(c echo.Context) error {
	log.Warn().Msg("[Test Client Helper] Non-production endpoint GET /migration/security/publicKey invoked")

	keyPair, err := transx.GetKeyStore().GenerateKeyPair(transx.GetKeyExpiry())
	if err != nil {
		log.Error().Err(err).Msg("[Test Client Helper] Failed to generate encryption key")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Key generation failed"))
	}

	bundle, err := keyPair.ExportPublicBundle()
	if err != nil {
		log.Error().Err(err).Msg("[Test Client Helper] Failed to export public key bundle")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Key export failed"))
	}

	return c.JSON(http.StatusOK, model.SuccessResponse(bundle))
}

// ScanObjectStorage godoc
// @ID ScanObjectStorage
// @Summary [TEST ONLY] Scan cloud account object storage buckets
// @Description [WARNING - NON-PRODUCTION TEST HELPER] This endpoint is provided strictly for local testing, evaluation, and Beetle UX Lab prototyping.
// @Description Do NOT send cloud credentials to this endpoint in production workflows; production clients must scan S3 buckets directly on the client side.
// @Description
// @Description Scans and lists object storage buckets in the specified cloud provider account and region via the MinIO S3 SDK (supported CSPs: `aws`, `gcp`, `alibaba`, `tencent`, `ibm`, `ncp`).
// @Tags [Test] Client Helper (Non-Production)
// @Accept json
// @Produce json
// @Param request body ScanObjectStorageRequest true "CSP credentials and target cloud information for bucket scanning"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[[]ScannedBucketSummary] "List of discovered object storage buckets"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters, unsupported CSP, or missing credentials"
// @Failure 500 {object} model.ApiResponse[any] "Failed to list buckets from S3 API"
// @Router /migration/middleware/objectStorage/scan [post]
func ScanObjectStorage(c echo.Context) error {
	log.Warn().Msg("[Test Client Helper] Non-production endpoint POST /migration/middleware/objectStorage/scan invoked")

	var req ScanObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("[Test Client Helper] Failed to bind scan request")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if strings.TrimSpace(req.Csp) == "" || strings.TrimSpace(req.Region) == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP and region are required"))
	}

	globalEndpoint, err := getS3EndpointForCsp(req.Csp, "")
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}
	regionEndpoint, err := getS3EndpointForCsp(req.Csp, req.Region)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// Decrypt sensitive credential fields via fieldsec when keyId is provided.
	if req.KeyId != "" {
		sensitiveFields := []string{
			"accessKeyId",
			"secretAccessKey",
			"tenantId",
			"subscriptionId",
			"s3AccessKey",
			"s3SecretKey",
		}
		if decrypted, decErr := fieldsec.DecryptWithStore(req, sensitiveFields, transx.GetKeyStore(), req.KeyId, "keyId"); decErr == nil {
			req = decrypted.(ScanObjectStorageRequest)
			log.Info().Str("keyId", req.KeyId).Msg("[Test Client Helper] Successfully decrypted credential fields via fieldsec")
		} else {
			log.Warn().Err(decErr).Str("keyId", req.KeyId).Msg("[Test Client Helper] Failed to decrypt credential fields via fieldsec")
		}
	}

	log.Info().
		Str("csp", req.Csp).
		Str("region", req.Region).
		Msg("[Test Client Helper] Scanning object storage buckets for CSP account")

	// Resolve S3 access credentials with S3AccessKey/S3SecretKey taking precedence.
	accessKey := req.S3AccessKey
	if accessKey == "" {
		accessKey = req.AccessKeyId
	}
	secretKey := req.S3SecretKey
	if secretKey == "" {
		secretKey = req.SecretAccessKey
	}

	if accessKey == "" || secretKey == "" {
		log.Warn().Str("csp", req.Csp).Msg("[Test Client Helper] Access Key and Secret Key are required for bucket scanning")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Access Key and Secret Key are required to scan cloud buckets"))
	}

	// Build global S3 client for account-wide bucket listing and location detection.
	globalClient, err := minio.New(globalEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
		Region: "us-east-1",
	})
	if err != nil {
		log.Error().Err(err).Str("csp", req.Csp).Msg("[Test Client Helper] Failed to create MinIO S3 global client")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to initialize S3 client: %v", err)))
	}

	ctx := c.Request().Context()
	realBuckets, err := globalClient.ListBuckets(ctx)
	if err != nil {
		log.Error().Err(err).Str("csp", req.Csp).Str("endpoint", globalEndpoint).Msg("[Test Client Helper] MinIO S3 ListBuckets failed")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse(fmt.Sprintf("Failed to list buckets from %s (%s): %v", req.Csp, globalEndpoint, err)))
	}

	// Build region-specific client for inspecting buckets in the target region.
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

	// Limit concurrent bucket inspections to 10 workers to avoid CSP rate throttling.
	maxConcurrency := 10
	sem := make(chan struct{}, maxConcurrency)

	for _, b := range realBuckets {
		wg.Add(1)
		go func(bucket minio.BucketInfo) {
			defer wg.Done()

			sem <- struct{}{}
			defer func() { <-sem }()

			// Resolve bucket region via X-Amz-Bucket-Region for AWS or GetBucketLocation otherwise.
			bucketRegion := ""
			if strings.ToLower(strings.TrimSpace(req.Csp)) == csp.AWS {
				bucketRegion = getAwsBucketRegion(ctx, bucket.Name)
			}
			if bucketRegion == "" {
				loc, locErr := globalClient.GetBucketLocation(ctx, bucket.Name)
				if locErr == nil {
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

			// Skip buckets located outside the requested target region.
			if targetRegion != "" && bucketRegion != targetRegion {
				log.Debug().
					Str("bucket", bucket.Name).
					Str("bucketRegion", bucketRegion).
					Str("targetRegion", targetRegion).
					Msg("[Test Client Helper] Skipping bucket in different region")
				return
			}

			creationTime := ""
			if !bucket.CreationDate.IsZero() {
				creationTime = bucket.CreationDate.Format("2006-01-02T15:04:05Z")
			}

			versioningEnabled := false
			if vConfig, vErr := regionClient.GetBucketVersioning(ctx, bucket.Name); vErr == nil {
				versioningEnabled = vConfig.Enabled()
			}

			encryptionType := "None"
			if encConfig, encErr := regionClient.GetBucketEncryption(ctx, bucket.Name); encErr == nil && len(encConfig.Rules) > 0 {
				encryptionType = "AES256"
			}

			// Sample up to 1000 objects to estimate object count and size quickly.
			var sizeBytes int64
			var objectCount int64
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

	// Sort scanned buckets alphabetically by BucketName for deterministic output.
	sort.Slice(buckets, func(i, j int) bool {
		return strings.ToLower(buckets[i].BucketName) < strings.ToLower(buckets[j].BucketName)
	})

	log.Info().Int("count", len(buckets)).Str("csp", req.Csp).Str("targetRegion", req.Region).Msg("[Test Client Helper] Successfully scanned and filtered buckets for target region")
	return c.JSON(http.StatusOK, model.SuccessResponse(buckets))
}

// InspectObjectStorage godoc
// @ID InspectObjectStorage
// @Summary [TEST ONLY] Collect deep metadata from selected source object storage buckets
// @Description [WARNING - NON-PRODUCTION TEST HELPER] This endpoint is provided strictly for local testing, evaluation, and Beetle UX Lab prototyping.
// @Description Do NOT send cloud credentials to this endpoint in production workflows; production clients must inspect source S3 buckets directly on the client side and submit the resulting `SourceObjectStorage` model to `POST /recommendation/middleware/objectStorage`.
// @Description
// @Description Deeply inspects and extracts feature and usage metadata (totalSizeBytes, objectCount, versioning, encryption, CORS rules, policy, tags, creationDate) from selected buckets (supported CSPs: `aws`, `gcp`, `alibaba`, `tencent`, `ibm`, `ncp`).
// @Description
// @Description [Note] Extracted fields conform to Beetle's Recommendation API input specification (`SourceObjectStorage`):
// @Description - Versioning: Extracted via `GetBucketVersioning`. If disabled or unavailable, `versioningEnabled` is false.
// @Description - Encryption: Extracted via `GetBucketEncryption`. If rules are absent or unavailable, `encryptionEnabled` is false.
// @Description - CORS: Extracted via `GetBucketCors`. Active rules are mapped into `corsRule`; otherwise `corsEnabled` is false and `corsRule` is nil.
// @Description - Public Access: Extracted via `GetBucketPolicy`. If a wildcard public principal statement is detected, `isPublic` is true.
// @Description - Tags: Extracted via `GetBucketTagging`.
// @Description - CreationDate: Extracted via bucket listing creation timestamp formatted in RFC 3339.
// @Description - AccessFrequency: Defaults to `"frequent"` (Standard storage tier baseline for recommendation).
// @Tags [Test] Client Helper (Non-Production)
// @Accept json
// @Produce json
// @Param request body InspectObjectStorageRequest true "Parameters and selected bucket names for deep object storage inspection"
// @Param X-Request-Id header string false "Unique request ID"
// @Success 200 {object} model.ApiResponse[storagemodel.SourceObjectStorage] "Successfully collected source object storage model for selected buckets"
// @Failure 400 {object} model.ApiResponse[any] "Invalid request parameters, unsupported CSP, or missing credentials"
// @Failure 500 {object} model.ApiResponse[any] "Failed to inspect object storage buckets"
// @Router /migration/middleware/objectStorage/inspect [post]
func InspectObjectStorage(c echo.Context) error {
	log.Warn().Msg("[Test Client Helper] Non-production endpoint POST /migration/middleware/objectStorage/inspect invoked")

	var req InspectObjectStorageRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Invalid request format"))
	}

	if strings.TrimSpace(req.Csp) == "" {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("CSP is required"))
	}
	if len(req.SelectedBucketNames) == 0 {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("At least one bucket name in selectedBucketNames is required"))
	}

	endpoint, err := getS3EndpointForCsp(req.Csp, req.Region)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse(err.Error()))
	}

	// Decrypt sensitive credential fields via fieldsec when keyId is provided.
	if req.KeyId != "" {
		sensitiveFields := []string{
			"accessKeyId",
			"secretAccessKey",
			"tenantId",
			"subscriptionId",
			"s3AccessKey",
			"s3SecretKey",
		}
		if decrypted, decErr := fieldsec.DecryptWithStore(req, sensitiveFields, transx.GetKeyStore(), req.KeyId, "keyId"); decErr == nil {
			req = decrypted.(InspectObjectStorageRequest)
			log.Info().Str("keyId", req.KeyId).Msg("[Test Client Helper] Successfully decrypted credential fields for bucket inspection")
		} else {
			log.Warn().Err(decErr).Str("keyId", req.KeyId).Msg("[Test Client Helper] Failed to decrypt credential fields for bucket inspection")
		}
	}

	// Resolve S3 access credentials with S3AccessKey/S3SecretKey taking precedence.
	accessKey := req.S3AccessKey
	if accessKey == "" {
		accessKey = req.AccessKeyId
	}
	secretKey := req.S3SecretKey
	if secretKey == "" {
		secretKey = req.SecretAccessKey
	}

	if accessKey == "" || secretKey == "" {
		log.Warn().Str("csp", req.Csp).Msg("[Test Client Helper] Access Key and Secret Key are required for bucket inspection")
		return c.JSON(http.StatusBadRequest, model.SimpleErrorResponse("Access Key and Secret Key are required to inspect object storage metadata"))
	}

	reqCtx := c.Request().Context()
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Error().Err(err).Msg("[Test Client Helper] Failed to initialize MinIO client for bucket inspection")
		return c.JSON(http.StatusInternalServerError, model.SimpleErrorResponse("Failed to initialize object storage client: "+err.Error()))
	}

	// Pre-fetch account bucket list to obtain CreationDate values.
	bucketCreationDates := make(map[string]string)
	if bucketsList, listErr := minioClient.ListBuckets(reqCtx); listErr == nil {
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

			sem <- struct{}{}
			defer func() { <-sem }()

			versioningEnabled := false
			if vConfig, vErr := minioClient.GetBucketVersioning(reqCtx, bucketName); vErr == nil && vConfig.Enabled() {
				versioningEnabled = true
			}

			encryptionEnabled := false
			if encConfig, encErr := minioClient.GetBucketEncryption(reqCtx, bucketName); encErr == nil && encConfig != nil && len(encConfig.Rules) > 0 {
				encryptionEnabled = true
			}

			isPublic := false
			if policyStr, pErr := minioClient.GetBucketPolicy(reqCtx, bucketName); pErr == nil {
				if strings.Contains(policyStr, `"Principal":"*"`) || strings.Contains(policyStr, `"Principal": "*"`) || strings.Contains(policyStr, `"Principal":{"AWS":"*"}`) {
					isPublic = true
				}
			}

			var tagsMap map[string]string
			if tagging, tErr := minioClient.GetBucketTagging(reqCtx, bucketName); tErr == nil && tagging != nil {
				tagsMap = tagging.ToMap()
			}

			// Extract active CORS rules when configured on the bucket.
			corsEnabled := false
			var corsRules []storagemodel.CORSRule
			if corsConfig, cErr := minioClient.GetBucketCors(reqCtx, bucketName); cErr == nil && corsConfig != nil && len(corsConfig.CORSRules) > 0 {
				corsEnabled = true
				corsRules = make([]storagemodel.CORSRule, 0, len(corsConfig.CORSRules))
				for _, r := range corsConfig.CORSRules {
					corsRules = append(corsRules, storagemodel.CORSRule{
						AllowedOrigin: r.AllowedOrigin,
						AllowedMethod: r.AllowedMethod,
						AllowedHeader: r.AllowedHeader,
						ExposeHeader:  r.ExposeHeader,
						MaxAgeSeconds: r.MaxAgeSeconds,
					})
				}
			}

			creationDate := bucketCreationDates[bucketName]

			var totalSizeBytes int64
			var objectCount int64
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

	// Sort inspected bucket results alphabetically by BucketName for deterministic output.
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

	log.Info().Int("count", len(results)).Str("csp", req.Csp).Msg("[Test Client Helper] Successfully inspected bucket metadata into SourceObjectStorage model")
	return c.JSON(http.StatusOK, model.SuccessResponse(sourceModel))
}
