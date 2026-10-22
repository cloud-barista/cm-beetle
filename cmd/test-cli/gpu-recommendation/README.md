# CM-Beetle GPU Recommendation Test CLI

Automated testing tool for evaluating CM-Beetle's GPU-accelerated infrastructure recommendation and 6-dimension match rate evaluation without provisioning actual cloud VMs.

---

## 1. Sample Source Workload Specification (Input)

The default test scenario uses an on-premise AI inference / lightweight training server equipped with an NVIDIA Tesla T4 GPU:

- **Source JSON Fixture**: [testconf/source-infra-gpu-t4.json](testconf/source-infra-gpu-t4.json)
- **Test Config Reference**: [testconf/test-config.yaml](testconf/test-config.yaml) (`beetle.requestBodyFile`)
- **API Request Mapping**: Sent as the `onpremiseInfraModel` field in the request body of `POST /beetle/recommendation/infra`

| Category             | Specification      | Details                                               |
| :------------------- | :----------------- | :---------------------------------------------------- |
| **Server ID**        | `nvidiat4-1`       | Standalone AI workload node                           |
| **CPU**              | 4 vCPUs            | 1 Socket, 2 Cores, 4 Threads (Architecture: `x86_64`) |
| **Memory**           | 16 GiB             | DDR4 RAM                                              |
| **Root Disk**        | 100 GB             | SSD                                                   |
| **Operating System** | Ubuntu 22.04.4 LTS | Codename: `jammy`, Kernel Linux 5.15+                 |
| **GPU Accelerator**  | 1x NVIDIA Tesla T4 | 16 GB VRAM, Driver: `570.195.03`, CUDA: `12.8`        |

---

## 2. Sample Recommendation Results by Target CSP (Output)

The following table summarizes the actual recommended VM specs, GPU configurations, OS images, and multi-dimensional match rates evaluated by CM-Beetle across the supported CSPs for the input source workload (**`nvidiat4-1`**: 4 vCPU, 16 GiB RAM, 1x Tesla T4 16GB, Ubuntu 22.04):

| Target CSP / Region                   | Recommended Node Spec                                                                                                                         | Recommended OS Image                                                                                         |              Overall Match Rate               |        Status        | Notes                                                                                                                                  |
| :------------------------------------ | :-------------------------------------------------------------------------------------------------------------------------------------------- | :----------------------------------------------------------------------------------------------------------- | :-------------------------------------------: | :------------------: | :------------------------------------------------------------------------------------------------------------------------------------- |
| **AWS**<br>`ap-northeast-2`           | **`g4dn.xlarge`**<br>4 vCPU, 16 GiB RAM<br>GPU: 1x NVIDIA Tesla T4 (16 GB)<br>(`aws+ap-northeast-2+g4dn.xlarge`)                          | `ami-012a353bb3afb92ee`<br>(Canonical Ubuntu Server 22.04 LTS Base)                                         | Min: 100.0%<br>Max: 100.0%<br>**Avg: 100.0%** |   `highly-matched`   | **Exact 100% Match across all 6D**<br>CPU, Mem, Image (Clean Base), GPUCount, GPUVRAM, GPUModel all 100% exact match.                |
| **GCP**<br>`asia-northeast3`          | **`g2-standard-4`**<br>4 vCPU, 16 GiB RAM<br>GPU: 1x NVIDIA L4 (24 GB)<br>(`gcp+asia-northeast3+g2-standard-4`)                           | `ubuntu-2204-jammy-v20260820`<br>(Google Cloud Ubuntu 22.04 LTS Base)                                       |  Min: 75.0%<br>Max: 100.0%<br>**Avg: 92.7%**  | `partially-matched`  | **L4 Match with Clean Base OS**<br>CPU 100%, Mem 97.7%, Image 83.3%, GPUCount 100%, VRAM 100%, GPUModel 75% (L4 vs T4).               |
| **NCP**<br>`kr`                       | **`gp1t4-g1-h50`**<br>4 vCPU, 16 GiB RAM<br>GPU: 1x NVIDIA Tesla T4 (16 GB)<br>(`ncp+kr+gp1t4-g1-h50`)                                    | `23214590`<br>(Ubuntu 22.04 Standard Base)                                                                   |  Min: 66.7%<br>Max: 100.0%<br>**Avg: 91.1%**  | `partially-matched`  | **Native T4 Match with Clean Base OS**<br>CPU 100%, Mem 80.0%, Image 66.7%, GPUCount 100%, GPUVRAM 100%, GPUModel 100%.              |
| **NHN Cloud**<br>`kr1`                | **`g2.t4.c4m32`**<br>4 vCPU, 32 GiB RAM<br>GPU: 1x NVIDIA Tesla T4 (16 GB)<br>(`nhn+kr1+g2.t4.c4m32`)                                     | `0f07c795-2a46-44fc-a61b-fa0d96763ce2`<br>(Ubuntu Server 22.04 LTS Base)                                    |  Min: 50.0%<br>Max: 100.0%<br>**Avg: 86.1%**  | `partially-matched`  | **Native T4 Match with Clean Base OS**<br>CPU 100%, Mem 50% (32GB bundle), Image 66.7%, GPUCount 100%, GPUVRAM 100%, GPUModel 100%. |
| **Alibaba Cloud**<br>`ap-northeast-2` | **`ecs.gn7i-c8g1.2xlarge`**<br>8 vCPU, 32 GiB RAM<br>GPU: 1x NVIDIA A10 (24 GB)<br>(`alibaba+ap-northeast-2+ecs.gn7i-c8g1.2xlarge`)       | `ubuntu_22_04_x64_20G_alibase_20260810.vhd`<br>(Alibaba Ubuntu 22.04 Base)                                  |  Min: 50.0%<br>Max: 100.0%<br>**Avg: 74.2%**  | `partially-matched`  | **A10 Match with Clean Base OS**<br>8 vCPU, 32GB RAM, 1x A10 24GB, Image 66.7%, GPUCount 100%, GPUVRAM 100%, GPUModel 75%.         |
| **Azure**<br>`koreasouth`             | **`standard_nd96isr_h200_v5`**<br>96 vCPU, 1900 GiB RAM<br>GPU: 8x NVIDIA H200 (141 GB)<br>(`azure+koreasouth+standard_nd96isr_h200_v5`)  | `Canonical:0001-com-ubuntu-server-jammy:22_04-lts-gen2:22.04.202608060`<br>(Canonical Ubuntu Server 22.04) |  Min: 0.9%<br>Max: 100.0%<br>**Avg: 46.0%**   | `partially-matched`  | **Clean Base OS with Available GPU**<br>Canonical official base image selected for target GPU spec in region.                         |
| **Tencent Cloud**<br>`ap-seoul`       | N/A                                                                                                                                          | N/A                                                                                                          |                     0.0%                      | `nothing-to-recommend` | No compatible GPU spec and OS image pair registered in the local catalog.                                                             |
| **IBM Cloud**<br>`au-syd`             | N/A                                                                                                                                          | N/A                                                                                                          |                     0.0%                      | `nothing-to-recommend` | No compatible GPU spec and OS image pair registered in the local catalog.                                                             |

