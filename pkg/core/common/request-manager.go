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

// Package common provides common utilities for the CM-Beetle system.
// This file (request-manager.go) manages HTTP requests:
// - API Request Tracking: Tracking API requests coming into Beetle (RequestMap, RequestDetails)
// - Internal HTTP Client: Executing API calls to Tumblebug (ExecuteHttpRequest, caching, circuit breaker)
package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/lkvstore"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

// ============================================================================
// API Request Tracking
// ----------------------------------------------------------------------------
// Tracks API requests coming into Beetle.
// Features: Request Tracking, Progress Updates, Status Management, Persistence
// ============================================================================

// RequestInfo stores the essential details of an HTTP request.
type RequestInfo struct {
	Method string            `json:"method"`         // HTTP method (GET, POST, etc.), indicating the request's action type.
	URL    string            `json:"url"`            // The URL the request is made to.
	Header map[string]string `json:"header"`         // Key-value pairs of the request headers.
	Body   any               `json:"body,omitempty"` // Optional: request body
}

// RequestDetails contains detailed information about an HTTP request and its processing status.
type RequestDetails struct {
	StartTime     time.Time   `json:"startTime"`     // The time when the request was received by the server.
	EndTime       time.Time   `json:"endTime"`       // The time when the request was fully processed.
	Status        string      `json:"status"`        // The current status of the request (e.g., "Handling", "Error", "Success").
	RequestInfo   RequestInfo `json:"requestInfo"`   // Extracted information about the request.
	ResponseData  any         `json:"responseData"`  // The data sent back in response to the request.
	ErrorResponse string      `json:"errorResponse"` // A message describing any error that occurred during request processing.
}

// ProgressInfo contains the progress information of a request.
type ProgressInfo struct {
	Title string    `json:"title"`
	Info  any       `json:"info"`
	Time  time.Time `json:"time"`
}

// RequestStatus constants define the possible states of a request.
// Status Flow: Handling → Success/Error
// - Handling: Request is currently being processed
// - Success: Request completed successfully
// - Error: Request failed with an error
// Note: These values are aligned with CB-Tumblebug's status values.
const (
	// RequestStatusHandling indicates the request is currently being processed.
	RequestStatusHandling = "Handling"
	// RequestStatusSuccess indicates the request completed successfully.
	RequestStatusSuccess = "Success"
	// RequestStatusError indicates the request failed with an error.
	RequestStatusError = "Error"

	// Deprecated: Legacy status constants (kept for reference)
	// RequestStatusReceived   = "received"
	// RequestStatusInProgress = "in-progress"
	// RequestStatusCompleted  = "completed"
	// RequestStatusFailed     = "failed"
)

// requestKeyPrefix is the key prefix for request details in lkvstore
const requestKeyPrefix = "/beetle/request/"

// RequestFilter defines filter criteria for querying requests
type RequestFilter struct {
	Status string    // Filter by status (Handling, Success, Error)
	Method string    // Filter by HTTP method
	URL    string    // Filter by URL (partial match)
	Since  time.Time // Filter by start time (requests after this time)
}

// SetRequest stores request details with the given request ID.
// It uses lkvstore for persistence across server restarts.
func SetRequest(reqID string, details RequestDetails) error {
	return lkvstore.Put(requestKeyPrefix+reqID, details)
}

// GetRequest retrieves request details by request ID.
// Returns the details and a boolean indicating if the request was found.
func GetRequest(reqID string) (RequestDetails, bool) {
	value, ok := lkvstore.Get(requestKeyPrefix + reqID)
	if !ok {
		return RequestDetails{}, false
	}

	return convertToRequestDetails(value)
}

// HasRequest checks if a request with the given ID exists.
func HasRequest(reqID string) bool {
	_, ok := lkvstore.Get(requestKeyPrefix + reqID)
	return ok
}

// RemoveRequest deletes request details by request ID.
func RemoveRequest(reqID string) {
	lkvstore.Delete(requestKeyPrefix + reqID)
}

