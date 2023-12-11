package value

import "math/big"

const (
	__STRING_VALUE          = "string"
	__ERROR_VALUE           = "error"
	__INTEGER_VALUE         = "integer"
	__ARRAY_VALUE           = "array"
	__NULL_VALUE            = "null"
	__BOOLEAN_VALUE         = "boolean"
	__DOUBLE_VALUE          = "double"
	__BIGNUMBER_VALUE       = "bignumber"
	__VERBATIM_STRING_VALUE = "verbatim_string"
	__MAP_VALUE             = "map"
	__SET_VALUE             = "set"
	__PUSH_VALUE            = "push"
)

type (
	Value interface {
		Type() string
		IsNull() bool
		Equals(other Value) bool
		String() (string, bool)
		Error() (string, bool)
		Integer() (int64, bool)
		Float() (float64, bool)
		BigNumber() (*big.Rat, bool)
		Map() (map[string]Value, bool)
		Array() ([]Value, bool)
	}

	ScalarValue interface {
		Value

		Encoding() string
	}
)

var (
	_NullBulkString *String = nil
	_NullArray      *Array  = nil
	_EmptyArray     *Array  = new(Array)
)
