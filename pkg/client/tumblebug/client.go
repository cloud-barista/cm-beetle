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

// Package tbclient provides client functions to interact with CB-Tumblebug API
package tbclient

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

// * [Note]
// * This Tumblebug Client is used to interact with the CB-Tumblebug API.
// * The Client contains the Tumblebug APIs required for computing infrastructure migration.
// * Other APIs can be added as needed.

// Global Client Instance
var globalClient *TumblebugClient
var once sync.Once

// TumblebugClient is the factory for creating request sessions.
// It holds the shared configuration and connection pool.
type TumblebugClient struct {
	restyClient *resty.Client
}

// ApiConfig holds the configuration for the Tumblebug client.
type ApiConfig struct {
	RestUrl       string
	Username      string
	Password      string
	Timeout       time.Duration
	CacheDuration time.Duration
	HttpTransport http.RoundTripper
}

// Init initializes the global Tumblebug client.
func Init(cfg ApiConfig) {
	once.Do(func() {
		globalClient = NewClient(cfg)
		log.Info().Msg("Initialized Tumblebug Client")
	})
}

// NewClient creates a new TumblebugClient instance.
func NewClient(cfg ApiConfig) *TumblebugClient {
	if cfg.RestUrl == "" {
		log.Fatal().Msg("Tumblebug REST URL is required")
	}
	if cfg.Username == "" || cfg.Password == "" {
		log.Fatal().Msg("Tumblebug API credentials (username/password) are required")
	}

	r := resty.New()
	r.SetBaseURL(cfg.RestUrl)
	r.SetBasicAuth(cfg.Username, cfg.Password)

	// Set Timeout
	if cfg.Timeout > 0 {
		r.SetTimeout(cfg.Timeout)
	} else {
		r.SetTimeout(10 * time.Minute)
	}

	// Set Transport
	if cfg.HttpTransport != nil {
		r.SetTransport(cfg.HttpTransport)
	} else {
		// Default to TumblebugTransport with default settings
		cacheDuration := cfg.CacheDuration
		if cacheDuration == 0 {
			cacheDuration = VeryShortDuration
		}
		r.SetTransport(&TumblebugTransport{
			Transport:     http.DefaultTransport,
			CacheDuration: cacheDuration,
		})
	}

	return &TumblebugClient{
		restyClient: r,
	}
}

// NewDefaultClient creates a new default client.
func NewDefaultClient() *TumblebugClient {
	cfg := ApiConfig{
		RestUrl:  config.Tumblebug.RestUrl,
		Username: config.Tumblebug.API.Username,
		Password: config.Tumblebug.API.Password,
	}
	// Fallback if RestUrl is empty, try Endpoint
	if cfg.RestUrl == "" {
		cfg.RestUrl = config.Tumblebug.Endpoint
	}
	return NewClient(cfg)
}

// Session represents a single request scope.
// It holds the request-specific configuration (headers, context, etc.).
type Session struct {
	req *resty.Request
}

// NewSession creates a new Session for making requests.
func (c *TumblebugClient) NewSession() *Session {
	return &Session{
		req: c.restyClient.R(),
	}
}

// NewSession creates a new Session using the global client.
func NewSession() *Session {
	if globalClient == nil {
		once.Do(func() {
			cfg := ApiConfig{
				RestUrl:  config.Tumblebug.RestUrl,
				Username: config.Tumblebug.API.Username,
				Password: config.Tumblebug.API.Password,
			}
			// Fallback if RestUrl is empty, try Endpoint
			if cfg.RestUrl == "" {
				cfg.RestUrl = config.Tumblebug.Endpoint
			}
			globalClient = NewClient(cfg)
			log.Info().Msg("Initialized Tumblebug Client (Lazy)")
		})
	}
	return globalClient.NewSession()
}

// SetHeader sets a single header.
func (s *Session) SetHeader(key, value string) *Session {
	s.req.SetHeader(key, value)
	return s
}

// SetHeaders sets multiple headers.
func (s *Session) SetHeaders(headers map[string]string) *Session {
	s.req.SetHeaders(headers)
	return s
}

// SetBody sets the request body.
func (s *Session) SetBody(body any) *Session {
	s.req.SetBody(body)
	return s
}

// SetResult sets the result object to unmarshal the response into.
func (s *Session) SetResult(res any) *Session {
	s.req.SetResult(res)
	return s
}

// SetContext sets the context for the request.
func (s *Session) SetContext(ctx context.Context) *Session {
	s.req.SetContext(ctx)
	return s
}

// SetTraceInfo injects the OpenTelemetry trace context into the request headers.
// This enables distributed tracing by propagating the trace ID and span ID to the Tumblebug API.
// It uses the global TextMapPropagator to ensure compatibility with the configured tracing backend.
//
// Note: This method modifies the request headers in-place (merging with existing headers).
// It does not affect the request context or other headers, so it is safe to call regardless of the order of SetHeader or SetContext.
func (s *Session) SetTraceInfo(ctx context.Context) *Session {
	// [OpenTelemetry]
	// TODO: Uncomment the following code when OTel is integrated.
	/*
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(s.req.Header))
	*/
	return s
}

// Execute executes the request with the given method and URL.
func (s *Session) Execute(method, url string) (*resty.Response, error) {
	return s.req.Execute(method, url)
}

// Get executes a GET request.
func (s *Session) Get(url string) (*resty.Response, error) {
	return s.req.Get(url)
}

// Post executes a POST request.
func (s *Session) Post(url string) (*resty.Response, error) {
	return s.req.Post(url)
}

// Put executes a PUT request.
func (s *Session) Put(url string) (*resty.Response, error) {
	return s.req.Put(url)
}

// Delete executes a DELETE request.
func (s *Session) Delete(url string) (*resty.Response, error) {
	return s.req.Delete(url)
}

// Head executes a HEAD request.
func (s *Session) Head(url string) (*resty.Response, error) {
	return s.req.Head(url)
}

// SetCacheDuration sets the cache duration for the request.
func (s *Session) SetCacheDuration(d time.Duration) *Session {
	ctx := s.req.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	s.req.SetContext(SetCacheDuration(ctx, d))
	return s
}

// NoCache disables caching for the request.
func (s *Session) NoCache() *Session {
	return s.SetCacheDuration(0)
}
