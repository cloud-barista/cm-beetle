# CM-Beetle integration test summary (with NLB)

> [!IMPORTANT]
> This document provides an overall summary of automated integration test results for all provider-region pairs.

## Execution details

- **Test Date**: July 1, 2026
- **Start Time**: 15:08:49 KST
- **End Time**: 15:45:03 KST
- **Total Execution Duration**: 39m13s
- **CM-Beetle Version**: 3950cc5
- **imdl Version**: unknown
- **CB-Tumblebug Version**: Unknown (Fallback to Latest)
- **CB-Spider Version**: Unknown (Fallback to Latest)
- **CB-MapUI Version**: Unknown (Fallback to Latest)

## High-level test status

| Metric | Count | Description |
|--------|-------|-------------|
| **Total CSP Pairs** | **10** | Number of unique CSP-Region configurations evaluated |
| Passed CSP Pairs | 4 | Pairs where all test steps succeeded |
| Failed CSP Pairs | 0 | Pairs where at least one test step failed |
| Skipped CSP Pairs | 6 | Pairs that were disabled in config |
| **Total Test Steps** | **130** | Total individual endpoint tests triggered |
| Passed Steps | 56 | Individual tests that succeeded |
| Failed Steps | 0 | Individual tests that failed |

## Provider-specific summary

| Provider-Region | Status | Duration | Steps Passed | Details |
|-----------------|--------|----------|--------------|---------|
| **AWS-Seoul** | ✅ **PASS** | 7m21s | 14 / 14 | [View Report](beetle-test-results-aws.md) |
| **Alibaba-Seoul** | ✅ **PASS** | 4m31s | 14 / 14 | [View Report](beetle-test-results-alibaba.md) |
| **Azure-Busan** | ✅ **PASS** | 11m44s | 14 / 14 | [View Report](beetle-test-results-azure.md) |
| **NCP-Seoul** | ✅ **PASS** | 15m30s | 14 / 14 | [View Report](beetle-test-results-ncp.md) |

