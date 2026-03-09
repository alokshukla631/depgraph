// Package cli implements the command-line interface for depgraph.
package cli

import (
	"github.com/spf13/cobra"
)

// Version is set at build time via ldflags.
var Version = "dev"

var rootCmd = &cobra.Command{
	Use:   "depgraph",
	Short: "A linter for your distributed system architecture",
	Long: `depgraph analyzes service dependency graphs from Docker Compose files
and Kubernetes manifests to detect structural risks like circular
dependencies, single points of failure, and missing health checks.`,
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().Bool("no-color", false, "Disable colored output")
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}
