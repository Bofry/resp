package internal

import (
	"bytes"
	"io"
	"testing"

	"github.com/FastHCA/resp/value"
)

func Test_IntegerReader(t *testing.T) {
	var reader CompositeByteReader

	{
		buf := []byte(":1324\r")
		reader = NewMultiByteReader(
			bytes.NewReader(buf),
		)
	}

	firstbyte, err := reader.ReadByte()
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}

	expectedFirstbyte := IntegerReader.NotationByte()
	if firstbyte != expectedFirstbyte {
		t.Errorf("assert 'Integer.FirstByte()':: expected '%+v', got '%+v'", expectedFirstbyte, firstbyte)
	}

	offset, v, err := IntegerReader.Read(reader)
	expectedOffset := 5
	if offset != expectedOffset {
		t.Errorf("assert 'IntegerReader.Read()' offset:: expected '%+v', got '%+v'", expectedOffset, offset)
	}
	expectedErr := io.EOF
	if err != expectedErr {
		t.Errorf("should %+v, but got %+v", expectedErr, err)
	}
	_ = v

	{
		reader.Append(bytes.NewReader([]byte{'\n'}))
		reader.Seek(int64(-offset), io.SeekCurrent)

		offset, v, err := IntegerReader.Read(reader)
		if err != nil {
			t.Errorf("should no error, but got %+v", err)
		}

		expectedOffset := 6
		if offset != expectedOffset {
			t.Errorf("assert 'IntegerReader.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
		}
		expectedValue := value.NewInteger(1324)
		if v != expectedValue {
			t.Errorf("assert 'IntegerReader.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
		}
	}
}
