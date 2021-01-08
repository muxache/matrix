package matrix

import "fmt"

//PrintByRow print all rows in the matrix
func (matrix Matrix) PrintByRow() {
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

//MatrixSize returns numbers of rows and collums of matrix
func (matrix Matrix) MatrixSize() (int, int) {
	if len(matrix) == 0 {
		return 0, 0
	}
	return len(matrix), len(matrix[0])
}
