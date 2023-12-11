package resp

import (
	"io"

	"github.com/FastHCA/resp/internal"
	"github.com/FastHCA/resp/value"
)

var _ Data = new(DataWriteFunc)

type DataWriteFunc func(writer io.Writer) error

func (fn DataWriteFunc) Write(writer io.Writer) error {
	return fn(writer)
}

//-----------------------------------------

func Array(data ...Data) Data {
	return DataWriteFunc(func(writer io.Writer) error {
		return internal.ArrayPacker.Pack(writer, data...)
	})
}

func BulkString(data string) Data {
	return DataWriteFunc(func(writer io.Writer) error {
		var value = value.NewString(data)
		return internal.BulkStringPacker.Pack(writer, value)
	})
}

func Integer(v int64) Data {
	return DataWriteFunc(func(writer io.Writer) error {
		var value = value.NewInteger(v)
		return internal.IntegerPacker.Pack(writer, value)
	})
}

func NullArray() Data {
	return DataWriteFunc(func(writer io.Writer) error {
		return internal.NullArrayPacker.Pack(writer)
	})
}

func NullBulkString() Data {
	return DataWriteFunc(func(writer io.Writer) error {
		return internal.NullBulkStringPacker.Pack(writer)
	})
}

func SimpleString(data string) Data {
	return DataWriteFunc(func(writer io.Writer) error {
		var value = value.NewString(data)
		return internal.SimpleStringPacker.Pack(writer, value)
	})
}

func SimpleError(errmsg string) Data {
	return DataWriteFunc(func(writer io.Writer) error {
		var value = value.NewError(errmsg)
		return internal.SimpleErrorPacker.Pack(writer, value)
	})
}
