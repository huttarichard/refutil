package refutil

import (
	"testing"
)

func TestSearchValue(t *testing.T) {
	kv := SearchValue([]uint{1, 2, 3}, 3)
	if kv == nil {
		t.Fatal()
	}
	if kv.Index() != 2 {
		t.Fatal()
	}
	kv = SearchValue([]uint{1, 2, 3}, 5)
	if kv != nil {
		t.Fatal()
	}
	kv = SearchSameValue([]uint{1, 2, 3}, 1)
	if kv != nil {
		t.Fatal()
	}
	kv = SearchSameValue([]uint{1, 2, 3}, uint(1))
	if kv.Index() != 0 {
		t.Fatal()
	}
}

func TestSearchKey(t *testing.T) {
	kv := SearchKey([]uint{1, 2, 3}, 2)
	if kv == nil {
		t.Fatal()
	}
	if kv.Index() != 2 {
		t.Fatal()
	}
	kv = SearchKey([]uint{1, 2, 3}, 5)
	if kv != nil {
		t.Fatal()
	}
	kv = SearchSameKey([]uint{1, 2, 3}, 100)
	if kv != nil {
		t.Fatal()
	}
	kv = SearchSameKey([]uint{1, 2, 3}, uint(0))
	if kv != nil {
		t.Fatal()
	}
	kv = SearchSameKey([]uint{1, 2, 3}, 1)
	if kv == nil {
		t.Fatal()
	}
}

func TestSearch(t *testing.T) {
	kv := Search([]uint{1, 2, 3}, func(k *KeyValue) bool {
		return k.Value.String() == "2"
	})
	if kv == nil {
		t.Fatal()
	}
	if kv.Index() != 1 {
		t.Fatal()
	}
	if kv.Value.String() != "2" {
		t.Fatal()
	}
}
