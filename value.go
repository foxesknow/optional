package optional

import (
	"errors"
	"fmt"
)

// https://fsharp.github.io/fsharp-core-docs/reference/fsharp-core-optionmodule.html

type Value[T any] struct {
	hasValue bool
	value    T
}

var noValue = errors.New("no value")

func Some[T any](value T) Value[T] {
	return Value[T]{
		hasValue: true,
		value:    value,
	}
}

func None[T any]() Value[T] {
	var v Value[T]
	return v
}

func Map[T, V any](v Value[T], mapper func(T) V) Value[V] {
	if v.IsSome() {
		return Some(mapper(v.value))
	} else {
		return None[V]()
	}
}

func Unpack[T any](v Value[Value[T]]) Value[T] {
	return v.value
}

func (v Value[T]) IsSome() bool {
	return v.hasValue
}

func (v Value[T]) IsNone() bool {
	return !v.hasValue
}

func (v Value[T]) Get() (T, error) {
	if v.IsSome() {
		return v.value, nil
	} else {
		var def T
		return def, noValue
	}
}

func (v Value[T]) OrElse(defaultValue T) T {
	if v.IsSome() {
		return v.value
	} else {
		return defaultValue
	}
}

func (v Value[T]) OrElseWith(factory func() T) T {
	if v.IsSome() {
		return v.value
	} else {
		return factory()
	}
}

func (v Value[T]) ToSlice() []T {
	if v.IsSome() {
		return []T{v.value}
	} else {
		return []T{}
	}
}

func (v Value[T]) String() string {
	if v.IsSome() {
		return fmt.Sprintf("Some(%v)", v.value)
	} else {
		return "None"
	}
}
