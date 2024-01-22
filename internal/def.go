package internal

import (
	"errors"
	"io"

	"github.com/FastHCA/resp/value"
)

const (
	__SIMPLE_STRING_TYPE   byte = '+' // simple
	__SIMPLE_ERROR_TYPE    byte = '-' // simple
	__INTEGER_TYPE         byte = ':' // simple
	__BULK_STRING_TYPE     byte = '$' // bulk
	__ARRAY_TYPE           byte = '*' // aggregate
	__NULL_TYPE            byte = '_' // simple
	__BOOLEAN_TYPE         byte = '#' // simple
	__DOUBLE_TYPE          byte = ',' // simple
	__BIGNUMBER_TYPE       byte = '(' // simple
	__BULK_ERROR_TYPE      byte = '!' // bulk
	__VERBATIM_STRING_TYPE byte = '=' // bulk
	__MAP_TYPE             byte = '%' // aggregate
	__SET_TYPE             byte = '~' // aggregate
	__PUSH_TYPE            byte = '>' // aggregate

	_CR         byte = '\r'
	_LF         byte = '\n'
	_TERMINATOR      = "\r\n"
)

const (
	SimpleStringReader = _SimpleStringReader(__SIMPLE_STRING_TYPE)
	SimpleErrorReader  = _SimpleErrorReader(__SIMPLE_ERROR_TYPE)
	IntegerReader      = _IntegerReader(__INTEGER_TYPE)
	BulkStringReader   = _BulkStringReader(__BULK_STRING_TYPE)
	ArrayReader        = _ArrayReader(__ARRAY_TYPE)

	// Null
	// Boolean
	// Double
	// BigNumber = _IntegerResolver(BIGNUMBER_TYPE)
	// VerbatimString
	// Map

	// BulkErrorReader = _BulkStringReader(__BULK_ERROR_TYPE)
	// SetReader       = _ArrayReader(__SET_TYPE)
	// PushReader      = _ArrayReader(__PUSH_TYPE)

	SimpleStringPacker   = _SimpleStringPacker(__SIMPLE_STRING_TYPE)
	SimpleErrorPacker    = _SimpleErrorPacker(__SIMPLE_ERROR_TYPE)
	IntegerPacker        = _IntegerPacker(__INTEGER_TYPE)
	BulkStringPacker     = _BulkStringPacker(__BULK_STRING_TYPE)
	ArrayPacker          = _ArrayPacker(__ARRAY_TYPE)
	NullBulkStringPacker = _BulkNullPacker(__BULK_STRING_TYPE)
	NullArrayPacker      = _BulkNullPacker(__ARRAY_TYPE)
)

type (
	ByteReader interface {
		Len() int
		Size() int64

		io.Reader
		io.ByteReader
		io.Seeker
	}

	DataReader interface {
		NotationByte() byte
		Read(reader ByteReader) (int, value.Value, error)
	}

	ValuePacker interface {
		Pack(writer io.Writer, val value.Value) error
	}

	DataPacker interface {
		Pack(writer io.Writer, data ...Data) error
	}

	Data interface {
		Write(writer io.Writer) error
	}
)

var (
	_ErrSimplePackerPackNullValue = errors.New("failed call pack() with null value")

	_DataReaderTable = map[byte]DataReader{
		__SIMPLE_STRING_TYPE: SimpleStringReader,
		__SIMPLE_ERROR_TYPE:  SimpleErrorReader,
		__INTEGER_TYPE:       IntegerReader,
		__BULK_STRING_TYPE:   BulkStringReader,
		__ARRAY_TYPE:         ArrayReader,
	}
)
