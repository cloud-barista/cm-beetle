# CB-Tumblebug v0.12.9 Synchronization - Breaking Changes

> **AI-assisted**: Generated with [GitHub Copilot](https://github.com/features/copilot) in VS Code (Model: Claude Sonnet 4.6 · High)

**Date**: 2026-04-28  
**Target Version**: CB-Tumblebug v0.12.9 (commit: 01de7bdea9e54b4000831e683916d2f24c163925)  
**Previous Version**: CB-Tumblebug v0.12.5 (commit: accd857011f30e34196cabc7a1388a8b3e68d4d7)

## Overview

CB-Tumblebug v0.12.9 introduced major breaking changes by renaming all infrastructure-related models:

- **MCI** (Multi-Cloud Infrastructure) → **Infra**
- **VM** (Virtual Machine) → **Node**
- **SubGroup** → **NodeGroup**

This synchronization updates `imdl/cloud-model/copied-tb-model.go` to match these changes, resulting in **breaking changes** throughout the cm-beetle codebase.

## Synchronized Struct Renames

### Primary Infrastructure Structs

| Old Name (v0.12.5)  | New Name (v0.12.9)    | Description                     |
| ------------------- | --------------------- | ------------------------------- |
| `MciReq`            | `InfraReq`            | Infrastructure creation request |
| `MciInfo`           | `InfraInfo`           | Infrastructure information      |
| `MciDynamicReq`     | `InfraDynamicReq`     | Dynamic infrastructure request  |
| `MciCmdReq`         | `InfraCmdReq`         | Infrastructure command request  |
| `MciSshCmdResult`   | `InfraSshCmdResult`   | SSH command result set          |
| `MciCreationErrors` | `InfraCreationErrors` | Infrastructure creation errors  |

### Node/VM Related Structs

| Old Name (v0.12.5)         | New Name (v0.12.9)          | Description                 |
| -------------------------- | --------------------------- | --------------------------- |
| `VmInfo`                   | `NodeInfo`                  | Node (VM) information       |
| `VmCreationError`          | `NodeCreationError`         | Node creation error         |
| `CreateSubGroupReq`        | `CreateNodeGroupReq`        | Node group creation request |
| `CreateSubGroupDynamicReq` | `CreateNodeGroupDynamicReq` | Dynamic node group request  |

### Field Renames in Structs

#### InfraReq (was MciReq)

- `SubGroups []CreateSubGroupReq` → `NodeGroups []CreateNodeGroupReq`
- `PostCommand MciCmdReq` → `PostCommand InfraCmdReq`
- Example: `"mci01"` → `"infra01"`
- Comments: "VMs" → "Nodes", "MCI" → "Infra"

#### InfraInfo (was MciInfo)

- `Vm []VmInfo` → `Node []NodeInfo`
- `NewVmList []string` → `NewNodeList []string`
- `PostCommand MciCmdReq` → `PostCommand InfraCmdReq`
- `PostCommandResult MciSshCmdResult` → `PostCommandResult InfraSshCmdResult`
- `CreationErrors *MciCreationErrors` → `CreationErrors *InfraCreationErrors`

#### InfraDynamicReq (was MciDynamicReq)

- `SubGroups []CreateSubGroupDynamicReq` → `NodeGroups []CreateNodeGroupDynamicReq`
- `PostCommand MciCmdReq` → `PostCommand InfraCmdReq`
- Comments: "SubGroups" → "NodeGroups", "VM" → "Node"

#### CreateNodeGroupReq (was CreateSubGroupReq)

- `SubGroupSize int` → `NodeGroupSize int`
- `VmUserName string` → `NodeUserName string`
- `VmUserPassword string` → `NodeUserPassword string`
- Comments: "SubGroup" → "NodeGroup", "VMs" → "Nodes"

#### CreateNodeGroupDynamicReq (was CreateSubGroupDynamicReq)

- `SubGroupSize int` → `NodeGroupSize int`
- `VmUserPassword string` → `NodeUserPassword string`
- Comments: "SubGroup" → "NodeGroup", "VM" → "Node", "MCI" → "Infra"

#### NodeInfo (was VmInfo)

- `SubGroupId string` → `NodeGroupId string`
- `VmUserName string` → `NodeUserName string`
- `VmUserPassword string` → `NodeUserPassword string`
- Comments: "VM" → "Node"

#### SshCmdResult

- `MciId string` → `InfraId string`
- `VmId string` → `NodeId string`
- `VmIp string` → `NodeIp string`

#### InfraCreationErrors (was MciCreationErrors)

- `VmObjectCreationErrors []VmCreationError` → `NodeObjectCreationErrors []NodeCreationError`
- `VmCreationErrors []VmCreationError` → `NodeCreationErrors []NodeCreationError`
- `TotalVmCount int` → `TotalNodeCount int`
- `SuccessfulVmCount int` → `SuccessfulNodeCount int`
- `FailedVmCount int` → `FailedNodeCount int`

#### NodeCreationError (was VmCreationError)

- `VmName string` → `NodeName string`

#### StatusCountInfo

- Comments: "VMs" → "Nodes"

#### CommandStatusInfo

- Comments: "VM" → "Node"

## REST API Endpoint Changes

In addition to model renames, CB-Tumblebug v0.12.9 changed all MCI-related REST API paths:

| Old Path (v0.12.5)              | New Path (v0.12.9)                  |
| ------------------------------- | ----------------------------------- |
| `POST /ns/{nsId}/mci`           | `POST /ns/{nsId}/infra`             |
| `GET /ns/{nsId}/mci`            | `GET /ns/{nsId}/infra`              |
| `GET /ns/{nsId}/mci/{mciId}`    | `GET /ns/{nsId}/infra/{infraId}`    |
| `DELETE /ns/{nsId}/mci/{mciId}` | `DELETE /ns/{nsId}/infra/{infraId}` |
| `POST /ns/{nsId}/mciDynamic`    | `POST /ns/{nsId}/infraDynamic`      |

These changes caused `404 Not Found` errors at runtime even after the model struct rename was complete.

> **Updated file**: `pkg/client/tumblebug/infra-provisioning.go` (formerly `mci.go`) — all URL strings updated.

## Additional Model Changes

### ImageStatus Constants

Three `ImageStatus` constants were missing from `copied-tb-model.go` and were added:

```go
ImageCreating   ImageStatus = "Creating"
ImageFailed     ImageStatus = "Failed"
ImageDeleting   ImageStatus = "Deleting"
```

Full set: `ImageCreating`, `ImageAvailable`, `ImageFailed`, `ImageUnavailable`, `ImageDeleting`, `ImageDeprecated`, `ImageNA`.

## Files Synchronized

### TB Model Files

- ✅ `imdl/cloud-model/copied-tb-model.go` - All TB structs renamed, synchronized, and `ImageStatus` constants completed
- ✅ `imdl/cloud-model/model.go` - Updated struct references
- ✅ `imdl/cloud-model/vm-infra-info.go` - Updated struct references

### Application Code Files

- ✅ `pkg/core/migration/vm-infra.go` - Updated `MciDynamicReq` → `InfraDynamicReq`, `MciReq` → `InfraReq`, `MciInfo` → `InfraInfo`, all field references
- ✅ `pkg/core/recommendation/vm-infra.go` - Updated all MCI/SubGroup/VM references
- ✅ `pkg/api/rest/controller/migration.go` - Updated request body struct embedding
- ✅ `cmd/test-cli/main.go` - Updated response struct field accesses and status checking logic
- ✅ `pkg/client/tumblebug/infra-provisioning.go` - Updated all REST API endpoint paths (`/mci` → `/infra`)

## Migration Path

### Step 1: Update Core Types

Replace all occurrences in your codebase:

```go
// Struct Names
cloudmodel.MciReq          → cloudmodel.InfraReq
cloudmodel.MciInfo         → cloudmodel.InfraInfo
cloudmodel.MciDynamicReq   → cloudmodel.InfraDynamicReq
cloudmodel.MciCmdReq       → cloudmodel.InfraCmdReq
cloudmodel.MciSshCmdResult → cloudmodel.InfraSshCmdResult
cloudmodel.MciCreationErrors → cloudmodel.InfraCreationErrors

cloudmodel.VmInfo          → cloudmodel.NodeInfo
cloudmodel.VmCreationError → cloudmodel.NodeCreationError
cloudmodel.CreateSubGroupReq → cloudmodel.CreateNodeGroupReq
cloudmodel.CreateSubGroupDynamicReq → cloudmodel.CreateNodeGroupDynamicReq
```

### Step 2: Update Field References

```go
// Field Names
.SubGroups    → .NodeGroups
.SubGroupSize → .NodeGroupSize
.SubGroupId   → .NodeGroupId
.Vm           → .Node
.NewVmList    → .NewNodeList
.VmUserName   → .NodeUserName
.VmUserPassword → .NodeUserPassword
.MciId        → .InfraId
.VmId         → .NodeId
.VmIp         → .NodeIp
.VmObjectCreationErrors → .NodeObjectCreationErrors
.VmCreationErrors → .NodeCreationErrors
.TotalVmCount → .TotalNodeCount
.SuccessfulVmCount → .SuccessfulNodeCount
.FailedVmCount → .FailedNodeCount
.VmName       → .NodeName
```

### Step 3: Update TB Model Imports

When updating `go.mod` to use TB v0.12.9, update `tbmodel` references:

```go
// Old (v0.12.5)
tbmodel.MciReq
tbmodel.MciInfo
tbmodel.MciDynamicReq

// New (v0.12.9)
tbmodel.InfraReq
tbmodel.InfraInfo
tbmodel.InfraDynamicReq
```

### Step 4: Update Model Conversions

```go
// Example conversion update
// Old:
modelconv.ConvertWithValidation[cloudmodel.MciReq, tbmodel.MciReq](req)

// New:
modelconv.ConvertWithValidation[cloudmodel.InfraReq, tbmodel.InfraReq](req)
```

### Step 5: Update Comments and Strings

- Update code comments: "MCI" → "Infra", "VM" → "Node", "SubGroup" → "NodeGroup"
- Update log messages and error strings
- Update API documentation and examples

## Testing Checklist

- [x] All Go files compile without errors (`make build` passes)
- [ ] Unit tests pass
- [x] Integration tests with CB-Tumblebug v0.12.9 work (REST API endpoints verified at runtime)
- [x] API endpoint responses use new field names
- [x] CLI commands work with new struct names
- [x] Documentation reflects new naming

## go.mod Update Required

Update the CB-Tumblebug dependency in `go.mod`:

```go
// Old
github.com/cloud-barista/cb-tumblebug v0.12.5

// New
github.com/cloud-barista/cb-tumblebug v0.12.9
```

## API Compatibility Note

⚠️ **Breaking API Change**: API requests and responses will use new field names:

- JSON field names remain unchanged (backward compatible at JSON level)
- But struct field names in Go code have changed (breaking for Go clients)

## References

- CB-Tumblebug v0.12.9 Release: https://github.com/cloud-barista/cb-tumblebug/releases/tag/v0.12.9
- TB Model Sync Instructions: `.github/instructions/tb-sync.instructions.md`
- SyncTB Prompt: `.github/prompts/sync-tb.prompt.md`

## Summary

This synchronization maintains alignment with CB-Tumblebug's infrastructure model evolution. The rename from MCI/VM to Infra/Node reflects a more general and accurate terminology for multi-cloud infrastructure management. While this introduces breaking changes, it ensures cm-beetle remains compatible with the latest CB-Tumblebug API and model conventions.
