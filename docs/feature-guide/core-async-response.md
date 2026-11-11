# Async Responses

## What is it?

Long-running migration and recommendation APIs support an optional asynchronous mode. By default they
respond synchronously, exactly as before. Send `Prefer: respond-async` to get an immediate `202
Accepted` instead, then poll for completion.

## How to Use It

```
POST /beetle/recommendation/infra
Prefer: respond-async

→ 202 Accepted
  Preference-Applied: respond-async
  { "reqId": "...", "status": "Handling", "statusUrl": "/beetle/request/{reqId}" }
```

```
GET /beetle/request/{reqId}

→ { "status": "Success" | "Error" | "Handling", ... }
```

Poll `GET /request/{reqId}` until `status` is `Success` or `Error`. For frequent UI polling while `status` is `Handling`, use `GET /request/{reqId}?brief=true` to receive `status`, `subRequestSummary`, and `subRequests` without transferring large request or response bodies, then call `GET /request/{reqId}` once upon completion. Omit the `Prefer` header (or send any other value) to get the original synchronous response; nothing changes for existing callers.

## Supported Endpoints

| Endpoint | Notes |
| --- | --- |
| `POST /migration/data` | Sync by default; stores `transx.MigrationResult` (including `hookOutputs`) in `responseData` |
| `POST /migration/ns/{nsId}/infra` | Emits `vNet`, `securityGroup`, `sshKey`, and `infra` steps in `subRequests` |
| `POST /migration/ns/{nsId}/infraWithDefaults` | Emits `vNet`, `securityGroup`, `sshKey`, and `infra` steps in `subRequests` |
| `DELETE /migration/ns/{nsId}/infra/{infraId}` | Worth using async: deletion includes several settle waits |
| `POST /migration/ns/{nsId}/k8sCluster` | Managed Kubernetes cluster provisioning |
| `POST /migration/middleware/ns/{nsId}/infra/{infraId}/nlb` | Emits NLB provisioning sub-requests |
| `DELETE /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}` | Worth using async: includes a fixed 15s settle wait |
| `POST /migration/middleware/ns/{nsId}/objectStorage` | Stores `MigratedObjectStorageResult` in `responseData` |
| `DELETE /migration/middleware/ns/{nsId}/objectStorage/{osId}` | Supports `?option=empty` and `?option=force` |
| `POST /migration/middleware/ns/{nsId}/rdbms` | Managed relational database provisioning |
| `POST /recommendation/infra` | Multi-candidate VM infrastructure recommendation |
| `POST /recommendation/infraWithDefaults` | Default-option VM infrastructure recommendation |
| `POST /recommendation/infraWithNlb` | VM infrastructure + NLB recommendation |
| `POST /recommendation/multiInfra` | Multi-group / multi-CSP infrastructure recommendation |
| `POST /recommendation/multiInfraWithNlb` | Multi-group infrastructure + NLB recommendation |
| `POST /recommendation/middleware/objectStorage` | Managed Object Storage recommendation |
| `POST /recommendation/middleware/rdbms` | Managed RDBMS recommendation |

## Status Values

| Status | Meaning |
| --- | --- |
| `Handling` | Still running (`subRequestSummary` and `subRequests` reflect live CB-Tumblebug step progress) |
| `Success` | Completed; result in `responseData` |
| `Error` | Failed; message in `errorResponse` |

## Context Detachment and Limits

- `common.RunAsync(ctx, reqId, fn)` detaches parent HTTP cancellation using `context.WithoutCancel(ctx)` while keeping `reqId` attached to `ctx`. Client disconnects after `202 Accepted` never abort in-flight cloud provisioning, and all outbound CB-Tumblebug sub-requests remain correlated to `reqId`.
- At most 20 async jobs run concurrently (shared across all endpoints above). Beyond that, a request
  returns `503` with a `Retry-After` header instead of `202`: retry after that delay, or retry without
  `Prefer: respond-async` to run synchronously.
- Only the `respond-async` preference token is recognized. Other `Prefer` tokens (e.g. `wait=N`) are
  ignored; there is no time-bounded hybrid mode.

## Related

- [Request Tracking and Sub-Request Progress](core-request-tracking.md) - `subRequests` correlation, `?brief=true`, and `?withTumblebug=true`
- [API Response Policy](../api-response-policy.md) - the underlying status-code and middleware contract
- [Async Support Plan](../plan/async-support-plan-for-recommendation-migration-apis.md) - design background and implementation notes
