package graphs

type WeightedGraph struct {
	Vertices map[string]map[string]int
}

func NewWeightedGraph() *WeightedGraph {
	return &WeightedGraph{
		Vertices: make(map[string]map[string]int),
	}
}

func (g *WeightedGraph) AddEdge(vertex1, vertex2 string, weight int) {
	if _, exists := g.Vertices[vertex1]; !exists {
		g.Vertices[vertex1] = make(map[string]int)
	}
	if _, exists := g.Vertices[vertex2]; !exists {
		g.Vertices[vertex2] = make(map[string]int)
	}

	g.Vertices[vertex1][vertex2] = weight
	g.Vertices[vertex2][vertex1] = weight
}

func (g *WeightedGraph) RemoveEdge(vertex1, vertex2 string) {
	delete(g.Vertices[vertex1], vertex2)
	delete(g.Vertices[vertex2], vertex1)
}

func (g *WeightedGraph) GetNeighbors(vertex string) map[string]int {
	return g.Vertices[vertex]
}

func (g *WeightedGraph) HasEdge(vertex1, vertex2 string) bool {
	_, exists := g.Vertices[vertex1][vertex2]
	return exists
}

func (g *WeightedGraph) GetEdgeWeight(vertex1, vertex2 string) (int, bool) {
	weight, exists := g.Vertices[vertex1][vertex2]
	return weight, exists
}
