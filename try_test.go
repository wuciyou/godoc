package godoc

import (
	"testing"
)

func TestTry(t *testing.T) {
	Try(func() {
		panic("testTry")
	}, func(e interface{}) {
		t.Log(e)
	})
}
