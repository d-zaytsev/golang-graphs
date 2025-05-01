package algo

import (
	g "dzaytsev/golang-graphs/graphs"
)

type NetworkTaskData[E any] struct {
	// graph
	g *g.FlowNetwork[E]
	// source
	s g.FlowNetworkVertex
	// sink
	t g.FlowNetworkVertex
}

func (data *NetworkTaskData[E]) DFS() ([]g.FlowNetworkVertex, bool) {
	res, res_code := DFS(data.s, data.t, *data.g, make([]g.FlowNetworkVertex, 0), make(map[g.FlowNetworkVertex]bool))

	return res, res_code
}

func DFS[E any](cur, t g.FlowNetworkVertex, graph g.FlowNetwork[E], path []g.FlowNetworkVertex, visited map[g.FlowNetworkVertex]bool) ([]g.FlowNetworkVertex, bool) {
	new_path := append(path, cur)

	if cur == t {
		return new_path, true
	}

	if visited[cur] {
		return nil, false
	}

	visited[cur] = true

	neighbors, err := graph.GetNeighbors(cur)

	if err != nil {
		return nil, false
	}

	for node := range neighbors {
		res, res_code := DFS(node, t, graph, new_path, visited)

		if res_code {
			return res, true
		}
	}

	return nil, false
}
