# pkg/core/validation

Checks whether a target migration model (`cloudmodel.RecommendedInfra`) is
internally consistent and can be provisioned, without creating or modifying
any CSP/Tumblebug resource. Backs both the standalone `POST
/validation/ns/{nsId}/infra` API and the pre-flight gate that
`pkg/core/migration` runs immediately before provisioning, so the two never
drift apart.

## Entry point

```go
result := validation.ValidateTargetInfra(nsId, targetInfraModel, useExisting)
// result.Valid  bool
// result.Issues []ValidationIssue{Code, Severity, Path, Message}
err := result.Err() // non-nil summary error, or nil if Valid
```

## What it checks

- **Naming & referential integrity** (`common.ValidateComposedNames`): name
  format, and internal references resolve within the submitted model.
- **Required fields**, which differ by `useExisting`: node group
  `vNetId`/`sshKeyId`/`securityGroupIds` (true) vs. `targetVNet.name` /
  `targetSshKey.name` / security group `name` (false).
- **Resource name collision / availability** against Tumblebug: must NOT
  already exist (`useExisting=false`), or must exist under the _same
  CSP/region connection_ the node group requests — not just under the same ID
  (`useExisting=true`, `CONNECTION_MISMATCH` otherwise).
- **VM spec/image compatibility** per node group (Tumblebug spec/image
  lookup + `pkg/compat`).
- **Infra (MCI) name collision**: `targetInfra.name` must not already exist.

## Running the tests

```bash
go test ./pkg/core/validation/... -v -cover
```

Plain `go test ./pkg/core/validation/...` (no `-v`) prints a single `ok ...`
line — the quietest form to paste into a PR or chat.

No live Tumblebug, Spider, or CSP credentials are needed. `infra_test.go`'s
`TestMain` starts an `httptest.Server` that serves canned JSON for the
handful of Tumblebug endpoints this package reads (VNet/SshKey/SecurityGroup/
Infra/Spec/Image), then points `tbclient` at it via `tbclient.Init(...)`.
Any request path not in the `fixtures` map 404s automatically, which is
exactly the "resource does not exist" case most scenarios need — so only the
"found" fixtures require an explicit map entry.

`tbclient.NewSession()` uses a single process-global client guarded by
`sync.Once`, so `tbclient.Init` can only meaningfully run once per test
binary. All scenarios therefore share one fake server; `baseTarget()`
builds a fresh, fully-valid model per test, and small `targetOption` helpers
(`withVNetName`, `withSshKeyName`, `withSecurityGroupName`) repoint it at a
specific fixture (or a nonexistent name) while keeping referential integrity
intact.

### Adding a new scenario

1. If it needs a resource to already exist, add a fixture to the `fixtures`
   map in `infra_test.go` (path → status/body). Not-found cases need nothing.
2. Add a case to the `tests` table in `TestValidateTargetInfra`: `useExisting`,
   the `targetOption`s to reach that state, and the expected `wantValid`/
   `wantCode`.
3. Give fixture names that read as intent (`vnet-mismatch`,
   `vnet-does-not-exist`) — the fixtures map doubles as documentation of what
   each stands in for.
