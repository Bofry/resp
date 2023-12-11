package internal

import (
	"bufio"
	"fmt"
	"io"

	"github.com/FastHCA/resp/value"
)

var _ ValuePacker = _SimpleStringPacker(0)

type _SimpleStringPacker byte

// Pack implements ValuePacker.
func (p _SimpleStringPacker) Pack(writer io.Writer, val value.Value) error {
	w := acquireBufioWriter(writer)
	defer w.Flush()

	// is null?
	if val == nil || val.IsNull() {
		return _ErrSimplePackerPackNullValue
	}

	// notation byte
	err := p.writeNotationByte(w)
	if err != nil {
		return err
	}

	// content
	{
		content, ok := val.String()
		if !ok {
			return fmt.Errorf("SimpleStringPacker cannot pack unsupported type '%v'", val.Type())
		}

		err = p.validateContent(content)
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

func (p _SimpleStringPacker) writeNotationByte(w *bufio.Writer) error {
	err := w.WriteByte(byte(p))
	if err != nil {
		return err
	}
	return nil
}

func (_SimpleStringPacker) writeContent(w *bufio.Writer, content string) error {
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

func (p _SimpleStringPacker) validateContent(s string) error {
	return validateSimpleValue([]byte(s))
}
