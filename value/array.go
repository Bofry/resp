package value

import (
	"math"
	"math/big"
	"reflect"
)

var (
	_ Value = new(Array)
)

type Array struct {
	content []Value
}

func NewArray(values ...Value) *Array {
	return &Array{
		content: values,
	}
}

// Type implements Value.
func (*Array) Type() string {
	return __ARRAY_VALUE
}

// IsNull implements Value.
func (v *Array) IsNull() bool {
	return v == nil
}

// Equals implements Value.
func (v *Array) Equals(other Value) bool {
	if v == nil {
		return other == nil || other.IsNull()
	}
	if other == nil || other.IsNull() {
		return false
	}

	switch other.(type) {
	case *Array:
		return reflect.DeepEqual(v, other)
	}
	return false
}

// Array implements Value.
func (v *Array) Array() ([]Value, bool) {
	return v.content, true
}

// BigNumber implements Value.
func (*Array) BigNumber() (*big.Rat, bool) {
	return nil, false
}

// Error implements Value.
func (*Array) Error() (string, bool) {
	return "", false
}

// Float implements Value.
func (*Array) Float() (float64, bool) {
	return math.NaN(), false
}

// Integer implements Value.
func (*Array) Integer() (int64, bool) {
	return 0, false
}

// Map implements Value.
func (v *Array) Map() (map[string]Value, bool) {
	if v == nil || v.IsNull() {
		return nil, true
	}

	// odd?
	if (len(v.content) & 1) == 1 {
		return nil, false
	}

	var content = make(map[string]Value)
	for i := 0; i < len(v.content); i += 2 {
		k, ok := v.content[i].String()
		if !ok {
			return nil, false
		}
		val := v.content[i+1]
		content[k] = val
	}
	return content, true
}

// String implements Value.
func (*Array) String() (string, bool) {
	return "", false
}
