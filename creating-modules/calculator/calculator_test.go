package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	x, y := 5, 8
	result := x + y

	num, err := Add(x, y)

	if num != result || err != nil {
		t.Fatalf(`expected %v as result for inputs %v and %v, instead got %v`, result, x, y, num)
	}
}

func TestAddError(t *testing.T) {
	x, y := 0, 0

	num, err := Add(x, y)

	if num != 0 || err == nil {
		t.Fatalf("Expected an error here but received value")
	}
}