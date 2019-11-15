package refutil

import "testing"

func TestContains(t *testing.T) {
	if !Contains([]uint{1, 2, 3}, 1) {
		t.Fatal()
	}
	if ContainsSameValue([]uint{1, 2, 3}, 1) {
		t.Fatal()
	}
	if !ContainsSameValue([]uint{1, 2, 3}, uint(1)) {
		t.Fatal()
	}
	if !ContainsSameKey([]uint{1, 2, 3}, 1) {
		t.Fatal()
	}
	if ContainsSameKey([]uint{1, 2, 3}, 20) {
		t.Fatal()
	}
	if !ContainsKey([]uint{1, 2, 3}, uint(1)) {
		t.Fatal()
	}
}
