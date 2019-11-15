package refutil

import (
	"reflect"
)

// CanIndirect will determine if type can be dereferenced
// refer to https://golang.org/pkg/reflect/#Type
// method Elem()
func (v Type) CanIndirect() bool {
	if v.Type == nil {
		return false
	}
	switch v.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Array:
		return true
	default:
		return false
	}
}

// CanIndirectType will determine if type can be dereferenced
// refer to https://golang.org/pkg/reflect/#Type
// method Elem()
func CanIndirectType(t reflect.Type) bool {
	return NewType(t).CanIndirect()
}

// Indirect will check if type can be dereferenced
// and do so, otherwise return type itself
// panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
func (v Type) Indirect() Type {
	if v.CanIndirect() {
		return NewType(v.Elem())
	}
	return v
}

// IndirectType will check if type can be dereferenced
// and do so, otherwise return type itself
// panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
func IndirectType(t reflect.Type) Type {
	return NewType(t).Indirect()
}

// IndirectTypeOf is shortcut to TypeOf().Elem().
// only indirect if possible otherwise it returns current type
func IndirectTypeOf(x interface{}) Type {
	return NewData(x).Type().Indirect()
}

// CanIndirect will determine if value can be dereferenced
// refer to https://golang.org/pkg/reflect/#Value.Elem
func (v Value) CanIndirect() bool {
	return v.IsKind(reflect.Ptr)
}

// CanIndirectValue will determine if value can be dereferenced
// refer to https://golang.org/pkg/reflect/#Value.Elem
func CanIndirectValue(v reflect.Value) bool {
	return NewValue(v).CanIndirect()
}

// Indirect will check if value can be dereferenced
// and dereference otherwise return value itself
func (v Value) Indirect() Value {
	if v.CanIndirect() {
		return NewValue(v.Elem())
	}
	return v
}

// IndirectValue will check if value can be dereferenced
// and dereference otherwise return value itself
func IndirectValue(v reflect.Value) Value {
	return NewValue(v).Indirect()
}

// IndirectValueOf is shortcut to ValueOf().Elem()
func IndirectValueOf(v interface{}) Value {
	return NewData(v).Value().Indirect()
}

// CanIndirect will determine if value can be dereferenced
// refer to https://golang.org/pkg/reflect/#Value.Elem
func (v Data) CanIndirect() bool {
	return v.Value().CanIndirect()
}

// CanIndirect will determine if value can be dereferenced
// refer to https://golang.org/pkg/reflect/#Value.Elem
func CanIndirect(v interface{}) bool {
	return NewData(v).CanIndirect()
}

// Indirect will check if value can be dereferenced
// and dereference otherwise return value itself
func (v Data) Indirect() Data {
	val := v.Value()
	if !val.CanIndirect() {
		return v
	}
	data := NewData(val.Indirect().InterfaceOrNil())
	return data
}

// Indirect will return actual value insted of pointer
// if values is passed value is returned
func Indirect(v interface{}) interface{} {
	return NewData(v).Indirect().Interface()
}
