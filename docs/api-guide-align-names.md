# API Guide: Align Names

_Subtitle: Propagate Name Changes to Dependent Resources_

This guide covers the **Align Names** API (`POST /naming/alignment`), which propagates a resource name change
to all dependent child resources within a `RecommendedVmInfra` model.

## Overview

When a parent resource (e.g., VNet) is renamed, child resources that reference it (e.g., Security Groups, SubGroups)
must also be updated to maintain internal consistency. The Align Names API automates this propagation.

CM-Beetle uses a **Late Binding** strategy: the recommendation model holds base names (e.g., `vnet-01`), and a `nameSeed` prefix is applied at migration time via the `?nameSeed=` query parameter.

> For the full NameSeed concept, rules, and workflow, see [NameSeed: Late Binding for Resource Names](feature-guide/naming-nameseed.md).

> [!NOTE]
> **One resource at a time.** Each rename specifies exactly one `oldName` → `newName`.
> Bulk alignment is not supported: when multiple resources of the same type exist (e.g., two subnets),
> each must be renamed individually to avoid ambiguity in identifying the target and its dependents.

---

## 1. Resource Name Alignment API

If you manually rename a resource in the model, use this API to propagate the change to all dependent child resources.

**Endpoint**: `POST /beetle/naming/alignment`

| Query Parameter | Type     | Description                                                |
| --------------- | -------- | ---------------------------------------------------------- |
| `resourceType`  | `string` | One of: `vNet`, `subnet`, `securityGroup`, `sshKey`, `mci` |
| `oldName`       | `string` | Current name of the resource (before change)               |
| `newName`       | `string` | New name to assign                                         |

**Body**: `RecommendedVmInfra` JSON

### Propagation Rules

| If you rename... | CM-Beetle automatically updates...         |
| ---------------- | ------------------------------------------ |
| `vNet`           | `SecurityGroup.vNetId`, `SubGroup.vNetId`  |
| `subnet`         | `SubGroup.subnetId`                        |
| `sshKey`         | `SubGroup.sshKeyId`                        |
| `securityGroup`  | Entries in `SubGroup.securityGroupIds`     |
| `mci`            | The MCI name itself (no child propagation) |

---

## 3. Validation API

Verify the internal consistency of a model without making changes.

**Endpoint**: `POST /beetle/naming/validation`  
**Body**: `RecommendedVmInfra` JSON  
**Function**: Checks that all IDs (`vNetId`, `subnetId`, etc.) reference existing resources within the same model.

---

## 4. Preview API

Dry-run: apply a NameSeed to every resource name in the model and inspect the result before migration.

**Endpoint**: `POST /beetle/naming/preview`

| Query Parameter | Type     | Required | Description                                                    |
| --------------- | -------- | -------- | -------------------------------------------------------------- |
| `nameSeed`      | `string` | No       | Prefix to apply to all resource names (e.g., `blue`)          |

**Body**: `RecommendedVmInfra` JSON  
**Function**: Returns the model with `{nameSeed}-` prepended to every resource name. No resources are created.

> Example: `nameSeed=blue` + base name `vnet-01` → `blue-vnet-01`.

---

## 5. Example Usage

> [!TIP]
> The credentials below (`default:default`) are **development-only** defaults.  
> Replace them with your configured credentials in other environments.

### Sample Model

<details>
<summary>sample-infra.json (click to expand)</summary>

```json
{
  "targetVNet": {
    "name": "vnet-01",
    "subnetInfoList": [{ "name": "subnet-01" }]
  },
  "targetSecurityGroupList": [{ "name": "sg-01", "vNetId": "vnet-01" }],
  "targetSshKey": { "name": "sshkey-01" },
  "targetVmInfra": {
    "name": "mci-01",
    "subGroups": [
      {
        "name": "subgroup-01",
        "vNetId": "vnet-01",
        "subnetId": "subnet-01",
        "sshKeyId": "sshkey-01",
        "securityGroupIds": ["sg-01"]
      }
    ]
  }
}
```

</details>

### Example 1: Rename VNet

> Set the `BODY` variable once and reuse it across all examples below.

```bash
BODY='{
  "targetVNet": {"name": "vnet-01", "subnetInfoList": [{"name": "subnet-01"}]},
  "targetSecurityGroupList": [{"name": "sg-01", "vNetId": "vnet-01"}],
  "targetSshKey": {"name": "sshkey-01"},
  "targetVmInfra": {
    "name": "mci-01",
    "subGroups": [{
      "name": "subgroup-01",
      "vNetId": "vnet-01", "subnetId": "subnet-01",
      "sshKeyId": "sshkey-01", "securityGroupIds": ["sg-01"]
    }]
  }
}'

curl -u default:default -s -X POST \
  "http://localhost:8056/beetle/naming/alignment?resourceType=vNet&oldName=vnet-01&newName=my-vnet-01" \
  -H "Content-Type: application/json" \
  -d "$BODY"
```

**Verified response (key fields):**

```json
{
  "targetVNet": { "name": "my-vnet-01" },
  "targetSecurityGroupList": [{ "vNetId": "my-vnet-01" }],
  "targetVmInfra": { "subGroups": [{ "vNetId": "my-vnet-01" }] }
}
```

### Example 2: Rename Subnet

```bash
curl -u default:default -s -X POST \
  "http://localhost:8056/beetle/naming/alignment?resourceType=subnet&oldName=subnet-01&newName=my-subnet-01" \
  -H "Content-Type: application/json" \
  -d "$BODY"
```

**Verified response (key fields):**

```json
{
  "targetVNet": { "subnetInfoList": [{ "name": "my-subnet-01" }] },
  "targetVmInfra": { "subGroups": [{ "subnetId": "my-subnet-01" }] }
}
```

### Example 3: Rename Security Group

```bash
curl -u default:default -s -X POST \
  "http://localhost:8056/beetle/naming/alignment?resourceType=securityGroup&oldName=sg-01&newName=my-sg-01" \
  -H "Content-Type: application/json" \
  -d "$BODY"
```

**Verified response (key fields):**

```json
{
  "targetSecurityGroupList": [{ "name": "my-sg-01" }],
  "targetVmInfra": { "subGroups": [{ "securityGroupIds": ["my-sg-01"] }] }
}
```

### Example 4: Preview Names with NameSeed

```bash
curl -u default:default -s -X POST \
  "http://localhost:8056/beetle/naming/preview?nameSeed=blue" \
  -H "Content-Type: application/json" \
  -d "$BODY"
```

**Verified response (key fields):**

```json
{
  "targetVNet": { "name": "blue-vnet-01" },
  "targetSecurityGroupList": [{ "name": "blue-sg-01" }],
  "targetSshKey": { "name": "blue-sshkey-01" },
  "targetVmInfra": { "name": "blue-mci-01" }
}
```

> This is a dry-run. Pass the same model and `?nameSeed=blue` to `POST /migration/.../infra?nameSeed=blue` when ready to create resources.
