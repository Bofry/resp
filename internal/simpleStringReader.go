package internal

import (
	"bytes"
	"fmt"

	"github.com/FastHCA/resp/value"
)

var _ DataReader = _SimpleStringReader(0)

type _SimpleStringReader byte

// NotationByte implements DataReader.
func (p _SimpleStringReader) NotationByte() byte {
	return byte(p)
}

// Read implements DataReader.
func (_SimpleStringReader) Read(r ByteReader) (int, value.Value, error) {

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
		case _LF:
			return buf.Len() + 1, nil, fmt.Errorf("read invalid character %q at %d", b, offset)
		default:
			buf.WriteByte(b)
		}
	}
	return buf.Len() + len(_TERMINATOR), value.NewString(buf.String()), nil
}
