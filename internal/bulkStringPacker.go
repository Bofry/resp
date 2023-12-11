package internal

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/FastHCA/resp/value"
)

var _ ValuePacker = _BulkStringPacker(0)

type _BulkStringPacker byte

// Pack implements ValuePacker.
func (p _BulkStringPacker) Pack(writer io.Writer, val value.Value) error {
	w := acquireBufioWriter(writer)
	defer w.Flush()

	// notation byte
	err := p.writeNotationByte(w)
	if err != nil {
		return err
	}

	// is null?
	if val == nil || val.IsNull() {
		return p.writeNull(w)
	}

	// content
	{
		content, ok := val.String()
		if !ok {
			return fmt.Errorf("BulkStringPacker cannot pack unsupported type '%v'", val.Type())
		}

		// size
		err = p.writeSize(w, len(content))
		if err != nil {
			return err
		}

		err = p.writeContent(w, content)
		if err != nil {
			return err
		}
	}
	return nil
}

func (_BulkStringPacker) writeContent(w *bufio.Writer, content string) error {
	_, err := w.Write([]byte(content))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(_TERMINATOR))
	if err != nil {
		return err
	}
	return nil
}

func (p _BulkStringPacker) writeNotationByte(w *bufio.Writer) error {
	err := w.WriteByte(byte(p))
	if err != nil {
		return err
	}
	return nil
}

func (_BulkStringPacker) writeNull(w *bufio.Writer) error {
	_, err := w.Write([]byte("-1"))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(_TERMINATOR))
	if err != nil {
		return err
	}
	return nil
}

func (_BulkStringPacker) writeSize(w *bufio.Writer, size int) error {
	content := strconv.FormatInt(int64(size), 10)

	_, err := w.Write([]byte(content))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(_TERMINATOR))
	if err != nil {
		return err
	}

	return nil
}
