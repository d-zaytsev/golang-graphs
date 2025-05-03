package algo

import (
	g "dzaytsev/golang-graphs/graphs"
	"fmt"
	"sort"
	"strings"
)

func (data *NetworkTaskData) PrintNetwork() string {
	var builder strings.Builder
	builder.WriteString("Flow Network:\n")

	vertices := make([]g.FlowNetworkVertex, 0, len(data.g.Vertices))

	for v := range data.g.Vertices {
		vertices = append(vertices, v)
	}

	sort.Slice(vertices, func(i, j int) bool {
		return vertices[i] < vertices[j]
	})

	for _, u := range vertices {
		neighbors := data.g.Vertices[u]
		if len(neighbors) == 0 {
			continue
		}

		builder.WriteString(fmt.Sprintf("Vertex %d → ", u))

		neighborList := make([]g.FlowNetworkVertex, 0, len(neighbors))
		for v := range neighbors {
			neighborList = append(neighborList, v)
		}
		sort.Slice(neighborList, func(i, j int) bool {
			return neighborList[i] < neighborList[j]
		})

		for i, v := range neighborList {
			edge := neighbors[v]
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(fmt.Sprintf("%d [%.1f/%.1f]",
				v, edge.Flow, edge.Capacity))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func PrintPath(path []g.FlowNetworkVertex) string {
	if len(path) == 0 {
		return "Empty path"
	}

	var builder strings.Builder

	for i, vertex := range path {
		if i > 0 {
			builder.WriteString(" → ")
		}
		builder.WriteString(fmt.Sprintf("%d", vertex))
	}

	return builder.String()
}
