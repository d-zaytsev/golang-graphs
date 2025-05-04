package eclParser

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"

	"github.com/Salvatore112/graph_analysis_algorithms/graphs"
)

func ReadECLgraph(filename string) (*graphs.WeightedGraph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var nodes, edges int32

	if err := binary.Read(file, binary.LittleEndian, &nodes); err != nil {
		return nil, fmt.Errorf("failed to read nodes: %v", err)
	}
	if nodes < 1 {
		return nil, fmt.Errorf("invalid number of nodes: %d", nodes)
	}

	if err := binary.Read(file, binary.LittleEndian, &edges); err != nil {
		return nil, fmt.Errorf("failed to read edges: %v", err)
	}
	if edges < 0 {
		return nil, fmt.Errorf("invalid number of edges: %d", edges)
	}

	nindex := make([]int32, nodes+1)
	if err := binary.Read(file, binary.LittleEndian, nindex); err != nil {
		return nil, fmt.Errorf("failed to read nindex: %v", err)
	}

	if nindex[0] != 0 || nindex[nodes] != edges {
		return nil, fmt.Errorf("invalid index array structure")
	}

	nlist := make([]int32, edges)
	if err := binary.Read(file, binary.LittleEndian, nlist); err != nil {
		return nil, fmt.Errorf("failed to read nlist: %v", err)
	}

	var eweight []int
	if _, err := file.Seek(0, io.SeekCurrent); err == nil {
		eweightBuffer := make([]int32, edges)
		if err := binary.Read(file, binary.LittleEndian, eweightBuffer); err == nil {
			eweight = make([]int, edges)
			for i, w := range eweightBuffer {
				eweight[i] = int(w)
			}
		}
	}
	if len(eweight) > 0 && len(eweight) != int(edges) {
		return nil, fmt.Errorf("not enought weights (edges are %d, but weights are %d)", edges, len(eweight))
	}

	g := graphs.NewWeightedGraph()
	rand := rand.New(rand.NewSource(64))

	for uId := int32(0); uId < nodes; uId++ {
		start := nindex[uId]
		end := nindex[uId+1]

		if start > end || end > edges {
			return nil, fmt.Errorf("invalid index range for vertex %d", uId)
		}

		u := strconv.Itoa(int(uId))
		for _, neighborId := range nlist[start:end] {
			if neighborId < 0 || neighborId >= nodes {
				return nil, fmt.Errorf("invalid neighbor index %d for vertex %d", neighborId, uId)
			}
			v := strconv.Itoa(int(neighborId))
			weight := 0
			if len(eweight) > 0 {
				weight = eweight[start]
			} else {
				weight = rand.Int() % 1000

			}
			g.AddEdge(u, v, weight)
			start++
		}
	}

	return g, nil
}
