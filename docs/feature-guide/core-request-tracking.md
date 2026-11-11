# Request Tracking and Sub-Request Progress

## Overview

CM-Beetle tracks incoming API requests in `lkvstore` using a unique `reqId` (from the `X-Request-Id` request/response header) and correlates every outbound mutating or recommendation call to CB-Tumblebug through **Transport-AOP Sub-Request Tracking**.

This design gives CLI callers, Beetle UX Lab, and Cloud-Migrator Portal (`cm-butterfly`) real-time step-level visibility into multi-resource recommendation and migration pipelines without scattering state-update code across domain functions or introducing circular package dependencies.

---

## Request Lifecycle and Skip Policy

### Status Lifecycle

Tracked requests transition through three canonical states aligned with CB-Tumblebug:

| Status | Meaning |
| --- | --- |
| `Handling` | Request is currently running (synchronously or inside `common.RunAsync`) |
| `Success` | Request completed successfully; final payload is stored in `responseData` |
| `Error` | Request failed; error message is stored in `errorResponse` (and optional `retry` metadata) |

### Middleware Skip Policy

To prevent read-heavy polling, documentation assets, and health probes from polluting `lkvstore`, `ShouldSkipRequestTracking` (`pkg/api/rest/middlewares/request-skip.go`) excludes:

- Request inspection endpoints: `/beetle/request`, `/beetle/request/*`, `/beetle/requests`, `/beetle/requests/*`
- Health and protocol checks: `/beetle/readyz`, `/beetle/httpVersion`
- Documentation and UI paths: `/beetle/api*`, `/beetle/docs*`, `/beetle/swagger*`, `/beetle/ui*`
- Reverse proxy and browser metadata paths: `/tumblebug*`, `/.well-known*`, `*/favicon.ico`

---

## Transport-AOP Sub-Request Tracking Architecture

### Why Transport-AOP Instead of Inline Progress Updates

A single CM-Beetle API call such as `POST /beetle/migration/ns/{nsId}/infra` fans out into a sequence of CB-Tumblebug operations (`VNet` -> `Security Groups` -> `SSH Key` -> `VM Infra`). Updating migration state manually inside every core function requires repetitive boilerplate and careful locking around `RequestDetails.ResponseData`.

Instead, CM-Beetle uses a **single-direction context + HTTP client hook architecture**:

```
REST Controller / Middleware (pkg/api/rest/...)
   │  WithRequestID(ctx, reqId)
   ▼
Core Domain Functions (pkg/core/{recommendation,migration,...})
   │  ctx passed as 1st parameter -> tbclient.NewSession().SetContext(ctx)
   ▼
Tumblebug HTTP Transport (pkg/client/tumblebug/client-transport.go & client-subrequest.go)
   │  TumblebugTransport.RoundTrip -> beginSubRequestTrace
   ▼
In-Memory Sub-Request Tracker (pkg/core/common/subrequest.go)
```

Key architectural properties:

1. **Zero Circular Dependencies**: `pkg/core/common` owns the `reqId` context key and `SubRequest` tracker; `pkg/client/tumblebug` imports `common` to record outbound calls; `pkg/core/*` passes `ctx` downward without any reverse dependency.
2. **Thread-Safe Snapshot & Finalization**: `subRequestTracker` guards per-request slices with `sync.Mutex`. Live polling calls `SnapshotSubRequests(reqId)` (returning a race-free copy), while request completion calls `FinalizeSubRequests(reqId)` once to persist the completed slice into `lkvstore` and release in-memory tracking state.
3. **Async Context Detachment**: When `Prefer: respond-async` is used, `common.RunAsync(ctx, reqId, fn)` derives its background context via `context.WithoutCancel(ctx)`. If the HTTP client disconnects after receiving `202 Accepted`, background cloud provisioning continues uninterrupted while still retaining `reqId` in `ctx`.

### Outbound CB-Tumblebug Header Correlation

Whenever `tbclient.NewSession().SetContext(ctx)` executes an HTTP request with a tracked `reqId` in `ctx`:

| Outbound Header | When Sent | Format | Purpose |
| --- | --- | --- | --- |
| `X-Request-Source` | Every outbound CB-Tumblebug call | `cm-beetle:{reqId}` | Groups all Tumblebug calls triggered by one Beetle request (`GET /tumblebug/requests?source=cm-beetle:{reqId}`) |
| `X-Request-Id` | Mutating / action calls (`POST`, `PUT`, `DELETE`) | `{reqId}-{seq}` (e.g., `abc123-1`, or `abc123-1-r2` on `429` retry) | Assigns a deterministic sub-request ID in CB-Tumblebug (`GET /tumblebug/request/{subReqId}`) |

---

## API Reference

### 1. Get Single Request Details

```
GET /beetle/request/{reqId}
```

