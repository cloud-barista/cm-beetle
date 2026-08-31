---
mode: agent
model: Claude Sonnet 4.5
description: "Synchronize CB-Tumblebug models and docker-compose files with specified version"
---

# SyncTB - CB-Tumblebug Model and Deployment Synchronization

Synchronize TB models in `copied-tb-model.go` and `copied-tb-k8s-model.go`, along with deployment files (`assets/`, `conf/`, `init/templates/`, `scripts/`, `interface/mcp/`), with the specified CB-Tumblebug version.

## Target Version

${input:version:CB-Tumblebug version (e.g., v0.12.30, v0.13.1, main, latest)}

## Process Overview

This prompt will help synchronize CB-Tumblebug models and deployment assets by:

1. **Current Version Detection**: Extract current TB version from `copied-tb-model.go` and `copied-tb-k8s-model.go` header comments
2. **Repository Setup**: Clone or navigate to CB-Tumblebug repository
3. **Git Diff Analysis**: Execute git diff to identify all changed structs and assets between versions
4. **Struct Dependency Mapping**: Map all existing structs in `copied-tb-model.go` and `copied-tb-k8s-model.go`
5. **Comprehensive Synchronization**: Update ALL structs in `copied-tb-model.go` (Infra, Spec, Image, Disk, SecurityGroup, SSHKey, VNet, NLB) and `copied-tb-k8s-model.go` (K8sCluster, K8sNodeGroup, Token/Kubeconfig)
6. **Version Files Update**: Update `docker-compose.yaml` and `go.mod` with target version
7. **Docker-Compose Files Sync**: Compare and update docker-compose related files:
   - `assets/` (DB dump, `k8sclusterinfo.yaml`, `diskinfo.yaml`, `rdbmsinfo.yaml`, `cloudinfo.yaml`, etc.)
   - `init/templates/` (standardized `infra-*.json`, `k8scluster-across.json`, `sg-*.json`, `vnet-*.json`)
   - `scripts/` (`restore-assets.sh`, `backup-assets.sh`, `scripts/lib/pg-backend.sh`)
   - `interface/mcp/` (`tb-mcp.py`, `Dockerfile`, `README.md`)
8. **SYNC.md Documentation**: Update `SYNC.md` with detailed change summary for the target version
9. **CM-Beetle Breaking Changes Documentation**: If the sync introduces a breaking change for Beetle API/imdl consumers, document it in `docs/changes/BREAKING_CHANGES_v{cm_beetle_version}.md`
10. **Validation**: Ensure Go compilation (`go build ./imdl/... ./cmd/... ./pkg/...`) and compose config validation (`docker compose config`)

## Synchronization Rules

**CRITICAL GUIDELINES** — this is the single source of truth for what to include, exclude, rename, and
update. Do not duplicate these rules elsewhere in this prompt.

### 1. Dependency Chain Rule

- **ALWAYS** synchronize ALL structs currently present in copied-tb-model.go that changed in git diff
- **ONLY** add new structs that are **direct or indirect dependencies** of existing structs (trace the
  chain: `ExistingStruct → NewDependency → SubDependency → ...`)
- **NEVER** add standalone new structs that have no dependency chain to an existing struct
- For each struct found in the git diff, decide with this process:
  1. **Name Change Detection**: Is this an existing cm-model struct that was renamed (e.g., `TbMciReq` → `MciReq`)?
  2. **Trace Back**: Can this struct be reached from any existing cm-model struct through field references?
  3. **Decision**: Include ONLY if a dependency path exists or it's a renamed existing struct; otherwise EXCLUDE

**Example Dependency Chain Analysis:**

```go
// ✅ INCLUDE: TbMciInfo (existing, renamed to MciInfo) → MciCreationErrors (existing) → VmCreationError (existing)
// ✅ INCLUDE: Name change detected and dependency chain exists

// ✅ INCLUDE: TbMciReq (existing, renamed to MciReq) → TbVmReq (existing, renamed to VmReq)
// ✅ INCLUDE: Both structs renamed but dependency chain maintained

// ❌ EXCLUDE: ReviewMciDynamicReqInfo (standalone new struct)
// ❌ EXCLUDE: No existing struct references this new struct

// ✅ INCLUDE: CreateSubGroupDynamicReq ← IF this is renamed TbVmDynamicReq (existing struct)
// ❌ EXCLUDE: CreateSubGroupDynamicReq ← IF this is completely new struct with no dependency path
```

