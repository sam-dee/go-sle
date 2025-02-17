package main

import (
	"math"
	solvers2 "sle_solver/solvers"
	"sle_solver/utils"
	"testing"
)

func TestSolvers(t *testing.T) {
	tests := []struct {
		name      string
		matrix    [][]float64
		constants []float64
		expected  []float64
	}{
		{
			name: "Case 2x2",
			matrix: [][]float64{
				{1, 2},
				{3, 4},
			},
			constants: []float64{5, 11},
			expected:  []float64{1, 2},
		},
		{
			name: "Case 3x3",
			matrix: [][]float64{
				{2, 1, -1},
				{-3, -1, 2},
				{-2, 1, 2},
			},
			constants: []float64{8, -11, -3},
			expected:  []float64{2, 3, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matrix := utils.NewDenseMatrix(len(tt.matrix), len(tt.matrix[0]))
			for i, row := range tt.matrix {
				for j, val := range row {
					matrix.Set(i, j, val)
				}
			}

			constants := utils.NewDenseMatrix(len(tt.constants), 1)
			for i, val := range tt.constants {
				constants.Set(i, 0, val)
			}

			expected := utils.NewDenseMatrix(len(tt.expected), 1)
			for i, val := range tt.expected {
				expected.Set(i, 0, val)
			}

			solvers := map[string]solvers2.LinearEquationSolver{
				"GaussMethodSync":      &solvers2.GaussMethodSync{},
				"GaussMethodParallel":  &solvers2.GaussMethodParallel{},
				"CramerMethodSync":     &solvers2.CramerMethodSync{},
				"CramerMethodParallel": &solvers2.CramerMethodParallel{},
			}

			for name, solver := range solvers {
				t.Run(name, func(t *testing.T) {
					result, err := solver.Solve(matrix, constants)
					if err != nil {
						t.Fatalf("Solver %s failed: %v", name, err)
					}

					for i := 0; i < len(tt.expected); i++ {
						if math.Abs(result.Get(i, 0)-expected.Get(i, 0)) > 1e-6 {
							t.Errorf("Solver %s produced incorrect result at index %d: got %.6f, want %.6f",
								name, i, result.Get(i, 0), expected.Get(i, 0))
						}
					}
				})
			}
		})
	}
}
