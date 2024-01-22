package internal

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestMultiByteReader_Len(t *testing.T) {
	{
		r := NewMultByteReader()

		var expectedLen int = 0
		if r.Len() != expectedLen {
			t.Errorf("assert 'MultByteReader.Len() offset':: expected '%+v', got '%+v'", expectedLen, r.Len())
		}
	}

	{
		r1 := bytes.NewReader([]byte("foo"))

		r := NewMultByteReader(r1)

		var expectedLen int = 3
		if r.Len() != expectedLen {
			t.Errorf("assert 'MultByteReader.Len() offset':: expected '%+v', got '%+v'", expectedLen, r.Len())
		}
	}

	{
		r1 := bytes.NewReader([]byte("foo"))
		r2 := bytes.NewReader([]byte("bar"))

		r := NewMultByteReader(r1, r2)

		var expectedLen int = 6
		if r.Len() != expectedLen {
			t.Errorf("assert 'MultByteReader.Len() offset':: expected '%+v', got '%+v'", expectedLen, r.Len())
		}
	}
}

func TestMultiByteReader_Size(t *testing.T) {
	{
		r := NewMultByteReader()

		var expectedLen int = 0
		if r.Len() != expectedLen {
			t.Errorf("assert 'MultByteReader.Len() offset':: expected '%+v', got '%+v'", expectedLen, r.Len())
		}
	}

	{
		r1 := bytes.NewReader([]byte("foo"))

		r := NewMultByteReader(r1)

		var expectedSize int64 = 3
		if r.Size() != expectedSize {
			t.Errorf("assert 'MultByteReader.Size() offset':: expected '%+v', got '%+v'", expectedSize, r.Size())
		}
	}

	{
		r1 := bytes.NewReader([]byte("foo"))
		r2 := bytes.NewReader([]byte("bar"))

		r := NewMultByteReader(r1, r2)

		var expectedSize int64 = 6
		if r.Size() != expectedSize {
			t.Errorf("assert 'MultByteReader.Size() offset':: expected '%+v', got '%+v'", expectedSize, r.Size())
		}
	}
}

