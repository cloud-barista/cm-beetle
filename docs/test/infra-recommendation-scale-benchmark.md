# 컴퓨팅 인프라 추천 엔진 대규모 스케일 벤치마크 (`1,000 -> 10,000 Nodes`)

- **작성일**: 2026-11-17
- **대상 엔드포인트**: `POST /beetle/recommendation/infra?desiredCsp=aws&desiredRegion=ap-northeast-2&limit=3&minMatchRate=85.0` (`Prefer: respond-async`)
- **대상 아키텍처**:
  - 요청 범위 HTTP 조회 캐시 단일화 ([`request-cache.go`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/client/tumblebug/request-cache.go), [`client-transport.go`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/client/tumblebug/client-transport.go), [`request-id.go`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/api/rest/middlewares/request-id.go))
  - 비동기 실행(`Prefer: respond-async` -> `202 Accepted`) 및 `GET /beetle/request/{reqId}` 폴링 (`ui/`)
- **CPU / 환경**: `linux/amd64`, `Intel(R) Core(TM) Ultra 7 356H`, 실기동 `CM-Beetle` (`localhost:8056`) + `CB-Tumblebug` (`localhost:1323`, initialized)

---

## 1. 테스트 워크로드 및 측정 방법

- **합성 온프레미스 인프라 모델 (`1,000 -> 10,000` Nodes, `+1,000` 단계, 누적 `55,000` 노드)**:
  - **하드웨어 프로파일 (`12종` 순환)**: `vCPU` (`2, 4, 8`), `Memory` (`4, 8, 16, 32, 64 GB`), `RootDisk` (`30, 50, 100 GB`), `OS` (`Ubuntu 22.04.5 LTS`, `x86_64`)
  - **보안그룹 (`4종` 압축)**: `web` (`22, 80, 443`), `was` (`22, 8080, 8443`), `db` (`22, 3306, 5432`), `worker` (`22, 9092`) -> `4개` SG로 중복 제거
  - **서브넷 분할 (`240대`당 `/24` 1개)**: `1,000대`(`5개`)부터 `10,000대`(`42개`, `10.1.1.0/24 ~ 10.1.42.0/24`)까지 자동 분할
- **측정 방식**:
  - 각 단계마다 `Prefer: respond-async`로 요청을 제출해 `202 Accepted`를 수신하고, `GET /beetle/request/{reqId}?brief=true&withSubRequests=false`를 `250ms` 주기로 폴링해 완료(`Success`)를 확인한 뒤 `brief=false&withSubRequests=true`로 전체 응답(`responseData`)과 `subRequestSummary`를 수집했습니다.
  - **서버 처리 시간 (`srv`)**: `GET /beetle/request/{reqId}`의 `endTime - startTime`
  - **End-to-End 시간 (`e2e`)**: `POST` 제출부터 폴링 완료 및 전체 후보 응답 JSON 수신 완료까지의 클라이언트 총 소요 시간

---

## 2. 라이브 실측 결과 (`CM-Beetle` + `CB-Tumblebug` 연동, `1,000 -> 10,000` Nodes)

전 구간(`10단계`, 누적 추천 노드 `55,000대`, 누적 Tumblebug 하위 호출 `360회`)에서 **에러 0건(`err=0`)**으로 100% 완주했습니다.

| 노드 수 (`N`) | 요청 / 응답 (`MB`) | E2E 소요 시간 | 서버 처리 시간 | 처리량 (`e2e` 기준) | Tumblebug 하위 호출 (`성공/전체/실패`) | 유효 추천 VM | 토폴로지 (`Subnets / SGs`) |
|---:|---:|---:|---:|---:|:---:|:---:|:---:|
| **1,000** | `0.88 / 2.44` | **`6.97s`** | `6.88s` | `143.4 대/초` | `36 / 36 / 0` | `1,000 / 1,000` | `5 / 4` |
| **2,000** | `1.77 / 4.67` | **`8.34s`** | `7.93s` | `239.9 대/초` | `36 / 36 / 0` | `2,000 / 2,000` | `9 / 4` |
| **3,000** | `2.65 / 6.91` | **`11.71s`** | `11.48s` | `256.1 대/초` | `36 / 36 / 0` | `3,000 / 3,000` | `13 / 4` |
| **4,000** | `3.54 / 9.14` | **`15.47s`** | `14.91s` | `258.5 대/초` | `36 / 36 / 0` | `4,000 / 4,000` | `17 / 4` |
| **5,000** | `4.42 / 11.38` | **`22.04s`** | `21.41s` | `226.9 대/초` | `36 / 36 / 0` | `5,000 / 5,000` | `21 / 4` |
| **6,000** | `5.31 / 13.61` | **`25.06s`** | `24.44s` | `239.4 대/초` | `36 / 36 / 0` | `6,000 / 6,000` | `25 / 4` |
| **7,000** | `6.19 / 15.85` | **`30.93s`** | `30.24s` | `226.3 대/초` | `36 / 36 / 0` | `7,000 / 7,000` | `30 / 4` |
| **8,000** | `7.08 / 18.09` | **`34.70s`** | `33.84s` | `230.6 대/초` | `36 / 36 / 0` | `8,000 / 8,000` | `34 / 4` |
| **9,000** | `7.96 / 20.32` | **`39.27s`** | `38.42s` | `229.2 대/초` | `36 / 36 / 0` | `9,000 / 9,000` | `38 / 4` |
| **10,000** | `8.85 / 22.56` | **`42.37s`** | `41.37s` | **`236.0 대/초`** | **`36 / 36 / 0`** | **`10,000 / 10,000`** | **`42 / 4`** |

