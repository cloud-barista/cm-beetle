# API Migration Guide — CM-Beetle (TB v0.12.9 Alignment)

> **AI-assisted**: Generated with [GitHub Copilot](https://github.com/features/copilot) in VS Code (Model: Claude Sonnet 4.6)

This document describes all **breaking API changes** introduced by the TB v0.12.9 alignment,
merged on **2026-04-29** (base commit `7d0dc85`).
These changes apply to all builds and releases from that merge onward.

> **This is a clean break.** Old paths and field names are removed immediately.
> No backwards-compatibility shims or dual keys are provided.

---

## Summary of Changes

The outdated vocabulary `mci / vmInfra / vmSpec / vmOsImage` has been replaced uniformly with
`infra / infra / spec / osImage`, matching CB-Tumblebug v0.12.9 terminology.

---

## BC-1 — URL Path Changes

| Old Path                                       | New Path                                           |
| ---------------------------------------------- | -------------------------------------------------- |
| `POST /migration/ns/{nsId}/mci`                | `POST /migration/ns/{nsId}/infra`                  |
| `POST /migration/ns/{nsId}/mciWithDefaults`    | `POST /migration/ns/{nsId}/infraWithDefaults`      |
| `GET /migration/ns/{nsId}/mci`                 | `GET /migration/ns/{nsId}/infra`                   |
| `GET /migration/ns/{nsId}/mci/{mciId}`         | `GET /migration/ns/{nsId}/infra/{infraId}`         |
| `DELETE /migration/ns/{nsId}/mci/{mciId}`      | `DELETE /migration/ns/{nsId}/infra/{infraId}`      |
| `POST /recommendation/vmInfra`                 | `POST /recommendation/infra`                       |
| `POST /recommendation/mciWithDefaults`         | `POST /recommendation/infraWithDefaults`           |
| `GET /summary/target/ns/{nsId}/mci/{mciId}`    | `GET /summary/target/ns/{nsId}/infra/{infraId}`    |
| `POST /report/migration/ns/{nsId}/mci/{mciId}` | `POST /report/migration/ns/{nsId}/infra/{infraId}` |
| `/recommendation/resources/vmSpecs`            | `/recommendation/resources/specs`                  |
| `/recommendation/resources/vmOsImages`         | `/recommendation/resources/osImages`               |

**Client action required**: Update all request URLs in your client code.

---

## BC-2 — URL Path Parameter Changes

| Old Parameter | New Parameter |
| ------------- | ------------- |
| `{mciId}`     | `{infraId}`   |

**Client action required**: Rename the path parameter variable in your URL templates.

---

## BC-3 — JSON Response Field Changes

| Old Field                    | New Field                |
| ---------------------------- | ------------------------ |
| `targetVmInfra`              | `targetInfra`            |
| `targetVmInfraList`          | `targetInfraList`        |
| `targetVmSpec`               | `targetSpec`             |
| `targetVmSpecList`           | `targetSpecList`         |
| `targetVmOsImage`            | `targetOsImage`          |
| `targetVmOsImageList`        | `targetOsImageList`      |
| `recommendedVmSpecList`      | `recommendedSpecList`    |
| `recommendedVmOsImageList`   | `recommendedOsImageList` |
| `recommendedVmInfraModel`    | `recommendedInfraModel`  |
| `mciId` (in report metadata) | `infraId`                |

**Client action required**: Update JSON deserialization mappings and struct tags.

---

## BC-4 — Swagger Schema Type Name Changes

| Old Type                        | New Type                      |
| ------------------------------- | ----------------------------- |
| `RecommendedVmInfra`            | `RecommendedInfra`            |
| `RecommendedVmInfraModel`       | `RecommendedInfraModel`       |
| `RecommendedVmInfraDynamic`     | `RecommendedInfraDynamic`     |
| `RecommendedVmInfraDynamicList` | `RecommendedInfraDynamicList` |
| `RecommendedVmSpec`             | `RecommendedSpec`             |
| `RecommendedVmSpecList`         | `RecommendedSpecList`         |
| `RecommendedVmOsImage`          | `RecommendedOsImage`          |
| `RecommendedVmOsImageList`      | `RecommendedOsImageList`      |

**Client action required**: Regenerate any auto-generated SDK or type stubs from the updated Swagger spec.

---

## BC-5 — Naming Enum Value Change

The `resourceType` enum used by the naming/alignment API has changed:

| Old Value | New Value |
| --------- | --------- |
| `mci`     | `infra`   |

**Client action required**: Update any `resourceType=mci` enum references to `resourceType=infra`.

---

## BC-6 — Removed Endpoints

The following endpoint has been **removed entirely** (it was already deprecated):

| Removed Endpoint                      | Replacement                                                         |
| ------------------------------------- | ------------------------------------------------------------------- |
| `POST /recommendation/mci`            | Use `POST /recommendation/infra`                                    |
| `POST /recommendation/containerInfra` | No direct replacement (K8s recommendation uses dedicated endpoints) |

---

## Before / After Examples

### Migrate Infrastructure

**Before (v0.4.x):**

```bash
curl -X POST http://localhost:8056/beetle/migration/ns/mig01/mci \
  -H "Content-Type: application/json" \
  -d '{ "onpremiseInfraModel": { ... } }'
```

**After (this change):**

```bash
curl -X POST http://localhost:8056/beetle/migration/ns/mig01/infra \
  -H "Content-Type: application/json" \
  -d '{ "onpremiseInfraModel": { ... } }'
```

---

### Get Specific Infrastructure

**Before (v0.4.x):**

```bash
curl http://localhost:8056/beetle/migration/ns/mig01/mci/infra01
```

**After (this change):**

```bash
curl http://localhost:8056/beetle/migration/ns/mig01/infra/infra01
```

---

### Recommend VM Infrastructure

**Before (v0.4.x):**

```bash
curl -X POST http://localhost:8056/beetle/recommendation/vmInfra \
  -H "Content-Type: application/json" \
  -d '{ ... }'
```

**After (this change):**

```bash
curl -X POST http://localhost:8056/beetle/recommendation/infra \
  -H "Content-Type: application/json" \
  -d '{ ... }'
```

---

### Target Infrastructure Summary

**Before (v0.4.x):**

```bash
curl "http://localhost:8056/beetle/summary/target/ns/mig01/mci/infra01?format=md"
```

**After (this change):**

```bash
curl "http://localhost:8056/beetle/summary/target/ns/mig01/infra/infra01?format=md"
```

---

### Migration Report

**Before (v0.4.x):**

```bash
curl -X POST http://localhost:8056/beetle/report/migration/ns/mig01/mci/infra01 \
  -H "Content-Type: application/json" \
  -d '{ "onpremiseInfraModel": { ... } }'
```

**After (this change):**

```bash
curl -X POST http://localhost:8056/beetle/report/migration/ns/mig01/infra/infra01 \
  -H "Content-Type: application/json" \
  -d '{ "onpremiseInfraModel": { ... } }'
```

---

### Recommend VM Specs

**Before (v0.4.x):**

```bash
curl -X POST http://localhost:8056/beetle/recommendation/resources/vmSpecs \
  -H "Content-Type: application/json" \
  -d '{ ... }'
```

**After (this change):**

```bash
curl -X POST http://localhost:8056/beetle/recommendation/resources/specs \
  -H "Content-Type: application/json" \
  -d '{ ... }'
```

---

## Deprecation and Sunset Policy

This is a **clean break** change. Old paths are removed immediately with no redirect:

- No 307 redirects are provided for old paths.
- Requests to old paths will return `404 Not Found`.
- Clients must update to new paths before pulling builds from commit `7d0dc85` or later.

---

## Checklist for Client Upgrade

- [ ] Update all API request URLs (`/mci` → `/infra`, `/vmInfra` → `/infra`, etc.)
- [ ] Update path parameter names (`{mciId}` → `{infraId}`)
- [ ] Update JSON response field mappings (see BC-3)
- [ ] Regenerate SDK stubs from the new Swagger spec (see BC-4)
- [ ] Update `resourceType=mci` enum to `resourceType=infra` (see BC-5)
- [ ] Remove calls to `POST /recommendation/mci` and `POST /recommendation/containerInfra` (see BC-6)
