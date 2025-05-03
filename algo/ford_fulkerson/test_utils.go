package algo

import (
	g "dzaytsev/golang-graphs/graphs"
)

func buildTaskData1() (*NetworkTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 10})
	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 5})
	network.AddEdge(1, 2, g.FlowNetworkEdge[float64]{Capacity: 15})
	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 10})
	network.AddEdge(3, 1, g.FlowNetworkEdge[float64]{Capacity: 20})

	test_data, err := MakeNetworkTaskData(network, 0, 3)

	if err != nil {
		return nil, err
	}

	return test_data, nil

}

func buildTaskData2() (*NetworkTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 10})
	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 5})
	network.AddEdge(1, 2, g.FlowNetworkEdge[float64]{Capacity: 15})
	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 10})
	network.AddEdge(3, 1, g.FlowNetworkEdge[float64]{Capacity: 20})

	test_data, err := MakeNetworkTaskData(network, 0, 3)

	if err != nil {
		return nil, err
	}

	return test_data, nil

}

func buildTaskData3() (*NetworkTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 1000})
	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 1000})

	network.AddEdge(1, 2, g.FlowNetworkEdge[float64]{Capacity: 1})
	network.AddEdge(1, 3, g.FlowNetworkEdge[float64]{Capacity: 1000})

	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 1000})

	test_data, err := MakeNetworkTaskData(network, 0, 3)

	if err != nil {
		return nil, err
	}

	return test_data, nil

}
