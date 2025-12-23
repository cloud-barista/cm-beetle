package main

import (
	"fmt"
	"io"
	"os"

	"github.com/cloud-barista/cm-beetle/deepdiffgo/pkg/diff"
	"github.com/cloud-barista/cm-beetle/deepdiffgo/pkg/loader"
	"github.com/cloud-barista/cm-beetle/deepdiffgo/pkg/report"
	"github.com/cloud-barista/cm-beetle/deepdiffgo/pkg/resolver"
	"github.com/spf13/cobra"
)

func main() {
	var outputFile string
	var format string
	var oldDesc string
	var newDesc string

	var rootCmd = &cobra.Command{
		Use:   "deepdiffgo [old_spec] [new_spec]",
		Short: "DeepDiffGo compares two Swagger/OpenAPI specifications",
		Long: `DeepDiffGo is a tool for comparing two Swagger/OpenAPI specifications.
  - It detects changes in paths, methods, parameters, responses, and schemas.
  - It supports both local files and remote URLs (http/https).`,
		Example: `  # Compare remote URL with local file
  deepdiffgo https://example.com/v1/swagger.yaml new_swagger.yaml

  # Output as Markdown to a file
  deepdiffgo old.yaml new.yaml -f markdown -o report.md

  # Add descriptions for each spec file
  deepdiffgo old.yaml new.yaml --old-desc "v1.0.0" --new-desc "main/abc123"`,
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			oldPath := args[0]
			newPath := args[1]

			fmt.Printf("Loading %s...\n", oldPath)
			oldSpec, err := loader.Load(oldPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading old spec: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("Loading %s...\n", newPath)
			newSpec, err := loader.Load(newPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading new spec: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Resolving references...")
			if err := resolver.New(oldSpec).Resolve(); err != nil {
				fmt.Fprintf(os.Stderr, "Error resolving old spec: %v\n", err)
				os.Exit(1)
			}
			if err := resolver.New(newSpec).Resolve(); err != nil {
				fmt.Fprintf(os.Stderr, "Error resolving new spec: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Comparing...")
			diffReport, err := diff.Diff(oldSpec, newSpec)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error comparing specs: %v\n", err)
				os.Exit(1)
			}

			diffReport.Spec1 = oldPath
			diffReport.Spec1Desc = oldDesc
			diffReport.Spec2 = newPath
			diffReport.Spec2Desc = newDesc

			var w io.Writer = os.Stdout
			if outputFile != "" {
				f, err := os.Create(outputFile)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
					os.Exit(1)
				}
				defer f.Close()
				w = f
			}

			switch format {
			case "json":
				report.WriteJSON(w, diffReport)
			case "markdown":
				report.WriteMarkdown(w, diffReport)
			default:
				report.WriteText(w, diffReport)
			}
		},
	}

	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file path (default stdout)")
	rootCmd.Flags().StringVarP(&format, "format", "f", "text", "Output format: text, markdown, json")
	rootCmd.Flags().StringVar(&oldDesc, "old-desc", "", "Description for old spec (e.g., tag version, branch/SHA)")
	rootCmd.Flags().StringVar(&newDesc, "new-desc", "", "Description for new spec (e.g., tag version, branch/SHA)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
