# CM-Beetle Test CLI

Automated testing tool for CM-Beetle APIs across multiple cloud service providers.

## Overview

This CLI automates testing of CM-Beetle's infrastructure recommendation and migration APIs across multiple CSP environments.

### Features

- **Multi-CSP Support**: Test across AWS, Azure, GCP, Alibaba Cloud, and NCP simultaneously
- **Complete Workflow**: End-to-end testing from recommendation to migration to cleanup
- **Comprehensive Testing**: 7 API tests + 2 report generations + 1 cleanup operation per CSP-Region pair
- **SSH Connectivity Test**: Verify actual accessibility of migrated VMs
- **Automated Reports**: Generate detailed Markdown test reports for each CSP
- **Infrastructure Summaries**: Source and target infrastructure analysis
- **Migration Reports**: Detailed comparison and migration analysis
- **Real Data**: Uses actual on-premise infrastructure data for realistic testing
- **Error Handling**: Skips remaining tests if early test fails to prevent cascading issues

### Test Workflow

#### Initial Setup

Before testing CSP-Region pairs, the CLI performs:

- **CM-Beetle Readiness Check**: `GET /beetle/readyz` - Verify service availability
- **Source Infrastructure Summary**: `POST /beetle/summary/source` - Generate summary of on-premise infrastructure

#### Per CSP-Region Pair Tests

Each CSP-Region pair executes the following sequential operations:

1. `POST /beetle/recommendation/mci` - Infrastructure recommendation
2. `POST /beetle/migration/ns/{nsId}/mci` - Infrastructure migration
3. `GET /beetle/migration/ns/{nsId}/mci` - List MCIs
4. `GET /beetle/migration/ns/{nsId}/mci?option=id` - List MCI IDs
5. `GET /beetle/migration/ns/{nsId}/mci/{mciId}` - Get specific MCI
6. **Remote Command Test** - SSH connectivity and accessibility check for migrated VMs
7. **Target Infrastructure Summary**: `GET /beetle/summary/target/ns/{nsId}/mci/{mciId}` - Generate summary of migrated infrastructure
8. **Migration Report**: `POST /beetle/report/migration/ns/{nsId}/mci/{mciId}` - Comprehensive migration analysis report
9. `DELETE /beetle/migration/ns/{nsId}/mci/{mciId}?option=terminate` - Delete MCI

**Note**: If any test fails, remaining tests for that CSP-Region pair are skipped.

## Quick Start

### 1. Build

```bash
make test-cli-build
```

### 2. Run

```bash
make test-cli
```

### 3. Check Results

- Console: Real-time progress and summary
- Files: `testresult/beetle-test-results-{csp}.md`

## Configuration

### Main Config

## Configuration

### Main Config

`testdata/config-multi-csp-and-region-pair.json`:

```json
{
  "beetleUrl": "http://localhost:8056",
  "namespaceId": "mig01",
  "authConfigFile": "testdata/auth-config.json",
  "desiredCspRegionPairs": [
    { "csp": "aws", "region": "ap-northeast-2", "name": "AWS-Seoul" },
    { "csp": "azure", "region": "koreacentral", "name": "Azure-Korea" },
    { "csp": "gcp", "region": "asia-northeast3", "name": "GCP-Seoul" },
    { "csp": "alibaba", "region": "ap-northeast-2", "name": "Alibaba-Seoul" },
    { "csp": "ncp", "region": "kr", "name": "NCP-Korea" }
  ],
  "requestBodyFile": "testdata/recommendation-request.json"
}
```

### Authentication Config

`testdata/auth-config.json` (excluded from version control):

```json
{
  "basicAuthUsername": "your-username",
  "basicAuthPassword": "your-password"
}
```

**Note**: Copy `testdata/template-auth-config.json` to `testdata/auth-config.json` and update with your credentials.

### Test Data

`testdata/recommendation-request.json` - Contains actual on-premise infrastructure data for testing.

### Test Data

`testdata/recommendation-request.json` - Contains actual on-premise infrastructure data for testing.

## Usage

### Basic Commands

```bash
# Build and run all tests
make test-cli-build
make test-cli

# Show help
make test-cli-help

# Direct execution
cd cmd/test-cli
./test-cli
./test-cli -verbose
./test-cli -config my-config.json
```

