package optional

import (
	"testing"
)

func TestDefaultState_None(t *testing.T) {
	var value Value[int]
	if value.IsSome() {
		t.Error("should have no value")
	}

	if !value.IsNone() {
		t.Error("should be none")
	}
}

func TestGet_HasValue(t *testing.T) {
	var value = Some(10)
	if !value.IsSome() {
		t.Error("should have a value")
	}

	if value.IsNone() {
		t.Error("value is not none")
	}

	i, err := value.Get()

	if err != nil {
		t.Error("error should be nil")
	}

	if i != 10 {
		t.Error("should have got 10")
	}
}

func TestGet_NotInitializes(t *testing.T) {
	var value Value[int]
	i, err := value.Get()

	if err == nil {
		t.Error("error should not be nil")
	}

	if i != 0 {
		t.Error("should have got 0")
	}
}

func TestMap_HasValue(t *testing.T) {
	v := Some("hello")
	length := Map(&v, func(x string) int { return len(x) })

	if length.OrElse(0) != 5 {
		t.Error("should have 5")
	}
}

func TestMap_NoValue(t *testing.T) {
	v := None[string]()
	length := Map(&v, func(x string) int { return len(x) })

	if length.IsSome() {
		t.Error("length should not have a value")
	}
}
