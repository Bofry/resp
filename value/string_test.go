package value_test

import (
	"math/big"
	"testing"

	"github.com/FastHCA/resp/value"
)

func TestString(t *testing.T) {
	v := value.NewString("16")

	decimal, ok := v.BigNumber()

	expectedOk := true
	if ok != expectedOk {
		t.Errorf("assert 'SimpleStringValue.Decimal() ok':: expected '%+v', got '%+v'", expectedOk, ok)
	}
	expectedDecimal := big.NewRat(16, 1)
	if decimal.RatString() != expectedDecimal.RatString() {
		t.Errorf("assert 'SimpleStringValue.Decimal() decimal':: expected '%+v', got '%+v'", expectedDecimal, decimal)
	}
}
