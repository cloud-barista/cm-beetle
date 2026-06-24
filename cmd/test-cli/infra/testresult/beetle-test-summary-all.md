# CM-Beetle Test Suite Summary

> [!NOTE]
> This document summarizes all CM-Beetle integration tests from a single test run.

## Test Run Information

- **Date**: June 24, 2026
- **Start Time**: 18:52:24 KST
- **Total Duration**: 18m27.197s
- **Execution Mode**: parallel
- **CM-Beetle**: v0.5.2+ (1c7e6cd)
- **CB-Tumblebug**: v0.12.19
- **CB-Spider**: v0.12.32

## Overall Results

**Status**: ✅ All Passed

| Metric                 | Count |
| ---------------------- | ----- |
| Total CSP-Region Pairs | 10    |
| Passed Pairs           | 6     |
| Failed Pairs           | 0     |
| Skipped Pairs          | 4     |
| Total API Tests        | 90    |
| Passed API Tests       | 54    |
| Failed API Tests       | 0     |

## Per CSP-Region Pair Results

| #   | Display Name    | CSP     | Region          | Status  | Duration   | Passed | Failed | Skipped |
| --- | --------------- | ------- | --------------- | ------- | ---------- | ------ | ------ | ------- |
| 1   | AWS-Seoul       | AWS     | ap-northeast-2  | ✅ PASS | 3m25.067s  | 9/9    | 0/9    | 0/9     |
| 2   | Azure-Busan     | AZURE   | koreasouth      | ✅ PASS | 4m49.076s  | 9/9    | 0/9    | 0/9     |
| 3   | GCP-Seoul       | GCP     | asia-northeast3 | ✅ PASS | 18m17.204s | 9/9    | 0/9    | 0/9     |
| 4   | Alibaba-Seoul   | ALIBABA | ap-northeast-2  | ✅ PASS | 2m35.207s  | 9/9    | 0/9    | 0/9     |
| 5   | IBMCloud-Sydney | IBM     | au-syd          | ✅ PASS | 6m35.558s  | 9/9    | 0/9    | 0/9     |
| 6   | NCP-Seoul       | NCP     | kr              | ✅ PASS | 13m54.87s  | 9/9    | 0/9    | 0/9     |

## Detailed Results Per CSP-Region Pair

### 1. AWS-Seoul (AWS, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-24 18:52:29 KST
- **Duration**: 3m25.067s
- **Namespace**: mig01

| Step | Endpoint / Description                                     | Status  | Duration  |
| ---- | ---------------------------------------------------------- | ------- | --------- |
| 1    | `POST /beetle/recommendation/infra`                        | ✅ PASS | 5.608s    |
| 2    | `POST /beetle/migration/ns/mig01/infra`                    | ✅ PASS | 42.359s   |
| 3    | `GET /beetle/migration/ns/mig01/infra`                     | ✅ PASS | 41ms      |
| 4    | `GET /beetle/migration/ns/mig01/infra?option=id`           | ✅ PASS | 6ms       |
| 5    | `GET /beetle/migration/ns/mig01/infra/{{infraId}}`         | ✅ PASS | 17ms      |
| 6    | Remote SSH Accessibility                                   | ✅ PASS | 0s        |
| 7    | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}`    | ✅ PASS | 5.313s    |
| 8    | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.345s    |
| 9    | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`      | ✅ PASS | 1m24.039s |

### 2. Azure-Busan (AZURE, koreasouth) — ✅ PASS

- **Start Time**: 2026-06-24 18:52:32 KST
- **Duration**: 4m49.076s
- **Namespace**: mig01

| Step | Endpoint / Description                                     | Status  | Duration  |
| ---- | ---------------------------------------------------------- | ------- | --------- |
| 1    | `POST /beetle/recommendation/infra`                        | ✅ PASS | 4.908s    |
| 2    | `POST /beetle/migration/ns/mig01/infra`                    | ✅ PASS | 1m36.241s |
| 3    | `GET /beetle/migration/ns/mig01/infra`                     | ✅ PASS | 76ms      |
| 4    | `GET /beetle/migration/ns/mig01/infra?option=id`           | ✅ PASS | 4ms       |
| 5    | `GET /beetle/migration/ns/mig01/infra/{{infraId}}`         | ✅ PASS | 12ms      |
| 6    | Remote SSH Accessibility                                   | ✅ PASS | 0s        |
| 7    | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}`    | ✅ PASS | 5.306s    |
| 8    | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.283s    |
| 9    | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`      | ✅ PASS | 1m50.696s |

### 3. GCP-Seoul (GCP, asia-northeast3) — ✅ PASS

