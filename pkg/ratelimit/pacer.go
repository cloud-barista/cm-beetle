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

// Package ratelimit provides small, reusable rate-limiting primitives built on top of
// the standard golang.org/x/time/rate token-bucket limiter.
//
// Rate limiting is enforced in one of two ways, and this package offers one type for each:
//   - Pacer shapes traffic: it delays calls so a resource with a known rate limit
//     (e.g. an API capped at N req/s) is never called faster than that limit. Use it on
//     calls you make to someone else.
//   - Cooldown polices traffic: it rejects a repeat action on the same key that arrives
//     before its interval has elapsed. Use it on calls others make to you.
//
// Neither type is tied to any particular transport or caller; wire them up where needed
// and translate the result into whatever response shape (HTTP 503/429 + Retry-After, etc.)
// is appropriate for that caller.
package ratelimit

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

// Pacer spaces out calls to a resource with a fixed, known rate limit. Callers wait their
// turn rather than being rejected, so a burst is spread over time instead of failing.
//
// Architecture: [callers] -> [token bucket: one call per interval] -> [resource]
type Pacer struct {
	interval time.Duration
	limiter  *rate.Limiter
}

// NewPacer creates a Pacer admitting one call per interval.
func NewPacer(interval time.Duration) *Pacer {
	return &Pacer{
		interval: interval,
		// Burst stays at 1 so an idle period can't accumulate tokens that would then
		// fire back-to-back and breach the limit.
		limiter: rate.NewLimiter(rate.Every(interval), 1),
	}
}

// Wait blocks until the caller's slot comes up. It returns *ErrLimited immediately,
// without waiting, when that slot would arrive after ctx's deadline, and ctx.Err() when
// ctx is canceled while waiting.
func (p *Pacer) Wait(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	reservation := p.limiter.Reserve()
	if !reservation.OK() {
		return &ErrLimited{RetryAfter: p.interval}
	}

	delay := reservation.Delay()
	if delay <= 0 {
		return nil
	}
	if deadline, ok := ctx.Deadline(); ok && time.Now().Add(delay).After(deadline) {
		reservation.Cancel() // give the slot back to whoever can actually wait for it
		return &ErrLimited{RetryAfter: delay}
	}

	timer := time.NewTimer(delay)
	defer timer.Stop()
	select {
	case <-timer.C:
		return nil
	case <-ctx.Done():
		reservation.Cancel()
		return ctx.Err()
	}
}

// Interval returns the delay enforced between admitted calls.
func (p *Pacer) Interval() time.Duration {
	return p.interval
}
