# 가용 영역(Zone) 인지형 추천 및 마이그레이션 후속 고도화 계획서 (보류 과제 모음)

- **작성 및 개정일**: 2026-11-12
- **선행 기준 문서**: [zone-and-availability-aware-recommendation-plan.md](./zone-and-availability-aware-recommendation-plan.md) (Phase 1: 정적 가용성 필터링 및 단일 서브넷 유효 Zone 정합성)
- **상태**: 보류 (Phase 1 정적 가용성 정합성 구현 및 실증 완료 후 하부 멀티존 제약 해소와 연계하여 순차 추진)

---

## 1. 개요

본 문서는 [zone-and-availability-aware-recommendation-plan.md](./zone-and-availability-aware-recommendation-plan.md)에서 **Phase 1 정적 가용성 정합성(비가용 스펙 필터링, 단일 기본 서브넷 유효 `AvailableZone` 지정, 사전 검증, `AIMX` CLI 정합화)** 만 우선 반영하기로 확정함에 따라, 동적 서브넷 추가/전환 및 확장 영역으로 보류한 **3가지 후속 과제**를 정의합니다.

1. **후속 과제 1**: 동적 멀티 서브넷 배치 및 `CB-Tumblebug` 정식 API(`retryFailedNodes`) 기반 실시간 `ZoneCapacity` Failover
2. **후속 과제 2**: `Beetle UX Lab` (`2. Target Cloud Optimizer`) 가용 영역 및 가용성 상태 UI 고도화
3. **후속 과제 3**: `KT Cloud` 및 `OpenStack` 활성화 (10종 CSP 전면 확대)

---

## 2. 후속 과제별 상세 설계

### 2.1 후속 과제 1: 동적 멀티 서브넷 배치 및 `CB-Tumblebug` 정식 API 기반 실시간 `ZoneCapacity` Failover

하부 `CB-Spider` 및 `CB-Tumblebug`의 CSP별 멀티존 서브넷 제약(`IBM` VPC Address Prefix 정책, `NCP` 멀티존 서브넷 정책 등)이 일관되게 정리된 이후, 임시 예외 분기 없이 정석 아키텍처로 동적 멀티존 지원과 실시간 용량 부족 자동 복구를 구현합니다.

1. **이종 NodeGroup 간 배타적 AZ 대응 멀티 서브넷 구성 (`pkg/core/recommendation`)**:
   - 복합 워크로드(예: 일반 CPU NodeGroup + 특정 AZ 전용 GPU NodeGroup)에서 모든 NodeGroup을 동시에 수용하는 단일 공통 AZ가 존재하지 않을 때, 필요한 AZ에 한해서만 비중첩 CIDR 서브넷을 구성하고 각 NodeGroup의 `SubnetId`를 매핑합니다 (`common.ApplyNameSeed` 및 `common.ValidateComposedNames` 무결성 유지).
2. **`CB-Tumblebug v0.13.8+` 정식 노드 재시도 API 위임형 Zone Failover (`pkg/client/tumblebug`, `pkg/core/migration`)**:
   - `CM-Beetle`이 임의로 `DeleteInfra` 후 전체 재생성하거나 수동으로 CIDR을 분할하지 않고, `CB-Tumblebug`에 내장된 공식 API를 호출하여 계층 간 책임을 분리합니다:
     - `POST /tumblebug/ns/{nsId}/infra/{infraId}/retryFailedNodesReview` (`ZoneCapability.Shiftable` 동적 확인 및 `Failure.Class == "ZoneCapacity"` 판별)
     - `POST /tumblebug/ns/{nsId}/infra/{infraId}/retryFailedNodes` (`targets[].zone`에 스펙의 차순위 `AvailableZone` 전달, `KeepFailedNodes: false`)
   - `CB-Tumblebug` 내부의 `resolveZoneOverride`가 기존 VNet 내 대체 Zone 서브넷 유무를 확인해 온디맨드로 서브넷을 확장(`CreateSubnet`)하고 실패 노드만 교체하므로, **단일 노드 전체 실패와 복수 NodeGroup 부분 실패(정상 `Running` VM 보존) 모두 단일 메커니즘으로 처리**됩니다.
   - `AccountQuota`, `Auth`, `ImageInvalid`, `SpecInvalid` 등 비재시도성 오류는 즉시 중단(`Fail Fast`)하고, 요청된 동일 스펙과 동일 리전을 유지(`No Fallback` 원칙)합니다.

---

### 2.2 후속 과제 2: `Beetle UX Lab` 가용 영역(`AvailableZones`) 및 가용성 상태 UI 고도화

백엔드가 `SubnetInfoList[0].Zone`을 자동 주입하는 것을 넘어, 사용자가 웹 UI(`Beetle UX Lab`)에서 각 후보 스펙의 가용 영역 정보를 직관적으로 확인하고 제어할 수 있도록 개선합니다 (`docs/test/aimx-kotaemon-rag-migration-test-guide.md` 시나리오 B 라이브 시연 편의성 강화).

