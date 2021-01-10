package matrix

//NewFractionalMatrix makes a new matrix with fractional elements
func NewFractionalMatrix(matrix [][]float64) FractionalMatrix {
	var fractionalMatrix FractionalMatrix
	fractionalMatrix = make([][]Element, len(matrix))
	for i := range matrix {
		fractionalMatrix[i] = make([]Element, len(matrix[i]))
		for j := range matrix[i] {
			fractionalMatrix[i][j].Number = int64(matrix[i][j])
			fractionalMatrix[i][j].Fraction = 1
		}
	}
	return fractionalMatrix
}
