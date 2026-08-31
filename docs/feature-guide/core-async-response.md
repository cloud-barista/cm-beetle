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

Poll `GET /request/{reqId}` until `status` is `Success` or `Error`. Omit the `Prefer` header (or send
any other value) to get the original synchronous response — nothing changes for existing callers.

## Supported Endpoints

| Endpoint | Notes |
| --- | --- |
| `POST /migration/data` | Sync by default (breaking change from earlier versions, which were always async) |
| `POST /migration/ns/{nsId}/infra` | |
| `POST /migration/ns/{nsId}/infraWithDefaults` | |
| `DELETE /migration/ns/{nsId}/infra/{infraId}` | Worth using async — deletion includes several settle waits |
| `POST /recommendation/infraWithNlb` | |
| `POST /recommendation/infra` | |
| `POST /recommendation/infraWithDefaults` | |
| `POST /migration/middleware/ns/{nsId}/infra/{infraId}/nlb` | |
| `DELETE /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}` | Worth using async — includes a fixed 15s settle wait |
| `POST /migration/middleware/ns/{nsId}/objectStorage` | |

## Status Values

| Status | Meaning |
| --- | --- |
| `Handling` | Still running |
| `Success` | Completed; result in `responseData` |
| `Error` | Failed; message in `errorResponse` |

## Limits

- At most 20 async jobs run concurrently (shared across all endpoints above). Beyond that, a request
  returns `503` with a `Retry-After` header instead of `202` — retry after that delay, or retry without
  `Prefer: respond-async` to run synchronously.
- Only the `respond-async` preference token is recognized. Other `Prefer` tokens (e.g. `wait=N`) are
  ignored — there is no time-bounded hybrid mode.

## Related

- [API Response Policy](../api-response-policy.md) — the underlying status-code and middleware contract
- [Async Support Plan](../plan/async-support-plan-for-recommendation-migration-apis.md) — design background and implementation notes
