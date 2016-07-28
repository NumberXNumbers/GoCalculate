package gcv

import (
	"reflect"
	"testing"
)

func TestMakeValueSetValueTypePrint(t *testing.T) {
	var floatTestValueA float64
	valueTestAa := MakeValue(floatTestValueA)

	if valueTestAa.Type() != Real {
		t.Fail()
	}

	var floatTestValueB float32
	valueTestAb := MakeValue(floatTestValueB)

	if valueTestAb.Type() != Real {
		t.Fail()
	}

	var intTestValueA int
	valueTestBa := MakeValue(intTestValueA)

	if valueTestBa.Type() != Real {
		t.Fail()
	}

	var intTestValueB int32
	valueTestBb := MakeValue(intTestValueB)

	if valueTestBb.Type() != Real {
		t.Fail()
	}

	var intTestValueC int64
	valueTestBc := MakeValue(intTestValueC)

	if valueTestBc.Type() != Real {
		t.Fail()
	}

	var complexTestValueA complex128
	valueTestCa := MakeValue(complexTestValueA)

	if valueTestCa.Type() != Real {
		t.Fail()
	}

	var complexTestValueB complex64
	valueTestB := MakeValue(complexTestValueB)

	if valueTestB.Type() != Real {
		t.Fail()
	}

	var stringTestValue string
	valueTestD := MakeValue(stringTestValue)

	if valueTestD.Type() != Real {
		t.Fail()
	}

	valueTestE := NewValue()

	if valueTestE.Type() != Real {
		t.Fail()
	}

	complexTestValueA = 1 + 5i
	valueTestF := MakeValue(complexTestValueA)

	if valueTestF.Type() != Complex {
		t.Fail()
	}

	complexTestValueB = 1 - 5i
	valueTestG := MakeValue(complexTestValueB)

	if valueTestG.Type() != Complex {
		t.Fail()
	}

	valueTestH := NewValue()
	value := MakeValue(complexTestValueA)
	valueTestH.Set(value)

	if valueTestH.Type() != Complex {
		t.Fail()
	}
}

func TestRealImagValue(t *testing.T) {
	floatTestValue := 5.0
	valueTestA := MakeValue(floatTestValue)

	if valueTestA.Real() != 5 {
		t.Fail()
	}

	intTestValue := 5
	valueTestB := MakeValue(intTestValue)

	if valueTestB.Real() != 5 {
		t.Fail()
	}

	complexTestValue := 5.0 + 2.0i
	valueTestC := MakeValue(complexTestValue)

	if valueTestC.Real() != 5 && valueTestC.Imag() != 2 {
		t.Fail()
	}

	stringTestValue := "Hi"
	valueTestD := MakeValue(stringTestValue)

	if valueTestD.Real() != 0 {
		t.Fail()
	}
}

func TestValue(t *testing.T) {
	floatTestValue := 5.0
	valueTestA := MakeValue(floatTestValue)

	if valueTestA.Real() != 5.0 {
		t.Fail()
	}

	intTestValue := 5
	valueTestB := MakeValue(intTestValue)

	if valueTestB.Real() != 5.0 {
		t.Fail()
	}

	complexTestValue := 5.0 + 2.0i
	valueTestC := MakeValue(complexTestValue)

	if valueTestC.Real() != 5.0 {
		t.Fail()
	}

	stringTestValue := "Hi"
	valueTestD := MakeValue(stringTestValue)

	if valueTestD.Real() != 0 {
		t.Fail()
	}
}

func TestComplex(t *testing.T) {
	floatTestValue := 5.0
	valueTestA := MakeValue(floatTestValue)

	if valueTestA.Complex() != 5.0+0.0i {
		t.Fail()
	}

	intTestValue := 5
	valueTestB := MakeValue(intTestValue)

	if valueTestB.Complex() != 5.0+0.0i {
		t.Fail()
	}

	complexTestValue := 5.0 + 2.0i
	valueTestC := MakeValue(complexTestValue)

	if valueTestC.Complex() != 5.0+2.0i {
		t.Fail()
	}

	stringTestValue := "Hi"
	valueTestD := MakeValue(stringTestValue)

	if valueTestD.Complex() != 0+0i {
		t.Fail()
	}
}

