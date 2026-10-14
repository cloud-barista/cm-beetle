package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// printComparisonSummary prints a concise Source vs. Target GPU & AI Service comparison table to stdout.
func printComparisonSummary(state *MigrationState, aiCfg AIServiceConfig) {
	s := state.Source
	t := state.Target
	dbMatch := "N/A"
	if s.WebUIDbSha256 != "" && t.WebUIDbSha256 != "" {
		if s.WebUIDbSha256 == t.WebUIDbSha256 {
			dbMatch = "MATCH (100%)"
		} else {
			dbMatch = "MISMATCH"
		}
	}

	srcUrl := firstNonEmpty(s.WebUIUrl, fmt.Sprintf("http://%s:%d", s.PublicIP, aiCfg.OpenWebUIPort))
	tgtUrl := firstNonEmpty(t.WebUIUrl, fmt.Sprintf("http://%s:%d", t.PublicIP, aiCfg.OpenWebUIPort))

	fmt.Println("\n--------------------------------------------------------------------------------")
	fmt.Println("📊 AIMX (AI Migration X) — Source vs. Target GPU AI Service Summary")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Metric", "Source GPU VM", "Target GPU VM")
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Infra / Node ID", fmt.Sprintf("%s / %s", s.InfraID, s.NodeID), fmt.Sprintf("%s / %s", t.InfraID, t.NodeID))
	fmt.Printf("  %-26s | %-28s | %-28s\n", "CSP / Region", fmt.Sprintf("%s / %s", s.Csp, s.Region), fmt.Sprintf("%s / %s", t.Csp, t.Region))
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Instance Spec", s.SpecID, t.SpecID)
	fmt.Printf("  %-26s | %-28s | %-28s\n", "GPU Accelerator", fmt.Sprintf("%dx %s (%.0fGB)", s.GpuCount, s.GpuModel, s.GpuVramGB), fmt.Sprintf("%dx %s (%.0fGB)", t.GpuCount, t.GpuModel, t.GpuVramGB))
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Driver / CUDA", fmt.Sprintf("%s / %s", s.DriverVersion, s.CudaVersion), fmt.Sprintf("%s / %s", t.DriverVersion, t.CudaVersion))
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Gemma Model", s.ModelName, t.ModelName)
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Ollama Processor", s.OllamaProcessor, t.OllamaProcessor)
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Inference Speed", fmt.Sprintf("%.2f tokens/sec", s.EvalTokensPerSec), fmt.Sprintf("%.2f tokens/sec", t.EvalTokensPerSec))
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Open WebUI HTTP Status", fmt.Sprintf("HTTP %d (v%s)", s.WebUIHttpStatus, firstNonEmpty(s.WebUIVersion, "-")), fmt.Sprintf("HTTP %d (v%s)", t.WebUIHttpStatus, firstNonEmpty(t.WebUIVersion, "-")))
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Open WebUI /api/config", fmt.Sprintf("ok=%v (ext=%v)", s.WebUIConfigOK, s.ExternalWebUIReachable), fmt.Sprintf("ok=%v (ext=%v)", t.WebUIConfigOK, t.ExternalWebUIReachable))
	fmt.Printf("  %-26s | %-28s | %-28s\n", "Data Volume Size", s.VolumeSizeHuman, t.VolumeSizeHuman)
	fmt.Printf("  %-26s | %-59s\n", "WebUI SQLite SHA-256", dbMatch)
	fmt.Printf("  %-26s | %-59s\n", "transx Transfer Duration", fmt.Sprintf("%s (strategy: %s)", state.DataTransferDuration, state.DataTransferStrategy))
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("🌐 [Source Open WebUI Browser URL] 👉 %s\n", srcUrl)
	fmt.Printf("🌐 [Target Open WebUI Browser URL] 👉 %s\n", tgtUrl)
	fmt.Println("--------------------------------------------------------------------------------")
}

