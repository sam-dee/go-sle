package solvers

import (
	"errors"
	"math"
	"sle_solver/utils"
	"sync"
)

type GaussMethodParallel struct{}

func (g *GaussMethodParallel) Solve(matrix utils.Matrix, constants utils.Matrix) (utils.Matrix, error) {
	rows, _ := matrix.Size()
	var wg sync.WaitGroup

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

		if math.Abs(matrix.Get(i, i)) < 1e-10 {
			return nil, errors.New("system is singular or underdetermined")
		}

		// parallel elimination of column i
		wg.Add(rows - i - 1)
		for k := i + 1; k < rows; k++ {
			go func(k int) {
				defer wg.Done()
				factor := matrix.Get(k, i) / matrix.Get(i, i)
				matrix.AddScaledRow(i, k, -factor)
				constants.AddScaledRow(i, k, -factor)
			}(k)
		}
		wg.Wait()
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
