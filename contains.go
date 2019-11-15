package refutil

// ContainsKey will search for specific key in list
// and return bool if found.  It panics if list is not searchable
func (v Value) ContainsKey(compare Comparator, element interface{}) bool {
	keyval := v.SearchKey(compare, element)
	return keyval != nil
}

// ContainsKey will search for specific key in list
// and return bool if found. It panics if list is not searchable
func (v Data) ContainsKey(compare Comparator, element interface{}) bool {
	return v.v.ContainsKey(compare, element)
}

// ContainsKey will search for specific key in list
// and return bool if found. It panics if list is not searchable
func ContainsKey(source interface{}, element interface{}) bool {
	return NewData(source).ContainsKey(Equal, element)
}

// ContainsSameKey will search for specific key in list
// and return bool if found. It panics if list is not searchable.
// Unlike ContainsKey compared keys needs to be deep equal
func ContainsSameKey(source interface{}, element interface{}) bool {
	return NewData(source).ContainsKey(DeepEqual, element)
}

// ContainsValue will search for specific value in list
// and return bool if found. It panics if list is not searchable
func (v Value) ContainsValue(compare Comparator, element interface{}) bool {
	keyval := v.SearchValue(compare, element)
	return keyval != nil
}

// ContainsValue will search for specific value in list
// and return bool if found.
// It panics if list is not searchable
func (v Data) ContainsValue(compare Comparator, element interface{}) bool {
	return v.v.ContainsValue(compare, element)
}

// ContainsValue will search for specific value in list
// and return bool if found.
// it panics if list is not searchable
func ContainsValue(source interface{}, element interface{}) bool {
	return NewData(source).ContainsValue(Equal, element)
}

// ContainsSameValue will search for specific value in list
// and return bool if found. It panics if list is not searchable.
// Unlike ContainsValue compared keys needs to be deep equal
func ContainsSameValue(source interface{}, element interface{}) bool {
	return NewData(source).ContainsValue(DeepEqual, element)
}

// Contains will search for specific value in list
// and return bool if found. It panics if list is not searchable
func Contains(source interface{}, element interface{}) bool {
	return ContainsValue(source, element)
}
