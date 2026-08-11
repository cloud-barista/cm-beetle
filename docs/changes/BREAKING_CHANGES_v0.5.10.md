# Breaking Changes for v0.5.10

**Release Version:** CM-Beetle v0.5.9 → v0.5.10  
**Release Date:** 2026-08-07  
**Dependency Updates:** CB-Tumblebug v0.12.25 → v0.12.30  
**Migration Required:** Yes - Breaking changes in API request/response format

---

# Part 1: Quick Reference

> **Target Audience:** Beetle API users, `imdl/cloud-model` package consumers  
> **Goal:** Understand breaking changes and required migration actions

---

## 📋 Breaking Changes Summary (High Priority)

> **Target:** Beetle API users actively using `/infra` endpoint

| Category       | Affected API / Model                     | Impact                                  | Required Action                              |
| -------------- | ---------------------------------------- | --------------------------------------- | -------------------------------------------- |
| **API**        | POST `/beetle/migration/ns/{nsId}/infra` | 🔴 **BREAKING** - Request format change | `postCommand` → `postCommands` (array)       |
| **Logic**      | Default security group auto-generation   | 🟡 **BEHAVIOR** - Ports closed          | -                                            |
| **Deployment** | CB-Spider metadata storage               | 🟡 **DEPLOYMENT** - Database migration  | Re-register CSP credentials after deployment |
| **Model**      | `imdl/cloud-model` package               | ✅ **ADDITIVE** - New structs/fields    | Update imports, rebuild                      |

> **Note:** SG recommendation and migration logic may change to address Tumblebug's security group creation policy change (only allow 22).

---

## 📋 Low Priority Changes (`/infraWithDefaults` endpoint)

> **Target:** `/infraWithDefaults` is not actively used by upstream subsystems. These changes are for reference only.

| Category | Affected API                                         | Impact                                  | Required Action                        |
| -------- | ---------------------------------------------------- | --------------------------------------- | -------------------------------------- |
| **API**  | POST `/beetle/migration/ns/{nsId}/infraWithDefaults` | 🔴 **BREAKING** - Request format change | `postCommand` → `postCommands` (array) |
| **API**  | POST `/beetle/migration/ns/{nsId}/infraWithDefaults` | 🔴 **BREAKING** - Field removed         | Remove `nodeUserPassword` field        |

> **Note:** `/infra` endpoint retains `nodeUserPassword` field. Only `/infraWithDefaults` (dynamic provisioning) removes it.

---

## ⚠️ Breaking Change #1: `postCommand` → `postCommands` (Array)

### What Changed

**API Endpoints:**

- `POST /beetle/migration/ns/{nsId}/infra` - Request body field changed
- `POST /beetle/migration/ns/{nsId}/infraWithDefaults` - Request body field changed

**Before (v0.5.9):**

```json
{
  "name": "my-infra",
  "postCommand": {
    "userName": "cb-user",
    "command": ["apt-get update"]
  }
}
```

**After (v0.5.10):**

```json
{
  "name": "my-infra",
  "postCommands": [
    {
      "userName": "cb-user",
      "command": ["apt-get update"]
    }
  ],
  "postCommandAsync": false
}
```

### Why This Changed

- Support multi-phase bootstrap commands (sequential execution)
- Target specific node groups or labels per phase
- Control error handling per phase

### Migration Action

**For API Consumers:**

1. Change JSON field name: `postCommand` → `postCommands`
2. Wrap command in array: `"postCommands": [{...}]`

> Response bodies are unaffected — `postCommand`/`postCommandResult` in the Infra response keep their existing (singular) field names.

**For `imdl/cloud-model` Users:**

```go
// Before (v0.5.9)
import cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

infraReq := cloudmodel.InfraReq{
    Name: "my-infra",
    PostCommand: cloudmodel.InfraCmdReq{
        UserName: "cb-user",
        Command:  []string{"apt-get update"},
    },
}
```

```go
// After (v0.5.10)
import cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

infraReq := cloudmodel.InfraReq{
    Name: "my-infra",
    PostCommands: []cloudmodel.PostCommandReq{
        {
            InfraCmdReq: cloudmodel.InfraCmdReq{
                UserName: "cb-user",
                Command:  []string{"apt-get update"},
            },
        },
    },
}
```

### Search Pattern

```bash
# Find usages in your codebase
grep -r '"postCommand"' --include="*.json" --include="*.go"
```

---

