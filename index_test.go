package refutil

import (
	"testing"
)

func TestSearchIndex(t *testing.T) {
	arr := []uint{1, 2, 3}
	data := NewData(arr)
	index := data.SearchIndex(Equal, uint(2))
	if index != 1 {
		t.Fatal()
	}
	index = data.SearchIndex(Equal, 2)
	if index != 1 {
		t.Fatal()
	}
	index = data.SearchIndex(Equal, nil)
	if index != -1 {
		t.Fatal()
	}
	data = NewData(nil)
	index = data.SearchIndex(Equal, nil)
	if index != -1 {
		t.Fatal()
	}
	defer func() {
		if r := recover(); r != ErrArgumentNotIndexable {
			t.Fatal()
		}
	}()
	data = NewData("")
	data.SearchIndex(Equal, 1)
}

func TestSearchIndex2(t *testing.T) {
	arr := []uint{1, 2, 3}
	if i := Index(arr, 1); i != 0 {
		t.Fatal()
	}
}

func TestSearchIndexSame(t *testing.T) {
	arr := []uint{1, 2, 3}
	if i := IndexSame(arr, 1); i != -1 {
		t.Fatal()
	}
	if i := IndexSame(arr, uint(0)); i != -1 {
		t.Fatal()
	}
}
