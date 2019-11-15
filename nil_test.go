package refutil

import (
	"testing"
	"time"

	"github.com/coinfinitygroup/refutil/test"
)

func TestNil(t *testing.T) {
	var header []string
	var header2 = []string{"ok"}
	nilString := test.Sample(test.SampleStringNil)
	str := test.Sample(test.SampleStringPtr)
	tests := []struct {
		x    interface{}
		want bool
	}{
		{&header, true},
		{nil, true},
		{(*time.Time)(nil), true},
		{header, true},
		{header2, false},
		{(*time.Time)(&time.Time{}), false},
		{1, false},
		{0, false},
		{nilString, true},
		{str, false},
	}
	for _, tt := range tests {
		t.Run("test nil", func(t *testing.T) {
			if got := IsNil(tt.x); got != tt.want {
				t.Errorf("Nil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUntyped(t *testing.T) {
	v := NewValue(nil)
	if !v.Untyped().IsNil() {
		t.Fatalf("untyped is not nil")
	}
	v2 := NewValue(struct{}{})
	if v2.Untyped() != v2 {
		t.Fatal()
	}
}
