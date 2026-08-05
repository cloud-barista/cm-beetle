# ratelimit

Small, reusable rate-limiting primitives built on [`golang.org/x/time/rate`](https://pkg.go.dev/golang.org/x/time/rate).

This package knows nothing about HTTP, Tumblebug, or Beetle. It provides the two ways a
rate limit can be enforced; callers wire them up and translate the results into their own responses.

| Type       | Strategy                         | Question it answers                                 | Typical use                                     |
| ---------- | -------------------------------- | --------------------------------------------------- | ----------------------------------------------- |
| `Pacer`    | Shaping — **delay** the excess   | "May _anyone_ call this shared resource right now?" | Stay under a downstream API's global rate limit |
| `Cooldown` | Policing — **reject** the excess | "May this _specific key_ act again yet?"            | Stop a user from polling one resource too often |

Rule of thumb: shape calls **you** make to someone else (you control the pace, so waiting
is better than failing); police calls **others** make to you (you can't slow them down, so
refuse and tell them when to come back).

## Where each one sits

The two guard opposite edges of a process and are used independently — separate instances,
no shared state, separate configuration. Using one never implies the other. Beetle happens
to use both, one at each edge:

```text
               inbound                             outbound
  clients ---> [ Cooldown ] ---> Beetle logic ---> [ Pacer ] ---> CB-Tumblebug
               reject a repeat                     delay until    (2 req/s)
               that came too soon                  a slot frees up
```

- **`Cooldown` faces the callers.** Beetle cannot slow a client down, so an early repeat is
  refused outright (HTTP `429` + `Retry-After`). It protects Beetle and the infrastructure
  behind it from a client polling harder than intended.
- **`Pacer` faces the subsystems Beetle calls.** Here Beetle is the one being restrained, and
  it controls its own send rate, so a call waits its turn instead of failing. It fails
  (HTTP `503` + `Retry-After`) only when the wait budget runs out before a slot opens.

## Pacer

`Pacer` shapes traffic: it spreads concurrent callers over time so a shared downstream
resource is never called faster than a fixed interval. Callers are admitted one at a time,
in arrival order.

```text
[callers] -> [token bucket: 1 call per interval] -> [resource]
```

```go
pacer := ratelimit.NewPacer(600 * time.Millisecond) // ~1.67 calls/s

// Wait blocks until this caller's slot arrives, bounded by the context.
ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
defer cancel()

if err := pacer.Wait(ctx); err != nil {
    return err // *ErrLimited if the slot lands after the deadline
}
resp, err := callDownstream()
```

`Wait` reserves a slot up front, so waits are bounded and fair. If the reserved slot would
land after the context deadline, it fails **immediately** with `*ErrLimited` and releases
the reservation, instead of holding the caller for the full budget only to time out.

The interval is fixed. Transient rate limiting from the downstream resource is better
handled by retrying the affected request (for example, an HTTP client retry on `429`) than
by slowing every caller down.

## Cooldown

`Cooldown` polices traffic: it enforces "at most one action per key per interval" and
rejects anything that arrives early instead of delaying it. Each key gets its own limiter
with a burst of 1; unused keys are evicted in the background so memory stays bounded.

```go
cd := ratelimit.NewCooldown(
    3*time.Minute,  // interval: one call per key every 3 minutes
    1*time.Hour,    // maxAge: evict keys idle this long
    10*time.Minute, // cleanupInterval: how often to evict
)
defer cd.Stop()

if allowed, retryAfter := cd.Allow(nsId + ":" + infraId); !allowed {
    return fmt.Errorf("retry after %v", retryAfter)
}
```

A rejected call does **not** consume the key's slot, so `retryAfter` stays accurate no
matter how often a client retries.

Pass `0` for `maxAge` or `cleanupInterval` to skip the background goroutine entirely.

## ErrLimited

`Pacer.Wait` returns `*ErrLimited` when it gives up. It carries `RetryAfter`, the duration
to advertise to the caller. Two helpers cover the usual response wiring:

```go
if retryAfter, ok := ratelimit.RetryAfter(err); ok {
    w.Header().Set("Retry-After", strconv.Itoa(ratelimit.RetryAfterSeconds(retryAfter)))
    w.WriteHeader(http.StatusServiceUnavailable)
}
```

- `RetryAfter(err)` — unwraps any `*ErrLimited` in the error chain.
- `RetryAfterSeconds(d)` — rounds **up** to whole seconds, never below `1`, as `Retry-After` requires.

`Cooldown.Allow` returns `retryAfter` directly instead of an error, since a rejected call
is a normal outcome there rather than a failure.

## Concurrency

Both types are safe for concurrent use. A `Pacer` or `Cooldown` is meant to be created
once and shared by every caller that contends for the same resource — one per resource,
not one per request.

## On the names

_Pacing_ and _cooldown_ are the established terms for these two policies, so the types
borrow them rather than inventing anything. `golang.org/x/time/rate` exposes the same two
modes as `Wait` and `Allow`, which is exactly the method each type specializes in.

The `-er` asymmetry is intentional. Go reserves that suffix for **actors**, mainly
single-method interfaces (`io.Reader`); concrete types are named for what they **are**
(`sync.Once`, `sync.WaitGroup`, `semaphore.Weighted`). A `Pacer` is an actor — it paces your
calls for you. A `Cooldown` is a rule you consult (`cooldown.Allow(key)`), so `Cooldowner`
would apply the letter of a convention that doesn't fit.

## Users in this repository

- [pkg/client/tumblebug/call-pacer.go](../client/tumblebug/call-pacer.go) — a `Pacer` keeping Beetle under CB-Tumblebug's 2 req/s limit on infra read APIs. See [docs/tumblebug-call-pacer.md](../../docs/tumblebug-call-pacer.md).
- [pkg/api/rest/middlewares/ssh-check-cooldown.go](../api/rest/middlewares/ssh-check-cooldown.go) — a `Cooldown` allowing an SSH readiness check once per infrastructure every 10 seconds.
