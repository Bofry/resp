package internal

import (
	"fmt"

	"github.com/FastHCA/resp/value"
)

func Resolve(reader ByteReader) (int, value.Value, error) {
	return resolve(reader)
}

func resolve(reader ByteReader) (int, value.Value, error) {
	var (
		offset int = 0
	)

	// read notation byte
	notation, err := reader.ReadByte()
	offset++
	if err != nil {
		return offset, nil, err
	}

	// get resolver
	resolver, ok := _DataReaderTable[notation]
	if !ok {
		return offset, nil, fmt.Errorf("resolve invalid notation byte '%c' at %d", notation, offset)
	}

	// resolve value
	n, value, err := resolver.Read(reader)
	offset += n

	return offset, value, err
}
