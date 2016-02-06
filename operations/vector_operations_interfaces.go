package operations

import "github.com/traviscox1990/GoCalculate/coreTypes"

const (
	// RowVector for operations
	RowVector = coreTypes.RowVector
	// ColVector for operations
	ColVector = coreTypes.ColVector
)

// Vector is an interface over coreTypes Vector for operations
type Vector interface {
	coreTypes.Vector
}

// VectorComplex is an interface over coreTypes VectorComplex for operations
type VectorComplex interface {
	coreTypes.VectorComplex
}

func makeComplexVector(length int, vectorType string) (vector VectorComplex) {
	vector = coreTypes.MakeComplexVector(length, vectorType)
	return
}

func makeVector(length int, vectorType string) (vector Vector) {
	vector = coreTypes.MakeVector(length, vectorType)
	return
}
