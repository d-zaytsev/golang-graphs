package algo

import (
	"bufio"
	g "dzaytsev/golang-graphs/graphs"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Statistics struct {
	mean    float64
	stdDev  float64
	ciLower float64
	ciUpper float64
	median  float64
	min     float64
	max     float64
}

type BuilderFunc func() (*MaxFlowTaskData, error)
type AlgorithmFunc func(*MaxFlowTaskData) (float64, error)

const (
	RUN_COUNT      = 100
	CONFIDENCE     = 0.95
	DIMACS_DATASET = "???.max"
)

func FordFulkersonIrrationalNetworkExperiment() {
	runExperiment("Ford-Fulkerson on Irrational Network", buildIrrationalNetwork,
		func(data *MaxFlowTaskData) (float64, error) {
			return data.FordFulkerson()
		})
}

func EdmondsKarpIrrationalNetworkExperiment() {
	runExperiment("EdmondsKarp on Irrational Network", buildIrrationalNetwork,
		func(data *MaxFlowTaskData) (float64, error) {
			return data.EdmondsKarp()
		})
}

func FordFulkersonHierarchicalNetworksExperiment() {
	// Сравнение на множестве разных по размеру сетей

	for layers_count := 5; layers_count < 20; layers_count += 5 {
		for nodes_count := 10; nodes_count < 100; nodes_count += 20 {
			name := fmt.Sprintf("Ford-Fulkerson on Hierarchical Network (layers=%d, nodes=%d)", layers_count, nodes_count)

			runExperiment(name,
				func() (*MaxFlowTaskData, error) {
					return buildHierarchicalNetwork(layers_count, nodes_count)
				},
				func(data *MaxFlowTaskData) (float64, error) {
					return data.FordFulkerson()
				})
		}
	}
}

func FordFulkersonHierarchicalNetworkExperiment() {
	runExperiment("Ford-Fulkerson on Hierarchical Network",
		func() (*MaxFlowTaskData, error) {
			return buildHierarchicalNetwork(10, 25)
		},
		func(data *MaxFlowTaskData) (float64, error) {
			return data.FordFulkerson()
		})
}

func EdmondsKarpHierarchicalNetworkExperiment() {
	runExperiment("EdmondsKarp on Hierarchical Network",
		func() (*MaxFlowTaskData, error) {
			return buildHierarchicalNetwork(10, 25)
		},
		func(data *MaxFlowTaskData) (float64, error) {
			return data.EdmondsKarp()
		})
}

func CapacityScalingHierarchicalNetworkExperiment() {
	runExperiment("Ford-Fulkerson(Capacity Scaling) on Hierarchical Network",
		func() (*MaxFlowTaskData, error) {
			return buildHierarchicalNetwork(10, 25)
		},
		func(data *MaxFlowTaskData) (float64, error) {
			return data.CapacityScalingFordFulkerson()
		})

	runExperiment("EdmondsKarp(Capacity Scaling) on Hierarchical Network",
		func() (*MaxFlowTaskData, error) {
			return buildHierarchicalNetwork(10, 25)
		},
		func(data *MaxFlowTaskData) (float64, error) {
			return data.CapacityScalingEdmondsKarp()
		})
}

func DinicHierarchicalNetworkExperiment() {
	runExperiment("Dinic on Hierarchical Network",
		func() (*MaxFlowTaskData, error) {
			return buildHierarchicalNetwork(10, 25)
		},
		func(data *MaxFlowTaskData) (float64, error) {
			return data.Dinic()
		})
}

func runExperiment(experimentName string, builder BuilderFunc, algorithm AlgorithmFunc) error {
	fmt.Printf("Running experiment: %s...\n", experimentName)

	time_results := make([]float64, 0)

	for i := 0; i < RUN_COUNT; i++ {
		data, err := builder()
		if err != nil {
			fmt.Printf("Error building data in run %d: %v\n", i, err)
			return err
		}

		start := time.Now()
		_, err = algorithm(data)
		end := time.Since(start)

		if err != nil {
			fmt.Printf("Error in algorithm during run %d: %v\n", i, err)
			return err
		}

		time_results = append(time_results, end.Seconds())
	}

	stats := calculateStatistics(time_results, CONFIDENCE)

	fmt.Printf("Results from %d runs:\n", len(time_results))
	fmt.Printf("→ Average time: %.6f seconds\n", stats.mean)
	fmt.Printf("→ Standard deviation: %.6f seconds\n", stats.stdDev)
	fmt.Printf("→ Confidence interval (%.1f%%): [%.6f, %.6f] seconds\n",
		CONFIDENCE*100, stats.ciLower, stats.ciUpper)
	fmt.Printf("→ Median: %.6f seconds\n", stats.median)
	fmt.Printf("→ Min: %.6f seconds\n", stats.min)
	fmt.Printf("→ Max: %.6f seconds\n", stats.max)

	return nil
}

func calculateStatistics(results []float64, confidence float64) Statistics {
	n := float64(len(results))

	sum := 0.0
	for _, val := range results {
		sum += val
	}
	mean := sum / n

	varianceSum := 0.0
	for _, val := range results {
		varianceSum += math.Pow(val-mean, 2)
	}
	variance := varianceSum / (n - 1)
	stdDev := math.Sqrt(variance)

	z := 1.96
	if confidence == 0.99 {
		z = 2.576
	} else if confidence == 0.90 {
		z = 1.645
	}

	margin := z * stdDev / math.Sqrt(n)
	ciLower := mean - margin
	ciUpper := mean + margin

	sortedResults := make([]float64, len(results))
	copy(sortedResults, results)
	sort.Float64s(sortedResults)

	min := sortedResults[0]
	max := sortedResults[len(sortedResults)-1]

	var median float64
	if len(sortedResults)%2 == 0 {
		median = (sortedResults[len(sortedResults)/2-1] + sortedResults[len(sortedResults)/2]) / 2
	} else {
		median = sortedResults[len(sortedResults)/2]
	}

	return Statistics{
		mean:    mean,
		stdDev:  stdDev,
		ciLower: ciLower,
		ciUpper: ciUpper,
		median:  median,
		min:     min,
		max:     max,
	}
}

func buildIrrationalNetwork() (*MaxFlowTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	for i := 0; i <= 5; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	r := (math.Sqrt(5) - 1) / 2

	network.AddEdge(0, 1, g.FlowNetworkEdge[float64]{Capacity: 100})
	network.AddEdge(0, 2, g.FlowNetworkEdge[float64]{Capacity: 100})
	network.AddEdge(0, 4, g.FlowNetworkEdge[float64]{Capacity: 100})

	network.AddEdge(1, 3, g.FlowNetworkEdge[float64]{Capacity: 100})

	network.AddEdge(2, 5, g.FlowNetworkEdge[float64]{Capacity: r})
	network.AddEdge(2, 3, g.FlowNetworkEdge[float64]{Capacity: 100})

	network.AddEdge(4, 5, g.FlowNetworkEdge[float64]{Capacity: 1})
	network.AddEdge(4, 1, g.FlowNetworkEdge[float64]{Capacity: 1})

	network.AddEdge(5, 3, g.FlowNetworkEdge[float64]{Capacity: 100})

	test_data, err := MakeNetworkTaskData(network, 0, 3)

	if err != nil {
		return nil, err
	}

	return test_data, nil
}

func buildHierarchicalNetwork(numLayers, nodesPerLayer int) (*MaxFlowTaskData, error) {
	network := g.MakeFlowNetwork[float64]()

	totalNodes := numLayers * nodesPerLayer

	for i := 0; i < totalNodes; i++ {
		network.AddVertex(g.FlowNetworkVertex(i))
	}

	for layer := 0; layer < numLayers-1; layer++ {
		for i := 0; i < nodesPerLayer; i++ {
			from := layer*nodesPerLayer + i

			connections := rand.Intn(nodesPerLayer) + 2
			for j := 0; j < connections; j++ {
				to := (layer+1)*nodesPerLayer + rand.Intn(nodesPerLayer)
				capacity := rand.Float64()*999 + 1 // Capacity in [1, 1000]
				network.AddEdge(g.FlowNetworkVertex(from), g.FlowNetworkVertex(to), g.FlowNetworkEdge[float64]{Capacity: capacity})
			}

			if rand.Float64() < 0.2 && i < nodesPerLayer-1 {
				next := layer*nodesPerLayer + i + 1
				capacity := rand.Float64()*50 + 5
				network.AddEdge(g.FlowNetworkVertex(from), g.FlowNetworkVertex(next), g.FlowNetworkEdge[float64]{Capacity: capacity})
			}
		}
	}

	source := 0
	sink := totalNodes - 1

	test_data, err := MakeNetworkTaskData(network, g.FlowNetworkVertex(source), g.FlowNetworkVertex(sink))
	if err != nil {
		return nil, err
	}

	return test_data, nil
}

func buildDimacsNetwork() (*MaxFlowTaskData, error) {
	net, err := readDimacsMaxFlowFile(DIMACS_DATASET)

	if err != nil {
		return nil, err
	}

	return net, nil
}

func readDimacsMaxFlowFile(filePath string) (*MaxFlowTaskData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	network := g.MakeFlowNetwork[float64]()
	var sourceNode, sinkNode int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || !strings.Contains("pcna", string(line[0])) {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "p":
			if len(fields) < 4 || fields[1] != "max" {
				continue
			}
			vertexCount, _ := strconv.Atoi(fields[2])

			for i := 1; i <= vertexCount; i++ {
				network.AddVertex(g.FlowNetworkVertex(i))
			}
		case "n":
			if len(fields) < 3 {
				continue
			}
			nodeID, _ := strconv.Atoi(fields[1])
			switch fields[2] {
			case "s":
				sourceNode = nodeID
			case "t":
				sinkNode = nodeID
			}
		case "a":
			if len(fields) < 4 {
				continue
			}
			from, _ := strconv.Atoi(fields[1])
			to, _ := strconv.Atoi(fields[2])
			capacity, _ := strconv.ParseFloat(fields[3], 64)
			network.AddEdge(g.FlowNetworkVertex(from), g.FlowNetworkVertex(to), g.FlowNetworkEdge[float64]{Capacity: capacity})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if sourceNode == 0 || sinkNode == 0 {
		return nil, fmt.Errorf("source or sink node not defined")
	}

	return MakeNetworkTaskData(network, g.FlowNetworkVertex(sourceNode), g.FlowNetworkVertex(sinkNode))
}
