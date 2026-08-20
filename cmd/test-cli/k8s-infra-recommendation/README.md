# CM-Beetle K8s Infra Recommendation Test CLI

Automated test tool for `POST /recommendation/k8sCluster` — the API that turns an on-premise
Kubernetes cluster description into a target-cloud K8s infra design.

Scoped to the recommendation call only: no validation, migration, or cleanup, since the endpoint
provisions nothing. Migration is covered separately by `cmd/test-cli/k8s-infra-migration`.

## What It Checks

Each on-premise fixture is sent to every enabled CSP/region pair, and the response is checked
against expectations declared in `testconf/test-config.yaml`:

- **Status code** — required for every scenario (200 for positive cases, 400 for negative ones)
- **Node group count** — how many node groups the worker set was folded into
- **Total node size** — sum of `desiredNodeSize` across node groups, i.e. worker count reproduction
- **Version prefix** — optional check on the recommended Kubernetes version

For any 200 response, these structural checks always run:

- `targetK8sCluster.name` and `.version` are non-empty
- every node group has a non-empty `specId`
- every node group has `desiredNodeSize >= 1`

## Quick Start

```bash
# From the repository root
make test-k8s-infra-recommendation
```

This copies `testconf/template-test-config.yaml` to `testconf/test-config.yaml` on first run.
Edit it to point at your Beetle endpoint and enable the CSP/region pairs you want.

Credentials go in `testconf/auth-config.json` (copy from `testconf/template-auth-config.json`).

## Configuration

`testconf/test-config.yaml` has three parts.

### `test.set.mode`

`parallel` (default) runs one goroutine per target so scenarios against a single CSP stay ordered
and readable. `sequential` runs everything in order.

### `test.cases` — target CSP/region pairs

```yaml
cases:
  - csp: aws
    region: ap-northeast-2
    name: AWS-Seoul
    execute: true
```

Recommendation costs nothing, so enabling several at once is cheap.

### `test.scenarios` — fixtures and expectations

```yaml
scenarios:
  - file: testconf/scenarios/honeybee-k8s-workers5.json
    name: workers5
    execute: true
    expect:
      statusCode: 200
      nodeGroupCount: 1
      totalNodeSize: 5
```

`expect.statusCode` is required. `nodeGroupCount`, `totalNodeSize`, and `versionPrefix` are
optional — omit one to skip that check.

Keeping expectations in config rather than in code matters here: the CLI must not encode what a
good recommendation looks like per CSP, because that judgement is exactly what is under test.

## Fixtures

`testconf/scenarios/` holds 15 on-premise models captured from cm-honeybee `/infra/refined`,
each isolating one recommendation behaviour.

| Fixture | Workers | Distinct specs | Isolates |
| --- | --- | --- | --- |
| `honeybee-k8s-refined-infra.json` | 2 | 1 | Baseline |
| `honeybee-k8s-workers0.json` | 0 | — | Control-plane only — no recommendation possible |
| `honeybee-k8s-workers1.json` | 1 | 1 | Single worker |
| `honeybee-k8s-workers5.json` | 5 | 1 | Same-spec workers fold into one group |
| `honeybee-k8s-tiny.json` | 1 | 1 | Below minimum viable worker spec (2 vCPU / 4 GiB) |
| `honeybee-k8s-3groups.json` | 3 | 3 | Three specs produce three node groups |
| `honeybee-k8s-hetero.json` | 2 | 2 | Two specs produce two node groups |
| `honeybee-k8s-arm64.json` | 2 | 1 | arm64 spec and matching node image |
| `honeybee-k8s-mixed-arch.json` | 2 | 1 (arch differs) | Architecture-based grouping |
| `honeybee-k8s-samecpu-diffdisk.json` | 2 | 1 (disk differs) | Whether disk size splits node groups |
| `honeybee-k8s-spec-small.json` | 2 | 1 | Upscaling to a viable spec |
| `honeybee-k8s-spec-large.json` | 2 | 1 | Large spec, no upscaling |
| `honeybee-k8s-v134.json` | 2 | 1 | Version mapping |
| `honeybee-k8s-v199.json` | 2 | 1 | Non-existent version falls back to latest |
| `honeybee-k8s-norole.json` | 0 | — | No `role: worker` present |

`nodeGroupCount` is deliberately left unset for `mixed-arch` and `samecpu-diffdisk` — the correct
grouping for those inputs is still an open design question, so pinning a number would freeze the
current behaviour rather than test it.

### Negative fixtures

`testconf/scenarios/negative/` holds three cm-honeybee payloads captured **before**
`/infra/refined`. All three describe the same source cluster at different pipeline stages:

```
honeybee source_group response    →  honeybee-k8s-raw-source-group.json     (top-level array)
  └ honeybee connection_info      →  honeybee-k8s-raw-connection-info.json  (full node labels)
      └ labels/node_info trimmed  →  beetle-k8s-recommendation-request.json
```

None of them use the `onpremiseInfraModel` wrapper the API expects, so feeding them directly is a
realistic user mistake. The API should reject each with a clean 400.

## CLI Options

```bash
cd cmd/test-cli/k8s-infra-recommendation
go run main.go -config testconf/test-config.yaml
```

| Option    | Default                     | Description                     |
| --------- | --------------------------- | ------------------------------- |
| `-config` | `testconf/test-config.yaml` | Path to test configuration file |

Exit code is non-zero if any case fails, so the CLI can gate CI.

## Results

- Console: per-case pass/fail with the individual check outcomes
- File: `testresult/k8s-infra-recommendation-test-report.md`

The report opens with a scenario × target matrix, so a regression confined to one CSP or one
scenario is visible at a glance, followed by per-case detail with the full response body.
CSP account identifiers (Azure subscription IDs, GCP project IDs, email addresses) are masked.

## Prerequisites

1. CM-Beetle running at the configured endpoint
2. CB-Tumblebug reachable by CM-Beetle, with the target CSPs' spec/image catalogs loaded
3. Credentials configured in `testconf/auth-config.json`

No CSP write permissions are needed — nothing is created.
