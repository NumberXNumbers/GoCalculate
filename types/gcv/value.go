package gcv

import "fmt"

const (
	// Complex is the complex128 value
	Complex = "complex128"
	// Float is the float64 value
	Float = "float64"
	// Int is the int value
	Int = "int"
)

// Value is the main return type for the GoCalulate Framework
type Value interface {
	// returns the int of the raw value
	Int() int

	// returns the float64 of the raw value
	Float64() float64

	// returns the complex128 of the raw value
	Complex128() complex128

	// allows you to reset the raw value
	SetRawValue(val interface{})

	// returns the type of raw value
	GetValueType() string

	Copy() Value

	// prints the type of the value value
	PrintType()

	// prints the raw value
	PrintRaw()
}

type value struct {
	rawValue  interface{}
	valueType string
}

func (v *value) Int() int {
	var intValue int
	switch v.valueType {
	case Float:
		intValue = int(v.rawValue.(float64))
	case Complex:
		intValue = int(real(v.rawValue.(complex128)))
	default:
		intValue = v.rawValue.(int)
	}
	return intValue
}

func (v *value) Float64() float64 {
	var floatValue float64
	switch v.valueType {
	case Int:
		floatValue = float64(v.rawValue.(int))
	case Complex:
		floatValue = float64(real(v.rawValue.(complex128)))
	default:
		floatValue = v.rawValue.(float64)
	}
	return floatValue
}

func (v *value) Complex128() complex128 {
	var complexValue complex128
	switch v.valueType {
	case Int:
		complexValue = complex128(complex(float64(v.rawValue.(int)), 0.0))
	case Float:
		complexValue = complex128(complex(v.rawValue.(float64), 0.0))
	default:
		complexValue = v.rawValue.(complex128)
	}
	return complexValue
}

func (v *value) GetValueType() string { return v.valueType }

func (v *value) SetRawValue(val interface{}) {
	switch val.(type) {
	case int:
		v.valueType = Int
		break
	case int32:
		v.valueType = Int
		val = int(val.(int32))
		break
	case int64:
		v.valueType = Int
		val = int(val.(int64))
		break
	case float64:
		v.valueType = Float
		break
	case float32:
		v.valueType = Float
		val = float64(val.(float32))
		break
	case complex128:
		v.valueType = Complex
		break
	case complex64:
		v.valueType = Complex
		val = complex128(val.(complex64))
		break
	default:
		v.valueType = Int
		val = 0
	}
	v.rawValue = val
}

func (v *value) Copy() Value {
	value := new(value)
	value.SetRawValue(v.rawValue)
	value.valueType = v.GetValueType()
	return value
}

func (v *value) PrintType() { fmt.Println(v.valueType) }

func (v *value) PrintRaw() { fmt.Println(v.rawValue) }

// NewValue returns a new Value
func NewValue(val interface{}) Value {
	value := new(value)
	value.SetRawValue(val)
	return value
}