## ⚠️ Breaking Change #2: `nodeUserPassword` Field Removed (Low Priority)

> **⚠️ Low Priority:** `/infraWithDefaults` is not actively used by upstream subsystems. This change is for reference only.

### What Changed

**API Endpoint:**

- `POST /beetle/migration/ns/{nsId}/infraWithDefaults` - Request body field removed

> **Important:** This only affects `/infraWithDefaults` (dynamic node groups). The `nodeUserPassword` field on the **static `/infra` endpoint's node groups is unaffected and still exists**.

**Before (v0.5.9):**

```json
{
  "nodeGroups": [
    {
      "name": "worker-group",
      "nodeUserPassword": "mypassword"
    }
  ]
}
```

**After (v0.5.10):**

```json
{
  "nodeGroups": [
    {
      "name": "worker-group"
    }
  ]
}
```

### Why This Changed

- Linux VMs use SSH keys exclusively (more secure)
- Windows VM passwords are auto-generated internally

### Migration Action

**For API Consumers:**

1. Remove `nodeUserPassword` field from request bodies
2. Linux: Access VMs via SSH keys (as before)
3. Windows: Retrieve auto-generated password from CSP console if needed

**For `imdl/cloud-model` Users:**

```go
// Before (v0.5.9)
nodeGroup := cloudmodel.CreateNodeGroupDynamicReq{
    Name:             "worker-group",
    NodeUserPassword: "mypassword",  // ❌ REMOVED
}
```

```go
// After (v0.5.10)
nodeGroup := cloudmodel.CreateNodeGroupDynamicReq{
    Name: "worker-group",
    // SSH keys only
}
```

### Search Pattern

```bash
# Find usages in your codebase
grep -r '"nodeUserPassword"' --include="*.json" --include="*.go"
```

---

## ℹ️ Logic Change: Default Security Group Now SSH-Only (Decision Needed)

Not an API or model change — no request/response field changed. But CM-Beetle's own recommendation/migration code needs a decision on how to handle it.

### What Changed

CB-Tumblebug's auto-generated security group (created when a Node group's infra request carries no explicit security group) now only opens TCP port 22 (SSH), instead of opening all TCP/UDP ports + ICMP.

**Previous Behavior (v0.5.9):**

- All TCP/UDP ports 1-65535 + ICMP were open

**New Behavior (v0.5.10):**

- Only TCP 22 (SSH) is open
- All other ports are closed by default

### Where This Applies in CM-Beetle

- **Static path** (`POST /infra`, via `recommendation.RecommendSecurityGroup` + `migration.checkAndSupportSSHAccessRule`): Unaffected. CM-Beetle always builds an explicit `FirewallRules` list before calling Tumblebug, so Tumblebug's auto-generation logic is never triggered here.
- **Dynamic path** (`POST /infraWithDefaults`, via `migration.CreateInfraWithDefaults`): Affected. CM-Beetle forwards `InfraDynamicReq` to Tumblebug as-is and does not set `SgTemplateId` or any firewall rules, so callers who omit `sgTemplateId` now silently get an SSH-only security group instead of the previous open-all default.

### Action Needed (CM-Beetle Maintainers)

Decide, in `pkg/core/migration/infra.go` (`CreateInfraWithDefaults`) and/or `pkg/core/recommendation`, whether to:

1. **Support the previous behavior** — set a default `SgTemplateId` (or equivalent firewall rules) for the dynamic path so existing callers keep the old open-access behavior, or
2. **Exclude/accept the new default** — leave the dynamic path as a pure passthrough and require callers to set `sgTemplateId` explicitly for non-SSH access (current state).

### Options for API Callers (Until the Above Is Decided)

#### Option A: Add Ports Manually After Creation

```bash
# 1. Get security group ID from infra details
curl "http://beetle:8056/beetle/migration/ns/{nsId}/infra/{infraId}?option=accessinfo"

# 2. Add required ports via Tumblebug API
curl -X POST "http://tumblebug:1323/tumblebug/ns/{nsId}/resources/securityGroup/{sgId}/rules" \
  -H 'Content-Type: application/json' \
  -d '{
    "firewallRules": [
      {"ports": "80",   "protocol": "TCP", "direction": "inbound", "cidr": "0.0.0.0/0"},
      {"ports": "443",  "protocol": "TCP", "direction": "inbound", "cidr": "0.0.0.0/0"}
    ]
  }'
```

#### Option B: Use Security Group Template

