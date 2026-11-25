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

	"github.com/cloud-barista/cm-beetle/pkg/core/common"
	"github.com/rs/zerolog/log"
)

type contextKey string

const cacheDurationKey contextKey = "cacheDuration"

// SetCacheDuration adds cache duration to the context
func SetCacheDuration(ctx context.Context, d time.Duration) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, cacheDurationKey, d)
}

// GetCacheDuration retrieves cache duration from the context
func GetCacheDuration(ctx context.Context) (time.Duration, bool) {
	if ctx == nil {
		return 0, false
	}
	d, ok := ctx.Value(cacheDurationKey).(time.Duration)
	return d, ok
}

// TumblebugTransport is a custom http.RoundTripper that adds caching and circuit breaker logic.
type TumblebugTransport struct {
	Transport     http.RoundTripper
	CacheDuration time.Duration

	mu              sync.Mutex
	cache           map[string]CacheItem
	requestCounter  map[string]int
	circuitBreakers map[string]CircuitBreakerState
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

const maxClientCacheEntries = 1024

func (t *TumblebugTransport) ensureInitLocked() {
	if t.cache == nil {
		t.cache = make(map[string]CacheItem)
	}
	if t.requestCounter == nil {
		t.requestCounter = make(map[string]int)
	}
	if t.circuitBreakers == nil {
		t.circuitBreakers = make(map[string]CircuitBreakerState)
	}
}

const (
	// Circuit breaker thresholds
	circuitBreakerFailureThreshold = 10               // Number of failures before opening circuit
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

	reqCache := requestCacheFromContext(req.Context())
	explicitDuration, hasExplicitDuration := GetCacheDuration(req.Context())
	noCacheRequested := hasExplicitDuration && explicitDuration == 0
	cacheableInRequest := reqCache != nil && !noCacheRequested && isRequestScopedCacheable(req)

	var reqCacheKey requestCacheKey
	if cacheableInRequest {
		reqCacheKey = buildRequestCacheKey(method, url, bodyBytes)
		if cachedResp, hit := reqCache.get(reqCacheKey, req); hit {
			log.Trace().Str("method", method).Str("url", url).Msg("Request cache hit")
			return cachedResp, nil
		}
	} else if reqCache != nil {
		reqCache.recordBypass()
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
		if !noCacheRequested {
			if cachedItem, found := t.getCachedItem(requestKey); found {
				log.Trace().Msgf("Cache hit! Expires: %v", time.Since(cachedItem.ExpiresAt))
				if cacheableInRequest && cachedItem.StatusCode >= http.StatusOK && cachedItem.StatusCode < http.StatusMultipleChoices {
					reqCache.set(reqCacheKey, cachedItem.StatusCode, cachedItem.Header, cachedItem.ResponseBytes)
				}

				// Construct response from cache
				resp := &http.Response{
					StatusCode: cachedItem.StatusCode,
					Header:     cachedItem.Header.Clone(),
					Body:       io.NopCloser(bytes.NewBuffer(cachedItem.ResponseBytes)),
					Request:    req,
				}
				return resp, nil
			}
		}

		// Check Circuit Breaker
		if t.checkCircuitBreaker(requestKey) {
			return nil, fmt.Errorf("API call temporarily blocked due to circuit breaker protection (repeated failures detected), please try again later (API: %s)", requestKey)
		}

		// Limit Concurrency
		concurrencyLimit := 20
		retryWait := 2 * time.Second
		retryLimit := 8
		retryCount := 0

		for {
			if !t.limitConcurrentRequests(requestKey, concurrencyLimit) {
				if retryCount >= retryLimit {
					log.Debug().Msgf("too many same requests after %d retries: %s", retryLimit, requestKey)
					return nil, fmt.Errorf("too many same requests: %s", requestKey)
				}
				if err := common.SleepWithContext(req.Context(), retryWait); err != nil {
					return nil, err
				}

				// Return newly populated cache entry if an in-flight request completed while waiting.
				if !noCacheRequested {
					if cachedItem, found := t.getCachedItem(requestKey); found {
						log.Debug().Msg("Got the cached result while waiting")
						if cacheableInRequest && cachedItem.StatusCode >= http.StatusOK && cachedItem.StatusCode < http.StatusMultipleChoices {
							reqCache.set(reqCacheKey, cachedItem.StatusCode, cachedItem.Header, cachedItem.ResponseBytes)
						}
						resp := &http.Response{
							StatusCode: cachedItem.StatusCode,
							Header:     cachedItem.Header.Clone(),
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

	if cacheableInRequest {
		reqCache.recordMiss()
	}

	finishSubReq := beginSubRequestTrace(req, bodyBytes)

	// Execute Request
	start := time.Now()
	resp, err := transport.RoundTrip(req)
	duration := time.Since(start)

	if err != nil {
		finishSubReq(0, nil, err)
		if method == "GET" {
			t.requestDone(requestKey)
			if req.Context().Err() == nil {
				t.recordCircuitBreakerFailure(requestKey)
			}
		}
		log.Error().Err(err).Str("method", method).Str("url", url).Dur("latency", duration).Msg("Internal Call Failed")
		return nil, err
	}

	// Read response body for caching and logging
	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		finishSubReq(resp.StatusCode, nil, err)
		if method == "GET" {
			t.requestDone(requestKey)
		}
		return nil, err
	}
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewBuffer(respBodyBytes))

	if resp.StatusCode >= 400 {
		finishSubReq(resp.StatusCode, respBodyBytes, nil)
		if method == "GET" {
			t.requestDone(requestKey)
			if resp.StatusCode != http.StatusNotFound {
				t.recordCircuitBreakerFailure(requestKey)
			}
		}
		if resp.StatusCode == http.StatusNotFound {
			log.Debug().Int("status", resp.StatusCode).Str("method", method).Str("url", url).Dur("latency", duration).Msg("Resource Not Found")
		} else {
			log.Error().Int("status", resp.StatusCode).Str("method", method).Str("url", url).Dur("latency", duration).Msg("Internal Call Error")
		}
		// We still return the response so the caller can handle the error body
		return resp, nil
	}

	finishSubReq(resp.StatusCode, nil, nil)

	if cacheableInRequest && resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		reqCache.set(reqCacheKey, resp.StatusCode, resp.Header, respBodyBytes)
	}

	// Success
	if method == "GET" {
		t.requestDone(requestKey)
		t.recordCircuitBreakerSuccess(requestKey)

		if !noCacheRequested {
			cacheDuration := t.CacheDuration
			if hasExplicitDuration {
				cacheDuration = explicitDuration
			}
			if cacheDuration > 0 {
				t.setCachedItem(requestKey, CacheItem{
					ResponseBytes: respBodyBytes,
					StatusCode:    resp.StatusCode,
					Header:        resp.Header,
					ExpiresAt:     time.Now().Add(cacheDuration),
				})
			}
		}
	}

	log.Debug().Int("status", resp.StatusCode).Str("method", method).Str("url", url).Dur("latency", duration).Msg("Internal Call OK")
	return resp, nil
}

func (t *TumblebugTransport) getCachedItem(requestKey string) (CacheItem, bool) {
	now := time.Now()
	t.mu.Lock()
	defer t.mu.Unlock()

	t.ensureInitLocked()
	item, found := t.cache[requestKey]
	if !found {
		return CacheItem{}, false
	}
	if !now.Before(item.ExpiresAt) {
		delete(t.cache, requestKey)
		return CacheItem{}, false
	}
	return item, true
}

func (t *TumblebugTransport) setCachedItem(requestKey string, item CacheItem) {
	now := time.Now()
	if !now.Before(item.ExpiresAt) {
		return
	}
	item.Header = item.Header.Clone()

	t.mu.Lock()
	defer t.mu.Unlock()

	t.ensureInitLocked()
	if len(t.cache) >= maxClientCacheEntries {
		for k, v := range t.cache {
			if !now.Before(v.ExpiresAt) {
				delete(t.cache, k)
			}
		}
	}
	t.cache[requestKey] = item
}

// checkCircuitBreaker checks if the circuit breaker is open for a given request key.
func (t *TumblebugTransport) checkCircuitBreaker(requestKey string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.ensureInitLocked()
	breaker, found := t.circuitBreakers[requestKey]
	if !found || !breaker.IsOpen {
		return false
	}

	if time.Since(breaker.LastFailure) > circuitBreakerOpenDuration {
		delete(t.circuitBreakers, requestKey)
		log.Debug().Msgf("API protection reset, service resumed: %s", requestKey)
		return false
	}

	log.Debug().Msgf("API protection is active: %s", requestKey)
	return true
}

// recordCircuitBreakerFailure records a failure for circuit breaker.
func (t *TumblebugTransport) recordCircuitBreakerFailure(requestKey string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.ensureInitLocked()
	breaker := t.circuitBreakers[requestKey]
	breaker.FailureCount++
	breaker.LastFailure = time.Now()

	if breaker.FailureCount >= circuitBreakerFailureThreshold {
		breaker.IsOpen = true
		log.Warn().Msgf("API protection activated due to consecutive failures: %s (failures: %d, blocked for 30 seconds)", requestKey, breaker.FailureCount)
	}

	t.circuitBreakers[requestKey] = breaker
}

// recordCircuitBreakerSuccess resets failure count on successful request.
func (t *TumblebugTransport) recordCircuitBreakerSuccess(requestKey string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.ensureInitLocked()
	if _, found := t.circuitBreakers[requestKey]; found {
		delete(t.circuitBreakers, requestKey)
		log.Debug().Msgf("API failure counter reset due to successful response: %s", requestKey)
	}
}

// limitConcurrentRequests limits the number of concurrent requests to the given limit.
func (t *TumblebugTransport) limitConcurrentRequests(requestKey string, limit int) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.ensureInitLocked()
	currentCount := t.requestCounter[requestKey]
	if currentCount >= limit {
		log.Debug().Msgf("[%d] requests for %s", currentCount, requestKey)
		return false
	}

	t.requestCounter[requestKey] = currentCount + 1
	return true
}

// requestDone decreases the request counter and removes empty entries.
func (t *TumblebugTransport) requestDone(requestKey string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.requestCounter == nil {
		return
	}
	currentCount := t.requestCounter[requestKey]
	if currentCount <= 1 {
		delete(t.requestCounter, requestKey)
		return
	}
	t.requestCounter[requestKey] = currentCount - 1
}
