---
mode: agent
model: Claude Sonnet 4.5
description: "Synchronize CB-Tumblebug models and docker-compose files with specified version"
---

# SyncTB - CB-Tumblebug Model Synchronization

Synchronize TB models in `copied-tb-model.go` with the specified CB-Tumblebug version.

## Target Version

${input:version:CB-Tumblebug version (e.g., v0.11.2, v0.12.0, latest)}

## Process Overview

This prompt will help synchronize CB-Tumblebug models by:

1. **Current Version Detection**: Extract current TB version from copied-tb-model.go
2. **Repository Setup**: Clone CB-Tumblebug repository temporarily
3. **Git Diff Analysis**: Execute git diff to identify all changed structs between versions
4. **Struct Dependency Mapping**: Map all existing structs in copied-tb-model.go and their dependencies
5. **Comprehensive Synchronization**: Update ALL structs that exist in copied-tb-model.go and their dependencies
6. **Version Files Update**: Update docker-compose.yaml and go.mod with target version
7. **Docker-Compose Files Sync**: Compare and update docker-compose related files (assets, config, init scripts)
8. **SYNC.md Documentation**: Update SYNC.md with detailed change summary for the target version
9. **Cleanup**: Remove temporary repository and return to original directory
10. **Validation**: Ensure compilation and proper serialization

## Synchronization Principles

**CRITICAL GUIDELINES**:

### 1. Dependency-Based Synchronization Rule

- **ALWAYS** synchronize ALL structs currently present in copied-tb-model.go
- **ONLY** add new structs that are **direct or indirect dependencies** of existing structs
- **NEVER** add standalone new structs that have no dependency chain to existing structs
- **FOLLOW dependency chains**: If existing struct A uses new struct B, and B uses new struct C, include both B and C

### 2. Struct Dependency Chain Analysis

- Map ALL existing structs in copied-tb-model.go before synchronization
- For each existing struct, identify ALL field types that reference other structs
- Trace dependency chains: `ExistingStruct → NewDependency → SubDependency → ...`
- **REJECT** new structs that cannot be traced back to any existing struct through dependency chains

### 3. Operations Scope

- **HEADER UPDATE**: Always update the version header with target version, commit hash, and date, regardless of struct changes (MANDATORY)
- **UPDATE**: Modify existing structs to match target version (always required)
- **CREATE**: Add new structs ONLY if they are dependencies of existing/updated structs
- **DELETE**: Remove structs that no longer exist in target version (with impact analysis)
- **RENAME**: Handle struct name changes (e.g., Tb prefix removal) with complete replacement
- **REMAP**: Update field type references when struct names change

### 4. Dependency Chain Filtering

- **INCLUDE**: New structs referenced in fields of existing structs
- **INCLUDE**: New structs referenced in fields of already-included dependency structs
- **EXCLUDE**: New structs that exist in CB-Tumblebug but have no dependency path to existing cm-model structs
- **EXCLUDE**: Standalone new functionality that doesn't integrate with existing structs

### 5. Struct Name Change Detection (CRITICAL)

**NEW REQUIREMENT**: Detect and handle struct name changes (especially Tb prefix removal):

- **Pattern Recognition**: Identify when existing struct names have been renamed in CB-Tumblebug
- **Mapping Strategy**: Create mapping between old names (e.g., `TbMciReq`) and new names (e.g., `MciReq`)
- **Replacement Strategy**: Replace struct definitions entirely when name changes are detected
- **Comment Preservation**: Maintain all existing TB-sourced field documentation during name changes
- **Reference Updates**: Update all struct type references within other structs when names change

#### Common Naming Patterns to Detect:

- **Tb Prefix Removal**: `TbStructName` → `StructName`
- **Functional Renaming**: `TbVmDynamicReq` → `CreateSubGroupDynamicReq`
- **Field Renaming**: `CommonSpec` → `SpecId`, `CommonImage` → `ImageId`

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
  - `assets/assets.dump.gz`: Compare MD5 checksums
  - `assets/cloudimage.csv`: Check for new/updated image entries
  - `assets/k8sclusterinfo.yaml`: Verify K8s version updates
  - `assets/cloudinfo.yaml`, `cloudspec.csv`: Check for CSP updates
  - `conf/cloud_conf.yaml`: Review configuration changes
  - `init/*.py`, `init/*.sh`: Check initialization script updates
  - `scripts/*.sh`: Verify operational script changes

- **Use `run_in_terminal`**: Copy updated files when needed:
  ```bash
  # Example: Update specific files
  cp $TB_PATH/assets/cloudimage.csv $BEETLE_PATH/assets/
  cp $TB_PATH/assets/k8sclusterinfo.yaml $BEETLE_PATH/assets/
  ```

