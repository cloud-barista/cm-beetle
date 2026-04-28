---
mode: agent
model: Claude Sonnet 4.5
description: "Synchronize CB-Tumblebug models in copied-tb-model.go with specified version"
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
6. **Cleanup**: Remove temporary repository and return to original directory
7. **Validation**: Ensure compilation and proper serialization

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

### 3. Direct Git Diff Execution

Execute git diff commands directly:

- **Use `run_in_terminal`**: Run: `git diff [current_version]..${input:version} -- src/core/model/` in the CB-Tumblebug repository
- **Use `get_terminal_output`**: Capture and analyze diff output line by line
- **Use `grep_search`**: Parse struct modifications from diff hunks
- **CRITICAL**: Check for struct name changes (especially Tb prefix removal patterns)
- **Pattern Detection**: Look for rename patterns like `type TbStructName` → `type StructName`
- Focus on files containing models used in copied-tb-model.go

### Step 4: Direct Model Synchronization

Directly apply identified changes to copied-tb-model.go:

- **CRITICAL**: Use ONLY git diff output as the source of truth for all struct changes
- **Single Source**: copied-tb-model.go is the only maintained source for TB model definitions
- **Use `replace_string_in_file`** to update struct definitions
- Apply field additions, removals, and type changes from git diff
- Update validation tags and JSON serialization tags
- Update version header with target version and commit hash
- Preserve cm-model specific documentation enhancements

### Step 5: Cleanup and Validation

- **Use `run_in_terminal`**: Remove cloned CB-Tumblebug repository: `rm -rf cb-tumblebug/`
- **Use `run_in_terminal`**: Return to cm-model directory
- **Use `get_errors`**: Compile and validate synchronized models

## Analysis Steps

### 1. Current State Analysis

- **Use `read_file`** to extract current TB version from [copied-tb-model.go](../../imdl/cloud-model/copied-tb-model.go) header comment
- **Use `grep_search`** to inventory ALL existing struct definitions in copied-tb-model.go
- **Use `grep_search`** to map struct dependencies and relationships within copied-tb-model.go
- **Use `create_directory`** to set up temporary workspace for CB-Tumblebug repository cloning

### 2. Repository Setup and Git Diff Analysis

**Repository Setup:**

- **Use `create_directory`**: Clone CB-Tumblebug repository in a temporary directory: `/tmp/sync-tb-${input:version}/`
- **Use `run_in_terminal`**: Navigate to cloned repository (`cd /tmp/sync-tb-${input:version}/cb-tumblebug`)
- **Use `run_in_terminal`**: Checkout target version: `git checkout ${input:version}`

**Comprehensive Git Diff Execution:**

- **Use `run_in_terminal`**: Execute comprehensive diff: `git diff [current_version]..${input:version} -- src/core/model/`
- **Use `get_terminal_output`**: Capture complete diff output for analysis
- **Focus**: ALL model files without exclusion:
  - `src/core/model/mci.go` (MCI-related structs)
  - `src/core/model/vnet.go` (VNet-related structs)
  - `src/core/model/sshkey.go` (SSH key structs)
  - `src/core/model/spec.go` (Specification structs)
  - `src/core/model/image.go` (Image-related structs)
  - `src/core/model/securitygroup.go` (Security group structs)
  - `src/core/model/subnet.go` (Subnet structs)
  - `src/core/model/common.go` (Common types and constants)
  - `src/core/model/config.go` (Configuration structs)
  - All other model files that contain struct changes

### 3. Dependency Chain Impact Analysis

**Phase 1: Existing Struct Inventory**

- **Use `grep_search`**: List ALL struct names currently in copied-tb-model.go: `type.*struct`
- Create inventory of current structs: TbMciReq, TbVNetReq, TbMciInfo, etc.
- **Use `semantic_search`**: Map field types and dependencies within each existing struct

**Phase 2: Dependency Chain Mapping**

For each existing struct, identify all struct-type fields:

```bash
# Example dependency mapping for existing structs:
# TbMciInfo → MciSshCmdResult (existing) + CreationErrors (potential new dependency)
# TbMciReq → TbVmReq (existing) + MciCmdReq (existing) + PolicyTypes (potential constants)
# TbVmInfo → Location (existing) + RegionInfo (existing) + ConnConfig (existing) + KeyValue (existing)
```

