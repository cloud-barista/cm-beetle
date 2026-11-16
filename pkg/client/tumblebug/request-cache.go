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
	"crypto/sha256"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"sync"
)

type requestCacheContextKey struct{}

type requestCacheKey struct {
	method   string
	url      string
	bodyHash [32]byte
}

type requestCacheEntry struct {
	statusCode int
	header     http.Header
	body       []byte
}

// RequestCacheStats reports request-scoped cache activity within one request context.
type RequestCacheStats struct {
	Hits     int
	Misses   int
	Bypassed int
}

// RequestCache stores read-only Tumblebug HTTP responses for the lifetime of a single request.
type RequestCache struct {
	mu       sync.RWMutex
	entries  map[requestCacheKey]requestCacheEntry
	hits     int
	misses   int
	bypassed int
}

// WithRequestCache attaches a request-scoped Tumblebug query cache to ctx if not already present.
func WithRequestCache(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if existing := requestCacheFromContext(ctx); existing != nil {
		return ctx
	}
	cache := &RequestCache{
		entries: make(map[requestCacheKey]requestCacheEntry),
	}
	return context.WithValue(ctx, requestCacheContextKey{}, cache)
}

func requestCacheFromContext(ctx context.Context) *RequestCache {
	if ctx == nil {
		return nil
	}
	cache, _ := ctx.Value(requestCacheContextKey{}).(*RequestCache)
	return cache
}

// GetRequestCacheStats returns hit, miss, and bypass counts for the request cache in ctx.
func GetRequestCacheStats(ctx context.Context) (RequestCacheStats, bool) {
	cache := requestCacheFromContext(ctx)
	if cache == nil {
		return RequestCacheStats{}, false
	}
	return cache.Stats(), true
}

// Stats returns a snapshot of the cache counters.
func (c *RequestCache) Stats() RequestCacheStats {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return RequestCacheStats{
		Hits:     c.hits,
		Misses:   c.misses,
		Bypassed: c.bypassed,
	}
}

func (c *RequestCache) recordMiss() {
	c.mu.Lock()
	c.misses++
	c.mu.Unlock()
}

func (c *RequestCache) recordBypass() {
	c.mu.Lock()
	c.bypassed++
	c.mu.Unlock()
}

func (c *RequestCache) get(key requestCacheKey, req *http.Request) (*http.Response, bool) {
	c.mu.RLock()
	entry, ok := c.entries[key]
	c.mu.RUnlock()
	if !ok {
		return nil, false
	}

	c.mu.Lock()
	c.hits++
	c.mu.Unlock()

	return &http.Response{
		StatusCode: entry.statusCode,
		Header:     entry.header.Clone(),
		Body:       io.NopCloser(bytes.NewReader(entry.body)),
		Request:    req,
	}, true
}

func (c *RequestCache) set(key requestCacheKey, statusCode int, header http.Header, body []byte) {
	bodyCopy := make([]byte, len(body))
	copy(bodyCopy, body)

	c.mu.Lock()
	c.entries[key] = requestCacheEntry{
		statusCode: statusCode,
		header:     header.Clone(),
		body:       bodyCopy,
	}
	c.mu.Unlock()
}

func buildRequestCacheKey(method, rawURL string, bodyBytes []byte) requestCacheKey {
	normalizedBody := canonicalizeCacheBody(bodyBytes)
	return requestCacheKey{
		method:   method,
		url:      rawURL,
		bodyHash: sha256.Sum256(normalizedBody),
	}
}

func canonicalizeCacheBody(bodyBytes []byte) []byte {
	if len(bodyBytes) == 0 {
		return nil
	}
	dec := json.NewDecoder(bytes.NewReader(bodyBytes))
	dec.UseNumber()
	var parsed any
	if err := dec.Decode(&parsed); err != nil {
		return bodyBytes
	}
	// Reject trailing non-whitespace tokens so non-JSON payloads hash verbatim.
	if dec.More() {
		return bodyBytes
	}
	canonical, err := json.Marshal(parsed)
	if err != nil {
		return bodyBytes
	}
	return canonical
}

// isRequestScopedCacheable permits only stateless catalog GETs and read-only query POSTs.
func isRequestScopedCacheable(req *http.Request) bool {
	if req == nil || req.URL == nil {
		return false
	}
	trimmed := "/" + strings.TrimLeft(req.URL.Path, "/")
	normalized := "/" + strings.TrimLeft(strings.TrimPrefix(trimmed, "/tumblebug"), "/")
	parts := strings.Split(strings.Trim(normalized, "/"), "/")

	switch req.Method {
	case http.MethodGet, http.MethodHead:
		return isCacheableReadGetPath(normalized, parts)
	case http.MethodPost:
		return isCacheableQueryPostPath(normalized, parts)
	default:
		return false
	}
}

func isCacheableReadGetPath(normalized string, parts []string) bool {
	switch normalized {
	case "/k8sClusterInfo",
		"/availableK8sVersion",
		"/checkK8sNodeGroupsOnK8sCreation",
		"/rdbms/support",
		"/rdbms/capability",
		"/objectStorage/support":
		return true
	}
	if strings.HasPrefix(normalized, "/provider/") || strings.HasPrefix(normalized, "/connConfig/") {
		return true
	}
	// Allow immutable catalog lookups (/ns/{nsId}/resources/spec/* and /image/*) while excluding mutable state.
	if len(parts) >= 5 && parts[0] == "ns" && parts[2] == "resources" {
		return parts[3] == "spec" || parts[3] == "image"
	}
	return false
}

func isCacheableQueryPostPath(normalized string, parts []string) bool {
	switch normalized {
	case "/recommendSpec",
		"/k8sClusterRecommendNode",
		"/specImagePairReview":
		return true
	}
	// Match /ns/{nsId}/resources/searchImage and /ns/{nsId}/resources/rdbms/validate.
	if len(parts) == 4 && parts[0] == "ns" && parts[2] == "resources" && parts[3] == "searchImage" {
		return true
	}
	if len(parts) == 5 && parts[0] == "ns" && parts[2] == "resources" && parts[3] == "rdbms" && parts[4] == "validate" {
		return true
	}
	return false
}
