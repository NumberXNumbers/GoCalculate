package vops

import (
	"errors"
	"math"
	"math/cmplx"
)

// VectorScalarMulti multiplies a real vector by a scalar
func VectorScalarMulti(scalar Value, vector Vector) Vector {
	newVector := makeVector(vector.Len(), vector.Type())
	for i := 0; i < vector.Len(); i++ {
		newVector.Set(i, newValue(scalar.Complex128()*vector.Get(i).Complex128()))
	}

	return newVector
}

// AngleTheta returns the angle theta between Vector A and Vector B using the dot product
func AngleTheta(vectorA Vector, vectorB Vector) (Value, error) {
	normA := vectorA.Norm()
	normB := vectorB.Norm()
	var theta Value

	if normA.Int() == 0 || normB.Int() == 0 {
		return theta, errors.New("Either Vector A or Vector B is the zero vector")
	}

	dotProduct, err := InnerProduct(vectorA, vectorB)

	if err != nil {
		return theta, err
	}

	if dotProduct.GetValueType() == Complex {
		theta = newValue(cmplx.Acos(dotProduct.Complex128() / (normA.Complex128() * normB.Complex128())))
	} else {
		theta = newValue(math.Acos(dotProduct.Float64() / (normA.Float64() * normB.Float64())))
	}

	return theta, nil
}

// InnerProduct returns the inner product for real Vectors
func InnerProduct(vectorA Vector, vectorB Vector) (Value, error) {
	var product Value

	if vectorA.Len() != vectorB.Len() {
		return product, errors.New("Length of vectors does not match")
	}

	if vectorA.Type() != RowVector || vectorB.Type() != ColVector {
		return product, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	if vectorA.Elements().CoreType() == Complex || vectorB.Elements().CoreType() == Complex {
		var complexProduct complex128
		for i := 0; i < vectorA.Len(); i++ {
			complexProduct += vectorA.Get(i).Complex128() * vectorB.Get(i).Complex128()
		}
		product = newValue(complexProduct)
	} else {
		var floatProduct float64
		for i := 0; i < vectorA.Len(); i++ {
			floatProduct += vectorA.Get(i).Float64() * vectorB.Get(i).Float64()
		}
		product = newValue(floatProduct)
	}

	return product, nil
}

// OuterProduct returns the outer product for real Vectors
func OuterProduct(vectorA Vector, vectorB Vector) (Matrix, error) {
	if vectorA.Type() != ColVector || vectorB.Type() != RowVector {
		return nil, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	matrix := makeMatrix(vectorA.Len(), vectorB.Len())

	for i := 0; i < vectorA.Len(); i++ {
		for j := 0; j < vectorB.Len(); j++ {
			matrix.Set(i, j, vectorA.Get(i).Float64()*vectorB.Get(j).Float64())
		}
	}

	return matrix, nil
}

// OuterProductComplex returns the outer product for real Vectors
func OuterProductComplex(vectorA Vector, vectorB Vector) (MatrixComplex, error) {
	if vectorA.Type() != ColVector || vectorB.Type() != RowVector {
		return nil, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	matrix := makeComplexMatrix(vectorA.Len(), vectorB.Len())

	for i := 0; i < vectorA.Len(); i++ {
		for j := 0; j < vectorB.Len(); j++ {
			matrix.Set(i, j, vectorA.Get(i).Complex128()*vectorB.Get(j).Complex128())
		}
	}

	return matrix, nil
}

// VectorAddition adds two real vectors together
func VectorAddition(vectorA Vector, vectorB Vector) (Vector, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Len() != vectorB.Len() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := makeVector(vectorA.Len(), vectorA.Type())

	if vectorA.Elements().CoreType() == Complex || vectorB.Elements().CoreType() == Complex {
		for i := 0; i < vectorA.Len(); i++ {
			vector.Set(i, newValue(vectorA.Get(i).Complex128()+vectorB.Get(i).Complex128()))
		}
	} else {
		for i := 0; i < vectorA.Len(); i++ {
			vector.Set(i, newValue(vectorA.Get(i).Float64()+vectorB.Get(i).Float64()))
		}
	}

	return vector, nil
}

// VectorSubtraction subtracts two real vectors together
func VectorSubtraction(vectorA Vector, vectorB Vector) (Vector, error) {
	if vectorA.Type() != vectorB.Type() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Len() != vectorB.Len() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := makeVector(vectorA.Len(), vectorA.Type())

	if vectorA.Elements().CoreType() == Complex || vectorB.Elements().CoreType() == Complex {
		for i := 0; i < vectorA.Len(); i++ {
			vector.Set(i, newValue(vectorA.Get(i).Complex128()-vectorB.Get(i).Complex128()))
		}
	} else {
		for i := 0; i < vectorA.Len(); i++ {
			vector.Set(i, newValue(vectorA.Get(i).Float64()-vectorB.Get(i).Float64()))
		}
	}

	return vector, nil
}
