package graphs

import (
	"fmt"
)

type FlowNetworkVertex int
type FlowNetworkEdge[E comparable] struct {
	Capacity E
	Flow     E
}

type FlowNetwork[E comparable] struct {
	Vertices map[FlowNetworkVertex]map[FlowNetworkVertex]*FlowNetworkEdge[E]
}

func MakeFlowNetwork[E comparable]() *FlowNetwork[E] {
	return &FlowNetwork[E]{
		Vertices: make(map[FlowNetworkVertex]map[FlowNetworkVertex]*FlowNetworkEdge[E]),
	}
}

func (g *FlowNetwork[E]) AddVertex(vertex FlowNetworkVertex) error {
	_, exists := g.Vertices[vertex]

	if exists {
		return fmt.Errorf("Can't add new vertex '%v', it is already exists.", vertex)
	}

	g.Vertices[vertex] = make(map[FlowNetworkVertex]*FlowNetworkEdge[E])

	return nil
}

func (g *FlowNetwork[E]) AddEdge(vertex1, vertex2 FlowNetworkVertex, edge FlowNetworkEdge[E]) error {
	_, exists1 := g.Vertices[vertex1]
	_, exists2 := g.Vertices[vertex2]

	if !exists1 || !exists2 {
		return fmt.Errorf("Can't add edge: vertex %v or %v doesn't exist.", vertex1, vertex2)
	}

	g.Vertices[vertex1][vertex2] = &edge

	return nil
}

func (g *FlowNetwork[E]) RemoveEdge(vertex1, vertex2 FlowNetworkVertex) error {
	_, exists := g.Vertices[vertex1][vertex2]

	if !exists {
		return fmt.Errorf("Can't remove edge (%v,%v). It doesn't exist.", vertex1, vertex2)
	}

	delete(g.Vertices[vertex1], vertex2)

	return nil
}

func (g *FlowNetwork[E]) GetNeighbors(vertex FlowNetworkVertex) (map[FlowNetworkVertex]*FlowNetworkEdge[E], error) {
	neighbors, exists := g.Vertices[vertex]

	if !exists {
		return nil, fmt.Errorf("Can't get neighbors of %v. It doesn't exist.", vertex)
	}

	return neighbors, nil
}

func (g *FlowNetwork[E]) HasVertex(vertex FlowNetworkVertex) bool {
	_, exists := g.Vertices[vertex]

	return exists
}

func (g *FlowNetwork[E]) HasEdge(vertex1, vertex2 FlowNetworkVertex) bool {
	_, exists := g.Vertices[vertex1][vertex2]
	return exists
}

func (g *FlowNetwork[E]) GetEdge(vertex1, vertex2 FlowNetworkVertex) (*FlowNetworkEdge[E], bool) {
	edge, exists := g.Vertices[vertex1][vertex2]
	return edge, exists
}
