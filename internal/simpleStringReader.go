package internal

import (
	"bytes"
	"fmt"
	"io"

	"github.com/FastHCA/resp/value"
)

var _ DataReader = _SimpleStringReader(0)

type _SimpleStringReader byte

// NotationByte implements DataReader.
func (p _SimpleStringReader) NotationByte() byte {
	return byte(p)
}

// Read implements DataReader.
func (_SimpleStringReader) Read(reader io.Reader) (int, value.Value, error) {
	r := acquireBufioReader(reader)

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

			b, err = r.ReadByte()
			offset++

			if err != nil {
				return offset, nil, err
			}
			if b != _LF {
				return offset, nil, fmt.Errorf("read invalid terminator '%c' at %d", b, offset)
			}
			kontinue = false
			break
		case _LF:
			return buf.Len(), nil, fmt.Errorf("read invalid character '%c' at %d", b, offset)
		default:
			buf.WriteByte(b)
		}
	}
	return buf.Len() + len(_TERMINATOR), value.NewString(buf.String()), nil
}
