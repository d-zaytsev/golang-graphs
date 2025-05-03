package algo

import (
	g "dzaytsev/golang-graphs/graphs"
)

type MaxFlowTestCase struct {
	name         string
	builder      func() (*MaxFlowTaskData, error)
	expectedFlow float64
}

func buildTestCases() []MaxFlowTestCase {
	tests := []MaxFlowTestCase{
		{"Simple Acyclic Network", buildSimpleTaskData, 10},
		{"Cyclic Network", buildCyclicTaskData, 10},
		{"Wiki Network", buildWikiTaskData, 2000},
		{"No Path from Source to Sink", buildNoPathTaskData, 0},
		{"Single Path, Single Bottleneck", buildSinglePathTaskData, 1},
		{"Parallel Paths with Different Capacities", buildParallelPathsTaskData, 15},
		{"Parallel Paths with Different Capacities #2", buildParallelPathsTaskData2, 10},
		{"All Edges Have the Same Capacity", buildUniformCapacityTaskData, 2},
	}

	return tests
}

func buildSimpleTaskData() (*MaxFlowTaskData, error) {
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

func buildCyclicTaskData() (*MaxFlowTaskData, error) {
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

func buildWikiTaskData() (*MaxFlowTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 1000})
	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 1000})

	network.AddEdge(1, 2, g.FlowNetworkEdge[float64]{Capacity: 1})
	network.AddEdge(1, 3, g.FlowNetworkEdge[float64]{Capacity: 1000})

	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 1000})

	return MakeNetworkTaskData(network, 0, 3)
}

func buildNoPathTaskData() (*MaxFlowTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 5})
	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 10})

	return MakeNetworkTaskData(network, 0, 3)
}

func buildSinglePathTaskData() (*MaxFlowTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 5; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 100})
	network.AddEdge(1, 2, g.FlowNetworkEdge[float64]{Capacity: 100})
	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 1})
	network.AddEdge(3, 4, g.FlowNetworkEdge[float64]{Capacity: 100})
	network.AddEdge(4, 5, g.FlowNetworkEdge[float64]{Capacity: 100})

	return MakeNetworkTaskData(network, 0, 5)
}

func buildParallelPathsTaskData() (*MaxFlowTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 5})
	network.AddEdge(1, 3, g.FlowNetworkEdge[float64]{Capacity: 5})

	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 10})
	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 10})

	return MakeNetworkTaskData(network, 0, 3)
}

func buildParallelPathsTaskData2() (*MaxFlowTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 3; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 5})
	network.AddEdge(1, 3, g.FlowNetworkEdge[float64]{Capacity: 5})

	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 5})
	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 5})

	return MakeNetworkTaskData(network, 0, 3)

}

func buildUniformCapacityTaskData() (*MaxFlowTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 5; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	cap := g.FlowNetworkEdge[float64]{Capacity: 1}
	network.AddEdge(0, 1, cap)
	network.AddEdge(0, 2, cap)
	network.AddEdge(1, 3, cap)
	network.AddEdge(2, 3, cap)
	network.AddEdge(2, 4, cap)
	network.AddEdge(3, 5, cap)
	network.AddEdge(4, 5, cap)

	return MakeNetworkTaskData(network, 0, 5)
}
