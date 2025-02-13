package graphs

import "fmt"

type MultiGraph struct {
	vertices map[string]map[string]int
}

func NewMultiGraph() *MultiGraph {
	return &MultiGraph{
		vertices: make(map[string]map[string]int),
	}
}

func (g *MultiGraph) AddEdge(vertex1, vertex2 string) {
	if g.vertices[vertex1] == nil {
		g.vertices[vertex1] = make(map[string]int)
	}
	if g.vertices[vertex2] == nil {
		g.vertices[vertex2] = make(map[string]int)
	}

	g.vertices[vertex1][vertex2]++
	g.vertices[vertex2][vertex1]++
}

func (g *MultiGraph) RemoveEdge(vertex1, vertex2 string) {
	if g.vertices[vertex1][vertex2] > 0 {
		g.vertices[vertex1][vertex2]--
		g.vertices[vertex2][vertex1]--
	}
}

func (g *MultiGraph) GetNeighbors(vertex string) []string {
	neighbors := []string{}
	for neighbor := range g.vertices[vertex] {
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

func (g *MultiGraph) HasEdge(vertex1, vertex2 string) bool {
	return g.vertices[vertex1][vertex2] > 0
}

func (g *MultiGraph) String() string {
	result := ""
	for vertex1, neighbors := range g.vertices {
		for vertex2, count := range neighbors {
			for i := 0; i < count; i++ {
				result += fmt.Sprintf("%s -- %s\n", vertex1, vertex2)
			}
		}
	}
	return result
}