func TestMultiByteReader_ReadByte(t *testing.T) {
	{
		r := NewMultByteReader()

		b, err := r.ReadByte()
		expectetErr := io.EOF
		if err != expectetErr {
			t.Errorf("assert 'MultByteReader.ReadByte() error':: expected '%+v', got '%+v'", expectetErr, err)
		}
		expectetByte := byte(0)
		if b != expectetByte {
			t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
		}
	}

	{
		r1 := bytes.NewReader([]byte("foo"))

		r := NewMultByteReader(r1)

		{
			b, err := r.ReadByte()
			if err != nil {
				t.Errorf("should no error, but got %+v", err)
			}
			expectetByte := byte('f')
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
		{
			b, err := r.ReadByte()
			if err != nil {
				t.Errorf("should no error, but got %+v", err)
			}
			expectetByte := byte('o')
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
		{
			b, err := r.ReadByte()
			if err != nil {
				t.Errorf("should no error, but got %+v", err)
			}
			expectetByte := byte('o')
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
		{
			b, err := r.ReadByte()
			expectetErr := io.EOF
			if err != expectetErr {
				t.Errorf("assert 'MultByteReader.ReadByte() error':: expected '%+v', got '%+v'", expectetErr, err)
			}
			expectetByte := byte(0)
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
	}

	{
		r1 := bytes.NewReader([]byte("foo"))
		r2 := bytes.NewReader([]byte("bar"))

		r := NewMultByteReader(r1, r2)

		{
			b, err := r.ReadByte()
			if err != nil {
				t.Errorf("should no error, but got %+v", err)
			}
			expectetByte := byte('f')
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
		{
			b, err := r.ReadByte()
			if err != nil {
				t.Errorf("should no error, but got %+v", err)
			}
			expectetByte := byte('o')
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
		{
			b, err := r.ReadByte()
			if err != nil {
				t.Errorf("should no error, but got %+v", err)
			}
			expectetByte := byte('o')
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
		{
			b, err := r.ReadByte()
			if err != nil {
				t.Errorf("should no error, but got %+v", err)
			}
			expectetByte := byte('b')
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
		{
			b, err := r.ReadByte()
			if err != nil {
				t.Errorf("should no error, but got %+v", err)
			}
			expectetByte := byte('a')
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
		{
			b, err := r.ReadByte()
			if err != nil {
				t.Errorf("should no error, but got %+v", err)
			}
			expectetByte := byte('r')
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
		{
			b, err := r.ReadByte()
			expectetErr := io.EOF
			if err != expectetErr {
				t.Errorf("assert 'MultByteReader.ReadByte() error':: expected '%+v', got '%+v'", expectetErr, err)
			}
			expectetByte := byte(0)
			if b != expectetByte {
				t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetByte, b)
			}
		}
	}
}

func TestMultiByteReader_Read(t *testing.T) {
	r1 := bytes.NewReader([]byte("foo"))
	r2 := bytes.NewReader([]byte(""))
	r3 := bytes.NewReader([]byte("barbaz"))
	var buf []byte = make([]byte, 9)

	r := NewMultByteReader(r1, r2, r3)

	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}
	expectetN := 9
	if n != expectetN {
		t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetN, n)
	}

	expectetBuf := []byte("foobarbaz")
	if !reflect.DeepEqual(buf, expectetBuf) {
		t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetBuf, buf)
	}
}

func TestMultiByteReader_Read_WithNested(t *testing.T) {
	r1 := bytes.NewReader([]byte("foo"))
	r2 := NewMultByteReader(
		bytes.NewReader([]byte("")),
		bytes.NewReader([]byte("barbaz")),
	)

	var buf []byte = make([]byte, 9)

	r := NewMultByteReader(r1, r2)

	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}
	expectetN := 9
	if n != expectetN {
		t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetN, n)
	}

	expectetBuf := []byte("foobarbaz")
	if !reflect.DeepEqual(buf, expectetBuf) {
		t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetBuf, buf)
	}
}

func TestMultiByteReader_Seek(t *testing.T) {
	r1 := bytes.NewReader([]byte("foo"))
	r2 := bytes.NewReader([]byte(""))
	r3 := bytes.NewReader([]byte("barbaz"))

	{
		r := NewMultByteReader(r1, r2, r3)

		offset, err := r.Seek(2, io.SeekStart)
		if err != nil {
			t.Errorf("should no error, but got %+v", err)
		}
		expectetOffset := int64(2)
		if offset != expectetOffset {
			t.Errorf("assert 'MultByteReader.Seek() offset':: expected '%+v', got '%+v'", expectetOffset, offset)
		}
	}

	{
		r := NewMultByteReader(r1, r2, r3)

		offset, err := r.Seek(5, io.SeekStart)
		if err != nil {
			t.Errorf("should no error, but got %+v", err)
		}
		expectetOffset := int64(5)
		if offset != expectetOffset {
			t.Errorf("assert 'MultByteReader.Seek() offset':: expected '%+v', got '%+v'", expectetOffset, offset)
		}
	}
}

func TestMultiByteReader_Seek_And_Read(t *testing.T) {
	r1 := bytes.NewReader([]byte("foo"))
	r2 := bytes.NewReader([]byte(""))
	r3 := bytes.NewReader([]byte("barbaz"))

	{
		var buf []byte = make([]byte, 7)

		r := NewMultByteReader(r1, r2, r3)

		_, err := r.Seek(2, io.SeekStart)
		if err != nil {
			t.Errorf("should no error, but got %+v", err)
		}

		n, err := r.Read(buf)
		if err != nil {
			t.Errorf("should no error, but got %+v", err)
		}
		expectetN := 7
		if n != expectetN {
			t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetN, n)
		}

		expectetBuf := []byte("obarbaz")
		if !reflect.DeepEqual(buf, expectetBuf) {
			t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetBuf, buf)
		}
	}

	{
		var buf []byte = make([]byte, 4)

		r := NewMultByteReader(r1, r2, r3)

		_, err := r.Seek(5, io.SeekStart)
		if err != nil {
			t.Errorf("should no error, but got %+v", err)
		}

		n, err := r.Read(buf)
		if err != nil {
			t.Errorf("should no error, but got %+v", err)
		}
		expectetN := 4
		if n != expectetN {
			t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetN, n)
		}

		expectetBuf := []byte("rbaz")
		if !reflect.DeepEqual(buf, expectetBuf) {
			t.Errorf("assert 'MultByteReader.ReadByte() byte':: expected '%+v', got '%+v'", expectetBuf, buf)
		}
	}
}
