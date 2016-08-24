package gcf

import (
	"errors"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

// Calculate will return Calculations or error
func Calculate(inputs ...interface{}) (Const, error) {
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
						return nil, errors.New("Mismatch of Parentheses found")
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
			return nil, errors.New("Input type not supported")
		}
	}

	for len(tempOpsStack) > 0 {
		topIndexInTempOpsStack := len(tempOpsStack) - 1
		topIndexInPostfixStack := len(postfixStack) - 1
		var operation string
		operation, tempOpsStack = tempOpsStack[topIndexInTempOpsStack], tempOpsStack[:topIndexInTempOpsStack]
		if operation == "(" {
			return nil, errors.New("Mismatch of Parentheses found")
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
					return nil, errors.New("Not enough operands")
				}

				operand1, operandStack = operandStack[len(operandStack)-1], operandStack[:len(operandStack)-1]
				result, err := h(operand1)
				if err != nil {
					return nil, err
				}

				operandStack = append(operandStack, result)
			} else if h, ok := binaryFuncs[operation]; ok {
				if len(operandStack) < 2 {
					return nil, errors.New("Not enough operands")
				}

				operand2, operandStack = operandStack[len(operandStack)-1], operandStack[:len(operandStack)-1]
				operand1, operandStack = operandStack[len(operandStack)-1], operandStack[:len(operandStack)-1]
				result, err := h(operand1, operand2)
				if err != nil {
					return nil, err
				}

				operandStack = append(operandStack, result)
			} else {
				return nil, errors.New("Operation not supported")
			}
		}
		i++
	}

	if len(operandStack) > 1 {
		return nil, errors.New("To many operands left over after calculation")
	}

	return operandStack[0], nil
}

// MustCalculate is the same as calculate, but will panic
func MustCalculate(inputs ...interface{}) Const {
	constant, err := Calculate(inputs...)
	if err != nil {
		panic(err)
	}
	return constant
}
