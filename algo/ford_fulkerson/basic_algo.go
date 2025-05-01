package algo

import (
	g "dzaytsev/golang-graphs/graphs"
	"fmt"
	"math"
)

type NetworkTaskData struct {
	// graph
	g *g.FlowNetwork[float64]
	// source
	s g.FlowNetworkVertex
	// sink
	t g.FlowNetworkVertex
}

func (data NetworkTaskData) DFS() ([]g.FlowNetworkVertex, bool) {
	res, res_code := DFS(data.s, data.t, *data.g, make([]g.FlowNetworkVertex, 0), make(map[g.FlowNetworkVertex]bool))

	return res, res_code
}

func (data NetworkTaskData) FordFulkerson() float64 {
	// result
	path_capacity_sum := 0.0

	for true {
		path, res_code := data.DFS()

		// Can continue
		if res_code {
			path_capacity, _ := data.GetPathCapacity(path)
			path_capacity_sum += path_capacity
			data.UpdatePathCapacity(path, path_capacity)
		} else {
			break
		}
	}

	return path_capacity_sum
}

func (data *NetworkTaskData) UpdatePathCapacity(path []g.FlowNetworkVertex, capacity float64) (float64, error) {
	if len(path) == 0 {
		return -1, fmt.Errorf("Path is empty")
	}

	node_before := path[0]

	for i := 1; i < len(path); i++ {
		// Back edges
		back_edge, back_exist := data.g.GetEdge(path[i], node_before)

		if !back_exist {
			// create new back edge
			data.g.AddEdge(path[i], node_before, g.FlowNetworkEdge[float64]{Capacity: capacity, Flow: 0})
		} else {
			// update capacity
			back_edge.Capacity += capacity
		}

		// Front edges
		front_edge, front_exist := data.g.GetEdge(node_before, path[i])

		if !front_exist {
			data.g.AddEdge(node_before, path[i], g.FlowNetworkEdge[float64]{Capacity: capacity, Flow: 0})
		} else {
			// update capacity
			front_edge.Capacity -= capacity

			// remove 0 capacity nodes
			if front_edge.Capacity <= 0 {
				data.g.RemoveEdge(node_before, path[i])
			}
		}

		node_before = path[i]
	}

	return capacity, nil
}

func (data NetworkTaskData) GetPathCapacity(path []g.FlowNetworkVertex) (float64, error) {
	if len(path) == 0 {
		return -1, fmt.Errorf("Path is empty")
	}

	node_before := path[0]

	var min_capacity = math.MaxFloat64

	for i := 1; i < len(path); i++ {
		edge, exist := data.g.GetEdge(node_before, path[i])

		if !exist {
			return -1, fmt.Errorf("Can't find edge from vertex '%v' to vertex '%v'.", path[0], path[1])
		}

		if edge.Capacity < min_capacity {
			min_capacity = edge.Capacity
		}
		node_before = path[i]
	}

	return min_capacity, nil
}

func DFS(cur, t g.FlowNetworkVertex, graph g.FlowNetwork[float64], path []g.FlowNetworkVertex, visited map[g.FlowNetworkVertex]bool) ([]g.FlowNetworkVertex, bool) {
	new_path := append(path, cur)

	if cur == t {
		return new_path, true
	}

	if visited[cur] {
		return nil, false
	} else {
		visited[cur] = true
	}

	neighbors, err := graph.GetNeighbors(cur)

	if err != nil {
		return nil, false
	}

	for _, node := range neighbors {
		res, res_code := DFS(node, t, graph, new_path, visited)

		if res_code {
			return res, true
		}
	}

	return nil, false
}
