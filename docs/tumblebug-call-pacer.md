# Tumblebug Call Pacer

CB-Tumblebug rate-limits some of its read APIs. The pacer keeps CM-Beetle under that limit by
spacing out its own outbound calls, so Beetle rarely triggers the limit in the first place.

This covers the Beetle side. For the generic mechanism it builds on, see
[pkg/ratelimit/README.md](../pkg/ratelimit/README.md).

## The limit being satisfied

| Route                          | Limit            | Applies per |
| ------------------------------ | ---------------- | ----------- |
| All routes                     | 50 req/s         | Source IP   |
| `GET /ns/:nsId/infra`          | 2 req/s, burst 2 | Source IP   |
| `GET /ns/:nsId/infra/:infraId` | 2 req/s, burst 2 | Source IP   |

- **The limit applies per source IP**, not per namespace or infrastructure — Echo's default
  extractor is `ctx.RealIP()`, so one Beetle process shares one budget.
- **Each route has its own limiter store**, so the two 2 req/s budgets are independent.
- **Tumblebug's `429` carries no `Retry-After`**, so a client cannot learn the advised wait.

Verified against cb-tumblebug v0.12.27 (`src/interface/rest/server/server.go`); `go.mod` pins
v0.12.25. Re-check after an upgrade.

## Two layers

**Pacing** keeps Beetle below the limit: one call every 625 ms (~1.6 req/s).

**A bounded retry** recovers from the `429`s pacing cannot prevent — Beetle may not be the only
client behind its IP, and the global 50 req/s bucket is shared across all routes. Keep both: the
retry is what made removing an earlier adaptive-slowdown design safe.

## Where the code lives

| File                                                                                        | Role                                                                |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| [pkg/ratelimit/pacer.go](../pkg/ratelimit/pacer.go)                                         | Generic `Pacer`. Knows nothing about Tumblebug.                     |
| [pkg/ratelimit/errors.go](../pkg/ratelimit/errors.go)                                       | `ErrLimited`, `RetryAfter`, `RetryAfterSeconds`                     |
| [pkg/client/tumblebug/call-pacer.go](../pkg/client/tumblebug/call-pacer.go)                 | Tumblebug wiring: constants, singleton, wait budget, `Session.pace` |
| [pkg/client/tumblebug/client.go](../pkg/client/tumblebug/client.go)                         | resty retry configuration for `429`                                 |
| [pkg/client/tumblebug/infra-provisioning.go](../pkg/client/tumblebug/infra-provisioning.go) | The paced call sites                                                |
| [pkg/config/config.go](../pkg/config/config.go)                                             | `RetrievalConfig`: the two operator knobs                           |
| [pkg/api/rest/controller/migration.go](../pkg/api/rest/controller/migration.go)             | Translates `ErrLimited` into `503` + `Retry-After`                  |

## Constants

Defined in [call-pacer.go](../pkg/client/tumblebug/call-pacer.go).

| Constant                               | Value     | Why                                                                        |
| -------------------------------------- | --------- | -------------------------------------------------------------------------- |
| `defaultTumblebugRequestsPerSec`       | 2.0       | Tumblebug's own limit; the one value to change if Tumblebug changes it     |
| `pacingHeadroom`                       | 0.8       | Targets 1.6 req/s (one per 625 ms), leaving margin under the limit         |
| `min` / `maxTumblebugRequestsPerSec`   | 0.1 / 100 | Bounds on the configured rate                                              |
| Burst (in `NewPacer`)                  | 1         | Stops an idle period banking tokens that then fire back-to-back            |
| `defaultPacingWait`                    | 8 s       | Budget when the caller sets no deadline                                    |
| `maxPacingWait`                        | 120 s     | Cap on the configured wait                                                 |
| `retryOn429Count` / `Wait` / `MaxWait` | 3/1 s/5 s | Honors `Retry-After` when present; Tumblebug sends none, so spacing is 1 s |

Only `defaultTumblebugRequestsPerSec` states someone else's constraint, and it is the one an
operator can override. Beetle uses burst 1 where Tumblebug allows 2, deliberately conservative.

## How a paced call runs

```text
ReadInfra()
  └─ s.pace("ReadInfra ns=… infra=…")      ← the string is only a log label, not a key
       ├─ ctx := s.req.Context()
       ├─ no deadline? wrap it with pacingWaitBudget()
       └─ getPacer().Wait(ctx)
            ├─ reservation := limiter.Reserve()    ← the slot is claimed now
            ├─ delay <= 0             → proceed immediately
            ├─ delay > ctx deadline   → Cancel(), return *ErrLimited{RetryAfter: delay}
            └─ otherwise              → wait on a timer, or Cancel() and return ctx.Err()
  └─ HTTP GET, retried up to 3 times on 429
```

