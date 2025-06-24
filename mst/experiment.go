package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Salvatore112/graph_analysis_algorithms/graphs"
	mst "github.com/Salvatore112/graph_analysis_algorithms/mst/algos"
	"github.com/Salvatore112/graph_analysis_algorithms/mst/eclParser"

	"github.com/montanaflynn/stats"
)

const (
	GRAPHS_DIR    = "experiment-graphs"
	N_EXPERIMENTS = 10
)

func GetFilesInDirectory(directory string) ([]string, error) {
	if directory == "" {
		directory = "."
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

func measureExecutionTime(mstAlgorithm mst.MSTAlogorithm, g *graphs.WeightedGraph) float64 {
	startTime := time.Now()
	mstAlgorithm(g)
	elapsedTime := time.Since(startTime).Seconds()
	return elapsedTime
}

func runExperiment(g *graphs.WeightedGraph, graphName string) {
	vertexCount := len(g.Vertices)
	edgeCount := len(g.GetEdges())

	algorithmFunctions := map[string]mst.MSTAlogorithm{
		"PrimMST":    mst.PrimMST,
		"BoruvkaMST": mst.BoruvkaMST,
		"KruskalMST": mst.KruskalMST,
	}
	executionTimes := make([]float64, N_EXPERIMENTS)

	for name, algorithm := range algorithmFunctions {
		for i := range N_EXPERIMENTS {
			executionTimes[i] = measureExecutionTime(algorithm, g)
		}
		mean, _ := stats.Mean(executionTimes)
		sd, _ := stats.StandardDeviation(executionTimes)
		fmt.Printf("%s,%d,%d,%s,%.2f,%.2f\n", graphName, vertexCount, edgeCount, name, mean, sd)
	}

}

func main() {
	fileNames, err := GetFilesInDirectory(GRAPHS_DIR)

	if err != nil {
		log.Fatalf("Error getting file list: %v", err)
	}
	fmt.Printf("Graph,Vertices,Edges,Algorithm,mean(s),s.d.\n")
	for _, fileName := range fileNames {
		filePath := GRAPHS_DIR + "/" + fileName
		graph, err := eclParser.ReadECLgraph(filePath)
		if err != nil {
			log.Printf("Error reading graph from file %s: %v\n", filePath, err)
			continue
		}

		runExperiment(graph, fileName)
	}
}
