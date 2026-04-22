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

### Repository Setup

This project follows a three-tier Git workflow:

- **upstream**: The official repository (`github.com/cloud-barista/cm-beetle`)
- **origin**: Your personal fork of the upstream repository
- **local**: Your local clone of the origin fork

**Normal code change workflow:**

```
local commit → git push origin <branch> → PR to upstream/main → merge
```

**Tagging workflow is different.** Tags are not pushed via PR — they are applied to a specific commit and pushed **directly to upstream**. The correct sequence is:

1. Ensure the relevant changes are already merged into `upstream/main` via PR.
2. Fetch the latest state of `upstream/main` to your local repository.
3. Create an annotated tag locally, pointing to `upstream/main`.
4. Push the tag directly to `upstream` (not `origin`).

> **Note:** Pushing a tag to `origin` is not required. The Go module proxy resolves versions from the upstream (official) repository's tags. If you do not have push access to `upstream`, ask a maintainer with push access to create the tag.

## 4. Quick Guide: Two Independent Workflows

> [!IMPORTANT]
>
> 1. Release and tag `transx` first.
> 2. Then update Beetle `go.mod` to that tag.

Why: `go.mod` can use `github.com/cloud-barista/cm-beetle/transx v0.1.0` only after the tag `transx/v0.1.0` exists on `upstream`.

Dependency rule: Workflow B starts only after Workflow A tag is available on `upstream`.

Development note: Feature development can continue between Workflow A and Workflow B. After `transx/v0.1.0` is tagged, apply that version in Beetle `go.mod` when you start Workflow B.

Typical timeline:

1. Workflow A complete (`transx` tag published).
2. Ongoing Beetle feature development and staging work.
3. Workflow B starts (update `go.mod` to tagged `transx` version), then PR/merge/release.

### Workflow A: Release and Tag transx

Goal: create `transx/v0.1.0` on `upstream`.

```bash
# 1) Develop transx on a branch
git checkout upstream/main -b feat-transx-v0.1.0
# edit transx/
git add transx/
git commit -m "feat(transx): ..."
git push origin feat-transx-v0.1.0

# 2) Open PR: origin/feat-transx-v0.1.0 -> upstream/main
#    Wait until merged.

# 3) Tag the merge result on upstream/main
git fetch upstream
git log upstream/main --oneline -5
git tag -a transx/v0.1.0 upstream/main -m "transx: release v0.1.0"

# Optional (safer if upstream/main has moved):
# git tag -a transx/v0.1.0 <merge_commit_sha> -m "transx: release v0.1.0"

# 4) Push tag to upstream
git push upstream transx/v0.1.0

# 5) Verify
git show transx/v0.1.0
```

Note: tag `upstream/main` (merge result), not your old branch commit hash.
If another PR is merged before you tag, use the exact merge commit SHA instead of `upstream/main`.

### Workflow B: Update Beetle to the Tagged Version

Goal: upgrade Beetle dependency to `transx v0.1.0`, then release Beetle.

```bash
# 1) Start a new branch
git fetch upstream
git checkout upstream/main -b feat-beetle-use-transx-v0.1.0

# 2) Update dependency
go get github.com/cloud-barista/cm-beetle/transx@v0.1.0
go mod tidy

# 3) Continue Beetle development and staging work
make build
make swag
go test ./...

# 4) Commit and push
# Option A: keep dependency bump and feature work in separate commits (recommended)
#   git add go.mod go.sum
#   git commit -m "chore(deps): bump transx to v0.1.0"
#   git add <feature-files>
#   git commit -m "feat(...): ..."
#
# Option B: single commit for small changes
git add .
git commit -m "chore(deps): bump transx to v0.1.0"
git push origin feat-beetle-use-transx-v0.1.0

# 5) Open PR -> upstream/main and merge
# 6) Create GitHub Release for Beetle (vX.Y.Z)
```

### One-Line Checklist

- Workflow A complete: `transx/v0.1.0` exists on `upstream`.
- Workflow B complete: Beetle PR merged with updated `go.mod`.
- Final: Beetle GitHub Release created.