| Query Parameter | Type | Default | Description |
| --- | --- | --- | --- |
| `brief` | `bool` | `false` | Omits `requestInfo.body` and `responseData` for lightweight high-frequency polling while preserving `status`, `subRequestSummary`, and `subRequests` |
| `withSubRequests` | `bool` | `true` | Includes the ordered `subRequests` array (`false` returns only `subRequestSummary`) |
| `withTumblebug` | `bool` | `false` | Queries CB-Tumblebug `GET /request/{subReqId}` (bounded by a 3s timeout) to populate `upstreamStatus`, `upstreamProgress`, and `upstreamError` |

#### Example Response (`GET /beetle/request/{reqId}?brief=true&withTumblebug=true`)

```json
{
  "success": true,
  "data": {
    "reqId": "1762849201442",
    "startTime": "2026-11-11T16:20:01.442+09:00",
    "endTime": "0001-01-01T00:00:00Z",
    "status": "Handling",
    "tumblebugSource": "cm-beetle:1762849201442",
    "requestInfo": {
      "method": "POST",
      "url": "/beetle/migration/ns/mig01/infra",
      "header": {
        "Prefer": "respond-async",
        "X-Request-Id": "1762849201442"
      }
    },
    "subRequestSummary": {
      "total": 4,
      "handling": 1,
      "success": 3,
      "error": 0,
      "currentAction": "POST infra (mig01-aws-infra)"
    },
    "subRequests": [
      {
        "seq": 1,
        "subReqId": "1762849201442-1",
        "method": "POST",
        "targetPath": "/ns/mig01/resources/vNet",
        "resourceType": "vNet",
        "resourceHint": "mig01-aws-vnet",
        "status": "Success",
        "statusCode": 201,
        "startTime": "2026-11-11T16:20:01.510+09:00",
        "endTime": "2026-11-11T16:20:08.112+09:00",
        "durationMs": 6602,
        "upstreamStatus": "Success"
      },
      {
        "seq": 4,
        "subReqId": "1762849201442-4",
        "method": "POST",
        "targetPath": "/ns/mig01/infra",
        "resourceType": "infra",
        "resourceHint": "mig01-aws-infra",
        "status": "Handling",
        "startTime": "2026-11-11T16:20:24.905+09:00",
        "upstreamStatus": "Handling"
      }
    ]
  }
}
```

### 2. List Tracked Requests

```
GET /beetle/requests
```

Returns `{ "success": true, "data": { "requests": [...] } }` sorted by `startTime` **descending** (newest first).

| Query Parameter | Type | Default | Description |
| --- | --- | --- | --- |
| `status` | `string` | `""` | Filter by status (`Handling`, `Success`, `Error`) |
| `method` | `string` | `""` | Filter by HTTP method (`GET`, `POST`, `PUT`, `DELETE`) |
| `url` | `string` | `""` | Filter by request URL substring |
| `source` | `string` | `""` | Filter by `X-Request-Source` header or `tumblebugSource` substring |
| `time` | `int` | `""` | Filter requests started within the last N minutes |
| `limit` | `int` | `0` | Cap the number of returned requests (`0` = unlimited) |
| `brief` | `bool` | `false` | Omit `requestInfo.body` and `responseData` across listed items |
| `savefile` | `string` | `false` | Set `"true"` to write filtered requests to `<beetleRoot>/log/request_log_<timestamp>.json` |

### 3. Delete Tracked Request Records

- `DELETE /beetle/request/{reqId}`: Removes a single request record (and any in-memory sub-request entries) and returns `SimpleSuccessResponse`.
- `DELETE /beetle/requests`: Clears all request records from `lkvstore` and memory, returning `SimpleSuccessResponse`.

---

## Recommended Polling Pattern for Portals (Beetle UX Lab / CM-Butterfly)

1. **Submit with Async Preference**: Send `Prefer: respond-async` and capture `{ reqId }` from the `202 Accepted` response.
2. **Lightweight Progress Polling**: Poll `GET /beetle/request/{reqId}?brief=true` at an interval of at least 10 seconds (`>= 10s`) while `status == "Handling"`. Render step transitions directly from `subRequestSummary` and `subRequests` (`vNet` -> `securityGroup` -> `sshKey` -> `infra` -> `nlb`), while using a local client clock for elapsed seconds.
3. **Node-Level Live Status (Optional)**: While a `subRequests` item with `resourceType == "infra"` is in `Handling`, poll `GET /beetle/migration/ns/{nsId}/infra/{infraId}` at an interval of at least 10 seconds (`>= 10s`) to display individual node provisioning states and public/private IP allocation.
4. **Final Result Fetch**: Once `status` becomes `Success` or `Error`, fetch `GET /beetle/request/{reqId}` (without `brief=true`) once to retrieve the complete `responseData` or `errorResponse` (and pass `?withTumblebug=true` when diagnosing upstream Tumblebug errors).

---

## Related

- [Async Responses](core-async-response.md) - `Prefer: respond-async` contract and supported endpoints
- [API Response Policy](../api-response-policy.md) - HTTP status codes and response envelopes
