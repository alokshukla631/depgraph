// Package model defines the core data types for depgraph.
package model

import "time"

// EdgeType classifies how a dependency was discovered.
type EdgeType string

const (
	EdgeExplicit EdgeType = "explicit" // depends_on, links
	EdgeImplicit EdgeType = "implicit" // env var URL pattern
	EdgeNetwork  EdgeType = "network"  // shared network
	EdgeVolume   EdgeType = "volume"   // shared volume
)

// Severity indicates the impact level of a finding.
type Severity string

const (
	SeverityCritical      Severity = "critical"
	SeverityHigh          Severity = "high"
	SeverityMedium        Severity = "medium"
	SeverityLow           Severity = "low"
	SeverityInformational Severity = "info"
)

// SeverityRank returns a numeric rank for sorting (lower = more severe).
func (s Severity) Rank() int {
	switch s {
	case SeverityCritical:
		return 0
	case SeverityHigh:
		return 1
	case SeverityMedium:
		return 2
	case SeverityLow:
		return 3
	case SeverityInformational:
		return 4
	default:
		return 5
	}
}

// Port represents a port mapping.
type Port struct {
	Host      int
	Container int
	Protocol  string // "tcp" or "udp"
}

// Volume represents a volume mount.
type Volume struct {
	Source string
	Target string
}

// HealthCheckConfig describes a service's health check configuration.
type HealthCheckConfig struct {
	Type     string        // "http", "tcp", "cmd"
	Endpoint string        // URL, port, or command
	Interval time.Duration
	Timeout  time.Duration
	Retries  int
}

// ResourceLimits describes CPU/memory constraints.
type ResourceLimits struct {
	CPULimit      string
	MemoryLimit   string
	CPURequest    string
	MemoryRequest string
}

// ProbeResult holds the outcome of a live health probe.
type ProbeResult struct {
	Reachable  bool
	StatusCode int
	Latency    time.Duration
	CheckedAt  time.Time
	Error      string
}