```json
{
  "name": "web-infra",
  "sgTemplateId": "sg-usecase-web",
  "nodeGroups": [...]
}
```

**Available Templates:**
| Template ID | Open Ports | Use Case |
| ------------------ | -------------------------------- | -------------------- |
| `sg-default` | TCP 22 | SSH only |
| `sg-usecase-web` | TCP 22, 80, 443, 8080, 8443 | Web servers |
| `sg-k8s` | All TCP/UDP + ICMP | Kubernetes clusters |
| `sg-openall` | All TCP/UDP + ICMP | Dev/test only |

---

## 🔧 Deployment Changes

### CB-Spider PostgreSQL METADB

**What Changed:**

- CB-Spider now uses PostgreSQL instead of SQLite for metadata storage

**Impact:**

- Existing Spider SQLite data will not be migrated
- All CSP credentials must be re-registered after deployment

**Required Actions:**

1. Update `deployments/docker-compose/.env`:

```bash
# Add Spider PostgreSQL credentials
SP_POSTGRES_USER=spider
SP_POSTGRES_PASSWORD=spider
SP_POSTGRES_DB=cb_spider
```

2. Update `docker-compose.yaml` to include `cb-spider-postgres` service (already done in v0.5.10)

3. After deployment:

```bash
docker compose down
docker compose up -d

# Re-register CSP credentials
curl -X POST "http://spider:1024/spider/credential/..." \
  -d @credential.json
```

---

## ✅ Migration Checklist

### High Priority (Required for `/infra` endpoint users)

**Code Changes:**

- [ ] Search codebase: `grep -r '"postCommand"'` (exclude `postCommands`)
- [ ] Update all `postCommand` → `postCommands` for `/infra` endpoint (wrap in array)
- [ ] Build and test: `go build ./...`

**Deployment:**

- [ ] Update `.env` with Spider PostgreSQL credentials
- [ ] Verify `docker-compose.yaml` includes `cb-spider-postgres` service
- [ ] Deploy: `docker compose down && docker compose up -d`
- [ ] Re-register all CSP credentials

**Testing:**

- [ ] Test infrastructure creation with new `postCommands` format

### Low Priority (Reference only - for `/infraWithDefaults` users)

> `/infraWithDefaults` is not actively used by upstream subsystems

**Code Changes:**

- [ ] Search codebase: `grep -r '"nodeUserPassword"'` (calls to `infraWithDefaults` only)
- [ ] Remove all `nodeUserPassword` field references (note: `/infra` still has this field)
- [ ] Update `postCommand` → `postCommands` for `/infraWithDefaults` endpoint

### CM-Beetle Codebase (Decision Needed)

- [ ] Decide whether `CreateInfraWithDefaults` (`pkg/core/migration/infra.go`) should default `SgTemplateId` to preserve the previous open-all behavior, or leave it to callers
- [ ] If supporting the previous behavior, implement the default and add a test covering the no-`sgTemplateId` case
- [ ] If excluding it, document the new SSH-only default clearly in the API reference

### Optional (New Features)

- [ ] Review multi-phase bootstrap capabilities
- [ ] Review multi-AZ subnet distribution (`distributeSubnets`)
- [ ] Review credential verification diagnostics (`verifiedMessage`)

---

## 🆕 New Capabilities (Non-Breaking)

### Multi-Phase Bootstrap Commands

Execute sequential commands with phase-based targeting:

```json
{
  "postCommands": [
    {
      "command": ["apt-get update"],
      "userName": "cb-user"
    },
    {
      "command": ["kubeadm init"],
      "userName": "cb-user",
      "nodeGroupId": "control"
    },
    {
      "command": ["kubeadm join ..."],
      "userName": "cb-user",
      "labelSelector": "role=worker",
      "continueOnError": true
    }
  ],
  "postCommandAsync": true
}
```

### Multi-AZ Distribution

Automatically distribute VMs across availability zones:

```json
{
  "nodeGroups": [
    {
      "name": "web-tier",
      "desiredNodeCount": 6,
      "distributeSubnets": true
    }
  ]
}
```

### Credential Diagnostics

Get detailed error messages for credential verification failures:

```json
{
  "verified": false,
  "verifiedMessage": "The client secret has expired. Issue a new secret in the Azure portal. (CSP error: AADSTS7000222)"
}
```

---

> **📖 For detailed technical reference and troubleshooting, see Part 2 below**

---

# Part 2: Technical Details & AI Reference