// writeMarkdownReport writes a comprehensive Markdown report to testresult/aimx-migration-report-{csp}-{region}.md.
func writeMarkdownReport(cfg TestConfig, target TestCase, state *MigrationState) (string, error) {
	outDir := filepath.Dir(cfg.Beetle.StateFile)
	if outDir == "" || outDir == "." {
		outDir = "testresult"
	}
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create report directory %s: %w", outDir, err)
	}

	reportFile := filepath.Join(outDir, fmt.Sprintf("aimx-migration-report-%s-%s.md",
		strings.ToLower(target.Csp), strings.ToLower(target.Region)))

	s := state.Source
	t := state.Target
	srcUrl := firstNonEmpty(s.WebUIUrl, fmt.Sprintf("http://%s:%d", s.PublicIP, cfg.AIService.OpenWebUIPort))
	tgtUrl := firstNonEmpty(t.WebUIUrl, fmt.Sprintf("http://%s:%d", t.PublicIP, cfg.AIService.OpenWebUIPort))

	dbStatus := "Pending"
	if s.WebUIDbSha256 != "" && t.WebUIDbSha256 != "" {
		if s.WebUIDbSha256 == t.WebUIDbSha256 {
			dbStatus = fmt.Sprintf("Verified Identical (`%s`)", shortHash(t.WebUIDbSha256))
		} else {
			dbStatus = fmt.Sprintf("Differed (Source `%s` vs Target `%s`)", shortHash(s.WebUIDbSha256), shortHash(t.WebUIDbSha256))
		}
	}

	var b strings.Builder
	b.WriteString("# AIMX (AI Migration X) — AI Service Migration Verification Report\n\n")
	b.WriteString(fmt.Sprintf("- **Generated At**: `%s`\n", time.Now().Format(time.RFC3339)))
	b.WriteString(fmt.Sprintf("- **Target Cloud**: `%s` (`%s`)\n", target.Csp, target.Region))
	b.WriteString(fmt.Sprintf("- **AI Model & App Stack**: Ollama (`%s`) + Open WebUI (`:%d`)\n", cfg.AIService.ModelName, cfg.AIService.OpenWebUIPort))
	b.WriteString(fmt.Sprintf("- **`transx` Data Transfer Duration**: `%s` (Strategy: `%s`, TransferModelWeights: `%v`, EncryptedKeys: `%v`)\n",
		state.DataTransferDuration, state.DataTransferStrategy, cfg.AIService.TransferModelWeights, cfg.AIService.UseDataEncryption))
	b.WriteString(fmt.Sprintf("- **Source Open WebUI URL**: [%s](%s)\n", srcUrl, srcUrl))
	b.WriteString(fmt.Sprintf("- **Target Open WebUI URL**: [%s](%s)\n\n", tgtUrl, tgtUrl))

	b.WriteString("## 1. Source vs. Target GPU Infrastructure & AI Service Comparison\n\n")
	b.WriteString("| Layer | Metric | Source GPU Environment | Target GPU Environment |\n")
	b.WriteString("| :--- | :--- | :--- | :--- |\n")
	b.WriteString(fmt.Sprintf("| **Infra** | Infra ID / Node ID | `%s` / `%s` | `%s` / `%s` |\n", s.InfraID, s.NodeID, t.InfraID, t.NodeID))
	b.WriteString(fmt.Sprintf("| **Infra** | CSP / Region | `%s` / `%s` | `%s` / `%s` |\n", s.Csp, s.Region, t.Csp, t.Region))
	b.WriteString(fmt.Sprintf("| **Infra** | Public IP | `%s` | `%s` |\n", s.PublicIP, t.PublicIP))
	b.WriteString(fmt.Sprintf("| **Infra** | Instance Spec | `%s` | `%s` |\n", s.SpecID, t.SpecID))
	b.WriteString(fmt.Sprintf("| **Infra** | Machine Image | `%s` | `%s` |\n", s.ImageID, t.ImageID))
	b.WriteString(fmt.Sprintf("| **GPU** | Accelerator Model | `%dx %s (%.0f GB VRAM)` | `%dx %s (%.0f GB VRAM)` |\n", s.GpuCount, s.GpuModel, s.GpuVramGB, t.GpuCount, t.GpuModel, t.GpuVramGB))
	b.WriteString(fmt.Sprintf("| **GPU** | Driver / CUDA Version | `%s` / `CUDA %s` | `%s` / `CUDA %s` |\n", s.DriverVersion, s.CudaVersion, t.DriverVersion, t.CudaVersion))
	b.WriteString(fmt.Sprintf("| **Model** | Gemma Model Tag | `%s` | `%s` |\n", s.ModelName, t.ModelName))
	b.WriteString(fmt.Sprintf("| **Model** | Ollama Processor (`ollama ps`) | `%s` | `%s` |\n", s.OllamaProcessor, t.OllamaProcessor))
	b.WriteString(fmt.Sprintf("| **Model** | Inference Throughput | **`%.2f tokens/sec`** | **`%.2f tokens/sec`** |\n", s.EvalTokensPerSec, t.EvalTokensPerSec))
	b.WriteString(fmt.Sprintf("| **App** | Open WebUI Browser Link | [%s](%s) (HTTP `%d`) | [%s](%s) (HTTP `%d`) |\n", srcUrl, srcUrl, s.WebUIHttpStatus, tgtUrl, tgtUrl, t.WebUIHttpStatus))
	b.WriteString(fmt.Sprintf("| **App** | Open WebUI `/api/config` Probe | `ok=%v` (`v%s`, external=`%v`) | `ok=%v` (`v%s`, external=`%v`) |\n", s.WebUIConfigOK, firstNonEmpty(s.WebUIVersion, "-"), s.ExternalWebUIReachable, t.WebUIConfigOK, firstNonEmpty(t.WebUIVersion, "-"), t.ExternalWebUIReachable))
	b.WriteString(fmt.Sprintf("| **App** | Open WebUI DB Integrity (`webui.db`) | `%s` | %s |\n", shortHash(s.WebUIDbSha256), dbStatus))
	b.WriteString(fmt.Sprintf("| **Data** | Service Directory Size (`%s`) | `%s` | `%s` |\n\n", cfg.AIService.RootDir, s.VolumeSizeHuman, t.VolumeSizeHuman))

	b.WriteString("## 2. Gemma GPU Inference Verification Output\n\n")
	b.WriteString(fmt.Sprintf("- **Prompt**: `%s`\n", cfg.AIService.TestPrompt))
	b.WriteString(fmt.Sprintf("- **Source Response**:\n  > %s\n\n", oneLine(s.SampleResponse)))
	b.WriteString(fmt.Sprintf("- **Target Response**:\n  > %s\n\n", oneLine(t.SampleResponse)))

	if len(state.StageDurations) > 0 {
		b.WriteString("## 3. Stage Execution Timings\n\n")
		b.WriteString("| Stage | Elapsed Duration |\n")
		b.WriteString("| :--- | :--- |\n")
		for _, k := range []string{"setup-source", "recommend", "migrate", "verify"} {
			if d, ok := state.StageDurations[k]; ok {
				b.WriteString(fmt.Sprintf("| `%s` | `%s` |\n", k, d))
			}
		}
	}

	if err := os.WriteFile(reportFile, []byte(b.String()), 0644); err != nil {
		return "", fmt.Errorf("failed to write markdown report %s: %w", reportFile, err)
	}
	return reportFile, nil
}

// shortHash returns the first 16 hex characters of a SHA-256 digest for table readability.
func shortHash(h string) string {
	if len(h) > 16 {
		return h[:16] + "..."
	}
	if h == "" {
		return "N/A"
	}
	return h
}

// oneLine collapses newlines in an LLM response string for blockquote formatting.
func oneLine(s string) string {
	cleaned := strings.ReplaceAll(strings.TrimSpace(s), "\n", " ")
	if cleaned == "" {
		return "(no response captured)"
	}
	return cleaned
}
