package refutil

// Search will iterate over struct map or slice
// providing key value of every element and
// return first result when iterator return true
func (v Value) Search(iterator SearchFunc) *KeyValue {
	var found *KeyValue
	v.Iterate(func(it *Iterator) {
		c := it.Current()
		if iterator(c) {
			found = c
			it.Stop()
		}
	})
	return found
}

// Search will iterate over struct map or slice
// providing key value of every element and
// return first result when iterator return true
func (v Data) Search(iterator SearchFunc) *KeyValue {
	return v.v.Search(iterator)
}

// Search will iterate over struct map or slice
// providing key value of every element and
// return first result when iterator return true
func Search(source interface{}, iterator SearchFunc) *KeyValue {
	return NewData(source).Search(iterator)
}

// SearchKey will search for specific key in struct slice or map
func (v Value) SearchKey(compare Comparator, element interface{}) *KeyValue {
	val := NewValue(element)
	return v.Search(func(keyVal *KeyValue) bool {
		if keyVal.Key.Compare(compare, val) {
			return true
		}
		return false
	})
}

// SearchKey will search for specific key in struct slice or map
func (v Data) SearchKey(compare Comparator, element interface{}) *KeyValue {
	return v.v.SearchKey(compare, element)
}

// SearchKey will search for specific key in struct slice or map
func SearchKey(source interface{}, element interface{}) *KeyValue {
	d := NewData(source).SearchKey(Equal, element)
	if d == nil {
		return nil
	}
	return d
}

// SearchSameKey will search for specific key in struct slice or map.
// Unlike SearchKey it the compared keys needs to be deep equal
func SearchSameKey(source interface{}, element interface{}) *KeyValue {
	d := NewData(source).SearchKey(DeepEqual, element)
	if d == nil {
		return nil
	}
	return d
}

// SearchValue will search for specific value in struct slice or map
func (v Value) SearchValue(compare Comparator, element interface{}) *KeyValue {
	val := NewValue(element)
	return v.Search(func(keyVal *KeyValue) bool {
		if keyVal.Value.Compare(compare, val) {
			return true
		}
		return false
	})
}

// SearchValue will search for specific value in struct slice or map
func (v Data) SearchValue(compare Comparator, element interface{}) *KeyValue {
	return v.v.SearchValue(compare, element)
}

// SearchValue will search for specific value in struct slice or map
func SearchValue(source interface{}, element interface{}) *KeyValue {
	d := NewData(source).SearchValue(Equal, element)
	if d == nil {
		return nil
	}
	return d
}

// SearchSameValue will search for specific value in struct slice or map.
// Unlike SearchKey it compared keys needs to be deep equal
func SearchSameValue(source interface{}, element interface{}) *KeyValue {
	d := NewData(source).SearchValue(DeepEqual, element)
	if d == nil {
		return nil
	}
	return d
}
