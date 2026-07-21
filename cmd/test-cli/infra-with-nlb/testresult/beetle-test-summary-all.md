# CM-Beetle integration test summary (with NLB)

> [!IMPORTANT]
> This document provides an overall summary of automated integration test results for all provider-region pairs.

## Execution details

- **Test Date**: July 21, 2026
- **Start Time**: 13:13:55 KST
- **End Time**: 13:21:25 KST
- **Total Execution Duration**: 8m7s
- **CM-Beetle Version**: v0.5.5+ (726f2e8)
- **imdl Version**: v0.1.10+ (726f2e8)
- **CB-Tumblebug Version**: v0.12.25
- **CB-Spider Version**: v0.12.35
- **CB-MapUI Version**: v0.12.50

## High-level test status

| Metric | Count | Description |
|--------|-------|-------------|
| **Total CSP Pairs** | **10** | Number of unique CSP-Region configurations evaluated |
| Passed CSP Pairs | 1 | Pairs where all test steps succeeded |
| Failed CSP Pairs | 0 | Pairs where at least one test step failed |
| Skipped CSP Pairs | 9 | Pairs that were disabled in config |
| **Total Test Steps** | **130** | Total individual endpoint tests triggered |
| Passed Steps | 14 | Individual tests that succeeded |
| Failed Steps | 0 | Individual tests that failed |

## Provider-specific summary

| Provider-Region | Status | Duration | Steps Passed | Details |
|-----------------|--------|----------|--------------|---------|
| **AWS-Seoul** | ✅ **PASS** | 8m2s | 14 / 14 | [View Report](beetle-test-results-aws.md) |

