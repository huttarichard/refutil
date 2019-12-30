package refutil

import (
	"reflect"

	"testing"

	"github.com/huttarichard/refutil/test"
)

func TestKind(t *testing.T) {
	var s string
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)
	unexported := test.Sample(test.SampleInt)
	val2 := reflect.ValueOf(unexported)
	typ2 := reflect.TypeOf(unexported)
	unexported2 := test.Sample(test.SampleStruct)
	val3 := reflect.ValueOf(unexported2)
	typ3 := reflect.TypeOf(unexported2)
	tests := []struct {
		expected Kinder
		equals   reflect.Kind
		want     bool
	}{
		{typ, reflect.String, true},
		{val, reflect.String, true},
		{typ, reflect.Map, false},
		{typ, reflect.Chan, false},
		{typ, reflect.Bool, false},
		{val, reflect.Map, false},
		{val, reflect.Chan, false},
		{val, reflect.Bool, false},
		{val2, reflect.Bool, false},
		{typ2, reflect.Bool, false},
		{val3, reflect.Ptr, true},
		{typ3, reflect.Ptr, true},
		{val3, reflect.Bool, false},
		{typ3, reflect.Bool, false},
	}
	for _, tt := range tests {
		t.Run("test len", func(t *testing.T) {
			l := IsKind(tt.expected, tt.equals)
			if l != tt.want {
				t.Errorf("IsKind() = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestKindOf(t *testing.T) {
	var s string
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)
	tests := []struct {
		expected Kinder
		equals   []reflect.Kind
		want     bool
	}{
		{typ, []reflect.Kind{reflect.String, reflect.Map}, true},
		{val, []reflect.Kind{reflect.String, reflect.Map}, true},
		{typ, []reflect.Kind{reflect.String}, true},
		{val, []reflect.Kind{reflect.String}, true},
		{typ, []reflect.Kind{reflect.Map}, false},
		{val, []reflect.Kind{reflect.Map}, false},
	}
	for _, tt := range tests {
		t.Run("test len", func(t *testing.T) {
			l := KindOneOf(tt.expected, tt.equals...)
			if l != tt.want {
				t.Errorf("KindOneOf() = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestKindIs(t *testing.T) {
	t1 := NewType(struct{}{})
	if !t1.IsKind(reflect.Struct) {
		t.Fatal()
	}
	if !t1.KindOneOf(reflect.Struct, reflect.Ptr) {
		t.Fatal()
	}
	v1 := NewValue(struct{}{})
	if !v1.IsKind(reflect.Struct) {
		t.Fatal()
	}
	if !v1.KindOneOf(reflect.Struct, reflect.Ptr) {
		t.Fatal()
	}
	d1 := NewData(struct{}{})
	if !d1.IsKind(reflect.Struct) {
		t.Fatal()
	}
	if !d1.KindOneOf(reflect.Struct, reflect.Ptr) {
		t.Fatal()
	}
}
