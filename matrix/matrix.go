package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix represents one or more rows and column
type Matrix struct {
	_Rows, _Cols int
	Matrix       [][]int
}

// New returns new Matrix object
func New(m string) (*Matrix, error) {
	var elements [][]string
	for _, row := range strings.Split(m, "\n") {
		row = strings.Trim(row, " ")
		elements = append(elements, strings.Split(row, " "))
	}
	mat, err := validateMatrix(elements)
	if err != nil {
		return nil, err
	}
	return &mat, nil
}

// Set sets value of r and c in a matrix
func (m *Matrix) Set(rowIdx, colIdx, value int) bool {
	if rowIdx < 0 || colIdx < 0 ||
		rowIdx >= m._Rows || colIdx >= m._Cols {
		return false
	}
	if rowIdx > m._Rows || colIdx > m._Cols {
		return false
	}
	m.Matrix[rowIdx][colIdx] = value
	return true
}

// Rows returns rows from Matrix object
func (m *Matrix) Rows() [][]int {
	var newMat [][]int
	for r := 0; r <= m._Rows-1; r++ {
		var row []int
		for c := 0; c <= m._Cols-1; c++ {
			row = append(row, m.Matrix[r][c])
		}
		newMat = append(newMat, row)
	}
	return newMat
}

// Cols returns column from Matrix object
func (m *Matrix) Cols() [][]int {
	var newMat [][]int
	for c := 0; c <= m._Cols-1; c++ {
		var col []int
		for r := 0; r <= m._Rows-1; r++ {
			col = append(col, m.Matrix[r][c])
		}
		newMat = append(newMat, col)
	}
	return newMat
}

func parseIntSlice(row []string) ([]int, error) {
	var rowsInt []int
	for _, elem := range row {
		elemInt, err := strconv.ParseInt(elem, 10, 64)
		if err != nil {
			return []int{}, err
		}
		rowsInt = append(rowsInt, int(elemInt))
	}
	return rowsInt, nil
}

func validateMatrix(matrix [][]string) (Matrix, error) {
	var rowsInt64 [][]int
	var rows, cols int
	var areRowsEven bool

	for i := 0; i <= len(matrix)-1; i++ {

		// tricking to make sure that rows with
		// uneven sizes are not part of valid matrix
		if i == 0 {
			cols = len(matrix[i])
		}
		areRowsEven = cols == len(matrix[i])
		if !areRowsEven {
			return Matrix{}, fmt.Errorf("Uneven rows found")
		}

		rowInt64, err := parseIntSlice(matrix[i])
		if err != nil {

			return Matrix{}, err
		}

		rowsInt64 = append(rowsInt64, rowInt64)
	}
	rows = len(rowsInt64)

	return Matrix{_Rows: rows, _Cols: cols, Matrix: rowsInt64}, nil
}
