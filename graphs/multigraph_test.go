package graphs

import (
	"testing"
)

func TestAddEdgeInMultiGraph(t *testing.T) {
	graph := NewMultiGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "B")

	if !graph.HasEdge("A", "B") {
		t.Errorf("Expected edge between A and B to exist")
	}
	if !graph.HasEdge("B", "A") {
		t.Errorf("Expected edge between B and A to exist")
	}

	if countEdges(graph, "A", "B") != 2 {
		t.Errorf("Expected 2 edges between A and B, got %d", countEdges(graph, "A", "B"))
	}
}

func TestRemoveEdgeInMultiGraph(t *testing.T) {
	graph := NewMultiGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "B")

	graph.RemoveEdge("A", "B")

	if countEdges(graph, "A", "B") != 1 {
		t.Errorf("Expected 1 edge between A and B, got %d", countEdges(graph, "A", "B"))
	}

	graph.RemoveEdge("A", "B")

	if graph.HasEdge("A", "B") {
		t.Errorf("Expected no edge between A and B")
	}
	if graph.HasEdge("B", "A") {
		t.Errorf("Expected no edge between B and A")
	}
}

func TestGetNeighborsInMultiGraph(t *testing.T) {
	graph := NewMultiGraph()
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

func TestHasEdgeInMultiGraph(t *testing.T) {
	graph := NewMultiGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "B")

	if !graph.HasEdge("A", "B") {
		t.Errorf("Expected edge between A and B to exist")
	}
	if !graph.HasEdge("B", "A") {
		t.Errorf("Expected edge between B and A to exist")
	}

	if graph.HasEdge("A", "C") {
		t.Errorf("Expected no edge between A and C")
	}
}

func TestMultipleEdgesBetweenVertices(t *testing.T) {
	graph := NewMultiGraph()
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "B")

	if countEdges(graph, "A", "B") != 2 {
		t.Errorf("Expected 2 edges between A and B, got %d", countEdges(graph, "A", "B"))
	}
}

func countEdges(graph *MultiGraph, vertex1, vertex2 string) int {
	return graph.Vertices[vertex1][vertex2]
}
