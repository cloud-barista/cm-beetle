# CM-Beetle Docs

Documentation site for CM-Beetle, computing infrastructure migration.

## Feature Guides

- **Core & System**
  - [Async Responses](feature-guide/core-async-response.md) – Non-blocking async response model and polling workflow.
  - [NameSeed: Late Binding for Resource Names](feature-guide/core-naming-rule.md) – Concept, rules, and workflow for deterministic resource naming.
  - [Validation Guide](feature-guide/core-validation.md) – Pre-flight dry-run validation rules for target infrastructure.
- **Infrastructure**
  - [Computing Infrastructure](feature-guide/infra-computing.md) – VM, VNet, Subnet, and Security Group migration support.
  - [Kubernetes Infrastructure](feature-guide/infra-kubernetes.md) – Managed Kubernetes cluster recommendation and migration.
- **Managed Middleware**
  - [Managed Middleware Overview](feature-guide/middleware-overview.md) – Multi-cloud support matrix for managed middleware.
  - [Network Load Balancer (NLB)](feature-guide/middleware-nlb.md) – NLB backend topology grouping and migration.
  - [Managed Object Storage](feature-guide/middleware-object-storage.md) – Object storage bucket recommendation, feature adjustment, and migration.
  - [Managed RDBMS (RDS)](feature-guide/middleware-rdbms.md) – Relational database recommendation, referential validation, and migration.
- **Data**
  - [Data Transfer](feature-guide/data-transfer.md) – Server filesystem and data transfer capabilities and workflow.

## API Guides

- [API Guide: Align Names](api-guide-align-names.md) – Propagate name changes to dependent resources (Late Binding, NameSeed).
- [API Development Guide](api-development-guide.md) – Guidelines for developing CM-Beetle APIs.
- [API Response Policy](api-response-policy.md) – Guidelines for API response message formatting.
- [Module Import Guide](module-import-guide.md) – How to import and use CM-Beetle modules.

## Installation & Setup

- [Installation and Execution](installation-and-execution.md) – How to install and run CM-Beetle.

## Testing & Integration

- [Test Results: Data Migration](test-results-data-migration.md) – Data migration test results and examples.
- [Test Results: Object Storage](test-results-object-storage.md) – Object storage test results and examples.
- [Beetle v0.3.0 Integration Testing](beetle-v0.3.0-integration-and-testing-with-tumblebug-honeybee-and-model.md) – Integration test results with CB-Tumblebug, Honeybee, and Model.
- [Beetle v0.4.0 Integration Testing](beetle-v0.4.0-integration-and-testing-with-tumblebug-honeybee-and-model.md) – Integration test results for v0.4.0.

## Dependency Upgrade Guides

- [API Upgrade Guide: TB v0.12.9](api-upgrade-guide-tb-v0.12.9.md) – API changes and upgrade steps for CB-Tumblebug v0.12.9.

## CM-Beetle Breaking Changes

- **[Breaking Changes](changes/)** – CM-Beetle release notes for breaking changes affecting Beetle API/`imdl` consumers
  - [v0.5.10 Breaking Changes](changes/BREAKING_CHANGES_v0.5.10.md) – `postCommand` → `postCommands`, `nodeUserPassword` removed from dynamic node groups

## Developer Resources

- [Tumblebug Call Pacer](tumblebug-call-pacer.md) – How CM-Beetle stays under CB-Tumblebug's rate limit, and how to maintain it.
- [AI Context](ai-context-for-us.md) – Context information for AI-assisted development.
