---
applyTo: "**/*.go"
---

# Go Coding Standards

## Style & Conventions

- **Style:** Follow **Effective Go** and **Go Code Review Comments**.
- **Formatting:** Use `gofmt` or `goimports`.
- **Linting:** Ensure code passes `golangci-lint` (run `make lint`).

## Logging & Error Handling

- **Logging:** Use `zerolog`.
  - Pattern: `log.Error().Err(err).Msg("message")`
- **Error Handling:**
  - Handle errors explicitly.
  - Return meaningful error messages and HTTP status codes using `model.Response`.

## Configuration

- **Viper:** Use `viper` for configuration management.

## Comments

> **Note**: This section provides detailed guidelines referenced in `.github/copilot-instructions.md`.
> All comments must be written in English (project-wide standard).

### General Principles

- **Be concise**: Write short, clear comments that convey essential information only.
- **Be precise**: Include technical details (numbers, limits, ratios) when relevant.
- **Be actionable**: Focus on **what** the code does, not how (implementation is in the code).
- **Avoid redundancy**: Don't comment on obvious code behavior.

### Function Comments

- State what the function does in **one or two sentences** maximum.
- Don't repeat the function name verbatim at the start.
- Avoid multi-line godoc blocks with bullet points; prefer named return values.

**Example:**

```go
// ✅ Good: Clear and concise
// GetDeleteRateLimiter returns the singleton instance (thread-safe, initialized once).
func GetDeleteRateLimiter() *DeleteInfraRateLimiter { ... }

// ❌ Bad: Verbose and repetitive
// GetDeleteRateLimiter returns the singleton DeleteInfraRateLimiter instance.
// It is thread-safe and initialized only once using sync.Once.
func GetDeleteRateLimiter() *DeleteInfraRateLimiter { ... }
```

### Type Comments

- Focus on **purpose** and **architecture** in 3-5 lines.
- Include critical constraints (limits, rates, patterns).
- Use diagrams or flow notation when helpful.

**Example:**

```go
// ✅ Good: Architecture visible at a glance
// DeleteInfraRateLimiter prevents TB's rate limit (2 req/s) by queueing and pacing requests.
//
// Architecture: [Goroutines] → [Queue: 50 max] → [Pacer: 600ms] → [TB API]
// - Queue full: Reject immediately
// - Rate exceeded: Slow down adaptively
type DeleteInfraRateLimiter struct { ... }
```

### Constant & Variable Comments

- Use **inline comments** for constants/variables (not separate lines).
- Include units, limits, or rationale in parentheses.

**Example:**

```go
// ✅ Good: Inline with technical details
const (
    maxQueueSize       = 50                      // Max concurrent requests (~30s wait)
    defaultInterval    = 600 * time.Millisecond  // 1.67 req/s (20% below TB's 2 req/s)
    minAllowedInterval = 550 * time.Millisecond  // Speed limit: 1.82 req/s
)

// ❌ Bad: Verbose separate comments
// maxQueueSize limits the number of delete requests that can wait concurrently.
// This prevents unbounded goroutine accumulation and potential OOM.
// With 600ms interval, 50 requests = ~30 seconds max wait time.
const maxQueueSize = 50
```

### Inline Comments

- Use sparingly for non-obvious logic only.
- Keep them short (5 words or less when possible).
- Remove comments that merely restate the code.

**Example:**

```go
// ✅ Good: Adds value
defer func() { <-rl.queue }()  // Release queue slot

// ❌ Bad: States the obvious
i++  // Increment i by 1
```

### Comment Review Checklist

Before committing, verify each comment:

- [ ] Is it **necessary**? (Does it add information not obvious from code?)
- [ ] Is it **concise**? (Can it be shorter without losing meaning?)
- [ ] Is it **accurate**? (Does it include specific numbers/limits if relevant?)
- [ ] Is it **clear**? (Would a new developer understand immediately?)

## Import Conventions

```go
// External dependencies with specific versions
// CB-Tumblebug imports (version specified in go.mod)
tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
tbresource "github.com/cloud-barista/cb-tumblebug/src/interface/rest/server/resource"
"github.com/cloud-barista/cb-tumblebug/src/core/common/netutil"

// Internal model imports (imdl - internalized infrastructure models)
cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"

// Internal packages
"github.com/cloud-barista/cm-beetle/pkg/config"
"github.com/cloud-barista/cm-beetle/pkg/core/common"
```
