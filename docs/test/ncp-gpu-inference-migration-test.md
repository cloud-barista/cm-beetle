# NCP GPU 추론서비스 이관 테스트 이슈 조치 및 검증 결과

- **작성일**: 2026-10-28
- **대상 환경**: CM-Beetle (`HEAD`), CB-Tumblebug (`0.13.7`), CB-Spider (`0.13.17`)

---

## 1. 공유된 이슈 요약 및 원인 분석

GPU 추론서비스 이관 테스트 중 NCP(`kr`) 타겟 구간에서 발생한 이슈와 원인은 다음과 같습니다.

| 구분         | 테스트 구간                          | 기존 증상                                                    | 원인 분석                                                                                                                                                                                                                                    |
| :----------- | :----------------------------------- | :----------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **3-1**      | `GCP (L4 24GB, Ubuntu 22.04) -> NCP` | 추천 실패 (`nothing-to-recommend`, `targetSpecList: null`)   | NCP 3세대 L4(`gp1l4-g3` 등) 및 A100P 스펙의 카탈로그 가격이 `-1`로 등록되어 있어 Beetle의 비용 필터(`CostPerHour >= 0`)에서 전량 제외됨. 또한 `acceleratorModel`이 `"NVIDIA NA"`로 표기되어 있었음.                                          |
| **3-2**      | `AWS (T4 16GB, Ubuntu 22.04) -> NCP` | 생성 전 검증 실패 (`gp1t4-g1-h50` + `23789321` incompatible) | 16GB 이상 조회 시 가격이 있는 1세대 T4(XEN)만 남고 L4(KVM, 가격 `-1`)는 탈락함. 이후 T4(XEN)와 호환되는 Ubuntu 이미지가 없어 매칭에 실패했으나, Fallback 로직이 첫 사양(`gp1t4-g1-h50`, XEN)과 첫 이미지(Ubuntu KVM)를 강제 결합하여 반환함. |
| **3-4**      | `GCP -> NCP T4` (수동 사양 지정)     | 생성 실패 (`No hypervisor was found that matches...`)        | NCP 1세대 T4(`gp1t4-g1-h50`)는 `XEN` 전용 스펙이며, NCP는 Ubuntu 22.04/24.04 이미지를 3세대 `KVM` 전용으로만 제공함(XEN용 리눅스 이미지는 Rocky Linux 8.8만 존재).                                                                           |
| **3-3, 3-5** | `GCP -> NCP L4` (수동 사양 지정)     | 생성 실패 (`Product Type: GPU - L4 - G3 Creation Limit: 0`)  | NCP 계정의 3세대 L4 GPU 서버 생성 한도(Quota)가 `0`으로 설정되어 있어 CSP API에서 생성을 거부함 (한도 상향 신청 완료, 승인 대기).                                                                                                            |

---

## 2. CM-Beetle 조치 사항

