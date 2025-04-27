package graphs

import "fmt"

type FlowNetworkVertex int
type FlowNetworkEdge struct {
	capacity float32
	flow     float32
}

type FlowNetwork struct {
	vertices map[FlowNetworkVertex]map[FlowNetworkVertex]FlowNetworkEdge
}

func MakeFlowNetwork() *FlowNetwork {
	return &FlowNetwork{
		vertices: make(map[FlowNetworkVertex]map[FlowNetworkVertex]FlowNetworkEdge),
	}
}

func (g *FlowNetwork) AddEdge(vertex1, vertex2 FlowNetworkVertex, edge FlowNetworkEdge) (err error) {
	_, exists1 := g.vertices[vertex1]
	_, exists2 := g.vertices[vertex2]

	if !exists1 || !exists2 {
		return fmt.Errorf("Failed to add edge.")
	}

	g.vertices[vertex1][vertex2] = edge

	return nil
}

func (g *FlowNetwork) RemoveEdge(vertex1, vertex2 FlowNetworkVertex) {
	delete(g.vertices[vertex1], vertex2)
}

func (g *FlowNetwork) GetNeighbors(vertex FlowNetworkVertex) map[FlowNetworkVertex]FlowNetworkEdge {
	return g.vertices[vertex]
}

func (g *FlowNetwork) HasEdge(vertex1, vertex2 FlowNetworkVertex) bool {
	_, exists := g.vertices[vertex1][vertex2]
	return exists
}

func (g *FlowNetwork) GetEdge(vertex1, vertex2 FlowNetworkVertex) (FlowNetworkEdge, bool) {
	edge, exists := g.vertices[vertex1][vertex2]
	return edge, exists
}
