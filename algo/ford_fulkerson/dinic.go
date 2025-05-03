package algo

import (
	g "dzaytsev/golang-graphs/graphs"
	"fmt"
	"slices"
)

const INF_DIST = -2

func (data *MaxFlowTaskData) Dinic() (float64, error) {
	for true {
		t_dist, err := data.getDistance(data.t)

		if err != nil {
			return -1, err
		}

		if t_dist == INF_DIST {
			break
		}

		path, err := data.levelGraphBFS()

		if err != nil {
			return -1, err
		}

		path_capacity, err := data.getPathMinCapacity(path)

		if err != nil {
			return -1, err
		}

		data.updateFlow(path, path_capacity)
	}

	neighbors, _ := data.g.GetNeighbors(data.s)
	res := 0.0

	for node := range neighbors {
		res += data.GetFlow(data.s, node)
	}

	return res, nil
}

func (data *MaxFlowTaskData) levelGraphBFS() ([]g.FlowNetworkVertex, error) {
	// 1. Find G_l (distances)
	distances := make(map[g.FlowNetworkVertex]int)
	vertices := data.g.Vertices

	for vertex := range vertices {
		dist, err := data.getDistance(vertex)

		if err != nil {
			return nil, err
		}

		distances[vertex] = dist
	}

	// 2. Find Blocking flow (DFS)
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

			isEdgeInLevelGraph, err := data.inLevelGraph(u, v, distances)

			if err != nil {
				return nil, err
			}

			if !isEdgeInLevelGraph {
				fmt.Println("edge is not in level graph")
				continue
			}

			visited[v] = &u
			stack = append(stack, v)

			if v == data.t {
				new_path := visitedToPath(v, visited, make([]g.FlowNetworkVertex, 0))
				slices.Reverse(new_path)

				return new_path, nil
			}
		}

	}

	return nil, fmt.Errorf("There is no blocking flow!")
}

func (data *MaxFlowTaskData) inLevelGraph(u, v g.FlowNetworkVertex, distances map[g.FlowNetworkVertex]int) (bool, error) {
	u_dist, u_exist := distances[u]
	v_dist, v_exist := distances[v]

	if !u_exist || !v_exist {
		return false, fmt.Errorf("Edge (%v,%v) doesn't exist", u, v)
	}

	return v_dist == u_dist+1, nil
}

func (data *MaxFlowTaskData) getDistance(v g.FlowNetworkVertex) (int, error) {
	if !data.g.HasVertex(v) {
		return -1, fmt.Errorf("Vertex %v doesn't exist", v)
	} else if data.s == v {
		return 0, nil
	}

	queue := []g.FlowNetworkVertex{data.s}

	visited := make(map[g.FlowNetworkVertex]int)
	visited[data.s] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		neighbours, _ := data.g.GetNeighbors(u)

		for t := range neighbours {
			_, exist := visited[t]

			if exist {
				continue
			}

			c_f := data.getResidualEdgeCapacity(u, t)

			if c_f <= 0 {
				continue
			}

			visited[t] = visited[u] + 1
			queue = append(queue, t)

			if t == v {
				return visited[t], nil
			}
		}

	}

	return INF_DIST, nil
}
