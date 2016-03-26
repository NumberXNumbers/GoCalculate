package gcv

import "testing"

func TestNewValueSetRawValueGetValueTypePrint(t *testing.T) {
    var floatTestValue float64
    valueTestA := NewValue(floatTestValue)

    if valueTestA.GetValueType() != "float64" {
        t.Fail()
    }

    valueTestA.PrintType()
    valueTestA.PrintRaw()

    var intTestValue int
    valueTestB := NewValue(intTestValue)

    if valueTestB.GetValueType() != "int" {
        t.Fail()
    }

    var complexTestValue complex128
    valueTestC := NewValue(complexTestValue)

    if valueTestC.GetValueType() != "complex128" {
        t.Fail()
    }

    var stringTestValue string
    valueTestD := NewValue(stringTestValue)

    if valueTestD.GetValueType() != "Unknown" {
        t.Fail()
    }
}

func TestIntValue(t *testing.T) {
    floatTestValue := 5.0
    valueTestA := NewValue(floatTestValue)

    if valueTestA.Int() != 5 {
        t.Fail()
    }

    intTestValue := 5
    valueTestB := NewValue(intTestValue)

    if valueTestB.Int() != 5 {
        t.Fail()
    }

    complexTestValue := 5.0+2.0i
    valueTestC := NewValue(complexTestValue)

    if valueTestC.Int() != 5 {
        t.Fail()
    }

    stringTestValue := "Hi"
    valueTestD := NewValue(stringTestValue)

    if valueTestD.Int() != 0 {
        t.Fail()
    }
}

func TestFloat64Value(t *testing.T) {
    floatTestValue := 5.0
    valueTestA := NewValue(floatTestValue)

    if valueTestA.Float64() != 5.0 {
        t.Fail()
    }

    intTestValue := 5
    valueTestB := NewValue(intTestValue)

    if valueTestB.Float64() != 5.0 {
        t.Fail()
    }

    complexTestValue := 5.0+2.0i
    valueTestC := NewValue(complexTestValue)

    if valueTestC.Float64() != 5.0 {
        t.Fail()
    }

    stringTestValue := "Hi"
    valueTestD := NewValue(stringTestValue)

    if valueTestD.Float64() != 0 {
        t.Fail()
    }
}

func TestComplex128(t *testing.T) {
    floatTestValue := 5.0
    valueTestA := NewValue(floatTestValue)

    if valueTestA.Complex128() != 5.0+0.0i {
        t.Fail()
    }

    intTestValue := 5
    valueTestB := NewValue(intTestValue)

    if valueTestB.Complex128() != 5.0+0.0i {
        t.Fail()
    }

    complexTestValue := 5.0+2.0i
    valueTestC := NewValue(complexTestValue)

    if valueTestC.Complex128() != 5.0+2.0i {
        t.Fail()
    }

    stringTestValue := "Hi"
    valueTestD := NewValue(stringTestValue)

    if valueTestD.Complex128() != 0+0i {
        t.Fail()
    }
}
