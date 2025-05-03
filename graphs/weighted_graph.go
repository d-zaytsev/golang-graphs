package graphs

type WeightedGraph struct {
	Vertices map[string]map[string]int
}

// Edge represents an edge in a weighted graph.
//
// Fields:
//
//	U: The name of the first vertex (string).
//	V: The name of the second vertex (string).
//	Weight: The weight of the edge (int).
type Edge struct {
	U      string
	V      string
	Weight int
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

// Returns a slice of all edges in the graph.
//
// IMPORTANT: The order of edges in the returned slice is NOT guaranteed to be deterministic.
//
// The function relies on the invariant of the Vertices map (u < v for any edge (u, v))
// to ensure that it only adds one representation of each undirected edge to the result.
func (g *WeightedGraph) GetEdges() []Edge {
	edges := make([]Edge, 0)
	seen := make(map[[2]string]struct{})

	for u := range g.Vertices {
		for v, weight := range g.Vertices[u] {
			a, b := u, v
			if a > b {
				a, b = b, a
			}
			key := [2]string{a, b}

			if _, exists := seen[key]; !exists {
				seen[key] = struct{}{}
				edges = append(edges, Edge{a, b, weight})
			}
		}
	}

	return edges
}