> **Target Audience:** AI assistants, infrastructure operators, troubleshooting

---

## 🔍 Affected Components Detail

### Beetle API Endpoints (Breaking - High Priority)

| Method | Endpoint                            | Breaking Change                          | Details                             |
| ------ | ----------------------------------- | ---------------------------------------- | ----------------------------------- |
| POST   | `/beetle/migration/ns/{nsId}/infra` | Request Body: postCommand → postCommands | Single field → array, see Change #1 |

### Beetle API Endpoints (Breaking - Low Priority / Reference Only)

> `/infraWithDefaults` is not actively used by upstream subsystems.

| Method | Endpoint                                        | Breaking Change                          | Details                             |
| ------ | ----------------------------------------------- | ---------------------------------------- | ----------------------------------- |
| POST   | `/beetle/migration/ns/{nsId}/infraWithDefaults` | Request Body: postCommand → postCommands | Single field → array, see Change #1 |
| POST   | `/beetle/migration/ns/{nsId}/infraWithDefaults` | Request Body: nodeUserPassword removed   | See Change #2                       |

> **Note:** `/infra` endpoint retains `nodeUserPassword` field in `CreateNodeGroupReq`. Only `/infraWithDefaults` (using `CreateNodeGroupDynamicReq`) removes it.
>
> Response bodies (`GET`/`POST` responses returning `VmInfraInfo`) are unaffected — `postCommand`/`postCommandResult` keep their existing field names.

### Beetle Internal Logic Changes (Decision Needed)

| Code Path                                                                                                                            | Change                        | Details                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `pkg/core/recommendation/resource-sg.go` (`RecommendSecurityGroup`) + `pkg/core/migration/infra.go` (`checkAndSupportSSHAccessRule`) | Unaffected                    | Always builds explicit `FirewallRules` before calling Tumblebug; Tumblebug's auto-default is never triggered                                                                                                   |
| `pkg/core/migration/infra.go` (`CreateInfraWithDefaults`)                                                                            | Affected — no code change yet | Forwards `InfraDynamicReq` to Tumblebug as-is; if `SgTemplateId` is empty, callers now silently get an SSH-only SG instead of the previous open-all default. See Logic Change section for the decision needed. |

### imdl/cloud-model Package

| Struct                      | Breaking Change          | Migration                                       |
| --------------------------- | ------------------------ | ----------------------------------------------- |
| `InfraReq`                  | PostCommand removed      | Use `PostCommands []PostCommandReq`             |
| `InfraDynamicReq`           | PostCommand removed      | Use `PostCommands []PostCommandReq`             |
| `CreateNodeGroupDynamicReq` | NodeUserPassword removed | Remove field assignment                         |
| `AddNodeGroupDynamicReq`    | New struct (additive)    | Optional: Use for adding nodes to infra         |
| `PostCommandReq`            | New struct (additive)    | Required: Wrapper for InfraCmdReq               |
| `PostCommandStatus`         | New enum (additive)      | Optional: Check command execution status        |
| `ConnConfig`                | New field (additive)     | Optional: Use `VerifiedMessage` for diagnostics |

---

## 📊 Model Changes Reference

### Modified Structs

#### InfraReq

```diff
type InfraReq struct {
    Name string `json:"name"`
    NodeGroups []CreateNodeGroupReq `json:"nodeGroups"`

-   PostCommand InfraCmdReq `json:"postCommand"`
+   PostCommands []PostCommandReq `json:"postCommands,omitempty"`
+   PostCommandAsync bool `json:"postCommandAsync,omitempty"`
}
```

#### InfraDynamicReq

```diff
type InfraDynamicReq struct {
    Name string `json:"name"`
    NodeGroups []CreateNodeGroupDynamicReq `json:"nodeGroups"`

-   PostCommand InfraCmdReq `json:"postCommand"`
+   PostCommands []PostCommandReq `json:"postCommands,omitempty"`
+   PostCommandAsync bool `json:"postCommandAsync,omitempty"`
}
```

#### CreateNodeGroupDynamicReq

```diff
type CreateNodeGroupDynamicReq struct {
    Name string `json:"name"`
    RootDiskType string `json:"rootDiskType,omitempty"`

-   NodeUserPassword string `json:"nodeUserPassword,omitempty"`
+   DistributeSubnets bool `json:"distributeSubnets,omitempty"`
}
```

### New Structs (Additive)

#### PostCommandReq

