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
	"sync"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/config"
	"github.com/cloud-barista/cm-beetle/pkg/ratelimit"
	"github.com/rs/zerolog/log"
)

const (
	// Tumblebug's limit on the read APIs Beetle paces. It applies per source IP, so one process
	// gets one budget. Configurable; see docs/tumblebug-call-pacer.md.
	defaultTumblebugRequestsPerSec = 2.0
	minTumblebugRequestsPerSec     = 0.1
	maxTumblebugRequestsPerSec     = 100.0

	// Fraction of Tumblebug's limit Beetle actually targets, absorbing clock skew and any other
	// client on this IP. At 2 req/s that is one call every 625ms.
	pacingHeadroom = 0.8

	// How long a call waits for a slot when its context carries no deadline of its own.
	defaultPacingWait = 8 * time.Second
	maxPacingWait     = 120 * time.Second

	// Bounded retry for the 429s pacing can't prevent. Tumblebug sends no Retry-After, so
	// retryOn429Wait is the effective spacing.
	retryOn429Count   = 3
	retryOn429Wait    = 1 * time.Second
	retryOn429MaxWait = 5 * time.Second
)

// getPacer returns the process-wide pacer for Tumblebug's read APIs. One pacer for the whole
// process, because Tumblebug's limit applies per source IP. Reads config on the first paced call.
var getPacer = sync.OnceValue(func() *ratelimit.Pacer {
	limit := tumblebugRequestsPerSec()
	interval := pacingInterval(limit)
	log.Info().Msgf("Tumblebug call pacer initialized (Tumblebug limit: %g req/s, paced at %g req/s / %v)",
		limit, limit*pacingHeadroom, interval)
	return ratelimit.NewPacer(interval)
})

// tumblebugRequestsPerSec returns the configured Tumblebug limit, clamped to a sane range.
func tumblebugRequestsPerSec() float64 {
	perSec := config.Tumblebug.Retrieval.RequestsPerSec
	switch {
	case perSec <= 0:
		return defaultTumblebugRequestsPerSec
	case perSec < minTumblebugRequestsPerSec:
		log.Warn().Msgf("Tumblebug rate limit %g req/s below the %g minimum; using %g", perSec, minTumblebugRequestsPerSec, minTumblebugRequestsPerSec)
		return minTumblebugRequestsPerSec
	case perSec > maxTumblebugRequestsPerSec:
		log.Warn().Msgf("Tumblebug rate limit %g req/s above the %g maximum; using %g", perSec, maxTumblebugRequestsPerSec, maxTumblebugRequestsPerSec)
		return maxTumblebugRequestsPerSec
	default:
		return perSec
	}
}

// pacingInterval converts Tumblebug's limit into the gap between Beetle's calls, applying
// pacingHeadroom so Beetle aims below that limit rather than at it.
func pacingInterval(requestsPerSec float64) time.Duration {
	return time.Duration(float64(time.Second) / (requestsPerSec * pacingHeadroom))
}

// pacingWaitBudget returns the configured wait, capped at maxPacingWait.
func pacingWaitBudget() time.Duration {
	sec := config.Tumblebug.Retrieval.MaxWaitSec
	if sec <= 0 {
		return defaultPacingWait
	}
	if d := time.Duration(sec) * time.Second; d < maxPacingWait {
		return d
	}
	return maxPacingWait
}

// pace blocks until this call's turn comes up, returning *ratelimit.ErrLimited rather
// than waiting when that turn falls outside the caller's budget. Set a deadline via
// Session.SetContext to choose that budget; without one, the configured default applies.
func (s *Session) pace(logCtx string) error {
	ctx := s.req.Context()
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, pacingWaitBudget())
		defer cancel()
	}

	start := time.Now()
	if err := getPacer().Wait(ctx); err != nil {
		log.Warn().Err(err).Msgf("Tumblebug call pacer refused a slot (%s)", logCtx)
		return err
	}
	if waited := time.Since(start); waited > 10*time.Millisecond {
		log.Debug().Msgf("Paced Tumblebug call (%s): waited %v", logCtx, waited)
	}
	return nil
}
