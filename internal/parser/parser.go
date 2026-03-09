// Package parser defines the interface and registry for config file parsers.
package parser

import "github.com/alokshukla631/depgraph/internal/model"

// ServiceDefinition is the raw output of a parser before graph construction.
type ServiceDefinition struct {
	Name          string
	Image         string
	Ports         []model.Port
	Networks      []string
	Volumes       []model.Volume
	EnvVars       map[string]string
	DependsOn     []string // explicit dependencies
	HealthCheck   *model.HealthCheckConfig
	Replicas      int
	Resources     *model.ResourceLimits
	RestartPolicy string
	SourceFile    string
	SourceType    string // "compose", "k8s"
}

// Parser reads a config file and returns raw service definitions.
type Parser interface {
	Name() string
	CanParse(filePath string) bool
	Parse(filePath string) ([]ServiceDefinition, error)
}
