# CM-Beetle Test Suite Summary

> [!NOTE]
> This document summarizes all CM-Beetle integration tests from a single test run.

## Test Run Information

- **Date**: June 17, 2026
- **Start Time**: 18:49:20 KST
- **Total Duration**: 19m17.399s
- **Execution Mode**: parallel
- **CM-Beetle**: v0.5.2
- **CB-Tumblebug**: v0.12.15
- **CB-Spider**: v0.12.30

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
| 1 | AWS-Seoul | AWS | ap-northeast-2 | ✅ PASS | 2m49.381s | 9/9 | 0/9 | 0/9 |
| 2 | Azure-Busan | AZURE | koreasouth | ✅ PASS | 5m3.389s | 9/9 | 0/9 | 0/9 |
| 3 | GCP-Seoul | GCP | asia-northeast3 | ✅ PASS | 19m8.658s | 9/9 | 0/9 | 0/9 |
| 4 | Alibaba-Seoul | ALIBABA | ap-northeast-2 | ✅ PASS | 2m48.52s | 9/9 | 0/9 | 0/9 |
| 5 | IBMCloud-Sydney | IBM | au-syd | ✅ PASS | 7m24.91s | 9/9 | 0/9 | 0/9 |
| 6 | NCP-Seoul | NCP | kr | ✅ PASS | 13m53.858s | 9/9 | 0/9 | 0/9 |

## Detailed Results Per CSP-Region Pair

### 1. AWS-Seoul (AWS, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-17 18:49:23 KST
- **Duration**: 2m49.381s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 2.551s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 45.292s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 44ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 6ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 14ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.315s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.283s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 47.922s |

### 2. Azure-Busan (AZURE, koreasouth) — ✅ PASS

- **Start Time**: 2026-06-17 18:49:25 KST
- **Duration**: 5m3.389s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 5.404s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m45.036s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 285ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 16ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.281s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.309s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 1m55.171s |

### 3. GCP-Seoul (GCP, asia-northeast3) — ✅ PASS

- **Start Time**: 2026-06-17 18:49:27 KST
- **Duration**: 19m8.658s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 12.933s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 8m30.247s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 40ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 5ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 37ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.278s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.32s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 9m1.853s |

### 4. Alibaba-Seoul (ALIBABA, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-17 18:49:28 KST
- **Duration**: 2m48.52s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 4.243s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m4.707s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 36ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 16ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.27s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.272s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 27.235s |

### 5. IBMCloud-Sydney (IBM, au-syd) — ✅ PASS

- **Start Time**: 2026-06-17 18:49:29 KST
- **Duration**: 7m24.91s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 22.616s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 3m41.513s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 21ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 14ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.307s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.27s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 1m39.119s |

### 6. NCP-Seoul (NCP, kr) — ✅ PASS

- **Start Time**: 2026-06-17 18:49:32 KST
- **Duration**: 13m53.858s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 2m26.642s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 7m47.528s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 2.7s |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 19ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.402s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.398s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 2m11.983s |

---

*Generated by CM-Beetle Test CLI on 2026-06-17 18:49:20 KST*
