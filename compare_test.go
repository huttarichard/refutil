package refutil

import "testing"

func TestComparator(t *testing.T) {
	d := NewData("test")
	if !d.Compare(DeepEqual, NewValue("test")) {
		t.Fatal()
	}
	d = NewData(nil)
	if d.Compare(DeepEqual, NewValue("test")) {
		t.Fatal()
	}
}
