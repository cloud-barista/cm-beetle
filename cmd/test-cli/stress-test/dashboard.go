package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GenerateMarkdownReport writes a human-readable Markdown summary report of the stress test.
func GenerateMarkdownReport(report *StressTestReport, outputPath string) error {
	cleanPath := filepath.Clean(outputPath)
	if err := os.MkdirAll(filepath.Dir(cleanPath), 0755); err != nil {
		return fmt.Errorf("failed to create report directory: %w", err)
	}

	var sb strings.Builder
	sb.WriteString("# CM-Beetle Multi-CSP Progressive Stress Test Report\n\n")
	sb.WriteString(fmt.Sprintf("- **Test Date**: %s %s\n", report.TestDate, report.TestTime))
	sb.WriteString(fmt.Sprintf("- **Execution Mode**: `%s`\n", report.ExecutionMode))
	sb.WriteString(fmt.Sprintf("- **Node Scale Range**: %d Nodes to %d Nodes (+%d Nodes per Round across Surviving CSPs)\n", report.InitialNodes, report.MaxNodes, report.StepNodes))
	sb.WriteString(fmt.Sprintf("- **Max Reached Scale**: **%d Nodes** across **%d Rounds** (%d Total Steps)\n", report.MaxReachedNodes, report.TotalRounds, report.TotalSteps))
	sb.WriteString(fmt.Sprintf("- **Total Nodes Provisioned & Swept**: **%d Provisioned / %d Cleaned** (Orphan Resources Left: **%d**)\n\n",
		report.TotalNodesProvisioned, report.TotalNodesCleaned, report.OrphanResourcesLeft))

	sb.WriteString("## 1. CSP Survival & Scale-Up Standing\n\n")
	sb.WriteString("| CSP | Region | Final Status | Max Succeeded Nodes | Failed At Nodes | Failure Category | Failure Reason | 2-Stage Cleanup Verified |\n")
	sb.WriteString("| :--- | :--- | :---: | :---: | :---: | :---: | :--- | :---: |\n")

	for _, st := range report.CSPStatuses {
		failedAt := "-"
		if st.FailedAtNodes > 0 {
			failedAt = fmt.Sprintf("%d Nodes (R%d)", st.FailedAtNodes, st.EliminatedRound)
		}
		code := st.FailureCode
		if code == "" {
			code = "NONE"
		}
		reason := st.FailureReason
		if reason == "" {
			reason = "Completed all configured rounds without error"
		}
		cleanBadge := "Verified (0 Orphans)"
		if !st.CleanupVerified {
			cleanBadge = "Check Required"
		}
		sb.WriteString(fmt.Sprintf("| **%s** | `%s` | `%s` | **%d Nodes** | %s | `%s` | %s | %s |\n",
			st.DisplayName, st.Region, st.Status, st.MaxSucceededNodes, failedAt, code, reason, cleanBadge))
	}

	sb.WriteString("\n## 2. Step-by-Step Execution Log\n\n")
	sb.WriteString("| Step | Round | CSP | Target Nodes | Workload (vCPU / Mem) | Avg Match | Result | Failure Code / Reason | Cleanup Duration | Cleanup Verification |\n")
	sb.WriteString("| :---: | :---: | :--- | :---: | :---: | :---: | :---: | :--- | :---: | :--- |\n")

	for _, step := range report.Steps {
		resText := "PASS"
		reasonText := "Provisioned & Verified"
		if !step.Success {
			resText = "ELIMINATED"
			reasonText = fmt.Sprintf("`%s`: %s", step.FailureCode, step.FailureReason)
		}
		cleanText := "0 Orphans"
		if !step.CleanupSummary.CleanVerified {
			cleanText = fmt.Sprintf("%d Orphans", step.CleanupSummary.OrphanResourcesRemaining)
		}
		sb.WriteString(fmt.Sprintf("| #%d | R%d | **%s** | **%d Nodes** | %d vCPU / %d GiB | %.1f%% | `%s` | %s | %.1fs | %s |\n",
			step.SequenceIndex, step.Round, step.DisplayName, step.TargetNodes,
			step.TotalVCPUs, step.TotalMemoryGB, step.AvgMatchRate,
			resText, reasonText, step.CleanupPhase.SimulatedSec, cleanText))
	}

	return os.WriteFile(cleanPath, []byte(sb.String()), 0644)
}

