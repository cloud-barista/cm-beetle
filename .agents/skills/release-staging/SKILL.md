---
name: release-staging
description: "Analyze and guide CM-Beetle release staging tasks for a target version (submodules, go.mod, deployments, UI, Swagger, and examples)."
---

# Release Staging - CM-Beetle

Perform a comprehensive staging analysis and execution for a CM-Beetle release (e.g., `v0.6.4`).

## Process Overview

1. **Pre-flight analysis**: Detect current state vs. target version across submodules, dependencies, deployment manifests, UI, Swagger, and example docs.
2. **Report discrepancies**: Surface items that need to be updated or decided.
3. **Execute approved changes**: Tag unreleased submodules if needed, update `go.mod`, and apply version/production updates to manifests, UI, and docs.
4. **Verify**: Confirm Go build (with and without `go.work`), Swagger sync, and UI lint/build pass cleanly.

---

## Step 1: Pre-flight Analysis

Collect the current state of the following and compare against the target version `<version>` (e.g., `v0.6.4`, and stripped Docker/NPM version `0.6.4`).

### 1.1 Independent Submodules (`imdl`, `transx`) & Go Dependency Version (`go.mod`)

1. **Check Submodule Commits Since Latest Tags**:
   // turbo
   Run `git tag -l "imdl/*" --sort=-v:refname | head -5` and `git tag -l "transx/*" --sort=-v:refname | head -5`.
   // turbo
   Run `git log <latest_imdl_tag>..HEAD --oneline -- imdl/` and `git log <latest_transx_tag>..HEAD --oneline -- transx/`.
   - **Flag** if `imdl/` or `transx/` has new commits on `upstream/main` that have not yet been tagged and updated in root `go.mod` (which causes `GOWORK=off go build ./cmd/cm-beetle` to fail).

2. **Read `go.mod` and Extract**:
   - `github.com/cloud-barista/cb-tumblebug` version
   - `github.com/cloud-barista/cm-beetle/imdl` version
   - `github.com/cloud-barista/cm-beetle/transx` version

3. **Cross-check Service Image Versions**:
   Check `deployments/docker-compose/docker-compose.yaml` and `deployments/docker-compose/docker-compose.ui.yaml`:
   - `cb-tumblebug` image tag
   - `cb-spider` image tag
   - `cb-tumblebug-etcd` image tag
   - `cm-honeybee` image tag
   - **Flag** any mismatch between Go dependency version and deployed service version.

### 1.2 CM-Beetle Image, UI, and Example Version Strings

Search for current version strings across:
- `deployments/docker-compose/docker-compose.yaml` (`image: cloudbaristaorg/cm-beetle:<version_no_v>`)
- `deployments/kubernetes/base/cm-beetle/deployment.yaml` (`image: cloudbaristaorg/cm-beetle:<version_no_v>`)
- Any overlays under `deployments/kubernetes/overlays/`
- `ui/package.json` and `ui/package-lock.json` (`"version": "<version_no_v>"`)
- `ui/src/components/layout/AppLayout.tsx` (`CM-Beetle <version>`)
- `examples/README.md`, `examples/data-migration/README.md`, and `examples/object-storage-migration/README.md` (Swagger URLs referencing `https://raw.githubusercontent.com/cloud-barista/cm-beetle/<version>/api/swagger.yaml`)

### 1.3 Docker Compose Production Settings

Check `deployments/docker-compose/docker-compose.yaml` for development settings that must be commented out for releases:
- `build:` section under `cm-beetle` service (must be commented out for production releases)
- `BEETLE_LOGLEVEL=debug` (must be commented out so default `info` level is used in production)

### 1.4 Swagger API Version & Spec Sync

- Read `cmd/cm-beetle/main.go` and report the current `@version` annotation value (`latest` is acceptable and does not need updating).
- Check if `ui/swagger-beetle.yaml` differs from `api/swagger.yaml` (`diff -q ui/swagger-beetle.yaml api/swagger.yaml`).

### 1.5 Git Status Summary

// turbo
Run `git status --short` and categorize:
- Modified tracked files (test results, docs, source)
- Untracked files (classify: release artifact vs. internal/AI files)
- Report the count of modified test result files separately (they are expected to be staged as-is).

---

## Step 2: Report - Items Requiring Action

Present a concise table of findings:

| Priority | File / Component | Current Value | Required Value | Action |
| :--- | :--- | :--- | :--- | :--- |
| 🔴 Critical | ... | ... | ... | Must fix before release |
| 🟡 Doc / Version | ... | ... | ... | Version string or spec sync |
| 🟢 OK | ... | ... | - | No action needed |

Ask the user to confirm which items to proceed with before making any changes.

---

## Step 3: Execute Approved Changes

For each item confirmed by the user, apply the following:

### 3.1 Independent Submodule Tagging & `go.mod` Update (if applicable)

