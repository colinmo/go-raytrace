package main

import (
	"fmt"
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

var IdentityMatrix = BuildMatrixFromString("|1|0|0|0|\n|0|1|0|0|\n|0|0|1|0|\n|0|0|0|1|")

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
func (m *Matrix) MultiplyTuple(m2 Tuple) Matrix {
	return m.MultiplyMatrix(m2.ToMatrix())
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
