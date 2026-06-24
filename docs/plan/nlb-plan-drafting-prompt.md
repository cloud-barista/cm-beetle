사용자의 소스 컴퓨팅 환경의 NLB를 목표 클라우드 환경으로 마이그레이션하기 위한 계획을 작성합니다.

---

## A. 자원 정보

1. 소스 컴퓨팅 환경에서 추출되는 NLB 정보는 다음 링크를 참조 및 분석 바랍니다. Raw 데이터입니다.
   https://github.com/cloud-barista/cm-honeybee/discussions/55

2. 이 중 마이그레이션에 필요한 항목(properties)만 선별하여 소스 모델(Schema, struct)을 도출합니다.
   기존 모델 패턴은 `imdl` 디렉토리를 참조 및 분석 바랍니다.

3. 2번에서 도출한 모델은 추후 cm-honeybee 프로젝트와 협의하여 반영합니다. 단순 참고하시기 바랍니다.

4. 목표 클라우드 NLB를 생성하기 위해 Tumblebug API를 활용합니다. 아래 Swagger API 문서를 참조 및 분석 바랍니다.
   https://raw.githubusercontent.com/cloud-barista/cb-tumblebug/v0.12.15/src/interface/rest/docs/swagger.yaml

5. 목표 클라우드 NLB 모델은 Tumblebug에서 NLB 생성에 필요한 항목 + @(컨텍스트 정보)로 구성합니다.

---

## B. 필수 설계 결정 사항

다음 사항은 초안 작성 전에 반드시 명시되어야 합니다. 정보가 없을 경우 각 항목에 대해 분석 후 합리적인 기본값을 제안하고 명시합니다.

### B1. 기존 API와의 관계

- 기존 `/recommendation/infra`, `/migration/infra`에 영향을 주는지 여부를 분석합니다.
- 신규 통합 엔드포인트가 필요한지, 독립 엔드포인트로 충분한지 결정하고 근거를 명시합니다.
  (기존 엔드포인트 목록 및 명명 패턴은 `pkg/api/rest/server.go`를 참조합니다.)
- 기존 API의 하위 호환성 유지 여부를 명시합니다.

### B2. 구현 범위 경계

- 이번 계획에 포함되는 기능을 명시합니다.
- 명시적으로 제외하거나 보류하는 기능을 명시합니다.

### B3. 마이그레이션 순서 및 전제 조건

- 이 자원을 생성하기 전에 먼저 존재해야 하는 자원 목록(Tumblebug 기준)을 명시합니다.
- 다른 마이그레이션 단계와의 수행 순서를 명시합니다.

### B4. 소스 데이터 정규화 책임 분담

- cm-honeybee: raw data 수집 및 정규화 담당 범위
- cm-beetle: 정규화된 모델 소비 담당 범위
- 협의가 필요한 항목과 단순 참고 항목을 구분합니다.

---

## C. 기능 정의

6. 추천(Recommendation)은 소스에서 추출·정제한 정보를 입력받아, 목표 클라우드 자원 생성에 필요한 항목으로 정리·매핑하는 기능/API입니다. 목표 클라우드 자원 모델이 반환됩니다.

7. 마이그레이션(Migration)은 추천 결과로 반환된 목표 클라우드 자원 모델을 입력받아 자원을 생성·설정하는 기능/API입니다. 생성·설정된 자원 정보가 반환됩니다.

---

## D. 계획 문서 구성 요건

계획 문서는 아래 섹션을 포함합니다.

| 섹션 | 내용 |
|------|------|
| 개요 | 전체 흐름 ASCII 다이어그램 (소스 환경 → CM-Beetle → 목표 클라우드) |
| 소스 raw data 분석 | 수집 데이터 예시, 마이그레이션 필요 항목 선별 표 |
| 소스 모델 | Go struct 정의 |
| 목표 모델 | Go struct 정의 (Tumblebug 필드와 대응) |
| 소스 → 목표 매핑 | 필드별 변환 규칙 표 |
| 추천 API | 엔드포인트, 입출력 모델, 처리 흐름 (Mermaid 또는 ASCII) |
| 마이그레이션 API | 엔드포인트, 입출력 모델, 처리 흐름 (Mermaid 또는 ASCII) |
| 구현 단계 | 체크리스트 형식 |
| 제약 사항 | 표 형식 |

**작성 기준:**
- 모델은 Go struct 코드 스타일로 작성합니다 (json 태그, validate 태그 포함).
- JSON 예시는 실제 데이터 기반으로 작성합니다 (placeholder string 지양).
- 처리 흐름은 Mermaid 또는 ASCII 다이어그램으로 표현합니다.

---

## E. 출력물

위 내용을 바탕으로 `docs/plan/nlb-plan-for-recommendation-and-migration.md`를 작성합니다.