If `imdl/` or `transx/` has unreleased commits already merged on `upstream/main`:

```bash
# 1. Fetch latest upstream
git fetch upstream

# 2. Create annotated submodule tags on upstream/main
git tag -a imdl/vX.Y.Z upstream/main -m "imdl: release vX.Y.Z"
git tag -a transx/vX.Y.Z upstream/main -m "transx: release vX.Y.Z"

# 3. Push submodule tags directly to upstream
git push upstream imdl/vX.Y.Z transx/vX.Y.Z

# 4. Update root go.mod and go.sum
go get github.com/cloud-barista/cm-beetle/imdl@vX.Y.Z github.com/cloud-barista/cm-beetle/transx@vX.Y.Z
```

### 3.2 Version String Updates (Manifests, UI, Examples)

1. **Deployment Manifests** (strip the `v` prefix for Docker tags, e.g., `v0.6.4` -> `0.6.4`):
   - `deployments/docker-compose/docker-compose.yaml`: Uncomment and set `image: cloudbaristaorg/cm-beetle:<version_no_v>`
   - `deployments/kubernetes/base/cm-beetle/deployment.yaml`: Set `image: cloudbaristaorg/cm-beetle:<version_no_v>`

2. **Beetle UX Lab (`ui/`)**:
   - `ui/package.json` and `ui/package-lock.json`: Set `"version": "<version_no_v>"`
   - `ui/src/components/layout/AppLayout.tsx`: Set header badge to `CM-Beetle <version>`
   - If `ui/swagger-beetle.yaml` is out of date, sync it with `cp api/swagger.yaml ui/swagger-beetle.yaml`.

3. **Example Documentation Swagger Links**:
   - Update `https://raw.githubusercontent.com/cloud-barista/cm-beetle/<old_version>/api/swagger.yaml` to `<version>` in:
     - `examples/README.md`
     - `examples/data-migration/README.md`
     - `examples/object-storage-migration/README.md`

### 3.3 Docker Compose Production Settings

Ensure `deployments/docker-compose/docker-compose.yaml` is configured for production:

1. **Comment out the build section** under `cm-beetle` service:
   ```yaml
   # build:
   #   context: ${PROJECT_ROOT}
   #   dockerfile: Dockerfile
   ```

2. **Ensure debug log level is commented out**:
   ```yaml
   # - BEETLE_LOGLEVEL=debug
   ```

### 3.4 CB-Tumblebug Go Dependency Upgrade (if `go.mod` is behind `docker-compose.yaml`)

If `cb-tumblebug` in `go.mod` is behind the `docker-compose.yaml` image version, confirm with the user and run:

```bash
go get github.com/cloud-barista/cb-tumblebug@<target-version>
go mod tidy
make build
```

---

## Step 4: Post-Change Verification

After all changes are applied, verify:

1. **Inspect `git diff`** to confirm all version strings and manifests are accurate.
2. **Verify Go Builds** (both workspace and standalone module resolution):
   ```bash
   go build ./cmd/cm-beetle
   GOWORK=off go build ./cmd/cm-beetle
   rm -f cm-beetle
   ```
3. **Verify UI Lint & Production Build**:
   ```bash
   npm --prefix ui run lint
   npm --prefix ui run build
   ```
4. If any API handler annotations were modified, run `make swag` and confirm `api/swagger.yaml` is up to date.

---

## Step 5: Staging Commit & Release Execution Guidance

Summarize the files ready to be staged (explicitly listing paths instead of `git add .`) and suggest a commit message following the project's conventional commit format:

```text
release: staging <version> with deployment updates

- Upgrade imdl to vX.Y.Z and transx to vX.Y.Z
- Set CM-Beetle image to <version_no_v> across Compose and Kubernetes
- Comment out local build section in Compose for production
- Update Beetle UX Lab UI version, layout badge, and Swagger spec
- Update Swagger spec URLs to <version> across examples
```

Remind the user of the full release execution sequence:

```text
1. Push branch to origin fork -> Open PR to upstream/main -> /approve and wait for auto-merge
2. git fetch upstream && git tag -a <version> upstream/main -m "Release <version>"
3. git push upstream <version>
4. Write GitHub Release Notes (including Breaking Changes)
```

---

## Constraints

- **Do not** edit `api/swagger.yaml`, `api/swagger.json`, or `api/docs.go` directly (these are generated by `make swag`).
- **Do not** modify test result files (`testresult/`); treat them as correct and stage as-is if modified.
- **Do not** add or modify `CHANGELOG.md` (GitHub Release Notes replace it).
- **Do not** include internal/AI files (`.agents/`, `.github/agents/`, `docs/plan/*.md`, `plan*.md`, `.github/prompts/*.prompt.md`, scratch files) in the staging commit.
- All code comments and commit messages must be in English.
