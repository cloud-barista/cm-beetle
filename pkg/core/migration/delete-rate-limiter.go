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

// Package migration provides client-side rate limiting for infrastructure operations
package migration

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

// DeleteInfraRateLimiter prevents TB's rate limit (2 req/s) by queueing and pacing requests.
//
// Architecture: [Goroutines] → [Queue: 50 max] → [Pacer: 600ms] → [TB API]
// - Queue full: Reject immediately
// - Rate exceeded: Slow down adaptively
// - Consistent success: Speed up gradually (down to 550ms)
//
// Implementation: Channel semaphore (queue) + Mutex (pacing)
type DeleteInfraRateLimiter struct {
	mu              sync.Mutex
	lastCallTime    time.Time
	minInterval     time.Duration
	currentInterval time.Duration
	queue           chan struct{} // Semaphore for queue capacity
	consecutiveOk   int
	consecutiveFail int
}

var (
	deleteRateLimiter *DeleteInfraRateLimiter
	rateLimiterOnce   sync.Once
)

const (
	maxQueueSize                = 50                      // Max concurrent waiting requests (~30s wait)
	defaultInterval             = 600 * time.Millisecond  // 1.67 req/s (20% below TB's 2 req/s)
	minAllowedInterval          = 550 * time.Millisecond  // Speed limit: 1.82 req/s
	maxAllowedInterval          = 1000 * time.Millisecond // Slow limit: 1 req/s
	consecutiveSuccessThreshold = 5                       // Successes needed to speed up
	intervalAdjustmentStep      = 50 * time.Millisecond   // Gradual adjustment step
	failureIntervalPenalty      = 200 * time.Millisecond  // Penalty on rate limit hit
)

// GetDeleteRateLimiter returns the singleton instance (thread-safe, initialized once).
func GetDeleteRateLimiter() *DeleteInfraRateLimiter {
	rateLimiterOnce.Do(func() {
		deleteRateLimiter = &DeleteInfraRateLimiter{
			minInterval:     defaultInterval,
			currentInterval: defaultInterval,
			queue:           make(chan struct{}, maxQueueSize),
		}
		log.Debug().Msgf("DeleteInfraRateLimiter initialized (interval: %v, max queue: %d, target rate: %.2f req/s)",
			defaultInterval, maxQueueSize, 1000.0/float64(defaultInterval.Milliseconds()))
	})
	return deleteRateLimiter
}

// WaitForSlot blocks until safe to call ReadInfra.
// Returns error if queue is full (>50 concurrent requests).
func (rl *DeleteInfraRateLimiter) WaitForSlot() error {
	// Try entering queue (non-blocking)
	select {
	case rl.queue <- struct{}{}:
		defer func() { <-rl.queue }()
		log.Debug().Msgf("Entered delete rate limiter queue (queue size: %d/%d)", len(rl.queue), maxQueueSize)
	default:
		err := fmt.Errorf("delete request queue is full (max: %d concurrent requests). Please retry later or reduce batch size", maxQueueSize)
		log.Warn().Msg(err.Error())
		return err
	}

	// Apply time-based pacing
	rl.mu.Lock()
	defer rl.mu.Unlock()

	elapsed := time.Since(rl.lastCallTime)
	if elapsed < rl.currentInterval {
		waitTime := rl.currentInterval - elapsed
		currentRate := 1000.0 / float64(rl.currentInterval.Milliseconds())
		log.Debug().Msgf("Rate limiting ReadInfra call (waiting %v, current rate: %.2f req/s)", waitTime, currentRate)
		time.Sleep(waitTime)
	}

	rl.lastCallTime = time.Now()
	return nil
}

// ReportSuccess gradually speeds up the limiter after consistent success.
func (rl *DeleteInfraRateLimiter) ReportSuccess() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.consecutiveOk++
	rl.consecutiveFail = 0

	if rl.consecutiveOk >= consecutiveSuccessThreshold && rl.currentInterval > minAllowedInterval {
		oldInterval := rl.currentInterval
		rl.currentInterval -= intervalAdjustmentStep
		if rl.currentInterval < minAllowedInterval {
			rl.currentInterval = minAllowedInterval
		}
		newRate := 1000.0 / float64(rl.currentInterval.Milliseconds())
		log.Debug().Msgf("Rate limiter speeding up: %v → %v (rate: %.2f req/s, consecutive ok: %d)",
			oldInterval, rl.currentInterval, newRate, rl.consecutiveOk)
		rl.consecutiveOk = 0
	}
}

// ReportFailure immediately slows down the limiter on rate limit error.
func (rl *DeleteInfraRateLimiter) ReportFailure() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.consecutiveFail++
	rl.consecutiveOk = 0

	oldInterval := rl.currentInterval
	rl.currentInterval += failureIntervalPenalty
	if rl.currentInterval > maxAllowedInterval {
		rl.currentInterval = maxAllowedInterval
	}
	newRate := 1000.0 / float64(rl.currentInterval.Milliseconds())
	log.Warn().Msgf("Rate limit detected - rate limiter slowing down: %v → %v (rate: %.2f req/s, consecutive fails: %d)",
		oldInterval, rl.currentInterval, newRate, rl.consecutiveFail)
}

// GetStats returns current limiter statistics for monitoring.
func (rl *DeleteInfraRateLimiter) GetStats() map[string]interface{} {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	return map[string]interface{}{
		"currentInterval":   rl.currentInterval.String(),
		"currentRate":       fmt.Sprintf("%.2f req/s", 1000.0/float64(rl.currentInterval.Milliseconds())),
		"queueSize":         len(rl.queue),
		"queueCapacity":     maxQueueSize,
		"consecutiveOk":     rl.consecutiveOk,
		"consecutiveFail":   rl.consecutiveFail,
		"lastCallTime":      rl.lastCallTime.Format(time.RFC3339),
		"timeSinceLastCall": time.Since(rl.lastCallTime).String(),
	}
}
