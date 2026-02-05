package optional

import (
	"encoding/json"
	"errors"
	"fmt"
)

// https://fsharp.github.io/fsharp-core-docs/reference/fsharp-core-optionmodule.html

// Represents an optional value
type Value[T any] struct {
	hasValue bool
	value    T
}

var noValue = errors.New("no value")

// Creates an optional value that has a value
func Some[T any](value T) Value[T] {
	return Value[T]{
		hasValue: true,
		value:    value,
	}
}

// Creates an optional value that has no value
func None[T any]() Value[T] {
	var v Value[T]
	return v
}

// Calls the mapper function if "v" contains a value, otherwise returns none.
func Map[T, R any](v Value[T], mapper func(T) R) Value[R] {
	if v.hasValue {
		return Some(mapper(v.value))
	}

	return None[R]()
}

// Calls the mapper function if "v1" and "v2" contain a value, otherwise returns none.
func Map2[T1, T2, V any](v1 Value[T1], v2 Value[T2], mapper func(T1, T2) V) Value[V] {
	if v1.hasValue && v2.hasValue {
		return Some(mapper(v1.value, v2.value))
	}

	return None[V]()
}

// Returns the inner optional value
func Unpack[T any](v Value[Value[T]]) Value[T] {
	return v.value
}

// Reports if the optional value contains an actual value
func (v Value[T]) IsSome() bool {
	return v.hasValue
}

// Reports if the optional value does not contain a value
func (v Value[T]) IsNone() bool {
	return !v.hasValue
}

// Attempts to get the value held in the optional value
func (v Value[T]) Get() (T, error) {
	if v.hasValue {
		return v.value, nil
	}

	var def T
	return def, noValue
}

// Returns the value held in the optional. If there is no value then panic.
func (v Value[T]) MustGet() T {
	if v.hasValue {
		return v.value
	}

	panic("no value in optional value")
}

// Returns the value held. If there is no value then returns 'defaultValue'
func (v Value[T]) OrElse(defaultValue T) T {
	if v.hasValue {
		return v.value
	}

	return defaultValue
}

// Returns the value if held. If there is no value then calls the factory function to create a value
func (v Value[T]) OrElseWith(factory func() T) T {
	if v.hasValue {
		return v.value
	}

	return factory()
}

// Returns a slice containing the value. If no value then an empty slice is returned.
func (v Value[T]) ToSlice() []T {
	if v.hasValue {
		return []T{v.value}
	}

	return []T{}
}

// Returns the optional value as a string
func (v Value[T]) String() string {
	if v.hasValue {
		return fmt.Sprintf("Some(%v)", v.value)
	}

	return "None"
}

// Marshals the optional value to json. If there is no value then `null` is written to the json
func (v Value[T]) MarshalJSON() ([]byte, error) {
	if v.hasValue {
		return json.Marshal(v.value)
	}

	return json.Marshal(nil)
}

// Populates a value from json
func (v *Value[T]) UnmarshalJSON(data []byte) error {
	if len(data) > 3 && data[0] == 'n' && data[1] == 'u' && data[2] == 'l' && data[3] == 'l' {
		v.hasValue = false
		var def T
		v.value = def

		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	v.hasValue = true
	v.value = value
	return nil
}

// Allows for ranged iteration. If there is no value then no iteration takes place
func (v Value[T]) Iterate(yield func(T) bool) {
	if v.hasValue {
		yield(v.value)
	}
}
