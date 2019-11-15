package test

import "testing"

func TestUnexported(t *testing.T) {
	Sample(SampleInt)
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal()
		}
	}()
	Sample(SampleType(9999))
}

func TestZeroer(t *testing.T) {
	z := &sampleZeroer{positive: true}
	if !z.IsZero() {
		t.Fatal()
	}
	x := Sample(SampleZeroer)
	if x.(zeroer).IsZero() {
		t.Fatal()
	}
}
