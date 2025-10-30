# Analyzer - File Data Migration Module

## Overview

The analyzer module provides file data migration capabilities with directory scanning, metadata extraction, and pattern-based filtering.

## How It Works (User Perspective)

### 1. Browse and Select Directory

Users navigate through directories using a file explorer interface:

```
[UI] Home Directory (/home/user)
  📁 Documents  →  [User clicks]
  📁 Downloads
  📁 Pictures
```

```
[UI] Documents Directory
  📁 Projects  →  [User clicks]
  📁 Archive
  📄 notes.txt
```

```
[UI] Projects Directory  →  [User selects this for migration]
  📁 webapp
  📁 api
  📄 README.md
```

**Behind the scenes**: `ListDirectory()` shows immediate contents at each level.

### 2. Configure Migration Options

Users configure what to migrate:

```
[UI] Migration Configuration
  ☑ Include subdirectories

  Include files:
    ☑ *.txt      (text files)
    ☑ *.md       (markdown files)
    ☑ docs/**    (all files in docs/)

  Exclude files:
    ☑ *.log      (log files)
    ☑ .git/**    (git directory)
    ☑ *.tmp      (temporary files)
```

**Behind the scenes**: `FilterOptions` with include/exclude patterns.

### 3. Generate Migration Plan

System scans and creates a detailed plan:

```
[UI] Migration Plan Summary
  Source: /home/user/Documents/Projects
  Files: 247 files
  Total Size: 15.3 MB

  [View Details] [Start Migration]
```

**Behind the scenes**: `CreateMigrationPlan()` scans directory, applies filters, and collects metadata.

### 4. Review File List

Users can review detailed file information:

```
[UI] File List (247 files)
  📄 README.md          15 KB    2025-10-29 14:23
  📄 config.yaml         2 KB    2025-10-28 09:15
  📁 src/
    📄 main.go          45 KB    2025-10-30 11:05
    📄 utils.go         12 KB    2025-10-27 16:42
  📁 docs/
    📄 guide.md         78 KB    2025-10-25 10:30
```

**Behind the scenes**: Each file has complete metadata (timestamps, size, permissions, owner).

## Key Features

- **Directory browsing**: Navigate hierarchically through directories
- **Metadata extraction**: Collect comprehensive file information (size, timestamps, permissions, owner/group)
- **Pattern filtering**: Include/exclude files using glob patterns (`*.txt`, `data/**`, `**/test/**`)
- **Migration planning**: Generate complete migration plans with file lists and statistics

## File Metadata Information

The analyzer extracts comprehensive metadata for each file and directory. This information is essential for migration planning and analysis.

| Metadata Item     | Description                                                          |
| ----------------- | -------------------------------------------------------------------- |
| **Path**          | Full absolute path of the file or directory                          |
| **Name**          | File or directory name (without path)                                |
| **Size**          | Size in bytes (0 for directories)                                    |
| **IsDir**         | Boolean flag indicating if the item is a directory                   |
| **Mode**          | Permission mode (e.g., "0755", "0644")                               |
| **ModTime**       | Last modification time (when content was changed)                    |
| **AccessTime**    | Last access time (when file was read or opened)                      |
| **ChangeTime**    | Last status change time (when metadata was modified)                 |
| **Owner**         | Owner user ID (UID on Linux/Unix)                                    |
| **Group**         | Group ID (GID on Linux/Unix)                                         |
| **MimeType**      | MIME type of the file (e.g., "text/plain", "application/json")       |
| **Extension**     | File extension including dot (e.g., ".txt", ".json")                 |
| **IsSymlink**     | Boolean flag indicating if the item is a symbolic link               |
| **SymlinkTarget** | Target path if the item is a symbolic link                           |
| **Checksum**      | Optional checksum value (MD5/SHA256) for file integrity verification |

## API Reference

### Core Functions