**Phase 3: Git Diff Analysis with Dependency Focus**

- **Use `run_in_terminal`** + **`get_terminal_output`**: Execute git diff for each model file
- **Priority**: Focus on changes to existing structs first
- **Dependency Tracing**: For each field change in existing structs, identify if it introduces new struct dependencies
- **Chain Following**: If a new dependency is found, recursively analyze its dependencies

**Phase 4: Dependency Chain Validation**

```bash
# For each new struct found in git diff, validate dependency chain:
# 1. Is this struct referenced by any existing struct? → INCLUDE
# 2. Is this struct referenced by any already-included dependency struct? → INCLUDE
# 3. Does this struct have no dependency path to existing structs? → EXCLUDE
```

**Diff Analysis Process:**

- **Use `get_terminal_output`** to capture complete diff output for each file
- **Use `grep_search`** to identify specific struct definitions and field patterns
- Parse diff hunks to identify:
  - Added lines (prefixed with `+`)
  - Removed lines (prefixed with `-`)
  - Context lines for struct identification
  - Context lines for struct identification
    **Git Diff Parsing:**

- **Use `grep_search`** to parse git diff output for:
  - Struct definitions: `type.*struct`
  - Added lines: lines starting with `+`
  - Removed lines: lines starting with `-`
  - Modified struct fields and their types

**Change Classification:**

- Identify changes to structs that exist in copied-tb-model.go
- Identify new struct dependencies introduced by existing struct changes
- Identify removed struct dependencies no longer needed

### 4. Dependency-Based Synchronization Process

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

#### D. Complete Field Synchronization

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

#### E. Dependency Struct Addition

For NEW structs referenced by existing structs that appear in git diff:

1. **Complete Addition**: Add the ENTIRE new struct definition from CB-Tumblebug source
2. **All Fields**: Include ALL fields with complete documentation
3. **Proper Placement**: Add in logical order near related structs
4. **Full Documentation**: Include ALL comments, examples, and validation tags from TB source

#### F. File Operations

Execute file editing operations using VS Code tools:

- **Use `multi_replace_string_in_file`** to apply multiple struct changes from git diff simultaneously (PREFERRED)
- **Use `replace_string_in_file`** for individual struct changes when needed
- **Use `read_file`** to verify changes and ensure proper context
- **Use `get_errors`** to validate Go compilation after changes
- **Use `grep_search`** to verify all structs are properly synchronized
- Maintain proper Go syntax and formatting
- Preserve existing cm-model documentation patterns

### 5. Repository Cleanup

After successful synchronization:

- **Use `run_in_terminal`** to remove the cloned CB-Tumblebug repository
- **Use `list_dir`** to verify cleanup and directory restoration
- **Use `read_file`** to validate final changes in copied-tb-model.go
- Return to original working directory

### 6. Documentation and Version Update

**Header Information Update:**

**CRITICAL: Perform this step even if no other changes were made to the file.**

