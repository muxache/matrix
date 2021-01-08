package matrix

import (
	"errors"
	"math"
)

//SumWith sums two matrices
func (matrix Matrix) SumWith(mat2 Matrix) (Matrix, error) {
	var bodyOfMatrix [][]float64
	if len(matrix) != len(mat2) {
		return Matrix{}, errors.New("Matrices have deferent numbers of rows")
	}
	bodyOfMatrix = make([][]float64, len(matrix))
	for i := range matrix {
		bodyOfMatrix[i] = make([]float64, len(matrix[i]))
		if len(matrix[i]) != len(mat2[i]) {
			return Matrix{}, errors.New("Matrices have different number of columns")
		}
		for j := range matrix[i] {
			bodyOfMatrix[i][j] = matrix[i][j] + mat2[i][j]
		}
	}
	return NewMatrix(bodyOfMatrix), nil
}

//MultiplyWith returns composition of two matrices
func (matrix Matrix) MultiplyWith(mat2 Matrix) (Matrix, error) {
	m1rows, m1columns := matrix.MatrixSize()
	m2rows, m2columns := mat2.MatrixSize()
	var newMatrix [][]float64
	if m1rows == m2columns && m2rows == m1columns {
		newMatrix = make([][]float64, m1rows)
		for rows := range matrix {
			newMatrix[rows] = make([]float64, m1rows)
			for i := range matrix {
				for j := range matrix[i] {
					newMatrix[rows][i] += matrix[rows][j] * mat2[j][i]
				}

			}
		}
	} else {
		return nil, errors.New("Matrices have deferent numbers of rows and columns")
	}
	return newMatrix, nil
}

//Transpose swaps rows and collums
func (matrix Matrix) Transpose() Matrix {
	rows, colums := matrix.MatrixSize()
	var bodyOfMatrix [][]float64
	bodyOfMatrix = make([][]float64, colums)
	for j := 0; j < colums; j++ {
		bodyOfMatrix[j] = make([]float64, rows)
		for i := range matrix {
			bodyOfMatrix[j][i] = matrix[i][j]
		}
	}
	return bodyOfMatrix
}

//Determinant calculates determinant of matrix
func (matrix Matrix) Determinant() (float64, error) {
	var res float64
	rows, colums := matrix.MatrixSize()
	if rows != colums || matrix == nil {
		return 0, errors.New("Matrix isn't square")
	}
	if rows == 2 {
		return matrix[0][0]*matrix[1][1] + matrix[0][1]*matrix[1][0], nil
	}
	var newMatrix [][]float64
	newMatrix = make([][]float64, rows-1)
	var exI int = 0
	for exJ := 0; exJ < len(matrix); exJ++ {
		for i := range matrix {
			if i != exI {
				newRow := make([]float64, rows-1)
				for j := range matrix[i] {
					if j != exJ {
						jt := j
						if j > exJ {
							jt--
						}
						newRow[jt] = matrix[i][j]
					}
				}
				newMatrix[i-1] = newRow
			}
		}
		ch := make(chan float64)
		go AlgebraicComplement(exI+1, exJ+1, matrix[exI][exJ], newMatrix, ch)
		res += <-ch
	}
	return res, nil
}

//Minor calculates minor of matrix
func Minor(matrix Matrix) float64 {
	var res float64
	rows, colums := matrix.MatrixSize()
	if rows != colums {
		return 0
	}
	if rows == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}
	var newMatrix [][]float64
	newMatrix = make([][]float64, rows-1)
	var exI int = 0
	for exJ := 0; exJ < len(matrix); exJ++ {
		for i := range matrix {
			if i != exI {
				newRow := make([]float64, rows-1)
				for j := range matrix[i] {
					if j != exJ {
						jt := j
						if j > exJ {
							jt--
						}
						newRow[jt] = matrix[i][j]
					}
				}
				newMatrix[i-1] = newRow
			}
		}
		res += math.Pow(-1, float64((exI+1)+(exJ+1))) * matrix[exI][exJ] * Minor(newMatrix)

	}
	return res
}

//AlgebraicComplement calculating algebraic complement and writes in channel
func AlgebraicComplement(row, column int, element float64, minor [][]float64, ch chan float64) {
	ch <- math.Pow(-1, float64((row)+(column))) * element * Minor(minor)
}
