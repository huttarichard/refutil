package refutil

// KeyValue provides key value. Key just like
// value can be any interface. If slice key will
// be int wrapped in Value
type KeyValue struct {
	Key   Value
	Value Value
}

// Index will return int as an index taken
// from key. This works only for slices
func (kv *KeyValue) Index() int {
	return kv.Key.Interface().(int)
}
