# Breaking Changes for v0.6.6

**Release Version:** CM-Beetle v0.6.5 -> v0.6.6  
**Release Date:** 2026-11-25  
**Dependency Updates:** CB-Tumblebug v0.13.8 -> v0.14.0 (CB-Spider v0.13.20 -> v0.14.0, CB-MapUI v0.13.14 -> v0.14.0)  
**Migration Required:** Yes - Breaking changes in Infra API and `imdl/cloud-model` request/response fields

---

# Part 1: Quick Reference

> **Target Audience:** Beetle API users, `imdl/cloud-model` package consumers  
> **Goal:** Understand breaking changes and required migration actions

---

## Breaking Changes Summary

| Category | Affected API / Model | Impact | Required Action |
| --- | --- | --- | --- |
| **API** | `POST /beetle/migration/ns/{nsId}/infra` | **BREAKING** - Request/response bootstrap schema changed | Migrate `targetInfra.postCommands` array to `targetInfra.postScript.phases`; use `Prefer: respond-async` instead of `postCommandAsync` |
| **API** | `POST /beetle/migration/ns/{nsId}/infraWithDefaults` | **BREAKING** - Request/response bootstrap schema changed & field removed | Migrate `postCommands` to `postScript.phases`; remove `installMonAgent`; use `Prefer: respond-async` for async execution |
| **API** | `GET /beetle/migration/ns/{nsId}/infra/{infraId}` | **BREAKING** - Response fields renamed/removed | Read `postScriptResults`, `postScriptStatus`, `postScriptRequestId`; remove `installMonAgent`, `monAgentStatus` references |
| **API** | `GET /beetle/summary/target/ns/{nsId}/infra/{infraId}` | **BREAKING** - Summary field removed | Remove `overview.installMonAgent` consumption |
| **Model** | `imdl/cloud-model` package | **BREAKING** - Struct/field renames and removals | Update `PostCommands` -> `PostScript`, `PostCommandReq` -> `PostScriptPhase`, remove `InstallMonAgent`/`MonAgentStatus`, rebuild |

> **Note:** For asynchronous infrastructure migration, CM-Beetle intentionally preserves synchronous post-script completion within Tumblebug so the target infra is guaranteed ready for subsequent data/software migration steps. Use the standard `Prefer: respond-async` request header (`202 Accepted` + `GET /beetle/request/{reqId}`) when non-blocking migration execution is required.

---

## Breaking Change #1: `postCommands` -> `postScript` (`PostScriptReq`)

### What Changed

**API Endpoints:**

- `POST /beetle/migration/ns/{nsId}/infra` - `targetInfra.postCommands` replaced by `targetInfra.postScript`; `postCommandAsync` removed from JSON body
- `POST /beetle/migration/ns/{nsId}/infraWithDefaults` - `postCommands` replaced by `postScript`; `postCommandAsync` removed from JSON body
- `GET /beetle/migration/ns/{nsId}/infra/{infraId}` (and `POST` creation responses) - `postCommandResults` / `postCommandStatus` / `postCommandRequestId` renamed to `postScriptResults` / `postScriptStatus` / `postScriptRequestId`

**Before (v0.6.5):**

```json
{
  "targetInfra": {
    "name": "my-infra",
    "installMonAgent": "no",
    "postCommands": [
      {
        "userName": "cb-user",
        "command": ["uname -a"]
      }
    ],
    "postCommandAsync": false
  }
}
```

**After (v0.6.6):**

```json
{
  "targetInfra": {
    "name": "my-infra",
    "postScript": {
      "phases": [
        {
          "userName": "cb-user",
          "command": ["uname -a"]
        }
      ]
    }
  }
}
```

*(Note: For `POST /beetle/migration/ns/{nsId}/infraWithDefaults`, `postScript` sits at the root of the JSON request body.)*

### Why This Changed

- Align request/response naming across `initScript` (first-boot cloud-init) and `postScript` (post-boot SSH phased execution) in CB-Tumblebug `v0.14.0`.
- Support phase-level template and script sequences (`templateId`, `scripts`, `env`) within `PostScriptPhase` (`InfraCmdReq`).
- Remove `postCommandAsync` from the JSON payload (`PostScriptAsync` has `json:"-"` in `imdl/cloud-model`); CM-Beetle provides unified async orchestration via `Prefer: respond-async`.

### Migration Action

**For API Consumers:**

1. Replace `"postCommands": [...]` with `"postScript": {"phases": [...]}`.
2. Remove `"postCommandAsync"` from JSON request payloads; send HTTP header `Prefer: respond-async` if asynchronous migration is needed and poll `GET /beetle/request/{reqId}`.
3. Update response parsers from `postCommandResults` / `postCommandStatus` / `postCommandRequestId` to `postScriptResults` / `postScriptStatus` / `postScriptRequestId`.

