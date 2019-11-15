package refutil

import (
	"testing"
)

type tPkg struct{}

func TestPathTo(t *testing.T) {
	if len(PathTo(&tPkg{})) == 0 {
		t.FailNow()
	}
}
