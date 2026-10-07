# [제안] Damselfly 모델 메타데이터 기반 계보(Lineage) 관리

> **관련 패키지**: `github.com/cloud-barista/cm-beetle/imdl/lineage`

---

## 1. 제안 사항

CM-Damselfly의 모델 저장 및 응답 메타데이터(Envelope)에 모델 간의 파생 계보(`lineage`) 객체를 추가하고, 3대 모델 타입(`source`, `planning`, `target`)의 파생 관계를 관리하도록 확장을 제안합니다.

---

## 2. 제안 이유

1. **소스-to-목표(Source-to-Target) 추적성 확보**: 온프레미스 원천 탐색부터 마이그레이션 계획, 타깃 클라우드 추천, 최종 배포까지 이어지는 모델 파생 관계를 단절 없이 추적하고 포털에 도식화하기 위함입니다.
2. **소스/목표 정보인 모델의 순수성 보존**: `OnpremInfra`, `RecommendedInfra` 등은 마이그레이션에 필요한 정보를 담은 모델입니다. 순수성을 유지하기 위해 모델 메타데이터 레이어에서 계보를 독립 관리하여 소스-목표간 추척성을 제공하는 것이 바람직합니다.
3. **자원 계층 및 생애주기 분리**:
   - **마이그레이션 계보 (Migration Lineage) (Damselfly)**: 모델 간의 부모-자식 파생 관계 관리
   - **배포 인프라 Label (Beetle)**: 실제 클라우드 프로비저닝 시점의 배포 일시(`cm-migrated-at`) 주입
   - **노드 Label (Beetle)**: 개별 원본 서버와 타깃 노드 간의 매핑(`cm-source-machine-ids`) 주입

---

## 3. 구현 방안 (안)

### 1) Damselfly Model Metadata/Envelope 확장

모델 메타데이터 객체에 `lineage` 필드를 추가합니다:

```json
{
  "id": "tgt-mdl-201",
  "name": "Target Recommendation",
  "modelType": "target",
  "lineage": {
    "sourceGroupId": "sg-ecommerce-prod",
    "sourceModelId": "src-mdl-101",
    "planningModelId": "plan-mdl-101",
    "targetModelId": "tgt-mdl-201"
  },
  "cloudInfraModel": { ... }
}
```

### 2) 공통 계보 패키지 활용 (`imdl/lineage`)

CM-Beetle에서 외부 의존성 없는 순수 Go 패키지를 제공하며, 모델 식별자만 포함합니다:

```go
type MigrationLineage struct {
	SourceGroupId   string `json:"sourceGroupId,omitempty"`
	SourceModelId   string `json:"sourceModelId,omitempty"`
	PlanningModelId string `json:"planningModelId,omitempty"`
	TargetModelId   string `json:"targetModelId,omitempty"`
}
```

### 3) 단계별 계보 전파 흐름

- **소스 모델 (`source`)**: 원천 탐색 정보 등록 시 `sourceGroupId`, `sourceModelId` 설정
- **계획 모델 (`planning`)**: 포털에서 수정/저장 시 부모 모델 ID (`sourceModelId`) 및 신규 `planningModelId` 설정
- **목표 모델 (`target`)**: Beetle 추천 확정 시 부모 모델 ID (`planningModelId`) 및 신규 `targetModelId` 설정

---

## 4. 기대 효과

1. **관심사의 분리**: 소스/목표 모델과 마이그레이션 관리를 위한 모델 메타데이터가 명확히 분리되어 모델 유지보수성이 향상됩니다.
2. **포털 현황 도식화 지원**: 포털(Butterfly)에서 모델을 조회해서, 원천 온프레미스부터 최종 추천까지의 흐름을 파이프라인으로 시각화할 수 있습니다.
3. **SW 이관 연계 자동화**: Grasshopper 등 후속 서브시스템이 마이그레이션 계보(Migration Lineage)와 Node Label을 기반으로 대상 소스 머신을 자동 식별할 수 있습니다.
