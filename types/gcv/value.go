package gcv

import "fmt"

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

    // prints the type of the value value
    PrintType()

    // prints the raw value
    PrintRaw()
}

type value struct {
    rawValue interface{}
    valueType string
}

func (v *value) Int() int {
    var intValue int
    switch v.valueType {
    case "int":
        intValue = v.rawValue.(int)
    case "float64":
        intValue = int(v.rawValue.(float64))
    case "complex128":
        intValue = int(real(v.rawValue.(complex128)))
    default:
        intValue = 0
    }
    return intValue
}

func (v *value) Float64() float64 {
    var floatValue float64
    switch v.valueType {
    case "int":
        floatValue = float64(v.rawValue.(int))
    case "float64":
        floatValue = v.rawValue.(float64)
    case "complex128":
        floatValue = float64(real(v.rawValue.(complex128)))
    default:
        floatValue = 0.0
    }
    return floatValue
}

func (v *value) Complex128() complex128 {
    var complexValue complex128
    switch v.valueType {
    case "int":
        complexValue = complex128(complex(float64(v.rawValue.(int)), 0.0))
    case "float64":
        complexValue = complex128(complex(v.rawValue.(float64), 0.0))
    case "complex128":
        complexValue = v.rawValue.(complex128)
    default:
        complexValue = 0+0i
    }
    return complexValue
}

func (v *value) GetValueType() string { return v.valueType }

func (v *value) SetRawValue(val interface{}) {
    switch val.(type) {
    case int, int32, int64:
        v.valueType = "int"
    case float64, float32:
        v.valueType = "float64"
    case complex128, complex64:
        v.valueType = "complex128"
    default:
        v.valueType = "Unknown"
    }
    v.rawValue = val
}

func (v *value) PrintType() { fmt.Println(v.valueType) }

func (v *value) PrintRaw() { fmt.Println(v.rawValue) }

// NewValue returns a new Value
func NewValue(val interface{}) Value {
    value := new(value)
    value.SetRawValue(val)
    return value
}
