package refutil

import (
	"fmt"
	"testing"
)

type YourStruct struct {
	A string
}

func (s *YourStruct) Equal(v interface{}) bool {
	return s.A == fmt.Sprintf("%s", v)
}

type DiffStruct struct {
	A string
}

func (s DiffStruct) String() string {
	return s.A
}

func TestDeepEqual(t *testing.T) {
	var s *string
	var x *string
	var y *string
	var z *uint
	tests := []struct {
		expected interface{}
		equals   interface{}
		want     bool
	}{
		{uint(0), 0, false},
		{s, s, true},
		{s, nil, false},
		{emptyError(), emptyError(), true},
		{emptyError(), nil, false},
		{nil, emptyError(), false},
		{struct{}{}, struct{}{}, true},
		{struct{}{}, nil, false},
		{"Hello World", "Hello World", true},
		{123, 123, true},
		{123.5, 123.5, true},
		{[]byte("Hello World"), []byte("Hello World"), true},
		{nil, nil, true},
		{map[int]int{5: 10}, map[int]int{10: 20}, false},
		{map[int]int{10: 20}, map[int]int{10: 20}, true},
		{&map[int]int{10: 20}, map[int]int{10: 20}, false},
		{'x', "x", false},
		{"x", 'x', false},
		{0, 0.1, false},
		{0.1, 0, false},
		{uint32(10), int32(10), false},
		{x, y, true},
		{x, z, false},
	}
	for _, tt := range tests {
		t.Run("test deepequal", func(t *testing.T) {
			l := DeepEqual(tt.expected, tt.equals)
			if l != tt.want {
				t.Errorf("DeepQual() = %v, want %v", l, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	var s *string
	var x *string
	var y *string
	var z *uint
	var double *string
	tests := []struct {
		expected interface{}
		equals   interface{}
		want     bool
	}{
		{struct{}{}, nil, false},
		{s, nil, true},
		{s, s, true},
		{emptyError(), emptyError(), true},
		{emptyError(), nil, true},
		{nil, emptyError(), true},
		{struct{}{}, struct{}{}, true},
		{struct{}{}, struct{ A string }{}, false},
		{struct{}{}, struct{ A string }{A: "ok"}, false},
		{"Hello World", "Hello World", true},
		{map[int]int{10: 20}, map[int]int{10: 20}, true},
		{&map[int]int{10: 20}, map[int]int{10: 20}, true},
		{map[int]int{5: 10}, map[int]int{10: 20}, false},
		{&map[int]int{5: 10}, map[int]int{10: 20}, false},
		{uint32(10), int32(10), true},
		{0, nil, false},
		{nil, 0, false},
		{y, z, true},
		{nil, z, true},
		{x, nil, true},
		{(*string)(nil), nil, true},
		{&double, double, true},
		{&YourStruct{A: "same"}, &DiffStruct{A: "same"}, true},
		{&YourStruct{A: "same"}, &DiffStruct{A: "diff"}, false},
	}
	for _, tt := range tests {
		t.Run("test equal", func(t *testing.T) {
			l := Equal(tt.expected, tt.equals)
			if l != tt.want {
				t.Errorf("Equal() = %v, want %v", l, tt.want)
			}
		})
	}
}

func ExampleEqual() {
	fmt.Printf("%v\n", Equal(uint(0), 0))
	// implements `Equal(v interface{}) bool`
	// checkout type `YourStruct`
	a := &YourStruct{A: "same"}
	b := &DiffStruct{A: "same"}
	fmt.Printf("%v\n", Equal(a, b))
	// Output:
	// true
	// true
}

func ExampleDeepEqual() {
	fmt.Printf("%v\n", DeepEqual(uint(0), 0))
	fmt.Printf("%v\n", DeepEqual(0, 0))
	// Output:
	// false
	// true
}
