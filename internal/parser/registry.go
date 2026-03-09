package parser

import "fmt"

// Registry holds registered parsers and routes files to the correct one.
type Registry struct {
	parsers []Parser
}

// NewRegistry creates an empty parser registry.
func NewRegistry() *Registry {
	return &Registry{}
}

// Register adds a parser to the registry.
func (r *Registry) Register(p Parser) {
	r.parsers = append(r.parsers, p)
}

// FindParser returns the first parser that can handle the given file.
// Returns an error if no parser matches.
func (r *Registry) FindParser(filePath string) (Parser, error) {
	for _, p := range r.parsers {
		if p.CanParse(filePath) {
			return p, nil
		}
	}
	return nil, fmt.Errorf("no parser found for file: %s", filePath)
}

// ParseFile finds the appropriate parser and parses the file.
func (r *Registry) ParseFile(filePath string) ([]ServiceDefinition, error) {
	p, err := r.FindParser(filePath)
	if err != nil {
		return nil, err
	}
	return p.Parse(filePath)
}
