package matrix

//NewMatrix initiate a new matrix
func NewMatrix(matrix [][]float64) Matrix {
	if len(matrix) < 1 {
		return Matrix{}
	}
	if len(matrix[0]) < 1 {
		return Matrix{}
	}

	return Padding(matrix)
}

//GenerateMatrix genetate zeroMatrix
func GenerateMatrix(i, j int) Matrix {
	var newMatrix [][]float64
	for i == 0 {
		newMatrix = make([][]float64, i)
		for j == 0 {
			newMatrix[i] = make([]float64, j)
			newMatrix[i][j] = 0
			j--
		}
		i--
	}
	return newMatrix
}

//Padding the matrix adds the 0 in the empty colums
func Padding(matrix [][]float64) Matrix {
	var maxColums int
	var filledMatrix [][]float64
	for m := range matrix {
		if len(matrix[m]) > maxColums {
			maxColums = len(matrix[m])
		}
	}
	filledMatrix = make([][]float64, len(matrix))
	for i := range matrix {
		filledMatrix[i] = make([]float64, maxColums)
		for j := range matrix[i] {
			filledMatrix[i][j] = matrix[i][j]
			if len(matrix[i]) < maxColums {
				for k := len(matrix[i]); k < maxColums; k++ {
					filledMatrix[i][k] = 0
				}
			}
		}
	}
	return filledMatrix
}
