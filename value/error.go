package value

import (
	"math"
	"math/big"
)

var _ Value = Error{}

type Error struct {
	content string
}

func NewError(v string) Error {
	return Error{
		content: v,
	}
}

// Type implements Value.
func (Error) Type() string {
	return __ERROR_VALUE
}

// IsNull implements Value.
func (Error) IsNull() bool {
	return false
}

// Equals implements Value.
func (v Error) Equals(other Value) bool {
	switch another := other.(type) {
	case Error:
		return v.content == another.content
	}
	return false
}

// Encoding implements Value.
func (Error) Encoding() string {
	return ""
}

// Array implements Value.
func (Error) Array() ([]Value, bool) {
	return nil, false
}

// BigNumber implements Value.
func (Error) BigNumber() (*big.Rat, bool) {
	return nil, false
}

// Error implements Value.
func (v Error) Error() (string, bool) {
	return v.content, true
}

// Float implements Value.
func (Error) Float() (float64, bool) {
	return math.NaN(), false
}

// Integer implements Value.
func (Error) Integer() (int64, bool) {
	return 0, false
}

// Map implements Value.
func (Error) Map() (map[string]Value, bool) {
	return nil, false
}

// String implements Value.
func (Error) String() (string, bool) {
	return "", false
}
