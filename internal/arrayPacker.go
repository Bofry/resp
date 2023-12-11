package internal

import (
	"bufio"
	"io"
	"strconv"
)

var _ DataPacker = _ArrayPacker(0)

type _ArrayPacker byte

// Pack implements DataPacker.
func (p _ArrayPacker) Pack(writer io.Writer, data ...Data) error {
	w := acquireBufioWriter(writer)
	defer w.Flush()

	// notation byte
	err := p.writeNotationByte(w)
	if err != nil {
		return err
	}

	// size
	size := len(data)
	err = p.writeSize(w, size)
	if err != nil {
		return err
	}

	// content
	for _, packer := range data {
		err = packer.Write(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p _ArrayPacker) writeNotationByte(w *bufio.Writer) error {
	err := w.WriteByte(byte(p))
	if err != nil {
		return err
	}
	return nil
}

func (_ArrayPacker) writeSize(w *bufio.Writer, size int) error {
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
