package internal

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/FastHCA/resp/value"
)

func Test_BulkStringReader_Read_Well(t *testing.T) {
	buf := []byte("$5\r\nhello\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := BulkStringReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'BulkString.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := BulkStringReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 10
	if offset != expectedOffset {
		t.Errorf("assert 'BulkString.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedValue := value.NewString("hello")
	if !reflect.DeepEqual(v, expectedValue) {
		t.Errorf("assert 'BulkString.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}

func Test_BulkStringReader_Read_NullValue(t *testing.T) {
	buf := []byte("$-1\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := BulkStringReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'BulkString.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := BulkStringReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 4
	if offset != expectedOffset {
		t.Errorf("assert 'BulkString.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	var expectedValue *value.String = nil
	if !reflect.DeepEqual(v, expectedValue) {
		t.Errorf("assert 'BulkString.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}

func Test_BulkStringReader_Read_EmptyString(t *testing.T) {
	buf := []byte("$0\r\n\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := BulkStringReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'BulkString.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := BulkStringReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 5
	if offset != expectedOffset {
		t.Errorf("assert 'BulkString.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	var expectedValue value.Value = value.NewString("")
	if !reflect.DeepEqual(v, expectedValue) {
		t.Errorf("assert 'BulkString.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}

func Test_BulkStringReader_Read_InvalidLength(t *testing.T) {
	buf := []byte("$-2\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := BulkStringReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'BulkString.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	_, _, err = BulkStringReader.Read(reader)
	if err == nil {
		t.Errorf("should get error")
	}
}

func Test_BulkStringReader_Read_InvalidContent(t *testing.T) {
	buf := []byte("$0\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := BulkStringReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'BulkString.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	_, _, err = BulkStringReader.Read(reader)
	if err == nil {
		t.Errorf("should get error")
	}
}
