package algo

import (
	g "dzaytsev/golang-graphs/graphs"
	"fmt"
	"math"
)

type MaxFlowTaskData struct {
	// graph
	g *g.FlowNetwork[float64]
	// source
	s g.FlowNetworkVertex
	// sink
	t g.FlowNetworkVertex
}

func (data *MaxFlowTaskData) GetCapacity(vertex1, vertex2 g.FlowNetworkVertex) float64 {
	exists := data.g.HasEdge(vertex1, vertex2)
	edge, _ := data.g.GetEdge(vertex1, vertex2)

	if !exists {
		return 0
	} else {
		return edge.Capacity
	}
}

func (data *MaxFlowTaskData) GetFlow(vertex1, vertex2 g.FlowNetworkVertex) float64 {
	exists := data.g.HasEdge(vertex1, vertex2)
	edge, _ := data.g.GetEdge(vertex1, vertex2)

	if !exists {
		return 0
	} else {
		return edge.Flow
	}
}

func (data *MaxFlowTaskData) SetFlow(vertex1, vertex2 g.FlowNetworkVertex, value float64) {
	exists := data.g.HasEdge(vertex1, vertex2)
	edge, _ := data.g.GetEdge(vertex1, vertex2)

	if value <= 0 {
		value = 0
	}

	if !exists {
		data.g.AddEdge(vertex1, vertex2, g.FlowNetworkEdge[float64]{Flow: value})
		return
	}

	edge.Flow = value
}

func MakeNetworkTaskData(network *g.FlowNetwork[float64], s, t g.FlowNetworkVertex) (*MaxFlowTaskData, error) {
	if !network.HasVertex(s) {
		return nil, fmt.Errorf("Source doesn't exist")
	} else if !network.HasVertex(t) {
		return nil, fmt.Errorf("Target doesn't exist")
	}

	return &MaxFlowTaskData{
		g: network,
		s: s,
		t: t,
	}, nil
}

func (data *MaxFlowTaskData) FordFulkerson() (float64, error) {
	for true {

		path, res_code := data.residualNetworkDFS()

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

func (data *MaxFlowTaskData) residualNetworkDFS() ([]g.FlowNetworkVertex, bool) {
	res, res_code := dfsHelper(data.s, data.t, data, make([]g.FlowNetworkVertex, 0), make(map[g.FlowNetworkVertex]bool))

	return res, res_code
}

func (data *MaxFlowTaskData) updateFlow(path []g.FlowNetworkVertex, min_capacity float64) error {
	if len(path) == 0 {
		return fmt.Errorf("Path is empty")
	}

	for i := 1; i < len(path); i++ {
		u := path[i-1]
		v := path[i]

		// f(u,v) = f(u,v) + c_f(p)
		data.SetFlow(u, v, data.GetFlow(u, v)+min_capacity)

		// f(v,u) = f(v,u) - c_f(p)
		data.SetFlow(v, u, data.GetFlow(v, u)-min_capacity)
	}

	return nil
}

func (data *MaxFlowTaskData) getPathMinCapacity(path []g.FlowNetworkVertex) (float64, error) {
	if len(path) == 0 {
		return -1, fmt.Errorf("Path is empty")
	}

	var min_capacity = math.MaxFloat64

	for i := 1; i < len(path); i++ {
		// c_f(p) = min {c_f (u,v)}
		u := path[i-1]
		v := path[i]

		c_f := data.getResidualEdgeCapacity(u, v)

		if c_f < min_capacity {
			min_capacity = c_f
		}
	}

	return min_capacity, nil
}

func dfsHelper(u, t g.FlowNetworkVertex, data *MaxFlowTaskData, path []g.FlowNetworkVertex, visited map[g.FlowNetworkVertex]bool) ([]g.FlowNetworkVertex, bool) {
	new_path := append(path, u)

	if u == t {
		return new_path, true
	}

	if visited[u] {
		return nil, false
	} else {
		visited[u] = true
	}

	neighbors, err := data.g.GetNeighbors(u)

	if err != nil {
		return nil, false
	}

	for v := range neighbors {
		c_f := data.getResidualEdgeCapacity(u, v)

		if c_f <= 0 {
			continue
		}

		res, res_code := dfsHelper(v, t, data, new_path, visited)

		if res_code {
			return res, true
		}
	}

	return nil, false
}

func (data *MaxFlowTaskData) getResidualEdgeCapacity(u, v g.FlowNetworkVertex) float64 {
	if data.GetCapacity(u, v) > 0 {
		return data.GetCapacity(u, v) - data.GetFlow(u, v)
	} else if data.GetFlow(v, u) > 0 {
		return data.GetFlow(v, u)
	}

	return 0
}
