package operations

import (
	"errors"
	"math"
	"math/cmplx"

	ct "github.com/traviscox1990/GoCalculate/coreTypes"
)

// VectorScalarMulti multiplies a real vector by a scalar
func VectorScalarMulti(scalar float64, vector ct.Vector) ct.Vector {
	newVector := ct.MakeVector(vector.Dim(), vector.Type())
	for i := 0; i < vector.Dim(); i++ {
		newVector.Set(i, scalar*vector.Get(i))
	}

	return newVector
}

// VectorComplexScalarMulti multiplies a complex vector by a scalar
func VectorComplexScalarMulti(scalar complex128, vector ct.VectorComplex) ct.VectorComplex {
	newVector := ct.MakeComplexVector(vector.Dim(), vector.Type())
	for i := 0; i < vector.Dim(); i++ {
		newVector.Set(i, scalar*vector.Get(i))
	}

	return newVector
}

// AngleTheta returns the angle theta between Vector A and Vector B using the dot product
func AngleTheta(vectorA ct.Vector, vectorB ct.Vector) (float64, error) {
	normA := vectorA.Norm()
	normB := vectorB.Norm()
	var theta float64

	if normA == 0 || normB == 0 {
		return theta, errors.New("Either Vector A or Vector B is the zero vector")
	}

	dotProduct, err := InnerProduct(vectorA, vectorB)

	if err != nil {
		return theta, err
	}

	theta = math.Acos(dotProduct / (normA * normB))

	return theta, nil
}

// AngleThetaComplex returns the angle theta between VectorComplex A and VectorComplex B using the dot product
func AngleThetaComplex(vectorA ct.VectorComplex, vectorB ct.VectorComplex) (complex128, error) {
	normA := vectorA.Norm()
	normB := vectorB.Norm()
	var theta complex128

	if normA == 0 || normB == 0 {
		return theta, errors.New("Either Vector A or Vector B is the zero vector")
	}

	dotProduct, err := InnerProductComplex(vectorA, vectorB)

	if err != nil {
		return theta, err
	}

	theta = cmplx.Acos(dotProduct / (normA * normB))

	return theta, nil
}

// InnerProduct returns the inner product for real Vectors
func InnerProduct(vectorA ct.Vector, vectorB ct.Vector) (float64, error) {
	var product float64

	if vectorA.Dim() != vectorB.Dim() {
		return product, errors.New("Length of vectors does not match")
	}

	if vectorA.Type() != ct.RowVector || vectorB.Type() != ct.ColVector {
		return product, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	for i := 0; i < vectorA.Dim(); i++ {
		product += vectorA.Get(i) * vectorB.Get(i)
	}

	return product, nil
}

// InnerProductComplex returns the inner product for complex Vectors
func InnerProductComplex(vectorA ct.VectorComplex, vectorB ct.VectorComplex) (complex128, error) {
	var product complex128

	if vectorA.Dim() != vectorB.Dim() {
		return product, errors.New("Length of vectors does not match")
	}

	if vectorA.Type() != ct.RowVector || vectorB.Type() != ct.ColVector {
		return product, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	for i := 0; i < vectorA.Dim(); i++ {
		product += vectorA.Get(i) * vectorB.Get(i)
	}

	return product, nil
}

// OuterProduct returns the outer product for real Vectors
func OuterProduct(vectorA ct.Vector, vectorB ct.Vector) (ct.Matrix, error) {
	if vectorA.Type() != ct.ColVector || vectorB.Type() != ct.RowVector {
		return nil, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	matrix := ct.MakeMatrix(vectorA.Dim(), vectorB.Dim())

	for i := 0; i < vectorA.Dim(); i++ {
		for j := 0; j < vectorB.Dim(); j++ {
			matrix.Set(i, j, vectorA.Get(i)*vectorB.Get(j))
		}
	}

	return matrix, nil
}

// OuterProductComplex returns the outer product for real Vectors
func OuterProductComplex(vectorA ct.VectorComplex, vectorB ct.VectorComplex) (ct.MatrixComplex, error) {
	if vectorA.Type() != ct.ColVector || vectorB.Type() != ct.RowVector {
		return nil, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	matrix := ct.MakeComplexMatrix(vectorA.Dim(), vectorB.Dim())

	for i := 0; i < vectorA.Dim(); i++ {
		for j := 0; j < vectorB.Dim(); j++ {
			matrix.Set(i, j, vectorA.Get(i)*vectorB.Get(j))
		}
	}

	return matrix, nil
}

// VectorAddition adds two real vectors together
func VectorAddition(vectorA ct.Vector, vectorB ct.Vector) (ct.Vector, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Dim() != vectorB.Dim() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := ct.MakeVector(vectorA.Dim(), vectorA.Type())

	for i := 0; i < vectorA.Dim(); i++ {
		vector.Set(i, vectorA.Get(i)+vectorB.Get(i))
	}

	return vector, nil
}

// VectorComplexAddition adds two complex vectors together
func VectorComplexAddition(vectorA ct.VectorComplex, vectorB ct.VectorComplex) (ct.VectorComplex, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Dim() != vectorB.Dim() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := ct.MakeComplexVector(vectorA.Dim(), vectorA.Type())

	for i := 0; i < vectorA.Dim(); i++ {
		vector.Set(i, vectorA.Get(i)+vectorB.Get(i))
	}

	return vector, nil
}

// VectorSubtraction subtracts two real vectors together
func VectorSubtraction(vectorA ct.Vector, vectorB ct.Vector) (ct.Vector, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Dim() != vectorB.Dim() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := ct.MakeVector(vectorA.Dim(), vectorA.Type())

	for i := 0; i < vectorA.Dim(); i++ {
		vector.Set(i, vectorA.Get(i)-vectorB.Get(i))
	}

	return vector, nil
}

// VectorComplexSubtraction subtracts two complex vectors together
func VectorComplexSubtraction(vectorA ct.VectorComplex, vectorB ct.VectorComplex) (ct.VectorComplex, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Dim() != vectorB.Dim() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := ct.MakeComplexVector(vectorA.Dim(), vectorA.Type())

	for i := 0; i < vectorA.Dim(); i++ {
		vector.Set(i, vectorA.Get(i)-vectorB.Get(i))
	}

	return vector, nil
}
