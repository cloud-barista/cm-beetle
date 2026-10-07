# Migration Lineage and Resource Traceability

## What is Migration Lineage?

**Migration Lineage** records and tracks the relationships between the information and models required throughout the Cloud-Barista cloud migration lifecycle:

> [!NOTE]
> **Why Lineage is Managed as Metadata**:  
> Migration Lineage records the relationships between the information and models required for migration (linking source groups, source models, planning models, and target models). Managing these relationships as external metadata (via model envelopes and cloud resource labels) rather than embedding them directly within infrastructure resource models (`OnpremInfra`, `RecommendedInfra`) keeps the technical models pure, decoupled, and reusable across environments.

```
[Phase 1: Discovery]           Honeybee Source Group (sourceGroupId)
                                      │
                                      ▼
[Phase 2: Source Model]        Damselfly Source Model (sourceModelId: Immutable Fact)
                                      │
                                      ▼
[Phase 3: Planning Model]      Damselfly Planning Model (planningModelId: Modernization Plan)
                                      │
                                      ▼
[Phase 4: Target Model]        Damselfly Target Model (targetModelId: Recommended Blueprint)
                                      │
                                      ▼
[Phase 5: Provisioning]        Beetle Deployed Cloud Infra (Tumblebug & Cloud Resource Labels)
```

It establishes:
1. **Pass-by-Reference & Traceability**: Links deployed cloud infrastructure back to its planning blueprint, source model, and discovery source group.
2. **Standardized Node Label (`cm-source-machine-ids`)**: Unambiguously identifies which source machine(s) each target Node was recommended from, enabling downstream data synchronization (e.g., CM-Grasshopper).

---

## Lineage Data Model (`imdl/lineage`)

CM-Beetle provides a standalone Go package [`imdl/lineage`](../../imdl/lineage/lineage.go) with zero external dependencies:

```go
type MigrationLineage struct {
    SourceGroupId   string `json:"sourceGroupId,omitempty" example:"sg-ecommerce-prod"`
    SourceModelId   string `json:"sourceModelId,omitempty" example:"src-mdl-101"`
    PlanningModelId string `json:"planningModelId,omitempty" example:"plan-mdl-101"`
    TargetModelId   string `json:"targetModelId,omitempty" example:"tgt-mdl-201"`
}
```

### Label Conversion Helpers
- `lin.ToLabelMap()`: Converts non-empty lineage fields into standard cloud resource label pairs.
- `lineage.FromLabelMap(labels)`: Reconstructs a `MigrationLineage` struct from a resource label map.

---

## Cloud Resource Label Specifications

All labels adhere to multi-cloud labeling constraints (lowercase alphanumeric and hyphens; valid across AWS, GCP, and Azure).

### 1. Top-Level Infrastructure Labels (`InfraReq.Label`)

Stamped at provisioning time by CM-Beetle on the target multi-cloud infrastructure:

| Label Key | Type | Description | Example |
| :--- | :--- | :--- | :--- |
| `cm-target-model-id` | `string` | ID of the target recommendation model used for provisioning | `"tgt-mdl-201"` |
| `cm-planning-model-id` | `string` | ID of the migration planning model | `"plan-mdl-101"` |
| `cm-source-model-id` | `string` | ID of the immutable source model | `"src-mdl-101"` |
| `cm-source-group-id` | `string` | ID of the Honeybee discovery source group | `"sg-ecommerce-prod"` |
| `cm-migrated-at` | `string` | UTC timestamp of provisioning completion (ISO-8601, RFC 3339) | `"2026-10-07T10:30:00Z"` |

### 2. Node & NodeGroup Level Label (`NodeGroup.Label`)

Stamped at recommendation and provisioning time on each NodeGroup (and copied onto individual Nodes):

| Label Key | Type | Description | Example |
| :--- | :--- | :--- | :--- |
| **`cm-source-machine-ids`** | `string` | Comma-separated source machine IDs (authoritative mapping key) | `"srv-01,srv-02"` |

> [!IMPORTANT]
> **Single Standard (Breaking Change)**:  
> The legacy singular key `sourceMachineId` and plural `sourceMachineIds` have been unified into **`cm-source-machine-ids`**.  
> In accordance with CM-Beetle's *No Fallback* principle, no silent degradation or dual-key fallback is maintained.

---

## How to Use It

### 1. Stamping Lineage during Migration

Pass `lineage` as an optional field in `MigrateInfraRequest`:

```http
POST /beetle/migration/ns/mig01/infra
Content-Type: application/json

{
  "lineage": {
    "sourceGroupId": "sg-ecommerce-prod",
    "sourceModelId": "src-mdl-101",
    "planningModelId": "plan-mdl-101",
    "targetModelId": "tgt-mdl-201"
  },
  "targetCloud": { "csp": "aws", "region": "ap-northeast-2" },
  "targetInfra": { ... }
}
```

CM-Beetle automatically:
1. Generates runtime deployment timestamp: `Label["cm-migrated-at"] = time.Now().UTC().Format(time.RFC3339)`.
2. Injects `cm-target-model-id`, `cm-planning-model-id`, `cm-source-model-id`, and `cm-source-group-id` into the CB-Tumblebug infrastructure creation request.
3. Propagates provider tags/labels to cloud resources (AWS EC2 Tags, GCP Labels, Azure Tags).

### 2. Backward Compatibility (Consumer-First Readiness)

The `lineage` field is marked `omitempty`. If callers (Portal or Damselfly) invoke migration without `lineage`, CM-Beetle processes the request normally and stamps `cm-migrated-at` without failing or rejecting the call.

---

## Subsystem Interoperability & 1:1 Mapping Guide

### CM-Grasshopper (Data Synchronization)
CM-Grasshopper reads `cm-source-machine-ids` from the deployed target Node to determine which source server to replicate data from.

For a NodeGroup with multiple homogeneous nodes (e.g. `NodeGroupSize > 1`):
- `NodeGroup.Label`: `{"cm-source-machine-ids": "srv-01,srv-02,srv-03"}`
- **Index-Based 1:1 Matching**: Match the $i$-th target Node in the group to the $i$-th machine ID in the comma-separated list:
  - `node[0]` $\longleftrightarrow$ `srv-01`
  - `node[1]` $\longleftrightarrow$ `srv-02`
  - `node[2]` $\longleftrightarrow$ `srv-03`

### CM-Damselfly & Portal (Butterfly)
- **Damselfly**: Stores model envelopes containing `lineage` metadata (`sourceGroupId`, `sourceModelId`, `planningModelId`, `targetModelId`).
- **Portal**: Renders source migration lineage badges and displays the migration markdown report generated by CM-Beetle.

---

## Related Documents

- [Lineage Architecture & Orchestration Plan](../plan/lineage-plan.md) — Comprehensive multi-subsystem RFC and specification
- [Breaking Change Notice for Grasshopper](../announcements/breaking-change-grasshopper.md) — Migration guide for data synchronization
- [Damselfly Model Metadata Request](../announcements/damselfly-model-metadata-lineage.md) — Model envelope schema extension
- [Butterfly Portal UX Request](../announcements/butterfly-portal-lineage-and-ux-request.md) — Form-based UI and lineage visualization
