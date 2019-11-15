package refutil

import (
	"testing"
)

func TestCanInterface(t *testing.T) {
	d := NewData(nil)
	if d.CanInterface() {
		t.Fatal()
	}
	d2 := NewData("")
	if !d2.CanInterface() {
		t.Fatal()
	}
	d3 := NewData(struct{}{})
	if !d3.CanInterface() {
		t.Fatal()
	}
}