### 하위 API (`SubRequest`) 호출 구성 (전 구간 동일)
- `POST /recommendSpec`: **`9`회**
- `POST /ns/system/resources/searchImage`: **`1`회**
- `POST /specImagePairReview`: **`26`회**
- 합계: **`36`회 (`POST` 기준, 10개 구간 모두 동일)** + `GET /provider/aws/region/ap-northeast-2` `1`회

---

## 3. 사양 다양성 스윕 (`N = 10,000` Nodes 고정, `K = 12 -> 500` 고유 사양)

동일 `10,000`대 규모에서 온프레미스 고유 사양(`HW × OS × Arch`) 수 `K`를 `12`종에서 `500`종까지 늘렸을 때의 캐시 효율, 엔드포인트별 호출당 응답 시간([`SubRequest.DurationMs`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/core/common/subrequest.go#L93)), 외부 API 대기 시간 대 로컬 처리 시간 비중을 측정했습니다.

- **프로파일 구성 원칙**:
  - `K = 12`: 위 Scale-Up 기준선과 동일(`12`종 x86_64 사양, 단일 OS Ubuntu 22.04)
  - `K = 50 -> 500`: `CB-Tumblebug` AWS 서울(`ap-northeast-2`) 카탈로그에서 매칭률이 확인된(`0.75 <= memory/vCPU <= 10.0`, 스펙 풀 내 `>= 90%` 매칭 가능 사양) 서로 다른 [`buildCpuDeploymentPlan`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/core/recommendation/resource-node-spec.go) 소수 구간(`x86_64` + `arm64` 약 80:20 혼합)과 기본 이미지 카탈로그(`includeBasicImageOnly=true`)에 등록된 `7`개 Ubuntu 버전(`20.04, 22.04, 24.04, 24.10, 25.04, 25.10, 26.04`) × `2`개 아키텍처를 교차 조합해 캐시 키가 실제로 분리되도록 구성했습니다.
- **측정 환경**: `CM-Beetle` (`BEETLE_LOGLEVEL=info`, `BEETLE_LOGWRITER=file`), `CB-Tumblebug` (`localhost:1323`), `2.0s` 폴링(각 단계 완료 후 `DELETE /request/{reqId}`로 `lkvstore` 누적 메모리 정리), `SubRequest` 지연 2초 이상 이상치 `0`건.

### 3-1. 고유 사양 수(`K`)별 소요 시간 및 시간 분해 (`N = 10,000` 고정)

| 고유 사양 `K` | 유효 `HW 범위` / `OS+Arch` | E2E 소요 시간 | 서버 처리 시간 | Tumblebug 대기 (`tb_wall`) | 로컬 매칭/직렬화 (`srv - tb_wall`) | 하위 호출 (`성공/전체/실패`) | 평균 매칭률 (`matched / total`) |
|---:|:---:|---:|---:|---:|---:|:---:|:---|
| **12** | `9 / 1` | **`46.35s`** | `44.17s` | `3.95s` (`8.9%`) | `40.21s` (`91.1%`) | `36 / 36 / 0` | `94.5%` (`6,667 / 10,000`) |
| **50** | `28 / 5` | **`44.56s`** | `42.15s` | `7.36s` (`17.5%`) | `34.78s` (`82.5%`) | `161 / 161 / 0` | `98.0%` (`8,800 / 10,000`) |
| **100** | `43 / 8` | **`53.35s`** | `51.61s` | `14.91s` (`28.9%`) | `36.71s` (`71.1%`) | `305 / 305 / 0` | `98.6%` (`9,200 / 10,000`) |
| **150** | `54 / 10` | **`55.72s`** | `54.39s` | `17.75s` (`32.6%`) | `36.64s` (`67.4%`) | `421 / 421 / 0` | `98.6%` (`9,266 / 10,000`) |
| **250** | `72 / 12` | **`61.53s`** | `59.35s` | `24.49s` (`41.3%`) | `34.86s` (`58.7%`) | `621 / 621 / 0` | `98.5%` (`9,240 / 10,000`) |
| **500** | `107 / 14` | **`72.52s`** | `71.80s` | **`38.56s` (`53.7%`)** | **`33.24s` (`46.3%`)** | **`1,024 / 1,024 / 0`** | `98.3%` (`9,320 / 10,000`) |

*(참고: `K = 12`에서 `K = 50`으로 넘어갈 때 서버 처리 시간이 `44.17s -> 42.15s`로 소폭 낮아진 것은 첫 실행의 Go 힙/Tumblebug 웜업 지연이 `K = 12`의 초기 로컬 시간 `40.21s` 및 첫 `/recommendSpec` 호출에 반영되었기 때문이며, 외부 Tumblebug 대기 시간 `tb_wall` 자체는 `3.95s -> 7.36s -> 14.91s -> 17.75s -> 24.49s -> 38.56s`로 단조 증가합니다.)*

### 3-2. 엔드포인트별 호출 수 및 호출당 응답 시간 (`SubRequest.DurationMs`)

| 고유 사양 `K` | `/recommendSpec` (횟수 / 평균 / `p50` / `p95`) | `/searchImage` (횟수 / 평균 / `p50` / `p95`) | `/specImagePairReview` (횟수 / 평균 / `p50` / `p95`) | 응답 크기 (`MB`) |
|---:|:---|:---|:---|---:|
| **12** | `9`회 / `243.3ms` / `188.0ms` / `279.0ms` | `1`회 / `264.0ms` / `264.0ms` / `264.0ms` | `26`회 / `57.0ms` / `36.0ms` / `118.0ms` | `22.55` |
| **50** | `28`회 / `45.9ms` / `55.0ms` / `82.0ms` | `5`회 / `233.0ms` / `119.0ms` / `176.0ms` | `128`회 / `37.8ms` / `30.0ms` / `89.0ms` | `22.98` |
| **100** | `43`회 / `32.1ms` / `17.0ms` / `83.0ms` | `8`회 / `606.1ms` / `623.0ms` / `661.0ms` | `254`회 / `33.6ms` / `29.0ms` / `36.0ms` | `23.29` |
| **150** | `54`회 / `17.8ms` / `14.0ms` / `38.0ms` | `10`회 / `588.0ms` / `601.0ms` / `661.0ms` | `357`회 / `30.0ms` / `28.0ms` / `34.0ms` | `23.39` |
| **250** | `72`회 / `16.1ms` / `12.0ms` / `37.0ms` | `12`회 / `557.2ms` / `558.0ms` / `646.0ms` | `537`회 / `30.4ms` / `29.0ms` / `35.0ms` | `23.54` |
| **500** | **`107`회 / `22.9ms` / `14.0ms` / `76.0ms`** | **`14`회 / `555.2ms` / `523.5ms` / `666.0ms`** | **`903`회 / `30.8ms` / `29.0ms` / `34.0ms`** | `24.01` |

- **AWS 서울 카탈로그 상한 및 수렴 확인**: `CB-Tumblebug`의 AWS(`ap-northeast-2`) 스펙 풀은 총 `762`개(`x86_64` `520` + `arm64` `242`), 고유 정수 `(vCPU, memoryGiB)` 쌍 `130`개(`95` + `35`)로 구성되며 [`calculateOptimalRange`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/core/recommendation/resource-node-spec.go) 기준 `87`개 소수 범위(`>= 90%` 매칭 커스텀 사양까지 확장해도 최대 약 `200`개 수준) 안에서 수렴합니다. 이 때문에 `K = 500` 조건에서도 실제 Tumblebug 호출은 산술 추정치(`3 × 500 = 1,500`회)보다 **31.7% 적은 `1,024`회**로 억제되었습니다.
- **병목 제거 검증 (`BEETLE_LOGLEVEL=info`, `BEETLE_LOGWRITER=file`)**: 기본 `debug` + `stdout/file` 동시 로깅 환경에서는 `10,000`대 1회 처리 중에도 수백 MB 디스크/콘솔 I/O가 누적될 수 있으나, `info`/`file` 로깅으로 분리한 상태에서는 누적 `60,000`대(`12 -> 500` 6연속) 실행 시 무응답 없이 에러 `0`건으로 완주했습니다.

---

## 4. 보조 오프라인 엔진 마이크로벤치마크 (`httptest.Server`, RTT `0ms`)

네트워크 왕복(`CB-Tumblebug`)을 제외한 `CM-Beetle` 순수 Go 엔진/메모리 할당 특성을 분리 관측하기 위해 [`BenchmarkRecommendInfraCandidates_Scale`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/core/recommendation/cache_test.go) (`-benchtime=1x`, `-benchmem`)를 함께 측정했습니다.

| 노드 수 (`N`) | 순수 엔진 시간 (`0ms` RTT) | 요청 캐시 `Hits / Misses` | 완료 시점 Heap / `VmRSS` (`MB`) | 누적 할당 (`B/op` / `allocs/op`) |
|---:|---:|---:|---:|---:|
| **1,000** | `357.7 ms` | `5,963 / 38` | `11.29 / 40.23` | `368.1 MB / 4.13M` |
| **5,000** | `1,458.6 ms` | `29,963 / 38` | `97.19 / 152.5` | `1,819.7 MB / 20.58M` |
| **10,000** | `2,966.5 ms` | `59,963 / 38` | `191.7 / 281.6` | `3,635.2 MB / 41.16M` |

---

## 5. 실측 기반 핵심 인사이트

1. **`1,000`대 `20분 13초` -> `6.88초` (`E2E 6.97초`), `4,000`대까지 `20초` 이내 달성**:
   - 기존([infra-recommendation-stress-test.md](file:///home/ubuntu/dev/cloud-barista/cm-beetle/docs/test/infra-recommendation-stress-test.md)) `1,000`대 `20분 13초`(`5,000`회 호출, `0.82 대/초`) 대비, 서버 기준 **약 `176배`** 단축되었습니다.
   - `10,000`대(`50,000`회 호출 규모)도 `K = 12` 조건에서는 **`36`회 실호출**(`99.93%` 호출 절감)만으로 **`42.37초`**(초당 `236.0`대)에 완료했습니다.
2. **외부 API 고정 비용 + 실환경 JSON 역직렬화 선형 구간(`~3.5~3.8ms/노드`)의 발견**:
   - `1,000`대(`6.88s`)와 `2,000`대(`7.93s`)의 차이는 약 **`1.05초`**에 불과해 저규모에서는 초기 Tumblebug 호출 `36`회 왕복(`~3~4초`)이 지배적입니다.
   - `10,000`대 기준 로컬 매칭 및 JSON 처리 시간은 사양 다양성(`K`)과 무관하게 **`~33~40초`(노드당 약 `3.3~4.0ms`)** 범위로 일정하게 유지됩니다.
3. **사양 다양성 `41.7`배 증가(`K = 12 -> 500`) 시 총 소요 시간은 `1.56`배(`46.35s -> 72.52s`)만 증가**:
   - `K`가 `12 -> 50 -> 100 -> 150 -> 250 -> 500`으로 늘어도 소수 구간 양자화(`500 -> 107` `/recommendSpec`, 약 `16~46 ms/call`), OS/Arch 통합(`500 -> 14` `/searchImage`, 약 `233~606 ms/call`), CSP 후보 카탈로그 수렴(`903` `/specImagePairReview`, 약 `30~38 ms/call`) 덕분에 Tumblebug 대기 시간은 `3.95s -> 38.56s`로 억제되며 전체 응답 시간은 **72.52초**(`137.9 대/초`)에서 안정적으로 방어됩니다.
4. **비동기(`Prefer: respond-async`) 오버헤드 최소화 확인**:
   - `10,000`대 기준 서버 완료 대비 클라이언트 `E2E` 차이(`202` 수락 + 폴링 + `22.5~24.0 MB` JSON 수신)는 **`0.72 ~ 2.37초`** 수준에 불과해 게이트웨이 타임아웃을 완전히 피하면서도 사용자 지연이 거의 없습니다.
5. **전역 TTL 배제 및 다음 병목 (`Next Bottleneck`)**:
   - `SoldOut`/`UnavailableZones`/`specImagePairReview`의 실시간 정합성을 해치지 않도록 **단일 요청 수명주기**에만 캐시를 제한했고, [`TumblebugTransport`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/client/tumblebug/client-transport.go#L51) + [`RequestIdAndDetailsIssuer`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/api/rest/middlewares/request-id.go#L24) 한 곳으로 단일화해 추천 외 도메인까지 일괄 가속했습니다.
   - `10,000`대에서는 후보 3개 응답 JSON만 **`22.55 ~ 24.01 MB`**에 달하고 [`GetRequest`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/core/common/request-manager.go#L142-L154) 폴링 시 [`convertToRequestDetails`](file:///home/ubuntu/dev/cloud-barista/cm-beetle/pkg/core/common/request-manager.go#L371-L393)가 요청 본문 JSON을 반복 재파싱하는 비용이 확인되었으므로, 차기 고도화 포인트는 **동일 스펙 노드 그룹 통합/응답 페이로드 경량화** 및 **`lkvstore` `brief` 폴링 경량화**입니다.
