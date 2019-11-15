package refutil

import "testing"

type testStinger struct{}

func (s *testStinger) String() string {
	return "stringer"
}

type testStinger2 struct{}

func (s testStinger2) String() string {
	return "stringer"
}

func TestString(t *testing.T) {
	x := "test"
	x2 := &testStinger{}
	tests := []struct {
		x        interface{}
		expected string
	}{
		{"1", "1"},
		{&x, "test"},
		{&x2, "stringer"},
		{&testStinger{}, "stringer"},
		{testStinger2{}, "stringer"},
		{map[string]string{}, "map[]"},
		{nil, ""},
		{(*string)(nil), ""},
	}
	for _, tt := range tests {
		t.Run("test can indirect type", func(t *testing.T) {
			if got := String(tt.x); got != tt.expected {
				t.Errorf("String() = %v, want %v", got, tt.expected)
			}
		})
	}
}