> [!NOTE]
>
> - **Evaluation Dimensions (6D)**: CPU, Memory, Image, GPUCount, GPUVRAM, GPUModel.
> - **Status Classification**:
>   - `highly-matched`: All active dimensions meet or exceed `minMatchRate` (default: 85.0%).
>   - `partially-matched`: Valid candidate where some dimensions deviate (e.g., higher memory bundle or different GPU model tier).
>   - `nothing-to-recommend`: No compatible spec-image pairs found in target region.

---

## 3. Quick Start

Run recommendation tests across all configured CSPs:

```bash
# Execute from the cm-beetle root directory
make test-gpu-recommendation
```

Or run directly from this directory:

```bash
cd cmd/test-cli/gpu-recommendation
go run main.go -config testconf/test-config.yaml
```

---

## 4. Configuration

`testconf/test-config.yaml`:

```yaml
test:
  set:
    mode: sequential # sequential or parallel
  cases:
    - csp: aws
      region: ap-northeast-2
      name: AWS-Seoul
      execute: true
    - csp: azure
      region: koreasouth
      name: Azure-Busan
      execute: true
    - csp: gcp
      region: asia-northeast3
      name: GCP-Seoul
      execute: true
    - csp: alibaba
      region: ap-northeast-2
      name: Alibaba-Seoul
      execute: true
    - csp: tencent
      region: ap-seoul
      name: Tencent-Seoul
      execute: true
    - csp: ibm
      region: au-syd
      name: IBMCloud-Sydney
      execute: true
    - csp: ncp
      region: kr
      name: NCP-Seoul
      execute: true
    - csp: nhn
      region: kr1
      name: NHNCloud-Pangyo
      execute: true

beetle:
  endpoint: http://localhost:8056
  requestBodyFile: testconf/source-infra-gpu-t4.json
  authConfigFile: testconf/auth-config.json
```

---

## 5. Output Reports

- **Console**: Displays real-time progress, candidate details, and 6-dimension scores.
- **Markdown Report**: Written to `testresult/gpu-recommendation-test-report.md`.

---

## 6. Architecture Strategy: OS Image Matching & GPU Stack Setup

1. **Standard Base OS Image Priority**:
   - Standard clean base OS images (e.g. Canonical Ubuntu 22.04 LTS) are prioritized as the primary recommendation target across all CSPs.
   - Eliminates artificial non-GPU image penalties, preserving genuine 100% OS similarity for matching base distributions.
   - Prevents catalog search failures (`nothing-to-recommend`) in CSPs or regions where pre-baked GPU images are not registered.
2. **GPU Image as Graceful Fallback**:
   - Pre-baked GPU images (`IsGPUImage=true`) are relegated to a secondary fallback when no basic OS image is registered in the target catalog.
3. **Roadmap: Idempotent GPU Node Bootstrap Script & API**:
   - Provide an idempotent initialization script and dedicated API (`POST /beetle/.../gpu-bootstrap`) that detects, configures, and verifies the full accelerator stack:
     - NVIDIA Kernel Driver
     - CUDA Toolkit and Runtime
     - NVIDIA Container Toolkit (`nvidia-docker2` runtime)
     - Fabric Manager (for multi-GPU NVLink nodes)
   - Guarantees a predictable, reproducible execution environment regardless of whether the node was booted from a base OS or a pre-baked image.
