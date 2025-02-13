package graphs

type WeightedGraph struct {
	vertices map[string]map[string]int
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		vertices: make(map[string]map[string]int),
	}
}

func (g *WeightedGraph) AddEdge(vertex1, vertex2 string, weight int) {
	if _, exists := g.vertices[vertex1]; !exists {
		g.vertices[vertex1] = make(map[string]int)
	}
	if _, exists := g.vertices[vertex2]; !exists {
		g.vertices[vertex2] = make(map[string]int)
	}

	g.vertices[vertex1][vertex2] = weight
	g.vertices[vertex2][vertex1] = weight
}

func (g *WeightedGraph) RemoveEdge(vertex1, vertex2 string) {
	delete(g.vertices[vertex1], vertex2)
	delete(g.vertices[vertex2], vertex1)
}

func (g *WeightedGraph) GetNeighbors(vertex string) map[string]int {
	return g.vertices[vertex]
}

func (g *WeightedGraph) HasEdge(vertex1, vertex2 string) bool {
	_, exists := g.vertices[vertex1][vertex2]
	return exists
}

func (g *WeightedGraph) GetEdgeWeight(vertex1, vertex2 string) (int, bool) {
	weight, exists := g.vertices[vertex1][vertex2]
	return weight, exists
}
