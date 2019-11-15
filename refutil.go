package refutil

// Data is structure supporting this package but also have
// application outside. It helps to simplify working with reflect
type Data struct {
	v Value
	t Type
}

// NewData creates new Data struct
func NewData(v interface{}) Data {
	return Data{
		v: NewValue(v),
		t: NewType(v),
	}
}

// Type will return reflect.TypeOf in wrapper Type
// and cache for later use
func (v Data) Type() Type {
	return v.t
}

// Value will return reflect.ValueOf in wrapper Value
// and cache for later use
func (v Data) Value() Value {
	return v.v
}
