package graphs

import (
	"testing"
)

func TestAddEdgeInWeightedOrientedGraph(t *testing.T) {
	graph := NewWeightedOrientedGraph()
	graph.AddEdge("A", "B", 10)

	if !graph.HasEdge("A", "B") {
		t.Errorf("Expected edge from A to B to exist")
	}

	if graph.HasEdge("B", "A") {
		t.Errorf("Expected no edge from B to A")
	}

	weight, exists := graph.GetEdgeWeight("A", "B")
	if !exists || weight != 10 {
		t.Errorf("Expected edge from A to B with weight 10, but got %d", weight)
	}
}

func TestRemoveEdgeInWeightedOrientedGraph(t *testing.T) {
	graph := NewWeightedOrientedGraph()
	graph.AddEdge("A", "B", 10)
	graph.RemoveEdge("A", "B")

	if graph.HasEdge("A", "B") {
		t.Errorf("Expected edge from A to B to be removed")
	}

	if graph.HasEdge("B", "A") {
		t.Errorf("Expected no edge from B to A after removal")
	}
}

func TestGetNeighborsInWeightedOrientedGraph(t *testing.T) {
	graph := NewWeightedOrientedGraph()
	graph.AddEdge("A", "B", 10)
	graph.AddEdge("A", "C", 5)

	neighbors := graph.GetNeighbors("A")
	expectedNeighbors := map[string]int{
		"B": 10,
		"C": 5,
	}

	if len(neighbors) != len(expectedNeighbors) {
		t.Errorf("Expected %d neighbors, got %d", len(expectedNeighbors), len(neighbors))
	}

	for neighbor, weight := range expectedNeighbors {
		if w, exists := neighbors[neighbor]; !exists || w != weight {
			t.Errorf("Expected neighbor %s with weight %d, but got weight %d", neighbor, weight, w)
		}
	}
}

func TestHasEdgeInWeightedOrientedGraph(t *testing.T) {
	graph := NewWeightedOrientedGraph()
	graph.AddEdge("A", "B", 10)

	if !graph.HasEdge("A", "B") {
		t.Errorf("Expected edge from A to B to exist")
	}

	if graph.HasEdge("B", "A") {
		t.Errorf("Expected no edge from B to A")
	}

	if graph.HasEdge("A", "C") {
		t.Errorf("Expected no edge from A to C")
	}
}

func TestGetEdgeWeightInWeightedOrientedGraph(t *testing.T) {
	graph := NewWeightedOrientedGraph()
	graph.AddEdge("A", "B", 10)

	weight, exists := graph.GetEdgeWeight("A", "B")
	if !exists || weight != 10 {
		t.Errorf("Expected weight 10 for edge A -> B, but got weight %d", weight)
	}

	_, exists = graph.GetEdgeWeight("B", "A")
	if exists {
		t.Errorf("Expected no edge from B to A")
	}
}
