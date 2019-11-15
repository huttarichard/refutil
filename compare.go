package refutil

// Compare will use comparator to compare to another Value
func (v Value) Compare(comparator Comparator, value Value) bool {
	if v.IsNil() || value.IsNil() {
		return v.InterfaceOrNil() == value.InterfaceOrNil()
	}
	return comparator(v.InterfaceOrNil(), value.InterfaceOrNil())
}

// Compare will use comparator to compare to another Value
func (v Data) Compare(comparator Comparator, value Value) bool {
	return v.v.Compare(comparator, value)
}
