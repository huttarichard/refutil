package refutil

import (
	"testing"
)

func TestStructIndexer(t *testing.T) {
	str := struct {
		Test1 string
		Test2 int
	}{
		Test1: "test",
		Test2: 1,
	}
	data := NewValue(str)
	indexer := NewStructIndexer(data)
	if len(indexer.Keys()) != 2 {
		t.Fatal()
	}
	k1 := indexer.Keys()[0].String()
	k2 := indexer.Keys()[1].String()
	if k1 != "Test1" || k2 != "Test2" {
		t.Fatal()
	}
	key1, v1 := data.At(0)
	if v1.String() != "test" || key1.String() != "Test1" {
		t.Fatal()
	}
	key2, v2 := data.At(1)
	if v2.Interface().(int) != 1 || key2.String() != "Test2" {
		t.Fatal()
	}
}

func TestMapIndexer(t *testing.T) {
	str := map[string]interface{}{
		"Test1": "test",
		"Test2": 1,
	}
	data := NewValue(str)
	indexer := NewMapIndexer(data)
	if len(indexer.Keys()) != 2 {
		t.Fatal()
	}
	k1 := indexer.Keys()[0].String()
	k2 := indexer.Keys()[1].String()
	if k1 != "Test1" || k2 != "Test2" {
		t.Fatal()
	}
	key1, v1 := data.At(0)
	if v1.String() != "test" || key1.String() != "Test1" {
		t.Fatal()
	}
	key2, v2 := data.At(1)
	if v2.Interface().(int) != 1 || key2.String() != "Test2" {
		t.Fatal()
	}
}

func TestSliceIndexer(t *testing.T) {
	str := []string{"1", "2"}
	data := NewValue(str)
	indexer := NewSliceIndexer(data)
	if len(indexer.Keys()) != 2 {
		t.Fatal()
	}
	k1 := indexer.Keys()[0].String()
	k2 := indexer.Keys()[1].String()
	if k1 != "0" || k2 != "1" {
		t.Fatal()
	}
	key1, v1 := data.At(0)
	if v1.String() != "1" || key1.String() != "0" {
		t.Fatal()
	}
	key2, v2 := data.At(1)
	if v2.String() != "2" || key2.String() != "1" {
		t.Fatal()
	}
}

func TestCanIndex(t *testing.T) {
	if !NewValue([]uint{1, 2, 3}).CanIndex() {
		t.Fatal()
	}
	if !NewValue("123").CanIndex() {
		t.Fatal()
	}
	if NewValue(struct{}{}).CanIndex() {
		t.Fatal()
	}
}

func TestIndex(t *testing.T) {
	val := NewValue([]uint{1, 2, 3})
	if val.Index(1).InterfaceOrNil().(uint) != uint(2) {
		t.Fatal()
	}
}

func TestCanUseIndexer(t *testing.T) {
	if !NewValue([]uint{}).CanUseIndexer() {
		t.Fatal()
	}
	if NewValue("").CanUseIndexer() {
		t.Fatal()
	}
}

func TestIndexer(t *testing.T) {
	val := NewValue([]uint{1, 2, 3})
	_ = val.Indexer().(*SliceIndexer)
	val = NewValue(map[string]string{})
	_ = val.Indexer().(*MapIndexer)
	val = NewValue(struct{}{})
	_ = val.Indexer().(*StructIndexer)
	defer func() {
		if r := recover(); r != ErrArgumentNotIndexable {
			t.Fatal()
		}
	}()
	NewValue("").Indexer()
}

func TestIndexerAt(t *testing.T) {
	val := NewValue([]uint{1, 2, 3})
	k, v := val.At(0)
	if k.String() != "0" || v.String() != "1" {
		t.Fatal()
	}
}

func TestDataIndexerAt(t *testing.T) {
	val := NewData([]uint{1, 2, 3})
	k, v := val.At(0)
	if k.String() != "0" || v.String() != "1" {
		t.Fatal()
	}
}