[Discussion #183 (Node Spec-Image Compatibility and Filtering in CM-Beetle)](https://github.com/cloud-barista/cm-beetle/discussions/183)의 정책(NCP KVM 전용 추천 및 비호환 리소스 사전 배제)에 맞추어 다음 3가지 개선을 적용했습니다.

1. **NCP KVM 전용 사전 필터링 및 가격 미등록(`-1`) 완화 (`pkg/core/recommendation/resource-node-spec-gpu.go`)**:
   - **1단계 (NCP KVM 필터)**: NCP 타겟 추천 시 구형 1세대 `XEN` 스펙(`gp1t4-g1-*`)을 사전 배제하고 3세대 `KVM` 스펙(`gp1l4-g3` 등)만 후보로 선별합니다.
   - **2단계 (가격 필터 완화)**: 선별된 후보 중 유효 가격(`CostPerHour >= 0`)을 가진 스펙이 하나라도 있으면 기존처럼 엄격히 필터링하되, 조회된 모든 KVM 후보 스펙의 가격이 `-1`(카탈로그 가격 미등록 상태)인 경우에는 후보를 버리지 않고 유지합니다(`filterPricedSpecs`).
2. **GPU 모델명 메타데이터 자동 복원 (`pkg/core/recommendation/resource-node-spec-gpu.go`)**:
   - 카탈로그의 `acceleratorModel`이 `"NVIDIA NA"` 등으로 누락된 경우, 스펙 상세 속성(`Details` 내 `ServerSpecDescription`: `"NVIDIA L4 GPU 1EA, GPUMemory 24GB..."`)에서 실제 GPU 모델명(`"NVIDIA L4"`)을 추출하여 복원합니다.
   - 이를 통해 오케스트레이터의 벤더/GPU 보유 여부 확인 조건을 충족하고, 동일 GPU 모델(`L4 -> L4`) 우선순위 정렬이 정상 동작합니다.
3. **비호환 사양 및 이미지 강제 결합 Fallback 제거 (`pkg/core/recommendation/infra-gpu.go` 외 3곳)**:
   - 호환되는 사양-이미지 쌍(`FindCompatibleVmSpecAndImagePairs`)이 없을 때 비호환 조합(`specList[0] + imageList[0]`)을 강제로 묶어 반환하던 Fallback을 제거하고, 호환 쌍이 전무할 경우 명확히 `status: "nothing-to-recommend"`를 반환하도록 통일했습니다.

---

## 3. 조치 후 테스트 및 검증 결과

동일한 소스 노드 조건으로 CM-Beetle 추천 API(`POST /beetle/recommendation/infra`) 및 사전 검증 API(`POST /beetle/validation/ns/default/infra?useExisting=false`)를 호출하여 확인한 결과는 다음과 같습니다.

| 구분          | 테스트 구간                                            | 조치 전 결과                                    | 조치 후 추천 결과 (1순위 후보)                                                                                   | 사전 검증 (`ValidateInfra`)                    |
| :------------ | :----------------------------------------------------- | :---------------------------------------------- | :--------------------------------------------------------------------------------------------------------------- | :--------------------------------------------- |
| **3-1**       | `GCP (L4 24GB, Ubuntu 22.04) -> NCP (kr)`              | `nothing-to-recommend`                          | **`ncp+kr+gp1l4-g3`** (4 vCPU, 48GiB, 1x `NVIDIA L4` 24GB, KVM) + **`23214590`** (`ubuntu-22.04-base`, KVM)      | **`valid: true`** (`issues: []`)               |
| **3-2**       | `AWS (T4 16GB, Ubuntu 22.04) -> NCP (kr)`              | `gp1t4-g1-h50` (XEN) + Ubuntu (KVM) 비호환 오류 | **`ncp+kr+gp1l4-g3`** (4 vCPU, 48GiB, 1x `NVIDIA L4` 24GB, KVM) + **`23214590`** (`ubuntu-22.04-base`, KVM)      | **`valid: true`** (`issues: []`)               |
| **기존 PASS** | `AWS (T4 16GB, Ubuntu 22.04) -> GCP (asia-northeast3)` | `g2-standard-4` 추천 및 이관 성공               | **`gcp+asia-northeast3+g2-standard-4`** (4 vCPU, 15.625GiB, 1x `NVIDIA L4` 24GB) + `ubuntu-2204-jammy-v20260820` | **`valid: true`** (`issues: []`)               |
| **3-4**       | `gp1t4-g1-h50` (XEN) + `23789321` (KVM) 수동 지정      | 생성 요청 후 NCP 하이퍼바이저 불일치 실패       | 추천 후보에서 원천 배제됨                                                                                        | **`valid: false`** (`SPEC_IMAGE_INCOMPATIBLE`) |

### 3.1 `GCP (L4 24GB) -> NCP (kr)` 및 `AWS (T4 16GB) -> NCP (kr)` 추천 후보 요약

두 시나리오 모두 NCP `kr` 리전에서 Ubuntu 22.04(KVM)와 정식 호환(`CorrespondingImageIds` 일치)되는 3세대 KVM L4 GPU 사양 3종이 우선순위대로 추천됩니다.

- **Candidate #1**: `ncp+kr+gp1l4-g3` (vCPU: 4, Memory: 48GiB, GPU: `1x NVIDIA L4 24GB`, Hypervisor: `KVM`) + Image `23214590` (`ubuntu-22.04-base (Hypervisor:KVM)`)
- **Candidate #2**: `ncp+kr+gp1l8-g3` (vCPU: 8, Memory: 48GiB, GPU: `1x NVIDIA L4 24GB`, Hypervisor: `KVM`) + Image `23214590` (`ubuntu-22.04-base (Hypervisor:KVM)`)
- **Candidate #3**: `ncp+kr+gp2l8-g3` (vCPU: 8, Memory: 96GiB, GPU: `2x NVIDIA L4 24GB`, Hypervisor: `KVM`) + Image `23214590` (`ubuntu-22.04-base (Hypervisor:KVM)`)

> [!NOTE]
>
> - **시간당 비용(`costPerHour`) 표기 안내**: CB-Tumblebug 카탈로그에 해당 리전/세대 스펙의 가격 정보가 아직 등록되지 않은 경우, 추천된 스펙(`targetSpecList`)의 `costPerHour` 값이 제공되지 않을 수 있습니다(`-1`로 반환). 이 경우에도 하드웨어 적합도(GPU 제조사, 개수, VRAM, 칩셋 모델, 호스트 vCPU/메모리) 및 OS 이미지 호환성 검증은 정상 수행되며, 추후 카탈로그에 가격 정보가 반영되면 별도 조치 없이 실제 비용이 표시됩니다.
> - **프로비저닝 한도 안내**: NCP 계정의 L4(G3) GPU 서버 생성 한도 상향이 승인되면, 위 추천 결과를 그대로 사용하여 `GCP -> NCP` 및 `AWS -> NCP` 프로비저닝을 진행할 수 있습니다.

---

## 4. 재현 및 확인용 API 호출 예제

### 4.1 `GCP (L4 24GB, Ubuntu 22.04) -> NCP (kr)` 추천 확인

```bash
curl -s -u default:default -X POST http://localhost:8056/beetle/recommendation/infra \
  -H "Content-Type: application/json" \
  -d '{
    "desiredCspAndRegionPair": { "csp": "ncp", "region": "kr" },
    "onpremiseInfraModel": {
      "nodes": [
        {
          "machineId": "lxc-gcp-10271025-gpu-infra-01",
          "cpu": { "architecture": "x86_64", "cpus": 1, "cores": 2, "threads": 4, "vendor": "GenuineIntel" },
          "memory": { "type": "RAM", "totalSize": 16 },
          "rootDisk": { "type": "SSD", "totalSize": 100 },
          "os": { "id": "ubuntu", "versionId": "22.04", "versionCodename": "jammy" },
          "gpuCards": [
            { "driverIndex": "0", "vendor": "NVIDIA", "model": "NVIDIA L4", "memoryTotalGB": 24 }
          ]
        }
      ]
    }
  }' | jq '.data[] | {status, description, nodeGroup: .targetInfra.nodeGroups[0], spec: .targetSpecList[0] | {id, cspSpecName, vCPU, memoryGiB, acceleratorModel, acceleratorCount, acceleratorMemoryGB}, image: .targetOsImageList[0] | {id, cspImageName, osDistribution}}'
```

### 4.2 `AWS (T4 16GB, Ubuntu 22.04) -> NCP (kr)` 추천 확인

```bash
curl -s -u default:default -X POST http://localhost:8056/beetle/recommendation/infra \
  -H "Content-Type: application/json" \
  -d '{
    "desiredCspAndRegionPair": { "csp": "ncp", "region": "kr" },
    "onpremiseInfraModel": {
      "nodes": [
        {
          "machineId": "lexai-src-compose-aws",
          "cpu": { "architecture": "x86_64", "cpus": 1, "cores": 2, "threads": 4, "vendor": "GenuineIntel" },
          "memory": { "type": "RAM", "totalSize": 16 },
          "rootDisk": { "type": "SSD", "totalSize": 100 },
          "os": { "id": "ubuntu", "versionId": "22.04", "versionCodename": "jammy" },
          "gpuCards": [
            { "driverIndex": "0", "vendor": "NVIDIA", "model": "Tesla T4", "memoryTotalGB": 16 }
          ]
        }
      ]
    }
  }' | jq '.data[] | {status, description, nodeGroup: .targetInfra.nodeGroups[0], spec: .targetSpecList[0] | {id, cspSpecName, vCPU, memoryGiB, acceleratorModel, acceleratorCount, acceleratorMemoryGB}, image: .targetOsImageList[0] | {id, cspImageName, osDistribution}}'
```
