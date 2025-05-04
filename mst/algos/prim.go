package mst

import (
	"math"

	"github.com/Salvatore112/graph_analysis_algorithms/graphs"
)

const NO_PARENT_ID_PRIM = -1
const START_VERTEX_INDEX = 0

func PrimMST(g *graphs.WeightedGraph) (mst *graphs.WeightedGraph) {
	mst = graphs.NewWeightedGraph()
	idToVertex := make([]string, len(g.Vertices))
	vertexToID := map[string]int{}
	i := 0
	key := make([]int, len(g.Vertices))
	parent := make([]int, len(g.Vertices))
	for vertex := range g.Vertices {
		idToVertex[i] = vertex
		vertexToID[vertex] = i
		key[i] = math.MaxInt
		parent[i] = NO_PARENT_ID_PRIM
		i++
	}
	key[START_VERTEX_INDEX] = 0
	q := NewPQ[int, int]()

	for i := range len(g.Vertices) {
		q.Push(Node[int, int]{i, key[i]})
	}

	for q.Len() > 0 {
		u := q.Pop().(Node[int, int])
		for vKey, weight := range g.GetNeighbors(idToVertex[u.Value]) {
			vId := vertexToID[vKey]
			if q.Contains(vId) && weight < key[vId] {
				parent[vId] = u.Value
				key[vId] = weight
				q.Update(vId, key[vId])
			}
		}
	}

	for i, parentId := range parent {
		if parentId != NO_PARENT_ID_PRIM {
			mst.AddEdge(idToVertex[i], idToVertex[parentId], key[i])
		}
	}
	return mst
}
