# NameSeed: Late Binding for Resource Names

## What is NameSeed?

`nameSeed` is a short prefix that CM-Beetle applies to **all resource names at the moment of migration** — not during recommendation.

```
nameSeed=blue  +  base name vnet-01  →  blue-vnet-01
```

It is passed exclusively as a **query parameter** on migration APIs:

```
POST /beetle/migration/ns/{nsId}/infra?nameSeed=blue
POST /beetle/migration/ns/{nsId}/objectStorage?nameSeed=blue
```

---

## Why Late Binding?

The recommendation model holds **base names** (e.g., `vnet-01`, `sg-01`) that are human-readable and easy to inspect or adjust.
Pinning a final deployment identity at the last moment keeps the model reusable across environments:

| Environment | Same model + different seed | Result          |
| ----------- | --------------------------- | --------------- |
| Dev         | `?nameSeed=dev`             | `dev-vnet-01`   |
| Staging     | `?nameSeed=stg`             | `stg-vnet-01`   |
| Production  | `?nameSeed=prod`            | `prod-vnet-01`  |
| (none)      | _(omit query param)_        | `vnet-01`       |

---

## What NameSeed Is NOT

| Misconception                              | Reality                                                   |
| ------------------------------------------ | --------------------------------------------------------- |
| A field in the request body                | It is a query parameter only                              |
| Stored in the recommendation model         | The model always holds base names; seed is never persisted |
| A resource ID or namespace                 | It is a pure name prefix                                  |
| Required                                   | Optional — omitting it leaves base names unchanged        |

---

## How It Works

`{nameSeed}-{baseName}` is applied to every resource name in the model at migration time.
Internal references (e.g., `vNetId`, `subnetId`, `sshKeyId`) are updated consistently so the model stays coherent.

```
Base model                  After ?nameSeed=blue
──────────────────────────  ─────────────────────────────
vNet.name       = vnet-01   vNet.name       = blue-vnet-01
sg.vNetId       = vnet-01   sg.vNetId       = blue-vnet-01
sg.name         = sg-01     sg.name         = blue-sg-01
subnet.name     = snet-01   subnet.name     = blue-snet-01
nodeGroup.name  = ng-01     nodeGroup.name  = blue-ng-01
```

---

## Validation Rules

| Rule                            | Detail                                    |
| ------------------------------- | ----------------------------------------- |
| Optional                        | Empty string → no prefix applied          |
| Max length                      | 20 characters                             |
| Allowed characters              | Alphanumeric (`a-z`, `A-Z`, `0-9`) and hyphens (`-`) |
| First character                 | Must be alphanumeric (not a hyphen)       |

Invalid examples: `_blue`, `-prod`, `my seed`, `averylongseedthatexceedslimit`

---

## Typical Workflow

```
┌──────────────────┐   ┌──────────────────┐   ┌──────────────────┐   ┌──────────────────┐
│  1. Recommend    │ → │  2. Align        │ → │  3. Preview      │ → │  4. Migrate      │
│                  │   │   (optional)     │   │   (optional)     │   │                  │
│ Returns base     │   │ Rename a         │   │ Dry-run: shows   │   │ Applies seed and │
│ names (no seed)  │   │ resource and     │   │ final names with │   │ creates resources│
│                  │   │ propagate refs   │   │ seed applied     │   │                  │
└──────────────────┘   └──────────────────┘   └──────────────────┘   └──────────────────┘
POST /recommendation/  POST /naming/           POST /naming/           POST /migration/
infra                  alignment               preview?nameSeed=blue   infra?nameSeed=blue
```

Use the **Preview** step to verify names before committing to resource creation.

---

## Related

- [Align Names API Guide](../api-guide-align-names.md) — rename a resource and propagate references
- [Object Storage Migration Feature Guide](object-storage-migration-feature-guide.md) — NameSeed usage for object storage
