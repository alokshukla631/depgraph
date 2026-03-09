package output

import "fmt"

// Registry holds registered output formatters.
type Registry struct {
	formatters map[string]Formatter
}

// NewRegistry creates an empty output registry.
func NewRegistry() *Registry {
	return &Registry{
		formatters: make(map[string]Formatter),
	}
}

// Register adds a formatter to the registry.
func (r *Registry) Register(f Formatter) {
	r.formatters[f.Name()] = f
}

// Get returns the formatter with the given name.
func (r *Registry) Get(name string) (Formatter, error) {
	f, ok := r.formatters[name]
	if !ok {
		return nil, fmt.Errorf("unknown output format: %s", name)
	}
	return f, nil
}

// Names returns the list of registered formatter names.
func (r *Registry) Names() []string {
	names := make([]string, 0, len(r.formatters))
	for n := range r.formatters {
		names = append(names, n)
	}
	return names
}