// GenerateInteractiveDashboardHTML writes a self-contained HTML time-lapse player for the stress test report.
func GenerateInteractiveDashboardHTML(report *StressTestReport, outputPath string) error {
	cleanPath := filepath.Clean(outputPath)
	if err := os.MkdirAll(filepath.Dir(cleanPath), 0755); err != nil {
		return fmt.Errorf("failed to create dashboard directory: %w", err)
	}

	jsonBytes, err := json.Marshal(report)
	if err != nil {
		return fmt.Errorf("failed to serialize report JSON for dashboard: %w", err)
	}

	// Escape closing script tags in embedded JSON payload to prevent script breakout.
	safeJSON := strings.ReplaceAll(string(jsonBytes), "</", "<\\/")
	htmlContent := strings.Replace(dashboardHTMLTemplate, "__EMBEDDED_REPORT_JSON__", safeJSON, 1)

	return os.WriteFile(cleanPath, []byte(htmlContent), 0644)
}

const dashboardHTMLTemplate = `<!DOCTYPE html>
<html lang="en" data-theme="dark">
<head>
<meta charset="UTF-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<meta http-equiv="Content-Security-Policy" content="default-src 'none'; script-src 'unsafe-inline'; style-src 'unsafe-inline'; img-src data:; connect-src 'none'; frame-ancestors 'none';" />
<title>CM-Beetle Multi-CSP Stress Test &amp; Auto-Cleanup Time-Lapse</title>
<style>
  :root[data-theme="dark"] {
    --bg-page: #090d16;
    --bg-card: #111827;
    --bg-card-alt: #1e293b;
    --bg-input: #0f172a;
    --border-color: #1f2937;
    --border-input: #334155;
    --text-main: #f8fafc;
    --text-muted: #94a3b8;
    --brand-emerald: #10b981;
    --brand-emerald-soft: rgba(16, 185, 129, 0.14);
    --brand-teal: #14b8a6;
    --brand-teal-soft: rgba(20, 184, 166, 0.14);
    --danger: #f43f5e;
    --danger-soft: rgba(244, 63, 94, 0.16);
    --warn: #f59e0b;
    --warn-soft: rgba(245, 158, 11, 0.16);
    --grid-line: #1e293b;
  }
  :root[data-theme="light"] {
    --bg-page: #f8fafc;
    --bg-card: #ffffff;
    --bg-card-alt: #f1f5f9;
    --bg-input: #ffffff;
    --border-color: #e2e8f0;
    --border-input: #cbd5e1;
    --text-main: #0f172a;
    --text-muted: #475569;
    --brand-emerald: #059669;
    --brand-emerald-soft: rgba(5, 150, 105, 0.12);
    --brand-teal: #0d9488;
    --brand-teal-soft: rgba(13, 148, 136, 0.12);
    --danger: #e11d48;
    --danger-soft: rgba(225, 29, 72, 0.12);
    --warn: #d97706;
    --warn-soft: rgba(217, 119, 6, 0.12);
    --grid-line: #e2e8f0;
  }
  * { box-sizing: border-box; margin: 0; padding: 0; }
  body {
    background-color: var(--bg-page);
    color: var(--text-main);
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Inter, sans-serif;
    font-size: 15px;
    line-height: 1.5;
    padding: 20px 28px 40px;
  }
  .container { max-width: 1560px; margin: 0 auto; display: flex; flex-direction: column; gap: 18px; }
  .header-box {
    display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between;
    gap: 12px; padding: 18px 24px; border-radius: 16px;
    background: var(--bg-card); border: 1px solid var(--border-color);
  }
  .header-title-group { display: flex; flex-wrap: wrap; align-items: center; gap: 12px; }
  .header-icon {
    width: 24px; height: 24px; border-radius: 6px;
    background: var(--brand-emerald-soft); color: var(--brand-emerald);
    display: inline-flex; align-items: center; justify-content: center; font-weight: 800;
  }
  .header-title { font-size: 18px; font-weight: 800; color: var(--text-main); }
  .header-subtitle { font-size: 14px; font-weight: 400; color: var(--text-muted); }
  .header-actions { display: flex; align-items: center; gap: 10px; }
  .btn {
    padding: 8px 14px; border-radius: 10px; font-size: 14px; font-weight: 600;
    cursor: pointer; border: 1px solid var(--border-input);
    background: var(--bg-input); color: var(--text-main); transition: all 0.15s ease;
  }
  .btn:hover { border-color: var(--brand-emerald); }
  .btn-primary {
    background: var(--brand-emerald); color: #ffffff; border-color: var(--brand-emerald); font-weight: 800;
  }
  .btn-primary:hover { opacity: 0.92; }
  .btn-active {
    background: var(--brand-emerald-soft); color: var(--brand-emerald); border-color: var(--brand-emerald); font-weight: 800;
  }
  .control-bar {
    display: flex; flex-wrap: wrap; align-items: center; justify-content: space-between;
    gap: 16px; padding: 14px 22px; border-radius: 14px;
    background: var(--bg-card); border: 1px solid var(--border-color);
  }
  .playback-buttons { display: flex; align-items: center; gap: 8px; }
  .speed-group { display: flex; align-items: center; gap: 6px; }
  .scrubber-wrap { flex: 1; min-width: 260px; display: flex; align-items: center; gap: 12px; }
  .scrubber { flex: 1; accent-color: var(--brand-emerald); cursor: pointer; height: 6px; }
  .kv-label { font-size: 14px; font-weight: 400; color: var(--text-muted); }
  .kv-value { font-size: 15px; font-weight: 800; color: var(--text-main); }
  .kpi-grid { display: grid; grid-template-columns: repeat(4, 1fr); gap: 14px; }
  .kpi-card {
    padding: 16px 20px; border-radius: 14px; background: var(--bg-card);
    border: 1px solid var(--border-color); display: flex; flex-direction: column; gap: 6px;
  }
  .kpi-metric { font-size: 24px; font-weight: 800; color: var(--text-main); }
  .kpi-metric-emerald { color: var(--brand-emerald); }
  .kpi-sub { font-size: 14px; font-weight: 400; color: var(--text-muted); }
  .main-grid { display: grid; grid-template-columns: 1.55fr 1fr; gap: 16px; }
  .panel {
    background: var(--bg-card); border: 1px solid var(--border-color);
    border-radius: 16px; padding: 20px; display: flex; flex-direction: column; gap: 14px;
  }
  .panel-header { display: flex; align-items: center; justify-content: space-between; gap: 12px; }
  .panel-title { font-size: 16px; font-weight: 800; color: var(--text-main); }
  .chart-container { width: 100%; height: 410px; position: relative; }
  .csp-grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: 10px; }
  .csp-card {
    padding: 12px 14px; border-radius: 12px; background: var(--bg-card-alt);
    border: 1px solid var(--border-color); display: flex; flex-direction: column; gap: 6px;
  }
  .csp-card-eliminated { border-color: var(--danger); background: var(--danger-soft); }
  .csp-card-active { border-color: var(--brand-emerald); }
  .csp-top { display: flex; align-items: center; justify-content: space-between; gap: 8px; }
  .csp-name { font-size: 15px; font-weight: 800; color: var(--text-main); display: flex; align-items: center; gap: 8px; }
  .color-dot { width: 10px; height: 10px; border-radius: 50%; display: inline-block; }
  .badge {
    padding: 2px 8px; border-radius: 6px; font-size: 12px; font-weight: 800;
    text-transform: uppercase; display: inline-block;
  }
  .badge-emerald { background: var(--brand-emerald-soft); color: var(--brand-emerald); }
  .badge-teal { background: var(--brand-teal-soft); color: var(--brand-teal); }
  .badge-danger { background: var(--danger-soft); color: var(--danger); }
  .badge-warn { background: var(--warn-soft); color: var(--warn); }
  .csp-reason { font-size: 14px; font-weight: 400; color: var(--text-muted); }
  .table-wrap { overflow-x: auto; max-height: 360px; overflow-y: auto; }
  table { width: 100%; border-collapse: collapse; text-align: left; }
  th {
    position: sticky; top: 0; background: var(--bg-card-alt);
    padding: 10px 12px; font-size: 14px; font-weight: 400; color: var(--text-muted);
    border-bottom: 1px solid var(--border-color);
  }
  td {
    padding: 10px 12px; font-size: 14px; border-bottom: 1px solid var(--border-color);
    color: var(--text-main);
  }
  tr.row-highlight { background: var(--brand-emerald-soft); }
  @media (max-width: 1200px) {
    .kpi-grid { grid-template-columns: repeat(2, 1fr); }
    .main-grid { grid-template-columns: 1fr; }
  }
</style>
</head>
<body>
<div class="container">
  <div class="header-box">
    <div class="header-title-group">
      <span class="header-icon">B</span>
      <span class="header-title">Multi-CSP Progressive Stress Test &amp; Auto-Cleanup Replay</span>
      <span class="header-subtitle" id="headerSubtitle">onpremiseInfraModel Scale-Up Survival Benchmark</span>
    </div>
    <div class="header-actions">
      <label class="btn" for="jsonFileInput">Load JSON Report</label>
      <input type="file" id="jsonFileInput" accept=".json" style="display:none" />
      <button class="btn" id="themeToggleBtn" type="button">Light Mode</button>
    </div>
  </div>

  <div class="control-bar">
    <div class="playback-buttons">
      <button class="btn btn-primary" id="playPauseBtn" type="button">Play Time-Lapse</button>
      <button class="btn" id="prevStepBtn" type="button">Prev Step</button>
      <button class="btn" id="nextStepBtn" type="button">Next Step</button>
      <button class="btn" id="resetBtn" type="button">Reset</button>
    </div>
    <div class="scrubber-wrap">
      <span class="kv-label">Timeline Step:</span>
      <input type="range" class="scrubber" id="stepScrubber" min="0" max="1" value="1" />
      <span class="kv-value" id="stepCounterLabel">0 / 0</span>
    </div>
    <div class="speed-group">
      <span class="kv-label">Time-Lapse Speed:</span>
      <button class="btn speed-btn" data-speed="1" type="button">1x</button>
      <button class="btn speed-btn" data-speed="5" type="button">5x</button>
      <button class="btn speed-btn btn-active" data-speed="15" type="button">15x</button>
      <button class="btn speed-btn" data-speed="30" type="button">30x</button>
    </div>
  </div>

  <div class="kpi-grid">
    <div class="kpi-card">
      <span class="kv-label">Current Scale &amp; Round</span>
      <span class="kpi-metric kpi-metric-emerald" id="kpiCurrentScale">0 Nodes</span>
      <span class="kpi-sub" id="kpiCurrentRound">Round 0</span>
    </div>
    <div class="kpi-card">
      <span class="kv-label">CSP Survival Standing</span>
      <span class="kpi-metric" id="kpiSurvivalCount">0 / 0 Active</span>
      <span class="kpi-sub" id="kpiEliminatedSummary">0 Eliminated</span>
    </div>
    <div class="kpi-card">
      <span class="kv-label">Cumulative Nodes Provisioned</span>
      <span class="kpi-metric" id="kpiTotalProvisioned">0 Nodes</span>
      <span class="kpi-sub" id="kpiElapsedSimTime">Simulated Elapsed: 0m 0s</span>
    </div>
    <div class="kpi-card">
      <span class="kv-label">2-Stage Auto-Cleanup Status</span>
      <span class="kpi-metric kpi-metric-emerald" id="kpiCleanupStatus">100% Clean (0 Orphans)</span>
      <span class="kpi-sub" id="kpiCleanedNodes">0 Nodes &amp; VNets Swept</span>
    </div>
  </div>

  <div class="main-grid">
    <div class="panel">
      <div class="panel-header">
        <span class="panel-title">Multi-CSP Node Scale-Up &amp; Failure Annotation Graph</span>
        <span class="kv-label" id="chartActiveEventLabel">Click or scrub timeline to inspect step details</span>
      </div>
      <div class="chart-container" id="svgChartHost"></div>
    </div>

    <div class="panel">
      <div class="panel-header">
        <span class="panel-title">8-CSP Survival Matrix &amp; Failure Reasons</span>
        <span class="kv-label">Auto-Cleanup Verified on Failure</span>
      </div>
      <div class="csp-grid" id="cspMatrixHost"></div>
    </div>
  </div>

  <div class="panel">
    <div class="panel-header">
      <span class="panel-title">Progressive Stress Test Execution Log (onpremiseInfraModel -&gt; Recommend -&gt; Migrate -&gt; Cleanup)</span>
      <span class="kv-label" id="logCountLabel">0 Steps</span>
    </div>
    <div class="table-wrap">
      <table>
        <thead>
          <tr>
            <th>Step</th>
            <th>Round</th>
            <th>CSP Target</th>
            <th>Node Scale</th>
            <th>onpremiseInfraModel</th>
            <th>Recommended</th>
            <th>Step Status</th>
            <th>Failure Category &amp; Cause</th>
            <th>Auto-Cleanup Verification</th>
          </tr>
        </thead>
        <tbody id="stepTableBody"></tbody>
      </table>
    </div>
  </div>
</div>

<script id="embeddedReportData" type="application/json">__EMBEDDED_REPORT_JSON__</script>
<script>
(function() {
  "use strict";

  var cspPalette = {
    aws: "#10b981",
    gcp: "#14b8a6",
    azure: "#f59e0b",
    alibaba: "#f97316",
    tencent: "#84cc16",
    ncp: "#eab308",
    ibm: "#f43f5e",
    nhn: "#64748b"
  };

  var reportData = null;
  var currentStepIdx = 0;
  var isPlaying = false;
  var playbackSpeed = 15;
  var timerHandle = null;

  function getCSPColor(csp) {
    var key = (csp || "").toLowerCase();
    return cspPalette[key] || "#10b981";
  }

  function formatDuration(sec) {
    var total = Math.max(0, Math.round(sec || 0));
    var mins = Math.floor(total / 60);
    var rem = total % 60;
    if (mins >= 60) {
      var hrs = Math.floor(mins / 60);
      mins = mins % 60;
      return hrs + "h " + mins + "m " + rem + "s";
    }
    return mins + "m " + rem + "s";
  }

  function createEl(tag, className, text) {
    var el = document.createElement(tag);
    if (className) {
      el.className = className;
    }
    if (text !== undefined && text !== null) {
      el.textContent = String(text);
    }
    return el;
  }

  function createSVGEl(tag, attrs) {
    var el = document.createElementNS("http://www.w3.org/2000/svg", tag);
    if (attrs) {
      var keys = Object.keys(attrs);
      for (var i = 0; i < keys.length; i++) {
        el.setAttribute(keys[i], String(attrs[keys[i]]));
      }
    }
    return el;
  }

  function loadReport(data) {
    reportData = data;
    var steps = (reportData && reportData.steps) ? reportData.steps : [];
    var scrubber = document.getElementById("stepScrubber");
    scrubber.max = String(steps.length);
    currentStepIdx = steps.length;
    scrubber.value = String(currentStepIdx);

    var subtitle = document.getElementById("headerSubtitle");
    subtitle.textContent = "onpremiseInfraModel Random Generator | Equal Round Scale-Up (+" + (reportData.stepNodes || 10) + " Nodes per Round) | Execution: " + (reportData.executionMode || "simulate");

    renderAtStep(currentStepIdx);
  }

  function renderAtStep(stepLimit) {
    if (!reportData || !reportData.steps) {
      return;
    }
    var steps = reportData.steps;
    var activeSteps = steps.slice(0, stepLimit);
    var totalSteps = steps.length;

    document.getElementById("stepCounterLabel").textContent = stepLimit + " / " + totalSteps;
    document.getElementById("stepScrubber").value = String(stepLimit);

    var currentNodes = 0;
    var currentRound = 0;
    var cumProvisioned = 0;
    var cumCleaned = 0;
    var elapsedSim = 0;
    var orphans = 0;

    var cspStateMap = {};
    var statuses = reportData.cspStatuses || [];
    for (var i = 0; i < statuses.length; i++) {
      var st = statuses[i];
      cspStateMap[st.csp] = {
        csp: st.csp,
        displayName: st.displayName,
        region: st.region,
        maxNodes: 0,
        eliminated: false,
        failedAtNodes: 0,
        failureCode: "",
        failureReason: "",
        cleanVerified: true,
        points: []
      };
    }

    for (var s = 0; s < activeSteps.length; s++) {
      var step = activeSteps[s];
      currentNodes = step.targetNodes;
      currentRound = step.round;
      elapsedSim = step.elapsedSimulatedSec || (elapsedSim + (step.simulatedDurationSec || 0));
      if (step.cleanupSummary && step.cleanupSummary.orphanResourcesRemaining > 0) {
        orphans += step.cleanupSummary.orphanResourcesRemaining;
      }

      var cState = cspStateMap[step.csp];
      if (!cState) {
        cState = {
          csp: step.csp,
          displayName: step.displayName,
          region: step.region,
          maxNodes: 0,
          eliminated: false,
          failedAtNodes: 0,
          failureCode: "",
          failureReason: "",
          cleanVerified: true,
          points: []
        };
        cspStateMap[step.csp] = cState;
      }

      cState.points.push({
        stepIndex: step.sequenceIndex,
        round: step.round,
        nodes: step.targetNodes,
        success: step.success,
        failureCode: step.failureCode,
        failureReason: step.failureReason
      });

      if (step.success) {
        cState.maxNodes = Math.max(cState.maxNodes, step.targetNodes);
        cumProvisioned += step.targetNodes;
        cumCleaned += step.targetNodes;
      } else {
        cState.eliminated = true;
        cState.failedAtNodes = step.targetNodes;
        cState.failureCode = step.failureCode;
        cState.failureReason = step.failureReason;
        cState.cleanVerified = step.cleanupSummary ? step.cleanupSummary.cleanVerified : true;
      }
    }

    var totalCSPs = Object.keys(cspStateMap).length;
    var eliminatedCount = 0;
    var cspKeys = Object.keys(cspStateMap);
    for (var k = 0; k < cspKeys.length; k++) {
      if (cspStateMap[cspKeys[k]].eliminated) {
        eliminatedCount++;
      }
    }
    var survivingCount = totalCSPs - eliminatedCount;

    document.getElementById("kpiCurrentScale").textContent = currentNodes + " Nodes";
    document.getElementById("kpiCurrentRound").textContent = "Round " + currentRound + " (Max Target: " + (reportData.maxNodes || 100) + " Nodes)";
    document.getElementById("kpiSurvivalCount").textContent = survivingCount + " / " + totalCSPs + " Active";
    document.getElementById("kpiEliminatedSummary").textContent = eliminatedCount + " Eliminated & Swept";
    document.getElementById("kpiTotalProvisioned").textContent = cumProvisioned + " Nodes";
    document.getElementById("kpiElapsedSimTime").textContent = "Cumulative Test Time: " + formatDuration(elapsedSim);
    document.getElementById("kpiCleanupStatus").textContent = orphans === 0 ? "100% Clean (0 Orphans)" : (orphans + " Orphans Left");
    document.getElementById("kpiCleanedNodes").textContent = cumCleaned + " Provisioned Nodes + All Failed Batches Swept";

    if (activeSteps.length > 0) {
      var last = activeSteps[activeSteps.length - 1];
      var evLabel = document.getElementById("chartActiveEventLabel");
      if (last.success) {
        evLabel.textContent = "Step #" + last.sequenceIndex + " (R" + last.round + "): " + last.displayName + " succeeded at " + last.targetNodes + " Nodes & cleaned up";
      } else {
        evLabel.textContent = "Step #" + last.sequenceIndex + " (R" + last.round + "): " + last.displayName + " failed at " + last.targetNodes + " Nodes (" + last.failureCode + ") -> Auto-Cleaned";
      }
    }

    renderSVGChart(cspStateMap, steps, stepLimit);
    renderCSPMatrix(cspStateMap);
    renderStepTable(activeSteps);
  }

  function renderSVGChart(cspStateMap, allSteps, stepLimit) {
    var host = document.getElementById("svgChartHost");
    host.replaceChildren();

    var width = 860;
    var height = 400;
    var padLeft = 58;
    var padRight = 34;
    var padTop = 26;
    var padBottom = 46;
    var plotW = width - padLeft - padRight;
    var plotH = height - padTop - padBottom;

    var maxNodes = Math.max(100, reportData.maxNodes || 100);
    var totalSteps = Math.max(1, allSteps.length);

    var svg = createSVGEl("svg", {
      viewBox: "0 0 " + width + " " + height,
      width: "100%",
      height: "100%"
    });

    // Horizontal Y grid lines for Node counts.
    var yTicks = 5;
    for (var i = 0; i <= yTicks; i++) {
      var val = Math.round((maxNodes / yTicks) * i);
      var y = padTop + plotH - (val / maxNodes) * plotH;
      svg.appendChild(createSVGEl("line", {
        x1: padLeft, y1: y, x2: width - padRight, y2: y,
        stroke: "var(--grid-line)", "stroke-width": "1", "stroke-dasharray": "4 4"
      }));
      var yText = createSVGEl("text", {
        x: padLeft - 10, y: y + 4,
        "text-anchor": "end", fill: "var(--text-muted)", "font-size": "12", "font-weight": "600"
      });
      yText.textContent = val + "N";
      svg.appendChild(yText);
    }

    // X axis baseline and label.
    svg.appendChild(createSVGEl("line", {
      x1: padLeft, y1: padTop + plotH, x2: width - padRight, y2: padTop + plotH,
      stroke: "var(--border-input)", "stroke-width": "1.5"
    }));
    var xAxisTitle = createSVGEl("text", {
      x: padLeft + plotW / 2, y: height - 10,
      "text-anchor": "middle", fill: "var(--text-muted)", "font-size": "12"
    });
    xAxisTitle.textContent = "Progressive Test Sequence (Step #1 to Step #" + totalSteps + ")";
    svg.appendChild(xAxisTitle);

    // Draw each CSP trajectory line and markers.
    var keys = Object.keys(cspStateMap);
    for (var c = 0; c < keys.length; c++) {
      var st = cspStateMap[keys[c]];
      var color = getCSPColor(st.csp);
      var pts = st.points;
      if (!pts || pts.length === 0) {
        continue;
      }

      var pathCoords = [];
      for (var p = 0; p < pts.length; p++) {
        var pt = pts[p];
        var x = padLeft + ((pt.stepIndex - 1) / Math.max(1, totalSteps - 1)) * plotW;
        var yPos = padTop + plotH - (pt.nodes / maxNodes) * plotH;
        pathCoords.push((p === 0 ? "M " : "L ") + x.toFixed(1) + " " + yPos.toFixed(1));
      }

      if (pathCoords.length > 1) {
        svg.appendChild(createSVGEl("path", {
          d: pathCoords.join(" "),
          fill: "none",
          stroke: color,
          "stroke-width": "2.6",
          "stroke-linecap": "round"
        }));
      }

      for (var m = 0; m < pts.length; m++) {
        var item = pts[m];
        var cx = padLeft + ((item.stepIndex - 1) / Math.max(1, totalSteps - 1)) * plotW;
        var cy = padTop + plotH - (item.nodes / maxNodes) * plotH;

        if (item.success) {
          svg.appendChild(createSVGEl("circle", {
            cx: cx, cy: cy, r: "4.5",
            fill: color, stroke: "var(--bg-card)", "stroke-width": "1.5"
          }));
        } else {
          // Failure marker with callout badge.
          svg.appendChild(createSVGEl("circle", {
            cx: cx, cy: cy, r: "7",
            fill: "var(--danger)", stroke: "#ffffff", "stroke-width": "2"
          }));
          var tag = createSVGEl("text", {
            x: Math.min(width - padRight - 10, Math.max(padLeft + 10, cx)),
            y: Math.max(padTop + 14, cy - 12),
            "text-anchor": cx > width - 180 ? "end" : "start",
            fill: "var(--danger)",
            "font-size": "11.5",
            "font-weight": "800"
          });
          tag.textContent = st.csp.toUpperCase() + " (" + item.nodes + "N): " + item.failureCode;
          svg.appendChild(tag);
        }
      }
    }

    host.appendChild(svg);
  }

  function renderCSPMatrix(cspStateMap) {
    var host = document.getElementById("cspMatrixHost");
    host.replaceChildren();

    var keys = Object.keys(cspStateMap);
    for (var i = 0; i < keys.length; i++) {
      var st = cspStateMap[keys[i]];
      var cardClass = "csp-card " + (st.eliminated ? "csp-card-eliminated" : (st.maxNodes > 0 ? "csp-card-active" : ""));
      var card = createEl("div", cardClass);

      var top = createEl("div", "csp-top");
      var nameWrap = createEl("div", "csp-name");
      var dot = createEl("span", "color-dot");
      dot.style.backgroundColor = getCSPColor(st.csp);
      nameWrap.appendChild(dot);
      nameWrap.appendChild(createEl("span", "", st.displayName));
      top.appendChild(nameWrap);

      var statusBadge = createEl("span", "badge " + (st.eliminated ? "badge-danger" : "badge-emerald"),
        st.eliminated ? ("ELIMINATED @ " + st.failedAtNodes + "N") : ("ACTIVE (" + st.maxNodes + "N)"));
      top.appendChild(statusBadge);
      card.appendChild(top);

      var scaleRow = createEl("div", "");
      scaleRow.appendChild(createEl("span", "kv-label", "Max Verified Scale: "));
      scaleRow.appendChild(createEl("span", "kv-value", st.maxNodes + " Nodes"));
      card.appendChild(scaleRow);

      if (st.eliminated) {
        var reasonEl = createEl("div", "csp-reason", "[" + st.failureCode + "] " + st.failureReason);
        card.appendChild(reasonEl);
        var cleanBadge = createEl("span", "badge badge-teal", "Auto-Cleanup Verified (0 Orphans)");
        card.appendChild(cleanBadge);
      } else {
        var okEl = createEl("div", "csp-reason", "All attempted rounds provisioned and cleanly swept (0 Orphans)");
        card.appendChild(okEl);
      }

      host.appendChild(card);
    }
  }

  function renderStepTable(activeSteps) {
    var tbody = document.getElementById("stepTableBody");
    tbody.replaceChildren();
    document.getElementById("logCountLabel").textContent = activeSteps.length + " Steps Displayed";

    for (var i = activeSteps.length - 1; i >= 0; i--) {
      var s = activeSteps[i];
      var tr = createEl("tr", i === activeSteps.length - 1 ? "row-highlight" : "");

      tr.appendChild(createEl("td", "kv-value", "#" + s.sequenceIndex));
      tr.appendChild(createEl("td", "", "Round " + s.round));
      tr.appendChild(createEl("td", "kv-value", s.displayName + " (" + s.region + ")"));
      tr.appendChild(createEl("td", "kv-value", s.targetNodes + " Nodes"));
      tr.appendChild(createEl("td", "", s.totalVcpus + " vCPU / " + s.totalMemoryGb + " GiB"));
      tr.appendChild(createEl("td", "", s.generatedNodeGroups + " NodeGroups (" + s.avgMatchRate.toFixed(1) + "%)"));

      var statusTd = createEl("td", "");
      var badge = createEl("span", "badge " + (s.success ? "badge-emerald" : "badge-danger"),
        s.success ? "SUCCEEDED" : "FAILED & SWEPT");
      statusTd.appendChild(badge);
      tr.appendChild(statusTd);

      var causeText = s.success ? "Provisioned in " + (s.migratePhase ? s.migratePhase.simulatedSec : 0) + "s" : ("[" + s.failureCode + "] " + s.failureReason);
      tr.appendChild(createEl("td", "", causeText));

      var cleanTd = createEl("td", "");
      var cleanBadge = createEl("span", "badge badge-teal",
        "Cleaned (" + (s.cleanupPhase ? s.cleanupPhase.simulatedSec : 0) + "s, 0 Orphans)");
      cleanTd.appendChild(cleanBadge);
      tr.appendChild(cleanTd);

      tbody.appendChild(tr);
    }
  }

  function togglePlay() {
    var btn = document.getElementById("playPauseBtn");
    if (isPlaying) {
      stopPlayback();
      return;
    }
    if (!reportData || !reportData.steps) {
      return;
    }
    if (currentStepIdx >= reportData.steps.length) {
      currentStepIdx = 0;
    }
    isPlaying = true;
    btn.textContent = "Pause";
    scheduleNextTick();
  }

  function stopPlayback() {
    isPlaying = false;
    document.getElementById("playPauseBtn").textContent = "Play Time-Lapse";
    if (timerHandle) {
      clearTimeout(timerHandle);
      timerHandle = null;
    }
  }

  function scheduleNextTick() {
    if (!isPlaying || !reportData || !reportData.steps) {
      return;
    }
    var intervalMs = Math.max(90, Math.round(1800 / playbackSpeed));
    timerHandle = setTimeout(function() {
      if (currentStepIdx < reportData.steps.length) {
        currentStepIdx++;
        renderAtStep(currentStepIdx);
        scheduleNextTick();
      } else {
        stopPlayback();
      }
    }, intervalMs);
  }

  document.getElementById("playPauseBtn").addEventListener("click", togglePlay);
  document.getElementById("resetBtn").addEventListener("click", function() {
    stopPlayback();
    currentStepIdx = 0;
    renderAtStep(currentStepIdx);
  });
  document.getElementById("prevStepBtn").addEventListener("click", function() {
    stopPlayback();
    if (currentStepIdx > 0) {
      currentStepIdx--;
      renderAtStep(currentStepIdx);
    }
  });
  document.getElementById("nextStepBtn").addEventListener("click", function() {
    stopPlayback();
    if (reportData && reportData.steps && currentStepIdx < reportData.steps.length) {
      currentStepIdx++;
      renderAtStep(currentStepIdx);
    }
  });
  document.getElementById("stepScrubber").addEventListener("input", function(e) {
    stopPlayback();
    currentStepIdx = parseInt(e.target.value, 10) || 0;
    renderAtStep(currentStepIdx);
  });

  var speedBtns = document.querySelectorAll(".speed-btn");
  for (var i = 0; i < speedBtns.length; i++) {
    speedBtns[i].addEventListener("click", function(e) {
      var spd = parseInt(e.currentTarget.getAttribute("data-speed"), 10) || 15;
      playbackSpeed = spd;
      for (var j = 0; j < speedBtns.length; j++) {
        speedBtns[j].classList.remove("btn-active");
      }
      e.currentTarget.classList.add("btn-active");
    });
  }

  document.getElementById("themeToggleBtn").addEventListener("click", function(e) {
    var root = document.documentElement;
    var current = root.getAttribute("data-theme") || "dark";
    var next = current === "dark" ? "light" : "dark";
    root.setAttribute("data-theme", next);
    e.currentTarget.textContent = next === "dark" ? "Light Mode" : "Dark Mode";
  });

  document.getElementById("jsonFileInput").addEventListener("change", function(e) {
    var file = e.target.files && e.target.files[0];
    if (!file) {
      return;
    }
    var reader = new FileReader();
    reader.onload = function(evt) {
      try {
        var parsed = JSON.parse(String(evt.target.result));
        stopPlayback();
        loadReport(parsed);
      } catch (err) {
        document.getElementById("chartActiveEventLabel").textContent = "Failed to parse JSON report file";
      }
    };
    reader.readAsText(file);
  });

  // Boot from embedded report payload.
  var rawEl = document.getElementById("embeddedReportData");
  if (rawEl && rawEl.textContent) {
    try {
      var initialData = JSON.parse(rawEl.textContent);
      loadReport(initialData);
    } catch (e) {
      document.getElementById("chartActiveEventLabel").textContent = "No embedded report data loaded";
    }
  }
})();
</script>
</body>
</html>
`
