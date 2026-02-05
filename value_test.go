package optional

import (
	"testing"
)

type CustomData struct {
	Name string
	Age  int
}

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
	length := Map(v, func(x string) int { return len(x) })

	if length.OrElse(0) != 5 {
		t.Error("should have 5")
	}
}

func TestMap_NoValue(t *testing.T) {
	v := None[string]()
	length := Map(v, func(x string) int { return len(x) })

	if length.IsSome() {
		t.Error("length should not have a value")
	}
}

func TestUnpack_NoValue(t *testing.T) {
	var v Value[Value[int]]

	if (Unpack(v)).IsSome() {
		t.Error("should be none")
	}
}

func TestUnpack_HasValue(t *testing.T) {
	v := Some(Some(10))
	p := Unpack(v)

	if !p.IsSome() {
		t.Error("should be some")
	}

	if p.OrElse(-1) != 10 {
		t.Error("should be 10")
	}
}

func TestOrElse(t *testing.T) {
	v := None[int]()
	if v.OrElse(10) != 10 {
		t.Error("expected 10")
	}

	if Some(20).OrElse(99) != 20 {
		t.Error("expected 20")
	}
}

func TestToSlice_NoValue(t *testing.T) {
	v := None[string]()
	slice := v.ToSlice()

	if len(slice) != 0 {
		t.Error("slice should be empty")
	}
}

func TestToSlice_Value(t *testing.T) {
	v := Some("Jack")
	slice := v.ToSlice()

	if len(slice) != 1 {
		t.Error("slice should have one item")
	}
}

func TestOrElseWith(t *testing.T) {
	v := None[int]()
	if v.OrElseWith(func() int { return 10 }) != 10 {
		t.Error("expected 10")
	}

	if Some(20).OrElseWith(func() int { return 99 }) != 20 {
		t.Error("expected 20")
	}
}

func TestString_None(t *testing.T) {
	v := None[int]()
	if v.String() != "None" {
		t.Error("expected None")
	}
}

func TestString_Int(t *testing.T) {
	v := Some(10)
	if v.String() != "Some(10)" {
		t.Error("expected Some(10)")
	}
}

func TestString_String(t *testing.T) {
	v := Some("Jack")
	if v.String() != "Some(Jack)" {
		t.Error("expected Some(Jack)")
	}
}

func TestString_Struct(t *testing.T) {
	customData := CustomData{
		Name: "Jack",
		Age:  41,
	}

	v := Some(customData)
	if s := v.String(); s == "None" {
		t.Error("expected something")
	}
}
