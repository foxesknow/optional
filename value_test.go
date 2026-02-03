package optional

import "testing"

func TestDefaultState_HasValue(t *testing.T) {
	var value Value[int]
	if value.HasValue() {
		t.Error("should have no value")
	}
}

func TestDefaultState_IsMissing(t *testing.T) {
	var value Value[int]

	if !value.IsMissing() {
		t.Error("should be missing")
	}
}

func TestGet_HasValue(t *testing.T) {
	var value = Some(10)
	if !value.HasValue() {
		t.Error("should have a value")
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
