package refutil

import (
	"reflect"
)

// Type is wrapper around reflect.Type
type Type struct {
	reflect.Type
}

// NewType returns new type from interface{}
func NewType(v interface{}) Type {
	if t, ok := v.(reflect.Type); ok {
		return Type{Type: t}
	}
	return Type{Type: reflect.TypeOf(v)}
}

// New will create new pointer of this type
func (t Type) New() Value {
	return NewValue(reflect.New(t.Type))
}

// NewSlice will create empty slice (not slice header!)
func (t Type) NewSlice() Value {
	sliceType := reflect.SliceOf(t.Elem())
	zeroLenSlice := reflect.MakeSlice(sliceType, 0, 0)
	return NewValue(zeroLenSlice.Convert(t))
}
