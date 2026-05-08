# CM-Beetle Data Migration Test CLI

An automated end-to-end test CLI for CM-Beetle's data migration feature.
It creates source and target Object Storages via CM-Beetle, uploads dummy data, migrates it across CSP-Region pairs, verifies the result, and generates a per-pair Markdown report.

## Prerequisites

- CM-Beetle running at the endpoint configured in `test-config.yaml` (default: `http://localhost:8056`)
- CB-Tumblebug running and accessible (default: `http://localhost:1323/tumblebug`)
- CB-Spider running behind CB-Tumblebug
- Valid CSP credentials registered in CB-Spider for every CSP-Region pair under `test.cases`
- Go 1.21+ (only needed when running directly with `go run`)

## Directory structure

```
cmd/test-cli/data/
├── main.go                           # Test CLI entry point
├── dummydata/                        # Auto-generated dummy files (created at runtime)
├── log/                              # Log files from each test run
├── testconf/
│   ├── template-test-config.yaml     # Template — copy to test-config.yaml and edit
│   ├── test-config.yaml              # Active configuration (git-ignored)
│   ├── template-auth-config.json     # Template — copy to auth-config.json and edit
│   ├── auth-config.json              # CM-Beetle basic auth credentials (git-ignored)
│   └── object-storage-creation-req.json  # Bucket specification for OS creation
└── testresult/                       # Per-pair Markdown reports (created at runtime)
```

## Configuration

### 1. `testconf/test-config.yaml`

Copy the template and fill in your values:

```bash
cp testconf/template-test-config.yaml testconf/test-config.yaml
```

Key fields:

| Field                        | Description                                                                                                     |
| ---------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `test.set.mode`              | `sequential` or `parallel` — execution mode for target pairs                                                    |
| `test.source.csp`            | CSP for the shared source Object Storage (e.g. `aws`)                                                           |
| `test.source.region`         | Region for the shared source Object Storage (e.g. `ap-northeast-2`)                                             |
| `test.source.nameSeed`       | Name prefix for the source OS (defaults to `migration.nameSeed + "s"`)                                          |
| `test.upload.filter`         | Glob patterns to include/exclude when uploading dummy data                                                      |
| `test.cases`                 | List of target CSP-Region pairs; set `execute: true` to enable                                                  |
| `beetle.endpoint`            | CM-Beetle REST API base URL                                                                                     |
| `beetle.namespaceId`         | CB-Tumblebug namespace to use                                                                                   |
| `beetle.authConfigFile`      | Path to `auth-config.json`                                                                                      |
| `beetle.osCreationReqFile`   | Path to `object-storage-creation-req.json`                                                                      |
| `tumblebug.endpoint`         | Tumblebug URL used by the test CLI (host network) for direct `transx` uploads                                   |
| `tumblebug.serverEndpoint`   | Tumblebug URL placed inside migration request bodies (used by CM-Beetle server, typically a Docker network URL) |
| `tumblebug.auth`             | Tumblebug basic auth credentials                                                                                |
| `testData.baseDir`           | Directory for generated dummy files (default: `./dummydata`)                                                    |
| `testData.cleanup`           | Delete dummy files after the test run (`true`/`false`)                                                          |
| `migration.nameSeed`         | Name prefix for target Object Storages                                                                          |
| `migration.poll.intervalSec` | Polling interval while waiting for async migration (seconds)                                                    |
| `migration.poll.timeoutSec`  | Maximum wait time for migration completion (seconds)                                                            |

### 2. `testconf/auth-config.json`

Copy the template and fill in your CM-Beetle credentials:

```bash
cp testconf/template-auth-config.json testconf/auth-config.json
```

```json
{
  "basicAuthUsername": "your-api-username",
  "basicAuthPassword": "your-api-password"
}
```

### 3. `testconf/object-storage-creation-req.json`

Defines the bucket specification used when creating both source and target Object Storages.
Edit `bucketName` to match the bucket name registered in CB-Spider for your CSP accounts:

```json
{
  "targetObjectStorages": [
    {
      "sourceBucketName": "test-source",
      "bucketName": "cm-beetle-test-bucket",
      "versioningEnabled": false,
      "corsEnabled": false
    }
  ]
}
```

## How to run

Run from the `cmd/test-cli/data/` directory:

```bash
cd cmd/test-cli/data

# Basic run
go run main.go -config testconf/test-config.yaml

# Capture all output to a log file as well
go run main.go -config testconf/test-config.yaml 2>&1 | tee log/test-run-$(date +%Y%m%d-%H%M%S).log
```

You can also build first:

```bash
go build -o test-cli-data main.go
./test-cli-data -config testconf/test-config.yaml
```

