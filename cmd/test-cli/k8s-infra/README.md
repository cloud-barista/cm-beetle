# CM-Beetle K8s Infra Unified Test CLI

Unified test tool for CM-Beetle's Kubernetes recommendation, validation, and migration lifecycle APIs (`cmd/test-cli/k8s-infra`).

Consolidates the previously separate `k8s-infra-recommendation` and `k8s-infra-migration` CLIs into a single binary with flexible execution modes.

---

## Execution Modes

| Mode | Flag | Cloud Cost | Description |
|---|---|:---:|---|
| **Migration (default)** | `-mode migration` (or `make test-k8s-infra`) | Incurs cloud cost | Full E2E lifecycle on real CSPs: Recommend → Validate → Migrate → List → Get & Verify → Workload → Delete → Residual check. |
| **Recommendation** | `-mode recommendation` (or `make test-k8s-infra MODE=recommendation`) | **$0 (Free)** | Offline high-speed test of 20 on-premise scenario fixtures against configured CSPs. |
| **Validation** | `-mode validation` (or `make test-k8s-infra MODE=validation`) | **$0 (Free)** | Standalone Pre-flight Validation API test (`POST /beetle/validation/ns/{nsId}/k8sCluster`) without provisioning. |
| **All** | `-mode all` (or `make test-k8s-infra MODE=all`) | Incurs cloud cost | Runs all recommendation scenarios first; if all pass, proceeds to E2E migration. |

---

## Quick Start

### 1. Recommendation Test (Free, Instant)
```bash
# Run 20 scenario fixtures against enabled CSP targets (zero cloud cost)
make test-k8s-infra MODE=recommendation

# Or via backward-compatible alias
make test-k8s-infra-recommendation
```

### 2. Validation Test (Free, Instant)
```bash
# Test pre-flight validation API contract
make test-k8s-infra MODE=validation
```

### 3. Full E2E Migration Test (Provisions Real Clusters)
```bash
# Run full E2E migration pipeline (Recommend -> Validate -> Migrate -> Workload -> Delete)
make test-k8s-infra

# Or via backward-compatible alias
make test-k8s-infra-migration
```

---

## E2E Migration Lifecycle Steps

| # | Step | API / Target |
|---|---|---|
| **1** | Recommend | `POST /beetle/recommendation/k8sCluster` |
| **2** | **Pre-flight Validation** | `POST /beetle/validation/ns/{nsId}/k8sCluster` |
| **3** | Migrate | `POST /beetle/migration/ns/{nsId}/k8sCluster` |
| **4** | List | `GET /beetle/migration/ns/{nsId}/k8sCluster` |
| **5** | Get & Verify | `GET /beetle/migration/ns/{nsId}/k8sCluster/{id}` (verified against recommendation) |
| **6** | Workload Verification | CB-Tumblebug kubeconfig & bearer token → Cluster API server (Nginx Deployment & LB) |
| **7** | Delete | `DELETE /beetle/migration/ns/{nsId}/k8sCluster/{id}` |
| **8** | Residual Check | Direct CB-Tumblebug inspection (VNet / SecurityGroup / SshKey) |

> [!NOTE]
> If a step fails, subsequent migration steps are skipped except cleanup: **Step 7 (Delete) always runs once a cluster exists** to prevent orphaned billable cloud resources.

---

## Configuration (`testconf/test-config.yaml`)

On first run, `template-test-config.yaml` is copied to `testconf/test-config.yaml`.

```yaml
test:
  executionMode: migration # migration | recommendation | validation | all
  set:
    mode: parallel # parallel or sequential
    startDelaySeconds: 10
  cases:
    - csp: aws
      region: ap-northeast-2
      name: AWS-Seoul
      execute: false # Enable target CSPs carefully for migration
  scenarios:
    - file: testconf/scenarios/honeybee-k8s-refined-infra.json
      name: baseline
      execute: true
      expect:
        statusCode: 200
        nodeGroupCount: 1
        totalNodeSize: 2

beetle:
  endpoint: http://localhost:8056
  namespaceId: mig01
  authConfigFile: testconf/auth-config.json
  requestBodyFile: testconf/recommendation-request.json
  nameSeed: mig
```

---

## Output Reports

Execution results are recorded in `cmd/test-cli/k8s-infra/testresult/`:
- **Recommendation report**: `k8s-infra-recommendation-test-report.md`
- **Per-CSP migration report**: `k8s-infra-migration-test-results-<csp>.md`
- **Overall migration summary**: `k8s-infra-migration-test-summary-all.md`

