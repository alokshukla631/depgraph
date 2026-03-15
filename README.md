# depgraph

A linter for your distributed system architecture. Analyze service dependency graphs for structural risks before they cause outages.

## What It Does

`depgraph` parses your Docker Compose and Kubernetes manifests to build a service dependency graph, then analyzes it for:

- **Circular dependencies** between services
- **Single points of failure** — services many others depend on with no replicas
- **Missing health checks** — services with no readiness/liveness probes
- **Missing resource limits** — services without CPU/memory constraints
- **Implicit dependencies** — hidden connections via environment variables (e.g., `DATABASE_URL`)
- **Risk scores** — a 0-100 score per service with actionable recommendations

## Quick Start

```bash
# Analyze a Docker Compose file
depgraph analyze docker-compose.yml

# Output as JSON
depgraph analyze -f json docker-compose.yml

# Generate a Graphviz diagram
depgraph analyze -f dot -o graph.dot docker-compose.yml
dot -Tpng graph.dot -o graph.png

# Analyze Kubernetes manifests
depgraph analyze k8s-manifests/

# Probe running services for health status
depgraph probe docker-compose.yml
```

## Install

```bash
# From source
go install github.com/alokshukla631/depgraph/cmd/depgraph@latest

# Or build from source
git clone https://github.com/alokshukla631/depgraph.git
cd depgraph
make build
./bin/depgraph --help
```

## Status

**Under active development** — see [DESIGN.md](docs/DESIGN.md) for the architecture and roadmap.

## Contributing

Contributions welcome! depgraph is designed to be extensible — you can add new parsers, analyzers, and output formats by implementing Go interfaces. See [CONTRIBUTING.md](docs/CONTRIBUTING.md) for details.

## License

Apache 2.0 — see [LICENSE](LICENSE).
