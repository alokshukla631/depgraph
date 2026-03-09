package model

import "testing"

func TestNewServiceGraph(t *testing.T) {
	g := NewServiceGraph()
	if g.NodeCount() != 0 {
		t.Errorf("expected 0 nodes, got %d", g.NodeCount())
	}
	if g.EdgeCount() != 0 {
		t.Errorf("expected 0 edges, got %d", g.EdgeCount())
	}
}

func TestAddNode(t *testing.T) {
	g := NewServiceGraph()
	err := g.AddNode(&ServiceNode{Name: "web"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.NodeCount() != 1 {
		t.Errorf("expected 1 node, got %d", g.NodeCount())
	}
}

func TestAddDuplicateNode(t *testing.T) {
	g := NewServiceGraph()
	_ = g.AddNode(&ServiceNode{Name: "web"})
	err := g.AddNode(&ServiceNode{Name: "web"})
	if err == nil {
		t.Fatal("expected error for duplicate node")
	}
}

func TestAddEdge(t *testing.T) {
	g := NewServiceGraph()
	_ = g.AddNode(&ServiceNode{Name: "web"})
	_ = g.AddNode(&ServiceNode{Name: "db"})
	g.AddEdge(Edge{From: "web", To: "db", Type: EdgeExplicit})

	if g.EdgeCount() != 1 {
		t.Errorf("expected 1 edge, got %d", g.EdgeCount())
	}

	deps := g.GetDependencies("web")
	if len(deps) != 1 || deps[0] != "db" {
		t.Errorf("expected web -> db, got %v", deps)
	}

	dependents := g.GetDependents("db")
	if len(dependents) != 1 || dependents[0] != "web" {
		t.Errorf("expected db dependents = [web], got %v", dependents)
	}
}

func TestInOutDegree(t *testing.T) {
	g := NewServiceGraph()
	_ = g.AddNode(&ServiceNode{Name: "api"})
	_ = g.AddNode(&ServiceNode{Name: "redis"})
	_ = g.AddNode(&ServiceNode{Name: "worker"})

	g.AddEdge(Edge{From: "api", To: "redis", Type: EdgeImplicit})
	g.AddEdge(Edge{From: "worker", To: "redis", Type: EdgeImplicit})

	if g.InDegree("redis") != 2 {
		t.Errorf("expected in-degree 2 for redis, got %d", g.InDegree("redis"))
	}
	if g.OutDegree("api") != 1 {
		t.Errorf("expected out-degree 1 for api, got %d", g.OutDegree("api"))
	}
}

func TestEdgesByType(t *testing.T) {
	g := NewServiceGraph()
	_ = g.AddNode(&ServiceNode{Name: "a"})
	_ = g.AddNode(&ServiceNode{Name: "b"})
	_ = g.AddNode(&ServiceNode{Name: "c"})

	g.AddEdge(Edge{From: "a", To: "b", Type: EdgeExplicit})
	g.AddEdge(Edge{From: "a", To: "c", Type: EdgeImplicit})
	g.AddEdge(Edge{From: "b", To: "c", Type: EdgeExplicit})

	counts := g.EdgesByType()
	if counts[EdgeExplicit] != 2 {
		t.Errorf("expected 2 explicit, got %d", counts[EdgeExplicit])
	}
	if counts[EdgeImplicit] != 1 {
		t.Errorf("expected 1 implicit, got %d", counts[EdgeImplicit])
	}
}

func TestSeverityRank(t *testing.T) {
	if SeverityCritical.Rank() >= SeverityHigh.Rank() {
		t.Error("critical should rank lower than high")
	}
	if SeverityHigh.Rank() >= SeverityMedium.Rank() {
		t.Error("high should rank lower than medium")
	}
	if SeverityMedium.Rank() >= SeverityLow.Rank() {
		t.Error("medium should rank lower than low")
	}
}
