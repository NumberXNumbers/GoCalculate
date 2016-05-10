package vops

import (
	"errors"
	"math"
	"math/cmplx"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

// ScalarMultiplication multiplies a real vector by a scalar
func ScalarMultiplication(scalar gcv.Value, vector v.Vector) v.Vector {
	newVector := v.NewVector(vector.Space(), vector.Len())
	for i := 0; i < vector.Len(); i++ {
		newVector.Set(i, gcv.MakeValue(scalar.Complex()*vector.Get(i).Complex()))
	}

	return newVector
}

// AngleTheta returns the angle theta between Vector A and Vector B using the dot product
func AngleTheta(vectorA v.Vector, vectorB v.Vector) (gcv.Value, error) {
	normA := vectorA.Norm()
	normB := vectorB.Norm()
	var theta gcv.Value

	if normA.Real() == 0 || normB.Real() == 0 {
		return theta, errors.New("Either Vector A or Vector B is the zero vector")
	}

	dotProduct, err := InnerProduct(vectorA, vectorB)

	if err != nil {
		return theta, err
	}

	if dotProduct.Type() == gcv.Complex {
		theta = gcv.MakeValue(cmplx.Acos(dotProduct.Complex() / (normA.Complex() * normB.Complex())))
	} else {
		theta = gcv.MakeValue(math.Acos(dotProduct.Real() / (normA.Real() * normB.Real())))
	}

	return theta, nil
}

// InnerProduct returns the inner product for real Vectors
func InnerProduct(vectorA v.Vector, vectorB v.Vector) (gcv.Value, error) {
	var product gcv.Value

	if vectorA.Len() != vectorB.Len() {
		return product, errors.New("Length of vectors does not match")
	}

	if vectorA.Space() != v.RowSpace || vectorB.Space() != v.ColSpace {
		return product, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	if vectorA.Elements().Type() == gcv.Complex || vectorB.Elements().Type() == gcv.Complex {
		var complexProduct complex128
		for i := 0; i < vectorA.Len(); i++ {
			complexProduct += vectorA.Get(i).Complex() * vectorB.Get(i).Complex()
		}
		product = gcv.MakeValue(complexProduct)
	} else {
		var floatProduct float64
		for i := 0; i < vectorA.Len(); i++ {
			floatProduct += vectorA.Get(i).Real() * vectorB.Get(i).Real()
		}
		product = gcv.MakeValue(floatProduct)
	}

	return product, nil
}

// OuterProduct returns the outer product for real Vectors
func OuterProduct(vectorA v.Vector, vectorB v.Vector) (m.Matrix, error) {
	if vectorA.Space() != v.ColSpace || vectorB.Space() != v.RowSpace {
		return nil, errors.New("One or both vector types are not consistent with the vector inner product")
	}

	matrix := m.NewMatrix(vectorA.Len(), vectorB.Len())

	matrixType := matrix.Type()
	if vectorA.Type() == gcv.Complex || vectorB.Type() == gcv.Complex {
		matrixType = gcv.Complex
	}

	for i := 0; i < vectorA.Len(); i++ {
		for j := 0; j < vectorB.Len(); j++ {
			if matrixType == gcv.Complex {
				matrix.Set(i, j, gcv.MakeValue(vectorA.Get(i).Complex()*vectorB.Get(j).Complex()))
				continue
			}
			matrix.Set(i, j, gcv.MakeValue(vectorA.Get(i).Real()*vectorB.Get(j).Real()))
		}
	}

	return matrix, nil
}

// Addition adds two real vectors together
func Addition(vectorA v.Vector, vectorB v.Vector) (v.Vector, error) {
	if vectorA.Space() != vectorB.Space() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Len() != vectorB.Len() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := v.NewVector(vectorA.Space(), vectorA.Len())

	if vectorA.Elements().Type() == gcv.Complex || vectorB.Elements().Type() == gcv.Complex {
		for i := 0; i < vectorA.Len(); i++ {
			vector.Set(i, gcv.MakeValue(vectorA.Get(i).Complex()+vectorB.Get(i).Complex()))
		}
	} else {
		for i := 0; i < vectorA.Len(); i++ {
			vector.Set(i, gcv.MakeValue(vectorA.Get(i).Real()+vectorB.Get(i).Real()))
		}
	}

	return vector, nil
}

// Subtraction subtracts two real vectors together
func Subtraction(vectorA v.Vector, vectorB v.Vector) (v.Vector, error) {
	if vectorA.Space() != vectorB.Space() {
		return nil, errors.New("Vectors are not of same type. Must be both be either column vectors or row vectors")
	}

	if vectorA.Len() != vectorB.Len() {
		return nil, errors.New("Vectors are not same dimensions")
	}

	vector := v.NewVector(vectorA.Space(), vectorA.Len())

	if vectorA.Elements().Type() == gcv.Complex || vectorB.Elements().Type() == gcv.Complex {
		for i := 0; i < vectorA.Len(); i++ {
			vector.Set(i, gcv.MakeValue(vectorA.Get(i).Complex()-vectorB.Get(i).Complex()))
		}
	} else {
		for i := 0; i < vectorA.Len(); i++ {
			vector.Set(i, gcv.MakeValue(vectorA.Get(i).Real()-vectorB.Get(i).Real()))
		}
	}

	return vector, nil
}
