package refutil

import (
	"testing"
	"time"

	"github.com/huttarichard/refutil/test"
)

type zeroerTest struct {
	Positive bool
}

func (p zeroerTest) IsZero() bool {
	return p.Positive
}

type zeroerTest2 struct {
	Positive bool
}

func (p *zeroerTest2) IsZero() bool {
	return p.Positive
}

func TestIsZero(t *testing.T) {
	type testZero struct {
		Value  interface{}
		Result bool
	}
	type testZeroStr struct {
		Zero bool
	}
	var emptyArr []string
	var emptyArr2 []*string
	type customArr []string
	type customArr2 []*string
	var str = "test"
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}
	var cases = []testZero{
		{customArr{}, true},
		{customArr2{}, true},
		{"", true},
		{[]string{}, true},
		{[]int{}, true},
		{[]uint{}, true},
		{[]float64{}, true},
		{struct{}{}, true},
		{0, true},
		{uint(0), true},
		{nil, true},
		{new(time.Time), true},
		{(*testZeroStr)(nil), true},
		{(*testZeroStr)(&testZeroStr{}), true},
		{map[string]string{}, true},
		{make(chan struct{}), true},
		{make(chan struct{}, 1), true},
		{time.Time{}, true},
		{zeroerTest{true}, true},
		{&zeroerTest{true}, true},
		{emptyArr, true},
		{emptyArr2, true},
		{[]*string{}, true},
		{&zeroerTest{false}, false},
		{zeroerTest{false}, false},
		{chWithValue, false},
		{map[string]string{"k": "k"}, false},
		{(*testZeroStr)(&testZeroStr{Zero: true}), false},
		{new(time.Time), true},
		{"1", false},
		{[]string{"1"}, false},
		{[]*string{&str}, false},
		{[]int{1}, false},
		{[]uint{1}, false},
		{[]float64{1}, false},
		{struct{ Test string }{Test: "1"}, false},
		{1, false},
		{uint(1), false},
		{NewValue(nil), true},
		{NewValue(struct{}{}), true},
		{NewValue(""), true},
	}
	for _, c := range cases {
		if c.Result {
			if !IsZero(c.Value) {
				t.Fatalf("'%v' is not zero", c.Value)
			}
		} else {
			if IsZero(c.Value) {
				t.Fatalf("'%+#v' is zero", c.Value)
			}
		}
	}
}

func TestZeroer(t *testing.T) {
	z := test.Sample(test.SampleEmptyZeroer)
	v := NewValue(z)
	_, ok := v.Zeroer()
	if ok {
		t.Fatal()
	}
}
