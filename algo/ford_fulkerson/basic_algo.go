package algo

import (
	g "dzaytsev/golang-graphs/graphs"
)

type NetworkTaskData struct {
	// graph
	g g.FlowNetwork[float32]
	// source
	s g.FlowNetworkVertex
	// sink
	t g.FlowNetworkVertex
}
