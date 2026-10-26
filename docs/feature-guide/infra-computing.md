# Infrastructure

This document covers recommendation and migration support for computing infrastructure resources (VM, VNet, Subnet, Security Group, SSH Key) across supported Cloud Service Providers (CSPs), along with the architecture and execution control model of the CM-Beetle computing infrastructure migration pipeline.

## Support Status

| Resource        | AWS | Azure | GCP | Alibaba | Tencent | IBM | (OpenStack) | NCP | NHN | (KT) |
| --------------- | :-: | :---: | :-: | :-----: | :-----: | :-: | :---------: | :-: | :-: | :--: |
| VNet            | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |     📅      | ✅  | ✅  |  -   |
| Subnet          | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |     📅      | ✅  | ✅  |  -   |
| Security Group  | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |     📅      | ✅  | ✅  |  -   |
| SSH Key         | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |     📅      | ✅  | ✅  |  -   |
| Node Spec       | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |     📅      | ✅  | ✅  |  -   |
| OS Image        | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |     📅      | ✅  | ✅  |  -   |
| Virtual Machine | ✅  |  ✅   | ✅  |   ✅    |   ✅    | ✅  |     📅      | ✅  | ✅  |  -   |

## Supported Operations

| Operation      | Supported Resources                                              |
| -------------- | ---------------------------------------------------------------- |
| Recommendation | VNet, Subnet, Security Group, VM Spec, OS Image, Virtual Machine |
| Migration      | VNet, Subnet, Security Group, SSH Key, Virtual Machine           |

## Migration Pipeline Architecture & Execution Flow

The CM-Beetle computing infrastructure migration pipeline automates provisioning, non-blocking execution control, outbound rate limiting, and end-to-end verification across target CSPs:

```mermaid
flowchart TD
    subgraph ClientLayer["[Client / Portal Layer]"]
        Req["1. Migration Request\nPOST /migration/ns/{nsId}/infra\n(Prefer: respond-async, useExisting=true)"]
        Poll["2. Async Status Polling\nGET /request/{reqId}"]
        VerifyReq["4. SSH Readiness Check\nGET .../infra/{infraId}/ssh-ready"]
    end

    subgraph PipelineEngine["[CM-Beetle Computing Infra Migration Pipeline]"]
        direction TB
        
        subgraph ExecControl["1) Execution Control (Non-blocking Async)"]
            AsyncHandler["HTTP 202 Accepted (Immediate Return)\nWorker Pool: Max 20 Concurrent Jobs"]
            StatusStore["Request State Store (LKVStore)\nHandling -> Success / Error"]
        end

        subgraph IdempotentEngine["2) Idempotent Resource Provisioning"]
            DetectExisting{"Detect Existing Resources\n(vNet / Subnet / SG / SSH Key)"}
            ReuseResource["Reuse Existing Resources on Matching Connection\n(Avoid Creation Collisions)"]
            CreateNew["Provision New Resources with NameSeed"]
            DetectExisting -->|Found| ReuseResource
            DetectExisting -->|Not Found| CreateNew
        end

        subgraph OutboundPacer["3) Outbound Rate Limiting & Retry Control"]
            Pacer["Outbound Call Pacer\n(4.0 req/s Token Bucket Shaping)"]
            RetryLimit["HTTP 429 Bounded Retry with Exponential Backoff\n(Retry-After Header Compliance)"]
            Pacer --> RetryLimit
        end

        subgraph ValidationEngine["4) Result Verification & Inbound Protection"]
            SSHCooldown["SSH Check Cooldown Middleware\n(1 Check per 30s per Infrastructure)"]
            SSHProbe["Per-Node SSH Handshake & Command Probing\n(10s Probe Interval, 30s Max Timeout)"]
            SSHCooldown --> SSHProbe
        end
    end

    subgraph Downstream["[Downstream Subsystems & CSPs]"]
        Tumblebug["CB-Tumblebug & CSP APIs\n(Infra Provisioning & Remote Cmd)"]
        TargetVMs["Target Cloud VMs\n(Cloud-init & SSH Access Ready)"]
    end

    %% Flow Connections
    Req --> AsyncHandler
    AsyncHandler --> StatusStore
    Poll -.-> StatusStore
    
    AsyncHandler --> DetectExisting
    ReuseResource --> Pacer
    CreateNew --> Pacer
    
    RetryLimit --> Tumblebug
    Tumblebug --> TargetVMs
    
    VerifyReq --> SSHCooldown
    SSHProbe --> TargetVMs
```

### Core Pipeline Capabilities

1. **Idempotent Resource Provisioning (`useExisting`)**:
   - By default (`useExisting=true`), the pipeline detects pre-existing network and security primitives (`vNet`, `subnet`, `securityGroup`, `sshKey`) in the target namespace under the requested CSP connection.
   - Reuses matching resources instead of recreating them, preventing name collisions and enabling safe repeated executions.
   - For details, see [Target Infrastructure Validation](core-validation.md) and [NameSeed: Late Binding for Resource Names](core-naming-rule.md).

2. **Non-blocking Async Execution & State Tracking (`Prefer: respond-async`)**:
   - Supports RFC 7240 `Prefer: respond-async` headers for long-running infrastructure provisioning.
   - Immediately returns `HTTP 202 Accepted` along with a unique request ID (`reqId`), executing the provisioning pipeline in a bounded background worker pool (maximum 20 concurrent jobs).
   - Callers track execution progress via `GET /beetle/request/{reqId}` until status transitions from `Handling` to `Success` or `Error`.
   - For details, see [Async Responses](core-async-response.md).

3. **Outbound Rate Limiting & Retry Control (Pacer & Bounded Retry)**:
   - Enforces proactive traffic shaping via a token bucket `Pacer` targeting 4.0 req/s (burst 1), keeping CM-Beetle safely under downstream CB-Tumblebug rate limits (5.0 req/s on infra endpoints).
   - Recovers from unexpected HTTP 429 responses using exponential backoff and bounded retries adhering to `Retry-After` headers.
   - For details, see [Tumblebug Call Pacer](../tumblebug-call-pacer.md) and [pkg/ratelimit/README.md](../../pkg/ratelimit/README.md).

4. **Target VM Remote Accessibility Verification (`ssh-ready` API)**:
   - Provides `GET /migration/ns/{nsId}/infra/{infraId}/ssh-ready?option=detail` to confirm that all provisioned VMs have completed cloud-init and are actively accepting remote SSH commands.
   - Protected by `SSHCheckCooldown` middleware (maximum 1 check per 30 seconds per infrastructure) to prevent SSH server overload.
   - For details, see [GPU Infrastructure Migration API Guide (Section 6.1)](../gpu-infra-migration-api-guide.md#61-타겟-인프라-ssh-접속-준비-상태-확인-api-get-ssh-ready).

## Roadmap

- **Planned**: OpenStack
