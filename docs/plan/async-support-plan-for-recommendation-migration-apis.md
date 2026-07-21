# Plan: Optional Async Responses for Recommendation & Migration APIs

## Overview

Users need async responses for CM-Beetle APIs that take a long time to complete (VM provisioning,
multi-region/multi-CSP recommendation lookups). CM-Beetle already has a working async pattern
(`POST /migration/data`) and a request-tracking infrastructure (`GET /request/{reqId}`). This plan
extends that same pattern to the recommendation and migration-infra APIs, **without breaking existing
synchronous callers**.

Status: All phases implemented — Phase 0 (Data), Phase 1 (migration infra, including DeleteInfra), Phase 2 (infraWithNlb), and Phase 3 (recommendation infra / infraWithDefaults). Extended beyond the original scope to `MigrateNlbs`, `DeleteNlb` (fixed 15s settle wait, [nlb.go:197-199](../../pkg/core/migration/nlb.go#L197-L199)), and `MigrateObjectStorage` (N buckets processed sequentially) after a latency review of Object Storage/NLB/individual-resource APIs — see [Extended Scope](#extended-scope-object-storage--nlb) below. `RecommendVmSpecs`/`RecommendVmOsImages` were reviewed and intentionally left synchronous. `docs/feature-guide/async-responses.md` (63 lines) has been written, completing the plan.

---

## Breaking-Change Policy (Platform Integration in Progress)

A platform-wide integration effort is currently underway, so **the recommendation and migration-infra
APIs (Phases 1-3 below) must have zero breaking change** — any client that never sends the new opt-in
signal must see byte-for-byte identical behavior to today.

**`POST /migration/data` (Data) is explicitly excluded from this constraint.** It is not part of the
current integration effort, so — in the interest of making the _whole_ API surface consistent — it is
deliberately being changed to follow the same `Prefer: respond-async` opt-in convention as every other
endpoint, even though this changes its current (always-202) behavior. See [Phase 0](#target-apis-and-rollout-order).

| Target endpoint                                                     | Default (no `Prefer` header) behavior after the change                              | Verified by                                                                              |
| ------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `POST /migration/ns/{nsId}/infra`, `/infraWithDefaults`             | Identical 201 + `MigrateInfraResponse`                                              | `preferRespondAsync(c) == false` → falls straight into the existing, untouched code path |
| `POST /recommendation/infra`, `/infraWithDefaults`, `/infraWithNlb` | Identical 200 response                                                              | Same                                                                                     |
| `POST /migration/data`                                              | **Changes** — now returns 200 synchronously by default instead of unconditional 202 | Intentional, out of scope for this constraint                                            |

**Implementation discipline required for Phases 1-3** (design safety is not the same as implementation safety):

1. **Insert, never restructure.** The async branch must be added as a pure early-return block right after
   existing input validation. The pre-existing synchronous code must not be reordered, reindented, or
   extracted into a shared function in the same change — a reviewer should be able to confirm safety just
   by checking that the diff is addition-only with zero lines touched in the sync path.
2. **Malformed/absent `Prefer` header must default to sync.** `preferRespondAsync` returns `false` unless
   the exact `respond-async` token is present — already designed this way, re-confirmed here as a hard
   requirement, not just an implementation detail.
3. **Swagger changes are additive and have an in-repo precedent.** Declaring more than one possible
   response for the same endpoint is not new to this codebase: `ListInfra` already declares two distinct
   `@Success 200` responses depending on the `option` query param
   ([migration.go:195-196](../../pkg/api/rest/controller/migration.go#L195-L196)). Adding
   `@Success 202 {object} model.ApiResponse[model.AsyncJobResponse]` alongside the existing `@Success 200/201`
   follows the same established pattern.

---

## Existing Infrastructure (reused, not modified)

| Component                      | Location                                                                                                                        | Role                                                                                                                                          |
| ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `RequestDetails` / `lkvstore`  | [pkg/core/common/request-manager.go](../../pkg/core/common/request-manager.go)                                                  | Persists per-request status (`sync.Map` + file-backed)                                                                                        |
| Status enum                    | `common.RequestStatusHandling / Success / Error` (request-manager.go:74-87)                                                     | Deliberately kept to 3 values — see [Status Values](#2-status-values-no-changes)                                                                |
| `model.AsyncJobResponse`       | [pkg/api/rest/model/async.go](../../pkg/api/rest/model/async.go)                                                                | `{reqId, status, statusUrl}` — the 202 response body                                                                                          |
| `RequestIdAndDetailsIssuer`    | [pkg/api/rest/middlewares/request-id.go](../../pkg/api/rest/middlewares/request-id.go)                                          | Issues/validates `X-Request-Id`, sets initial `Handling` status before the handler runs                                                       |
| `ResponseBodyDump`             | [pkg/api/rest/middlewares/response-dump.go](../../pkg/api/rest/middlewares/response-dump.go)                                    | Finalizes status to `Success`/`Error` after the handler returns — **skips this entirely if the response status is 202** (response-dump.go:47) |
| Reference implementation       | `MigrateData` / `executeMigrationAsync` in [migration-data.go](../../pkg/api/rest/controller/migration-data.go) (lines 119-230) | The only existing async handler; the pattern to replicate                                                                                     |
| `common.UpdateRequestProgress` | request-manager.go:330                                                                                                          | Appends incremental progress entries to `ResponseData` — implemented but currently unused anywhere                                            |

**No changes to middleware _logic_ are required.** HTTP 202 is already a reserved, codebase-wide signal
meaning "async job accepted; a background goroutine will finalize the status." Grep confirms
`StatusAccepted` is used in exactly two places today (the handler that sets it, and the middleware guard
that skips it). The one exception is a one-line, additive change to `RequestIdAndDetailsIssuer`'s CORS
header list — see [CORS Header Exposure](#4-cors-header-exposure-for-preference-applied) below.

---

## Design Decisions

### 1. Sync and async on the same endpoint (no new endpoints, no breaking change)

Confirmed requirement: both response modes must stay available, with minimal added complexity.

- Opt-in via the **`Prefer: respond-async` request header** (RFC 7240), not a query parameter.
  Rationale: sync-vs-async is an HTTP-transport-level processing preference (how the request/response
  cycle itself behaves), not a domain/business-level option — the same category as `Accept` or
  `If-Match`, unlike e.g. `useExisting` (a business option on _what_ the migration does, correctly left
  as a query param — see [naming precedent](#naming-precedent-query-param-vs-header) below). `Prefer` is
  the standard mechanism for exactly this case, used the same way by Microsoft Graph, OData, and Azure
  REST APIs.
- When honored, the server echoes back **`Preference-Applied: respond-async`** on the 202 response, per
  RFC 7240 §3 — a one-line addition that lets clients confirm the async path was actually taken.
- The branch happens **inside the existing handler**, immediately after input validation, before any
  slow work starts. Both branches call the _same_ underlying core function — no duplicated business logic.
- One shared generic helper in `pkg/core/common` removes per-handler boilerplate, regardless of what the
  wrapped core function does or which package it belongs to (`recommendation`, `migration`, `transx` —
  see [Open Questions](#open-questions) for why this also covers `executeMigrationAsync`):

```go
// maxConcurrentAsyncJobs caps background jobs from RunAsync (shared by all async endpoints).
const maxConcurrentAsyncJobs = 20

var asyncJobSemaphore = make(chan struct{}, maxConcurrentAsyncJobs)

// RunAsync runs work in the background, recovers panics, and finalizes the request record.
// Returns false at capacity — caller must respond 503, not spawn anyway.
func RunAsync[T any](reqID string, work func() (T, error)) bool {
    select {
    case asyncJobSemaphore <- struct{}{}:
    default:
        return false
    }
    go func() {
        defer func() { <-asyncJobSemaphore }()
        defer func() {
            if r := recover(); r != nil {
                // Stack trace stays server-side; clients never see it (see note below).
                log.Error().Str("reqId", reqID).Str("stack", string(debug.Stack())).Msgf("panic: %v", r)
                finalizeError(reqID, fmt.Errorf("internal error (panic): %v", r))
            }
        }()
        result, err := work()
        finalizeRequest(reqID, result, err)
    }()
    return true
}

// preferRespondAsync reports whether Prefer includes "respond-async" (RFC 7240).
// Other tokens (e.g. wait=N) are ignored.
func preferRespondAsync(c echo.Context) bool {
    for _, token := range strings.Split(c.Request().Header.Get("Prefer"), ",") {
        if strings.TrimSpace(token) == "respond-async" {
            return true
        }
    }
    return false
}

func RecommendVmInfraCandidates(c echo.Context) error {
    // ... existing input validation, unchanged ...

    if preferRespondAsync(c) {
        reqID := c.Request().Header.Get(echo.HeaderXRequestID)
        started := common.RunAsync(reqID, func() ([]cloudmodel.RecommendedInfra, error) {
            return recommendation.RecommendVmInfraCandidates(csp, region, sourceInfra, limit, minMatchRate)
        })
        if !started {
            c.Response().Header().Set("Retry-After", "5")
            return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
                "Too many async jobs in progress; retry shortly, or retry without Prefer: respond-async"))
        }
        c.Response().Header().Set("Preference-Applied", "respond-async")
        return c.JSON(http.StatusAccepted, model.SuccessResponseWithMessage(
            model.AsyncJobResponse{
                ReqID:     reqID,
                Status:    common.RequestStatusHandling,
                StatusURL: fmt.Sprintf("/beetle/request/%s", reqID),
            },
            "Recommendation started. Use GET /request/{reqId} to check status."))
    }

    // existing synchronous path — unchanged
    recommendedInfraCandidates, err := recommendation.RecommendVmInfraCandidates(csp, region, sourceInfra, limit, minMatchRate)
    ...
}
```

The handler-level async branch is identical regardless of whether the wrapped core function is a single
call (`RecommendVmInfraCandidates`) or a multi-stage orchestration (`CreateVMInfraWithDefaults`) — all
complexity stays inside the closure passed to `RunAsync`, not duplicated per handler.

**Panic messages never reach the client.** `GET /request/{reqId}` returns `common.RequestDetails` verbatim,
including `ErrorResponse` ([api-request.go](../../pkg/api/rest/controller/api-request.go)) — so a raw
`debug.Stack()` in that field would leak internal file paths and call frames to any API caller. The full
stack trace goes to the server log only; `ErrorResponse` gets a short, generic message. This matches how
`middleware.Recover()` already behaves for synchronous handlers ([server.go:110](../../pkg/api/rest/server.go#L110)) —
logs the panic, never returns the stack trace to the client.

**Swagger annotation**: document with
`@Param Prefer header string false "Set to 'respond-async' for async processing (RFC 7240). Only this token is recognized; 'wait=N' and other Prefer tokens are ignored." Enums(respond-async)`,
plus `@Failure 503 {object} model.ApiResponse[any] "Too many concurrent async jobs; retry later or without Prefer: respond-async"`.
Swagger UI renders header params as a normal input field, exactly as it already does for `X-Request-Id`
throughout this codebase, so there's no loss of Swagger UI testability versus a query param.

**Trade-off, accepted deliberately**: one global pool means a burst of slow migration jobs can consume all
20 slots and cause an unrelated, cheap recommendation call to get `503` even though it would finish in
seconds. Per-endpoint or per-job-type pools would avoid this but add real configuration surface for a
problem that hasn't been observed yet — same reasoning as deferring progress reporting
([Design Decision 3](#3-progress-reporting-deferred-not-part-of-this-plan)). Revisit with separate pools
only if this starvation is actually observed in practice.

**Hardcoded constant, deliberately not a config value**: making this configurable via
`config.Beetle` would mean adding a field in four places — `pkg/config/config.go`, `conf/config.yaml`,
`conf/template-config.yaml`, and `conf/template-setup.env` (this repo's env-var binding is explicit per
field via `viper.BindEnv` in `bindEnvironmentVariables()`, not automatic, so a config-based value cannot
skip any of the four) — for a number with no evidence yet that it needs runtime tuning. A Go constant is
simpler and can be changed and redeployed if 20 turns out to be wrong. Revisit as a config value only if
real usage shows a concrete need to tune it without a rebuild.

#### Naming precedent: query param vs. header

There is no general convention of "control flags go in headers." The actual distinction is which layer
the flag belongs to:

- **HTTP-transport-level processing preference** (sync vs. async, response verbosity, conditional
  requests) → header. Standard examples: `Prefer: respond-async` / `Prefer: return=minimal` (Microsoft
  Graph, OData), `Idempotency-Key` (Stripe), `If-Match`.
- **Domain/business-level request option** (dry-run, force, upsert, reuse-existing-resources) → query
  parameter is the common convention. Examples: Kubernetes `?dryRun=All`, common `?force=true` /
  `?upsert=true` patterns.

`useExisting` on `MigrateInfra` (a business option on what the migration does) was correctly left as a
query parameter under this rule — it is not an exception that needs revisiting.

### 2. Status values: no changes

`Handling` / `Success` / `Error` stay as-is — this was a deliberate choice, not an oversight, and it
also keeps the enum aligned with CB-Tumblebug's status values (per existing code comment). New edge
cases introduced by making more APIs async (timeouts, orphaned jobs, deleted-mid-job records) are folded
into `Error` with a descriptive `ErrorResponse` message rather than new status values.

### 3. Progress reporting: deferred, not part of this plan

`common.UpdateRequestProgress` exists but is deliberately **not** wired into any of these APIs now. It
solves a different problem than "get an async response" (mid-flight visibility into a multi-stage job)
and, on inspection, its implementation scope varies unpredictably by function:

- `CreateVMInfraWithDefaults` ([vm-infra.go:94](../../pkg/core/migration/vm-infra.go#L94)) is a single
  Tumblebug call — nothing meaningful to report.
- `CreateInfra` ([vm-infra.go:123](../../pkg/core/migration/vm-infra.go#L123)) has all its stages
  (vNet → SSH key → security groups → infra) inline in one function — reporting would touch only this
  one function.
- `CreateInfraWithExisting` ([vm-infra.go:321-394](../../pkg/core/migration/vm-infra.go#L321-L394))
  delegates each stage to separate child functions (`useOrCreateNetwork`, `useOrCreateSshKey`,
  `useOrCreateSecurityGroup`, each looped per node group) — stage-level reporting only touches the
  parent function, but reporting anything finer (e.g. "reusing" vs. "creating" per resource) requires
  threading a parameter into all three child functions.

Because the actual scope ranges from "one function, one line" to "one parent + three child functions,"
with no immediate requirement driving the need for it, adding it now would be speculative complexity.
Revisit only if there is a concrete request for granular progress visibility.

### 4. CORS header exposure for `Preference-Applied`

`RequestIdAndDetailsIssuer` ([request-id.go:34](../../pkg/api/rest/middlewares/request-id.go#L34))
currently sets `Access-Control-Expose-Headers: X-Request-Id` only. Browser-based JS clients calling
cross-origin cannot read `Preference-Applied` unless it is also listed. This is the one middleware touch
in this whole plan — purely additive (extends a header value, doesn't remove or change existing behavior),
so it doesn't conflict with the "no middleware logic changes" statement in
[Existing Infrastructure](#existing-infrastructure-reused-not-modified):

```go
c.Response().Header().Set("Access-Control-Expose-Headers", "X-Request-Id, Preference-Applied")
```

Not strictly required for functionality — `reqId`/`status` are already present in the JSON response body
(which `data/main.go`'s test-cli already reads from, not the header), so this only matters for
browser-based JS clients that specifically want to read the confirmation header rather than the body.
Lower priority than items 1-3, but simple enough to include in the same change.

---

## Status Update Timing — Verified Sequence

Traced precisely to confirm no new middleware work or race conditions are introduced.

**Common prefix (every request):**

```
1. RequestIdAndDetailsIssuer (before handler)
   → SetRequest(reqID, {Status: Handling, RequestInfo, StartTime: now})
2. Handler runs
3. ResponseBodyDump (same goroutine, right after the handler writes its response)
   → if response status == 202: return   // guard, does nothing further
   → else: GetRequest → set Status Success/Error, EndTime, ResponseData/ErrorResponse → SetRequest
```

**Sync path:** step 3 finalizes the record in the same goroutine as the request — no concurrent writer,
no race.

**Async path:** step 3 is a no-op (202 guard). The background goroutine — running after the request
goroutine has already returned — later does its own `GetRequest → SetRequest`. Because the request
goroutine performs no further writes to that `reqID` after returning 202, this is a sequential handoff
between goroutines, not concurrent access — no write-write race.

### Confirmed safe, and known limitations

| Item                            | Finding                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| ------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `lkvstore` concurrency          | Backed by `sync.Map`; individual `Get`/`Put` calls are atomic, but the `GetRequest → SetRequest` _pair_ is not lock-protected (classic read-modify-write). Safe today because exactly one goroutine writes to a given `reqID` at a time in both sync and async paths — confirmed no goroutines/parallelism exist yet in `pkg/core/recommendation` or `pkg/core/migration`. Not a concern for this plan since progress reporting is out of scope (see [Progress Reporting](#3-progress-reporting-deferred-not-part-of-this-plan)); would need re-checking if that is ever added with parallel node-group processing. |
| Panic safety                    | `executeMigrationAsync` (existing reference implementation) has **no `recover()`** — a panic in the background goroutine currently crashes the process. **Addressed**: `RunAsync` has `recover()` built in, and `executeMigrationAsync` is refactored to use `RunAsync` in Phase 0 (see [Open Questions](#open-questions)), closing this gap for all endpoints at once.                                                                                                                                                                                                                                                                                                                                                    |
| Record deleted mid-job          | If `DELETE /request/{reqId}` is called while a background job is still running, the job's final `GetRequest` returns `!ok`; it logs an error and silently drops the result. Acceptable given today's usage; no change proposed.                                                                                                                                                                                                                                                                                                                                                                                     |
| Orphaned `Handling` records     | `CleanupOldRequests` intentionally never deletes `Handling` records regardless of age. A crash mid-job (or a normal shutdown, see [Open Questions](#open-questions)) leaves a permanent stuck record even with `recover()` in place, since `recover()` only handles in-process panics. No new status value proposed (see [Status Values](#2-status-values-no-changes)).                                                                                                                                                                                                                                                                                                           |
| Unbounded concurrent async jobs | Unlike sync calls (naturally throttled by clients holding a connection open), async lets a client fire many long-running jobs without waiting. No worker-pool/semaphore pattern existed in this codebase to bound this. **Addressed**: `RunAsync`'s `asyncJobSemaphore` (see [Design Decision 1](#1-sync-and-async-on-the-same-endpoint-no-new-endpoints-no-breaking-change)) caps concurrent background jobs at 20, rejecting new ones with `503` rather than queuing unboundedly.                                                                                                                                 |

**Conclusion:** middleware _logic_ requires zero changes; the only middleware touch is the additive CORS
header list update in [Design Decision 4](#4-cors-header-exposure-for-preference-applied). Handlers only
need to follow the existing "return 202 → middleware backs off" contract.

---

## Target APIs and Rollout Order

| Phase | Endpoint(s)                                                                      | Why this order                                                                                                                                                                                                                                                                                                                                                                                                             | Breaking change allowed?                       |
| ----- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------- |
| 0     | `POST /migration/data`                                                           | Consistency retrofit: bring Data onto the same opt-in `Prefer` convention as every other endpoint (see [Breaking-Change Policy](#breaking-change-policy-platform-integration-in-progress))                                                                                                                                                                                                                                 | **Yes** — not part of the platform integration |
| 1     | `POST /migration/ns/{nsId}/infra`, `POST /migration/ns/{nsId}/infraWithDefaults`, `DELETE /migration/ns/{nsId}/infra/{infraId}` | Real provisioning latency from synchronous Tumblebug/CSP calls (`CreateVNet`, `CreateSshKey`, `CreateSecurityGroup`, `CreateInfra`/`CreateInfraDynamic` in [vm-infra.go](../../pkg/core/migration/vm-infra.go)); pattern already proven via `MigrateData`. `DeleteVMInfra` (same file) was added to this phase too — it has four fixed `time.Sleep(3s)` settle-waits plus a vNet-deletion retry loop (up to 10 x 10s), a minimum ~12s and worst case 100s+, a stronger async case than the creation path. | No                                             |
| 2     | `POST /recommendation/infraWithNlb`                                              | Per-node-group loop over spec/image/security-group lookups ([infra-with-nlb.go:306](../../pkg/core/recommendation/infra-with-nlb.go))                                                                                                                                                                                                                                                                                      | No                                             |
| 3     | `POST /recommendation/infra`, `POST /recommendation/infraWithDefaults`           | Lower latency than Phase 2, lower urgency                                                                                                                                                                                                                                                                                                                                                                                  | No                                             |

---

## Extended Scope: Object Storage / NLB

After the four phases above shipped, a latency review covered every remaining recommendation/migration
API (object storage, NLB, individual resources like VNet/SG/SSH key/spec/image) to check whether any of
them deserved the same treatment. Findings and decisions:

| API | Latency source | Decision |
|---|---|---|
| `DELETE /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}` (`DeleteNlb`) | Fixed 15s settle wait after successful deletion ([nlb.go:197-199](../../pkg/core/migration/nlb.go#L197-L199)), same rationale as `DeleteInfra` | **Added** |
| `POST /migration/middleware/ns/{nsId}/infra/{infraId}/nlb` (`MigrateNlbs`) | N NLBs processed sequentially, each with a lookup + create call ([nlb.go:48-95](../../pkg/core/migration/nlb.go#L48-L95)) | **Added** |
| `POST /migration/middleware/ns/{nsId}/objectStorage` (`MigrateObjectStorage`) | N buckets created sequentially, up to 2 extra calls per bucket for versioning/CORS ([object-storage.go:64-169](../../pkg/core/migration/object-storage.go#L64-L169)) | **Added** |
| `RecommendVmSpecs` / `RecommendVmOsImages` | Per-node loop, each with up to 5 retry calls widening the search range ([resource-vm-spec.go:148-263](../../pkg/core/recommendation/resource-vm-spec.go#L148-L263)) | **Left synchronous** — individual calls are fast with no `time.Sleep`; latency only shows up with many nodes, a weaker case than the others |
| Object Storage/NLB single-item GET/DELETE, `RecommendVNet`/`RecommendSecurityGroups` (local computation only), VNet/SSHKey/SecurityGroup proxy CRUD ([migration-resource.go](../../pkg/api/rest/controller/migration-resource.go)) | Single external call or pure local computation | **Left synchronous** — no meaningful latency to hide |

`DeleteNlb`'s core function (`migration.DeleteNlb`) returns only `error`, like `transx.Transfer` in
[Phase 0](#target-apis-and-rollout-order) — its closure follows the same shape, returning a small
`map[string]any` confirmation message as `T` since there is no natural result value.

---

## Work Items per Phase

1. Add `preferRespondAsync(c)` check + `common.RunAsync(reqID, func() (T, error) {...})` branch to the
   target handler(s), reusing `model.AsyncJobResponse`. Handle the `false` return (capacity reached) with
   a `503` response, per [Design Decision 1](#1-sync-and-async-on-the-same-endpoint-no-new-endpoints-no-breaking-change).
2. Add the generic `common.RunAsync[T any]` helper — including the `maxConcurrentAsyncJobs` constant and
   `asyncJobSemaphore` — once, in Phase 0, and refactor `executeMigrationAsync` to use it in the same PR —
   see [Open Questions](#open-questions) (resolved: not optional, done in Phase 0). Kept as a Go constant,
   not a config value — see the trade-off note in
   [Design Decision 1](#1-sync-and-async-on-the-same-endpoint-no-new-endpoints-no-breaking-change).
3. Update Swagger annotations: add an `[Async]` usage block matching `MigrateData`'s current
   `@Description [How to Use]` section, document the `Prefer` header (noting `wait=N` is not supported),
   add `202` and `503` to `@Success`/`@Failure` alongside the existing success response (multiple
   `@Success` entries on one endpoint already has precedent — see
   [Breaking-Change Policy](#breaking-change-policy-platform-integration-in-progress) item 3).
4. **Phase 0 only**: add a new synchronous code path to `MigrateData` (today it has none — it is
   unconditionally async) that calls the same underlying migration logic inline and returns `200` directly
   when `Prefer: respond-async` is absent.
5. **Phase 0 only**: update `cmd/test-cli/data/main.go`'s migration call to add
   `.SetHeader("Prefer", "respond-async")` — required, not optional, since after Phase 0 the default
   response becomes `200` and this test's existing `resp.StatusCode() != 202` check (data/main.go:893)
   would otherwise start failing.
6. **Phase 0 only**: add `Preference-Applied` to `Access-Control-Expose-Headers` in
   `RequestIdAndDetailsIssuer` ([request-id.go:34](../../pkg/api/rest/middlewares/request-id.go#L34)) —
   see [Design Decision 4](#4-cors-header-exposure-for-preference-applied). The only middleware change in
   this plan; do it once, in Phase 0, since it applies globally regardless of phase.
7. Manual test: verify `Prefer: respond-async` and default (sync) against a running Tumblebug backend.
   For the `503`-at-capacity path: no automated test needed given no `_test.go` infra exists for these
   packages today ([Test CLI Impact](#test-cli-impact)) — a quick manual check (fire 20+ concurrent
   async requests with a script) is sufficient to confirm the semaphore rejects correctly.
8. No changes to `ResponseBodyDump` or the `RequestStatus*` enum. `RequestIdAndDetailsIssuer` gets only
   the one additive CORS line from item 6 above — no logic changes.
9. Documentation and guardrails — see [Documentation & Guardrails](#documentation--guardrails-contributor-and-ai-agent-facing) below.
10. **After all phases ship**: write `docs/feature-guide/async-responses.md` — see
    [Feature Guide](#feature-guide-post-implementation-user-facing) below. Keep it short (~50-100 lines,
    table/example-driven like `naming-nameseed.md`), not a long narrative like the existing
    migration feature guides. Include the `503`-at-capacity behavior and the `wait=N` non-support note.

---

## Documentation & Guardrails (Contributor and AI-Agent Facing)

Today the async convention exists in exactly one place (`MigrateData`) and is undocumented — a new
contributor (human or AI coding agent) can only learn it by reading `migration-data.go` source. Once this
plan lands, HTTP `202` becomes a codebase-wide reserved signal (see
[Existing Infrastructure](#existing-infrastructure-reused-not-modified)): any handler that returns `202`
for an unrelated reason will silently cause `ResponseBodyDump` to skip finalizing that request's status
record, leaving it permanently stuck at `Handling` — a trap that isn't obvious from reading the middleware
code alone. Three additions close this gap, so the rule is visible wherever someone is likely to be
working, not just in this plan document (which is a point-in-time proposal, not a living reference):

1. **`docs/api-response-policy.md`** — add an "Async Responses" section documenting the
   `Prefer: respond-async` → `202` → poll `GET /request/{reqId}` contract, the `503`-at-capacity behavior
   (20 concurrent jobs, see [Design Decision 1](#1-sync-and-async-on-the-same-endpoint-no-new-endpoints-no-breaking-change)),
   and that `wait=N` and other `Prefer` tokens are not supported. This file currently has no mention of
   `202` at all; it is the natural, permanent home for this convention, unlike the plan doc.
2. **Code comments** at the two points where the rule actually matters:
   - The `if c.Response().Status == http.StatusAccepted` guard in
     [response-dump.go](../../pkg/api/rest/middlewares/response-dump.go) — explain _why_ it skips, and that
     `202` must never be returned for a non-async reason.
   - The `common.RunAsync[T any]` definition itself — a one-line doc comment pointing back to the
     `docs/api-response-policy.md` section, so a contributor who finds the helper first (not the plan) still
     gets the contract.
3. **`CLAUDE.md`** (repo root — does not currently exist, to be created) — add an explicit,
   agent-readable instruction so AI coding agents pick up the same guardrail automatically in future
   sessions, not just human contributors reading docs:

   > HTTP `202 Accepted` is reserved repo-wide to mean "async job accepted; a background goroutine will
   > finalize the `/request/{reqId}` status record." Never return `202` from a handler for any other
   > reason — `ResponseBodyDump` middleware (`pkg/api/rest/middlewares/response-dump.go`) skips status
   > finalization for any `202` response, so doing so leaves that request's tracking record permanently
   > stuck at `Handling`. When adding async support to a new endpoint, use `common.RunAsync` and the
   > `Prefer: respond-async` header convention documented in `docs/api-response-policy.md`.

---

## Feature Guide (Post-Implementation, User-Facing)

The three items above ([Documentation & Guardrails](#documentation--guardrails-contributor-and-ai-agent-facing))
target contributors and AI agents working _on_ CM-Beetle. A separate, user-facing doc is also needed for
people integrating _with_ CM-Beetle's API — this belongs in `docs/feature-guide/`, added as a work item
once implementation is complete (not before, so it reflects the shipped behavior, not the proposal).

**Existing docs in this directory vary widely in length** — `vm-infrastructure.md` (26 lines) and
`naming-nameseed.md` (96 lines) are short, table- and example-driven; `data-migration-feature-guide.md`
(461 lines) and `object-storage-migration-feature-guide.md` (1168 lines) are long narrative walkthroughs.
**Follow the short style** — a long document works against its own purpose here, since the async
convention itself is simple (one header, one status code, one polling endpoint) and a long writeup would
make that simplicity harder to see, not easier.

Target: a new `docs/feature-guide/async-responses.md`, similarly scoped to `naming-nameseed.md` (roughly
50-100 lines), covering only:

1. **What it is** — one paragraph: opt-in via `Prefer: respond-async`, default stays synchronous.
2. **How to use it** — the request/response shapes, modeled on `naming-nameseed.md`'s terse
   request/response code blocks rather than prose:

   ```
   POST /beetle/recommendation/infra
   Prefer: respond-async
   → 202 Accepted, Preference-Applied: respond-async
     { "reqId": "...", "status": "Handling", "statusUrl": "/beetle/request/{reqId}" }

   GET /beetle/request/{reqId}
   → { "status": "Success" | "Error" | "Handling", ... }
   ```

3. **Which endpoints support it** — one table listing the Phase 0-3 endpoints from this plan.
4. **Status values** — one small table (`Handling` / `Success` / `Error`), no more than that (matching
   [Status Values](#2-status-values-no-changes) — do not expand this in the user-facing doc either).
5. **Limits** — one or two lines: at most 20 concurrent async jobs — retry after the `Retry-After`
   seconds on `503` — and only the `respond-async` preference token is recognized (`wait=N` is ignored,
   no time-bounded hybrid mode).
6. A short **Related** section linking to `docs/api-response-policy.md` for the underlying contract and
   to this plan doc for implementation background, following the same "Related" convention
   `naming-nameseed.md` already uses.

Explicitly out of scope for this guide: progress reporting (not implemented, see
[Progress Reporting](#3-progress-reporting-deferred-not-part-of-this-plan)), and internal implementation
details (`common.RunAsync`, middleware internals) — those belong in code comments and
`docs/api-response-policy.md`, not in a user-facing feature guide.

---

## Test CLI Impact

`cmd/test-cli/{data,infra,infra-with-nlb}/main.go` are independent `package main` tools (no shared
library between them — each is copied and adapted, per the existing `nlb-plan-for-recommendation-and-migration.md`
convention), so client-side impact was assessed directly against their current call sites.

### Existing sync tests: zero changes required

`infra/main.go`'s `runMigrationTest` ([main.go:955-964](../../cmd/test-cli/infra/main.go#L955-L964)) calls
`common.ExecuteHttpRequest` with no headers set. Since the new async path is opt-in via `Prefer`, and this
call never sends that header, the server keeps returning its current synchronous response — the existing
9–13 step test suites in `infra`/`infra-with-nlb` pass unmodified. This is a concrete, code-level
confirmation of the "no breaking change" requirement.

### Decision: sync-only test coverage for everything except Data

`cmd/test-cli/data/main.go` already implements the full async client flow for `MigrateData`
(lines 870-989) and remains the one place async is exercised end-to-end. **`infra` and
`infra-with-nlb` intentionally stay sync-only** — no new `runMigrationTestAsync`-style functions,
no `Prefer` header, no polling logic added to those tools, even though Phase 1/2 endpoints now
support async. This was a deliberate scope decision, not an oversight: the async code path in each
handler is a thin, identical branch around the same core function the sync test already exercises
(see [Design Decision 1](#1-sync-and-async-on-the-same-endpoint-no-new-endpoints-no-breaking-change)),
so the sync tests already cover the business logic; only `data/main.go` needed the one-line `Prefer`
header addition from Phase 0's work items, since `MigrateData` has no sync path to fall back to
without it.

**If async test coverage for `infra`/`infra-with-nlb` is wanted later**, reuse the pattern already
proven in `data/main.go` — the reqId-extraction, 202-branch, and poll-loop code, built with raw resty
(**not** `common.ExecuteHttpRequest`, which is a shared internal helper also used for Beetle→Tumblebug
calls, [request-manager.go:418](../../pkg/core/common/request-manager.go#L418)).

---

## Open Questions

- ~~Should `executeMigrationAsync` be refactored to use the new shared helpers in the same PR, or left
  as-is and only the new handlers use them?~~ **Resolved**: refactor it in the same PR as Phase 0, using
  `common.RunAsync`. `transx.Transfer(req)` returning only `error` (no result value) is not a blocker —
  the closure just builds the success payload itself and returns it as `T`:

  ```go
  if preferRespondAsync(c) {
      reqID := c.Request().Header.Get(echo.HeaderXRequestID)
      started := common.RunAsync(reqID, func() (map[string]any, error) {
          start := time.Now()
          if err := transx.Transfer(*req); err != nil {
              return nil, fmt.Errorf("data migration failed: %w (%s)", err, time.Since(start).Round(time.Millisecond))
          }
          elapsed := time.Since(start).Round(time.Millisecond)
          return map[string]any{
              "message":     fmt.Sprintf("Data migrated successfully (%s)", elapsed),
              "elapsedTime": elapsed.String(),
          }, nil
      })
      if !started {
          c.Response().Header().Set("Retry-After", "5")
          return c.JSON(http.StatusServiceUnavailable, model.SimpleErrorResponse(
              "Too many async jobs in progress; retry shortly, or retry without Prefer: respond-async"))
      }
      c.Response().Header().Set("Preference-Applied", "respond-async")
      return c.JSON(http.StatusAccepted, model.SuccessResponseWithMessage(
          model.AsyncJobResponse{ReqID: reqID, Status: common.RequestStatusHandling, StatusURL: fmt.Sprintf("/beetle/request/%s", reqID)},
          "Migration started. Use GET /request/{reqId} to check status."))
  }
  ```

  The custom `"Data migration failed: ... (163ms)"` error format and the `message`/`elapsedTime` success
  shape are preserved exactly — that formatting is domain-specific and belongs in the closure, not in
  `RunAsync` itself, the same way every other endpoint's closure owns its own result shape. This is folded
  into Phase 0's work items above since Data is already being touched there for the sync-path addition.

- **Deferred, explicitly out of scope for this plan**: graceful shutdown does not drain in-flight async
  jobs. `server.go`'s existing `e.Shutdown(ctx)` ([server.go:467-490](../../pkg/api/rest/server.go#L467-L490))
  only waits for in-progress HTTP request/response cycles — an async handler's HTTP cycle ends the moment
  it returns `202`, so Echo's shutdown has no visibility into the detached `RunAsync` goroutine still
  running afterward. This means a normal rolling restart (not just a crash) can orphan an in-flight async
  job at "Handling" — a more common trigger than the crash/panic scenario the existing
  [Orphaned `Handling` records](#confirmed-safe-and-known-limitations) limitation was written around, though
  the same underlying limitation. Explicitly deferred: no draining, waiting, or shutdown-time logging added
  in this plan. Revisit only if this proves disruptive in practice.
