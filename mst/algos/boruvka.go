package mst

import (
	"github.com/Salvatore112/graph_analysis_algorithms/graphs"
)

const NO_CC = -1

func BoruvkaMST(g *graphs.WeightedGraph) (mst *graphs.WeightedGraph) {
	mst = graphs.NewWeightedGraph()
	vertexToID := make(map[string]int)
	id := 0
	for v := range g.Vertices {
		vertexToID[v] = id
		id++
	}

	dsu := NewDSU(len(g.Vertices))
	edgesSet := make(map[graphs.WeightedEdge]struct{})
	for _, edge := range g.GetEdges() {
		edgesSet[edge] = struct{}{}
	}
	for {
		cheapest := make(map[int]graphs.WeightedEdge)
		for edge := range edgesSet {
			u := edge.U
			v := edge.V
			weight := edge.Weight
			rootU := dsu.Find(vertexToID[u])
			rootV := dsu.Find(vertexToID[v])
			if rootU != rootV {
				if chepeastEdge, exist := cheapest[rootU]; !exist || weight < chepeastEdge.Weight {
					cheapest[rootU] = edge
				}
				if chepeastEdge, exist := cheapest[rootV]; !exist || weight < chepeastEdge.Weight {
					cheapest[rootV] = edge
				}
			}
		}
		for _, edge := range cheapest {
			u := edge.U
			v := edge.V
			uId := vertexToID[u]
			vId := vertexToID[v]
			weight := edge.Weight
			rootU := dsu.Find(uId)
			rootV := dsu.Find(vId)
			if rootU != rootV {
				mst.AddEdge(u, v, weight)
				dsu.Union(uId, rootV)
			}
			delete(edgesSet, edge)
		}
		for edge := range edgesSet {
			u := edge.U
			v := edge.V
			uId := vertexToID[u]
			vId := vertexToID[v]
			rootU := dsu.Find(uId)
			rootV := dsu.Find(vId)
			if rootU == rootV {
				delete(edgesSet, edge)
			}
		}

		if len(cheapest) == 0 {
			break
		}
	}
	return mst
}
