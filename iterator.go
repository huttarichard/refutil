package refutil

// Iterator designed to manage iterations
// over indexer
type Iterator struct {
	value   Value
	indexer Indexer
	index   int
	stop    bool
}

// NewIterator creates new iterator on value
func NewIterator(v Value) *Iterator {
	return &Iterator{
		value:   v,
		indexer: v.Indexer(),
		index:   0,
		stop:    false,
	}
}

// Index return current iteration index
func (s *Iterator) Index() int {
	return s.index
}

// Len will provide how many keys we are iterating over
func (s *Iterator) Len() int {
	return len(s.indexer.Keys())
}

// HasNext will tell you if next element is available
func (s *Iterator) HasNext() bool {
	return s.index+1 < s.Len()
}

// CanNext consider if user wants to stop and if
// next element is available
func (s *Iterator) CanNext() bool {
	return !s.stop && s.HasNext()
}

// Current return current key value
func (s *Iterator) Current() *KeyValue {
	k, v := s.indexer.At(s.index)
	return &KeyValue{Key: k, Value: v}
}

// Stop will stop iteration
func (s *Iterator) Stop() {
	s.stop = true
}

func (s *Iterator) bump() {
	s.index++
}

// Iterator will create and return iterator struct
func (v Value) Iterator() *Iterator {
	return NewIterator(v)
}

// Iterator will create and return iterator struct
func (v Data) Iterator() *Iterator {
	return v.v.Iterator()
}

// Iterate will iterate over struct map or slice
// providing iterator object
func (v Value) Iterate(iterator IteratorFunc) {
	if v.IsNil() {
		return
	}
	it := v.Iterator()
	for {
		iterator(it)
		if !it.CanNext() {
			break
		}
		it.bump()
	}
}

// Iterate will iterate over struct map or slice
// providing key value of every element
func (v Data) Iterate(iterator IteratorFunc) {
	v.v.Iterate(iterator)
}
