package internal

import (
	"errors"
	"io"
)

var (
	_ ByteReader = new(MultByteReader)
)

type MultByteReader struct {
	readers   []ByteReader
	pos       int64
	container *ReaderContainer
}

func NewMultByteReader(readers ...ByteReader) *MultByteReader {

	container := NewReaderManager(readers...)

	return &MultByteReader{
		readers:   readers,
		container: container,
		pos:       0,
	}
}

// Len implements ByteReader.
func (r *MultByteReader) Len() int {
	var size int = 0

	for _, rd := range r.readers {
		size += rd.Len()
	}
	return size
}

// Read implements ByteReader.
func (r *MultByteReader) Read(p []byte) (n int, err error) {
	var reader = r.container.current()
	if reader == nil {
		return 0, io.EOF
	}

	var count int = 0
	for reader != nil {
		n, err := reader.Read(p[count:])
		if err == io.EOF {
			r.container.next()
			reader = r.container.current()
			continue
		}
		if err != nil {
			return 0, err
		}
		count += n

		if len(p) <= count {
			break
		}
		r.container.next()
		reader = r.container.current()
	}
	return count, nil
}

// ReadByte implements ByteReader.
func (r *MultByteReader) ReadByte() (byte, error) {
	var reader = r.container.current()
	for reader != nil {
		b, err := reader.ReadByte()
		if err == io.EOF {
			r.container.next()
			reader = r.container.current()
			continue
		}
		if err != nil {
			return b, err
		}

		return b, nil
	}

	return 0, io.EOF
}

// Seek implements ByteReader.
func (r *MultByteReader) Seek(offset int64, whence int) (int64, error) {
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

	r.container.reset()
	err := r.container.skip(pos)
	if err != nil {
		return 0, err
	}
	return pos, nil
}

// Size implements ByteReader.
func (r *MultByteReader) Size() int64 {
	var size int64 = 0

	for _, rd := range r.readers {
		size += rd.Size()
	}
	return size
}
