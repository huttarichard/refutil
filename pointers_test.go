package refutil

import (
	"reflect"
	"testing"
)

func TestCanIndirectType(t *testing.T) {
	tests := []struct {
		x    reflect.Type
		want bool
	}{
		{reflect.TypeOf([]uint{}), true},
		{reflect.TypeOf(false), false},
	}
	for _, tt := range tests {
		t.Run("test can indirect type", func(t *testing.T) {
			if got := CanIndirectType(tt.x); got != tt.want {
				t.Errorf("CanIndirectType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndirectType(t *testing.T) {
	tests := []struct {
		x        reflect.Type
		expected string
	}{
		{reflect.TypeOf([]uint{}), "uint"},
		{reflect.TypeOf([]*uint{}), "*uint"},
		{IndirectType(reflect.TypeOf([]*uint{})), "uint"},
		{reflect.TypeOf(false), "bool"},
	}
	for _, tt := range tests {
		t.Run("test can indirect type", func(t *testing.T) {
			if got := IndirectType(tt.x); got.String() != tt.expected {
				t.Errorf("IndirectType() = %v, want %v", got.String(), tt.expected)
			}
		})
	}
}

func TestIndirectTypeOf(t *testing.T) {
	tests := []struct {
		x        interface{}
		expected string
	}{
		{[]uint{}, "uint"},
		{[]*uint{}, "*uint"},
		{false, "bool"},
	}
	for _, tt := range tests {
		t.Run("test can indirect type of", func(t *testing.T) {
			if got := IndirectTypeOf(tt.x); got.String() != tt.expected {
				t.Errorf("IndirectTypeOf() = %v, want %v", got.String(), tt.expected)
			}
		})
	}
}

func TestCanIndirectValue(t *testing.T) {
	if CanIndirectValue(reflect.ValueOf([]uint{})) {
		t.Fail()
	}
	if !CanIndirectValue(reflect.ValueOf(&struct{}{})) {
		t.Fail()
	}
}

func TestIndirectValue(t *testing.T) {
	a := 1
	if IndirectValue(reflect.ValueOf(&a)).Type().String() != "int" {
		t.Fatalf("invalid indirect value of a")
	}
	var x *string
	if !NewValue(x).IsValid() {
		t.Fatalf("invalid indirect value of x pointer")
	}
	if IndirectValueOf(x).IsValid() {
		t.Fatalf("indirect value of x should not be valid")
	}
}

func TestIndirect(t *testing.T) {
	a := 1
	b := struct{}{}
	c := &b
	if Indirect(&a) != 1 {
		t.Fatalf("indirect of 1 is not 1")
	}
	if Indirect(&b) != b {
		t.Fatalf("indirect of &b is not b")
	}
	if Indirect(Indirect(&c)) != b {
		t.Fatalf("double indirect of &c is not b")
	}
}

func TestCanIndirect(t *testing.T) {
	t1 := NewData(nil)
	if t1.Type().CanIndirect() {
		t.Fatalf("nil is not indirectable")
	}
	if t1.CanIndirect() {
		t.Fatalf("nil is not indirectable")
	}
	if CanIndirect(nil) {
		t.Fatalf("nil is not indirectable")
	}
	t2 := NewData(&struct{}{})
	if !t2.Type().CanIndirect() {
		t.Fatalf("&struct{}{} is indirectable")
	}
	if !t2.CanIndirect() {
		t.Fatalf("&struct{}{} is indirectable")
	}
	if !CanIndirect(&struct{}{}) {
		t.Fatalf("&struct{}{} is indirectable")
	}
}
