package gcf

import "fmt"

var (
	binaryFuncs = map[string]func(Const, Const) Const{
		"+": Add,
		"-": Sub,
		"*": Mult,
		"/": Div,
	}
	orderOfOperations = map[string]uint{
		"exp": 3,
		"*":   2,
		"/":   2,
		"+":   1,
		"-":   1,
	}
	leftParen  = "("
	rightParen = ")"
)

// Function is the function type for GoCalculate
type Function struct {
	inputTypes map[int]Type
	args       []interface{}
	varNum     map[Var]int
	numVars    int
	regVars    []Var
}

func (f *Function) getVar(i int) Var {
	if f.typeInput(i) == Constant {
		return newConstVar(f.args[i].(Const))
	} else if f.typeInput(i) == Variable {
		return f.args[i].(Var)
	}
	e := fmt.Sprintf("Index %d, is not of type Var or Const", i)
	panic(e)
}

func (f *Function) getOp(i int) string {
	if f.typeInput(i) == Operation {
		return f.args[i].(string)
	}
	e := fmt.Sprintf("Index %d, is not of type Operations", i)
	panic(e)
}

func (f *Function) typeInput(x int) Type { return f.inputTypes[x] }

func evalHelper(f *Function, inputs ...interface{}) Const {
	var leftParens uint
	var rightParens uint

	var operandStack []Const
	var operatorStack []string

	i := 0
	for i < len(f.args) {
		if f.typeInput(i) == Constant || f.typeInput(i) == Variable {
			variable := f.getVar(i)
			operandStack = append(operandStack, variable.Eval(inputs[f.varNum[variable]]))
		} else if f.typeInput(i) == Operation {
			operation := f.getOp(i)
			if operation == leftParen {
				var subArgs []interface{}
				leftParens++
				for index := i + 1; index < len(f.args); index++ {
					if f.typeInput(index) == Operation && f.getOp(index) == rightParen {
						rightParens++
					} else if f.typeInput(index) == Operation && f.getOp(index) == leftParen {
						leftParens++
					}
					if rightParens != leftParens {
						subArgs = append(subArgs, f.args[index])
						continue
					}
					break
				}
				i = i + len(subArgs) + 1
				h := MakeFunc(f.regVars, subArgs...)
				result := evalHelper(h, inputs...)
				operandStack = append(operandStack, result)
			} else {
				operatorStack = append(operatorStack, f.getOp(i))
			}
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

// Eval will evaluate a function
func (f *Function) Eval(inputs ...interface{}) Const {
	if len(inputs) != f.numVars {
		panic("Number of inputs is not equal to the number of variables in function")
	}

	return evalHelper(f, inputs...)
}

// MakeFunc will make a gcf function struct
func MakeFunc(regVars []Var, inputs ...interface{}) *Function {
	function := new(Function)

	function.regVars = regVars
	var varNum = make(map[Var]int)
	var numVars int
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
		switch n.(type) {
		case string:
			inputType[i] = Operation
		case Const:
			inputType[i] = Constant
		case Var:
			if _, ok := varNum[n.(Var)]; !ok {
				e := fmt.Sprintf("Variable at index %d, was not registered", i)
				panic(e)
			}
			inputType[i] = Variable
		default:
			panic("Input type not supported")
		}
	}
	function.inputTypes = inputType
	function.numVars = numVars
	function.varNum = varNum
	function.args = inputs
	return function
}
