package internal

import (
	"bytes"
	"testing"
)

func TestReaderManager_current(t *testing.T) {
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

func TestReaderManager_forget(t *testing.T) {
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

	c.forget()
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

func TestReaderManager_append(t *testing.T) {
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

	r3 := bytes.NewReader([]byte("baz"))
	c.append(r3)
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
		expectedCurrent := r3
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
