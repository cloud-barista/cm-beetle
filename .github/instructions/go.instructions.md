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
- **No history/rationale narratives**: Never explain what a prior version did, why it was
  wrong, or "do not reintroduce X" warnings in code comments. That belongs in the commit
  message or PR description, not the source file. A comment describes the current code,
  not its diff.

**Example:**

```go
// ✅ Good: describes current behavior only
// isNcpImageCompatible checks if the image's CspImageName exactly matches one of the spec's CorrespondingImageIds.
func isNcpImageCompatible(spec SpecInfo, image ImageInfo) bool { ... }

// ❌ Bad: explains history and rationale instead of the code
// isNcpImageCompatible checks if NCP image is compatible with spec using CorrespondingImageIds.
// Matching mirrors CB-Tumblebug's own reference implementation (filterImagesByCorrespondingIds
// in src/core/resource/image.go): an exact equality check, with no permissive fallback.
// A prior version of this function tried to "recover" an ID via regex/Details lookups and
// treated extraction failures as compatible, which let unrelated pairs pass; do not
// reintroduce that behavior.
func isNcpImageCompatible(spec SpecInfo, image ImageInfo) bool { ... }
```

### Function Comments

- State what the function does in **one or two sentences** maximum.
- Don't repeat the function name verbatim at the start.
- Avoid multi-line godoc blocks with bullet points; prefer named return values.

**Example:**

```go
// ✅ Good: Clear and concise
// getPacer returns the process-wide pacer for TB's rate-limited read APIs.
var getPacer = sync.OnceValue(func() *ratelimit.Pacer { ... })

// ❌ Bad: Verbose and repetitive
// getPacer returns the process-wide ratelimit.Pacer instance for TB's read APIs.
// It is thread-safe and initialized only once using sync.OnceValue.
var getPacer = sync.OnceValue(func() *ratelimit.Pacer { ... })
```

### Type Comments

- Focus on **purpose** and **architecture** in 3-5 lines.
- Include critical constraints (limits, rates, patterns).
- Use diagrams or flow notation when helpful.

**Example:**

```go
// ✅ Good: Architecture visible at a glance
// Pacer spaces out calls to a resource with a fixed, known rate limit. Callers wait
// their turn rather than being rejected, so a burst is spread over time.
//
// Architecture: [callers] → [token bucket: one call per interval] → [resource]
type Pacer struct { ... }
```

### Constant & Variable Comments

- Use **inline comments** for constants/variables (not separate lines).
- Include units, limits, or rationale in parentheses.

**Example:**

```go
// ✅ Good: Inline with technical details
const (
    pacingInterval    = 600 * time.Millisecond  // 1.67 req/s (20% below TB's 2 req/s)
    defaultPacingWait = 8 * time.Second         // Wait budget when the context has no deadline
    maxPacingWait     = 120 * time.Second       // Upper clamp on the configurable budget
)

// ❌ Bad: Verbose separate comments
// defaultPacingWait is how long a call waits for a slot when its context
// carries no deadline of its own. It keeps callers from blocking forever
// when an operator has not configured a budget.
const defaultPacingWait = 8 * time.Second
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
