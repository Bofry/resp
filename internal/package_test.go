package internal

import (
	"bytes"
	"io"
	"testing"

	"github.com/FastHCA/resp/value"
)

func Test_IntegerReader(t *testing.T) {
	t.SkipNow()

	buf := []byte(":1324")
	buf = append(buf, byte(13))
	buf = append(buf, byte(0))
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
	t.Logf("offset:: %+v", offset)
	if err != nil {
		t.Errorf("should no error, but got %+v", err)
	}
	_ = v

	{
		reader.Seek(int64(-offset), io.SeekCurrent)

		offset, v, err := IntegerReader.Read(reader)
		t.Logf("offset:: %+v", offset)
		if err != nil {
			t.Errorf("should no error, but got %+v", err)
		}

		expectedOffset := 6
		if offset != expectedOffset {
			t.Errorf("assert 'Integer.Read() offset':: expected '%+v', got '%+v'", expectedOffset, offset)
		}
		expectedValue := value.NewInteger(1324)
		if v != expectedValue {
			t.Errorf("assert 'Integer.Read() value':: expected '%+v', got '%+v'", expectedValue, v)
		}
	}
}

// func TestXxx(t *testing.T) {
// 	r1 := bytes.NewReader([]byte("foo"))
// 	r2 := bytes.NewReader([]byte("bar"))
// 	r3 := bytes.NewReader([]byte("baz"))

// 	r := io.MultiReader(r1, r2, r3)

// }
