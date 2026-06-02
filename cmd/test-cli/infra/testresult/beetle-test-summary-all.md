# CM-Beetle Test Suite Summary

> [!NOTE]
> This document summarizes all CM-Beetle integration tests from a single test run.

## Test Run Information

- **Date**: June 2, 2026
- **Start Time**: 20:56:57 KST
- **Total Duration**: 18m14.241s
- **Execution Mode**: parallel
- **CM-Beetle**: imdl/v0.1.5+ (7e4bd21)
- **CB-Tumblebug**: v0.12.13
- **CB-Spider**: v0.12.26

## Overall Results

**Status**: ✅ All Passed

| Metric | Count |
|--------|-------|
| Total CSP-Region Pairs | 10 |
| Passed Pairs | 6 |
| Failed Pairs | 0 |
| Skipped Pairs | 4 |
| Total API Tests | 90 |
| Passed API Tests | 54 |
| Failed API Tests | 0 |

## Per CSP-Region Pair Results

| # | Display Name | CSP | Region | Status | Duration | Passed | Failed | Skipped |
|---|--------------|-----|--------|--------|----------|--------|--------|--------|
| 1 | AWS-Seoul | AWS | ap-northeast-2 | ✅ PASS | 3m17.579s | 9/9 | 0/9 | 0/9 |
| 2 | Azure-Busan | AZURE | koreasouth | ✅ PASS | 4m57.447s | 9/9 | 0/9 | 0/9 |
| 3 | GCP-Seoul | GCP | asia-northeast3 | ✅ PASS | 18m4.998s | 9/9 | 0/9 | 0/9 |
| 4 | Alibaba-Seoul | ALIBABA | ap-northeast-2 | ✅ PASS | 2m56.723s | 9/9 | 0/9 | 0/9 |
| 5 | IBMCloud-Sydney | IBM | au-syd | ✅ PASS | 6m8.262s | 9/9 | 0/9 | 0/9 |
| 6 | NCP-Seoul | NCP | kr | ✅ PASS | 13m20.277s | 9/9 | 0/9 | 0/9 |

## Detailed Results Per CSP-Region Pair

### 1. AWS-Seoul (AWS, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-02 20:57:02 KST
- **Duration**: 3m17.579s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 11.915s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m2.785s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 43ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 5ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 24ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.411s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.503s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 47.728s |

### 2. Azure-Busan (AZURE, koreasouth) — ✅ PASS

- **Start Time**: 2026-06-02 20:57:03 KST
- **Duration**: 4m57.447s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 13.084s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m36.82s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 126ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 16ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 28ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.584s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.604s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 1m47.036s |

### 3. GCP-Seoul (GCP, asia-northeast3) — ✅ PASS

- **Start Time**: 2026-06-02 20:57:06 KST
- **Duration**: 18m4.998s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 16.118s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 7m45.076s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 56ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 3ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 17ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.398s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.324s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 8m39.645s |

### 4. Alibaba-Seoul (ALIBABA, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-02 20:57:07 KST
- **Duration**: 2m56.723s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 7.581s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m9.011s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 42ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 9ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 20ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.617s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.392s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 27.428s |

### 5. IBMCloud-Sydney (IBM, au-syd) — ✅ PASS

- **Start Time**: 2026-06-02 20:57:10 KST
- **Duration**: 6m8.262s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 15.53s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 2m54.786s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 38ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 7ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 26ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.345s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.3s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 1m22.808s |

### 6. NCP-Seoul (NCP, kr) — ✅ PASS

- **Start Time**: 2026-06-02 20:57:11 KST
- **Duration**: 13m20.277s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 2m15.919s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 7m24.368s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 1.069s |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 7ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 26ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.954s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.433s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 2m14.464s |

---

*Generated by CM-Beetle Test CLI on 2026-06-02 20:56:57 KST*
