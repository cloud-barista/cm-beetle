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

## Import Conventions

```go
// External dependencies with specific versions
// CB-Tumblebug imports (version specified in go.mod)
tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"
tbresource "github.com/cloud-barista/cb-tumblebug/src/interface/rest/server/resource"
"github.com/cloud-barista/cb-tumblebug/src/core/common/netutil"

// cm-model imports - Two main packages (version specified in go.mod)
cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
onpremmodel "github.com/cloud-barista/cm-model/infra/on-premise-model"

// Internal packages
"github.com/cloud-barista/cm-beetle/pkg/config"
"github.com/cloud-barista/cm-beetle/pkg/core/common"
```