**For `imdl/cloud-model` Users:**

```go
// Before (v0.6.5)
import cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

infraReq := cloudmodel.InfraReq{
    Name:            "my-infra",
    InstallMonAgent: "no",
    PostCommands: []cloudmodel.PostCommandReq{
        {
            InfraCmdReq: cloudmodel.InfraCmdReq{
                UserName: "cb-user",
                Command:  []string{"uname -a"},
            },
        },
    },
}
```

```go
// After (v0.6.6)
import cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

infraReq := cloudmodel.InfraReq{
    Name: "my-infra",
    PostScript: &cloudmodel.PostScriptReq{
        Phases: []cloudmodel.PostScriptPhase{
            {
                InfraCmdReq: cloudmodel.InfraCmdReq{
                    UserName: "cb-user",
                    Command:  []string{"uname -a"},
                },
            },
        },
    },
}
```

### Search Pattern

```bash
# Find deprecated postCommand* usages in your codebase
grep -rnE "postCommand|PostCommand" --include="*.json" --include="*.go" --include="*.ts" --include="*.tsx"
```

---

## Breaking Change #2: `installMonAgent` & `monAgentStatus` Removed

### What Changed

**API Endpoints:**

- `POST /beetle/migration/ns/{nsId}/infra` - `targetInfra.installMonAgent` removed from `InfraReq` and `InfraInfo` response
- `POST /beetle/migration/ns/{nsId}/infraWithDefaults` - `installMonAgent` removed from `InfraDynamicReq` and `InfraInfo` response
- `GET /beetle/migration/ns/{nsId}/infra/{infraId}` - `installMonAgent` removed from `InfraInfo`; `monAgentStatus` removed from `node[]` (`NodeInfo`)
- `GET /beetle/summary/target/ns/{nsId}/infra/{infraId}` - `overview.installMonAgent` removed from `TargetInfraOverview` (and target summary Markdown row `Monitoring Agent`)

### Why This Changed

- CB-Tumblebug `v0.14.0` removed legacy `InstallMonAgent` and `MonAgentStatus` fields in favor of unified `initScript`-based observability agent provisioning (`initscript-otel-agent`) and `ntfy`-based node status reporting.

### Migration Action

1. Remove `installMonAgent` from Infra creation JSON payloads and Go `cloudmodel.InfraReq` / `cloudmodel.InfraDynamicReq` initializers.
2. Remove references to `installMonAgent` in `InfraInfo` / `TargetInfraOverview` and `monAgentStatus` in `NodeInfo` / `NodeSummary`.
3. For observability agent setup on newly provisioned nodes, configure `initScript` (`templateId: "initscript-otel-agent"` or custom `scripts`).

---

## Migration Checklist

### Code Changes

- [ ] Search codebase for `postCommands`, `postCommandAsync`, `postCommandResults`, `postCommandStatus`, `postCommandRequestId`, `installMonAgent`, `monAgentStatus`
- [ ] Update Infra creation requests to use `postScript.phases` and remove `installMonAgent` / `postCommandAsync`
- [ ] Update Infra response & target summary consumers for `postScript*` fields and removal of `installMonAgent` / `monAgentStatus`
- [ ] Update `imdl/cloud-model` Go callers (`PostScriptReq`, `PostScriptPhase`, `PostScriptPhaseResult`, `PostScriptStatus`) and rebuild

### Deployment

- [ ] Update `deployments/docker-compose/docker-compose.yaml` images to `cb-tumblebug:0.14.0`, `cb-spider:0.14.0`, `cb-mapui:0.14.0`
- [ ] Sync `deployments/docker-compose/cb-tumblebug/` (`assets/`, `init/`, `scripts/`, `interface/mcp/`)

### Testing

- [ ] Test `POST /beetle/migration/ns/{nsId}/infra` with updated `postScript` schema
- [ ] Test `GET /beetle/summary/target/ns/{nsId}/infra/{infraId}` against target summary consumers

### Optional (New Features)

- [ ] Adopt `initScript` (`InitScriptReq` / `InitScriptStep`) per Infra or per NodeGroup
- [ ] Inspect `initScriptStatus`, `initScriptMessage`, and `initScriptSteps` on `NodeInfo` / `NodeSummary`
- [ ] Use `SpecInfo.availabilityQuota` and `SpecInfo.acceleratorPartitionSize` in GPU/quota-aware workflows

---

## New Capabilities (Non-Breaking)

