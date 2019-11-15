package refutil

import "reflect"

// Value is wrapper around reflect.Value
type Value struct {
	reflect.Value
}

// NewValue will create new Value
func NewValue(v interface{}) Value {
	if t, ok := v.(reflect.Value); ok {
		return Value{Value: t}
	}
	if t, ok := v.(Value); ok {
		return Value{Value: t.Value}
	}
	if t, ok := v.(Type); ok {
		return Value{Value: t.Zero().Value}
	}
	return Value{Value: reflect.ValueOf(v)}
}

// Type will wrap original reflect.Type to Type
func (v Value) Type() Type {
	if !v.Value.IsValid() {
		return NewType(nil)
	}
	return NewType(v.Value.Type())
}
