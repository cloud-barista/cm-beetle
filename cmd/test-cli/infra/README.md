# CM-Beetle Infra Test CLI

Automated end-to-end testing tool for CM-Beetle infrastructure recommendation and migration APIs.

## Overview

This CLI automates testing of CM-Beetle's infrastructure recommendation and migration APIs across multiple CSP-Region pairs.

### Features

- **Multi-CSP Support**: Test across AWS, Azure, GCP, Alibaba Cloud, NCP, and more
- **Parallel / Sequential Execution**: Configurable via `test.set.mode` in `testconf/test-config.yaml`
- **Complete Workflow**: End-to-end from recommendation → migration → SSH connectivity check → cleanup
- **SSH Connectivity Test**: Verify actual accessibility of all migrated VMs via Tumblebug access info
- **Automated Reports**: Generate detailed Markdown test reports per CSP-Region pair
- **Infrastructure Summaries**: Source (on-premise) and target (cloud) infrastructure analysis
- **Migration Reports**: Detailed source-to-target comparison report
- **Error Handling**: Skips remaining tests if an early test fails to prevent cascading failures

### Test Workflow

#### Initial Step (Before CSP-Region Pair Tests)

- **CM-Beetle Readiness Check**: `GET /beetle/readyz` — Verify CM-Beetle is available
- **Source Infrastructure Summary**: `POST /beetle/summary/source` — Summarize the on-premise source infrastructure

#### Per CSP-Region Pair (9 Tests)

| #   | API / Description                                                                              |
| --- | ---------------------------------------------------------------------------------------------- |
| 1   | `POST /beetle/recommendation/infra` — Recommend target infrastructure (uses first candidate)   |
| 2   | `POST /beetle/migration/ns/{nsId}/infra` — Migrate infrastructure                              |
| 3   | `GET /beetle/migration/ns/{nsId}/infra` — List all infras                                      |
| 4   | `GET /beetle/migration/ns/{nsId}/infra?option=id` — List infra IDs                             |
| 5   | `GET /beetle/migration/ns/{nsId}/infra/{infraId}` — Get specific infra                         |
| 6   | Remote Command Test — SSH connectivity check for all migrated VMs                              |
| 7   | `GET /beetle/summary/target/ns/{nsId}/infra/{infraId}` — Target infrastructure summary         |
| 8   | `POST /beetle/report/migration/ns/{nsId}/infra/{infraId}` — Migration report                   |
| 9   | `DELETE /beetle/migration/ns/{nsId}/infra/{infraId}?option=terminate` — Delete infra (cleanup) |

> If any test from 1–8 fails, subsequent tests for that CSP-Region pair are skipped. Test 9 (cleanup) always runs when an `infraId` is available.

## Quick Start

### 1. Configure

```bash
# Copy templates (done automatically by make test-infra if not present)
cp testconf/template-test-config.yaml testconf/test-config.yaml
cp testconf/template-auth-config.json testconf/auth-config.json
```

Edit `testconf/test-config.yaml` to select target CSP-Region pairs and set the Beetle endpoint.

Edit `testconf/auth-config.json` to set API credentials.

### 2. Run

```bash
# From the repository root
make test-infra
```

### 3. Check Results

- Console: Real-time progress and final summary
- Files: `cmd/test-cli/infra/testresult/beetle-test-results-{csp}.md`

## Configuration

### Test Config

`testconf/test-config.yaml`:

```yaml
test:
  set:
    mode: parallel # parallel or sequential
  cases:
    - csp: aws
      region: ap-northeast-2
      name: AWS-Seoul
      execute: true
    - csp: azure
      region: koreasouth
      name: Azure-Busan
      execute: false
    # ... more CSP-Region pairs

beetle:
  endpoint: http://localhost:8056
  namespaceId: mig01
  authConfigFile: testconf/auth-config.json
  requestBodyFile: testconf/recommendation-request.json
```

### Authentication Config

`testconf/auth-config.json` (copy from `testconf/template-auth-config.json` and fill in credentials):

```json
{
  "beetleApiUsername": "your-beetle-api-username",
  "beetleApiPassword": "your-beetle-api-password",
  "tumblebugApiUsername": "your-tumblebug-username",
  "tumblebugApiPassword": "your-tumblebug-password",
  "tumblebugEndpoint": "http://localhost:1323"
}
```

