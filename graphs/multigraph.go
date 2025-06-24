package graphs

import "fmt"

type MultiGraph struct {
	Vertices map[string]map[string]int
}

func NewMultiGraph() *MultiGraph {
	return &MultiGraph{
		Vertices: make(map[string]map[string]int),
	}
}

func (g *MultiGraph) AddEdge(vertex1, vertex2 string) {
	if g.Vertices[vertex1] == nil {
		g.Vertices[vertex1] = make(map[string]int)
	}
	if g.Vertices[vertex2] == nil {
		g.Vertices[vertex2] = make(map[string]int)
	}

	g.Vertices[vertex1][vertex2]++
	g.Vertices[vertex2][vertex1]++
}

func (g *MultiGraph) RemoveEdge(vertex1, vertex2 string) {
	if g.Vertices[vertex1][vertex2] > 0 {
		g.Vertices[vertex1][vertex2]--
		g.Vertices[vertex2][vertex1]--
	}
}

func (g *MultiGraph) GetNeighbors(vertex string) []string {
	neighbors := []string{}
	for neighbor := range g.Vertices[vertex] {
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

func (g *MultiGraph) HasEdge(vertex1, vertex2 string) bool {
	return g.Vertices[vertex1][vertex2] > 0
}

func (g *MultiGraph) String() string {
	result := ""
	for vertex1, neighbors := range g.Vertices {
		for vertex2, count := range neighbors {
			for i := 0; i < count; i++ {
				result += fmt.Sprintf("%s -- %s\n", vertex1, vertex2)
			}
		}
	}
	return result
}