func TestMakeValuesAndValues(t *testing.T) {
	testValueA := MakeValue(0.5)
	testValueB := MakeValue(0.6)
	testValues := MakeValues(testValueA, testValueB)

	if !reflect.DeepEqual(testValues.Values()[0], testValueA) || !reflect.DeepEqual(testValues.Values()[1], testValueB) {
		t.Fail()
	}
}

func TestGetandSetValues(t *testing.T) {
	testValueA := MakeValue(0.5)
	testValueB := MakeValue(0.6)
	testValues := MakeValues(0.5, 0.6)

	if !reflect.DeepEqual(testValues.Get(0), testValueA) || !reflect.DeepEqual(testValues.Get(1), testValueB) {
		t.Fail()
	}

	newValueA := MakeValue(1 + 0.3i)
	newValueB := MakeValue(8)

	testValues.Set(0, newValueA)
	testValues.Set(1, newValueB)

	if !reflect.DeepEqual(testValues.Get(0), newValueA) || !reflect.DeepEqual(testValues.Get(1), newValueB) {
		t.Fail()
	}
}

func TestCopyValues(t *testing.T) {
	testValueA := MakeValue(0.5)
	copyA := testValueA.Copy()

	if copyA.Type() != Real {
		t.Fail()
	}

	if !reflect.DeepEqual(copyA, testValueA) {
		t.Fail()
	}

	testValueB := MakeValue(0.6)
	values := MakeValues(testValueA, testValueB)

	testValues := values.Copy()

	testValueB.Set(4)

	if !reflect.DeepEqual(testValues, values) {
		t.Fail()
	}

	testValueC := MakeValue(1 + 5i)
	copyC := testValueC.Copy()

	if copyC.Type() != Complex {
		t.Fail()
	}

	if !reflect.DeepEqual(copyC, testValueC) {
		t.Fail()
	}

	copyC.Set(2)

	if copyC.Type() != Real {
		t.Fail()
	}

	if reflect.DeepEqual(copyC, testValueC) {
		t.Fail()
	}
}

func TestAppendValue(t *testing.T) {
	testValueA := MakeValue(0.5)
	testValueB := MakeValue(0.6)
	testValuesA := MakeValues(testValueA, testValueB)

	testValueC := MakeValue(0.7)
	testValuesA.Append(testValueC)

	if !reflect.DeepEqual(testValuesA.Values()[2], testValueC) {
		t.Fail()
	}

	testValuesB := MakeValues()
	testValuesB.Append(testValueC)

	if !reflect.DeepEqual(testValuesB.Values()[0], testValueC) {
		t.Errorf("Expected %v, received %v", testValueC, testValuesB.Values()[0])
	}
}

func TestSubsetAndLenValues(t *testing.T) {
	testValueA := MakeValue(0.5)
	testValueB := MakeValue(0.6)
	testValueC := MakeValue(0.7)
	values := MakeValues(testValueA, testValueB, testValueC)

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
	testValueA := MakeValue(5)
	testValueB := MakeValue(0.6)
	testValueC := MakeValue(0.7 + 1i)
	values := MakeValues(testValueA, testValueB, testValueC)

	indexOfA := values.IndexOf(testValueA)
	if indexOfA != 0 {
		t.Fail()
	}

	indexOfB := values.IndexOf(testValueB)
	if indexOfB != 1 {
		t.Fail()
	}

	indexOfC := values.IndexOf(testValueC)
	if indexOfC != 2 {
		t.Fail()
	}

	indexOfNone := values.IndexOf(MakeValue(0.8))

	if indexOfNone != -1 {
		t.Fail()
	}
}