```go
type PostCommandReq struct {
    InfraCmdReq
    NodeGroupId     string `json:"nodeGroupId,omitempty"`
    NodeId          string `json:"nodeId,omitempty"`
    LabelSelector   string `json:"labelSelector,omitempty"`
    ContinueOnError bool   `json:"continueOnError,omitempty"`
}
```

#### PostCommandStatus

```go
type PostCommandStatus string
const (
    PostCommandStatusNone                PostCommandStatus = "None"
    PostCommandStatusCompleted           PostCommandStatus = "Completed"
    PostCommandStatusCompletedWithErrors PostCommandStatus = "CompletedWithErrors"
    PostCommandStatusFailed              PostCommandStatus = "Failed"
    PostCommandStatusSkipped             PostCommandStatus = "Skipped"
    PostCommandStatusRunning             PostCommandStatus = "Running"
)
```

---

## 🔧 Deployment Configuration

### Docker Compose Changes

**File:** `deployments/docker-compose/docker-compose.yaml`

```yaml
services:
  cb-spider:
    image: cloudbaristaorg/cb-spider:0.12.42
    depends_on:
      cb-spider-postgres:
        condition: service_healthy
    environment:
      - SPIDER_METADB_ENDPOINT=cb-spider-postgres:5432
      - SPIDER_METADB_USER=${SP_POSTGRES_USER}
      - SPIDER_METADB_PASSWORD=${SP_POSTGRES_PASSWORD}
      - SPIDER_METADB_DATABASE=${SP_POSTGRES_DB}

  cb-spider-postgres:
    image: postgres:16-alpine
    container_name: cb-spider-postgres
    volumes:
      - ./data/cb-spider-container/meta_db/postgres:/var/lib/postgresql/data
    environment:
      - POSTGRES_USER=${SP_POSTGRES_USER:-spider}
      - POSTGRES_PASSWORD=${SP_POSTGRES_PASSWORD:-spider}
      - POSTGRES_DB=${SP_POSTGRES_DB:-cb_spider}
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U $$POSTGRES_USER -d $$POSTGRES_DB"]
      interval: 10s
      timeout: 5s
      retries: 5

  cb-tumblebug:
    image: cloudbaristaorg/cb-tumblebug:0.12.30
```

**File:** `deployments/docker-compose/.env`

```bash
# CB-Spider PostgreSQL Setup (NEW)
SP_POSTGRES_USER=spider
SP_POSTGRES_PASSWORD=spider
SP_POSTGRES_DB=cb_spider
```

---

## 🚀 Migration Examples

### Example 1: Simple Infrastructure Migration (Low Priority - `/infraWithDefaults`)

> **⚠️ Low Priority:** `/infraWithDefaults` is not actively used by upstream subsystems. This example is for reference only.

**Before (v0.5.9 API Request):**

```bash
curl -X POST "http://beetle:8056/beetle/migration/ns/test/infraWithDefaults" \
  -u default:default \
  -H "Content-Type: application/json" \
  -d '{
    "name": "web-infra",
    "postCommand": {
      "userName": "cb-user",
      "command": ["apt-get update"]
    },
    "nodeGroups": [{
      "name": "web",
      "nodeUserPassword": "mypassword"
    }]
  }'
```

**After (v0.5.10 API Request):**

```bash
curl -X POST "http://beetle:8056/beetle/migration/ns/test/infraWithDefaults" \
  -u default:default \
  -H "Content-Type: application/json" \
  -d '{
    "name": "web-infra",
    "sgTemplateId": "sg-usecase-web",
    "postCommands": [{
      "userName": "cb-user",
      "command": ["apt-get update"]
    }],
    "nodeGroups": [{
      "name": "web"
    }]
  }'
```

### Example 2: Go Code Migration - `imdl/cloud-model` (Low Priority)

> **⚠️ Low Priority:** This example uses `InfraDynamicReq` (for `/infraWithDefaults`). For `InfraReq` (for `/infra`), only the `postCommand` → `postCommands` change applies (no `nodeUserPassword` removal).

**Before (v0.5.9):**

```go
import cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

req := cloudmodel.InfraDynamicReq{
    Name: "my-infra",
    PostCommand: cloudmodel.InfraCmdReq{
        UserName: "cb-user",
        Command:  []string{"apt-get update"},
    },
    NodeGroups: []cloudmodel.CreateNodeGroupDynamicReq{
        {
            Name:             "worker",
            NodeUserPassword: "mypassword",
        },
    },
}
```

