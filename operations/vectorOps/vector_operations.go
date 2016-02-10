package vectorOps

import (
	"errors"
	"math"
	"math/cmplx"
)

// VectorScalarMulti multiplies a real vector by a scalar
func VectorScalarMulti(scalar float64, vector Vector) Vector {
	newVector := makeVector(vector.Dim(), vector.Type())
	for i := 0; i < vector.Dim(); i++ {
		newVector.Set(i, scalar*vector.Get(i))
	}

	return newVector
}

// VectorComplexScalarMulti multiplies a complex vector by a scalar
func VectorComplexScalarMulti(scalar complex128, vector VectorComplex) VectorComplex {
	newVector := makeComplexVector(vector.Dim(), vector.Type())
	for i := 0; i < vector.Dim(); i++ {
		newVector.Set(i, scalar*vector.Get(i))
	}

	return newVector
}

// AngleTheta returns the angle theta between Vector A and Vector B using the dot product
func AngleTheta(vectorA Vector, vectorB Vector) (float64, error) {
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
func AngleThetaComplex(vectorA VectorComplex, vectorB VectorComplex) (complex128, error) {
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
func InnerProduct(vectorA Vector, vectorB Vector) (float64, error) {
	var product float64

	if vectorA.Dim() != vectorB.Dim() {
		return product, errors.New("Length of vectors does not match")
	}

	if vectorA.Type() != RowVector || vectorB.Type() != ColVector {
		return product, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	for i := 0; i < vectorA.Dim(); i++ {
		product += vectorA.Get(i) * vectorB.Get(i)
	}

	return product, nil
}

// InnerProductComplex returns the inner product for complex Vectors
func InnerProductComplex(vectorA VectorComplex, vectorB VectorComplex) (complex128, error) {
	var product complex128

	if vectorA.Dim() != vectorB.Dim() {
		return product, errors.New("Length of vectors does not match")
	}

	if vectorA.Type() != RowVector || vectorB.Type() != ColVector {
		return product, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	for i := 0; i < vectorA.Dim(); i++ {
		product += vectorA.Get(i) * vectorB.Get(i)
	}

	return product, nil
}

// OuterProduct returns the outer product for real Vectors
func OuterProduct(vectorA Vector, vectorB Vector) (Matrix, error) {
	if vectorA.Type() != ColVector || vectorB.Type() != RowVector {
		return nil, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	matrix := makeMatrix(vectorA.Dim(), vectorB.Dim())

	for i := 0; i < vectorA.Dim(); i++ {
		for j := 0; j < vectorB.Dim(); j++ {
			matrix.Set(i, j, vectorA.Get(i)*vectorB.Get(j))
		}
	}

	return matrix, nil
}

// OuterProductComplex returns the outer product for real Vectors
func OuterProductComplex(vectorA VectorComplex, vectorB VectorComplex) (MatrixComplex, error) {
	if vectorA.Type() != ColVector || vectorB.Type() != RowVector {
		return nil, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	matrix := makeComplexMatrix(vectorA.Dim(), vectorB.Dim())

	for i := 0; i < vectorA.Dim(); i++ {
		for j := 0; j < vectorB.Dim(); j++ {
			matrix.Set(i, j, vectorA.Get(i)*vectorB.Get(j))
		}
	}

	return matrix, nil
}

// VectorAddition adds two real vectors together
func VectorAddition(vectorA Vector, vectorB Vector) (Vector, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Dim() != vectorB.Dim() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := makeVector(vectorA.Dim(), vectorA.Type())

	for i := 0; i < vectorA.Dim(); i++ {
		vector.Set(i, vectorA.Get(i)+vectorB.Get(i))
	}

	return vector, nil
}

// VectorComplexAddition adds two complex vectors together
func VectorComplexAddition(vectorA VectorComplex, vectorB VectorComplex) (VectorComplex, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Dim() != vectorB.Dim() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := makeComplexVector(vectorA.Dim(), vectorA.Type())

	for i := 0; i < vectorA.Dim(); i++ {
		vector.Set(i, vectorA.Get(i)+vectorB.Get(i))
	}

	return vector, nil
}

// VectorSubtraction subtracts two real vectors together
func VectorSubtraction(vectorA Vector, vectorB Vector) (Vector, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Dim() != vectorB.Dim() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := makeVector(vectorA.Dim(), vectorA.Type())

	for i := 0; i < vectorA.Dim(); i++ {
		vector.Set(i, vectorA.Get(i)-vectorB.Get(i))
	}

	return vector, nil
}

// VectorComplexSubtraction subtracts two complex vectors together
func VectorComplexSubtraction(vectorA VectorComplex, vectorB VectorComplex) (VectorComplex, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Dim() != vectorB.Dim() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := makeComplexVector(vectorA.Dim(), vectorA.Type())

	for i := 0; i < vectorA.Dim(); i++ {
		vector.Set(i, vectorA.Get(i)-vectorB.Get(i))
	}

	return vector, nil
}