## Test phases

The CLI runs in three phases:

### Pre-flight (once, before all target tests)

1. Check CM-Beetle readiness (`GET /beetle/readyz`)
2. Generate dummy data files in `testData.baseDir` (skipped if files already exist)
3. Create source Object Storage via CM-Beetle (`POST /beetle/migration/middleware/ns/{nsId}/objectStorage`)
   - A 409 Conflict response is treated as success (idempotent reuse of an existing OS)
4. Upload dummy data from the local filesystem to the source OS using `transx.Transfer()` directly
   - The upload applies the file filter from `test.upload.filter`
   - Encryption is not applied at this stage (local → source OS)

### Per target CSP-Region pair (Steps 1–4)

Each enabled target case in `test.cases` executes the following steps:

| Step | API                                                                        | Description                                                 |
| ---- | -------------------------------------------------------------------------- | ----------------------------------------------------------- |
| 1    | `POST /beetle/migration/middleware/ns/{nsId}/objectStorage`                | Create target Object Storage                                |
| 2    | `POST /beetle/migration/data`                                              | Migrate data from source OS to target OS (encrypted, async) |
| 3    | `GET /beetle/migration/middleware/ns/{nsId}/objectStorage/{osId}/object`   | Verify migrated objects match source                        |
| 4    | `DELETE /beetle/migration/middleware/ns/{nsId}/objectStorage/{targetOsId}` | Delete target OS (cleanup)                                  |

**Encryption**: Step 2 fetches a one-time RSA public key from CM-Beetle (`GET /beetle/migration/data/encryptionKey`) and encrypts sensitive fields (Tumblebug auth passwords) before sending the migration request.

**Async polling**: Step 2 returns `202 Accepted`. The CLI polls `GET /beetle/request/{reqId}` at `migration.poll.intervalSec` intervals until the status is `Success` or `Error`, or until `migration.poll.timeoutSec` is reached.

**Parallel mode**: In `parallel` mode, all enabled cases run concurrently. Each case gets a unique name suffix (e.g. `data0101`, `data0102`) to avoid naming collisions.

### Post-flight (once, after all target tests)

Deletes the shared source Object Storage. This runs unconditionally via `defer`, so it executes even if tests fail or the process receives `SIGINT`/`SIGTERM`.

**Delete phases** (applied to both source and target OS deletions):

1. `DELETE ?option=empty` — empties the bucket (once)
2. `DELETE` (standard) — up to 3 retries with 10-second delays
3. `DELETE ?option=force` — fallback if standard deletion fails, up to 3 retries

## Output

### Logs

Structured JSON logs (zerolog) are printed to stderr. Pipe to a file with `tee` to persist them:

```bash
go run main.go -config testconf/test-config.yaml 2>&1 | tee log/test-run-$(date +%Y%m%d-%H%M%S).log
```

### Markdown reports

A per-pair report is saved to `testresult/` after each target case completes:

```
testresult/data-test-results-{sourceCsp}-to-{targetCsp}.md
```

Each report contains:

- Environment details (CM-Beetle version, CSP, region, OS IDs, name seed)
- Test scenario description
- Step-by-step results table with status, duration, and request/response details
- Collapsible blocks for long request/response bodies
- Object tree comparison (source vs. target) from the verification step
- Sensitive values (Azure subscription IDs, GCP project IDs, email addresses) are automatically masked

### Exit code

| Code | Meaning                                             |
| ---- | --------------------------------------------------- |
| `0`  | All enabled test cases passed                       |
| `1`  | One or more test cases failed, or pre-flight failed |

## Signal handling

Sending `SIGINT` (Ctrl+C) or `SIGTERM` to the process cancels remaining test cases and prints a warning. Post-flight cleanup (source OS deletion) still runs before the process exits.

## Dummy data structure

The CLI auto-generates the following file set in `testData.baseDir`:

| Directory               | Files                   | Size range    | Type   |
| ----------------------- | ----------------------- | ------------- | ------ |
| `small/`                | 10 × `.txt`, 5 × `.md`  | 1–50 KB       | Text   |
| `medium/`               | 5 × `.json`, 3 × `.csv` | 100 KB – 1 MB | Text   |
| `large/`                | 2 × `.bin`              | 2–5 MB        | Binary |
| `docs/`                 | 3 × `.txt`, 2 × `.md`   | 5–100 KB      | Text   |
| `nested/level1/level2/` | 2 × `.txt`              | 1–10 KB       | Text   |
| `nested/level1/level3/` | 2 × `.txt`              | 1–10 KB       | Text   |

If `testData.baseDir` already contains files, generation is skipped.
