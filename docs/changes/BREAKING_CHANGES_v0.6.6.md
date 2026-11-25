# Breaking Changes for v0.6.6

**Release Version:** CM-Beetle v0.6.5 -> v0.6.6  
**Release Date:** 2026-11-25  
**Target Audience:** REST API clients and upper subsystems  
**Migration Required:** Yes - Breaking changes in Infra, Request Tracking, Data Migration, and Managed RDBMS REST APIs

---

## Breaking Changes Summary

| #   | Affected API                                                                                                                                                               | Impact                                                            | Required Action                                                                                                                                                |
| --- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1   | `POST /beetle/migration/ns/{nsId}/{infra,infraWithDefaults,k8sCluster}`, `GET /beetle/migration/ns/{nsId}/infra/{infraId}`, Recommendation / Validation / Naming endpoints | **BREAKING** - Post-boot script schema changed                    | Migrate `postCommands` to `postScript.phases`, remove `postCommandAsync` (use `Prefer: respond-async` header), and read `postScript*` response fields          |
| 2   | Infra Migration (`POST`/`GET`) and `GET /beetle/summary/target/ns/{nsId}/infra/{infraId}`                                                                                  | **BREAKING** - Legacy monitoring fields removed                   | Remove `installMonAgent`, `monAgentStatus`, and `Monitoring Agent` target summary row references                                                               |
| 3   | `DELETE /beetle/request/{reqId}`, `GET /beetle/requests?savefile=true`, request tracking scope                                                                             | **BREAKING** - Response structure and tracked routes changed      | Expect `success: true` on `DELETE`, read `data.file_path`/`data.count` inside `ApiResponse` on `savefile=true`, and note non-operation routes are not recorded |
| 4   | `POST /beetle/migration/data`                                                                                                                                              | **BREAKING** - Typed sync response and strict HTTP 400 validation | Parse `data` as structured `MigrationResult` and use valid 12-byte AES-GCM nonces for encrypted payloads                                                       |
| 5   | `/beetle/migration/middleware/ns/{nsId}/rdbms*`                                                                                                                            | **BREAKING** - Upstream 4xx passthrough and private RDBMS skip    | Handle HTTP `404`/`422` (instead of `500`) and run logical DB/schema setup in-VPC for private AWS/IBM RDBMS                                                    |

---

## Breaking Change #1: `postCommands` -> `postScript` (`PostScriptReq`)

### What Changed

- `POST /beetle/migration/ns/{nsId}/infra` and `POST /beetle/migration/ns/{nsId}/k8sCluster`: `targetInfra.postCommands` is replaced by `targetInfra.postScript`; `postCommandAsync` is removed from the JSON body.
- `POST /beetle/migration/ns/{nsId}/infraWithDefaults`: `postCommands` is replaced by `postScript` at the root JSON object; `postCommandAsync` is removed.
- `POST /beetle/validation/...`, `POST /beetle/naming/...`, and `POST /beetle/recommendation/{infra,infraWithNlb,multiInfra,multiInfraWithNlb}`: `RecommendedInfra` schema uses `postScript` (`postCommands` and `postCommandAsync` removed).
- `GET /beetle/migration/ns/{nsId}/infra/{infraId}` (and `POST` creation responses): `postCommandResults`, `postCommandStatus`, and `postCommandRequestId` are renamed to `postScriptResults`, `postScriptStatus`, and `postScriptRequestId`.

**Before (v0.6.5):**

```json
{
  "targetInfra": {
    "name": "my-infra",
    "installMonAgent": "no",
    "postCommands": [
      {
        "userName": "cb-user",
        "command": ["uname -a"]
      }
    ],
    "postCommandAsync": false
  }
}
```

**After (v0.6.6):**

```json
{
  "targetInfra": {
    "name": "my-infra",
    "postScript": {
      "phases": [
        {
          "userName": "cb-user",
          "command": ["uname -a"]
        }
      ]
    }
  }
}
```

> **Note:** For non-blocking asynchronous infrastructure migration, send the standard HTTP header `Prefer: respond-async` (`202 Accepted`) and poll `GET /beetle/request/{reqId}`.

### Migration Action

1. Replace `"postCommands": [...]` with `"postScript": {"phases": [...]}`.
2. Remove `"postCommandAsync"` from JSON request payloads; use HTTP header `Prefer: respond-async` when asynchronous execution is needed.
3. Update response parsers from `postCommandResults` / `postCommandStatus` / `postCommandRequestId` to `postScriptResults` / `postScriptStatus` / `postScriptRequestId`.

---

## Breaking Change #2: `installMonAgent` and `monAgentStatus` Removed

### What Changed

