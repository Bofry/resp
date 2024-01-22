package internal

import (
	"io"
)

type ReaderContainer struct {
	readers []ByteReader
	index   int
}

func NewReaderManager(readers ...ByteReader) *ReaderContainer {
	return &ReaderContainer{
		readers: readers,
		index:   0,
	}
}

func (c *ReaderContainer) current() ByteReader {
	if len(c.readers) > c.index {
		return c.readers[c.index]
	}
	return nil
}

func (c *ReaderContainer) next() {
	if len(c.readers) > c.index {
		c.index++
	}
}

func (c *ReaderContainer) skip(n int64) error {
	var (
		skipped int64 = 0
		offset        = int64(c.index) + n
	)

	for _, r := range c.readers {
		size := int64(r.Size())

		if (offset - skipped) < size {
			pos := offset - skipped

			rd := c.readers[c.index]
			_, err := rd.Seek(pos, io.SeekStart)
			if err != nil {
				return err
			}
			break
		}
		skipped += size
		c.index++
	}

	for i := c.index + 1; i < len(c.readers); i++ {
		rd := c.readers[i]
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

func (c *ReaderContainer) reset() {
	c.index = 0
}
