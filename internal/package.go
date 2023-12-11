package internal

import (
	"bufio"
	"fmt"
	"io"

	"github.com/FastHCA/resp/value"
)

func Read(reader io.Reader) (int, value.Value, error) {
	r := acquireBufioReader(reader)
	return read(r)
}

func read(reader *bufio.Reader) (int, value.Value, error) {
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
