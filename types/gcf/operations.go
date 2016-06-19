package gcf

import (
	"reflect"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/gcv/gcvops"
	"github.com/NumberXNumbers/GoCalculate/types/m"
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

// Calculate will return Calculations
func Calculate(inputs ...interface{}) Const {
	var leftParens uint
	var rightParens uint

	var operandStack []Const
	var operatorStack []string

	i := 0
	for i < len(inputs) {
		switch inputs[i].(type) {
		case Const:
			constant := inputs[i].(Const)
			operandStack = append(operandStack, constant)
		case int, int32, int64, float32, float64, complex64, complex128, gcv.Value, v.Vector, m.Matrix:
			constant := MakeConst(inputs[i])
			operandStack = append(operandStack, constant)
		case string:
			operation := inputs[i].(string)
			if operation == leftParen {
				var subArgs []interface{}
				leftParens++
				for index := i + 1; index < len(inputs); index++ {
					if v := reflect.ValueOf(inputs[index]); v.Kind() == reflect.String && inputs[index].(string) == rightParen {
						rightParens++
					} else if v := reflect.ValueOf(inputs[index]); v.Kind() == reflect.String && inputs[index].(string) == leftParen {
						leftParens++
					}
					if rightParens != leftParens {
						subArgs = append(subArgs, inputs[index])
						continue
					}
					break
				}
				i = i + len(subArgs) + 1
				result := Calculate(subArgs...)
				operandStack = append(operandStack, result)
			} else {
				operatorStack = append(operatorStack, operation)
			}
		default:
			panic("Unsupported type")
		}

		if len(operatorStack) > len(operandStack) || len(operandStack) > len(operatorStack)+1 {
			panic("Operators-Operand mismatch error")
		}

		if len(operandStack) == 3 && len(operatorStack) == 2 {
			var operand1 Const
			var operand2 Const
			var operator string
			if orderOfOperations[operatorStack[0]] >= orderOfOperations[operatorStack[1]] {
				operand1 = operandStack[0]
				operand2 = operandStack[1]
				operandStack = operandStack[1:]

				operator = operatorStack[0]
				operatorStack = operatorStack[1:]

				h := binaryFuncs[operator]

				result := h(operand1, operand2)

				operandStack[0] = result
			} else {
				operand1 = operandStack[1]
				operand2 = operandStack[2]
				operandStack = operandStack[:1]

				operator = operatorStack[1]
				operatorStack = operatorStack[:1]

				h := binaryFuncs[operator]

				result := h(operand1, operand2)

				operandStack = append(operandStack, result)
			}

		}

		i++
	}

	if len(operandStack) == 2 && len(operatorStack) == 1 {
		operand1 := operandStack[0]
		operand2 := operandStack[1]

		operator := operatorStack[0]

		h := binaryFuncs[operator]

		return h(operand1, operand2)
	}
	return operandStack[0]
}
