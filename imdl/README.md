# IMDL - Infrastructure Models

## Overview

This directory contains **infrastructure models** for computing infra migration. These models were originally part of the `cm-model` repository but have been internalized to reduce external dependencies and improve maintainability.

## Directory Structure

```
imdl/
├── go.mod                    # Module definition (github.com/cloud-barista/cm-beetle/imdl)
├── README.md                 # This file
├── cloud-model/              # Cloud infrastructure models
│   ├── model.go              # Recommended cloud infrastructure models
│   ├── copied-tb-model.go    # CB-Tumblebug models (synchronized with specific TB versions)
│   └── vm-infra-info.go      # VM infrastructure information models
└── on-premise-model/         # On-premise infrastructure models
    ├── model.go              # Main on-premise infrastructure models
    ├── server.go             # Server hardware and OS models
    └── network.go            # Network-related models
```

## Model Categories

### 1. Cloud Models (`cloud-model/`)

**Purpose**: Model recommended cloud infrastructure configurations and VM infrastructure information.

**Key Types**:

- `RecommendedVmInfraModel`: Recommended VM infrastructure configuration
- `RecommendedVNet`, `RecommendedSecurityGroup`, `RecommendedVmSpec`: Cloud resource recommendations
- `VmInfraInfo`: VM infrastructure information
- TB-prefixed types (e.g., `TbMciReq`, `TbVmReq`): CB-Tumblebug framework models

**CB-Tumblebug Integration**:

- Models in `copied-tb-model.go` are synchronized with specific CB-Tumblebug versions
- These models are copied (not imported) to avoid circular dependencies
- Use the SyncTB prompt (`.github/prompts/sync-tb.prompt.md`) to update TB models

### 2. On-Premise Models (`on-premise-model/`)

**Purpose**: Model on-premise infrastructure for migration planning.

**Key Types**:

- `OnpremiseInfraModel`: Root model for on-premise infrastructure
- `ServerProperty`: Detailed server specifications (CPU, memory, disk, network, OS)
- `NetworkProperty`: Network configuration including IPv4/IPv6 networks
- Hardware component models: `CpuProperty`, `MemoryProperty`, `DiskProperty`
- Network models: `NetworkInterfaceProperty`, `RouteProperty`, `FirewallRuleProperty`
- `OsProperty`: Operating system information

## Import Usage

Import these models in your code using:

```go
import (
    cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
    onpremmodel "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
)
```

## Maintenance

### Updating CB-Tumblebug Models

To synchronize TB models in `cloud-model/copied-tb-model.go` with a newer CB-Tumblebug version:

1. Use the SyncTB prompt: `.github/prompts/sync-tb.prompt.md`
2. Specify the target TB version (e.g., `v0.12.0`, `latest`)
3. Follow the automated synchronization process
4. Verify build and tests after synchronization

See [tb-sync.instructions.md](../.github/instructions/tb-sync.instructions.md) for detailed synchronization guidelines.

## Design Principles

- **Zero External Dependencies**: This module has no external dependencies, only pure struct definitions
- **JSON Serialization**: All models use `json` tags for serialization
- **Validation Tags**: Include `validate` tags where appropriate
- **Comprehensive Documentation**: Field comments include examples, units, and validation patterns
- **CB-Tumblebug Compatibility**: Maintain compatibility with specific CB-Tumblebug versions

## History

- **Origin**: Migrated from `github.com/cloud-barista/cm-model/infra` (v0.0.21)
- **Rationale**: Internalized to reduce external dependencies and improve version control
- **Module**: `github.com/cloud-barista/cm-beetle/imdl`

## Version Management and Tagging

This module is independently versioned using Git tags with the `imdl/` prefix (e.g., `imdl/v0.1.0`).

### Releasing a New Version

After your changes to `imdl/` have been merged into `upstream/main`, create and push the version tag:

```bash
# 1) Fetch latest upstream and verify merge
git fetch upstream
git log upstream/main --oneline -5

# 2) Tag the merge result on upstream/main
git tag -a imdl/v0.1.0 upstream/main -m "imdl: release v0.1.0"
# Optional (safer if upstream/main has moved):
# git tag -a imdl/v0.1.0 <merge_commit_sha> -m "imdl: release v0.1.0"

# 3) Push tag to upstream
git push upstream imdl/v0.1.0

# 4) Verify tag
git show imdl/v0.1.0
```

> **Note:** Tag `upstream/main` (the merge result), not your old branch commit hash. If another PR is merged before you tag, use the exact merge commit SHA instead of `upstream/main`.

### Updating CM-Beetle Dependency

After creating and pushing the tag, update the main CM-Beetle project to use the new version:

```bash
# 1) Start a new branch for CM-Beetle update
git fetch upstream
# Optional (skip this if working in an existing branch):
# git checkout upstream/main -b feat-beetle-use-imdl-v0.1.0

# 2) Update dependency to the new version
go get github.com/cloud-barista/cm-beetle/imdl@v0.1.0
go mod tidy

# 3) Then follow standard development workflow: verify, test, commit, push, and open PR
```

For detailed instructions on the complete workflow, see [docs/module-import-guide.md](../docs/module-import-guide.md).