// RemoveAllRequests deletes all request tracking records.
func RemoveAllRequests() {
	requests, ok := lkvstore.GetKvWithPrefix(requestKeyPrefix)
	if !ok {
		return
	}

	for _, kv := range requests {
		lkvstore.Delete(kv.Key)
	}
}

// DefaultRequestRetentionPeriod is the default retention period for completed/failed requests.
const DefaultRequestRetentionPeriod = 7 * 24 * time.Hour // 1 week

// CleanupOldRequests removes completed or failed requests older than the specified duration.
// Requests with "Handling" status are not removed regardless of age.
// Returns the number of removed requests.
func CleanupOldRequests(maxAge time.Duration) int {
	kvList, ok := lkvstore.GetKvWithPrefix(requestKeyPrefix)
	if !ok {
		return 0
	}

	removedCount := 0
	now := time.Now()

	for _, kv := range kvList {
		details, ok := convertToRequestDetails(kv.Value)
		if !ok {
			continue
		}

		// Skip requests that are still being processed
		if strings.EqualFold(details.Status, RequestStatusHandling) {
			continue
		}

		// Check if the request is older than maxAge
		// Use EndTime if available, otherwise use StartTime
		referenceTime := details.EndTime
		if referenceTime.IsZero() {
			referenceTime = details.StartTime
		}

		if now.Sub(referenceTime) > maxAge {
			lkvstore.Delete(kv.Key)
			removedCount++
			log.Debug().
				Str("key", kv.Key).
				Str("status", details.Status).
				Time("endTime", details.EndTime).
				Msg("Removed old request record")
		}
	}

	if removedCount > 0 {
		log.Info().
			Int("removedCount", removedCount).
			Dur("maxAge", maxAge).
			Msg("Cleaned up old request records")
	}

	return removedCount
}

// StartRequestCleanupScheduler starts a background goroutine that periodically
// cleans up old request records. It runs cleanup at the specified interval.
func StartRequestCleanupScheduler(interval time.Duration, maxAge time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		log.Info().
			Dur("interval", interval).
			Dur("maxAge", maxAge).
			Msg("Request cleanup scheduler started")

		// Run initial cleanup
		CleanupOldRequests(maxAge)

		for range ticker.C {
			CleanupOldRequests(maxAge)
		}
	}()
}

// GetAllRequests retrieves all request details with optional filtering.
// If filter is nil, all requests are returned.
func GetAllRequests(filter *RequestFilter) []RequestDetails {
	var results []RequestDetails

	kvList, ok := lkvstore.GetKvWithPrefix(requestKeyPrefix)
	if !ok {
		return results
	}

	for _, kv := range kvList {
		details, ok := convertToRequestDetails(kv.Value)
		if !ok {
			continue
		}

		// Apply filters if provided
		if filter != nil {
			if filter.Status != "" && !strings.EqualFold(details.Status, filter.Status) {
				continue
			}
			if filter.Method != "" && !strings.EqualFold(details.RequestInfo.Method, filter.Method) {
				continue
			}
			if filter.URL != "" && !strings.Contains(strings.ToLower(details.RequestInfo.URL), strings.ToLower(filter.URL)) {
				continue
			}
			if !filter.Since.IsZero() && details.StartTime.Before(filter.Since) {
				continue
			}
		}

		results = append(results, details)
	}

	return results
}

// convertToRequestDetails converts lkvstore value to RequestDetails.
// Handles both map[string]any (from JSON) and RequestDetails types.
func convertToRequestDetails(value any) (RequestDetails, bool) {
	switch v := value.(type) {
	case RequestDetails:
		return v, true
	case map[string]any:
		// Convert map to RequestDetails via JSON
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			log.Error().Err(err).Msg("Failed to marshal value to JSON")
			return RequestDetails{}, false
		}

		var details RequestDetails
		if err := json.Unmarshal(jsonBytes, &details); err != nil {
			log.Error().Err(err).Msg("Failed to unmarshal JSON to RequestDetails")
			return RequestDetails{}, false
		}
		return details, true
	default:
		log.Error().Msgf("Unexpected value type: %T", value)
		return RequestDetails{}, false
	}
}

