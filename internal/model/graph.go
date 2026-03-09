package model

import "fmt"

// ServiceNode represents a single service in the dependency graph.
type ServiceNode struct {
	Name        string
	Image       string
	Ports       []Port
	Networks    []string
	Volumes     []Volume
	EnvVars     map[string]string
	HealthCheck *HealthCheckConfig
	Replicas    int
	Resources   *ResourceLimits
	LiveStatus  *ProbeResult
	SourceFile  string
	SourceType  string // "compose", "k8s"

	// RestartPolicy stores the restart policy (e.g., "always", "on-failure").
	RestartPolicy string
}

// Edge represents a dependency between two services.
type Edge struct {
	From         string
	To           string
	Type         EdgeType
	Protocol     string // "http", "grpc", "tcp", etc.
	Port         int
	SourceDetail string // e.g., env var name that caused implicit detection
}

// ServiceGraph is the core data structure holding all services and their relationships.
type ServiceGraph struct {
	Nodes map[string]*ServiceNode
	Edges []Edge

	// adjacency lists for fast lookup
	outEdges map[string][]string // service -> services it depends on
	inEdges  map[string][]string // service -> services that depend on it
}

// NewServiceGraph creates an empty graph.
func NewServiceGraph() *ServiceGraph {
	return &ServiceGraph{
		Nodes:    make(map[string]*ServiceNode),
		outEdges: make(map[string][]string),
		inEdges:  make(map[string][]string),
	}
}

// AddNode adds a service node to the graph. Returns an error if the name already exists.
func (g *ServiceGraph) AddNode(node *ServiceNode) error {
	if _, exists := g.Nodes[node.Name]; exists {
		return fmt.Errorf("duplicate service name: %s", node.Name)
	}
	g.Nodes[node.Name] = node
	return nil
}

// AddEdge adds a directed dependency edge from -> to.
func (g *ServiceGraph) AddEdge(e Edge) {
	g.Edges = append(g.Edges, e)
	g.outEdges[e.From] = append(g.outEdges[e.From], e.To)
	g.inEdges[e.To] = append(g.inEdges[e.To], e.From)
}

// GetDependencies returns the names of services that the given service depends on.
func (g *ServiceGraph) GetDependencies(name string) []string {
	return g.outEdges[name]
}

// GetDependents returns the names of services that depend on the given service.
func (g *ServiceGraph) GetDependents(name string) []string {
	return g.inEdges[name]
}

// InDegree returns the number of services that depend on the given service.
func (g *ServiceGraph) InDegree(name string) int {
	return len(g.inEdges[name])
}

// OutDegree returns the number of services the given service depends on.
func (g *ServiceGraph) OutDegree(name string) int {
	return len(g.outEdges[name])
}

// NodeCount returns the number of services in the graph.
func (g *ServiceGraph) NodeCount() int {
	return len(g.Nodes)
}

// EdgeCount returns the total number of dependency edges.
func (g *ServiceGraph) EdgeCount() int {
	return len(g.Edges)
}

// EdgesByType returns a count of edges grouped by EdgeType.
func (g *ServiceGraph) EdgesByType() map[EdgeType]int {
	counts := make(map[EdgeType]int)
	for _, e := range g.Edges {
		counts[e.Type]++
	}
	return counts
}
