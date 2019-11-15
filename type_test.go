package refutil

import (
	"reflect"
	"testing"
)

func TestNewType(t *testing.T) {
	t1 := NewType(nil)
	_ = Type(t1)
	t2 := NewType(t)
	_ = Type(t2)
	t3 := NewType(reflect.TypeOf(nil))
	_ = Type(t3)
	t4 := NewType(NewType(nil))
	_ = Type(t4)
}

func TestTypeNew(t *testing.T) {
	t1 := NewType(struct{}{})
	if !t1.New().IsKind(reflect.Ptr) {
		t.Fatalf("type.New is not ptr")
	}
	i := t1.New().Indirect().InterfaceOrNil()
	if !reflect.DeepEqual(i, struct{}{}) {
		t.Fatalf("type.New is empty struct")
	}
}

type exampleStruct struct {
	test string
}

func TestTypeZero(t *testing.T) {
	tests := []struct {
		x        interface{}
		expected interface{}
	}{
		{struct{}{}, struct{}{}},
		{0, 0},
		{1, 0},
		{uint(1), uint(0)},
		{"test", ""},
		{&exampleStruct{}, (*exampleStruct)(nil)},
		{exampleStruct{test: "1"}, exampleStruct{}},
	}
	for _, tt := range tests {
		t.Run("test can indirect type", func(t *testing.T) {
			n := NewType(tt.x).Zero().InterfaceOrNil()
			if !reflect.DeepEqual(n, tt.expected) {
				t.Errorf("IndirectType() = %v, want %v, got %v", tt.x, tt.expected, n)
			}
		})
	}
}

func TestNewSlice(t *testing.T) {
	t1 := NewType([]uint{})
	v := t1.NewSlice()
	if v.IsNil() {
		t.Fatalf("slice is nil")
	}
	if !reflect.DeepEqual(v.InterfaceOrNil(), []uint{}) {
		t.Fatalf("new slice not created")
	}
}

func TestNewSlice2(t *testing.T) {
	t1 := NewType([]uint{1})
	v := t1.NewSlice()
	if v.IsNil() {
		t.Fatalf("slice is nil")
	}
	if !reflect.DeepEqual(v.InterfaceOrNil(), []uint{}) {
		t.Fatalf("new slice not created")
	}
}
