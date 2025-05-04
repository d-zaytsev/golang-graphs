package algo

import (
	g "dzaytsev/golang-graphs/graphs"
	"fmt"
)

type MaxFlowTaskData struct {
	// graph
	g *g.FlowNetwork[float64]
	// source
	s g.FlowNetworkVertex
	// sink
	t g.FlowNetworkVertex
}

func (data *MaxFlowTaskData) GetCapacity(vertex1, vertex2 g.FlowNetworkVertex) float64 {
	exists := data.g.HasEdge(vertex1, vertex2)
	edge, _ := data.g.GetEdge(vertex1, vertex2)

	if !exists {
		return 0
	} else {
		return edge.Capacity
	}
}

func (data *MaxFlowTaskData) GetFlow(vertex1, vertex2 g.FlowNetworkVertex) float64 {
	exists := data.g.HasEdge(vertex1, vertex2)
	edge, _ := data.g.GetEdge(vertex1, vertex2)

	if !exists {
		return 0
	} else {
		return edge.Flow
	}
}

func (data *MaxFlowTaskData) SetFlow(vertex1, vertex2 g.FlowNetworkVertex, value float64) {
	exists := data.g.HasEdge(vertex1, vertex2)
	edge, _ := data.g.GetEdge(vertex1, vertex2)

	if value <= 0 {
		value = 0
	}

	if !exists {
		data.g.AddEdge(vertex1, vertex2, g.FlowNetworkEdge[float64]{Flow: value})
		return
	}

	edge.Flow = value
}

func MakeNetworkTaskData(network *g.FlowNetwork[float64], s, t g.FlowNetworkVertex) (*MaxFlowTaskData, error) {
	if !network.HasVertex(s) {
		return nil, fmt.Errorf("Source doesn't exist")
	} else if !network.HasVertex(t) {
		return nil, fmt.Errorf("Target doesn't exist")
	}

	return &MaxFlowTaskData{
		g: network,
		s: s,
		t: t,
	}, nil
}
