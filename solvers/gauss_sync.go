package solvers

import (
	"errors"
	"math"
	"sle_solver/utils"
)

type GaussMethodSync struct{}

func (g *GaussMethodSync) Solve(matrix utils.Matrix, constants utils.Matrix) (utils.Matrix, error) {
	rows, _ := matrix.Size()

	for i := 0; i < rows; i++ {
		// partial pivoting
		maxRow := i
		for k := i + 1; k < rows; k++ {
			if math.Abs(matrix.Get(k, i)) > math.Abs(matrix.Get(maxRow, i)) {
				maxRow = k
			}
		}
		matrix.SwapRows(i, maxRow)
		constants.SwapRows(i, maxRow)

		// check for zero pivot
		if math.Abs(matrix.Get(i, i)) < 1e-10 {
			return nil, errors.New("system is singular or underdetermined")
		}

		// Eliminate column i
		for k := i + 1; k < rows; k++ {
			factor := matrix.Get(k, i) / matrix.Get(i, i)
			matrix.AddScaledRow(i, k, -factor)
			constants.AddScaledRow(i, k, -factor)
		}
	}

	// back substitution
	solution := utils.NewDenseMatrix(rows, 1)
	for i := rows - 1; i >= 0; i-- {
		sum := constants.Get(i, 0)
		for j := i + 1; j < rows; j++ {
			sum -= matrix.Get(i, j) * solution.Get(j, 0)
		}
		if math.Abs(matrix.Get(i, i)) < 1e-10 {
			return nil, errors.New("system is singular or underdetermined")
		}
		solution.Set(i, 0, sum/matrix.Get(i, i))
	}
	return solution, nil
}
