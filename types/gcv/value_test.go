package gcv

import (
	"reflect"
	"testing"
)

func TestMakeValueSetValueGetValueTypePrint(t *testing.T) {
	var floatTestValueA float64
	valueTestAa := MakeValue(floatTestValueA)

	if valueTestAa.GetValueType() != Float {
		t.Fail()
	}

	var floatTestValueB float32
	valueTestAb := MakeValue(floatTestValueB)

	if valueTestAb.GetValueType() != Float {
		t.Fail()
	}

	var intTestValueA int
	valueTestBa := MakeValue(intTestValueA)

	if valueTestBa.GetValueType() != Int {
		t.Fail()
	}

	var intTestValueB int32
	valueTestBb := MakeValue(intTestValueB)

	if valueTestBb.GetValueType() != Int {
		t.Fail()
	}

	var intTestValueC int64
	valueTestBc := MakeValue(intTestValueC)

	if valueTestBc.GetValueType() != Int {
		t.Fail()
	}

	var complexTestValueA complex128
	valueTestCa := MakeValue(complexTestValueA)

	if valueTestCa.GetValueType() != Complex {
		t.Fail()
	}

	var complexTestValueB complex64
	valueTestB := MakeValue(complexTestValueB)

	if valueTestB.GetValueType() != Complex {
		t.Fail()
	}

	var stringTestValue string
	valueTestD := MakeValue(stringTestValue)

	if valueTestD.GetValueType() != Int {
		t.Fail()
	}

	valueTestE := NewValue()

	if valueTestE.GetValueType() != Int {
		t.Fail()
	}
}

func TestIntValue(t *testing.T) {
	floatTestValue := 5.0
	valueTestA := MakeValue(floatTestValue)

	if valueTestA.Int() != 5 {
		t.Fail()
	}

	intTestValue := 5
	valueTestB := MakeValue(intTestValue)

	if valueTestB.Int() != 5 {
		t.Fail()
	}

	complexTestValue := 5.0 + 2.0i
	valueTestC := MakeValue(complexTestValue)

	if valueTestC.Int() != 5 {
		t.Fail()
	}

	stringTestValue := "Hi"
	valueTestD := MakeValue(stringTestValue)

	if valueTestD.Int() != 0 {
		t.Fail()
	}
}

func TestFloat64Value(t *testing.T) {
	floatTestValue := 5.0
	valueTestA := MakeValue(floatTestValue)

	if valueTestA.Float64() != 5.0 {
		t.Fail()
	}

	intTestValue := 5
	valueTestB := MakeValue(intTestValue)

	if valueTestB.Float64() != 5.0 {
		t.Fail()
	}

	complexTestValue := 5.0 + 2.0i
	valueTestC := MakeValue(complexTestValue)

	if valueTestC.Float64() != 5.0 {
		t.Fail()
	}

	stringTestValue := "Hi"
	valueTestD := MakeValue(stringTestValue)

	if valueTestD.Float64() != 0 {
		t.Fail()
	}
}

func TestComplex128(t *testing.T) {
	floatTestValue := 5.0
	valueTestA := MakeValue(floatTestValue)

	if valueTestA.Complex128() != 5.0+0.0i {
		t.Fail()
	}

	intTestValue := 5
	valueTestB := MakeValue(intTestValue)

	if valueTestB.Complex128() != 5.0+0.0i {
		t.Fail()
	}

	complexTestValue := 5.0 + 2.0i
	valueTestC := MakeValue(complexTestValue)

	if valueTestC.Complex128() != 5.0+2.0i {
		t.Fail()
	}

	stringTestValue := "Hi"
	valueTestD := MakeValue(stringTestValue)

	if valueTestD.Complex128() != 0+0i {
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
	testValues := MakeValuesPure(0.5, 0.6)

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

	if copyA.GetValueType() != Float {
		t.Fail()
	}

	if !reflect.DeepEqual(copyA, testValueA) {
		t.Fail()
	}

	testValueB := MakeValue(0.6)
	values := MakeValues(testValueA, testValueB)

	testValues := values.Copy()

	testValueB.SetValue(4)

	if !reflect.DeepEqual(testValues, values) {
		t.Fail()
	}

	testValueC := MakeValue(1 + 5i)
	copyC := testValueC.Copy()

	if copyC.GetValueType() != Complex {
		t.Fail()
	}

	if !reflect.DeepEqual(copyC, testValueC) {
		t.Fail()
	}

	copyC.SetValue(2)

	if copyC.GetValueType() != Int {
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
