package graphs

type BasicGraph struct {
	vertices map[string][]string
}

func NewBasicGraph() *BasicGraph {
	return &BasicGraph{
		vertices: make(map[string][]string),
	}
}

func (g *BasicGraph) AddEdge(vertex1, vertex2 string) {
	g.vertices[vertex1] = append(g.vertices[vertex1], vertex2)
	g.vertices[vertex2] = append(g.vertices[vertex2], vertex1)
}

func (g *BasicGraph) RemoveEdge(vertex1, vertex2 string) {
	g.vertices[vertex1] = removeElement(g.vertices[vertex1], vertex2)
	g.vertices[vertex2] = removeElement(g.vertices[vertex2], vertex1)
}

func (g *BasicGraph) GetNeighbors(vertex string) []string {
	return g.vertices[vertex]
}

func (g *BasicGraph) HasEdge(vertex1, vertex2 string) bool {
	for _, v := range g.vertices[vertex1] {
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
