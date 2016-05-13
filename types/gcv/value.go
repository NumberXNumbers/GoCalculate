package gcv

// Type is the value type of Value
type Type int

const (
	// Real is for a real value
	Real Type = iota
	// Complex is for a complex value
	Complex
)

// Value is the main return type for the GoCalulate Framework
type Value interface {
	// returns the real part of a value
	Real() float64

	// returns the imaginary part of a value
	Imag() float64

	// returns the compelx representation of a value
	Complex() complex128

	// allows you to reset the value
	SetValue(val interface{})

	// returns the type of raw value
	Type() Type

	Copy() Value
}

type value struct {
	real      float64
	imaginary float64
	valueType Type
}

func (v *value) Real() float64 { return v.real }

func (v *value) Imag() float64 { return v.imaginary }

func (v *value) Complex() complex128 { return complex(v.Real(), v.Imag()) }

func (v *value) Type() Type { return v.valueType }

func (v *value) SetValue(val interface{}) {
	switch val.(type) {
	case int:
		v.valueType = Real
		v.real = float64(val.(int))
		v.imaginary = 0
		break
	case int32:
		v.valueType = Real
		v.real = float64(val.(int32))
		v.imaginary = 0
		break
	case int64:
		v.valueType = Real
		v.real = float64(val.(int64))
		v.imaginary = 0
		break
	case float64:
		v.valueType = Real
		floatVal := val.(float64)
		v.real = floatVal
		v.imaginary = 0
		break
	case float32:
		v.valueType = Real
		v.real = float64(val.(float32))
		v.imaginary = 0
		break
	case complex128:
		v.valueType = Complex
		complexVal := val.(complex128)
		v.real = real(complexVal)
		imagValue := imag(complexVal)
		if imagValue == 0 {
			v.valueType = Real
		}
		v.imaginary = imagValue
		break
	case complex64:
		v.valueType = Complex
		complexVal := complex128(val.(complex64))
		v.real = real(complexVal)
		imagValue := imag(complexVal)
		if imagValue == 0 {
			v.valueType = Real
		}
		v.imaginary = imagValue
		break
	default:
		v.valueType = Real
		v.real = 0
		v.imaginary = 0
	}
}

func (v *value) Copy() Value {
	value := new(value)
	value.valueType = v.Type()
	value.real = v.Real()
	value.imaginary = v.Imag()
	return value
}

// NewValue returns the 0 real value
func NewValue() Value {
	value := new(value)
	value.SetValue(0)
	return value
}

// MakeValue returns a Value with value val
func MakeValue(val interface{}) Value {
	value := new(value)
	value.SetValue(val)
	return value
}
