package matrix

import (
	"errors"
	"math"
)

//SumWith sums two matrices
func (matrix FractionalMatrix) SumWith(mat2 FractionalMatrix) (FractionalMatrix, error) {
	var bodyOfMatrix [][]Element
	if len(matrix) != len(mat2) {
		return nil, errors.New("Matrices have deferent numbers of rows")
	}
	bodyOfMatrix = make([][]Element, len(matrix))
	for i := range matrix {
		bodyOfMatrix[i] = make([]Element, len(matrix[i]))
		if len(matrix[i]) != len(mat2[i]) {
			return nil, errors.New("Matrices have different number of columns")
		}
		for j := range matrix[i] {
			bodyOfMatrix[i][j] = Addition(matrix[i][j], mat2[i][j])
		}
	}
	return bodyOfMatrix, nil
}

//MultiplyWith returns composition of two matrices
func (matrix FractionalMatrix) MultiplyWith(mat2 FractionalMatrix) (FractionalMatrix, error) {
	m1rows, m1columns := matrix.MatrixSize()
	m2rows, _ := mat2.MatrixSize()
	if m1columns == 0 && m2rows == 0 {
		return nil, errors.New("Matrices are zero")
	}
	var newMatrix [][]Element
	if m2rows == m1columns {
		newMatrix = make([][]Element, m1rows)
		for rows := range matrix {
			newMatrix[rows] = make([]Element, m1rows)
			for i := range matrix {
				for j := range matrix {
					newMatrix[rows][i] = Addition(newMatrix[rows][i], Multiply(matrix[rows][j], mat2[j][i]))
				}
			}
		}
	} else {
		return nil, errors.New("Matrices have deferent numbers of rows and columns")
	}
	return newMatrix, nil
}

//Transpose swaps rows and collums
func (matrix FractionalMatrix) Transpose() FractionalMatrix {
	rows, colums := matrix.MatrixSize()
	var bodyOfMatrix [][]Element
	bodyOfMatrix = make([][]Element, colums)
	for j := 0; j < colums; j++ {
		bodyOfMatrix[j] = make([]Element, rows)
		for i := range matrix {
			bodyOfMatrix[j][i] = matrix[i][j]
		}
	}
	return bodyOfMatrix
}

//Determinant calculates determinant of matrix
func (matrix FractionalMatrix) Determinant() (Element, error) {
	rows, colums := matrix.MatrixSize()
	if rows != colums || matrix == nil {
		return Element{}, errors.New("Matrix isn't square")
	}
	return MinorFractional(matrix), nil
}

//Minor calculates minor of matrix
func MinorFractional(matrix FractionalMatrix) Element {
	var res Element
	rows, colums := matrix.MatrixSize()
	if rows != colums {
		return Element{}
	}
	if rows == 2 {
		return Subtraction(Multiply(matrix[0][0], matrix[1][1]), Multiply(matrix[0][1], matrix[1][0]))
	}
	var newMatrix [][]Element
	newMatrix = make([][]Element, rows-1)
	var exI int = 0
	for exJ := 0; exJ < len(matrix); exJ++ {
		for i := range matrix {
			if i != exI {
				newRow := make([]Element, rows-1)
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
		//res += math.Pow(-1, float64((exI+1)+(exJ+1))) * matrix[exI][exJ] * Minor(newMatrix)
		ch := make(chan Element)
		go AlgebraicComplementFractional(exI+1, exJ+1, matrix[exI][exJ], newMatrix, ch)
		res = Addition(res, <-ch)
	}
	return res
}

//AlgebraicComplementFractional calculating algebraic complement and writes in channel
func AlgebraicComplementFractional(row, column int, element Element, minor [][]Element, ch chan Element) {
	ch <- MultiplyWithNumber(Multiply(element, MinorFractional(minor)), math.Pow(-1, float64((row)+(column))))
}
