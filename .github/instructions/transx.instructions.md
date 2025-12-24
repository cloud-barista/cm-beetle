---
applyTo: "transx/**"
---

# TransX Library Instructions

This directory (`transx/`) contains the **Data Migration Library**.
It is a self-contained Go module with its own `go.mod`.

## Key Constraints

- **Independent Module:** Treat this as a separate library.
- **No Upstream Dependencies:** Do NOT import packages from `cm-beetle/pkg/...` or `cm-beetle/cmd/...`.
- **Dependency Management:** Add dependencies to `transx/go.mod`, not the root `go.mod`.

## Functionality

- Supports `rsync` and `object-storage-api` transfer methods.
- Supports Direct and Relay migration modes.
