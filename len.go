package refutil

import (
	"reflect"
)

// CanLen tells you if possible to search for len of
// given reflect.Value.
// https://golang.org/pkg/reflect/#Value.Len
func (v Value) CanLen() bool {
	switch v.Indirect().Kind() {
	case reflect.Map, reflect.Chan, reflect.Array, reflect.Slice, reflect.String:
		return true
	default:
		return false
	}
}

// CanLen tells you if possible to search for len of
// given reflect.Value.
// https://golang.org/pkg/reflect/#Value.Len
func (v Data) CanLen() bool {
	return v.Value().CanLen()
}

// Len try to get length of object. It will wrap original
// value into new value so interface{} can work
func (v Value) Len() (length int) {
	x := reflect.ValueOf(v.Indirect().Interface())
	return x.Len()
}

// Len try to get length of object.
func (v Data) Len() (length int) {
	return v.Value().Len()
}

// CanLen tells you if possible to search for len of
// given object.
// https://golang.org/pkg/reflect/#Value.Len
func CanLen(x interface{}) bool {
	return NewValue(x).CanLen()
}

// Len try to get length of object.
func Len(x interface{}) (length int) {
	return NewData(x).Len()
}
