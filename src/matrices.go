package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Matrix struct {
	Rows  int
	Cols  int
	Cells map[int]map[int]float64
}

func NewMatrix(width, height int) Matrix {
	m := Matrix{
		Rows: width,
		Cols: height,
	}
	m.Cells = make(map[int]map[int]float64)
	for i := 0; i < width; i++ {
		m.Cells[i] = make(map[int]float64)
		for j := 0; j < height; j++ {
			m.Cells[i][j] = 0
		}
	}
	return m
}

func BuildMatrixFromString(stringmatrix string) Matrix {
	m := Matrix{}
	allRows := strings.Split(stringmatrix, "\n")
	m.Cols = len(allRows)
	aRow := strings.Split(strings.Trim(allRows[0], " |"), "|")
	m.Rows = len(aRow)
	m.Cells = make(map[int]map[int]float64)
	for x, row := range allRows {
		thisRow := strings.Split(strings.Trim(row, " |"), "|")
		m.Cells[x] = make(map[int]float64)
		for y, value := range thisRow {
			val, _ := strconv.ParseFloat(value, 64)
			m.Cells[x][y] = float64(val)
		}
	}
	return m
}

var IdentityMatrix = func() Matrix { return BuildMatrixFromString("|1|0|0|0|\n|0|1|0|0|\n|0|0|1|0|\n|0|0|0|1|") }

func NewTranslation(x, y, z float64) Matrix {
	newT := IdentityMatrix()
	newT.Cells[0][3] = x
	newT.Cells[1][3] = y
	newT.Cells[2][3] = z
	return newT
}

func NewScaling(x, y, z float64) Matrix {
	newT := IdentityMatrix()
	newT.Cells[0][0] = x
	newT.Cells[1][1] = y
	newT.Cells[2][2] = z
	return newT
}

func NewRotationX(radians float64) Matrix {
	newT := IdentityMatrix()
	newT.Cells[1][1] = math.Cos(radians)
	newT.Cells[1][2] = -math.Sin(radians)
	newT.Cells[2][1] = math.Sin(radians)
	newT.Cells[2][2] = math.Cos(radians)
	return newT
}

func NewRotationY(radians float64) Matrix {
	newT := IdentityMatrix()
	newT.Cells[0][0] = math.Cos(radians)
	newT.Cells[2][0] = -math.Sin(radians)
	newT.Cells[0][2] = math.Sin(radians)
	newT.Cells[2][2] = math.Cos(radians)
	return newT
}

func NewRotationZ(radians float64) Matrix {
	newT := IdentityMatrix()
	newT.Cells[0][0] = math.Cos(radians)
	newT.Cells[0][1] = -math.Sin(radians)
	newT.Cells[1][0] = math.Sin(radians)
	newT.Cells[1][1] = math.Cos(radians)
	return newT
}

func NewShearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	newT := IdentityMatrix()
	newT.Cells[0][1] = xy
	newT.Cells[0][2] = xz
	newT.Cells[1][0] = yx
	newT.Cells[1][2] = yz
	newT.Cells[2][0] = zx
	newT.Cells[2][1] = zy
	return newT
}

func (m *Matrix) GetRowCell(row, col int) float64 {
	return m.Cells[row][col]
}

