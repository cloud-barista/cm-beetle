# CM-Beetle K8s Infra Migration Test CLI

End-to-end test tool for CM-Beetle's K8s migration APIs. It runs the full lifecycle against a
real CSP: recommend → migrate → list → get (verified against the recommendation) → delete →
residual resource check.

> [!WARNING]
> Every enabled target provisions a **real managed K8s cluster and incurs cost**. Enable one
> CSP at a time until the flow is confirmed for it. Cleanup always runs once a cluster exists,
> even when an earlier step fails.

Recommendation-only checks live in `cmd/test-cli/k8s-infra-recommendation`, which is free to
run and needs no CSP write permissions.

## Test Flow

| # | Step | API |
| --- | --- | --- |
| 1 | Recommend | `POST /beetle/recommendation/k8sCluster` |
| 2 | Migrate | `POST /beetle/migration/ns/{nsId}/k8sCluster` |
| 3 | List | `GET /beetle/migration/ns/{nsId}/k8sCluster` |
| 4 | Get + **verify vs recommendation** | `GET /beetle/migration/ns/{nsId}/k8sCluster/{id}` |
| 5 | **Workload verification** | CB-Tumblebug kubeconfig + token -> the cluster's own K8s API |
| 6 | Delete | `DELETE /beetle/migration/ns/{nsId}/k8sCluster/{id}` |
| 7 | Residual resource check | CB-Tumblebug (VNet / SecurityGroup / SshKey) |

If a step fails, the remaining steps are skipped — except cleanup. **Step 6 always runs once a
cluster ID exists**, so a mid-run failure never leaves a billable cluster behind.

Beetle creates the cluster before its node groups, so a migration can fail *after* a running
cluster exists — a node group failure is the common case. The failure response does not
necessarily carry the cluster's ID. When step 2 fails, the CLI therefore looks the cluster up by
the recommended name and adopts it for cleanup, rather than assuming nothing was created.

### Step 5 proves the cluster is usable

Steps 1-4 confirm the cluster was *created* as designed. Step 5 confirms it *works*: it fetches
the kubeconfig from CB-Tumblebug, reaches the cluster's own API server, checks that the number
of Ready nodes matches the recommendation, then runs an nginx Deployment and waits for its pod
to reach `Running` before removing it.

This is the K8s counterpart of the SSH connectivity check `cmd/test-cli/infra` runs against
every migrated VM. A cluster that reports `Active` but cannot schedule a pod has not been
migrated in any useful sense.

Three things make this work without Docker or `kubectl`:

- The **dedicated `/kubeconfig` sub-resource** is used, not the generic cluster GET — the latter
  returns `AccessInfo.Kubeconfig` without the native flag and never becomes ready on some CSPs.
  It is polled, because it lags the `Active` status while the control-plane endpoint provisions.
- The **bearer token** comes from Tumblebug's `/token` endpoint, since CSP-native kubeconfigs
  authenticate through an exec plugin rather than a static token.
- The **cluster CA** from `certificate-authority-data` is trusted explicitly, because the API
  server's certificate is signed by the cluster's own CA.

Headlamp (used by cb-tumblebug's own `k8s-test`) is deliberately not used: it only forwards the
bearer token to the API server unchanged, so it adds a Docker dependency without adding reach.

With `workload.loadBalancerEnabled` (default on), the step goes further: it publishes the
Deployment through a `LoadBalancer` Service, waits for the CSP to assign an ingress address, and
fetches the page from outside the cluster. Scheduling a pod proves the cluster runs work; this
proves the work is *reachable*, which is a distinct failure mode — a cluster can come up healthy
and still fail to provision a working load balancer.

Two details make this reliable across CSPs:

- AWS publishes a **hostname**, other CSPs an **IP**; both `status.loadBalancer.ingress[].ip` and
  `.hostname` are accepted.
- An address appears **before** the load balancer forwards traffic — health checks must pass, and
  a hostname has to resolve — so the HTTP fetch is retried rather than attempted once.

Cleanup deletes the Service **before** the Deployment, because the Service is what holds the real,
billable cloud load balancer. A cleanup failure there is reported as an error rather than a note,
since a leaked load balancer keeps costing money after the run.

Expect roughly 2-6 extra minutes per target. Set `workload.loadBalancerEnabled: false` to stop at
pod `Running`, or `workload.enabled: false` to skip the whole step.

### Step 4 is the point of this CLI

Steps 1–3 confirm the API responded. Step 4 confirms the cluster **matches what was
recommended**: node group count, and per node group (matched by name, since CSPs may reorder)
the `specId` and `desiredNodeSize`. A migration that succeeds while silently dropping a node
group or substituting a spec would otherwise go unnoticed.

## Waiting: the CLI does not track per-CSP provisioning times

Cluster creation takes roughly 3–25 minutes depending on the CSP. **The CLI does not need to
know that.** Beetle already polls until the cluster reaches `Active` — 15 s × 160 = 40 min in
`pkg/core/migration/k8s-infra.go` — and a per-CSP timeout was deliberately rejected there as
false precision, since provisioning time also grows with cluster size.

So the only problem left for the CLI is *how to wait for a call that can take tens of minutes*:

- **`migration.mode: async`** (default) — sends `Prefer: respond-async`, gets `202` with a
  `reqId`, then polls `GET /request/{reqId}`. The HTTP connection stays short, progress is
  printed each poll, and the timeout is the CLI's to control. A `503` (Beetle's async job pool
  of 20 is full) is retried, not failed.
