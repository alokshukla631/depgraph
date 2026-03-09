// Package graph constructs a ServiceGraph from parsed service definitions.
package graph

import (
	"github.com/alokshukla631/depgraph/internal/model"
	"github.com/alokshukla631/depgraph/internal/parser"
)

// Build constructs a ServiceGraph from a slice of ServiceDefinitions.
func Build(defs []parser.ServiceDefinition) (*model.ServiceGraph, error) {
	g := model.NewServiceGraph()

	// First pass: add all nodes.
	for i := range defs {
		d := &defs[i]
		node := &model.ServiceNode{
			Name:          d.Name,
			Image:         d.Image,
			Ports:         d.Ports,
			Networks:      d.Networks,
			Volumes:       d.Volumes,
			EnvVars:       d.EnvVars,
			HealthCheck:   d.HealthCheck,
			Replicas:      d.Replicas,
			Resources:     d.Resources,
			RestartPolicy: d.RestartPolicy,
			SourceFile:    d.SourceFile,
			SourceType:    d.SourceType,
		}
		if err := g.AddNode(node); err != nil {
			return nil, err
		}
	}

	// Second pass: add explicit edges from depends_on.
	for _, d := range defs {
		for _, dep := range d.DependsOn {
			g.AddEdge(model.Edge{
				From: d.Name,
				To:   dep,
				Type: model.EdgeExplicit,
			})
		}
	}

	return g, nil
}
