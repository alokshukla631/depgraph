// Package output defines the interface and registry for output formatters.
package output

import (
	"io"

	"github.com/alokshukla631/depgraph/internal/model"
)

// Formatter renders the graph and findings to a writer.
type Formatter interface {
	Name() string
	Format(w io.Writer, graph *model.ServiceGraph, findings []model.Finding) error
}