### 2. Struct Rename Handling

- **Detect** struct renames in CB-Tumblebug, especially Tb prefix removal: `type TbStructName` → `type StructName`
- **Replace** the old struct definition entirely with the new name and content (not a partial edit)
- **Update** ALL field type references in other structs when a referenced struct's name changes
- **Preserve** all existing TB-sourced field documentation and examples during the rename

**Common Naming Patterns to Detect:**

- **Tb Prefix Removal**: `TbStructName` → `StructName`
- **Functional Renaming**: `TbVmDynamicReq` → `CreateSubGroupDynamicReq`
- **Field Renaming**: `CommonSpec` → `SpecId`, `CommonImage` → `ImageId`

### 3. Operations Scope

- **HEADER UPDATE** (MANDATORY, even with no struct changes): update the version header with target
  version, commit hash, and date — see format in Section 4 below
- **UPDATE**: modify every existing struct that changed in git diff — field additions, removals, type
  changes, validation tag changes, JSON tag changes, and updated examples (no exceptions for
  "complexity" or subjective necessity judgments)
- **CREATE**: add new structs ONLY if they are dependencies of existing/updated structs (Rule 1)
- **DELETE**: remove structs that no longer exist in the target version (with impact analysis)
- **RENAME**: handle struct name changes per Rule 2
- **NEVER DELETE** existing Tumblebug-synchronized field comments/examples — they are valuable
  documentation from the TB source
- **DO NOT** include "// \* Path:" comments — clear struct names and descriptions are enough

### 4. Version Header Update Format

Update the header comment in copied-tb-model.go to include the commit hash:

```go
// * To avoid circular dependencies, the following structs are copied from the cb-tumblebug framework.
// TODO: When the cb-tumblebug framework is updated, we should synchronize these structs.
// * Version: CB-Tumblebug ${input:version} (commit: [full_commit_hash])
// * Synchronized: [YYYY-MM-DD] (include notable changes or PR references)
```

### 5. New Dependency Struct Addition

For new structs referenced by existing structs (Rule 1) that appear in git diff:

- Add the ENTIRE new struct definition from the CB-Tumblebug source, with ALL fields
- Include ALL comments, examples, and validation tags from the TB source
- Place it in logical order near the related structs

### 6. Complete Field Synchronization Checklist

For every struct that exists in copied-tb-model.go and appears in the git diff, using ONLY the git
diff as the source of truth:

- [ ] Add ALL new fields exactly as shown in git diff `+` lines
- [ ] Remove ALL fields shown in git diff `-` lines
- [ ] Update ALL field types, tags, and comments per the diff
- [ ] Update field types when a referenced struct's name changed (Rule 2)
- [ ] Apply ALL validation tag changes (`validate:"required"`, etc.)
- [ ] Update ALL JSON serialization tags (`json:"fieldName"`, `omitempty`)
- [ ] Update ALL struct tag examples to match the TB source
- [ ] Preserve ALL existing Tumblebug field documentation and examples

## Tool Usage Guide

### Primary File Operations

- **`read_file`**: Read current TB version from copied-tb-model.go header and examine existing structs
- **`replace_string_in_file`**: Apply individual struct field changes, update version headers, and preserve existing documentation
- **`multi_replace_string_in_file`**: Apply multiple struct changes simultaneously for efficiency (PREFERRED for batch updates)
- **`get_errors`**: Validate Go compilation after synchronization changes

### Repository and Git Operations

- **`run_in_terminal`**: Execute git commands for cloning, checkout, and diff operations
- **`get_terminal_output`**: Retrieve git diff output and command results for analysis
- **`create_directory`**: Create temporary directories for CB-Tumblebug repository cloning

### Code Analysis and Search

- **`grep_search`**: Search for specific struct names, validation tags, and field patterns
- **`file_search`**: Locate model files and identify synchronization targets
- **`list_dir`**: Navigate repository structure and verify cleanup operations

### Dependency Analysis and Validation

