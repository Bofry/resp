package internal

import (
	"bufio"
	"io"
)

var _ DataPacker = _BulkNullPacker(0)

type _BulkNullPacker byte

// Pack implements Packer.
func (p _BulkNullPacker) Pack(writer io.Writer, data ...Data) error {
	w := acquireBufioWriter(writer)
	defer w.Flush()

	// notation byte
	err := p.writeNotationByte(w)
	if err != nil {
		return err
	}

	// content
	err = p.writeNull(w)
	if err != nil {
		return err
	}
	return nil
}

func (p _BulkNullPacker) writeNotationByte(w *bufio.Writer) error {
	err := w.WriteByte(byte(p))
	if err != nil {
		return err
	}
	return nil
}

func (p _BulkNullPacker) writeNull(w *bufio.Writer) error {
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
