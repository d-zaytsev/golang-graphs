package graphs

import (
	"fmt"
)

type FlowNetworkVertex int
type FlowNetworkEdge[E any] struct {
	Capacity E
	Flow     E
}

type FlowNetwork[E any] struct {
	vertices map[FlowNetworkVertex]map[FlowNetworkVertex]FlowNetworkEdge[E]
}

func MakeFlowNetwork[E any]() *FlowNetwork[E] {
	return &FlowNetwork[E]{
		vertices: make(map[FlowNetworkVertex]map[FlowNetworkVertex]FlowNetworkEdge[E]),
	}
}

func (g *FlowNetwork[E]) AddVertex(vertex FlowNetworkVertex) error {
	_, exists := g.vertices[vertex]

	if exists {
		return fmt.Errorf("Can't add new vertex '%v', it is already exists.", vertex)
	}

	g.vertices[vertex] = make(map[FlowNetworkVertex]FlowNetworkEdge[E])

	return nil
}

func (g *FlowNetwork[E]) AddEdge(vertex1, vertex2 FlowNetworkVertex, edge FlowNetworkEdge[E]) error {
	_, exists1 := g.vertices[vertex1]
	_, exists2 := g.vertices[vertex2]

	if !exists1 || !exists2 {
		return fmt.Errorf("Can't add edge: vertex %v or %v doesn't exist.", vertex1, vertex2)
	}

	g.vertices[vertex1][vertex2] = edge

	return nil
}

func (g *FlowNetwork[E]) RemoveEdge(vertex1, vertex2 FlowNetworkVertex) error {
	_, exists := g.vertices[vertex1][vertex2]

	if !exists {
		return fmt.Errorf("Can't remove edge (%v,%v). It doesn't exist.", vertex1, vertex2)
	}

	delete(g.vertices[vertex1], vertex2)

	return nil
}

func (g *FlowNetwork[E]) GetNeighbors(vertex FlowNetworkVertex) (map[FlowNetworkVertex]FlowNetworkEdge[E], error) {
	neighbors, exists := g.vertices[vertex]

	if !exists {
		return nil, fmt.Errorf("Can't get neighbors of %v. It doesn't exist.", vertex)
	}

	return neighbors, nil
}

func (g *FlowNetwork[E]) HasVertex(vertex FlowNetworkVertex) bool {
	_, exists := g.vertices[vertex]

	return exists
}

func (g *FlowNetwork[E]) HasEdge(vertex1, vertex2 FlowNetworkVertex) bool {
	_, exists := g.vertices[vertex1][vertex2]
	return exists
}

func (g *FlowNetwork[E]) GetEdge(vertex1, vertex2 FlowNetworkVertex) (*FlowNetworkEdge[E], bool) {
	edge, exists := g.vertices[vertex1][vertex2]
	return &edge, exists
}
