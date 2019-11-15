package refutil

// InterfaceOrNil will either return interface or
// nil if interface cannot be obtained
func (v Value) InterfaceOrNil() interface{} {
	if !v.Value.IsValid() {
		return nil
	}
	// if !v.Value.CanInterface() {
	// 	return nil
	// }
	return v.Value.Interface()
}

// CanInterface returns whether value can interface
func (v Value) CanInterface() bool {
	if !v.Value.IsValid() {
		return false
	}
	return v.Value.CanInterface()
}

// CanInterface returns whether value can interface
func (v Data) CanInterface() bool {
	return v.Value().CanInterface()
}

// Interface returns actual data originally inserted
func (v Data) Interface() interface{} {
	return v.Value().Interface()
}

// InterfaceOrNil will either return interface or
// nil if interface cannot be obtained
func (v Data) InterfaceOrNil() interface{} {
	return v.Value().InterfaceOrNil()
}
