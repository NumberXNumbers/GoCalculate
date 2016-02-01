package coreTypes

const (
	// ColVector is a vectorType Column Vector
	ColVector = "column"
	// RowVector is a vectorType Row Vector
	RowVector = "row"
)

// baseVector is basic vector interface
type baseVector interface {
	// Returns Type of vector. Either Column or Row vector
	Type() string

	// Returns dimensions of Vector
	Dim() int

	// Transpose of vector. i.e changes a row into a column vector and vice versa
	Trans()
}

type vectorBase struct {
	vectorType string
}

// implementation of Type method
func (v vectorBase) Type() string { return v.vectorType }
