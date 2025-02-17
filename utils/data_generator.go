package utils

import (
	"math/rand"
	"time"
)

func GenerateData(n int, dataType string) (Matrix, Matrix) {
	rand.Seed(time.Now().UnixNano())
	matrix := NewDenseMatrix(n, n)
	constants := NewDenseMatrix(n, 1)

	switch dataType {
	case "random":
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix.Set(i, j, rand.Float64()*10-5)
			}
			constants.Set(i, 0, rand.Float64()*10-5)
		}
	case "diagonal-dominant":
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					matrix.Set(i, j, rand.Float64()*10+10)
				} else {
					matrix.Set(i, j, rand.Float64()*5-2.5)
				}
			}
			constants.Set(i, 0, rand.Float64()*10-5)
		}
	case "sparse":
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if rand.Float64() < 0.2 {
					matrix.Set(i, j, rand.Float64()*10-5)
				} else {
					matrix.Set(i, j, 0)
				}
			}
			constants.Set(i, 0, rand.Float64()*10-5)
		}
	default:
		panic("Unsupported data type")
	}
	return matrix, constants
}