- **`migration.mode: sync`** — holds the connection until the cluster is `Active`. Useful for
  exercising the blocking path; `migration.timeoutSec` must exceed Beetle's own 40 min budget.

Deletion has **no async mode**, so it is a long synchronous call with its own timeout and a
retry loop (see below).

## Quick Start

```bash
# From the repository root
make test-k8s-infra-migration
```

This copies `testconf/template-test-config.yaml` to `testconf/test-config.yaml` on first run.
Edit it to enable exactly the CSP you intend to provision.

Credentials go in `testconf/auth-config.json` (copy from `testconf/template-auth-config.json`).
Tumblebug credentials are needed only for step 6.

## Configuration

`testconf/test-config.yaml`:

| Key | Default | Purpose |
| --- | --- | --- |
| `test.set.mode` | `parallel` | `parallel` (one goroutine per target) or `sequential` |
| `test.set.startDelaySeconds` | `10` | Stagger between parallel starts, to avoid bursting Tumblebug's read APIs |
| `migration.mode` | `async` | `async` or `sync` — see above |
| `migration.timeoutSec` | `3600` | HTTP timeout for the sync path |
| `poll.intervalSec` | `30` | Async status poll interval |
| `poll.timeoutSec` | `3600` | Async overall budget — must exceed Beetle's 40 min |
| `delete.timeoutSec` | `1800` | Beetle polls node group teardown for up to 20 min |
| `delete.retryIntervalSec` | `60` | Backoff between deletion attempts |
| `delete.maxRetries` | `10` | Deletion attempts before giving up |
| `verify.residualResources` | `true` | Run step 7 |
| `workload.enabled` | `true` | Run step 5 |
| `workload.kubeconfigTimeoutSec` | `600` | Kubeconfig readiness budget |
| `workload.podReadyTimeoutSec` | `180` | nginx pod `Running` budget |
| `workload.loadBalancerEnabled` | `true` | Publish via LoadBalancer and fetch from outside |
| `workload.lbAddressTimeoutSec` | `900` | Budget for the CSP to assign an ingress address |
| `workload.lbAccessTimeoutSec` | `300` | Budget for the endpoint to start serving |

## Known CSP behaviours — recorded, not failed

These are Beetle/CSP behaviours the CLI reports without failing the run. Fixing them is tracked
in `docs/k8s-migration/k8s-migration-improvement-analysis.md`, not here.

| Behaviour | Handling |
| --- | --- |
| Deletion rejected for minutes after `Active` (observed on NCP: `409 UPDATING`) | Retried on a fixed backoff, up to `delete.maxRetries` |
| Deleting an already-absent cluster returns 500 instead of a no-op success | Treated as cleanup success, noted in the report |
| VNet / SecurityGroup / SshKey survive cluster deletion (all CSPs) | Step 6 records what remains; informational only |

The primary target is **AWS**; a run that passes on AWS is treated as green.

## Results

- Console: per-step pass/fail with check outcomes and live polling progress
- Files, under `testresult/`:
  - `k8s-infra-migration-test-results-{csp}.md` — full lifecycle per CSP, with the
    recommendation and the created cluster attached
  - `k8s-infra-migration-test-summary-all.md` — one-row-per-target overview

CSP account identifiers (Azure subscription IDs, GCP project IDs, email addresses) are masked.

Exit code is non-zero if any step fails.

## Prerequisites

1. CM-Beetle running at the configured endpoint
2. CB-Tumblebug and CB-Spider reachable, with credentials registered for the target CSP
3. **CSP write permissions** — clusters, VPCs, security groups and SSH keys are created
4. `testconf/auth-config.json` filled in

## Known limitation: repeated runs

Resource names are derived from the source cluster and are not uniquified per run, and creation
is not idempotent. Sequential runs are fine because each run deletes its cluster, but two
concurrent runs against the same CSP/region will collide, and leftover prerequisites from an
earlier run (see step 6) can cause a later run to fail. This is Beetle behaviour tracked as
improvements #3 and #9 in `docs/k8s-migration/k8s-migration-improvement-analysis.md`.