- **대상 파일**: `ui/src/components/design/CloudInfraOptimizer.tsx`
- **주요 내용**:
  1. **추천 스펙 카드 가용성 배지 표시**:
     - 후보 모델의 `TargetSpecList`에 포함된 `availabilityStatus`(`Available`, `Unverified`) 및 `availableZones` 목록을 `Node Spec` 정보 영역에 배지로 시각화합니다.
  2. **Zonal vs Regional 서브넷 편집 UX 분기**:
     - Zonal Subnet CSP(`AWS`, `Alibaba`, `Tencent`, `IBM`, `NCP`, `KT`) 선택 시 해당 스펙의 `availableZones` 목록을 드롭다운으로 제공하여 미지원 AZ 오입력을 원천 방지합니다.
     - Regional Subnet CSP(`Azure`, `GCP`, `NHN`, `OpenStack`) 선택 시 Zone 입력란을 비활성화하거나 `Regional Subnet (Auto-assigned by CSP)` 안내 문구를 표시합니다.
  3. **디자인 시스템 준수**:
     - `.agents/AGENTS.md`의 브랜드 컬러(`Emerald` / `Teal`), 폰트 위계(`font-normal text-text-muted` / `font-extrabold text-text-main`, 최소 `text-sm`), 프로젝트 표준 용어(`Nodes`, `Node Groups`, `2. Target Cloud Optimizer`) 및 `npm run lint` 규칙을 엄격히 준수합니다.

---

### 2.3 후속 과제 3: `KT Cloud` 및 `OpenStack` 활성화 (10종 CSP 완결)

현재 `pkg/core/recommendation/infra.go`의 `supportedProviders`에서 비활성화(`false`)되어 있는 `KT Cloud`와 `OpenStack`을 정식 활성화하여 10종 멀티클라우드 지원 체계를 완성합니다.

1. **`KT Cloud` (`kt` / `ktvpc`) 활성화 및 제약 정합성 검증**:
   - **대상 파일**: `pkg/core/recommendation/infra.go`, `pkg/core/recommendation/assets/networkinfo.yaml`, `pkg/compat/compat-kt.go`
   - **주요 내용**:
     - `supportedProviders`에서 `csp.KT: true`로 전환합니다.
     - `networkinfo.yaml` 내 주석(`# TBD`) 처리된 `kt` 네트워크 규칙(RFC 1918 대역, VNet prefix `/16`~`/24`, Subnet prefix `/24`~`/28`, 예약 IP 5개)을 활성화합니다.
     - `KT Cloud` VPC의 Zonal 서브넷 `Zone` 필수 매핑, 루트 디스크 크기 제약(`50GB`), 방화벽 규칙 호환성을 실환경(`cmd/test-cli/infra`)에서 검증합니다.
2. **`OpenStack` 활성화 및 프라이빗 클라우드 카탈로그 예외 처리**:
   - **대상 파일**: `pkg/core/recommendation/infra.go`, `resource-node-spec.go`, `resource-node-spec-gpu.go`, `resource-k8s-node-spec.go`, `assets/networkinfo.yaml`, `pkg/compat/compat-openstack.go`
   - **주요 내용**:
     - `supportedProviders`에서 `csp.OpenStack: true`로 전환합니다.
     - 프라이빗 클라우드 특성상 퍼블릭 가격 단가(`CostPerHour`)가 `0` 또는 `-1`로 등록되거나 `AvailabilityStatus`가 `Unverified`로 유지되는 경우가 일반적이므로, `csp.ResolveCloudPlatform(provider) == csp.OpenStack`일 때 가격 미설정 스펙이 `filterPricedSpecs`에서 탈락하지 않도록 허용하고 스펙 적합도(vCPU / Memory / GPU) 기준으로 추천 순위를 산정합니다.
     - `networkinfo.yaml` 내 `openstack` 네트워크 제약 정의를 추가합니다.

---

## 3. 추진 일정 및 우선순위 제안

| 순서 | 과제명 | 선행 조건 | 주요 검증 방법 |
| :---: | :--- | :--- | :--- |
| **1순위** | 후속 과제 2: `Beetle UX Lab` 가용 영역 UI 고도화 | Phase 1 정적 가용성 정합성 병합 완료 | `npm --prefix ui run lint`, 시나리오 B 라이브 시연 리허설 |
| **2순위** | 후속 과제 1: 동적 멀티 서브넷 및 `retryFailedNodes` 기반 Zone Failover | 하부 멀티존 서브넷 정책 및 `retryFailedNodes` E2E 검증 | `httptest` 모의 검증 및 실 클라우드 용량 부족 전환 테스트 |
| **3순위** | 후속 과제 3: `KT Cloud` 및 `OpenStack` 활성화 | 테스트베드 커넥션 및 이미지 자산 확보 | `cmd/test-cli/infra`, 10종 CSP 추천 매트릭스 테스트 |