| Function                                                 | Purpose                                               |
| -------------------------------------------------------- | ----------------------------------------------------- |
| `ListDirectory(path)`                                    | List immediate directory contents (for UI navigation) |
| `CreateMigrationPlan(sourceDir, includeSubDir, filters)` | Generate migration plan with filtering                |
| `ScanDirectory(options)`                                 | Comprehensive directory scan with options             |
| `ExtractFileMetadata(path, collectChecksum)`             | Get detailed file metadata                            |

### Key Data Structures

```go
// Migration plan with file list and statistics
type MigrationPlan struct {
    SourceDir     string
    TotalFiles    int
    TotalSize     int64
    FileList      []FileMetadata
    FilterOptions FilterOptions
}

// File metadata (timestamps, permissions, owner, size, etc.)
type FileMetadata struct {
    Path          string    // Full absolute path
    Name          string    // File or directory name
    Size          int64     // Size in bytes (0 for directories)
    IsDir         bool      // True if directory
    Mode          string    // Permission mode (e.g., "0755")
    ModTime       time.Time // Last modification time
    AccessTime    time.Time // Last access time (if available)
    ChangeTime    time.Time // Last status change time (if available)
    Owner         string    // Owner user ID
    Group         string    // Group ID
    MimeType      string    // MIME type (for files)
    Extension     string    // File extension (e.g., ".txt")
    IsSymlink     bool      // True if symbolic link
    SymlinkTarget string    // Target if symbolic link
    Checksum      string    // Optional checksum (MD5/SHA256)
}

// Filter configuration
type FilterOptions struct {
    IncludePatterns []string  // e.g., []string{"*.txt", "docs/**"}
    ExcludePatterns []string  // e.g., []string{"*.log", ".git/**"}
}
```

## Usage Examples

### Example 1: UI Directory Navigation

```go
// Step 1: Show home directory
result, _ := analyzer.ListDirectory("")
// Display: result.Entries (files and directories)

// Step 2: User clicks "Documents"
result, _ := analyzer.ListDirectory("/home/user/Documents")
// Display: Documents contents

// Step 3: User clicks "Projects"
result, _ := analyzer.ListDirectory("/home/user/Documents/Projects")
// Display: Projects contents
// User selects this directory for migration
```

### Example 2: Complete Migration Workflow

```go
// User has selected: /home/user/Documents/Projects
// User configured filters in UI

filters := analyzer.FilterOptions{
    IncludePatterns: []string{"*.txt", "*.md", "docs/**"},
    ExcludePatterns: []string{"*.log", ".git/**"},
}

// Generate migration plan
plan, _ := analyzer.CreateMigrationPlan(
    "/home/user/Documents/Projects",  // Selected directory
    true,                              // Include subdirectories
    filters,                           // Apply filters
)

// Display to user:
fmt.Printf("Migration Plan:\n")
fmt.Printf("  Files: %d\n", plan.TotalFiles)
fmt.Printf("  Size: %d bytes\n", plan.TotalSize)

// Show file list
for _, file := range plan.FileList {
    fmt.Printf("  %s (%d bytes)\n", file.Path, file.Size)
}
```

## Filter Patterns

### Basic Patterns

- `*.txt` - All text files
- `*.log` - All log files
- `data/*` - Files directly in data/

### Recursive Patterns

- `data/**` - All files under data/ (recursive)
- `**/*.json` - All JSON files anywhere
- `**/test/**` - All files in any test/ directory

### Example Filter Configuration

```go
filters := FilterOptions{
    IncludePatterns: []string{
        "*.txt", "*.md",   // Text and markdown files
        "docs/**",         // All files in docs/
    },
    ExcludePatterns: []string{
        "*.log", "*.tmp",  // Log and temp files
        ".git/**",         // Git directory
        "**/test/**",      // Test directories
    },
}
```

## Testing

Run the example:

```bash
cd analyzer/examples/basic
go run main.go
```

Run tests:

```bash
cd analyzer
go test -v
```

## Platform Support

- **Linux/Unix**: Full support (owner, group, timestamps, permissions)
- **Windows/macOS**: Can be added with platform-specific implementations

## License

Apache 2.0 License