### First-Time Maintainer Checklist

- Confirm PR is merged to `upstream/main` before tagging.
- Tag `upstream` merge commit, not local feature branch commit.
- Push tag to `upstream`, not only `origin`.
- Confirm tag with `git show`.
- Start Beetle version bump only after tag is visible on `upstream`.

## 5. Local Development (Replace Directive)

If you are developing another project locally and want to use the local version of these packages without pushing tags, use the `replace` directive in your project's `go.mod`:

```go
module my-other-project

go 1.25.0

require github.com/cloud-barista/cm-beetle/transx v0.0.0-unused

replace github.com/cloud-barista/cm-beetle/transx => ../cm-beetle/transx
```

## 6. Summary Table

| Module         | Module Path                                     | Tag Format Example  |
| :------------- | :---------------------------------------------- | :------------------ |
| **Root**       | `github.com/cloud-barista/cm-beetle`            | `v0.5.0`            |
| **Analyzer**   | `github.com/cloud-barista/cm-beetle/analyzer`   | `analyzer/v0.1.0`   |
| **Transx**     | `github.com/cloud-barista/cm-beetle/transx`     | `transx/v0.1.0`     |
| **Deepdiffgo** | `github.com/cloud-barista/cm-beetle/deepdiffgo` | `deepdiffgo/v0.1.0` |

## 7. Real-World Reference Examples

Many large-scale Go projects use this same "multi-module monorepo" pattern. You can refer to these projects' release/tagging history:

- **[OpenTelemetry Go](https://github.com/open-telemetry/opentelemetry-go/tags)**: Manages numerous modules like `api`, `sdk`, and `exporters`. Tags look like `api/v1.33.0`, `sdk/v1.33.0`, etc.
- **[Google Cloud Go SDK](https://github.com/googleapis/google-cloud-go/tags)**: Each service (e.g., `storage`, `pubsub`, `firestore`) is a separate module. Tags follow the pattern `storage/v1.43.0`, `pubsub/v1.36.2`, etc.
- **[Go Tools (gopls)](https://github.com/golang/tools/tags)**: The `gopls` language server is a module within the `tools` repo. Tags are `gopls/v0.17.1`, etc.
- **[AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2/tags)**: Manages hundreds of service-specific modules. Tags are `service/s3/v1.66.0`, `service/dynamodb/v1.59.0`, etc.

## 8. Best Practices for Maintainers

### 8.1. Reducing Confusion

To ensure clarity for contributors and users:

- **Comprehensive Documentation**: Keep this guide updated and link to it from the main `README.md`.
- **Consistent Naming**: Use the module's subdirectory name as the tag prefix (e.g., `transx/v1.0.0` for `./transx`).
- **Semantic Versioning**: Follow SemVer strictly for each module independently.

### 8.2. Documenting in Release Notes

In the main Beetle release notes (under **"Related components"**), explicitly list the versions of internal modules that were tested and included:

- **transx**: `v0.1.0` ([Link to tag](https://github.com/cloud-barista/cm-beetle/tags/transx/v0.1.0))
- **analyzer**: `v0.1.5` ([Link to tag](https://github.com/cloud-barista/cm-beetle/tags/analyzer/v0.1.5))

### 8.3. Operational Efficiency (Automation)

Managing multiple tags manually can be tedious. You can automate this process:

- **Batch Tagging Script**: Use a simple shell script to tag all related modules at once if they are usually released together.
- **GitHub Actions**: Use actions like `mathieudutour/github-tag-action` or custom scripts in your CI pipeline to handle tagging based on file changes in specific directories.

```bash
# Example batch tagging script
# Run this after all changes are merged into upstream/main
git fetch upstream

VERSION="v0.5.0"
TAGS=("transx/$VERSION" "analyzer/$VERSION" "deepdiffgo/$VERSION" "$VERSION")
for TAG in "${TAGS[@]}"; do
  git tag -a "$TAG" upstream/main -m "Release $TAG"
  git push upstream "$TAG"
done
```
