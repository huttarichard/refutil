package refutil

import (
	"reflect"
)

// IsNil will check if you can ask reflect for IsNil
// and if its pointer it will recursively check until its resolved
func (v Value) IsNil() bool {
	if !v.IsValid() {
		return true
	}
	kind := v.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && v.Value.IsNil() {
		return true
	}
	if v.IsKind(reflect.Ptr) {
		return v.Indirect().IsNil()
	}
	return false
}

// Untyped will return value it self if not nil. If nil
// rather then using value's nil it uses raw nil. This
// is useful in cases where you want insted of (*string)(nil) just nil
func (v Value) Untyped() Value {
	if v.IsNil() {
		return NewValue(nil)
	}
	return v
}

// Untyped will return value it self if not nil. If nil
// rather then using value's nil it uses raw nil. This
// is useful in cases where you want instead of (*string)(nil) just nil
func (v Data) Untyped() Data {
	if v.IsNil() {
		return NewData(nil)
	}
	return v
}

// IsNil checks if type is nil
func (v Type) IsNil() bool {
	return v.Type == nil
}

// IsNil checks if a specified object is nil or not
// also check underlying type if its nil or not
func (v Data) IsNil() bool {
	return v.Value().IsNil()
}

// IsNil checks if a specified object is nil or not
// also check underlying type if its nil or not
func IsNil(object interface{}) bool {
	return NewData(object).IsNil()
}
