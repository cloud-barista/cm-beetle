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

// Package migration provides SSH readiness check rate limiting
package migration

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

// SSHCheckRecord tracks SSH readiness check for a specific infrastructure
type SSHCheckRecord struct {
	InfraKey      string    // Format: "nsId:infraId"
	LastCheckTime time.Time // Time of last check
	CheckCount    int       // Number of checks performed
}

// SSHCheckRateLimiter provides rate limiting for SSH readiness checks
// to prevent abuse and excessive SSH connection attempts
type SSHCheckRateLimiter struct {
	mu            sync.RWMutex
	records       map[string]*SSHCheckRecord // Key: "nsId:infraId"
	minInterval   time.Duration              // Minimum interval between checks for same infra
	cleanupTicker *time.Ticker               // Periodic cleanup of old records
}

var (
	// Global instance of SSH check rate limiter
	sshCheckLimiter *SSHCheckRateLimiter
	limiterOnce     sync.Once
)

// GetSSHCheckRateLimiter returns the singleton instance of SSH check rate limiter
func GetSSHCheckRateLimiter() *SSHCheckRateLimiter {
	limiterOnce.Do(func() {
		sshCheckLimiter = &SSHCheckRateLimiter{
			records:     make(map[string]*SSHCheckRecord),
			minInterval: 3 * time.Minute, // Minimum 3 minutes between checks for same infra
		}

		// Start background cleanup goroutine
		sshCheckLimiter.startCleanup()
	})
	return sshCheckLimiter
}

// CheckAllowed checks if an SSH readiness check is allowed for the given infrastructure
// Returns true if allowed, false if rate limited, along with time until next allowed check
func (l *SSHCheckRateLimiter) CheckAllowed(nsId string, infraId string) (bool, time.Duration, error) {
	if nsId == "" || infraId == "" {
		return false, 0, fmt.Errorf("nsId and infraId are required")
	}

	key := fmt.Sprintf("%s:%s", nsId, infraId)

	l.mu.Lock()
	defer l.mu.Unlock()

	record, exists := l.records[key]
	now := time.Now()

	if !exists {
		// First check for this infra - always allowed
		l.records[key] = &SSHCheckRecord{
			InfraKey:      key,
			LastCheckTime: now,
			CheckCount:    1,
		}
		log.Debug().Msgf("SSH check allowed (first check): %s", key)
		return true, 0, nil
	}

	// Check if enough time has passed since last check
	timeSinceLastCheck := now.Sub(record.LastCheckTime)
	if timeSinceLastCheck < l.minInterval {
		// Rate limited
		timeUntilNext := l.minInterval - timeSinceLastCheck
		log.Warn().Msgf("SSH check rate limited: %s (last check: %v ago, min interval: %v, retry after: %v)",
			key, timeSinceLastCheck.Round(time.Second), l.minInterval, timeUntilNext.Round(time.Second))
		return false, timeUntilNext, nil
	}

	// Allowed - update record
	record.LastCheckTime = now
	record.CheckCount++
	log.Debug().Msgf("SSH check allowed: %s (check count: %d, last check: %v ago)",
		key, record.CheckCount, timeSinceLastCheck.Round(time.Second))
	return true, 0, nil
}

// startCleanup starts a background goroutine to periodically clean up old records
// Records older than 1 hour are removed to prevent memory leak
func (l *SSHCheckRateLimiter) startCleanup() {
	l.cleanupTicker = time.NewTicker(10 * time.Minute)

	go func() {
		for range l.cleanupTicker.C {
			l.cleanup()
		}
	}()

	log.Info().Msg("SSH check rate limiter cleanup started (interval: 10 minutes)")
}

// cleanup removes records older than 1 hour
func (l *SSHCheckRateLimiter) cleanup() {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	maxAge := 1 * time.Hour
	removed := 0

	for key, record := range l.records {
		if now.Sub(record.LastCheckTime) > maxAge {
			delete(l.records, key)
			removed++
		}
	}

	if removed > 0 {
		log.Debug().Msgf("SSH check rate limiter cleanup: removed %d old records", removed)
	}
}

// Stop stops the background cleanup goroutine
func (l *SSHCheckRateLimiter) Stop() {
	if l.cleanupTicker != nil {
		l.cleanupTicker.Stop()
		log.Info().Msg("SSH check rate limiter cleanup stopped")
	}
}

// GetStats returns current statistics about the rate limiter
func (l *SSHCheckRateLimiter) GetStats() map[string]interface{} {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return map[string]interface{}{
		"total_records":  len(l.records),
		"min_interval":   l.minInterval.String(),
		"cleanup_active": l.cleanupTicker != nil,
	}
}
