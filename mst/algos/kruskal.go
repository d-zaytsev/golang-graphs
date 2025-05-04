package mst

import (
	"sort"

	"github.com/Salvatore112/graph_analysis_algorithms/graphs"
)

func KruskalMST(g *graphs.WeightedGraph) (mst *graphs.WeightedGraph) {
	mst = graphs.NewWeightedGraph()
	vertexToID := make(map[string]int)
	id := 0
	for v := range g.Vertices {
		vertexToID[v] = id
		id++
	}
	dsu := NewDSU(len(g.Vertices))
	edges := getSortedEdges(g)
	edgesAdded := 0
	verticesCount := len(g.Vertices)
	for _, edge := range edges {
		uID := vertexToID[edge.U]
		vID := vertexToID[edge.V]
		if dsu.Find(uID) != dsu.Find(vID) {
			mst.AddEdge(edge.U, edge.V, edge.Weight)
			dsu.Union(uID, vID)
			edgesAdded++
			if edgesAdded == verticesCount-1 {
				break
			}
		}
	}
	return mst
}

func getSortedEdges(g *graphs.WeightedGraph) []graphs.WeightedEdge {
	edges := g.GetEdges()
	sort.Slice(edges, func(i, j int) bool {
		w1, w2 := edges[i].Weight, edges[j].Weight

		if w1 != w2 {
			return w1 < w2
		} else {
			return edges[i].U < edges[j].U
		}
	})
	return edges
}
