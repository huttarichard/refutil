package refutil

import (
	"reflect"
)

// DeepEqual determines if two Objects are considered equal.
// This is what reflect does, except this method add nil check
// for speed improvement
func (v Data) DeepEqual(element interface{}) bool {
	data := v.InterfaceOrNil()
	if data == nil || element == nil {
		return data == element
	}
	return reflect.DeepEqual(data, element)
}

// DeepEqual determines if two Objects are considered equal.
// This is what reflect does, except this method add nil check
// for speed improvement
func DeepEqual(expected, actual interface{}) bool {
	return NewData(expected).DeepEqual(actual)
}

// Equal gets whether two Objects are equal.
// It don't cares about underlying type.
// If first argument is compliant with `Equalizer` interface
// then it will use `Equal(interface) bool` on type to compare.
// Otherwise it try to compare using `DeepEqual` and same types.
func (v Data) Equal(element interface{}) bool {
	data := v.InterfaceOrNil()
	equalizer, ok := data.(Equalizer)
	if ok {
		return equalizer.Equal(element)
	}
	el := NewData(element)
	if v.IsNil() || el.IsNil() {
		su := v.Indirect().Untyped().InterfaceOrNil()
		eu := el.Indirect().Untyped().InterfaceOrNil()
		return su == eu
	}
	if v.DeepEqual(element) {
		return true
	}
	at := el.Type()
	ev := v.Value().Indirect()
	if ev.IsValid() && ev.Type().ConvertibleTo(at.Type) {
		// Attempt comparison after type conversion
		return reflect.DeepEqual(ev.Convert(at).Interface(), element)
	}
	return false
}

// Equal gets whether two Objects are equal.
// It don't cares about underlying type.
// If first argument is compliant with `Equalizer` interface
// then it will use `Equal(interface) bool` on type to compare.
// Otherwise it try to compare using `DeepEqual` and same types.
func Equal(expected, actual interface{}) bool {
	return NewData(expected).Equal(actual)
}
