package resp_test

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/FastHCA/resp"
)

func TestArray_Well(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.Array(
		resp.BulkString("hello"),
		resp.BulkString("world"),
	)

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestArray_Nested(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.Array(
		resp.Array(
			resp.BulkString("hello"),
			resp.BulkString("world"),
		),
		resp.Array(
			resp.BulkString("foo"),
			resp.BulkString("bar"),
			resp.Integer(1024),
		),
	)

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte("*2\r\n*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n*3\r\n$3\r\nfoo\r\n$3\r\nbar\r\n:1024\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestArray_Empty(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.Array()

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte("*0\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestBulkString_Well(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.BulkString("hello")

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte("$5\r\nhello\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestBulkString_EmptyString(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.BulkString("")

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte("$0\r\n\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestInteger_PositiveValue(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.Integer(1024)

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte(":1024\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestInteger_NegativeValue(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.Integer(-1023)

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte(":-1023\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestNullArray(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.NullArray()

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte("*-1\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestNullBulkString(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.NullBulkString()

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte("$-1\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestSimpleString(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.SimpleString("OK")

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte("+OK\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}

func TestSimpleError(t *testing.T) {
	writer := new(bytes.Buffer)
	pakcer := resp.SimpleError("WRONGTYPE Operation against a key holding the wrong kind of value")

	err := pakcer.Write(writer)
	if err != nil {
		t.Errorf("should not err, but got:: %+v", err)
	}

	actual := writer.Bytes()

	expected := []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("assert `protocol`:: expected '%+v', got '%+v'", expected, actual)
	}
}
