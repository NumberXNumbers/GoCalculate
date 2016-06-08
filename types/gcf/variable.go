package gcf

import "fmt"

var (
	varTypeMap = map[Type]string{
		Value:    "Value",
		Vector:   "Vector",
		Matrix:   "Matrix",
		Constant: "Constant",
	}
)

// Var is the GoCalculate variable type
type Var interface {
	// Eval will take a variable and return a constant.
	// the variable type, either type value, vector, matrix or const
	// must match the passed in variable type, else Eval will panic
	Eval(x interface{}) Const
}

type variable struct {
	varType Type
}

func (v *variable) Eval(x interface{}) Const {
	constant := MakeConst(x)
	if v.varType != constant.Type() {
		error := fmt.Sprintf("Expected %v, received %v", varTypeMap[v.varType], varTypeMap[constant.Type()])
		panic(error)
	}
	return constant
}

// NewVar will make a new Variable of type varType
func NewVar(varType Type) Var {
	variable := new(variable)
	variable.varType = varType
	return variable
}

// type constVar is for a variable that when eval is called,
// will only return it's type Const it was assigned at creation of Var
type constVar struct {
	constant Const
}

func (c *constVar) Eval(x interface{}) Const {
	return c.constant
}

func newConstVar(c Const) Var {
	constVar := new(constVar)
	constVar.constant = c
	return constVar
}