### Options

- `-config`: Configuration file path (default: `testdata/config-multi-csp-and-region-pair.json`)
- `-verbose`: Enable detailed output

## Test Results

### Console Output

```
============================================================
SOURCE INFRASTRUCTURE SUMMARY
============================================================
✅ Source summary generated successfully

============================================================
Testing CSP-Region Pair 1/5: AWS-Seoul
============================================================

--- Test 1: POST /beetle/recommendation/mci ---
✅ Test 1 passed (Duration: 541ms)

--- Test 2: POST /beetle/migration/ns/mig01/mci ---
✅ Test 2 passed (Duration: 37.999s)

--- Test 3: GET /beetle/migration/ns/mig01/mci ---
✅ Test 3 passed (Duration: 125ms)

--- Test 4: GET /beetle/migration/ns/mig01/mci?option=id ---
✅ Test 4 passed (Duration: 98ms)

--- Test 5: GET /beetle/migration/ns/mig01/mci/{mciId} ---
✅ Test 5 passed (Duration: 142ms)

--- Test 6: Remote Command Accessibility Check ---
✅ Test 6 passed (Duration: 3.2s)
   - VM web-server-01: ✅ Connected successfully
   - VM was-server-01: ✅ Connected successfully

--- Target Infrastructure Summary ---
✅ Target summary generated successfully

--- Migration Report ---
✅ Migration report generated successfully

--- Test 7: DELETE /beetle/migration/ns/mig01/mci/{mciId} ---
✅ Test 7 passed (Duration: 25.3s)

============================================================
OVERALL TEST SUMMARY
============================================================
Total CSP-Region Pairs: 5
Successful Pairs: 4/5
Total Tests: 35
Passed Tests: 28/35
Overall Time: 3m45s
```

### Detailed Reports

Individual Markdown reports are generated in `testresult/`:

- `source-infra-summary.md` - On-premise infrastructure summary
- `beetle-test-results-aws.md` - AWS test results
- `beetle-test-results-azure.md` - Azure test results
- `beetle-test-results-gcp.md` - GCP test results
- `beetle-test-results-alibaba.md` - Alibaba Cloud test results
- `beetle-test-results-ncp.md` - NCP test results
- `target-infra-summary-{csp}-{mciId}.md` - Target infrastructure summaries (per CSP)
- `migration-report-{csp}-{mciId}.md` - Migration reports (per CSP)

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
  "requestBodyFile": "testdata/recommendation-request.json"
}
```

### Different Environment

```json
{
  "beetleUrl": "http://my-server:8080",
  "namespaceId": "my-namespace",
  "authConfigFile": "testdata/my-auth-config.json",
  "desiredCspRegionPairs": [...]
}
```

## Prerequisites

1. CM-Beetle server running at `http://localhost:8056`
2. Configured namespace in CM-Beetle
3. CSP credentials configured for each cloud provider
4. Network connectivity to CM-Beetle and cloud providers
5. **Authentication setup**: Copy `testdata/template-auth-config.json` to `testdata/auth-config.json` and configure credentials
6. **SSH access**: Ensure VMs can be accessed via SSH for Test 6 (Remote Command Test)

## Troubleshooting

### Build Issues

```bash
go mod tidy
make clean
make test-cli-build
```

### Connection Issues

```bash
# Check CM-Beetle server
curl http://localhost:8056/beetle/readyz

# Verify config
cat testdata/config-multi-csp-and-region-pair.json
```

### Test Failures

1. Check logs: `log/beetle.log`
2. Verify namespace exists
3. Validate CSP credentials
4. Check network connectivity

## Directory Structure

```
cmd/test-cli/
├── main.go                                    # Main CLI code
├── README.md                                  # This file
├── test-cli                                   # Built binary
├── testdata/
│   ├── config-multi-csp-and-region-pair.json # Main config
│   └── recommendation-request.json            # Test data
├── testresult/                                # Test reports
│   └── beetle-test-results-*.md
└── log/
    └── beetle.log                             # Test logs
```

Built with Go 1.25+ for quality assurance and multi-cloud compatibility testing of CM-Beetle.
