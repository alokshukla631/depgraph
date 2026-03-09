package prober

import (
	"context"
	"fmt"

	"github.com/alokshukla631/depgraph/internal/model"
)

// Registry holds registered probers.
type Registry struct {
	probers []Prober
}

// NewRegistry creates an empty prober registry.
func NewRegistry() *Registry {
	return &Registry{}
}

// Register adds a prober to the registry.
func (r *Registry) Register(p Prober) {
	r.probers = append(r.probers, p)
}

// FindProber returns the first prober that can handle the given node.
func (r *Registry) FindProber(node *model.ServiceNode) (Prober, error) {
	for _, p := range r.probers {
		if p.CanProbe(node) {
			return p, nil
		}
	}
	return nil, fmt.Errorf("no prober found for service: %s", node.Name)
}

// ProbeNode finds the appropriate prober and probes the service.
func (r *Registry) ProbeNode(ctx context.Context, node *model.ServiceNode) (*model.ProbeResult, error) {
	p, err := r.FindProber(node)
	if err != nil {
		return nil, err
	}
	return p.Probe(ctx, node)
}
