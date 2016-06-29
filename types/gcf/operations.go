package gcf

import (
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
	panic("One or More Types are not supported")
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
	panic("One or More Types are not supported")
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
	panic("One or More Types are not supported")
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

// Pow will raise one constant to the power of another constant
func Pow(constA Const, constB Const) Const {
	if constA.Type() == Value && constB.Type() == Value {
		return MakeConst(gcvops.Pow(constA.Value(), constB.Value()))
	}
	panic("One or More Types are not supported")
}

// Sqrt will find the square root of a Const
func Sqrt(constant Const) Const {
	if constant.Type() == Value {
		return MakeConst(gcvops.Sqrt(constant.Value()))
	}
	panic("Const Type is not supported for Sqrt")
}

// Sin will find the sine of a Const
func Sin(constant Const) Const {
	if constant.Type() == Value {
		return MakeConst(gcvops.Sin(constant.Value()))
	}
	panic("Const Type is not supported for Sin")
}

// Cos will find the cosine of a Const
func Cos(constant Const) Const {
	if constant.Type() == Value {
		return MakeConst(gcvops.Cos(constant.Value()))
	}
	panic("Const Type is not supported for Sin")
}

// Calculate will return Calculations
func Calculate(inputs ...interface{}) Const {
	var tempOpsStack []string
	var postfixStack []interface{}
	var inputType = make(map[int]Type)
	for i, n := range inputs {
		topIndexInPostfixStack := len(postfixStack) - 1
		switch n.(type) {
		case string:
			operation := n.(string)
			var finishComparing bool
			topIndexInTempOpsStack := len(tempOpsStack) - 1
			if len(tempOpsStack) == 0 ||
				(tempOpsStack[topIndexInTempOpsStack] == leftParen && operation != rightParen) {
				tempOpsStack = append(tempOpsStack, operation)
			} else if operation == leftParen {
				tempOpsStack = append(tempOpsStack, operation)
			} else if operation == rightParen {
				for !finishComparing {
					if len(tempOpsStack) == 0 {
						panic("Mismatch of Parentheses found")
					}
					topOperationInTempOpsStack := tempOpsStack[topIndexInTempOpsStack]
					if topOperationInTempOpsStack == leftParen {
						tempOpsStack = tempOpsStack[:topIndexInTempOpsStack]
						finishComparing = true
					} else {
						inputType[topIndexInPostfixStack+1] = Operation
						postfixStack, tempOpsStack = append(postfixStack, topOperationInTempOpsStack), tempOpsStack[:topIndexInTempOpsStack]
					}
					topIndexInTempOpsStack = len(tempOpsStack) - 1
					topIndexInPostfixStack = len(postfixStack) - 1
				}
			} else {
				topOperationInTempOpsStack := tempOpsStack[topIndexInTempOpsStack]
				var isPreviousUnary bool
				var isUnary bool
				if _, ok := unaryFuncs[topOperationInTempOpsStack]; ok {
					isPreviousUnary = true
				}
				if _, ok := unaryFuncs[operation]; ok {
					isUnary = true
				}
				if isPreviousUnary || orderOfOperations[operation] < orderOfOperations[topOperationInTempOpsStack] {
					for !finishComparing {
						if isUnary && isPreviousUnary {
							tempOpsStack = append(tempOpsStack, operation)
							finishComparing = true
						} else if (topOperationInTempOpsStack == leftParen ||
							orderOfOperations[operation] > orderOfOperations[topOperationInTempOpsStack] ||
							isUnary) &&
							!isPreviousUnary {
							tempOpsStack = append(tempOpsStack, operation)
							finishComparing = true
						} else if orderOfOperations[operation] == orderOfOperations[topOperationInTempOpsStack] {
							if operation == pow {
								tempOpsStack = append(tempOpsStack, operation)
								finishComparing = true
							} else {
								inputType[topIndexInPostfixStack+1] = Operation
								postfixStack, tempOpsStack = append(postfixStack, topOperationInTempOpsStack), tempOpsStack[:topIndexInTempOpsStack]
								topIndexInTempOpsStack = len(tempOpsStack) - 1
							}
						} else if orderOfOperations[operation] < orderOfOperations[topOperationInTempOpsStack] || isPreviousUnary {
							inputType[topIndexInPostfixStack+1] = Operation
							postfixStack, tempOpsStack = append(postfixStack, topOperationInTempOpsStack), tempOpsStack[:topIndexInTempOpsStack]
							topIndexInTempOpsStack = len(tempOpsStack) - 1
						}

						if len(tempOpsStack) == 0 {
							tempOpsStack = append(tempOpsStack, operation)
							finishComparing = true
						} else {
							topOperationInTempOpsStack = tempOpsStack[topIndexInTempOpsStack]
							topIndexInPostfixStack = len(postfixStack) - 1
							if _, ok := unaryFuncs[topOperationInTempOpsStack]; !ok {
								isPreviousUnary = false
							}
						}
					}
				} else if orderOfOperations[operation] > orderOfOperations[topOperationInTempOpsStack] {
					tempOpsStack = append(tempOpsStack, operation)
				} else if orderOfOperations[operation] == orderOfOperations[topOperationInTempOpsStack] {
					if operation == pow {
						tempOpsStack = append(tempOpsStack, operation)
					} else {
						inputType[topIndexInPostfixStack+1] = Operation
						postfixStack, tempOpsStack = append(postfixStack, topOperationInTempOpsStack), tempOpsStack[:topIndexInTempOpsStack]
						tempOpsStack = append(tempOpsStack, operation)
					}
				}
			}
		case int, int32, int64, float32, float64, complex64, complex128, gcv.Value, v.Vector, m.Matrix:
			postfixStack = append(postfixStack, MakeConst(inputs[i]))
			inputType[topIndexInPostfixStack+1] = Constant
		case Const:
			postfixStack = append(postfixStack, n)
			inputType[topIndexInPostfixStack+1] = Constant
		default:
			panic("Input type not supported")
		}
	}

	for len(tempOpsStack) > 0 {
		topIndexInTempOpsStack := len(tempOpsStack) - 1
		topIndexInPostfixStack := len(postfixStack) - 1
		var operation string
		operation, tempOpsStack = tempOpsStack[topIndexInTempOpsStack], tempOpsStack[:topIndexInTempOpsStack]
		if operation == "(" {
			panic("Mismatch of Parentheses found")
		}
		inputType[topIndexInPostfixStack+1] = Operation
		postfixStack = append(postfixStack, operation)
	}

	var operand1 Const
	var operand2 Const
	var operandStack []Const
	i := 0
	for i < len(postfixStack) {
		if inputType[i] == Constant {
			operandStack = append(operandStack, postfixStack[i].(Const))
		} else if inputType[i] == Operation {
			operation := postfixStack[i].(string)
			if h, ok := unaryFuncs[operation]; ok {
				if len(operandStack) == 0 {
					panic("Not enough operands")
				} else {
					operand1, operandStack = operandStack[len(operandStack)-1], operandStack[:len(operandStack)-1]
					result := h(operand1)
					operandStack = append(operandStack, result)
				}
			} else if h, ok := binaryFuncs[operation]; ok {
				if len(operandStack) < 2 {
					panic("Not enough operands")
				} else {
					operand2, operandStack = operandStack[len(operandStack)-1], operandStack[:len(operandStack)-1]
					operand1, operandStack = operandStack[len(operandStack)-1], operandStack[:len(operandStack)-1]
					result := h(operand1, operand2)
					operandStack = append(operandStack, result)

				}
			} else {
				panic("Operation not supported")
			}
		}
		i++
	}

	if len(operandStack) > 1 {
		panic("To many operands left over after calculation")
	}

	return operandStack[0]
}
