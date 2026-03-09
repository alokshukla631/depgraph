// Package prober defines the interface and registry for live service probers.
package prober

import (
	"context"

	"github.com/alokshukla631/depgraph/internal/model"
)

// Prober checks the live status of a service.
type Prober interface {
	Name() string
	CanProbe(node *model.ServiceNode) bool
	Probe(ctx context.Context, node *model.ServiceNode) (*model.ProbeResult, error)
}
