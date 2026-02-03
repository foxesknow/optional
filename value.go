package optional

import (
	"errors"
	"fmt"
)

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

func Map[T, V any](v *Value[T], mapper func(T) V) Value[V] {
	if v.IsSome() {
		return Some(mapper(v.value))
	} else {
		return None[V]()
	}
}

func (v *Value[T]) IsSome() bool {
	return v != nil && v.hasValue
}

func (v *Value[T]) IsNone() bool {
	return v == nil || !v.hasValue
}

func (v *Value[T]) Get() (T, error) {
	if v.IsSome() {
		return v.value, nil
	} else {
		var def T
		return def, noValue
	}
}

func (v *Value[T]) OrElse(defaultValue T) T {
	if v.IsSome() {
		return v.value
	} else {
		return defaultValue
	}
}

func (v *Value[T]) String() string {
	if v.IsSome() {
		if s, ok := interface{}(v.value).(fmt.Stringer); ok {
			return fmt.Sprintf("Some(%v)", s.String())
		} else {
			return "Some(...)"
		}
	} else {
		return "none"
	}
}