- [ ] **`run_in_terminal`**: Get target version commit hash: `git rev-parse HEAD`
- [ ] **`replace_string_in_file`**: Update version header with commit hash and synchronization date
- [ ] **CRITICAL**: **DO NOT** include individual struct path comments (// \* Path: src/core/model/...)
- [ ] **CRITICAL**: Focus on clear struct names and descriptions only
- [ ] **`read_file`**: Verify ALL valuable documentation is preserved during synchronization

### 7. Final Validation Checklist

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
- [ ] **`run_in_terminal`**: Execute dependency analysis: `python3 scripts/analyze_dependencies.py`

### 8. Dependency Analysis Report

**Final Step - Dependency Verification:**

- **Use `run_in_terminal`**: Execute `python3 scripts/analyze_dependencies.py` to generate dependency analysis
- **Analyze Output**: Review struct relationships and ensure proper dependency chains are maintained
- **Verify Results**: Confirm no orphaned structs exist and all dependencies are valid
- **Document Findings**: Include any notable dependency changes in the synchronization summary

## Files to Update

- [copied-tb-model.go](../../imdl/cloud-model/copied-tb-model.go)

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

### Complete Synchronization Guidelines

- **Existing Struct Rule**: ALL structs in copied-tb-model.go MUST be updated according to git diff
- **Struct Name Changes**: MUST detect and handle struct renames (especially Tb prefix removal)
- **Reference Updates**: MUST update all field type references when struct names change
- **Dependency Chain Rule**: Add new structs ONLY if they have dependency chains to existing structs
- **No Orphaned Structs**: Do NOT include standalone new structs without dependency paths to existing structs
- **Full Operations**: Perform CREATE (dependencies only), UPDATE (existing structs), DELETE, RENAME operations as needed
- **Complete Documentation**: Include ALL field comments, examples, and validation tags from TB source

### Dependency Chain Guidelines

- **Trace Dependencies**: For each new struct in git diff, verify if it's referenced by existing or dependency-connected structs
- **Name Change Detection**: Check if "new" structs are actually renamed existing structs (e.g., `TbMciReq` → `MciReq`)
- **Follow Chains**: Include multi-level dependencies: `ExistingStruct → NewStruct1 → NewStruct2 → ...`
- **Exclude Orphans**: Reject new structs that cannot be traced back to existing structs
- **Examples of Valid Chains**:
  - `TbMciInfo (existing, renamed to MciInfo) → MciCreationErrors (existing) → VmCreationError (existing)` ✅
  - `TbSpecInfo (existing, renamed to SpecInfo) → NewSpecExtension (new)` ✅
- **Examples of Invalid Chains**:
  - `ReviewMciDynamicReqInfo (standalone new)` ❌
  - `ProvisioningLog (standalone new)` ❌
- **Rename Detection Patterns**:
  - `TbMciReq` → `MciReq` (Tb prefix removal) ✅
  - `TbVmDynamicReq` → `CreateSubGroupDynamicReq` (functional rename) ✅

### Tool Usage Best Practices

- **Terminal Operations**: Use `run_in_terminal` for all git commands and `get_terminal_output` for capturing results
- **File Modifications**: Always use `replace_string_in_file` with sufficient context (3-5 lines before/after)
- **Validation**: Run `get_errors` after each significant change to catch compilation issues early
- **Search Operations**: Combine `grep_search` and `semantic_search` for comprehensive code analysis
- **Safety Checks**: Use `list_dir` and `read_file` to verify operations and cleanup

## Expected Output

1. **Current State Analysis**: Complete inventory of existing structs in copied-tb-model.go
2. **Repository Setup**: Clone CB-Tumblebug repository to temporary directory `/tmp/sync-tb-${input:version}/`
3. **Name Change Detection**: Identify struct renames (especially Tb prefix removal patterns)
4. **Dependency Chain Analysis**: Identify git diff changes and trace dependency chains from existing structs
5. **Change Classification**: Categorize changes into existing struct updates vs. dependency-connected new structs vs. renamed structs
6. **Selective Synchronization**: Apply ALL changes to existing structs, handle renames, and add ONLY dependency-connected new structs
7. **Reference Updates**: Update all field type references when struct names change
8. **Full Documentation Update**: Update ALL path references and version headers
9. **Compilation Verification**: Ensure ALL changes compile without errors
10. **Repository Cleanup**: Removal of temporary CB-Tumblebug repository
11. **Final Validation**: Confirmation that ALL existing structs are synchronized, renames handled, and dependencies complete
12. **Dependency Analysis Report**: Detailed dependency analysis showing struct relationships and validation

## Execution Steps

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

### Phase 3: Dependency Chain Validation

15. **`run_in_terminal`** + **`list_dir`**: Remove temporary CB-Tumblebug repository
16. **`get_errors`** + **`read_file`**: Run final validation in cm-model directory
17. **`grep_search`**: Verify NO orphaned structs exist - all new structs must connect to existing structs or be renamed existing structs
18. **Struct Name Consistency Check**: Verify all struct name changes are applied consistently across all field type references
19. **Dependency Path Review**: Generate summary showing dependency chains for all included new structs and handled renames
20. **Reference Integrity Check**: Confirm all field type references use updated struct names
21. **`run_in_terminal`**: Execute dependency analysis: `python3 scripts/analyze_dependencies.py` for comprehensive dependency validation

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
