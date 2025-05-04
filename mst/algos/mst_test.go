package mst

import (
	"maps"
	"testing"

	"github.com/Salvatore112/graph_analysis_algorithms/graphs"
)

var edges map[graphs.WeightedEdge]struct{}
var edgesExpected1 map[graphs.WeightedEdge]struct{}
var edgesExpected2 map[graphs.WeightedEdge]struct{}

func init() {
	edges = map[graphs.WeightedEdge]struct{}{
		{U: "1", V: "2", Weight: 10}: {},
		{U: "1", V: "5", Weight: 14}: {},
		{U: "2", V: "3", Weight: 17}: {},
		{U: "2", V: "4", Weight: 15}: {},
		{U: "2", V: "5", Weight: 13}: {},
		{U: "3", V: "4", Weight: 19}: {},
		{U: "3", V: "7", Weight: 15}: {},
		{U: "4", V: "5", Weight: 15}: {},
		{U: "4", V: "7", Weight: 16}: {},
		{U: "5", V: "6", Weight: 20}: {},
		{U: "6", V: "7", Weight: 19}: {},
	}

	edgesExpected1 = map[graphs.WeightedEdge]struct{}{
		{U: "1", V: "2", Weight: 10}: {},
		{U: "2", V: "4", Weight: 15}: {},
		{U: "2", V: "5", Weight: 13}: {},
		{U: "3", V: "7", Weight: 15}: {},
		{U: "4", V: "7", Weight: 16}: {},
		{U: "6", V: "7", Weight: 19}: {},
	}
	
	edgesExpected2 = map[graphs.WeightedEdge]struct{}{
		{U: "1", V: "2", Weight: 10}: {},
		{U: "2", V: "5", Weight: 13}: {},
		{U: "3", V: "7", Weight: 15}: {},
		{U: "4", V: "7", Weight: 16}: {},
		{U: "4", V: "5", Weight: 15}: {},
		{U: "6", V: "7", Weight: 19}: {},
	}
}


func TestMST(t *testing.T) {
	type args struct {
		edges        map[graphs.WeightedEdge]struct{}
		mstAlgorithm MSTAlogorithm
	}
	tests := []struct {
		name          string
		args          args
		edgesExpected [](map[graphs.WeightedEdge]struct{})
	}{
		{
			name:          "kruskal_test1",
			args:          args{edges, KruskalMST},
			edgesExpected: [](map[graphs.WeightedEdge]struct{}){edgesExpected1, edgesExpected2},
		},
		{
			name:          "prim_test1",
			args:          args{edges, PrimMST},
			edgesExpected: [](map[graphs.WeightedEdge]struct{}){edgesExpected1, edgesExpected2},
		},
		{
			name:          "boruvka_test1",
			args:          args{edges, BoruvkaMST},
			edgesExpected: [](map[graphs.WeightedEdge]struct{}){edgesExpected1, edgesExpected2},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			graph := graphs.NewWeightedGraph()
			for edge := range edges {
				graph.AddEdge(edge.U, edge.V, edge.Weight)
			}
			mst := tt.args.mstAlgorithm(graph)
			resultEdges := make(map[graphs.WeightedEdge]struct{})
			for _, e := range mst.GetEdges() {
				resultEdges[e] = struct{}{}
			}
			res := false
			for _, edgesExpected := range tt.edgesExpected {
				if maps.Equal(resultEdges, edgesExpected) {
					res = true
					break
				}
			}
			if !res {
				t.Errorf("Expected edges %v, got %v", tt.edgesExpected, resultEdges)
			}
		})
	}
}
