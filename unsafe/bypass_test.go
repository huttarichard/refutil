package unsafe

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/huttarichard/refutil/test"
)

func float64bits(f float64) unsafe.Pointer {
	return (unsafe.Pointer(&f))
}

func TestBypass(t *testing.T) {
	x := test.Sample(test.SampleStruct)
	y := reflect.ValueOf(x).Elem()
	field := ReflectValue(y.FieldByName("interf"))
	if !reflect.DeepEqual(field.Interface(), "test") {
		t.Fatal()
	}
	a := unsafe.Pointer(&struct{ test string }{test: "test"})
	v := reflect.ValueOf(a)
	ReflectValue(v)
}