`Pacer.Wait` reserves a slot up front and then decides. Preserve two behaviors if you modify it:

- **Fail fast.** When the reserved slot lands after the deadline, return immediately instead of
  blocking for the whole budget and failing anyway.
- **Cancel the reservation** on every path that gives up. This returns the slot to the bucket so
  the next caller's estimate stays accurate.

## The wait budget

A paced call waits at most one budget for its slot. The budget comes from the first source that
applies:

1. **The caller's own deadline**, set with `Session.SetContext`. Today only `DeleteInfra` does
   this, using 30 s because deletion tolerates a longer wait than an interactive read.
2. **`pacingWaitBudget()`** — `config.Tumblebug.Retrieval.MaxWaitSec`, capped at 120 s.
3. **`defaultPacingWait`** (8 s) when the configured value is unset or not positive.

> **Gotcha:** the context you pass to `SetContext` is the same context resty uses for the HTTP
> request. A deadline therefore covers the pacing wait **and** the request itself, not the
> pacing wait alone. Setting a tight deadline on a slow endpoint produces a transport timeout,
> not an `ErrLimited`.

### There is no queue: how waiting actually works

This is the part that surprises people, so it is worth spelling out.

**There is no queue anywhere in this design.** No list of waiting callers, no counter, no worker
draining a backlog. What exists is a single timestamp inside the token bucket marking when the
next call may go out.

Think of a bakery's ticket machine. You take a numbered ticket that says _"your turn: 3:05:12"_,
and then you go sit down. Nobody stands in line, and the shop keeps no record of who is waiting —
your ticket alone tells you when to come back. Two separate things are happening:

| Concern               | Mechanism                                                                                     | Where                    |
| --------------------- | --------------------------------------------------------------------------------------------- | ------------------------ |
| **Who goes when**     | `Reserve()` advances the bucket's timestamp by one interval and hands the caller that instant | `golang.org/x/time/rate` |
| **Whether to bother** | Compare that instant against the caller's own `ctx` deadline                                  | `Pacer.Wait`             |

So the ordering comes from timestamp arithmetic, and the **deadline is the admission test** — it
decides whether a caller accepts the slot it was handed or gives it back. Each caller then sleeps
on its own `time.Timer`, independently, in its own goroutine.

Three consequences follow, and they are the whole reason the design works:

1. **Order is fixed at arrival, not at execution.** `Reserve()` is called the moment the request
   arrives, so slots are handed out first-come-first-served even though nothing tracks the order.
2. **Rejection is instant, never a timeout.** A caller whose slot falls past its deadline returns
   in microseconds. It does **not** block for the full budget and then fail.
3. **Rejected callers cost nothing.** `Cancel()` returns the slot to the bucket, so giving up
   does not push anyone else further back.

Because admission is bounded by _time_ rather than by _depth_, the N-th simultaneous caller faces
`(N-1) × 625 ms` and is admitted only while that fits its own budget:

| Wait budget             | Concurrent waiters admitted |
| ----------------------- | --------------------------- |
| 8 s (default)           | 13                          |
| 30 s (`DeleteInfra`)    | 49                          |
| 120 s (`maxPacingWait`) | 193                         |

Each caller is measured against its own deadline, so a long-budget call can be admitted at the
very moment a short-budget one is rejected. Patient callers automatically inherit the slots
impatient ones decline.

This replaced an earlier fixed `maxQueueSize`. A queue depth means a different wait time whenever
the interval changes; a time budget means what callers actually care about.

#### Worked example: a bulk delete during normal browsing

Twenty `DeleteInfra` requests arrive at once. Each one's paced `ReadInfra` carries a 30 s budget,
so all twenty are admitted, starting 625 ms apart — the last begins at ~11.9 s. Note that only
the _read_ is staggered: each delete releases its slot immediately afterward and then runs its
real work (the `DELETE`, settle waits, vNet retries) fully in parallel with the others.

Now a `GET .../infra` arrives while those reads are pending. It carries only an 8 s budget, so it
is quoted a slot past its deadline and gets an instant `503` — not a hang, but not data either.
It recovers once the backlog falls back under 13:

$$t_{\text{blocked}} = (N - 12.8) \times 625\,\text{ms}$$

For N = 20 that is ~4.5 s; for a full 49 it is ~22.6 s. This is a deliberate trade-off, not a
bug: deletes were given 30 s precisely because they tolerate waiting, and the cost is that a
large teardown can briefly crowd out interactive reads. If that ever becomes a real problem, the
fix is to stop sharing one bucket — give deletes a separate, lower-priority pacer.

