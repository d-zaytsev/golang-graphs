package algo

import (
	g "dzaytsev/golang-graphs/graphs"
	"math"
	"slices"
)

func (data *MaxFlowTaskData) CapacityScalingFordFulkerson() (float64, error) {
	U, err := data.findMaxGraphCapacity()

	if err != nil {
		return -1, err
	}

	delta := calcDelta(U)

	for delta > 0 {

		path, res_code := data.heuristicResidualNetworkDFS(delta)

		if res_code {
			path_capacity, err := data.getPathMinCapacity(path)

			if err != nil {
				return -1, err
			}

			err = data.updateFlow(path, path_capacity)

			if err != nil {
				return -1, err
			}
		} else {
			delta /= 2
		}
	}

	neighbors, err := data.g.GetNeighbors(data.s)

	if err != nil {
		return -1, err
	}

	res := 0.0

	for node := range neighbors {
		res += data.GetFlow(data.s, node)
	}

	return res, nil
}

func (data *MaxFlowTaskData) heuristicResidualNetworkDFS(delta float64) ([]g.FlowNetworkVertex, bool) {
	stack := make([]g.FlowNetworkVertex, 0)
	visited := make(map[g.FlowNetworkVertex]*g.FlowNetworkVertex)

	stack = append(stack, data.s)
	visited[data.s] = nil

	for len(stack) > 0 {
		n := len(stack) - 1

		u := stack[n]
		stack = stack[:n]

		neighbours, _ := data.g.GetNeighbors(u)

		for v := range neighbours {
			_, exist := visited[v]

			if exist {
				continue
			}

			c_f := data.getResidualEdgeCapacity(u, v)

			if c_f <= 0 || c_f < delta {
				continue
			}

			visited[v] = &u
			stack = append(stack, v)

			if v == data.t {
				new_path := visitedToPath(v, visited, make([]g.FlowNetworkVertex, 0))
				slices.Reverse(new_path)

				return new_path, true
			}
		}

	}

	return nil, false

}

func calcDelta(U float64) float64 {
	return math.Pow(2, math.Floor(math.Log2(U)))
}

func (data *MaxFlowTaskData) findMaxGraphCapacity() (float64, error) {
	max_capacity := -1.0

	stack := make([]g.FlowNetworkVertex, 0)
	visited := make(map[g.FlowNetworkVertex]bool)

	stack = append(stack, data.s)
	visited[data.s] = true

	for len(stack) > 0 {
		n := len(stack) - 1

		u := stack[n]
		stack = stack[:n]

		neighbours, err := data.g.GetNeighbors(u)

		if err != nil {
			return -1, err
		}

		for v := range neighbours {
			_, exist := visited[v]

			if exist {
				continue
			}

			c_f := data.getResidualEdgeCapacity(u, v)

			if c_f <= 0 {
				continue
			} else if c_f > max_capacity {
				max_capacity = c_f
			}

			visited[v] = true
			stack = append(stack, v)
		}
	}

	return max_capacity, nil
}
