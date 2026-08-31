# Plan: Test CLI for Infra-with-NLB

## Overview

Create `cmd/test-cli/infra-with-nlb/main.go` by adapting the existing `cmd/test-cli/infra/main.go`.
The NLB-aware test CLI covers the full recommendation → VM migration → NLB migration → cleanup flow.

---

## Directory Structure

```
cmd/test-cli/infra-with-nlb/
├── main.go                          # (to be created)
├── testconf/
│   ├── recommendation-request.json  # Source infra model with NLBs (already exists)
│   ├── test-config.yaml             # Test config (CSP/region pairs, beetle endpoint)
│   ├── auth-config.json             # Auth credentials (gitignored)
│   └── template-*.yaml / *.json    # Templates for auth and test config
└── testresult/
    ├── beetle-test-results-aws.md   # Per-CSP report (auto-generated)
    └── beetle-test-summary-all.md  # Overall summary (auto-generated)
```

---

## Key Differences from `infra` Test CLI

| Item | `infra` | `infra-with-nlb` |
|---|---|---|
| Recommendation endpoint | `POST /recommendation/infra` | `POST /recommendation/infraWithNlb` |
| Request type | `RecommendInfraRequest` | `RecommendInfraWithNlbRequest` |
| Input loader | `onpremiseInfraModel` + `nameSeed` | `sourceInfra` + `nameSeed` (same JSON shape, different field name) |
| Total tests | 9 | 13 |
| NLB migration steps | — | Tests 7–9, 12 |

---

## Request / Response Types

### Recommendation

```
POST /beetle/recommendation/infraWithNlb?desiredCsp={csp}&desiredRegion={region}&limit=2

Request body: controller.RecommendInfraWithNlbRequest
{
  "desiredCsp":    "aws",
  "desiredRegion": "ap-northeast-2",
  "sourceInfra":   { ... OnpremInfra with NLBs ... }
}

Response: model.ApiResponse[[]cloudmodel.RecommendedInfra]
  RecommendedInfra.TargetNlbList []NlbReq  ← passed to NLB migration
```

### NLB Migration

```
POST /beetle/migration/middleware/ns/{nsId}/infra/{infraId}/nlb
Body: cloudmodel.RecommendedNlb { TargetNlbList []NlbReq }
Response: model.ApiResponse[cloudmodel.MigratedNlbResult]
  MigratedNlbResult.NlbList []MigratedNlbInfo

GET  /beetle/migration/middleware/ns/{nsId}/infra/{infraId}/nlb
Response: model.ApiResponse[[]cloudmodel.MigratedNlbInfo]

GET  /beetle/migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}
Response: model.ApiResponse[cloudmodel.MigratedNlbInfo]

DELETE /beetle/migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}
Response: 204 No Content  (includes 15s settle wait inside the handler)
```

---

## Test Flow (13 Steps per CSP-Region Pair)

| # | Endpoint / Description | Key Notes |
|---|---|---|
| 1 | `POST /recommendation/infraWithNlb` | Use `limit=2`; pick candidate[0] for migration |
| 2 | `POST /migration/ns/{nsId}/infra` | Same as infra test; use `nameSeed` |
| 3 | `GET /migration/ns/{nsId}/infra` | List all infras |
| 4 | `GET /migration/ns/{nsId}/infra?option=id` | List infra IDs; identify `infraId` |
| 5 | `GET /migration/ns/{nsId}/infra/{infraId}` | Get specific infra details |
| 6 | Remote SSH Accessibility Check | Same SSH retry logic as infra test |
| 7 | `POST /migration/middleware/ns/{nsId}/infra/{infraId}/nlb` | Body: `{targetNlbList}` from step 1 response; capture `nlbId` list |
| 8 | `GET /migration/middleware/ns/{nsId}/infra/{infraId}/nlb` | List NLBs |
| 9 | `GET /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}` | Get first NLB |
| 10 | Target Infrastructure Summary | Same as infra test |
| 11 | Migration Report | Same as infra test |
| 12 | `DELETE /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}` | Cleanup NLBs; handler has 15s settle wait |
| 13 | `DELETE /migration/ns/{nsId}/infra/{infraId}` | Cleanup VM infra; retry x10 / 10s interval |

> **Cleanup order matters**: NLBs (test 12) must be deleted before VM infra (test 13) to avoid VNet DependencyViolation.
> Tests 12 and 13 always run regardless of earlier failures.

---

## CSPTestReport — Added Fields

```go
type CSPTestReport struct {
    // ... (same as infra) ...
    NlbRecommendationResponse *model.ApiResponse[[]cloudmodel.RecommendedInfra]
    MigratedNlbResult         *cloudmodel.MigratedNlbResult
    NlbListResponse           []cloudmodel.MigratedNlbInfo
    GetNlbResponse            *cloudmodel.MigratedNlbInfo
}
```

---

## Input File: `recommendation-request.json`

Same top-level shape as the infra test:
```json
{
  "nameSeed": "my",
  "sourceInfra": { ... OnpremInfra with nlbs[] ... }
}
```

The loader reads `sourceInfra` (not `onpremiseInfraModel`) and passes it as `RecommendInfraWithNlbRequest.SourceInfra`.

---

## Reused from `infra/main.go` (unchanged)

- `TestConfig`, `TestCase`, `AuthConfig`, `TestResults`, `TestSuite` structs
- `loadConfig`, `loadAuthConfig`
- `checkBeetleReadiness`
- `runListInfraTest`, `runListInfraIdsTest`, `runGetInfraTest`, `runDeleteInfraTest`
- `runRemoteCommandTest`, `testSSHConnectivity`
- `runSourceSummaryTest`, `runTargetSummaryTest`, `runMigrationReportTest`
- `animatedSleep`, `maskSensitiveInfo`, `printTestCaseBanner`, `printFinalSummary`
- `generateOverallSummaryMarkdown`, `buildOverallSummaryContent`
- `convertMapTo*` helpers
- `getBeetleVersion`, `getImdlVersion`, `getVersionFromDockerCompose`

## Modified from `infra/main.go`

- `loadSourceInfraModel`: reads `sourceInfra` field (was `onpremiseInfraModel`)
- `runRecommendationTest`: calls `/infraWithNlb`, uses `RecommendInfraWithNlbRequest`
- `runTestCase`: 13-step flow; adds NLB steps 7–9 and 12
- `TotalTests = 13`
- `generateMarkdownContent`: add NLB test case sections (7–9, 12); renumber 10–13

## New functions

- `runMigrateNlbsTest`: `POST .../nlb`
- `runListNlbsTest`: `GET .../nlb`
- `runGetNlbTest`: `GET .../nlb/{nlbId}`
- `runDeleteNlbTest`: `DELETE .../nlb/{nlbId}` (204 response)

---

## Markdown Report

Per-CSP report: `testresult/beetle-test-results-{csp}.md`
Overall summary: `testresult/beetle-test-summary-all.md`

Both follow the same structure as the infra test report, with NLB test sections added.
