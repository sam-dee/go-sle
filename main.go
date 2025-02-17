package main

import (
	"fmt"
)

func main() {
	//sizes := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 8192}
	sizes := []int{2, 4, 6, 8}
	dataTypes := []string{"random"}

	filename := "benchmarks.csv"

	RunParametrizedBenchmarks(sizes, dataTypes, filename)

	fmt.Printf("\nBenchmark results exported to %s\n", filename)
}
