package graphs

import (
	"testing"
)

func TestAddEdgeInFlowNetwork(t *testing.T) {
	graph := MakeFlowNetwork[float32]()

	graph.AddVertex(1)
	graph.AddVertex(2)

	graph.AddEdge(1, 2, FlowNetworkEdge[float32]{0, 1})
	graph.AddEdge(2, 1, FlowNetworkEdge[float32]{0, 1})

	if !graph.HasEdge(1, 2) || !graph.HasEdge(2, 1) {
		t.FailNow()
	}
}

func TestRemoveEdgeInFlowNetwork(t *testing.T) {
	graph := MakeFlowNetwork[float32]()

	graph.AddVertex(1)
	graph.AddVertex(2)

	graph.AddEdge(1, 2, FlowNetworkEdge[float32]{0, 1})
	graph.AddEdge(2, 1, FlowNetworkEdge[float32]{0, 1})

	graph.RemoveEdge(1, 2)
	graph.RemoveEdge(2, 1)

	if graph.HasEdge(1, 2) || graph.HasEdge(2, 1) {
		t.FailNow()
	}
}
