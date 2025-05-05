package main

import (
	exp "dzaytsev/golang-graphs/algo/ford_fulkerson"
	"fmt"
)

func main() {
	exp.FordFulkersonIrrationalNetworkExperiment()
	fmt.Println()
	exp.EdmondsKarpIrrationalNetworkExperiment()
	fmt.Println()

	exp.FordFulkersonHierarchicalNetworkExperiment()
	fmt.Println()
	exp.EdmondsKarpHierarchicalNetworkExperiment()
	fmt.Println()
	exp.DinicHierarchicalNetworkExperiment()
	fmt.Println()
	exp.CapacityScalingHierarchicalNetworkExperiment()

}