## What a caller sees on failure

`pace` returns `*ratelimit.ErrLimited`, which propagates unwrapped to the controller:

```go
if retryAfter, ok := ratelimit.RetryAfter(err); ok {
    c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
    return c.JSON(http.StatusServiceUnavailable, ...)
}
```

The status is **503, not 429**, because `pace` returns before any HTTP request is made —
Tumblebug is never contacted. This is Beetle's own admission control, and the API client did
nothing wrong. See [API Response Policy](api-response-policy.md).

## Coverage

**Paced:** `ReadAllInfra`, `ReadInfra`, `ReadInfraIDs`.

**Not paced, and correctly so:** `CreateInfra`, `CreateInfraDynamic`, `DeleteInfra`,
`InfraRecommendSpec`, `RemoteCommandToInfra`. Those routes carry no Tumblebug limiter.

**Known gap:** `ReadInfraAccessInfo` calls `GET /ns/:nsId/infra/:infraId`, the same rate-limited
route as `ReadInfra`, but does not call `pace`. The test CLIs under `cmd/test-cli/` already call
it, so this gap is live rather than theoretical. Add `s.pace(...)` to it.

## Operating it

Two knobs, in `conf/config.yaml` under `tumblebug.retrieval` or as environment variables:

| Key                | Environment variable                          | Default | Meaning                                       |
| ------------------ | --------------------------------------------- | ------- | --------------------------------------------- |
| `requests_per_sec` | `BEETLE_TUMBLEBUG_RETRIEVAL_REQUESTS_PER_SEC` | 2       | Tumblebug's limit, in Tumblebug's own unit    |
| `max_wait_sec`     | `BEETLE_TUMBLEBUG_RETRIEVAL_MAX_WAIT_SEC`     | 8       | How long a call waits for a slot before a 503 |

The group is named `retrieval`, not `ratelimit`, because it describes the tight limit on
Tumblebug's retrieval APIs — **not** its 50 req/s global limit. Setting 50 here would pace Beetle
at 40 req/s into a 2 req/s route and produce constant 429s.

State Tumblebug's **actual** limit and do not pre-subtract a safety margin: Beetle applies
`pacingHeadroom` itself, so 2 yields a 1.6 req/s pace. The knob exists so a Tumblebug policy
change needs a config edit rather than a Beetle release.

Log lines to look for:

| Message                                                                                   | Level | Meaning                                                        |
| ----------------------------------------------------------------------------------------- | ----- | -------------------------------------------------------------- |
| `Tumblebug call pacer initialized (Tumblebug limit: 2 req/s, paced at 1.6 req/s / 625ms)` | debug | Emitted once, on first paced call; echoes the effective config |
| `Tumblebug rate limit … below the … minimum; using …`                                     | warn  | The configured rate was out of range and clamped               |
| `Paced Tumblebug call (…): waited 1.2s`                                                   | debug | A call waited for its slot; logged only above 10 ms            |
| `Tumblebug call pacer refused a slot (…)`                                                 | warn  | The wait budget ran out                                        |
| `Pacing budget exhausted: ReadInfra (…)`                                                  | warn  | Same event, logged at the call site                            |

The initialization line is the quickest way to confirm an operator's config took effect.
Frequent warnings mean Beetle's own concurrency is too high or `max_wait_sec` is too low. They do
**not** mean Tumblebug is rejecting requests.

## Maintenance notes

- **If Tumblebug changes its limit,** set `tumblebug.retrieval.requests_per_sec` to the new value.
  The interval and the logged pace derive from it; no code change is needed.
- **Do not give the pacer a per-namespace or per-infrastructure key.** Tumblebug applies its
  limit per source IP, so N separate pacers would each run at 1.6 req/s into one shared budget.
- **When adding a paced method,** call `s.pace(...)` before building the request, and say so in
  the method's doc comment.
- **Testing:** `getPacer` is a `sync.OnceValue` singleton and cannot be reset between tests. Test
  `ratelimit.Pacer` directly rather than through `tbclient`.
- **Horizontal scaling caveat:** the design assumes one Beetle process equals one source IP. If
  several replicas share one egress IP through NAT or a gateway, Tumblebug sees a single identity
  aggregating all of their traffic, and the total can exceed 2 req/s even though every replica
  paces itself correctly. In that topology, set `requests_per_sec` to Tumblebug's limit **divided
  by** the replica count, so each replica paces proportionally slower, or move pacing to a shared
  component.
