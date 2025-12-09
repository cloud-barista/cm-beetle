package report

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"github.com/cloud-barista/cm-beetle/deepdiffgo/pkg/diff"
)

func WriteText(w io.Writer, report *diff.DiffReport) {
	if len(report.APIChanges) == 0 {
		fmt.Fprintln(w, "No changes found.")
		return
	}

	fmt.Fprintln(w, "DeepDiffGo Report")
	if report.Spec1 != "" && report.Spec2 != "" {
		fmt.Fprintln(w, "Comparing:")
		fmt.Fprintf(w, "  - Old: %s\n", report.Spec1)
		fmt.Fprintf(w, "  - New: %s\n", report.Spec2)
	}
	fmt.Fprintln(w, "=================")

	// Sort API changes for consistent output
	sortChanges(report)

	for _, apiChange := range report.APIChanges {
		switch apiChange.Type {
		case diff.Added:
			fmt.Fprintf(w, "[+] %s %s (New API)\n", apiChange.Method, apiChange.Path)
		case diff.Removed:
			fmt.Fprintf(w, "[-] %s %s (Deprecated API)\n", apiChange.Method, apiChange.Path)
		case diff.Modified:
			fmt.Fprintf(w, "[*] %s %s\n", apiChange.Method, apiChange.Path)
			for _, change := range apiChange.Changes {
				prefix := ""
				switch change.Type {
				case diff.Added:
					prefix = "+"
				case diff.Removed:
					prefix = "-"
				case diff.Modified:
					prefix = "*"
				}

				msg := change.Message
				if change.Type == diff.Modified {
					msg = fmt.Sprintf("%s (From: %v, To: %v)", change.Message, change.From, change.To)
				}

				fmt.Fprintf(w, "    %s %s: %s\n", prefix, change.Path, msg)
			}
		}
	}
}

func WriteMarkdown(w io.Writer, report *diff.DiffReport) {
	if len(report.APIChanges) == 0 {
		fmt.Fprintln(w, "No changes found.")
		return
	}

	fmt.Fprintln(w, "## DeepDiffGo Report")
	if report.Spec1 != "" && report.Spec2 != "" {
		fmt.Fprintln(w, "**Comparing:**")
		fmt.Fprintf(w, "- Old: `%s`\n", report.Spec1)
		fmt.Fprintf(w, "- New: `%s`\n", report.Spec2)
	}
	fmt.Fprintln(w, "")

	sortChanges(report)

	// Group changes by type
	var modified, added, removed []diff.APIChange
	for _, apiChange := range report.APIChanges {
		switch apiChange.Type {
		case diff.Modified:
			modified = append(modified, apiChange)
		case diff.Added:
			added = append(added, apiChange)
		case diff.Removed:
			removed = append(removed, apiChange)
		}
	}

	// Write Modified APIs section
	if len(modified) > 0 {
		fmt.Fprintln(w, "### Modified APIs")
		fmt.Fprintln(w, "")
		for _, apiChange := range modified {
			fmt.Fprintf(w, "#### `%s` %s\n", apiChange.Method, apiChange.Path)
			for _, change := range apiChange.Changes {
				msg := change.Message
				if change.Type == diff.Modified {
					msg = fmt.Sprintf("%s (From: `%v`, To: `%v`)", change.Message, change.From, change.To)
				}

				fmt.Fprintf(w, "- `%s`: %s\n", change.Path, msg)
			}
			fmt.Fprintln(w, "")
		}
	}

	// Write Added APIs section
	if len(added) > 0 {
		fmt.Fprintln(w, "### Added APIs")
		fmt.Fprintln(w, "")
		for _, apiChange := range added {
			fmt.Fprintf(w, "#### `%s` %s\n", apiChange.Method, apiChange.Path)
			fmt.Fprintln(w, "")
		}
	}

	// Write Removed APIs section
	if len(removed) > 0 {
		fmt.Fprintln(w, "### Removed APIs")
		fmt.Fprintln(w, "")
		for _, apiChange := range removed {
			fmt.Fprintf(w, "#### `%s` %s\n", apiChange.Method, apiChange.Path)
			fmt.Fprintln(w, "")
		}
	}
}

func WriteJSON(w io.Writer, report *diff.DiffReport) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(report)
}

func sortChanges(report *diff.DiffReport) {
	sort.Slice(report.APIChanges, func(i, j int) bool {
		if report.APIChanges[i].Path != report.APIChanges[j].Path {
			return report.APIChanges[i].Path < report.APIChanges[j].Path
		}
		return report.APIChanges[i].Method < report.APIChanges[j].Method
	})
}
