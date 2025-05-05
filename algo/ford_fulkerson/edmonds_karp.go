package algo

import (
	g "dzaytsev/golang-graphs/graphs"
	"slices"
)

func (data *MaxFlowTaskData) EdmondsKarp() (float64, error) {
	for {

		path, res_code := data.resudialNetworkBFS()

		// Can find path
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
			break
		}
	}

	neighbors, _ := data.g.GetNeighbors(data.s)
	res := 0.0

	for node := range neighbors {
		res += data.GetFlow(data.s, node)
	}

	return res, nil
}

func (data *MaxFlowTaskData) resudialNetworkBFS() ([]g.FlowNetworkVertex, bool) {
	queue := []g.FlowNetworkVertex{data.s}

	// parent of each node
	visited := make(map[g.FlowNetworkVertex]*g.FlowNetworkVertex)
	visited[data.s] = nil

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		neighbours, _ := data.g.GetNeighbors(u)

		for v := range neighbours {
			_, exist := visited[v]

			if exist {
				// skip visited nodes
				continue
			}

			c_f := data.getResidualEdgeCapacity(u, v)

			if c_f <= 0 {
				continue
			}

			visited[v] = &u
			queue = append(queue, v)

			if v == data.t {
				result := visitedToPath(v, visited, make([]g.FlowNetworkVertex, 0))
				slices.Reverse(result)

				return result, true

			}
		}
	}

	return nil, false
}

func visitedToPath(cur g.FlowNetworkVertex, visited map[g.FlowNetworkVertex]*g.FlowNetworkVertex, path []g.FlowNetworkVertex) []g.FlowNetworkVertex {
	next, exist := visited[cur]

	new_path := append(path, cur)

	if !exist || next == nil {
		return new_path
	} else {
		return visitedToPath(*next, visited, new_path)
	}
}
