package graphs

type DirectedGraph struct {
	vertices map[string][]string
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		vertices: make(map[string][]string),
	}
}

func (g *DirectedGraph) AddEdge(vertex1, vertex2 string) {
	g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
}

func (g *DirectedGraph) RemoveEdge(vertex1, vertex2 string) {
	g.vertices[vertex1] = removeElement(g.vertices[vertex1], vertex2)
}

func (g *DirectedGraph) GetNeighbors(vertex string) []string {
	return g.vertices[vertex]
}

func (g *DirectedGraph) HasEdge(vertex1, vertex2 string) bool {
	for _, v := range g.vertices[vertex1] {
		if v == vertex2 {
			return true
		}
	}
	return false
}
