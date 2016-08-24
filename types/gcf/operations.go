package gcf

import (
	"errors"

	"github.com/NumberXNumbers/GoCalculate/types/gcv/gcvops"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/m/mops"
	"github.com/NumberXNumbers/GoCalculate/types/v"
	"github.com/NumberXNumbers/GoCalculate/types/v/vops"
)

// Add will add two constants together
func Add(constA Const, constB Const) (Const, error) {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Add(constA.Value(), constB.Value())), nil
	}

	if constA.Type() == Vector && constB.Type() == Vector {
		vector, err := vops.Add(constA.Vector(), constB.Vector())
		if err != nil {
			return nil, err
		}
		return MakeConst(vector), nil
	}

	if constA.Type() == Matrix && constB.Type() == Matrix {
		matrix, err := mops.Add(constA.Matrix(), constB.Matrix())
		if err != nil {
			return nil, err
		}
		return MakeConst(matrix), nil
	}
	return nil, errors.New("One or More Types are not supported")
}

// MustAdd is the same as Add but will panic
func MustAdd(constA Const, constB Const) Const {
	constant, err := Add(constA, constB)
	if err != nil {
		panic(err)
	}
	return constant
}

// Sub will subtract two constants together
func Sub(constA Const, constB Const) (Const, error) {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Sub(constA.Value(), constB.Value())), nil
	}
	if constA.Type() == Vector && constB.Type() == Vector {
		vector, err := vops.Sub(constA.Vector(), constB.Vector())
		if err != nil {
			return nil, err
		}
		return MakeConst(vector), nil
	}
	if constA.Type() == Matrix && constB.Type() == Matrix {
		matrix, err := mops.Sub(constA.Matrix(), constB.Matrix())
		if err != nil {
			return nil, err
		}
		return MakeConst(matrix), nil
	}
	return nil, errors.New("One or More Types are not supported")
}

// MustSub is the same as Sub but will panic
func MustSub(constA Const, constB Const) Const {
	constant, err := Sub(constA, constB)
	if err != nil {
		panic(err)
	}
	return constant
}

// Div will divide two constants together
func Div(constA Const, constB Const) (Const, error) {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Div(constA.Value(), constB.Value())), nil
	}
	if constA.Type() == Vector && constB.Type() == Value {
		return MakeConst(vops.SDiv(constB.Value(), constA.Vector())), nil
	}
	if constA.Type() == Matrix && constB.Type() == Value {
		return MakeConst(mops.SDiv(constB.Value(), constA.Matrix())), nil
	}
	return nil, errors.New("One or More Types are not supported")
}

// MustDiv is the same as Div but will panic
func MustDiv(constA Const, constB Const) Const {
	constant, err := Div(constA, constB)
	if err != nil {
		panic(err)
	}
	return constant
}

// Mult will multiply two constants together
func Mult(constA Const, constB Const) (Const, error) {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Mult(constA.Value(), constB.Value())), nil
	}
	if constA.Type() == Vector && constB.Type() == Value {
		return MakeConst(vops.SMult(constB.Value(), constA.Vector())), nil
	}
	if constA.Type() == Value && constB.Type() == Vector {
		return MakeConst(vops.SMult(constA.Value(), constB.Vector())), nil
	}
	if constA.Type() == Vector && constB.Type() == Vector {
		vectorA := constA.Vector()
		vectorB := constB.Vector()
		if vectorA.Space() == v.RowSpace {
			vector, err := vops.InnerProduct(vectorA, vectorB)
			if err != nil {
				return nil, err
			}
			return MakeConst(vector), nil
		}
		matrix, err := vops.OuterProduct(vectorA, vectorB)
		if err != nil {
			return nil, err
		}
		return MakeConst(matrix), nil
	}
	if constA.Type() == Matrix && constB.Type() == Value {
		return MakeConst(mops.SMult(constB.Value(), constA.Matrix())), nil
	}
	if constA.Type() == Value && constB.Type() == Matrix {
		return MakeConst(mops.SMult(constA.Value(), constB.Matrix())), nil
	}
	if constA.Type() == Vector && constB.Type() == Matrix {
		vector, err := mops.VMMult(constA.Vector(), constB.Matrix())
		if err != nil {
			return nil, err
		}
		return MakeConst(vector), nil
	}
	if constA.Type() == Matrix && constB.Type() == Vector {
		vector, err := mops.MVMult(constB.Vector(), constA.Matrix())
		if err != nil {
			return nil, err
		}
		return MakeConst(vector), nil
	}
	matrix, err := mops.MultSimple(constA.Matrix(), constB.Matrix())
	if err != nil {
		return nil, err
	}
	return MakeConst(matrix), nil
}

// MustMult is the same as Mult but will panic
func MustMult(constA Const, constB Const) Const {
	constant, err := Mult(constA, constB)
	if err != nil {
		panic(err)
	}
	return constant
}

// Pow will raise one constant to the power of another constant
// for matrix consts, it is assumed that the value it will be raised to is an integer
func Pow(constA Const, constB Const) (Const, error) {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Pow(constA.Value(), constB.Value())), nil
	}
	if constA.Type() == Matrix && constB.Type() == Value {
		matrix, err := mops.Pow(constA.Matrix(), int(constB.Value().Real()))
		if err != nil {
			return nil, err
		}
		return MakeConst(matrix), nil
	}
	return nil, errors.New("One or More Types are not supported")
}

// MustPow is the same as Pow but will panic
func MustPow(constA Const, constB Const) Const {
	con, err := Pow(constA, constB)
	if err != nil {
		panic(err)
	}
	return con
}

// Sqrt will find the square root of a Const
func Sqrt(constant Const) (Const, error) {
	if constant.Type() == Value {
		return MakeConst(gcvops.Sqrt(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Sqrt")
}

// MustSqrt is the same as Sqrt but will panic
func MustSqrt(constant Const) Const {
	con, err := Sqrt(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Conj will find the conjuage of a Const
func Conj(constant Const) (Const, error) {
	if constant.Type() == Value {
		return MakeConst(gcvops.Conj(constant.Value())), nil
	}
	if constant.Type() == Vector {
		return MakeConst(v.MakeConjVector(constant.Vector())), nil
	}
	if constant.Type() == Matrix {
		return MakeConst(m.MakeConjMatrix(constant.Matrix())), nil
	}
	return nil, errors.New("Const Type is not supported for Conj")
}

// MustConj is the same as Conj but will panic
func MustConj(constant Const) Const {
	con, err := Conj(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Sin will find the sine of a Const
func Sin(constant Const) (Const, error) {
	if constant.Type() == Value {
		return MakeConst(gcvops.Sin(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Sin")
}

// MustSin is the same as Sin but will panic
func MustSin(constant Const) Const {
	con, err := Sin(constant)
	if err != nil {
		panic(err)
	}
	return con
}

// Cos will find the cosine of a Const
func Cos(constant Const) (Const, error) {
	if constant.Type() == Value {
		return MakeConst(gcvops.Cos(constant.Value())), nil
	}
	return nil, errors.New("Const Type is not supported for Cos")
}

// MustCos is the same as Cos but will panic
func MustCos(constant Const) Const {
	con, err := Cos(constant)
	if err != nil {
		panic(err)
	}
	return con
}
