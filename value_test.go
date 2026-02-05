package optional

import (
	"encoding/json"
	"testing"
)

type CustomData struct {
	Name    string        `json:"name"`
	Age     int           `json:"age"`
	Address Value[string] `json:"address"`
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

func TestMustGet_NotInitializes(t *testing.T) {
	var value Value[int]

	defer func() { recover() }()
	value.MustGet()
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

func TestMap2(t *testing.T) {
	adder := func(v1 int, v2 int) int { return v1 + v2 }

	if v := Map2(Some(1), Some(2), adder); v.MustGet() != 3 {
		t.Error("should have 3")
	}

	if v := Map2(Some(1), None[int](), adder); !v.IsNone() {
		t.Error("should be none")
	}

	if v := Map2(None[int](), Some(1), adder); !v.IsNone() {
		t.Error("should be none")
	}

	if v := Map2(None[int](), None[int](), adder); !v.IsNone() {
		t.Error("should be none")
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

func TestJson_Int(t *testing.T) {
	v := Some(10)
	b, _ := json.Marshal(v)
	s := string(b)

	if s != "10" {
		t.Error("should be 10")
	}

	var roundTrip Value[int]
	if err := json.Unmarshal(b, &roundTrip); err != nil {
		t.Error(err)
	}

	if i := roundTrip.MustGet(); i != 10 {
		t.Error("should be 10")
	}
}

func TestJson_Int_None(t *testing.T) {
	v := None[int]()
	b, _ := json.Marshal(v)
	s := string(b)

	if s != "null" {
		t.Error("should be null")
	}

	var roundTrip Value[int]
	if err := json.Unmarshal(b, &roundTrip); err != nil {
		t.Error(err)
	}

	if !roundTrip.IsNone() {
		t.Error("should be none")
	}
}

func TestJson_EmbeddedInString_None(t *testing.T) {
	v := CustomData{
		Name: "Jack",
		Age:  41,
	}
	b, _ := json.Marshal(v)

	var roundTrip CustomData
	if err := json.Unmarshal(b, &roundTrip); err != nil {
		t.Error(err)
	}

	if roundTrip.Address.IsSome() {
		t.Error("round trip Address should be none")
	}
}

func TestJson_EmbeddedInString_Some(t *testing.T) {
	v := CustomData{
		Name:    "Jack",
		Age:     41,
		Address: Some("The Island"),
	}
	b, _ := json.Marshal(v)

	var roundTrip CustomData
	if err := json.Unmarshal(b, &roundTrip); err != nil {
		t.Error(err)
	}

	if address, _ := roundTrip.Address.Get(); address != "The Island" {
		t.Error("round trip Address is wrong")
	}
}

func TestJson_CustomData_Some(t *testing.T) {
	v := Some(CustomData{
		Name:    "Jack",
		Age:     41,
		Address: Some("The Island"),
	})
	b, _ := json.Marshal(v)

	var roundTrip Value[CustomData]
	if err := json.Unmarshal(b, &roundTrip); err != nil {
		t.Error(err)
	}

	if !roundTrip.IsSome() {
		t.Error("round should be some")
	}

	value := roundTrip.MustGet()
	if value.Name != "Jack" {
		t.Error("expected Jack")
	}

	if value.Age != 41 {
		t.Error("expected 10")
	}

	if value.Address.IsNone() {
		t.Error("expected Address")
	}
}

func TestJson_Iterate_Some(t *testing.T) {
	total := 100
	v := Some(5)

	for i := range v.Iterate {
		total += i
	}

	if total != 105 {
		t.Error("expected 105")
	}
}

func TestJson_Iterate_None(t *testing.T) {
	total := 100
	v := None[int]()
	looped := false

	for i := range v.Iterate {
		total += i
		looped = true
	}

	if looped {
		t.Error("should not have iterated")
	}

	if total != 100 {
		t.Error("expected 100")
	}
}
