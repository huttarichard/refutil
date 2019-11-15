package refutil

import "reflect"

// Equalizer is way how to check if object
// can compare it self to another object. Used
// only with `Equal` method
type Equalizer interface {
	Equal(v interface{}) bool
}

// Kinder as interface for reflect.Value or reflect.Type Kind() method
type Kinder interface {
	Kind() reflect.Kind
}

// Comparator is used as abstraction for method
// like IsEqual which compare two values
type Comparator func(interface{}, interface{}) bool

// Zeroer ...
type Zeroer interface {
	IsZero() bool
}

// Indexer provides unified searching for key, value
// in map, struct or slice.
type Indexer interface {
	At(i int) (Value, Value)
	Keys() []Value
}

// IteratorFunc function iterates over Indexer
// providing key and value
type IteratorFunc func(*Iterator)

// SearchFunc function iterates over Indexer
// providing key and value and return result if
// result of function is true
type SearchFunc func(*KeyValue) bool
