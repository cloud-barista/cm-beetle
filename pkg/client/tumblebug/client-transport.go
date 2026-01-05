/*
Copyright 2024 The Cloud-Barista Authors.
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

package tbclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type contextKey string

const cacheDurationKey contextKey = "cacheDuration"

// SetCacheDuration adds cache duration to the context
func SetCacheDuration(ctx context.Context, d time.Duration) context.Context {
	return context.WithValue(ctx, cacheDurationKey, d)
}

// GetCacheDuration retrieves cache duration from the context
func GetCacheDuration(ctx context.Context) (time.Duration, bool) {
	d, ok := ctx.Value(cacheDurationKey).(time.Duration)
	return d, ok
}

// TumblebugTransport is a custom http.RoundTripper that adds caching and circuit breaker logic.
type TumblebugTransport struct {
	Transport     http.RoundTripper
	CacheDuration time.Duration
}

// CacheItem is a struct to store cached item
type CacheItem struct {
	ResponseBytes []byte
	StatusCode    int
	Header        http.Header
	ExpiresAt     time.Time
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

const (
	// MaxDebugBodyLength is the maximum length of response body for debug logs
	MaxDebugBodyLength = 50000
)

// RoundTrip executes a single HTTP transaction, returning a Response for the provided Request.
func (t *TumblebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := t.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	method := req.Method
	url := req.URL.String()

	// Read body for cache key generation (if needed)
	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		// Restore body for the actual request
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// Generate cache key for GET method only
	requestKey := ""
	if method == "GET" {
		if len(bodyBytes) > 0 {
			requestKey = fmt.Sprintf("%s_%s_%s", method, url, string(bodyBytes))
		} else {
			requestKey = fmt.Sprintf("%s_%s", method, url)
		}

		// Check Cache
		if item, found := clientCache.Load(requestKey); found {
			if cachedItem, ok := item.(CacheItem); ok {
				if time.Now().Before(cachedItem.ExpiresAt) {
					log.Trace().Msgf("Cache hit! Expires: %v", time.Since(cachedItem.ExpiresAt))

					// Construct response from cache
					resp := &http.Response{
						StatusCode: cachedItem.StatusCode,
						Header:     cachedItem.Header,
						Body:       io.NopCloser(bytes.NewBuffer(cachedItem.ResponseBytes)),
						Request:    req,
					}
					return resp, nil
				} else {
					clientCache.Delete(requestKey)
				}
			}
		}

		// Check Circuit Breaker
		if checkCircuitBreaker(requestKey) {
			return nil, fmt.Errorf("API call temporarily blocked due to circuit breaker protection (repeated failures detected), please try again later (API: %s)", requestKey)
		}

		// Limit Concurrency
		concurrencyLimit := 20
		retryWait := 2 * time.Second
		retryLimit := 8
		retryCount := 0

		for {
			if !limitConcurrentRequests(requestKey, concurrencyLimit) {
				if retryCount >= retryLimit {
					log.Debug().Msgf("too many same requests after %d retries: %s", retryLimit, requestKey)
					return nil, fmt.Errorf("too many same requests: %s", requestKey)
				}
				time.Sleep(retryWait)

				// Check cache again while waiting
				if item, found := clientCache.Load(requestKey); found {
					if cachedItem, ok := item.(CacheItem); ok {
						requestDone(requestKey) // Release the count we tried to acquire? No, limitConcurrentRequests returns false if not acquired.
						// Wait, limitConcurrentRequests increments if successful. If false, it didn't increment.
						// But if we are waiting, someone else is running.
						// If we find cache, we return it.

						log.Debug().Msg("Got the cached result while waiting")
						resp := &http.Response{
							StatusCode: cachedItem.StatusCode,
							Header:     cachedItem.Header,
							Body:       io.NopCloser(bytes.NewBuffer(cachedItem.ResponseBytes)),
							Request:    req,
						}
						return resp, nil
					}
				}
				retryCount++
			} else {
				break
			}
		}
	}

	// Execute Request
	start := time.Now()
	resp, err := transport.RoundTrip(req)
	duration := time.Since(start)

	if err != nil {
		if method == "GET" {
			requestDone(requestKey)
			recordCircuitBreakerFailure(requestKey)
		}
		log.Error().Err(err).Str("method", method).Str("url", url).Dur("latency", duration).Msg("Internal Call Failed")
		return nil, err
	}

	// Read response body for caching and logging
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewBuffer(respBodyBytes))

	if resp.StatusCode >= 400 {
		if method == "GET" {
			requestDone(requestKey)
			recordCircuitBreakerFailure(requestKey)
		}
		log.Error().Int("status", resp.StatusCode).Str("method", method).Str("url", url).Dur("latency", duration).Msg("Internal Call Error")
		// We still return the response so the caller can handle the error body
		return resp, nil
	}

	// Success
	if method == "GET" {
		requestDone(requestKey)
		recordCircuitBreakerSuccess(requestKey)

		// Determine cache duration
		cacheDuration := t.CacheDuration
		if d, ok := GetCacheDuration(req.Context()); ok {
			cacheDuration = d
		}

		// Cache the response
		clientCache.Store(requestKey, CacheItem{
			ResponseBytes: respBodyBytes,
			StatusCode:    resp.StatusCode,
			Header:        resp.Header,
			ExpiresAt:     time.Now().Add(cacheDuration),
		})
	}

	log.Debug().Int("status", resp.StatusCode).Str("method", method).Str("url", url).Dur("latency", duration).Msg("Internal Call OK")
	return resp, nil
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