func (m *Matrix) EqualsMatrix(m2 Matrix) bool {
	if m.Rows == m2.Rows &&
		m.Cols == m2.Cols {
		for x, row := range m.Cells {
			for y := range row {
				if !epsilonEquals(m.Cells[x][y], m2.Cells[x][y]) {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (m *Matrix) EqualsTuple(m2 Tuple) bool {
	if m.Rows == 4 &&
		m.Cols == 1 {
		return epsilonEquals(m.Cells[0][0], m2.X) &&
			epsilonEquals(m.Cells[1][0], m2.Y) &&
			epsilonEquals(m.Cells[2][0], m2.Z) &&
			epsilonEquals(m.Cells[3][0], m2.W)
	}
	return false
}

func (m *Matrix) MultiplyMatrix(m2 Matrix) Matrix {
	result := NewMatrix(m.Rows, m2.Cols)
	for row := 0; row < result.Rows; row++ {
		for col := 0; col < result.Cols; col++ {
			result.Cells[row][col] = 0
			for x := 0; x < result.Rows; x++ {
				result.Cells[row][col] = result.Cells[row][col] + m.Cells[row][x]*m2.Cells[x][col]
			}
		}
	}
	return result
}
func (m *Matrix) MultiplyTuple(m2 Tuple) Tuple {
	x := m.MultiplyMatrix(m2.ToMatrix())
	return NewTuple(
		x.Cells[0][0],
		x.Cells[1][0],
		x.Cells[2][0],
		x.Cells[3][0],
	)
}

func (m *Matrix) ToString() string {
	output := ""
	for row := 0; row < m.Rows; row++ {
		output = output + "|"
		for col := 0; col < m.Cols; col++ {
			output = output + fmt.Sprintf("%f|", m.Cells[row][col])
		}
		output = output + "\n"
	}
	return output
}

func (m *Matrix) Transpose() Matrix {
	n := NewMatrix(m.Cols, m.Rows)
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			n.Cells[col][row] = m.Cells[row][col]
		}
	}
	return n
}

func (m *Matrix) Determinant() float64 {
	if m.Rows == 2 && m.Cols == 2 {
		return m.Cells[0][0]*m.Cells[1][1] - m.Cells[0][1]*m.Cells[1][0]
	}
	det := 0.0
	for i := range m.Cells {
		det = det + m.Cells[0][i]*m.Cofactor(0, i)
	}
	return det
}

func (m *Matrix) Submatrix(subRow, subCol int) Matrix {
	newRowCount := m.Rows - 1
	newColCount := m.Cols - 1
	bob := NewMatrix(newRowCount, newColCount)
	for i := 0; i < subRow; i++ {
		bob.Cells[i] = make(map[int]float64)
		for j := 0; j < subCol; j++ {
			bob.Cells[i][j] = m.Cells[i][j]
		}
		for j := subCol; j < newColCount; j++ {
			bob.Cells[i][j] = m.Cells[i][j+1]
		}
	}
	for i := subRow; i < newRowCount; i++ {
		bob.Cells[i] = make(map[int]float64)
		for j := 0; j < subCol; j++ {
			bob.Cells[i][j] = m.Cells[i+1][j]
		}
		for j := subCol; j < newColCount; j++ {
			bob.Cells[i][j] = m.Cells[i+1][j+1]
		}
	}
	return bob
}

func (m *Matrix) Minor(row, col int) float64 {
	b := m.Submatrix(row, col)
	return b.Determinant()
}

func (m *Matrix) Cofactor(row, col int) float64 {
	b := m.Minor(row, col)
	if (row+col)%2 == 0 {
		return b
	}
	return -b
}

func (m *Matrix) Invertable() bool {
	return m.Determinant() != 0
}

func (m *Matrix) Inverse() Matrix {
	if !m.Invertable() {
		log.Fatal("Cannot invert this matrix")
	}

	m2 := NewMatrix(m.Rows, m.Cols)
	determinent := m.Determinant()

	for i := range m.Cells {
		for j := range m.Cells[i] {
			c := m.Cofactor(i, j)
			m2.Cells[j][i] = c / determinent
		}
	}
	return m2
}

func ViewTransform(from, to, up Tuple) Matrix {
	forward := to.Subtract(from).Normalize()
	left := forward.CrossProduct(up.Normalize())
	trueUp := left.CrossProduct(forward)

	orientation := NewMatrix(4, 4)
	orientation.Cells = make(map[int]map[int]float64)
	orientation.Cells[0] = map[int]float64{0: left.X, 1: left.Y, 2: left.Z, 3: 0}
	orientation.Cells[1] = map[int]float64{0: trueUp.X, 1: trueUp.Y, 2: trueUp.Z, 3: 0}
	orientation.Cells[2] = map[int]float64{0: -forward.X, 1: -forward.Y, 2: -forward.Z, 3: 0}
	orientation.Cells[3] = map[int]float64{0: 0, 1: 0, 2: 0, 3: 1}
	return orientation.MultiplyMatrix(NewTranslation(-from.X, -from.Y, -from.Z))
}