### Step 7: SYNC.md Documentation

Document all docker-compose file changes in SYNC.md:

- **Use `read_file`**: Read current SYNC.md to understand format
- **Use `replace_string_in_file`** or **`multi_replace_string_in_file`**: Add new version section at the top:

  ```markdown
  ## v${input:version} Sync (YYYY-MM-DD)

  Based on TB v${input:version} `[commit_short_hash]` (tagged release).

  | File                    | Action                                |
  | ----------------------- | ------------------------------------- |
  | `assets/assets.dump.gz` | **Updated** — MD5 changed to `[hash]` |
  | `assets/cloudimage.csv` | **Updated** — [describe changes]      |
  | ...                     | ...                                   |
  ```

- **Change Categories**:
  - **Updated**: File content changed, copy from TB
  - **No change**: File identical between versions
  - **New**: File added in TB, consider adding
  - **Removed**: File deleted in TB, consider removing

- **Use `get_errors`**: Verify markdown syntax if possible

### Step 8: Cleanup and Validation

- **Use `run_in_terminal`**: Remove cloned CB-Tumblebug repository: `rm -rf /tmp/sync-tb-${input:version}/`
- **Use `run_in_terminal`**: Return to cm-beetle directory
- **Use `get_errors`**: Compile and validate synchronized models
- **Use `run_in_terminal`**: Execute dependency analysis: `python3 scripts/analyze_dependencies.py`

## Synchronization Rules

#### A. Mandatory Synchronization Rules

**Rule 1: Update ALL Existing Structs**

- **MUST** update every struct that exists in copied-tb-model.go if changed in git diff
- **MUST** include all field additions, modifications, and deletions
- **MUST** handle struct name changes (especially Tb prefix removal)
- **NO** exceptions for "complexity" or subjective necessity judgments

**Rule 2: Struct Name Change Handling (CRITICAL)**

- **MUST** detect when existing struct names have changed in CB-Tumblebug
- **PATTERN**: Look for `type TbStructName` → `type StructName` (Tb prefix removal)
- **REPLACEMENT**: Completely replace old struct definition with new one
- **REFERENCES**: Update ALL field type references when struct names change
- **PRESERVE**: Maintain all existing TB-sourced documentation during name changes

**Rule 3: Include ONLY Dependency Chain Structs**

- **MUST** add new struct types referenced by existing structs (direct dependencies)
- **MUST** add new struct types referenced by direct dependencies (indirect dependencies)
- **MUST** include nested types, array element types, pointer target types only if they connect to existing structs
- **EXCLUDE** new structs that have no dependency path to any existing struct

**Rule 4: Dependency Chain Validation Process**

For each struct found in CB-Tumblebug git diff:

1. **Name Change Detection**: Check if existing cm-model struct has been renamed (e.g., `TbMciReq` → `MciReq`)
2. **Trace Back**: Can this struct be reached from any existing cm-model struct through field references?
3. **Dependency Path**: Is there a chain: `ExistingStruct → ... → NewStruct`?
4. **Decision**: Include ONLY if dependency path exists or if it's a renamed existing struct, otherwise EXCLUDE

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

#### B. Version Header Update

**CRITICAL: This step is MANDATORY and must be performed even if no struct changes are detected.**

Update the header comment in copied-tb-model.go to include commit hash:

```go
// * To avoid circular dependencies, the following structs are copied from the cb-tumblebug framework.
// TODO: When the cb-tumblebug framework is updated, we should synchronize these structs.
// * Version: CB-Tumblebug ${input:version} (commit: [full_commit_hash])
// * Synchronized: [YYYY-MM-DD] (include notable changes or PR references)
```

**Note**: Do NOT include individual struct path comments. Focus on clear struct names and descriptions only.

#### C. Complete Field Synchronization

For EVERY struct that exists in copied-tb-model.go AND appears in git diff:

1. **Git Diff as Source**: Use ONLY git diff output for struct changes (single source of truth)
2. **Name Change Detection**: Check if struct name has changed (e.g., `TbMciReq` → `MciReq`)
3. **Complete Replacement**: If name changed, replace entire struct definition with new name and content
4. **Field Additions**: Add ALL new fields exactly as shown in git diff `+` lines
5. **Field Removals**: Remove ALL fields shown in git diff `-` lines
6. **Field Modifications**: Update ALL field types, tags, and comments based on diff changes
7. **Type Reference Updates**: Update field types when referenced struct names change
8. **Validation Tag Updates**: Apply ALL validation tag changes (`validate:"required"`, etc.)
9. **JSON Tag Updates**: Update ALL JSON serialization tags (`json:"fieldName"`, `omitempty`)
10. **Example Updates**: Update ALL struct tag examples to match TB source
11. **Comment Preservation**: Maintain ALL existing Tumblebug field documentation and examples
12. **Header Update**: Update version header with target version and commit hash
13. **Documentation**: Preserve clear struct names and descriptions without path references

