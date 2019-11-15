package refutil

import (
	"reflect"
)

// Zeroer is like assertion with adition check
// if value can be used for type assertion to Zeroer
func (v Value) Zeroer() (Zeroer, bool) {
	if !v.IsValid() {
		return nil, false
	}
	if v.IsNil() {
		return nil, false
	}
	z, k := v.Interface().(Zeroer)
	return z, k
}

// IsZero return true if underlying type is equal to its zero value
func (v Value) IsZero() bool {
	if zeroer, k := v.Zeroer(); k {
		return zeroer.IsZero()
	}
	if v.IsNil() {
		return true
	}
	t := v.Type()
	switch t.Kind() {
	case reflect.Map:
		return v.Len() == 0
	case reflect.Chan:
		return v.Len() == 0
	case reflect.Slice:
		s := t.NewSlice()
		return reflect.DeepEqual(v.Interface(), s.Interface())
	case reflect.Ptr:
		return v.Indirect().IsZero()
	default:
		return reflect.DeepEqual(v.Interface(), t.Zero().InterfaceOrNil())
	}
}

// IsZero return true if underlying type is equal to its zero value
func (v Data) IsZero() bool {
	return v.Value().IsZero()
}

// IsZero return true if underlying type is equal to its zero value
func IsZero(v interface{}) bool {
	return NewData(v).IsZero()
}

// Zero will return zero of type
func (t Type) Zero() Value {
	if t.IsNil() {
		return NewValue(nil)
	}
	return NewValue(reflect.Zero(t.Type))
}

// Zero will return zero of its own type
func (v Value) Zero() Value {
	return v.Type().Zero()
}
