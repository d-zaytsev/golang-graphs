package graphs

import (
	"testing"
)

func TestAddEdgeInBasicGraph(t *testing.T) {
	graph := NewBasicGraph()
	graph.AddEdge("A", "B")

	if !graph.HasEdge("A", "B") {
		t.Errorf("Expected edge between A and B to exist")
	}
	if !graph.HasEdge("B", "A") {
		t.Errorf("Expected edge between B and A to exist")
	}
}

func TestRemoveEdgeInBasicGraph(t *testing.T) {
	graph := NewBasicGraph()
	graph.AddEdge("A", "B")
	graph.RemoveEdge("A", "B")

	if graph.HasEdge("A", "B") {
		t.Errorf("Expected edge between A and B to be removed")
	}
	if graph.HasEdge("B", "A") {
		t.Errorf("Expected edge between B and A to be removed")
	}
}

func TestGetNeighborsInBasicGraph(t *testing.T) {
	graph := NewBasicGraph()
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

func TestHasEdgeInBasicGraph(t *testing.T) {
	graph := NewBasicGraph()
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

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
