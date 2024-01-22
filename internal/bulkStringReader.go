package internal

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/FastHCA/resp/value"
)

var _ DataReader = _BulkStringReader(0)

type _BulkStringReader byte

// NotationByte implements DataReader.
func (v _BulkStringReader) NotationByte() byte {
	return byte(v)
}

// Read implements DataReader.
func (_BulkStringReader) Read(r ByteReader) (int, value.Value, error) {
	var (
		size   int64
		offset int
	)

	// read length
	{
		var (
			buf      bytes.Buffer
			kontinue = true
		)
		for kontinue {
			b, err := r.ReadByte()
			if err != nil {
				return buf.Len(), nil, err
			}

			switch b {
			case _CR:
				offset = buf.Len()
				offset++

				b, err = r.ReadByte()
				if err != nil {
					return offset, nil, err
				}
				offset++

				if b != _LF {
					return offset, nil, fmt.Errorf("read invalid terminator %q at %d", b, offset)
				}
				kontinue = false
				break
			default:
				if !(b == '-' || (b >= '0' && b <= '9')) {
					return buf.Len() + 1, nil, fmt.Errorf("read invalid character %q at %d", b, offset)
				}
				buf.WriteByte(b)
			}
		}
		offset = buf.Len() + len(_TERMINATOR)

		content := buf.String()
		n, err := strconv.ParseInt(content, 10, 64)
		if err != nil {
			return offset, nil, err
		}
		// export
		size = n
	}

	if size < 0 {
		// Null bulk strings
		if size == -1 {
			return offset, value.NullBulkString(), nil
		}
		return offset, nil, fmt.Errorf("read invalid length '%d' at %d", size, offset)
	}

	// read content
	var content []byte
	if size > 0 {
		content = make([]byte, size)

		n, err := r.Read(content)
		offset += n
		if err != nil {
			return offset, nil, err
		}
	}

	// check terminator
	{
		var terminator []byte = make([]byte, len(_TERMINATOR))
		n, err := r.Read(terminator)
		offset += n
		if err != nil {
			return offset, nil, errors.Join(fmt.Errorf("read invalid terminator '%v' at %d", terminator, offset-n+1), err)
		}
	}

	return offset, value.NewString(string(content)), nil
}
