package algo

import (
	g "dzaytsev/golang-graphs/graphs"
	"slices"
	"testing"
)

func CreateTestFloatNetwork() *NetworkTaskData {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 10, Flow: 0})
	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 5, Flow: 0})
	network.AddEdge(1, 2, g.FlowNetworkEdge[float64]{Capacity: 15, Flow: 0})
	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 10, Flow: 0})
	network.AddEdge(3, 1, g.FlowNetworkEdge[float64]{Capacity: 20, Flow: 0})

	return &NetworkTaskData{
		g: network,
		s: 0,
		t: 3,
	}
}

func TestDFS(t *testing.T) {
	test_network := CreateTestFloatNetwork()

	res, res_code := test_network.DFS()

	if !res_code {
		t.Errorf("Incorrect result code!")
	}

	correct_path := []g.FlowNetworkVertex{0, 1, 2, 3}

	if slices.Compare(res, correct_path) != 0 {
		t.Errorf("Incorrect result path!")
	}
}

func TestFordFulkersonAlgo(t *testing.T) {
	test_network := CreateTestFloatNetwork()

	res := test_network.FordFulkerson()

	if res != 10 {
		t.Errorf("Incorrect result path: %v!", res)
	}
}
