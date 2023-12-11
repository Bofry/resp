package value_test

import (
	"reflect"
	"testing"

	"github.com/FastHCA/resp/value"
)

func TestArray_Map(t *testing.T) {
	arr := value.NewArray(
		value.NewString("foo"),
		value.NewString("FOO"),
		value.NewString("bar"),
		value.NewString("BAR"),
	)

	m, ok := arr.Map()

	expectedOK := true
	if ok != expectedOK {
		t.Errorf("assert 'ArrayValue.Map() ok':: expected '%+v', got '%+v'", expectedOK, ok)
	}

	_ = m

	expectedMap := map[string]value.Value{
		"foo": value.NewString("FOO"),
		"bar": value.NewString("BAR"),
	}
	if !reflect.DeepEqual(m, expectedMap) {
		t.Errorf("assert 'ArrayValue.Map() map':: expected '%+v', got '%+v'", expectedMap, m)
	}
}
