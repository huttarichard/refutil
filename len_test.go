package refutil

import (
	"testing"

	"github.com/coinfinitygroup/refutil/test"
)

func TestLen(t *testing.T) {
	var header []string
	ok := "1"
	ok2 := map[string]string{"1": "1", "2": "2"}
	ok3 := test.Sample(test.SampleStringSlice)
	var arr [2]string
	tests := []struct {
		x          interface{}
		wantLength int
	}{
		{[]uint{}, 0},
		{[]uint{1}, 1},
		{"", 0},
		{"1", 1},
		{"10", 2},
		{map[string]string{}, 0},
		{map[string]string{"1": "1"}, 1},
		{ok2, 2},
		{&ok2, 2},
		{header, 0},
		{&ok, 1},
		{ok3, 1},
		{&ok3, 1},
		{arr, 2},
	}
	for _, tt := range tests {
		t.Run("test len", func(t *testing.T) {
			l := Len(tt.x)
			if l != tt.wantLength {
				t.Errorf("Len() = %v, want %v", l, tt.wantLength)
			}
		})
	}
}

func TestCanLen(t *testing.T) {
	var header []string
	ok := "1"
	ok2 := map[string]string{"1": "1", "2": "2"}
	ok3 := test.Sample(test.SampleStringSlice)
	ok4 := test.Sample(test.SampleChanBool)
	var arr [2]string
	tests := []struct {
		x      interface{}
		canLen bool
	}{
		{[]uint{}, true},
		{[]uint{1}, true},
		{"", true},
		{"1", true},
		{"10", true},
		{map[string]string{}, true},
		{map[string]string{"1": "1"}, true},
		{ok2, true},
		{&ok2, true},
		{header, true},
		{&ok, true},
		{ok3, true},
		{&ok3, false},
		{&ok4, false}, // in this case its interface rather then slice
		{struct{}{}, false},
		{make(chan bool), true},
		{arr, true},
		{ok4, true},
		{&ok4, false},
	}
	for _, tt := range tests {
		t.Run("test len", func(t *testing.T) {
			l := CanLen(tt.x)
			if l != tt.canLen {
				t.Errorf("CanLen() = %v, want %v for %+#v", l, tt.canLen, tt.x)
			}
		})
	}
}

func TestCanLenForData(t *testing.T) {
	d := NewData(nil)
	if d.CanLen() {
		t.Fatal()
	}
	d2 := NewData("")
	if !d2.CanLen() {
		t.Fatal()
	}
}
