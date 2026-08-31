# Beetle Async Response — Manual Test Checklist

Assumes Beetle is up at `http://localhost:8056/beetle` (per `deployments/docker-compose/docker-compose.yaml`)
with Tumblebug reachable as its backend dependency. If `BEETLE_API_AUTH_ENABLED=true`, add
`-u "$BEETLE_API_USERNAME:$BEETLE_API_PASSWORD"` to every curl call below.

## 0. Readiness

```bash
curl -sf http://localhost:8056/beetle/readyz && echo OK
```

## 1. Sync default unchanged (no `Prefer` header)

```bash
curl -si -X POST http://localhost:8056/beetle/recommendation/infra \
  -H "Content-Type: application/json" \
  -d '{"desiredCspAndRegionPair":{"csp":"aws","region":"ap-northeast-2"},"OnpremiseInfraModel":{}}' \
  | head -20
```

Expect: `200` (or `400`/`500` depending on body validity) — same as before this change, no `202`.

## 2. Async opt-in (`Prefer: respond-async`)

```bash
RESP=$(curl -si -X POST http://localhost:8056/beetle/recommendation/infra \
  -H "Content-Type: application/json" \
  -H "Prefer: respond-async" \
  -d '{"desiredCspAndRegionPair":{"csp":"aws","region":"ap-northeast-2"},"OnpremiseInfraModel":{}}')
echo "$RESP" | head -20
```

Expect: `202 Accepted`, header `Preference-Applied: respond-async`, body `{"data":{"reqId":"...", "status":"Handling", "statusUrl":"..."}}`.

```bash
REQ_ID=$(echo "$RESP" | grep -o '"reqId":"[^"]*"' | cut -d'"' -f4)
curl -s http://localhost:8056/beetle/request/$REQ_ID | python3 -m json.tool
```

Poll a few times a couple seconds apart — `status` should move from `Handling` to `Success` or `Error`,
and stay there (not revert, not disappear).

## 3. `wait=N` is ignored (no hybrid mode)

```bash
curl -si -X POST http://localhost:8056/beetle/recommendation/infra \
  -H "Content-Type: application/json" \
  -H "Prefer: wait=5" \
  -d '{"desiredCspAndRegionPair":{"csp":"aws","region":"ap-northeast-2"},"OnpremiseInfraModel":{}}' \
  | head -5
```

Expect: behaves exactly like test #1 (sync, no `202`) — `wait=N` alone must not trigger async.

## 4. Capacity limit (`503` + `Retry-After`)

Fire more than 20 concurrent async jobs against a slower endpoint (migration infra is a good target
since it's naturally slower, giving the jobs time to pile up):

```bash
for i in $(seq 1 25); do
  curl -s -o /dev/null -w "%{http_code} " -X POST \
    "http://localhost:8056/beetle/migration/ns/mig01/infra?useExisting=false" \
    -H "Content-Type: application/json" -H "Prefer: respond-async" \
    -d '{"TargetInfra":{"name":"loadtest-'"$i"'","nodeGroups":[]}}' &
done
wait
echo
```

Expect: mostly `202`, with some `503` once the 20-job cap is hit, each `503` response including a
`Retry-After` header. (Exact split depends on how fast each job finishes — this call site has no real
node groups so it should error out and free its slot quickly; adjust the loop count/target endpoint if
you want to reproduce it more reliably.)

## 5. Spot-check the other endpoint groups

Repeat test #1/#2 for at least one endpoint per group to catch anything endpoint-specific:
- `POST /migration/ns/{nsId}/infraWithDefaults`
- `DELETE /migration/ns/{nsId}/infra/{infraId}` (expect `200`/`204` sync default unchanged)
- `POST /recommendation/infraWithNlb`
- `POST /migration/middleware/ns/{nsId}/infra/{infraId}/nlb`
- `DELETE /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}`
- `POST /migration/middleware/ns/{nsId}/objectStorage`
- `POST /migration/data` (this one is the exception — confirm the *default* is now sync `200`, not `202`)
