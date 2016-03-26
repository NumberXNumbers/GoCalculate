package gcv

import (
    "testing"
    "reflect"
)

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

func TestNewValuesAndValues(t *testing.T) {
    testValueA := NewValue(0.5)
    testValueB := NewValue(0.6)
    testValues := NewValues(testValueA, testValueB)

    if !reflect.DeepEqual(testValues.Values()[0], testValueA) || !reflect.DeepEqual(testValues.Values()[1], testValueB) {
        t.Fail()
    }
}

func TestGetandSetValues(t *testing.T) {
    testValueA := NewValue(0.5)
    testValueB := NewValue(0.6)
    testValues := NewValues(testValueA, testValueB)

    if !reflect.DeepEqual(testValues.Get(0), testValueA) || !reflect.DeepEqual(testValues.Get(1), testValueB) {
        t.Fail()
    }

    newValueA := NewValue(1+0.3i)
    newValueB := NewValue(8)

    testValues.Set(0, newValueA)
    testValues.Set(1, newValueB)

    if !reflect.DeepEqual(testValues.Get(0), newValueA) || !reflect.DeepEqual(testValues.Get(1), newValueB) {
        t.Fail()
    }
}

func TestCopyValues(t *testing.T) {
    testValueA := NewValue(0.5)
    testValueB := NewValue(0.6)
    values := NewValues(testValueA, testValueB)

    testValues := values.Copy()

    if !reflect.DeepEqual(testValues, values) {
        t.Fail()
    }
}

func TestAppendValue(t *testing.T) {
    testValueA := NewValue(0.5)
    testValueB := NewValue(0.6)
    testValues := NewValues(testValueA, testValueB)

    testValueC := NewValue(0.7)
    testValues.Append(testValueC)

    if !reflect.DeepEqual(testValues.Values()[2], testValueC) {
        t.Fail()
    }
}

func TestSubsetAndLenValues(t *testing.T) {
    testValueA := NewValue(0.5)
    testValueB := NewValue(0.6)
    testValueC := NewValue(0.7)
    values := NewValues(testValueA, testValueB, testValueC)

    lenA := values.Len()

    testValues := values.Subset(1, 2)

    lenB := testValues.Len()

    if !reflect.DeepEqual(testValues.Values()[0], testValueB) || !reflect.DeepEqual(testValues.Values()[1], testValueC) {
        t.Fail()
    }

    if lenB != lenA-1 {
        t.Fail()
    }
}

func TestIndexOfValues(t *testing.T) {
    testValueA := NewValue(0.5)
    testValueB := NewValue(0.6)
    testValueC := NewValue(0.7)
    values := NewValues(testValueA, testValueB, testValueC)

    indexOfC := values.IndexOf(testValueC)

    if indexOfC != 2 {
        t.Fail()
    }

    indexOfNone := values.IndexOf(NewValue(0.8))

    if indexOfNone != -1 {
        t.Fail()
    }
}
