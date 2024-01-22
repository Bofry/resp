package internal

import (
	"io"
)

type ReaderManager struct {
	readers []ByteReader
	index   int
}

func NewReaderManager(readers ...ByteReader) *ReaderManager {
	return &ReaderManager{
		readers: readers,
		index:   0,
	}
}

func (m *ReaderManager) current() ByteReader {
	if len(m.readers) > m.index {
		return m.readers[m.index]
	}
	return nil
}

func (m *ReaderManager) next() {
	if len(m.readers) > m.index {
		m.index++
	}
}

func (m *ReaderManager) skip(n int64) error {
	var (
		skipped int64 = 0
		offset        = int64(m.index) + n
	)

	for _, r := range m.readers {
		size := int64(r.Size())

		if (offset - skipped) < size {
			pos := offset - skipped

			rd := m.readers[m.index]
			_, err := rd.Seek(pos, io.SeekStart)
			if err != nil {
				return err
			}
			break
		}
		skipped += size
		m.index++
	}

	for i := m.index + 1; i < len(m.readers); i++ {
		rd := m.readers[i]
		offset, err := rd.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
		if rd.Size() > 0 && offset == 0 {
			break
		}

		_, err = rd.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *ReaderManager) reset() {
	m.index = 0
}

func (m *ReaderManager) forget() {
	if m.index > 0 {
		m.readers = m.readers[m.index:]
		m.index = 0
	}
}

func (m *ReaderManager) append(readers ...ByteReader) {
	m.readers = append(m.readers, readers...)
}
