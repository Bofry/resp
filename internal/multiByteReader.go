package internal

import (
	"errors"
	"io"
)

var (
	_ ByteReader          = new(MultiByteReader)
	_ CompositeByteReader = new(MultiByteReader)
)

type MultiByteReader struct {
	pos     int64
	manager *ReaderManager
}

func NewMultByteReader(readers ...ByteReader) *MultiByteReader {

	manager := NewReaderManager(readers...)

	return &MultiByteReader{
		manager: manager,
		pos:     0,
	}
}

// Len implements ByteReader.
func (r *MultiByteReader) Len() int {
	var size int = 0

	for _, rd := range r.readers() {
		size += rd.Len()
	}
	return size
}

// Read implements ByteReader.
func (r *MultiByteReader) Read(p []byte) (n int, err error) {
	var reader = r.manager.current()
	if reader == nil {
		return 0, io.EOF
	}

	var count int = 0
	for reader != nil {
		n, err := reader.Read(p[count:])
		if err == io.EOF {
			r.manager.next()
			reader = r.manager.current()
			continue
		}
		if err != nil {
			return 0, err
		}
		count += n

		if len(p) <= count {
			break
		}
		r.manager.next()
		reader = r.manager.current()
	}
	r.pos += int64(count)

	return count, nil
}

// ReadByte implements ByteReader.
func (r *MultiByteReader) ReadByte() (byte, error) {
	var reader = r.manager.current()
	for reader != nil {
		b, err := reader.ReadByte()
		if err == io.EOF {
			r.manager.next()
			reader = r.manager.current()
			continue
		}
		if err != nil {
			return b, err
		}

		r.pos++
		return b, nil
	}

	return 0, io.EOF
}

// Seek implements ByteReader.
func (r *MultiByteReader) Seek(offset int64, whence int) (int64, error) {
	var pos int64

	switch whence {
	case io.SeekStart:
		pos = offset
	case io.SeekCurrent:
		pos = r.pos + offset
	case io.SeekEnd:
		pos = r.Size() + offset
	default:
		return 0, errors.New("resp.MultByteReader.Seek: invalid whence")
	}
	if pos < 0 {
		return 0, errors.New("resp.MultByteReader.Seek: negative position")
	}
	r.pos = pos

	r.manager.reset()
	err := r.manager.skip(pos)
	if err != nil {
		return 0, err
	}
	return pos, nil
}

// Size implements ByteReader.
func (r *MultiByteReader) Size() int64 {
	var size int64 = 0

	for _, rd := range r.readers() {
		size += rd.Size()
	}
	return size
}

// Append implements CompositeByteReader.
func (r *MultiByteReader) Append(readers ...ByteReader) error {
	r.manager.append(readers...)
	return nil
}

// Forget implements CompositeByteReader.
func (r *MultiByteReader) Forget() error {
	r.manager.forget()
	return nil
}

func (r *MultiByteReader) readers() []ByteReader {
	return r.manager.readers
}
