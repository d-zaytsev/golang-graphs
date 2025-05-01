package graphs

type DirectedGraph struct {
	Vertices map[string][]string
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		Vertices: make(map[string][]string),
	}
}

func (g *DirectedGraph) AddEdge(vertex1, vertex2 string) {
	g.Vertices[vertex1] = append(g.Vertices[vertex1], vertex2)
}

func (g *DirectedGraph) RemoveEdge(vertex1, vertex2 string) {
	g.Vertices[vertex1] = removeElement(g.Vertices[vertex1], vertex2)
}

func (g *DirectedGraph) GetNeighbors(vertex string) []string {
	return g.Vertices[vertex]
}

func (g *DirectedGraph) HasEdge(vertex1, vertex2 string) bool {
	for _, v := range g.Vertices[vertex1] {
		if v == vertex2 {
			return true
		}
	}
	return false
}
