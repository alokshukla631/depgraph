package analyzer

import "github.com/alokshukla631/depgraph/internal/model"

// Registry holds registered analyzers and runs them in order.
type Registry struct {
	analyzers []Analyzer
}

// NewRegistry creates an empty analyzer registry.
func NewRegistry() *Registry {
	return &Registry{}
}

// Register adds an analyzer to the registry.
func (r *Registry) Register(a Analyzer) {
	r.analyzers = append(r.analyzers, a)
}

// RunAll executes all registered analyzers against the graph and returns
// all findings combined.
func (r *Registry) RunAll(graph *model.ServiceGraph) ([]model.Finding, error) {
	var allFindings []model.Finding
	for _, a := range r.analyzers {
		findings, err := a.Analyze(graph)
		if err != nil {
			return nil, err
		}
		allFindings = append(allFindings, findings...)
	}
	return allFindings, nil
}
