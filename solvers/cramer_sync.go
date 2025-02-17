package solvers

import (
	"errors"
	"math"
	"sle_solver/utils"
)

type CramerMethodSync struct{}

func (c *CramerMethodSync) Solve(matrix utils.Matrix, constants utils.Matrix) (utils.Matrix, error) {
	rows, _ := matrix.Size()
	detA := utils.Det(matrix)
	if math.Abs(detA) < 1e-10 {
		return nil, errors.New("matrix is singular and the system is not solvable")
	}

	solution := utils.NewDenseMatrix(rows, 1)
	for i := 0; i < rows; i++ {
		modifiedMatrix := matrix.Copy()
		for j := 0; j < rows; j++ {
			modifiedMatrix.Set(j, i, constants.Get(j, 0))
		}
		detModified := utils.Det(modifiedMatrix)
		solution.Set(i, 0, detModified/detA)
	}
	return solution, nil
}
