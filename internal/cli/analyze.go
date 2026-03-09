package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze [flags] <file>...",
	Short: "Parse config files and report structural risks",
	Long: `Analyze Docker Compose files and Kubernetes manifests to build a service
dependency graph and detect structural risks such as circular dependencies,
single points of failure, missing health checks, and more.`,
	Args: cobra.MinimumNArgs(1),
	RunE: runAnalyze,
}

func init() {
	analyzeCmd.Flags().StringP("format", "f", "table", "Output format: table, json, dot, html")
	analyzeCmd.Flags().StringP("output", "o", "", "Write output to file instead of stdout")
	analyzeCmd.Flags().String("severity", "low", "Minimum severity to report: critical, high, medium, low")
	analyzeCmd.Flags().Bool("skip-implicit", false, "Skip implicit dependency detection from env vars")
	analyzeCmd.Flags().Bool("no-risk-score", false, "Skip risk score calculation")
	rootCmd.AddCommand(analyzeCmd)
}

func runAnalyze(cmd *cobra.Command, args []string) error {
	// TODO: Wire up parser registry, graph builder, analyzers, and output.
	// This will be implemented in subsequent CRs.
	fmt.Printf("depgraph %s — Service Dependency Analyzer\n\n", Version)
	for _, f := range args {
		fmt.Printf("  Analyzing: %s\n", f)
	}
	fmt.Println("\n  (analysis engine not yet wired — coming in M2)")
	return nil
}
