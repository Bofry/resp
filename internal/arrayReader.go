package internal

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/FastHCA/resp/value"
)

var _ DataReader = _ArrayReader(0)

type _ArrayReader byte

// NotationByte implements DataReader.
func (p _ArrayReader) NotationByte() byte {
	return byte(p)
}

// Read implements DataReader.
func (_ArrayReader) Read(r ByteReader) (int, value.Value, error) {
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
			return offset, value.NullArray(), err
		}
		// export
		size = n
	}

	if size < 0 {
		// Null arrays
		if size == -1 {
			return offset, nil, nil
		}
		return offset, nil, fmt.Errorf("read invalid length '%d' at %d", size, offset)
	}

	// read content
	var content []value.Value
	if size > 0 {
		content = make([]value.Value, 0, size)

		for i := 0; i < int(size); i++ {
			n, elem, err := resolve(r)
			offset += n

			if err != nil {
				return offset, nil, err
			}

			content = append(content, elem)
		}
	}

	return offset, value.NewArray(content...), nil
}
