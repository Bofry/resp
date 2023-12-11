package value

import (
	"math"
	"math/big"

	"github.com/cstockton/go-conv"
)

var (
	_ ScalarValue = new(String)
)

type String struct {
	content string
}

func NewString(v string) *String {
	return &String{
		content: v,
	}
}

// Type implements Value.
func (*String) Type() string {
	return __STRING_VALUE
}

// IsNull implements Value.
func (v *String) IsNull() bool {
	return v == nil
}

// Equals implements Value.
func (v *String) Equals(other Value) bool {
	if v == nil {
		return other == nil || other.IsNull()
	}
	if other == nil || other.IsNull() {
		return false
	}

	switch another := other.(type) {
	case *String:
		return v.content == another.content
	}
	return false
}

// Encoding implements ScalarValue.
func (*String) Encoding() string {
	return ""
}

// Array implements Value.
func (v *String) Array() ([]Value, bool) {
	return nil, false
}

// BigNumber implements Value.
func (v *String) BigNumber() (*big.Rat, bool) {
	r := new(big.Rat)
	_, ok := r.SetString(v.content)
	if !ok {
		return nil, false
	}
	return r, true
}

// Error implements Value.
func (*String) Error() (string, bool) {
	return "", false
}

// Float implements Value.
func (v *String) Float() (float64, bool) {
	f, err := conv.Float64(v)
	if err != nil {
		return math.NaN(), false
	}
	return f, true
}

// Integer implements Value.
func (v *String) Integer() (int64, bool) {
	i, err := conv.Int64(v)
	if err != nil {
		return 0, false
	}
	return i, true
}

// Map implements Value.
func (*String) Map() (map[string]Value, bool) {
	return nil, false
}

// String implements Value.
func (v *String) String() (string, bool) {
	return v.content, true
}
