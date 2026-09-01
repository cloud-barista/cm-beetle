# IMDL - Infrastructure Models

## Overview

This directory contains **infrastructure models** for computing infra migration. These models were originally part of the `cm-model` repository but have been internalized to reduce external dependencies and improve maintainability.

### Property Similarity and Sharing Feasibility Analysis

| Resource       | Similarity |   Shared   | Key Reason                                                                                                                |
| -------------- | :--------: | :--------: | ------------------------------------------------------------------------------------------------------------------------- |
| VNet           |    Low     |     ❌     | Source: multiple `CidrBlocks []string` + gateway list; Target: single CIDR + `SubnetInfoList` + `ConnectionName`          |
| Subnet         |  Very Low  |     ❌     | Source has no independent subnet model; target `SubnetReq` is a child of `VNetReq`                                        |
| Security Group |   Medium   |     ❌     | Source: separate `SrcCIDR`/`DstCIDR` + `Action` (allow/deny); Target: single `CIDR`, deny rules unsupported               |
| SSH Key        |    None    |     ❌     | No SSH key model on the source side; target is create/register-only                                                       |
| Spec           |    Low     |     ❌     | Source: physical core count (`Cpus × Threads`), GHz, vendor string; Target: logical vCPU count, TB Spec ID                |
| Image          |    Low     |     ❌     | Source: `PrettyName`, `VersionCodename`; Target: `CspImageName`, `CspImageId`, `OSPlatform`                               |
| Node (VM)      |  Very Low  |     ❌     | Source: direct hardware description; Target: Spec/Image/VNet/SG/SshKey ID reference system                                |
| Object Storage |    High    | Properties | Property types (`CORSRule`, etc.) are shared; top-level models (`SourceObjectStorage`/`TargetObjectStorage`) are separate |

## Directory Structure

```
imdl/
├── go.mod                    # Module definition (github.com/cloud-barista/cm-beetle/imdl)
├── README.md                 # This file
├── cloud-model/              # Cloud infrastructure models (source)
│   ├── model.go              # Recommended cloud infrastructure models
│   ├── copied-tb-model.go    # CB-Tumblebug models (synchronized with specific TB versions)
│   └── infra-info.go         # Infrastructure information models
├── on-premise-model/         # On-premise infrastructure models (source)
│   ├── model.go              # Main on-premise infrastructure models
│   ├── server.go             # Server hardware and OS models
│   └── network.go            # Network-related models
├── storage-model/            # Object storage models (source + target)
│   ├── object-storage.go     # Property building-block structs (BucketFeatureProperty, CORSRule, …)
│   └── model.go              # Top-level models (SourceObjectStorage, TargetObjectStorage, RecommendedObjectStorage)
└── rdbms-model/              # Managed RDBMS models (source + target)
    ├── rdbms-source.go       # Source RDBMS & database properties
    ├── rdbms-target.go       # Target RDBMS instance specs & databases
    ├── copied-tb-rdbms-model.go # CB-Tumblebug RDBMS API models (synchronized with TB)
    └── model.go              # Top-level models (SourceRDBMSModel, RecommendedRDBMSModel, CloudProperty)
```

## Model Categories

### 1. Cloud Models (`cloud-model/`)

**Purpose**: Model recommended cloud infrastructure configurations and infrastructure information.

**Key Types**:

- `RecommendedVmInfraModel`: Recommended infrastructure configuration
- `RecommendedVNet`, `RecommendedSecurityGroup`, `RecommendedVmSpec`: Cloud resource recommendations
- `VmInfraInfo`: Infrastructure information
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

### 3. Storage Models (`storage-model/`)

**Purpose**: Model object storage buckets for both source (as observed) and target (to provision). Core property types are shared between source and target; top-level models are separate.

**Key Types**:

- `SourceObjectStorage`: Source bucket state — composes feature, usage, and metadata properties
- `TargetObjectStorage`: Target bucket spec — composes `BucketSpecProperty`; derived after CSP feature-support validation
- `RecommendedObjectStorage`: Recommendation result and direct input to the migration API
- `CORSRule`: Shared property type composed into both source and target structs

