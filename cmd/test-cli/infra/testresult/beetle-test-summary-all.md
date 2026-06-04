# CM-Beetle Test Suite Summary

> [!NOTE]
> This document summarizes all CM-Beetle integration tests from a single test run.

## Test Run Information

- **Date**: June 4, 2026
- **Start Time**: 15:43:37 KST
- **Total Duration**: 18m44.694s
- **Execution Mode**: parallel
- **CM-Beetle**: v0.5.0+ (2aeaf75)
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
| 1 | AWS-Seoul | AWS | ap-northeast-2 | ✅ PASS | 3m34.669s | 9/9 | 0/9 | 0/9 |
| 2 | Azure-Busan | AZURE | koreasouth | ✅ PASS | 6m18.1s | 9/9 | 0/9 | 0/9 |
| 3 | GCP-Seoul | GCP | asia-northeast3 | ✅ PASS | 18m36.066s | 9/9 | 0/9 | 0/9 |
| 4 | Alibaba-Seoul | ALIBABA | ap-northeast-2 | ✅ PASS | 2m55.272s | 9/9 | 0/9 | 0/9 |
| 5 | IBMCloud-Sydney | IBM | au-syd | ✅ PASS | 6m54.197s | 9/9 | 0/9 | 0/9 |
| 6 | NCP-Seoul | NCP | kr | ✅ PASS | 13m45.47s | 9/9 | 0/9 | 0/9 |

## Detailed Results Per CSP-Region Pair

### 1. AWS-Seoul (AWS, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-04 15:43:42 KST
- **Duration**: 3m34.669s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 13.617s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m10.975s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 42ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 5ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 19ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.32s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.306s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 55.315s |

### 2. Azure-Busan (AZURE, koreasouth) — ✅ PASS

- **Start Time**: 2026-06-04 15:43:44 KST
- **Duration**: 6m18.1s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 14.21s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 2m23.979s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 508ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 10ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 22ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.301s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.284s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 2m11.58s |

### 3. GCP-Seoul (GCP, asia-northeast3) — ✅ PASS

- **Start Time**: 2026-06-04 15:43:45 KST
- **Duration**: 18m36.066s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 16.421s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 8m34.262s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 74ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 6ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 91ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.712s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.646s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 8m20.604s |

### 4. Alibaba-Seoul (ALIBABA, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-04 15:43:48 KST
- **Duration**: 2m55.272s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 7.878s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m6.668s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 48ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 5ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 21ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.293s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.297s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 28.391s |

### 5. IBMCloud-Sydney (IBM, au-syd) — ✅ PASS

- **Start Time**: 2026-06-04 15:43:50 KST
- **Duration**: 6m54.197s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 23.241s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 3m8.867s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 37ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 46ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.327s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.303s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 1m43.539s |

### 6. NCP-Seoul (NCP, kr) — ✅ PASS

- **Start Time**: 2026-06-04 15:43:52 KST
- **Duration**: 13m45.47s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 2m19.522s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 7m44.938s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 2.845s |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 3ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 23ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.418s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.545s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 2m12.471s |

---

*Generated by CM-Beetle Test CLI on 2026-06-04 15:43:37 KST*
