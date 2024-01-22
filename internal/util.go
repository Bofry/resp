package internal

import (
	"bufio"
	"fmt"
	"io"
)

func acquireBufioWriter(writer io.Writer) *bufio.Writer {
	w, ok := writer.(*bufio.Writer)
	if !ok {
		w = bufio.NewWriter(writer)
	}
	return w
}

func validateSimpleValue(b []byte) error {
	for i, c := range b {
		if c == '\r' || c == '\n' {
			return fmt.Errorf("invalid character %q at %d", c, i)
		}
	}
	return nil
}
