package m

import (
	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/gcv/gcvops"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

func sMult(scalar gcv.Value, vector v.Vector) v.Vector {
	newVector := v.NewVector(vector.Space(), vector.Len())
	for i := 0; i < vector.Len(); i++ {
		newVector.Set(i, gcv.MakeValue(scalar.Complex()*vector.Get(i).Complex()))
	}

	return newVector
}

func sDiv(scalar gcv.Value, vector v.Vector) v.Vector {
	newVector := vector.Copy()
	for i := 0; i < newVector.Len(); i++ {
		newVector.Set(i, gcvops.Div(vector.Get(i), scalar))
	}
	return newVector
}

func sub(vectorA v.Vector, vectorB v.Vector) v.Vector {
	vector := v.NewVector(vectorA.Space(), vectorA.Len())
	for i := 0; i < vectorA.Len(); i++ {
		vector.Set(i, gcvops.Sub(vectorA.Get(i), vectorB.Get(i)))
	}

	return vector
}
