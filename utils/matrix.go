package utils

type Matrix interface {
	Get(i, j int) float64
	Set(i, j int, value float64)
	Size() (int, int)
	Copy() Matrix
	SwapRows(i, j int)
	AddScaledRow(src, dest int, factor float64) // add scaled row src to dest
}

type DenseMatrix struct {
	data [][]float64
	rows int
	cols int
}

func NewDenseMatrix(rows, cols int) *DenseMatrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}
	return &DenseMatrix{data: data, rows: rows, cols: cols}
}

func (m *DenseMatrix) Get(i, j int) float64 {
	return m.data[i][j]
}

func (m *DenseMatrix) Set(i, j int, value float64) {
	m.data[i][j] = value
}

func (m *DenseMatrix) Size() (int, int) {
	return m.rows, m.cols
}

func (m *DenseMatrix) Copy() Matrix {
	copyData := make([][]float64, m.rows)
	for i := range m.data {
		copyData[i] = make([]float64, m.cols)
		copy(copyData[i], m.data[i])
	}
	return &DenseMatrix{data: copyData, rows: m.rows, cols: m.cols}
}

func (m *DenseMatrix) SwapRows(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *DenseMatrix) AddScaledRow(src, dest int, factor float64) {
	for j := range m.data[dest] {
		m.data[dest][j] += factor * m.data[src][j]
	}
}

func Det(matrix Matrix) float64 {
	rows, cols := matrix.Size()
	if rows != cols {
		panic("Det is only defined for square matrices")
	}

	if rows == 1 {
		return matrix.Get(0, 0)
	}

	if rows == 2 {
		return matrix.Get(0, 0)*matrix.Get(1, 1) - matrix.Get(0, 1)*matrix.Get(1, 0)
	}

	// expand along the first row
	det := 0.0
	for col := 0; col < cols; col++ {
		// create the submatrix by excluding the first row and current column
		subMatrix := NewDenseMatrix(rows-1, cols-1)
		subRow := 0
		for i := 1; i < rows; i++ { // skip the first row
			subCol := 0
			for j := 0; j < cols; j++ {
				if j == col { // skip the current column
					continue
				}
				subMatrix.Set(subRow, subCol, matrix.Get(i, j))
				subCol++
			}
			subRow++
		}

		sign := 1.0
		if col%2 != 0 {
			sign = -1.0
		}
		det += sign * matrix.Get(0, col) * Det(subMatrix)
	}

	return det
}
