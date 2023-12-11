package internal

import (
	"bytes"
	"testing"

	"github.com/FastHCA/resp/value"
)

func Test_SimpleErrorReader_Read(t *testing.T) {
	buf := []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := SimpleErrorReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'SimpleError.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := SimpleErrorReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 67
	if offset != expectedOffset {
		t.Errorf("assert 'SimpleError.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedValue := value.NewError("WRONGTYPE Operation against a key holding the wrong kind of value")
	if v != expectedValue {
		t.Errorf("assert 'SimpleError.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}