**After (v0.5.10):**

```go
import cloudmodel "github.com/cloud-barista/cm-beetle/imdl/cloud-model"

req := cloudmodel.InfraDynamicReq{
    Name:         "my-infra",
    SgTemplateId: "sg-usecase-web",
    PostCommands: []cloudmodel.PostCommandReq{
        {
            InfraCmdReq: cloudmodel.InfraCmdReq{
                UserName: "cb-user",
                Command:  []string{"apt-get update"},
            },
        },
    },
    NodeGroups: []cloudmodel.CreateNodeGroupDynamicReq{
        {
            Name: "worker",
        },
    },
}
```

---

## 🔎 Troubleshooting

### Issue: "Field 'postCommand' not found"

**Cause:** You're sending v0.5.9 request format to v0.5.10 API

**Solution:**

- Change `postCommand` → `postCommands`
- Wrap in array: `"postCommands": [{...}]`

### Issue: "Port 80/443 not accessible after deployment"

**Cause:** Using `POST /infraWithDefaults` without `sgTemplateId` now gets an SSH-only security group (see Logic Change section)

**Workarounds (until the default-handling decision is made):**

1. Add ports manually after creation (see Logic Change section, Options)
2. Use security group template: `"sgTemplateId": "sg-usecase-web"`

### Issue: "Spider metadata lost after upgrade"

**Cause:** Spider switched from SQLite to PostgreSQL

**Solution:**

- Re-register CSP credentials
- No data migration available

### Issue: "Field 'nodeUserPassword' not found"

**Cause:** Field removed in v0.5.10

**Solution:**

- Remove field from requests
- Linux VMs use SSH keys (no action needed)
- Windows VMs have auto-generated passwords (retrieve from CSP console)

---

## 📚 Reference

### Official Documentation

- [CB-Tumblebug v0.12.30 Release](https://github.com/cloud-barista/cb-tumblebug/releases/tag/v0.12.30)
- [Upgrade Guide v0.12.25→v0.12.30](https://github.com/cloud-barista/cb-tumblebug/discussions/2664)

### CM-Beetle Documentation

- [API Development Guide](../api-development-guide.md)
- [API Response Policy](../api-response-policy.md)
- [Installation and Execution](../installation-and-execution.md)

### Related Changes

- [CB-Tumblebug Sync Status](../../deployments/docker-compose/cb-tumblebug/SYNC.md)

---

## 🔖 Document Summary

### Breaking Changes - High Priority (Action Required)

1. 🔴 **API Request Format:** `postCommand` → `postCommands` (array) — `/infra` endpoint
2. 🟡 **Auto-SG Behavior:** Default security groups now SSH-only (affects dynamic path if not setting `sgTemplateId`)

### Breaking Changes - Low Priority (Reference Only)

> `/infraWithDefaults` is not actively used by upstream subsystems

1. 🔴 **API Request Format:** `postCommand` → `postCommands` (array) — `/infraWithDefaults` endpoint
2. 🔴 **API Request Format:** `nodeUserPassword` field removed — `/infraWithDefaults` only (note: `/infra` still has this field)

### Logic Changes (Decision Needed)

1. ℹ️ Auto-generated security groups now SSH-only by default — CM-Beetle's `CreateInfraWithDefaults` (dynamic path) does not yet set a default `sgTemplateId`; maintainers must decide whether to support the previous open-all behavior or accept the new default

### Deployment Changes (Action Required)

1. 🔧 **Spider PostgreSQL:** Re-register CSP credentials after upgrade
2. 🔧 **Version Updates:** TB 0.12.30, Spider 0.12.42, MapUI 0.12.56

### New Features (Optional)

1. ✅ Multi-phase bootstrap with targeting
2. ✅ Multi-AZ subnet distribution
3. ✅ Credential verification diagnostics
4. ✅ 10+ new structs in `imdl/cloud-model`

---

## Version Information

| Component    | v0.5.9   | v0.5.10  | Notes                        |
| ------------ | -------- | -------- | ---------------------------- |
| CM-Beetle    | v0.5.9   | v0.5.10  | This release                 |
| CB-Tumblebug | v0.12.25 | v0.12.30 | Model changes + new features |
| CB-Spider    | 0.12.35  | 0.12.42  | PostgreSQL METADB support    |
| CB-MapUI     | 0.12.50  | 0.12.56  | UI improvements              |
