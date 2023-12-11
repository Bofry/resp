package value

import (
	"math"
	"math/big"
	"strconv"

	"github.com/cstockton/go-conv"
)

var _ Value = Integer{}

type Integer struct {
	content int64
}

func NewInteger(v int64) Integer {
	return Integer{
		content: v,
	}
}

// Type implements Value.
func (Integer) Type() string {
	return __INTEGER_VALUE
}

// IsNull implements Value.
func (Integer) IsNull() bool {
	return false
}

// Equals implements Value.
func (v Integer) Equals(other Value) bool {
	switch another := other.(type) {
	case Integer:
		return v.content == another.content
	}
	return false
}

// Encoding implements Value.
func (Integer) Encoding() string {
	return ""
}

// Array implements Value.
func (Integer) Array() ([]Value, bool) {
	return nil, false
}

// BigNumber implements Value.
func (v Integer) BigNumber() (*big.Rat, bool) {
	r := new(big.Rat)
	r.SetInt64(v.content)
	return r, true
}

// Error implements Value.
func (Integer) Error() (string, bool) {
	return "", false
}

// Float implements Value.
func (v Integer) Float() (float64, bool) {
	f, err := conv.Float64(v)
	if err != nil {
		return math.NaN(), false
	}
	return f, true
}

// Integer implements Value.
func (v Integer) Integer() (int64, bool) {
	return v.content, true
}

// Map implements Value.
func (Integer) Map() (map[string]Value, bool) {
	return nil, false
}

// String implements Value.
func (v Integer) String() (string, bool) {
	return strconv.FormatInt(v.content, 10), true
}