- **`run_in_terminal`**: Execute `python3 scripts/analyze_dependencies.py` for comprehensive dependency analysis
- **Dependency Validation**: Verify struct relationships and ensure proper dependency chains are maintained
- **Orphan Detection**: Identify any standalone structs that lack proper dependency connections
- **Relationship Mapping**: Analyze field type references and struct usage patterns

## Detailed Workflow

### Step 1: Current State Assessment

- **Use `read_file`** to parse current TB version from [copied-tb-model.go](../../imdl/cloud-model/copied-tb-model.go) header comment
- **Use `run_in_terminal`** to save current working directory (`pwd`)
- **Use `create_directory`** to create temporary directory for CB-Tumblebug repository

### Step 2: Repository Operations

- **Use `run_in_terminal`**: Clone CB-Tumblebug repository: `git clone https://github.com/cloud-barista/cb-tumblebug.git`
- **Use `run_in_terminal`**: Navigate to cloned repository (`cd cb-tumblebug`)
- **Use `read_file`**: Identify current version (from copied-tb-model.go header)
- **Use `run_in_terminal`**: Checkout target version: `git checkout ${input:version}`

### Step 3: Git Diff Analysis

Execute git diff commands directly:

- **Use `run_in_terminal`**: Run: `git diff [current_version]..${input:version} -- src/core/model/` in the CB-Tumblebug repository
- **Use `get_terminal_output`**: Capture and analyze diff output line by line
- **Use `grep_search`**: Parse struct modifications from diff hunks
- **CRITICAL**: Check for struct name changes (especially Tb prefix removal patterns)
- **Pattern Detection**: Look for rename patterns like `type TbStructName` → `type StructName`
- Focus on files containing models used in copied-tb-model.go

### Step 4: Model Synchronization

Directly apply identified changes to copied-tb-model.go:

- **CRITICAL**: Use ONLY git diff output as the source of truth for all struct changes
- **Single Source**: copied-tb-model.go is the only maintained source for TB model definitions
- **Use `replace_string_in_file`** to update struct definitions
- Apply field additions, removals, and type changes from git diff
- Update validation tags and JSON serialization tags
- Update version header with target version and commit hash
- Preserve cm-model specific documentation enhancements

### Step 5: Version Files Update

Update version references in docker-compose.yaml and go.mod:

- **Use `run_in_terminal`**: Extract service versions from CB-Tumblebug's docker-compose.yaml:

  ```bash
  cd /tmp/sync-tb-${input:version}/cb-tumblebug

  # Check if docker-compose.yaml exists
  if [ -f "docker-compose.yaml" ]; then
    # Extract cb-spider version
    SPIDER_VERSION=$(grep -A 1 "cb-spider:" docker-compose.yaml | grep "image:" | sed 's/.*cloudbaristaorg\/cb-spider:\([0-9.]*\).*/\1/')

    # Extract cb-mapui version
    MAPUI_VERSION=$(grep -A 1 "cb-mapui:" docker-compose.yaml | grep "image:" | sed 's/.*cloudbaristaorg\/cb-mapui:\([0-9.]*\).*/\1/')

    echo "Extracted from TB docker-compose.yaml:"
    echo "  CB-Spider version: $SPIDER_VERSION"
    echo "  CB-MapUI version: $MAPUI_VERSION"
  else
    echo "Warning: docker-compose.yaml not found in CB-Tumblebug repository"
    echo "Please check CB-Tumblebug release notes manually for compatible versions"
  fi
  ```

- **Use `read_file`**: Read current versions from CM-Beetle's docker-compose.yaml

- **Use `replace_string_in_file`**: Update cb-tumblebug image version:

  ```yaml
  cb-tumblebug:
    image: cloudbaristaorg/cb-tumblebug:${input:version}
  ```

- **Use `replace_string_in_file`**: Update cb-spider image version (if extracted successfully):

  ```yaml
  cb-spider:
    image: cloudbaristaorg/cb-spider:$SPIDER_VERSION
  ```

- **Use `replace_string_in_file`**: Update cb-mapui image version (if extracted successfully):

  ```yaml
  cb-mapui:
    image: cloudbaristaorg/cb-mapui:$MAPUI_VERSION
  ```

