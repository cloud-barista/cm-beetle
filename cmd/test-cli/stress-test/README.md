# Multi-CSP Progressive Stress Test & Zero-Residue Cleanup CLI

Executes progressive multi-CSP scale-up stress testing (`Plan` -> `Recommend` -> `Migrate` -> `2-Stage Auto-Cleanup`) using randomly generated `onpremiseInfraModel` workloads (`10 Nodes` to `100+ Nodes`), and outputs:
- `testresult/stress-test-data.json`: Checkpoint-persisted raw step and time-series data
- `testresult/stress-test-summary.md`: Round-by-round Markdown summary table
- `testresult/stress-test-dashboard.html`: Standalone interactive time-lapse replay dashboard (`1x`, `5x`, `15x`, `30x` speeds, dark/light theme, SVG scale-up & failure annotations)

## Scale-Up Progression Rule

- **Round 1 (`10 Nodes`)**: Tests all 8 CSPs (`AWS`, `Azure`, `GCP`, `Alibaba`, `IBM`, `NCP`, `Tencent`, `NHN`) at `10 Nodes`.
- **Round 2 (`20 Nodes`) ~ Round N**: Tests all surviving CSPs by increasing the target scale by `+10 Nodes` per round.
- **Zero-Residue Cleanup on Failure**: Whenever a CSP encounters a failure, its resources are immediately swept via 2-stage cleanup (`MCI` -> `SecurityGroup` -> `SSHKey` -> `VNet`) and the CSP is excluded from subsequent rounds.

## Usage

```bash
cd cmd/test-cli/stress-test

# 1. Run in simulation mode (+10 Nodes per round across surviving CSPs)
go run . -config testconf/template-test-config.yaml -mode simulate

# 2. Run against a live CM-Beetle & CB-Tumblebug server (Real provisioning & 2-Stage Force Cleanup)
go run . -config testconf/template-test-config.yaml -mode real

# 3. Re-generate Markdown and HTML dashboard from an existing JSON dataset
go run . -from-json testresult/stress-test-data.json
```
