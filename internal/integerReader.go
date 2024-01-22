package internal

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/FastHCA/resp/value"
)

var _ DataReader = _IntegerReader(0)

type _IntegerReader byte

// NotationByte implements DataReader.
func (v _IntegerReader) NotationByte() byte {
	return byte(v)
}

// Read implements DataReader.
func (_IntegerReader) Read(r ByteReader) (int, value.Value, error) {
	var (
		buf      bytes.Buffer
		offset   int
		kontinue = true
	)

	for kontinue {
		b, err := r.ReadByte()

		if err != nil {
			return buf.Len(), nil, err
		}
		offset++

		switch b {
		case _CR:
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
			if !(b == '+' || b == '-' || (b >= '0' && b <= '9')) {
				return offset, nil, fmt.Errorf("read invalid character %q at %d", b, offset)
			}
			buf.WriteByte(b)
		}
	}

	var res value.Value
	{
		content := buf.String()
		i, err := strconv.ParseInt(content, 10, 64)
		if err != nil {
			return offset, nil, err
		}
		// export
		res = value.NewInteger(i)
	}
	return offset, res, nil
}