- **Use `read_file`**: Read current version from go.mod

- **Use `replace_string_in_file`**: Update cb-tumblebug dependency in go.mod:

  ```go
  require (
      github.com/cloud-barista/cb-tumblebug ${input:version}
  ```

- **Use `run_in_terminal`**: Verify go.mod changes and tidy dependencies:

  ```bash
  cd /home/ubuntu/dev/cloud-barista/cm-beetle
  go mod tidy
  ```

- **Use `get_errors`**: Check for any Go module dependency errors

### Step 6: Docker-Compose Files Synchronization

Compare and synchronize docker-compose deployment files:

- **Use `run_in_terminal`**: Compare assets files:

  ```bash
  cd /tmp/sync-tb-${input:version}/cb-tumblebug
  TB_PATH=$(pwd)
  BEETLE_PATH=/home/ubuntu/dev/cloud-barista/cm-beetle/deployments/docker-compose/cb-tumblebug

  # Check each directory for changes
  for dir in assets conf init scripts; do
    diff -qr $BEETLE_PATH/$dir $TB_PATH/$dir 2>&1 | grep "differ\|Only in"
  done

  # Check MCP files (different path)
  diff -qr $BEETLE_PATH/interface/mcp $TB_PATH/src/interface/mcp 2>&1 | grep "differ\|Only in"
  ```

- **Use `run_in_terminal`**: Review specific file differences:

  ```bash
  # Example: Check cloudimage.csv changes
  diff -u $BEETLE_PATH/assets/cloudimage.csv $TB_PATH/assets/cloudimage.csv
  ```

- **File-by-File Analysis**:
  - `assets/assets.dump.gz` & `.info`: Compare MD5 checksums and sync 39MB dump
  - `assets/cloudimage.csv`, `cloudspec.csv`: Check for new/updated image/spec entries
  - `assets/k8sclusterinfo.yaml`: Verify K8s version updates
  - `assets/diskinfo.yaml`: Verify disk specifications and IOPS metadata
  - `assets/rdbmsinfo.yaml`: Verify managed RDBMS multi-cloud matrices
  - `assets/cloudinfo.yaml`, `extractionpatterns.yaml`: Check for CSP updates
  - `init/templates/`: Ensure 1:1 match with standardized `infra-*.json`, `k8scluster-across.json`, `sg-*.json`, `vnet-*.json` templates
  - `scripts/lib/pg-backend.sh`: Essential PostgreSQL backend detection library used by restore-assets.sh
  - `scripts/restore-assets.sh`, `scripts/backup-assets.sh`: Asset backup/restore scripts
  - `interface/mcp/*`: Stateless HTTP MCP server and proxy scripts

- **Use `run_in_terminal`**: Copy updated files when needed:
  ```bash
  # Example: Update specific files
  cp $TB_PATH/assets/*.yaml $BEETLE_PATH/assets/
  cp $TB_PATH/assets/*.csv $BEETLE_PATH/assets/
  cp $TB_PATH/assets/assets.dump.gz* $BEETLE_PATH/assets/
  rm -rf $BEETLE_PATH/init/templates && cp -r $TB_PATH/init/templates $BEETLE_PATH/init/
  mkdir -p $BEETLE_PATH/scripts/lib && cp $TB_PATH/scripts/lib/pg-backend.sh $BEETLE_PATH/scripts/lib/
  cp $TB_PATH/scripts/restore-assets.sh $TB_PATH/scripts/backup-assets.sh $BEETLE_PATH/scripts/
  cp -r $TB_PATH/src/interface/mcp/* $BEETLE_PATH/interface/mcp/
  ```

### Step 7: SYNC.md Documentation

**IMPORTANT**: Since the repository is version-controlled with Git, only keep the latest sync information in SYNC.md. Previous version history can be accessed through Git history.

Document all docker-compose file changes in SYNC.md:

- **Use `read_file`**: Read current SYNC.md to understand format
- **Use `replace_string_in_file`**: **Replace entire content** with the new version section:

  ```markdown
  ## v${input:version} Sync (YYYY-MM-DD)

  Based on TB v${input:version} `[commit_short_hash]` (tagged release).

  | File                    | Action                                |
  | ----------------------- | ------------------------------------- |
  | `assets/assets.dump.gz` | **Updated** — MD5 changed to `[hash]` |
  | `assets/cloudimage.csv` | **Updated** — [describe changes]      |
  | ...                     | ...                                   |
  ```

