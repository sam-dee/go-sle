package solvers

import "sle_solver/utils"

type LinearEquationSolver interface {
	Solve(matrix utils.Matrix, constants utils.Matrix) (utils.Matrix, error)
}
