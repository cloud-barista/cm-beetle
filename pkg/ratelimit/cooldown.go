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

package ratelimit

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Cooldown rejects repeated actions on the same key that arrive less than interval apart
// (e.g. "this API may be called for a given resource once every N minutes"). Unlike Pacer
// it never delays a caller: an action is either allowed now or refused with the wait time.
// It is implemented as one rate.Limiter (burst 1) per key.
type Cooldown struct {
	mu       sync.Mutex
	limiters map[string]*cooldownEntry
	interval time.Duration
	maxAge   time.Duration
	stopCh   chan struct{}
}

type cooldownEntry struct {
	limiter  *rate.Limiter
	lastUsed time.Time
}

// NewCooldown creates a Cooldown enforcing interval between allowed calls per key.
// Entries unused for longer than maxAge are periodically evicted (checked every
// cleanupInterval) to bound memory; pass maxAge or cleanupInterval as 0 to disable
// cleanup (e.g. for short-lived Cooldowns such as in tests).
func NewCooldown(interval, maxAge, cleanupInterval time.Duration) *Cooldown {
	c := &Cooldown{
		limiters: make(map[string]*cooldownEntry),
		interval: interval,
		maxAge:   maxAge,
	}
	if maxAge > 0 && cleanupInterval > 0 {
		c.stopCh = make(chan struct{})
		go c.cleanupLoop(cleanupInterval)
	}
	return c
}

// Allow reports whether an action for key is allowed now. If not, retryAfter is the
// duration to wait before the next allowed call.
func (c *Cooldown) Allow(key string) (allowed bool, retryAfter time.Duration) {
	c.mu.Lock()
	entry, ok := c.limiters[key]
	if !ok {
		entry = &cooldownEntry{limiter: rate.NewLimiter(rate.Every(c.interval), 1)}
		c.limiters[key] = entry
	}
	entry.lastUsed = time.Now()
	limiter := entry.limiter
	c.mu.Unlock()

	reservation := limiter.Reserve()
	if !reservation.OK() {
		// Only happens if burst (1) is smaller than the requested tokens (1), so this
		// branch is unreachable in practice; treat it as "not allowed" defensively.
		return false, 0
	}
	if delay := reservation.Delay(); delay > 0 {
		reservation.Cancel() // give the token back; caller didn't consume a slot
		return false, delay
	}
	return true, 0
}

// cleanupLoop periodically evicts entries idle for longer than maxAge, until Stop is called.
func (c *Cooldown) cleanupLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.evictStale()
		case <-c.stopCh:
			return
		}
	}
}

func (c *Cooldown) evictStale() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, entry := range c.limiters {
		if now.Sub(entry.lastUsed) > c.maxAge {
			delete(c.limiters, key)
		}
	}
}

// Stop stops the background cleanup goroutine, if any.
func (c *Cooldown) Stop() {
	if c.stopCh != nil {
		close(c.stopCh)
	}
}
