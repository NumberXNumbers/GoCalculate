package gcv

// Values works as GoCalulate's main list of values
type Values interface {
	// Returns the index of val. If val is not in Values it returns -1.
	IndexOf(val Value) int

	// Set the Value val at index
	Set(index int, val Value)

	// Append Value val to Values
	Append(val Value)

	// Returns the Value at index
	Get(index int) Value

	// Returns the Raw Value slice. Used mainly for use with range
	Values() []Value

	// Returns a subset of Values from start to finish
	Subset(start, finish int) Values

	// Returns a copy of Values.
	Copy() Values

	// Return Length of Values
	Len() int

	// Returns the Core Type of the Values. i.e the highest ranking Type
	Type() string
}

type values struct {
	vals     []Value
	length   int
	coreType string
}

func (v *values) Type() string { return v.coreType }

func (v *values) setValues(vals []Value) {
	v.vals = make([]Value, len(vals))
	v.length = len(vals)
	v.coreType = Int
	for index, val := range vals {
		v.vals[index] = val.Copy()
		if v.Type() != Complex && len(v.Type()) < len(val.GetValueType()) {
			v.coreType = val.GetValueType()
		}
	}
}

func (v *values) Len() int { return v.length }

func (v *values) Values() []Value { return v.vals }

func (v *values) Set(index int, val Value) { v.vals[index] = val }

func (v *values) Get(index int) Value { return v.vals[index] }

func (v *values) Append(val Value) {
	v.setValues(append(v.Values(), val))
}

func (v *values) Copy() Values {
	vals := new(values)
	vElements := make([]Value, len(v.vals))
	for index, val := range v.Values() {
		vElements[index] = val.Copy()
	}
	vals.length = v.Len()
	vals.coreType = v.Type()
	vals.vals = vElements
	return vals
}

func (v *values) Subset(start, finish int) Values {
	vals := new(values)
	subVals := make([]Value, len(v.vals[start:finish+1]))
	copy(subVals, v.vals[start:finish+1])
	vals.setValues(subVals)
	return vals
}

func (v *values) IndexOf(val Value) int {
	for index, value := range v.Values() {
		if value.GetValueType() == Complex && value.Complex128() == val.Complex128() {
			return index
		}
		if value.GetValueType() == Float && value.Float64() == val.Float64() {
			return index
		}
		if value.GetValueType() == Int && value.Int() == val.Int() {
			return index
		}
	}
	return -1
}

// NewValues returns a Values type
func NewValues(vals ...Value) Values {
	newValues := new(values)
	if vals == nil {
		vals = make([]Value, 0)
	}
	newValues.setValues(vals)
	return newValues
}
