package resp

import (
	"bytes"
	"io"
	_ "unsafe"

	"github.com/FastHCA/resp/value"
)

//go:linkname Resolve github.com/FastHCA/resp/internal.Resolve
func Resolve(reader io.Reader) (int, value.Value, error)

func Marshal(data Data) ([]byte, error) {
	var buf = new(bytes.Buffer)

	err := data.Write(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
