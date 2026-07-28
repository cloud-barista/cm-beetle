# Target Infrastructure Validation Feature Guide

## Overview

CM-Beetle provides a validation API that checks whether a target infrastructure model (the model a user edits after receiving a recommendation) can actually be migrated — **without creating or modifying any resource**. This lets a Portal (or any API client) catch problems while the user is still editing, instead of only finding out after clicking "migrate".

- **Validates** a `RecommendedInfra` model against the same rules migration execution enforces immediately before provisioning
- **Never creates, modifies, or deletes** any CSP/Tumblebug resource — read-only checks only
- **Returns every problem found in one call** (not just the first one), so a UI can highlight all invalid fields at once
- **Mirrors the intended migration mode** via the same `useExisting` flag the migration API takes (see [API Reference](#api-reference) below), since what counts as "valid" differs by mode

> [!NOTE]
> This feature is in **Preview**. It currently covers computing infrastructure (VNet, SSH key, security groups, VM spec/image, NodeGroups). Other target resource types (e.g. managed NLB, Object Storage, databases) are expected to be added under the same `[Validation]` API group later.

> [!IMPORTANT]
> Because Tumblebug/CSP state can change after this call returns, a `valid: true` result is a best-effort snapshot, not a guarantee. The migration API re-runs this exact same validation immediately before provisioning, so a request that was valid a minute ago can still fail migration if something changed in between (e.g. someone else created a same-named resource).

---

## API Reference

**Endpoint:** `POST /validation/ns/{nsId}/infra`

| Parameter     | Type    | Location | Required | Description                                                                                                                                                                      |
| ------------- | ------- | -------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `nsId`        | string  | path     | Yes      | Namespace ID (e.g. `mig01`)                                                                                                                                                      |
| `useExisting` | boolean | query    | No       | Validate as if reusing existing resources instead of creating new ones (default: `true`). Should match the value intended for the actual `POST /migration/ns/{nsId}/infra` call. |
| request body  | JSON    | body     | Yes      | The target `RecommendedInfra` model to validate — same shape as the migration API's request body                                                                                 |

**Response — always `200 OK`:**

```json
{
  "success": true,
  "data": {
    "valid": false,
    "issues": [
      {
        "code": "CONNECTION_MISMATCH",
        "severity": "error",
        "path": "targetVNet (id: vnet-01)",
        "message": "vNet 'vnet-01' already exists in namespace 'mig01', but under connection 'aws-ap-northeast-2', not the requested 'gcp-asia-northeast3' - reusing it would place resources in a different CSP/region"
      }
    ]
  }
}
```

`400 Bad Request` is reserved for a malformed request itself (bad JSON, missing `nsId`) — a **failed validation is not a request error**, it's a normal, successfully-answered result. Always check the `valid` field, not the HTTP status code.

---

## What Gets Validated

| #   | Category                                       | `useExisting` mode | Issue code                    | Severity | Triggered when                                                                                                                                                                                                          |
| --- | ---------------------------------------------- | ------------------ | ----------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 1   | Name format & referential integrity            | both               | `REFERENTIAL_INTEGRITY`       | error    | A name isn't 3–63 alphanumeric/hyphen characters, or an internal reference doesn't resolve within the submitted model (e.g. a NodeGroup's `securityGroupIds` entry has no matching `name` in `targetSecurityGroupList`) |
| 2   | Infra (MCI) name required                      | both               | `REQUIRED_FIELD_MISSING`      | error    | `targetInfra.name` is empty                                                                                                                                                                                             |
| 3   | Required fields for creating fresh resources   | `false`            | `REQUIRED_FIELD_MISSING`      | error    | `targetVNet.name`, `targetSshKey.name`, or a security group's `name` is empty                                                                                                                                           |
| 4   | Required fields for reusing existing resources | `true`             | `REQUIRED_FIELD_MISSING`      | error    | A NodeGroup's `vNetId`, `sshKeyId`, or `securityGroupIds` is empty                                                                                                                                                      |
| 5   | Resource name collision (fresh creation)       | `false`            | `RESOURCE_ALREADY_EXISTS`     | error    | The VNet/SSH key/security group to be created already exists in the namespace                                                                                                                                           |
| 6   | CSP/region connection match (reuse)            | `true`             | `CONNECTION_MISMATCH`         | error    | The VNet/SSH key/security group exists, but under a different Tumblebug connection (CSP+region+credential) than the NodeGroup requests                                                                                  |
| 7   | Fallback creation data present (reuse)         | `true`             | `RESOURCE_NOT_AVAILABLE`      | error    | The resource doesn't exist yet, and the accompanying `Target*` data isn't enough to create it instead (e.g. `targetVNet.cidrBlock` is empty)                                                                            |
| 8   | Spec/image required fields                     | both               | `REQUIRED_FIELD_MISSING`      | error    | A NodeGroup's `specId`, `imageId`, or `connectionName` is empty                                                                                                                                                         |
| 9   | Connection name format                         | both               | `INVALID_CONNECTION_NAME`     | error    | A NodeGroup's `connectionName` isn't in `csp-region` format (e.g. `aws-ap-northeast-2`)                                                                                                                                 |
| 10  | Spec/image lookup                              | both               | `SPEC_OR_IMAGE_LOOKUP_FAILED` | error    | The `specId` or `imageId` doesn't resolve to a real Tumblebug spec/image                                                                                                                                                |
| 11  | Spec/image compatibility                       | both               | `SPEC_IMAGE_INCOMPATIBLE`     | error    | The resolved spec and image are incompatible for that CSP (e.g. an `x86_64` spec paired with an `arm64` image)                                                                                                          |
| 12  | Infra (MCI) name collision                     | both               | `RESOURCE_ALREADY_EXISTS`     | error    | `targetInfra.name` already exists in the namespace, regardless of mode                                                                                                                                                  |

All issues currently returned are `severity: "error"` — `warning` is a reserved severity for future checks that shouldn't block migration by themselves (see [`pkg/core/validation/model.go`](../../pkg/core/validation/model.go)).

### Why checks 3–7 depend on `useExisting`

Migration behaves differently depending on `useExisting`, so validation mirrors that instead of applying one fixed rule:

- **`false` (create fresh):** CM-Beetle always creates a brand-new VNet/SSH key/security groups. If something with that name already exists, creating it would collide — so existence itself is the error (#5).
- **`true` (reuse existing):** CM-Beetle reuses a resource by ID if one is found, otherwise falls back to creating it from the accompanying `Target*` data. A resource that doesn't exist yet is only an error when that fallback data is also missing (#7).

  Finding a resource by ID isn't automatically safe to reuse, though. Tumblebug only guarantees the ID is unique — not that the resource sits in the CSP/region you're actually targeting. Reusing a same-named resource from the wrong region would silently attach new VMs to the wrong network, so its connection must match what the NodeGroup asks for (#6).

### Path format

`path` points at the offending field so a UI can highlight it directly. It's either a JSON-pointer-style field path (`targetVNet.name`, `targetInfra.nodeGroups[0].specId`) or, for resource-level checks that aren't tied to one field, a short `field (id: value)` description (`targetVNet (id: vnet-01)`). Treat it as a hint for display, not a strict machine-parseable schema — it's a plain string, not indexed by consumers today.

---

## Example: valid target model

```bash
curl -X POST "http://localhost:8056/beetle/validation/ns/mig01/infra?useExisting=false" \
  -H "Content-Type: application/json" \
  -d @target-infra.json
```

```json
{
  "success": true,
  "data": {
    "valid": true,
    "issues": []
  }
}
```

## Example: multiple problems in one response

```json
{
  "success": true,
  "data": {
    "valid": false,
    "issues": [
      {
        "code": "RESOURCE_ALREADY_EXISTS",
        "severity": "error",
        "path": "targetVNet.name",
        "message": "vNet 'vnet-01' already exists in namespace 'mig01'"
      },
      {
        "code": "SPEC_IMAGE_INCOMPATIBLE",
        "severity": "error",
        "path": "targetInfra.nodeGroups[0]",
        "message": "VM spec 'aws+t3.medium' and image 'aws+ubuntu-22.04-arm64' are incompatible for CSP 'aws' in nodegroup 'ng1'"
      }
    ]
  }
}
```

Both issues are returned together in one call — the Portal doesn't need to fix one field, call the API again, then find the next problem.

---

## References

### Source code

- API handler: [pkg/api/rest/controller/validation.go](../../pkg/api/rest/controller/validation.go)
- Validation logic: [pkg/core/validation/](../../pkg/core/validation/) ([package README](../../pkg/core/validation/README.md))
- Migration execution (runs the same checks internally before provisioning): [pkg/core/migration/infra.go](../../pkg/core/migration/infra.go)

### Related documentation

- [API Guide: Align Names](../api-guide-align-names.md) — naming/referential-integrity utility APIs this feature reuses
