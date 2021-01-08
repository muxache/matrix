package matrix

//PropertyMatrix contains number of rows and collums, elemenst of matrix and descriprion as well.
type PropertyMatrix struct {
	Body        Matrix
	Rows        int
	Collums     int
	Description string
}

//Matrix contains a matrix
type Matrix [][]float64
