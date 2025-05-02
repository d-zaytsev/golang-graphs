package algo

import (
	g "dzaytsev/golang-graphs/graphs"
	"testing"
)

func TestAlgo1(t *testing.T) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 10})
	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 5})
	network.AddEdge(1, 2, g.FlowNetworkEdge[float64]{Capacity: 15})
	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 10})
	network.AddEdge(3, 1, g.FlowNetworkEdge[float64]{Capacity: 20})

	test_data, _ := MakeNetworkTaskData(network, 0, 3)

	res := test_data.FordFulkerson()

	if res != 10 {
		t.Errorf("Incorrect result path: %v!", res)
	}
}

func TestAlgo2(t *testing.T) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 5; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 4})
	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 7})

	network.AddEdge(1, 3, g.FlowNetworkEdge[float64]{Capacity: 4})
	network.AddEdge(1, 4, g.FlowNetworkEdge[float64]{Capacity: 8})

	network.AddEdge(2, 1, g.FlowNetworkEdge[float64]{Capacity: 4})
	network.AddEdge(2, 4, g.FlowNetworkEdge[float64]{Capacity: 2})

	network.AddEdge(3, 5, g.FlowNetworkEdge[float64]{Capacity: 12})

	network.AddEdge(4, 3, g.FlowNetworkEdge[float64]{Capacity: 4})
	network.AddEdge(4, 5, g.FlowNetworkEdge[float64]{Capacity: 5})

	test_data, _ := MakeNetworkTaskData(network, 0, 5)

	res := test_data.FordFulkerson()

	if res != 10 {
		t.Errorf("Incorrect result path: %v!", res)
	}
}

func TestAlgo3(t *testing.T) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 1000})
	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 1000})

	network.AddEdge(1, 2, g.FlowNetworkEdge[float64]{Capacity: 1})
	network.AddEdge(1, 3, g.FlowNetworkEdge[float64]{Capacity: 1000})

	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 1000})

	test_data, _ := MakeNetworkTaskData(network, 0, 3)

	res := test_data.FordFulkerson()

	if res != 2000 {
		t.Errorf("Incorrect result path: %v!", res)
	}
}
