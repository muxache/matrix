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

//FractionalMatrix has structure with number and fraction
type FractionalMatrix [][]Element

//Element part of the Fractional Matrix
type Element struct {
	Number   int64
	Fraction int64
}
