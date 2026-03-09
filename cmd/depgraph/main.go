// Package main is the entry point for the depgraph CLI tool.
// depgraph is a linter for distributed system architectures that analyzes
// service dependency graphs for structural risks.
package main

import (
	"fmt"
	"os"

	"github.com/alokshukla631/depgraph/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
