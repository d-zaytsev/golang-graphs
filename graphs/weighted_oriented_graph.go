package graphs

type WeightedOrientedGraph struct {
	vertices map[string]map[string]int
}

func NewWeightedOrientedGraph() *WeightedOrientedGraph {
	return &WeightedOrientedGraph{
		vertices: make(map[string]map[string]int),
	}
}

func (g *WeightedOrientedGraph) AddEdge(vertex1, vertex2 string, weight int) {
	if _, exists := g.vertices[vertex1]; !exists {
		g.vertices[vertex1] = make(map[string]int)
	}
	if _, exists := g.vertices[vertex2]; !exists {
		g.vertices[vertex2] = make(map[string]int)
	}

	g.vertices[vertex1][vertex2] = weight
}

func (g *WeightedOrientedGraph) RemoveEdge(vertex1, vertex2 string) {
	delete(g.vertices[vertex1], vertex2)
}

func (g *WeightedOrientedGraph) GetNeighbors(vertex string) map[string]int {
	return g.vertices[vertex]
}

func (g *WeightedOrientedGraph) HasEdge(vertex1, vertex2 string) bool {
	_, exists := g.vertices[vertex1][vertex2]
	return exists
}

func (g *WeightedOrientedGraph) GetEdgeWeight(vertex1, vertex2 string) (int, bool) {
	weight, exists := g.vertices[vertex1][vertex2]
	return weight, exists
}
