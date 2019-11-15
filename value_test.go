package refutil

import (
	"reflect"
	"testing"
)

func TestNewValue(t *testing.T) {
	v1 := NewValue(NewType(struct{}{}))
	if !reflect.DeepEqual(v1.InterfaceOrNil(), struct{}{}) {
		t.Fatal()
	}
	v2 := NewValue(NewType(nil))
	if !reflect.DeepEqual(v2.InterfaceOrNil(), nil) {
		t.Fatal()
	}
	t1 := NewValue(nil)
	_ = Value(t1)
	t2 := NewValue(t)
	_ = Value(t2)
	t3 := NewValue(reflect.ValueOf(nil))
	_ = Value(t3)
	t4 := NewValue(NewValue(nil))
	_ = Value(t4)
}

func TestValueType(t *testing.T) {
	v := NewValue(nil)
	if !v.Type().IsNil() {
		t.Fatal()
	}
}
