package graphs

type BasicGraph struct {
	Vertices map[string][]string
}

func NewBasicGraph() *BasicGraph {
	return &BasicGraph{
		Vertices: make(map[string][]string),
	}
}

func (g *BasicGraph) AddEdge(vertex1, vertex2 string) {
	g.Vertices[vertex1] = append(g.Vertices[vertex1], vertex2)
	g.Vertices[vertex2] = append(g.Vertices[vertex2], vertex1)
}

func (g *BasicGraph) RemoveEdge(vertex1, vertex2 string) {
	g.Vertices[vertex1] = removeElement(g.Vertices[vertex1], vertex2)
	g.Vertices[vertex2] = removeElement(g.Vertices[vertex2], vertex1)
}

func (g *BasicGraph) GetNeighbors(vertex string) []string {
	return g.Vertices[vertex]
}

func (g *BasicGraph) HasEdge(vertex1, vertex2 string) bool {
	for _, v := range g.Vertices[vertex1] {
		if v == vertex2 {
			return true
		}
	}
	return false
}

func removeElement(slice []string, element string) []string {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
