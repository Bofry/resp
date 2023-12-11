package internal

import (
	"bytes"
	"testing"

	"github.com/FastHCA/resp/value"
)

func Test_IntegerReader_Read(t *testing.T) {
	buf := []byte(":1024\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := IntegerReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'Integer.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := IntegerReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 6
	if offset != expectedOffset {
		t.Errorf("assert 'Integer.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedValue := value.NewInteger(1024)
	if v != expectedValue {
		t.Errorf("assert 'Integer.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}

func Test_IntegerReader_Read_With_PositiveValue(t *testing.T) {
	buf := []byte(":+1024\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := IntegerReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'Integer.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := IntegerReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 7
	if offset != expectedOffset {
		t.Errorf("assert 'Integer.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedValue := value.NewInteger(1024)
	if v != expectedValue {
		t.Errorf("assert 'Integer.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}

func Test_IntegerReader_Read_With_NegativeValue(t *testing.T) {
	buf := []byte(":-1023\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := IntegerReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'Integer.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := IntegerReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 7
	if offset != expectedOffset {
		t.Errorf("assert 'Integer.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedValue := value.NewInteger(-1023)
	if v != expectedValue {
		t.Errorf("assert 'Integer.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}