### 4. RDBMS Models (`rdbms-model/`)

**Purpose**: Model managed relational database instances (MySQL, MariaDB) for both source and target cloud environments.

**Key Types**:

- `SourceRDBMS`: Source DB instance state with compute, storage, engine version, and inner database properties
- `TargetRDBMSInstance`: Target cloud RDS configuration derived from recommendation and CSP capability
- `RecommendedRDBMS`: Recommendation result and direct input to the RDBMS migration API
- `RDBMSInfo`: CB-Tumblebug synchronized RDBMS instance information model

## Import Usage

Import these models in your code using:

```go
import (
    cloudmodel    "github.com/cloud-barista/cm-beetle/imdl/cloud-model"
    onpremmodel   "github.com/cloud-barista/cm-beetle/imdl/on-premise-model"
    storagemodel  "github.com/cloud-barista/cm-beetle/imdl/storage-model"
    rdbmsmodel    "github.com/cloud-barista/cm-beetle/imdl/rdbms-model"
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

### Repository Setup and Remote Model

This project follows a three-tier Git remote workflow:

- **`upstream`**: Official organization repository (`github.com/cloud-barista/cm-beetle`)
- **`origin`**: Your personal fork (`github.com/<your-account>/cm-beetle`)
- **`local`**: Your local clone

```
Normal Code Changes: Local Commit → git push origin <branch> → Open PR to upstream/main → Merge
Tagging (Maintainers): git tag -a imdl/vX.Y.Z upstream/main → git push upstream imdl/vX.Y.Z
```

---

### Step-by-Step Release Workflow

#### Step 1: Commit and PR `imdl` Changes

1. Create a local branch from `upstream/main` (or `upstream/dev` if using a development branch):
   ```bash
   git fetch upstream
   git checkout -b feat/update-imdl upstream/main
   ```
2. Commit `imdl/` changes, push to your **`origin`** fork (not `upstream`), and open a PR:
   ```bash
   git add imdl/
   git commit -m "update(imdl): sync models with cb-tumblebug vX.Y.Z"
   git push -u origin feat/update-imdl
   ```
3. Merge the PR into `upstream/main` (or `upstream/dev`).

#### Step 2: Tag the Released `imdl` Version

After the PR is merged into `upstream`, create and push the annotated tag directly to `upstream`:

```bash
# 1) Fetch latest upstream and verify merge
git fetch upstream
git log upstream/main --oneline -5

# 2) Check recent imdl tags to determine the next version
git tag -l "imdl/*" --sort=-v:refname | head -5

# 3) Set next version (e.g., export NEXT_IMDL_TAG="imdl/v0.1.13")
export NEXT_IMDL_TAG=  # Update this value (e.g., export NEXT_IMDL_TAG="imdl/v0.1.13")
echo "Next imdl tag: $NEXT_IMDL_TAG"

# 4) Tag the merge result on upstream/main (or upstream/dev)
git tag -a $NEXT_IMDL_TAG upstream/main -m "imdl: release ${NEXT_IMDL_TAG#imdl/}"

# 5) Push tag directly to upstream
git push upstream $NEXT_IMDL_TAG

# 6) Verify tag
git show $NEXT_IMDL_TAG
```

> **Note:** Tags are pushed directly to `upstream` (not `origin`) so the Go module proxy can resolve them. Pushing a tag creates a tag object and does NOT directly push commits to branches.

#### Step 3: Update CM-Beetle Dependency (`go.mod`)

After publishing the tag, update CM-Beetle's root `go.mod` via a new branch and PR:

```bash
# 1) Start a new branch from upstream/main (or upstream/dev)
git fetch upstream
git checkout -b chore/update-imdl-${NEXT_IMDL_TAG#imdl/} --no-track upstream/main

# 2) Update dependency to the newly tagged version
go get github.com/cloud-barista/cm-beetle/imdl@${NEXT_IMDL_TAG#imdl/}
go mod tidy

# 3) Follow standard development workflow: verify, test, commit, push to your origin fork, and open a PR
```

For detailed instructions on the complete multi-module workflow, see [docs/module-import-guide.md](../docs/module-import-guide.md).
