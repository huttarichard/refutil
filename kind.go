package refutil

import "reflect"

// KindOneOf return true if type is one of the kind supplied
func KindOneOf(t Kinder, kinds ...reflect.Kind) bool {
	searched := t.Kind()
	for _, k := range kinds {
		if k == searched {
			return true
		}
	}
	return false
}

// IsKind return true if type or value is specific kind
func IsKind(t Kinder, kind reflect.Kind) bool {
	return t.Kind() == kind
}

// IsKind is sugar to simplify conditions
func (t Type) IsKind(k reflect.Kind) bool {
	return IsKind(t, k)
}

// KindOneOf is method to compare type against multiple
// reflect.Type with logical OR
func (t Type) KindOneOf(k ...reflect.Kind) bool {
	return KindOneOf(t, k...)
}

// IsKind is sugar to simplify conditions
func (t Value) IsKind(k reflect.Kind) bool {
	return IsKind(t, k)
}

// KindOneOf is method to compare type against multiple
// reflect.Type with logical OR
func (t Value) KindOneOf(k ...reflect.Kind) bool {
	return KindOneOf(t, k...)
}

// IsKind reports if data is specific kind
func (v Data) IsKind(k reflect.Kind) bool {
	return v.Value().IsKind(k)
}

// KindOneOf reports if data is one if expected kinds
func (v Data) KindOneOf(k ...reflect.Kind) bool {
	return v.Value().KindOneOf(k...)
}
