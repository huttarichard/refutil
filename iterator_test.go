package refutil

import (
	"testing"
)

func TestIterator(t *testing.T) {
	iterator := NewData([]uint{1, 2, 3}).Iterator()
	if iterator.Index() != 0 {
		t.Fatal()
	}
	if iterator.Len() != 3 {
		t.Fatal()
	}
	if iterator.HasNext() != true {
		t.Fatal()
	}
	if iterator.CanNext() != true {
		t.Fatal()
	}
	keyval := iterator.Current()
	if keyval.Index() != 0 {
		t.Fatal()
	}
	if keyval.Key.String() != "0" {
		t.Fatal()
	}
	if keyval.Value.String() != "1" {
		t.Fatal()
	}
	iterator.bump()
	iterator.bump()
	if iterator.HasNext() != false {
		t.Fatal()
	}
	keyval = iterator.Current()
	if keyval.Index() != 2 {
		t.Fatal()
	}
	if keyval.Key.String() != "2" {
		t.Fatal()
	}
	if keyval.Value.String() != "3" {
		t.Fatal()
	}
	iterator = NewData([]uint{1, 2, 3}).Iterator()
	iterator.Stop()
	if iterator.HasNext() != true {
		t.Fatal()
	}
	if iterator.CanNext() != false {
		t.Fatal()
	}
}

func TestIteration(t *testing.T) {
	var i int
	NewData([]uint{1, 2, 3}).Iterate(func(x *Iterator) {
		kv := x.Current()
		if kv.Index() != i {
			t.Fatal()
		}
		i++
	})
	var arr []int
	NewData([]uint{1, 2, 3}).Iterate(func(x *Iterator) {
		kv := x.Current()
		arr = append(arr, kv.Index())
		if kv.Index() == 1 {
			x.Stop()
		}
	})
	if len(arr) != 2 {
		t.Fatal()
	}
}
