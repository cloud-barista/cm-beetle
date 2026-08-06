# CM-Beetle Multi-Infra Recommendation Test CLI

Automated test tool for `POST /recommendation/multiInfra` and `POST /recommendation/multiInfraWithNlb` —
the cross-CSP comparison APIs that recommend one best-match candidate per target CSP/region pair.

Modeled on `cmd/test-cli/infra-with-nlb`'s structure (config loading, readiness check, markdown
report generation), but scoped to exactly the recommendation call itself: no validation,
migration, SSH check, or cleanup, since neither endpoint provisions anything.

## What It Checks

For each endpoint, the CLI sends one request covering all configured targets and verifies the
response accurately reflects the input:

1. **Request succeeds** — `200` with `success: true`.
2. **Input/output accuracy** — the response has exactly one result item per requested target,
   in request order, matched via `targetCloud` (`csp`/`region`).

## Quick Start

```bash
# From the repository root
make test-multi-infra-recommendation
```

This copies `testconf/template-test-config.yaml` to `testconf/test-config.yaml` on first run —
edit it to point at your Beetle endpoint and select at least 2 CSP/region cases with `execute: true`.

## Configuration

`testconf/test-config.yaml`:

```yaml
test:
  cases:
    - csp: aws
      region: ap-northeast-2
      name: AWS-Seoul
      execute: true
    - csp: azure
      region: koreasouth
      name: Azure-Busan
      execute: true

beetle:
  endpoint: http://localhost:8056
  requestBodyFile: testconf/source-infra.json
  requestBodyFileWithNlb: testconf/source-infra-with-nlb.json
  authConfigFile: testconf/auth-config.json
```

- `requestBodyFile` / `requestBodyFileWithNlb`: `{"nameSeed": ..., "sourceInfra": {...}}` fixtures
  (the latter's `sourceInfra` includes `nlbs`), used as the `sourceInfra` field for the two tests.
- Cases with `execute: true` become the request's `desiredCspAndRegionPairs` (capped at 10).
  At least 2 are required.

## Results

- Console: real-time pass/fail per endpoint and per-target accuracy detail
- File: `testresult/multi-infra-recommendation-test-report.md` — full request/response bodies
  (collapsible) and the accuracy check for both endpoints

Exit code is non-zero if either endpoint's call or accuracy check fails.
