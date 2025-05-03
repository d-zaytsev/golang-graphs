package graphs

import (
	"maps"
	"testing"
)

func TestAddEdgeInWeightedGraph(t *testing.T) {
	graph := NewWeightedGraph()
	graph.AddEdge("A", "B", 10)

	if !graph.HasEdge("A", "B") {
		t.Errorf("Expected edge between A and B to exist")
	}
	if !graph.HasEdge("B", "A") {
		t.Errorf("Expected edge between B and A to exist")
	}

	weight, exists := graph.GetEdgeWeight("A", "B")
	if !exists || weight != 10 {
		t.Errorf("Expected edge between A and B with weight 10, but got %d", weight)
	}
}

func TestGetEdgesInWeightedGraph(t *testing.T) {
	edgesExpected := map[Edge]struct{}{
		{"A", "B", 10}: {},
		{"A", "C", 30}: {},
		{"B", "C", 20}: {},
	}

	graph := NewWeightedGraph()
	for edge := range edgesExpected {
		graph.AddEdge(edge.U, edge.V, edge.Weight)
	}
	edges := graph.GetEdges()
	resultEdges := make(map[Edge]struct{})
	for _, e := range edges {
		resultEdges[e] = struct{}{}
	}
	if !maps.Equal(resultEdges, edgesExpected) {
		t.Errorf("Expected edges %v, got %v", edgesExpected, resultEdges)
	}
}

func TestRemoveEdgeInWeightedGraph(t *testing.T) {
	graph := NewWeightedGraph()
	graph.AddEdge("A", "B", 10)
	graph.RemoveEdge("A", "B")

	if graph.HasEdge("A", "B") {
		t.Errorf("Expected edge between A and B to be removed")
	}
	if graph.HasEdge("B", "A") {
		t.Errorf("Expected edge between B and A to be removed")
	}
}

func TestGetNeighborsInWeightedGraph(t *testing.T) {
	graph := NewWeightedGraph()
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

func TestHasEdgeInWeightedGraph(t *testing.T) {
	graph := NewWeightedGraph()
	graph.AddEdge("A", "B", 10)

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

func TestGetEdgeWeightInWeightedGraph(t *testing.T) {
	graph := NewWeightedGraph()
	graph.AddEdge("A", "B", 10)

	weight, exists := graph.GetEdgeWeight("A", "B")
	if !exists || weight != 10 {
		t.Errorf("Expected weight 10 for edge A-B, but got weight %d", weight)
	}

	_, exists = graph.GetEdgeWeight("A", "C")
	if exists {
		t.Errorf("Expected no edge between A and C")
	}
}
