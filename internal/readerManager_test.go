package internal

import (
	"bytes"
	"testing"
)

func TestReaderManager(t *testing.T) {
	r1 := bytes.NewReader([]byte("foo"))
	r2 := bytes.NewReader([]byte("bar"))

	c := NewReaderManager(r1, r2)

	{
		current := c.current()
		expectedCurrent := r1
		if current != expectedCurrent {
			t.Errorf("assert 'ReaderContainer.current()':: expected '%+v', got '%+v'", expectedCurrent, current)
		}
	}

	c.next()
	{
		current := c.current()
		expectedCurrent := r2
		if current != expectedCurrent {
			t.Errorf("assert 'ReaderContainer.current()':: expected '%+v', got '%+v'", expectedCurrent, current)
		}
	}

	c.next()
	{
		current := c.current()
		var expectedCurrent ByteReader = nil
		if current != expectedCurrent {
			t.Errorf("assert 'ReaderContainer.current()':: expected '%+v', got '%+v'", expectedCurrent, current)
		}
	}
}
