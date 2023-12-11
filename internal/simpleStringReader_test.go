package internal

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/FastHCA/resp/value"
)

func Test_SimpleStringReader_Read(t *testing.T) {
	buf := []byte("+OK\r\n")

	reader := bytes.NewReader(buf)

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := SimpleStringReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'SimpleString.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := SimpleStringReader.Read(reader)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedOffset := 4
	if offset != expectedOffset {
		t.Errorf("assert 'SimpleString.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedValue := value.NewString("OK")
	if !reflect.DeepEqual(v, expectedValue) {
		t.Errorf("assert 'SimpleString.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
	}
}