- `POST /beetle/migration/ns/{nsId}/infra`: `targetInfra.installMonAgent` removed from request and `InfraInfo` response.
- `POST /beetle/migration/ns/{nsId}/infraWithDefaults`: `installMonAgent` removed from request and `InfraInfo` response.
- `GET /beetle/migration/ns/{nsId}/infra/{infraId}`: `installMonAgent` removed from `InfraInfo`; `monAgentStatus` removed from `node[]` (`NodeInfo`).
- `GET /beetle/summary/target/ns/{nsId}/infra/{infraId}`: `overview.installMonAgent` removed from JSON response, and the `Monitoring Agent` row removed from Markdown output.

### Why This Changed

CB-Tumblebug `v0.14.0` removed legacy `installMonAgent` / `monAgentStatus` fields in favor of unified `initScript`-based observability agent provisioning (`templateId: "initscript-otel-agent"`).

### Migration Action

1. Remove `installMonAgent` from Infra creation request payloads.
2. Remove client references to `installMonAgent` (`InfraInfo`, `TargetInfraOverview`) and `monAgentStatus` (`NodeInfo`, `NodeSummary`).
3. To install an observability agent during node boot, use `initScript` (`templateId: "initscript-otel-agent"` or custom `scripts`).

---

## Breaking Change #3: `/beetle/request*` Response Formats and Tracking Scope

### What Changed

- `DELETE /beetle/request/{reqId}` now returns `"success": true` on successful deletion (previously returned `"success": false`).
- `GET /beetle/requests?savefile=true` now wraps the result in the standard `ApiResponse` envelope instead of returning top-level fields:
  - **Before (v0.6.5):** `{"message": "...", "file_path": "...", "count": N}`
  - **After (v0.6.6):** `{"success": true, "message": "...", "data": {"file_path": "...", "count": N}}`
- `GET /beetle/requests` no longer records non-operation routes (`/beetle/readyz`, `/beetle/httpVersion`, `/beetle/request*`, `/beetle/requests*`, `/beetle/api*`, `/beetle/docs*`, `/beetle/swagger*`, `/tumblebug*`, and `OPTIONS` preflight requests).
- Sensitive headers (`Authorization`, `Cookie`, `Set-Cookie`, `X-Api-Key`, tokens) and secret JSON fields in stored request records are redacted (`X-Request-Id` and tracing metadata remain visible).

### Migration Action

1. Update callers of `DELETE /beetle/request/{reqId}` to expect `"success": true`.
2. Update callers of `GET /beetle/requests?savefile=true` to read `data.file_path` and `data.count`.
3. Do not rely on `GET /beetle/requests` to query health checks, docs, request-tracking endpoints, or `/tumblebug/*` proxy calls.

---

## Breaking Change #4: `POST /beetle/migration/data` Response Schema and Encryption Validation

### What Changed

- Synchronous `POST /beetle/migration/data` (HTTP 200) now returns a structured `MigrationResult` object inside `ApiResponse.data` (including executed transfer details and bounded 8 KB UTF-8 hook outputs) instead of an untyped payload. On partial failure where hook outputs are preserved, HTTP 500 may also include `data`.
- Encrypted request payloads (using `GET /beetle/migration/data/encryptionKey`) enforce strict 12-byte AES-GCM nonce validation, and decryption failures consistently return **HTTP 400 Bad Request**.

### Migration Action

1. Update clients consuming synchronous `POST /beetle/migration/data` responses to parse `data` as `MigrationResult`.
2. Ensure clients encrypting data migration payloads generate valid 12-byte AES-GCM nonces and handle HTTP `400` on invalid or expired key payloads.

---

## Breaking Change #5: Managed RDBMS 4xx Passthrough and Private Mode Behavior

### What Changed

