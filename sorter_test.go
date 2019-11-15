package refutil

import "testing"

func TestSorter(t *testing.T) {
	v1 := NewValue("2")
	v2 := NewValue("1")
	sorter := ValueSorter([]Value{v1, v2})
	sorter.Sort()
	if sorter[0].String() != "1" {
		t.Fatal()
	}
	if sorter[1].String() != "2" {
		t.Fatal()
	}
}
