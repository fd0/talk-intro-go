package pad_test

import (
	"testing"
	"github.com/fd0/pad"
)

var padTests = []struct {
	input string
	len   int
	want  string
}{
	{"foobar", 10, "foobar****"},
	{"foobarbarquux", 8, "foobarbarquux"},
}

func TestRight(t *testing.T) {
	for i, test := range padTests {
		output := pad.Right(test.input, test.len)
		if output != test.want {
			t.Errorf("test %d failed: want %q, got %q", i, test.want, output)
		}
	}
}
