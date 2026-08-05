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
	"errors"
	"fmt"
	"time"
)

// ErrLimited reports that a call was refused because a rate limit would otherwise be
// exceeded. RetryAfter is how long the caller should wait before trying again.
type ErrLimited struct {
	RetryAfter time.Duration
}

func (e *ErrLimited) Error() string {
	return fmt.Sprintf("rate limited, retry after %v", e.RetryAfter)
}

// RetryAfter reports whether err is (or wraps) an ErrLimited and, if so, how long the
// caller should wait. It lets any package translate a rate-limit failure into its own
// response without importing whichever package produced the error.
func RetryAfter(err error) (retryAfter time.Duration, ok bool) {
	var limited *ErrLimited
	if !errors.As(err, &limited) {
		return 0, false
	}
	return limited.RetryAfter, true
}

// RetryAfterSeconds converts d into whole seconds for a Retry-After header, rounding up
// so the value never advertises a shorter wait than the limiter actually enforces.
func RetryAfterSeconds(d time.Duration) int {
	seconds := int(d / time.Second)
	if d%time.Second != 0 {
		seconds++
	}
	if seconds < 1 {
		return 1
	}
	return seconds
}
