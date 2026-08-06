---
mode: agent
model: Claude Sonnet 4.5
description: "Analyze and guide CM-Beetle release staging tasks for a target version"
---

# Release Staging — CM-Beetle

Perform a comprehensive staging analysis and execution for a CM-Beetle release.

## Target Version

${input:version:Target release version (e.g., v0.5.1, v0.6.0)}

## Process Overview

1. **Pre-flight analysis**: Detect current state vs. target version across all key files
2. **Report discrepancies**: Surface items that need to be updated or decided
3. **Execute approved changes**: Apply version string updates to deployment manifests and docs
4. **Verify**: Confirm build and swagger are consistent

---

## Step 1: Pre-flight Analysis

Collect the current state of the following and compare against the target version `${input:version}`.

### 1.1 Go Dependency Version (go.mod)

Read `go.mod` and extract:

- `github.com/cloud-barista/cb-tumblebug` version
- `github.com/cloud-barista/cm-beetle/imdl` version
- `github.com/cloud-barista/cm-beetle/transx` version

Cross-check these against the service image versions in `deployments/docker-compose/docker-compose.yaml`:

- `cb-tumblebug` image tag
- `cb-spider` image tag
- etcd image tag

**Flag** any mismatch between Go dependency version and deployed service version.

### 1.2 CM-Beetle Image Version Strings

Search for `cm-beetle:` image references and report current values:

- `deployments/docker-compose/docker-compose.yaml`
- `deployments/kubernetes/base/cm-beetle/deployment.yaml`
- Any overlays under `deployments/kubernetes/overlays/`

### 1.3 Docker Compose Production Settings

Check `deployments/docker-compose/docker-compose.yaml` for development settings that should be commented out:

- `build:` section under `cm-beetle` service (should be commented out for releases)
- `BEETLE_LOGLEVEL=debug` (should be commented out; default `info` is used in production)

### 1.4 Swagger API Version

Read `cmd/cm-beetle/main.go` and report the current `@version` annotation value.
Note: `latest` is acceptable and does not need updating.

### 1.5 Git Status Summary

Run `git status --short` and categorize:

- Modified tracked files (test results, docs, source)
- Untracked files (classify: release artifact vs. internal/AI files)

Report the count of modified test result files separately (they are expected to be staged as-is).

---

## Step 2: Report — Items Requiring Action

Present a concise table of findings:

| Priority    | File | Current Value | Required Value | Action                  |
| ----------- | ---- | ------------- | -------------- | ----------------------- |
| 🔴 Critical | ...  | ...           | ...            | Must fix before release |
| 🟡 Doc      | ...  | ...           | ...            | Version string update   |
| 🟢 OK       | ...  | ...           | —              | No action needed        |

Ask the user to confirm which items to proceed with before making any changes.

---

## Step 3: Execute Approved Changes

For each item confirmed by the user, apply the following:

### Version String Updates (Deployment Manifests)

Update `cm-beetle` image tags to `${input:version}` (strip the `v` prefix for Docker tags, e.g., `v0.5.1` → `0.5.1`):

- `deployments/docker-compose/docker-compose.yaml`
- `deployments/kubernetes/base/cm-beetle/deployment.yaml`

### Docker Compose Production Settings

Ensure `deployments/docker-compose/docker-compose.yaml` is configured for production:

1. **Comment out the build section** under `cm-beetle` service:
   ```yaml
   # build:
   #   context: ${PROJECT_ROOT}
   #   dockerfile: Dockerfile
   ```

2. **Comment out debug log level**:
   ```yaml
   # - BEETLE_LOGLEVEL=debug
   ```
   (This allows the default `info` level to be used)

### Go Dependency Upgrade (if go.mod version is behind docker-compose)

If `cb-tumblebug` in `go.mod` is behind the docker-compose image version, instruct the user to run:

```bash
go get github.com/cloud-barista/cb-tumblebug@<target-version>
go mod tidy
make build
```

Do **not** run this automatically — inform the user and wait for confirmation.

---

## Step 4: Post-Change Verification

After all changes are applied, verify:

1. Read back each modified file and confirm the new version string is correct.
2. Remind the user to run `make build` if Go dependencies were changed.
3. Remind the user to run `make swag` if any API handler annotations were modified.

---

## Step 5: Staging Commit Guidance

Summarize the files ready to be staged and suggest a commit message following the project's conventional commit format:

```
release: staging <version> with [key highlights]
```

Example body bullets:

- Upgrade CB-Tumblebug to vX.Y.Z
- Upgrade CB-Spider to vX.Y.Z
- Update deployment manifests to vX.Y.Z
- Refresh test results for N CSPs

Remind the user of the full release execution sequence:

```
1. PR → upstream/main merge
2. git tag -a <version> upstream/main -m "Release <version>"
3. git push upstream <version>
4. GitHub Release Notes 작성 (Breaking Changes 포함)
```

---

## Constraints

- **Do not** edit `api/swagger.yaml`, `api/swagger.json`, or `api/docs.go` directly — these are generated by `make swag`.
- **Do not** modify test result files (`testresult/`) — treat them as correct and stage as-is.
- **Do not** add or modify `CHANGELOG.md` — GitHub Release Notes replace it.
- **Do not** include internal/AI files (`.agents/`, `.github/agents/`, `plan*.md`, `plan.md`, `.github/prompts/*.prompt.md`) in the staging commit.
- All code comments and messages must be in English.
