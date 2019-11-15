// Copyright (c) 2015 Dave Collins <dave@davec.name>

// NOTE: Due to the following build constraints, this file will only be compiled
// when either the code is running on Google App Engine or "-tags disableunsafe"
// is added to the go build command line.
// +build appengine disableunsafe

package unsafe

import "reflect"

const (
	// UnsafeDisabled is a build-time constant which specifies whether or
	// not access to the unsafe package is available.
	UnsafeDisabled = true
)

// ReflectValue typically converts the passed reflect.Value into a one
// that bypasses the typical safety restrictions preventing access to
// unaddressable and unexported data.  However, doing this relies on access to
// the unsafe package.  This is a stub version which simply returns the passed
// reflect.Value when the unsafe package is not available.
func ReflectValue(v reflect.Value) reflect.Value {
	return v
}
