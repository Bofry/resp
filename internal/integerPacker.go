package internal

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/FastHCA/resp/value"
)

var _ ValuePacker = _IntegerPacker(0)

type _IntegerPacker byte

// Pack implements ValuePacker.
func (p _IntegerPacker) Pack(writer io.Writer, val value.Value) error {
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
		number, ok := val.Integer()
		if !ok {
			return fmt.Errorf("SimpleErrorPacker cannot pack unsupported type '%v'", val.Type())
		}

		err = p.writeContent(w, number)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p _IntegerPacker) writeNotationByte(w *bufio.Writer) error {
	err := w.WriteByte(byte(p))
	if err != nil {
		return err
	}
	return nil
}

func (_IntegerPacker) writeContent(w *bufio.Writer, number int64) error {
	content := strconv.FormatInt(number, 10)

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
