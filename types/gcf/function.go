package gcf

import (
	"fmt"

	"github.com/NumberXNumbers/GoCalculate/types/gcv"
	"github.com/NumberXNumbers/GoCalculate/types/m"
	"github.com/NumberXNumbers/GoCalculate/types/v"
)

var (
	leftParen  = "("
	rightParen = ")"
	pow        = "^"
	unaryFuncs = map[string]func(Const) Const{
		"Sin":  Sin,
		"Sqrt": Sqrt,
		"Cos":  Cos,
	}
	binaryFuncs = map[string]func(Const, Const) Const{
		"+": Add,
		"-": Sub,
		"*": Mult,
		"/": Div,
		pow: Pow,
	}
	orderOfOperations = map[string]uint{
		pow: 3,
		"*": 2,
		"/": 2,
		"+": 1,
		"-": 1,
	}
)

// Function is the function type for GoCalculate
type Function struct {
	inputTypes map[int]Type
	Args       []interface{}
	varNum     map[Var]int
	numVars    int
	regVars    []Var
}

func (f *Function) getVar(i int) Var {
	if f.typeInput(i) == Constant {
		return newConstVar(f.Args[i].(Const))
	} else if f.typeInput(i) == Variable {
		return f.Args[i].(Var)
	}
	e := fmt.Sprintf("Index %d, is not of type Var or Const", i)
	panic(e)
}

func (f *Function) getOp(i int) string {
	if f.typeInput(i) == Operation {
		return f.Args[i].(string)
	}
	e := fmt.Sprintf("Index %d, is not of type Operations", i)
	panic(e)
}

func (f *Function) typeInput(x int) Type { return f.inputTypes[x] }

// Eval will evaluate a function
func (f *Function) Eval(inputs ...interface{}) Const {
	if len(inputs) != f.numVars {
		panic("Number of inputs is not equal to the number of variables in function")
	}
	var operand1 Const
	var operand2 Const
	var operandStack []Const
	i := 0
	for i < len(f.Args) {
		if f.typeInput(i) == Constant || f.typeInput(i) == Variable {
			variable := f.getVar(i)
			if len(inputs) != 0 {
				operandStack = append(operandStack, variable.Eval(inputs[f.varNum[variable]]))
			} else {
				operandStack = append(operandStack, variable.Eval(0))
			}
		} else if f.typeInput(i) == Operation {
			operation := f.getOp(i)
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

// MakeFunc will make a gcf function struct
func MakeFunc(regVars []Var, inputs ...interface{}) *Function {
	function := new(Function)

	function.regVars = regVars
	var varNum = make(map[Var]int)
	var numVars int
	var tempOpsStack []string
	var postfixStack []interface{}
	for i, v := range regVars {
		if _, ok := varNum[v]; !ok {
			varNum[v] = numVars
			numVars++
			continue
		}
		e := fmt.Sprintf("Error registering variables. Variable at index %d, is a duplicate", i)
		panic(e)
	}
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
		case Var:
			if _, ok := varNum[n.(Var)]; !ok {
				e := fmt.Sprintf("Variable at index %d, was not registered", i)
				panic(e)
			}
			postfixStack = append(postfixStack, n)
			inputType[topIndexInPostfixStack+1] = Variable
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

	function.inputTypes = inputType
	function.numVars = numVars
	function.varNum = varNum
	function.Args = postfixStack
	return function
}
