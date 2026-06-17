# CM-Beetle Test Suite Summary

> [!NOTE]
> This document summarizes all CM-Beetle integration tests from a single test run.

## Test Run Information

- **Date**: June 17, 2026
- **Start Time**: 19:45:06 KST
- **Total Duration**: 21m37.52s
- **Execution Mode**: parallel
- **CM-Beetle**: v0.5.2+ (8c6611e)
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
| 1 | AWS-Seoul | AWS | ap-northeast-2 | ✅ PASS | 2m54.489s | 9/9 | 0/9 | 0/9 |
| 2 | Azure-Busan | AZURE | koreasouth | ✅ PASS | 4m55.214s | 9/9 | 0/9 | 0/9 |
| 3 | GCP-Seoul | GCP | asia-northeast3 | ✅ PASS | 21m27.564s | 9/9 | 0/9 | 0/9 |
| 4 | Alibaba-Seoul | ALIBABA | ap-northeast-2 | ✅ PASS | 2m54.307s | 9/9 | 0/9 | 0/9 |
| 5 | IBMCloud-Sydney | IBM | au-syd | ✅ PASS | 6m34.943s | 9/9 | 0/9 | 0/9 |
| 6 | NCP-Seoul | NCP | kr | ✅ PASS | 13m52.484s | 9/9 | 0/9 | 0/9 |

## Detailed Results Per CSP-Region Pair

### 1. AWS-Seoul (AWS, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-17 19:45:11 KST
- **Duration**: 2m54.489s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 4.035s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 48.927s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 44ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 155ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.277s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.31s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 48.37s |

### 2. Azure-Busan (AZURE, koreasouth) — ✅ PASS

- **Start Time**: 2026-06-17 19:45:13 KST
- **Duration**: 4m55.214s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 4.866s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m38.524s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 44ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 19ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 64ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.848s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.302s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 1m48.556s |

### 3. GCP-Seoul (GCP, asia-northeast3) — ✅ PASS

- **Start Time**: 2026-06-17 19:45:16 KST
- **Duration**: 21m27.564s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 12.107s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 10m2.63s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 35ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 5ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 19ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.306s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.351s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 9m50.47s |

### 4. Alibaba-Seoul (ALIBABA, ap-northeast-2) — ✅ PASS

- **Start Time**: 2026-06-17 19:45:15 KST
- **Duration**: 2m54.307s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 4.549s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 1m8.859s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 33ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 14ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.426s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.516s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 27.862s |

### 5. IBMCloud-Sydney (IBM, au-syd) — ✅ PASS

- **Start Time**: 2026-06-17 19:45:16 KST
- **Duration**: 6m34.943s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 17.372s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 3m4.209s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 94ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 5ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 20ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.307s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.321s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 1m31.844s |

### 6. NCP-Seoul (NCP, kr) — ✅ PASS

- **Start Time**: 2026-06-17 19:45:18 KST
- **Duration**: 13m52.484s
- **Namespace**: mig01

| Step | Endpoint / Description | Status | Duration |
|------|------------------------|--------|----------|
| 1 | `POST /beetle/recommendation/infra` | ✅ PASS | 2m25.728s |
| 2 | `POST /beetle/migration/ns/mig01/infra` | ✅ PASS | 7m49.03s |
| 3 | `GET /beetle/migration/ns/mig01/infra` | ✅ PASS | 35ms |
| 4 | `GET /beetle/migration/ns/mig01/infra?option=id` | ✅ PASS | 4ms |
| 5 | `GET /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 18ms |
| 6 | Remote SSH Accessibility | ✅ PASS | 0s |
| 7 | `GET /beetle/summary/target/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.293s |
| 8 | `POST /beetle/report/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 5.303s |
| 9 | `DELETE /beetle/migration/ns/mig01/infra/{{infraId}}` | ✅ PASS | 2m12.933s |

---

*Generated by CM-Beetle Test CLI on 2026-06-17 19:45:06 KST*
