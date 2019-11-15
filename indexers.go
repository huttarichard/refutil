package refutil

import (
	"reflect"
)

// StructIndexer implements Indexer for
// structs
type StructIndexer struct {
	Source Value
	keys   []Value
}

var _ Indexer = (*StructIndexer)(nil)

// NewStructIndexer will create StructIndexer
func NewStructIndexer(source Value) *StructIndexer {
	var keys []Value
	t := source.Type()
	for i := 0; i < source.NumField(); i++ {
		keys = append(keys, NewValue(t.Field(i).Name))
	}
	s := ValueSorter(keys)
	s.Sort()
	return &StructIndexer{Source: source, keys: s}
}

// At will return key value for index i
func (v *StructIndexer) At(i int) (Value, Value) {
	key := v.keys[i]
	field := v.Source.FieldByName(key.String())
	return NewValue(key), NewValue(field)
}

// Keys will return list of keys
func (v *StructIndexer) Keys() []Value {
	return v.keys
}

// MapIndexer implements Indexer for
// maps
type MapIndexer struct {
	Source Value
	keys   []Value
}

var _ Indexer = (*MapIndexer)(nil)

// NewMapIndexer will create MapIndexer
func NewMapIndexer(source Value) *MapIndexer {
	keys := source.MapKeys()
	var arr []Value
	for _, k := range keys {
		arr = append(arr, NewValue(k))
	}
	s := ValueSorter(arr)
	s.Sort()
	return &MapIndexer{Source: source, keys: s}
}

// At will return key value for index i
func (v *MapIndexer) At(i int) (Value, Value) {
	key := v.keys[i]
	field := v.Source.MapIndex(key.Value)
	return key, NewValue(field)
}

// Keys will return list of keys
func (v *MapIndexer) Keys() []Value {
	return v.keys
}

// SliceIndexer implements Indexer for
// slices or arrays
type SliceIndexer struct {
	Source Value
	keys   []Value
}

var _ Indexer = (*SliceIndexer)(nil)

// NewSliceIndexer will create SliceIndexer
func NewSliceIndexer(source Value) *SliceIndexer {
	var keys []Value
	for i := 0; i < source.Len(); i++ {
		keys = append(keys, NewValue(i))
	}
	return &SliceIndexer{Source: source, keys: keys}
}

// At will return key value for index i
func (v *SliceIndexer) At(i int) (Value, Value) {
	return NewValue(i), v.Source.Index(i)
}

// Keys will return list of keys
func (v *SliceIndexer) Keys() []Value {
	return v.keys
}

// CanIndex tells you if you can get value from index.
// Checkout https://golang.org/pkg/reflect/#Value.Index
func (v Value) CanIndex() bool {
	nv := v.Indirect()
	if nv.KindOneOf(reflect.String, reflect.Slice, reflect.Array) {
		return true
	}
	return false
}

// Index is just wrapper of regular index
// method. It will indirect value if is pointer
func (v Value) Index(i int) Value {
	nv := v.Indirect()
	return NewValue(nv.Value.Index(i))
}

// CanUseIndexer return bool if Indexer Method can be used
func (v Value) CanUseIndexer() bool {
	nv := v.Indirect()
	if nv.KindOneOf(reflect.Struct, reflect.Map, reflect.Slice, reflect.Array) {
		return true
	}
	return false
}

// Indexer will return proper Indexer interface
func (v Value) Indexer() Indexer {
	nv := v.Indirect()
	switch nv.Kind() {
	case reflect.Struct:
		return NewStructIndexer(nv)
	case reflect.Map:
		return NewMapIndexer(nv)
	case reflect.Slice, reflect.Array:
		return NewSliceIndexer(nv)
	default:
		panic(ErrArgumentNotIndexable)
	}
}

// At returns key value of given object if `Indexer`
// interface can be used
func (v Value) At(i int) (Value, Value) {
	return v.Indexer().At(i)
}

// At returns key value of given object if `Indexer`
// interface can be used
func (v Data) At(i int) (Value, Value) {
	k, x := v.v.Indexer().At(i)
	return k, x
}
