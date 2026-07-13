# CM-Beetle integration test summary (with NLB)

> [!IMPORTANT]
> This document provides an overall summary of automated integration test results for all provider-region pairs.

## Execution details

- **Test Date**: July 13, 2026
- **Start Time**: 20:03:25 KST
- **End Time**: 20:03:36 KST
- **Total Execution Duration**: 10s
- **CM-Beetle Version**: b418c24
- **imdl Version**: unknown
- **CB-Tumblebug Version**: Unknown (Fallback to Latest)
- **CB-Spider Version**: Unknown (Fallback to Latest)
- **CB-MapUI Version**: Unknown (Fallback to Latest)

## High-level test status

| Metric | Count | Description |
|--------|-------|-------------|
| **Total CSP Pairs** | **10** | Number of unique CSP-Region configurations evaluated |
| Passed CSP Pairs | 0 | Pairs where all test steps succeeded |
| Failed CSP Pairs | 1 | Pairs where at least one test step failed |
| Skipped CSP Pairs | 9 | Pairs that were disabled in config |
| **Total Test Steps** | **130** | Total individual endpoint tests triggered |
| Passed Steps | 0 | Individual tests that succeeded |
| Failed Steps | 1 | Individual tests that failed |
| Skipped Steps | 13 | Tests skipped due to pre-requisite step failure |

## Provider-specific summary

| Provider-Region | Status | Duration | Steps Passed | Details |
|-----------------|--------|----------|--------------|---------|
| **Alibaba-Singapore** | ❌ **FAIL** | 5s | 0 / 14 | [View Report](beetle-test-results-alibaba.md) |

