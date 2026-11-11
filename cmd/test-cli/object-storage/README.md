# CM-Beetle Object Storage Migration Test CLI

An automated end-to-end test CLI for CM-Beetle's Object Storage recommendation and migration workflow.
It runs a 6-step lifecycle test across configured CSP-Region pairs, verifies bucket creation and cleanup, and generates a per-CSP Markdown report.

## Prerequisites

- CM-Beetle running at the endpoint configured in `test-config.yaml` (default: `http://localhost:8056`)
- CB-Tumblebug running and accessible (default: `http://localhost:1323/tumblebug`)
- CB-Spider running behind CB-Tumblebug
- Valid CSP credentials and connection configs registered in CB-Spider / CB-Tumblebug for each enabled CSP-Region pair under `test.cases`
- Target namespace (`beetle.namespaceId`) created in CB-Tumblebug
- Go 1.24+ (when running directly with `go run` or `make test-os`)

## Directory Structure

```text
cmd/test-cli/object-storage/
├── main.go                             # Test CLI entry point
├── README.md                           # This guide
├── testconf/
│   ├── template-test-config.yaml       # Template: copy to test-config.yaml and edit
│   ├── test-config.yaml                # Active configuration (git-ignored)
│   ├── template-auth-config.json       # Template: copy to auth-config.json and edit
│   ├── auth-config.json                # CM-Beetle Basic Auth credentials (git-ignored)
│   └── recommendation-request.json     # Source bucket specifications for recommendation
└── testresult/                         # Per-CSP Markdown reports (created at runtime)
```

## Configuration

### 1. `testconf/test-config.yaml`

Copy the template and configure the target CSP-Region pairs:

```bash
cp testconf/template-test-config.yaml testconf/test-config.yaml
```

Key fields:

| Field | Description |
| --- | --- |
| `test.set.mode` | `sequential` or `parallel` execution mode across enabled CSP-Region pairs |
| `test.cases` | List of target CSP-Region pairs (`csp`, `region`, `name`, `execute`) |
| `beetle.endpoint` | CM-Beetle REST API server URL (for example, `http://localhost:8056`) |
| `beetle.namespaceId` | Target CB-Tumblebug namespace ID (for example, `mig01`) |
| `beetle.authConfigFile` | Path to `auth-config.json` |
| `beetle.requestBodyFile` | Path to `recommendation-request.json` |
| `migration.nameSeed` | Name seed passed via `?nameSeed=` during migration (in `parallel` mode, a numeric suffix such as `01`, `02` is appended per case to prevent bucket name collisions) |

### 2. `testconf/auth-config.json`

Copy the template and fill in CM-Beetle Basic Auth credentials:

```bash
cp testconf/template-auth-config.json testconf/auth-config.json
```

```json
{
  "basicAuthUsername": "your-api-username",
  "basicAuthPassword": "your-api-password"
}
```

### 3. `testconf/recommendation-request.json`

Defines `sourceObjectStorages` used to request target bucket recommendations (`POST /beetle/recommendation/middleware/objectStorage`).

## How to Run

From the repository root:

```bash
make test-os
```

Or directly from `cmd/test-cli/object-storage/`:

```bash
cd cmd/test-cli/object-storage

# Run with default config
go run main.go -config testconf/test-config.yaml

# Override nameSeed for a single run (useful to avoid global S3 bucket name collisions)
go run main.go -config testconf/test-config.yaml -nameSeed myteam2026
```

## 6-Step Test Workflow

For each CSP-Region pair with `execute: true`, the CLI executes:

1. **Recommend (`POST /beetle/recommendation/middleware/objectStorage`)**: Generates target bucket specifications (`os-01`, `os-02`, ...) for the target CSP and region.
2. **Migrate (`POST /beetle/migration/middleware/ns/{nsId}/objectStorage?nameSeed=...`)**: Creates target buckets, applies Versioning and CORS configurations, and parses the `201 Created` `MigratedObjectStorageResult` response (`status`, `warnings`, `objectStorage`).
3. **List (`GET /beetle/migration/middleware/ns/{nsId}/objectStorage`)**: Lists object storages registered in the namespace.
4. **Exist (`HEAD /beetle/migration/middleware/ns/{nsId}/objectStorage/{osId}`)**: Verifies that the first migrated bucket exists (`200 OK`).
5. **Get (`GET /beetle/migration/middleware/ns/{nsId}/objectStorage/{osId}`)**: Retrieves detailed metadata for the first migrated bucket.
6. **Cleanup (`DELETE /beetle/migration/middleware/ns/{nsId}/objectStorage/{osId}`)**: Always runs even if an earlier step failed. Attempts standard deletion first, and automatically retries with `?option=force` if standard deletion fails (for example, when object versions remain).

## Output Reports

After each CSP-Region test case finishes, a sanitized Markdown report (with subscription IDs, project IDs, and email addresses masked) is written to:

```text
cmd/test-cli/object-storage/testresult/os-test-results-{csp}.md
```
