package main

import (
	"gonum.org/v1/gonum/mat"
)

func identityMatrix() *mat.Dense {
	values := []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	return mat.NewDense(4, 4, values)
}

func subMatrix(m *mat.Dense, row, col int) *mat.Dense {
	r, c := m.Dims()
	new := mat.NewDense(r-1, c-1, nil)
	i2, j2 := 0, 0
	for i := 0; i < r; i++ {
		if i == row {
			continue
		}
		j2 = 0
		for j := 0; j < c; j++ {
			if j == col {
				continue
			}
			new.Set(i2, j2, m.At(i, j))
			j2++
		}
		i2++
	}
	return new
}

func minor(m *mat.Dense, row, col int) float64 {
	z := subMatrix(m, row, col)
	return mat.Det(z)
}

func cofactor(m *mat.Dense, row, col int) float64 {
	z := minor(m, row, col)
	if (row+col)%2 == 1 {
		z *= -1
	}
	return z
}

func isInvertable(m *mat.Dense) bool {
	return mat.Det(m) != 0
}