- Managed RDBMS endpoints (`/beetle/migration/middleware/ns/{nsId}/rdbms*`) now return upstream CB-Tumblebug HTTP `4xx` status codes directly (such as `404 Not Found` and `422 Unprocessable Entity`) instead of wrapping them as HTTP `500 Internal Server Error`.
- In private RDBMS mode (`publicAccess=false`), on SQL-only CSPs (`aws`, `ibm`), remote logical database creation is skipped rather than timing out when attempting to reach a private VPC endpoint from outside.
- **OpenStack Note:** Managed RDBMS **MariaDB 11.4** is verified; **MySQL 8.4** is currently marked **In Progress (`🚧`)** due to an upstream OpenStack Trove `EnableRootUser` error (`Plugin 'mysql_native_password' is not loaded`, see [cb-spider#1941](https://github.com/cloud-barista/cb-spider/issues/1941)).

### Migration Action

1. Update RDBMS API clients to handle HTTP `404` and `422` directly instead of matching HTTP `500` error bodies.
2. When provisioning private AWS/IBM RDBMS (`publicAccess=false`), execute logical database, schema, and data setup from inside the target VPC/VNet (for example, via `POST /beetle/migration/data` hooks or CM-Centipede).

---

## API Endpoint Change Matrix

| Method | Endpoint                                                                   | Breaking Change           | Summary                                                                                               |
| ------ | -------------------------------------------------------------------------- | ------------------------- | ----------------------------------------------------------------------------------------------------- |
| POST   | `/beetle/migration/ns/{nsId}/infra`                                        | Request & response fields | `targetInfra.postCommands` -> `targetInfra.postScript`; `installMonAgent`, `postCommandAsync` removed |
| POST   | `/beetle/migration/ns/{nsId}/infraWithDefaults`                            | Request & response fields | `postCommands` -> `postScript`; `installMonAgent`, `postCommandAsync` removed                         |
| POST   | `/beetle/migration/ns/{nsId}/k8sCluster`                                   | Request & response fields | `targetInfra.postCommands` -> `targetInfra.postScript`; `installMonAgent`, `postCommandAsync` removed |
| POST   | `/beetle/validation/ns/{nsId}/{infra,k8sCluster}`                          | Request schema            | `RecommendedInfra` updated (`postScript`, no `postCommands`/`installMonAgent`)                        |
| POST   | `/beetle/naming/{alignment,preview,validation}`                            | Request schema            | `RecommendedInfra` updated (`postScript`, no `postCommands`/`installMonAgent`)                        |
| POST   | `/beetle/recommendation/{infra,infraWithNlb,multiInfra,multiInfraWithNlb}` | Response schema           | Returned `RecommendedInfra[]` omits `postCommands`, `postCommandAsync`, `installMonAgent`             |
| GET    | `/beetle/migration/ns/{nsId}/infra/{infraId}`                              | Response fields           | `postCommand*` -> `postScript*`; `installMonAgent`, `monAgentStatus` removed                          |
| GET    | `/beetle/summary/target/ns/{nsId}/infra/{infraId}`                         | Response field            | `overview.installMonAgent` and Markdown `Monitoring Agent` row removed                                |
| GET    | `/beetle/requests?savefile=true`                                           | Response envelope         | Wrapped in `ApiResponse` (`data.file_path`, `data.count`)                                             |
| DELETE | `/beetle/request/{reqId}`                                                  | Response flag             | Returns `"success": true`                                                                             |
| POST   | `/beetle/migration/data`                                                   | Response schema & errors  | Sync HTTP 200 returns structured `MigrationResult` in `data`; decryption errors return HTTP `400`     |
| ALL    | `/beetle/migration/middleware/ns/{nsId}/rdbms*`                            | Error status codes        | Preserves upstream HTTP `404` / `422` directly                                                        |

---

## New API Capabilities (Non-Breaking)

- **First-boot `initScript`**: Supported on Infra and NodeGroup creation payloads, with node-level status fields (`initScriptStatus`, `initScriptMessage`, `initScriptSteps`) in `NodeInfo` and `NodeSummary`.
- **Template & Multi-Step Commands**: `InfraCmdReq` supports `templateId`, `scripts`, and `env` (`command` is required only when neither `templateId` nor `scripts` is set).
- **Spec Quota & GPU Partitioning**: `SpecInfo` adds `availabilityQuota` and `acceleratorPartitionSize`.
- **New `postScriptStatus` Value**: Added `"Interrupted"`.
- **Sub-Request Trace & Lightweight Polling**: `GET /beetle/request/{reqId}` adds `brief`, `withSubRequests`, and `withTumblebug` query parameters (`withTumblebug=true` enriches active sub-requests with live `upstreamStatus` and `upstreamProgress`); `GET /beetle/requests` adds `source`, `limit`, and `brief`.
- **Symmetric Data Migration Hooks**: `POST /beetle/migration/data` supports `preCmd` and `postCmd` on both `source` and `destination` with failure recovery on `source.postCmd`.
- **On-Premise `listeningPorts`**: On-premise `NodeProperty` adds optional `listeningPorts` (`[]int`).

---

## Client Migration Checklist

- [ ] Search client code for `postCommands`, `postCommandAsync`, `postCommandResults`, `postCommandStatus`, `postCommandRequestId`, `installMonAgent`, and `monAgentStatus`
- [ ] Update Infra / K8s creation payloads to use `postScript.phases` and remove `installMonAgent` and `postCommandAsync`
- [ ] Pass `Prefer: respond-async` header if asynchronous migration execution is required
- [ ] Update Infra and Target Summary response parsers for `postScript*` fields and removal of `installMonAgent` / `monAgentStatus`
- [ ] Update `DELETE /beetle/request/{reqId}` (`success: true`) and `GET /beetle/requests?savefile=true` (`data.file_path`, `data.count`) callers
- [ ] Update `POST /beetle/migration/data` parsers for `MigrationResult` and HTTP `400` decryption errors
- [ ] Update RDBMS API error handlers to accept HTTP `404` and `422`