#### D. Dependency Struct Addition

For NEW structs referenced by existing structs that appear in git diff:

1. **Complete Addition**: Add the ENTIRE new struct definition from CB-Tumblebug source
2. **All Fields**: Include ALL fields with complete documentation
3. **Proper Placement**: Add in logical order near related structs
4. **Full Documentation**: Include ALL comments, examples, and validation tags from TB source

#### E. File Operations

Execute file editing operations using VS Code tools:

- **Use `multi_replace_string_in_file`** to apply multiple struct changes from git diff simultaneously (PREFERRED)
- **Use `replace_string_in_file`** for individual struct changes when needed
- **Use `read_file`** to verify changes and ensure proper context
- **Use `get_errors`** to validate Go compilation after changes
- **Use `grep_search`** to verify all structs are properly synchronized
- Maintain proper Go syntax and formatting
- Preserve existing cm-model documentation patterns

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
- [ ] **`read_file`**: Verify SYNC.md updated with new version section
- [ ] **Docker-Compose Files**: Confirm all detected file changes documented in SYNC.md
- [ ] **`get_errors`**: Check for any broken file references in docker-compose configuration

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
- **Dependency Analysis**: Always run `python3 scripts/analyze_dependencies.py` for final validation
- **🚨 CRITICAL SAFEGUARD**: **NEVER DELETE Tumblebug-synchronized field comments** - These contain valuable examples and documentation from CB-Tumblebug source that must be preserved
- **🚨 CLEAN DOCUMENTATION**: **DO NOT** include "// \* Path:" comments - Focus on clear struct names and descriptions only

## Execution Workflow

### Phase 1: Dependency Chain Analysis

1. **`read_file`**: Read current version from copied-tb-model.go header
2. **`grep_search`**: Inventory ALL existing struct definitions in copied-tb-model.go
3. **`create_directory`** + **`run_in_terminal`**: Create temporary directory and clone CB-Tumblebug repository
4. **`run_in_terminal`**: Checkout target version in CB-Tumblebug repository
5. **`run_in_terminal`** + **`get_terminal_output`**: Execute comprehensive git diff commands

### Phase 2: Selective Synchronization

6. **`grep_search`**: Analyze git diff output to identify ALL changes to existing structs
7. **Struct Name Change Detection**: Identify struct renames (e.g., `TbMciReq` → `MciReq`)
8. **Dependency Chain Tracing**: For each new struct in git diff, verify dependency path to existing structs OR check if it's a renamed existing struct
9. **`read_file`**: **BEFORE EDITING** - Read current struct documentation to preserve existing comments
10. **`multi_replace_string_in_file`**: Apply ALL struct changes simultaneously using batch editing (PREFERRED)
11. **`multi_replace_string_in_file`**: Update all field type references when struct names change in one operation
12. **`multi_replace_string_in_file`**: Add ONLY dependency-connected new structs with complete definitions
13. **`replace_string_in_file`**: Update version header and ALL source path comments
14. **`get_errors`**: Validate Go syntax and compilation after each major change

### Phase 2.5: Version Files Update

14.1. **`run_in_terminal`**: Extract cb-spider and cb-mapui versions from TB's docker-compose.yaml
14.2. **`read_file`**: Read current version from go.mod
14.3. **`replace_string_in_file`**: Update cb-tumblebug version in go.mod require section
14.4. **`run_in_terminal`**: Run `go mod tidy` to update dependencies
14.5. **`read_file`**: Read current versions from CM-Beetle's docker-compose.yaml
14.6. **`replace_string_in_file`**: Update cb-tumblebug image version in docker-compose.yaml
14.7. **`replace_string_in_file`**: Update cb-spider image version (using extracted version)
14.8. **`replace_string_in_file`**: Update cb-mapui image version (using extracted version)
14.9. **`get_errors`**: Verify no Go module errors after version update

### Phase 3: Docker-Compose Sync & Documentation

15. **`run_in_terminal`**: Compare docker-compose deployment files
16. **`run_in_terminal`**: Copy updated files when changes detected
17. **`replace_string_in_file`**: Add new version section to SYNC.md

### Phase 4: Cleanup & Validation

18. **`run_in_terminal`**: Remove temporary CB-Tumblebug repository
19. **`get_errors`**: Run final compilation validation
20. **`grep_search`**: Verify NO orphaned structs exist
21. **`run_in_terminal`**: Execute dependency analysis: `python3 scripts/analyze_dependencies.py`

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
