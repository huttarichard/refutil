package refutil

// SearchIndex will search for specific element in list
// and return index. If not found returns -1.
// it panics if list is not searchable
func (v Value) SearchIndex(compare Comparator, element interface{}) int {
	keyval := v.SearchValue(compare, element)
	if keyval == nil {
		return -1
	}
	return keyval.Index()
}

// SearchIndex will search for specific element in list
// and return index. If not found returns -1.
// it panics if list is not searchable
func (v Data) SearchIndex(compare Comparator, element interface{}) int {
	return v.v.SearchIndex(compare, element)
}

// Index will search for specific element in list
// and return index. If not found returns -1. If list is
// not list it will panic
func Index(list interface{}, element interface{}) int {
	return NewData(list).SearchIndex(Equal, element)
}

// IndexSame will search for specific element in list
// and return index. If not found returns -1. If list is
// not list it will panic.
// But unlike Index it will compare values so they must be
// deep equal.
func IndexSame(list interface{}, element interface{}) int {
	return NewData(list).SearchIndex(DeepEqual, element)
}
