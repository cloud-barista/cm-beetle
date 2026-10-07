# [수정 요청] NodeGroup 원천 머신 식별자 Label 키 표준화

> **적용 버전**: CM-Beetle 차기 릴리스

---

## 1. 수정 요청 사항

Beetle이 배포한 인프라의 Node Label을 참조할 때, 기존 단수/복수 Label 키(`sourceMachineId`, `sourceMachineIds`) 대신 단일 표준 키인 **`cm-source-machine-ids`**를 참조하도록 수정을 요청드립니다.

| 기존 Label 키                          | 신규 표준 Label 키          | 값 형식                                          |
| :------------------------------------- | :-------------------------- | :----------------------------------------------- |
| `sourceMachineId` / `sourceMachineIds` | **`cm-source-machine-ids`** | 쉼표 구분 머신 ID 목록 (단일 머신인 경우 1개 ID) |

---

## 2. 수정 요청 이유

1. **NodeGroup 구조적 특성 반영**: NodeGroup은 동종 사양의 Node $N$대로 구성되는 특성을 가지므로, 서버와 개별 노드 간의 엄격한 1:1 Label 매핑이 불가하여 해당 NodeGroup에 속한 소스 머신 ID 목록(`cm-source-machine-ids`)으로 통합 관리하기 위함입니다.
2. **단일 표준 적용**: 기존 단수형(`sourceMachineId`)과 복수형(`sourceMachineIds`)의 혼재로 인한 매핑 누락을 방지하고 컴포넌트 간 결합도를 단순화하기 위함입니다.

(참고 - Grasshopper는 Tumblebug이 관리하는 Label을 사용) 멀티클라우드 Label 규격 준수: GCP 등 주요 CSP에서 Label 키의 대문자 및 특수문자가 금지되어 있어 소문자 하이픈(`cm-source-machine-ids`) 단일 규격으로 통일했습니다.

---

## 3. 구현 방안 및 1:1 매칭 기준 (안)

### 1) Label 참조 코드 변경

```go
// 변경 전:
machineId := node.Label["sourceMachineId"]

// 변경 후:
machineIds := node.Label["cm-source-machine-ids"]
```

### 2) NodeGroup 내 복수 Node 1:1 매칭 가이드

NodeGroup은 동일 사양으로 여러 대의 Node를 일괄 생성하므로, 해당 NodeGroup의 Label에는 소스 서버 머신 ID 목록이 들어갑니다.

- **예시**: `{"cm-source-machine-ids": "srv-01,srv-02,srv-03"}`

Grasshopper가 소스 서버와 타깃 Node를 1:1로 매핑하여 데이터를 복제할 수 있도록 아래 매칭 방안을 안내드립니다:

- **순서 기반 인덱스 매칭 (기본 권장안)**:
  NodeGroup 내 생성된 Node 목록의 순서(`node[i]`)와 `cm-source-machine-ids`를 쉼표로 분리한 배열(`machineIds[i]`)을 인덱스 기준으로 1:1 매칭합니다.
  - `node[0]` $\longleftrightarrow$ `srv-01`
  - `node[1]` $\longleftrightarrow$ `srv-02`
  - `node[2]` $\longleftrightarrow$ `srv-03`

- **무상태 노드 풀 순차 할당**:
  웹 서버 등 동종 무상태 노드의 경우 순서에 구애받지 않고 유휴 Node에 순차적으로 워커 작업을 할당합니다.

> **추후 협의 사항**:
> 상태 저장 서버 등 특정 머신과의 엄격한 고정 매핑이 요구되거나, 개별 Node 레벨 단일 Label 주입 등 추가 지원이 필요한 경우 향후 인터페이스 협의를 통해 보완할 수 있습니다.

---

## 4. 기대 효과

1. **CSP Label 정합성 보장**: AWS, GCP, Azure 등 모든 대상 클라우드에서 에러 없이 Label이 안전하게 전파됩니다.
2. **SW 이관 자동화 신뢰성 향상**: Grasshopper가 원본 소스 머신을 단일 키로 명확히 식별하고 1:1 매칭을 수립하여 데이터 복제 파이프라인의 안정성이 향상됩니다.