- **First-boot `initScript` (`InitScriptReq`, `InitScriptStep`)**: Supported on `InfraDynamicReq`, `CreateNodeGroupReq`, and `CreateNodeGroupDynamicReq`, with node-level status reporting (`initScriptStatus`, `initScriptMessage`, `initScriptSteps`) in `NodeInfo` and `NodeSummary`.
- **Template & multi-step `InfraCmdReq`**: `InfraCmdReq` now supports `templateId`, `scripts` (`[]InitScriptStep`), and `env` (`map[string]string`); `command` is required only when neither `templateId` nor `scripts` is provided.
- **Spec Quota & GPU Partitioning**: `SpecInfo` (and `SpecSummary`) adds `acceleratorPartitionSize` (`float32`), and `SpecInfo` adds `availabilityQuota` (`*AvailabilityQuota`).
- **New `PostScriptStatus` enum value**: Added `PostScriptStatusInterrupted = "Interrupted"`.

---

# Part 2: Technical Details & AI Reference

> **Target Audience:** AI assistants, infrastructure operators, troubleshooting

---

## Affected Components Detail

### Beetle API Endpoints (Breaking)

| Method | Endpoint | Breaking Change | Details |
| --- | --- | --- | --- |
| POST | `/beetle/migration/ns/{nsId}/infra` | Request & response fields changed | `targetInfra.postCommands` -> `targetInfra.postScript`, `installMonAgent` & `postCommandAsync` removed |
| POST | `/beetle/migration/ns/{nsId}/infraWithDefaults` | Request & response fields changed | `postCommands` -> `postScript`, `installMonAgent` & `postCommandAsync` removed |
| GET | `/beetle/migration/ns/{nsId}/infra/{infraId}` | Response fields renamed/removed | `postCommand*` -> `postScript*`, `installMonAgent` & `monAgentStatus` removed |
| GET | `/beetle/summary/target/ns/{nsId}/infra/{infraId}` | Response field removed | `overview.installMonAgent` removed |

### `imdl/cloud-model` Package

| Struct | Change | Migration |
| --- | --- | --- |
| `InfraReq` | `PostCommands`, `PostCommandAsync`, `InstallMonAgent` removed; `PostScript *PostScriptReq`, `PostScriptAsync (json:"-")` added | Assign `&PostScriptReq{Phases: ...}` |
| `InfraDynamicReq` | `PostCommands`, `PostCommandAsync`, `InstallMonAgent` removed; `PostScript`, `PostScriptAsync (json:"-")`, `InitScript *InitScriptReq` added | Assign `&PostScriptReq{Phases: ...}` |
| `AddNodeGroupDynamicReq` | `PostCommands`, `PostCommandAsync` removed; `PostScript`, `PostScriptAsync (json:"-")` added | Assign `&PostScriptReq{Phases: ...}` |
| `CreateNodeGroupReq` / `CreateNodeGroupDynamicReq` | `InitScript *InitScriptReq` added (additive) | Optional first-boot script config |
| `InfraInfo` | `InstallMonAgent`, `PostCommandResults`, `PostCommandStatus`, `PostCommandRequestId` removed; `PostScript*` equivalents added | Update response field reads |
| `InfraInfoSummary` | `InstallMonAgent` removed | Remove field read |
| `NodeInfo` / `NodeSummary` | `MonAgentStatus` removed; `InitScript*` fields added | Remove `MonAgentStatus` read |
| `PostScriptReq` | New struct wrapping `Phases []PostScriptPhase` | Required wrapper for `PostScript` |
| `PostScriptPhase` / `PostScriptPhaseResult` / `PostScriptStatus` | Renamed from `PostCommandReq` / `PostCommandPhaseResult` / `PostCommandStatus` (`Interrupted` added) | Update type names |
| `InfraCmdReq` | `TemplateId`, `Scripts`, `Env` added; `Command` tag updated | Additive |
| `SpecInfo` / `SpecSummary` / `AvailabilityQuota` | `AcceleratorPartitionSize` and `AvailabilityQuota` added | Additive |

---

## Troubleshooting

### Issue: JSON request still sends `"postCommands"` or `"installMonAgent"`

**Cause:** Caller is sending a v0.6.5 request schema to v0.6.6 API.

**Solution:**

- Replace `"postCommands": [...]` with `"postScript": {"phases": [...]}`.
- Remove `"installMonAgent"` and `"postCommandAsync"` from the JSON payload.
- Pass HTTP header `Prefer: respond-async` if asynchronous migration execution is desired.

---

## Version Information

| Component | v0.6.5 | v0.6.6 | Notes |
| --- | --- | --- | --- |
| CM-Beetle | v0.6.5 | v0.6.6 | Stabilization & TB v0.14.0 sync |
| CB-Tumblebug | v0.13.8 | v0.14.0 | `postScript`, `initScript`, quota/partitioning, monAgent removal |
| CB-Spider | v0.13.20 | v0.14.0 | Aligned with TB v0.14.0 |
| CB-MapUI | v0.13.14 | v0.14.0 | Aligned with TB v0.14.0 |
