package path

import (
	"path"
	"testing"
)

func TestMatch(t *testing.T) {
	t.Log(path.Match("adc*", "adcdadf"))
}
