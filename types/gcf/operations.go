package gcf

import (
	"github.com/NumberXNumbers/GoCalculate/types/gcv/gcvops"
	"github.com/NumberXNumbers/GoCalculate/types/m/mops"
	"github.com/NumberXNumbers/GoCalculate/types/v"
	"github.com/NumberXNumbers/GoCalculate/types/v/vops"
)

// Add will add two constants together
func Add(constA Const, constB Const) Const {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Add(constA.Value(), constB.Value()))
	}

	if constA.Type() == Vector && constB.Type() == Vector {
		vector, err := vops.Add(constA.Vector(), constB.Vector())
		if err != nil {
			panic(err)
		}
		return MakeConst(vector)
	}

	if constA.Type() == Matrix && constB.Type() == Matrix {
		matrix, err := mops.Add(constA.Matrix(), constB.Matrix())
		if err != nil {
			panic(err)
		}
		return MakeConst(matrix)
	}
	panic("One or More Types are incorrect")
}

// Sub will subtract two constants together
func Sub(constA Const, constB Const) Const {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Sub(constA.Value(), constB.Value()))
	}
	if constA.Type() == Vector && constB.Type() == Vector {
		vector, err := vops.Sub(constA.Vector(), constB.Vector())
		if err != nil {
			panic(err)
		}
		return MakeConst(vector)
	}
	if constA.Type() == Matrix && constB.Type() == Matrix {
		matrix, err := mops.Sub(constA.Matrix(), constB.Matrix())
		if err != nil {
			panic(err)
		}
		return MakeConst(matrix)
	}
	panic("One or More Types are incorrect")
}

// Div will divide two constants together
func Div(constA Const, constB Const) Const {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Div(constA.Value(), constB.Value()))
	}
	if constA.Type() == Vector && constB.Type() == Value {
		return MakeConst(vops.SDiv(constB.Value(), constA.Vector()))
	}
	if constA.Type() == Matrix && constB.Type() == Value {
		return MakeConst(mops.SDiv(constB.Value(), constA.Matrix()))
	}
	panic("One or More Types are incorrect")
}

// Mult will multiply two constants together
func Mult(constA Const, constB Const) Const {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Mult(constA.Value(), constB.Value()))
	}
	if constA.Type() == Vector && constB.Type() == Value {
		return MakeConst(vops.SMult(constB.Value(), constA.Vector()))
	}
	if constA.Type() == Value && constB.Type() == Vector {
		return MakeConst(vops.SMult(constA.Value(), constB.Vector()))
	}
	if constA.Type() == Vector && constB.Type() == Vector {
		vectorA := constA.Vector()
		vectorB := constB.Vector()
		if vectorA.Space() == v.RowSpace {
			vector, err := vops.InnerProduct(vectorA, vectorB)
			if err != nil {
				panic(err)
			}
			return MakeConst(vector)
		}
		matrix, err := vops.OuterProduct(vectorA, vectorB)
		if err != nil {
			panic(err)
		}
		return MakeConst(matrix)
	}
	if constA.Type() == Matrix && constB.Type() == Value {
		return MakeConst(mops.SMult(constB.Value(), constA.Matrix()))
	}
	if constA.Type() == Value && constB.Type() == Matrix {
		return MakeConst(mops.SMult(constA.Value(), constB.Matrix()))
	}
	if constA.Type() == Vector && constB.Type() == Matrix {
		vector, err := mops.VMMult(constA.Vector(), constB.Matrix())
		if err != nil {
			panic(err)
		}
		return MakeConst(vector)
	}
	if constA.Type() == Matrix && constB.Type() == Vector {
		vector, err := mops.MVMult(constB.Vector(), constA.Matrix())
		if err != nil {
			panic(err)
		}
		return MakeConst(vector)
	}
	matrix, err := mops.MultSimple(constA.Matrix(), constB.Matrix())
	if err != nil {
		panic(err)
	}
	return MakeConst(matrix)
}