// ExtractRequestInfo extracts necessary information from http.Request
func ExtractRequestInfo(r *http.Request) RequestInfo {
	headerInfo := make(map[string]string)

	// Headers to ignore
	ignorePatterns := []string{
		"Accept", "Authorization", "Connection", "Referer",
		"Sec-", "User-Agent",
	}

	for name, headers := range r.Header {
		shouldIgnore := false
		for _, pattern := range ignorePatterns {
			if strings.HasPrefix(name, pattern) {
				shouldIgnore = true
				break
			}
		}

		if !shouldIgnore {
			headerInfo[name] = headers[0]
		}
	}

	var bodyObject any
	if r.Body != nil { // Check if the body is not nil
		bodyBytes, err := io.ReadAll(r.Body)
		if err == nil {
			json.Unmarshal(bodyBytes, &bodyObject) // Try to unmarshal to a JSON object

			// Important: Write the body back for further processing
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	return RequestInfo{
		Method: r.Method,
		URL:    r.URL.String(),
		Header: headerInfo,
		Body:   bodyObject, // Use the JSON object if parsing was successful, otherwise it's nil
	}
}

// UpdateRequestProgress updates the handling status of the request.
func UpdateRequestProgress(reqID string, progressData any) {
	details, ok := GetRequest(reqID)
	if !ok {
		log.Warn().Str("reqID", reqID).Msg("Request ID not found")
		return
	}

	var responseData []any
	if details.ResponseData != nil {
		// Safe type assertion for existing ResponseData
		if existing, ok := details.ResponseData.([]any); ok {
			responseData = existing
		} else {
			// If ResponseData is not []any, wrap existing data in a slice
			responseData = []any{details.ResponseData}
		}
	}

	// Append the new progressData to the existing ResponseData
	responseData = append(responseData, progressData)
	details.ResponseData = responseData

	if err := SetRequest(reqID, details); err != nil {
		log.Error().Err(err).Str("reqID", reqID).Msg("Failed to update request progress")
	}
}

// ============================================================================
// Internal HTTP Client
// ----------------------------------------------------------------------------
// Handles internal API calls from Beetle to downstream services (e.g., Tumblebug).
// Features: Caching, Circuit Breaker, Concurrency Control
// ============================================================================

// CacheItem is a struct to store cached item
type CacheItem[T any] struct {
	Response  T
	ExpiresAt time.Time
}

// CircuitBreakerState represents the state of a circuit breaker for a specific request
type CircuitBreakerState struct {
	FailureCount int
	LastFailure  time.Time
	IsOpen       bool
}

// clientCache is a map for cache items of internal calls
var clientCache = sync.Map{}

// clientRequestCounter is a map for request counters of internal calls
var clientRequestCounter = sync.Map{}

// clientCircuitBreakers tracks circuit breaker states for different request keys
var clientCircuitBreakers = sync.Map{}

const (
	// Circuit breaker thresholds
	circuitBreakerFailureThreshold = 5                // Number of failures before opening circuit
	circuitBreakerOpenDuration     = 30 * time.Second // How long to keep circuit open
)

const (
	// VeryShortDuration is a duration for very short-term cache
	VeryShortDuration = 1 * time.Second
	// ShortDuration is a duration for short-term cache
	ShortDuration = 2 * time.Second
	// MediumDuration is a duration for medium-term cache
	MediumDuration = 5 * time.Second
	// LongDuration is a duration for long-term cache
	LongDuration = 10 * time.Second
)

// NoBody is a constant for empty body
const NoBody = "NOBODY"

// MaxDebugBodyLength is the maximum length of response body for debug logs
const MaxDebugBodyLength = 50000

// SetUseBody returns false if the given body is NoBody
func SetUseBody(requestBody any) bool {
	if str, ok := requestBody.(string); ok {
		return str != NoBody
	}
	return true
}

// ExecuteHttpRequest performs the HTTP request and fills the result (var requestBody interface{} = nil for empty body)
func ExecuteHttpRequest[B any, T any](
	client *resty.Client,
	method string,
	url string,
	headers map[string]string,
	useBody bool, // New parameter to specify if body should be used
	body *B,
	result *T, // Generic type
	cacheDuration time.Duration,
) error {

	// Perform the HTTP request using Resty
	setRestyDebug := false // Disable Resty debug, use custom logging instead

	// Record request start time
	requestStartTime := time.Now()

	// Log the request in zerologger style (use trace for GET, debug for others)
	if method == "GET" {
		requestLogEvent := log.Trace().
			Str("Method", method).
			Str("URI", url)
		if useBody && body != nil {
			if bodyBytes, err := json.Marshal(body); err == nil {
				requestLogEvent = requestLogEvent.RawJSON("requestBody", bodyBytes)
			}
		}
		requestLogEvent.Msg("Internal Call Start")
	} else {
		requestLogEvent := log.Debug().
			Str("Method", method).
			Str("URI", url)
		if useBody && body != nil {
			if bodyBytes, err := json.Marshal(body); err == nil {
				requestLogEvent = requestLogEvent.RawJSON("requestBody", bodyBytes)
			}
		}
		requestLogEvent.Msg("Internal Call Start")
	}

	// Generate cache key for GET method only
	requestKey := ""
	if method == "GET" {

		if useBody {
			// Serialize the body to JSON
			bodyString, err := json.Marshal(body)
			if err != nil {
				return fmt.Errorf("JSON marshaling failed: %w", err)
			}
			// Create cache key using both URL and body
			requestKey = fmt.Sprintf("%s_%s_%s", method, url, string(bodyString))
		} else {
			// Create cache key using only URL
			requestKey = fmt.Sprintf("%s_%s", method, url)
		}

		if item, found := clientCache.Load(requestKey); found {
			// Ensure safe type assertion
			cachedItem, ok := item.(CacheItem[T])
			if !ok {
				log.Error().Msgf("Type assertion failed for cache item: expected CacheItem[%T], got %T", *result, item)
				clientCache.Delete(requestKey) // Delete invalid cache item
				return fmt.Errorf("type assertion failed for cache item")
			}

			if time.Now().Before(cachedItem.ExpiresAt) {
				log.Trace().Msgf("Cache hit! Expires: %v", time.Since(cachedItem.ExpiresAt))
				*result = cachedItem.Response

				return nil
			} else {
				clientCache.Delete(requestKey)
			}
		}

		// Check circuit breaker before making actual requests
		if checkCircuitBreaker(requestKey) {
			return fmt.Errorf("API call temporarily blocked due to circuit breaker protection (repeated failures detected), please try again later (API: %s)", requestKey)
		}

		// Limit the number of concurrent requests
		concurrencyLimit := 20       // Increased from 10 to 20 for better throughput
		retryWait := 2 * time.Second // Reduced from 5 to 2 seconds for faster retries
		retryLimit := 8              // Increased from 3 to 8 for more resilience
		retryCount := 0
		// try to wait for the upcoming cached result when sending queue is full
		for {
			if !limitConcurrentRequests(requestKey, concurrencyLimit) {
				if retryCount >= retryLimit {
					log.Debug().Msgf("too many same requests after %d retries: %s", retryLimit, requestKey)
					return fmt.Errorf("too many same requests: %s", requestKey)
				}
				time.Sleep(retryWait)

				if item, found := clientCache.Load(requestKey); found {
					cachedItem, ok := item.(CacheItem[T])
					if !ok {
						log.Error().Msgf("Type assertion failed for cache item while waiting: expected CacheItem[%T], got %T", *result, item)
						clientCache.Delete(requestKey) // Delete invalid cache item
						return fmt.Errorf("type assertion failed for cache item while waiting")
					}
					*result = cachedItem.Response
					// release the request count for parallel requests limit
					requestDone(requestKey)
					log.Debug().Msg("Got the cached result while waiting")
					return nil
				}
				retryCount++
			} else {
				break
			}
		}
	}

	// Perform the HTTP request using Resty
	client.SetDebug(setRestyDebug)
	// SetAllowGetMethodPayload should be set to true for GET method to allow payload
	// NOTE: Need to be removed when cb-spider api stops using GET method with payload
	client.SetAllowGetMethodPayload(true)
	req := client.R().SetHeader("Content-Type", "application/json").SetResult(result)

	if headers != nil {
		req = req.SetHeaders(headers)
	}

	if useBody {
		req = req.SetBody(body)
	}

	var resp *resty.Response
	var err error

	// Execute HTTP method based on the given type
	switch method {
	case "GET":
		resp, err = req.Get(url)
	case "POST":
		resp, err = req.Post(url)
	case "PUT":
		resp, err = req.Put(url)
	case "DELETE":
		resp, err = req.Delete(url)
	default:
		return fmt.Errorf("unsupported rest method: %s", method)
	}

	if err != nil {
		// Log error response in zerologger style (use trace for GET, debug for others)
		duration := time.Since(requestStartTime)

		if method == "GET" {
			log.Trace().
				Str("Method", method).
				Str("URI", url).
				Dur("latency", duration).
				Str("error", err.Error()).
				Msg("Internal Call Failed")
		} else {
			log.Debug().
				Str("Method", method).
				Str("URI", url).
				Dur("latency", duration).
				Str("error", err.Error()).
				Msg("Internal Call Failed")
		}

		if method == "GET" {
			requestDone(requestKey)
			// Record circuit breaker failure for GET requests
			recordCircuitBreakerFailure(requestKey)
		}
		cleanedError := cleanErrorMessage(err.Error())
		cleanedURL := cleanURL(url)
		return fmt.Errorf("%s (from %s)", cleanedError, cleanedURL)
	}

	if resp.IsError() {
		// Log HTTP error response in zerologger style (use trace for GET, debug for others)
		duration := time.Since(requestStartTime)

		if method == "GET" {
			errorLogEvent := log.Trace().
				Str("Method", method).
				Str("URI", url).
				Dur("latency", duration).
				Int("status", resp.StatusCode())

			if len(resp.Body()) > 0 {
				errorLogEvent = errorLogEvent.RawJSON("responseBody", resp.Body())
			}
			errorLogEvent.Msg("Internal Call Error")
		} else {
			errorLogEvent := log.Debug().
				Str("Method", method).
				Str("URI", url).
				Dur("latency", duration).
				Int("status", resp.StatusCode())

			if len(resp.Body()) > 0 {
				if shouldTruncateBody(resp.Body()) {
					// Body is too large for debug level, show truncation message
					truncationMsg := createTruncationMessage(len(resp.Body()))
					errorLogEvent = errorLogEvent.RawJSON("responseBody", []byte(truncationMsg))

					// Log full body at trace level
					log.Trace().
						Str("Method", method).
						Str("URI", url).
						Dur("latency", duration).
						Int("status", resp.StatusCode()).
						RawJSON("responseBody", resp.Body()).
						Msg("Internal Call Error (Full Response)")
				} else {
					// Body is small enough for debug level
					errorLogEvent = errorLogEvent.RawJSON("responseBody", resp.Body())
				}
			}
			errorLogEvent.Msg("Internal Call Error")
		}

		if method == "GET" {
			requestDone(requestKey)
			// Record circuit breaker failure for GET requests
			recordCircuitBreakerFailure(requestKey)
		}
		cleanedBody := cleanErrorMessage(string(resp.Body()))
		cleanedURL := cleanURL(url)
		return fmt.Errorf("%s (from %s (%s))", cleanedBody, cleanedURL, resp.Status())
	}

	// Log successful response in zerologger style (use trace for GET, debug for others)
	duration := time.Since(requestStartTime)

	if method == "GET" {
		successLogEvent := log.Trace().
			Str("Method", method).
			Str("URI", url).
			Dur("latency", duration).
			Int("status", resp.StatusCode())

		if len(resp.Body()) > 0 {
			successLogEvent = successLogEvent.RawJSON("responseBody", resp.Body())
		}
		successLogEvent.Msg("Internal Call OK")
	} else {
		successLogEvent := log.Debug().
			Str("Method", method).
			Str("URI", url).
			Dur("latency", duration).
			Int("status", resp.StatusCode())

		if len(resp.Body()) > 0 {
			if shouldTruncateBody(resp.Body()) {
				// Body is too large for debug level, show truncation message
				truncationMsg := createTruncationMessage(len(resp.Body()))
				successLogEvent = successLogEvent.RawJSON("responseBody", []byte(truncationMsg))

				// Log full body at trace level
				log.Trace().
					Str("Method", method).
					Str("URI", url).
					Dur("latency", duration).
					Int("status", resp.StatusCode()).
					RawJSON("responseBody", resp.Body()).
					Msg("Internal Call OK (Full Response)")
			} else {
				// Body is small enough for debug level
				successLogEvent = successLogEvent.RawJSON("responseBody", resp.Body())
			}
		}
		successLogEvent.Msg("Internal Call OK")
	}

	// Update the cache for GET method only
	if method == "GET" {

		// release the request count for parallel requests limit
		requestDone(requestKey)

		// Record circuit breaker success for GET requests
		recordCircuitBreakerSuccess(requestKey)

		// Check if result is nil
		if result == nil {
			log.Trace().Msg("Result is nil, not caching")
		} else {
			clientCache.Store(requestKey, CacheItem[T]{Response: *result, ExpiresAt: time.Now().Add(cacheDuration)})
		}
	}

	return nil
}

// checkCircuitBreaker checks if the circuit breaker is open for a given request key
func checkCircuitBreaker(requestKey string) bool {
	if item, found := clientCircuitBreakers.Load(requestKey); found {
		if breaker, ok := item.(CircuitBreakerState); ok {
			if breaker.IsOpen {
				// Check if enough time has passed to reset the circuit breaker
				if time.Since(breaker.LastFailure) > circuitBreakerOpenDuration {
					// Reset circuit breaker
					breaker.IsOpen = false
					breaker.FailureCount = 0
					clientCircuitBreakers.Store(requestKey, breaker)
					log.Debug().Msgf("API protection reset, service resumed: %s", requestKey)
					return false
				}
				log.Debug().Msgf("API protection is active: %s", requestKey)
				return true
			}
		}
	}
	return false
}

// recordCircuitBreakerFailure records a failure for circuit breaker
func recordCircuitBreakerFailure(requestKey string) {
	var breaker CircuitBreakerState
	if item, found := clientCircuitBreakers.Load(requestKey); found {
		if existing, ok := item.(CircuitBreakerState); ok {
			breaker = existing
		}
	}

	breaker.FailureCount++
	breaker.LastFailure = time.Now()

	if breaker.FailureCount >= circuitBreakerFailureThreshold {
		breaker.IsOpen = true
		log.Warn().Msgf("API protection activated due to consecutive failures: %s (failures: %d, blocked for 30 seconds)", requestKey, breaker.FailureCount)
	}

	clientCircuitBreakers.Store(requestKey, breaker)
}

// recordCircuitBreakerSuccess resets failure count on successful request
func recordCircuitBreakerSuccess(requestKey string) {
	if item, found := clientCircuitBreakers.Load(requestKey); found {
		if breaker, ok := item.(CircuitBreakerState); ok {
			if breaker.FailureCount > 0 {
				breaker.FailureCount = 0
				breaker.IsOpen = false
				clientCircuitBreakers.Store(requestKey, breaker)
				log.Debug().Msgf("API failure counter reset due to successful response: %s", requestKey)
			}
		}
	}
}

// limitConcurrentRequests limits the number of Concurrent requests to the given limit
func limitConcurrentRequests(requestKey string, limit int) bool {
	count, _ := clientRequestCounter.LoadOrStore(requestKey, 0)
	currentCount := count.(int)

	if currentCount >= limit {
		log.Debug().Msgf("[%d] requests for %s", currentCount, requestKey)
		return false
	}

	clientRequestCounter.Store(requestKey, currentCount+1)
	return true
}

// requestDone decreases the request counter
func requestDone(requestKey string) {
	count, _ := clientRequestCounter.Load(requestKey)
	if count == nil {
		return
	}
	currentCount := count.(int)

	if currentCount > 0 {
		clientRequestCounter.Store(requestKey, currentCount-1)
	}
}

// shouldTruncateBody checks if the response body should be truncated for debug logging
func shouldTruncateBody(body []byte) bool {
	return len(body) > MaxDebugBodyLength
}

// createTruncationMessage creates a message indicating body was truncated
func createTruncationMessage(bodyLength int) string {
	// Fast integer approximation: ~7 bytes per word in JSON
	bodyWords := bodyLength / 7
	limitWords := MaxDebugBodyLength / 7

	return fmt.Sprintf(`{"message":"Response body %d words (%d bytes) exceeds limit %d words (%d bytes). Use TRACE level to log and check content."}`,
		bodyWords, bodyLength, limitWords, MaxDebugBodyLength)
}

// cleanErrorMessage removes unwanted characters from error messages
func cleanErrorMessage(message string) string {
	// Remove escaped quotes, newlines, and other escape sequences
	cleaned := strings.ReplaceAll(message, "\\\"", "")
	cleaned = strings.ReplaceAll(cleaned, "\"", "")
	cleaned = strings.ReplaceAll(cleaned, "\\n", " ")
	cleaned = strings.ReplaceAll(cleaned, "\n", " ")
	cleaned = strings.ReplaceAll(cleaned, "\\t", " ")
	cleaned = strings.ReplaceAll(cleaned, "\t", " ")
	cleaned = strings.ReplaceAll(cleaned, "\\r", " ")
	cleaned = strings.ReplaceAll(cleaned, "\r", " ")

	// Remove multiple spaces and trim
	for strings.Contains(cleaned, "  ") {
		cleaned = strings.ReplaceAll(cleaned, "  ", " ")
	}

	cleaned = strings.TrimSpace(cleaned)

	// Remove surrounding curly braces if present
	if len(cleaned) >= 2 && strings.HasPrefix(cleaned, "{") && strings.HasSuffix(cleaned, "}") {
		cleaned = strings.TrimSpace(cleaned[1 : len(cleaned)-1])

		// Remove common JSON field prefixes like "message:", "error:", etc.
		if strings.HasPrefix(cleaned, "message:") {
			cleaned = strings.TrimSpace(cleaned[8:]) // Remove "message:" (8 characters)
		} else if strings.HasPrefix(cleaned, "error:") {
			cleaned = strings.TrimSpace(cleaned[6:]) // Remove "error:" (6 characters)
		} else if strings.HasPrefix(cleaned, "details:") {
			cleaned = strings.TrimSpace(cleaned[8:]) // Remove "details:" (8 characters)
		} else if strings.HasPrefix(cleaned, "info:") {
			cleaned = strings.TrimSpace(cleaned[5:]) // Remove "info:" (5 characters)
		}
	}

	return cleaned
}

// cleanURL removes protocol prefix from URL for cleaner error messages
func cleanURL(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url[7:] // Remove "http://" (7 characters)
	} else if strings.HasPrefix(url, "https://") {
		return url[8:] // Remove "https://" (8 characters)
	}
	return url
}
