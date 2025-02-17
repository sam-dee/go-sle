package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sle_solver/solvers"
	"sle_solver/utils"
	"testing"
	"time"
)

func BenchmarkSolvers(b *testing.B, solver solvers.LinearEquationSolver, size int, dataType string) {
	// exclude setup time
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		matrix, constants := utils.GenerateData(size, dataType)
		_, err := solver.Solve(matrix, constants)
		if err != nil {
			b.Fatalf("Error solving: %v", err)
		}
	}
}

func RunParametrizedBenchmarks(sizes []int, dataTypes []string, filename string) {
	var solvers = map[string]solvers.LinearEquationSolver{
		"GaussMethodSync":      &solvers.GaussMethodSync{},
		"GaussMethodParallel":  &solvers.GaussMethodParallel{},
		"CramerMethodSync":     &solvers.CramerMethodSync{},
		"CramerMethodParallel": &solvers.CramerMethodParallel{},
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Solver", "DataType", "Size", "TimePerOperation"}
	writer.Write(header)

	for name, solver := range solvers {
		for _, dataType := range dataTypes {
			fmt.Printf("\nBenchmarking %s with %s data:\n", name, dataType)
			for _, size := range sizes {
				fmt.Printf("Matrix Size: %d\n", size)

				benchmark := testing.Benchmark(func(b *testing.B) {
					BenchmarkSolvers(b, solver, size, dataType)
				})

				timePerOp := time.Duration(benchmark.NsPerOp())
				row := []string{name, dataType, fmt.Sprintf("%d", size), fmt.Sprintf("%v", timePerOp.Seconds())}
				writer.Write(row)

				fmt.Printf("Time per operation: %v\n", timePerOp)
			}
		}
	}
}