- **Start Time**: 2026-06-24 18:52:34 KST
- **Duration**: 18m17.204s
- **Namespace**: mig01

| Step | Endpoint / Description                                     | Status  | Duration  |
| ---- | ---------------------------------------------------------- | ------- | --------- |
| 1    | `POST /beetle/recommendation/infra`                        | ✅ PASS | 12.755s   |
| 2    | `POST /beetle/migration/ns/mig01/infra`                    | ✅ PASS | 8m3.005s  |
| 3    | `GET /beetle/migration/ns/mig01/infra`                     | ✅ PASS | 77ms      |
| 4    | `GET /beetle/migration/ns/mig01/infra?option=id`           | ✅ PASS | 3ms       |
| 5    | `GET /beetle/migration/ns/mig01/infra/{{infraId}}`         | ✅ PASS | 29ms      |
| 6    | Remote SSH Accessibility                                   | ✅ PASS | 0s        |
| 7    | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}`    | ✅ PASS | 5.317s    |
| 8    | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.315s    |
| 9    | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`      | ✅ PASS | 8m37.854s |

### 4. Alibaba-Seoul (ALIBABA, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-24 18:52:35 KST
- **Duration**: 2m35.207s
- **Namespace**: mig01

| Step | Endpoint / Description                                     | Status  | Duration |
| ---- | ---------------------------------------------------------- | ------- | -------- |
| 1    | `POST /beetle/recommendation/infra`                        | ✅ PASS | 4.488s   |
| 2    | `POST /beetle/migration/ns/mig01/infra`                    | ✅ PASS | 51.704s  |
| 3    | `GET /beetle/migration/ns/mig01/infra`                     | ✅ PASS | 50ms     |
| 4    | `GET /beetle/migration/ns/mig01/infra?option=id`           | ✅ PASS | 6ms      |
| 5    | `GET /beetle/migration/ns/mig01/infra/{{infraId}}`         | ✅ PASS | 20ms     |
| 6    | Remote SSH Accessibility                                   | ✅ PASS | 0s       |
| 7    | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}`    | ✅ PASS | 5.32s    |
| 8    | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.294s   |
| 9    | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`      | ✅ PASS | 26.733s  |

### 5. IBMCloud-Sydney (IBM, au-syd) — ✅ PASS

- **Start Time**: 2026-06-24 18:52:36 KST
- **Duration**: 6m35.558s
- **Namespace**: mig01

| Step | Endpoint / Description                                     | Status  | Duration |
| ---- | ---------------------------------------------------------- | ------- | -------- |
| 1    | `POST /beetle/recommendation/infra`                        | ✅ PASS | 19.297s  |
| 2    | `POST /beetle/migration/ns/mig01/infra`                    | ✅ PASS | 2m45.66s |
| 3    | `GET /beetle/migration/ns/mig01/infra`                     | ✅ PASS | 46ms     |
| 4    | `GET /beetle/migration/ns/mig01/infra?option=id`           | ✅ PASS | 8ms      |
| 5    | `GET /beetle/migration/ns/mig01/infra/{{infraId}}`         | ✅ PASS | 30ms     |
| 6    | Remote SSH Accessibility                                   | ✅ PASS | 0s       |
| 7    | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}`    | ✅ PASS | 5.335s   |
| 8    | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.306s   |
| 9    | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`      | ✅ PASS | 1m51.29s |

### 6. NCP-Seoul (NCP, kr) — ✅ PASS

- **Start Time**: 2026-06-24 18:52:38 KST
- **Duration**: 13m54.87s
- **Namespace**: mig01

| Step | Endpoint / Description                                     | Status  | Duration  |
| ---- | ---------------------------------------------------------- | ------- | --------- |
| 1    | `POST /beetle/recommendation/infra`                        | ✅ PASS | 2m14.947s |
| 2    | `POST /beetle/migration/ns/mig01/infra`                    | ✅ PASS | 7m33.298s |
| 3    | `GET /beetle/migration/ns/mig01/infra`                     | ✅ PASS | 1.785s    |
| 4    | `GET /beetle/migration/ns/mig01/infra?option=id`           | ✅ PASS | 4ms       |
| 5    | `GET /beetle/migration/ns/mig01/infra/{{infraId}}`         | ✅ PASS | 20ms      |
| 6    | Remote SSH Accessibility                                   | ✅ PASS | 0s        |
| 7    | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}`    | ✅ PASS | 5.271s    |
| 8    | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.272s    |
| 9    | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}`      | ✅ PASS | 2m40.815s |

---

_Generated by CM-Beetle Test CLI on 2026-06-24 18:52:24 KST_