- **File Replacement Strategy**:
  - Delete ALL previous version sections
  - Keep ONLY the new v${input:version} section
  - Maintain file header/description if present
  - Git history provides access to previous sync records

- **Change Categories**:
  - **Updated**: File content changed, copy from TB
  - **No change**: File identical between versions
  - **New**: File added in TB, consider adding
  - **Removed**: File deleted in TB, consider removing

- **Use `get_errors`**: Verify markdown syntax if possible

### Step 7.5: CM-Beetle Breaking Changes Documentation

Document breaking changes for Beetle API/imdl consumers under `docs/changes/`. Do not create or
update files under `docs/sync/`.

**Note on versioning**: `${input:version}` above refers to the CB-Tumblebug version. This step uses a
separate, unrelated version — CM-Beetle's own release version, referred to below as
`{cm_beetle_version}` (e.g., `v0.5.10`). Never write the TB version into the breaking-changes filename.

First, determine whether this sync introduces a breaking change **from a CM-Beetle API/imdl consumer's
perspective** (not from CB-Tumblebug's perspective):

- A field renamed, removed, or retyped in a struct that Beetle's REST API exposes in a request or
  response body (i.e., an `imdl/cloud-model` struct embedded by a request/response type in
  `pkg/api/rest/controller/*.go`)
- A CB-Tumblebug behavior change that changes what an existing Beetle endpoint does at runtime, even
  without a field-level change (e.g., a changed default value)
- A required deployment change (e.g., a new dependent service, a new required environment variable)

**If none of the above apply** (the sync only adds new structs/fields with no impact on existing Beetle
endpoints), **skip this step** — do not create a breaking-changes document for a purely additive sync.

**Worked example:**

```text
// ✅ BREAKING: CreateNodeGroupDynamicReq.NodeUserPassword field removed
// ✅ BREAKING: This struct is embedded in a Beetle request body → API consumers must update requests

// ✅ BREAKING: TB's auto-generated security group now opens only TCP 22 (was all ports)
// ✅ BREAKING: No field changed, but callers of an existing Beetle endpoint get different runtime
//    behavior with the same request body → document as a Logic Change (see structure below)

// ❌ NOT BREAKING: New `PostCommandStatus` enum added, referenced only by a new optional field
// ❌ NOT BREAKING: Existing Beetle requests/responses are unaffected → skip this step
```

**If a breaking change is found:**

- **Use `run_in_terminal`**: Check the current CM-Beetle release (`git describe --tags --abbrev=0`) and
  confirm the target `{cm_beetle_version}` with the user (e.g., current `v0.5.9` → target `v0.5.10`)
- **Use `file_search`**: Find the most recent file in `docs/changes/`
- **Use `read_file`**: Read that file in full and use it as the concrete template for structure, tone,
  and conventions (e.g., blank `Required Action` + `> **Note:**` for internal follow-ups) — do not rely
  on the abstract structure description below as a substitute for reading an actual example
- **Use `create_file`** (or edit, if a doc for `{cm_beetle_version}` already exists): Create
  `docs/changes/BREAKING_CHANGES_v{cm_beetle_version}.md`
- **Verify Beetle's actual API surface before writing anything**: read `pkg/api/rest/server.go` for the
  real route registrations and the matching controller file for the actual request/response types.
  **Never assume CB-Tumblebug's route path or field names carry over unchanged** — Beetle may use a
  different route name (e.g., `infraWithDefaults` instead of TB's `infraDynamic`) or expose only a
  subset of a TB struct's fields
- **Structure** (match the template file read above):
  - Header: CM-Beetle release version range, release date, CB-Tumblebug dependency version range
  - **Part 1 — Quick Reference** (for Beetle API users and `imdl/cloud-model` consumers):
    - Breaking Changes Summary table with columns `Category` (API / Logic / Deployment / Model),
      `Affected API / Model` (Beetle's actual route, not TB's), `Impact`, `Required Action`
    - Leave `Required Action` blank (`-`) when there is nothing an API/imdl user needs to do; put
      internal engineering follow-ups in a `> **Note:**` line below the table instead
    - One `## ⚠️ Breaking Change #N` section per breaking item, with Before/After JSON for API
      consumers and Before/After Go using actual `cloudmodel` type names for `imdl` consumers
    - A separate `## ℹ️ Logic Change` section for runtime behavior changes that need no field-level
      migration but still deserve a call-out
    - A Migration Checklist grouped by audience (Code Changes, Deployment, Testing, Optional)
  - **Part 2 — Technical Details & AI Reference**: endpoint/struct diff tables, full model diffs,
    curl/Go examples, troubleshooting, and a Document Summary
- **Use `get_errors`** and a follow-up `read_file` to confirm the doc is well-formed

### Step 8: Cleanup and Validation

- **Use `run_in_terminal`**: Remove cloned CB-Tumblebug repository: `rm -rf /tmp/sync-tb-${input:version}/`
- **Use `run_in_terminal`**: Return to cm-beetle directory
- **Use `get_errors`**: Compile and validate synchronized models
- **Use `run_in_terminal`**: Execute dependency analysis: `python3 scripts/analyze_dependencies.py`

## Final Validation Checklist

After synchronization (use appropriate tools for each validation):

- [ ] **`list_dir`**: Temporary CB-Tumblebug repository removed
- [ ] **`run_in_terminal`**: Working directory restored to cm-model
- [ ] **`get_errors`**: No compilation errors detected
- [ ] **`grep_search`**: All existing structs synchronized with git diff changes
- [ ] **`grep_search`**: All new dependency structs added ONLY if connected to existing structs
- [ ] **`grep_search`**: Verify NO "// \* Path:" comments remain in the file
- [ ] **`read_file`**: Confirm version header includes commit hash and synchronization date
- [ ] **Dependency Chain Verification**: No standalone new structs included without dependency path
- [ ] **`read_file`**: Documentation is preserved and enhanced
- [ ] **Manual Review**: Backward compatibility maintained where possible
- [ ] **`grep_search`**: Source path comments are accurate and reflect target version
- [ ] **`read_file`**: Version header reflects target version with change summary
- [ ] **CRITICAL**: **`grep_search`**: Verify NO orphaned structs exist (all new structs must trace back to existing structs)
- [ ] **CRITICAL**: **Dependency Path Validation**: Each new struct has clear dependency chain to existing cm-model structs
- [ ] **`grep_search`**: Confirm ALL dependency structs are present
- [ ] **CRITICAL**: **`read_file`**: Verify ALL Tumblebug-synchronized field comments and examples are preserved
- [ ] **CRITICAL**: **`grep_search`**: Confirm Path line numbers match actual CB-Tumblebug source file locations
- [ ] **CRITICAL**: **`read_file`**: Ensure no valuable documentation was unintentionally deleted during synchronization
- [ ] **`read_file`**: Verify go.mod cb-tumblebug version updated to target version
- [ ] **`run_in_terminal`**: Confirm `go mod tidy` completed without errors
- [ ] **`read_file`**: Verify docker-compose.yaml cb-tumblebug image version updated to target version
- [ ] **`grep_search`**: Verify docker-compose.yaml cb-spider version matches TB's docker-compose.yaml
- [ ] **`grep_search`**: Verify docker-compose.yaml cb-mapui version matches TB's docker-compose.yaml
- [ ] **`run_in_terminal`**: Execute dependency analysis: `python3 scripts/analyze_dependencies.py`
- [ ] **`read_file`**: Verify SYNC.md updated with new version only (old history removed)
- [ ] **Docker-Compose Files**: Confirm all detected file changes documented in SYNC.md
- [ ] **`get_errors`**: Check for any broken file references in docker-compose configuration
- [ ] **Breaking Change Assessment**: Confirmed whether this sync affects Beetle's actual API endpoints or `imdl` consumers
- [ ] **`read_file`**: If a breaking change was found, `docs/changes/BREAKING_CHANGES_v{cm_beetle_version}.md` reflects Beetle's actual routes/structs (not raw TB routes/structs)
- [ ] **`file_search`**: No new files were created under `docs/sync/`

## Files to Update

- [copied-tb-model.go](../../imdl/cloud-model/copied-tb-model.go)
- [go.mod](../../go.mod) - Update cb-tumblebug dependency version
- [docker-compose.yaml](../../deployments/docker-compose/docker-compose.yaml) - Update cb-tumblebug, cb-spider, cb-mapui image versions
- [SYNC.md](../../deployments/docker-compose/cb-tumblebug/SYNC.md)
- Docker-compose deployment files (conditional, based on detected changes):
  - `deployments/docker-compose/cb-tumblebug/assets/*`
  - `deployments/docker-compose/cb-tumblebug/conf/*`
  - `deployments/docker-compose/cb-tumblebug/init/*`
  - `deployments/docker-compose/cb-tumblebug/scripts/*`
  - `deployments/docker-compose/cb-tumblebug/interface/mcp/*`
- `docs/changes/BREAKING_CHANGES_v{cm_beetle_version}.md` (conditional — only when the sync introduces a
  breaking change for Beetle API/imdl consumers; see Step 7.5)

**Do not** create or update files under `docs/sync/` (e.g., `SYNC_TB_*.md`).

## Reference Guidelines

Follow the patterns and guidelines defined in:

- [copilot-instructions.md](../copilot-instructions.md) - CM-Beetle Project Overview
- [tb-sync.instructions.md](../instructions/tb-sync.instructions.md) - TB Synchronization Guidelines
- [copied-tb-model.go](../../imdl/cloud-model/copied-tb-model.go) - Current synchronized TB models (single source of truth)

**⚠️ CRITICAL**: **ALWAYS** use git diff output as the authoritative source for struct changes during synchronization. Do not rely on documentation or external references for struct definitions.

## Important Notes

- **Maintainer-Only Process**: Only maintainers should initiate TB model synchronization
- **Git-Based Comparison**: Uses git diff for accurate change detection between versions
- **Temporary Repository**: CB-Tumblebug repository is cloned temporarily and cleaned up after use
- **Working Directory Safety**: Process saves and restores original working directory
- **Complete Synchronization**: ALL existing structs MUST be synchronized according to git diff
- **No Arbitrary Filtering**: NEVER skip structs based on subjective complexity judgments
- **Dependency Inclusion**: MUST include ALL dependency structs required by existing structs
- **Documentation Critical**: Maintain comprehensive change documentation
- **SYNC.md Policy**: Keep only the latest version in SYNC.md (delete old history), as Git provides version control
- **Dependency Analysis**: Always run `python3 scripts/analyze_dependencies.py` for final validation
- **Beetle Breaking Changes Doc**: Only create `docs/changes/BREAKING_CHANGES_v{cm_beetle_version}.md` when the
  sync actually breaks something for Beetle API/imdl consumers; skip it for purely additive syncs
- **`docs/sync/` Is Off-Limits**: Never create files under `docs/sync/`
- **🚨 CRITICAL SAFEGUARD**: **NEVER DELETE Tumblebug-synchronized field comments** - These contain valuable examples and documentation from CB-Tumblebug source that must be preserved
- **🚨 CLEAN DOCUMENTATION**: **DO NOT** include "// \* Path:" comments - Focus on clear struct names and descriptions only

## Dependency Analysis Script

The `scripts/analyze_dependencies.py` script provides comprehensive dependency analysis for all structs in the cloudmodel package:

**Key Features**:

- **Struct Inventory**: Lists all structs across copied-tb-model.go, model.go, and vm-infra-info.go
- **Dependency Mapping**: Shows internal dependencies within each struct
- **Reference Tracking**: Identifies which structs are referenced by others
- **Orphan Detection**: Finds unreferenced structs that may be candidates for removal
- **Cross-File Analysis**: Analyzes dependencies between different model files

**Usage Options**:

- `python3 scripts/analyze_dependencies.py` - Basic analysis (recommended for SyncTB)
- `python3 scripts/analyze_dependencies.py --verbose` - Detailed dependency information
- `python3 scripts/analyze_dependencies.py --unused-only` - Show only unreferenced structs

**SyncTB Integration**: The script is automatically executed at the end of the synchronization process to validate that all dependency chains are properly maintained and no orphaned structs exist after the update.
