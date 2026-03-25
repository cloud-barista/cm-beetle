# Importing and Versioning Internal Packages

The `cm-beetle` repository manages several independent Go modules as nested packages. These include:

- **analyzer**: `github.com/cloud-barista/cm-beetle/analyzer`
- **transx**: `github.com/cloud-barista/cm-beetle/transx`
- **deepdiffgo**: `github.com/cloud-barista/cm-beetle/deepdiffgo`

Since these are independent modules (each has its own `go.mod`), they can be imported into other projects separately.

## 1. Importing from Other Projects

To import these packages into another Go project, use the full module path:

```go
import (
    "github.com/cloud-barista/cm-beetle/analyzer"
    "github.com/cloud-barista/cm-beetle/transx"
)
```

## 2. Referencing Versions (Git Tags)

Go's module system uses Git tags for versioning. For nested modules in a single repository, you must prefix the tag with the module's path relative to the repository root.

### Example: Versioning `transx`

If you want to release version `v0.1.0` of the `transx` package, you should create a Git tag named:
`transx/v0.1.0`

When another project wants to use this specific version, it will look like this in their `go.mod`:

```go
require github.com/cloud-barista/cm-beetle/transx v0.1.0
```

### Note on Root Tags

Tags that are just `vX.Y.Z` (without a prefix) refer to the **root module** of the repository (`github.com/cloud-barista/cm-beetle`). These tags do **not** automatically apply to the nested modules. Each nested module should have its own set of prefixed tags for independent versioning.

## 3. How to Create Tags

You can create these prefixed tags using the Git CLI. This is the recommended method for sub-modules to avoid cluttering the repository's main Release history.

### Using Git CLI (Recommended)

```bash
# 1. Create an annotated tag with a simple release note
# -a: create an annotated tag, -m: include a message (release note)
git tag -a transx/v0.1.0 -m "Release notes: brief summary of changes"

# 2. Push the tag to your remote (origin or upstream)
git push origin transx/v0.1.0
git push upstream transx/v0.1.0
```

### Important: Git Tag vs. GitHub Release

- **Git Tag (CLI)**: This is a **technical requirement** for Go modules. Using **annotated tags** (`-a`) allows you to include a simple message (release note) within the tag itself, which can be viewed via `git show <tagname>`.
- **GitHub Release (UI)**: This is a **human-facing announcement**. It includes changelogs and assets.
  - **Recommendation**: For internal sub-modules like `transx`, it is often better to use **Option B (Git CLI)** only. This avoids "noise" in the repository's main Release history.
  - **Main Beetle releases** should continue using the GitHub UI for visibility.

## 3. Local Development (Replace Directive)

If you are developing another project locally and want to use the local version of these packages without pushing tags, use the `replace` directive in your project's `go.mod`:

```go
module my-other-project

go 1.25.0

require github.com/cloud-barista/cm-beetle/transx v0.0.0-unused

replace github.com/cloud-barista/cm-beetle/transx => ../cm-beetle/transx
```

## 4. Summary Table

| Module         | Module Path                                     | Tag Format Example  |
| :------------- | :---------------------------------------------- | :------------------ |
| **Root**       | `github.com/cloud-barista/cm-beetle`            | `v0.5.0`            |
| **Analyzer**   | `github.com/cloud-barista/cm-beetle/analyzer`   | `analyzer/v0.1.0`   |
| **Transx**     | `github.com/cloud-barista/cm-beetle/transx`     | `transx/v0.1.0`     |
| **Deepdiffgo** | `github.com/cloud-barista/cm-beetle/deepdiffgo` | `deepdiffgo/v0.1.0` |

## 5. Real-World Reference Examples

Many large-scale Go projects use this same "multi-module monorepo" pattern. You can refer to these projects' release/tagging history:

- **[OpenTelemetry Go](https://github.com/open-telemetry/opentelemetry-go/tags)**: Manages numerous modules like `api`, `sdk`, and `exporters`. Tags look like `api/v1.33.0`, `sdk/v1.33.0`, etc.
- **[Google Cloud Go SDK](https://github.com/googleapis/google-cloud-go/tags)**: Each service (e.g., `storage`, `pubsub`, `firestore`) is a separate module. Tags follow the pattern `storage/v1.43.0`, `pubsub/v1.36.2`, etc.
- **[Go Tools (gopls)](https://github.com/golang/tools/tags)**: The `gopls` language server is a module within the `tools` repo. Tags are `gopls/v0.17.1`, etc.
- **[AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2/tags)**: Manages hundreds of service-specific modules. Tags are `service/s3/v1.66.0`, `service/dynamodb/v1.59.0`, etc.

## 6. Best Practices for Maintainers

### 6.1. Reducing Confusion

To ensure clarity for contributors and users:

- **Comprehensive Documentation**: Keep this guide updated and link to it from the main `README.md`.
- **Consistent Naming**: Use the module's subdirectory name as the tag prefix (e.g., `transx/v1.0.0` for `./transx`).
- **Semantic Versioning**: Follow SemVer strictly for each module independently.

### 6.2. Documenting in Release Notes

In the main Beetle release notes (under **"Related components"**), explicitly list the versions of internal modules that were tested and included:

- **transx**: `v0.1.0` ([Link to tag](https://github.com/cloud-barista/cm-beetle/tags/transx/v0.1.0))
- **analyzer**: `v0.1.5` ([Link to tag](https://github.com/cloud-barista/cm-beetle/tags/analyzer/v0.1.5))

### 6.3. Operational Efficiency (Automation)

Managing multiple tags manually can be tedious. You can automate this process:

- **Batch Tagging Script**: Use a simple shell script to tag all related modules at once if they are usually released together.
- **GitHub Actions**: Use actions like `mathieudutour/github-tag-action` or custom scripts in your CI pipeline to handle tagging based on file changes in specific directories.

```bash
# Example batch tagging script
VERSION="v0.5.0"
TAGS=("transx/$VERSION" "analyzer/$VERSION" "deepdiffgo/$VERSION" "$VERSION")
for TAG in "${TAGS[@]}"; do
  git tag "$TAG"
  git push origin "$TAG"
done
```
