package internal

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/FastHCA/resp/value"
)

func Test_ArrayResolver_Resolve_Well(t *testing.T) {
	buf := []byte("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := ArrayReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'Array.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := ArrayReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 25
	if offset != expectedOffset {
		t.Errorf("assert 'Array.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedValue := value.NewArray(
		value.NewString("hello"),
		value.NewString("world"),
	)
	if !reflect.DeepEqual(v, expectedValue) {
		t.Errorf("assert 'Array.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}

func Test_ArrayResolver_Resolve_NullValue(t *testing.T) {
	buf := []byte("*-1\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := ArrayReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'Array.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := ArrayReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 4
	if offset != expectedOffset {
		t.Errorf("assert 'Array.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	var expectedValue value.Value = nil
	if v != expectedValue {
		t.Errorf("assert 'Array.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}

func Test_ArrayResolver_Resolve_EmptyArray(t *testing.T) {
	buf := []byte("*0\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := ArrayReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'Array.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := ArrayReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 3
	if offset != expectedOffset {
		t.Errorf("assert 'Array.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedValue := value.NewArray()
	if !reflect.DeepEqual(v, expectedValue) {
		t.Errorf("assert 'Array.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}

func Test_ArrayResolver_Resolve_InvalidLength(t *testing.T) {
	buf := []byte("*-2\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := ArrayReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'Array.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	_, _, err = BulkStringReader.Read(reader)
	if err == nil {
		t.Errorf("should get error")
	}
}

func Test_ArrayResolver_Resolve_Nested(t *testing.T) {
	buf := []byte("*2\r\n*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n*3\r\n$3\r\nfoo\r\n$3\r\nbar\r\n:1024\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := ArrayReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'Array.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := ArrayReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 58
	if offset != expectedOffset {
		t.Errorf("assert 'Array.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedValue := value.NewArray(
		value.NewArray(
			value.NewString("hello"),
			value.NewString("world"),
		),
		value.NewArray(
			value.NewString("foo"),
			value.NewString("bar"),
			value.NewInteger(1024),
		),
	)
	if !reflect.DeepEqual(v, expectedValue) {
		t.Errorf("assert 'Array.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}
