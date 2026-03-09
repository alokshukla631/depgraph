// Package analyzer defines the interface and registry for graph analyzers.
package analyzer

import "github.com/alokshukla631/depgraph/internal/model"

// Analyzer inspects the service graph and returns findings.
type Analyzer interface {
	Name() string
	Analyze(graph *model.ServiceGraph) ([]model.Finding, error)
}
