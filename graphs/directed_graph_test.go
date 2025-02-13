package graphs

import (
	"testing"
)

func TestAddEdgeInDirectedGraph(t *testing.T) {
	graph := NewDirectedGraph()
	graph.AddEdge("A", "B")

	if !graph.HasEdge("A", "B") {
		t.Errorf("Expected edge from A to B to exist")
	}
	if graph.HasEdge("B", "A") {
		t.Errorf("Expected no edge from B to A in directed graph")
	}
}

func TestRemoveEdgeInDirectedGraph(t *testing.T) {
	graph := NewDirectedGraph()
	graph.AddEdge("A", "B")
	graph.RemoveEdge("A", "B")

	if graph.HasEdge("A", "B") {
		t.Errorf("Expected edge from A to B to be removed")
	}
	if graph.HasEdge("B", "A") {
		t.Errorf("Expected no edge from B to A in directed graph")
	}
}

func TestGetNeighborsInDirectedGraph(t *testing.T) {
	graph := NewDirectedGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")

	neighbors := graph.GetNeighbors("A")
	expectedNeighbors := []string{"B", "C"}

	if len(neighbors) != len(expectedNeighbors) {
		t.Errorf("Expected %d neighbors, got %d", len(expectedNeighbors), len(neighbors))
	}

	for _, neighbor := range expectedNeighbors {
		if !contains(neighbors, neighbor) {
			t.Errorf("Expected neighbor %s not found", neighbor)
		}
	}
}

func TestHasEdgeInDirectedGraph(t *testing.T) {
	graph := NewDirectedGraph()
	graph.AddEdge("A", "B")

	if !graph.HasEdge("A", "B") {
		t.Errorf("Expected edge from A to B to exist")
	}
	if graph.HasEdge("B", "A") {
		t.Errorf("Expected no edge from B to A in directed graph")
	}
	if graph.HasEdge("A", "C") {
		t.Errorf("Expected no edge from A to C")
	}
}
