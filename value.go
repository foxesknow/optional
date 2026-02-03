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

func (v *Value[T]) HasValue() bool {
	return v != nil && v.hasValue
}

func (v *Value[T]) IsMissing() bool {
	return v == nil || !v.hasValue
}

func (v *Value[T]) Get() (T, error) {
	if v.HasValue() {
		return v.value, nil
	} else {
		var def T
		return def, noValue
	}
}

func (v *Value[T]) OrElse(defaultValue T) T {
	if v.HasValue() {
		return v.value
	} else {
		return defaultValue
	}
}

func (v *Value[T]) String() string {
	if v.HasValue() {
		if s, ok := interface{}(v.value).(fmt.Stringer); ok {
			return fmt.Sprintf("Some(%v)", s.String())
		} else {
			return "Some(...)"
		}
	} else {
		return "none"
	}
}
