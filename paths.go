package refutil

// PathTo will return actual path to package
// where interface{} is defined.
func (v Type) PathTo() string {
	return v.Indirect().Type.PkgPath()
}

// PathTo will return actual path to package
// where interface{} is defined.
func (v Data) PathTo() string {
	return v.Type().PathTo()
}

// PathTo will return actual path to package
// where interface{} is defined.
func PathTo(v interface{}) string {
	return NewData(v).PathTo()
}
