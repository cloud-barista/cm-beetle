---
applyTo: "analyzer/**"
---

# Analyzer Module Instructions

This directory (`analyzer/`) contains the **File System Analysis & Metadata Extraction Module**.
It is a self-contained Go module with its own `go.mod`.

## Key Constraints

- **Independent Module:** Treat this as a separate library.
- **No Upstream Dependencies:** Do NOT import packages from `cm-beetle/pkg/...` or `cm-beetle/cmd/...`.
- **Dependency Management:** Add dependencies to `analyzer/go.mod`, not the root `go.mod`.

## Functionality

- **Directory Scanning:** Recursively list directory contents (`ListDirectory`).
- **Metadata Extraction:** Analyze file properties (size, permissions, modification time).
- **Pattern Filtering:** Filter files based on glob patterns (include/exclude).
- **Usage:** Primarily used to analyze source infrastructure before migration.
