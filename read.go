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

//PrintByRow printing fraction matrix
func (f FractionalMatrix) PrintByRow() {
	for i := range f {
		for j := range f[i] {
			fmt.Printf("%d/%d ", f[i][j].Number, f[i][j].Fraction)
		}
		fmt.Println()
	}
}

//MatrixSize returns numbers of rows and collums of matrix
func (f FractionalMatrix) MatrixSize() (int, int) {
	if len(f) == 0 {
		return 0, 0
	}
	return len(f), len(f[0])
}
