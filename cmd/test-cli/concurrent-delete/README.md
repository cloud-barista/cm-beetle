# Concurrent Delete Test

Tests concurrent infrastructure deletion to verify rate limiting behavior.

## Purpose

This test validates that CM-Beetle's rate limiter correctly handles multiple simultaneous infrastructure deletion requests without triggering CB-Tumblebug's rate limit (2 req/s on ReadInfra API).

## Test Flow

1. **Phase 1: Recommendation & Migration**
   - Creates 5 infrastructures (one per CSP)
   - Uses timestamp-based nameSeed: `cdtest-YYYYMMDD-HHMMSS-XX`
   - Each infrastructure is created sequentially

2. **Phase 2: Stabilization**
   - Waits 10 seconds for infrastructure to stabilize

3. **Phase 3: Concurrent Deletion**
   - Deletes all 5 infrastructures simultaneously using goroutines
   - Rate limiter should activate and pace requests at ~600ms intervals

## Configuration

Edit `testconf/test-config.yaml`:

```yaml
test:
  cases:
    - csp: aws
      region: ap-northeast-2
      name: AWS-Seoul
      execute: true  # Set to false to skip
    # ... (up to 5 CSPs)

beetle:
  endpoint: http://localhost:8056
  namespaceId: mig01
  requestBodyFile: testconf/recommendation-request.json
```

**Note**: The test will use the first 5 CSPs where `execute: true`.

## Usage

```bash
# 1. Ensure CM-Beetle is running
make run

# 2. Run the test
cd cmd/test-cli/concurrent-delete
go run main.go -config testconf/test-config.yaml

# 3. Monitor rate limiter logs (in another terminal)
tail -f cmd/cm-beetle/log/cm-beetle.log | grep -E "rate|queue|ReadInfra"
```

## Expected Behavior

### Successful Rate Limiting

```log
[DEBUG] DeleteInfraRateLimiter initialized (interval: 600ms, max queue: 50, target rate: 1.67 req/s)
[DEBUG] Entered delete rate limiter queue (queue size: 1/50)
[DEBUG] Entered delete rate limiter queue (queue size: 2/50)
[DEBUG] Entered delete rate limiter queue (queue size: 3/50)
[DEBUG] Rate limiting ReadInfra call (waiting 234ms, current rate: 1.67 req/s)
[DEBUG] ReadInfra succeeded (nsId: mig01, infraId: cdtest-..., attempt: 1)
```

### Timing Expectations

- **5 concurrent deletions**: ~3 seconds total
  - Request 1: 0ms (immediate)
  - Request 2: 600ms wait
  - Request 3: 1200ms wait
  - Request 4: 1800ms wait
  - Request 5: 2400ms wait

- **No TB rate limit errors** (HTTP 429) should occur

### Adaptive Behavior

If rate limit is detected (unlikely with proper implementation):
```log
[WARN] Rate limit hit on ReadInfra, retrying after 1s (attempt: 1/3)
[WARN] Rate limit detected - rate limiter slowing down: 600ms → 800ms
```

After consistent success:
```log
[DEBUG] Rate limiter speeding up: 600ms → 550ms (rate: 1.82 req/s, consecutive ok: 5)
```

## Troubleshooting

**Issue**: All deletions fail immediately
- Check if CM-Beetle is running: `curl http://localhost:8056/beetle/health`
- Verify namespace exists in CB-Tumblebug

**Issue**: "Queue full" errors
- Should not occur with 5 concurrent requests (queue size: 50)
- If it does, there may be other concurrent operations running

**Issue**: Infrastructure creation fails
- Check CSP credentials in CB-Tumblebug
- Verify `testconf/test-config.yaml` CSP/Region values
- Ensure `testconf/recommendation-request.json` exists

## Cleanup

If test is interrupted, manually delete infrastructures:

```bash
# List infrastructures
curl http://localhost:8056/beetle/migration/ns/mig01/infra?option=id

# Delete specific infrastructure
curl -X DELETE "http://localhost:8056/beetle/migration/ns/mig01/infra/{infraId}?option=terminate"
```
