package refutil

import (
	"fmt"
)

// String returns
func (v Value) String() string {
	if !v.CanInterface() {
		return ""
	}
	if v.IsNil() {
		return ""
	}
	switch x := v.Interface().(type) {
	case fmt.Stringer:
		return x.String()
	case string:
		return x
	case *string:
		return *x
	default:
		if v.CanIndirect() {
			return v.Indirect().String()
		}
		return fmt.Sprintf("%v", x)
	}
}

// String returns
func (v Data) String() string {
	return v.Value().String()
}

// String return string from value
func String(v interface{}) string {
	return NewData(v).String()
}
