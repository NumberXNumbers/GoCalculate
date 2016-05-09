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
	SetValue(val interface{})

	// returns the type of raw value
	GetValueType() string

	Copy() Value

	// prints the type of the value value
	PrintType()

	// prints the raw value
	PrintRaw()
}

type value struct {
	rawValue     interface{}
	intValue     int
	floatValue   float64
	complexValue complex128
	valueType    string
}

func (v *value) Int() int { return v.intValue }

func (v *value) Float64() float64 { return v.floatValue }

func (v *value) Complex128() complex128 { return v.complexValue }

func (v *value) GetValueType() string { return v.valueType }

func (v *value) SetValue(val interface{}) {
	switch val.(type) {
	case int:
		v.valueType = Int
		intVal := val.(int)
		v.intValue = intVal
		v.floatValue = float64(intVal)
		v.complexValue = complex128(complex(float64(intVal), 0))
		break
	case int32:
		v.valueType = Int
		intVal := int(val.(int32))
		v.intValue = intVal
		v.floatValue = float64(intVal)
		v.complexValue = complex128(complex(float64(intVal), 0))
		break
	case int64:
		v.valueType = Int
		intVal := int(val.(int64))
		v.intValue = intVal
		v.floatValue = float64(intVal)
		v.complexValue = complex128(complex(float64(intVal), 0))
		break
	case float64:
		v.valueType = Float
		floatVal := val.(float64)
		v.intValue = int(floatVal)
		v.floatValue = floatVal
		v.complexValue = complex128(complex(floatVal, 0))
		break
	case float32:
		v.valueType = Float
		floatVal := float64(val.(float32))
		v.intValue = int(floatVal)
		v.floatValue = floatVal
		v.complexValue = complex128(complex(floatVal, 0))
		break
	case complex128:
		v.valueType = Complex
		complexVal := val.(complex128)
		v.intValue = int(real(complexVal))
		v.floatValue = real(complexVal)
		v.complexValue = complexVal
		break
	case complex64:
		v.valueType = Complex
		complexVal := complex128(val.(complex64))
		v.intValue = int(real(complexVal))
		v.floatValue = real(complexVal)
		v.complexValue = complexVal
		break
	default:
		v.valueType = Int
		val = 0
		v.intValue = 0
		v.floatValue = 0.0
		v.complexValue = 0 + 0i
	}
	v.rawValue = val
}

func (v *value) Copy() Value {
	value := new(value)
	value.SetValue(v.rawValue)
	value.valueType = v.GetValueType()
	return value
}

func (v *value) PrintType() { fmt.Println(v.valueType) }

func (v *value) PrintRaw() { fmt.Println(v.rawValue) }

// NewValue returns a new Value
func NewValue(val interface{}) Value {
	value := new(value)
	value.SetValue(val)
	return value
}
