# Scripts Directory

This directory contains utility scripts for analyzing and managing the CM-Beetle project.

## Available Scripts

### 🔍 `analyze_dependencies.py`

Analyzes the struct dependencies across the entire `cloudmodel` package (`copied-tb-model.go`, `model.go`, and `vm-infra-info.go`) and provides comprehensive insights into struct relationships and usage patterns.

#### Features

- **Package-wide Analysis**: Analyzes all Go files in the `imdl/cloud-model` directory
- **Cross-file Dependencies**: Identifies dependencies between structs across different files
- **Independence Detection**: Identifies structs with no dependencies on custom types
- **Usage Tracking**: Checks which structs are referenced by other structs within the package
- **File Location Tracking**: Shows which file each struct is defined in
- **Cleanup Candidates**: Identifies completely unused structs across the entire package

#### Usage

```bash
# Basic analysis (all files in cloudmodel package)
python3 scripts/analyze_dependencies.py

# Detailed analysis with dependency chains
python3 scripts/analyze_dependencies.py --verbose

# Show only unused structs (for cleanup)
python3 scripts/analyze_dependencies.py --unused-only
```

#### Sample Output

```
🔍 CB-Tumblebug Model Dependency Analysis (CloudModel Package)
=================================================================

📊 Statistics:
   Total files analyzed: 3
   copied-tb-model.go: 25 structs, 3 string types
   model.go: 12 structs, 0 string types
   vm-infra-info.go: 4 structs, 0 string types
   Package total: 41 structs, 3 string types, 44 custom types

✅ REFERENCED STRUCTS (used by other structs) [16]:
   📄 From copied-tb-model.go:
      • ConnConfig ← RegionZoneInfo, RegionDetail
      • ImageInfo ← RecommendedVmInfra (model.go)
      • MciReq ← RecommendedVmInfra (model.go)
      ...

🏝️  UNREFERENCED STRUCTS (not used by other structs) [11]:
   📄 From copied-tb-model.go:
      • FirewallRuleReq
      • IdList
      • MigratedVmInfraModel
      ...
```

#### Command Line Options

- `--verbose, -v`: Show detailed dependency information including reference chains and file locations
- `--unused-only, -u`: Show only completely unused structs (useful for cleanup operations)

#### Use Cases

1. **Before CB-Tumblebug Sync**: Understand current dependency structure across all cloudmodel files
2. **Code Cleanup**: Identify unused structs that might be removed from any file
3. **Refactoring**: Understand impact of struct changes across file boundaries
4. **Architecture Review**: Analyze cross-file dependencies and coupling
5. **Documentation**: Generate comprehensive dependency documentation
6. **Quality Assurance**: Verify struct usage consistency across the package

#### Requirements

- Python 3.6+
- Access to `imdl/cloud-model/` directory with Go source files
- Analyzes: `copied-tb-model.go`, `model.go`, `vm-infra-info.go`

#### Technical Details

- Parses Go source files to extract struct definitions and type declarations
- Analyzes field types to identify dependencies on custom types
- Tracks cross-file references to determine struct usage
- Handles complex type references: pointers, arrays, and combinations

### 🔐 `test-field-encryption-and-decryption.sh`

Tests the field-level encryption and decryption functionality for sensitive data.

### 📚 `ask-ai-api-guide.sh`

Helper script for querying AI about API usage and guidelines.

### 📊 `system-info/`

Directory containing system information gathering scripts and utilities.