| Field                                           | Description                                                      |
| ----------------------------------------------- | ---------------------------------------------------------------- |
| `beetleApiUsername` / `beetleApiPassword`       | Credentials for CM-Beetle REST API (Basic Auth)                  |
| `tumblebugApiUsername` / `tumblebugApiPassword` | Credentials for CB-Tumblebug REST API (used in Test 6 SSH check) |
| `tumblebugEndpoint`                             | CB-Tumblebug base URL                                            |

### Request Body (On-Premise Infra Model)

`testconf/recommendation-request.json` — On-premise infrastructure data used as the source for recommendation and migration tests.

## CLI Options

```bash
cd cmd/test-cli/infra
go run main.go -config testconf/test-config.yaml
```

| Option    | Default                     | Description                     |
| --------- | --------------------------- | ------------------------------- |
| `-config` | `testconf/test-config.yaml` | Path to test configuration file |

## Test Results

### Console Output Example

```
============================================================
SOURCE INFRASTRUCTURE SUMMARY
============================================================

============================================================
Testing CSP-Region Pair 1/1: AWS-Seoul
============================================================

--- Test 1: POST /beetle/recommendation/infra ---
✅ Test 1 passed

--- Test 2: POST /beetle/migration/ns/mig01/infra ---
✅ Test 2 passed

--- Test 3: GET /beetle/migration/ns/mig01/infra ---
✅ Test 3 passed

--- Test 4: GET /beetle/migration/ns/mig01/infra?option=id ---
✅ Test 4 passed

--- Test 5: GET /beetle/migration/ns/mig01/infra/{infraId} ---
✅ Test 5 passed

--- Test 6: Remote Command Accessibility Check ---
✅ Test 6 passed

--- Test 7: Target Infrastructure Summary ---
✅ Test 7 passed

--- Test 8: Migration Report ---
✅ Test 8 passed

--- Test 9: DELETE /beetle/migration/ns/mig01/infra/{infraId} ---
✅ Test 9 passed

============================================================
OVERALL TEST SUMMARY
============================================================
Total CSP-Region Pairs: 1
Successful Pairs: 1
Total Tests: 9
Passed Tests: 9
```

### Generated Report Files

All files are saved under `testresult/` (relative to `cmd/test-cli/infra/`):

| File                                   | Description                                   |
| -------------------------------------- | --------------------------------------------- |
| `beetle-summary-source.md`             | Source (on-premise) infrastructure summary    |
| `beetle-test-results-{csp}.md`         | Full test results per CSP-Region pair         |
| `beetle-summary-target-{csp}.md`       | Target (cloud) infrastructure summary per CSP |
| `beetle-report-mig-source-to-{csp}.md` | Migration report per CSP                      |

## Custom Configuration

### Test Specific CSPs

Create custom config file:

```json
{
  "beetleUrl": "http://localhost:8056",
  "namespaceId": "mig01",
  "desiredCspRegionPairs": [
    { "csp": "aws", "region": "ap-northeast-2" },
    { "csp": "azure", "region": "koreacentral" }
  ],
  "requestBodyFile": "testconf/recommendation-request.json"
}
```

### Different Environment

```json
{
  "beetleUrl": "http://my-server:8080",
  "namespaceId": "my-namespace",
  "authConfigFile": "testconf/my-auth-config.json",
  "desiredCspRegionPairs": [...]
}
```

## Prerequisites

1. CM-Beetle server running at `http://localhost:8056`
2. Configured namespace in CM-Beetle
3. CSP credentials configured for each cloud provider
4. Network connectivity to CM-Beetle and cloud providers
5. **Authentication setup**: Copy `testconf/template-auth-config.json` to `testconf/auth-config.json` and configure credentials
6. **SSH access**: Ensure VMs can be accessed via SSH for Test 6 (Remote Command Test)

## Troubleshooting

### Build Issues

```bash
go mod tidy
make clean
make test-infra-build
```

### Connection Issues

```bash
# Check CM-Beetle server
curl http://localhost:8056/beetle/readyz

# Verify config
cat testconf/config-multi-csp-and-region-pair.json
```

### Test Failures

1. Check logs: `log/beetle.log`
2. Verify namespace exists
3. Validate CSP credentials
4. Check network connectivity

## Directory Structure

```
cmd/test-cli/infra/
├── main.go                                    # Main CLI code
├── README.md                                  # This file
├── test-infra                                 # Built binary
├── testconf/
│   ├── config-multi-csp-and-region-pair.json # Main config
│   └── recommendation-request.json            # Test data
├── testresult/                                # Test reports
│   └── beetle-test-results-*.md
└── log/
    └── beetle.log                             # Test logs
```

Built with Go 1.26.2+ for quality assurance and multi-cloud compatibility testing of CM-Beetle.
