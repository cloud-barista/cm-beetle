# Rate Limiting Test

Exercises the two limits Beetle applies to Tumblebug retrievals: the call pacer's wait budget
(seen as an immediate `503`) and the pacing of concurrent deletions.

## Purpose

Beetle paces its calls to Tumblebug's read APIs at ~1.6 req/s so it stays under Tumblebug's
2 req/s per-source-IP limit. This test confirms two things:

- A burst larger than the wait budget can absorb is **refused immediately** with `503`, before
  any Tumblebug call is made.
- Concurrent deletions are **spaced, not rejected**, because each carries a longer wait budget.

See [docs/tumblebug-call-pacer.md](../../../docs/tumblebug-call-pacer.md) for the design.

## Test Flow

1. **Recommendation & Migration** — creates up to 5 infrastructures (one per CSP), sequentially.
2. **Stabilization** — waits 10 s.
3. **Retrieval Burst** — fires `-retrieval-burst` (default 30) simultaneous `GET .../infra/{infraId}`
   requests and reports how many were admitted vs. refused, and how fast the refusals came back.
4. **Concurrent Deletion** — deletes every infrastructure at once using goroutines.

## Configuration

Edit `testconf/test-config.yaml`:

```yaml
test:
  cases:
    - csp: aws
      region: ap-northeast-2
      name: AWS-Seoul
      execute: true # Set to false to skip
    # ... (up to 5 CSPs)

beetle:
  endpoint: http://localhost:8056
  namespaceId: mig01
  requestBodyFile: testconf/recommendation-request.json
```

The test uses the first 5 CSPs where `execute: true`.

## Usage

```bash
# 1. Ensure CM-Beetle is running
make run

# 2. Run the test
make test-rate-limiting

# Or directly, with a different burst size
cd cmd/test-cli/rate-limiting
go run main.go -config testconf/test-config.yaml -retrieval-burst 40

# 3. Monitor pacer logs (in another terminal)
tail -f cmd/cm-beetle/log/cm-beetle.log | grep -i pacer
```

## Expected Behavior

### Retrieval burst

A burst passes through two independent limits. Beetle's server-wide edge limiter admits 20 at
once and answers `429` to the rest; of those that get through, the pacer admits about 13 within
the 8 s wait budget and refuses the remainder with `503` — before any Tumblebug call, so those
answer in milliseconds:

```log
Sending 30 simultaneous retrievals across 5 infrastructure(s)...
  Admitted     (2xx): 13, slowest 7.61s
  Paced out    (503):  7, slowest 24ms, Retry-After: 8
  Edge limited (429): 10 (Beetle's own 20 req/s limiter, never reached the pacer)
✅ Refusals returned immediately (within 1s)
```

The `429` count is expected, not a failure. If nothing is refused with `503`, raise
`-retrieval-burst`.

### Concurrent deletion

Each deletion carries a 30 s wait budget, so all of them are admitted and merely staggered by
625 ms. Five deletions therefore take ~3 s to get past the pacer, and **no** Tumblebug `429`
should appear. If Tumblebug still answers `429` (another client shares the same source IP), the
request is retried up to 3 times.

## Troubleshooting

**Everything fails immediately**

- Check CM-Beetle is running: `curl http://localhost:8056/beetle/health`
- Verify the namespace exists in CB-Tumblebug

**Deletions get `503` too**

- Other concurrent operations are competing for slots
- Raise `BEETLE_TUMBLEBUG_RETRIEVAL_MAX_WAIT_SEC` (capped at 120 s)

**Infrastructure creation fails**

- Check CSP credentials in CB-Tumblebug
- Verify CSP/Region values in `testconf/test-config.yaml`
- Ensure `testconf/recommendation-request.json` exists

## Cleanup

If the test is interrupted, delete the infrastructures manually:

```bash
# List infrastructures
curl http://localhost:8056/beetle/migration/ns/mig01/infra?option=id

# Delete a specific one
curl -X DELETE "http://localhost:8056/beetle/migration/ns/mig01/infra/{infraId}?option=terminate"
```
