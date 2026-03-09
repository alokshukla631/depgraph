package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var probeCmd = &cobra.Command{
	Use:   "probe [flags] <file>...",
	Short: "Analyze and actively probe running services for health",
	Long: `Probe performs the same analysis as 'analyze', then actively checks running
services for health status and latency. Requires services to be running.`,
	Args: cobra.MinimumNArgs(1),
	RunE: runProbe,
}

func init() {
	probeCmd.Flags().StringP("format", "f", "table", "Output format: table, json, dot, html")
	probeCmd.Flags().StringP("output", "o", "", "Write output to file instead of stdout")
	probeCmd.Flags().String("severity", "low", "Minimum severity to report: critical, high, medium, low")
	probeCmd.Flags().Bool("skip-implicit", false, "Skip implicit dependency detection from env vars")
	probeCmd.Flags().Bool("no-risk-score", false, "Skip risk score calculation")
	probeCmd.Flags().Duration("timeout", 5*time.Second, "Per-service probe timeout")
	probeCmd.Flags().Int("concurrency", 10, "Max concurrent probes")
	probeCmd.Flags().Int("requests", 5, "Number of probe requests per service")
	rootCmd.AddCommand(probeCmd)
}

func runProbe(cmd *cobra.Command, args []string) error {
	// TODO: Wire up prober after analyze pipeline.
	fmt.Printf("depgraph %s — Service Dependency Analyzer (probe mode)\n\n", Version)
	for _, f := range args {
		fmt.Printf("  Analyzing + probing: %s\n", f)
	}
	fmt.Println("\n  (probe engine not yet wired — coming in M6)")
	return nil
}
